# TODO: 마교 교주가 정파 맹주로 사기 친다

> **Project root:** `/home/drajin/work/cocrates.ai/work/마교-교주-정파-맹주-사기`
> **Updated:** 2026-08-01

## Snapshot

| Done | In progress | Pending | Blocked | Skipped |
|------|-------------|---------|---------|---------|
| 84   | 1           | 0       | 0       | 0       |

**Current focus:** T-085 — sequential continuity republication audit/revision, Chapter 151 next.
**Recommended next:** Audit Chapter 151 against the released Chapter 150 audit ledger, then continue sequentially.

## Active

- [x] **T-001** `done` — ① define: `overview.md` 작성 (전체 규모/모드/릴리스 단위/검증 기준)
  - Phase: ① define | Artifact: `overview.md`
  - Notes: 릴리스 단위 = 에피소드(화). 첨부의 'Chapter' = skill 'Part', 첨부의 '화' = skill 'chapter'로 매핑. 결말 = Option A. 완료 2026-08-01.
- [x] **T-002** `done` — ① gate: `overview.md` Architect 승인
  - Notes: Architect-approved gate 2026-08-01. 요구사항 명확(Publisher brief lock은 원안 제공으로 충족).
- [x] **T-003** `done` — ② plan: `series.md` 작성 (Part Catalog 17개 — 첨부 Chapter 1~17)
  - Notes: 완료 2026-08-01. 17파트/200화.
- [x] **T-004** `done` — ② gate: `series.md` Architect 승인
  - Notes: Architect-approved gate 2026-08-01. Part Catalog 정합 확인.
- [x] **T-005** `done` — ③ architecture: `characters.md` + `characters/*` (Part 1 출연진 프로필)
  - Notes: 완료 2026-08-01. 남궁태의/남궁휘/남궁혁/남궁진/설향/노의원 6개 프로필.
- [x] **T-006** `done` — ③ architecture: `locations.md` + `locations/*` (Part 1 장소 프로필)
  - Notes: 완료 2026-08-01. 남궁세가(취운각 병상·낙양 본가).
- [x] **T-007** `done` — ③ architecture: `world-bible.md` + `world/*` (무공 체계/천마신교/독)
  - Notes: 완료 2026-08-01. 회귀 원리·골혈재조공·혼백산·천마신교.
- [x] **T-008** `done` — ③ architecture: `stagings.md` + `stagings/*` (병상/야습 스테이징)
  - Notes: 완료 2026-08-01. p1-taeui-sickroom, p1-taeui-night-assault.
- [x] **T-009** `done` — ③ architecture: `parts/001~017` 챕터(화) 목록 파일 17개 (첨부 회차 구성 준수)
  - Notes: 완료 2026-08-01. 파트 17개 전부 작성(화 번호는 첨부 원안과 동일: 1~200).
- [x] **T-010** `done` — ③ architecture: `chapters/001~005` 챕터 카탈로그 (thin)
  - Notes: 완료 2026-08-01. Role/Conflict/Chapter Arc/Hooks/Character Arcs.
- [x] **T-011** `done` — ③ gate: architecture 일괄 Architect 승인
  - Notes: Architect-approved gate 2026-08-01. Prior-Design Consistency 선확인(parts↔catalogs 정합).

## Backlog

### ④ chapter design — Part 1 (1~5화)
- [x] **T-012** `done` — ④ design: `chapters/001-죽은-천마-망나니로-눈을-뜨다.md` (episode/scene 전개 설계)
  - Notes: 완료 2026-08-01. 3에피소드/6씬, ~5,000자. Manuscript Readiness 통과.
- [x] **T-013** `done` — ④ design: `chapters/002-보약이라는-이름의-독.md` (episode/scene 전개 설계)
  - Notes: 완료 2026-08-01. 혼백산 확정·남궁휘 가면·죽는 척 전환.
- [x] **T-014** `done` — ④ design: `chapters/003-뼈를-다시-엮다.md` (episode/scene 전개 설계)
  - Notes: 완료 2026-08-01. 골혈재조공 실행·노의원 귀뺨·"기적" 파문.
- [x] **T-015** `done` — ④ design: `chapters/004-정파-놈들-수준.md` (episode/scene 전개 설계)
  - Notes: 완료 2026-08-01. 남궁혁·남궁진 첫 등장·가문 관찰·밤의 각오.
