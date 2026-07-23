# JSON 작업 방식

## Concern
JSON 객체/파일의 조회, 검색, 변경 기능을 어떤 방식으로 제공해야 하는가?

## Status
approved

## Context
- 파일 기반 데이터 모델 채택 (JSON 파일 단위)
- 복합 JSON 구조 지원 필요 (중첩 객체, 배열, 혼합 타입)
- 필드 레벨 접근이 핵심 기능
- 에이전트가 직관적으로 사용할 수 있어야 함

## Decision
**기본 경로 + JSONPath 혼합 지원 (Option C)** 사용자 승인: 단순 경로(기본)와 JSONPath(고급)를 모두 지원하며, 문법에 따라 자동 감지하여 처리.

## Options

### Option A — 단순 경로 기반 접근

**문법**:
```
 필드 경로: 점(.) 구분자
 배열 접근: 대괄호[] 인덱스
 중첩 접근: 조합
```

**예시**:
```bash
# 조회
$ jsondb get overview.json title
$ jsondb get overview.json nested.field
$ jsondb get overview.json items[0].name

# 변경
$ jsondb set overview.json title "새 제목"
$ jsondb set overview.json items[0].name "새 이름"

# 삭제
$ jsondb delete overview.json items[1]
```

**장점**:
- **단순성**: 배우기 쉬운 문법
- **직관성**: 파일 시스템 경로와 유사
- **구현 용이**: 파싱이 간단

**단점**:
- **제한적**: 복잡한 쿼리 불가
- **배열 조작**: 조건부 접근 어려움
- **검색 기능**: 단순 값 비교만 가능

---

### Option B — JSONPath 기반 접근

**문법** (RFC 9535 준수):
```
 필드 경로: $. 필드명
 배열 접근: $[] 인덱스 또는 필터
 조건 검색: [?(@.필드==값)]
```

**예시**:
```bash
# 조회
$ jsondb get overview.json "$.title"
$ jsondb get overview.json "$.nested.field"
$ jsondb get overview.json "$.items[0].name"

# 변경
$ jsondb set overview.json "$.title" "새 제목"
$ jsondb set overview.json "$.items[0].name" "새 이름"

# 검색
$ jsondb search overview.json "$.tags[?(@=='important')]"
$ jsondb search overview.json "$.items[?(@.active==true)]"
```

**장점**:
- **표준**: RFC 9535 표준 문법
- **강력**: 복잡한 쿼리 가능
- **유연**: 조건부 검색, 필터링 지원

**단점**:
- **복잡성**: 문법이 복잡함
- **학습 곡선**: 에이전트가 학습 필요
- **구현 부담**: 완전한 JSONPath 파서 구현 필요

---

### Option C — 혼합 접근 (단순 + 고급)

**문법**:
```
기본 접근: 점(.) 구분자 (단순한 경우)
고급 접근: JSONPath 문법 (복잡한 경우)
자동 선택: 문법에 따라 자동 감지
```

**예시**:
```bash
# 기본 접근 (단순)
$ jsondb get overview.json title
$ jsondb get overview.json nested.field

# 고급 접근 (복잡한 경우)
$ jsondb get overview.json "$.items[?(@.active==true)].name"

# 변경
$ jsondb set overview.json title "새 제목"
$ jsondb set overview.json "$.items[0].name" "새 이름"
```

**장점**:
- **유연성**: 단순한 경우와 복잡한 경우 모두 지원
- **점진적 도입**: 기본 기능 먼저, 고급 기능은 필요 시
- **에이전트 친화적**: 단순한 문법으로 시작 가능

**단점**:
- **복잡성**: 두 가지 문법 관리 필요
- **일관성**: 문법 혼합으로 인한 혼동 가능

---

### Option D — 함수형 접근

**문법**:
```
메서드 체이닝 방식
filter, map, reduce 등 함수형 연산
```

**예시**:
```bash
# 조회
$ jsondb query overview.json 'get("title")'
$ jsondb query overview.json 'get("items").filter(active==true).get("name")'

# 변경
$ jsondb query overview.json 'set("title", "새 제목")'
$ jsondb query overview.json 'get("items")[0].set("name", "새 이름")'
```

