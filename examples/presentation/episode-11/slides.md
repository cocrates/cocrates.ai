# Slide Plan — Episode 11: 구조 기반 산출물 생성 활동 - 스킬

**Total slides:** 11
**Estimated pace:** ~1.5분/슬라이드 (총 ~17분)
**Template:** 타이틀 → 학습목표 → 전체 흐름(2) → 4개 스킬(4) → 연결과 요약(1) → 통찰+학습확인(1) → 다음편연결(1)

---

## 도입 (Slides 1-2) — 약 2분

### Slide 1: Title — 구조 기반 산출물 생성 활동 - 스킬
- **Section:** 도입
- **Purpose/goal:** Ep10(원리) → Ep11(스킬) 전환을 명확히 하고, 이제 파이프라인이 어떻게 실제로 동작하는지 들어간다는 기대감을 형성한다.
- **Content:**
  - **주제:** 구조 기반 산출물 생성 활동 - 스킬
  - **부제:** 스킬이 없으면? spec-driven-generation이 구조적으로 산출물을 생성하게 한다
  - **에피소드 연결:** Ep10 파이프라인 원리 ✅ → **Ep11 4개 스킬의 내부 메커니즘**
  - **핵심 전환 메시지:** "Ep10에서 우리는 Generation 파이프라인의 '왜(Why)'를 배웠습니다. 이제 '어떻게(How)' — 실제로 이 파이프라인을 구동하는 4개 스킬의 내부로 들어갑니다."
- **Type:** 전환/타이틀

### Slide 2: 학습 목표
- **Section:** 도입
- **Purpose/goal:** 청중이 오늘 발표를 통해 얻게 될 3가지 핵심 목표를 제시한다.
- **Content:**
  - **도입 질문:** "Ep10에서 배운 파이프라인 원리 — 이게 실제로 코드로는 어떻게 구현되어 있을까요?"
  - **3가지 학습 목표:**
    1. 🔄 Cocrates가 요청을 받으면 어떤 스킬을 쓸지 **라우팅**하는가
    2. 🚦 **spec-driven-generation**이 spec이 없을 때 어떻게 **spec 작성을 유도**하는가
    3. 🛠️ 4개 스킬 각각의 **핵심 원칙과 절대 금지 사항**
- **Type:** 학습 목표

---

## 본론 1: 전체 흐름 (Slides 3-4) — 약 3분

### Slide 3: 🔄 Intent-to-Skill Routing — 요청이 스킬이 되는 순간
- **Section:** 본론 1 — 전체 흐름
- **Purpose/goal:** 사용자의 "생성해줘" 요청이 Cocrates 내부에서 어떻게 처리되는지 첫 단계를 이해시킨다. 스킬 유무가 첫 번째 분기점임을 전달한다.
- **Content:**
  - **문제:** 사용자가 "A를 생성해줘"라고 하면, Cocrates는 무엇부터 할까?
  - **라우팅 규칙:**
    1. 최종 산출물 유형(final deliverable type) 식별
    2. 해당 유형에 매칭되는 생성 스킬 검색 (가장 구체적인 타입부터)
    3. 있음 → 해당 스킬 사용 (구조가 이미 스킬에 명세됨)
    4. 없음 → **spec-driven-generation으로 폴백**
  - **비유:** 전문 병원 vs 응급실. 전문과목(스킬)이 있으면 전문의 진료, 없으면 응급실(spec-driven-generation)에서 기본적인 처치.
  - **💡 포인트:** 모든 요청이 spec-driven-generation을 거치는 것은 아니다 — **스킬이 없을 때만 fallback**
- **Type:** 개념/비유

### Slide 4: 🚦 spec-driven-generation Fallback — "spec이 없으면 spec을 만든다"
- **Section:** 본론 1 — 전체 흐름
- **Purpose/goal:** spec-driven-generation이 단순 생성기가 아니라, Spec Readiness Gate를 통해 spec이 없으면 spec 작성을 유도하는 메타-프로세스임을 전달한다.
- **Content:**
  - **spec-driven-generation의 이중 역할:**
    1. **Spec Readiness Gate:** Spec이 준비되었는지 확인
    2. **생성:** 준비된 Spec 기준으로 산출물 생성
  - **Gate 통과 실패 시 (Spec 없음):**
    ```
    ASR 식별 → (대안 있으면) ADR → 사용자 결정 → Spec 작성 → 생성
    ```
  - **8가지 보편 ASR 범주 개관:** 목적/독자, 형태, 범위, 품질, 제약, 구조, 의존성, 프로세스
  - **전원주택 비유 연결:** 건축주가 "집 지어줘" → 설계사무소(spec-driven-generation)가 먼저 "설계도부터 만들죠"라고 말하는 상황
  - **💡 포인트:** "생성하기 전에 멈출 줄 안다" — 이것이 spec-driven-generation의 가장 위대한 점
- **Type:** 개념/개관

---

## 본론 2: 4대 전담 스킬 (Slides 5-8) — 약 6분