- [x] **T-016** `done` — ④ design: `chapters/005-망나니의-눈에-마교-교주가-들어앉다.md` (episode/scene 전개 설계)
  - Notes: 완료 2026-08-01. 야습 역공·심문·쪽지 — Part 2 연결.
- [x] **T-017** `done` — ④ gate: 001~005 설계 일괄 Architect 승인 (Consistency+Readiness)
  - Notes: Architect-approved gate 2026-08-01. 각 설계의 Prior-Design Consistency Gate·Manuscript Readiness 전부 통과 확인. 동원(형제 처소) 위치 추가 — 아키텍처 정합화.

### ⑥ design evaluation (Part 1)
- [x] **T-018** `done` — ⑥ design eval: `evaluations/001~005` (Target Reader·Genre Critic·Plot Expert·Reader-Editor·Setting/Lore·Character Critic) + Schema 게이트
  - Notes: 완료 2026-08-01. 3개 비평 그룹 병렬 실행. Schema ✅ 전 챕터 통과. 결과서: `evaluations/001~005-design.md`.

### ⑦ revision (if needed)
- [x] **T-019** `done` — ⑦ design revision (평가 반영, 필요 시)
  - Notes: 완료 2026-08-01. High 5건 + Med 8건 전면 반영.

### ⑤ generation — Part 1 (1~5화)
- [x] **T-020** `done` — ⑤ manuscript: `manuscripts/001-죽은-천마-망나니로-눈을-뜨다.md` (4,647자)
- [x] **T-021** `done` — ⑤ manuscript: `manuscripts/002-보약이라는-이름의-독.md` (4,523자)
- [x] **T-022** `done` — ⑤ manuscript: `manuscripts/003-뼈를-다시-엮다.md` (4,514자)
- [x] **T-023** `done` — ⑤ manuscript: `manuscripts/004-정파-놈들-수준.md` (4,506자)
- [x] **T-024** `done` — ⑤ manuscript: `manuscripts/005-망나니의-눈에-마교-교주가-들어앉다.md` (4,513자)
- [x] **T-025** `done` — ⑤ gate: 원고 일괄 Architect 승인 (Design-Fidelity)
  - Notes: Architect-approved gate 2026-08-01. 전 화 Fidelity/Engagement/Literary 통과.

### ⑥ manuscript evaluation + ⑦ revision (Part 1)
- [x] **T-026** `done` — ⑥ manuscript eval: `evaluations/001~005` (Literary Critic 추가)
  - Notes: 완료 2026-08-01. 6페르소나 일괄 평가.
- [x] **T-027** `done` — ⑦ manuscript revision (필요 시)
  - Notes: 완료 2026-08-01. 002 S2 수정 적용.

### ⑧ release — Part 1
- [x] **T-028** `done` — ⑧ release: `continuity/001~005-summary.md` + `story-so-far.md` + `unresolved-threads.md`
  - Notes: 완료 2026-08-01. Part 1 릴리스 확정.

### Part 2 (6~15화: 남궁세가 청소 작업) — Released (45,484자)
- [x] **T-029** `done` — ③ 보강: Part 2 아키텍처 캐스케이드 (장로 2명 프로필·노의원 span·동원/약방 뷰·스테이징 3종)
  - Notes: 완료 2026-08-01. 신규 장로 프로필 및 스테이징 3종 작성 완결.
- [x] **T-030** `done` — ④ design: `chapters/006~010` (살아난 시한폭탄 ~ 가회의 날)
  - Notes: 완료 2026-08-01. 각 3에피소드/6씬 설계 완료.
- [x] **T-031** `done` — ④ design: `chapters/011~015` (증거 따위 필요 없고 ~ 청소 완료 보고)
  - Notes: 완료 2026-08-01. 각 3에피소드/6씬 설계 완료.
- [x] **T-032** `done` — ④ gate: 006~015 설계 일괄 Architect 승인
  - Notes: 완료 2026-08-01. Architect-approved gate (`evaluations/006~015-design-gate.md`).
- [x] **T-033** `done` — ⑥ design eval: `evaluations/006~015-design.md` (6페르소나 + Schema)
  - Notes: 완료 2026-08-01. 6페르소나 3그룹 병렬 평가 완료 (`evaluations/006~015-design.md`).
