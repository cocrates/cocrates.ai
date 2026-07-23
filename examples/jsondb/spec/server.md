# Server Module Specification

## Requirement

jsondb의 Server 모듈은 HTTP 인터페이스를 제공하여 에이전트와 CLI가 jsondb에 접근할 수 있도록 한다. RESTful API와 JSON Patch/Merge Patch를 지원한다.

## Context

- 에이전트가 HTTP API를 통해 직접 접근
- CLI는 Server에 HTTP 요청 전달
- 엔진 모듈과 연동
- JSON 요청/응답

## Decisions

- **프로토콜**: HTTP/1.1
- **포맷**: JSON
- **기본 포트**: 8080
- **API 버전**: `/api/v1/`
- **JSON Patch**: `github.com/evanphx/json-patch` 라이브러리 사용
- **인증**: 선택적 (API 키)

## Requirements

### 파일 관리 API

- `GET /api/v1/files` — 파일 목록 조회
- `POST /api/v1/files` — 새 파일 생성
- `DELETE /api/v1/files/{path}` — 파일 삭제
- `HEAD /api/v1/files/{path}` — 파일 존재 여부 확인

### 데이터 조회 API

- `GET /api/v1/data/{path}` — 전체 데이터 조회
- `GET /api/v1/data/{path}?field={field}` — 특정 필드 조회
- `GET /api/v1/data/{path}?query={jsonpath}` — JSONPath 쿼리

### 데이터 변경 API

- `PUT /api/v1/data/{path}?field={field}` — 필드 값 설정
- `POST /api/v1/data/{path}` — 데이터 추가 (배열)
- `PATCH /api/v1/data/{path}` — 부분 업데이트
- `DELETE /api/v1/data/{path}?field={field}` — 필드 삭제

### JSON Patch 지원

- `Content-Type: application/json-patch+json` — JSON Patch (RFC 6902)
- `Content-Type: application/merge-patch+json` — JSON Merge Patch (RFC 7396)
- 기본: JSON Merge Patch로 처리

### 응답 형식

- 성공: `{"success": true, "data": {...}}`
- 에러: `{"success": false, "error": {"code": "...", "message": "..."}}`

### 에러 코드

- `FILE_NOT_FOUND` (404)
- `FIELD_NOT_FOUND` (404)
- `INVALID_PATH` (400)
- `INVALID_QUERY` (400)
- `TYPE_MISMATCH` (400)
- `PERMISSION_DENIED` (403)
- `STORAGE_ERROR` (500)
- `VERSION_CONFLICT` (409)

### 에이전트 통합

- 에이전트는 HTTP API를 통해 서버에 직접 접근
- CLI는 사람용 인터페이스 (에이전트가 사용하지 않음)
- HTTP API가 에이전트 통합의 핵심 인터페이스

## Constraints

- Go 언어로 구현
- `github.com/evanphx/json-patch` 라이브러리 사용
- 엔진 모듈과 연동
- RESTful API 설계 준수

## Out of Scope

- WebSocket 실시간 동기화
- gRPC 지원
- Rate Limiting
- OAuth/JWT 인증

## Open Questions

- 없음

## Related

- `jsondb/adr/http-api.md`
- `jsondb/adr/json-patch.md`
- `jsondb/adr/interface-type.md`
- `jsondb/spec/PRD.md`

## Tags

`server`, `http`, `api`, `rest`, `json-patch`
