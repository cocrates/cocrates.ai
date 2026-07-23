# 서버 동시 쓰기 직렬화

## Concern
Long-lived Engine 구조에서 같은 파일에 대한 동시 쓰기 요청을 어떻게 직렬화할 것인가?

## Status
approved

## Context
- `server-engine-lifecycle.md` ADR에서 파일별 Long-lived Engine 채택
- 서버가 시작되면 파일별 Engine 인스턴스를 lazy하게 생성하여 유지
- 여러 HTTP 요청이 동시에 같은 파일에 대한 쓰기 작업 수행 가능
- 기존 Engine은 `sync.RWMutex`를 보유하여 단일 Engine 인스턴스 내부에서는 직렬화 보장
- **미해결 질문**: 파일별 독립 Engine 구조에서 동시 쓰기 요청의 안전한 처리 방식

## Decision
**Option B — 파일 경로 기반 RWLock** 사용자 승인: 서버 레벨에서 파일 경로를 키로 하는 RWLock 맵을 유지하여 파일 수준에서 명시적 직렬화 보장.

## Options

### Option A — Engine 내부 Mutex에 의존

**방식**: 기존 Engine의 `sync.RWMutex`에 전적으로 의존. 별도의 추가 잠금 없이 Engine이 직렬화를 처리.

```go
// 현재 구조 유지
func (s *Server) handleDataSet(w http.ResponseWriter, r *http.Request, filePath string) {
    me, err := s.getOrCreateEngine(filePath)
    if err != nil {
        s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
        return
    }
    
    // Engine.Set()이 내부 mutex로 직렬화
    if err := me.engine.Set(field, value); err != nil {
        s.writeError(w, http.StatusBadRequest, "SET_FAILED", err.Error())
        return
    }
    
    // Engine.Save()가 내부 mutex로 직렬화
    if err := me.engine.Save(); err != nil {
        s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", err.Error())
        return
    }
}
```

**장점**:
- 추가 잠금 로직 불필요
- 기존 Engine 구조 최대한 재사용
- 구현 복잡도 낮음

**단점**:
- Set + Save 사이에 다른 요청이 끼어들 수 있음 (graceful)
- 파일 수준 직렬화가 아닌 메모리 수준 직렬화

---

### Option B — 파일 경로 기반 RWLock

**방식**: 서버 레벨에서 파일 경로를 키로 하는 RWLock 맵을 유지. 쓰기 요청 시 Lock, 읽기 요청 시 RLock.

```go
type Server struct {
    engines    map[string]*managedEngine
    fileLocks  map[string]*sync.RWMutex
    mu         sync.RWMutex  // map 접근용
    addr       string
    root       string
}

func (s *Server) getFileLock(filePath string) *sync.RWMutex {
    s.mu.RLock()
    lock, ok := s.fileLocks[filePath]
    s.mu.RUnlock()
    if ok {
        return lock
    }
    
    s.mu.Lock()
    defer s.mu.Unlock()
    if lock, ok = s.fileLocks[filePath]; ok {
        return lock
    }
    lock = &sync.RWMutex{}
    s.fileLocks[filePath] = lock
    return lock
}

func (s *Server) handleDataGet(w http.ResponseWriter, r *http.Request, filePath string) {
    lock := s.getFileLock(filePath)
    lock.RLock()
    defer lock.RUnlock()
    
    me, err := s.getOrCreateEngine(filePath)
    if err != nil {
        s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
        return
    }
    
    result, err := me.engine.Get("")
    // ...
}

func (s *Server) handleDataSet(w http.ResponseWriter, r *http.Request, filePath string) {
    lock := s.getFileLock(filePath)
    lock.Lock()
    defer lock.Unlock()
    
    me, err := s.getOrCreateEngine(filePath)
    if err != nil {
        s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
        return
    }
    
    if err := me.engine.Set(field, value); err != nil {
        s.writeError(w, http.StatusBadRequest, "SET_FAILED", err.Error())
        return
    }
    
    if err := me.engine.Save(); err != nil {
        s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", err.Error())
        return
    }
}
```

