# Script: Slide 06 — 🔀 Intent-To-Skill Routing — 의도가 스킬을 선택한다

**Time:** ~1.5분

**[SLIDE 06: 매핑 테이블 (Intent → Skill). 상단: "단순 키워드 매칭이 아닌 의도(Intent) 추론"]**

자, 이제 Cocrates의 두뇌라고 할 수 있는 **Intent-To-Skill Routing**을 살펴보겠습니다.

기본 원리는 간단합니다. Cocrates는 사용자의 요청을 단순 키워드로 매칭하지 않습니다. 대신 **사용자의 근본 의도, Intent를 추론**하여 최적의 스킬을 매핑합니다.

화면에 보이는 테이블을 봐주세요.

사용자가 "개념을 배우고 싶다"는 의도를 가지고 질문하면, Cocrates는 `education` 스킬을 로드합니다. "배운 내용을 정리하고 싶다"면 `knowledge-capture`, "이해도를 확인받고 싶다"면 `reflection` 스킬이 발동됩니다.

선택지를 비교하고 결정해야 할 때는 `adr-writing`, 명세를 정의해야 할 때는 `spec-writing`, 산출물을 생성해야 할 때는 `spec-driven-generation`, 검증이 필요할 때는 `spec-driven-verification` 스킬이 각각 로드됩니다.

그리고 "스킬을 만들고 싶다"는 명시적 요청이 있으면 `generating-skill-creation`이라는 메타 스킬이 발동됩니다.

이것이 Cocrates가 다양한 요청을 정확히 처리할 수 있는 비결입니다. 여러분이 "블룸의 분류학에 대해 알려줘"라고만 해도, Cocrates는 이것을 단순한 정보 요청이 아니라 **"개념을 배우고 싶다"** 는 의도로 해석해서 `education` 스킬을 로드하는 거죠.

만약 매핑되는 스킬이 없으면, 기본 생성 방식인 `spec-driven-generation`이 처리합니다.

이제 이 Intent-To-Skill Routing을 통해 선택된 스킬들이 실제로 어떤 활동을 수행하는지, Cocrates의 두 가지 Core Activity 파이프라인을 살펴보겠습니다.

→ [SLIDE 07: 🏗️ Core Activity 1 — 산출물 생성 파이프라인]
