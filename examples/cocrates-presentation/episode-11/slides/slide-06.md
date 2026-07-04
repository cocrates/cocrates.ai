# Slide Content Specification — Slide 6

## Slide 6: 📋 spec-writing — "무엇을 만들기로 했는가"

**Section:** 본론 2 — 4개 스킬 상세
**Type:** 개념/상세
**Time:** ~1분 30초

### Visual Design
- **Background:** 동일 템플릿
- **Layout:** 상단 제목 + 좌측 원칙 + 우측 Spec vs ADR 비교
  - 상단: 스킬명
  - 좌측 40%: 핵심 원칙 + 금지 규칙
  - 우측 60%: Spec vs ADR 비교 + 3가지 역할

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | 📋 spec-writing — "무엇을 만들기로 했는가" | 상단 |
| 핵심 원칙 | **자체 완결성 (Self-Containment)** — 이 문서 하나면 OK | 좌측 상단 |
| 금지 1 | ❌ **ADR 참조 금지** — "adr/01.md 참고" 금지 | 좌측 중단 |
| 금지 2 | ❌ **장황한 줄글 금지** — 1 불릿 = 1 검증 가능 요구사항 | 좌측 하단 |
| Spec vs ADR | ADR = 결정의 **이유** (왜) / Spec = 결정된 **사실** (무엇을) | 우측 상단 비교표 |
| 역할 1 | 📝 생성의 유일한 기준점 | 우측 하단, 3개 불릿 |
| 역할 2 | ✅ 검증의 유일한 기준점 | 우측 하단 |
| 역할 3 | 🤝 의사소통의 공통 참조점 | 우측 하단 |
| Key Point | 💡 Spec은 ADR(과거)과 Generation(미래) 사이의 **현재 설계도** | 하단 강조 박스 |

### Visual Notes
- ADR vs Spec 비교: 두 개의 카드를 나란히 배치, 각각의 성격을 아이콘으로 표현
- 자체 완결성 원칙에 가장 큰 시각적 비중

### Transitions
- **In:** 제목 → 좌측(원칙+금지) → 우측(비교+역할)
- **Out:** Key Point 강조
