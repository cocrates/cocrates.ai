# Slide 11: 🤫 Undocumented ASR — AI의 침묵의 기본값

## Key Message
Spec이 명시하지 않은 빈틈을 AI가 자기 마음대로 채운 설계 결정이 Undocumented ASR이다. 검증을 통해 이 빈틈을 발견하고 Spec에 추가해야 한다.

## Content
- **Headlines / text:**
  - **Undocumented ASR이란?** Spec이 명시하지 않은 빈틈을 AI가 자기 마음대로 채운 설계 결정
  - **검증 중 6건 발견:**
    - 🔗 URL 인코딩 누락 — 경로에 특수문자가 들어가면?
    - 🔄 데이터 수정 시 스키마 재검증 생략 — 무결성 검증 누락
    - ... (기타 4건)
  - **왜 발생하는가?** Spec이 모든 것을 다 명시할 수는 없다. AI는 빈틈을 자신의 '기본값'으로 채운다.
- **💡 교훈:** Spec이 완벽하더라도, 명시되지 않은 공백은 AI가 임의로 채운다. 검증을 통해 이 빈틈을 발견하고 Spec에 추가해야 한다.

## Visual Elements
- 중앙에 Spec 문서 — 군데군데 빈 구멍(구멍 뚫린 종이 이미지)
- 빈 구멍에서 AI가 무언가를 채워넣는 손 이미지
- 발견된 6건은 리스트로 표시
- "침묵의 기본값" 컨셉: AI가 조용히 결정을 내리는 이미지

## Layout Notes
- Spec의 빈틈을 시각적으로 표현
- 6건의 ASR은 간결하게 리스트 (너무 많은 텍스트 지양)
- 교훈 메시지는 하단에 강조

## Reference
- Script: `scripts/script-11.md`
- Slide plan: `slides.md` → Slide 11
