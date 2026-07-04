# Slide Content Specification — Slide 4

## Slide 4: 🚦 spec-driven-generation Fallback — "spec이 없으면 spec을 만든다"

**Section:** 본론 1 — 전체 흐름
**Type:** 개념/개관
**Time:** ~1분 30초

### Visual Design
- **Background:** 동일 템플릿
- **Layout:** 상단 제목 + 중앙 이중 역할 + 하단 비유
  - 상단: 제목 + 부제
  - 중앙 50%: spec-driven-generation의 이중 역할 (좌측/우측 비교)
  - 하단 30%: Gate 통과 실패 시 플로우 + 비유

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | 🚦 spec-driven-generation — "spec이 없으면 spec을 만든다" | 상단 |
| 역할 1 | **Spec Readiness Gate** — Spec이 준비되었는지 확인 | 좌측 카드 |
| 역할 2 | **생성** — 준비된 Spec 기준으로 산출물 생성 | 우측 카드 |
| Gate 실패 플로우 | ASR 식별 → (대안 있으면) ADR → 사용자 결정 → Spec 작성 → 생성 | 하단, 작은 플로우 |
| 8가지 ASR 범주 | 목적/독자, 형태, 범위, 품질, 제약, 구조, 의존성, 프로세스 | 하단, 작은 태그 |
| 비유 | 건축주 "집 지어줘" → 설계사무소 "설계도부터 만들죠" | 하단 발표자 노트 |
| Key Point | 💡 **"생성하기 전에 멈출 줄 안다"** | 하단 강조 박스 |

### Visual Notes
- 이중 역할을 좌우로 나란히 배치, 시각적 대비
- Gate 실패 플로우는 Slide 9의 전체 파이프라인 프리뷰 역할

### Transitions
- **In:** 제목 → 역할 카드 → Gate 플로우 순서
- **Out:** Key Point 표시 후 유지
