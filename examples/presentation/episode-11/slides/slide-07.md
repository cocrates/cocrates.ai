# Slide Content Specification — Slide 7

## Slide 7: ⚙️ spec-driven-generation — "설계도대로 짓는다"

**Section:** 본론 2 — 4개 스킬 상세
**Type:** 개념/상세
**Time:** ~1분 30초

### Visual Design
- **Background:** 동일 템플릿
- **Layout:** 상단 제목 + 중앙 Gate 상세 플로우 + 하단 비유
  - 상단: 스킬명 + 핵심 원칙
  - 중앙 50%: ASR Readiness Gate 상세 (단계별 박스)
  - 하단 30%: 비유 (건축 현장)

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | ⚙️ spec-driven-generation — "설계도대로 짓는다" | 상단 |
| 핵심 원칙 | **Spec이 유일한 헌법이자 근거** | 제목 아래 |
| Gate Step 1 | Spec 없음 → 8가지 보편 ASR 범주 스캔 | 중앙, 박스 1 |
| Gate Step 2 | Open ASR 중 대안 있음 → adr-writing 호출 | 중앙, 박스 2 |
| Gate Step 3 | 모든 결정 완료 → spec-writing 호출 → Spec 완성 | 중앙, 박스 3 |
| Gate Step 4 | Spec 완성 → **생성 시작** | 중앙, 박스 4 (강조) |
| Spec 우선 원칙 | 최초 프롬프트보다 Spec을 우선시 | 중앙, 작은 박스 |
| 비유 | "일단 벽돌부터" → "설계도부터 확인합시다" | 하단, 텍스트 |
| Key Point | 💡 **"생성하기 전에 멈출 줄 아는" 용기** | 하단 강조 박스 |

### Visual Notes
- Gate 단계를 가로 또는 세로 플로우로 표현
- Step 4 (생성 시작)에만 초록색 강조 — 나머지는 대기 상태임을 암시

### Transitions
- **In:** 제목 → Gate 단계 순차 등장 (1→2→3→4)
- **Out:** 비유 + Key Point