### Slide 5: 📝 adr-writing — "왜 이런 결정을 내렸는가"
- **Section:** 본론 2 — 4개 스킬 상세
- **Purpose/goal:** ADR의 핵심 원칙(1 ADR = 1 Concern)과 3가지 금지 규칙을 전달한다. ADR이 결정을 강요하는 것이 아니라 결정을 가능하게 하는 장치임을 강조한다.
- **Content:**
  - **핵심 원칙:** 1 ADR = 1 Concern (하나의 ADR 파일 = 하나의 관심사)
  - **ADR이 다루는 것:** 기술부터 문서 구조, 발표 전개 방식까지 — 산출물에 영향을 미치는 모든 구조적 결정
  - **세 가지 금지 규칙 (Constraints):**
    1. ❌ **가짜 대안 제시 금지** — 모든 대안은 실질적이고 장단점이 팽팽해야 함
    2. ❌ **산문 스타일 금지** — 이름 + 2~4개 불릿으로만 트레이드오프 요약
    3. ❌ **무단 승인 금지** — 사용자의 명시적 구두 승인 후에만 `proposed` → `approved`
  - **ADR 예시:** "보고서 서론 구조 — chronological vs problem-first vs executive summary"
  - **💡 포인트:** ADR은 결정을 **강요하는** 것이 아니라 결정을 **가능하게 하는** 장치다
- **Type:** 개념/상세

### Slide 6: 📋 spec-writing — "무엇을 만들기로 했는가"
- **Section:** 본론 2 — 4개 스킬 상세
- **Purpose/goal:** Spec의 핵심 원칙(자체 완결성)과 2가지 금지 규칙을 전달한다. Spec이 ADR(과거)과 Generation(미래) 사이의 현재 설계도임을 강조한다.
- **Content:**
  - **핵심 원칙:** 자체 완결성 (Self-Containment)
  - **Spec vs ADR:**
    - ADR = 결정의 **이유** (왜)
    - Spec = 결정된 **사실** (무엇을)
  - **두 가지 금지 규칙 (Constraints):**
    1. ❌ **ADR 참조 금지** — "상세는 adr/01.md 참고" 금지. 결정 결과만 복사
    2. ❌ **산문 스타일 금지** — 하나의 불릿 = 하나의 명확한 요구사항 (검증 가능해야 함)
  - **Spec의 세 가지 역할:**
    - 📝 생성의 유일한 기준점
    - ✅ 검증의 유일한 기준점
    - 🤝 의사소통의 공통 참조점
  - **💡 포인트:** Spec은 '과거 기록(ADR)'과 '미래 결과물(Generation)' 사이의 **현재 설계도**
- **Type:** 개념/상세

### Slide 7: ⚙️ spec-driven-generation — "설계도대로 짓는다"
- **Section:** 본론 2 — 4개 스킬 상세
- **Purpose/goal:** ASR Readiness Gate의 상세 동작과 Spec 우선 원칙을 전달한다. Slide 4에서 개관한 내용을 이 스킬의 관점에서 심화한다.
- **Content:**
  - **핵심 원칙:** Spec이 유일한 헌법이자 근거
  - **ASR Readiness Gate (상세):**
    - Spec 없음 → 8가지 보편 ASR 범주 스캔
    - Open ASR 중 대안 있음 → adr-writing 호출
    - 모든 결정 완료 → spec-writing 호출 → Spec 완성 → 생성
  - **Spec 우선 원칙:** 최초 프롬프트보다 Spec을 우선시
  - **비유:** 건축 현장에서 "일단 벽돌부터 쌓아봐"라는 요청에도 현장소장이 "설계도부터 확인합시다"라고 말하는 상황
  - **💡 포인트:** "생성하기 전에 멈출 줄 아는" 용기 — 이것이 AI를 올바르게 쓰는 첫걸음
- **Type:** 개념/상세

### Slide 8: 🔍 spec-driven-verification — "설계도와 대조하라"
- **Section:** 본론 2 — 4개 스킬 상세
- **Purpose/goal:** 검증의 이중 목적(Deviation + Undocumented ASR)과, 두 경우 모두 사용자가 재검토하고 학습해야 할 대상임을 전달한다. Undocumented ASR의 3단계 처리(Ignore/Accept/Escalate)를 설명한다.
- **Content:**
  - **핵심 원칙:** 이중 목적 검증 — Deviation + Undocumented ASR
  - **두 가지 검증 대상:**
    1. ⚠️ **Deviation (이탈):** Spec 명시와 다른 구현
       - 예: "PostgreSQL 쓰라니까 MongoDB"
       - 원인: 설계자의 실수, 구현 중 발견된 충돌, AI의 문제 해결 과정
       - **→ 사용자 재검토:** 설계가 잘못된 것일 수도, 구현이 잘못된 것일 수도
    2. ❓ **Undocumented ASR (미문서화 ASR):** Spec에 없었지만 AI가 생성물에 구현한 내용
       - 예: AI가 임의로 하위 섹션 추가, AI가 새로운 요구사항이라고 제시
       - **→ 사용자 재검토:** 중요성과 확신 정도에 따라 3가지 처리
  - **Undocumented ASR 처리 3단계:**
    1. **🙅 무시 (Ignore):** 중요하지 않음 → 그대로 둔다. Spec 미포함
    2. **✅ 수용 (Accept):** 중요하고 타당 → Spec에 포함
    3. **🔍 ADR 검토 (Escalate):** 중요하지만 판단 어려움 → ADR로 결정
  - **검증의 숨은 목적 — 사용자 학습:**
    - ASR을 놓친 경험, 잘못된 의사결정 경험 → **사용자 역량 개선**
    - "아, 이걸 미리 결정했어야 했구나" → 다음 설계에서 같은 실수 반복 방지
  - **💡 포인트:** 검증은 생성물을 검사하는 동시에 **사용자의 다음 설계를 더 나아지게 만드는 피드백 루프**
