# 서버 엔진 생명주기 관리

## Concern
서버 핸들러가 매 HTTP 요청마다 Engine의 Open/Close를 반복하여 성능 저하와 Race Condition이 발생한다. 이 문제를 어떻게 해결해야 하는가?

## Status
approved

## Context
- 현재 서버 핸들러의 모든 요청이 `s.engine.Open(path)` → 작업 → `defer s.engine.Close()` 패턴 사용
- Engine은 단일 `data map[string]interface{}`를 보유
- **Race Condition**: Request A가 Open(fileA) 후 Request B가 Open(fileB)하면 fileA 데이터가 덮어씌워짐
- **성능**: 읽기 전용 요청도 매번 디스크 I/O + JSON 파싱 반복
- **동시성 제어(ASR-006)**: 파일 잠금 + RWLock 혼합이 승인되었으나, 요청 레벨에서 미적용
- **엔진 아키텍처(ASR-007)**: 모듈 기반 설계가 승인되었으나, 생명주기 관리 미설계

## Decision
**Option D — Long-lived Engine + Write-through Cache** 사용자 승인: 파일별 Engine을 서버 생명주기 동안 유지하고, 읽기는 메모리에서, 쓰기는 즉시 디스크에 반영하는 방식으로 Race Condition과 성능 문제를 동시에 해결.

## Options

### Option A — 파일별 Long-lived Engine

**방식**: 서버 시작 시 활성 파일 목록을 기반으로 Engine 인스턴스를 미리 생성하고, 서버 종료 시 닫는다. 요청이 들어오면 기존 Engine을 재사용한다.

```go
type Server struct {
    engines map[string]engine.Engine  // 파일 경로별 Engine 관리
    mu      sync.RWMutex
    addr    string
    root    string
}

func (s *Server) getEngine(filePath string) engine.Engine {
    s.mu.RLock()
    eng, ok := s.engines[filePath]
    s.mu.RUnlock()
    if ok {
        return eng
    }
    // 없으면 새로 생성
    s.mu.Lock()
    defer s.mu.Unlock()
    // Double-check
    if eng, ok = s.engines[filePath]; ok {
        return eng
    }
    backend := storage.NewDirectIOBackend()
    eng = engine.New(backend)
    eng.Open(filepath.Join(s.root, filePath))
    s.engines[filePath] = eng
    return eng
}
```

**장점**:
- 파일 디스크립터 재사용 → `open()`/`close()` 오버헤드 제거
- JSON 파일을 한 번만 파싱 → 반복 I/O 제거
- Race Condition 해소 (파일별 독립 Engine)

**단점**:
- 메모리 사용량 증가 (활성 파일 수 × 데이터 크기)
- 파일 외부 변경 감지 불가 (다른 프로세스가 수정 시 stale 데이터)
- 엔진 수명 관리 복잡도 (open/close lifecycle)

---

### Option B — 요청별 독립 Engine (Pool)

**방식**: 매 요청마다 독립적인 Engine 인스턴스를 생성하고 사용 후 반환한다. Engine 풀을 통해 인스턴스 재사용을 최적화한다.

```go
type EnginePool struct {
    pool    sync.Pool
    root    string
}

func (p *EnginePool) Get(filePath string) (engine.Engine, func()) {
    eng := p.pool.Get().(engine.Engine)
    eng.Open(filepath.Join(p.root, filePath))
    
    release := func() {
        eng.Close()
        p.pool.Put(eng)
    }
    return eng, release
}

// 핸들러 사용 예
func (s *Server) handleDataGet(w http.ResponseWriter, r *http.Request, filePath string) {
    eng, release := s.enginePool.Get(filePath)
    defer release()
    
    result, err := eng.Get("")
    // ...
}
```

**장점**:
- Race Condition 해소 (요청별 독립 인스턴스)
- 구현이 직관적
- 메모리 관리 용이 (요청 종료 시 자동 해제)

**단점**:
- 여전히 매 요청마다 파일 I/O + JSON 파싱
- `open()`/`close()` 오버헤드 잔존
- 성능 개선 없음 (정합성만 해결)

---

### Option C — In-Memory Cache + Lazy Invalidation