- [x] **T-034** `done` — ⑦ design revision (평가 반영)
  - Notes: 완료 2026-08-01. 32개 리비전 항목 전면 반영.
- [x] **T-035** `done` — ⑤ manuscript: `manuscripts/006~008`
  - Notes: 완료 2026-08-01. 006(4,815자), 007(4,521자), 008(4,517자).
- [x] **T-036** `done` — ⑤ manuscript: `manuscripts/009~011`
  - Notes: 완료 2026-08-01. 009(4,505자), 010(4,500자), 011(4,512자).
- [x] **T-037** `done` — ⑤ manuscript: `manuscripts/012~015`
  - Notes: 완료 2026-08-01. 012(4,506자), 013(4,558자), 014(4,537자), 015(4,513자).
- [x] **T-038** `done` — ⑤ gate: 006~015 원고 일괄 Architect 승인 (Design-Fidelity)
  - Notes: 완료 2026-08-01. Architect-approved gate (`evaluations/006~015-manuscript-gate.md`).
- [x] **T-039** `done` — ⑥ manuscript eval + ⑦ revision (Part 2)
  - Notes: 완료 2026-08-01. 6페르소나 통합 원고 평가 완결 (`evaluations/006~015-manuscript.md`). ALL S-TIER PASS.
- [x] **T-040** `done` — ⑧ release: continuity Part 2 (summary + story-so-far + unresolved-threads 갱신)
  - Notes: 완료 2026-08-01. `continuity/006~015-summary.md` 작성 및 `story-so-far.md`, `unresolved-threads.md` Part 2 릴리스 동기화 완결.

### 이후 파트 (T-041+)
- [x] **T-041** `done` — ④~⑧ Part 3 (16~30화): 위선자 첫째 형과 맹주의 방관
  - Notes: 완료 2026-08-01. Part 3 release records T-056~062.
- [x] **T-042** `done` — ④~⑧ Part 4 (31~50화): "마교보다 더한 놈이 나타났다"
  - Notes: 완료 2026-08-01. 20화, 92,393자, 청년단 최고 위치 도달.
- [x] **T-043** `done` — ④~⑧ Part 5 (51~62화): 흙수저 무사와의 첫 만남 (2단계 시작)
  - Notes: 완료 2026-08-01. 12화, 56,585자, 청운 영입 및 2단계 전환.
- [x] **T-044** `done` — ④~⑧ Part 6 (63~75화): 대문파의 착복과 음해 적발
  - Notes: 완료 2026-08-01. 13화, 59,953자, 청풍문 누명 해결 및 고진 패배.
- [x] **T-045** `done` — ④~⑧ Part 7 (76~88화): 천하무도대회 & 위선자 챔피언 참교육
  - Notes: 완료 2026-08-01. 13화, 59,357자, 제갈운 폭로 및 약침 단서.
- [x] **T-046** `done` — ④~⑧ Part 8 (89~100화): 강호 흑시장 털어먹기
  - Notes: 완료 2026-08-01. 12화, 55,877자, 귀시장 70% 협상 및 청문회 훅.
- [x] **T-047** `done` — ④~⑧ Part 9 (101~112화): 구파일방 청문회 박살내기
  - Notes: 완료 2026-08-01. 12화, 55,?자, 청문회 역전.
- [x] **T-048** `done` — ④~⑧ Part 10 (113~125화): 사파와의 마교식 교섭
  - Notes: 완료 2026-08-01. 13화, 혈천곡 사흘 협정.
- [x] **T-049** `done` — ④~⑧ Part 11 (126~138화): 황실과 관가의 수작질 차단
  - Notes: 완료 2026-08-01. 13화, 관무쌍청 회복.
- [x] **T-050** `done` — ④~⑧ Part 12 (139~150화): 신무단 출범 & 3단계 복선
  - Notes: 완료 2026-08-01. 12화, 신무단과 정마 접선 복선.
- [x] **T-051** `done` — ④~⑧ Part 13 (151~162화): 정·마 밀약의 실체 파헤치기
  - Notes: 완료 2026-08-01. 12화, 밀통 장부 확보.
