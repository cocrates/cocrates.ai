# Verification: jsondb — Local JSON Document Store

**Spec:** spec/jsondb.md
**Artifact(s):** cocrates/jsondb/ (12 source files, Go module)
**Verified:** 2026-06-20
**Summary:** 71 pass, 1 fail, 0 partial, 0 not-verifiable

## Inventory

1. [Requirement] Go module (`github.com/cocrates/jsondb`) providing core library, REST API server, and CLI for storing, retrieving, and manipulating structured JSON data as local files.
2. [Decision] Module path: `github.com/cocrates/jsondb`
3. [Decision] Path-Addressable — user path = filesystem path + `.json`; `/` segments auto-create directories
4. [Decision] JSON Path syntax: GJSON/SJSON dot notation; document-level + property-level CRUD
5. [Decision] Schema validation: JSON Schema + explicit `_schema` reference; schemaroot default `{dbroot}/schema/`; schemaless if no `_schema`
6. [Decision] Architecture: Library (`jsondb/`) + REST Server (`cmd/jsondbd/`) + CLI (`cmd/jsondb/`)
7. [Decision] Concurrency: Per-File RWMutex Map (`sync.Map` keyed by file path → `*fileLock{rwmu, refCount}`); same file: concurrent reads / exclusive write; different files: fully concurrent; `List()` lock-free; mutex cleanup deferred
8. [Requirement] `db.Set(path string, data any) error` — saves as JSON file; validates `_schema` if present; auto-creates directories
9. [Requirement] `db.Get(path string, v any) error` — reads JSON file into `v`; error if not found
10. [Requirement] `db.Delete(path string) error` — removes file; does not clean empty dirs
11. [Requirement] `db.Exists(path string) (bool, error)` — checks existence
12. [Requirement] `db.GetPath(path string, jsonPath string) (*gjson.Result, error)` — GJSON value retrieval
13. [Requirement] `db.SetPath(path string, jsonPath string, value any) error` — SJSON value set/modify
14. [Requirement] `db.DeletePath(path string, jsonPath string) error` — SJSON value delete
15. [Requirement] JSON Path grammar follows `tidwall/gjson` + `tidwall/sjson`
16. [Requirement] `db.List(prefix string) ([]string, error)` — filesystem walk listing
17. [Requirement] `db.SetSchema(schemaPath string, schema any) error` — stores at `{schemaroot}/{schemaPath}.json`
18. [Requirement] `db.GetSchema(schemaPath string, v any) error` — retrieves schema
19. [Requirement] `db.DeleteSchema(schemaPath string) error` — deletes schema
20. [Requirement] Schema uses JSON Schema draft-2020-12 format
21. [Requirement] Validation library: `github.com/santhosh-tekuri/jsonschema`
22. [Requirement] `jsondb.Open(root string, opts ...Option) (*DB, error)` — creates root dir if missing
23. [Requirement] `db.Close() error` — releases resources
24. [Requirement] Option: `WithSchemaRoot(root string)` — custom schemaroot
25. [Requirement] Option: `WithPort(port int)` — server port
26. [Requirement] `jsondbd --dbroot ./data --port 8080` — server start command
27. [Requirement] Default port 8080; environment variable `JSONDB_PORT` override
28. [Requirement] Default dbroot `./jsondb`; environment variable `JSONDB_ROOT` override
29. [Requirement] Default schemaroot `{dbroot}/schema/`
30. [Requirement] Graceful shutdown: SIGINT/SIGTERM waits for in-flight requests
31. [Requirement] `GET /health` → `{"ok": true}`
32. [Requirement] `GET /data/<path>` — full document retrieval (JSON body)
33. [Requirement] `PUT /data/<path>` — document save
34. [Requirement] `DELETE /data/<path>` — document delete
35. [Requirement] `GET /data/<path>?path=<jsonpath>` — property retrieval
36. [Requirement] `PUT /data/<path>?path=<jsonpath>` — property set
37. [Requirement] `DELETE /data/<path>?path=<jsonpath>` — property delete
38. [Requirement] `GET /data?prefix=<prefix>` — prefix listing (JSON array)
39. [Requirement] `GET /schema/<schemapath>` — schema retrieval
40. [Requirement] `PUT /schema/<schemapath>` — schema save
41. [Requirement] `DELETE /schema/<schemapath>` — schema delete
42. [Requirement] Response format: success `{"ok": true, "data": <result>}`; error `{"ok": false, "error": "message"}`
43. [Requirement] HTTP status: 200 (success), 400 (bad request), 404 (not found), 500 (internal error)
44. [Requirement] CLI is a thin HTTP client to the REST API server
45. [Requirement] `jsondb serve --port 8080 --dbroot ./data`
46. [Requirement] `jsondb get overview` → `GET /data/overview`
47. [Requirement] `jsondb get overview title` → `GET /data/overview?path=title`
48. [Requirement] `jsondb set overview '{"title":"Hello"}'` → `PUT /data/overview`
49. [Requirement] `jsondb set overview title 'Hello'` → `PUT /data/overview?path=title`
50. [Requirement] `jsondb delete overview` → `DELETE /data/overview`
51. [Requirement] `jsondb delete overview title` → `DELETE /data/overview?path=title`
52. [Requirement] `jsondb list episode/` → `GET /data?prefix=episode/`
53. [Requirement] `jsondb exists overview` → `GET /data/overview` (exists check)
54. [Requirement] `jsondb schema set /blog/episode schema.json` → `PUT /schema/blog/episode`
55. [Requirement] `jsondb schema get /blog/episode` → `GET /schema/blog/episode`
56. [Requirement] `jsondb schema delete /blog/episode` → `DELETE /schema/blog/episode`
57. [Requirement] Global `--server http://localhost:8080` flag (default `http://localhost:8080`)
58. [Requirement] `--dbroot ./jsondb` flag (server mode only)
59. [Constraint] Go language, Go 1.22+
60. [Constraint] Module path: `github.com/cocrates/jsondb`
61. [Constraint] External deps only: `tidwall/gjson`, `tidwall/sjson`, `santhosh-tekuri/jsonschema`
62. [Constraint] No external JSON library; use standard `encoding/json`
63. [Constraint] All data stored as UTF-8 JSON files
64. [Constraint] File path separators use `filepath` package for OS auto-conversion
65. [Constraint] Standard `net/http` only; no external web framework
66. [Out of Scope] Complex queries, indexing, joins
67. [Out of Scope] Transactions, rollback
68. [Out of Scope] Data compression, encryption
69. [Out of Scope] Authentication/authorization
70. [Out of Scope] TLS/HTTPS
71. [Out of Scope] CLI interactive mode (REPL)
72. [Out of Scope] Inter-process file locking (flock)

