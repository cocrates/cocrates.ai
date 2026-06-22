# 구조 기반 워크플로우 생성 스킬 — generating-skill-creation

---

지난 두 편에서는 구조 기반 산출물 생성의 핵심 파이프라인을 살펴보았습니다. ADR(선택지 분석과 결정) → Spec(결정의 통합) → Generation(Spec 기반 생성) → Verification(항목별 검증)의 4단계 흐름과 각 스킬의 역할을 이해했습니다.

이 접근법은 근본적으로 **구조가 명확하지 않은 상황**을 전제로 합니다. 산출물의 구조가 모호하거나, 결정되지 않은 사항이 많을 때 — Cocrates는 ASR을 식별하고, ADR로 대안을 분석하고, Spec으로 통합한 뒤에야 생성합니다.

그런데 질문이 생깁니다. **산출물 구조가 이미 명확하다면 어떨까요?**

예를 들어, "발표자료를 만들어줘"라고 요청했을 때 — 발표자료는 이미 명확한 구조를 가집니다. 슬라이드 덱, 섹션, 페이지, 각 페이지의 거버닝 메시지와 서포트... 굳이 매번 ASR부터 시작할 필요가 없습니다. 더 적합한, 더 구체적인 워크플로우가 존재할 수 있습니다.

바로 이런 경우를 위해 Cocrates Harness는 **artifact-specific skill(산출물 유형별 스킬)** 을 제공합니다. document-authoring, presentation-authoring, blog-series-authoring 같은 스킬들은 각 산출물 유형에 최적화된 점진 구체화 워크플로우를 정의합니다.

그렇다면 이 스킬들은 어떻게 탄생했을까요? 이 질문의 답이 **generating-skill-creation**입니다.

---

## 1. generating-skill-creation — 스킬을 만드는 스킬

generating-skill-creation은 **메타-스킬(Meta-Skill)** 입니다. 스킬 자체를 만드는 스킬입니다. 최종 산출물을 생성하는 것이 아니라, 그 산출물을 생성하는 **절차 그 자체를 설계하고 SKILL.md 파일로 작성**합니다.

### 발동 조건 — 오직 명시적 요청일 때만

generating-skill-creation은 사용자가 **"~을 생성하는 스킬을 만들어줘"** 라고 명시적으로 요청할 때만 작동합니다.

이 조건이 왜 중요할까요? 일반적인 생성 요청("발표자료 만들어줘", "보고서 작성해줘")에서 적합한 스킬이 없으면 어떻게 될까요?

이미 Ep7에서 보았듯이, 그 경우에는 **spec-driven-generation의 ASR Readiness Gate**가 처리합니다. 스킬이 없어도, 산출물 생성을 요청받으면 Cocrates는 기본적으로:
1. 산출물 유형을 추론하고
2. ASR을 식별하고
3. 필요한 ADR을 작성하고
4. Spec으로 통합하고
5. 생성합니다

즉, **스킬이 없다고 해서 generating-skill-creation이 발동하는 것은 아닙니다.** spec-driven-generation이라는 기본 파이프라인이 작동할 뿐입니다.

generating-skill-creation은 이와 근본적으로 다릅니다. 사용자가 "이런 유형의 산출물을 생성하는 **절차 자체를 시스템에 추가**해줘"라고 요청할 때만 활성화됩니다. 전자는 **단일 산출물**을 대상으로 하고, 후자는 **산출물 유형을 위한 일반화된 절차**를 대상으로 합니다.

### spec-driven-generation과의 차이

| | spec-driven-generation | generating-skill-creation |
|---|----------------------|--------------------------|
| **발동 조건** | "이 산출물을 생성해줘" + 스킬 없음 | "이 산출물을 생성하는 스킬을 만들어줘" |
| **대상** | 단일 산출물 하나 | 반복 가능한 생성 절차 (스킬) |
| **결과** | 최종 문서, 코드, 이미지 등 | `.opencode/skills/{slug}/SKILL.md` |
| **재사용성** | 매번 처음부터 ASR 식별 | 한 번 만들면 같은 유형 요청에 재사용 |

---

## 2. 산출물을 분해하는 다섯 가지 구성 요소

새 스킬을 설계할 때 generating-skill-creation이 가장 먼저 하는 일은 **최종 산출물을 다섯 가지 차원으로 분해**하는 것입니다. 산출물 유형에 따라 이름은 달라질 수 있지만, 역할은 동일합니다:

| 구성 요소 | 역할 |
|-----------|------|
| **과업과 제약 (Assignment & Constraints)** | 범위, 품질 기준, 금기 사항, 성공 조건 |
| **맥락과 규칙 (Context & Rules)** | 도메인 규칙, 세계관, 스타일 가이드 |
| **개체 (Entities)** | 재사용 가능한 구성 요소 (캐릭터, 모듈, 섹션 등) |
| **공간과 배치 (Space & Placement)** | 위치, 장면, 레이아웃 영역 |
| **구조와 흐름 (Structure & Flow)** | 계층적 콘텐츠 구성 |

예를 들어, 블로그 시리즈 생성 스킬을 설계한다면:

- **과업과 제약**: 시리즈 주제, 예상 독자, 에피소드 수, 분량
- **맥락과 규칙**: 블로그 플랫폼 특성, 톤 앤 매너, 마크다운 규칙
- **개체**: 각 에피소드, 개요(outline) 문서, episodes.md 파일
- **공간과 배치**: 에피소드의 순서와 배치 — 시리즈 아크에서의 위치
- **구조와 흐름**: 시리즈 전체의 전개 — 서론→본론→결론, 각 에피소드의 역할

이 다섯 가지 차원을 모두 식별하고, 각각이 어느 단계에서 어느 수준까지 구체화될지 결정하는 것이 스킬 설계의 첫걸음입니다.

---

## 3. Snowflake 5단계 — 점진 구체화

generating-skill-creation이 새 스킬의 생성 절차를 구체화하는 방식은 **Snowflake 5단계**입니다:

| 단계 | 목적 | 승인 게이트 |
|------|------|------------|
| **define** | 과업, 범위, 제약, 성공 기준 확정 | 모든 후속 결정의 기준선 |
| **plan** | 스킬의 골격, 방향, 주요 구성 요소 개관 | 전체 리듬과 방향 승인 |
| **architecture design** | 구성 요소별 카탈로그, 계층 구조, 상호 참조 | 구조적 일관성 승인 |
| **detail design** | 생성에 필요한 단위별 명세 확정 | **생성 전 최종 게이트** |
| **generation** | 확정된 명세에 따라 SKILL.md 생성 및 조립 | 단위별 및 최종 품질 검토 |

여기서 가장 중요한 규칙은 하나입니다:

> **detail design이 완전히 확정되기 전까지 generation 단계로 넘어가지 않는다.**

이 게이트가 없다면, generating-skill-creation은 '그냥 만들어보고 수정'하는 일회성 프롬프트와 다를 바 없습니다. 구조를 검토하고, 단계별로 승인받고, 최종 명세가 완성된 후에야 비로소 SKILL.md를 생성합니다.

### 단계별 구성 요소 구체화

모든 구성 요소가 같은 속도로 구체화되지는 않습니다. generating-skill-creation이 새 스킬을 설계할 때 **각 구성 요소가 각 단계에서 어느 수준까지 구체화되는지**를 정의하는 것이 핵심입니다.

예를 들어, **개체(Entities)** 구성 요소의 구체화 과정:

| 단계 | 개체의 구체화 수준 |
|------|------------------|
| define | 개체 유형만 식별 (예: "에피소드", "개요") |
| plan | 핵심 역할과 관계 요약 |
| architecture design | 카탈로그, 관계 맵, 계층 구조 |
| detail design | 개체별 상세 명세, 변형, 상태 |
| generation | 개체별 산출물 생성 |

구조와 흐름이 개체나 공간을 참조할 때는 먼저 카탈로그를 완성하고, 상세 명세는 필요할 때 채우는 **지연 구체화(lazy refinement)** 전략을 사용합니다. 모든 것을 한 번에 완벽하게 정의하려고 하면 설계가 무거워지고 사용자의 피로도가 높아지기 때문입니다.

---

## 4. Meta Snowflake — 스킬 저작 7단계

generating-skill-creation 자신이 새 SKILL.md를 만들 때는 다음 7단계로 진행합니다. 이것을 **Meta Snowflake**라고 부릅니다:

