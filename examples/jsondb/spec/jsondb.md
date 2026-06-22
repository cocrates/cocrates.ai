# jsondb — Local JSON Document Store

## Requirement
Go module (`github.com/cocrates/jsondb`) providing a core library, REST API server, and CLI for storing, retrieving, and manipulating structured JSON data as local files. Designed for AI-generated intermediate artifacts that need structured persistence beyond raw markdown.

## Context
- AI가 생성하는 중간 산출물(overview, outline, episodes 등)을 구조화된 JSON으로 저장
- Markdown과 달리 JSON이면 필요한 속성만 추출/변경 가능
- 여러 프로세스(AI 에이전트)가 동시에 같은 DB에 접근 가능해야 함
- 단순한 로컬 파일 기반 — 외부 DB 의존성 없음
- 샘플 프로젝트 (Cocrates와 무관)

## Decisions
- **Module path**: `github.com/cocrates/jsondb`
- **저장 모델**: Path-Addressable — 사용자 지정 path가 그대로 파일 경로가 됨
  - `db.Set("overview", data)` → `{dbroot}/overview.json`
  - `db.Set("episode/e1", data)` → `{dbroot}/episode/e1.json`
  - path의 `/` 세그먼트 기준 디렉토리 자동 생성
- **JSON Path 문법**: GJSON/SJSON (`tidwall/gjson` + `tidwall/sjson`)
  - dot notation: `metadata.author`, `scenes.1.name`
  - document-level CRUD와 property-level CRUD를 모두 제공
- **스키마 검증**: JSON Schema + explicit `_schema` reference
  - `{schemaroot}` 디렉토리에 스키마 파일 저장 (기본값: `{dbroot}/schema/`)
  - 데이터 파일의 `_schema` 필드가 `{schemaroot}` 기준 경로로 스키마 참조
  - `_schema`가 없으면 schemaless (검증 없이 자유 저장)
  - `_schema`가 있으면 write 시 JSON Schema 검증 수행
- **아키텍처**: Library + REST Server
  - `jsondb/` — core library (thread-safe engine)
  - `cmd/jsondbd/` — REST API server daemon
  - `cmd/jsondb/` — CLI (REST client)
- **동시성 모델**: Per-File RWMutex Map
  - `sync.Map` keyed by 정규화된 파일 경로 → `*fileLock{rwmu, refCount}`
  - 같은 파일: 동시 읽기(RLock) 허용, 쓰기(Lock)는 독점
  - 다른 파일: 완전 동시 처리 가능
  - `List()`만 lock 없이 filesystem walk (best-effort)
  - 초기 버전: mutex cleanup 생략

## Requirements

### Library Core API (`jsondb` package)

#### Document-Level CRUD
- `db.Set(path string, data any) error` — path에 JSON 파일로 저장. `data`에 `_schema`가 있으면 검증. 디렉토리는 자동 생성.
- `db.Get(path string, v any) error` — path의 JSON 파일을 읽어 `v`에 unmarshal. 파일이 없으면 error.
- `db.Delete(path string) error` — path의 JSON 파일 삭제. 빈 디렉토리는 정리하지 않음.
- `db.Exists(path string) (bool, error)` — path에 파일 존재 여부 확인.

#### Property-Level CRUD (JSON Path)
- `db.GetPath(path string, jsonPath string) (*gjson.Result, error)` — 문서 내 특정 JSON 값 조회
- `db.SetPath(path string, jsonPath string, value any) error` — 문서 내 특정 JSON 값 설정/변경
- `db.DeletePath(path string, jsonPath string) error` — 문서 내 특정 JSON 값 삭제
- JSON Path 문법은 `tidwall/gjson` + `tidwall/sjson` 규칙을 따름

#### Listing
- `db.List(prefix string) ([]string, error)` — 주어진 prefix 아래 모든 문서 path 반환 (filesystem walk)

#### Schema Management
- `db.SetSchema(schemaPath string, schema any) error` — `{schemaroot}/{schemaPath}.json`에 JSON Schema 저장
- `db.GetSchema(schemaPath string, v any) error` — Schema 조회
- `db.DeleteSchema(schemaPath string) error` — Schema 삭제
- Schema는 JSON Schema 표준(draft-2020-12) 포맷 사용
- 검증 라이브러리: `github.com/santhosh-tekuri/jsonschema`

