# Expected Questions & Answer Plans — Episode 07

---

## Q1: "Skills가 독립적이라면, 같은 기능이 여러 스킬에 중복될 수 있지 않나요?"

- **Context:** Slide 4(두 가지 기둥)에서 Skills의 독립성을 설명할 때 나올 수 있는 질문. 중복 코드 우려.
- **Answer strategy:**
  - "네, 가능합니다. 하지만 의도적인 설계입니다."
  - "각 스킬은 완전히 독립적이어야 하므로, 공통 기능이 필요하면 각 스킬이 자체적으로 구현합니다."
  - "중복은 드라이(Dry) 원칙에 위배되는 것처럼 보일 수 있지만, 스킬의 독립성과 진화 가능성이라는 더 큰 가치를 위해 중복을 허용합니다."
  - "실제로 대부분의 스킬은 고유한 워크플로우를 가지므로 중복은 생각보다 적습니다. 만약 진짜 공통 기능이 필요하다면, Agent Prompt의 Principle 섹션에서 공통 원칙으로 정의하고 각 스킬이 참조하는 방식으로 해결합니다."
- **Reference:** `cocrates.md` — Principle 섹션, Ep6 — 스킬 생성 파이프라인

---

## Q2: "Intent-To-Skill Routing이 항상 정확한가요? 의도를 잘못 추론하면 어떻게 되나요?"

- **Context:** Slide 6(Routing)에서 나올 수 있는 실용적 질문. AI의 의도 추론 정확도에 대한 우려.
- **Answer strategy:**
  - "완벽하지는 않습니다. 하지만 Cocrates는 Socratic 접근을 통해 이 문제를 완화합니다."
  - "첫째, Cocrates는 사용자의 요청을 추론할 때 단순히 매핑만 하지 않고, 사용자에게 확인을 구합니다. 예를 들어 '학습을 원하시는 건가요, 문서 작성을 원하시는 건가요?'라고 되묻는 식입니다."
  - "둘째, 매핑 테이블은 사용자의 피드백을 통해 지속적으로 개선됩니다. 잘못된 매핑이 발견되면 해당 스킬의 라우팅 조건을 수정합니다."
  - "셋째, 매핑되는 스킬이 없거나 불확실한 경우, 기본 생성 방식(`spec-driven-generation`)으로 fallback하여 사용자와 함께 요구사항을 명확히 하는 과정을 거칩니다."
- **Reference:** `cocrates.md` — Request Handling, Intent-To-Skill Routing 표

---

## Q3: "6개 섹션 중에 가장 중요한 섹션은 무엇인가요?"

- **Context:** Slide 5(6개 섹션 개관)에서 나올 수 있는 심층 질문.
- **Answer strategy:**
  - "상황에 따라 다르지만, 시스템 설계 관점에서 가장 중요한 섹션은 **Principle**과 **Request Handling**입니다."
  - "**Principle(Harness Ignorance)** 은 Cocrates의 정체성을 결정합니다. 이 원칙이 없으면 Cocrates는 다른 AI와 다를 바 없는 단순 정보 제공자가 됩니다."
  - "**Request Handling(Intent-To-Skill Routing)** 은 Cocrates가 다양한 요청을 정확히 처리할 수 있게 하는 실질적 엔진입니다."
  - "하지만 **Success Criteria**도 빼놓을 수 없습니다. '사용자가 산출물을 설명할 수 있는가?'라는 기준이 Cocrates의 모든 행동을 사용자 중심으로 유지하게 합니다."
- **Reference:** `cocrates.md` 전체 구조

---

## Q4: "Part 3(구조 편)에서 어떤 순서로 진행되나요? 지금 Ep7의 위치는?"

- **Context:** Slide 11(다음 에피소드 연결)에서 시리즈 전체 로드맵에 대한 질문.
- **Answer strategy:**
  - "Part 3는 총 6개 에피소드(Ep7~Ep12)로 구성됩니다."
  - "**Ep7(현재):** Cocrates Harness 전체 구조 개관 — Agent + Skills, 6개 섹션, Routing, Core Activities"
  - "**Ep8~9:** 첫 번째 Core Activity인 학습 파이프라인의 원리(Ep8)와 스킬(Ep9)"
  - "**Ep10~11:** 두 번째 Core Activity인 산출물 생성 파이프라인의 원리(Ep10)와 스킬(Ep11)"
  - "**Ep12:** 스킬을 생성하는 메타 스킬(generating-skill-creation)"
  - "즉, Ep7이 전체 지도를 보여준다면, Ep8~12는 지도의 각 영역을 확대해서 보는 방식입니다."
