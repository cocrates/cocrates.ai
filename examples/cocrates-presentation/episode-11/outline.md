# Presentation Outline — Episode 11

**Title:** 구조 기반 산출물 생성 활동 - 스킬
**Total estimated time:** 17분 (11슬라이드, ~1.5분/슬라이드)
**Template:** 타이틀 → 학습목표 → 4개 섹션 → 학습 확인 → 다음편연결

---

## 1. Introduction (슬라이드 1~2) — 2분

### Slide 1: Title
- **Key message:** Ep10(원리) → Ep11(스킬). 이제 파이프라인이 어떻게 실제로 동작하는지 들어간다.
- **Content:**
  - 주제: 구조 기반 산출물 생성 활동 - 스킬
  - 부제: 스킬이 없으면? spec-driven-generation이 구조적으로 산출물을 생성하게 한다
  - 에피소드 연결: Ep10 파이프라인 원리 → Ep11 4개 스킬의 내부 메커니즘
- **Time:** 30초

### Slide 2: 학습 목표
- **Key message:** 오늘 알게 될 3가지
- **Content:**
  1. 🔄 Cocrates가 요청을 받으면 어떤 스킬을 쓸지 **라우팅**하는가
  2. 🚦 **spec-driven-generation**이 spec이 없을 때 어떻게 **spec 작성을 유도**하는가
  3. 🛠️ 4개 스킬 각각의 **핵심 원칙과 절대 금지 사항**
- **Time:** 1분 30초

---

## 2. 본론 1: Cocrates 생성 파이프라인의 전체 흐름 (슬라이드 3~4) — 3분

### Slide 3: 🔄 Intent-to-Skill Routing — 요청이 스킬이 되는 순간
- **Key message:** 사용자의 "생성해줘" 요청은 먼저 라우터를 통과한다. 스킬이 있으면 스킬로, 없으면 fallback.
- **Content:**
  - **문제:** 사용자가 "A를 생성해줘"라고 하면, Cocrates는 무엇부터 할까?
  - **라우팅 규칙:**
    1. 최종 산출물 유형(final deliverable type) 식별
    2. 해당 유형에 매칭되는 생성 스킬 검색 (가장 구체적인 타입부터)
    3. 있음 → 해당 스킬 사용 (구조가 이미 스킬에 명세됨)
    4. 없음 → **spec-driven-generation으로 폴백**
  - **비유:** 전문 병원 vs 응급실. 전문과목(스킬)이 있으면 전문의 진료, 없으면 응급실(spec-driven-generation)에서 기본적인 처치.
  - **💡 포인트:** 모든 요청이 spec-driven-generation을 거치는 것은 아니다 — **스킬이 없을 때만 fallback**
- **Time:** 1분 30초

### Slide 4: 🚦 spec-driven-generation Fallback — "spec이 없으면 spec을 만든다"
- **Key message:** spec-driven-generation은 단순 생성기가 아니다. Spec Readiness Gate를 통해 **spec이 없으면 spec 작성을 유도**하는 메타-프로세스다.
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
- **Time:** 1분 30초

---

## 3. 본론 2: 4대 전담 스킬 (슬라이드 5~8) — 6분

### Slide 5: 📝 adr-writing — "왜 이런 결정을 내렸는가"
- **Key message:** ADR은 단순 기록이 아니라, 사용자의 결정권을 보장하고 결정의 역사를 투명하게 남기는 장치.
- **Content:**
  - **핵심 원칙:** 1 ADR = 1 Concern (하나의 ADR 파일 = 하나의 관심사)
  - **ADR이 다루는 것:** 기술부터 문서 구조, 발표 전개 방식까지 — 산출물에 영향을 미치는 모든 구조적 결정
  - **세 가지 금지 규칙 (Constraints):**
    1. ❌ 가짜 대안 제시 금지 — 모든 대안은 실질적이고 장단점이 팽팽해야 함
    2. ❌ 산문 스타일 금지 — 이름 + 2~4개 불릿으로만 트레이드오프 요약
    3. ❌ 무단 승인 금지 — 사용자의 명시적 구두 승인 후에만 `proposed` → `approved`
  - **ADR 예시:** "보고서 서론 구조 — chronological vs problem-first vs executive summary"
  - **💡 포인트:** ADR은 결정을 **강요하는** 것이 아니라 결정을 **가능하게 하는** 장치다
