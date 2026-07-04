# Slide Plan — Episode 07: Cocrates Harness 구조 — 아키텍처 딥다이브

**Total slides:** 11
**Estimated pace:** ~1.25분/슬라이드 (총 ~15분)
**Template:** 타이틀 → 학습목표 → 콘텐츠(에피소드 주도) → 학습 확인 → 다음편연결

---

## 도입 (Slides 1-2) — 약 2.5분

### Slide 1: Title — Cocrates Harness 구조
- **Section:** 도입
- **Purpose/goal:** Ep6(스킬 생성 실습)에서 Ep7(아키텍처 딥다이브)로의 전환을 알린다. Cocrates를 "사용"하는 것에서 "이해"하는 것으로 시청자의 시각을 전환시킨다. Cocrates는 완성된 하네스가 아니라 지속적으로 진화해야 하는 존재임을 암시한다.
- **Content:**
  - **주제:** Cocrates Harness 구조 — 아키텍처 딥다이브
  - **부제:** 하나의 거대한 프롬프트가 품을 수 없는 'AI 주방'의 비밀
  - **에피소드 연결:** Ep6까지 Cocrates를 **사용**했다면, Ep7부터는 Cocrates가 **내부적으로 어떻게 돌아가는지** 구조를 뜯어본다
  - **핵심 전환 메시지:**
    - "Cocrates는 최종 완성된 하네스가 아닙니다"
    - "자신의 업무 패턴에 맞춰 지속적으로 발전시키고 진화시켜야 합니다"
    - "이를 위해서 Cocrates Harness의 구조를 정확하게 이해하는 것이 필요합니다"
- **Type:** 전환/타이틀

### Slide 2: 학습 목표
- **Section:** 도입
- **Purpose/goal:** 시청자가 이 발표를 통해 얻게 될 3가지 핵심 목표를 제시한다. "Cocrates는 어떻게 이렇게 다양한 활동을 할 수 있을까?"라는 호기심을 자극한다.
- **Content:**
  - **도입 질문 (청중 자극):** "여러분, Cocrates가 Ep4에서는 학습을, Ep5에서는 코드 생성을, Ep6에서는 스킬 생성을 했습니다. 하나의 AI가 이렇게 다른 활동을 어떻게 가능하게 할까요?"
  - **3가지 학습 목표:**
    1. 🏛️ Cocrates가 Agent + Skills 구조를 채택한 이유를 설명할 수 있다
    2. 📜 Cocrates Agent Prompt의 6개 섹션과 각 역할을 설명할 수 있다
    3. 🔀 Intent-To-Skill Routing이 의도를 기반으로 어떻게 스킬을 선택하는지 설명할 수 있다
- **Type:** 학습 목표

---

## 본론 (Slides 3-9) — 약 9분

### Slide 3: 🍳 "만능 칼의 함정" — 하나의 프롬프트로는 안 된다
- **Section:** 본론 — 만능 칼의 함정 (2.1)
- **Purpose/goal:** 왜 Cocrates가 Agent + Skills 구조를 채택했는지에 대한 근본적인 이유를 전달한다. "하나의 프롬프트로 모든 산출물을 처리할 수 없다"는 통찰을 주방 비유로 이해시킨다.
- **Content:**
  - **비유:** 전문 주방에서 생선 회 뜨기, 스테이크 굽기, 디저트 플레이팅은 각각 다른 칼과 조리 순서가 필요 — 하나의 만능 칼로는 모든 요리의 품질을 낼 수 없다
  - **AI의 현실 — 산출물 유형별 구조적 접근의 차이:**
    - 보고서/문서: 논리 계층 구조 (목차, 섹션, 단락)
    - 발표자료: 페이지 단위 레이아웃 + 거버닝 메시지
    - 학습 활동: 질문-피드백 흐름 (턴제 미션)
    - 블로그 시리즈: 에피소드와 스토리라인
  - **결론:** 하나의 프롬프트로 모든 산출물 처리 → 지나치게 일반적이거나 특정 유형에 치우침 → Agent + Skills 분리가 필요
- **Type:** 개념/비유

### Slide 4: 🏛️ Cocrates Harness의 두 가지 기둥
- **Section:** 본론 — Agent + Skills 구조 (2.2)
- **Purpose/goal:** Cocrates 아키텍처의 가장 핵심적인 구조 — Agent와 Skills의 분리 원칙을 전달한다. 각각의 역할과 장점을 명확히 이해시킨다.
- **Content:**
  - **Cocrates Agent** (`cocrates.md`): 공통 원칙과 컨트롤 타워 — 최상위 헌법
    - 정체성, 원칙, 구조, 요청 처리, 핵심 활동, 성공 기준
  - **Skills** (`.opencode/skills/*/SKILL.md`): 독립된 전문가 팀 — 산출물별 최적화 워크플로우
    - 교육, 문서 작성, ADR, 검증 등 각 활동의 구체적 절차
  - **분리 원칙:** 공통의 통제 체계는 유지하되, 개별 워크플로우를 독립적으로 확장
  - **장점 3가지:**
    1. 각 산출물 유형에 최적화된 워크플로우 제공
    2. 개별 스킬의 추가/수정/개선이 독립적으로 가능
    3. 새로운 산출물 유형이 필요하면 스킬만 추가하면 됨
