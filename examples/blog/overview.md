# AI를 올바르게 쓰는 법 — Cocrates Harness

## 핵심 원칙 (Core Principle)

> **The unexamined code is not worth generating.**
> (검토되지 않은 산출물은 생성할 가치가 없다.)

**code**란 AI를 사용해서 만들고자 하는 **최종 산출물**을 의미한다. '검토(examine)'의 본질은 **이해하고, 판단하고, 승인하는 것**이다. 검토의 목적은 **무지의 제거(Harnessing Ignorance)** 에 있다. AI가 구조를 설계해도 내가 이해하지 못했다면 그것은 검토된 것이 아니며, 내게 무지만 남긴다.

Cocrates는 이 원칙을 실현하는 하네스다. 사용자가 AI를 올바른 방법으로 사용할 수 있도록 워크플로우를 안내하고, 모든 단계에서 사용자의 이해와 승인을 요구한다.

## 용어 정의

이 시리즈에서 사용하는 핵심 용어를 먼저 명확히 정의한다.

| 용어 | 의미 | 비고 |
|------|------|------|
| **Cocrates** | 전체 시스템을 지칭하는 이름. "Cocrates는 무엇인가?"라는 질문에 대한 답은 아래 Cocrates Harness와 같다. | |
| **Cocrates Harness** | 전체 시스템 = **Cocrates Agent + Skills** | 원문: *"Cocrates is a **Cocrates agent plus skills** harness."* |
| **Cocrates Agent** | 핵심 에이전트. `cocrates.md` 파일로 정의되며, 원칙·의도 인식·스킬 선택·태스크 관리·가드레일을 담당한다. | "You are **Cocrates**"의 주체 |
| **Skills** | 작업별 절차. `.opencode/skills/*/SKILL.md`로 관리되며 추가·수정·확장 가능하다. | |

즉, **Cocrates Harness = Cocrates Agent + Skills**이며, 이 시리즈의 **3부(구조)** 는 바로 이 Cocrates Harness를 소개한다.

### Cocrates Agent의 역할

Cocrates Agent (`cocrates.md`)는 두 가지 **Core Activity** 파이프라인을 정의한다:
- **Artifact Generation**: 설계(ADR → Spec) → Spec 기반 생성 → Spec 기반 검증
- **Learning**: Education → Knowledge Capture → Reflection

## 독자 (Audience)

- **주 독자:** 소프트웨어 개발자(초급 개발자 또는 대학생)
- **함께 읽을 독자:** AI 도구를 업무에 활용하는 모든 지식 노동자
- **사전 지식:** LLM 사용 경험, 자신의 분야에서 문서·산출물을 만든 경험
- **읽기 목적:** Cocrates Agent의 구조와 각 스킬의 원리를 이해하고, 이를 자신의 AI 활용에 적용하는 방법을 익히는 것

## 로그라인 (Logline)

같은 LLM을 쓰는데도 누군가는 조수로, 누군가는 수십 명의 AI 팀을 이끄는 총괄 디렉터로 활용한다. 그 차이는 모델이 아니라 **사용자의 능력**에 있다.
Cocrates는 사용자가 AI를 올바른 방법으로 사용할 수 있도록 워크플로우를 안내하는 에이전트다.

## 시리즈 구조

| 구간 | 주제 | 내용 |
|------|----|------|
| 1. 소개 | Cocrates Harness의 필요성과 원칙 | [Ep1~2] 문제 인식과 핵심 원칙 선언 |
| 2. 실습 | 먼저 경험한다 | [Ep3~6] 설치 → 학습 실습 → 생성 실습 → 스킬 생성 실습 |
| 3. 구조 | 원리를 이해한다 | [Ep7~12] 시스템 구조, 학습 및 산출물 생성 활동과 스킬 |
| 4. 결론 | 지속적 개선과 선언 | [Ep13~14] 진화, 사용자 선언문 |

