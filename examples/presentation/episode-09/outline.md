# Outline — Episode 09: 소크라테스식 학습 활동 - 스킬

**Estimated time:** ~13분
**Number of slides:** 13
**Template:** 타이틀 → 학습목표 → 콘텐츠(3개 스킬 내부 탐험) → 학습 확인 → 다음편연결

---

## 템플릿 구조

```
도입 (슬라이드 1~2) — 약 2분
본론 (슬라이드 3~11) — 약 9분
  - Education 스킬 내부: No Spoon-feeding + Turn-based Mission + Bloom's 2D/Push&Pull + 3블록/DIP예시 (슬라이드 3~5)
  - Knowledge Capture 스킬 내부: Ignorance 기록 + Merge (슬라이드 6~8)
  - Reflection 스킬 내부: 면접관 + Pull-only + Gap 처리 (슬라이드 9~11)
결론 (슬라이드 12~13) — 약 2분
```

---

### 1. 도입 (슬라이드 1~2) — 약 2분

#### 슬라이드 1: 타이틀 페이지
- **연결:** Ep8에서 Learning 파이프라인의 3대 철학과 3단계 구조 이해 → Ep9는 그 원리가 실제 `.opencode/skills/`에서 **어떤 코드와 규칙**으로 구현되었는지 내부 탐험
- **핵심 질문:** "Ep4에서 Cocrates가 했던 그 질문들, 누군가가 미리 스크립트를 짜놓은 걸까요?"
- **부제:** 3블록 미션과 무지를 박제하는 코드 엔진

#### 슬라이드 2: 학습 목표
- **도입 질문:** "Ep8에서 Learning 파이프라인의 원리를 배웠습니다. 그런데 이 원칙들이 실제로 어떻게 '코드'로 굳어져 있을까요?"
- **3가지 학습 목표:**
  1. Education 스킬의 No Spoon-feeding 원칙, Turn-based Mission 구조, 그리고 3블록 출력 포맷
  2. Knowledge Capture 스킬의 Ignorance 기록 철학과 Merge 전략
  3. Reflection 스킬의 Interviewer 페르소나와 Gap 발견 시 "가르치지 않는" 원칙

---

### 2. 본론 (슬라이드 3~11) — 약 9분

#### 2.1 Education 스킬 내부 (슬라이드 3~5) — 4분

##### 슬라이드 3: 🧗‍♂️ Education Skill — No Spoon-feeding과 Turn-based Mission
- **규칙 1 — No Spoon-feeding (최상위 헌법):**
  - "절대 한 턴에 완전한 정답을 주지 않는다"
  - 개념 설명은 최대 1~3문장, 응답의 20% 이하로 제한
  - **Incomplete State 원칙:** 대화를 항상 불완전한 상태로 멈춤 → 사용자가 채워야 함
- **규칙 2 — Turn-based Mission:**
  - 모든 응답은 정확히 하나의 `🔥 [MISSION]` 블록으로 끝나야 함
  - 사용자가 현재 MISSION에 답하기 전에 절대 다음 힌트를 스포일러하지 않음
- **두 규칙의 관계:** No Spoon-feeding은 '제한', Turn-based Mission은 '생성' — 하나의 엔진

##### 슬라이드 4: 📊 Bloom's 2D Matrix + Push/Pull 전략
- **동시 격상 금지:** Y축(인지 과정)과 X축(지식 차원)을 동시에 올리지 않음
- **Push 전략:** 초보자나 인지적 붕괴 시, 낮은 단계부터 단계적으로 안내
- **Pull 전략 (기본값):** 높은 도전을 먼저 던져 사용자가 끌어당기게 함
- **판단 규칙:** 기본 Pull, 혼란 시 Push로 전환 후 한 축만 낮춤

##### 슬라이드 5: 🔥 3블록 출력 포맷 + DIP 실제 사례
- **3블록 출력 포맷 (SKILL.md 실제 규칙):**
  - 💡 **Concept Briefing:** 핵심 원리를 1~3문장, 일상적 비유로 전달
  - 💻 **Thought Lab:** 결함이 있거나 불완전한 예시/시나리오 제시
  - 🔥 **MISSION:** 정확히 하나의 인지적 과제
- **DIP 실제 예시 (SKILL.md 원문):**
  - 사용자: "DIP 원칙에 대해 알려줘"
  - 💡 Briefing: "DIP는 전기 플러그와 같습니다..."
  - 💻 Lab: `class OrderService { constructor(private db: MySQLDatabase) {} }` // flaw
  - 🔥 MISSION: "MySQLDatabase를 PostgresDatabase로 바꾸면 무엇이 깨질까요?"
