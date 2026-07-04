# Script — Slide 10 (2분 30초)

## Slide: 💡 핵심 통찰 + 학습 확인

**Time:** ~2분 30초

"오늘 배운 내용을 **3가지 핵심 통찰**로 요약합니다.

첫째, **'스킬이 없으면 spec-driven-generation이 폴백된다'**
이것이 Cocrates의 기본 동작입니다.
스킬이 있으면 스킬을 쓰고, 없으면 spec-driven-generation이 ASR→ADR→Spec→생성의 흐름을 자동으로 시작합니다.

둘째, **'생성하기 전에 멈출 줄 안다'**
Spec Readiness Gate가 구조 없는 생성을 차단합니다.
AI가 마구잡이로 생성하는 것을 막고, 사용자가 의도한 방향으로 가도록 통제합니다.

셋째, **'각 스킬은 파일로만 소통한다'**
adr/ → spec/ → 생성물 → verification/
이 명확한 파일 기반 인터페이스가 각 스킬의 독립성과 모듈성을 보장합니다.

(잠시 멈춤)

자, 이제 스스로 질문에 답해보시겠어요?

질문 하나: 사용자가 '보고서 작성해줘'라고 요청하면, Cocrates는 어떤 순서로 동작할까요?
— 라우팅 → 스킬 검색 → 없음 → spec-driven-generation 폴백 → Gate → ASR 식별... 기억나시나요?

질문 둘: spec-driven-generation이 spec 없을 때 하는 3단계는?
— ASR 식별 → ADR(대안 있으면) → Spec 작성. 맞습니다.

질문 셋: adr-writing의 3가지 금지 규칙은?
— 가짜 대안 금지, 산문 스타일 금지, 무단 승인 금지.

질문 넷: Undocumented ASR 발견 시 3가지 처리 방법은?
— 무시(Ignore), 수용(Accept), ADR 검토(Escalate).

(잠시 멈춤)
모든 질문에 답하셨다면, 오늘 발표의 100%를 이해하신 겁니다."
