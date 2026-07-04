# Presentation Definition — Episode 09

## Topic
소크라테스식 학습 활동 - 스킬: 3블록 미션과 무지(Ignorance)를 박제하는 코드 엔진

## Purpose
Ep8에서 Learning 파이프라인의 3대 철학(소크라테스 산파술, 블룸 분류학, ZPD)과 3단계 구조(Education → Knowledge Capture → Reflection)를 이해했다. Ep9에서는 이 원리가 실제 `.opencode/skills/` 안에서 **구체적인 코드 구조와 워크플로우 규칙**으로 어떻게 구현되어 있는지 내부를 들여다본다. Ep4의 실습 경험과 Ep8의 원리가 실제 스킬 파일에서 어떻게 연결되는지 확인한다.

## Audience
- Ep1~Ep8을 시청한 학습자 (Learning 파이프라인 원리 이해자)
- Ep4의 실습을 경험하고 "그때 Cocrates가 왜 그렇게 행동했는지" 스킬 내부 구조가 궁금한 사람
- `.opencode/skills/` 구조에 관심이 있는 개발자

## Logistics
- **Time limit:** ~13분 (13슬라이드)
- **Format:** 온라인 교육 영상 (시리즈의 일부)
- **Tone/style:** 개념적 원리(Ep8)에서 한 단계 더 내려가 구체적인 구현체를 살펴보는 탐험 느낌. 코드 블록 예시를 보여주되, 실제로 다 읽을 필요는 없다는 점을 안내. 스킬 파일의 "행동 규칙"에 초점.

## Success Criteria
시청자가 다음을 설명할 수 있어야 함:
1. Education 스킬의 No Spoon-feeding 원칙, Turn-based Mission 구조, 3블록 출력 포맷(Concept Briefing → Thought Lab → MISSION)
2. Knowledge Capture 스킬의 Ignorance 기록 철학과 Merge 전략
3. Reflection 스킬의 Interviewer 페르소나와 Gap 발견 시 행동 원칙 (가르치지 않는다)

## Constraints & Taboos
- 각 스킬의 전체 SKILL.md 전문을 다 읽지 않음 — 핵심 워크플로우와 규칙만 발췌
- Ep8의 3대 철학과 3단계 구조를 이미 알고 있다는 전제하에 진행 (반복 설명 금지)
- 코드 블록을 지나치게 자세히 설명하지 않음 — "이런 규칙이 있다"는 개념적 이해에 초점
- Education 스킬의 Bloom's 2D Matrix나 Push & Pull은 소개하되, Ep8의 내용과의 연결선에서만 다룸

## External References
- `examples/blog/episodes/09-socratic-learning-skills.md`
- `examples/presentation/episode-08/` — Ep8 원리 편 참조
- `examples/presentation/episode-04/` — Ep4 실습 참조 (교육 스킬의 실제 작동 경험)
- `.opencode/skills/education/SKILL.md`
- `.opencode/skills/knowledge-capture/SKILL.md`
- `.opencode/skills/reflection/SKILL.md`

## Estimated Structure
- **도입 (2슬라이드, ~2분):** Ep8 연결 + 학습 목표 (스킬 내부 탐험)
- **본론 (9슬라이드, ~9분):** Education 스킬 내부(No Spoon-feeding + Turn-based Mission → Bloom's 2D + Push&Pull → 3블록 출력 포맷 + DIP 실제 사례) → Knowledge Capture 스킬 내부(Ignorance 기록, Merge) → Reflection 스킬 내부(면접관, Gap 처리) → 세 스킬의 유기적 연결
- **결론 (2슬라이드, ~2분):** 핵심 요약 + Ep10 연결
