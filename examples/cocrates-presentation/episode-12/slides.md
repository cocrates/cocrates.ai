# Slide Plan — Episode 12: 구조 기반 워크플로우 생성 스킬

**Total slides:** 11
**Estimated pace:** ~1.5분/슬라이드 (총 ~16분)
**Template:** 타이틀 → 학습목표 → 본론1(4) → 본론2(2) → 본론3(1) → 학습확인 → 다음편연결

---

## 도입 (Slides 1-2) — 약 2분

### Slide 1: Title — 구조 기반 워크플로우 생성 스킬
- **Section:** 도입
- **Purpose/goal:** Ep11(4개 스킬) → Ep12(메타-스킬) 전환을 명확히 하고, '스킬을 만드는 스킬'이라는 새로운 개념에 대한 기대감을 형성한다.
- **Content:**
  - **주제:** 구조 기반 워크플로우 생성 스킬
  - **부제:** generating-skill-creation — 스킬을 만드는 메타-스킬
  - **에피소드 연결:** Ep11 4대 스킬 ✅ → **Ep12 스킬을 생성하는 스킬**
  - **핵심 전환 메시지:** "Ep11에서 우리는 4개의 전담 스킬이 어떻게 구조적 생성을 구현하는지 배웠습니다. 그런데, 만약 필요한 스킬이 없다면요? 스킬 자체를 어떻게 만들까요? 오늘은 그 질문에 답합니다."
- **Type:** 전환/타이틀
- **Visual:** 중앙 집중형 타이틀. 시리즈 공통 템플릿(블루/네이비 그라데이션). 제목 아래 부제, 하단에 에피소드 표시.

### Slide 2: 학습 목표
- **Section:** 도입
- **Purpose/goal:** 청중이 오늘 발표를 통해 얻게 될 4가지 목표를 제시한다.
- **Content:**
  - **도입 질문:** "Ep10~11에서 구조 기반 생성의 '왜(Why)'와 '어떻게(How)'를 배웠습니다. 그렇다면 이 구조 자체를 스킬로 만드는 방법은 무엇일까요?"
  - **4가지 학습 목표:**
    1. 🧬 **generating-skill-creation**의 정체성 — "스킬을 만드는 스킬"
    2. 🔍 산출물을 분해하는 **5대 프리즘** (Artifact Components)
    3. ❄️ **Snowflake 5단계** — define → plan → architecture design → detail design → generation
    4. 🏗️ **Meta Snowflake 7단계** — Kernel → Components → Frame → Outline → Spec → Skill → Verification
- **Type:** 학습 목표
- **Visual:** 4개 목표를 번호나 아이콘과 함께 리스트 형태로. 각각 간단한 키워드만 표시.

---

## 본론 1: 메타-스킬의 정체성 (Slides 3-6) — 약 6분

### Slide 3: 🧬 두 가지 접근법 — 구조를 아는 것과 모르는 것
- **Section:** 본론 1 — 메타-스킬의 정체성
- **Purpose/goal:** 산출물 생성의 두 가지 접근법(구조를 모를 때 vs 알 때)을 비교하고, generating-skill-creation이 왜 필요한지 동기를 부여한다. 여행 계획과 보고서 사례로 직관적 이해를 돕는다.
- **Content:**
  - **사례: 여행 계획 수립**
    - **구조를 모를 때 (spec-driven-generation):** 호텔, 식당, 활동을 개별 검토 → 막연하게 계획 수립
      - ASR: "어떤 호텔?" / ADR: "시내 vs 교외" / 결정 → Spec 반영
    - **구조를 알 때 (스킬 사용):** 여행 계획 = **timeline 구조**
      - 숙박 → 조식 → 활동 → 중식 → 활동 → ... (일별 timeline)
      - 여행지1 → 여행지2 → 여행지3 → ... (경로 구조)
      - → 먼저 여행지 경로를 결정하고, 각 여행지에서 구체적 timeline 계획
      - → 각 timeline 단계에서 호텔, 식당, 활동 검토 (더 구체적인 설계 결정 가능)
  - **메타-스킬:** 이 'timeline 구조'를 **재사용 가능한 템플릿(스킬)** 으로 만드는 것
  - **보고서 사례 연결 (Ep10):** 문서의 골격 → 섹션 요약 → 섹션 내용 작성 = 구조를 알 때의 snowflake 접근
  - **💡 포인트:** 두 접근법의 차이는 '구조를 아는가 모르는가' — generating-skill-creation은 **알게 된 구조를 영구적인 스킬로 전환**한다
- **Type:** 개념/비유
- **Visual:** 좌우 비교 레이아웃. 왼쪽: "구조를 모를 때 → 개별 검토(ASR/ADR)", 오른쪽: "구조를 알 때 → timeline 구조". 하단에 메타-스킬 개념 연결.

