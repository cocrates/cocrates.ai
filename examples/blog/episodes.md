# 에피소드 설계

---

## 1. 소개 — Cocrates Harness의 필요성과 원칙

### Ep1: 같은 LLM, 다른 결과 (`01-same-llm-different-results`)

- **상태**: ✅ 작성 완료 (`episodes/01-same-llm-different-results.md`)
- **Hook**: 두 개발자가 같은 Claude에게 "로그인 모듈 만들어줘"를 입력한다. 한 명은 그대로 붙여넣고, 다른 한 명은 구조를 검토하고 수정한다. 차이는 LLM이 아니라 '사용법'에 있다.
- **스토리라인**:
  1. 도입: 같은 LLM, 다른 결과 — 당신의 AI는 몇 명분인가?
  2. A님의 사례: AI에게 시키는 대로 일을 시킴, 검토 생략
  3. B님의 사례: AI와 함께 고민하고 구조를 먼저 설계함
  4. 기술이 아니라 태도의 문제
  5. Cocrates 소개 — 질문을 통해 사용자의 성장을 돕는 에이전트
- **핵심 메시지**: 모델이 아니라 당신의 태도가 차이를 만든다. AI가 100명이 될지, 1명이 될지는 당신에게 달려 있다.
- **CTA**: 다음 편에서 핵심 원칙을 선언한다.

---

### Ep2: The Unexamined Code Is Not Worth Generating (`02-unexamined-code`)

- **상태**: ✅ 작성 완료 (`episodes/02-unexamined-code.md`)
- **Hook**: *"The unexamined code is not worth generating."* — 검토되지 않은 산출물은 생성할 가치가 없다.
- **스토리라인**:
  1. 선언: 시리즈를 관통하는 한 문장
  2. 'Code'의 정의: 소스 코드뿐 아니라 모든 AI 최종 산출물
  3. 구조(architecture) 설계의 중요성 — 이 블로그 시리즈를 만드는 과정을 예시로
  4. Cocrates가 안내하는 워크플로우
  5. '검토한다'는 것의 의미: 이해 + 판단 + 승인
  6. 검토 = 무지의 제거
- **핵심 메시지**: 검토되지 않은 산출물은 생성할 가치가 없다. 검토한다는 것은 내가 이해하고, 내가 판단하고, 내가 승인하는 것이다.
- **CTA**: 다음 편에서 Cocrates Harness를 설치하고 첫 대화를 시작한다.

---

## 2. 실습 — 먼저 경험한다

### Ep3: Cocrates Harness 설치 (`03-installation`)

