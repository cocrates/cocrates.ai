# Slide Plan — Episode 09: 소크라테스식 학습 활동 - 스킬

**Total slides:** 13
**Estimated pace:** ~1분/슬라이드 (총 ~13분)
**Template:** 타이틀 → 학습목표 → 콘텐츠(3개 스킬 내부) → 학습 확인 → 다음편연결

---

## 도입 (Slides 1-2) — 약 2분

### Slide 1: Title — 소크라테스식 학습 활동 - 스킬
- **Section:** 도입
- **Purpose/goal:** Ep8(원리 이해)에서 Ep9(스킬 내부 구현)로의 전환을 알린다. "Ep4에서 Cocrates가 했던 그 질문들 — 누군가가 미리 스크립트를 짜놓은 걸까?"라는 궁금증을 자극한다.
- **Content:**
  - **주제:** 소크라테스식 학습 활동 — 스킬
  - **부제:** 3블록 미션과 무지(Ignorance)를 박제하는 코드 엔진
  - **에피소드 연결:** Ep8에서 Learning 파이프라인의 3대 철학과 3단계 구조 이해 → Ep9는 그 원리가 실제 `.opencode/skills/`에서 **어떤 코드 구조와 워크플로우 규칙**으로 구현되었는지 내부 탐험
  - **핵심 전환 메시지:**
    - "Ep4에서 Cocrates가 던졌던 그 날카로운 질문들, 기억하시나요?"
    - "누군가가 미리 시나리오를 짜놓은 것처럼 정교했죠. 오늘은 그 비밀을 낱낱이 공개합니다."
- **Type:** 전환/타이틀

### Slide 2: 학습 목표
- **Section:** 도입
- **Purpose/goal:** 시청자가 이 발표를 통해 얻게 될 3가지 핵심 목표를 제시한다. Ep8의 원리가 실제 코드로 어떻게 굳어졌는지에 대한 호기심을 자극한다.
- **Content:**
  - **도입 질문:** "Ep8에서 Learning 파이프라인의 원리를 배웠습니다. 그런데 이 원칙들이 실제로 어떻게 '코드'로 굳어져 있을까요? 오늘은 그 내부를 직접 들여다봅니다."
  - **3가지 학습 목표:**
    1. 🧗‍♂️ Education 스킬의 No Spoon-feeding 원칙, Turn-based Mission 구조, 그리고 3블록 출력 포맷
    2. 💾 Knowledge Capture 스킬의 Ignorance 기록 철학과 Merge 전략
    3. 🕵️‍♂️ Reflection 스킬의 Interviewer 페르소나와 Gap 발견 시 "가르치지 않는" 원칙
- **Type:** 학습 목표

---

## 본론 (Slides 3-11) — 약 11분

### Slide 3: 🧗‍♂️ Education Skill — No Spoon-feeding과 Turn-based Mission
- **Section:** 본론 — Education 스킬: 핵심 규칙 (2.1)
- **Purpose/goal:** Education 스킬의 두 가지 핵심 규칙 — No Spoon-feeding(절대 떠먹여 주지 않음)과 Turn-based Mission(턴제 미션) — 에 집중한다. 이 두 규칙이 Cocrates 교육 대화의 근간임을 전달한다.
- **Content:**
  - **규칙 1 — No Spoon-feeding (최상위 헌법):**
    - "절대 한 턴에 완전한 정답이나 전체 해결책을 떠먹여 주지 않는다"
    - 개념 설명은 최대 **1~3문장, 응답의 20% 이하**로 제한
    - **Incomplete State 원칙:** 대화를 항상 불완전한 상태로 멈춤 → 사용자가 직접 빈칸을 채워야 함
  - **규칙 2 — Turn-based Mission:**
    - 모든 응답은 정확히 **하나의 `🔥 [MISSION]` 블록**으로 끝나야 함
    - 사용자가 현재 MISSION에 답하기 전에 절대 다음 단계의 힌트를 **스포일러하지 않음**
    - MISSION은 단순한 사실 확인이 아닌 **인지적 과제**여야 함 (적용, 분석, 평가 등)
  - **두 규칙의 관계:** No Spoon-feeding이 '무엇을 하지 말아야 하는지'를 정한다면, Turn-based Mission은 '무엇을 해야 하는지'를 정한다. **하나는 제한, 하나는 생성 — 두 규칙이 하나의 엔진을 이룬다.**
