# Presentation Definition — Episode 11

## Topic
구조 기반 산출물 생성 활동 - 스킬: 4대 전담 스킬의 설계 원칙과 'spec이 없으면 spec을 만들게 한다'는 핵심 동작

## Purpose
사용자가 "A를 생성해줘"라고 요청하면, Cocrates는 먼저 **A 생성에 특화된 스킬이 있는지 확인**한다. 있다면 해당 스킬(이미 구조가 내장됨)로 생성한다. **없다면 `spec-driven-generation` 스킬로 폴백(fallback)**한다.

`spec-driven-generation`은 **spec이 명확한지 먼저 확인**한다. 명확하지 않으면, 구조를 검토하여 spec을 작성하도록 유도한다 — 이것이 핵심이다. 즉, ASR을 식별하고 → ADR을 통해 사용자가 결정하게 하고 → 결정된 내용을 Spec에 포함시킨다. Spec이 완성된 후에야 비로소 Spec을 근거로 산출물을 생성한다.

이 과정에서 협력하는 **4개 핵심 스킬**(adr-writing, spec-writing, spec-driven-generation, spec-driven-verification) 각각의 핵심 원칙, 산출물, 절대 금지 사항(Constraints)을 이해하고, 이 스킬들이 어떻게 협력하여 '검증 가능한 산출물'을 생성하는지 파악한다.

## Audience
- **Background:** 소프트웨어 개발자(초급~중급), AI 도구를 업무에 활용하는 지식 노동자
- **Knowledge level:** Ep1~10을 통해 Cocrates 구조와 Generation 파이프라인 원리를 학습한 상태
- **Interests/expectations:** 원리를 넘어 실제 스킬의 구체적인 '어떻게(How)'를 알고 싶어 함

## Logistics
- **Time limit:** 18분 (14슬라이드, ~1.3분/슬라이드)
- **Format:** 시리즈 내 독립 발표 세션 (in-person)
- **Tone/style:** Part 3(구조) — 개념적 깊이를 유지하되 비유로 추상성 낮춤

## Success Criteria
1. 청중이 **"스킬이 없으면 spec-driven-generation이 폴백된다"**는 Cocrates의 기본 동작을 설명할 수 있다
2. 청중이 `spec-driven-generation`이 **spec이 없을 때 spec 작성을 유도하는 과정**(ASR 식별 → ADR → Spec)을 설명할 수 있다
3. 청중이 4개 스킬 각각의 **핵심 원칙과 절대 금지 사항(Constraints)**을 말할 수 있다

## Constraints & Taboos
- Ep10 원리를 이미 학습했다고 가정 — ASR/ADR/Spec/Generation 기본 개념 재설명 금지
- 지나치게 스킬 내부 코드 구현까지 들어가지 않음 (원칙과 메커니즘에 초점)
- 각 스킬의 SKILL.md 파일 경로는 언급하되 내용 전체를 옮겨 싣지 않음

## External References
- `examples/blog/episodes/11-architecture-driven-generation-skills.md` (블로그 원문)
- `.opencode/skills/adr-writing/SKILL.md`
- `.opencode/skills/spec-writing/SKILL.md`
- `.opencode/skills/spec-driven-generation/SKILL.md`
- `.opencode/skills/spec-driven-verification/SKILL.md`
- `examples/presentation/episode-10/` (Ep10 발표자료 — Generation 파이프라인 원리)

## Ep10과의 관계
| 구분 | Ep10 (원리) | Ep11 (스킬) |
|------|-------------|-------------|
| 초점 | Why — 왜 구조가 필요한가 | How — 어떻게 구현하는가 |
| 대상 | 파이프라인 개념 | 파이프라인을 구동하는 4개 스킬 + **fallback 메커니즘** |
| 핵심 통찰 | "그냥 써줘"는 위험하다 | **스킬이 없으면`spec-driven-generation`이 spec 작성을 유도한다** |
| 비유 | 전원주택 건축, 여행 계획 | 전담 부대(Specialist Teams) + 긴급 대응팀(fallback) |
