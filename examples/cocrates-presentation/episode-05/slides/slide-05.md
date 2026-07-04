# Slide 05: ADR ② — 아키텍처 논쟁

## Key Message
AI가 "쉽고 간단하다"고 제안한 Library-Only 방식이 미래의 확장성(다중 프로세스) 요구를 만족하지 못했다. 쉽고 간단하다 ≠ 올바른 설계다.

## Content
- **Headlines / text:**
  - **ADR ②: 시스템 구조는 어떻게 할까?**
  - **🤖 AI의 제안:** "단일 프로세스 Library-Only가 **쉽고 간단**합니다"
    - 🏷️ 핵심 문제: **쉽고 간단하다 (Simple)**
  - **🧑 사용자의 브레이크:** "여러 AI 에이전트가 동시에 접근할 거야. Client-Server 구조로 가야 해."
  - **✅ 결과:** REST API 서버(jsondbd) + CLI 도구 포함 아키텍처로 전면 수정
- **💡 교훈:** **쉽고 간단하다 ≠ 올바른 설계다.** AI가 만들기 쉬운 방식이 실제 환경의 요구를 만족하지 않을 수 있다.

## Visual Elements
- 단일 프로세스 아이콘(좌측 작게) → Client-Server 아이콘(우측 크게) 확장对比
- 연결선: Library(단순) → REST API + CLI(복잡) 확장되는 느낌
- "쉽고 간단하다" 키워드 강조

## Layout Notes
- ADR ①과 동일한 템플릿 구조 유지 (일관성)
- AI 제안 → 사용자 브레이크 → 결과 → 교훈 순서
- "코드 3줄" 인용은 따옴표로 강조

## Reference
- Script: `scripts/script-05.md`
- Slide plan: `slides.md` → Slide 05
