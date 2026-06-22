# Concurrency Model — Per-File Mutex Map Design

## Concern
What level of concurrency safety should jsondb provide, and what synchronization mechanism best supports the REST server architecture?

## Status
approved

## Context
- jsondb는 Library + REST Server 구조로 결정됨 (architecture-client-server ADR)
- 서버는 여러 클라이언트(프로세스)의 동시 요청을 처리
- DB 전체를 단일 lock으로 보호하면 병목 발생 — 다른 파일을 읽는 요청까지 차단
- AI 산출물 저장 시나리오: 여러 클라이언트가 서로 다른 파일을 읽고 쓰는 경우가 대부분
- 동일 파일에 대한 동시 쓰기는 드물지만 안전해야 함
- Go의 `sync.Map` + `sync.RWMutex` 조합으로 구현 가능

## Decision
**Option B — Per-File RWMutex Map**.
- `sync.Map`에 파일 경로별 `sync.RWMutex` 관리
- 같은 파일: reads 공유, write 독점
- 다른 파일: 완전 동시 처리 가능

## Options

### Option A — DB-level RWMutex
- `DB` struct에 `sync.RWMutex` 하나로 전체 DB 보호
- 읽기: RLock (공유), 쓰기: Lock (독점)
- Pro: 가장 단순한 구현
- Con: **서버 환경에서 병목** — 파일 A 읽는 중에 파일 B 읽기/쓰기 모두 차단
- Con: REST server의 동시 처리 이점을 살리지 못함

### Option B — Per-File RWMutex Map (권장)
- 파일 경로별 개별 `sync.RWMutex`를 `sync.Map`으로 관리
- 공유 자원: Mutex map 자체는 thread-safe (`sync.Map`)
- Pro: **서로 다른 파일은 완전히 동시 처리** — 서버 성능 최적
- Pro: 같은 파일에 대한 동시 읽기는 RLock으로 허용
- Pro: 한 파일 쓰기가 다른 파일 작업을 전혀 막지 않음
- Con: Mutex map 메모리 관리 고려 필요
- Con: `List()` 같은 다중 파일 작업은 별도 처리 필요

### Option C — No Internal Sync
- 호출자 책임
- Con: 서버 환경에서 race condition 위험

### Option D — flock (Process-Level)
- `syscall.Flock`으로 OS 파일 lock
- Con: POSIX 전용, 서버 내부 고루틴 동시성에는 부적합

## Tradeoffs (Server 환경 기준)
| | A (DB Lock) | B (Per-File) | C (No Sync) | D (flock) |
|---|------------|-------------|-------------|----------|
| 동시 처리량 | ★ (직렬화) | ★★★★★ (파일별) | ★★★★★ | ★★★ |
| 구현 단순함 | ★★★★★ | ★★★ | ★★★★★ | ★★ |
| 같은 파일 동시 읽기 | ★ (쓰기시 전체block) | ★★★★ (RLock) | ★★★★★ | ★★★★ |
| REST server 적합성 | ★ (병목) | ★★★★★ | ★ (위험) | ★★★ |

---

## 상세 설계: Per-File RWMutex Map

### 자료 구조

```go
type DB struct {
    muMap sync.Map  // map[string]*fileLock  (key: canonicalized file path)

    root       string
    schemaRoot string
}

type fileLock struct {
    rwmu   sync.RWMutex
    refCount int32        // 참조 카운트 — cleanup 용도
}
```

### Lock 획득 흐름

```
1. path 정규화: filepath.Clean(path) + filepath.Join(db.root, path) + ".json"
2. sync.Map.Load()로 기존 fileLock 조회
3. 없으면: 새 fileLock 생성 → sync.Map.LoadOrStore()로 저장
   - LoadOrStore로 race condition 방지
   - 두 고루틴이 동시에 생성해도 하나만 저장됨
4. fileLock.refCount++ (atomic)
5. 읽기: rwmu.RLock(),  쓰기: rwmu.Lock()
6. 작업 수행
7. rwmu.RUnlock() or rwmu.Unlock()
8. fileLock.refCount-- (atomic)
```

### Path 정규화 규칙

```go
func (db *DB) lockKey(dataPath string) string {
    // 1. path separator 통일 (forward slash로)
    // 2. filepath.Clean으로 . .. 정리
    // 3. db.root 기준 절대 경로로 변환
    // 4. .json 확장자 없이 key 생성
    clean := filepath.Clean(filepath.Join(db.root, dataPath))
    return filepath.ToSlash(clean)
}
```

> **주의**: 같은 파일을 가리키는 모든 경로 표현이 동일한 key로 정규화되어야 함
> - `episode/e1` = `./episode/e1` = `episode/./e1` = `episode/../episode/e1`

### Mutex Cleanup 전략 (선택)

작업 종료 후 refCount가 0이고 일정 시간이 지난 fileLock은 map에서 제거 가능.
그러나 **초기 버전에서는 cleanup 생략** — 이유:
- `sync.RWMutex` 크기는 매우 작음 (56 bytes)
- 일반적으로 수백~수천 개의 key로 충분
- cleanup 로직이 오히려 복잡도만 증가
- 필요시 이후 도입 (leak detection 후)

### List() 작업 처리

`List(prefix)`는 filesystem을 직접 walk. 이때는:
- **No lock** — filesystem 자체의 일관성으로 충분
- walk 중 생성/삭제된 파일은 결과에서 누락되거나 추가될 수 있음
- 정합성이 중요한 작업이 아니므로 best-effort로 충분
- 향후 필요시 snapshot 방식 도입 가능

### Schema 파일 Lock

Schema도 파일로 저장되므로 동일한 Mutex Map으로 관리:
- Schema key prefix: `_schema_` (일반 데이터와 구분)
- `db.SetSchema("/blog/episode", schema)` → lockKey = `_schema_/blog/episode`

### Deadlock 안전성

jsondb의 모든 public method는 **단일 파일**에만 접근:
- `Get(path)` → 파일 하나만 read lock
- `Set(path, data)` → 파일 하나만 write lock
- `GetPath(path, jsonPath)` → 파일 하나만 read lock
- `SetPath(path, jsonPath, value)` → 파일 하나만 write lock

→ 복수 파일 lock이 없으므로 **deadlock 가능성 없음**

```go
// 예: SetPath — 단일 파일만 lock
func (db *DB) SetPath(path, jsonPath string, value any) error {
    lock := db.getOrCreateLock(path)
    lock.rwmu.Lock()
    defer lock.rwmu.Unlock()

    // read file → sjson.Set → write file
    data, err := os.ReadFile(db.filePath(path))
    // ...
}
```

## Recommendation (optional)
- **Option B — Per-File RWMutex Map** 으로 결정. 서버 아키텍처에서 최적의 성능을 냄.

## Consequences
- `sync.Map`을 key: 정규화된 파일경로, value: `*fileLock` 으로 구성
- 모든 public method는 작업 전 `getOrCreateLock()` 호출
- 읽기: RLock, 쓰기: Lock (관리자는 panic 방지를 위해 RWMutex 규칙 철저히 준수)
- List()는 lock 없이 filesystem walk
- 초기 버전: mutex cleanup 생략 (메모리 사용량 모니터링 후 필요시 도입)

## Related
- architecture-client-server ADR (REST server 구조)
- go-package-api ADR

## Tags
`concurrency`, `mutex-map`, `per-file-lock`, `thread-safety`

## Approved
- 2026-06-19: Option B (Per-File RWMutex Map), user confirmed for server architecture