### Slide 4: 🔍 산출물 5대 프리즘 — 소설 창작으로 이해하기
- **Section:** 본론 1 — 메타-스킬의 정체성
- **Purpose/goal:** generating-skill-creation이 새 스킬을 설계할 때 대상 산출물을 5가지 차원으로 분해한다는 것을 소설 창작 예시로 전달한다.
- **Content:**
  - **소설 창작의 5대 프리즘:**
    1. 📋 **과업과 제약** — 장르(판타지/미스터리), 분량(300쪽/시리즈), 마감일, 출판사 가이드라인
    2. 🎨 **맥락과 규칙** — 세계관(중세/현대/미래), 매직 시스템 규칙, 톤(진지/유쾌), 언어 스타일
    3. 🧩 **개체 (등장인물)** — 주인공, 조연, 악당, 장면, 플롯 포인트, 테마
    4. 📍 **공간과 배치** — 챕터 구성, 1/2/3막 구조, 페이지 레이아웃
    5. 🔗 **구조와 흐름** — 발단 → 전개 → 위기 → 절정 → 결말 (5막 구조), 서브플롯
  - **같은 프리즘, 다른 산출물:** 블로그 시리즈 = 주제(과업), 톤앤매너(맥락), 에피소드(개체), 시리즈 순서(공간), 서론-본론-결론(구조)
  - **💡 포인트:** 소설이든 보고서든 게임이든 — **이 5대 프리즘으로 모든 산출물 유형을 분석**할 수 있다
- **Type:** 개념/프레임워크
- **Visual:** 5개 카드 또는 컬럼 레이아웃. 각 프리즘을 소설 예시와 함께 표시. 하단에 "같은 프리즘, 다른 산출물" 연결선.

### Slide 5: ❄️ Snowflake 5단계 — 소설 창작의 점진 구체화
- **Section:** 본론 1 — 메타-스킬의 정체성
- **Purpose/goal:** 새 스킬의 생성 절차가 5단계 Snowflake로 점진 구체화된다는 것과, detail design → generation 게이트의 중요성을 소설 창작 예시로 전달한다.
- **Content:**
  - **소설 창작 Snowflake:**
    - **define:** 장르(판타지 스릴러), 목표 독자(성인), 분량(400쪽), 핵심 테마 결정
    - **plan:** 전체 스토리 아크, 3막 구조 개관, 주요 등장인물 개요
    - **architecture design:** 챕터별 개요, 각 막의 전환점, 인물 관계도, 서브플롯 구조
    - **🚨 detail design:** 각 챕터의 장면 리스트, 대화/설명/액션 비율, 세부 플롯 포인트
    - **generation:** 실제 집필
  - **게이트 설명:** detail design(장면 리스트)이 완료되기 전까지 generation(집필) 금지
    - 게이트가 없으면? → 막연한 설정만 있고 이야기는 진전되지 않는 '세계관 덤프' 소설
    - 게이트가 있으면? → 장면 리스트라는 설계도에 따라 체계적으로 집필
  - **💡 포인트:** 이 Snowflake은 presentation-authoring(이 발표자료)도, 소설 창작 스킬도 **동일한 구조로 설계**될 수 있다
- **Type:** 개념/프로세스
- **Visual:** Snowflake 5단계를 왼쪽에서 오른쪽으로 진행하는 프로세스 다이어그램. 각 단계 아래 소설 예시 텍스트. detail design과 generation 사이에 🚨 게이트 아이콘 강조.

### Slide 6: 📊 단계별 구성 요소 구체화 — 소설 창작의 구성 요소별 진화
- **Section:** 본론 1 — 메타-스킬의 정체성
- **Purpose/goal:** 각 구성 요소가 동일한 속도로 구체화되지 않는다는 점을 소설 창작 예시로 전달한다. 어떤 요소는 define에서 완성되고, 어떤 요소는 generation에서 비로소 완성된다.
- **Content:**
  - **소설 창작 Per-Stage Refinement 표:**
    | 구성 요소 | define | plan | arch design | detail design | generation |
    |-----------|--------|------|-------------|---------------|------------|
    | **과업과 제약** | 장르/분량 ✅ 확정 | — | — | — | — |
    | **맥락과 규칙** | 세계관 방향 | 톤/분위기 요약 | 세계관 규칙 카탈로그 | 세부 규칙/변형 | — |
    | **개체(인물)** | 역할 유형 | 관계 요약 | 인물 카탈로그/계층 | 인물별 상세 설정 | ✅ 인물별 챕터 |
    | **공간(챕터)** | 필요성 확인 | 3막 구조 개관 | 챕터별 카탈로그 | 챕터별 장면 리스트 | ✅ 집필 |
    | **구조와 흐름** | 5막 구조 | 상위 요약 | 전체 플롯 계층 | 단위별 상세 | ✅ 최종 조립 |
  - **관찰점:**
    - 과업과 제약: define에서 확정되고 이후 변하지 않음 (장르/분량)
    - 개체(인물): 역할 유형 → 관계 → 카탈로그 → 상세 설정 → 인물별 챕터로 점진 완성
    - 구조와 흐름(플롯): 5막 구조 → 상위 요약 → 전체 계층 → 단위별 상세 → 최종 조립
  - **지연 구체화(Lazy Refinement):** 플롯이 인물을 참조할 때 → 먼저 인물 카탈로그(architecture)를 완성하고, 인물별 상세(detail)는 필요할 때 채움
  - **💡 포인트:** 모든 것을 한 번에 완벽히 설계하려 하지 않는다 — **적시에 적절한 수준으로 구체화**
