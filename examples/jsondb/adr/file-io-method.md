# 파일 I/O 방식

## Concern
jsondb에서 파일을 읽고 쓸 때 어떤 I/O 방식을 사용해야 하는가?

## Status
approved

## Context
- 파일 기반 데이터 모델 채택
- 필드 레벨 접근이 핵심 기능 (전체 파일 로드 없이)
- 여러 프로세스 동시 접근 가능
- Go 언어로 구현

## Decision
**단계적 접근 + Storage Backend 추상화** 사용자 승인: Phase 1은 직접 I/O로 시작하되, 향후 mmap/WAL로 변경이 용이하도록 Storage Backend 인터페이스를 추상화.

## Options

### Option A — 직접 파일 I/O (Direct I/O)

**방식**:
```go
// 파일 열기 → 읽기 → 닫기
data, err := os.ReadFile("overview.json")
json.Unmarshal(data, &obj)

// 수정 후 쓰기
data, _ = json.Marshal(obj)
os.WriteFile("overview.json", data, 0644)
```

**장점**:
- **단순성**: 표준 라이브러리로 구현 가능
- **예측성**: 명확한 읽기/쓰기 동작
- **디버깅 용이**: 실제 파일 상태와 일치

**단점**:
- **성능**: 매번 파일 열기/닫기 오버헤드
- **전체 로드**: JSON 파싱을 위해 전체 파일 읽기 필요
- **동시성**: 파일 레벨 잠금만 가능
- **쓰기 효율**: 변경 시 전체 파일 덮어쓰기

**성능 특성**:
```
읽기: ~1-10ms (파일 크기에 따라)
쓰기: ~1-10ms (전체 파일 크기)
동시성: flock() 으로 제어
```

---

### Option B — 메모리 매핑 (mmap)

**방식**:
```go
// 파일을 메모리에 매핑
f, _ := os.Open("overview.json")
data, _ := syscall.Mmap(int(f.Fd()), 0, stat.Size(), 
    syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)

// 메모리에서 직접 접근
// 수정 시 자동으로 파일에 반영
```

**장점**:
- **성능**: 커널이 페이지 캐싱으로 최적화
- **랜덤 접근**: 파일의 임의 위치를 빠르게 접근
- **쓰기 효율**: 수정된 페이지만 디스크에 기록
- **동시성**: OS 수준에서 처리

**단점**:
- **복잡성**: mmap API 사용법 필요
- **페이지 폴트**: 초기 접근 시 지연 발생
- **플랫폼 의존**: OS별 mmap 구현 차이
- **파일 크기 변경**: 매핑 후 파일 확장 어려움

**성능 특성**:
```
읽기: ~0.1-1ms (페이지 캐시 히트 시)
쓰기: ~0.1-1ms (변경된 페이지만 기록)
동시성: OS 수준 페이지 잠금
```

---

### Option C — 버퍼링된 I/O (Buffered I/O)

**방식**:
```go
// 버퍼 기반 읽기
reader := bufio.NewReader(file)
reader.Read(buf)

// 버퍼 기반 쓰기
writer := bufio.NewWriter(file)
writer.Write(data)
writer.Flush()
```

**장점**:
- **시스템 콜 최소화**: 여러 읽기/쓰기를 버퍼에서 처리
- **유연성**: 버퍼 크기 조정 가능
- **표준 라이브러리**: Go에 내장

**단점**:
- **버퍼 관리**: 버퍼 플러시 타이밍 관리 필요
- **캐싱 불일치**: 버퍼와 실제 파일 상태 차이 가능
- **동시성**: 기본적으로 제한적

**성능 특성**:
```
읽기: ~0.5-5ms (버퍼 히트 시 빠름)
쓰기: ~0.5-5ms (플러시 시점에 기록)
동시성: 추가 잠금 메커니즘 필요
```

---

### Option D — mmap + 인덱스 조합

**방식**:
```go
// 1. 파일을 mmap으로 매핑
// 2. 필드별 오프셋 인덱스 유지
// 3. 인덱스로 빠른 접근

type FieldIndex struct {
    FieldName string
    Offset    int64  // 파일 내 위치
    Length    int    // 필드 데이터 길이
}

// 인덱스 기반 접근
func (db *JsonDB) Get(field string) interface{} {
    idx := db.index[field]
    data := db.mmap[idx.Offset : idx.Offset+int64(idx.Length)]
    return json.Unmarshal(data)
}
```

**장점**:
- **최고 성능**: mmap + 인덱스로 빠른 접근
- **필드 레벨**: 전체 파싱 없이 필드만 접근
- **읽기 쓰기 분리**: 읽기는 인덱스, 쓰기는 mmap

**단점**:
- **가장 복잡**: 인덱스 관리 필요
- **동기화**: 인덱스와 데이터 동기화 필수
- **구현 부담**: 인덱스 갱신 로직 구현

