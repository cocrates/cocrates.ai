# Slide 06: 🔀 Intent-To-Skill Routing — 의도가 스킬을 선택한다

## Key Message
Cocrates의 두뇌. 단순 키워드 매칭이 아닌 사용자의 근본 의도(Intent)를 추론하여 최적의 스킬을 매핑한다.

## Content
- **원리:** 단순 키워드 매칭 ✗ → 사용자의 근본 의도(Intent) 추론 ✓
- **주요 매핑 테이블:**
  | 의도 (Intent) | 스킬 (Skill) |
  |------|------|
  | 개념을 배우고 싶다 | `education` |
  | 배운 내용을 정리하고 싶다 | `knowledge-capture` |
  | 이해도를 확인받고 싶다 | `reflection` |
  | 선택지를 비교하고 싶다 | `adr-writing` |
  | 명세를 정의하고 싶다 | `spec-writing` |
  | 산출물을 생성하고 싶다 | `spec-driven-generation` |
  | 검증하고 싶다 | `spec-driven-verification` |
  | 스킬을 만들고 싶다 | `generating-skill-creation` |
- **라우팅 규칙:** 의도 추론 → 적합한 스킬 탐색 → 있으면 로드하여 처리, 없으면 기본 생성(spec-driven-generation)
- **포인트:** "알려줘" → 의도 추론 → `education` 스킬 자동 로드

## Visual Elements
- 상단: "키워드 매칭 ❌" vs "의도 추론 ✅" 대비
- 중앙: 매핑 테이블 (좌: Intent / 우: Skill)
- 하단: 라우팅 규칙 순서도 (간략)
- 강조: education, spec-driven-generation 행

## Layout Notes
- 테이블이 핵심 — 가독성 좋게 포맷팅
- 테이블 좌우 대비 (의도는 자연어, 스킬은 코드명)
- 발표 시 각 행을 하나씩 가리키며 설명 가능하도록 여백 확보

## Reference
- Script: `scripts/script-06.md`
- Slide plan: `slides.md` → Slide 06
