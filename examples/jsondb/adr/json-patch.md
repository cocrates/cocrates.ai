# JSON Patch/Merge Patch 지원

## Concern
부분 업데이트(Partial Update)를 어떤 방식으로 제공해야 하는가?

## Status
approved

## Context
- RESTful API 설계에서 PATCH 메서드 사용
- 필드 레벨 변경이 핵심 기능
- Go 언어로 구현
- 표준 기반 접근 선호

## Decision
**JSON Patch (RFC 6902) + JSON Merge Patch (RFC 7396) 모두 지원** 사용자 승인: `github.com/evanphx/json-patch` 라이브러리 사용.

## 공식 표준

### 1. RFC 6902 - JSON Patch

**목적**: JSON 문서에 대한 변경 작업 시퀀스를 정의

**연산 타입**:
| 연산 | 설명 | 예시 |
|------|------|------|
| `add` | 필드 추가 | `{"op": "add", "path": "/title", "value": "새 제목"}` |
| `remove` | 필드 삭제 | `{"op": "remove", "path": "/title"}` |
| `replace` | 필드 교체 | `{"op": "replace", "path": "/title", "value": "새 제목"}` |
| `move` | 필드 이동 | `{"op": "move", "from": "/old", "path": "/new"}` |
| `copy` | 필드 복사 | `{"op": "copy", "from": "/source", "path": "/target"}` |
| `test` | 값 테스트 | `{"op": "test", "path": "/title", "value": "기존 제목"}` |

**장점**:
- **표준**: RFC로 정의된 공식 표준
- **강력**: 복잡한 변경 시퀀스 가능
- **원자성**: 여러 변경을 하나의 트랜잭션으로 처리

**단점**:
- **복잡성**: 구문이 복잡함
- **JSON 형식**: 별도의 Content-Type 필요

### 2. RFC 7396 - JSON Merge Patch

**목적**: JSON 객체를 병합하는 간단한 방법

**동작 원리**:
- `null` 값은 필드 삭제
- 그 외 값은 덮어쓰기 (병합)

**장점**:
- **단순성**: 배우기 쉬움
- **직관성**: 자연스러운 JSON 병합

**단점**:
- **제한적**: 배열 조작 불가
- **단순 병합만 가능**: 복잡한 변경 불가

## 비교

| 기준 | JSON Patch | JSON Merge Patch |
|------|-----------|------------------|
| **표준** | RFC 6902 | RFC 7396 |
| **단순성** | ★★☆☆☆ | ★★★★★ |
| **강력함** | ★★★★★ | ★★★☆☆ |
| **배열 조작** | ✅ 가능 | ❌ 불가 |
| **원자성** | ✅ 지원 | ❌ 불가 |
| **Content-Type** | `application/json-patch+json` | `application/merge-patch+json` |

## Go 구현

### 라이브러리 선택

**`github.com/evanphx/json-patch`** 사용:
- JSON Patch + Merge Patch 모두 지원
- 1.5k+ Stars로 검증
- MIT 라이선스
- 활발한 유지보수

### 사용 예시

#### JSON Patch (RFC 6902)

```go
import (
    "github.com/evanphx/json-patch"
)

// JSON Patch 문서 생성
patch := jsonpatch.PatchOperations{
    {Op: "replace", Path: "/title", Value: "새 제목"},
    {Op: "add", Path: "/tags/-", Value: "updated"},
    {Op: "remove", Path: "/description"},
}

// 원본 JSON
original := []byte(`{"title": "기존 제목", "description": "내용"}`)

// 패치 적용
patched, err := patch.Apply(original)
if err != nil {
    // 에러 처리
}
```

#### JSON Merge Patch (RFC 7396)

```go
import (
    "encoding/json"
    "github.com/evanphx/json-patch"
)

// 원본 JSON
original := []byte(`{"title": "기존 제목", "description": "내용"}`)

// 머지 패치
patchData := map[string]interface{}{
    "title": "새 제목",
    "description": nil,  // 삭제
}
patchBytes, _ := json.Marshal(patchData)

// 패치 생성 및 적용
patch, _ := jsonpatch.DecodePatch(patchBytes)
patched, err := patch.Apply(original)
```

## API 설계

### Content-Type에 따른 분기

```http
# JSON Patch 사용
PATCH /api/v1/data/overview.json
Content-Type: application/json-patch+json

[
  {"op": "replace", "path": "/title", "value": "새 제목"},
  {"op": "add", "path": "/tags/-", "value": "updated"}
]

# JSON Merge Patch 사용
PATCH /api/v1/data/overview.json
Content-Type: application/merge-patch+json

{
  "title": "새 제목",
  "description": null
}
```

### 서버 처리 로직

```go
func (h *Handler) HandlePatch(w http.ResponseWriter, r *http.Request) {
    contentType := r.Header.Get("Content-Type")
    
    switch contentType {
    case "application/json-patch+json":
        // JSON Patch 처리
        handleJSONPatch(w, r)
    case "application/merge-patch+json":
        // JSON Merge Patch 처리
        handleMergePatch(w, r)
    default:
        // 기본: JSON Merge Patch로 처리
        handleMergePatch(w, r)
    }
}
```

## Downstream Concerns
- [ ] **에러 처리**: 잘못된 패치에 대한 에러 메시지
- [ ] **버전 관리**: ETag/If-Match로 동시성 제어
- [ ] **검증**: 패치 적용 전 유효성 검사
- [ ] **로깅**: 패치 변경 이력 기록

## Tags
`json-patch`, `rfc6902`, `rfc7396`, `partial-update`

## Approved
- 2026-07-21: JSON Patch + Merge Patch 모두 지원, json-patch 라이브러리 사용, 사용자 승인