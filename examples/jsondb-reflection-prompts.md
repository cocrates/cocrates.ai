# jsondb Reflection Session Log

- **Date:** 2026-06-29
- **Project:** examples/jsondb
- **Purpose:** jsondb 프로젝트에 대한 이해도 평가 (Reflection)
- **Skill Used:** reflection → knowledge-capture

---

## 대화 흐름

### 1. Opening: 프로젝트 본질과 목적

| 역할 | 내용 |
|------|------|
| **AI** | jsondb가 해결하려는 핵심 문제, 파일 시스템 선택 이유, AI 맥락이 설계에 남긴 흔적에 대해 질문 |
| **User** | AI 생성 마크다운 파일(overview, outline, sections)의 파일 단위 비효율 → JSON 체계화 + 부분 검색/변경이 목적이라고 응답 |
| **판정** | ✅ Solid — 핵심 동기와 AI 맥락 정확히 인지 |

### 2. 저장소 모델 — 경로와 파일 시스템 매핑

| 역할 | 내용 |
|------|------|
| **AI** | `episode/e1` 경로의 디스크 매핑, 문서 간 간섭 여부, 매핑 방식의 장단점 질문 |
| **User** | `episode/e1.json` 저장, 파일 기반 lock으로 간섭 없음, 디렉토리 구조와 동일해 관리 용이. **단점(제약)은 모르겠다고 응답** |
| **AI** | 파일 시스템 제약(상대 경로, OS 문자 제한, 케이스 민감성)에 대한 후속 질문 |
| **User** | **`../../../config.json` Path Traversal 위험을 스스로 추론** — 상대 경로 지원하지 않도록 방어 필요 제안 |
| **판정** | ✅ Solid — 경로 매핑 정확. 단점 추론에서 Path Traversal 위험 발견은 고급 통찰 |

### 3. 동시성 모델 — Per-File RWMutex

| 역할 | 내용 |
|------|------|
| **AI** | 전역 lock 대신 파일별 lock 선택 이유, 단일 파일 lock 정책의 deadlock-free 이점, lock 미사용 함수 질문 |
| **User** | "파일별 lock이 성능 개선, 단일 파일만 lock → deadlock-free 가능", lock 미사용 함수는 **List()** 라고 정확히 응답 |
| **판정** | ✅ Solid — 동시성 설계 의도와 deadlock-free 원리 완전 이해 |

### 4. 스키마 시스템

| 역할 | 내용 |
|------|------|
| **AI** | 스키마 참조 방식, schemaless 처리, SetPath() 스키마 재검증 여부 질문 |
| **User** | `$schema`(실제는 `_schema`)로 혼동했지만 개념은 인지. Schemaless 문서는 자유 — 정확. SetPath() 검증은 "필요해 보임"이나 실제 구현은 모르며 "사양에 명시 필요" 제안 |
| **판정** | ⚠️ Partial — 개념은 알고 있으나 정확한 필드명(`_schema`)과 SetPath() 동작 차이는 미확인 Gap |

### 5. 아키텍처 — Library + REST Server

| 역할 | 내용 |
|------|------|
| **AI** | 두 방식을 모두 제공하는 이유, 각각의 사용 시나리오, Go 1.22 ServeMux 새 기능 질문 |
| **User** | "HTTP만 있어도 충분, CLI와 curl로 접근 가능. Go 1.22 기능은 모름" |
| **AI** | Go 라이브러리 임포트 시 이점(동일 프로세스, 직렬화/네트워크 오버헤드 없음)과 두 CLI 바이너리(cmd/jsondb + cmd/jsondbd) 분리 이유를 추가 질문 |
| **User** | 여기까지 진행 후 종료 요청 |
| **판정** | ⚠️ Partial — CLI 사용 시나리오는 이해했으나 Library 임포트의 장점과 두 바이너리 분리 이유는 검토되지 않음 |

---

## 이 대화를 통해 알게된 "설계 검토되고 사양으로 정리되었어야 하는 사항"

Reflection 과정에서 드러난, 사전에 **ADR 또는 Spec에서 명시적으로 검토/결정되지 않았거나 문서화되지 않은 설계 사항**들입니다.

### 1. Path Traversal 방어 정책 누락 ⚠️ (위험)

- **문제:** `a/../../../config.json` 같은 상대 경로를 `filePath()`가 root 경계 검증 없이 `filepath.Join`으로 처리 → root 바깥 파일 접근 가능
- **lockKey()도 동일한 문제** — `filepath.Clean`은 하지만 root로 시작하는지 검증하지 않음
- **필요한 결정:** 상대 경로 차단 여부, 경로 유효성 검증 정책, root 범위를 벗어나는 경로 처리 방식
- **영향:** 보안 취약점 (의도하지 않은 파일 읽기/쓰기)

### 2. 경로 형식 제약 사항 미문서화

- **문제:** OS별 파일명 제한 문자(Windows: `\ / : * ? " < > |` 등), 케이스 민감성 차이(macOS는 case-insensitive, Linux는 case-sensitive), 최대 경로 길이 등에 대한 방어나 명시가 없음
- **필요한 결정:** 지원하는 경로 문자 집합, 케이스 처리 정책, 잘못된 경로에 대한 에러 처리

### 3. SetPath() 스키마 재검증 정책 미정의

- **문제:** `Set()`은 스키마 검증을 하지만, `SetPath()`는 하지 않음. 이 차이가 의도된 설계인지, 일관성 없는 동작인지 Spec에 명시되지 않음
- **SetPath()로 `_schema` 필드를 직접 변경하면?** — 스키마 참조가 바뀌어도 검증 없이 통과
- **필요한 결정:** SetPath()에서도 스키마 재검증을 수행할지, 아니면 "저수준 우회"로 의도적으로 허용할지

### 4. List()의 일관성 보장 수준 미문서화

- **문제:** `List()`는 lock 없이 `filepath.Walk`를 사용. 동시 쓰기 중에 호출되면 불완전한 목록을 반환할 수 있으나, 이 "best-effort" 동작이 Spec에 명시되지 않음
- **필요한 결정:** List()의 일관성 보장 수준을 명시하고, 필요하다면 snapshot lock 도입 검토

### 5. CLI URL 구성의 Spec-구현 불일치

- **문제:** `verification/jsondb.md`에서 발견됨. Schema 경로가 `/`로 시작하면 URL에 `//`가 발생 (`/schema//blog/episode`). 기능적으로는 동작하지만 Spec에 문서화된 형식과 다름
- **필요한 결정:** URL 정규화 규칙 명시 또는 구현 수정

---

*이 문서는 Reflection 세션의 대화 흐름과 그 과정에서 발견된 설계 검토 필요 사항을 정리합니다.*