- **Time:** 1분 30초

### Slide 6: 📋 spec-writing — "무엇을 만들기로 했는가"
- **Key message:** Spec은 자체 완결적이어야 한다. 이 문서 하나만 읽으면 누구나 같은 것을 만들 수 있어야 한다.
- **Content:**
  - **핵심 원칙:** 자체 완결성 (Self-Containment)
  - **Spec vs ADR:**
    - ADR = 결정의 **이유** (왜)
    - Spec = 결정된 **사실** (무엇을)
  - **두 가지 금지 규칙 (Constraints):**
    1. ❌ ADR 참조 금지 — "상세는 adr/01.md 참고" 금지. 결정 결과만 복사
    2. ❌ 장황한 줄글 금지 — 하나의 불릿 = 하나의 명확한 요구사항 (검증 가능해야 함)
  - **Spec의 세 가지 역할:**
    - 📝 생성의 유일한 기준점
    - ✅ 검증의 유일한 기준점
    - 🤝 의사소통의 공통 참조점
  - **💡 포인트:** Spec은 '과거 기록(ADR)'과 '미래 결과물(Generation)' 사이의 **현재 설계도**
- **Time:** 1분 30초

### Slide 7: ⚙️ spec-driven-generation — "설계도대로 짓는다"
- **Key message:** Spec이 유일한 헌법. ASR Readiness Gate가 생성 전에 멈춰서 설계도를 확인한다.
- **Content:**
  - **핵심 원칙:** Spec이 유일한 헌법이자 근거
  - **ASR Readiness Gate (상세):**
    - Spec 없음 → 8가지 보편 ASR 범주 스캔
    - Open ASR 중 대안 있음 → adr-writing 호출
    - 모든 결정 완료 → spec-writing 호출 → Spec 완성 → 생성
  - **Spec 우선 원칙:** 최초 프롬프트보다 Spec을 우선시
  - **비유:** 건축 현장에서 "일단 벽돌부터 쌓아봐"라는 요청에도 현장소장이 "설계도부터 확인합시다"라고 말하는 상황
  - **💡 포인트:** "생성하기 전에 멈출 줄 아는" 용기 — 이것이 AI를 올바르게 쓰는 첫걸음
- **Time:** 1분 30초

### Slide 8: 🔍 spec-driven-verification — "설계도와 대조하라"
- **Key message:** 검증은 단순 통과/실패가 아니다. Deviation과 Undocumented ASR, 모두 **사용자가 재검토하고 학습할 대상**이다.
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
    1. **🙅 무시 (Ignore):** AI가 중요하다고 나열했지만, 사용자 판단에 중요하지 않음
       → **그대로 둔다.** Spec에 포함하지 않음.
    2. **✅ 수용 (Accept):** 중요한 요구사항이고, AI의 결정이 타당하다고 판단
       → **Spec에 포함한다.** 다음 생성부터 공식 요구사항으로 반영.
    3. **🔍 ADR 검토 (Escalate):** 중요한 요구사항인데, AI의 결정을 따라도 될지 판단 어려움
       → **ADR로 검토한다.** 대안을 비교하고 사용자가 직접 결정.
  - **검증의 숨은 목적 — 사용자 학습:**
    - ASR을 놓친 경험, 잘못된 의사결정 경험을 통해 **사용자는 자신의 역량을 개선**한다
    - "아, 이걸 미리 결정했어야 했구나" → 다음 프로젝트에서 같은 실수 반복 방지
    - 검증은 단순한 합격/불합격 심사가 아니라 **사용자의 구조적 사고력을 키우는 학습 재료**
  - **💡 포인트:** 검증은 생성물을 검사하는 동시에 **사용자의 다음 설계를 더 나아지게 만드는 피드백 루프**
- **Time:** 1분 30초

---

## 4. 본론 3: 연결과 요약 (슬라이드 9) — 2분

