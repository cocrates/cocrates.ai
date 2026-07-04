# Slide 04: ADR ① — 저장 모델 논쟁

## Key Message
AI가 "일반적"이라고 제안한 Collection-Document 방식이 사용자의 구체적인 요구(Path-Addressable)와 맞지 않았다. 일반적이다 ≠ 내 요구에 맞다.

## Content
- **Headlines / text:**
  - **ADR ①: 데이터를 어떻게 저장할까?**
  - **🤖 AI의 제안:** "MongoDB처럼 Collection-Document 구조가 **일반적**이고 쉽습니다"
    - 🏷️ 핵심 문제: **일반적이다 (Generic)**
  - **🧑 사용자의 브레이크:** "내가 지정한 경로가 그대로 파일명이 되는 Path-Addressable 방식이 필요해"
    - 예: `Set("episode/e1.json")` → 실제 `episode/e1.json` 파일
  - **✅ 결과:** 사용자의 요구대로 경로 매핑 구조로 결정
- **💡 교훈:** **일반적이다 ≠ 내 요구에 맞다.** AI가 가장 흔히 접한 패턴이 최적은 아니다.

## Visual Elements
- 가운데 저울/균형 이미지: 한쪽은 "일반적인 패턴", 다른 한쪽은 "내 요구사항"
- AI 말풍선(좌측) → 사용자 브레이크(우측) → 결정(하단) 순서 플로우
- "일반적이다" 키워드 강조 (굵은 글씨, 다른 색상)

## Layout Notes
- ADR 논쟁 형식: AI 제안 → 사용자 피드백 → 결과 → 교훈 순서
- "핵심 문제"는 강조 박스로 표시
- 하단 교훈은 인용문 스타일

## Reference
- Script: `scripts/script-04.md`
- Slide plan: `slides.md` → Slide 04
