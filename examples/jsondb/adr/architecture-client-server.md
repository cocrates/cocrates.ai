# Architecture — Client-Server vs Library-Only

## Concern
jsondb가 여러 프로세스의 동시 접근을 지원하려면 어떤 아키텍처가 적합한가?

## Status
approved

## Context
- 기존 설계: Go library + CLI (단일 프로세스 가정, RWMutex로 고루틴 안전)
- 문제: "동시에 여러 프로세스에서 요청하는 경우" — library-only로는 프로세스 간 동시성 해결 불가
- 사용자 제안: REST API 서버를 두고, 복수 프로세스가 HTTP로 요청
- RWMutex는 서버 내부 엔진에서 thread-safe 유지
- AI 산출물 저장 시나리오에서 여러 프로세스(AI 에이전트)가 동시에 같은 DB에 접근할 가능성

## Decision
**Option B — Library + REST Server**.
- `jsondb/` core library (thread-safe engine)
- `cmd/jsondb/` CLI (REST client or direct library use)
- `cmd/jsondbd/` REST API server daemon
- REST API: JSON over HTTP on localhost

## Options

### Option A — Library-Only + CLI (현재 설계)
- 구조: `jsondb/` library + `cmd/jsondb/` CLI
- 동시성: 단일 프로세스 내 고루틴 안전 (RWMutex)
- 다중 프로세스: **지원하지 않음** (운영체제 파일 write 충돌 가능)
- 사용: Go 코드에서 `import` 하여 사용, 또는 CLI로 단일 명령 실행
- Pro: 가장 단순함. 외부 의존성 없음. 데몬 관리 불필요.
- Pro: 제로 네트워크 오버헤드
- Con: **여러 프로세스가 동시에 같은 DB 접근 시 문제**
- Con: Go가 아닌 언어에서 사용하려면 cgo 또는 subprocess 호출 필요

### Option B — Library + REST Server (권장)
- 구조:
  ```
  jsondb/           — core library (thread-safe engine, RWMutex)
  cmd/jsondb/       — CLI client (REST API 호출)
  cmd/jsondbd/      — REST API server daemon
  ```
- REST API: `GET /data/<path>`, `PUT /data/<path>`, `DELETE /data/<path>` 등
- 통신: JSON over HTTP (localhost, port 지정)
- 동시성: 서버 내부는 RWMutex로 thread-safe, 여러 클라이언트는 HTTP worker goroutine으로 처리
- Pro: **여러 프로세스/언어에서 동시 접근 가능**
- Pro: 라이브러리도 그대로 유지 — Go로 직접 import해서 쓸 수도 있음
- Pro: 서버는 필요할 때만 실행 (개발 시에는 라이브러리만 사용 가능)
- Pro: HTTP 기반이므로 모니터링, 로깅, 인증 추가 용이
- Con: 네트워크 오버헤드 (로컬host라도 직렬화/역직렬화)
- Con: 서버 프로세스 관리 필요 (시작/중지, 포트 충돌 등)
- Con: REST API 설계 및 유지보수 추가

### Option C — Server-Only (REST API 전용)
- 구조: `cmd/jsondbd/` 서버만 제공. 별도 library 없음.
- 모든 클라이언트는 HTTP로만 통신
- Pro: 단일 인터페이스 — 모두 HTTP
- Con: Go 코드에서 직접 import해서 쓸 수 없음 — 모든 사용이 HTTP
- Con: AI가 Go 코드를 생성할 때 항상 HTTP call 필요 — 복잡도 증가
- 평가: **유연성 부족**으로 비추천

### Option D — Library + gRPC Server
- REST 대신 gRPC 사용
- Pro: 엄격한 계약(.proto), 스트리밍 가능, 성능 우수
- Con: gRPC 도구 체인 필요(protoc, 코드 생성), 무거움
- Con: curl 등 범용 도구로 테스트 불편
- 평가: 간단한 로컬 DB에 **over-engineering**

## Tradeoffs

| 기준 | A (Library) | B (Lib+REST) | C (Server) | D (gRPC) |
|------|-------------|-------------|-----------|---------|
| 다중 프로세스 지원 | ★ | ★★★★★ | ★★★★★ | ★★★★★ |
| 구현 단순함 | ★★★★★ | ★★★ | ★★★ | ★★ |
| Go import 사용 | ★★★★★ | ★★★★★ | ★ (HTTP only) | ★ (gRPC client) |
| 범용성 (non-Go) | ★ | ★★★★ (HTTP) | ★★★★ (HTTP) | ★★★ (gRPC) |
| CLI 유용성 | ★★★ (직접 파일 조작) | ★★★★ (REST client) | ★★★★ | ★★★ |
| AI 코드 생성 용이 | ★★★★★ | ★★★★ | ★★★ | ★★★ |

## REST API 제안 (Option B 기준)

```
# Document CRUD
GET    /data/<path>              → 문서/속성 조회 (Accept: application/json)
PUT    /data/<path>              → 문서 저장 (body: JSON)
DELETE /data/<path>              → 문서 삭제

# Property CRUD (query parameter로 JSON path 지정)
GET    /data/<path>?path=<jsonpath>    → 특정 속성 조회
PUT    /data/<path>?path=<jsonpath>    → 특정 속성 설정 (body: value)
DELETE /data/<path>?path=<jsonpath>    → 특정 속성 삭제

# Schema
GET    /schema/<schemapath>      → 스키마 조회
PUT    /schema/<schemapath>      → 스키마 저장 (body: JSON Schema)
DELETE /schema/<schemapath>      → 스키마 삭제

# Listing
GET    /data?prefix=<prefix>     → 문서 목록

# Health
GET    /health                   → 서버 상태

# Response format: {"ok": true, "data": ..., "error": "..."}
```

## CLI 동작 방식 (Option B 기준)
```
# Server mode
jsondb serve --port 8080 --dbroot ./mydb

# Client mode (기본: localhost:8080)
jsondb get overview                          → GET /data/overview
jsondb get overview title                    → GET /data/overview?path=title
jsondb set overview title "Hello"             → PUT /data/overview?path=title
jsondb list episode/                          → GET /data?prefix=episode/
jsondb schema set /blog/episode schema.json   → PUT /schema/blog/episode
```

## Recommendation (optional)
- **Option B — Library + REST Server** 를 차기 버전 방향으로 제안합니다.
  - 1단계: 먼저 `jsondb/` library + `cmd/jsondb/` CLI (with server client mode) 구현
  - 2단계: `cmd/jsondbd/` 서버 구현 — library를 내부에서 사용 (thin HTTP wrapper)
  - 이렇게 하면 처음에는 library-only로 빠르게 개발하고, 이후 서버를 추가하는 식으로 점진적 확장 가능

## Consequences
- Library의 모든 기능은 thread-safe (RWMutex) — 서버와 직접 사용 모두 대응
- CLI는 서버 모드(`serve`)와 클라이언트 모드(기본)로 동작
- 서버는 Go 표준 `net/http` 사용 (외부 웹 프레임워크 불필요)
- REST API port는 `--port` 또는 환경변수로 설정

## Related
- concurrency-model ADR (RWMutex thread-safety)
- go-package-api ADR (package structure)

## Tags
`architecture`, `client-server`, `rest-api`, `concurrency`

## Approved
- 2026-06-19: Option B (Library + REST Server), user confirmed