**방식**: Engine과 Storage 사이에 캐시 레이어를 추가한다. 파일 수정 시 캐시를 무효화하고, 외부 변경 감지를 위해 파일 modification time을 주기적으로 확인한다.

```go
type CachedEngine struct {
    engines    map[string]*cachedEntry
    mu         sync.RWMutex
    root       string
    pollInterval time.Duration
}

type cachedEntry struct {
    engine    engine.Engine
    modTime   time.Time
    lastCheck time.Time
}

func (ce *CachedEngine) Get(filePath string) (interface{}, error) {
    entry := ce.getOrCreate(filePath)
    
    // 주기적으로 파일 변경 확인
    if time.Since(entry.lastCheck) > ce.pollInterval {
        ce.checkAndInvalidate(entry, filePath)
        entry.lastCheck = time.Now()
    }
    
    return entry.engine.Get("")
}

func (ce *CachedEngine) checkAndInvalidate(entry *cachedEntry, filePath string) {
    info, err := os.Stat(filepath.Join(ce.root, filePath))
    if err != nil {
        return
    }
    if info.ModTime().After(entry.modTime) {
        // 파일이 외부에서 변경됨 → 캐시 무효화
        entry.engine.Close()
        entry.engine.Open(filepath.Join(ce.root, filePath))
        entry.modTime = info.ModTime()
    }
}
```

**장점**:
- 읽기 성능 극대화 (캐시 �트 시 I/O 없음)
- 외부 변경 감지 가능 (파일 modification time)
- 기존 Engine 인터페이스 유지

**단점**:
- 구현 복잡도 높음 (캐시 정책, 무효화 로직)
- 메모리 사용량 증가
- polling 오버헤드 (주기적 파일 stat)
- 캐시 일관성 보장이 어려움 (write-through vs write-back)

---

### Option D — Long-lived Engine + Write-through Cache

**방식**: Option A(Long-lived Engine)과 Option C(In-Memory Cache)를 결합한다. 파일별 Engine을 서버 생명주기 동안 유지하되, 읽기는 메모리 캐시에서, 쓰기는 즉시 디스크에 반영(write-through)한다.

```go
type Server struct {
    engines    map[string]*managedEngine
    mu         sync.RWMutex
    addr       string
    root       string
}

type managedEngine struct {
    engine   engine.Engine
    modTime  time.Time
    dirty    bool
}

func (s *Server) getOrCreateEngine(filePath string) (*managedEngine, error) {
    fullPath := filepath.Join(s.root, filePath)
    
    s.mu.RLock()
    me, ok := s.engines[fullPath]
    s.mu.RUnlock()
    if ok {
        return me, nil
    }
    
    s.mu.Lock()
    defer s.mu.Unlock()
    
    // Double-check
    if me, ok = s.engines[fullPath]; ok {
        return me, nil
    }
    
    // 새 Engine 생성 및 열기
    backend := storage.NewDirectIOBackend()
    eng := engine.New(backend)
    if err := eng.Open(fullPath); err != nil {
        return nil, err
    }
    
    info, _ := os.Stat(fullPath)
    me = &managedEngine{
        engine:  eng,
        modTime: info.ModTime(),
    }
    s.engines[fullPath] = me
    return me, nil
}

func (s *Server) handleDataGet(w http.ResponseWriter, r *http.Request, filePath string) {
    me, err := s.getOrCreateEngine(filePath)
    if err != nil {
        s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
        return
    }
    // Close하지 않음 — Engine 유지
    
    result, err := me.engine.Get("")
    if err != nil {
        s.writeError(w, http.StatusNotFound, "FIELD_NOT_FOUND", err.Error())
        return
    }
    
    s.writeJSON(w, http.StatusOK, map[string]interface{}{
        "success": true,
        "data":    result,
    })
}

func (s *Server) handleDataSet(w http.ResponseWriter, r *http.Request, filePath string) {
    me, err := s.getOrCreateEngine(filePath)
    if err != nil {
        s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
        return
    }
    
    // Set + Save (write-through)
    if err := me.engine.Set(field, value); err != nil {
        s.writeError(w, http.StatusBadRequest, "SET_FAILED", err.Error())
        return
    }
    if err := me.engine.Save(); err != nil {
        s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", err.Error())
        return
    }
    
    // modTime 갱신
    info, _ := os.Stat(filepath.Join(s.root, filePath))
    me.modTime = info.ModTime()
    
    s.writeJSON(w, http.StatusOK, map[string]interface{}{"success": true})
}
```

