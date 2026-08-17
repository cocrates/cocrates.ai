# Design Evaluation: Episode 091 — 사흘의 약속

## Evaluation Scope & Target Reader
- **Target Reader:** 성인 남성향 회귀·빙의·환생 무협과 문파 장악물, 가족 반전과 복수형 사이다를 선호하는 웹소설 독자 (`overview.md`).
- **Evaluated design:** `episodes/091-사흘의-약속.md`
- **Architect G4 input:** Design Consistency ✅ · Generation Readiness ✅. This evaluation is mandatory Stage ⑤.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 / series-level criterion; not this late-arc design scope. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Early-arc criterion already satisfied outside Episode 091 scope. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 2의 전면전 대 소수 침투 대립, Scene 4의 침투조 확정이라는 실질적 사건과 다음 후크가 있다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1 criterion; Episode 091 is P3. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series-level late-arc payoff criterion; Episode 091 uses the already-established 약 그릇·가환 but does not complete the series payoff. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 미래 지식으로 원액 위치를 예언하지 않고, 현재 지형·병세·증거를 계산해 소수 침투를 선택한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Endgame criterion; not executed at design scope. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript quality belongs to Stages ⑥–⑦; forecast is checked below under Schema. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | Four unique `### Scene` sections use the required meta lines and flat bullet fields. |
| No skill/workflow dump after the design | ✅ | Episode body contains no pasted workflow procedure; only its own gate/readiness results. |
| Unique Scene headings; no pasted twin scenes | ✅ | Scene 1–4 have distinct functions: medical limit, strategy conflict, role allocation, named addition. |
| Canonical episode path | ✅ | Exact path `episodes/091-사흘의-약속.md`. |
| Field notation | ✅ | Required fields use `**Field:**` / `- **Field:**`. |
| Every scene has required fields | ✅ | Each scene contains POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, 6–8 outline lines, Unit budget, one Est. length. |
| Characters Appearing ↔ On stage union | ✅ | Union of Scene 1–2 cast plus Scene 4’s `진우의 스승의 제자` equals the six-name Appearing list; no ghost cast. |
| On stage includes speakers | ✅ | All dialogue-intent speakers and named actors in Beats are present on their scene’s On stage list. |
| Characters ⊆ `characters.md` | ✅ | Six profiles read OK this turn: `characters/서진우.md`, `서도현.md`, `남궁혁.md`, `가환.md`, `의원의-제자.md`, `진우의-스승의-제자.md`. |
| Summary/Hooks cast alignment | ✅ | Summary, Hooks, Seeds, Closing mention only the six Appearing characters. |
| No later-list cast debut | ✅ | The series row already names “옛 스승의 제자”; it was additively cataloged as `진우의 스승의 제자` before citation. |
| Locations ⊆ Key Locations | ✅ | `전쟁의-계곡` is a Key Locations slug in `locations.md`; both facets belong to that place. |
| Location facets ⊆ Multi-facet anchors | ✅ | `locations/전쟁의-계곡.md` read OK: anchors are `진입로 / 마을터 / 수원 / 절벽길`; design uses `마을터` and `절벽길` exactly. |
| Nested scene files absent | ✅ | Single canonical episode file; no nested scene directory. |
| No template residue | ✅ | No unresolved `{...}` body placeholders. |
| Prose forecast present | ✅ | Every Unit budget uses only dialogue/action/sensory/POV/transition typed integer terms. |
| Forecast ↔ Est. cross-check | ✅ | Independent recomputation: Sc1 `3×260+2×200+1×140+1×180+1×100=1,600`; Sc2 `4×260+2×200+1×140+1×180+1×100=1,800`; Sc3 `4×260+2×200+2×140+1×180+1×100=2,000`; Sc4 `4×260+2×200+1×140+1×180+1×100=1,800`. Each Est equals its subtotal and lies within outline density. |
| Dialogue intent vs outline speech | ✅ | Each scene with speech has non-`none` Dialogue intent and matching on-stage speakers. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene Est. fields: `1,600 + 1,800 + 2,000 + 1,800 = 7,200`; header addends: `1,600 + 1,800 + 2,000 + 1,800 = 7,200`. |
| Est. length sum ≥ Scale min | ✅ | 7,200 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 7,200 ≤ 8,000; central target band is met at its upper edge. |
| Cited staging/profile paths exist | ✅ | This-turn reads succeeded for `stagings/091-사흘의-작전.md`, `stagings/091-침투조-선정.md`, `locations/전쟁의-계곡.md`, and all six character profile paths. |
| Episode List plot | ✅ | Series Summary’s three clauses map to Scene 1 (three-day objective), Scene 2 (Hyuk vs. Jinwoo), and Scene 4 (Dohyun adds one name). |
| Hook evidence strength (internal) | ✅ | Body quotes: series Hook「침투조 명단에 도현이 직접 한 사람을 추가한다」; Summary「도현은 침투조 명단에 직접 한 사람을 추가한다」; Out「침투조 명단에 도현이 직접 한 사람이 추가된 채」; Seed「도현이 이름을 적으며」; Scene 4 Turn「도현은 ... 한 이름을 적는다」. Same obligation strength. |
| Hook scope | ✅ | Closing has one primary obligation—Dohyun’s direct addition—and no invented chase, faction arrival, or second reveal. |
| No design-paste / meta-only scenes | ✅ | Each scene has a concrete observable event and distinct causal function; no repeated Unit budget or Beat sentences. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | Medical limit → strategic conflict → geographic role allocation → named addition is a causal escalation. |
| Scene transitions chain | ✅ | S1’s battle proposal leads to S2; S2’s order to move leads to S3; S3’s blank list leads to S4; S4 exits into infiltration. |
| Scene sections complete | ✅ | All four Scene Index rows have complete generation briefs. |
| Generation Readiness | ✅ | All Schema rows pass; no unresolved design-field Apply finding. |
| Beat concreteness | ✅ | Blood comparison, troop proposal, map/facet check, list inscription, and conditional acceptance are observable actions. |
| Est. length budget | ✅ | Independent sum 7,200; Scale pass; no padding-only scene. |
| Prose forecast quality | ✅ | Unit types correspond to the medical verification, argument, map action, and naming reveal. |
| Episode List scope aligned | ✅ | No early reveal of the elixir location or mother’s survival; only planned infiltration setup is executed. |
| Prior hook addressed | ✅ | Scene 1 immediately converts the prior episode’s three-day promise into a verified deadline. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected | ✅ | Prior Design Alignment records Phase A, selective Phase B, staging details, continuity pair, and canonical episode path. |
| Series / overview tone & arc honored | ✅ | Cold, event-centered P3 strategy with family stake; no romance or training detour. |
| Episode List Summary / Hook honored | ✅ | All Summary clauses and the exact Hook obligation are quoted and mapped to scenes. |
| Hook internal consistency | ✅ | Summary / Arc close / Out / Seed / Scene 4 Turn agree on one direct addition. |
| Characters from architecture | ✅ | All six appear in Character Catalog and have profile paths; the new profile was added before design citation. |
| Profile-backed knowledge / recognition | ✅ | Dohyun’s choice is backed by his relationship in `진우의 스승의 제자.md`; the disciple’s knowledge is limited to the profile’s Core Drive and relationship, not an unsupported full identity reveal. |
| Locations from architecture | ✅ | Scene locations are exact Key Location slug + exact facet. |
| Location profile paths readable | ✅ | `locations/전쟁의-계곡.md` read OK this turn. |
| Location facets ⊆ anchors | ✅ | `마을터` and `절벽길` are exact anchors in the read profile. |
| Stagings from episode design | ✅ | Both stagings were authored in Stage ④, cite profile states, and match cast/blocking; both read OK this turn. |
| World rules / history consistent | ✅ | Contract effects remain delayed reaction, poison flare, memory/trace risk; no cost-free supernatural act is added. |
| No improvised entities or silent lore | ✅ | The only new character is additive architecture; no new faction/place/rule is introduced. |
| Continuity files used | ✅ | Immediate prior summary and story-so-far are explicitly loaded and cited. |
| Character/location state vs story-so-far | ✅ | Jinwoo remains 봉인-해제 파열; Dohyun 병세-노출; valley remains restricted and damaged; no state reset. |
| Unresolved threads categorized | ✅ | Eight active threads are advanced, planted, or held in the design. |
| No contradiction of released continuity | ✅ | Dohyun remains injured and the messenger remains ignorant of the elixir location. |
| Conflicts section empty or escalated | ✅ | The only catalog gap was escalated additively and closed before Key Events. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — selective load and all cited paths passed.
- Locations: ✅ — index, path, and facet evidence are separately present in the design file.
- Length / Prose forecast: ✅ — every scene has written/recomputed equality and header/field equality at 7,200.
- Episode List Summary: ✅ — each signature clause maps to one named Scene Beat.
- Hook to Next / Closing: ✅ — canonical Hook, Summary, Out, Seed, and closing Turn quote the same direct-addition obligation.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | “사흘 안에 원액을 훔치려면 진우는 누구를 믿고, 도현은 누구를 믿을 것인가?” |
| Personal stake present | ✅ | Dohyun’s death deadline and Jinwoo’s exposed seal make delay and visibility costly. |
| Episode Out hook | ✅ | Dohyun directly writes one person into the infiltration list; it is earned by Scenes 1–3’s selection process. |
| Exposition budget respected | ✅ | Medical information is limited to the decision-relevant reaction-delay distinction; mother/elixir location remains held. |
| Seed discipline | ✅ | One Plant and one Hint, with explicit Holds. |
| Scene-first Key Events | ✅ | Four scenes carry action, turn, sensory cue, dialogue intent, and transition. |
| Sensory-emotional on every scene | ✅ | Each scene pairs a concrete valley detail with a POV reaction. |
| Motifs planned across scenes | ✅ | Finger-counting and folded map recur across the arc. |
| Overview signature line | N/A | `overview.md` has no required signature dialogue line. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | Medical facts appear only as observed evidence under immediate tactical pressure. |
| Sensory-emotional pairing | ✅ | Cold light, ash, hoof sound, and red thread are each attached to a character calculation or reaction. |
| Dialogue voices + intent | ✅ | Hyuk presses openly, Jinwoo calculates tersely, Gahwan limits claims, Dohyun speaks slowly, and the disciple answers narrowly. |
| Reader-discovered meaning | ✅ | Reader infers Dohyun is still choosing and protecting through action; no thematic monologue is planned. |
| Antagonist plausibility | ✅ | The absent Black Wind Pavilion is treated as an alert-sensitive opponent, not a passive target. |
| Closing image specified | ✅ | Dohyun’s black finger stops on the new name. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | Opens on a dying father’s measurable deadline, not a recap. |
| Personal stake matches reader expectation | ✅ | Family survival, tactical reversal, and distrust of the father’s choice are immediate. |
| Pacing / density fits platform expectations | ✅ | Four concrete turns, 7,200-character forecast, no training or briefing-only scene. |
| Out hook makes this reader want the next episode | ✅ | A suspected traitor’s teacher-lineage enters the infiltration team by the father’s order. |
| No alienation of core audience without overview intent | ✅ | Maintains male-oriented martial action, strategic conflict, and family-revenge tension. |

