# 서버 종료 시 정리

## Concern
서버가 종료될 때 열린 모든 Engine 인스턴스를 어떻게 안전하게 정리하고 잠금을 해제할 것인가?

## Status
approved

## Context
- `server-engine-lifecycle.md` ADR에서 파일별 Long-lived Engine 채택
- 서버가 시작되면 파일별 Engine 인스턴스를 lazy하게 생성하여 유지
- 서버 종료 시 열린 모든 Engine을 정상적으로 닫아야 함
- 파일 잠금(flock)이 해제되어야 다른 프로세스가 접근 가능
- 비정상 종료 시 잠금이 남을 수 있음 (OS가 자동 해제하지만, 애플리케이션 수준 정리 필요)

## Decision
**Option A — signal 핸들러 + Graceful Shutdown** 사용자 승인: OS 시그널을 수신하여 graceful shutdown 수행, 모든 요청 처리 완료 후 Engine 정리.

## Options

### Option A — signal 핸들러 + Graceful Shutdown

**방식**: OS 시그널(SIGTERM, SIGINT)을 수신하여 graceful shutdown 수행. 모든 요청 처리 완료 후 Engine 정리.

```go
func (s *Server) Start() error {
    mux := http.NewServeMux()
    // ... 핸들러 등록 ...
    
    srv := &http.Server{
        Addr:    s.addr,
        Handler: mux,
    }
    
    // 시그널 핸들러 등록
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    go func() {
        <-sigChan
        fmt.Println("\nShutting down server...")
        s.shutdown()
        srv.Shutdown(context.Background())
    }()
    
    fmt.Printf("Server started on %s\n", s.addr)
    return srv.ListenAndServe()
}

func (s *Server) shutdown() {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    for path, me := range s.engines {
        fmt.Printf("Closing engine for %s\n", path)
        me.engine.Close()
        delete(s.engines, path)
    }
    
    for path, lock := range s.fileLocks {
        // RWMutex는 명시적 해제 불가, GC에 의존
        delete(s.fileLocks, path)
    }
    
    fmt.Println("All engines closed")
}
```

**장점**:
- Graceful shutdown 보장 (진행 중인 요청 완료 후 종료)
- 모든 Engine 정상 닫기
- 파일 잠금 해제
- 표준 Go 패턴

**단점**:
- Graceful shutdown 시간 제한 필요 (무한 대기 방지)
- 시그널 핸들러 구현 필요

---

### Option B — context cancellation 사용

**방식**: Server에 context를 전달하고, context 취소 시 모든 Engine 정리.

```go
func (s *Server) Start(ctx context.Context) error {
    mux := http.NewServeMux()
    // ... 핸들러 등록 ...
    
    srv := &http.Server{
        Addr:    s.addr,
        Handler: mux,
    }
    
    go func() {
        <-ctx.Done()
        fmt.Println("Context cancelled, shutting down...")
        s.shutdown()
        srv.Shutdown(context.Background())
    }()
    
    fmt.Printf("Server started on %s\n", s.addr)
    return srv.ListenAndServe()
}

// 사용 예
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
    cancel()
}()

server.Start(ctx)
```

**장점**:
- Context 기반으로 더 유연한 취소 관리
- 타임아웃 결합 용이
- 테스트 용이성

**단점**:
- context 관리 복잡도
- 이미 context를 사용하는 핸들러와 충돌 가능

---

### Option C — defer 기반 정리

**방식**: Server 시작 시 모든 Engine 생성을 defer로 등록. panic 복구와 함께 정리 보장.

```go
func (s *Server) Start() error {
    defer s.cleanupAll()
    
    mux := http.NewServeMux()
    // ... 핸들러 등록 ...
    
    fmt.Printf("Server started on %s\n", s.addr)
    return http.ListenAndServe(s.addr, mux)
}

func (s *Server) cleanupAll() {
    if r := recover(); r != nil {
        fmt.Printf("Panic recovered: %v\n", r)
    }
    
    s.mu.Lock()
    defer s.mu.Unlock()
    
    for path, me := range s.engines {
        fmt.Printf("Closing engine for %s\n", path)
        me.engine.Close()
        delete(s.engines, path)
    }
}
```

**장pts**:
- 구현이 가장 단순
- panic 발생 시에도 정리 보장
- defer로 자동 실행

**단점**:
- Graceful shutdown 미보장 (즉시 종료)
- 진행 중인 요청 손실 가능
- 시그널 처리 미지원

## Tradeoffs

| 기준 | Option A (signal 핸들러) | Option B (context) | Option C (defer) |
|------|------------------------|-------------------|-----------------|
| **구현 복잡도** | ★★★★☆ | ★★★☆☆ | ★★★★★ |
| **Graceful shutdown** | ★★★★★ | ★★★★★ | ★☆☆☆☆ |
| **안정성** | ★★★★★ | ★★★★☆ | ★★★☆☆ |
| **유연성** | ★★★☆☆ | ★★★★★ | ★★☆☆☆ |
| **테스트 용이성** | ★★★☆☆ | ★★★★★ | ★★★★☆ |

## Recommendation
**Option A (signal 핸들러 + Graceful Shutdown)** 추천 이유:

1. **Graceful shutdown**: 진행 중인 요청을 완료한 후 종료
2. **안정성**: 모든 Engine 정상 닫기, 파일 잠금 해제
3. **표준 패턴**: Go에서 흔히 사용되는 graceful shutdown 패턴
4. **실용성**: Option B보다 구현이 간단하면서 충분한 기능 제공

## Consequences
- `signal.Notify`로 SIGINT, SIGTERM 시그널 수신
- `http.Server.Shutdown()`으로 graceful shutdown 수행
- `Server.shutdown()` 메서드에서 모든 Engine.Close() 호출
- 종료 시 로그 출력으로 사용자에게 알림
- 타임아웃 설정으로 무한 대기 방지 (예: 30초)

## Related ASRs
- ASR-006 — 동시성 제어 — 잠금 해제 보장
- ASR-007 — 엔진 아키텍처 — Engine 생명주기 관리

## Downstream Concerns
- [ ] **비정상 종료 복구**: 서버가 비정상 종료된 후 재시작 시 잠금 상태 확인
- [ ] **종료 타임아웃 설정**: graceful shutdown 최대 대기 시간

## Related
- `jsondb/adr/server-engine-lifecycle.md` — 상위 ADR (Long-lived Engine 결정)
- `jsondb/adr/concurrency-control.md` — 기존 동시성 제어 결정

## Tags
`server`, `shutdown`, `cleanup`, `graceful`, `signal-handling`

## Approved
- 2026-07-21: Option A (signal 핸들러 + Graceful Shutdown), 사용자 승인