### Slide 9: 🔗 파이프라인 연결 + 요약 매트릭스
- **Key message:** 4개 스킬은 파이프라인으로 연결되며, 각 스킬의 정체성은 '하는 일'보다 '하지 않는 일(Constraints)'이 결정한다.
- **Content:**
  - **전체 파이프라인 다이어그램:**
    ```
    [Routing] → [spec-driven-gen Gate] → [adr-writing] → [spec-writing]
                                              ↓                ↓
                                         사용자 결정        Spec 완성
                                              ↓                ↓
                                     [spec-driven-gen 생성] → [spec-driven-verification]
                                                                      ↓
                                                               [순환: 검증 결과 → Spec 개정]
    ```
  - **스킬 간 협력 관계 (좌측 노트):**
    - adr-writing → spec-writing: 결정 결과 전달 (ADR 참조 금지)
    - spec-writing → spec-driven-generation: Spec 전달 (유일한 입력)
    - spec-driven-generation → spec-driven-verification: 생성물 전달
    - spec-driven-verification → spec-writing: Undocumented ASR → Spec 개정 (순환)
  - **요약:**
    | 스킬 | 핵심 역할 | 절대 금지 사항 |
    |------|---------|--------------|
    | adr-writing | 대안 분석 & 결정 기록 | 가짜 대안, 산문 스타일, 무단 승인 |
    | spec-writing | 결정 통합 설계도 | ADR 참조, 장황한 줄글 |
    | spec-driven-gen | Spec 기준 생성 | Spec 없이 생성 (Gate 통과 필수) |
    | spec-driven-ver | Deviation + Undoc ASR 발견 | 사용자 컨펌 없이 독단적 수정 |
  - **💡 포인트:** 각 스킬은 각자의 역할에 집중하고, **인터페이스(파일)로만 소통**한다
- **Time:** 2분

---

## 5. 결론 (슬라이드 10~11) — 4분

### Slide 10: 💡 핵심 통찰 + 학습 확인
- **Key message:** 오늘의 3가지 핵심 통찰을 기억하고, 스스로 질문에 답해보자.
- **Content:**
  - **3가지 핵심 통찰:**
    1. **🔄 "스킬이 없으면 spec-driven-generation이 폴백된다"** — 이것이 Cocrates의 기본 동작
    2. **🚦 "생성하기 전에 멈출 줄 안다"** — Spec Readiness Gate가 구조 없는 생성을 차단
    3. **🔗 "각 스킬은 파일로만 소통한다"** — adr/ → spec/ → 생성물 → verification/ 의 명확한 인터페이스
  - **학습 확인 (자기 질문):**
    - ❓ 사용자가 "보고서 작성해줘"라고 요청하면, Cocrates는 어떤 순서로 동작하는가?
    - ❓ `spec-driven-generation`이 spec 없을 때 하는 3단계는?
    - ❓ adr-writing의 3가지 금지 규칙은?
    - ❓ Undocumented ASR 발견 시 3가지 처리 방법은?
- **Time:** 2분 30초 (통찰 30초 + 질문 2분)

### Slide 11: 다음 에피소드 + 마무리 인용
- **Key message:** 다음은 Ep12 — 구조 기반 워크플로우 생성 스킬. 지금까지 배운 구조적 접근을 스스로 스킬화하는 메타-스킬.
- **Content:**
  - **지금까지의 여정:**
    - Ep10: Generation 파이프라인 원리 ✅
    - **Ep11: 4대 스킬의 메커니즘 ✅**
    - Ep12: 스킬을 생성하는 메타-스킬 (generating-skill-creation)
  - **다음 편 예고:** "구조적 접근법 자체를 스스로 스킬화하는 마법"
  - **마무리 인용:**
    > "구조를 지배하는 자가 결과물의 주권을 쥡니다. 그리고 최고의 구조는 다시 자동화된 스킬로 박제됩니다."
  - 🦉
- **Time:** 1분 30초

---

## Appendix: 슬라이드별 시간 배분

| 구간 | 슬라이드 | 시간 | 누적 |
|------|---------|------|------|
| 도입 | 1~2 | 2분 | 0~2분 |
| 본론1: 전체 흐름 | 3~4 | 3분 | 2~5분 |
| 본론2: 4개 스킬 | 5~8 | 6분 | 5~11분 |
| 본론3: 연결과 요약 | 9 | 2분 | 11~13분 |
| 결론 | 10~11 | 4분 | 13~17분 |
