# HTTP API 설계

## Concern
Server 모듈이 제공할 HTTP API의 엔드포인트와 요청/응답 형식은 어떤 형태로 설계해야 하는가?

## Status
approved

## Context
- 하이브리드 인터페이스 (CLI + Server)
- 엔진 모듈과 연동
- RESTful API 설계
- JSON 요청/응답

## Decision
**RESTful API 설계** 사용자 승인: 파일 관리 + 데이터 조회/변경 엔드포인트, JSON Patch/Merge Patch 지원.

## API 개요

### 기본 정보
- **프로토콜**: HTTP/1.1
- **포맷**: JSON
- **인증**: 선택적 (API 키 또는 토큰)
- **기본 포트**: 8080

### URL 구조
```
http://localhost:8080/api/v1/{resource}
```

## 엔드포인트 설계

### 1. 파일 관리

| 메서드 | 엔드포인트 | 설명 |
|--------|-----------|------|
| `GET` | `/api/v1/files` | 파일 목록 조회 |
| `POST` | `/api/v1/files` | 새 파일 생성 |
| `DELETE` | `/api/v1/files/{path}` | 파일 삭제 |
| `HEAD` | `/api/v1/files/{path}` | 파일 존재 여부 확인 |

### 2. 데이터 조회

| 메서드 | 엔드포인트 | 설명 |
|--------|-----------|------|
| `GET` | `/api/v1/data/{path}` | 전체 데이터 조회 |
| `GET` | `/api/v1/data/{path}?field={field}` | 특정 필드 조회 |
| `GET` | `/api/v1/data/{path}?query={jsonpath}` | JSONPath 쿼리 |

### 3. 데이터 변경

| 메서드 | 엔드포인트 | 설명 |
|--------|-----------|------|
| `PUT` | `/api/v1/data/{path}?field={field}` | 필드 값 설정 |
| `POST` | `/api/v1/data/{path}` | 데이터 추가 (배열) |
| `PATCH` | `/api/v1/data/{path}` | 부분 업데이트 (JSON Patch/Merge Patch) |
| `DELETE` | `/api/v1/data/{path}?field={field}` | 필드 삭제 |

> **참고**: PATCH 메서드의 상세 사양은 `jsondb/adr/json-patch.md` 참조

## 상세 설계

### 1. 파일 목록 조회

**요청**:
```http
GET /api/v1/files?dir={directory}
```

**응답**:
```json
{
  "success": true,
  "data": {
    "files": [
      {
        "name": "overview.json",
        "path": "overview.json",
        "size": 1024,
        "modified": "2026-07-21T10:30:00Z"
      },
      {
        "name": "settings.json",
        "path": "settings.json",
        "size": 512,
        "modified": "2026-07-21T09:15:00Z"
      }
    ],
    "directories": [
      {
        "name": "projects",
        "path": "projects"
      }
    ]
  }
}
```

### 2. 데이터 조회

**요청**:
```http
GET /api/v1/data/overview.json
GET /api/v1/data/overview.json?field=title
GET /api/v1/data/overview.json?query=$.items[?(@.active==true)]
```

**응답**:
```json
{
  "success": true,
  "data": {
    "title": "프로젝트 개요",
    "content": "이 프로젝트는...",
    "tags": ["important", "draft"]
  },
  "meta": {
    "path": "overview.json",
    "size": 1024,
    "modified": "2026-07-21T10:30:00Z"
  }
}
```

### 3. 데이터 변경

**요청**:
```http
PUT /api/v1/data/overview.json?field=title
Content-Type: application/json

{
  "value": "새 제목"
}
```

**응답**:
```json
{
  "success": true,
  "data": {
    "path": "overview.json",
    "field": "title",
    "old_value": "기존 제목",
    "new_value": "새 제목"
  }
}
```

### 4. 부분 업데이트

**요청**:
```http
PATCH /api/v1/data/overview.json
Content-Type: application/json

{
  "title": "새 제목",
  "tags": ["important", "updated"]
}
```

**응답**:
```json
{
  "success": true,
  "data": {
    "path": "overview.json",
    "updated_fields": ["title", "tags"],
    "version": 2
  }
}
```

## 에러 응답 형식

### 표준 에러 응답
```json
{
  "success": false,
  "error": {
    "code": "FIELD_NOT_FOUND",
    "message": "필드를 찾을 수 없습니다",
    "details": {
      "path": "overview.json",
      "field": "nonexistent"
    }
  }
}
```

### 에러 코드

| 코드 | HTTP 상태 | 설명 |
|------|----------|------|
| `FILE_NOT_FOUND` | 404 | 파일을 찾을 수 없음 |
| `FIELD_NOT_FOUND` | 404 | 필드를 찾을 수 없음 |
| `INVALID_PATH` | 400 | 잘못된 경로 |
| `INVALID_QUERY` | 400 | 잘못된 JSONPath 쿼리 |
| `TYPE_MISMATCH` | 400 | 타입 불일치 |
| `PERMISSION_DENIED` | 403 | 접근 권한 없음 |
| `STORAGE_ERROR` | 500 | 저장소 오류 |
| `VERSION_CONFLICT` | 409 | 버전 충돌 (동시성) |

## 인증/인가

### API 키 방식
```http
Authorization: Bearer {api-key}
```

### 설정
```json
{
  "auth": {
    "enabled": false,
    "type": "api-key",
    "keys": ["your-api-key-here"]
  }
}
```

## 성능 고려

### 캐싱
```http
Cache-Control: max-age=3600
ETag: "abc123"
```

### 배치 처리
```http
POST /api/v1/batch
Content-Type: application/json

{
  "operations": [
    {"method": "GET", "path": "overview.json", "field": "title"},
    {"method": "PUT", "path": "settings.json", "field": "theme", "value": "dark"}
  ]
}
```

## Downstream Concerns
- [ ] **WebSocket 지원**: 실시간 업데이트 알림
- [ ] **gRPC 지원**: 고성능 내부 통신
- [ ] **API 버전 관리**: 향후 API 변경 대응
- [ ] **Rate Limiting**: 요청 제한
- [ ] **JSON Patch 상세 사양**: `jsondb/adr/json-patch.md` 참조

## Tags
`http`, `api`, `rest`, `json`, `json-patch`

## Approved
- 2026-07-21: RESTful API 설계, 사용자 승인