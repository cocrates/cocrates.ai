# Presentation Definition — Episode 10

## Topic
구조 기반 산출물 생성 활동 - 원리: "그냥 작성해줘"가 낳는 블랙박스와의 결별

## Purpose
Ep7에서 Cocrates의 두 가지 Core Activity를 확인했다. Ep9까지 첫 번째 Core Activity(Learning)를 원리와 스킬 내부까지 완전 정복했다. Ep10에서는 두 번째 Core Activity인 **Artifact Generation 파이프라인의 원리**로 딥다이브한다. "보고서 한 장 작성해줘"라는 요청에 Cocrates가 왜 생성 버튼을 바로 누르지 않고, 구조 설계(ASR 식별 → ADR → Spec)부터 요구하는지 그 이유를 이해한다. 전원주택 건축 비유로 추상적인 개념을 직관적으로 전달한다.

## Audience
- Ep1~Ep9을 시청한 학습자 (Learning 파이프라인 완전 정복자)
- AI에게 "그냥 작성해줘"라고 시켰다가 엉뚱한 결과를 받아본 경험이 있는 사람
- 구조적 사고와 아키텍처 설계에 관심이 있는 사람

## Logistics
- **Time limit:** ~15분 (13~14슬라이드)
- **Format:** 온라인 교육 영상 (시리즈의 일부)
- **Tone/style:** Part 3의 공통 원칙 — 개념적 깊이를 유지하되 전원주택 건축 비유로 추상성을 낮춤. "AI가 생성해주는 시대에 왜 내가 구조를 설계해야 하지?"라는 청중의 저항감을 인정하고, 그 이유를 납득시키는 대화체.

## Success Criteria
시청자가 다음을 설명할 수 있어야 함:
1. "그냥 작성해줘"가 낳는 세 가지 문제 (논리 붕괴, 설명 책임 상실, 블랙박스화)
2. ASR(Architecturally Significant Requirement)의 개념과 침묵의 기본값(Silent Default)의 위험성
3. 구조 기반 생성 4단계 파이프라인: Concern/ASR 식별 → ADR → Spec → Generation & Verification

## Constraints & Taboos
- 각 스킬의 내부 워크플로우 상세는 다루지 않음 (Ep11에서 다룸)
- Learning 파이프라인의 내용과 혼동하지 않도록 명확히 구분
- 전원주택 비유를 일관되게 유지 (비유가 중간에 바뀌지 않도록)
- ADR과 Spec의 구체적인 템플릿/포맷은 Ep11에서 다루므로 여기서는 개념적 이해에 집중

## External References
- `examples/blog/episodes/10-architecture-driven-generation-activity.md`
- `examples/presentation/episode-07/` — Ep7 전체 구조 (Core Activity 2개)
- `examples/presentation/episode-09/` — Ep9 Learning 파이프라인 완료 확인

## Estimated Structure
- **도입 (2슬라이드, ~2분):** Ep9 연결 + 학습 목표 (Learning에서 Generation으로 전환)
- **본론 (9~10슬라이드, ~11분):** "그냥 작성해줘"의 문제 → ASR 개념 → 4단계 파이프라인 (Concern/ASR → ADR → Spec → Generation & Verification)
- **결론 (2슬라이드, ~2분):** 핵심 요약 + Ep11 연결
