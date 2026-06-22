# JSON Path Syntax — Partial Document Access

## Concern
Within a JSON document, what path expression syntax should be used to specify which value to read, update, or delete?

## Status
approved

## Context
- jsondb must support partial document access: getting/setting/deleting a specific value within a JSON document without touching the rest
- Example document at `episode/e1.json`:
  ```json
  {
    "title": "Episode 1",
    "scenes": [
      {"name": "Intro", "duration": 120},
      {"name": "Main", "duration": 300}
    ],
    "metadata": {"author": "AI", "tags": ["go", "json"]}
  }
  ```
- `db.Get("episode/e1", "metadata.author")` should return `"AI"`
- `db.Set("episode/e1", "scenes.1.duration", 350)` should update only that field
- Must handle: object keys, array indices, nested paths, special characters in keys
- The path syntax should be simple enough for AI to generate reliably
- Go ecosystem: prefer battle-tested libraries over custom implementations

## Decision
**Option A — GJSON/SJSON Path Syntax**. `tidwall/gjson`(Read) + `tidwall/sjson`(Write) 사용. Dot notation 기반: `metadata.author`, `scenes.1.name`.

## Options

### Option A — GJSON/SJSON Path Syntax (tidwall/gjson + tidwall/sjson)
- Use existing Go libraries: `tidwall/gjson` for reads, `tidwall/sjson` for writes
- Dot-separated path: `metadata.author`, `scenes.1.name`
- Array index via dot or bracket: `scenes.1` or `scenes[1]`
- Wildcard: `scenes.#.name` → all scene names
- Modifier: `scenes.#(duration>200)#.name`
- Pro: **battle-tested**, 10k+ GitHub stars, extremely fast (single-pass parsing)
- Pro: AI-friendly — dot notation은 LLM이 생성하기 쉬움
- Pro: read/sjson 모두 동일한 path 문법
- Con: non-standard (JSONPath와 다름), 학습 필요
- Con: 쿼리 기능은 방대하지만 jsondb는 CRUD만 필요

### Option B — Standard JSONPath (RFC 9535)
- Use Go library like `oliveagle/jsonpath` or `PaesslerAG/jsonpath`
- Standard syntax: `$.metadata.author`, `$.scenes[0].name`
- Pro: **표준** (RFC 9535), JSONPath를 아는 사용자에게 친숙
- Pro: 복잡한 필터 표현식 지원
- Con: `$` prefix가 모든 path에 필요 → 사소하지만 번거로움
- Con: JSONPath 구현마다 동작 차이 있음
- Con: JSONPath 표준은 아직 Go 생태계에서 완전히 성숙하지 않음

### Option C — JMESPath
- Query language for JSON, used by AWS CLI, `jmespath/go-jmespath`
- Syntax: `metadata.author`, `scenes[0].name`, `scenes[?duration > `200`].name`
- Pro: 표현력이 뛰어남 (projection, filter, multi-select)
- Con: CRUD보다는 **쿼리**에 특화 — Set/Delete 용도로는 적합하지 않음
- Con: 상대적으로 무거움 (원래 용도가 complex query)

### Option D — Custom Minimal Path Notation
- 직접 구현한 단순 경로 문법 (예: 점으로 구분, `a.b[0].c`)
- Pro: 완전한 제어, 필요한 만큼만 구현
- Con: **직접 구현, 테스트, 유지보수 부담**
- Con: edge case (escape, unicode key 등) 처리 누락 위험

## Tradeoffs
| | A (GJSON) | B (JSONPath) | C (JMESPath) | D (Custom) |
|---|-----------|-------------|-------------|------------|
| Go 생태계 성숙도 | ★★★★★ 최고 | ★★★ | ★★★ | N/A |
| CRUD 적합성 | ★★★★★ Read+Write 모두 지원 | ★★★ Read 위주 | ★★ Read 전용 | ★★★★ 맞춤 가능 |
| AI 생성 용이성 | ★★★★★ dot natation | ★★★★ $ 필요 | ★★★★ | ★★★★★ |
| 유지보수 부담 | ★★★★★ 최소 (라이브러리) | ★★★★ | ★★★★ | ★ (직접 개발) |
| 표현력 | ★★★★ (wildcard, filter) | ★★★★★ | ★★★★★ | ★★ (심플) |

## Recommendation (optional)
- **Option A — GJSON/SJSON** 을 강력히 추천합니다.
  - 이유 1: Read(`gjson`)와 Write(`sjson`)가 동일한 path 문법으로 완벽히 대응
  - 이유 2: dot notation은 AI가 생성하기 가장 쉬운 형식
  - 이유 3: Go 생태계에서 가장 성숙하고 빠른 JSON path 라이브러리
  - 이유 4: jsondb는 복잡한 쿼리가 아닌 CRUD가 목적 — GJSON의 기능으로 충분

## Consequences
- Import: `github.com/tidwall/gjson`, `github.com/tidwall/sjson`
- Path 예시: `metadata.author`, `scenes.1.name`, `scenes.1.duration`
- Get/Set/Delete 모두 동일한 path 문법 사용
- 유효하지 않은 path는 error 반환

## Related
- storage-model ADR (파일 수준 path 결정)

## Tags
`json-path`, `partial-access`, `gjson`, `sjson`

## Approved
- 2026-06-19: Option A (GJSON/SJSON), user confirmed
