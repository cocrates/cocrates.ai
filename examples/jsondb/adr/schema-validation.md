# Schema Validation — Definition Format and Data Association

## Concern
How should jsondb define schemas and associate them with data files, to prevent AI from writing arbitrary or malformed data?

## Status
approved

## Context
- jsondb의 핵심 동기: AI가 생성한 데이터를 구조화해서 저장하는 것
- "스키마 없이 AI가 멋대로 사용하면, 제대로 활용하기 어려울 것" (user quote)
- Schema는 AI가 생성할 수 있어야 함 — 사람이 직접 정의하는 것보다 AI가 함께 생성하는 경우가 많음
- 검증 시점: `db.Set()` 호출 시 (write-time validation)
- Schema는 선택적(optional)이어야 함
- 사용자 제안: `{schemaroot}` 폴더에 스키마를 정의하고, 데이터 파일에서 스키마 경로를 명시

## Decision
**Option A — Explicit Schema Reference**.
- 스키마 정의: JSON Schema, `{schemaroot}` 디렉토리 아래 저장
- 데이터 파일의 `_schema` 필드가 `{schemaroot}` 기준 경로로 스키마 참조
- `_schema`가 **없는 문서는 schemaless** — 검증 없이 자유롭게 저장
- `_schema`가 있는 문서만 해당 스키마로 검증

## Options

### Option A — Explicit Schema Reference (사용자 제안)
- **스키마 정의**: 표준 JSON Schema. `{schemaroot}` 디렉토리에 저장 (예: `{dbroot}/schema/`)
- **스키마 참조**: 데이터 파일의 특수 필드(`_schema`)가 스키마 경로를 명시
  ```
  // {dbroot}/schema/blog/episode.json  ← 스키마 파일
  // {dbroot}/episode/e1.json           ← 데이터 파일
  {
    "_schema": "/blog/episode.json",
    "title": "Episode 1",
    "scenes": [...]
  }
  ```
- `_schema` 경로는 `{schemaroot}` 기준으로 해석 (위 예시에서 `/blog/episode.json` → `{schemaroot}/blog/episode.json`)
- Pro: **하나의 스키마로 여러 문서 검증** 가능 — `e1.json`, `e2.json`이 같은 `/blog/episode.json` 참조
- Pro: **직관적** — 파일만 봐도 어떤 스키마를 따라야 하는지 알 수 있음
- Pro: `{schemaroot}`가 명시적 — 다른 프로젝트의 스키마 참조도 가능
- Pro: AI가 데이터 생성 시 `_schema` 필드를 포함하도록 하기 쉬움
- Con: 모든 데이터 파일에 `_schema` 필드가 포함되어야 함 (귀찮을 수 있음)
- Con: `_schema` 필드는 jsondb의 특수 필드 — 데이터와 구분 필요

### Option B — Path Mirror Convention
- 스키마가 데이터 path를 미러링한 위치에 자동 매핑
  - 데이터: `{dbroot}/episode/e1.json`
  - 스키마: `{dbroot}/_schema/episode/e1.json`
- 데이터 파일에 별도 명시 불필요 — path 기반으로 자동 연결
- Pro: 데이터 파일에 _schema 필드가 없어도 됨
- Con: 문서 하나당 스키마 하나 — 같은 구조의 문서(e1, e2)가 같은 스키마를 공유하려면 중복 or symlink 필요
- Con: path가 바뀌면 스키마 연결도 바뀜

### Option C — Go Struct Tags
- Go 구조체에 `json` + `validate` 태그로 스키마 정의
- `go-playground/validator` 라이브러리 사용
- Pro: 타입 안전, 컴파일 타임 검증
- Con: **런타임에 스키마 변경 불가** — AI 생성 데이터에 부적합
- Con: 새 데이터 타입마다 Go 코드 수정 필요

## Tradeoffs
| | A (Explicit Reference) | B (Path Mirror) | C (Go Struct) |
|---|------------------------|-----------------|---------------|
| 스키마 공유 | ★★★★★ 다수 문서 공유 용이 | ★★ 문서당 하나 | ★★★★ 타입으로 공유 |
| 데이터만 봐도 스키마 식별 | ★★★★★ _schema 필드 | ★ path 추적 필요 | ★ 해당 없음 |
| AI 생성 용이성 | ★★★★★ `_schema` 포함 | ★★★★★ 자동 매핑 | ★ Go 코드 필요 |
| 간결성 (데이터 파일) | ★★★ _schema 필드 추가 | ★★★★★ 추가 필드 없음 | ★★★★★ |
| 유연성 | ★★★★★ 런타임 변경 자유 | ★★★★★ | ★ 컴파일 타임 고정 |

## Recommendation (optional)
- **Option A — Explicit Schema Reference** 를 추천합니다.
  - 사용자 제안과 정확히 일치
  - 가장 큰 장점: **하나의 스키마로 여러 문서 검증** — 블로그 시리즈의 모든 에피소드가 동일한 `episode.json` 스키마를 공유
  - `_schema` 필드는 선택적 — 없는 문서는 검증 없이 저장
  - JSON Schema 포맷을 채택하므로 AI 생성에 최적

## Consequences
- `{schemaroot}` 기본값: `{dbroot}/schema/` (configurable)
- 데이터 파일의 `_schema` 값은 `{schemaroot}` 기준 상대 경로 (예: `/blog/episode.json`)
- `db.SetSchema(path, schema)` → `{schemaroot}/{path}.json` 에 저장
- 검증: `db.Set(dataPath, data)` 호출 시 data의 `_schema` 필드를 읽어 schema 로드 → 검증
- `_schema` 필드는 jsondb가 내부적으로 사용 — `db.Get()` 결과에서 제외할지 여부는 API 설계에서 결정

## Related
- storage-model ADR (path-addressable storage)

## Tags
`schema`, `validation`, `json-schema`, `data-integrity`

## Approved
- 2026-06-19: Option A (Explicit Schema Reference), user confirmed — `_schema`가 없으면 schemaless로 자유 저장
