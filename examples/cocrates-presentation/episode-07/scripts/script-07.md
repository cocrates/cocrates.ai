# Script: Slide 07 — 🏗️ Core Activity 1 — 산출물 생성 파이프라인

**Time:** ~1분

**[SLIDE 07: 파이프라인 흐름도 — ADR → Spec → Generation → Verification. 각 단계 아이콘]**

Cocrates의 첫 번째 핵심 활동은 **산출물 생성**입니다. 이 파이프라인은 네 단계로 구성됩니다: ADR, Spec, Generation, Verification.

**ADR**은 선택지 분석과 결정 단계입니다. "왜 이렇게 할까?"라는 질문에 답합니다. Ep5에서 jsondb 사례의 3개 ADR을 기억하시나요? 저장 모델, 아키텍처, 동시성 — 각각의 선택지를 분석하고 사용자가 결정을 내렸습니다.

**Spec**은 결정된 모든 요구사항을 하나로 통합합니다. "무엇을 만들까?"라는 질문에 답하는 자체 완결적 명세서입니다.

**Generation**은 승인된 Spec만을 근거로 최종 산출물을 생성합니다. Spec이 유일한 기준입니다.

**Verification**은 결과물이 Spec과 일치하는지 항목별로 검증합니다. Ep5의 72개 검증 항목, 기억나시나요?

이 파이프라인을 담당하는 스킬들은 `adr-writing`, `spec-writing`, `spec-driven-generation`, `spec-driven-verification`입니다.

→ [SLIDE 08: 🎓 Core Activity 2 — 학습 파이프라인]
