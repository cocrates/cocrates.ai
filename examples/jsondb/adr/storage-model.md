# Storage Model — Path-to-Filesystem Mapping Strategy

## Concern
Should jsondb enforce a fixed collection-document hierarchy, or should the user-specified path directly determine the filesystem path?

## Status
approved

## Context
- jsondb의 용도: AI가 생성한 중간 산출물(overview, outline, episodes, episode별 상세 등)을 구조화된 JSON으로 저장
- 이런 산출물은 계층 구조가 일정하지 않음 — `episode/e1`, `episode/e2/deep` 등 프로젝트마다 다양함
- 저장 경로는 인간과 AI가 직관적으로 이해할 수 있어야 함 (파일 브라우저에서 바로 찾을 수 있어야)
- Partial document access (JSON path 기반 read/update)는 별도 ADR에서 다룸
- 저장소가 collection-document 스키마를 강제하면, AI 산출물의 다양한 구조를 표현하기 어려움

## Decision
**Option A — Path-Addressable**. 사용자가 지정한 path가 곧 filesystem 경로가 되는 모델.
- `db.Set("overview", data)` → `{dbroot}/overview.json`
- `db.Set("episode/e1", data)` → `{dbroot}/episode/e1.json`
- 디렉토리는 path의 `/` 세그먼트 기준 자동 생성

## Options

### Option A — Path-Addressable (key = filesystem path)
- 사용자가 지정한 path가 그대로 파일 경로가 됨
- `db.Set("overview", data)` → `{dbroot}/overview.json`
- `db.Set("episode/e1", data)` → `{dbroot}/episode/e1.json`
- 경로의 디렉토리部分是 자동 생성
- Pro: **완전한 자유도** — 사용자가 원하는 구조를 그대로 반영
- Pro: 직관적 — 파일 시스템에서 바로 찾을 수 있음
- Con: 계층 구조의 일관성을 사용자에게 전적으로 위임
- Con: "컬렉션 목록 보기" 같은 연산은 filesystem walk 필요

### Option B — Collection-Document (강제 2단계 계층)
- `{dbroot}/{collection}/{document-id}.json` 형태 강제
- `db.Set("episodes", "e1", data)` 형태의 API
- Pro: 일관된 구조 — 모든 데이터가 collection 하위에 정리됨
- Pro: collection 단위 연산(목록, 전체삭제)이 자연스러움
- Con: 유연성 부족 — `episode/e1` 같은 자유로운 경로 표현 불가
- Con: 2단계를 넘는 깊이 표현이 어려움 (episode/e1/scene-3 같은 경우?)

### Option C — Hybrid (Path-Addressable + 별도 Collection 개념)
- 기본은 path-addressable이지만, path의 첫 번째 세그먼트를 "collection"으로 간주하는 메타정보 제공
- 예: `episode/e1` → collection=`episode`, id=`e1`
- `db.ListCollections()` → path 첫 세그먼트 기준으로 수집
- Pro: path 주소 지정의 자유로움 + collection 개념의 편의를 모두 제공
- Con: 경계 모호 — `episode/e1/scene/deep` 같은 깊은 경로에서 collection 정의가 불분명
- Con: 설계 복잡도 증가

## Tradeoffs
| | A (Path-Addressable) | B (Collection-Doc) | C (Hybrid) |
|---|----------------------|--------------------|------------|
| 구조 자유도 | ★★★★★ 완전 자유 | ★★ 2단계 고정 | ★★★★ 대부분 자유 |
| 직관성 (파일시스템) | ★★★★★ path=파일경로 | ★★★ collection/dir | ★★★★ 대부분 직관 |
| Collection 연산 | ★★ filesystem walk 필요 | ★★★★★ 네이티브 | ★★★★ 부분 지원 |
| 설계 단순함 | ★★★★★ 매우 단순 | ★★★★ 단순 | ★★★ 중간 |
| 인간이 구조 설계 | 필요 (path를 사용자가 결정) | 불필요 (틀 고정) | 약간 필요 |

## Recommendation (optional)
- **Option A — Path-Addressable** 을 추천합니다. jsondb의 핵심 동기는 "AI 산출물을 구조화해서 저장"하는 것이고, 산출물의 구조는 프로젝트마다, 심지어 산출물마다 다릅니다. 강제된 계층은 오히려 방해가 됩니다. path-addressable이면:
  - 블로그 시리즈: `overview.json`, `outline.json`, `episodes.json`, `episode/intro.json`, `episode/deep-dive.json`
  - 연구 노트: `notes.json`, `references.json`, `chapter/draft-1.json`
  - 등 완전한 자유도

## Consequences
- API: `db.Set(path, data)`, `db.Get(path)`, `db.Delete(path)` — path는 항상 string
- 디렉토리는 path의 `/` 세그먼트 기준으로 자동 생성
- 경로 구분자는 항상 `/` (OS에 따라 자동 변환)
- `db.List(prefix)` 같은 prefix 기반 조회로 collection-like 기능 제공 가능

## Related
- Path syntax ADR (partial document access within a JSON file)

## Tags
`storage`, `path-mapping`, `filesystem`

## Approved
- 2026-06-19: Option A (Path-Addressable), user confirmed
