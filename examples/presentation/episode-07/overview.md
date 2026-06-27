# Presentation Definition — Episode 07

## Topic
Cocrates Harness 구조 — 하나의 거대한 프롬프트가 품을 수 없는 'AI 주방'의 비밀

## Purpose
지금까지 사용자가 Cocrates를 **사용**하는 법을 배웠다면, 이제 Cocrates가 **내부적으로 어떻게 돌아가는지** 구조를 이해한다. Agent + Skills 아키텍처가 왜 필요한지, Cocrates Agent Prompt가 어떤 6개 섹션으로 구성되어 있는지, Intent-To-Skill Routing이 어떻게 작동하는지를 설명한다. Part 3(구조 편)의 첫 번째 에피소드로서, 이후 Ep8~12에서 다룰 개별 스킬 원리의 기반을 제공한다.

## Audience
- Ep1~Ep6을 시청한 학습자 (Cocrates 사용 경험자)
- Cocrates가 "어떻게" 작동하는지 궁금한 사람
- AI 시스템 아키텍처에 관심이 있는 개발자

## Logistics
- **Time limit:** ~15분 (10~12슬라이드)
- **Format:** 온라인 교육 영상 (시리즈의 일부)
- **Tone/style:** 전문적이나 비유(주방/만능 칼)로 접근성을 높임. Part 3의 첫 편이므로 개념 설명에 집중.

## Success Criteria
시청자가 다음을 설명할 수 있어야 함:
1. Cocrates가 Agent + Skills 구조를 채택한 이유
2. Cocrates Agent Prompt의 6개 섹션과 각 역할
3. Intent-To-Skill Routing이 의도를 기반으로 어떻게 스킬을 선택하는지

## Constraints & Taboos
- 개별 스킬의 상세 워크플로우는 다루지 않음 (Ep8~12에서 다룸)
- 설치 방법은 다루지 않음 (Ep3에서 다룸)
- 실습 시나리오는 다루지 않음 (Ep4~6에서 다룸)
- Cocrates Agent Prompt의 모든 라인을 읽지 않음 — 핵심만 구조적으로 전달

## External References
- `cocrates.md` (Cocrates Agent Prompt 원본)
- `examples/blog/episodes/07-cocrates-harness-structure.md`
- `examples/presentation/docs/ai-native-engineer.md`
- Ep6 Slide 14 (다음 에피소드 연결 참조)

## Estimated Structure
- **도입 (2슬라이드, ~3분):** Ep6 연결 + 학습 목표 + "만능 칼의 함정"
- **본론 (6~8슬라이드, ~10분):** Agent + Skills 필요성 → Cocrates Agent 6대 섹션 → Intent-To-Skill Routing → Core Activities
- **결론 (2슬라이드, ~2분):** 핵심 요약 + Ep8 연결
