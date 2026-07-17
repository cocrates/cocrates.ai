# TODO: 초지능은 정말 위험한가? — 균형 잡힌 검토

> **Workspace:** `docs/ai-superintelligence-safety/`
> **Updated:** 2026-07-17

## Snapshot

| Done | In progress | Pending | Blocked | Skipped |
|------|-------------|---------|---------|---------|
| 0    | 1           | 5       | 0       | 0       |

**Current focus:** T-001
**Recommended next:** T-001 — 문서 정의 (overview.md) 확정

## Active

- [ ] **T-001** `in_progress` — 문서 정의 (overview.md) 검토 및 확정
  - Phase: ① define
  - Artifact: `docs/ai-superintelligence-safety/overview.md`
  - Depends: —
  - Notes: 사용자가 overview.md를 검토하고 승인하면 다음 단계로 진행

## Backlog

### Phase ② plan
- [ ] **T-002** `pending` — 문서 구조 설계 (outline.md)
  - Phase: ② plan
  - Artifact: `docs/ai-superintelligence-safety/outline.md`
  - Depends: T-001 ✓

### Phase ③ architecture design
- [ ] **T-003** `pending` — 섹션별 상세 설계 (sections.md)
  - Phase: ③ architecture design
  - Artifact: `docs/ai-superintelligence-safety/sections.md`
  - Depends: T-002 ✓

### Phase ④ detail design
- [ ] **T-004** `pending` — 각 섹션 내용 작성
  - Phase: ④ detail design
  - Artifact: `docs/ai-superintelligence-safety/sections/*.md`
  - Depends: T-003 ✓

### Phase ⑤ generation
- [ ] **T-005** `pending` — 최종 문서 통합
  - Phase: ⑤ generation
  - Artifact: `docs/ai-superintelligence-safety/ai-superintelligence-safety.md`
  - Depends: T-004 ✓

### Phase ⑥ validation/revision
- [ ] **T-006** `pending` — 검증 및 보완 (Q&A + validation)
  - Phase: ⑥ validation/revision
  - Artifact: `docs/ai-superintelligence-safety/q-and-a.md`, `docs/ai-superintelligence-safety/validation.md`
  - Depends: T-005 ✓

### Cross-cutting
- [ ] **T-007** `pending` — 리서치 (ai-2040.com 및 최신 자료 수집)
  - Phase: research
  - Artifact: `docs/ai-superintelligence-safety/docs/*.md`
  - Depends: —
  - Notes:任何 단계에서 진행 가능, T-001 승인 후 시작

## Completed

(아직 없음)

## Notes
- 사용자가 "결론은 독자가 판단해야 한다"고 명시 — 논쟁을 균형 있게 소개하는 데 집중
- 최신 연구/사례 포함 필수
- ai-2040.com 외에도 다양한 출처 리서치 필요