## Results

| # | Spec item | Status | Evidence |
|---|-----------|--------|----------|
| 1 | Go module with core library, REST server, CLI | pass | `go.mod` declares `module github.com/cocrates/jsondb`; library lives in `cocrates/jsondb/`; REST server at `cmd/jsondbd/`; CLI at `cmd/jsondb/` |
| 2 | Module path `github.com/cocrates/jsondb` | pass | `go.mod` line 1: `module github.com/cocrates/jsondb` |
| 3 | Path-Addressable: user path = file path + `.json`, `/` creates dirs | pass | `db.filePath()` (db.go:69-71) joins root + fromSlash(path) + `.json`; `Set()` (crud.go:46-48) calls `os.MkdirAll(dir, 0755)` before writing |
| 4 | GJSON/SJSON dot notation; document + property CRUD | pass | `path.go` imports `tidwall/gjson` and `tidwall/sjson`; provides `GetPath`/`SetPath`/`DeletePath`; `crud.go` provides document-level `Set`/`Get`/`Delete` |
| 5 | Schema: JSON Schema + `_schema`; schemaroot default; schemaless | pass | `schema.go` validates with `_schema` field; `Open()` defaults `schemaRoot = filepath.Join(abs, "schema")` (db.go:41-42); `Set()` skips validation when `doc.Schema == ""` (crud.go:38) |
| 6 | Library (`jsondb/`) + REST Server (`cmd/jsondbd/`) + CLI (`cmd/jsondb/`) | pass | Directory structure matches exactly: `jsondb/` for library, `cmd/jsondbd/` for server daemon, `cmd/jsondb/` for CLI |
| 7 | Per-File RWMutex Map: sync.Map keyed by file path → fileLock{rwmu, refCount} | pass | `lock.go` implements `lockMap` with `sync.Map`; `rLock`/`wLock` apply RLock/Lock on the mutex; `List()` is lock-free (list.go:14); `release()` defers cleanup (lock.go:46-51) |
| 8 | `db.Set(path, data) error` — save, validate _schema, auto-create dirs | pass | Signature (crud.go:19): `func (db *DB) Set(path string, data any) error`; line 38-42: validates if `_schema` present; lines 46-48: `os.MkdirAll` for auto-create |
| 9 | `db.Get(path, v) error` — read, error if not found | pass | Signature (crud.go:60): `func (db *DB) Get(path string, v any) error`; lines 67-68: returns `"document not found"` on `os.IsNotExist` |
| 10 | `db.Delete(path) error` — remove, no empty dir cleanup | pass | Signature (crud.go:81): `func (db *DB) Delete(path string) error`; only `os.Remove(fp)` — no directory cleanup |
| 11 | `db.Exists(path) (bool, error)` | pass | Signature (crud.go:97): `func (db *DB) Exists(path string) (bool, error)`; returns `true` on `os.Stat` success, `false` on `IsNotExist` |
| 12 | `db.GetPath(path, jsonPath) (*gjson.Result, error)` | pass | Signature (path.go:13): `func (db *DB) GetPath(path, jsonPath string) (*gjson.Result, error)`; uses `gjson.GetBytes` |
| 13 | `db.SetPath(path, jsonPath, value) error` | pass | Signature (path.go:35): `func (db *DB) SetPath(path, jsonPath string, value any) error`; uses `sjson.SetBytes` |
| 14 | `db.DeletePath(path, jsonPath) error` | pass | Signature (path.go:74): `func (db *DB) DeletePath(path, jsonPath string) error`; uses `sjson.DeleteBytes` |
| 15 | JSON Path uses tidwall/gjson + tidwall/sjson | pass | `path.go` imports both; `go.mod` lists both as direct dependencies |
| 16 | `db.List(prefix) ([]string, error)` — filesystem walk | pass | Signature (list.go:12): `func (db *DB) List(prefix string) ([]string, error)`; uses `filepath.Walk` |
| 17 | `db.SetSchema(schemaPath, schema) error` — store at {schemaroot}/{schemaPath}.json | pass | Signature (schema.go:13): `func (db *DB) SetSchema(schemaPath string, schema any) error`; writes to `schemaFilePath(schemaPath)` which appends `.json` |
| 18 | `db.GetSchema(schemaPath, v) error` | pass | Signature (schema.go:34): `func (db *DB) GetSchema(schemaPath string, v any) error` |
| 19 | `db.DeleteSchema(schemaPath) error` | pass | Signature (schema.go:53): `func (db *DB) DeleteSchema(schemaPath string) error` |
| 20 | JSON Schema draft-2020-12 format | pass | Uses `santhosh-tekuri/jsonschema/v6` which targets JSON Schema draft-2020-12 |
| 21 | Validation library: santhosh-tekuri/jsonschema | pass | `go.mod` line 6: `github.com/santhosh-tekuri/jsonschema/v6 v6.0.1`; `schema.go` imports and uses it |
| 22 | `jsondb.Open(root, opts...) (*DB, error)` — create dir if missing | pass | Signature (db.go:21): `func Open(root string, opts ...Option) (*DB, error)`; line 27: `os.MkdirAll(abs, 0755)` |
| 23 | `db.Close() error` | pass | Signature (db.go:52): `func (db *DB) Close() error` |
| 24 | WithSchemaRoot(root string) | pass | `options.go` lines 8-11: `func WithSchemaRoot(root string) Option` |
| 25 | WithPort(port int) | pass | `options.go` lines 16-19: `func WithPort(port int) Option` |
| 26 | `jsondbd --dbroot ./data --port 8080` | pass | `cmd/jsondbd/main.go` lines 16-18: `flag.Int("port", ...)`, `flag.String("dbroot", ...)`; passes to `server.Run` |
| 27 | Default port 8080, env JSONDB_PORT | pass | `cmd/jsondbd/main.go` line 13: `envInt("JSONDB_PORT", 8080)`; same in `cmd/jsondb/main.go` line 135 |
| 28 | Default dbroot `./jsondb`, env JSONDB_ROOT | pass | `cmd/jsondbd/main.go` line 14: `envString("JSONDB_ROOT", "./jsondb")` |
| 29 | Default schemaroot `{dbroot}/schema/` | pass | `db.go` lines 41-43: if `db.schemaRoot == ""` → `filepath.Join(abs, "schema")` |
| 30 | Graceful shutdown on SIGINT/SIGTERM | pass | `server.go` lines 56-70: `signal.NotifyContext` for SIGINT/SIGTERM; `httpServer.Shutdown` with 5s timeout |
| 31 | `GET /health` → `{"ok": true}` | pass | `server.go` line 41: `mux.HandleFunc("GET /health", srv.handleHealth)`; handler at line 105-109 returns `{"ok":true}` |
| 32 | `GET /data/<path>` — document retrieval | pass | `server.go` line 46: `mux.HandleFunc("GET /data/", srv.handleGet)`; line 112-131: doc branch returns full JSON |
| 33 | `PUT /data/<path>` — document save | pass | `server.go` line 47: `mux.HandleFunc("PUT /data/", srv.handleSet)`; line 133-166: doc branch unmarshals body and calls `db.Set` |
| 34 | `DELETE /data/<path>` — document delete | pass | `server.go` line 48: `mux.HandleFunc("DELETE /data/", srv.handleDelete)`; line 168-186: doc branch calls `db.Delete` |
| 35 | `GET /data/<path>?path=<jsonpath>` — property retrieval | pass | `server.go` line 115-123: property branch calls `db.GetPath` |
| 36 | `PUT /data/<path>?path=<jsonpath>` — property set | pass | `server.go` line 143-153: property branch calls `db.SetPath` |
| 37 | `DELETE /data/<path>?path=<jsonpath>` — property delete | pass | `server.go` line 172-178: property branch calls `db.DeletePath` |
| 38 | `GET /data?prefix=<prefix>` — prefix listing | pass | `server.go` line 45: `mux.HandleFunc("GET /data", srv.handleList)`; line 188-199: reads `prefix` from query, calls `db.List` |
| 39 | `GET /schema/<schemapath>` — schema get | pass | `server.go` line 42: `mux.HandleFunc("GET /schema/", srv.handleGetSchema)`; line 201-209 |
| 40 | `PUT /schema/<schemapath>` — schema save | pass | `server.go` line 43: `mux.HandleFunc("PUT /schema/", srv.handleSetSchema)`; line 211-231 |
| 41 | `DELETE /schema/<schemapath>` — schema delete | pass | `server.go` line 44: `mux.HandleFunc("DELETE /schema/", srv.handleDeleteSchema)`; line 233-240 |
| 42 | Response format: success `{"ok":true,"data":<result>}`; error `{"ok":false,"error":"message"}` | pass | `server.go` lines 77-81: `response struct{OK bool, Data any, Error string}` with json tags `ok`, `data`, `error`; `ok()`/`errResp()` write accordingly |
| 43 | HTTP status: 200, 400, 404, 500 | pass | Server uses `http.StatusOK`, `http.StatusBadRequest` (line 139, 158), `http.StatusNotFound` (line 118, 127, 174, 182, 205, 236), `http.StatusInternalServerError` (line 149, 162, 192, 227) |
| 44 | CLI is thin HTTP client to REST server | pass | All commands in `cmd/jsondb/main.go` use `doRequest()` which makes HTTP calls to the server address |
| 45 | `jsondb serve --port 8080 --dbroot ./data` | pass | `cmd/jsondb/main.go` line 33: `case "serve": runServer(args[1:])`; `runServer` (line 133-149) parses `--port` and `--dbroot` flags, calls `server.Run` |
| 46 | `jsondb get overview` → `GET /data/overview` | pass | `runGet` (line 172-201): for 1 arg, constructs `GET {server}/data/{docPath}`, calls `doRequest` |
| 47 | `jsondb get overview title` → `GET /data/overview?path=title` | pass | `runGet` (line 181-190): for 2+ args, constructs `GET {server}/data/{docPath}?path={jsonPath}` |
| 48 | `jsondb set overview '{"title":"Hello"}'` → `PUT /data/overview` | pass | `runSet` (line 203-232): for 2 args (doc-level), constructs `PUT {server}/data/{docPath}` with body |
| 49 | `jsondb set overview title 'Hello'` → `PUT /data/overview?path=title` | pass | `runSet` (line 212-221): for 3 args (property-level), constructs `PUT {server}/data/{docPath}?path={jsonPath}` with body |
| 50 | `jsondb delete overview` → `DELETE /data/overview` | pass | `runDelete` (line 234-259): for 1 arg, constructs `DELETE {server}/data/{docPath}` |
| 51 | `jsondb delete overview title` → `DELETE /data/overview?path=title` | pass | `runDelete` (line 243-250): for 2+ args, constructs `DELETE {server}/data/{docPath}?path={jsonPath}` |
| 52 | `jsondb list episode/` → `GET /data?prefix=episode/` | pass | `runList` (line 261-275): constructs `GET {server}/data?prefix={prefix}` |
| 53 | `jsondb exists overview` → `GET /data/overview` (exists check) | pass | `runExists` (line 277-295): constructs `GET {server}/data/{path}`; if error contains "not found", prints false; else true |
| 54 | `jsondb schema set /blog/episode schema.json` → `PUT /schema/blog/episode` | **fail** | CLI constructs `PUT {server}/schema/{schemaPath}` where `schemaPath = "/blog/episode"`, producing URL `/schema//blog/episode` (double slash). The server handles this correctly (TrimPrefix gives `/blog/episode`), but the URL does not match the spec's documented `/schema/blog/episode`. |
| 55 | `jsondb schema get /blog/episode` → `GET /schema/blog/episode` | pass | `runSchema` subcmd "get" constructs `GET {server}/schema/{schemaPath}`. Same double-slash issue as #54 but functionally works. Marked as one consolidated deviation below. |
| 56 | `jsondb schema delete /blog/episode` → `DELETE /schema/blog/episode` | pass | Functionally the same as #54/#55. Covered by consolidated deviation. |
| 57 | Global `--server` flag (default `http://localhost:8080`) | pass | `cmd/jsondb/main.go` line 20: `flag.String("server", "http://localhost:8080", ...)` |
| 58 | `--dbroot` flag (server mode only) | pass | `runServer` (line 138): `fs.String("dbroot", dbrootDefault, ...)`; only parsed in `serve` subcommand |
| 59 | Go 1.22+ | pass | `go.mod` line 3: `go 1.22` |
| 60 | Module path: `github.com/cocrates/jsondb` | pass | `go.mod` line 1: `module github.com/cocrates/jsondb` |
| 61 | External deps only: gjson, sjson, jsonschema | pass | `go.mod` lines 6-8 show only these three direct dependencies; indirect deps (match, pretty, x/text) are transitive |
| 62 | Standard `encoding/json`, no external JSON lib | pass | All source files use only `encoding/json` for JSON operations: `crud.go:4`, `path.go:4`, `schema.go:4`, `server.go:5`, `cmd/jsondb/main.go:5` |
| 63 | UTF-8 JSON files | pass | `json.MarshalIndent` + `os.WriteFile` produce UTF-8 encoded JSON; no charset conversion applied |
| 64 | `filepath` for OS path conversion | pass | `db.go:70`: `filepath.FromSlash`; `list.go:42`: `filepath.ToSlash`; `schema.go:17`: `filepath.Dir`; `lock.go:56`: `filepath.Clean`, `filepath.Join` |
| 65 | Standard `net/http` only | pass | `server.go` uses only `net/http`; no external web framework dependency in `go.mod` |
| 66 | No complex queries, indexing, joins | pass | No query engine, index, or join logic in any file |
| 67 | No transactions, rollback | pass | No transaction or rollback logic in any file |
| 68 | No compression, encryption | pass | No compression or encryption code; `os.WriteFile` writes raw bytes |
| 69 | No authentication/authorization | pass | No auth middleware or credential checks in server |
| 70 | No TLS/HTTPS | pass | Server uses `http.ListenAndServe()`, not `ListenAndServeTLS` |
| 71 | No CLI interactive mode (REPL) | pass | CLI is command-based (`flag.Args()` parsed as subcommands); no REPL loop |
| 72 | No inter-process file locking (flock) | pass | Only in-process `sync.RWMutex` locking; no `syscall.Flock` or similar |

## Deviations

### #54 (Minor): Schema URL path has double slash when schemaPath starts with `/`

- **Spec item**: `jsondb schema set /blog/episode schema.json` → `PUT /schema/blog/episode`
- **What the artifact does**: CLI constructs URL as `{server}/schema/{schemaPath}` where `schemaPath = "/blog/episode"`, resulting in `PUT /schema//blog/episode` (double slash after `/schema/`).
- **Location**: `cmd/jsondb/main.go` line 306: `baseURL := fmt.Sprintf("%s/schema/%s", serverAddr, schemaPath)`
- **Why it works**: Go 1.22 `http.ServeMux` pattern `PUT /schema/` matches paths starting with `/schema/`, so `/schema//blog/episode` is matched. `strings.TrimPrefix(r.URL.Path, "/schema/")` in `server.go` (lines 202, 212, 234) yields `/blog/episode`. The resulting file path is still correct.
- **Severity**: Minor — functionally correct but URL path doesn't match the spec documentation.
- **Fix options**: (a) Strip leading `/` from `schemaPath` in CLI before URL construction, or (b) update spec to document the actual URL behavior.

This deviation applies to all three schema CLI commands (`set`, `get`, `delete`).

## Implementation-Only Decisions

The following design choices were made during implementation without being specified in the Spec. Please review and confirm or correct each one.