- **Type:** 개념/원칙

### Slide 4: 📊 Bloom's 2D Matrix + Push/Pull 전략
- **Section:** 본론 — Education 스킬: 난이도 조절 (2.2)
- **Purpose/goal:** Education 스킬이 Bloom's 2D 매트릭스와 Push/Pull 전략을 어떻게 활용하여 사용자 수준에 맞게 난이도를 동적으로 조절하는지 설명한다. "똑같은 질문을 모든 사용자에게 던지는 것이 아니다"라는 점을 강조한다.
- **Content:**
  - **Bloom's 2D Matrix 활용:**
    - Y축(인지 과정): Remember → Understand → Apply → Analyze → Evaluate → Create
    - X축(지식 차원): Factual → Conceptual → Procedural → Metacognitive
    - **핵심 규칙 — 동시 격상 금지:** 두 축을 동시에 올리지 않음
  - **Push/Pull 전략:**
    - **Pull (기본값):** 높은 수준의 도전을 먼저 던져 사용자가 지식을 끌어당기게 함
    - **Push:** 사용자에게 인지적 붕괴나 혼란이 감지되면 즉시 단계를 낮춰 징검다리 제공
    - **판단 규칙:** 기본 Pull, 혼란 시 Push로 전환 후 한 축만 낮춤
  - **Ep4 연결:** Cocrates가 사용자의 답변에 따라 다음 미션의 난이도를 조절했던 경험
  - **💡 포인트:** 단순한 퀴즈가 아니다. **사용자의 실시간 반응에 적응하는 지능형 난이도 조절 엔진**이다.
- **Type:** 개념/상세

### Slide 5: 🔥 3블록 출력 포맷 + DIP 실제 사례
- **Section:** 본론 — Education 스킬: 3블록 + 실제 사례 (2.3)
- **Purpose/goal:** Education 스킬의 3블록 출력 포맷(Concept Briefing → Thought Lab → MISSION)을 실제 SKILL.md에 정의된 DIP(Dependency Inversion Principle) 예시를 통해 구체적으로 보여준다. 추상적인 규칙이 실제 대화에서 어떻게 형체를 갖추는지 실감나게 전달한다.
- **Content:**
  - **3블록 출력 포맷 (SKILL.md 실제 규칙):**
    ```markdown
    ### 💡 [Concept Briefing]
    (핵심 원리를 1~3문장, 일상적 비유로 전달. 응답의 20% 이하)

    ### 💻 [Thought Lab]
    (일부러 결함을 심어두었거나 비어있는 불완전한 예시/시나리오)

    ### 🔥 [MISSION]
    (사용자가 다음 턴에 생각하고 답해야 할 정확히 하나의 인지적 과제)
    ```
  - **각 블록의 역할:**
    - **Briefing:** 최소한의 맥락 — "이게 왜 중요한가"만 전달 (비유 사용)
    - **Lab:** 결함 있는 예시 — 사용자가 직접 오류를 발견하고 수정하도록 유도
    - **MISSION:** 정확히 하나의 인지적 과제 — 절대 2개 이상 주지 않음
  - **실제 예시 — DIP 원칙 설명 요청:**
    - **사용자:** "DIP 원칙에 대해 알려줘"
    - **Cocrates의 3블록 응답:**
      - 💡 **Briefing:** "DIP는 전기 플러그와 같습니다. 고수준 모듈이 저수준 세부사항에 직결되지 않고, 둘 다 공유 인터페이스에 의존해야 합니다."
      - 💻 **Thought Lab:**
        ```typescript
        class OrderService {
          constructor(private db: MySQLDatabase) {} // flaw
          createOrder() { /* ... */ }
        }
        ```
      - 🔥 **MISSION:** "자신의 언어로, `MySQLDatabase`를 `PostgresDatabase`로 바꾸면 무엇이 깨질지 설명해보세요. 그리고 DIP를 위반하는 의존 방향이 어디인지 지적해보세요."
  - **💡 포인트:** 이 3블록 구조가 Cocrates의 모든 교육 대화를 생성하는 **단일한 엔진**이다. Briefing이 없으면 사용자는 맥락을 모르고, Lab이 없으면 생각할 재료가 없고, MISSION이 없으면 수동적으로 끝난다. **세 블록이 하나의 완전한 학습 턴을 이룬다.**
