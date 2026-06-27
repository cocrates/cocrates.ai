---
sidebar_position: 12
---
# EP12. 구조 기반 워크플로우 생성 스킬

![Cocrates Skill Creation Principles](/img/12-workflow-creation-skill.png)

---

지금까지 우리는 Cocrates의 두 가지 생성 방식을 배웠다.

**첫 번째 방식:** 요청한 A에 전용 스킬이 있으면 그 스킬로 생성한다. 예를 들어, *"블로그 시리즈를 생성해줘"* → `blog-series-authoring` 스킬 발동. 간단하고 빠르다.

**두 번째 방식:** 전용 스킬이 없으면 **spec-driven-generation**으로 폴백한다. ASR 식별 → ADR → Spec → Generation → Verification. 구조 기반 생성 파이프라인이 가동된다.

그런데 Ep11을 읽으면서 이런 생각이 들지 않았는가?

*"좋다. 근데... 내가 필요한 스킬이 아예 없으면?"*

예를 들어 *"정기적인 코드 리뷰 리포트를 생성해줘"* 라는 스킬이 필요하다고 해보자. Cocrates에는 그런 스킬이 없다. spec-driven-generation으로 폴백해서 매번 새로 만들 수도 있지만, 같은 패턴의 작업을 반복할 거라면 **스킬 자체를 만들어 버리는 게 낫다.**

오늘 우리가 배울 것은 바로 이것이다. **스킬을 만드는 스킬 — generating-skill-creation.**

---

## 🔍 메타-스킬: 스킬을 만드는 스킬

generating-skill-creation은 Cocrates의 **메타-스킬(Meta-Skill)** 이다. 메타-스킬이란, 일반 스킬처럼 산출물을 생성하는 스킬이 아니라 **스킬 그 자체를 설계하고 생성하는 스킬**이다.

일반 스킬과 메타-스킬의 차이를 비유로 이해해보자.

- 일반 스킬: *"물고기를 잡아줘"* → 물고기를 잡는 기술
- 메타-스킬: *"물고기 잡는 법을 가르쳐줘"* → 낚시를 가르치는 기술

generating-skill-creation은 5대 Artifact Components를 기반으로 동작한다.

**5대 Artifact Components:**

1. **Assignment & Constraints:** 이 스킬이 해결해야 할 과제와 제약 조건은 무엇인가?
2. **Context & Rules:** 스킬이 동작할 맥락과 따라야 할 규칙은 무엇인가?
3. **Entities:** 스킬이 다루는 주요 개념과 객체는 무엇인가?
4. **Space & Placement:** 결과물이 어디에 위치하고 어떻게 배치되는가?
5. **Structure & Flow:** 스킬의 단계별 흐름과 구조는 무엇인가?

이 5가지 구성 요소를 정의하는 것이 스킬 생성의 첫걸음이다.

---

## ❄️ Snowflake 5단계 — 스킬을 점진적으로 구체화

5대 Components가 식별되면, Cocrates는 **Snowflake 5단계**를 통해 스킬을 점진적으로 구체화한다.

**단계 1 — Define:** *"이 스킬은 무엇을 하는가?"* Kernel을 한 문장으로 정의한다. 스킬의 핵심 목적을 압축한다.

**단계 2 — Plan:** *"어떤 순서로 진행할 것인가?"* 스킬의 전체 단계와 순서를 계획한다.

**단계 3 — Architecture Design:** *"구성 요소 간 관계는 무엇인가?"* Entities, Space, Flow 등 아키텍처를 설계한다.

**단계 4 — Detail Design:** *"각 단계의 구체적인 동작은 무엇인가?"* 각 단계의 입력, 출력, 승인 조건, 금지사항을 상세히 정의한다.

**단계 5 — Generation:** *"SKILL.md 파일을 생성한다."* 모든 설계가 완료된 후에야 최종 파일을 생성한다.

여기서 가장 중요한 규칙이 하나 있다.

---

## 🚫 절대 규칙 — Detail Design 완성 전에는 Generation 금지

이 규칙은 Cocrates의 생성 철학 중에서도 가장 엄격하게 지켜진다.

> **"Detail Design이 완전히 확정되기 전까지 Generation 단계로 넘어가지 않는다."**

왜일까? Ep2에서 선언한 원칙을 기억하는가? *"The unexamined code is not worth generating."* 스킬은 AI의 행동 방식을 영구적으로 결정한다. 잘못된 스킬은 잘못된 결과를 반복해서 양산한다.

그래서 Cocrates는 Detail Design 단계에서 다음 질문들을 철저히 검토한다.

- 각 단계의 입력과 출력이 명확한가?
- 승인 게이트가 모든 단계에 있는가?
- 금지사항이 구체적인가?
- 예외 상황 처리는 어떻게 할 것인가?

이 모든 질문에 답이 나오기 전까지는 단 한 줄의 SKILL.md도 생성하지 않는다.

---

## 🔄 Meta Snowflake 7단계 — 더 큰 그림

Snowflake 5단계는 스킬을 설계하는 기본 프로세스다. 하지만 generating-skill-creation은 이보다 더 큰 그림을 그린다. **Meta Snowflake 7단계**다.

```
Kernel → Components → Frame → Outline → Spec → Skill → Verification
```

이 7단계를 순서대로 설명하면 이렇다.

**Kernel:** 한 문장으로 스킬의 목적을 정의한다.

