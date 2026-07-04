# 에피소드 설계

---

## 1. 소개 — Cocrates Harness의 필요성과 원칙

### Ep1: 같은 LLM, 다른 결과 (`01-same-llm-different-results`)

- **상태**: ✅ 작성 완료 (`episodes/01-same-llm-different-results.md`)
- **Hook**: 두 개발자가 같은 Claude에게 "로그인 모듈 만들어줘"를 입력한다. 한 명은 그대로 붙여넣고, 다른 한 명은 구조를 검토하고 수정한다. 차이는 LLM이 아니라 **사용자의 태도**에 있다. "당신의 AI는 몇 명의 팀원입니까?"
- **스토리라인**:
  1. **도입**: 같은 LLM, 다른 결과 — AI-assisted(도구) vs AI-native(팀원)의 극명한 차이
  2. **AI-assisted Engineer의 사례**: AI에게 시키는 대로 일을 시킴, 검토 생략, 코드는 나오지만 이해도는 낮아짐
  3. **AI-native Engineer의 사례**: AI와 함께 고민하고 구조를 먼저 설계함. AI가 내놓은 결과를 **검토(U→A→E→A)** 한 후 승인
  4. **OpenAI 실제 사례 (Ryan Lopopolo 팀)**: 3명이 5개월 만에 100만 줄, Zero-Code 애플리케이션 구축 — AI가 팀원이 될 때의 가능성을 증명
  5. **기술이 아니라 태도의 문제**: 사용자의 태도가 AI의 활용도를 결정한다. AI-native Engineer는 AI를 도구가 아닌 **일을 수행하는 팀원**으로 대한다
  6. **Cocrates 소개 — Co + Socrates**: 상호 소크라테스 관계. 사용자가 AI에게 소크라테스가 되고(검토·판단), AI가 사용자에게 소크라테스가 되어(무지를 깨닫게 함) 성장을 돕는 에이전트
- **핵심 메시지**: 모델이 아니라 **당신의 태도**가 차이를 만든다. AI-assisted(도구)에서 AI-native(팀원)로 전환하라. AI가 100명이 될지, 1명이 될지는 당신에게 달려 있다. Cocrates는 이 전환을 돕기 위해 설계되었다.
- **CTA**: 다음 편에서 핵심 원칙을 선언한다 — "검토되지 않은 산출물은 생성할 가치가 없다."

---

### Ep2: The Unexamined Code Is Not Worth Generating (`02-unexamined-code`)

- **상태**: ✅ 작성 완료 (`episodes/02-unexamined-code.md`)
- **Hook**: *"The unexamined code is not worth generating."* — 검토되지 않은 산출물은 생성할 가치가 없다. 이 문장은 소크라테스의 "The unexamined life is not worth living"에서 영감을 받았다. 삶이 그러하듯, AI가 만든 산출물도 **검토되지 않으면 생성할 가치가 없다.**
- **스토리라인**:
  1. **선언**: 시리즈를 관통하는 한 문장 — "The unexamined code is not worth generating."
  2. **'Code'의 확장 정의**: 소스 코드뿐 아니라 AI가 만드는 모든 최종 산출물 — 문서, 보고서, 슬라이드, 블로그, 학습 노트 등
  3. **검토(U→A→E→A)의 4단계**:
     - **U(Understand / 이해)**: AI의 산출물을 일고 이해한다. 모르는 개념이나 용어가 있으면 먼저 파악한다.
     - **A(Analyze / 분석)**: 산출물의 구조와 논리, 가정을 분석한다. "이게 정말 맞는가?"를 검증한다.
     - **E(Evaluate / 평가)**: 산출물이 내 요구사항과 품질 기준을 만족하는지 평가한다. 대안과 비교한다.
     - **A(Approve / 승인)**: 모든 검토가 끝난 후에만 승인한다. 승인하지 못하겠다면 다시 설계로 돌아간다.
  4. **검토의 본질 — 무지의 제거(Harnessing Ignorance)**: AI가 구조를 설계해도 내가 이해하지 못했다면 그것은 검토된 것이 아니며, 내게 무지만 남긴다. 검토의 궁극적 목적은 무지를 제거하는 것이다
  5. **Cocrates = 가드레일**: 사용자가 AI를 주도적으로 사용하고, 이해하지 못한 채 넘어가지 않도록 안내하는 가드레일. 사용자의 주도권을 지켜주는 리듬
  6. **이 블로그 시리즈를 만드는 과정을 예시로**: 이 시리즈 자체가 Architecture-Based Design → Review → Approval의 과정을 거쳐 만들어졌다
- **핵심 메시지**: 검토되지 않은 산출물은 생성할 가치가 없다. 검토한다는 것은 **이해(U) + 분석(A) + 평가(E) + 승인(A)** — 내가 이해하고, 내가 판단하고, 내가 승인하는 것이다. 검토의 본질은 무지의 제거다.
- **CTA**: 다음 편에서 Cocrates Harness를 설치하고 첫 대화를 시작한다.

---

## 2. 실습 — 먼저 경험한다

### Ep3: Cocrates Harness 설치 (`03-installation`)