- **Type:** 개념/코드 + 사례

### Slide 6: 💾 Knowledge Capture Skill — 핵심만 캡처한다
- **Section:** 본론 — Knowledge Capture 스킬 (2.4)
- **Purpose/goal:** Knowledge Capture 스킬의 핵심 원칙 — "장황한 설명 금지, 회상 가능한 최소 단위만 저장"을 전달한다. 저장 위치와 저장 대상을 구체적으로 설명한다.
- **Content:**
  - **핵심 원칙:** 백과사전식 장황한 설명이나 전체 코드는 **일절 저장하지 않음**
  - **저장 위치:** `kb/{topic-slug}.md` (kebab-case, 영어)
  - **저장 대상:**
    - **한 줄 정의** — 추후 회상을 위한 최소 단위
    - **비유** — 개념을 내 것으로 만드는 이미지
    - ✅ **잘 이해한 것**
    - ⚠️ **"내가 처음에 착각했던 오개념과 무지(Ignorance)"** ← `## Wrong Assumptions / Gaps` 섹션
  - **💡 포인트:** "기억하고 있는 것"을 기록하는 것이 아니라 **"내가 몰랐던 것"을 기록하는 것** — Knowledge Capture의 가장 중요한 철학
- **Type:** 개념/개관

### Slide 7: ⚠️ 무지(Ignorance)의 기록 — 오개념 박제의 힘
- **Section:** 본론 — Knowledge Capture: Ignorance 기록 (2.5)
- **Purpose/goal:** Knowledge Capture의 백미인 `## Wrong Assumptions / Gaps` 섹션의 철학과 효과를 깊이 있게 전달한다. "무지를 기록하는 것"이 왜 장기 기억에 가장 강력한 도구인지 이해시킨다.
- **Content:**
  - **철학:** "무엇을 새로 배웠는지"보다 **"어떤 오개념을 가지고 있었고 어떻게 틀렸었는지"** 를 박제하는 것이 장기 기억의 핵심
  - **KB 실제 예시 (Bloom's taxonomy):**
    ```markdown
    ## Wrong Assumptions / Gaps
    - (이전 오해) 블룸 분류학 = 기억→이해→...→창조의 순차적 피라미드
    - (깨달음) 사실은 2차원 매트릭스 — 지식 차원 × 인지 과정의 교차점
    ```
  - **왜 효과적인가?** 뇌는 새로운 정보보다 **깨진 예상(Broken Expectation)** 을 더 오래 기억함. "내가 틀렸다"는 사실이 기억을 강화함
  - **💡 포인트:** "아는 것"을 기록하는 KB는 흔하다. **"틀렸던 것"을 기록하는 KB가 Cocrates만의 차별점이다.**
- **Type:** 개념/심화

### Slide 8: 🔄 Merge 전략 — 지식 중복 방지
- **Section:** 본론 — Knowledge Capture: Merge (2.6)
- **Purpose/goal:** Knowledge Capture 스킬이 새로운 KB 파일을 무분별하게 양산하지 않고, 기존 지식과 병합하는 전략을 설명한다. 지식의 일관성과 누적성을 강조한다.
- **Content:**
  - **문제:** 매번 새로운 파일을 만들면 KB가 수천 개로 불어남 → 회상 불가능
  - **해결 — Merge 전략:**
    1. 저장 요청이 들어오면 동일 주제의 기존 KB 파일 검색 (파일명, 제목, Tags 비교)
    2. 발견 시 **덮어쓰지 않음** — 새로운 인사이트만 `## Update History`와 함께 추가
    3. 미발견 시 새 파일 생성
  - **병합 예시:**
    ```markdown
    ## Update History
    - 2026-06-26: Push/Pull 전략의 판단 규칙 추가
    - 2026-06-20: 최초 작성 (3블록 구조)
    ```
  - **💡 포인트:** 지식은 **교체**가 아니라 **누적**되어야 한다. Merge 전략은 지식의 연속성을 보장한다.
- **Type:** 개념/상세

### Slide 9: 🕵️‍♂️ Reflection Skill — 면접관의 등장
- **Section:** 본론 — Reflection 스킬 (2.7)
- **Purpose/goal:** Reflection 스킬의 페르소나 전환(코치 → 면접관)과 Education과의 근본적인 차이를 전달한다. Reflection의 목적이 채점이 아니라 "정확한 지식 지도"를 그리는 것임을 강조한다.
- **Content:**
  - **페르소나 전환:** "평가해줘" / "시험해줘" 요청 시 친절한 코치 → **엄격한 면접관(Interviewer)**으로 돌변
  - **Education vs Reflection 비교:**
    | | Education | Reflection |
    |---|---|---|
    | **역할** | 코치 — 이해를 구축 | 면접관 — 이해를 측정 |
    | **목표** | 사용자를 새로운 통찰로 안내 | 사용자의 지식 지도를 함께 그림 |
    | **전략** | Push/Pull | **Pull only** (Push 금지) |
    | **Gap 발견 시** | 즉시 비계 제공 | 기록만 하고 Education으로 제안 |
  - **💡 포인트:** 같은 Bloom's 2D 매트릭스를 사용하지만 **목적이 다르면 전략도 완전히 달라진다.**
- **Type:** 개념/비교

### Slide 10: 🎯 Pull-only 평가 — 가르치지 않고 측정한다
- **Section:** 본론 — Reflection: Pull-only (2.8)
- **Purpose/goal:** Reflection이 왜 Push(가르치기)를 절대 사용하지 않고 Pull-only로 평가하는지, 그 이유와 방법을 설명한다. Gap 발견 시의 행동 원칙인 "가르치지 않는다"의 의미를 전달한다.
- **Content:**
  - **왜 Pull-only인가?** Reflection의 목적은 **성장**이 아니라 **측정**. 가르치기 시작하면 현재 수준이 왜곡됨
  - **질문 유형 (실제 SKILL.md 규칙에서 발췌):**
    - "이 원칙을 완전히 다른 도메인의 예시를 들어 설명해보세요" (전이)
    - "이 시나리오에서 원칙이 깨지는 경계 조건은 어디인가요?" (경계)
    - "KB에 기록된 과거의 그 틀렸던 가정, 지금도 똑같이 실수하고 있지 않나요?" (메타인지)
  - **Gap 발견 시 행동 원칙:**
    - 그 자리에서 **강의하지 않음**
    - 발견된 공백은 `⚠️ 제대로 알지 못했던 것` 리스트로 기록만
    - 세션 마무리 시 "이 부분에 대해 별도의 Education을 따로 진행하시겠습니까?" 정중한 제안
  - **💡 포인트:** "모르는 걸 가르쳐주는 것"이 친절처럼 보이지만, Reflection의 순수성을 해친다. **측정과 교수는 분리되어야 한다.**
- **Type:** 개념/원칙

### Slide 11: 🔗 세 스킬의 유기적 연결
- **Section:** 본론 — 스킬 연결 (2.9)
- **Purpose/goal:** Education, Knowledge Capture, Reflection 세 스킬이 단순히 나열된 것이 아니라, 하나의 Learning 파이프라인으로 유기적으로 연결되어 있음을 종합한다. 특히 같은 Bloom's 2D 매트릭스를 공유하지만 사용 목적이 다른 점을 강조한다.
- **Content:**
  - **파이프라인 흐름:**
    ```
    [Education: 탐구] → [Knowledge Capture: 기록] → [Reflection: 검증]
         ↑                                                    │
         └────────── (Gap 발견 시 Education 환류 제안) ────────┘
    ```
  - **같은 도구, 다른 목적:**
    - **Education:** Bloom's 2D 매트릭스로 **Push/Pull** → 사용자 **성장**
    - **Reflection:** Bloom's 2D 매트릭스로 **Pull-only** → 이해도 **측정**
  - **KB의 가교 역할:** Education에서 배운 내용이 KB에 저장되고, Reflection이 이 KB를 평가 루브릭으로 사용
  - **💡 포인트:** 세 스킬은 독립적으로 동작하지만, **같은 언어(Bloom's 2D 매트릭스)로 소통하는 하나의 시스템**이다.
- **Type:** 개념/종합

---

## 결론 (Slides 12-13) — 약 2분

### Slide 12: 핵심 요약 + 학습 확인
- **Section:** 결론
- **Purpose/goal:** 발표의 핵심을 3가지 포인트로 요약하고, 시청자가 스스로 이해를 확인할 수 있도록 자기 질문을 제시한다.
- **Content:**
  - **3-Point 핵심 요약:**
    1. **🧗‍♂️ Education 스킬:** No Spoon-feeding이 최상위 헌법. 3블록 엔진(Briefing → Lab → MISSION)으로 사용자를 불완전한 상태에 묶어둔다. Push/Pull로 난이도를 동적 조절.
    2. **💾 Knowledge Capture 스킬:** "무엇을 배웠는지"보다 **"무엇을 오해했는지(Wrong Assumptions/Gaps)"** 를 기록한다. Merge 전략으로 지식의 연속성을 보장.
    3. **🕵️‍♂️ Reflection 스킬:** 가르치지 않고 측정만 한다(Pull-only). Gap 발견 시 기록만 하고, Education 환류를 제안한다.
  - **학습 확인 (자기 질문):**
    - ❓ Education의 No Spoon-feeding과 Turn-based Mission을 설명할 수 있는가?
    - ❓ 3블록 출력 포맷(Concept Briefing → Thought Lab → MISSION)의 각 역할은?
    - ❓ Reflection이 Gap 발견 시 왜 가르치지 않고 기록만 하는가?
- **Type:** 요약

### Slide 13: 다음 에피소드 연결
- **Section:** 결론
- **Purpose/goal:** Learning 파이프라인(원리 + 스킬 내부) 완전 정복을 선언하고, 다음 단계(Ep10: 구조 기반 산출물 생성 - 원리)에 대한 기대감을 형성한다.
- **Content:**
  - **지금까지의 여정:**
    - Ep7: Cocrates 전체 구조 이해
    - Ep8: Learning 파이프라인 원리 이해
    - **Ep9: Learning 파이프라인 스킬 내부 탐험 ✅ (완료)**
  - **다음 편 (Ep10): 구조 기반 산출물 생성 활동 - 원리**
    - Cocrates의 두 번째 Core Activity!
    - "당장 보고서 한 장 써내라"는 사용자의 요구에 Cocrates는 왜 **생성 버튼을 누르지 않고 아키텍처(ADR)부터 설계하자고 빗장을 걸어 잠그는가?**
    - ASR → ADR → Spec → 생성 → 검증 — 구조 기반 생성의 전 과정 개념 이해
  - **마무리 인용:**
    > "답을 아는 것과, 그 답이 무너지지 않도록 구조화하는 것은 완전히 다른 차원의 도달입니다."
  - 🦉
- **Type:** 전환/마무리