#### DB Lifecycle
- `jsondb.Open(root string, opts ...Option) (*DB, error)` — DB 열기/생성. root 디렉토리가 없으면 생성.
- `db.Close() error` — DB 닫기 (리소스 정리)
- Option: `jsondb.WithSchemaRoot(root string)` — schemaroot 커스터마이징
- Option: `jsondb.WithPort(port int)` — 서버 포트 (서버 모드에서 사용)

### REST API Server (`cmd/jsondbd/`)

#### Server Lifecycle
- `jsondbd --dbroot ./data --port 8080` — 서버 시작
- 기본 포트: 8080 (환경변수 `JSONDB_PORT`로도 설정 가능)
- 기본 dbroot: `./jsondb` (환경변수 `JSONDB_ROOT`로도 설정 가능)
- 기본 schemaroot: `{dbroot}/schema/`
- Graceful shutdown: SIGINT/SIGTERM 시 연결 중인 요청 완료 후 종료

#### REST Endpoints

```
# Health
GET    /health                          → {"ok": true}

# Document CRUD
GET    /data/<path>                     → 전체 문서 조회 (JSON body)
PUT    /data/<path>                     → 문서 저장 (body: JSON)
DELETE /data/<path>                     → 문서 삭제

# Property CRUD (query param ?path=<jsonpath>)
GET    /data/<path>?path=<jsonpath>      → 특정 속성 조회 (raw value)
PUT    /data/<path>?path=<jsonpath>      → 특정 속성 설정 (body: raw value)
DELETE /data/<path>?path=<jsonpath>      → 특정 속성 삭제

# Listing
GET    /data?prefix=<prefix>            → prefix 하위 문서 목록 (JSON array)

# Schema CRUD
GET    /schema/<schemapath>             → 스키마 조회
PUT    /schema/<schemapath>             → 스키마 저장 (body: JSON Schema)
DELETE /schema/<schemapath>             → 스키마 삭제
```

#### Response Format
```json
// 성공
{"ok": true, "data": <결과>}

// 오류
{"ok": false, "error": "message"}
```

- HTTP status: 200(성공), 400(잘못된 요청), 404(없음), 500(내부 오류)

### CLI (`cmd/jsondb/`)

CLI는 REST API 서버에 HTTP 요청을 보내는 thin client.

```bash
# Server mode (CLI가 직접 서버 실행)
jsondb serve --port 8080 --dbroot ./data

# Client mode (기존 서버에 요청)
jsondb get overview                          → GET /data/overview
jsondb get overview title                    → GET /data/overview?path=title
jsondb set overview '{"title":"Hello"}'      → PUT /data/overview
jsondb set overview title 'Hello'            → PUT /data/overview?path=title
jsondb delete overview                       → DELETE /data/overview
jsondb delete overview title                 → DELETE /data/overview?path=title
jsondb list episode/                         → GET /data?prefix=episode/
jsondb exists overview                       → GET /data/overview (exists check)
jsondb schema set /blog/episode schema.json  → PUT /schema/blog/episode
jsondb schema get /blog/episode              → GET /schema/blog/episode
jsondb schema delete /blog/episode           → DELETE /schema/blog/episode
```

#### 글로벌 플래그
- `--server http://localhost:8080` (기본값: `http://localhost:8080`)
- `--dbroot ./jsondb` (server mode에서만 사용)

## Constraints
- **Go** 언어 (Go 1.22+)
- Module path: `github.com/cocrates/jsondb`
- 외부 의존성: `tidwall/gjson`, `tidwall/sjson`, `santhosh-tekuri/jsonschema` 만 사용
- 외부 JSON 라이브러리 사용 금지 (표준 `encoding/json` 사용)
- 모든 데이터는 UTF-8 JSON 파일로 저장
- 파일 경로 구분자는 OS에 따라 자동 변환 (`filepath` 패키지 사용)
- 표준 `net/http` 사용 (외부 웹 프레임워크 불가)

## Out of Scope
- 복잡한 쿼리, 인덱싱, Join
- 트랜잭션, 롤백
- 데이터 압축, 암호화
- 인증/인가 (초기 버전)
- TLS/HTTPS (초기 버전은 localhost 전제)
- CLI interactive mode (REPL)
- 프로세스 간 파일 lock (flock 등) — 서버가 단일 접점

## Open Questions
- (없음 — 모든 결정 완료)

## Related
- adr/storage-model.md
- adr/json-path-syntax.md
- adr/schema-validation.md
- adr/go-package-api.md
- adr/architecture-client-server.md
- adr/concurrency-model.md

## Tags
`jsondb`, `go`, `json-storage`, `local-db`, `rest-api`