| 구간 | 편 | 내용 | 역할 |
|------|----|------|------|
| **1. 소개** | 1 | 같은 LLM, 다른 결과 | 문제 인식 — 사용자 태도의 차이가 결과를 만든다 |
| | 2 | The Unexamined Code Is Not Worth Generating | 핵심 원칙 선언 — 검토의 본질은 무지의 제거 |
| **2. 실습** | 3 | Cocrates Harness 설치 | opencode/cocrates harness plugin 설치, 환경 설정, 기본 동작 확인 |
| | 4 | 소크라테스식 학습 활동 - 실습 | Education → Knowledge Capture → Reflection 실제 워크플로우 경험 |
| | 5 | 구조 기반 산출물 생성 활동 - 실습 | adr → spec → generation/verification 실제 워크플로우 경험 |
| | 6 | 구조 기반 워크플로우 생성 실습 | generating-skill-creation 스킬 실제 사용 |
| **3. 구조** | 7 | Cocrates Harness 구조 | Cocrates Agent + Skills 구조, Cocrates Agent Prompt |
| | 8 | 소크라테스식 학습 활동 | 능동적 학습을 위한 활동 및 교육 철학, 연결된 스킬 개관 |
| | 9 | 소크라테스식 학습 활동 - 스킬 | Education → Knowledge Capture → Reflection 스킬의 핵심 내용 |
| | 10 | 구조 기반 산출물 생성 활동 | 구조적 의사결정(ADR)과 Spec, 연결된 스킬 개관 |
| | 11 | 구조 기반 산출물 생성 활동 - 스킬 | adr → spec → generation/verification 스킬의 핵심 내용 |
| | 12 | 구조 기반 워크플로우 생성 스킬 | generating-skill-creation 스킬의 핵심 내용 |
| **4. 결론** | 13 | 올바른 Cocrates Harness 활용 | Cocrates Harness를 진화시켜야 한다 — 사용자도 에이전트도 지속적 개선 |
| | 14 | Cocrates Harness 사용자 선언문 | 끊임없이 개선하겠다, 포기하지 않겠다는 사용자의 선언 |

## 요약 (Summary)

같은 LLM을 두고 어떤 사람은 단순한 코드 생성 조수로 쓰고, 어떤 사람은 수십 명의 AI 개발팀을 지휘하는 총괄 디렉터처럼 활용한다. 이 차이는 모델의 성능이 아니라 사용자의 역량에서 비롯된다.

Cocrates는 이 역량을 키우기 위해 설계된 에이전트 하네스다. 불확실성을 체계적인 탐구로 전환하고, 사용자가 아키텍처 기반 설계, 검토 및 승인 과정을 거쳐 결과물을 완벽하게 이해할 수 있도록 안내한다.

이 시리즈는 Cocrates의 필요성과 원칙(소개)에서 시작해, 먼저 설치와 실습을 통해 경험한 뒤(실습), 시스템 구조와 각 활동의 원리를 이해하고(구조), 지속적 개선의 다짐으로 마무리한다(결론).

## 제약 및 맥락

| 항목 | 내용 |
|------|------|
| **에피소드 수** | 14편 (소개 2 + 실습 4 + 구조 6 + 결론 2) |
| **편당 분량** | 읽기 8~15분 (한국어 2,500~4,500자 목표) |
| **서사** | 소개(필요성·원칙) → 실습(설치·경험) → 구조(시스템·원리) → 결론(진화·선언) |
| **톤** | 부드러운 대화체. 개념의 깊이는 유지하되 누구나 이해할 수 있도록. |
| **범위 안** | Cocrates 구조, 학습 활동(개념/스킬), 산출물 생성 활동(개념/스킬), 설치·실습·지속적 개선·사용자 선언 |
| **범위 밖** | 일반론적 하네스 개념, 모델 비교, 타 에이전트 프레임워크와의 비교 |
| **핵심 메시지** | ① 검토되지 않은 산출물은 생성할 가치가 없다 ② 검토 = 이해 + 판단 + 승인 ③ Cocrates는 이를 위한 워크플로우를 안내한다 ④ 사용자도 에이전트도 함께 진화해야 한다 |
