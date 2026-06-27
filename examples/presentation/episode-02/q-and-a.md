# Expected Questions & Answer Plans — Episode 02

## Q1: '"Code"라는 표현이 개발자만을 위한 것처럼 느껴지는데, 비개발자도 적용할 수 있나요?'
- **Context:** Slide 5("Code의 확장") 이후 또는 Slide 10(요약)에서 'Code'라는 용어에 대해 이질감을 느낄 때
- **Answer strategy:**
  - Cocrates는 'Code'를 "모든 AI 생성 최종 산출물"로 확장 정의했음을 재확인
  - 보고서, 슬라이드, 블로그, 학습 노트까지 포함된다는 점을 구체적 예시로 재언급
  - "용어는 'Code'지만, 적용 범위는 AI를 사용하는 모든 사람"이라는 점을 강조
- **Reference:** slides.md Slide 5, slides/slide-05.md, outline.md Slide 5

## Q2: '검토(U→A→E→A)가 너무 번거롭다. 간단한 작업에도 이 4단계를 모두 거쳐야 하나요?'
- **Context:** Slide 7(U→A→E→A) 설명 중 또는 Slide 9(Cocrates 리듬)에서 Cocrates의 깐깐함에 대한 반응
- **Answer strategy:**
  - 모든 작업에 동일한 깊이의 검토가 필요한 것은 아님
  - **중요도에 따라 검토 깊이를 조정**할 수 있음 — 예: "이해 + 승인"만으로 충분한 작업도 있음
  - 하지만 **처음에는 4단계를 의식적으로 연습**하는 것이 중요 (자동화된 습관이 될 때까지)
  - "뼈대를 검토하는 1시간이, 일주일짜리 버그 수정을 막아준다" — 시간 대비 효과 관점에서 접근
- **Reference:** slides.md Slide 7, Slide 9, slides/slide-07.md, slides/slide-09.md

## Q3: 'Cocrates Harness는 구체적으로 어떻게 설치하나요? Ep3를 기다려야 하나요?'
- **Context:** Slide 12(다음 편 연결)에서 Cocrates Harness 설치에 대한 기대감 형성 후
- **Answer strategy:**
  - Ep3에서 상세 설치 가이드를 다룰 예정임을 안내
  - "생각보다 정말 간단합니다" — 부담을 덜어주는 톤
  - 설치에 앞서 '원칙(무엇)'을 먼저 이해하는 것이 중요함을 강조 (지금 에피소드의 역할)
  - 궁금증이 크다면 docs/ai-native-engineer.md에서 개념적 설명을 미리 볼 수 있음을 안내
- **Reference:** slides.md Slide 12, overview.md

## Q4: 'Ep1에서 AI-assisted와 AI-native를 배웠는데, 이 에피소드가 Ep1과 어떻게 연결되나요?'
- **Context:** Slide 1(타이틀)의 Ep1 복습 또는 시리즈 전체 흐름에 대한 질문
- **Answer strategy:**
  - Ep1 = "무엇이 다른가" (AI-assisted vs AI-native의 차이)
  - Ep2 = "핵심 역량은 무엇인가" (검토 = U→A→E→A)
  - Ep3 = "어떻게 실천할 것인가" (Cocrates Harness 설치와 첫 대화)
  - 즉, **Ep1이 '진단'이라면 Ep2는 '원칙 선언'** — 개념 이해에서 실천으로 가는 과정
- **Reference:** slides.md Slide 1, overview.md, docs/episode-02.md

## Q5: '"무지의 제거"라는 개념이 추상적으로 느껴집니다. 구체적인 예시가 있나요?'
- **Context:** Slide 8(무지의 제거) 설명 중 또는 후
- **Answer strategy:**
  - 실제 예시: "AI가 추천한 아키텍처 패턴을 검토하다가, 내가 몰랐던 디자인 패턴을 발견하고 학습하게 된 사례"
  - 또는: "AI가 생성한 SQL 쿼리를 분석하다가, 인덱스의 중요성을 깨닫게 된 경우"
  - 핵심: **"아, 이런 게 있었구나"** 라는 순간이 바로 무지가 제거되는 순간
  - 질문을 던지세요: "왜 이렇게 만들었지?", "다른 방법은 없을까?", "내가 모르는 게 뭘까?"
- **Reference:** slides.md Slide 8, slides/slide-08.md, docs/episode-02.md

## Q6: '12분짜리 영상에 12개 슬라이드가 너무 많지 않나요?'
- **Context:** 전체 에피소드 분량에 대한 피드백
- **Answer strategy:**
  - 55초/슬라이드는 온라인 교육 콘텐츠의 일반적인 페이스
  - 각 슬라이드의 내용 밀도가 높지 않고, 하나의 메시지에 집중
  - Slide 11(학습 확인)은 시청자가 영상을 멈추고 생각하는 시간을 포함하여 오히려 '여유'를 주는 장치
  - "짧고 강력하게"가 아니라 "이해할 시간을 주면서" 접근
- **Reference:** overview.md, slides.md

## Q7: 'Cocrates가 검토를 강조하는데, AI가 검토를 대신해줄 수는 없나요?'
- **Context:** Cocrates의 철학 전반에 대한 질문
- **Answer strategy:**
  - **AI는 검토를 도와줄 수 있지만, '승인'은 인간만이 할 수 있음**
  - AI는 "이 부분이 일반적인 패턴과 다릅니다"라고 알려줄 수 있지만, "프로젝트의 목적에 부합하는가?"는 인간이 판단
  - AI-native Engineer = AI를 팀원으로 활용하되, **최종 책임과 승인은 인간**
  - 즉, AI는 Analyze와 Evaluate를 도와줄 수 있지만, Approve는 인간의 몫
- **Reference:** docs/ai-native-engineer.md, slides.md Slide 7, Slide 9
