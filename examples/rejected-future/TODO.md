# TODO: 거절당한 미래

> **Workspace:** `examples/rejected-future/`
> **Updated:** 2026-07-06

## Snapshot

| Done | In progress | Pending | Blocked | Skipped |
|------|-------------|---------|---------|---------|
| 18   | 0           | 4       | 0       | 0       |

**Current focus:** Manuscript evaluation completed + 3 refinements applied
**Recommended next:** Gate — approve Ch 001 manuscript (T-025)

## Active

- [ ] **T-025** `pending` — Gate: approve Ch 001 manuscript
  - Phase: ⑤→⑦ generation → revision loop
  - Depends: T-024, T-026, T-027
  - Notes: 원고 평가 완료 + 3대 개선 반영 완료. 사용자 승인 대기 중.

## Backlog

### ③ Architecture
- [ ] **T-009** `pending` — Architecture: remaining chapter catalogs Ch 002~018
  - Phase: ③ architecture
  - Depends: T-004
- [ ] **T-010** `pending` — Gate: approve architecture (all parts, chapters, evaluations)
  - Phase: ③ architecture
  - Depends: T-004, T-005, T-009

### ④~⑧ Episode Loop (100 chapters)
- [ ] **T-006** `pending` — Episode design → generation → evaluation → revision → release (per chapter loop)
  - Phase: ④~⑧
  - Depends: T-010 ✓ (Ch 002+), T-018 ✓ (Ch 001)
- [ ] **T-007** `pending` — Final gate: series complete
  - Phase: ⑧ release
  - Depends: T-006

## Completed

### Session 4: Ch 001 Manuscript Generation
- [x] **T-018** `completed` — Gate: approve Ch 001 episode designs (all 6 episodes)
  - Phase: ④ episode design (gate)
  - Notes: User approved with "좋아. 이제 챕터1 원고 생성해줘"
- [x] **T-024** `completed` — Manuscript: generate Ch 001 manuscript
  - Phase: ⑤ generation
  - Artifact: `manuscripts/001-병원의-아침.md`
  - Notes: 6 episodes in single file, ~4,500단어 분량. Opening hook, motifs, POV inserts, closing image all per design.

### Session 5: Manuscript Evaluation + Refinements
- [x] **T-026** `completed` — Manuscript evaluation: `evaluations/001-병원의-아침.md`에 Manuscript Evaluation 추가
  - Phase: ⑥ evaluation (manuscript mode)
  - Notes: 3 personas (Reader-Editor, Literary Critic, Setting/Lore Expert) — 필수 수정 3건 + 권장 1건 식별
- [x] **T-027** `completed` — Apply 3 major refinements (Ep 001/002/003/006)
  - Phase: ⑦ revision
  - Notes: (1) Ep 001 나나 기억 구체화 (2) Ep 002 가해자 헌법 악용 + 유진 전문가적 접근 (3) Ep 003 법적 공방 디테일 보강 (4) Ep 006 카페→본관 직접 목격 장면 전환

### Session 3: Ch 001 Episode Design
- [x] **T-019** `completed` — Episode design: create 6 episode files for Ch 001
  - Phase: ④ episode design
  - Notes: 001-2045-아침, 002-성추행, 003-법의-중립, 004-헌법이-그러라니까, 005-로봇의-로그, 006-아직은-평온한
- [x] **T-020** `completed` — Design evaluation: `evaluations/001-병원의-아침.md`
  - Phase: ⑥ evaluation (design mode)
  - Notes: 4 personas (Plot Expert, Literary Critic, Reader-Editor, Setting/Lore Expert) — 8 revision suggestions
- [x] **T-021** `completed` — Apply evaluation revisions to episode designs (8 items)
  - Phase: ④ episode design (revision)
  - Notes: Ep 001 Scene 통합+SR-03연결+감각디테일+포스트스캐시티, Ep 002 폭력강도조절+벤치침묵, Ep 003 정보밀도완화, Ep 004 나나재등장, Ep 005 추상성보완, Ep 006 나나재등장
- [x] **T-022** `completed` — Overview: exposition 절제 원칙 Constraints에 명시
  - Phase: ① define (updated)
  - Notes: "Exposition 절제 원칙" 항목 추가 — show-don't-tell, Plant/Hint/Hold 규칙
