# Slide 06: ADR ③ — 동시성 모델 논쟁

## Key Message
AI가 "코드 3줄"로 강조한 DB-level Lock이 성능(처리량, 응답 시간)이라는 핵심 품질을 전혀 고려하지 못했다. AI는 품질을 충분히 반영하지 못한다.

## Content
- **Headlines / text:**
  - **ADR ③: 동시 접근을 어떻게 제어할까?**
  - **🤖 AI의 제안:** "DB 전체 락(DB-level Lock)이면 코드 3줄이면 끝납니다"
    - 🏷️ 핵심 문제: **품질을 충분히 반영하지 못한다 (Quality-blind)**
  - **🧑 사용자의 브레이크:** "파일 하나 읽을 때 전체 DB를 멈추면 병목이잖아! 파일별 락(Per-File Mutex)으로 해줘."
  - **✅ 결과:** `sync.Map`을 활용한 고성능 동시성 제어
- **💡 교훈:** **AI는 품질을 충분히 반영하지 못한다.** 성능, 확장성, 유지보수성 같은 품질 속성은 AI가 스스로 고려하지 않는다. 사용자가 지적해야 반영된다.

## Visual Elements
- 자물쇠 이미지: 하나의 큰 자물쇠(DB-level Lock) → 여러 개의 작은 자물쇠(Per-File Mutex)
- 성능 그래프 또는 속도 비교 시각화 (Lock 대비 Mutex의 성능 우위)
- "코드 3줄"은 따옴표로 강조

## Layout Notes
- ADR ①, ②와 동일한 템플릿 구조 유지 (일관성)
- AI 제안 → 사용자 브레이크 → 결과 → 교훈 순서
- "Code 3줄" 부분은 시청자의 기억에 남도록 강조

## Reference
- Script: `scripts/script-06.md`
- Slide plan: `slides.md` → Slide 06
