# 동시성 제어 방식

## Concern
여러 에이전트(프로세스)가 동시에 jsondb에 접근할 때 데이터 무결성과 일관성을 어떻게 보장해야 하는가?

## Status
approved

## Context
- 파일 기반 데이터 모델
- 하이브리드 인터페이스 (CLI + Server)
- 여러 에이전트가 동시에 접근 가능
- 필드 레벨 접근이 핵심 기능

## Decision
**파일 잠금 + 메모리 잠금 혼합 (Option C)** 사용자 승인: 프로세스 간 동기화는 파일 잠금, 서버 내부 스레드 간 동기화는 RWLock으로 처리.

## 동작 시나리오

### 시나리오 1: 단일 클라이언트
```
에이전트 A → CLI/Server → Engine → 파일
```
- 동시성 문제 없음
- 잠금 불필요

### 시나리오 2: 읽기 경합
```
에이전트 A → 읽기 → 파일 ← 읽기 ← 에이전트 B
```
- 읽기는 병렬 가능
- 데이터 변경 없음 → 안전

### 시나리오 3: 쓰기 경합
```
에이전트 A → 쓰기 → 파일 ← 쓰기 ← 에이전트 B
```
- 데이터 손상 위험
- 잠금 필요

### 시나리오 4: 읽기-쓰기 경합
```
에이전트 A → 읽기 → 파일 ← 쓰기 ← 에이전트 B
```
- 불일치 위험
- 잠금 필요

## Options

### Option A — 파일 잠금 (File Locking)

**방식**:
```go
import "syscall"

func (b *DirectIOBackend) Lock() error {
    return syscall.Flock(int(b.file.Fd()), syscall.LOCK_EX)
}

func (b *DirectIOBackend) Unlock() error {
    return syscall.Flock(int(b.file.Fd()), syscall.LOCK_UN)
}
```

**동작 원리**:
- 쓰기 시 exclusively lock (LOCK_EX)
- 읽기 시 shared lock (LOCK_SH)
- 잠금 해제 시 자동 해제

**장점**:
- **단순성**: OS 수준 잠금, 구현 용이
- **안정성**: 프로세스 비정상 종료 시 자동 해제
- **성능**: 읽기 병렬화 가능

**단점**:
- **파일 수준**: 전체 파일 잠금, 필드 수준 잠금 불가
- **잠금 경쟁**: 빈번한 쓰기 시 병목
- **데드락 위험**: 잠금 순서 불일치 시 발생 가능

**성능 특성**:
```
읽기: 병렬 가능 (shared lock)
쓰기: 직렬화 (exclusive lock)
잠금 오버헤드: ~0.1ms
```

---

### Option B — 읽기-쓰기 잠금 (RWLock)

**방식**:
```go
type RWLock struct {
    mu       sync.RWMutex
    filePath string
}

func (l *RWLock) RLock() {
    l.mu.RLock()
}

func (l *RWLock) RUnlock() {
    l.mu.RUnlock()
}

func (l *RWLock) Lock() {
    l.mu.Lock()
}

func (l *RWLock) Unlock() {
    l.mu.Unlock()
}
```

**동작 원리**:
- 읽기: RLock (여러 스레드 가능)
- 쓰기: Lock (단일 스레드)
- 메모리 수준 잠금

**장점**:
- **높은 성능**: 파일 I/O 없이 잠금
- **유연성**: 세밀한 잠금 제어 가능
- **단순성**: Go sync 패키지 활용

**단점**:
- **프로세스 간 불가**: 단일 프로세스 내부에서만 동작
- **서버 모드 전용**: CLI 모드에서는 의미 없음
- **영속성 없음**: 프로세스 종료 시 잠금 해제

**성능 특성**:
```
읽기: 병렬 가능
쓰기: 직렬화
잠금 오버헤드: ~0.01ms (메모리 수준)
```

---

### Option C — 파일 잠금 + 메모리 잠금 혼합

**방식**:
```
Server 모드: 파일 잠금 (프로세스 간) + RWLock (내부 스레드)
CLI 모드: 파일 잠금만
```

**동작 원리**:
- 서버: 파일 잠금으로 프로세스 간 동기화
- 서버 내부: RWLock으로 스레드 간 동기화
- CLI: 파일 잠금만 사용

**장점**:
- **완전한 동기화**: 프로세스 간 + 스레드 간 모두 처리
- **유연성**: 모드에 따라 적절한 잠금
- **성능**: 최적의 잠금 전략 적용

**단점**:
- **복잡성**: 두 가지 잠금 관리 필요
- **구현 부담**: 혼합 로직 구현

**성능 특성**:
```
읽기: 병렬 가능 (파일 + 메모리)
쓰기: 직렬화 (파일 + 메모리)
잠금 오버헤드: ~0.1-0.2ms
```

---

### Option D — 버전 기반 잠금 (Optimistic Locking)

**방식**:
```go
type VersionedData struct {
    data    map[string]interface{}
    version int64
}

func (d *VersionedData) Update(path string, value interface{}, expectedVersion int64) error {
    if d.version != expectedVersion {
        return ErrVersionConflict
    }
    // 업데이트
    d.version++
    return nil
}
```

**동작 원리**:
- 읽기 시 version 번호 확인
- 쓰기 시 version 변경 여부 확인
- 변경 시 재시도

**장점**:
- **높은 병렬성**: 잠금 없이 읽기 가능
- **단순성**: 잠금 메커니즘 불필요
- **성능**: 잠금 오버헤드 없음

**단점**:
- **충돌 시 재시도**: 빈번한 쓰기 시 성능 저하
- **구현 복잡**: version 관리 필요
- **데이터 일시적 불일치**: 충돌 감지까지 불일치 가능

**성능 특성**:
```
읽기: 완전 병렬
쓰기: 충돌 시 재시도
잠금 오버헤드: 0ms
```

## Tradeoffs

| 기준 | 파일 잠금 | RWLock | 혼합 | 버전 기반 |
|------|----------|--------|------|----------|
| **구현 복잡도** | ★★★★★ | ★★★★☆ | ★★★☆☆ | ★★★☆☆ |
| **프로세스 간 동기화** | ★★★★★ | ★☆☆☆☆ | ★★★★★ | ★★☆☆☆ |
| **성능** | ★★★☆☆ | ★★★★★ | ★★★★☆ | ★★★★★ |
| **안정성** | ★★★★★ | ★★★☆☆ | ★★★★★ | ★★★☆☆ |
| **단순성** | ★★★★★ | ★★★★☆ | ★★★☆☆ | ★★★☆☆ |

## Recommendation
**파일 잠금 + 메모리 잠금 혼합 (Option C)** 추천 이유:

1. **완전한 동기화**: 프로세스 간 + 스레드 간 모두 처리
2. **안정성**: 파일 잠금으로 프로세스 비정상 종료 대응
3. **성능**: RWLock으로 내부 스레드 동기화 최적화
4. **유연성**: CLI/Server 모드에 따라 적절한 잠금

**구현 전략**:
1. Phase 1: 파일 잠금만 구현 (기본 기능)
2. Phase 2: RWLock 추가 (성능 최적화)

## Downstream Concerns
- [ ] **잠금 타임아웃**: 잠금 대기 시간 제한
- [ ] **데드락 방지**: 잠금 순서 관리
- [ ] **잠금 해제 전략**: 비정상 종료 시 잠금 해제
- [ ] **모니터링**: 잠금 상태 모니터링

## Tags
`concurrency`, `locking`, `file-lock`, `rwlock`

## Approved
- 2026-07-21: Option C (파일 잠금 + 메모리 잠금 혼합), 사용자 승인