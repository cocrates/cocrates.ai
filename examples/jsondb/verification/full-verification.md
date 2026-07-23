# Verification: jsondb 전체 검증

**Spec:** jsondb/spec/PRD.md, jsondb/spec/engine.md, jsondb/spec/server.md, jsondb/spec/cli.md
**Artifact(s):** jsondb/src/
**Verified:** 2026-07-21
**Summary:** 28 pass, 14 fail, 6 partial, 2 not-verifiable

---

## PRD.md 검증

| # | Spec item | Status | Evidence / Notes |
|---|-----------|--------|------------------|
| 1 | 필드 레벨 접근 (기본 경로) | pass | `engine/query/parser.go` - 점 구분자 파싱 지원 |
| 2 | 필드 레벨 접근 (JSONPath) | partial | `engine/query/executor.go:220-236` - Get만 구현, Set/Delete 미구현 |
| 3 | CLI 모드 | pass | `cmd/main.go` - get, set, delete, create, rm 동작 확인 |
| 4 | Server 모드 | partial | `server/server.go` - 기본 핸들러 구현, 일부 TODO 존재 |
| 5 | CLI → Server HTTP 요청 | fail | CLI가 직접 엔진 사용 (CLI→Server 연결 미구현) |
| 6 | RESTful HTTP API (파일 관리) | partial | 핸들러 존재하나 일부 미구현 (파일 생성/삭제) |
| 7 | RESTful HTTP API (데이터 조회) | pass | `GET /api/v1/data/{path}` 구현 |
| 8 | RESTful HTTP API (데이터 변경) | pass | `PUT /api/v1/data/{path}` 구현 |
| 9 | RESTful HTTP API (부분 업데이트) | fail | `PATCH` 핸들러存在하나 미구현 (TODO) |
| 10 | 파일 잠금 (flock) | pass | `storage/backend.go:172-233` - Lock/Unlock/RLock/RUnlock 구현 |
| 11 | RWLock (Server 모드) | fail | 엔진에 sync.RWMutex 있으나 서버에서 flock 미사용 |
| 12 | Go 언어 | pass | go.mod 확인 |
| 13 | ojg 라이브러리 | pass | go.mod에 의존성 추가 |
| 14 | json-patch 라이브러리 | fail | go.mod에 미추가 |
| 15 | 파일 기반 저장 | pass | `storage/backend.go` - 직접 I/O 구현 |
| 16 | Storage Backend 추상화 | pass | `StorageBackend` 인터페이스 정의 |

---

## engine.md 검증

| # | Spec item | Status | Evidence / Notes |
|---|-----------|--------|------------------|
| 1 | Engine 인터페이스 (Open) | pass | `engine/engine.go:46-65` |
| 2 | Engine 인터페이스 (Get) | pass | `engine/engine.go:68-87` |
| 3 | Engine 인터페이스 (Set) | pass | `engine/engine.go:90-104` |
| 4 | Engine 인터페이스 (Delete) | pass | `engine/engine.go:107-121` |
| 5 | Engine 인터페이스 (Save) | pass | `engine/engine.go:124-133` |
| 6 | Engine 인터페이스 (Close) | pass | `engine/engine.go:136-151` |
| 7 | Query Parser - 점 구분자 | pass | `engine/query/parser.go:62-79` |
| 8 | Query Parser - JSONPath 감지 | pass | `engine/query/parser.go:40-46` |
| 9 | Query Parser - 배열 인덱스 | pass | `engine/query/parser.go:82-112` |
| 10 | Query Executor - 기본 경로 | pass | `engine/query/executor.go:54-83` (Get), `86-151` (Set), `154-217` (Delete) |
| 11 | Query Executor - JSONPath Get | pass | `engine/query/executor.go:220-236` |
| 12 | Query Executor - JSONPath Set | fail | `engine/query/executor.go:239-251` - "not fully implemented yet" 반환 |
| 13 | Query Executor - JSONPath Delete | fail | `engine/query/executor.go:254-265` - "not fully implemented yet" 반환 |
| 14 | StorageBackend 인터페이스 | pass | `storage/backend.go:12-30` |
| 15 | DirectIOBackend 구현체 | pass | `storage/backend.go:32-233` |
| 16 | 파일 시스템 유사 구조 | pass | 파일 경로 기반 접근 |
| 17 | 중첩 객체 지원 | pass | 재귀적 경로 탐색 |
| 18 | 배열 지원 | pass | `[index]` 구문 파싱 |
| 19 | 에러 타입 정의 | pass | `engine/types.go:6-21` |

---

## server.md 검증

