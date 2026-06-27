# Outline — Episode 07: Cocrates Harness 구조 — 아키텍처 딥다이브

**Estimated pace:** ~1.25분/슬라이드 (총 ~15분)
**Number of slides:** 11

---

## 템플릿 구조

```
도입 (슬라이드 1~2) — 약 2.5분
본론 (슬라이드 3~9) — 약 9분
결론 (슬라이드 10~11) — 약 2.5분
```

---

### 1. 도입 (슬라이드 1~2) — 약 2.5분

#### 슬라이드 1: 타이틀 페이지
- **연결:** Ep6까지 Cocrates를 **사용**했다면, Ep7부터는 Cocrates가 **내부적으로 어떻게 돌아가는지** 구조를 뜯어본다
- **핵심 전환:**
  - "Cocrates는 최종 완성된 하네스가 아닙니다"
  - "자신의 업무 패턴에 맞춰 지속적으로 발전시키고 진화시켜야 합니다"
  - "이를 위해서 Cocrates Harness의 구조를 정확하게 이해하는 것이 필요합니다"
- **부제:** 하나의 거대한 프롬프트가 품을 수 없는 'AI 주방'의 비밀

#### 슬라이드 2: 학습 목표
- **도입 질문:** "여러분, Cocrates가 Ep4에서는 학습을, Ep5에서는 코드 생성을, Ep6에서는 스킬 생성을 했습니다. 하나의 AI가 이렇게 다른 활동을 어떻게 가능하게 할까요?"
- **3가지 학습 목표:**
  1. Cocrates가 Agent + Skills 구조를 채택한 이유를 설명할 수 있다
  2. Cocrates Agent Prompt의 6개 섹션과 각 역할을 설명할 수 있다
  3. Intent-To-Skill Routing이 의도를 기반으로 어떻게 스킬을 선택하는지 설명할 수 있다

---

### 2. 본론 (슬라이드 3~9) — 약 9분

#### 2.1 만능 칼의 함정 (슬라이드 3) — 1.5분

##### 슬라이드 3: 🍳 "만능 칼의 함정" — 하나의 프롬프트로는 안 된다
- **비유:** 전문 주방에서 생선 회 뜨기, 스테이크 굽기, 디저트 플레이팅은 각각 다른 칼과 조리 순서가 필요
- **AI의 현실:** 다양한 산출물 유형(코드, 문서, 발표자료, 학습 활동, 블로그 시리즈 등)은 각각 완전히 다른 구조적 접근이 필요
  - 보고서/문서: 논리 계층 구조 (목차, 섹션, 단락)
  - 발표자료: 페이지 단위 레이아웃 + 거버닝 메시지
  - 학습 활동: 질문-피드백 흐름 (턴제 미션)
- **결론:** 하나의 프롬프트로 모든 것을 처리 → 지나치게 일반적이거나 특정 유형에 치우침

#### 2.2 Agent + Skills 구조 (슬라이드 4) — 1.5분

##### 슬라이드 4: 🏛️ Cocrates Harness의 두 가지 기둥
- **Cocrates Agent** (`cocrates.md`): 공통 원칙과 컨트롤 타워 — 최상위 헌법
- **Skills** (`.opencode/skills/*/SKILL.md`): 독립된 전문가 팀 — 산출물별 최적화 워크플로우
- **분리 원칙:** 공통의 통제 체계는 유지하되, 개별 워크플로우를 독립적으로 확장
- **장점:** (1) 각 산출물 유형에 최적화 (2) 개별 추가/수정/개선 가능 (3) 새로운 유형은 스킬만 추가

#### 2.3 Cocrates Agent 6대 섹션 (슬라이드 5) — 1.5분

##### 슬라이드 5: 📜 Cocrates Agent Prompt — 6개 섹션 개관
- **전체 구조:** Persona → Principle → Harness Architecture → Request Handling → Core Activities → Success Criteria
- **간략 설명:**
  1. **Persona:** "불확실성을 체계적인 탐구로 전환" — 에이전트의 역할과 태도
  2. **Principle:** Harness Ignorance — 무지의 통제, 검토 없는 생성 금지
  3. **Harness Architecture:** Agent + Skills 분리 구조 선언
  4. **Request Handling:** 의도 기반 스킬 라우팅 (다음 슬라이드에서 상세)
  5. **Core Activities:** 산출물 생성 + 학습 두 파이프라인
  6. **Success Criteria:** 사용자가 산출물을 스스로 설명할 수 있는가

#### 2.4 Intent-To-Skill Routing (슬라이드 6) — 1.5분

##### 슬라이드 6: 🔀 Intent-To-Skill Routing — 의도가 스킬을 선택한다
- **원리:** 단순 키워드 매칭이 아닌 사용자의 근본 의도(Intent) 추론
- **주요 매핑 (표):**
  | 의도 | 스킬 |
  |------|------|
  | 개념을 배우고 싶다 | `education` |
  | 정리/저장하고 싶다 | `knowledge-capture` |
  | 이해도를 확인받고 싶다 | `reflection` |
  | 선택지를 비교하고 싶다 | `adr-writing` |
  | 명세를 정의하고 싶다 | `spec-writing` |
  | 산출물을 생성하고 싶다 | `spec-driven-generation` |
  | 검증하고 싶다 | `spec-driven-verification` |
  | 스킬을 만들고 싶다 | `generating-skill-creation` |