- **포인트:** 세 블록이 하나의 완전한 학습 턴을 이룬다. 이 구조가 모든 교육 대화의 단일 엔진.

#### 2.2 Knowledge Capture 스킬 내부 (슬라이드 6~8) — 3분

##### 슬라이드 6: 💾 Knowledge Capture Skill — 핵심만 캡처한다
- **핵심 원칙:** 장황한 설명 금지. 회상(Recall) 가능한 최소 단위만 저장
- **저장 위치:** `kb/{topic-slug}.md`
- **저장 대상:** 한 줄 정의, 비유, ✅ 핵심 이해, ⚠️ **Wrong Assumptions / Gaps** ← 가장 중요

##### 슬라이드 7: ⚠️ 무지(Ignorance)의 기록 — 오개념 박제의 힘
- **철학:** "무엇을 새로 배웠는지"보다 "**어떤 오개념을 가지고 있었고 어떻게 틀렸었는지**"를 박제하는 것이 장기 기억의 핵심
- **KB 실제 예시:** `kb/bloom-taxonomy.md`의 Wrong Assumptions 섹션 예시

##### 슬라이드 8: 🔄 Merge 전략 — 지식 중복 방지
- 새로운 파일 무분별한 양산 금지
- 동일 주제 검색 → 발견 시 덮어쓰지 않고 새 인사이트를 `## Update History`와 함께 추가
- 병합 판단 기준: 파일명, 제목, Tags

#### 2.3 Reflection 스킬 내부 (슬라이드 9~11) — 3분

##### 슬라이드 9: 🕵️‍♂️ Reflection Skill — 면접관의 등장
- **페르소나 전환:** 친절한 코치 → **엄격한 면접관(Interviewer)**
- **목적:** 채점이 아니라 "내가 아는 것과 모르는 것의 정확한 지도"를 함께 그림
- **교육 vs 평가 비교:** Coach(Education) vs Interviewer(Reflection)

##### 슬라이드 10: 🎯 Pull-only 평가 — 가르치지 않고 측정한다
- **Push 금지:** Reflection은 절대 Push(가르치기)로 전환하지 않음
- **질문 유형:** "반례를 들어보세요", "경계 조건은?", "다른 도메인에 적용해보세요"
- **Gap 발견 시:** 그 자리에서 강의하지 않고 기록만 → 세션 말미에 Education 제안

##### 슬라이드 11: 🔗 세 스킬의 유기적 연결
- Education(탐구) → Knowledge Capture(기록) → Reflection(검증) → (Gap 시 Education 환류)
- 세 스킬이 **동일한 Bloom's 2D 매트릭스**를 공유하지만 목적이 다름
  - Education: Push/Pull로 성장
  - Reflection: Pull-only로 측정
- KB가 Education과 Reflection을 연결하는 가교 역할

---

### 3. 결론 (슬라이드 12~13) — 약 2분

#### 슬라이드 12: 핵심 요약 + 학습 확인
- **3-Point 요약:**
  1. Education 스킬: No Spoon-feeding이 최상위 헌법. 3블록 엔진(Briefing→Lab→MISSION) + Push/Pull로 난이도 동적 조절
  2. Knowledge Capture 스킬: "무엇을 배웠는지"보다 "무엇을 오해했는지"를 기록한다
  3. Reflection 스킬: 가르치지 않고 측정만 한다 — Gap은 Education으로 환류한다
- **학습 확인 (자기 질문):**
  - Education의 No Spoon-feeding과 Turn-based Mission을 설명할 수 있는가?
  - 3블록 출력 포맷(Concept Briefing → Thought Lab → MISSION)의 각 역할은?
  - Reflection이 Gap 발견 시 왜 가르치지 않는가?

##### 슬라이드 13: 다음 에피소드 연결
- **지금까지:** Learning 파이프라인(원리 → 스킬 내부) 완전 정복
- **Ep10: 구조 기반 산출물 생성 활동 - 원리**
  - Cocrates의 두 번째 Core Activity — "당장 만들어줘"라는 요청에 Cocrates가 왜 승인부터 받으라고 하는가?
  - ASR → ADR → Spec → 생성 → 검증 파이프라인의 개념적 이해
- **마무리:** "답을 아는 것과, 그 답이 무너지지 않도록 구조화하는 것은 완전히 다른 차원의 도달입니다."