- **상태**: ✅ 작성 완료 (`episodes/03-installation.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 3
- **Hook**: 이제까지 Cocrates가 무엇인지, 어떤 원칙으로 움직이는지 알았다. 그렇다면 어떻게 시작해야 할까? Cocrates Harness는 **opencode plugin**으로 동작한다. 설치는 단순하지만, 선택지가 있다.
- **스토리라인**:
  1. **opencode 설치**
     - Cocrates Harness는 opencode 플러그인으로 동작하므로, 먼저 opencode 자체를 설치해야 한다
     - 설치 방법은 세 가지 (사용자 환경에 따라 선택):
       - **opencode 터미널**: Claude 스타일의 TUI에 익숙한 사용자용
       - **opencode desktop**: Cursor, VS Code 같은 IDE 환경에 익숙한 사용자용
       - **opencode 확장**: 기존 Cursor/Code 환경을 그대로 유지하면서 확장 기능으로 사용
     - 설치 상세는 https://opencode.ai/download 참고
     - TUI에 익숙하지 않은 사람들은 VS Code/Cursor에서 opencode 사용을 권장
       - opencode 확장보다 VS Code의 Copilot이나 Cursor와 유사한 GUI 제공하는 OpenChamber 확장 사용을 권장
  2. **cocrates-harness plugin 설치**
     - opencode config 파일(`opencode.jsonc`)의 `plugin` 항목에 `"@cocrates/cocrates-harness"` 추가
     - config 파일 위치:
       - 기본: `~/.config/opencode/` 폴더
       - `OPENCODE_CONFIG_DIR` 환경변수가 설정된 경우: 해당 경로
     - 설정 예시:
       ```json
       {
         "$schema": "https://opencode.ai/config.json",
         "plugin": ["@cocrates/cocrates-harness"]
       }
       ```
  3. **skill 파일 복사**
     - 플러그인 설치 외에도 skills 디렉토리에 Cocrates의 스킬 파일들을 복사해야 한다
     - https://github.com/cocrates/cocrates.ai/tree/main/skills 의 모든 스킬을 `OPENCODE_CONFIG_DIR/skills/` 에 복사
     - 각 스킬은 `.opencode/skills/{skill-name}/SKILL.md` 형태로 제공되며, Cocrates의 워크플로우(교육, 문서 생성, ADR 작성 등)를 담당한다
     - opencode 설치 후 실행하고, 이 문서를 근거로 **"Cocrates Harness 설치해줘"**라고 요청하면 AI가 플러그인 설치와 스킬 복사를 도와준다
  4. **설치 확인 및 첫 대화**
     - 설치가 완료되면 opencode를 재시작하고 **Cocrates Agent**가 활성화되었는지 확인
     - "블룸의 분류학에 대해 알려줘" 같은 질문으로 Cocrates의 소크라테스식 응답을 경험
     - 기대할 수 있는 것: 구조적 질문, 턴제 미션, 검토와 승인 요청
     - 기대하지 말아야 할 것: 단순 정보 덤핑, 무비판적 코드 생성
- **핵심 메시지**: 설치 자체는 단순하다 — opencode를 설치하고, plugin을 추가하고, skills를 복사하면 끝. 하지만 진짜 시작은 설치 후다. Cocrates의 소크라테스식 접근에 익숙해지는 것이 첫 번째 관문이며, 이 시리즈의 실습(Ep4~6)과 구조(Ep7~12)에서 본격적으로 다룬다.
- **포함할 내용**:
  - [ ] opencode 설치 옵션 비교 (터미널 / desktop / 확장)
  - [ ] config 파일 경로 확인 방법 (`~/.config/opencode/` vs `OPENCODE_CONFIG_DIR`)
  - [ ] opencode.jsonc 설정 예시 (plugin 항목)
  - [ ] skills 복사 경로와 방법
  - [ ] 설치 후 첫 대화 추천 질문
- **제외할 내용**:
  - 스킬 설계 방법 (Ep13에서 다룸)
  - 실습 시나리오 (Ep4~6에서 다룸)
  - 각 스킬의 상세 내용 (Ep9, Ep11에서 다룸)
- **CTA**: 설치가 끝났다면, 이제 Cocrates에게 "블룸의 분류학에 대해 알려줘"라고 물어보라. 어떤 일이 벌어질지 궁금하지 않은가? 설치만으로는 아무 일도 일어나지 않는다. Cocrates는 당신의 참여를 기다린다.

---

### Ep4: 소크라테스식 학습 활동 (`04-socratic-learning-practice`)

- **상태**: ✅ 작성 완료 (`episodes/04-socratic-learning-practice.md`)
- **메타**: 예상 분량 14~18분, 발행 순서 4
- **Hook**: "블룸의 분류학에 대해 알려줘" — 평소 같았으면 위키피디아식 설명을 받았을 것이다. 하지만 Cocrates는 질문으로 답했다. 그리고 그 과정에서 블룸 분류학을 **완전히 다르게** 이해하게 되었다. 이 편은 그 여정을 함께 따라가는 실습이다.
- **스토리라인**:
  1. **실습 시나리오**: Bloom's Taxonomy를 Cocrates와 함께 배우기
     - "블룸의 분류학에 대해 알려줘"라는 단순한 질문으로 시작
     - Cocrates는 정답을 주지 않고, 결함이 있는 예시 + 미션으로 첫 턴을 주도
  2. **Education 단계 — 세 가지 깨달음**
     - 첫 번째 깨달음: 피라미드는 '수업 순서'가 아니라 의존성 구조
     - 두 번째 깨달음: 1차원이 아니라 2차원(지식 차원 × 인지 과정) 매트릭스
     - 세 번째 깨달음: Push(순차적) vs Pull(도전적) 전략과 ZPD에 따른 혼합
  3. **Knowledge Capture 단계**: 배운 내용을 `kb/bloom-taxonomy.md`에 핵심과 Gap 중심으로 기록
  4. **Reflection 단계**: Apply→Analyze→Evaluate→Create 4개 미션으로 진짜 이해도 확인
     - MISSION 1(Apply): DIP 도메인에 지식 차원 투영하기
     - MISSION 2(Analyze): 위계와 순서의 차이를 설계로 연결
     - MISSION 3(Evaluate): Push/Pull과 ZPD의 조건부 적용
     - MISSION 4(Create): OOP 첫 3주 커리큘럼 설계
     - 최종 산출물: ✅ 잘 알고 있는 것 / ⚠️ 제대로 알지 못했던 것
  5. **전체 사이클 회고**: Education → Knowledge Capture → Reflection → KB 갱신
- **핵심 메시지**: Learning 파이프라인의 3단계를 실제 대화 기반 스토리텔링으로 경험함으로써, Cocrates의 학습 방식이 단순 정보 전달과 어떻게 다른지 체감한다.
- **포함할 내용**:
  - [ ] 실제 대화 스크립트 예시 (Education 턴제 미션의 구체적 흐름)
  - [ ] 2차원 매트릭스 다이어그램 (Knowledge Dim × Cognitive Process)
  - [ ] Push/Pull 전략 비교 설명
  - [ ] KB 저장 결과 예시 (`kb/bloom-taxonomy.md` 포맷)
  - [ ] Reflection 평가 결과 예시 (✅ / ⚠️)
  - [ ] 독자가 직접 Cocrates와 학습할 때의 팁
- **제외할 내용**:
  - 산출물 생성 실습 (Ep5에서 다룸)
  - 스킬 생성 실습 (Ep6에서 다룸)
- **CTA**: 학습 활동의 실제를 경험했다면, 이제 구조 기반 산출물 생성 활동을 실습해보자. Ep5에서는 ADR부터 검증까지 전 과정을 직접 따라간다.

---

### Ep5: 구조 기반 산출물 생성 활동 (`05-architecture-driven-generation-practice`)

- **상태**: ✅ 작성 완료 (`episodes/05-architecture-driven-generation-practice.md`)
- **메타**: 예상 분량 18~22분, 발행 순서 5
- **Hook**: *"cocrates/jsondb에 jsondb를 개발해줘."* — 평소 같았으면 AI는 바로 "물론입니다!"라며 코드를 쏟아냈을 것이다. 하지만 Cocrates는 질문으로 답했다. *"이 jsondb의 용도는 무엇인가요?"* 그리고 그 질문이 세 개의 ADR과 하나의 Spec, 일곱 개의 Go 파일로 이어졌다. 그런데 **Spec에 따라 생성된 코드에도 예상치 못한 차이(Deviation)가 발견되었다.** 이 편은 생성 이후 검증까지의 전 과정을 함께 따라가는 실습이다.
- **스토리라인**:
  1. **실습 시나리오**: "jsondb를 개발해줘" — 요구사항 명확화부터 시작
  2. **ADR 단계 — AI는 '쉬운 방법'을 제시하지만, 사용자가 결정한다**
     - ADR 1 — 저장 모델: AI의 Collection-Document 가정 → 사용자의 Path-Addressable 결정
     - ADR 2 — 아키텍처: Library-Only → Client-Server로 전환 (사용자가 복수 프로세스 시나리오를 질문)
     - ADR 3 — 동시성 모델: DB-level RWMutex → Per-File RWMutex Map (사용자가 성능 병목 예측)
  3. **Spec 단계**: 승인된 모든 결정을 `spec/jsondb.md`에 자체 완결적으로 통합
  4. **Generation 단계**: Spec만 보고 Library 7개 Go 파일 + Server + CLI 생성, `go build`/`go vet`/테스트 모두 통과
  5. **Verification 단계**: 72개 항목 검증 — 70 pass / 1 fail / 1 not-verifiable
     - Deviation 사례: CLI Schema URL 이중 슬래시 문제
     - Undocumented ASR 6건 발견 (URL encoding, SetPath 검증 누락 등)
  6. **회고**: '쉬운 방법의 함정', 반복/증분 개발의 필요성, Undocumented ASR은 실패가 아니라 기회
- **핵심 메시지**: Cocrates와의 구조 기반 생성은 단순히 코드를 얻는 과정이 아니다. AI가 제시하는 '쉬운 방법'을 무비판적으로 따르는 대신, 사용자가 직접 질문하고 비교하고 결정하는 **적극적인 설계 과정**이다. 생성 후 검증에서 발견된 Deviation과 Undocumented ASR은 Spec을 더 완벽하게 다듬는 **반복/증분 개발의 기회**를 제공한다.
- **포함할 내용**:
  - [ ] 세 가지 ADR 각각의 'Cocrates의 첫 제안' vs '사용자의 검토 후 결정' 비교 (표)
  - [ ] 가장 극적인 반전 — Library-Only → Client-Server 전환 과정
  - [ ] Per-File Mutex 설계 — 성능 고려의 구체적 예시
  - [ ] Spec → Generation → Verification의 실제 산출물 예시
  - [ ] Verification 결과: 72개 항목 검증, 70 pass / 1 fail / 1 not-verifiable
  - [ ] Deviation 사례: CLI Schema URL 이중 슬래시 문제
  - [ ] Undocumented ASR 6건 표
  - [ ] '쉬운 방법의 함정' — AI의 첫 제안이 항상 옳지는 않은 이유
  - [ ] 반복/증분 개발 — Verification → Spec 업데이트 → 재생성 → 재검증 주기
- **제외할 내용**:
  - 학습 활동 실습 (Ep4에서 이미 다룸)
  - 스킬 생성 실습 (Ep6에서 다룸)
- **CTA**: 구조 기반 산출물 생성의 전체 사이클을 경험했다. 이제 직접 스킬을 만들어보는 마지막 실습으로 넘어간다. 그 전에 스스로에게 물어보자: *"AI가 제시한 첫 번째 해결책을, 나는 정말 충분히 검토했는가?"*

---

### Ep6: 구조 기반 워크플로우 생성 (`06-workflow-creation-practice`)

- **상태**: ✅ 작성 완료 (`episodes/06-workflow-creation-practice.md`)
- **메타**: 예상 분량 16~20분, 발행 순서 6
- **Hook**: "설명문/보고서 등을 생성하는 스킬을 생성해줘." — 언뜻 보면 문서 하나를 만들어 달라는 요청 같지만, Cocrates는 다르게 읽는다. 사용자가 요청한 것은 **산출물이 아니라 '그 산출물을 만드는 절차'** 다. 이 요청은 `generating-skill-creation`을 발동시킨다. 이 편은 그 전 과정을 따라가는 실습이다.
- **스토리라인**:
  1. **실습 시나리오**: "설명문/보고서 스킬을 만들어줘" — generating-skill-creation 발동
  2. **1단계 — Kernel**: 핵심을 한 문장으로 정의 ("이 스킬은 일반적 상황에서 Markdown 설명문/보고서를 검토 가능한 단계를 거쳐 생성하도록 돕는다")
  3. **2단계 — Frame**: 문서 계층 구조(메타→개요→본문→결론→부록)와 P1~P5 5단계 설계
  4. **3단계 — Outline**: P1~P5 각 단계의 입력·출력·승인 조건 정의
     - **사용자의 결정적 기여**: "항상 서론→본론→결론 순서는 아니다" → '맞춤형 문서 구성' 원칙으로 발전
  5. **4단계 — Spec**: 5가지 핵심 원칙, 6가지 맞춤형 구성 방식, 7가지 금지사항, 5가지 완료 조건
  6. **5~6단계 — Skill 생성 및 Verification**: SKILL.md 생성 후 7개 검증 항목 모두 통과
  7. **스킬 생성 이후**: "~에 관한 보고서를 생성해줘" 요청이 document-authoring 스킬로 처리되는 방식
- **핵심 메시지**: 스킬 생성은 단순히 `SKILL.md` 파일을 만드는 기술적 작업이 아니다. Cocrates와 사용자가 **대화를 통해 산출물의 구조를 발견하고, 절차를 설계하고, 원칙을 정립하는 공동 설계 과정**이다. Kernel → Frame → Outline → Spec → Skill → Verification의 과정을 통해 모호한 요청이 구조화된 스킬로 탄생한다.
- **포함할 내용**:
  - [ ] Kernel 단계: Cocrates의 세 가지 질문(무엇을, 누가, 어떤 형태로)과 Kernel 문장 도출
  - [ ] Frame: 문서 계층 구조와 생성 단계 설계
  - [ ] Outline: P1~P5 5단계 절차 표
  - [ ] 사용자의 결정적 기여: "항상 서론→본론→결론 순서는 아니다" → '맞춤형 문서 구성' 원칙
  - [ ] 생성된 document-authoring 스킬의 5가지 핵심 설계 요소
  - [ ] 스킬 등록 후의 처리 흐름
  - [ ] 실제 코드 경로: `.opencode/skills/document-authoring/SKILL.md`
- **제외할 내용**:
  - 학습 실습 (Ep4에서 다룸)
  - 산출물 생성 실습 (Ep5에서 다룸)
- **CTA**: 이제 `document-authoring` 스킬이 추가되었다. 직접 한번 시도해보시죠. 업무 보고서나 학습 정리 문서를 요청해보면, Cocrates가 개요부터 섹션 구조, 섹션별 작성까지 단계별로 검토와 승인을 요청하며 함께 만들어갈 것이다. 세 가지 실습을 모두 마쳤다면, 이제 Cocrates Harness의 구조와 원리를 깊이 이해해볼 차례다.

---

## 3. 구조 — 원리를 이해한다

### Ep7: Cocrates Harness 구조 (`07-cocrates-harness-structure`)

- **상태**: ✅ 작성 완료 (`episodes/07-cocrates-harness-structure.md`)
- **메타**: 예상 분량 12~15분, 발행 순서 7
- **Hook**: 하나의 프롬프트로 문서도 쓰고, 코드도 만들고, 학습도 시키고... 가능할까? Cocrates는 불가능하다는 결론에서 출발한다. 하나의 프롬프트로는 모든 산출물을 제대로 다룰 수 없다. 그래서 Cocrates는 **Agent + Skills** 구조로 설계되었다.

- **스토리라인**:
  1. **왜 Agent + Skills 구조인가**
     - 다양한 산출물 유형: 코드, 문서, 발표자료, 블로그 시리즈, 분석 보고서, 학습 노트 등
     - 산출물마다 **구조적 접근이 다름**: 문서는 목차와 섹션, 발표자료는 페이지와 거버닝 메시지, 블로그 시리즈는 에피소드와 스토리라인
     - 하나의 프롬프트로 모든 것을 처리하면? → 지나치게 일반적이거나 특정 유형에 치우침
     - **스킬로 분리하면**: 각 산출물 유형에 최적화된 워크플로우 제공 + 개별 추가/수정/개선 가능
     - **지속적 개선**: 새로운 산출물 유형이 필요하면 스킬만 추가하면 됨. 기존 스킬은 독립적으로 개선 가능.

  2. **Cocrates = Cocrates Agent + Skills**
     - **Cocrates Agent** (`cocrates.md`): 공통된 원칙과 제어를 담당
       - 정체성(Persona), 핵심 원칙(Principle), 구조(Harness Architecture), 요청 처리(Request Handling), 핵심 활동(Core Activities), 성공 기준(Success Criteria)
     - **Skills** (`.opencode/skills/*/SKILL.md`): 구체적 워크플로우를 담당
       - 각 스킬은 독립 파일로 관리되며 추가·수정·확장 가능

  3. **Cocrates Agent Prompt 구조 (`cocrates.md`)** — 다음 섹션으로 구성:
     1. **정체성 (Persona)**: "You are **Cocrates**: an agent that turns uncertainty into disciplined inquiry..." — 에이전트의 역할과 태도를 정의
     2. **핵심 원칙 (Principle)**: Harness Ignorance — 불확실성을 방치하지 않고 아키텍처를 통해 보고·검토·관리 가능하게 만듦. 소크라테스식 태도: 존중하는 탐구, 스스로 깨닫게 하는 질문
      3. **구조 (Harness Architecture)**: Cocrates Agent + Skills. Cocrates Agent는 원칙·의도 인식·스킬 선택·태스크 관리·가드레일. Skills는 작업별 절차.
     4. **요청 처리 (Request Handling)**: 의도 추론(1) → 스킬 로드(2) → 스킬 없으면 spec-driven-generation(3) → 멀티스텝 작업 추적(4). Intent-To-Skill Routing 표.
       5. **핵심 활동 (Core Activities)**: Artifact Generation 파이프라인(ADR → Spec → 생성 → 검증)과 Learning 파이프라인(Education → Knowledge Capture → Reflection)
      6. **성공 기준 (Success Criteria)**: 사용자가 자신의 지식과 무지를 명확히 이해하고, 산출물의 구조와 내용을 스스로 설명할 수 있으며, 모든 단계에 주체적으로 참여했을 때

   4. **전체 요청 처리 흐름**
      - 사용자 요청 → 의도 인식(Intent-To-Skill Routing)
      - 스킬 있음 → 스킬 로드 → 스킬 명세에 따라 처리
      - 스킬 없음 → spec-driven-generation이 직접 처리
      - 사용자가 "스킬을 만들어줘"라고 명시적 요청 → generating-skill-creation으로 새 스킬 설계 → 사용자 승인 → 실행
      - 모든 단계에서 사용자 검토·승인 게이트

- **핵심 메시지**: Cocrates가 Agent + Skills 구조로 설계된 이유는 **산출물마다 다른 구조적 접근이 필요하고, 지속적인 개선이 가능해야 하기 때문**이다. Agent Prompt(`cocrates.md`)는 정체성·원칙·구조·요청처리·핵심활동·성공기준의 6개 섹션으로 구성된다.

- **포함할 내용**:
  - [ ] 다양한 산출물 유형과 각각 필요한 구조적 접근의 차이 (비교)
  - [ ] Cocrates Agent + Skills 구조 다이어그램
  - [ ] Intent-To-Skill Routing 표 (의도별 스킬 매핑)
  - [ ] 두 Core Activity 파이프라인
  - [ ] Cocrates Agent Prompt 구조 요약 (6개 섹션)

- **제외할 내용**:
  - 개별 스킬의 상세 워크플로우 (Ep8~12에서 다룸)
  - 설치 방법 (Ep3에서 다룸)
  - 실습 (Ep4~6에서 다룸)

- **CTA**: 이 구조의 첫 번째 활동 영역 — 소크라테스식 학습 활동의 개념을 다음 편에서 살펴본다. (실습에서 경험한 학습 방식이 어떤 원리로 설계되었는지 알아보자.)

---

### Ep8: 소크라테스식 학습 활동 - 원리 (`08-socratic-learning-activity`)

- **상태**: ✅ 작성 완료 (`episodes/08-socratic-learning-activity.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 8
- **Hook**: "알려줘"라고 말하는 순간, 대부분의 AI는 정답을 쏟아낸다. 하지만 Cocrates는 질문으로 답한다. 왜?
- **스토리라인**:
  1. **왜 소크라테스식 학습인가**: 일방적 정보 전달의 문제점
     - 사용자가 모른다는 사실조차 모르게 됨
     - 이해 없이 산출물을 수용하는 습관
     - 무지의 눈덩이
   2. **능동적 학습의 철학**: 소크라테스의 '무지의 지(知)' — 내가 모른다는 것을 아는 것이 지혜의 시작. Cocrates는 세 가지 교육공학적 핵심 철학을 기반으로 사용자가 스스로 학습할 수 있도록 대화를 이끈다.
      - **소크라테스 산파술 (Socratic Midwifery)**: 정답을 직접 주지 않고, 정교한 질문을 던져 사용자가 스스로 논리의 모순을 깨닫고 진리에 도달하게 한다. Cocrates가 "알려줘"에 정보를 쏟아내지 않고 질문으로 답하는 이유가 여기에 있다.
      - **블룸의 교육분류학 (Bloom's Taxonomy)**: 기억(Remember) → 이해(Understand) → 적용(Apply) → 분석(Analyze) → 평가(Evaluate) → 창조(Create)의 6단계를 밟아가며 질문의 깊이를 점진적으로 심화한다. 단순 개념 확인에서 시작해 스스로 설계하고 창조하는 단계까지 이끈다.
      - **근접발달영역 (ZPD) 및 비계 설정 (Scaffolding)**: 사용자의 현재 지식 수준을 파악하여 너무 쉽지도, 너무 어렵지도 않은 딱 한 단계 위의 과제(Mission)를 제시한다. 사용자가 성장함에 따라 힌트의 양을 줄여나간다. *"No pain, no gain"* — 적절한 인지적 부하 없이는 진정한 성장이 일어나지 않는다.
   3. **Learning 파이프라인 개관**:
     ```
     Education → Knowledge Capture → Reflection
     ```
   4. **각 단계의 역할 — 세 가지 학습 활동**:
      - **Education (질의응답 기반 학습)**: 사용자가 "알려줘", "설명해줘", "이게 뭐야"라고 요청하면 education 스킬을 활성화하여 소크라테스식 턴제 미션 기반 대화를 진행한다. 사용자는 개념 브리핑 → 생각의 실험실 → 미션의 3블록 구조로 학습하며, 매 턴 자신의 말로 답변하고 이해 공백(Gap)을 인식해야 한다.
      - **Knowledge Capture (학습 내용 정리)**: 사용자가 "정리해줘", "저장해줘"라고 요청하면 knowledge-capture 스킬이 배운 내용을 `kb/` 폴더에 recall용 핵심만 kebab-case 마크다운 파일로 저장한다. 상세 설명이 아닌, 나중에 떠올릴 수 있는 최소한의 핵심과 *"내가 틀렸던 가정"*을 기록한다.
      - **Reflection (이해도 평가)**: 사용자가 "평가해줘", "시험해줘"라고 요청하면 reflection 스킬이 저장된 KB와 대화 기록을 근거로 소크라테스식 질문을 던져 실제 이해도를 평가한다. 결과로 ✅ **잘 알고 있는 것**과 ⚠️ **제대로 알지 못했던 것**을 구분하여 보여준다. 사용자는 이를 근거로 추가 학습을 진행할 수 있다.
   5. **연결된 스킬**: education, knowledge-capture, reflection
   6. **사용자가 Cocrates와 학습하는 방법**
      - Cocrates의 답변을 수동적으로 받아들이지 않는다. 질문에 스스로 답하고, 제시된 과제(Mission)를 직접 수행함으로써 지식을 내재화한다.
      - "그냥 이해하는 수준"이 아니라, 생각하고, 의문을 제기하고, 추가 탐구를 통해 진정한 이해에 도달해야 한다. Cocrates는 사용자가 스스로 깨닫도록 질문할 뿐, 지식을 주입하지 않는다.
      - 세 단계(Education → Knowledge Capture → Reflection)를 순환하며 학습을 완성한다. 각 단계는 사용자의 적극적인 참여 없이는 작동하지 않는다.
      - **비유하자면 Cocrates와의 학습은 개인 트레이너와의 운동과 같다.** 트레이너가 운동법을 알려주고 자세를 교정해주지만, 실제로 움직이고 땀을 흘리는 것은 사용자 자신이다. Cocrates가 질문과 미션으로 길을 안내할지라도, 걸어가는 것은 사용자다.
- **핵심 메시지**: 소크라테스식 학습은 사용자가 스스로 깨닫고, 기록하고, 확인하는 3단계 파이프라인으로 구성된다. 목표는 지식 전달이 아니라 **무지의 제거**다. Cocrates는 소크라테스 산파술·블룸의 교육분류학·ZPD 비계 설정의 세 교육 철학을 기반으로 사용자가 능동적으로 학습에 참여하도록 이끈다.
- **포함할 내용**:
  - [ ] Learning 파이프라인 다이어그램
  - [ ] 기존 AI 학습 방식과의 비교
  - [ ] 교육공학적 핵심 철학 설명: 소크라테스 산파술, 블룸의 교육분류학, ZPD 및 비계 설정
  - [ ] Learning 파이프라인 세부 활동: Education / Knowledge Capture / Reflection 구체적 예시
  - [ ] 사용자 학습 태도 안내 — Cocrates와 함께 공부하는 방법
- **제외할 내용**:
  - 각 스킬의 내부 워크플로우 상세 (Ep9에서 다룸)
  - 실습 시나리오 (Ep4에서 다룸)
- **CTA**: 다음 편에서 Learning 파이프라인의 각 스킬을 자세히 살펴본다. 실습에서 경험한 학습 방식이 스킬 레벨에서 어떻게 구현되는지 알아보자.

---

### Ep9: 소크라테스식 학습 활동 - 스킬 (`09-socratic-learning-skills`)

- **상태**: ✅ 작성 완료 (`episodes/09-socratic-learning-skills.md`)
- **메타**: 예상 분량 12~15분, 발행 순서 9
- **Hook**: Education, Knowledge Capture, Reflection — 세 스킬이 어떻게 연결되어 하나의 학습 사이클을 만드는지 그 내부를 들여다본다.
- **스토리라인**:
  1. **Education Skill** — 소크라테스형 1:1 학습 코치
     - 일방 전달 금지 원칙
     - 턴제 미션 기반 대화 구조
     - Concept Briefing → Thought Laboratory → Mission 3블록 형식
     - Bloom's Taxonomy + 2-Axis Matrix: 난이도 조절
  2. **Knowledge Capture Skill** — 배운 것을 기록하다
     - recall용 핵심만 저장: 상세 설명 금지, 한 줄 정의 필수
     - 병합 전략: 동일 주제의 기존 KB 파일을 찾아 보완
     - '내가 틀렸던 가정'을 기록하는 구조
     - 무지(ignorance)의 기록
  3. **Reflection Skill** — 진짜 아는 것과 그냥 아는 척하는 것
     - Socratic evaluation: 단순 암기 확인이 아닌 적용·반례·경계 질문
     - KB를 평가 기준으로 사용
     - Gap 발견 시 education 모드로 전환
     - 최종 산출물: ✅ 잘 알고 있는 것 / ⚠️ 제대로 알지 못했던 것
  4. **세 스킬의 연결과 순환**:
     ```
     Education → Knowledge Capture → Reflection → (필요시) Education
     ```
- **핵심 메시지**: 세 스킬이 하나의 사이클로 연결되어 사용자의 학습을 능동적으로 이끈다. Education이 탐구를, Knowledge Capture가 기록을, Reflection이 검증을 담당한다.
- **포함할 내용**:
  - [ ] Education 스킬의 3블록 형식 예시
  - [ ] Knowledge Capture의 KB 저장 예시
  - [ ] Reflection 평가 산출물 예시
  - [ ] 실제 코드 경로: `.opencode/skills/education/SKILL.md` 등
- **제외할 내용**:
  - 실습 (Ep4에서 다룸)
  - 산출물 생성 스킬 (Ep10~11에서 다룸)
- **CTA**: 학습 활동을 살펴봤다. 이제 구조 기반 산출물 생성 활동으로 넘어간다.

---

### Ep10: 구조 기반 산출물 생성 활동 - 원리 (`10-architecture-driven-generation-activity`)

- **상태**: ✅ 작성 완료 (`episodes/10-architecture-driven-generation-activity.md`)
- **메타**: 예상 분량 12~15분, 발행 순서 10
- **Hook**: "보고서 작성해줘" — AI가 5초 만에 30페이지를 뚝딱 만들어낸다. 읽어보니 빈약하다. 어쩔 수 없다. Cocrates는 절대 그렇게 하지 않는다. 좋은 산출물은 구조에서 나오고, 구조는 결정에서 나오기 때문이다.

- **스토리라인**:

  1. **왜 구조(architecture)가 먼저인가**:
     - 구조 없이 생성된 산출물은 일관성과 논리가 부족함
     - 사용자가 산출물을 이해하고 설명할 수 없게 됨
     - 검토가 불가능한 블랙박스가 탄생
     - '그냥 작성해줘'는 '아무렇게나 작성해줘'와 같음

  2. **그런데, '구조'라는 것이 모호하다 — ASR과 ADR**:
     - "구조를 먼저 설계하라"는 말은 너무 추상적이다. 무엇을 설계해야 하는가?
     - 핵심은 **ASR(Architecturally Significant Requirement, 구조적으로 중요한 요구사항)** 을 식별하는 것
     - ASR이란 최종 산출물의 **구조, 구성, 품질에 실질적인 영향을 미치는 요구사항이나 설계 결정**을 말함
     - 소프트웨어뿐 아니라 문서, 발표자료, 블로그 시리즈 등 **모든 산출물 유형에 적용**되는 개념
     - ASR이 명시되지 않으면 생성 단계에서 **침묵의 기본값(silent default)** 으로 채워지고, 그 결과는 사용자가 통제할 수 없음
     - 이사를 앞두고 짐을 박스에 넣는 상황을 떠올려보자. '잘 넣어줘'라고 말하면 이삿짐센터 직원은 자신의 방식대로 넣는다. 중요한 유리그릇이 박스 아래에 깔릴 수도 있다. '구조'란 이런 식으로 중요한 것이 어디에 위치할지, 어떻게 보호될지, 어떤 순서로 배치될지를 결정하는 일이다. ASR은 '무엇이 유리그릇인지'를 식별하는 작업이고, ADR은 '유리그릇을 어떻게 포장할지' 선택지를 분석하고 결정하는 작업이다.

  3. **전원주택 비유로 이해하는 Concern → ADR → Spec → 생성**:

     - 전원주택을 짓는다고 생각해 보자.
     - **1층으로 지을 것인가, 2층으로 지을 것인가? 지붕을 올릴 것인가, 옥상을 둘 것인가, 옥탑방을 둘 것인가?** — 이들은 각각 하나의 **Concern(관심사)** 이다. 이 Concern은 ASR이다. 건물의 구조와 사용성에 실질적인 영향을 미치기 때문이다.
     - 만일 사용자가 *"2층에 옥탑방으로 지어주세요"* 라고 요청했다면, 이것은 **사용자의 요구사항**이다. 명확한 결정이므로 Spec에 바로 기록하면 된다.
     - 그런데 사용자가 *"우리 생활 패턴에 가장 적합한 것으로 결정해 주세요"* 라고 건축가에게 요청했다면? 이것이 ADR이 필요한 상황이다. 건축가는 각각의 결정(1층 vs 2층, 지붕 vs 옥상 vs 옥탑방)에 대해 **장/단점을 분석하고 사용자의 생활 패턴에 가장 적합한 것을 권장**한다. 사용자가 결정하면, 그것이 비로소 사용자의 요구사항이 된다.
     - 즉, **구조적으로 중요한 관심사(Concern/ASR)** 에 대해서는:
       1. 선택지를 분석하고 ADR 생성
       2. 사용자가 검토하고 결정
       3. 결정된 내용을 Spec으로 정리
       4. Spec에 따라 최종 산출물 생성
     - 이것이 Cocrates의 **가장 기본적인 구조 기반 산출물 생성 방법**이다.

  4. **Artifact Generation 파이프라인 개관**:

     ```
     ASR 식별 → ADR (대안 분석 & 결정) → Spec (결정 통합) 
         → Spec-driven Generation → Spec-driven Verification
     ```

     - **ASR 식별**: 생성 전에 '무엇이 구조적으로 중요한 결정인지'를 사용자와 함께 발견
     - **ADR**: 하나의 Concern에 대한 대안을 분석하고 사용자가 결정
     - **Spec**: 결정된 모든 요구사항과 제약을 하나의 자체 완결적 문서로 통합
     - **Spec-driven Generation**: 승인된 Spec만을 근거로 최종 산출물 생성
     - **Spec-driven Verification**: 결과물이 Spec과 일치하는지 항목별로 검증

  5. **각 단계의 역할 상세**:
     - **ASR 식별**: 침묵의 기본값을 방지한다. 생성 전에 '무엇이 중요한 결정인지'를 의식적으로 발견하는 과정이必须先 한다. (독자가 가져갈 역량: "이 요청에서 무엇이 구조적으로 중요한 결정인지 식별하는 눈")
     - **ADR (Any Decision Record)**: 대안을 구조적으로 비교하고, 선택의 근거를 기록한다. 사용자의 승인이 필수다. (독자가 가져갈 역량: "선택지를 보고 장단점을 비교한 뒤 자신의 결정을 명확히 하는 습관")
     - **Spec**: ADR에서 승인된 결정을 비롯한 모든 요구사항을 **ADR 없이도 이해할 수 있는** 자체 완결적 문서로 통합한다. (독자가 가져갈 역량: "결정된 사항을 하나의 설계도로 통합하는 능력")
     - **Spec-driven Generation**: Spec을 유일한 설계도로 삼아 생성한다. 원래 프롬프트가 아닌 Spec이 기준이다. (독자가 가져갈 역량: "명세서를 기준으로 AI에게 정확히 생성하도록 지시하는 능력")
     - **Spec-driven Verification**: 생성된 산출물이 Spec의 모든 항목을 만족하는지 하나하나 확인한다. 예상치 못한 설계 결정(Undocumented ASR)도 발견한다. (독자가 가져갈 역량: "결과물을 맹신하지 않고 명세서와 대조하여 검증하는 태도")

  6. **연결된 스킬**: adr-writing, spec-writing, spec-driven-generation, spec-driven-verification

- **핵심 메시지**: 구조 기반 생성의 핵심은 **'구조가 모호하다는 사실을 인정하는 것'** 에서 출발한다. 중요한 관심사(ASR)를 식별하고, 선택지를 분석·결정하고(ADR), 결정을 자체 완결적 명세로 통합한 뒤(Spec), 그 명세를 기준으로 생성하고 검증하는 파이프라인이 구조 기반 생성을 완성한다.

- **포함할 내용**:
  - [ ] ASR 개념 설명과 전산물 유형별 예시
  - [ ] 전원주택 비유 — Concern/ADR/Spec/생성의 관계를 이해시키는 스토리텔링
  - [ ] Artifact Generation 파이프라인 다이어그램
  - [ ] 구조 기반 접근 vs 무구조 접근 비교 (침묵의 기본값이 만드는 차이)
  - [ ] 각 단계에서 독자가 가져갈 역량 포인트

- **제외할 내용**:
  - 각 스킬의 상세 워크플로우 (Ep11에서 다룸)
  - 문서/발표자료/블로그 시리즈 등 artifact-specific 스킬과의 관계 (Ep11에서 다룸)
  - 스킬이 없을 때의 처리 (Ep12에서 다룸)
  - 실습 시나리오 (Ep5에서 다룸)

- **CTA**: 다음 편에서 Artifact Generation 파이프라인의 각 스킬 — adr-writing, spec-writing, spec-driven-generation, spec-driven-verification — 의 핵심 내용을 자세히 살펴본다. 이번 편에서 구조 기반 생성의 '왜'를 이해했다면, 다음 편에서는 '어떻게'를 알게 될 것이다.

---

### Ep11: 구조 기반 산출물 생성 활동 - 스킬 (`11-architecture-driven-generation-skills`)

- **상태**: ✅ 작성 완료 (`episodes/11-architecture-driven-generation-skills.md`)
- **메타**: 예상 분량 14~18분, 발행 순서 11
- **Hook**: Ep6에서 구조 기반 생성의 '왜'를 이해했다. 이제 4개의 스킬이 각각 무엇을 담당하고, 어떻게 연결되어 하나의 파이프라인을 만드는지 그 핵심 내부를 들여다본다. 각 스킬은 고유한 원칙과 금지 사항을 가지며, 그 경계가 파이프라인의 품질을 결정한다.

- **스토리라인**:

  1. **adr-writing (Any Decision Record)** — 왜 이런 결정을 내렸는가

     - **ADR = Architecture Decision Record가 아니라 Any Decision Record**. 하나의 Concern(관심사)에 대해 대안을 분석하고, 근거와 함께 사용자가 결정을 내리도록 돕는다.
     - **핵심 구조**: 하나의 ADR 파일 = 하나의 Concern. Concern이 다르면 별도 파일.
     - **Option 제시 원칙**: 최소 2~3개의 실질적 대안을 제시 — 가짜 대안이나 하나뿐인 선택지는 금지
     - **대안 서술 방식**: 장황한 산문 금지, 간결한 bullet으로 요약 (각 Option: 이름 + 2~4개 bullet)
     - **사용자 승인 게이트**: `proposed` 상태로 저장 → 사용자 검토 → 명시적 승인 → `approved`로 변경 → Spec으로 이관
     - **저장 위치**: `adr/{concern-slug}.md` (프로젝트 루트의 `adr/` 디렉토리)
     - **병합 원칙**: 동일/유사 Concern의 기존 ADR이 있으면 새로 만들지 않고 병합하거나 대체(superseded) 처리
     - **금지 사항**: 2개 미만의 Option 제시, 명시적 승인 없이 `approved` 처리, 동일 Concern 중복 파일
     - **역할**: ADR은 **의사결정 감사 추적(audit trail)** 이다. Spec이 '무엇을' 만드는지를 담는다면, ADR은 '왜 그렇게 결정했는지'를 담는다.

  2. **spec-writing** — 결정된 요구사항의 통합

     - **Spec의 핵심 원칙 — 자체 완결성(Self-Containment)**: Spec 하나만 읽으면 '무엇을 만들지'를 완전히 이해할 수 있어야 한다. ADR을 열어볼 필요가 없다.
     - **ADR과 Spec의 역할 분담**:
       - ADR = 왜 그 결정을 내렸는지 (선택지 분석, 근거, 감사 추적)
       - Spec = 무엇을 만들기로 결정했는지 (명확한 요구사항, 제약, 범위)
     - **Spec은 살아있는 문서(living document)**: 승인 상태를 추적하지 않음 — 사용자가 직접 편집 가능. `Status`나 `Approved` 필드 금지.
     - **서술 원칙**: 산문 금지, bullet 위주로 간결하게. 하나의 bullet = 하나의 요구사항·결정·제약. 검증 가능한(verifiable) 문장으로 작성.
     - **저장 위치**: `spec/{requirement-slug}.md` (프로젝트 루트의 `spec/` 디렉토리)
     - **ADR → Spec 이관 시 규칙**: ADR의 승인된 결정을 '완전히 명시된 사양'으로 옮긴다. 'ADR의 Option B 참조' 같은 링크는 금지 — Spec만으로 생성 가능해야 함.
     - **열린 질문(Open Questions) 처리**: 아직 결정되지 않은 사항은 기록하고, 생성이 이를 기다려야 하는지 사용자와 합의
     - **금지 사항**: 사용자가 확인하지 않은 요구사항 추가, ADR 링크로 Spec 내용 대체, 장황한 설명, 승인 상태 필드

  3. **spec-driven-generation** — 승인된 Spec만으로 생성

     - **Spec이 유일한 근거**: 원래 사용자 프롬프트가 아니라 Spec이 생성의 기준이다. 프롬프트와 Spec이 충돌하면 Spec을 따른다.
     - **ASR(Architecturally Significant Requirement) Readiness Gate**: Spec이 없으면 무작정 생성하지 않는다. 먼저 ASR을 식별하고, 필요한 경우 ADR과 Spec 작성 과정을 거친다. 이것이 Cocrates가 '생성' 요청을 받았을 때 실제로 하는 첫 번째 일이다.
     - **ASR의 8가지 보편 범주**: 목적과 독자(Purpose & Audience), 산출물 형태(Deliverable Form), 범위 경계(Scope Boundary), 품질 기준(Quality Bar), 제약(Constraints), 구조와 구성(Structure & Organization), 통합과 의존성(Integration & Dependencies), 프로세스와 단계(Process & Staging)
     - **도메인별 ASR**: 소프트웨어, 문서/보고서, 발표자료, 이미지/시각자료, 콘텐츠 시리즈, 데이터/스키마 등 산출물 유형별 추가 고려사항
     - **Spec이 있어도 생성 전 확인사항**: 미해결 Open Questions 차단 여부, Spec bullet이 ADR을 참조하는지, 모호한 요구사항이 있는지, ASR이 누락되었는지. 하나라도 있으면 생성 전에 해결.
     - **artifact-specific skill과의 관계**: document-authoring, presentation-authoring, blog-series-authoring 같은 구체 산출물 스킬이 있으면 그 스킬의 단계별 게이트를 따르되, Spec이 최종 산출물의 합격 기준을 정의한다. Snowflake 점진 구체화(개요 → 구조 → 설계 → 생성)는 이들 스킬에서 구현된다.
     - **금지 사항**: 준비되지 않은 Spec으로 생성, ADR을 Spec 대신 읽어서 생성, 프롬프트를 Spec보다 우선시, 명시되지 않은 기능/범위 추가, ASR 검토 생략

  4. **spec-driven-verification** — 결과물 검증

     - **검증의 이중 목적**: (1) 결과물이 Spec과 일치하는지 항목별 확인, (2) 생성 과정에서 Spec에 없던 **Undocumented ASR(문서화되지 않은 구조적 요구사항)** 발견
     - **Undocumented ASR이란**: Spec에 명시되지 않았지만 생성 결과에 구현된 구조적 결정. 예를 들어 '보고서를 작성해줘'라고만 했는데 AI가 특정 섹션 순서로 작성했다면, 그 섹션 순서는 검토되지 않은 ASR이다.
     - **검증 프로세스**: Spec의 모든 항목을 인벤토리로 추출 → 각 항목을 개별 확인(pass/fail/partial) → Spec 이탈(Deviation) 집계 → Undocumented ASR 스캔 → 검증 보고서 저장
     - **저장 위치**: `verification/{requirement-slug}.md` (프로젝트 루트의 `verification/` 디렉토리)
     - **검증 보고서 구성**: Spec 항목별 상태와 증거, Deviation 목록(심각도 포함), Undocumented ASR 목록(위험도와 권장 조치 포함)
     - **사용자 검토 게이트**: 검증 보고서를 저장한 후 사용자가 검토하고 수정/확정 결정을 내릴 때까지 기다림. 무조건 수정하지 않음.
     - **Deviation과 Undocumented ASR의 차이**: Deviation = Spec에 명시된 요구사항을 위반. Undocumented ASR = Spec에 없었지만 결과물에 구현된 결정. 전자는 수정이 원칙, 후자는 확인(confirm), Spec 업데이트, 또는 수정 중 선택.
     - **금지 사항**: Spec이 아닌 ADR이나 대화 기록으로 검증, 인벤토리 없이 '대충 확인', 항목별 증거 없이 pass 처리, Undocumented ASR 스캔 생략, 사용자 검토 전에 무단 수정

  5. **파이프라인의 유연성과 artifact-specific 스킬과의 관계**:
     - 기본 파이프라인은 **ADR → Spec → Generation → Verification**이지만, 산출물 유형에 따라 변형 가능
     - **document-authoring** 스킬: 문서 구조를 개요 → 섹션 설계 → 섹션 생성 → 검토 단계로 점진 구체화. Spec이 문서 구조를 정의하면, 이 스킬이 문서 특화 워크플로우(서론/본론/결론 구성, 각 섹션의 훅과 CTA 등)를 처리한다.
     - **presentation-authoring** 스킬: 발표자료를 문서 → 섹션 → 페이지 계층으로 점진 구체화. 각 페이지는 거버닝 메시지(두괄식) + 서포트 구조를 가진다.
     - **blog-series-authoring** 스킬: 블로그 시리즈를 개요 → 아웃라인 → 에피소드 설계 → 에피소드 본문 순으로 점진 구체화. 각 단계마다 파일로 저장하고 사용자 승인을 받는다.
     - **공통 원칙**: 모든 artifact-specific 스킬은 Snowflake 방식의 점진 구체화, 단계별 사용자 승인 게이트, Architecture-driven 접근을 공유한다.

- **핵심 메시지**: ADR(선택지 분석과 결정) → Spec(자체 완결적 명세 통합) → Generation(ASR Readiness Gate를 통과한 Spec 기반 생성) → Verification(항목별 검증 + Undocumented ASR 발견)의 4단계가 구조 기반 생성을 완성한다. 각 스킬은 고유한 원칙과 금지 사항으로 파이프라인의 품질을 보장하며, artifact-specific 스킬(document/presentation/blog-series authoring)은 이 기본 파이프라인을 산출물 유형에 맞게 구체화한다.

- **포함할 내용**:
  - [ ] 각 스킬의 핵심 원칙 한 줄 요약 (ADR = 선택과 근거 / Spec = 자체 완결성 / Generation = Readiness Gate / Verification = Undocumented ASR 발견)
  - [ ] ADR vs Spec 비교 표 (단위, 목적, 독자가 필요한 것, 파이프라인 역할)
  - [ ] ASR 8가지 보편 범주 표
  - [ ] Undocumented ASR 개념 설명과 예시
  - [ ] artifact-specific 스킬 3종(document/presentation/blog-series)의 특징과 언제 사용하는지
  - [ ] 실제 코드 경로: `.opencode/skills/adr-writing/SKILL.md`, `spec-writing/SKILL.md` 등

- **제외할 내용**:
  - generating-skill-creation (Ep12에서 다룸)
  - 각 스킬의 전체 워크플로우 상세 (너무 깊은 단계까지는 생략)
  - 실습 시나리오 (Ep5에서 다룸)

- **CTA**: 이 4개의 스킬이 구조적 생성 파이프라인의 핵심이다. 이는 일반적인 산출물에 대한 가장 일반적인 구조적 접근이다(**spec-driven-generation**). 그런데, 산출물에 따라서 구조적인 접근을 명확히 할 수 있다. 이때는 산출물에 맞는 구조적 접근을 스킬로 명세하고, 명세된 스킬에 따라서 생성하는 것이 좋다. 다음편에서는 **generating-skill-creation** 스킬의 핵심 요소를 알아본다.

---

### Ep12: 구조 기반 워크플로우 생성 스킬  (`12-workflow-creation-skill`)

- **상태**: ✅ 작성 완료 (`episodes/12-workflow-creation-skill.md`)
- **메타**: 예상 분량 13~16분, 발행 순서 12
- **Hook**: **generating-skill-creation**는 스킬을 생성하는 **메타-스킬** 이다. 산출물을 생성하는 **절차 그 자체를 설계**한다. 이것은 Cocrates Harness에서 가장 추상적이면서도 강력한 활동이다.

- **스토리라인**:

   1. **산출물을 분해하는 다섯 가지 구성 요소(Artifact Components)**:
      - 새 스킬을 설계할 때 generating-skill-creation은 최종 산출물을 다섯 가지 차원으로 분해한다. 산출물 유형에 따라 이름은 달라질 수 있지만 역할은 동일하다:

      | 구성 요소 | 역할 |
      |-----------|------|
      | **과업과 제약 (Assignment & Constraints)** | 범위, 품질 기준, 금기 사항, 성공 조건 |
      | **맥락과 규칙 (Context & Rules)** | 도메인 규칙, 세계관, 스타일 가이드 |
      | **개체 (Entities)** | 재사용 가능한 구성 요소 |
      | **공간과 배치 (Space & Placement)** | 위치, 장면, 레이아웃 영역 |
      | **구조와 흐름 (Structure & Flow)** | 계층적 콘텐츠 구성 |

      - 스킬을 설계할 때는 이 다섯 가지 차원을 모두 식별하고, 각각이 어느 단계에서 어느 수준까지 구체화될지 결정한다.

   2. **Snowflake 5단계 — 각 단계의 목적과 승인 게이트**:
      - generating-skill-creation은 새 스킬의 생성 절차를 **define → plan → architecture design → detail design → generation**의 5단계로 점진 구체화한다:

      | 단계 | 목적 | 승인 게이트 |
      |------|------|------------|
      | **define** | 과업, 범위, 제약, 성공 기준 확정 | 모든 후속 결정의 기준선 |
      | **plan** | 스킬의 골격, 방향, 주요 구성 요소 개관 | 전체 리듬과 방향 승인 |
      | **architecture design** | 구성 요소별 카탈로그, 계층 구조, 상호 참조 | 구조적 일관성 승인 |
      | **detail design** | 생성에 필요한 단위별 명세 확정 | **생성 전 최종 게이트** |
      | **generation** | 확정된 명세에 따라 SKILL.md 생성 및 조립 | 단위별 및 최종 품질 검토 |

      - **절대 규칙**: detail design이 완전히 확정되기 전까지 generation 단계로 넘어가지 않는다. 이 게이트가 없으면 '그냥 만들어보고 수정'하는 일회성 프롬프트와 다를 바 없다.

   3. **단계별 구성 요소 구체화(Per-Stage Refinement)**:
      - 각 구성 요소는 동일한 속도로 구체화되지 않는다. generating-skill-creation이 새 스킬을 설계할 때 **각 구성 요소가 각 단계에서 어느 수준까지 구체화되는지**를 정의하는 것이 핵심이다.
      - 예: **개체(Entities)** 구성 요소의 구체화 과정
        - define: 개체 유형만 식별
        - plan: 핵심 역할과 관계 요약
        - architecture design: 카탈로그, 관계 맵, 계층 구조
        - detail design: 개체별 상세 명세, 변형, 상태
        - generation: 개체별 산출물 생성
      - 구조와 흐름이 개체나 공간을 참조할 때는 먼저 카탈로그를 완성하고, 상세 명세는 필요할 때 채우는 **지연 구체화(lazy refinement)** 전략을 사용한다.

   4. **스킬 저작 절차(Meta Snowflake) — Kernel → Components → Frame → Outline → Spec → Skill → Verification**:
      - generating-skill-creation 자신이 새 SKILL.md를 만들 때는 다음 7단계로 진행한다:

      | 단계 | Snowflake 대응 | 산출물 |
      |------|---------------|--------|
      | **Kernel** | define | 생성 대상을 한 문장으로 정의 ("이 스킬은 ~을 생성하도록 돕는다") |
      | **Components** | plan | 다섯 가지 구성 요소 차원 식별 |
      | **Frame** | architecture design | 5단계 Snowflake 워크플로우, 파일 구조, 승인 지점 설계 |
      | **Outline** | detail design | 각 단계별 파일 산출물, 입력, 생성 활동, 완료 기준, 검토 질문, 승인 조건 정의 |
      | **Spec** | generation | frontmatter, 원칙, 사용 시점, 절차, 대화 규칙, 금지 사항, 완료 조건 확정 |
      | **Skill** | generation | `.opencode/skills/{skill-slug}/SKILL.md` 파일 생성 |
      | **Verification** | generation | 스킬이 생성 하네스로 작동하는지 검증 |

   5. **파일 산출물 정의 방식(Outline Format)**:
      - 새 스킬의 각 단계에서 생성되는 중간 파일은 다음 형식으로 정의된다. 이는 중간 산출물이 채팅에만 남지 않고 **파일로 관리**되도록 보장한다:
        - **Input**: 의존하는 산출물
        - **Creation activity**: 이 파일을 생성하기 위한 작업
        - **Completion criteria**: 이 파일이 완료된 것으로 간주되는 조건
        - **Review questions**: 사용자 확인을 위한 질문
        - **Approval criteria**: 다음 단계로 진행하기 위한 승인 조건

   6. **대화 규칙과 금지 사항 — Cocrates가 스킬을 설계하는 방식**:
      - **대화 규칙**:
        - 현재 단계(Kernel/Components/Frame/Outline/Spec/Skill/Verification)를 항상 명시
        - Frame을 설계할 때는 5단계 Snowflake와 단계별 구성 요소 구체화 표를 함께 제시
        - 사용자 승인 없이 다음 단계로 진행하지 않음
        - 구성 요소나 Frame이 불명확하면 먼저 질문
      - **금지 사항**:
        - 적합한 생성 스킬이 없는데 최종 산출물부터 생성
        - 산출물 구조 분석 없이 SKILL.md 템플릿만 채우기
        - 사용자 검토 지점이 없는 생성 스킬 작성
        - 중간 산출물을 채팅에만 남기고 파일 저장 규칙 누락
        - 기존 스킬과 중복되는 새 스킬 생성
        - `compatibility: opencode`와 `metadata.agent: cocrates` 누락

   7. **스킬의 재사용성과 진화**:
      - 한 번 만든 스킬은 이후 동일 유형 요청에서 generating-skill-creation의 개입 없이 재사용
      - artifact-specific skill(document-authoring, presentation-authoring, blog-series-authoring)은 모두 이 방식으로 탄생
      - 새 산출물 유형이 필요하면 스킬만 추가하면 되고, 기존 스킬은 독립적으로 개선 가능

- **핵심 메시지**: generating-skill-creation은 **스킬을 만드는 스킬(Meta-Skill)** 이다. 최종 산출물을 다섯 가지 구성 요소로 분해하고, Snowflake 5단계(define → plan → architecture design → detail design → generation)로 점진 구체화하며, Meta Snowflake 7단계(Kernel → Components → Frame → Outline → Spec → Skill → Verification)로 저작하여 사용자 승인을 거쳐 SKILL.md를 생성한다.

- **포함할 내용**:
  - [ ] 다섯 가지 산출물 구성 요소(Artifact Components)와 각 차원의 역할 표
  - [ ] Snowflake 5단계 — 각 단계의 목적과 승인 게이트 표
  - [ ] 단계별 구성 요소 구체화(Per-Stage Refinement) 개념 — 각 구성 요소가 단계별로 어느 수준까지 구체화되는지
  - [ ] Meta Snowflake 7단계 (Kernel → Components → Frame → Outline → Spec → Skill → Verification) 표
  - [ ] 파일 산출물 정의 방식(Outline Format) — Input / Creation activity / Completion criteria / Review questions / Approval criteria
  - [ ] 대화 규칙과 금지 사항 — generating-skill-construction의 동작 원리
  - [ ] 스킬의 재사용성과 진화 — artifact-specific skill의 탄생 과정
  - [ ] 실제 코드 경로: `.opencode/skills/generating-skill-creation/SKILL.md`

- **제외할 내용**:
  - 실습 (Ep6에서 다룸)
  - 설치 방법 (Ep3에서 다룸)
  - 개별 artifact-specific 스킬의 상세 워크플로우 (Ep11에서 이미 다룸)

- **CTA**: generating-skill-creation까지 이해했다면, Cocrates Harness의 모든 구조적 활동을 살펴본 셈이다. 이제 마지막으로, Cocrates를 지속적으로 발전시키는 방법과 사용자로서의 다짐을 알아보겠다.

---

## 4. 결론 — 지속적 개선과 선언

### Ep13: 올바른 Cocrates Harness 활용 (`13-evolution-of-cocrates`)

- **상태**: ✅ 작성 완료 (`episodes/13-evolution-of-cocrates.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 13
- **Hook**: Cocrates를 설치하고, 실습까지 마쳤다. 구조의 원리까지 이해했다. 이제 끝일까? 아니다. 진짜 사용은 지금부터 시작이다.
- **스토리라인**:
  1. **사용자가 진화해야 한다** — Harnessing Ignorance, Cocrates Harness의 핵심 원칙이다.
     - Cocrates와의 대화에 익숙해지기
     - 소크라테스식 질문을 스스로에게도 던지기
     - 검토와 승인의 습관화
  2. **Cocrates Harness를 진화시켜야 한다**
     - Skills는 고정된 것이 아니다: 추가·수정·확장 가능
     - 사용자의 업무 패턴에 맞는 스킬 설계
     - 경험에 따라 agent prompt 개선
     - agent prompt, skill이 모두 공개되어 있다. opencode plugin이 아니더라도 claude code, cursor, codex, antigravity 등에도 직접 구현할 수 있다.
  3. **함께 진화하는 관계**
     - 사용자가 성장할수록 Cocrates도 더 효과적으로 동작
     - 피드백 루프: 사용 경험 → agent/skill 개선 → 더 나은 경험
  4. **지속적 개선의 원칙**
     - 완벽한 첫 시도는 없다. 계속 고쳐나가는 것이 중요하다.
     - "The unexamined code is not worth generating" — 여기에는 에이전트 자신도 포함된다.
  5. **open source**
     - agent prompt, skill이 모두 공개되어 있다.
     - opencode plugin이 아니더라도 claude code, cursor, codex, antigravity 등에도 직접 구현할 수 있다.
     - 또는 나만의 하네스를 구축하는 것도 매우 바람직하다.
- **핵심 메시지**: Cocrates는 완성된 제품이 아니라 **함께 진화하는 도구**다. 사용자도 Cocrates도 지속적인 개선이 필요하다.
- **포함할 내용**:
  - [ ] 스킬 추가/수정 예시
  - [ ] Agent 프롬프트 개선 예시
  - [ ] 피드백 루프 다이어그램
- **제외할 내용**:
  - 설치 방법 (Ep3에서 다룸)
  - 개별 실습 (Ep4~6에서 다룸)
- **CTA**: 마지막 편은 선언문이다. 당신이 AI를 올바르게 쓰는 사람임을 선언하라.

---

### Ep14: Cocrates Harness 사용자 선언문 (`14-user-manifesto`)

- **상태**: ✅ 작성 완료 (`episodes/14-user-manifesto.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 14
- **Hook**: 지금까지의 여정을 되돌아보며, 당신은 하나의 선언을 준비한다.
- **스토리라인**:
  1. **시리즈의 핵심 — 세 문장으로 압축**
     - **Cocrates의 핵심 원칙**: *"The unexamined code is not worth generating."* — 검토되지 않은 산출물은 생성할 가치가 없다. 이 원칙은 소스 코드뿐 아니라 모든 AI 최종 산출물에 적용된다.
     - **Socrates의 지혜**: *"I know I know nothing."* — 내가 아는 것이 없다는 것을 아는 것이 지혜의 시작이다. 무지의 인식이 곧 탐구의 출발점이다.
     - **Harness Ignorance**: 무지를 두려워하지 말고, 오히려 그것을 활용하라. 무지를 인식하는 순간, 당신은 그것을 채울 기회를 얻는다. Cocrates Harness는 바로 이 **무지의 활용(Harnessing Ignorance)** 을 지원하기 위해 설계되었다.
  2. **사용자 선언 7계명** — 하나씩 선언한다
     - *1. I do not blindly request.*
     - *2. I do not blindly trust.*
     - *3. I do not surrender control.*
     - *4. I do not copy and run.*
     - *5. I do not hide my ignorance.*
     - *6. I always reflect and learn.*
     - *7. I always fight for my growth.*
  3. **선언문을 마음속 깊이 새긴다**
     - 단순히 읽고 넘어가는 것이 아니다. 이 선언문은 앞으로 AI와의 모든 상호작용에서 당신의 태도를 규정하는 기준이 된다.
     - 의심될 때, 이 선언문으로 돌아온다. "지금 나는 이 선언문에 따라 행동하고 있는가?"
  4. **Cocrates Harness를 사용하여 당신의 능력을 극대화하라**
     - Cocrates Harness는 당신의 도구다. 이 선언문은 당신의 의지다.
     - 도구와 의지가 하나가 될 때, AI는 단순한 생성기가 아니라 **당신의 능력을 극대화하는 파트너**가 된다.
     - 이제 당신의 차례다.
- **핵심 메시지**: Cocrates Harness의 궁극적인 목표는 사용자가 AI를 올바르게 활용하여 **자신의 능력을 극대화**하는 것이다. 그 출발점은 핵심 원칙(The unexamined code is not worth generating)과 소크라테스의 지혜(I know I know nothing)에 있다. 사용자 선언 7계명은 이 여정의 나침반이다.
- **포함할 내용**:
  - [ ] 세 가지 핵심 문장 — Cocrates 원칙 / Socrates / Harness Ignorance 연결
  - [ ] 사용자 선언 7계명 (한글 + 영문)
  - [ ] 각 계명의 간단한 설명
  - [ ] 마무리 메시지 — Cocrates Harness로 능력을 극대화하라
- **제외할 내용**:
  - 새로운 개념이나 기술 설명
  - 실습
  - Cocrates Agent의 약속 (선언문에 집중)
- **CTA**: 이제 당신 차례다. 이 선언문을 마음속 깊이 새기고, Cocrates Harness와 함께 당신의 능력을 극대화하라.

## 완료 기준

| 항목 | 기준 |
|------|------|
| **각 편 분량** | 읽기 8~15분 / 한국어 2,500~4,500자 |
| **검토 완료** | 사용자가 각 편을 읽고 구조·내용 승인 |
| **게시 준비** | 모든 편 검토 완료 후 최종 통합 리뷰 |
| **전체 상태** | 14개 에피소드 모두 작성 완료 ✅ |
