# Presentation Outline — Episode 12

**Title:** 구조 기반 워크플로우 생성 스킬 — generating-skill-creation: 스킬을 만드는 메타 엔진
**Total estimated time:** 16분 (11슬라이드, ~1.5분/슬라이드)
**Template:** 타이틀 → 학습목표 → 3개 본론 섹션 → 학습 확인 → 다음편연결

---

## 1. Introduction (슬라이드 1~2) — 2분

### Slide 1: Title
- **Key message:** Ep11(4개 스킬) → Ep12(메타-스킬). 구조적 접근을 **스스로 스킬화하는 방법**으로 들어간다.
- **Content:**
  - 주제: 구조 기반 워크플로우 생성 스킬
  - 부제: generating-skill-creation — 스킬을 만드는 메타-스킬
  - 에피소드 연결: Ep11 4대 스킬 → **Ep12 스킬을 생성하는 스킬**
- **Time:** 30초

### Slide 2: 학습 목표
- **Key message:** 오늘 알게 될 4가지
- **Content:**
  1. 🧬 **generating-skill-creation**의 정체성 — "스킬을 만드는 스킬"
  2. 🔍 산출물을 분해하는 **5대 프리즘** (Artifact Components)
  3. ❄️ **Snowflake 5단계** — define → plan → architecture design → detail design → generation
  4. 🏗️ **Meta Snowflake 7단계** — Kernel → Components → Frame → Outline → Spec → Skill → Verification
- **Time:** 1분 30초

---

## 2. 본론 1: 메타-스킬의 정체성 (슬라이드 3~6) — 6분

### Slide 3: 🧬 두 가지 접근법 — 구조를 아는 것과 모르는 것
- **Key message:** 산출물 생성에는 두 가지 접근법이 있다. **구조를 모르면 ASR→ADR→Spec으로 구조를 만들고, 구조를 알면 그 구조를 기준으로 점진 구체화(snowflake)한다.** generating-skill-creation은 후자를 위한 **스킬을 설계하는 메타-스킬**이다.
- **Content:**
  - **사례: 여행 계획 수립**
    - **구조를 모를 때 (spec-driven-generation):** 호텔, 식당, 활동을 개별 검토 → 막연하게 계획 수립
      - ASR: "어떤 호텔?" / ADR: "시내 vs 교외" / 결정 → Spec 반영
    - **구조를 알 때 (스킬 사용):** 여행 계획 = **timeline 구조**
      - 숙박 → 조식 → 활동 → 중식 → 활동 → ... (일별 timeline)
      - 여행지1 → 여행지2 → 여행지3 → ... (경로 구조)
      - → 먼저 여행지 경로를 결정하고, 각 여행지에서 구체적 timeline 계획
      - → 각 timeline 단계에서 호텔, 식당, 활동 검토 (더 구체적인 설계 결정 가능)
  - **메타-스킬:** 이 '여행 계획의 구조(timeline)'를 **재사용 가능한 템플릿(스킬)** 으로 만드는 것
    - 일반 스킬: 구조를 따라 산출물 생성 (예: "일본 여행 계획 짜줘" → timeline 구조로 생성)
    - generating-skill-creation: **새로운 산출물 유형의 구조를 발견하고 스킬로 박제** (예: "출장용 여행 계획 스킬 만들어줘")
  - **보고서 사례 연결 (Ep10):** 문서의 골격 → 섹션 요약 → 섹션 내용 작성 = 구조를 알 때의 snowflake 접근
  - **💡 포인트:** 두 접근법의 차이는 '구조를 아는가 모르는가' — generating-skill-creation은 **알게 된 구조를 영구적인 스킬로 전환**한다
- **Time:** 1분 30초

### Slide 4: 🔍 산출물 5대 프리즘 — 소설 창작으로 이해하기
- **Key message:** 새 스킬을 설계할 때 generating-skill-creation은 대상 산출물을 **5가지 차원**으로 분해한다. 산출물 유형에 따라 이름은 달라져도 **역할은 동일**하다.
- **Content:**
  - **소설 창작의 5대 프리즘:**
    1. 📋 **과업과 제약** — 장르(판타지/미스터리), 분량(300쪽/시리즈), 마감일, 출판사 가이드라인
    2. 🎨 **맥락과 규칙** — 세계관(중세/현대/미래), 매직 시스템 규칙, 톤(진지/유쾌), 언어 스타일
    3. 🧩 **개체 (등장인물)** — 주인공, 조연, 악당, 장면, 플롯 포인트, 테마
    4. 📍 **공간과 배치** — 챕터 구성, 1/2/3막 구조, 페이지 레이아웃
    5. 🔗 **구조와 흐름** — 발단 → 전개 → 위기 → 절정 → 결말 (5막 구조), 서브플롯
  - **같은 프리즘, 다른 산출물:** 블로그 시리즈 스킬, 이미지 생성 스킬도 동일한 5대 차원으로 분석 가능
  - **💡 포인트:** 소설이든 보고서든 게임이든 — **이 5대 프리즘으로 모든 산출물 유형을 분석**할 수 있다