**장점**:
- 파일별 Engine 유지 → I/O 오버헤드 최소화
- Write-through → 데이터 일관성 보장
- Race Condition 해소 (파일별 독립 Engine)
- 외부 변경 감지 가능 (modTime 비교)

**단점**:
- 메모리 사용량 증가 (활성 파일 수 × 데이터 크기)
- 파일별 Engine 관리 복잡도
- 서버 시작 시 모든 파일을 열 필요 없음 (lazy)

## Tradeoffs

| 기준 | Option A (Long-lived) | Option B (Pool) | Option C (Cache) | Option D (Long-lived + Cache) |
|------|----------------------|-----------------|-------------------|-------------------------------|
| **성능** | ★★★★☆ | ★★☆☆☆ | ★★★★★ | ★★★★★ |
| **정합성** | ★★★★☆ | ★★★★★ | ★★★☆☆ | ★★★★★ |
| **구현 복잡도** | ★★★★☆ | ★★★★★ | ★★☆☆☆ | ★★★☆☆ |
| **메모리 효율** | ★★★☆☆ | ★★★★★ | ★★★☆☆ | ★★★☆☆ |
| **외부 변경 감지** | ★★☆☆☆ | ★★★★★ | ★★★★☆ | ★★★★☆ |
| **확장성** | ★★★☆☆ | ★★★★☆ | ★★★★★ | ★★★★☆ |

## Recommendation
**Option D (Long-lived Engine + Write-through Cache)** 추천 이유:

1. **성능**: 파일별 Engine 유지로 반복 I/O 파싱 제거
2. **정합성**: 파일별 독립 Engine으로 Race Condition 해소
3. **일관성**: Write-through로 쓰기 즉시 디스크 반영
4. **유연성**: Lazy 생성으로 불필요한 리소스 할당 방지
5. **Phase 1에 적합**: 기존 Engine 인터페이스 최대한 재사용, 점진적 개선 가능

**구현 전략**:
1. Phase 1: Option D 구현 (기본 기능)
2. Phase 2: 외부 변경 감지 강화 (inotify/kqueue)
3. Phase 3: 메모리 최적화 (LRU 기반 Engine 관리)

## Consequences
- 서버 시작 시 모든 파일을 열지 않음 (lazy loading)
- 파일별 Engine 인스턴스를 map으로 관리
- 읽기 요청: Engine.Get() 직접 호출 (I/O 없음)
- 쓰기 요청: Engine.Set() + Engine.Save() (write-through)
- 서버 종료 시 모든 Engine Close()
- `server.go`의 핸들러 패턴 변경 필요

## Related ASRs
- ASR-006 — 동시성 제어 — 요청 레벨 Race Condition 해결
- ASR-007 — 엔진 아키텍처 — Engine 생명주기 관리 설계 추가

## Downstream Concerns
- [ ] **파일 외부 변경 감지**: 다른 프로세스가 파일을 수정할 때如何感知? (inotify/kqueue)
- [ ] **메모리 관리**: 활성 파일 수 제한, LRU 기반 Engine 관리
- [ ] **서버 종료 시 정리**: 열린 Engine 정상 닫기, 잠금 해제
- [ ] **동시 쓰기 처리**: 같은 파일에 대한 동시 쓰기 요청 직렬화

## Related
- `jsondb/adr/concurrency-control.md` — 기존 동시성 제어 결정
- `jsondb/adr/engine-architecture.md` — 기존 엔진 아키텍처 결정
- `jsondb/spec/engine.md` — 엔진 모듈 스펙

## Tags
`server`, `engine`, `lifecycle`, `performance`, `concurrency`

## Approved
- 2026-07-21: Option D (Long-lived Engine + Write-through Cache), 사용자 승인