- **Type:** 개념/상세 + 원칙

---

## 본론 3: 연결과 요약 (Slide 9) — 약 2분

### Slide 9: 🔗 파이프라인 연결 + 요약 매트릭스
- **Section:** 본론 3 — 연결과 요약
- **Purpose/goal:** 4개 스킬이 입력-출력으로 연결된 파이프라인임을 시각적으로 보여주고, 각 스킬의 역할과 금지 사항을 매트릭스로 요약한다.
- **Content:**
  - **전체 파이프라인 다이어그램 (상단):**
    ```
    [Routing] → [spec-driven-gen Gate] → [adr-writing] → [spec-writing]
                                              ↓                ↓
                                         사용자 결정        Spec 완성
                                              ↓                ↓
                                    [spec-driven-gen 생성] → [spec-driven-verification]
                                                                     ↓
                                                              [순환: 검증 결과 → Spec 개정]
    ```
  - **요약 매트릭스 (중단):**
    | 스킬 | 핵심 역할 | 절대 금지 사항 |
    |------|---------|--------------|
    | adr-writing | 대안 분석 & 결정 기록 | 가짜 대안, 산문 스타일, 무단 승인 |
    | spec-writing | 결정 통합 설계도 | ADR 참조, 장황한 줄글 |
    | spec-driven-gen | Spec 기준 생성 | Spec 없이 생성 (Gate 통과 필수) |
    | spec-driven-ver | Deviation + Undoc ASR 발견 | 사용자 컨펌 없이 독단적 수정 |
  - **스킬 간 협력 관계:**
    - adr-writing → spec-writing: 결정 결과 전달 (ADR 참조 금지)
    - spec-writing → spec-driven-generation: Spec 전달 (유일한 입력)
    - spec-driven-generation → spec-driven-verification: 생성물 전달
    - spec-driven-verification → spec-writing: Undocumented ASR → Spec 개정 (순환)
  - **💡 포인트:** 각 스킬은 각자의 역할에 집중하고, **인터페이스(파일)로만 소통**한다
- **Type:** 요약/시각화

---

## 결론 (Slides 10-11) — 약 4분

### Slide 10: 💡 핵심 통찰 + 학습 확인
- **Section:** 결론
- **Purpose/goal:** 오늘의 3가지 핵심 통찰을 복습하고, 청중이 스스로 질문에 답하며 이해를 확인하도록 유도한다.
- **Content:**
  - **3가지 핵심 통찰:**
    1. **🔄 "스킬이 없으면 spec-driven-generation이 폴백된다"** — Cocrates의 기본 동작
    2. **🚦 "생성하기 전에 멈출 줄 안다"** — Spec Readiness Gate가 구조 없는 생성을 차단
    3. **🔗 "각 스킬은 파일로만 소통한다"** — adr/ → spec/ → 생성물 → verification/
  - **학습 확인 (자기 질문):**
    - ❓ 사용자가 "보고서 작성해줘"라고 요청하면, Cocrates는 어떤 순서로 동작하는가?
    - ❓ `spec-driven-generation`이 spec 없을 때 하는 3단계는?
    - ❓ adr-writing의 3가지 금지 규칙은?
    - ❓ Undocumented ASR 발견 시 3가지 처리 방법은?
- **Type:** 요약 + 학습 확인

### Slide 11: 다음 에피소드 + 마무리 인용
- **Section:** 결론
- **Purpose/goal:** Ep12(메타-스킬)에 대한 기대감을 형성하고 시리즈 전체 흐름을 점검한다. 마무리 인용으로 메시지를 각인시킨다.
- **Content:**
  - **지금까지의 여정:**
    - Ep10: Generation 파이프라인 원리 ✅
    - **Ep11: 4대 스킬의 메커니즘 ✅**
    - Ep12: 스킬을 생성하는 메타-스킬 (generating-skill-creation)
  - **다음 편 예고:** "구조적 접근법 자체를 스스로 스킬화하는 마법"
  - **마무리 인용:**
    > "구조를 지배하는 자가 결과물의 주권을 쥡니다. 그리고 최고의 구조는 다시 자동화된 스킬로 박제됩니다."
  - 🦉
- **Type:** 전환/마무리
