# Script — Slide 9 (2분)

## Slide: 🔗 파이프라인 연결 + 요약 매트릭스

**Time:** ~2분

"지금까지 4개 스킬을 각각 보았습니다.
이제 이 스킬들이 어떻게 연결되어 하나의 파이프라인을 이루는지 한눈에 보겠습니다.

(다이어그램을 가리키며)
사용자의 요청이 들어오면 먼저 **Routing**을 통과합니다.
스킬이 있으면 해당 스킬로 직행하고, 없으면 **spec-driven-generation의 Gate**로 갑니다.

Gate에서 spec이 없다고 판단되면 **adr-writing**으로 갑니다.
adr-writing에서 사용자가 결정을 내리면, 그 결정은 **spec-writing**으로 전달됩니다.
spec-writing이 Spec을 완성하면, 비로소 **spec-driven-generation**이 생성을 시작합니다.
생성된 결과물은 **spec-driven-verification**이 검증합니다.
검증 결과는 다시 spec-writing으로 피드백되어 Spec을 개정합니다. — 이것이 순환 구조입니다.

이제 매트릭스를 보겠습니다. (매트릭스를 가리키며)

각 스킬의 정체성은 '하는 일'보다 '하지 않는 일'이 결정한다고 말씀드렸죠.

adr-writing: 대안 분석과 결정 기록이 역할이지만,
'가짜 대안 금지, 산문 스타일 금지, 무단 승인 금지' — 이 세 가지 금지가 진짜 정체성입니다.

spec-writing: 결정 통합 설계도.
'ADR 참조 금지, 장황한 줄글 금지' — 이 두 가지가 없으면 그냥 장황한 문서가 됩니다.

spec-driven-generation: Spec 기준 생성.
'Spec 없이 생성 금지' — Gate 통과가 필수. 이 하나의 금지가 모든 품질을 보장합니다.

spec-driven-verification: Deviation과 Undocumented ASR 발견.
'사용자 컨펌 없이 독단적 수정 금지' — 이 금지가 검증을 단순한 체크리스트가 아니라 학습 도구로 만듭니다.

그리고 협력 관계를 보면 흥미롭습니다.
모든 스킬은 **파일로만 소통**합니다.
adr-writing의 결과는 adr/ 디렉토리의 파일.
spec-writing의 결과는 spec/ 디렉토리의 파일.
생성물은 output/ 디렉토리.
검증 결과는 verification/ 디렉토리.

각 스킬은 각자의 역할에 집중하고, 파일이라는 명확한 인터페이스로만 소통합니다.
이것이 모듈화의 정수입니다.

자, 이제 오늘 배운 내용을 정리하겠습니다."