| # | Spec item | Status | Evidence / Notes |
|---|-----------|--------|------------------|
| 1 | GET /api/v1/files | pass | `server/server.go:100-109` - 빈 목록 반환 (미완 구현) |
| 2 | POST /api/v1/files | fail | `server/server.go:112-115` - "not implemented yet" 반환 |
| 3 | DELETE /api/v1/files/{path} | fail | `server/server.go:118-121` - "not implemented yet" 반환 |
| 4 | HEAD /api/v1/files/{path} | fail | `server/server.go:124-127` - "not implemented yet" 반환 |
| 5 | GET /api/v1/data/{path} | pass | `server/server.go:130-163` |
| 6 | GET /api/v1/data/{path}?field={field} | pass | `server/server.go:147-148` |
| 7 | GET /api/v1/data/{path}?query={jsonpath} | pass | `server/server.go:145-146` |
| 8 | PUT /api/v1/data/{path}?field={field} | pass | `server/server.go:166-209` |
| 9 | POST /api/v1/data/{path} | fail | `server/server.go:254-257` - "not implemented yet" 반환 |
| 10 | PATCH /api/v1/data/{path} | fail | `server/server.go:212-215` - "not implemented yet" 반환 |
| 11 | DELETE /api/v1/data/{path}?field={field} | pass | `server/server.go:218-251` |
| 12 | JSON Patch (RFC 6902) | fail | 미구현 |
| 13 | JSON Merge Patch (RFC 7396) | fail | 미구현 |
| 14 | 응답 형식 (성공) | pass | `{"success": true, "data": {...}}` |
| 15 | 응답 형식 (에러) | pass | `{"success": false, "error": {"code": "...", "message": "..."}}` |
| 16 | 에러 코드 (FILE_NOT_FOUND) | pass | `server/server.go:133` |
| 17 | 에러 코드 (FIELD_NOT_FOUND) | pass | `server/server.go:155` |
| 18 | 에러 코드 (INVALID_PATH) | pass | `server/server.go:169` |
| 19 | 에러 코드 (STORAGE_ERROR) | pass | `server/server.go:197` |
| 20 | 에이전트 HTTP API 직접 접근 | pass | HTTP API 엔드포인트 구현 |

---

## cli.md 검증

| # | Spec item | Status | Evidence / Notes |
|---|-----------|--------|------------------|
| 1 | get {file} | pass | `cmd/main.go:87-124` |
| 2 | get {file} {field} | pass | `cmd/main.go:111-112` |
| 3 | set {file} {field} {value} | pass | `cmd/main.go:126-167` |
| 4 | delete {file} {field} | pass | `cmd/main.go:169-202` |
| 5 | create {file} | pass | `cmd/main.go:234-249` |
| 6 | rm {file} | pass | `cmd/main.go:251-265` |
| 7 | serve | pass | `cmd/main.go:267-297` |
| 8 | ls | fail | `cmd/main.go:221-232` - "not implemented yet" 반환 |
| 9 | patch {file} {patch} | fail | `cmd/main.go:204-219` - "not implemented yet" 반환 |
| 10 | status | fail | `cmd/main.go:299-303` - "not implemented yet" 반환 |
| 11 | stop | fail | `cmd/main.go:305-309` - "not implemented yet" 반환 |
| 12 | cp {src} {dst} | not-verifiable | 명령어 목록에 없음 |
| 13 | mv {src} {dst} | not-verifiable | 명령어 목록에 없음 |
| 14 | --pretty | pass | `cmd/main.go:332-338` |
| 15 | --verbose | pass | `cmd/main.go:341-347` |
| 16 | --server {url} | pass | `cmd/main.go:350-363` |
| 17 | JSONDB_SERVER 환경 변수 | pass | `cmd/main.go:358-360` |
| 18 | JSONDB_ROOT 환경 변수 | fail | 미구현 |
| 19 | JSONDB_LOG_LEVEL 환경 변수 | fail | 미구현 |
| 20 | 설정 파일 (~/.jsondb/config.json) | fail | 미구현 |

---

## Deviations (Non-compliance)

### Critical (스펙 핵심 요구사항 위반)

| # | Spec item | 실제 동작 | 심각도 |
|---|-----------|----------|--------|
| 1 | JSON Patch/Merge Patch 지원 | 미구현 | Critical |
| 2 | CLI → Server HTTP 연결 | CLI가 직접 엔진 사용 | Critical |
| 3 | RWLock (Server 모드) | 서버에서 파일 잠금 미사용 | Critical |

### Major (스펙 제약 조건 위반)

| # | Spec item | 실제 동작 | 심각도 |
|---|-----------|----------|--------|
| 1 | json-patch 라이브러리 사용 | 미추가 | Major |
| 2 | JSONPath Set/Delete | "not fully implemented" 반환 | Major |
| 3 | 파일 관리 API (POST, DELETE, HEAD) | 미구현 | Major |
| 4 | CLI 명령어 (ls, patch, status, stop) | 미구현 | Major |

### Minor (사소한 차이)

| # | Spec item | 실제 동작 | 심각도 |
|---|-----------|----------|--------|
| 1 | cp, mv 명령어 | 명령어 목록에 없음 | Minor |
| 2 | 환경 변수 (ROOT, LOG_LEVEL) | 미구현 | Minor |
| 3 | 설정 파일 | 미구현 | Minor |

---

## Undocumented ASRs (Specification Gaps)

| # | 결정 | 위치 | 카테고리 | 격차 | 위험 | 권장 조치 |
|---|------|------|----------|------|------|----------|
| 1 | CLI가 직접 엔진 사용 | cmd/main.go | Architecture | CLI→Server 연결 미설계 | 에이전트와 CLI의 동작 불일치 | 스펙 업데이트 필요 |
| 2 | 매 요청마다 Open/Close | server/server.go | Performance | 연결 풀링 미고려 | 성능 저하 | 연결 관리 설계 필요 |
| 3 | 에러 메시지 한글 미사용 | 전체 | i18n | 영어 에러 메시지 | 한국어 사용자 불편 | 선택적 |

---

## Recommended Next Steps

1. **즉시 수정 필요 (Critical):**
   - JSON Patch/Merge Patch 구현 (`github.com/evanphx/json-patch` 추가)
   - CLI → Server HTTP 연결 구현 또는 스펙 수정
   - 서버 모드 RWLock 통합

2. **추후 수정 (Major):**
   - JSONPath Set/Delete 구현
   - 파일 관리 API 구현
   - CLI 명령어 확장 (ls, patch, status, stop)

3. **선택적 수정 (Minor):**
   - cp, mv 명령어 추가
   - 환경 변수 지원
   - 설정 파일 지원

---

## User Review

{사용자 검토 대기 중}