**성능 특성**:
```
필드 읽기: ~0.01-0.1ms (인덱스 조회 후 접근)
필드 쓰기: ~0.1-1ms (데이터 + 인덱스 갱신)
동시성: 인덱스와 데이터에 대한 잠금 필요
```

---

### Option E — WAL (Write-Ahead Logging)

**방식**:
```go
// 1. 변경사항을 로그 파일에 먼저 기록
// 2. 로그 기반으로 메인 파일 갱신
// 3. 로그 파일 주기적 정리

func (db *JsonDB) Set(key string, value interface{}) {
    // 로그에 기록
    db.wal.Write(LogEntry{Key: key, Value: value})
    // 메인 파일 갱신
    db.updateMainFile(key, value)
    // 로그에서 제거
    db.wal.Remove(LogEntry)
}
```

**장점**:
- **복구 가능**: 비정상 종료 후 로그로 복구
- **동시성**: 읽기는 메인 파일, 쓰기는 로그로 분리
- **성능**: 쓰기는 로그에 빠르게 기록

**단점**:
- **복잡성**: 로그 관리, 컴팩션 필요
- **공간**: 로그 파일 추가 공간 사용
- **복구 시간**: 로그가 길면 복구 시간 증가

**성능 특성**:
```
읽기: ~0.1-1ms (메인 파일 접근)
쓰기: ~0.1-0.5ms (로그 기록)
동시성: 읽기/쓰기 분리로 높은 병렬성
```

## Tradeoffs

| 기준 | 직접 I/O | mmap | 버퍼링 | mmap+인덱스 | WAL |
|------|---------|------|--------|------------|-----|
| **구현 복잡도** | ★★★★★ | ★★★☆☆ | ★★★★☆ | ★★☆☆☆ | ★☆☆☆☆ |
| **읽기 성능** | ★★☆☆☆ | ★★★★★ | ★★★☆☆ | ★★★★★ | ★★★★★ |
| **쓰기 성능** | ★★☆☆☆ | ★★★★☆ | ★★★☆☆ | ★★★★☆ | ★★★★★ |
| **동시성** | ★★☆☆☆ | ★★★★☆ | ★★☆☆☆ | ★★★☆☆ | ★★★★★ |
| **복구 가능성** | ★☆☆☆☆ | ★★☆☆☆ | ★★☆☆☆ | ★★☆☆☆ | ★★★★★ |
| **메모리 효율** | ★★★★☆ | ★★★☆☆ | ★★★★☆ | ★★★☆☆ | ★★★☆☆ |

## Recommendation
**단계적 접근** 추천:

**Phase 1: 직접 I/O (Option A)**로 빠른 프로토타이핑 시작
- 가장 단순하고 빠른 구현
- 필드 레벨 접근은 JSON 경로 파싱으로 구현
- 동시성은 파일 잠금으로 처리

**Phase 2: mmap (Option B)**으로 성능 최적화
- 필드 접근이 빈번해지면 mmap으로 전환
- OS 캐싱 활용으로 성능 향상

**Phase 3: mmap+인덱스 (Option D)**로 고성능 달성
- 필드 접근이 매우 빈번하면 인덱스 추가
- 최고 수준의 성능

이유:
1. **MVP 우선**: 빠른 구현과 테스트 가능
2. **점진적 최적화**: 병목이 확인된 부분만 최적화
3. **유연성**: 사용 패턴에 따라 조정 가능

## Downstream Concerns
- [ ] **동시성 제어**: 파일 잠금 또는 mmap 잠금 전략
- [ ] **인덱싱 방식**: 필드별 인덱스 구현 방법
- [ ] **캐싱 전략**: 메모리 캐시 정책 (LRU, TTL 등)
- [ ] **에러 처리**: 파일 없음, 접근 거부 등 처리

## Storage Backend 추상화 설계

**목적**: 향후 mmap/WAL 등으로 변경이 용이하도록 인터페이스 추상화

```go
// StorageBackend 인터페이스
type StorageBackend interface {
    // 파일 열기
    Open(path string) error
    
    // 필드 읽기
    Get(fieldPath string) (interface{}, error)
    
    // 필드 쓰기
    Set(fieldPath string, value interface{}) error
    
    // 파일 저장
    Save() error
    
    // 파일 닫기
    Close() error
}

// 직접 I/O 구현체
type DirectIOBackend struct {
    data map[string]interface{}
    path string
}

// 향후 구현할 구현체들
// type MmapBackend struct { ... }
// type WALBackend struct { ... }
```

**장점**:
- **유연성**: 백엔드 교체 시 엔진 로직 변경 불필요
- **테스트 용이성**: Mock 백엔드로 테스트 가능
- **점진적 도입**: 필요한 시점에 고성능 백엔드 추가

## Tags
`file-io`, `performance`, `mmap`, `storage`

## Approved
- 2026-07-21: 단계적 접근 + Storage Backend 추상화, 사용자 승인