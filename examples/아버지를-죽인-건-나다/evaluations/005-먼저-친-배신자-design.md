# Design Evaluation: Episode 005 — 먼저 친 배신자

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope; Episode 005 does not re-execute the opening setup. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003 scope; this is an already-established early-arc criterion. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 공개 심문으로 내외당의 책임 구조를 압박하고 백무진의 협력 조건을 확보한다. Scene 3 Turn에서 미래 사망 기록을 물질적 증거로 제시한다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | Episode 005는 서도현의 내면을 공개하지 않고, 흑풍루 표식·장부·내부 배신 구조를 중심으로 P1의 복수 판을 유지한다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반 회수 기준이며 Episode 005 설계 범위를 넘어선다. 이 회차는 기존 운송표·표식 스레드를 전진시키고 후속 기록을 Plant한다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 미래의 배신 구조를 공개 심문에 활용하지만, 미래 사망 기록의 작성자·날짜 상세·사건 전문은 Hold한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 결말·후반 정서적 절정 기준이다. Episode 005는 P1 내부 장악 단계에 해당한다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 분량·밀도는 Stage ⑥/⑦에서 판정한다. Forecast arithmetic는 아래 Schema에서 별도로 검증한다. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | All three scenes use the canonical meta fields and flat bullet fields. |
| No skill/workflow dump after the design | ✅ | No workflow instruction sections such as `## Pre-Design Load` or copied stage templates appear as post-design content. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Exactly one each of `### Scene 1`, `### Scene 2`, and `### Scene 3`; situations and turns are distinct. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | Actual design path: `episodes/005-먼저-친-배신자.md`. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | Scene metadata and all required bullets use the required notation. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | Each scene has POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, 7-line Paragraph outline, typed Unit budget, and exactly one Est. length. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing cast is 서진우·윤태석·백무진·장로-대표; every scene On stage contains all four, with no ghost cast. |
| On stage includes speakers | ✅ | Every intentional speaker named in each scene's Beat and Dialogue intent is one of the four On-stage characters. Rear observers are ambient and have no attributable speech. |
| Characters ⊆ `characters.md` | ✅ | `characters.md` contains all four; this turn read the profiles `characters/서진우.md`, `characters/윤태석.md`, `characters/백무진.md`, and `characters/장로-대표.md`. |
| Summary/Hooks cast alignment | ✅ | Summary, Episode Arc, Seeds, and Closing refer only to appearing/catalogued characters; the final hook names 진우, 윤태석, and the record already tied to the appearing cast. |
| No later-list cast debut | ✅ | Episode 005's four named characters are catalogued for Episode 005 or earlier; no later Episode List cast is introduced. |
| Locations ⊆ `locations.md` Key Locations | ✅ | Sc1–Sc3 use `북문서가-본가`, which maps to the Key Locations row `북문서가 본가`; `장로회당 표결단` is an exact facet of that key location. |
| Location facets ⊆ Multi-facet anchors | ✅ | `locations/북문서가-본가.md` read OK; `장로회당 표결단` is an exact `Multi-facet anchors` label. |
| Nested `episodes/{slug}/` scene files absent | ✅ | The episode is a single file; no nested scene directory is cited or present. |
| No template residue | ✅ | No unresolved `{placeholder}` or instructional template braces remain. |
| Prose forecast present (outline + typed units) | ✅ | Every scene has a 7-line paragraph outline and only the five permitted unit types. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Scene 1: `3×260 + 3×180 + 2×130 + 2×150 + 1×80 = 1,960`; recomputed 1,960; Est. 2,000. Scene 2: `3×250 + 3×180 + 2×120 + 2×140 + 1×80 = 1,890`; recomputed 1,890; Est. 1,900. Scene 3: `3×260 + 3×180 + 2×130 + 2×150 + 1×90 = 1,970`; recomputed 1,970; Est. 2,000. Each Est. lies within its 7-line outline density band. |
| Dialogue intent vs outline speech | ✅ | Every scene contains speech in Beat/outline and has non-`none` Dialogue intent naming only On-stage speakers. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: `2,000 + 1,900 + 2,000 = 5,900`; header addends: `2,000 + 1,900 + 2,000 = 5,900`. |
| Est. length sum ≥ Scale min (hard) | ✅ | Recomputed scene-field sum is 5,900; overview Scale minimum is 4,000. |
| Est. length sum ≤ Scale max (hard) | ✅ | Recomputed scene-field sum is 5,900; overview Scale maximum is 8,000. |
| Cited staging/profile paths exist | ✅ | This turn read OK: `characters/서진우.md`, `characters/윤태석.md`, `characters/백무진.md`, `characters/장로-대표.md`, `locations/북문서가-본가.md`, and `stagings/005-내외당-공개-압박.md`. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary says 진우 publicly pressures the future-using internal disciples to gain the first internal force, while one pretends to follow 흑풍루 and hides another purpose. Scene 1–2 execute the public pressure and first internal alliance; Scene 2–3 execute 윤태석's false-command posture and hidden purpose. The Summary's future-death-record reveal maps to Scene 3 Turn. |
| Hook evidence strength (internal) | ✅ | Body quotes align at the same strength: `series.md` Hook to Next 「거짓 밀고자의 품에서 진우 자신의 미래 사망 기록이 나온다」; Episode Summary ends with the record emerging; Episode Out repeats it; Scene 3 Turn states the record emerges; Seeds classifies the physical record as `Plant / Hook`. |
| Hook scope (no Out creep) | ✅ | Last Transition out only folds the record under the black token and sets 백무진's next report; it adds no chase, faction arrival, or second reveal. |
| No design-paste / meta-only scenes | ✅ | Each scene has a distinct dramatic event: public accusation, conditional alliance and exposure of false terminology, then extraction of the hidden record. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | Evidence moves from public accusation → differentiated pressure → personal record reveal. |
| Scene transitions chain | ✅ | Scene 1 identifies the order of separate pressure; Scene 2 moves toward 윤태석's sleeve; Scene 3 closes the hearing after the record reveal. All `When` fields remain the morning after Episode 004's night meeting. |
| Scene sections complete | ✅ | Every Scene Index row has a matching complete scene section with all required fields. |
| Generation Readiness | ✅ | All Schema rows are ✅; no schema blocker or Pending adjudication applies. |
| Beat concreteness | ✅ | Every Beat names observable evidence handling, testimony, exposure, bargaining, or record extraction. |
| Est. length budget | ✅ | Independent total is 5,900, equal to the header and within 4,000–8,000; central target band is met. |
| Prose forecast quality | ✅ | Typed dialogue/action/sensory/POV/transition units correspond to the seven outline intents in each scene. |
| Episode List scope aligned | ✅ | Summary and Hook to Next are executed without adding an extra reveal or future payoff. |
| Prior hook addressed (ep 002+) | ✅ | Scene 1 directly uses Episode 004's black token and 북항 transport-slip fragment. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Design records overview, series, appearing profiles, used location, world aspects, staging, and Episode 004 continuity as loaded. This evaluation re-read the cited profiles and source artifacts. |
| Series / overview tone & arc honored | ✅ | Cold, event-driven pressure advances P1's internal takeover while preserving the revenge/family mystery contract. |
| Episode List Summary / Hook to Next honored | ✅ | `series.md` Episode 005 Summary and Hook are quoted and mapped in the Schema evidence above. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Episode Arc close, Out, Hook Seed, and Scene 3 Turn all preserve the same physical-record reveal. |
| Characters from architecture; profiles not redefined | ✅ | All four characters are in the catalog and their profile drives/voices are used: Jinwoo's evidence control, Taeseok's fixed gaze/hidden record motive, Mujin's conditional survival bargain, and the elder's procedural pressure. |
| Profile-backed knowledge / recognition | ✅ | 윤태석's hidden sibling-record motive and 흑풍루 coercion are supported by `characters/윤태석.md`; Jinwoo's future knowledge and observation habits are supported by `characters/서진우.md`. No unsupported identity recognition is planted. |
| Locations from architecture; profiles not redefined | ✅ | The sole Key Location maps to `locations/북문서가-본가.md`; no new place is invented in the Beats. |
| Location profile paths readable | ✅ | This-turn read OK: `locations/북문서가-본가.md`. |
| Location facets ⊆ Multi-facet anchors | ✅ | `장로회당 표결단` exactly matches the profile's anchor list. |
| Stagings from episode design; blocking not redefined | ✅ | This-turn read OK: `stagings/005-내외당-공개-압박.md`; the three scenes cite the same staging and retain its cast states, left/right positions, props, and no-observer-speech rule. |
| World rules / history consistent with bible | ✅ | The design uses documents, seals, transport routes, and distributed commands as clues without claiming a new contract function or contradicting the world bible. |
| No improvised entities or silent lore | ✅ | New character/staging needs were added before design citation; location facet and named cast are all catalogued. |
| Continuity files used (ep 002+) | ✅ | `continuity/story-so-far.md`, `continuity/004-회합의-손님-summary.md`, and `continuity/unresolved-threads.md` were loaded and cited. |
| Character/location state vs `story-so-far` | ✅ | Jinwoo retains the black token and transport slip; the public hearing adds an unreleased state without undoing Episode 004. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-001, TH-004, TH-005, and TH-006 are advanced; the sibling-record and death-record are planted; remaining questions are explicitly held. |
| No contradiction of released continuity | ✅ | No released event is reversed or rewritten. |
| Conflicts section empty or escalated (not ignored) | ✅ | Design states `None`; the loaded architecture and continuity reveal no conflict. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Loaded required artifacts | ✅ | This-turn load included `overview.md`, `series.md`, `characters.md` plus all four appearing profiles, `locations.md` plus `locations/북문서가-본가.md`, `world-bible.md` plus `world/무림-세력.md` and `world/혈맥계약과-약그릇.md`, `stagings/005-내외당-공개-압박.md`, and the Episode 004 continuity set. |
| Locations (index) | ✅ | Sc1–Sc3: `북문서가-본가` ∈ Key Locations; each uses the facet `장로회당 표결단`. |
| Locations (path) | ✅ | Exact paths read OK this turn: `locations/북문서가-본가.md` and `stagings/005-내외당-공개-압박.md`; all four cited character profile paths also read OK. |
| Locations (facets) | ✅ | Sc1–Sc3: `북문서가-본가 + 장로회당 표결단` ⊆ exact `Multi-facet anchors`. |
| Length / Prose forecast | ✅ | Sc1 written=1,960; recomputed=1,960; Est=2,000 · Sc2 written=1,890; recomputed=1,890; Est=1,900 · Sc3 written=1,970; recomputed=1,970; Est=2,000 · header addends `2,000+1,900+2,000=5,900`. |
| Episode List Summary | ✅ | `series.md` Summary clause 「내외당 제자들을 공개적으로 압박해 첫 내부 세력을 얻는다」→ Scenes 1–2 Beats; 「한 명은 흑풍루의 명령을 받은 척 다른 목적을 숨긴다」→ Scenes 2–3 Beats; 「거짓 밀고자의 품에서 진우 자신의 미래 사망 기록이 나온다」→ Scene 3 Turn. |
| Hook to Next / Closing | ✅ | Hook 「거짓 밀고자의 품에서 진우 자신의 미래 사망 기록이 나온다」; Summary 「그러나 거짓 밀고자의 품에서 진우 자신의 미래 사망 기록이 나온다」; Out 「거짓 밀고자의 품에서 진우 자신의 미래 사망 기록이 나온다」; Scene 3 Turn 「거짓 밀고자의 품에서 진우 자신의 미래 사망 기록이 나온다」. |
| Hook scope / internal consistency | ✅ | The closing obligation is one physical reveal; the final transition adds only the next report instruction and no independent scope-creep obligation. |
| Overview signature lines | N/A | `overview.md` contains constraints rather than a designated signature dialogue line; no unplaced signature line is required. |
| Profile-backed knowledge | ✅ | Taeseok's sibling-record drive and Jinwoo's future-knowledge basis are profile-backed. |
| Opening honors prior | ✅ | Scene 1 opens with the black token and transport-slip evidence from Episode 004. |
| Continuity states | ✅ | No released state is undone; the token and slip remain Jinwoo's evidence. |
| Unresolved threads | ✅ | Existing threads are explicitly picked up/advanced/planted/held in the design. |
| Characters / speakers | ✅ | Catalog, Appearing, On stage, Beat, and Dialogue intent agree. |
| Locations / staging / world | ✅ | Index, exact profile paths, exact facet anchor, staging block, and world rules all agree. |
| Length / forecast | ✅ | Exact products and sums independently recomputed; no range-only formula, duplicate Est., or density-band failure. |
| No cross-scene paste / no template residue | ✅ | Distinct scenes and no unresolved template text. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 「진우에게 흑풍루 표식을 건넨 손은 내당과 외당 중 어느 쪽의 손인가?」 is stated and not answered immediately. |
| Personal stake present | ✅ | Jinwoo must prevent the future betrayal and confront a physical record of his own future death. |
| Episode Out hook | ✅ | The record reveal is concrete, personal, and exactly matches the series Hook to Next. |
| Exposition budget respected | ✅ | New concepts are limited to inner/outer roles, token-command connection, and disposition records; they surface through interrogation. |
| Seed discipline | ✅ | One primary Plant/Hook (death record), one Plant (sibling record), and one Hint (token's command connection), with explicit Holds. |
| Scene-first Key Events (all required fields) | ✅ | All three scenes contain concrete causal Key Events and complete schema fields. |
| Sensory-emotional on every scene | ✅ | Ink/token texture, document-box latch, and paper rustle each trigger a Jinwoo reaction. |
| Motifs planned across scenes | ✅ | 펼쳐지는 문서 and 소매 안쪽 are planned across Scenes 1–3 / 2–3. Scene-level motif-touch fields are not explicit; carry the placement as a Stage ⑥ generation constraint. |
| Overview signature line | N/A | No designated overview signature dialogue line. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | Information is delivered as testimony, document comparison, and visible behavior under public procedural pressure rather than exposition. |
| Sensory-emotional pairing | ✅ | Each scene pairs a material detail with Jinwoo's controlled physical reaction. |
| Dialogue voices + Dialogue intent | ✅ | Jinwoo is clipped and controlled; Taeseok is low and formal; Mujin is rough and conditional; the elder is procedural. These match the loaded profiles. |
| Reader-discovered meaning | ✅ | The reader infers that future knowledge cannot prevent an already-documented death; the design explicitly Holds authorship and causality. |
| Antagonist plausibility | ✅ | The elder's procedural self-interest and Taeseok's survival/record motive create pressures beyond cartoon malice. |
| Closing image specified | ✅ | A thin record bearing Jinwoo's name and an unarrived death date lies beside the black token. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | The locked audience of adult male readers of regression, sect-takeover, and revenge web fiction receives an immediate public-pressure scene built around a prior-episode token. |
| Personal stake matches what this reader came for | ✅ | The protagonist uses foreknowledge for a satisfying reversal while the threat becomes personally lethal through his own future death record. |
| Pacing / density fits platform expectations | ✅ | Three escalating scenes, 5,900-character target, low exposition, and concrete turns fit the episode contract. |
| Out hook makes *this* reader want the next episode | ✅ | A documented future death directly promises the next episode's death-ledger investigation. |
| No alienation of core audience without overview intent | ✅ | No romance detour, prolonged training, or abstract family reconciliation replaces the action-forward revenge structure. |

## Design Critique (required personas)

#### Target Reader
- Stance: Adult male web-novel reader seeking regression advantage, sect power play, revenge pressure, and a strong serialized hook.
- Strengths: The design delivers public dominance, a conditional first ally, and a concrete personal danger in a clean escalation.
- Defects: —
- Reader impact: The black token payoff is delayed just enough for Scene 1 pressure, while the final death record creates a strong click-through reason.

#### Genre Critic
- Stance: Tests the regression martial-arts and sect-takeover contract without requiring premature lore exposition.
- Strengths: Jinwoo's future knowledge functions as an actionable advantage; the public interrogation and alliance bargain provide genre-satisfying reversals.
- Defects: —
- Reader impact: Readers receive the expected “先手” pleasure while the record reveal prevents the episode from feeling like a routine victory lap.

#### Plot Expert
- Stance: Checks causality, Episode List fidelity, Hook body alignment, and closing scope.
- Strengths: The black token and transport slip causally open the hearing; testimony contradictions lead to Mujin's bargain; the sleeve investigation leads to the death record. Summary, Out, Seeds, and Scene 3 Turn all state the same hook strength.
- Defects: —
- Reader impact: The reader can follow why each pressure move happens and is not asked to absorb an unrelated second cliffhanger.

#### Reader-Editor
- Stance: Checks serialization density, exposition restraint, scene handoffs, and closing-hook load.
- Strengths: Each scene has a distinct function and observable turn; the closing carries one primary obligation rather than a pile of unrelated promises.
- Defects: —
- Reader impact: The episode should skim cleanly at platform speed, with the document-handling details acting as readable visual anchors.

#### Literary Critic
- Stance: Checks motif, sensory-emotional design, restraint, and whether meaning is left for the reader to infer.
- Strengths: Documents and sleeve interiors create a coherent visible/invisible contrast; the closing image externalizes the theme without a planned thematic monologue.
- Defects: Motif placements are episode-level rather than attached to explicit scene `Motif touch` fields → severity Low → carry the two motif placements as a generation constraint in Stage ⑥.
- Reader impact: If the manuscript preserves the physical document/sleeve recurrence without explaining it, the episode's cold pressure will gain texture rather than becoming a sequence of procedural exchanges.

#### Character Critic
- Stance: Checks motive, voice, relationship pressure, and profile-backed knowledge/recognition.
- Strengths: Taeseok's false betrayal is supported by his sibling-record drive; Mujin's bargain follows his survival drive; Jinwoo's attention to sleeves and evidence follows his profile habits; no unsupported recognition claim is introduced.
- Defects: —
- Reader impact: The two potential allies remain distinct—one hides a private purpose, the other states a public condition—so the reader has reason to track both beyond this episode.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Motif placement is episode-level rather than scene-field explicit (Literary Critic) | Low | No | yes | The target reader benefits from recurring document/sleeve imagery, but the design already specifies the placements by scene; a generation handoff constraint is sufficient and avoids unnecessary design churn. | Stage ⑥ must visibly thread `펼쳐지는 문서` through Scenes 1–3 and `소매 안쪽` through Scenes 2–3 without thematic explanation. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |
