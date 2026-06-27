# AI를 올바르게 쓰는 법 — Cocrates Harness

## 핵심 원칙 (Core Principle)

> **The unexamined code is not worth generating.**
> (검토되지 않은 산출물은 생성할 가치가 없다.)

**code**란 AI를 사용해서 만들고자 하는 **최종 산출물**을 의미한다. 소스 코드뿐 아니라 AI가 만드는 보고서, 슬라이드, 블로그, 학습 노트 등 모든 것이 Code다.

**검토(examine)** 의 본질은 **이해(U) → 분석(A) → 평가(E) → 승인(A)** 의 4단계로 이루어진다. 검토의 목적은 **무지의 제거(Harnessing Ignorance)** 에 있다. AI가 구조를 설계해도 내가 이해하지 못했다면 그것은 검토된 것이 아니며, 내게 무지만 남긴다.

Cocrates는 이 원칙을 실현하는 하네스다. 사용자가 AI를 올바른 방법으로 사용할 수 있도록 워크플로우를 안내하고, 모든 단계에서 사용자의 이해와 승인을 요구한다.

---

## 시리즈의 근본 질문

> **"당신의 AI는 몇 명의 팀원입니까?"**

AI-native Engineer는 AI를 단순한 도구(AI-assisted)가 아니라 **일을 수행하는 팀원**으로 대한다. 지시 → AI가 수행 → 내가 검토(이해→분석→평가→승인). 당신은 몇 명의 AI 팀원을 리딩해서 업무를 수행할 수 있는가? **결국 당신의 역량이 중요하다.**

### AI-assisted Engineer vs AI-native Engineer

| | AI-assisted | AI-native |
|---|---|---|
| **AI의 역할** | 나를 도와주는 **도구** | 일을 수행하는 **팀원** |
| **사용자의 역할** | AI의 도움을 받아 직접 수행 | 지시하고, 검토(U→A→E→A)하고, 승인 |
| **검토 방식** | "AI가 잘 짰겠지" → 그냥 넘어감 | "내가 이해했는가?" → 검토 후 승인 |
| **결과** | 코드↑, 이해도↓, 무지↑ | 코드↑, 이해도↑, 역량↑ |

---

## 용어 정의

이 시리즈에서 사용하는 핵심 용어를 먼저 명확히 정의한다.

| 용어 | 의미 | 비고 |
|------|------|------|
| **Cocrates** | 전체 시스템을 지칭하는 이름. "Cocrates는 무엇인가?"라는 질문에 대한 답은 아래 Cocrates Harness와 같다. | |
| **Cocrates Harness** | 전체 시스템 = **Cocrates Agent + Skills** | 원문: *"Cocrates is a **Cocrates agent plus skills** harness."* |
| **Cocrates Agent** | 핵심 에이전트. `cocrates.md` 파일로 정의되며, 원칙·의도 인식·스킬 선택·태스크 관리·가드레일을 담당한다. | "You are **Cocrates**"의 주체 |
| **Skills** | 작업별 절차. `.opencode/skills/*/SKILL.md`로 관리되며 추가·수정·확장 가능하다. | |

즉, **Cocrates Harness = Cocrates Agent + Skills**이며, 이 시리즈의 **3부(구조)** 는 바로 이 Cocrates Harness를 소개한다.

### Cocrates라는 이름의 의미

> **Cocrates = Co(상호) + Socrates(소크라테스)**

**두 방향의 소크라테스 관계**가 핵심이다:
- **사용자 → AI:** 사용자가 소크라테스가 되어 AI에게 질문하고, AI의 결과물을 검토·판단하여 잘못된 생성을 방지한다.
- **AI → 사용자:** AI가 소크라테스가 되어, 사용자가 막연한 요청을 하거나 산출물을 모른 채 승인하지 않도록 방지한다. 대충 넘어가는 것을 제대로 이해하도록 하고, 무지를 알게 하여 스스로 역량을 키우도록 돕는다.

이 **"상호 소크라테스 관계"** 로 AI를 사용하게 하는 것이 Cocrates Harness의 핵심 철학이다.

### Cocrates Agent의 역할

Cocrates Agent (`cocrates.md`)는 두 가지 **Core Activity** 파이프라인을 정의한다:
- **Artifact Generation**: 설계(ADR → Spec) → Spec 기반 생성 → Spec 기반 검증
- **Learning**: Education → Knowledge Capture → Reflection

---

## 독자 (Audience)

- **주 독자:** 소프트웨어 개발자(초급 개발자 또는 대학생)
- **함께 읽을 독자:** AI 도구를 업무에 활용하는 모든 지식 노동자
- **사전 지식:** LLM 사용 경험, 자신의 분야에서 문서·산출물을 만든 경험
- **읽기 목적:** Cocrates Agent의 구조와 각 스킬의 원리를 이해하고, 이를 자신의 AI 활용에 적용하는 방법을 익히는 것

---

## 로그라인 (Logline)

같은 LLM을 쓰는데도 누군가는 조수(AI-assisted)로, 누군가는 수십 명의 AI 팀을 이끄는 총괄 디렉터(AI-native)로 활용한다. 그 차이는 모델이 아니라 **사용자의 태도와 역량**에 있다. Cocrates는 사용자가 AI를 올바른 방법으로 사용할 수 있도록 워크플로우를 안내하고, AI-native Engineer로 성장하도록 돕는 에이전트 하네스다.

---

## 시리즈 구조

| 구간 | 주제 | 내용 |
|------|----|------|
| **1. 소개** | Cocrates Harness의 필요성과 원칙 | [Ep1~2] 문제 인식 — AI-assisted vs AI-native의 차이, 핵심 원칙 선언 |
| **2. 실습** | 먼저 경험한다 | [Ep3~6] 설치 → 학습 실습 → 생성 실습 → 스킬 생성 실습 |
| **3. 구조** | 원리를 이해한다 | [Ep7~12] 시스템 구조, 학습 및 산출물 생성 활동과 스킬 |
| **4. 결론** | 지속적 개선과 선언 | [Ep13~14] 진화, 사용자 선언문 |

### Part 1. 소개 — Cocrates Harness의 필요성과 원칙

"같은 LLM, 같은 요청인데 왜 결과가 다를까?" — 이 질문에서 시작한다. AI-assisted Engineer(도구로 사용)와 AI-native Engineer(팀원으로 활용)의 차이를 명확히 구분하고, "검토되지 않은 산출물은 생성할 가치가 없다"는 핵심 원칙을 선언한다.

| 편 | 제목 | 핵심 질문 | 역할 |
|----|------|-----------|------|
| **1** | 같은 LLM, 다른 결과 | "당신의 AI는 몇 명의 팀원인가?" | 문제 인식 — 사용자 태도의 차이가 결과를 만든다 |
| **2** | The Unexamined Code Is Not Worth Generating | "검토(U→A→E→A)의 진짜 의미는 무엇인가?" | 핵심 원칙 선언 — 검토의 본질은 무지의 제거 |

### Part 2. 실습 — 먼저 경험한다

설치부터 시작해 Cocrates의 두 가지 핵심 활동(학습, 생성)과 워크플로우 생성을 직접 경험한다. "경험 먼저, 원리 나중" — 실습을 통해 Cocrates의 작동 방식을 몸으로 익힌다.

| 편 | 제목 | 핵심 경험 | 역할 |
|----|------|-----------|------|
| **3** | Cocrates Harness 설치 | opencode + Cocrates Plugin + Skills | AI-native Engineer의 첫걸음 |
| **4** | 소크라테스식 학습 활동 - 실습 | Education → Knowledge Capture → Reflection | Learning 파이프라인 직접 체험 |
| **5** | 구조 기반 산출물 생성 활동 - 실습 | ADR → Spec → Generation → Verification | Artifact Generation 파이프라인 직접 체험 |
| **6** | 구조 기반 워크플로우 생성 실습 | Kernel → Frame → Outline → Spec → Skill → Verification | 스킬 생성 직접 체험 |

### Part 3. 구조 — 원리를 이해한다

"사용"에서 "이해"로 전환. Cocrates Harness의 Agent + Skills 구조와 각 활동의 개념적 토대, 스킬의 핵심 원칙을 깊이 탐구한다.

| 편 | 제목 | 핵심 내용 | 역할 |
|----|------|-----------|------|
| **7** | Cocrates Harness 구조 | Agent + Skills, 6개 섹션 프롬프트, Intent-To-Skill Routing | 전체 시스템 조감도 |
| **8** | 소크라테스식 학습 활동 - 원리 | Socratic Maieutics, Bloom's Taxonomy, ZPD, Learning Pipeline | 학습 활동의 철학적·교육학적 기반 |
| **9** | 소크라테스식 학습 활동 - 스킬 | education, knowledge-capture, reflection 스킬 핵심 | 학습 스킬의 내부 구조 |
| **10** | 구조 기반 산출물 생성 활동 - 원리 | ASR, ADR, Spec, Snowflake, Verification as Gate | 생성 활동의 개념적 기반 |
| **11** | 구조 기반 산출물 생성 활동 - 스킬 | adr-writing, spec-writing, spec-driven-generation, spec-driven-verification | 생성 스킬의 내부 구조와 원칙 |
| **12** | 구조 기반 워크플로우 생성 스킬 | generating-skill-creation, Meta Snowflake 7단계 | 메타-스킬의 원리 |