## Design Critique (required personas)
#### Target Reader
- **Stance:** This reader wants tactical momentum, family pressure, and a sharp next-episode suspicion.
- **Strengths:** The three-day deadline is concrete; the full-war versus infiltration dispute gives immediate strategic friction; the father’s handwritten addition is a strong serial Out.
- **Defects:** —
- **Reader impact:** High retention potential because the new ally is both useful and distrusted.

#### Genre Critic
- **Stance:** Tests 회귀 무협·문파 장악물 promises of decisive planning and earned reversal.
- **Strengths:** Jinwoo’s future knowledge is subordinated to present evidence; Hyuk’s proposal creates a credible tactical split; no effortless elixir acquisition occurs.
- **Defects:** —
- **Reader impact:** Delivers genre satisfaction without spending the infiltration payoff early.

#### Plot Expert
- **Stance:** Checks causality, hook body alignment, and Hook scope.
- **Strengths:** Scene transitions form a clean chain; every Summary clause lands in a named scene; all Hook body surfaces preserve the direct-addition strength; Out has one obligation.
- **Defects:** —
- **Reader impact:** The next episode’s suspicion is caused by the current decision rather than appended as a teaser.

#### Reader-Editor
- **Stance:** Checks serialization, exposition restraint, and closing density.
- **Strengths:** Scene 1 turns medical explanation into pressure; Scenes 2–3 prevent the operation from becoming a generic briefing; the final name is specific and actionable.
- **Defects:** —
- **Reader impact:** Low skim risk; each scene changes the tactical question.