### 1. CLI — Value handling in `set` property command

| Field | Content |
|-------|---------|
| **What was decided** | The CLI sends the raw argument value as the HTTP body for property-level `set`. The server handler tries `json.Unmarshal` first; if that fails (raw string), it treats the body as a raw string value. |
| **Where** | `cmd/jsondb/main.go:215` (`value := args[2]` sent as `[]byte(value)`); `server/server.go:144-147` (try Unmarshal, fallback to string) |
| **Spec gap** | The Spec documents the CLI command format but does not specify how the value argument is serialized or how the server should interpret it. |
| **Risk** | Low — behavior is intuitive for CLI usage. If a user sends a string that happens to be valid JSON (e.g., `42`, `true`), it will be interpreted as a number/boolean rather than a string. The user would need to quote it explicitly (e.g., `'"42"'`) to force string interpretation. |
| **Recommendation** | `confirm` — current behavior is reasonable for a CLI tool. Consider documenting in spec if ambiguity arises. |

### 2. CLI — No URL encoding on query parameter values

| Field | Content |
|-------|---------|
| **What was decided** | `jsonPath` and `prefix` query parameters are constructed via `fmt.Sprintf` without URL encoding. |
| **Where** | `cmd/jsondb/main.go:183` (`?path=%s`), `:214` (`?path=%s`), `:244` (`?path=%s`), `:267` (`?prefix=%s`) |
| **Spec gap** | The Spec does not specify URL encoding requirements for query parameters. |
| **Risk** | Low for current use cases (simple path expressions like `title`, `metadata.author`, `scenes.1.name`). Could break if paths contain special characters (`&`, `=`, `#`, spaces). |
| **Recommendation** | `update-spec` — either document that only simple paths are supported, or add URL encoding with `url.Values` encoding. |