- **Type:** 개념/구조

### Slide 5: 📜 Cocrates Agent Prompt — 6개 섹션 개관
- **Section:** 본론 — Agent 6대 섹션 (2.3)
- **Purpose/goal:** Cocrates Agent(cocrates.md)의 내부 구성을 개괄적으로 보여준다. 6개 섹션이 순차적으로 어떻게 Agent의 행동 전반을 정의하는지 전달한다. (블로그 에피소드보다 간략하게 — 구조적 이해에 초점)
- **Content:**
  - **전체 구조를 한눈에:** Persona → Principle → Harness Architecture → Request Handling → Core Activities → Success Criteria
  - **각 섹션 간략 설명:**
    1. **Persona (정체성):** "불확실성을 체계적인 탐구로 전환하고, 사용자가 결과물을 완전히 이해할 때까지 안내하는 존재"
    2. **Principle (원칙):** Harness Ignorance — 무지의 통제, 검토 없는 생성 금지
    3. **Harness Architecture (구조):** Agent + Skills 분리 원칙 선언
    4. **Request Handling (요청 처리):** 의도 기반 스킬 라우팅 (다음 슬라이드)
    5. **Core Activities (핵심 활동):** 산출물 생성 + 학습 두 파이프라인
    6. **Success Criteria (성공 기준):** 사용자가 산출물을 스스로 설명할 수 있는가?
- **Type:** 개념/개관

### Slide 6: 🔀 Intent-To-Skill Routing — 의도가 스킬을 선택한다
- **Section:** 본론 — Intent-To-Skill Routing (2.4)
- **Purpose/goal:** Cocrates의 두뇌에 해당하는 의도 기반 라우팅의 원리와 실제 매핑 테이블을 전달한다. 단순 키워드 매칭이 아닌 근본 의도 추론의 중요성을 강조한다.
- **Content:**
  - **원리:** 단순 키워드 매칭이 아닌 **사용자의 근본 의도(Intent)를 추론**하여 최적의 스킬 매핑
  - **주요 매핑 테이블:**
    | 의도 | 스킬 |
    |------|------|
    | 개념을 배우고 싶다 | `education` |
    | 배운 내용을 정리하고 싶다 | `knowledge-capture` |
    | 이해도를 확인받고 싶다 | `reflection` |
    | 선택지를 비교하고 싶다 | `adr-writing` |
    | 명세를 정의하고 싶다 | `spec-writing` |
    | 산출물을 생성하고 싶다 | `spec-driven-generation` |
    | 검증하고 싶다 | `spec-driven-verification` |
    | 스킬을 만들고 싶다 | `generating-skill-creation` |
  - **라우팅 규칙 (간략):** 의도 추론 → 적합한 스킬 탐색 → 있으면 로드하여 처리, 없으면 기본 생성 방식(spec-driven-generation)으로 처리
  - **💡 포인트:** 이 표가 Cocrates가 다양한 요청을 정확히 처리할 수 있는 비결이다. 사용자가 "알려줘"라고만 해도, Cocrates는 의도를 추론해 `education` 스킬을 로드한다.
- **Type:** 개념/표

### Slide 7: 🏗️ Core Activity 1 — 산출물 생성 파이프라인
- **Section:** 본론 — Core Activity 1 (2.5)
- **Purpose/goal:** 두 가지 Core Activity 중 첫 번째인 산출물 생성 파이프라인의 전체 흐름을 개괄한다. Ep5에서 실습으로 경험한 내용을 구조적 관점에서 다시 조망한다.
- **Content:**
  - **파이프라인:** ADR → Spec → Generation → Verification
  - **각 단계 역할 (간략):**
    - **ADR:** 선택지 분석과 결정 ("왜 이렇게 할까?")
    - **Spec:** 결정된 모든 요구사항을 통합 ("무엇을 만들까?")
    - **Generation:** 승인된 Spec만을 근거로 생성
    - **Verification:** 결과물이 Spec과 일치하는지 항목별 검증
  - **연결:** Ep5에서 실습으로 경험한 파이프라인 — jsondb 사례의 3개 ADR, 72개 검증 항목
- **Type:** 개념/개관