- **Reference:** `overview.md` — Part 3 에피소드 목록

---

## Q5: "Slide 9에서 Cocrates는 완성된 하네스가 아니라며 진화를 강조했는데, 실제로 어떻게 진화시키나요?"

- **Context:** Slide 9(진화)에서 나올 수 있는 실무적 질문. 구체적인 방법론.
- **Answer strategy:**
  - "세 가지 방법이 있습니다."
  - "**첫째, 기존 스킬 수정.** 예를 들어 `document-authoring` 스킬의 Frame 단계에서 문서 레이아웃을 회사 템플릿에 맞게 변경하는 것입니다. 스킬 파일(SKILL.md)을 직접 편집하면 됩니다."
  - "**둘째, 새 스킬 생성.** Ep6에서 배운 `generating-skill-creation` 스킬을 사용하면 됩니다. '새로운 보고서 형식을 위한 스킬을 만들어줘'라고 요청하면 Cocrates가 Kernel→Frame→Outline→Spec→Skill 과정을 안내합니다."
  - "**셋째, Agent Prompt 조정.** `cocrates.md`의 Principle이나 Success Criteria 섹션을 자신의 팀 문화에 맞게 수정할 수 있습니다. 예를 들어 성공 기준에 '팀 리뷰를 거칠 것'을 추가하는 식입니다."
  - "중요한 것은, 구조를 이해해야 무엇을 어떻게 수정할지 알 수 있다는 점입니다. 이것이 Ep7의 핵심 메시지입니다."
- **Reference:** Ep6 — 스킬 생성 파이프라인, Ep12 — generating-skill-creation

---

## Q6: "이 구조는 opencode에만 적용되나요? 다른 AI 플랫폼에서도 사용할 수 있나요?"

- **Context:** 실무 적용에 관한 질문. opencode 특화 구조인지 범용적인지.
- **Answer strategy:**
  - "Cocrates Harness는 현재 opencode 플러그인으로 구현되어 있지만, 그 **원칙과 구조는 범용적**입니다."
  - "Agent + Skills 분리, Intent-To-Skill Routing, Core Activity 파이프라인 — 이 개념들은 모든 AI 플랫폼에 적용할 수 있는 설계 패턴입니다."
  - "예를 들어, Claude의 Projects나 GPTs에서도 유사한 구조를 구현할 수 있습니다. 시스템 프롬프트를 Agent로, 개별 instruction 파일을 Skills로 대응시키는 방식입니다."
  - "다만 opencode는 이러한 구조를 구현하기 위한 **플랫폼 지원**(plugin 시스템, skills 디렉토리, config 파일 등)을 기본 제공하므로 가장 적합한 환경입니다."
- **Reference:** Ep3 — opencode 설치, `cocrates.md` — 플랫폼 중립적 원칙

---

## Q7: "Ep7은 너무 개념적인 내용인데, 실제 코딩이나 실습은 없나요?"

- **Context:** Part 3의 첫 편으로, 실습(Ep4~6)에 익숙해진 청중이 개념 설명에 지루함을 느낄 수 있음.
- **Answer strategy:**
  - "네, Ep7은 의도적으로 개념 중심입니다. Part 3의 첫 편으로서 전체 지도를 보여주는 역할을 합니다."
  - "마치 새로운 도시를 탐험할 때, 처음에는 전체 지도를 보는 것이 효율적입니다. 그런 다음 각 동네를 하나씩 방문하는 거죠."
  - "Ep8부터는 다시 구체적인 내용으로 들어갑니다. Ep8~9는 학습 파이프라인의 원리와 스킬, Ep10~11은 생성 파이프라인의 원리와 스킬을 다루며, 각각 Ep4와 Ep5의 실습을 구조적 관점에서 재해석합니다."
  - "그리고 Ep12는 다시 실습으로 돌아와서, 메타 스킬을 직접 만들어보는 과정을 다룹니다."
- **Reference:** `overview.md` — Part 3 구성, Ep4~Ep6 실습 경험