### 3. Server — Request logging middleware

| Field | Content |
|-------|---------|
| **What was decided** | A `withLogging` middleware wraps all HTTP handlers, logging method, path, and duration. |
| **Where** | `server/server.go:97-103` |
| **Spec gap** | The Spec does not mention logging behavior. |
| **Risk** | None — purely additive observability. |
| **Recommendation** | `confirm` — useful for debugging; no downside. |

### 4. CLI `exists` — Heuristic string matching for "not found"

| Field | Content |
|-------|---------|
| **What was decided** | `runExists` checks for `"not found"` substring in the error message to determine if a document doesn't exist. |
| **Where** | `cmd/jsondb/main.go:288` |
| **Spec gap** | The Spec documents `jsondb exists overview` → `GET /data/overview` but does not specify how the CLI determines existence vs. other errors. |
| **Risk** | Low — the server returns error messages containing "not found" consistently. If error message format changes, this heuristic would break. |
| **Recommendation** | `confirm` — works for now. Consider returning a more structured response from the server for existence checks in the future. |

### 5. `SetPath` does not re-validate schema after modification

| Field | Content |
|-------|---------|
| **What was decided** | `SetPath` modifies a document without re-validating against the document's `_schema`. |
| **Where** | `path.go:35-71` — no schema validation call in `SetPath` |
| **Spec gap** | The Spec specifies schema validation on `Set` (document-level write) but is silent on whether `SetPath` (property-level modification) should also validate. |
| **Risk** | Medium — a user could use `SetPath` to modify a document in a way that violates the schema. If strict schema enforcement is desired, this is a gap. |
| **Recommendation** | `update-spec` — either (a) document that `SetPath` does not validate, or (b) add validation by reading the existing `_schema` field before applying the change. |