**장점**:
- 파일 수준에서 명시적 직렬화 보장
- 읽기 요청은 병렬 허용 (RLock)
- 쓰기 요청은 완전 직렬화 (Lock)
- Race Condition 완전 차단

**단점**:
- 추가 잠금 구조 필요
- 데드락 위험 (잠금 순서 불일치)
- 메모리 사용량 증가 (파일당 하나의 RWMutex)

---

### Option C — Set + Save 원자적 처리

**방식**: Set과 Save를 하나의 원자적 작업으로 묶어 다른 요청이 끼어들지 못하도록 함. Engine에 새로운 메서드 추가.

```go
// Engine 인터페이스 확장
type Engine interface {
    // ... 기존 메서드 ...
    
    // SetAndSave는 Set과 Save를 원자적으로 수행
    SetAndSave(fieldPath string, value interface{}) error
}

func (e *jsonEngine) SetAndSave(fieldPath string, value interface{}) error {
    e.mu.Lock()
    defer e.mu.Unlock()
    
    if !e.open {
        return ErrNotOpen
    }
    
    parsed, err := query.Parse(fieldPath)
    if err != nil {
        return fmt.Errorf("failed to parse path: %w", err)
    }
    
    if err := e.executor.Set(e.data, parsed, value); err != nil {
        return err
    }
    
    return e.storage.Write(e.data)
}
```

**장점**:
- Set + Save가 원자적으로 실행
- 다른 요청이 중간에 끼어들지 않음
- 기존 Engine 인터페이스 최소 변경

**단점**:
- Engine 인터페이스 변경 필요
- lock 점유 시간 증가 (Set + Save 전체)
- 읽기 요청의 응답 시간 증가

## Tradeoffs

| 기준 | Option A (Engine Mutex) | Option B (파일 경로 RWLock) | Option C (원자적 처리) |
|------|------------------------|---------------------------|----------------------|
| **구현 복잡도** | ★★★★★ | ★★★☆☆ | ★★★★☆ |
| **직렬화 보장** | ★★★☆☆ | ★★★★★ | ★★★★★ |
| **성능** | ★★★★☆ | ★★★★☆ | ★★★☆☆ |
| **안정성** | ★★★☆☆ | ★★★★☆ | ★★★★★ |
| **기존 코드 변경** | ★★★★★ | ★★★☆☆ | ★★☆☆☆ |

## Recommendation
**Option B (파일 경로 기반 RWLock)** 추천 이유:

1. **명시적 직렬화**: 파일 수준에서 쓰기 요청을 명확히 직렬화
2. **읽기 병렬화**: RLock으로 읽기 요청은 병렬 허용
3. **Race Condition 해결**: Set + Save 사이 끼어듦 문제 방지
4. **적절한 잠금 범위**: Option C보다 잠금 점유 시간이 짧음

## Consequences
- `Server` 구조에 `fileLocks map[string]*sync.RWMutex` 필드 추가
- `getFileLock(filePath)` 헬퍼 메서드 추가
- 읽기 핸들러: `lock.RLock()` / `lock.RUnlock()` 사용
- 쓰기 핸들러: `lock.Lock()` / `lock.Unlock()` 사용
- Engine의 내부 mutex는 유지 (이중 잠금 구조)
- 서버 시작 시 fileLocks 초기화, 종료 시 정리

## Related ASRs
- ASR-006 — 동시성 제어 — 파일 수준 동시 쓰기 직렬화
- ASR-007 — 엔진 아키텍처 — Engine 인터페이스 유지하면서 서버 레벨 잠금 추가

## Downstream Concerns
- [ ] **데드락 방지**: 잠금 순서 표준화 (항상 같은 순서로 잠금)
- [ ] **잠금 타임아웃**: 잠금 대기 시간 제한 (데드락 감지)

## Related
- `jsondb/adr/server-engine-lifecycle.md` — 상위 ADR (Long-lived Engine 결정)
- `jsondb/adr/concurrency-control.md` — 기존 동시성 제어 결정

## Tags
`server`, `concurrency`, `locking`, `write-serialization`

## Approved
- 2026-07-21: Option B (파일 경로 기반 RWLock), 사용자 승인
