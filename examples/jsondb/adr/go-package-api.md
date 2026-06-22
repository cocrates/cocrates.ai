# Go Package Structure and Public API

## Concern
How should the jsondb Go module be structured and what public API should it expose?

## Status
approved

## Context
- jsondb is a Go library for local file-based JSON document storage
- Core features: path-addressable CRUD, partial JSON path access, optional schema validation
- Target users: Go programs that need to persist structured JSON data locally
- AI-generated artifacts: the primary use case — API should be simple enough for AI to generate usage code
- Go convention: "simple, flat, composable"

## Decision
**Option C — Core Library + Standalone CLI**.
- 라이브러리 패키지: `jsondb/` — Go 프로그램에서 import하여 사용
- CLI 바이너리: `cmd/jsondb/` — 쉘에서 직접 데이터 조작
- 두 가지 CRUD 레벨 모두 제공:
  - **문서 전체 CRUD**: `Set`, `Get`, `Delete`
  - **문서 내 속성 CRUD**: `SetPath`, `GetPath`, `DeletePath`

## Options

### Option A — Flat Single Package (`jsondb`)
- 단일 Go package: `import "github.com/cocrates/jsondb"`
- Public API: `Open()`, `Close()`, `Set()`, `Get()`, `Delete()`, `List()`, `SetPath()`, `GetPath()`, `DeletePath()`, `SetSchema()`, `GetSchema()`
- Core type: `DB struct` — all methods on this type
- Pro: **Go의 권장 방식** — flat is better than nested
- Pro: 사용자 입장에서 단일 import로 모든 기능 사용
- Pro: AI가 import 한 줄로 사용법을 기억하기 쉬움
- Con: 패키지 내 코드가 커지면 파일을 분산시켜야 함 (한 패키지 여러 파일)

### Option B — Core + Subpackages
- `jsondb` (core CRUD) + `jsondb/schema` (validation) + `jsondb/store` (storage backends)
- Pro: 관심사 분리 — 각 패키지의 책임이 명확
- Pro: schema 기능이 필요 없는 사용자는 core만 import
- Con: 사용자 입장에서 여러 패키지 import 불편
- Con: 순환 참조 위험
- Con: 초기 단계에 과도한 추상화

### Option C — Core + Standalone CLI
- 라이브러리(`jsondb`) + 별도 CLI 바이너리(`cmd/jsondb/`)
- Pro: CLI로 직접 데이터 조작 가능 (`jsondb get episode/e1 metadata.author`)
- Con: CLI 설계와 유지보수 추가 부담
- Con: 초기 범위를 벗어남 — "먼저 라이브러리부터"

## Tradeoffs
| | A (Flat) | B (Subpackages) | C (CLI) |
|---|----------|-----------------|---------|
| 단순성 | ★★★★★ | ★★★ | ★★ |
| Import 편의 | ★★★★★ 단일 import | ★★★ 여러 import | ★★★★ |
| 관심사 분리 | ★★★ (파일로 분리) | ★★★★★ | ★★★★ |
| 초기 개발 속도 | ★★★★★ 빠름 | ★★ 느림 (설계 부담) | ★ 느림 |

## Recommendation (optional)
- 사용자 결정에 따라 **Option C (Core + CLI)** 로 확정.
  - 라이브러리는 flat package 유지 (`import "github.com/cocrates/jsondb"`)
  - CLI는 `cmd/jsondb/` 에 위치, library API의 thin wrapper 역할
  - CLI 명령어 예시:
    - `jsondb get overview` — 문서 전체 조회
    - `jsondb get overview title` — 특정 속성 조회
    - `jsondb set overview title "New Title"` — 속성 설정
    - `jsondb delete overview title` — 속성 삭제
    - `jsondb list episode/` — prefix 목록
    - `jsondb schema set /blog/episode schema.json` — 스키마 등록

## Proposed Public API

```go
// Open opens or creates a jsondb at the given root directory.
db, err := jsondb.Open(root string, opts ...Option)

// Close flushes and closes the database.
db.Close()

// --- Document-level CRUD ---

// Set stores data at the given path as a JSON file.
// If data has a _schema field, validates against that schema.
err := db.Set(path string, data any)

// Get retrieves and unmarshals the JSON document at path.
data, err := db.Get(path string, v any)

// Delete removes the JSON file at path.
err := db.Delete(path string)

// --- Partial-access CRUD (GJSON/SJSON path within a document) ---

// SetPath sets a specific value within a JSON document.
// path: document path (e.g. "episode/e1")
// jsonPath: GJSON path within the document (e.g. "metadata.author")
err := db.SetPath(path, jsonPath string, value any)

// GetPath retrieves a specific value from a JSON document.
result, err := db.GetPath(path, jsonPath string)  // returns *gjson.Result

// DeletePath removes a specific value from a JSON document.
err := db.DeletePath(path, jsonPath string)

// --- Listing ---

// List returns all document paths under the given prefix.
paths, err := db.List(prefix string)  // e.g., List("episode/")

// Exists returns true if a document exists at the given path.
ok, err := db.Exists(path string)

// --- Schema Management ---

// SetSchema stores a JSON Schema at schemaPath under schemaroot.
err := db.SetSchema(schemaPath string, schema any)

// GetSchema retrieves a JSON Schema.
schema, err := db.GetSchema(schemaPath string, v any)

// DeleteSchema removes a schema.
err := db.DeleteSchema(schemaPath string)

// --- Options ---
// jsondb.WithSchemaRoot(root string) Option  // custom schemaroot
```

## Consequences
- Module path: `github.com/cocrates/jsondb`
- 라이브러리: 단일 `jsondb` 패키지 (flat), 필요시 파일 분리
- CLI: `cmd/jsondb/` 디렉토리, library API의 thin wrapper
- 모든 CRUD는 `DB` struct의 메서드
- 문서 레벨과 속성 레벨 CRUD 모두 제공

## Related
- storage-model (path-addressable)
- json-path-syntax (GJSON/SJSON)
- schema-validation (explicit reference)

## Tags
`go-package`, `api-design`, `module-structure`

## Approved
- 2026-06-19: Option C (Core + CLI), user confirmed — document-level + property-level CRUD 모두 필요