- [x] **T-052** `done` — ④~⑧ Part 14 (163~174화): 가짜 천마와 천마신교 내분
  - Notes: 완료 2026-08-01. 12화, 천마신교 재장악.
- [x] **T-053** `done` — ④~⑧ Part 15 (175~186화): 무림맹 대청소 & 남궁세가 완전 정복
  - Notes: 완료 2026-08-01. 12화, 부자 결투와 장로 축출.
- [x] **T-054** `done` — ④~⑧ Part 16 (187~195화): 최종 결전 (정·마 썩은물 연합 vs 신무단)
  - Notes: 완료 2026-08-01. 9화, 통합 군단과 신무공 승리.
- [x] **T-055** `done` — ④~⑧ Part 17 (196~200화): 에필로그 — 무림의 새로운 절대자 (Option A 엔딩)
  - Notes: 완료 2026-08-01. 5화, 천하무림연합과 완결.

### Part 3 execution (016~030화)
- [x] **T-056** `done` — ③ architecture cascade: Part 3 catalog audit, cast/location/staging readiness
  - Notes: 완료 2026-08-01. 무림맹 본산·점창파 영지 프로필과 016 취운각 문서 확인 staging 추가; Part 3 설계와 cascade 정합화.
- [x] **T-057** `done` — ④ design: `chapters/016~030`
  - Notes: 완료 2026-08-01. 15개 장 모두 3 Episode/6 canonical scenes, Load/Consistency/Readiness PASS.
- [x] **T-058** `done` — ⑥ design evaluation + ⑦ revision: Part 3
  - Notes: 완료 2026-08-01. 6인 패널 평가 및 Architect adjudication 완료; `evaluations/016~030-design.md`.
- [x] **T-059** `done` — ⑤ manuscripts: `manuscripts/016~030`
  - Notes: 완료 2026-08-01. 15개 원고, 총 71,647자, 전부 4,500자 이상.
- [x] **T-060** `done` — ⑤ Design-Fidelity gate: Part 3 manuscripts
  - Notes: Architect-approved 2026-08-01; `evaluations/016~030-manuscript-gate.md`.
- [x] **T-061** `done` — ⑥ manuscript evaluation + ⑦ revision: Part 3
  - Notes: 완료 2026-08-01. Target Reader/Genre/Reader-Editor/Literary/Character/Setting 패널 adjudication PASS.
- [x] **T-062** `done` — ⑧ release: Part 3 continuity sync
  - Notes: 완료 2026-08-01. `continuity/016~030-summary.md`, story-so-far, unresolved-threads 동기화 완료.

### Part 4 + Part 5 execution (031~062화)
- [x] **T-063** `done` — ③ architecture cascade: Part 4/5 cast, locations, stagings, chapter-list audit
  - Notes: 완료 2026-08-01. 청운·청풍문 아키텍처 및 청년단/순찰 staging 추가; 중복 프로필 정리.
- [x] **T-064** `done` — ④ design: Part 4 `chapters/031~050` and Part 5 `chapters/051~062`
  - Notes: 완료 2026-08-01. 32개 설계 Load/Consistency/Readiness PASS.
- [x] **T-065** `done` — ⑥ design evaluation + ⑦ revision: Parts 4~5
  - Notes: 완료 2026-08-01. `evaluations/031~062-design.md`, Architect adjudication PASS.
- [x] **T-066** `done` — ⑤ manuscripts: `manuscripts/031~062`
  - Notes: 완료 2026-08-01. 32개 원고, Part 4 92,393자 + Part 5 56,585자.
- [x] **T-067** `done` — ⑤ Design-Fidelity gate: Parts 4~5
  - Notes: 완료 2026-08-01. `evaluations/031~062-manuscript-gate.md` PASS.
- [x] **T-068** `done` — ⑥ manuscript evaluation + ⑦ revision: Parts 4~5
  - Notes: 완료 2026-08-01. 6인 패널 adjudication PASS.
- [x] **T-069** `done` — ⑧ release: Parts 4~5 continuity sync
  - Notes: 완료 2026-08-01. `continuity/031~062-summary.md`, story-so-far, unresolved-threads 동기화.