- **Time:** 1분 30초

### Slide 5: ❄️ Snowflake 5단계 — 소설 창작의 점진 구체화
- **Key message:** 새 스킬의 생성 절차는 5단계로 점진 구체화되며, **detail design이 완료되기 전까지 generation은 절대 금지**.
- **Content:**
  - **소설 창작 Snowflake:**
    - **define:** 장르(판타지 스릴러), 목표 독자(성인), 분량(400쪽), 핵심 테마 결정
    - **plan:** 전체 스토리 아크, 3막 구조 개관, 주요 등장인물 개요
    - **architecture design:** 챕터별 개요, 각 막의 전환점, 인물 관계도, 서브플롯 구조
    - **detail design:** 각 챕터의 장면 리스트, 대화/설명/액션 비율, 세부 플롯 포인트
    - **generation:** 실제 집필
  - **🚨 절대 규칙:** detail design(장면 리스트)이 완료되기 전까지 generation(집필) 금지
    - 이 게이트가 없으면? 막연한 설정만 있고 이야기는 진전되지 않는 '세계관 덤프' 소설 탄생
  - **💡 포인트:** 이 Snowflake은 presentation-authoring 스킬(이 발표자료)과 소설 창작이 **동일한 구조**로 설계될 수 있음을 보여준다
- **Time:** 1분 30초

### Slide 6: 📊 단계별 구성 요소 구체화 — 소설 창작의 구성 요소별 진화
- **Key message:** 각 구성 요소는 **동일한 속도로 구체화되지 않는다.** 과업과 제약은 define에서 완성되지만, 등장인물(개체)은 generation에서 비로소 완성된다.
- **Content:**
  - **소설 창작 Per-Stage Refinement 표:**
    | 구성 요소 | define | plan | arch design | detail design | generation |
    |-----------|--------|------|-------------|---------------|------------|
    | **과업과 제약** | 장르/분량 ✅ | — | — | — | — |
    | **맥락과 규칙** | 세계관 방향 | 톤/분위기 요약 | 세계관 규칙 카탈로그 | 세부 규칙/변형 | — |
    | **개체(인물)** | 역할 유형 | 관계 요약 | 인물 카탈로그/계층 | 인물별 상세 설정 | ✅ 인물별 챕터 |
    | **공간(챕터)** | 필요성 확인 | 3막 구조 개관 | 챕터별 카탈로그 | 챕터별 장면 리스트 | ✅ 집필 |
    | **구조와 흐름** | 5막 구조 | 상위 요약 | 전체 플롯 계층 | 단위별 상세 | ✅ 최종 조립 |
  - **지연 구체화(Lazy Refinement):** 구조(플롯)가 개체(인물)를 참조할 때 먼저 인물 카탈로그(architecture)를 완성하고, 인물별 상세 설정(detail)은 필요할 때 채움
  - **💡 포인트:** 모든 것을 한 번에 완벽히 설계하려 하지 않는다 — **적시에 적절한 수준으로 구체화**
- **Time:** 1분 30초

---

## 3. 본론 2: Meta Snowflake 7단계 (슬라이드 7~8) — 3분

### Slide 7: 🏗️ Meta Snowflake 7단계 개관
- **Key message:** generating-skill-creation 자신이 새 SKILL.md를 만들 때는 7단계를 밟는다. 각 단계는 Snowflake 5단계에 매핑된다.
- **Content:**
  - **7단계 파이프라인:**
    ```
    [Kernel] → [Components] → [Frame] → [Outline] → [Spec] → [Skill] → [Verification]
    ```
  - **Snowflake 매핑:**
    | Meta 단계 | Snowflake | 핵심 활동 |
    |-----------|-----------|----------|
    | Kernel | define | 생성 대상을 **한 문장**으로 정의 |
    | Components | plan | 5대 차원 식별 |
    | Frame | architecture design | 워크플로우 + 파일 구조 + 승인 지점 설계 |
    | Outline | **detail design** 🚨 | 단계별 파일 산출물, 입력, 완료 기준, 승인 조건 정의 |
    | Spec → Skill → Verification | generation | 명세 확정 → 파일 생성 → 검증 |
  - **🚨 게이트:** Outline(detail design)이 완전히 확정되기 전까지 Spec/Skill(geneneration)으로 넘어가지 않음
  - **💡 포인트:** 이 7단계는 이 발표자료 슬라이드의 구조를 설계할 때와 **동일한 논리**다
- **Time:** 1분 30초