**Components:** 5대 Artifact Components를 식별하고 정의한다.

**Frame:** 스킬의 전체 구조와 뼈대를 설계한다. 어떤 단계로 구성될지 개괄적인 프레임을 잡는다.

**Outline:** Frame을 구체화한다. 각 단계의 세부 항목과 흐름을 상세하게 기술한다.

**Spec:** 모든 결정을 통합한 자체 완결적인 명세서를 작성한다. 이 Spec만 보면 누구나 스킬의 전체 동작 방식을 이해할 수 있어야 한다.

**Skill:** Spec을 근거로 실제 SKILL.md 파일을 생성한다.

**Verification:** 생성된 스킬이 Spec과 일치하는지, 누락된 부분은 없는지 검증한다.

이 7단계는 단순한 순서도, 병렬도 아니다. 각 단계는 이전 단계의 결과에 의존하며, 검증에서 문제가 발견되면 이전 단계로 피드백된다.

---

## 🎯 Per-Stage Refinement — 지연 구체화 전략

Meta Snowflake 7단계에서 한 가지 중요한 원칙이 있다.

**각 구성 요소는 동일한 속도로 구체화되지 않는다.**

어떤 요소는 초기 단계에서 매우 구체적으로 정의되고, 어떤 요소는 마지막 단계에서야 명확해진다. Cocrates는 이것을 **Lazy Refinement(지연 구체화)** 전략이라고 부른다.

예를 들어, *"Entities"* 는 Kernel과 Components 단계에서 대략적인 윤곽만 잡고, Outline과 Spec 단계에서 구체화된다. 반면 *"Structure & Flow"* 는 Frame 단계에서부터 상세히 설계되어야 한다.

이 전략이 중요한 이유는 **불필요한 조기 결정을 방지**하기 때문이다. 너무 이른 단계에서 세부 사항을 결정하면, 나중에 더 나은 대안이 떠올라도 변경하기 어렵다. 지연 구체화는 결정을 최대한 늦춤으로써 유연성을 확보한다.

---

## 💡 오늘의 깨달음: 스킬은 AI의 SOP다

생각해보면, 스킬은 AI에게 **SOP(Standard Operating Procedure, 표준 업무 절차)** 를 주는 것과 같다.

회사에서 신입 사원이 들어오면, 우리는 일하는 방법을 가르친다. *"이런 요청이 오면, 먼저 A를 확인하고, B를 검토한 후, C를 수행해"* 라는 식으로 말이다.

스킬이 바로 그 역할을 한다. AI에게 *"이런 유형의 요청이 오면, 이렇게 일해라"* 라는 SOP를 정의해주는 것이다.

그리고 한 번 정의된 스킬은 재사용 가능한 자산이 된다. 오늘 만든 스킬은 내일도, 다음 주에도, 1년 후에도 같은 품질로 동작한다. 

**일회성 요청은 소모품이다. 하지만 스킬은 자산이다.**

---

## 📌 오늘의 핵심

1. **generating-skill-creation은 메타-스킬이다.** 스킬을 만드는 스킬. 5대 Artifact Components를 기반으로 동작한다.
2. **Snowflake 5단계로 점진적 구체화한다.** Define → Plan → Architecture Design → Detail Design → Generation. 각 단계는 승인 게이트를 통과해야 다음으로 넘어간다.
3. **절대 규칙: Detail Design 완성 전 Generation 금지.** 잘못된 스킬은 잘못된 결과를 반복해서 양산한다.
4. **Meta Snowflake 7단계가 더 큰 그림이다.** Kernel → Components → Frame → Outline → Spec → Skill → Verification. 각 단계의 결과는 다음 단계의 입력이 된다.
5. **지연 구체화 전략.** 모든 구성 요소를 동시에 구체화하지 않는다. 필요한 순간에 필요한 만큼만 구체화한다.

---

**스스로에게 질문하자.**

- Snowflake 5단계와 Meta Snowflake 7단계의 차이를 설명할 수 있는가?
- 5대 Artifact Components를 말하고 각각을 설명할 수 있는가?
- Lazy Refinement가 왜 중요한 전략인가?

---

## 🎬 다음 편 예고

지금까지 우리는 Cocrates의 모든 것을 배웠다.

Ep1~Ep2에서 철학과 원칙을, Ep3에서 설치를, Ep4~Ep6에서 실습을, Ep7~Ep10에서 내부 구조와 원리를, Ep11~Ep12에서 핵심 스킬과 메타-스킬을 탐구했다.

**이제 우리는 Cocrates를 사용할 준비가 되었다.**

그런데... 진짜 중요한 질문이 아직 남아 있다.

*"이걸로 끝일까? Cocrates를 그냥 '잘' 사용하면 되는 걸까?"*

아니다. 진짜는 지금부터다. Cocrates를 사용하는 것만으로는 충분하지 않다.

**다음 편, Ep13의 주제는 이것이다.** 사용자도 진화해야 하고, Cocrates Harness도 진화해야 한다. 그리고 이 두 진화는 서로를 강화하는 피드백 루프를 형성한다.

*"Cocrates를 사용하는 것으로 충분하지 않다 — Cocrates를 진화시켜야 한다."*

---

*이 시리즈는 Cocrates Harness 프레임워크를 소개합니다. Cocrates는 소크라테스식 대화로 사용자가 주도권을 잡고 성장하도록 설계된 에이전트 하네스입니다.*
