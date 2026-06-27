# Slide 08: 📝 Spec — 결정을 하나의 헌법으로

## Key Message
3개의 ADR 결정을 하나의 설계 명세서(Spec)로 통합한다. Spec은 Self-Containment를 가지며, 유일한 진실 공급원(Single Source of Truth)이 된다.

## Content
- **Headlines / text:**
  - **Spec의 탄생:** 모든 ADR 결정을 **하나의 설계 명세서**로 통합
  - **Self-Containment (자기 완결성):** `spec/jsondb.md` 하나로 전체 구조 파악
    - 📋 10개 REST 엔드포인트 기능
    - 🔗 API 시그니처와 CLI 명령어 스펙
    - 🧪 테스트 요구사항까지 모두 포함
  - **Spec의 역할 (Single Source of Truth):**
    - 초기의 모호한 요청("jsondb 만들어줘")은 잊힌다
    - 오직 **통합된 Spec**만이 유일한 진실 공급원
    - 🏛️ Spec = **설계의 헌법**
- **💡 포인트:** Spec 없이 생성은 있을 수 없다. Spec이 곧 **설계의 완성**이다.

## Visual Elements
- 3개의 ADR 문서가 하나의 Spec 문서로 합쳐지는 애니메이션
- Spec 문서 아이콘에서 뻗어나가는 REST, CLI, Test 아이콘
- "헌법" 이미지 또는 상징 (법전 스타일)

## Layout Notes
- 중앙에 Spec 문서 이미지, 주변에 포함 내용 배치
- ADR → Spec 병합 과정을 시각화
- Self-Containment 개념은 체크리스트 형식으로

## Reference
- Script: `scripts/script-08.md`
- Slide plan: `slides.md` → Slide 08