- **Type:** 개념/데이터 시각화
- **Visual:** 표 형태로 구성 요소별 5단계 구체화 과정 표시. ✅와 —의 패턴을 관찰. 하단에 지연 구체화 화살표 다이어그램.

---

## 본론 2: Meta Snowflake 7단계 (Slides 7-8) — 약 3분

### Slide 7: 🏗️ Meta Snowflake 7단계 개관
- **Section:** 본론 2 — Meta Snowflake 7단계
- **Purpose/goal:** generating-skill-creation 자신이 새 SKILL.md를 만들 때 7단계를 밟는다는 것을 개관한다. 각 단계가 Snowflake 5단계에 어떻게 매핑되는지 보여준다.
- **Content:**
  - **7단계 파이프라인:**
    ```
    [Kernel] → [Components] → [Frame] → [Outline] → [Spec] → [Skill] → [Verification]
    ```
  - **Snowflake 매핑:**
    | Meta 단계 | Snowflake 대응 | 핵심 활동 |
    |-----------|----------------|----------|
    | **Kernel** | define | 생성 대상을 **한 문장**으로 정의 |
    | **Components** | plan | 5대 차원 식별 |
    | **Frame** | architecture design | 워크플로우 + 파일 구조 + 승인 지점 설계 |
    | **Outline** | **🚨 detail design** | 단계별 파일 산출물, 입력, 완료 기준, 승인 조건 정의 |
    | **Spec → Skill → Verification** | generation | 명세 확정 → 파일 생성 → 검증 |
  - **🚨 핵심 게이트:** Outline(detail design)이 완전히 확정되기 전까지 Spec/Skill(generation)로 넘어가지 않음
  - **비유 (소설 창작 연결):** Outline = 장면 리스트. 장면 리스트 없이 집필하지 않는 것처럼, Outline 없이 SKILL.md를 생성하지 않음
  - **💡 포인트:** 이 7단계는 이 발표자료 슬라이드의 구조를 설계할 때와 **동일한 논리**로 작동한다
- **Type:** 개념/개관
- **Visual:** 상단에 7단계 수평 파이프라인 다이어그램. 하단에 Snowflake 매핑 표. Outline과 Spec 사이에 🚨 게이트 표시 강조.

### Slide 8: 📋 각 단계별 상세 — Kernel부터 Verification까지
- **Section:** 본론 2 — Meta Snowflake 7단계
- **Purpose/goal:** 각 단계의 구체적인 목적과 산출물을 전달한다. 특히 Kernel(한 문장 정의)이 모든 후속 단계의 기준이 됨을 강조한다.
- **Content:**
  - **Kernel** — "이 스킬은 [대상 상황]에서 [산출물 유형]을 [검토 가능한 단계]를 거쳐 생성하도록 돕는다"
    - Kernel이 흔들리면 모든 후속 단계가 흔들림 → define 단계에서 가장 중요한 활동
  - **Components** — 5대 차원을 산출물 유형에 맞게 명명하고 범위 설정
  - **Frame** — 5단계 Snowflake 워크플로우 + 디렉토리 구조 + 승인 지점
  - **Outline (Detail Design)** — 각 단계별 파일 산출물 정의 (5가지 항목):
    ```
    - Input: 의존하는 산출물
    - Creation activity: 생성 작업
    - Completion criteria: 완료 조건
    - Review questions: 검토 질문
    - Approval criteria: 승인 조건
    ```
  - **Spec → Skill → Verification** — 확정된 명세로 SKILL.md 생성 및 검증
    - Spec: frontmatter, 원칙, 절차, 금지 사항, 완료 조건
    - Skill: `.opencode/skills/{slug}/SKILL.md` 파일 생성
    - Verification: 스킬이 생성 하네스로 작동하는지 검증
  - **💡 포인트:** Outline Format이 중요한 이유 — 중간 산출물이 **채팅에 휘발되지 않고 파일로 관리**되도록 보장