- **라우팅 규칙:** 의도 추론 → 스킬 로드 → 있으면 처리, 없으면 spec-driven-generation

#### 2.5 Core Activity: 산출물 생성 (슬라이드 7) — 1분

##### 슬라이드 7: 🏗️ Core Activity 1 — 산출물 생성 파이프라인
- **파이프라인:** ADR → Spec → Generation → Verification
- **각 단계 역할:**
  - ADR: 선택지 분석과 결정
  - Spec: 결정된 요구사항 통합
  - Generation: Spec 기반 생성
  - Verification: 항목별 검증
- **연결:** Ep5에서 실습으로 경험한 파이프라인

#### 2.6 Core Activity: 학습 (슬라이드 8) — 1분

##### 슬라이드 8: 🎓 Core Activity 2 — 학습 파이프라인
- **파이프라인:** Education → Knowledge Capture → Reflection
- **각 단계 역할:**
  - Education: 소크라테스식 질문-미션 기반 학습
  - Knowledge Capture: 배운 내용을 KB에 기록
  - Reflection: 실제 이해도 평가
- **연결:** Ep4에서 실습으로 경험한 파이프라인

#### 2.7 Cocrates는 완성된 하네스가 아니다 (슬라이드 9) — 1분

##### 슬라이드 9: 🔄 Cocrates는 계속 진화한다
- **핵심 메시지:** Cocrates는 최종 완성된 하네스가 아니다
- **사용자가 할 수 있는 것:**
  - 기존 스킬을 자신의 업무 패턴에 맞게 수정
  - 새로운 산출물 유형에 대한 스킬을 직접 생성 (Ep6에서 배운 방법)
  - Agent Prompt의 원칙을 자신의 맥락에 맞게 조정
- **구조 이해 = 커스터마이징과 진화의 첫걸음**
- **Ep6 연결:** Ep6에서 배운 `generating-skill-creation`으로 새 스킬을 만들 수 있다

---

### 3. 결론 (슬라이드 10~11) — 약 2.5분

#### 슬라이드 10: 핵심 요약 + 학습 확인
- **3-point 요약:**
  1. **🍳 만능 칼은 없다.** 산출물 유형마다 구조적 접근이 다르므로, 하나의 프롬프트로는 모든 것을 처리할 수 없다. Cocrates는 Agent(헌법)와 Skills(전문 부대)를 분리하여 이 문제를 해결한다.
  2. **📜 Cocrates Agent = 6개 섹션의 헌법.** Persona → Principle → Harness Architecture → Request Handling → Core Activities → Success Criteria. 이 중 Request Handling의 Intent-To-Skill Routing이 의도 기반 스킬 선택의 핵심이다.
  3. **🔄 Cocrates는 완성된 하네스가 아니다.** 구조를 이해해야 커스터마이징과 진화가 가능하다. 새로운 스킬을 추가하고 기존 스킬을 수정하며 자신의 업무 패턴에 맞게 발전시켜 나가야 한다.
- **학습 확인 (자기 질문):**
  - ❓ Cocrates가 Agent + Skills 구조를 채택한 이유를 설명할 수 있는가?
  - ❓ Cocrates Agent Prompt의 6개 섹션을 말할 수 있는가?
  - ❓ Intent-To-Skill Routing이 의도를 기반으로 어떻게 스킬을 선택하는지 설명할 수 있는가?

#### 슬라이드 11: 다음 에피소드 연결
- **지금까지:** Cocrates의 전체 구조(Agent + Skills, 6개 섹션, Routing, Core Activities)를 이해했다
- **Ep8: 소크라테스식 학습 활동 - 원리**
  - Part 3의 두 번째 에피소드 — 첫 번째 Core Activity(학습 파이프라인)의 원리로 딥다이브
  - Ep4에서 실습으로 경험한 학습 방식이 어떤 원리로 설계되었는지
  - 소크라테스 산파술, 블룸 분류학, ZPD가 어떻게 Cocrates에 구현되었는지
- **마지막 질문:**
  > "여러분은 Cocrates를 '사용'하는 것에서 멈출 것인가, 아니면 Cocrates를 '이해'하고 자신의 것으로 진화시킬 것인가?"

---

## 타임라인 요약

| 구간 | 슬라이드 | 시간 | 누적 |
|------|----------|------|------|
| 도입 | 1-2 | 2.5분 | 2.5분 |
| 만능 칼의 함정 | 3 | 1.5분 | 4분 |
| Agent + Skills 구조 | 4 | 1.5분 | 5.5분 |
| Agent Prompt 6대 섹션 | 5 | 1.5분 | 7분 |
| Intent-To-Skill Routing | 6 | 1.5분 | 8.5분 |
| Core Activity 1 (생성) | 7 | 1분 | 9.5분 |
| Core Activity 2 (학습) | 8 | 1분 | 10.5분 |
| Cocrates 진화 | 9 | 1분 | 11.5분 |
| 결론 | 10-11 | 2.5분 | ~14분 |
