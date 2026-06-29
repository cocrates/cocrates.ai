# jsondb

## One-line Definition
AI 생성 intermediate artifact(overview, outline, sections 등)를 JSON으로 체계화해 파일 단위가 아닌 부분 검색/변경이 가능하도록 하는 Go 기반 로컬 파일 시스템 JSON Document Store.

## Key Points
- **저장 모델:** 논리적 경로(`episode/e1`) → `{root}/episode/e1.json` 파일. 디렉토리 구조가 곧 문서 구조.
- **동시성 모델:** Per-file RWMutex. 서로 다른 파일은 완전 동시 접근 가능. 모든 public 메서드는 정확히 **하나의 파일**만 lock → deadlock 원천 불가능.
- **Lock-free 함수:** `List()`만 유일하게 lock 미사용. `filepath.Walk` 기반 best-effort.
- **스키마:** JSON Schema draft-2020-12. 문서 내 `_schema` 필드로 참조. `_schema` 없으면 schemaless(검증 생략).
- **SetPath 스키마 검증:** `SetPath()`는 스키마 재검증을 하지 않음. (Set()만 검증)
- **아키텍처:** Library + REST Server(Go 1.22 ServeMux) + CLI(jsondb/jsondbd). Go 프로세스 내 임포트 또는 HTTP 접근 모두 지원.
- **JSON Path:** `tidwall/gjson`(읽기), `tidwall/sjson`(쓰기/삭제). Dot notation (`metadata.author`, `scenes.1.name`).

## Wrong Assumptions / Gaps
- **스키마 참조 필드명:** `$schema`가 아니라 `_schema`. (개념은 맞았으나 정확한 이름 혼동)
- **SetPath 검증:** "검증이 필요할 것"이라고 추론했지만 실제로는 하지 않음. 사양에 명시 필요.
- **Library + HTTP 서버 이중 구조:** "HTTP만 있어도 충분"하다고 생각했으나, Go 라이브러리 임포트 시 동일 프로세스 내 직렬화/네트워크 오버헤드 없이 사용 가능한 이점이 있음.
- **Path Traversal:** `../../../config.json` 같은 상대 경로로 root 바깥 접근 가능성. `filepath.Join`이 Clean은 하지만 root 범위 검증은 없음.

## Related
- bloom-taxonomy

## Tags
`jsondb`, `go`, `json-document-store`, `file-based-db`, `ai-artifacts`

## Update History
- 2026-06-29: Reflection 평가 기반 최초 KB 작성. 핵심 개념, 이해한 부분, Gap 기록.