**장점**:
- **강력**: 복잡한 연산 가능
- **유연**: 다양한 조작 지원

**단점**:
- **복잡성**: 가장 높은 학습 곡선
- **비직관적**: 파일 시스템 접근과 다름
- **구현 부담**: 함수형 파서 구현 필요

## Tradeoffs

| 기준 | 단순 경로 | JSONPath | 혼합 | 함수형 |
|------|----------|----------|------|--------|
| **단순성** | ★★★★★ | ★★★☆☆ | ★★★★☆ | ★☆☆☆☆ |
| **직관성** | ★★★★★ | ★★★☆☆ | ★★★★☆ | ★★☆☆☆ |
| **강력함** | ★★☆☆☆ | ★★★★★ | ★★★★☆ | ★★★★★ |
| **구현 난이도** | ★★★★★ | ★★★☆☆ | ★★★☆☆ | ★☆☆☆☆ |
| **에이전트 친화적** | ★★★★★ | ★★★☆☆ | ★★★★☆ | ★★☆☆☆ |

## Recommendation
**혼합 접근 (Option C)** 추천 이유:

1. **단계적 도입**: 기본 기능(단순 경로) 먼저 구현
2. **확장성**: 필요 시 JSONPath 지원 추가
3. **에이전트 호환**: 단순한 문법으로 에이전트 통합 용이
4. **유연성**: 사용자 선택권 제공

**구현 순서**:
1. Phase 1: 단순 경로 기반 접근 (기본 기능)
2. Phase 2: JSONPath 지원 (고급 기능)

## JSONPath 연구 결과

### 에이전트 호환성 분석

| 항목 | 분석 | 결과 |
|------|------|------|
| **표준성** | RFC 9535 국제 표준 | ✅ 문제 없음 |
| **에이전트 호환성** | LangChain, AutoGPT 등에서 이미 사용 | ✅ 문제 없음 |
| **학습 곡선** | 기본 문법은 쉬움, 복잡한 쿼리는 학습 필요 | ⚠️ 일부 학습 필요 |
| **일관성** | 여러 에이전트 프레임워크에서 지원 | ✅ 문제 없음 |

**결론**: JSONPath는 에이전트가 이미 사용하는 표준이므로, 큰 문제 없음.

### Go 언어 JSONPath 구현 방법

**기존 라이브러리 사용 (추천)**:

| 라이브러리 | 특징 | Stars | 선택 추천 |
|-----------|------|-------|----------|
| `github.com/ohler55/ojg` | 고성능, JSONPath 완전 지원 | 1.5k+ | ⭐ 추천 |
| `github.com/PaesslerAG/jsonpath` | 간단한 API, 경량 | 500+ | 경량 시 |
| `github.com/yalp/jsonpath` | 기본 기능 지원 | 200+ | 기본만 필요 시 |
| `github.com/bhmj/jsonslice` | 고성능, 제로 카피 | 100+ | 성능 중시 시 |

**라이브러리 사용 예시** (`ojg` 기반):
```go
import (
    "github.com/ohler55/ojg"
    "github.com/ohler55/ojg/jp"
)

// JSONPath로 조회
obj, _ := oj.ParseString(jsonData)
result, _ := jp.Parse("$.items[?(@.active==true)]").GetFirst(obj)

// JSONPath로 변경
path, _ := jp.Parse("$.title")
path.Set(obj, "새 제목")
```

**커스텀 구현**: 복잡도가 높아서 비권장. 기존 라이브러리 사용 권장.

## Downstream Concerns
- [ ] **검색 기능**: 전체 파일 검색 vs 단일 파일 검색
- [ ] **배열 연산**: 추가, 삭제, 필터링 방식
- [ ] **타입 검증**: 잘못된 타입 접근 처리
- [ ] **에러 메시지**: 사용자 친화적 에러 메시지

## Tags
`json`, `query`, `jsonpath`, `api-design`

## Approved
- 2026-07-21: Option C (기본 경로 + JSONPath 혼합 지원), 사용자 승인