- **Type:** 개념/상세
- **Visual:** 7단계를 수직 리스트로 표시. 각 단계별 간단한 설명과 아이콘. Outline 단계는 박스 강조. 하단에 Outline Format 예시 코드 블록.

---

## 본론 3: 제약과 진화 (Slide 9) — 약 2분

### Slide 9: 🚫 금지 사항 + 🧬 재사용성과 진화
- **Section:** 본론 3 — 제약과 진화
- **Purpose/goal:** generating-skill-creation의 6가지 금지 사항과 스킬의 재사용성·진화 메커니즘을 전달한다. 금지 사항이 '품질을 보장하는 장치'임을 강조한다.
- **Content:**
  - **4가지 금지 사항 (Constraints):**
    1. ❌ **구조 분석 없이 SKILL.md 템플릿만 채우기** — 5대 프리즘 분석 생략
    2. ❌ **사용자 검토 지점 없는 스킬 작성** — 승인 게이트 없는 독단적 설계
    3. ❌ **중간 산출물을 파일로 저장하지 않음** — 채팅에만 남기고 휘발
    4. ❌ **기존 스킬과 중복되는 새 스킬** — 시스템 효율성 저하
  - **재사용성:**
    - 한 번 만든 스킬 → 이후 동일 유형 요청 시 **시스템이 자동 로드** — generating-skill-creation 재개입 불필요
    - **artifact-specific skill의 탄생:** document-authoring, presentation-authoring, blog-series-authoring 모두 이 방식으로 탄생
    - 새 산출물 유형 = **스킬만 추가**하면 끝. 기존 스킬은 독립적으로 개선 가능
  - **💡 포인트:** generating-skill-creation은 '만들고 사라지는' 스킬이 아니라 **스킬 생태계의 출발점**이다
- **Type:** 원칙/요약
- **Visual:** 좌측: 6가지 금지 사항을 빨간 ❌ 아이콘과 함께 리스트. 우측: 스킬 생태계 다이어그램 — generating-skill-creation이 새로운 스킬을 만들고, 그 스킬들이 시스템에 등록되어 재사용되는 순환 구조.

---

## 결론 (Slides 10-11) — 약 3분

### Slide 10: 💡 핵심 통찰 + 학습 확인
- **Section:** 결론
- **Purpose/goal:** 오늘의 3가지 핵심 통찰을 복습하고, 청중이 스스로 질문에 답하며 이해를 확인하도록 유도한다.
- **Content:**
  - **3가지 핵심 통찰:**
    1. **🧬 "generating-skill-creation은 메타-스킬이다"** — 최종 산출물이 아니라 절차(SKILL.md)를 생성
    2. **❄️ "모든 스킬은 Snowflake로 점진 구체화된다"** — define → plan → architecture design → detail design → generation
    3. **🚨 "detail design 완료 전 generation은 절대 금지"** — 이 게이트가 구조 없는 생성을 차단
  - **학습 확인 (소크라테스식 자기 질문):**
    - ❓ 산출물 5대 프리즘을 모두 말할 수 있는가? (힌트: 소설 창작 예시)
    - ❓ 산출물을 점진적으로 구체화 하는 Snowflake의 단계는? (힌트: Define부터 Generation까지)
    - ❓ Meta Snowflake 7단계의 순서는? (힌트: Kernel부터 Verification까지)
- **Type:** 요약 + 학습 확인
- **Visual:** 상단: 3가지 통찰을 카드 형태로. 하단: 4개 질문을 화살표나 말풍선으로 표시. 각 질문 아래 작은 힌트.

### Slide 11: 다음 에피소드 + 마무리 인용
- **Section:** 결론
- **Purpose/goal:** Part 3(구조)의 완료를 선언하고, Part 4(결론)로의 전환을 안내한다. 마무리 인용으로 메시지를 각인시킨다.
- **Content:**
  - **지금까지의 여정 (Part 3 완료):**
    - Ep7: Cocrates Harness 구조 ✅
    - Ep8~9: Learning 파이프라인 (원리 + 스킬) ✅
    - Ep10~11: Generation 파이프라인 (원리 + 스킬) ✅
    - **Ep12: 메타-스킬 (generating-skill-creation) ✅**
  - **다음 편 예고 (Part 4):**
    - Ep13: 올바른 Cocrates Harness 활용 — **사용자와 에이전트의 공진화**
    - Ep14: Cocrates 사용자 선언문 — **7계명**
  - **마무리 인용:**
    > "시스템에 스킬을 박제하는 순간, 당신의 노하우는 AI의 유전자에 영원히 각인됩니다."
  - 🦉 (올빼미 — Cocrates 상징)
- **Type:** 전환/마무리
- **Visual:** 상단: 지금까지 여정을 체크리스트 형태로. 중단: 다음 편预告. 하단: 인용문 강조. 오른쪽 아래: 올빼미 아이콘.
