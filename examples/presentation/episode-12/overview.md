# Presentation Definition — Episode 12

## Topic
구조 기반 워크플로우 생성 스킬 — generating-skill-creation: **스킬을 만드는 메타-스킬**

## Purpose
지금까지 우리는 두 가지 생성 방식을 배웠다:
1. **스킬이 있는 경우** → 해당 스킬의 내장된 구조(Snowflake)로 생성
2. **스킬이 없는 경우** → `spec-driven-generation`이 구조를 만들어 생성

그런데 "스킬 자체가 없으면 어떻게 만드는가?" 이 질문에 답하는 것이 **generating-skill-creation**이다.
이 에피소드에서는 generating-skill-creation이 **어떻게 산출물을 5대 구성 요소로 분해하고, Snowflake 5단계로 점진 구체화하며, Meta Snowflake 7단계로 저작하여 사용자 승인을 거쳐 SKILL.md를 생성하는지** 그 메커니즘을 이해한다.

## Audience
- **Background:** 소프트웨어 개발자(초급~중급), AI 도구를 업무에 활용하는 지식 노동자
- **Knowledge level:** Ep1~11을 통해 Cocrates 구조, Learning 파이프라인, Artifact Generation 파이프라인과 4개 스킬까지 학습한 상태
- **Interests/expectations:** "구조적 접근을 내 업무에 맞는 스킬로 직접 만들어보고 싶다" — 메타 인지 욕구가 있는 청중

## Logistics
- **Time limit:** 16분 (11슬라이드, ~1.5분/슬라이드)
- **Format:** 시리즈 내 독립 발표 세션 (in-person)
- **Tone/style:** Part 3(구조) — 개념적 깊이를 유지하되 비유로 추상성 낮춤. "요리법 vs 요리" 비유 활용

## Success Criteria
1. 청중이 **"스킬이 없을 때 만드는 스킬(generating-skill-creation)"** 의 정체성을 설명할 수 있다
2. 청중이 **산출물 5대 구성 요소(Assignment, Context, Entities, Space, Structure)** 를 말하고 각 역할을 설명할 수 있다
3. 청중이 **Meta Snowflake 7단계(Kernel → Components → Frame → Outline → Spec → Skill → Verification)** 의 흐름을 설명할 수 있다
4. 청중이 **생성 전 마지막 게이트(detail design 완료 후 generation)** 의 중요성을 이해한다

## Constraints & Taboos
- Ep6(스킬 생성 실습)의 구체적 대화 내용 재설명 금지 — 실습은 경험했다고 가정
- Ep10~11(Generation 파이프라인 원리+스킬) 개념을 이미 학습했다고 가정 — ASR/ADR/Spec 기본 개념 재설명 금지
- generating-skill-creation의 SKILL.md 파일 경로는 언급하되 전체 내용을 옮겨 싣지 않음
- 스킬 저작 절차의 '어떻게(How)'에 초점 — SKILL.md의 모든 상세 규칙까지 다루지 않음

## External References
- `examples/blog/episodes/12-workflow-creation-skill.md` (블로그 원문)
- `.opencode/skills/generating-skill-creation/SKILL.md`
- `examples/presentation/episode-06/` (Ep6 발표자료 — 스킬 생성 실습)
- `examples/presentation/episode-10/` (Ep10 발표자료 — Generation 파이프라인 원리)
- `examples/presentation/episode-11/` (Ep11 발표자료 — 4개 스킬)

## Ep11과의 관계
| 구분 | Ep11 (4개 스킬) | Ep12 (메타-스킬) |
|------|----------------|------------------|
| 초점 | How — 기존 스킬의 메커니즘 | Meta-How — **스킬을 만드는 방법** |
| 대상 | adr-writing, spec-writing, spec-driven-generation, spec-driven-verification | **generating-skill-creation** (단일 스킬) |
| 핵심 통찰 | 스킬이 없으면 spec-driven-generation이 fallback된다 | **스킬이 필요하면 generating-skill-creation이 스킬을 만든다** |
| 비유 | 전문 부대 + 응급실 | **요리 도구를 만드는 대장간** |