| 단계 | Snowflake 대응 | 산출물 |
|------|---------------|--------|
| **Kernel** | define | 생성 대상을 한 문장으로 정의 ("이 스킬은 ~을 생성하도록 돕는다") |
| **Components** | plan | 다섯 가지 구성 요소 차원 식별 |
| **Frame** | architecture design | 5단계 Snowflake 워크플로우, 파일 구조, 승인 지점 설계 |
| **Outline** | detail design | 각 단계별 파일 산출물, 입력, 생성 활동, 완료 기준, 검토 질문, 승인 조건 정의 |
| **Spec** | generation | frontmatter, 원칙, 사용 시점, 절차, 대화 규칙, 금지 사항, 완료 조건 확정 |
| **Skill** | generation | `.opencode/skills/{skill-slug}/SKILL.md` 파일 생성 |
| **Verification** | generation | 스킬이 생성 하네스로 작동하는지 검증 |

### Kernel — 한 문장으로 정의

Kernel 단계에서는 새 스킬이 다룰 생성 대상을 한 문장으로 정의합니다:

> "이 스킬은 {대상 사용자/상황}에서 {산출물 유형}을 {검토 가능한 단계}를 거쳐 생성하도록 돕는다."

이 한 문장이 스킬의 정체성을 결정합니다. 모호함이 남아 있다면 다음 단계로 넘어가지 않습니다.

### Outline — 파일 산출물 정의

Outline 단계에서는 각 단계에서 생성되는 중간 파일을 다음 형식으로 정의합니다:

- **Input**: 의존하는 산출물
- **Creation activity**: 이 파일을 생성하기 위한 작업
- **Completion criteria**: 이 파일이 완료된 것으로 간주되는 조건
- **Review questions**: 사용자 확인을 위한 질문
- **Approval criteria**: 다음 단계로 진행하기 위한 승인 조건

이 정의가 있어야 중간 산출물이 채팅에만 머무르지 않고, **파일로 관리**되며, **명확한 완료 기준**을 가집니다.

### Generation — SKILL.md 작성

Spec이 확정되면 SKILL.md를 생성합니다. frontmatter에는 반드시 다음이 포함되어야 합니다:

```yaml
---
name: {skill-slug}
description: >-
  {이 스킬이 무엇을 하고 언제 쓰이는지 구체적으로 설명}
compatibility: opencode
metadata:
  agent: cocrates
---
```

그리고 본문에는 핵심 원칙, 사용 시점, 작업 위치, 중간 산출물 정의, 단계별 절차, 대화 규칙, 금지 사항, 완료 조건이 포함됩니다.

---

## 5. 스킬의 재사용성과 진화

generating-skill-creation으로 스킬을 한 번 만들면, 이후 같은 유형의 요청에서는 generating-skill-creation의 개입 없이 **해당 스킬이 직접 로드되어 사용**됩니다. 이것이 스킬을 만드는 이유입니다. 매번 ASR부터 시작하는 spec-driven-generation보다 훨씬 효율적이고, 산출물 유형에 최적화된 워크플로우를 제공할 수 있습니다.

실제로 Cocrates Harness의 artifact-specific 스킬들 — document-authoring, presentation-authoring, blog-series-authoring — 은 모두 이 generating-skill-creation 방식으로 탄생했습니다.

새로운 산출물 유형이 필요하면 generating-skill-creation을 통해 스킬만 추가하면 됩니다. 기존 스킬은 독립적으로 개선할 수 있습니다. 이것이 Cocrates Harness가 **Agent + Skills** 구조로 설계된 이유 중 하나입니다.

---

## generating-skill-creation 요약

generating-skill-creation의 핵심을 한 문장으로 요약하면 이렇습니다:

> **최종 산출물을 생성하는 절차를, 검토 가능한 단계로 분해하고 승인 게이트를 두어, 반복 가능한 스킬로 만드는 메타-스킬.**

generating-skill-creation까지 이해했다면, Cocrates Harness의 **모든 구조적 활동**을 살펴본 셈입니다. 소개(Ep1~2)에서 문제 인식과 원칙을, 실습(Ep3~6)에서 설치와 세 가지 실습을, 구조(Ep7~12)에서 시스템 구조, 학습 활동, 산출물 생성 활동, 그리고 워크플로우 생성 스킬까지 — 이제 Cocrates가 무엇이고, 어떻게 작동하는지 전체 그림이 그려졌을 것입니다.

모든 원리를 이해했으니, 이제 마지막으로 Cocrates를 지속적으로 발전시키는 방법과 사용자로서의 다짐을 알아보겠습니다.

---

*이 시리즈는 Cocrates Harness 프레임워크를 소개합니다. Cocrates는 소크라테스식 대화로 사용자가 주도권을 잡고 성장하도록 설계된 에이전트 하네스입니다.*

*→ 다음 편: 올바른 Cocrates Harness 활용 — 함께 진화하는 도구*