- **상태**: ✅ 작성 완료 (`episodes/03-installation.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 3
- **Hook**: 이제까지 Cocrates가 무엇인지, 어떤 원칙으로 움직이는지 알았다. 그렇다면 어떻게 시작해야 할까? Cocrates Harness는 **opencode 플랫폼 위에서 동작하는 플러그인(Plugin)** 이다. 설치는 단순하지만, 선택지가 있다. 그리고 설치조차 **검토(U→A→E→A)하라는 Cocrates의 원칙**을 기억하라.
- **스토리라인**:
  1. **Cocrates = Plugin + Skill**
     - Cocrates Harness는 opencode 플러그인으로 설치되는 **Cocrates Plugin**과, 워크플로우를 제공하는 **Skills**로 구성된다
     - opencode = 극장(플랫폼), Cocrates = 배우(플러그인), Skills = 대본(워크플로우)
  2. **3가지 UI 선택지 (사용자 환경에 따라 선택)**:
     - **opencode 터미널**: Claude 스타일의 TUI, 터미널 환경에 익숙한 사용자
     - **opencode 데스크탑**: Cursor/VS Code 같은 GUI 환경을 선호하는 사용자
     - **VS Code + OpenChamber 확장**: 기존 VS Code 환경을 유지하면서 OpenChamber 확장으로 opencode 사용
  3. **설치의 3단계**:
     - 1단계: opencode 설치 (https://opencode.ai/download)
     - 2단계: Cocrates Plugin 설치 — `opencode.jsonc`의 `plugin` 항목에 `"@cocrates/cocrates-harness"` 추가
     - 3단계: Skills 복사 — `~/.config/opencode/skills/`에 모든 스킬 파일 복사
  4. **치트키 설치법**: opencode 설치 후, 이 문서를 근거로 **"Cocrates Harness 설치해줘"**라고 AI에게 요청하면 AI가 플러그인 설치와 스킬 복사를 도와준다. AI를 AI-native 방식으로 활용하는 첫걸음
  5. **설치 확인 및 첫 대화**: Cocrates Agent 활성화 확인 후 "블룸의 분류학에 대해 알려줘" 같은 질문으로 Cocrates의 소크라테스식 응답 경험
- **핵심 메시지**: Cocrates Harness = Plugin + Skill. 설치는 3단계(opencode → plugin → skills)로 단순하다. 하지만 진짜 시작은 설치 후다. **설치조차 검토(U→A→E→A)하라는 원칙**을 첫 대화에서 경험하게 될 것이다.
- **포함할 내용**:
  - [ ] opencode 설치 옵션 비교 (터미널 / 데스크탑 / VS Code+OpenChamber)
  - [ ] Plugin + Skill 구조 설명과 opencode=극장 비유
  - [ ] config 파일 경로 확인 방법 (`~/.config/opencode/` vs `OPENCODE_CONFIG_DIR`)
  - [ ] opencode.jsonc 설정 예시 (plugin 항목)
  - [ ] 치트키 설치법: AI에게 "Cocrates Harness 설치해줘" 요청
  - [ ] skills 복사 경로와 방법
  - [ ] 설치 후 첫 대화 추천 질문
- **제외할 내용**:
  - 스킬 설계 방법 (Ep13에서 다룸)
  - 실습 시나리오 (Ep4~6에서 다룸)
  - 각 스킬의 상세 내용 (Ep9, Ep11에서 다룸)
- **CTA**: 설치가 끝났다면, 이제 Cocrates에게 "블룸의 분류학에 대해 알려줘"라고 물어보라. 설치만으로는 아무 일도 일어나지 않는다. Cocrates는 당신의 참여를 기다린다.

---

### Ep4: 소크라테스식 학습 활동 (`04-socratic-learning-practice`)

- **상태**: ✅ 작성 완료 (`episodes/04-socratic-learning-practice.md`)
- **메타**: 예상 분량 14~18분, 발행 순서 4
- **Hook**: "블룸의 분류학에 대해 알려줘" — 평소 같았으면 위키피디아식 설명을 받았을 것이다. 하지만 Cocrates는 **질문으로 답했다**. 그리고 그 과정에서 블룸 분류학을 **완전히 다르게** 이해하게 되었다. **같은 질문, 완전히 다른 답변** — 이 편은 그 여정을 함께 따라가는 실습이다.
- **스토리라인**:
  1. **실습 시나리오**: Bloom's Taxonomy를 Cocrates와 함께 배우기
     - "블룸의 분류학에 대해 알려줘" — 일반 AI vs Cocrates의 응답 비교
     - Cocrates는 정답을 주지 않고, 결함이 있는 예시 + 미션으로 첫 턴을 주도
  2. **Education 단계 — 세 가지 깨달음 (2차원 매트릭스)**:
     - **첫 번째 깨달음**: 피라미드는 '수업 순서(진도)'가 아니라 **의존성 구조(Pull 전략)** — 상위 수준은 하위 수준을 전제한다
     - **두 번째 깨달음**: 1차원이 아니라 **2차원 매트릭스(Knowledge × Cognitive = 24개 셀)** — 사실적 지식과 개념적 지식이 같은 인지 과정에 있더라도 완전히 다른 접근이 필요하다
     - **세 번째 깨달음**: **Push(순차적) vs Pull(도전적)** 전략과 ZPD에 따른 혼합 — 학습자의 수준에 따라 밀고(Mission) 당기는(Push/Pull) 전략을 구사
  3. **Knowledge Capture 단계**: 배운 내용을 `kb/bloom-taxonomy.md`에 **핵심과 Gap 중심**으로 기록. Gap 기록(내가 틀렸던 가정)이 핵심 요약보다 강력함을 강조
  4. **Reflection 단계**: Apply→Analyze→Evaluate→Create 4개 미션으로 진짜 이해도 확인. **파이널 미션: Git 브랜치 전략 커리큘럼 설계(Create 수준)**
     - MISSION 1(Apply): DIP 도메인에 지식 차원 투영하기
     - MISSION 2(Analyze): 위계와 순서의 차이를 설계로 연결
     - MISSION 3(Evaluate): Push/Pull과 ZPD의 조건부 적용
     - MISSION 4(Create): Git 브랜치 전략 첫 3주 커리큘럼 설계
     - 최종 산출물: ✅ **확실한 영역** / ⚠️ **삐걱거리는 영역(내가 틀렸던 가정)**
  5. **전체 사이클 회고**: Education → Knowledge Capture → Reflection → KB 갱신. **Reflection은 면접관 모드로 변신** — 친절한 코치에서 엄격한 면접관으로
- **핵심 메시지**: Learning 파이프라인의 3단계를 실제 대화 기반 스토리텔링으로 경험함으로써, Cocrates의 학습 방식이 단순 정보 전달과 어떻게 다른지 체감한다. **2차원 매트릭스(Knowledge × Cognitive)** 라는 깨달음을 통해 블룸 분류학을 완전히 다르게 이해하게 된다.
- **포함할 내용**:
  - [ ] 일반 AI 응답 vs Cocrates 응답 비교 (같은 질문, 완전히 다른 답변)
  - [ ] 실제 대화 스크립트 예시 (Education 턴제 미션의 구체적 흐름)
  - [ ] 2차원 매트릭스 다이어그램 (Knowledge Dim × Cognitive Process = 24개 셀)
  - [ ] Push/Pull 전략 비교 설명
  - [ ] KB 저장 결과 예시 (`kb/bloom-taxonomy.md` 포맷) — Gap 기록 강조
  - [ ] Reflection 평가 결과 예시 (✅ 확실한 영역 / ⚠️ 삐걱거리는 영역)
  - [ ] 파이널 미션: Git 브랜치 전략 커리큘럼 설계
  - [ ] 독자가 직접 Cocrates와 학습할 때의 팁
- **제외할 내용**:
  - 산출물 생성 실습 (Ep5에서 다룸)
  - 스킬 생성 실습 (Ep6에서 다룸)
- **CTA**: 학습 활동의 실제를 경험했다면, 이제 구조 기반 산출물 생성 활동을 실습해보자. Ep5에서는 ADR부터 검증까지 전 과정을 직접 따라간다.

---

### Ep5: 구조 기반 산출물 생성 활동 (`05-architecture-driven-generation-practice`)

- **상태**: ✅ 작성 완료 (`episodes/05-architecture-driven-generation-practice.md`)
- **메타**: 예상 분량 18~22분, 발행 순서 5
- **Hook**: *"cocrates/jsondb에 jsondb를 개발해줘."* — 평소 같았으면 AI는 바로 **"물론입니다!"** 라며 코드를 쏟아냈을 것이다. 하지만 Cocrates는 질문으로 답했다. *"이 jsondb의 용도는 무엇인가요?"* 그리고 그 질문이 세 개의 ADR과 하나의 Spec, 일곱 개의 Go 파일로 이어졌다. 그런데 **생성된 코드에도 예상치 못한 차이(Deviation)가 발견되었다.** 이 편은 생성 이후 검증까지의 전 과정을 함께 따라가는 실습이다.
- **스토리라인**:
  1. **실습 시나리오**: "jsondb를 개발해줘" — 요구사항 명확화부터 시작
  2. **ADR 단계 — AI 제안의 3가지 문제점 + Living Cycle**:
     - **AI 제안이 부적절한 3가지 이유**:
       - **① 일반적(Generic)**: 모든 경우에 통용되는 모범 사례만 제시 — ADR 1(저장 모델): AI의 Collection-Document 가정 → 사용자의 Path-Addressable 결정
       - **② 쉽고 간단(Simple)**: AI는 가장 구현하기 쉬운 방법을 우선 제시 — ADR 2(아키텍처): Library-Only → Client-Server로 전환 (사용자가 복수 프로세스 시나리오를 질문)
       - **③ 품질 미반영(Quality-blind)**: 성능·확장성 같은 품질 속성을 고려하지 않음 — ADR 3(동시성 모델): DB-level RWMutex → Per-File RWMutex Map (사용자가 성능 병목 예측)
     - **Living Cycle**: 검증 피드백이 구조 설계(ADR)로 회귀하는 순환. 검증 → Spec 업데이트 → 재생성 → 재검증 = Waterfall이 아니라 **살아있는 순환**
  3. **Spec 단계**: 승인된 모든 결정을 `spec/jsondb.md`에 자체 완결적으로 통합
  4. **Generation 단계**: Spec만 보고 Library 7개 Go 파일 + Server + CLI 생성, `go build`/`go vet`/테스트 모두 통과
  5. **Verification 단계**: 72개 검증 항목 — **71 pass / 1 fail / 0 not-verifiable**
     - Deviation 사례: CLI Schema URL 이중 슬래시 문제
     - **Undocumented ASR 6건 발견** (URL encoding, SetPath 검증 누락 등)
  6. **회고**: '쉬운 방법의 함정', 반복/증분 개발의 필요성, Undocumented ASR은 실패가 아니라 **다음 Spec을 더 완벽하게 만들 기회**
- **핵심 메시지**: Cocrates와의 구조 기반 생성은 단순히 코드를 얻는 과정이 아니다. AI가 제시하는 '쉬운 방법'을 무비판적으로 따르는 대신, 사용자가 직접 질문하고 비교하고 결정하는 **적극적인 설계 과정**이다. AI 제안의 3가지 문제(Generic/Simple/Quality-blind)를 인식하면, AI의 첫 제안을 더 신중하게 검토할 수 있다. 생성 후 검증(Verification)은 Spec을 더 완벽하게 다듬는 **Living Cycle**의 시작점이다.
- **포함할 내용**:
  - [ ] 세 가지 ADR 각각의 'Cocrates의 첫 제안' vs '사용자의 검토 후 결정' 비교 (표)
  - [ ] 가장 극적인 반전 — Library-Only → Client-Server 전환 과정
  - [ ] Per-File Mutex 설계 — 성능 고려의 구체적 예시
  - [ ] AI 제안의 3가지 문제(Generic/Simple/Quality-blind) 설명
  - [ ] Living Cycle 다이어그램 — 검증 → ADR/Spac 업데이트 → 재생성 → 재검증
  - [ ] Spec → Generation → Verification의 실제 산출물 예시
  - [ ] Verification 결과: 72개 검증 항목, 71 pass / 1 fail
  - [ ] Undocumented ASR 6건 표
- **제외할 내용**:
  - 학습 활동 실습 (Ep4에서 이미 다룸)
  - 스킬 생성 실습 (Ep6에서 다룸)
- **CTA**: 구조 기반 산출물 생성의 전체 사이클을 경험했다. 이제 직접 스킬을 만들어보는 마지막 실습으로 넘어간다. 그 전에 스스로에게 물어보자: *"AI가 제시한 첫 번째 해결책을, 나는 정말 충분히 검토했는가?"*

---

### Ep6: 구조 기반 워크플로우 생성 (`06-workflow-creation-practice`)

- **상태**: ✅ 작성 완료 (`episodes/06-workflow-creation-practice.md`)
- **메타**: 예상 분량 16~20분, 발행 순서 6
- **Hook**: "설명문/보고서 등을 생성하는 스킬을 생성해줘." — 언뜻 보면 문서 하나를 만들어 달라는 요청 같지만, Cocrates는 다르게 읽는다. **"물고기를 잡아달라" vs "물고기 잡는 법을 배우라"** — 사용자가 요청한 것은 **산출물이 아니라 '그 산출물을 만드는 절차'** 다. 이 요청은 `generating-skill-creation`을 발동시킨다. 이 편은 그 전 과정을 따라가는 실습이다.
- **스토리라인**:
  1. **실습 시나리오**: "설명문/보고서 스킬을 만들어줘" — generating-skill-creation 발동
  2. **1단계 — Kernel**: 핵심을 한 문장으로 정의 ("이 스킬은 일반적 상황에서 Markdown 설명문/보고서를 검토 가능한 단계를 거쳐 생성하도록 돕는다")
  3. **2단계 — Frame**: 문서 계층 구조(메타→개요→본문→결론→부록)와 P1~P5 5단계 설계
  4. **3단계 — Outline**: P1~P5 각 단계의 입력·출력·승인 조건 정의
     - **사용자의 결정적 기여**: "항상 서론→본론→결론 순서는 아니다" → AI의 고정관념을 깨고 **'6가지 맞춤형 문서 구성 방식'** 으로 발전
  5. **4단계 — Spec**: 5가지 핵심 원칙, 6가지 맞춤형 구성 방식, 7가지 금지사항, 5가지 완료 조건
  6. **5~6단계 — Skill 생성 및 Verification**: SKILL.md 생성 후 **7개 검증 항목 전원 PASS**
  7. **U→A→J→A 프레임워크**: 이해(Understand) → 분석(Analyze) → 판단(Judge) → 승인(Approve). Cocrates의 모든 검토 활동의 공통 리듬
  8. **Snowflake 공통 패턴**: Ep4(학습) / Ep5(산출물) / Ep6(스킬 생성) — 모든 실습이 **작게 시작, 단계별 확장, 검토와 승인, 파일 저장**의 Snowflake 패턴을 공유한다
- **핵심 메시지**: 스킬 생성은 단순히 `SKILL.md` 파일을 만드는 기술적 작업이 아니다. Cocrates와 사용자가 **대화를 통해 산출물의 구조를 발견하고, 절차를 설계하고, 원칙을 정립하는 공동 설계 과정**이다. "보고서 써줘" 대신 "보고서 쓰는 법"을 AI에게 가르치는 것 — 이것이 AI-native Engineer의 사고방식이다.
- **포함할 내용**:
  - [ ] "물고기 잡는 법" 비유 — 산출물 vs 절차의 차이
  - [ ] Kernel 단계: Cocrates의 세 가지 질문(무엇을, 누가, 어떤 형태로)과 Kernel 문장 도출
  - [ ] Frame: 문서 계층 구조와 생성 단계 설계
  - [ ] Outline: P1~P5 5단계 절차 표
  - [ ] 사용자의 결정적 기여: "항상 서론→본론→결론 순서는 아니다" → 6가지 맞춤형 문서 구성 방식
  - [ ] Before/After 비교: 수동적 수용자 → 능동적 설계자
  - [ ] U→A→J→A 프레임워크 설명
  - [ ] Snowflake 공통 패턴 — Ep4/5/6의 공통 구조
  - [ ] 생성된 document-authoring 스킬의 핵심 설계 요소
  - [ ] 실제 코드 경로: `.opencode/skills/document-authoring/SKILL.md`
- **제외할 내용**:
  - 학습 실습 (Ep4에서 다룸)
  - 산출물 생성 실습 (Ep5에서 다룸)
- **CTA**: 이제 `document-authoring` 스킬이 추가되었다. 업무 보고서나 학습 정리 문서를 요청해보라. Cocrates가 개요부터 섹션 구조, 섹션별 작성까지 단계별로 검토와 승인을 요청하며 함께 만들어갈 것이다. 세 가지 실습을 모두 마쳤다면, 이제 Cocrates Harness의 구조와 원리를 깊이 이해해볼 차례다.

---

## 3. 구조 — 원리를 이해한다

### Ep7: Cocrates Harness 구조 (`07-cocrates-harness-structure`)

- **상태**: ✅ 작성 완료 (`episodes/07-cocrates-harness-structure.md`)
- **메타**: 예상 분량 12~15분, 발행 순서 7
- **Hook**: 하나의 프롬프트로 문서도 쓰고, 코드도 만들고, 학습도 시키고... 가능할까? Cocrates는 불가능하다는 결론에서 출발한다. 하나의 프롬프트로는 모든 산출물을 제대로 다룰 수 없다. 이것이 **One-Size-Fits-All 함정**이다. 주방에 칼이 하나뿐이라면? 토마토도 저 칼로, 빵도 저 칼로, 생선도 저 칼로 썰어야 한다. 결과는? 모든 요리가 엉망이 된다. 그래서 Cocrates는 **Agent + Skills** 구조로 설계되었다.
- **스토리라인**:
  1. **왜 Agent + Skills 구조인가**
     - **One-Size-Fits-All 함정**: 하나의 프롬프트로 모든 산출물 유형(코드/문서/발표자료/블로그/분석보고서/학습노트)을 처리하려면 지나치게 일반적이거나 특정 유형에 치우침
     - **스킬로 분리하면**: 각 산출물 유형에 최적화된 워크플로우 제공 + 개별 추가/수정/개선 가능
     - **지속적 개선**: 새로운 산출물 유형이 필요하면 스킬만 추가하면 됨
  2. **Cocrates = Cocrates Agent + Skills**
     - **Cocrates Agent** ([`cocrates.md`](pathname:///cocrates.md)): 공통된 원칙과 제어를 담당
       - 6개 섹션: Persona → Principle → Harness Architecture → Request Handling → Core Activities → Success Criteria
     - **Skills** (`.opencode/skills/*/SKILL.md`): 구체적 워크플로우를 담당
       - 각 스킬은 독립 파일로 관리되며 추가·수정·확장 가능
  3. **Cocrates Agent Prompt 6개 섹션 상세**
     1. **정체성 (Persona)**: "You are **Cocrates**: an agent that turns uncertainty into disciplined inquiry..."
     2. **핵심 원칙 (Principle)**: Harness Ignorance — 불확실성을 방치하지 않고 아키텍처를 통해 보고·검토·관리 가능하게
     3. **구조 (Harness Architecture)**: Cocrates Agent + Skills 분리 구조
     4. **요청 처리 (Request Handling)**: 의도 추론(1) → 스킬 로드(2) → 없으면 spec-driven-generation(3) → 멀티스텝 추적(4). 8개 의도 → 8개 스킬 매핑
     5. **핵심 활동 (Core Activities)**: Artifact Generation 파이프라인(ADR→Spec→생성→검증) + Learning 파이프라인(Education→Knowledge Capture→Reflection)
     6. **성공 기준 (Success Criteria)**: 사용자가 자신의 지식과 무지를 명확히 이해하고, 산출물의 구조와 내용을 스스로 설명 가능하며, 모든 단계에 주체적으로 참여
  4. **"사용 → 이해 → 진화" 사이클**: Ep3~6에서 경험 → Ep7~12에서 이해 → Ep13~14에서 진화
- **핵심 메시지**: Cocrates가 Agent + Skills 구조로 설계된 이유는 **산출물마다 다른 구조적 접근이 필요하고, 지속적인 개선이 가능해야 하기 때문**이다. Agent Prompt는 6개 섹션으로 구성되며, **Intent-To-Skill Routing**이 8개 의도를 적절한 스킬로 연결한다. Cocrates는 완성된 하네스가 아니라 사용자가 직접 스킬을 추가/수정/진화시킬 수 있는 **살아있는 시스템**이다.
- **포함할 내용**:
  - [ ] One-Size-Fits-All 함정 — 주방(칼) 비유
  - [ ] 다양한 산출물 유형과 각각 필요한 구조적 접근의 차이 (비교)
  - [ ] Cocrates Agent + Skills 구조 다이어그램
  - [ ] Intent-To-Skill Routing 표 (의도별 스킬 매핑)
  - [ ] 두 Core Activity 파이프라인
  - [ ] Cocrates Agent Prompt 구조 요약 (6개 섹션)
- **제외할 내용**:
  - 개별 스킬의 상세 워크플로우 (Ep8~12에서 다룸)
  - 설치 방법 (Ep3에서 다룸)
  - 실습 (Ep4~6에서 다룸)
- **CTA**: 이 구조의 첫 번째 활동 영역 — 소크라테스식 학습 활동의 개념을 다음 편에서 살펴본다.

---

### Ep8: 소크라테스식 학습 활동 - 원리 (`08-socratic-learning-activity`)

- **상태**: ✅ 작성 완료 (`episodes/08-socratic-learning-activity.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 8
- **Hook**: "알려줘"라고 말하는 순간, 대부분의 AI는 정답을 쏟아낸다. 하지만 Cocrates는 질문으로 답한다. 왜? **"알려줘"는 가장 위험한 세 글자**이기 때문이다. 정답을 받는 순간, 당신은 생각을 멈춘다. 모른다는 사실조차 인식하지 못한 채 넘어간다. 이것이 무지의 눈덩이가 시작되는 지점이다.
- **스토리라인**:
  1. **왜 소크라테스식 학습인가**: 일방적 정보 전달의 문제점
     - 사용자가 모른다는 사실조차 모르게 됨 (무지의 무지)
     - 이해 없이 산출물을 수용하는 습관
     - 무지의 눈덩이 — 쌓일수록 감당할 수 없게 됨
  2. **능동적 학습의 3가지 교육 철학**:
     - **소크라테스 산파술 (Socratic Midwifery)**: 정답을 직접 주지 않고, 정교한 질문을 던져 사용자가 스스로 논리의 모순을 깨닫고 진리에 도달하게 한다
     - **블룸의 교육분류학 (Bloom's Taxonomy)**: 기억→이해→적용→분석→평가→창조의 6단계를 밟으며 질문의 깊이를 점진적으로 심화. 2차원 매트릭스(Knowledge × Cognitive)
     - **근접발달영역 (ZPD) 및 비계 설정 (Scaffolding)**: 사용자의 현재 지식 수준을 파악하여 너무 쉽지도, 너무 어렵지도 않은 딱 한 단계 위의 과제(Mission)를 제시. *"No pain, no gain"* — 적절한 인지적 부하 없이는 진정한 성장이 없다
  3. **Learning 파이프라인 개관**:
     ```
     Education → Knowledge Capture → Reflection
     ```
  4. **각 단계의 역할**:
     - **Education (소크라테스형 1:1 학습 코치)**: "알려줘", "설명해줘" 요청 → education 스킬 활성화. Concept Briefing → Thought Laboratory → Mission 3블록 구조
     - **Knowledge Capture (배운 것을 기록)**: "정리해줘", "저장해줘" → `kb/` 폴더에 recall용 핵심만 저장. 상세 설명이 아닌, **"내가 틀렸던 가정"**을 기록
     - **Reflection (진짜 이해 확인)**: "평가해줘", "시험해줘" → KB와 대화 기록을 근거로 소크라테스식 질문. ✅ 잘 알고 있는 것 / ⚠️ 제대로 알지 못했던 것
- **핵심 메시지**: 소크라테스식 학습은 사용자가 **스스로 깨닫고(Education), 기록하고(Knowledge Capture), 확인하는(Reflection)** 3단계 파이프라인으로 구성된다. 목표는 지식 전달이 아니라 **무지의 제거**다. "알려줘"는 가장 위험한 세 글자 — Cocrates는 질문으로 답함으로써 사용자가 스스로 생각하게 만든다.
- **포함할 내용**:
  - [ ] Learning 파이프라인 다이어그램
  - [ ] 기존 AI 학습 방식 vs Cocrates 학습 방식 비교
  - [ ] 3가지 교육 철학 설명: 소크라테스 산파술, 블룸의 교육분류학(2차원 매트릭스), ZPD 및 비계 설정
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
     - 일방 전달 금지 원칙 (No Spoon-feeding)
     - 턴제 미션 기반 대화 구조: Concept Briefing → Thought Laboratory → Mission 3블록
     - Bloom's Taxonomy + 2-Axis Matrix로 난이도 조절
     - Push(순차적) / Pull(도전적) 전략
  2. **Knowledge Capture Skill** — 배운 것을 기록하다
     - recall용 핵심만 저장: 상세 설명 금지, 한 줄 정의 필수
     - 병합 전략: 동일 주제의 기존 KB 파일을 찾아 보완
     - **'내가 틀렸던 가정'과 'Gap'** 을 기록하는 구조 — 무지(ignorance)의 기록
  3. **Reflection Skill** — 진짜 아는 것과 그냥 아는 척하는 것
     - Socratic evaluation: 단순 암기 확인이 아닌 적용·반례·경계 질문
     - KB를 평가 기준으로 사용
     - Gap 발견 시 Education 모드로 전환 (순환 구조)
     - 최종 산출물: ✅ **확실한 영역** / ⚠️ **삐걱거리는 영역**
  4. **세 스킬의 연결과 순환**:
     ```
     Education → Knowledge Capture → Reflection → (필요시) Education
     ```
     각 스킬은 독립적이면서도 하나의 완전한 학습 사이클을 형성한다. Education이 탐구를, Knowledge Capture가 기록을, Reflection이 검증을 담당한다.
- **핵심 메시지**: 세 스킬이 하나의 사이클로 연결되어 사용자의 학습을 능동적으로 이끈다. Education(탐구) → Knowledge Capture(기록) → Reflection(검증) → (필요시) Education(추가 탐구) 순환.
- **포함할 내용**:
  - [ ] Education 스킬의 3블록 형식 예시 (Concept Briefing / Thought Laboratory / Mission)
  - [ ] Knowledge Capture의 KB 저장 예시 — Gap 기록 강조
  - [ ] Reflection 평가 산출물 예시 (✅ / ⚠️)
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
  1. **왜 구조(architecture)가 먼저인가**
     - 구조 없이 생성된 산출물은 일관성과 논리가 부족함
     - 사용자가 산출물을 이해하고 설명할 수 없게 됨
     - 검토가 불가능한 블랙박스가 탄생
     - '그냥 작성해줘'는 '아무렇게나 작성해줘'와 같음
  2. **'구조'를 구체화하라 — ASR과 ADR**
     - "구조를 먼저 설계하라"는 말은 너무 추상적이다. 핵심은 **ASR(Architecturally Significant Requirement)** 을 식별하는 것
     - ASR이란 최종 산출물의 구조·구성·품질에 실질적 영향을 미치는 요구사항
     - ASR이 명시되지 않으면 생성 단계에서 **침묵의 기본값(Silent Default)** 으로 채워지고, 그 결과는 사용자가 통제할 수 없음
     - 전원주택 비유: 이삿짐센터에 "잘 넣어줘" → 중요한 유리그릇이 박스 아래에 깔림. ASR은 '무엇이 유리그릇인지' 식별하는 작업
  3. **Concern → ADR → Spec → 생성 (전원주택 비유)**
     - 전원주택을 짓는다고 생각해 보자. "1층 vs 2층?", "지붕 vs 옥상 vs 옥탑방?" — 각각이 하나의 **Concern(관심사/ASR)**
     - 사용자가 명확한 결정을 내렸으면 → Spec에 기록
     - 결정을 내리지 못했으면 → ADR이 대안을 분석하고 권장 → 사용자 결정 → Spec에 통합
     - Spec에 따라 건축(생성)
  4. **Artifact Generation 파이프라인 개관**:
     ```
     ASR 식별 → ADR (대안 분석 & 결정) → Spec (결정 통합) 
         → Spec-driven Generation → Spec-driven Verification
     ```
    5. **AI 팀원이 작업한다. AI-native Engineer(팀장)로서 가져갈 각 단계별 역량**:
      - **ASR 식별**: AI가 이것저것 ASR이라고 던져 주는데, **중요한 것과 중요하지 않은 것을 구별하는 분별력**. 침묵의 기본값을 깨는 첫걸음이지만, 더 나아가 AI가 모두 중요하다고 말하는 것 중에서 진짜 중요한 것만 골라내는 안목이다.
      - **ADR**: AI가 분석·제안한 결정이 **정말 최적인지 분석/판단/결정하는 판단력**. AI는 대안과 장단점을 보여주지만, 최종 결정은 팀장인 당신의 몫이다. "이 분석이 충분한가? 다른 대안을 더 고려해야 하지 않나?"를 질문하는 능력.
      - **Spec**: AI가 통합·작성한 Spec에서 **더 빠진 것은 없는지, 추가로 검토할 것은 없는지, 이 Spec으로 최종 산출물이 어떻게 나올지 예상하는 통찰력**. Spec을 읽고 "이대로 생성하면 어떤 부분이 문제가 될까?"를 미리 보는 능력. 빠진 내용이 발견되면 ADR로 돌아가 검토하고 Spec에 추가하는 **순환 감각**도 포함된다.
      - **Generation / Verification**: 결과물을 맹신하지 않고 **검증하는 태도**. Spec의 항목 하나하나를 대조하며, 결과물에 Spec에서 논의되지 않은 설계 결정(Undocumented ASR)이 포함되지는 않았는지 확인한다. "되는 것 같다"가 아니라 "Spec의 3.2항을 만족하는가?"가 기준이다.
- **핵심 메시지**: 구조 기반 생성의 핵심은 **'구조가 모호하다는 사실을 인정하는 것'** 에서 출발한다. 중요한 관심사(ASR)를 식별하고, 선택지를 분석·결정하고(ADR), 결정을 자체 완결적 명세로 통합한 뒤(Spec), 그 명세를 기준으로 생성하고 검증하는 파이프라인이 구조 기반 생성을 완성한다. **침묵의 기본값(Silent Default)** 을 방지하는 것이 구조 기반 접근의 핵심이다.
- **포함할 내용**:
  - [ ] ASR 개념 설명과 산출물 유형별 예시
  - [ ] 침묵의 기본값(Silent Default) 개념
  - [ ] 전원주택 비유 — Concern/ADR/Spec/생성의 관계
  - [ ] Artifact Generation 파이프라인 다이어그램
  - [ ] 구조 기반 접근 vs 무구조 접근 비교
  - [ ] 각 단계에서 독자가 가져갈 역량 포인트
- **제외할 내용**:
  - 각 스킬의 상세 워크플로우 (Ep11에서 다룸)
  - artifact-specific 스킬과의 관계 (Ep11에서 다룸)
  - 스킬이 없을 때의 처리 (Ep12에서 다룸)
  - 실습 시나리오 (Ep5에서 다룸)
- **CTA**: 다음 편에서 Artifact Generation 파이프라인의 각 스킬 핵심 내용을 자세히 살펴본다. 이번 편에서 구조 기반 생성의 '왜'를 이해했다면, 다음 편에서는 '어떻게'를 알게 될 것이다.

---

### Ep11: 구조 기반 산출물 생성 활동 - 스킬 (`11-architecture-driven-generation-skills`)

- **상태**: ✅ 작성 완료 (`episodes/11-architecture-driven-generation-skills.md`)
- **메타**: 예상 분량 14~18분, 발행 순서 11
- **Hook**: Ep10에서 구조 기반 생성의 '왜'를 이해했다. 이제 4개의 스킬이 각각 무엇을 담당하고, 어떻게 연결되어 하나의 파이프라인을 만드는지 그 핵심 내부를 들여다본다. 각 스킬은 고유한 원칙과 금지 사항을 가지며, 그 경계가 파이프라인의 품질을 결정한다.
- **스토리라인**:
  1. **adr-writing (Any Decision Record)** — 하나의 Concern = 하나의 ADR. 최소 2~3개 실질적 대안. 간결한 bullet 스타일. `proposed` → 사용자 승인 → `approved`. ADR은 **의사결정 감사 추적(audit trail)**
  2. **spec-writing** — **자체 완결성(Self-Containment)**: Spec 하나만 읽으면 '무엇을 만들지' 완전히 이해 가능. ADR 링크 금지. 검증 가능한 문장. 살아있는 문서(living document) — 승인 상태 추적 금지
  3. **spec-driven-generation** — **Spec이 유일한 근거**: 원래 프롬프트가 아닌 Spec이 생성의 기준. **ASR Readiness Gate**: 8가지 보편 범주(목적과 독자, 산출물 형태, 범위 경계, 품질 기준, 제약, 구조와 구성, 통합과 의존성, 프로세스와 단계)로 ASR 식별. artifact-specific 스킬이 있으면 그 스킬의 단계별 게이트를 따름
  4. **spec-driven-verification** — **검증의 이중 목적**: (1) 결과물이 Spec과 일치하는지 항목별 확인, (2) **Undocumented ASR** 발견. Deviation(Spec 위반) vs Undocumented ASR(Spec에 없던 결정) 구분. 사용자 검토 게이트 — 무조건 수정하지 않음
  5. **파이프라인의 유연성과 artifact-specific 스킬**:
     - **document-authoring**: 문서 구조를 개요 → 섹션 설계 → 섹션 생성 → 검토로 점진 구체화
     - **presentation-authoring**: 발표자료를 문서 → 섹션 → 페이지 계층으로 점진 구체화. 각 페이지는 거버닝 메시지(두괄식) + 서포트 구조
     - **blog-series-authoring**: 블로그 시리즈를 개요 → 아웃라인 → 에피소드 설계 → 에피소드 본문 순으로 점진 구체화
     - **공통 원칙**: Snowflake 점진 구체화, 단계별 사용자 승인 게이트, Architecture-driven 접근
- **핵심 메시지**: ADR(선택지 분석과 결정) → Spec(자체 완결적 명세 통합) → Generation(ASR Readiness Gate 통과 후 Spec 기반 생성) → Verification(항목별 검증 + Undocumented ASR 발견)의 4단계가 구조 기반 생성을 완성한다. 각 스킬은 고유한 원칙과 금지 사항으로 파이프라인의 품질을 보장한다.
- **포함할 내용**:
  - [ ] 각 스킬의 핵심 원칙 한 줄 요약 (ADR = 선택과 근거 / Spec = 자체 완결성 / Generation = Readiness Gate / Verification = Undocumented ASR 발견)
  - [ ] ADR vs Spec 비교 표
  - [ ] ASR 8가지 보편 범주 표
  - [ ] Undocumented ASR 개념 설명과 예시
  - [ ] artifact-specific 스킬 3종의 특징과 언제 사용하는지
  - [ ] 실제 코드 경로: `.opencode/skills/adr-writing/SKILL.md` 등
- **제외할 내용**:
  - generating-skill-creation (Ep12에서 다룸)
  - 실습 시나리오 (Ep5에서 다룸)
- **CTA**: 이 4개의 스킬이 구조적 생성 파이프라인의 핵심이다. 이는 일반적인 산출물에 대한 가장 일반적인 구조적 접근이다. 그런데, 산출물 유형에 따라 구조적 접근을 스킬로 명세하면 더 효과적이다. 다음 편에서는 **generating-skill-creation** 스킬의 핵심 요소를 알아본다.

---

### Ep12: 구조 기반 워크플로우 생성 스킬 (`12-workflow-creation-skill`)

- **상태**: ✅ 작성 완료 (`episodes/12-workflow-creation-skill.md`)
- **메타**: 예상 분량 13~16분, 발행 순서 12
- **Hook**: **generating-skill-creation**는 스킬을 생성하는 **메타-스킬**이다. 산출물을 생성하는 **절차 그 자체를 설계**한다. 스킬 = AI의 SOP(Standard Operating Procedure). 이것은 Cocrates Harness에서 가장 추상적이면서도 강력한 활동이다.
- **스토리라인**:
  1. **산출물을 분해하는 5가지 구성 요소(Artifact Components)**:
     - 과업과 제약(Assignment & Constraints), 맥락과 규칙(Context & Rules), 개체(Entities), 공간과 배치(Space & Placement), 구조와 흐름(Structure & Flow)
  2. **Snowflake 5단계 — 각 단계의 목적과 승인 게이트**:
     - Define(기준선 확정) → Plan(골격과 방향 승인) → Architecture Design(구조적 일관성 승인) → Detail Design(생성 전 최종 게이트) → Generation(SKILL.md 생성)
     - **절대 규칙**: detail design이 완전히 확정되기 전까지 generation 단계로 넘어가지 않음
  3. **단계별 구성 요소 구체화(Per-Stage Refinement)**: 각 구성 요소는 동일한 속도로 구체화되지 않음. 지연 구체화(Lazy Refinement) 전략
  4. **Meta Snowflake 7단계**: Kernel → Components → Frame → Outline → Spec → Skill → Verification
  5. **스킬의 재사용성과 진화**: 한 번 만든 스킬은 이후 재사용. 새 산출물 유형이 필요하면 스킬만 추가. artifact-specific skill(document/presentation/blog-series authoring)은 모두 이 방식으로 탄생
- **핵심 메시지**: generating-skill-creation은 **스킬을 만드는 스킬(Meta-Skill)** 이다. 최종 산출물을 5가지 구성 요소로 분해하고, Snowflake 5단계로 점진 구체화하며, Meta Snowflake 7단계로 저작하여 사용자 승인을 거쳐 SKILL.md를 생성한다.
- **포함할 내용**:
  - [ ] 5가지 산출물 구성 요소(Artifact Components) 표
  - [ ] Snowflake 5단계 — 각 단계의 목적과 승인 게이트 표
  - [ ] Per-Stage Refinement 개념 설명
  - [ ] Meta Snowflake 7단계 표
  - [ ] 실제 코드 경로: `.opencode/skills/generating-skill-creation/SKILL.md`
- **제외할 내용**:
  - 실습 (Ep6에서 다룸)
  - 개별 artifact-specific 스킬 상세 워크플로우 (Ep11에서 다룸)
- **CTA**: generating-skill-creation까지 이해했다면, Cocrates Harness의 모든 구조적 활동을 살펴본 셈이다. 이제 마지막으로, Cocrates를 지속적으로 발전시키는 방법과 사용자로서의 다짐을 알아보겠다.

---

## 4. 결론 — 지속적 개선과 선언

### Ep13: 올바른 Cocrates Harness 활용 (`13-evolution-of-cocrates`)

- **상태**: ✅ 작성 완료 (`episodes/13-evolution-of-cocrates.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 13
- **Hook**: Cocrates를 설치하고, 실습까지 마쳤다. 구조의 원리까지 이해했다. 이제 끝일까? 아니다. 진짜 사용은 지금부터 시작이다. Cocrates는 완성된 제품이 아니라 **함께 진화하는 도구**다.
- **스토리라인**:
  1. **사용자가 진화해야 한다** — 소크라테스식 질문 내재화, 검토(U→A→E→A)와 승인의 습관화
  2. **Cocrates Harness도 진화해야 한다** — 스킬 추가·수정·확장, 에이전트 프롬프트 커스터마이징
  3. **상호 진화 — 피드백 루프**: 사용자가 성장할수록 Cocrates도 더 효과적으로 동작. 사용 경험 → Agent/Skill 개선 → 더 나은 경험 → 사용자 성장
  4. **오픈소스 정신**: github.com/cocrates/cocrates.ai, 모든 코드와 프롬프트 공개. opencode plugin이 아니더라도 Claude Code, Cursor, Codex 등에 직접 구현 가능. 또는 나만의 하네스를 구축하는 것도 매우 바람직
- **핵심 메시지**: Cocrates는 완성된 제품이 아니라 **함께 진화하는 도구**다. 사용자도 진화하고(소크라테스식 질문 내재화, 검토 습관화), Cocrates도 진화한다(스킬 추가/수정, 프롬프트 개선). **상호 진화**가 핵심이다.
- **포함할 내용**:
  - [ ] 스킬 추가/수정 예시
  - [ ] Agent 프롬프트 개선 예시
  - [ ] 피드백 루프 다이어그램 (사용자 → Cocrates 상호 진화)
- **제외할 내용**:
  - 설치 방법 (Ep3에서 다룸)
  - 개별 실습 (Ep4~6에서 다룸)
- **CTA**: 마지막 편은 선언문이다. 당신이 AI를 올바르게 쓰는 사람임을 선언하라.

---

### Ep14: Cocrates 사용자 선언문 (`14-user-manifesto`)

- **상태**: ✅ 작성 완료 (`episodes/14-user-manifesto.md`)
- **메타**: 예상 분량 10~12분, 발행 순서 14
- **Hook**: 지금까지의 여정을 되돌아보며, 당신은 하나의 선언을 준비한다. Cocrates Harness는 당신의 **도구**다. 이 선언문은 당신의 **의지**다. 도구와 의지가 하나가 될 때, AI는 단순한 생성기가 아니라 **당신의 능력을 극대화하는 파트너**가 된다.
- **스토리라인**:
  1. **시리즈의 핵심 — 세 문장으로 압축**
     - **Cocrates의 핵심 원칙**: *"The unexamined code is not worth generating."*
     - **Socrates의 지혜**: *"I know I know nothing."*
     - **Harness Ignorance**: 무지를 두려워하지 말고, 오히려 그것을 활용하라
  2. **사용자 선언 7계명 — AI 주권(AI Sovereignty) 선언**:
     - *1. I do not blindly request.* — 설계 없는 요구를 거부한다
     - *2. I do not blindly trust.* — 검증 없는 신뢰를 거부한다
     - *3. I do not outsource my sovereignty.* — 맥락의 주권을 외주 주지 않는다
     - *4. I do not hide my ignorance.* — 무지를 아는 척 기만하지 않는다
     - *5. I relentlessly fight against cognitive laziness.* — 인지적 나태함에 맞서 끝까지 투쟁한다
     - *6. I always strive for the optimal.* — 언제나 최적의 구조를 갈망한다
     - *7. I never build a closed fortress.* — 나만의 닫힌 성벽을 쌓지 않는다
  3. **선언문을 마음속 깊이 새긴다**: 단순히 읽고 넘어가는 것이 아니다. 의심될 때, 이 선언문으로 돌아온다. "지금 나는 이 선언문에 따라 행동하고 있는가?"
  4. **Cocrates Harness + User Manifesto = 무한한 역량 확장**: 도구(Cocrates Harness)와 의지(User Manifesto)가 하나가 될 때, 당신은 AI-native Engineer로 거듭난다. 사용자는 주권자로서 설계를 주도하고, 최적을 추구하며, 닫힌 성벽 없이 집단지성과 함께 진화한다.
- **핵심 메시지**: Cocrates Harness의 궁극적인 목표는 사용자가 AI를 올바르게 활용하여 **자신의 능력을 극대화**하는 것이다. 사용자 선언 7계명은 이 여정의 나침반이다. 업데이트된 7계명은 더욱 선명해진 사용자의 자세를 반영한다: 설계와 검증을 통해 주권을 행사하고, 무지를 인정하며, 인지적 나태와 싸우고, 최적을 추구하며, 열린 생태계에서 함께 진화한다. **Cocrates Harness(도구) + User Manifesto(의지) = AI-native Engineer(무한한 역량 확장)**.
- **포함할 내용**:
  - [ ] 세 가지 핵심 문장 — Cocrates 원칙 / Socrates / Harness Ignorance 연결
  - [ ] 사용자 선언 7계명 (한글 + 영문)
  - [ ] 각 계명의 간단한 설명 — 설계, 검증, 주권, 무지 인정, 인지적 나태 투쟁, 최적 추구, 열린 공유
  - [ ] 마무리 메시지 — Cocrates Harness로 능력을 극대화하라
- **제외할 내용**:
  - 새로운 개념이나 기술 설명
  - 실습
  - Cocrates Agent의 약속 (선언문에 집중)
- **CTA**: 이제 당신 차례다. 이 선언문을 마음속 깊이 새기고, Cocrates Harness와 함께 당신의 능력을 극대화하라.

---

## 완료 기준

| 항목 | 기준 |
|------|------|
| **각 편 분량** | 읽기 8~15분 / 한국어 2,500~4,500자 |
| **검토 완료** | 사용자가 각 편을 읽고 구조·내용 승인 |
| **게시 준비** | 모든 편 검토 완료 후 최종 통합 리뷰 |
| **전체 상태** | 14개 에피소드 모두 작성 완료 ✅ |