### Parts 6–8 execution (063~100화)
- [x] **T-070** `done` — ③ architecture cascade: Parts 6–8 cast, locations, stagings, chapter-list audit
  - Notes: 완료 2026-08-01. 고진·제갈운·흑야차 및 외부 장소/귀시장 아키텍처 추가·정합화.
- [x] **T-071** `done` — ④ design: `chapters/063~100`
  - Notes: 완료 2026-08-01. 38개 설계 Load/Consistency/Readiness PASS.
- [x] **T-072** `done` — ⑥ design evaluation + ⑦ revision: Parts 6–8
  - Notes: 완료 2026-08-01. `evaluations/063~100-design.md`, Architect adjudication PASS.
- [x] **T-073** `done` — ⑤ manuscripts: `manuscripts/063~100`
  - Notes: 완료 2026-08-01. 38개 원고, 총 175,187자; 089~094 과다 분량을 4,500~5,500자로 재작성.
- [x] **T-074** `done` — ⑤ Design-Fidelity gate: Parts 6–8
  - Notes: 완료 2026-08-01. `evaluations/063~100-manuscript-gate.md` PASS.
- [x] **T-075** `done` — ⑥ manuscript evaluation + ⑦ revision: Parts 6–8
  - Notes: 완료 2026-08-01. 6인 패널 adjudication PASS.
- [x] **T-076** `done` — ⑧ release: Parts 6–8 continuity sync
  - Notes: 완료 2026-08-01. `continuity/063~100-summary.md`, story-so-far, unresolved-threads 동기화.

## Notes

**계획 의도 (plan intent):**
- 장르/톤: 무협 회귀 사이다물. 통쾌·개그·속도감. 정파의 위선을 마교식으로 부수는 재미.
- 독자: 무협/사이다 웹소설 독자.
- 릴리스 단위: 에피소드(화) 1개 = 10~20분 분량(약 4,500~5,500자).
- **Part 1 (001~005화) 릴리스 완료:** 22,693자.
- **Part 2 (006~015화) 릴리스 완료:** 45,484자. (006~015 10개 화 전부 ≥ 4,500자 충족).
- Architect 게이트 정책: 스킬의 "user approval" 게이트는 Architect가 전문가 자격으로 승인/진행. Publisher 결정만 사용자 대기.
- **Part 3 (016~030화) 릴리스 완료:** 71,647자. 남궁혁 사회적 매장, 남궁진 첫 정면 대화, 무림맹 외부 표적 전환.
- **Part 4 (031~050화) 릴리스 완료:** 92,393자. 청년단 최고 위치, 정파 맹주 아들 신분의 역이용.
- **Part 5 (051~062화) 릴리스 완료:** 56,585자. 청운 영입, 진짜 협과 2단계 전환.
- **Part 6 (063~075화) 릴리스 완료:** 59,953자. 청풍문 누명 해결, 대문파 착복 폭로.
- **Part 7 (076~088화) 릴리스 완료:** 59,357자. 무도대회 약물·혈도 파괴 폭로.
- **Part 8 (089~100화) 릴리스 완료:** 55,877자. 귀시장 장악, 민초 환원, 청문회 동맹 훅.
- **Parts 9~17 (101~200화) 릴리스 완료:** 543,888자. 청문회부터 천하무림연합까지 전체 결말 완성.
- **2026-08-02 continuity republication audit:** Publisher requested sequential audit/revision beginning with Chapters 001~005. Because canonical `continuity/` currently represents the completed 200-chapter series, the audit uses `continuity/audit-part1/` as a sequential working ledger; canonical continuity will be reconciled after Chapter 005.

### Sequential continuity audit (Publisher-requested)
- [x] **T-084** `done` — 001~005 sequential continuity republication audit/revision/release
  - Notes: 완료 2026-08-02. Chapters 001~005 each delegated to a subagent in strict order; all five PASS after design-first revisions. Audit ledger and canonical continuity reconciled. Findings applied: 3 + 6 + 5 + 5 + 8 = 27.
- [ ] **T-085** `in_progress` — Chapters 006~200 sequential continuity republication audit/revision/release
  - Notes: Publisher requested full sequential quality pass. Chapters 006~150 completed serially; each chapter was delegated, revised/released, and written into `continuity/audit-serial/` before the next. No parallel chapter prose generation. Cumulative ledger integrity is being checked when subtask reports omit exact paths.
