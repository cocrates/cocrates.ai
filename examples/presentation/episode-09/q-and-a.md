# Expected Questions & Answer Plans — Episode 09

## Q1: "Education 스킬의 No Spoon-feeding은 이해하는데, 사용자가 정말 기본 개념이 전혀 없으면 어떻게 하나요? 1~3문장만으로는 설명이 부족할 것 같은데요."
- **Context:** Slide 3 시청 중. No Spoon-feeding 규칙이 너무 극단적이라고 느낀 청중.
- **Answer strategy:**
  - No Spoon-feeding이 "절대 설명하지 않는다"는 뜻이 아니라 "완전한 정답을 한 번에 주지 않는다"는 뜻임을 명확히 함.
  - 1~3문장의 Briefing은 맥락을 잡아주는 역할. 이후 Thought Lab과 MISSION을 통해 사용자가 스스로 채우도록 설계됨.
  - Cocrates는 사용자의 응답을 보고 추가 비계가 필요하면 Push 전략으로 전환할 수 있음. "모르면 방치한다"는 의미가 아님.
  - 예: DIP 예시에서 Briefing만으로 이해가 안 되면, "플러그와 소켓" 비유를 더 풀어서 설명하는 Push로 전환 가능.
- **Reference:** slides.md Slide 3 — Push/Pull 전략, Incomplete State 원칙.

## Q2: "Knowledge Capture의 `Wrong Assumptions / Gaps` 섹션 — 이건 사용자가 직접 작성하는 건가요, Cocrates가 작성하는 건가요?"
- **Context:** Slide 7 시청 중. Ignorance 기록의 실제 워크플로우가 궁금한 청중.
- **Answer strategy:**
  - Knowledge Capture 스킬은 사용자가 "정리해줘"라고 요청했을 때 작동함. Cocrates가 Education 세션에서 관찰한 사용자의 오개념과 깨달음을 바탕으로 초안을 생성하고, 사용자가 확인 및 승인하는 방식.
  - 즉, Cocrates가 "당신이 이전에 블룸 분류학을 순차적 피라미드로 오해했었고, 지금은 2차원 매트릭스로 이해했습니다"라고 제안하면, 사용자가 "맞아, 그랬어" 또는 "아니, 사실 이렇게 이해했어"라고 수정하는 협업 형태.
  - 핵심은 **사용자가 실제로 이해하고 동의한 내용만** KB에 저장된다는 점. AI가 임의로 생성한 내용은 저장되지 않음.
- **Reference:** slides.md Slide 7, docs/episode-09.md — Knowledge Capture 저장 조건.

## Q3: "Reflection 스킬이 Pull-only라고 했는데, 만약 사용자가 모든 질문에 대답을 못하면 어떻게 되나요? 아무것도 얻는 게 없지 않나요?"
- **Context:** Slide 10 시청 중. Reflection의 실용성에 대한 의문.
- **Answer strategy:**
  - Reflection의 목표는 "얼마나 아는지 증명하는 것"이 아니라 **"어디까지 알고, 어디서부터 모르는지의 경계를 찾는 것"** 입니다.
  - 사용자가 모든 질문에 답하지 못하더라도, 그 '못 답한 지점'이 정확히 어디인지가 명확해집니다. 이것이 Reflection의 가장 가치 있는 산출물입니다.
  - 못 답한 부분은 `⚠️ 제대로 알지 못했던 것` 리스트에 기록되고, 이후 "이 부분에 대해 Education을 따로 진행하시겠습니까?"라는 제안으로 이어집니다. 즉, 못 답한 것이 손실이 아니라 **다음 학습의 출발점**이 되는 겁니다.
- **Reference:** scripts/script-10.md — Gap 발견 시 행동 원칙.

## Q4: "Ep8과 Ep9의 차이가 뭔가요? Ep8에서 이미 3단계 파이프라인을 배웠는데, Ep9는 그걸 다시 스킬 레벨에서 설명한 것 같아서요."
- **Context:** 발표 후반. 두 에피소드의 차이를 명확히 구분하지 못한 청중.
- **Answer strategy:**
  - **Ep8**은 "왜(Why)" — 3대 교육 철학과 파이프라인의 개념적 원리.
  - **Ep9**는 "어떻게(How)" — 그 원리가 실제로 `.opencode/skills/` 디렉토리 안에서 **어떤 코드와 규칙**으로 명시되어 있는지.
  - 비유: Ep8이 건축의 '설계도'라면, Ep9는 그 설계도가 실제 '시공 매뉴얼'로 어떻게 적혀 있는지를 보는 것.
  - Ep9의 핵심 가치는 "Cocrates가 그렇게 행동한 이유가 **코드로 강제된 규칙** 때문임을 알게 되는 것". 이는 사용자가 Cocrates의 행동을 예측하고 더 효과적으로 활용하는 데 도움을 줌.
- **Reference:** overview.md — Ep8 vs Ep9 Purpose 비교.

## Q5: "세 스킬 중에서 가장 중요한 스킬은 무엇인가요?"
- **Context:** 발표 마무리 단계. 전체 우선순위를 알고 싶은 청중.
- **Answer strategy:**
  - 셋 다 중요하지만, 굳이 꼽자면 **Education**이 가장 기본입니다. Knowledge Capture와 Reflection이 아무리 훌륭해도, Education을 통해 배울 내용이 없으면 저장하고 검증할 지식이 없기 때문입니다.
  - 하지만 시스템의 완성도 측면에서는 **Reflection**이 가장 독특합니다. 많은 AI 학습 시스템이 Education(가르치기)과 Knowledge Capture(요약)까지만 제공하고, Reflection(엄격한 검증)은 생략합니다. Cocrates가 이 세 단계를 모두 갖춘 것이 가장 큰 차별점입니다.
  - 개인적으로는 **Knowledge Capture의 Ignorance 기록**이 가장 실용적입니다. 이 개념만 자기 학습에 적용해도 학습 효과가 크게 향상됩니다.
- **Reference:** slides.md Slide 11 — 세 스킬의 유기적 연결.

## Q6: "이 세 스킬은 다른 프로젝트에서도 사용할 수 있나요? 아니면 Cocrates 전용인가요?"
- **Context:** 실용적인 청중이 자신의 프로젝트에 적용 가능한지 질문.
- **Answer strategy:**
  - 원리 자체는 범용적입니다. Education의 3블록 구조, Knowledge Capture의 Ignorance 기록, Reflection의 Pull-only 평가는 어떤 학습 시스템에도 적용할 수 있는 설계 패턴입니다.
  - 다만, Cocrates는 이 세 스킬이 **하나의 프레임워크로 통합**되어 있다는 점이 차별점입니다. 각 스킬이 같은 Bloom's 2D 언어로 소통하고, KB가 Education과 Reflection을 연결하는 가교 역할을 하는 설계는 범용적으로 적용하기 어려운 통합 수준입니다.
  - 개별 스킬의 아이디어만 차용한다면: (1) 학습할 때 '요약'보다 '오개념 기록'에 집중하고, (2) 스스로에게 '반례를 들어보세요'라고 질문해보는 것만으로도 효과를 볼 수 있습니다.
- **Reference:** 전체 Ep9 내용 + Ep8 Learning 파이프라인 원리.