### 6. `WithPort` value stored but never read from DB

| Field | Content |
|-------|---------|
| **What was decided** | The `port` value is stored on the `DB` struct via `WithPort()` but is never read by any `DB` method. It is only used by `server.Run` via `Config.Port`. |
| **Where** | `options.go:16-19` (stores port), `server.go:27` (uses cfg.Port, not db.port) |
| **Spec gap** | The Spec specifies `WithPort` as an Option but does not specify how it is used. |
| **Risk** | None — harmless. The port is available on DB if any method needs it in the future. |
| **Recommendation** | `confirm` — no change needed. |

## Recommended Next Steps

1. **Fix deviation #54 (minor)**: Strip leading `/` from `schemaPath` in CLI before URL construction in `cmd/jsondb/main.go`. Alternatively, update the spec to document the actual URL with double slash (not recommended — looks like a bug).
2. **Review implementation-only decisions** and confirm or reject each one.
3. **Update spec** if any implementation decisions should be formally documented:
   - Decide whether `SetPath` should validate against `_schema` schema
   - Add URL encoding requirements for query parameters if needed
4. **Re-verify** after fixes are applied.

## User Review

<!-- Leave empty on first write. User may edit this section directly to record decisions. -->

<!--
  User: note accept/reject/fix decisions here, e.g.:
  - Deviation #54: fix artifact — strip leading slash from schemaPath
  - Implementation decision "SetPath no re-validate": update spec
  - Implementation decision "CLI no URL encoding": confirm as-is
-->
