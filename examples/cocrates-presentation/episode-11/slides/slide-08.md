# Slide Content Specification — Slide 8

## Slide 8: 🔍 spec-driven-verification — "설계도와 대조하라"

**Section:** 본론 2 — 4개 스킬 상세
**Type:** 개념/상세 + 원칙
**Time:** ~1분 30초

### Visual Design
- **Background:** 동일 템플릿
- **Layout:** 상단 제목 + 좌측 Deviation + 우측 Undocumented ASR
  - 상단: 제목 + 핵심 원칙
  - 좌측 50%: Deviation 설명 (예시 + 원인)
  - 우측 50%: Undocumented ASR 설명 (3단계 처리)
  - 하단: 숨은 목적 + Key Point

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | 🔍 spec-driven-verification — "설계도와 대조하라" | 상단 |
| 핵심 원칙 | 이중 목적 검증 — **Deviation + Undocumented ASR** | 제목 아래 |
| Deviation (좌측) | Spec 명시와 다른 구현 | 좌측 상단 |
| Deviation 예시 | "PostgreSQL 쓰라니까 MongoDB" | 좌측 예시 박스 |
| Deviation 원인 | 설계 실수 / 구현 충돌 / AI 판단 | 좌측 하단 |
| Deviation 처리 | → **사용자 재검토** (어느 쪽이 잘못됐는지) | 좌측 강조 |
| Undoc ASR (우측) | Spec에 없었지만 AI가 구현한 내용 | 우측 상단 |
| Undoc ASR 예시 | AI가 임의로 하위 섹션 추가 / 새로운 요구사항 제시 | 우측 예시 박스 |
| 처리 1 | 🙅 **무시 (Ignore)** — 중요하지 않음 → Spec 미포함 | 우측, 카드 1 |
| 처리 2 | ✅ **수용 (Accept)** — 중요하고 타당 → Spec에 포함 | 우측, 카드 2 |
| 처리 3 | 🔍 **ADR 검토 (Escalate)** — 중요하지만 판단 어려움 → ADR로 | 우측, 카드 3 |
| 숨은 목적 | ASR 놓친 경험 → **사용자 역량 개선** | 하단 좌측 |
| Key Point | 💡 검증은 생성물 검사 + **사용자의 다음 설계를 개선하는 피드백 루프** | 하단 강조 박스 |

### Visual Notes
- 좌우 분할이 명확히 구분되도록 색상 차이 (Deviation: 빨간 계열, Undoc ASR: 파란 계열)
- 3단계 처리 카드에 각각 아이콘 + 간결한 설명

### Transitions
- **In:** 제목 → 좌측(Deviation) → 우측(Undoc ASR, 단계별 순차)
- **Out:** 숨은 목적 → Key Point