#### Character Critic
- **Stance:** Checks motivation, distinct voices, and profile-backed knowledge.
- **Strengths:** Dohyun’s choice follows his protective drive; Jinwoo’s suspicion follows his unresolved teacher betrayal; the new disciple’s limited knowledge is supported by his profile and does not over-reveal.
- **Defects:** —
- **Reader impact:** The father-son conflict advances through a decision, not reconciliation dialogue.

#### Literary Critic
- **Stance:** Checks motif, sensory logic, and the non-expository closing image.
- **Strengths:** Finger-counting and the folded map give bodily and spatial form to trust/time; the ending image lets the name carry the emotional weight.
- **Defects:** Motif touches are episode-level rather than explicit scene fields.
- **Reader impact:** This is a Low craft risk, not a reader-blocking design defect; Stage ⑥ should preserve both motifs visibly.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | Episode motifs need visible scene-level placement (Literary Critic) | Low | No | yes | Reader benefit comes from retaining the finger-count and folded-map images during prose generation; changing the plot design is unnecessary. | Carry to Stage ⑥: preserve finger-counting in Scenes 1/2/4 and folded-map contact in Scenes 2/3/4; do not explain the motifs thematically. | Carry-⑥ |

No High or Med design findings. No Pending revision remains. The adjacent Scenes 3–4 share `절벽길` but have visibly distinct Functions and Turns: geographic/role selection versus named-person reveal; this is not a duplicate scene or a hidden catalog gap.

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

**Gate G5:** Architect approved 2025-02-14. The design evaluation passed with Generation-ready ✅; Carry-⑥ motif constraints are handed to Stage ⑥. Proceed immediately to manuscript generation.