### Slide 8: 🎓 Core Activity 2 — 학습 파이프라인
- **Section:** 본론 — Core Activity 2 (2.6)
- **Purpose/goal:** 두 번째 Core Activity인 학습 파이프라인의 전체 흐름을 개괄한다. Ep4에서 실습으로 경험한 내용을 구조적 관점에서 다시 조망한다.
- **Content:**
  - **파이프라인:** Education → Knowledge Capture → Reflection
  - **각 단계 역할 (간략):**
    - **Education:** 소크라테스식 질문-미션 기반 학습 ("직접 생각하게 한다")
    - **Knowledge Capture:** 배운 내용을 KB(knowledge base)에 핵심과 Gap 중심으로 기록
    - **Reflection:** 실제 이해도를 평가 — ✅ 잘 아는 것 vs ⚠️ 제대로 모르는 것
  - **연결:** Ep4에서 실습으로 경험한 파이프라인 — Bloom's taxonomy 미션 4개
- **Type:** 개념/개관

### Slide 9: 🔄 Cocrates는 완성된 하네스가 아니다
- **Section:** 본론 — Cocrates 진화 (2.7)
- **Purpose/goal:** 가장 중요한 메시지 — Cocrates는 최종 완성품이 아니라 지속적으로 진화해야 하는 플랫폼임을 전달한다. 구조를 이해하는 것이 커스터마이징과 진화의 첫걸음임을 강조한다.
- **Content:**
  - **핵심 메시지:** Cocrates는 최종 완성된 하네스가 아니다
  - **사용자가 할 수 있는 것 (구조 이해의 결실):**
    - 기존 스킬을 자신의 업무 패턴에 맞게 수정
    - 새로운 산출물 유형에 대한 스킬을 직접 생성 (Ep6의 `generating-skill-creation`으로)
    - Agent Prompt의 원칙을 자신의 맥락에 맞게 조정
  - **구조 이해 = 커스터마이징과 진화의 첫걸음**
  - **Ep6 연결:** Ep6에서 배운 스킬 생성 파이프라인이 바로 이 진화를 위한 핵심 도구
- **Type:** 개념/동기부여

---

## 결론 (Slides 10-11) — 약 2.5분

### Slide 10: 핵심 요약 + 학습 확인
- **Section:** 결론
- **Purpose/goal:** 발표의 핵심을 3가지 포인트로 요약하고, 시청자가 스스로 이해를 확인할 수 있도록 자기 질문을 제시한다.
- **Content:**
  - **3-Point 핵심 요약:**
    1. **🍳 만능 칼은 없다.** 산출물 유형마다 구조적 접근이 다르므로, 하나의 프롬프트로는 모든 것을 처리할 수 없다. Cocrates는 Agent(헌법)와 Skills(전문 부대)를 분리하여 이 문제를 해결한다.
    2. **📜 Cocrates Agent = 6개 섹션의 헌법.** Persona → Principle → Harness Architecture → Request Handling → Core Activities → Success Criteria. 이 중 Intent-To-Skill Routing이 의도 기반 스킬 선택의 핵심이다.
    3. **🔄 Cocrates는 완성된 하네스가 아니다.** 구조를 이해해야 커스터마이징과 진화가 가능하다. 새로운 스킬을 추가하고 기존 스킬을 수정하며 자신의 업무 패턴에 맞게 발전시켜 나가야 한다.
  - **학습 확인 (자기 질문):**
    - ❓ Cocrates가 Agent + Skills 구조를 채택한 이유를 설명할 수 있는가?
    - ❓ Cocrates Agent Prompt의 6개 섹션을 말할 수 있는가?
    - ❓ Intent-To-Skill Routing이 의도를 기반으로 어떻게 스킬을 선택하는지 설명할 수 있는가?
- **Type:** 요약

### Slide 11: 다음 에피소드 연결
- **Section:** 결론
- **Purpose/goal:** 전체 구조 이해의 완료를 선언하고, 다음 단계(Ep8: 학습 파이프라인 원리)에 대한 기대감을 형성한다. "사용 → 이해 → 진화"의 여정을 시청자에게 상기시킨다.
- **Content:**
  - **지금까지의 전환:** Cocrates를 **사용**하는 법 → Cocrates의 **구조** 이해 완료
  - **다음 편 (Ep8): 소크라테스식 학습 활동 - 원리**
    - Part 3의 두 번째 에피소드 — 첫 번째 Core Activity인 **학습 파이프라인의 원리**로 딥다이브
    - Ep4에서 실습으로 경험한 학습 방식이 어떤 원리로 설계되었는지
    - 소크라테스 산파술, 블룸 분류학, ZPD가 어떻게 Cocrates에 구현되었는지
  - **마지막 질문 (시청자에게):**
    > "여러분은 Cocrates를 '사용'하는 것에서 멈출 것인가, '이해'하고 자신의 것으로 진화시킬 것인가?"
  - 🦉
- **Type:** 전환/마무리