### Slide 8: 📋 각 단계별 상세 — Kernel부터 Verification까지
- **Key message:** 각 단계는 고유한 목적과 산출물이 있다. 특히 **Kernel(한 문장 정의)** 이 흔들리면 모든 후속 단계가 흔들린다.
- **Content:**
  - **Kernel** — "이 스킬은 [대상 상황]에서 [산출물 유형]을 [검토 가능한 단계]를 거쳐 생성하도록 돕는다"
  - **Components** — 5대 차원을 산출물 유형에 맞게 명명하고 범위 설정
  - **Frame** — 5단계 Snowflake 워크플로우 + 디렉토리 구조 + 승인 지점
  - **Outline (Detail Design)** — 각 단계별 파일 산출물 정의:
    ```
    - Input: 의존하는 산출물
    - Creation activity: 생성 작업
    - Completion criteria: 완료 조건
    - Review questions: 검토 질문
    - Approval criteria: 승인 조건
    ```
  - **Spec → Skill → Verification** — 확정된 명세로 SKILL.md 생성 및 검증
  - **💡 포인트:** Outline Format이 중요한 이유 — 중간 산출물이 **채팅에 휘발되지 않고 파일로 관리**됨
- **Time:** 1분 30초

---

## 4. 본론 3: 제약과 진화 (슬라이드 9) — 2분

### Slide 9: 🚫 금지 사항 + 🧬 재사용성과 진화
- **Key message:** generating-skill-creation의 힘은 '하는 일'보다 **'하지 않는 일(금지 사항)'** 과 **'재사용'** 에서 나온다.
- **Content:**
  - **4가지 금지 사항:**
    1. ❌ 구조 분석 없이 SKILL.md 템플릿만 채우기
    2. ❌ 사용자 검토 지점 없는 스킬 작성
    3. ❌ 중간 산출물을 파일로 저장하지 않음
    4. ❌ 기존 스킬과 중복되는 새 스킬
  - **재사용성:**
    - 한 번 만든 스킬 → 이후 동일 유형 요청 시 **시스템이 자동 로드**
    - **artifact-specific skill의 탄생:** document-authoring, presentation-authoring, blog-series-authoring은 모두 이 방식으로 탄생
    - 새 산출물 유형 = **스킬만 추가**하면 끝. 기존 스킬은 독립적으로 개선 가능
  - **💡 포인트:** generating-skill-creation은 '만들고 사라지는' 스킬이 아니라 **스킬 생태계의 출발점**
- **Time:** 2분

---

## 5. Conclusion (슬라이드 10~11) — 3분

### Slide 10: 💡 핵심 통찰 + 학습 확인
- **Key message:** 오늘의 3가지 핵심 통찰을 기억하고, 스스로 질문에 답해보자.
- **Content:**
  - **3가지 핵심 통찰:**
    1. **🧬 "generating-skill-creation은 메타-스킬이다"** — 산출물이 아니라 절차(SKILL.md)를 생성
    2. **❄️ "모든 스킬은 Snowflake로 점진 구체화된다"** — define → plan → architecture design → detail design → generation
    3. **🚨 "detail design 완료 전 generation은 절대 금지"** — 이 게이트가 구조 없는 스킬 생성을 차단
  - **학습 확인 (자기 질문):**
    - ❓ 산출물 5대 프리즘을 모두 말할 수 있는가?
    - ❓ 산출물을 점진적으로 구체화 하는 Snowflake의 단계는?
    - ❓ Meta Snowflake 7단계의 순서는?
- **Time:** 2분 (통찰 30초 + 질문 1분 30초)

### Slide 11: 다음 에피소드 + 마무리 인용
- **Key message:** Ep12(메타-스킬)까지 이해했다면 Cocrates Harness의 모든 구조적 활동을 살펴본 셈. 이제 마지막 — **지속적 개선과 사용자 선언**.
- **Content:**
  - **지금까지의 여정 (Part 3 완료):**
    - Ep7: Cocrates Harness 구조 ✅
    - Ep8~9: Learning 파이프라인 (원리 + 스킬) ✅
    - Ep10~11: Generation 파이프라인 (원리 + 스킬) ✅
    - **Ep12: 메타-스킬 (generating-skill-creation) ✅**
  - **다음 편 예고 (Part 4):**
    - Ep13: 올바른 Cocrates Harness 활용 — 지속적 개선
    - Ep14: Cocrates 사용자 선언문 — 7계명
  - **마무리 인용:**
    > "시스템에 스킬을 박제하는 순간, 당신의 노하우는 AI의 유전자에 영원히 각인됩니다."
- **Time:** 1분

---

## Appendix: 슬라이드별 시간 배분

| 구간 | 슬라이드 | 시간 | 누적 |
|------|---------|------|------|
| 도입 | 1~2 | 2분 | 0~2분 |
| 본론1: 메타-스킬 정체성 | 3~6 | 6분 | 2~8분 |
| 본론2: Meta Snowflake 7단계 | 7~8 | 3분 | 8~11분 |
| 본론3: 제약과 진화 | 9 | 2분 | 11~13분 |
| 결론 | 10~11 | 3분 | 13~16분 |
