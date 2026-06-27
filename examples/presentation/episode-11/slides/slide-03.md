# Slide Content Specification — Slide 3

## Slide 3: 🔄 Intent-to-Skill Routing — 요청이 스킬이 되는 순간

**Section:** 본론 1 — 전체 흐름
**Type:** 개념/비유
**Time:** ~1분 30초

### Visual Design
- **Background:** 동일 템플릿
- **Layout:** 좌측 다이어그램 + 우측 설명
  - 좌측 50%: 라우팅 플로우 다이어그램
  - 우측 50%: 비유 설명 (전문 병원)

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | 🔄 Intent-to-Skill Routing | 상단 좌측 |
| 다이어그램 | 사용자 요청 → 산출물 유형 식별 → 스킬 검색 (구체적→일반) → 있음:스킬 사용 / 없음:fallback | 좌측, 화살표 플로우 |
| 비유 | 전문 병원 vs 응급실 | 우측, 텍스트 블록 |
| Key Point | 💡 모든 요청이 spec-driven-generation을 거치는 것은 아니다 — **스킬이 없을 때만 fallback** | 하단 강조 박스 |

### Routing Diagram Detail
```
사용자 요청 ("보고서 작성해줘")
         ↓
    산출물 유형 식별 (final deliverable type)
         ↓
    생성 스킬 검색 (가장 구체적인 타입부터)
         ↓
    ┌─── 있음 ──→ 해당 스킬 사용 (직행)
    │
    └─── 없음 ──→ spec-driven-generation Fallback (← 강조)
```

### Visual Notes
- 화살표에 단계별 번호 (1-2-3-4)
- "없음" 분기점에 빨간색/주황색 강조

### Transitions
- **In:** 다이어그램 먼저 나타나고, 우측 비유가 이어서 등장
- **Out:** Key Point를 마지막에 표시하고 유지