### Part 4. 결론 — 지속적 개선과 선언

시리즈의 마무리. 사용자도 Cocrates Harness도 함께 진화해야 함을 선언하고, 사용자 선언문(7계명)으로 마무리한다.

| 편 | 제목 | 핵심 내용 | 역할 |
|----|------|-----------|------|
| **13** | 올바른 Cocrates Harness 활용 | 사용자와 에이전트의 상호 진화, 오픈소스 | 지속적 발전의 원칙 |
| **14** | Cocrates Harness 사용자 선언문 | 7가지 계명 — AI 주권 선언 | 시리즈의 마무리와 실천 다짐 |

---

## 시리즈 아크 (Series Arc)

```
Part 1. 소개 ── "그래, 맞아. 근데 어떻게 하지?"     (문제 인식 + 원칙)
    ↓
Part 2. 실습 ── "직접 해보니 궁금해지는구나."        (경험을 통한 이해)
    ↓
Part 3. 구조 ── "아, 실습에서 경험한 것이            (원리와 시스템 이해)
                이런 원리였구나."
    ↓
Part 4. 결론 ── "이제 내가 할 일을 알겠다."          (실천과 다짐)
```

---

## 요약 (Summary)

같은 LLM을 두고 어떤 사람은 단순한 코드 생성 조수(AI-assisted)로 쓰고, 어떤 사람은 수십 명의 AI 개발팀을 지휘하는 총괄 디렉터(AI-native)처럼 활용한다. 이 차이는 모델의 성능이 아니라 사용자의 **태도(AI를 도구로 보는가, 팀원으로 보는가)** 와 **검토 역량(U→A→E→A)** 에서 비롯된다.

Cocrates(Co+Socrates)는 이 역량을 키우기 위해 설계된 에이전트 하네스다. **상호 소크라테스 관계**를 통해 사용자가 AI를 주도적으로 사용(리딩)하게 하고, 아키텍처 기반 설계·검토·승인 과정을 거쳐 결과물을 완벽하게 이해하도록 안내한다. 궁극적 목표는 **무지의 제거(Harnessing Ignorance)** — 내가 무엇을 모르는지 알고 채워가는 성장의 과정이다.

이 시리즈는 Cocrates의 필요성과 원칙(소개)에서 시작해, 먼저 설치와 실습을 통해 경험한 뒤(실습), 시스템 구조와 각 활동의 원리를 이해하고(구조), 지속적 개선의 다짐으로 마무리한다(결론).

---

## 제약 및 맥락

| 항목 | 내용 |
|------|------|
| **에피소드 수** | 14편 (소개 2 + 실습 4 + 구조 6 + 결론 2) |
| **편당 분량** | 읽기 8~15분 (한국어 2,500~4,500자 목표) |
| **서사** | 소개(필요성·원칙) → 실습(설치·경험) → 구조(시스템·원리) → 결론(진화·선언) |
| **톤** | 부드러운 대화체. 개념의 깊이는 유지하되 누구나 이해할 수 있도록. |
| **공통 패턴** | 모든 활동(학습, 생성, 스킬)은 **Snowflake 점진적 확장** — 작게 시작, 단계별 확장, 검토와 승인, 파일 저장 |
| **범위 안** | Cocrates 구조, 학습 활동(개념/스킬), 산출물 생성 활동(개념/스킬), 설치·실습·지속적 개선·사용자 선언 |
| **범위 밖** | 일반론적 하네스 개념, 모델 비교, 타 에이전트 프레임워크와의 비교 |
| **핵심 메시지** | ① 검토되지 않은 산출물은 생성할 가치가 없다 ② 검토 = 이해(U) + 분석(A) + 평가(E) + 승인(A) ③ Cocrates는 상호 소크라테스 관계로 AI-native Engineer의 성장을 돕는다 ④ 사용자도 에이전트도 함께 진화해야 한다 |