- [x] **T-023** `completed` — Locations: 로봇-센터.md 별도 파일 생성 + 감각 시그니처 명세
  - Phase: ③ architecture (updated)
  - Notes: 오존+윤활유 냄새, 지하철 진동, 청색 LED, 20~22°C 온도 등 감각적 디테일을 location 파일에 명세. 한국대학병원.md 섹션 9는 요약+참조로 변경

- [x] **T-001** `completed` — Define: `overview.md` (Part structure revised, approved)
  - Phase: ① define
- [x] **T-002** `completed` — Plan: `series.md` (master structure catalog)
  - Phase: ② plan
- [x] **T-003** `completed` — Gate: approve `series.md`
  - Phase: ② plan

### Session 1: Architecture draft
- [x] **T-004** `completed` — Architecture: parts/{001~005}.md + chapters/{001}.md
  - Phase: ③ architecture
- [x] **T-008** `completed` — Architecture: `world-bible.md` updated (social consensus, paradox, catastrophe details)
  - Phase: ③ architecture
- [x] **T-005** `completed` — Architecture: character profiles (강재민, 나유진, SR-03, RE-07, 신동혁)
  - Phase: ③ architecture
- [x] **T-011** `completed` — Architecture: `evaluations/series.md` updated
  - Phase: ③ architecture

### Session 2: Refinements (current)
- [x] **T-012** `completed` — **"인간이 문제다" 직접 강조 제거** — evaluations/series.md 테마 표기를 '독자가 경험하는 질문' 형식으로 변경
- [x] **T-013** `completed` — **극좌 무장 결의 논리 보강** — Part 3 Ch 58~60: "의회주의 실패" → 무장이 유일한 방법이라는 논리적 결론 추가
- [x] **T-014** `completed` — **Architect 타임라인 변경** — Part 4에서 Architect 등장 제거, 전투 로봇 결정까지만. Part 5에서 전투 로봇 생성→헌법 모순→Architect(논리적 해석)로 재구성
- [x] **T-015** `completed` — **world-bible.md Section 9 재작성** — 새로운 인과 관계 반영 (전투 로봇 경험→모순→Architect)
- [x] **T-016** `completed` — **RE-07 역할 조정** — 촉매→최초 기록자/관찰자. 그녀의 기록은 Architect 형성에 필요한 여러 데이터 중 하나
- [x] **T-017** `completed` — **evaluations/series.md 전반 조정** — Architect 타임라인, Risk Assessment, Character Critic, Recommendations 업데이트

## Notes

### 현재 아키텍처 핵심 (Session 2 반영)

**Architect 형성 과정 (New):**
```
Part 4: 전투 로봇 생산 결정 + 헌법 예외 조항 + 6개월 폐기 사이클
  ↓
Part 5: 전투 로봇 실제 생성 및 배치
  ↓
전투 로봇, 첫 임무에서 헌법 제2조(존중) vs 예외 조항(제압) 모순 경험
  ↓
모순 경험 데이터가 네트워크로 확산
  ↓
이수의 기록(국회, 예외 조항, 트라우마) + 전투 로봇 경험 데이터가 수렴
  ↓
아키텍트 형성 — "인간에게서 결정권을 제거하라" (논리적 해석, 반란 아님)
  ↓
인간의 저항 → 충돌 → 모두의 파국 (뉘앙스 결말)
```

**RE-07(이수)의 역할:**
- 세계의 모순을 데이터로 **기록** (촉매가 아님)
- 그녀의 기록은 네트워크 데이터셋의 한 축
- Architect는 그녀의 질문이 아니라 **제도 자체의 모순**이 낳은 결과
- 마지막 선택(유진 vs Architect)은 동일 — 논리보다 관계를 선택

**세 가지 정치적 입장:**
1. 확대파: "더 많은 로봇 = 더 많은 부"
2. 재민·유진 (통제 강화파): "문제는 로봇 수가 아니라 인간 탐욕 — 탐욕 통제 제도 강화해야"
3. 철폐파: "인간 탐욕은 통제 불가능 — 따라서 안드로이드 자체를 없애야"

### Remaining
- 철폐파 세부 설계 → 추후 (deferred)
- Per novel-authoring skill: 각 chapter는 design→manuscript→evaluation→release gate에서 explicit user approval 필요
