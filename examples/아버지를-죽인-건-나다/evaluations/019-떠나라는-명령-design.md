# Design Evaluation: Episode 019 — 떠나라는 명령

> Target: `episodes/019-떠나라는-명령.md`
> Target Reader: 회귀·빙의·환생 무협과 문파 장악물, 가족 관계의 반전, 복수형 사이다를 선호하는 성인 남성향 웹소설 독자
> Evaluation basis: Stage ④ design, `overview.md`, full `series.md`, architecture indexes and cited profiles, `continuity/story-so-far.md`, `continuity/018-완성품-summary.md`

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 criterion; this late P1 episode is not its scope. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Early-arc placement criterion; Episode 019 instead deepens the already established father conflict. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 2 executes a concrete inner-hall control negotiation; Scene 3 executes father opposition and leaves the second threat-letter hook. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | Dohyun orders departure and seizes the sword, but his motive stays behaviorally ambiguous; no protective verdict is stated. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series/late-arc payoff criterion; this episode plants the second threat-letter without claiming the later payoff. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | Jinwoo uses only released evidence and current custody leverage; he does not predict Dohyun’s response or the letter’s contents. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Endgame criterion; this episode advances but does not resolve the series conflict. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript-stage criterion; design forecast is checked below under Schema, not marked as prose quality. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only | ✅ | One canonical episode file with three complete Scene sections and required field order. |
| No skill/workflow dump after the design | ✅ | Body contains episode design and concise gate evidence only; no copied workflow procedure. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scenes 1–3 have unique titles and distinct turns: negotiate, condition, disarm. |
| Canonical episode path | ✅ | Actual disk path: `episodes/019-떠나라는-명령.md`. |
| Field notation | ✅ | Meta fields use `**Field:**`; execution fields use `- **Field:**`. |
| Every scene has required meta + bullet fields | ✅ | POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, outline, Unit budget, and exactly one Est. appear in all scenes. |
| Characters Appearing ↔ On stage union | ✅ | Appearing is 서진우·서도현·남궁혁·백무진, exactly the union of Scene 1’s four-person roster and Scenes 2–3’s two-person rosters. |
| On stage includes speakers | ✅ | Scene 1’s named supporting actions (혁·무진) are on stage; all dialogue agents are rostered and no ghost cast appears. |
| Characters ⊆ `characters.md` | ✅ | `characters.md` rows and `characters/서진우.md`, `characters/서도현.md`, `characters/남궁혁.md`, `characters/백무진.md` were read this turn; all paths read OK. |
| Summary/Hooks cast alignment | ✅ | Summary, hooks, seeds, closing, and scene rosters use only the four catalogued appearing characters. |
| No later-list cast debut | ✅ | No new cast; only established Episode 018 participants are used. |
| Locations ⊆ `locations.md` Key Locations | ✅ | All scenes cite `북문서가-본가`, which is a Key Locations row with slug `북문서가-본가`. |
| Location facets ⊆ Multi-facet anchors | ✅ | `locations/북문서가-본가.md` read OK; anchors explicitly include `가주전-회랑 접속부`, additive `내당 문앞`, and `가주전 문앞`. |
| Nested scene files absent | ✅ | Single canonical episode file; no nested scene directory. |
| No template residue | ✅ | No raw instruction placeholders remain. |
| Prose forecast present | ✅ | Each scene has five allowed typed categories and integer products: 2,490 / 2,490 / 2,350. |
| Forecast ↔ Est. cross-check | ✅ | Independent recomputation: Sc1 5×250+2×180+2×120+2×140+80=2,490; Sc2 same=2,490; Sc3 4×250+3×180+2×120+3×140+80=2,350. Est. 2,400/2,400/2,300 is within ±20% and outline density. |
| Dialogue intent vs outline speech | ✅ | Every outline’s planned spoken exchange is represented in Dialogue intent; none is marked `none`. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: 2,400 + 2,400 + 2,300 = 7,100; header addends: 2,400 + 2,400 + 2,300 = 7,100. |
| Est. length sum ≥ Scale min | ✅ | 7,100 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 7,100 ≤ 8,000; central band target is met. |
| Cited staging/profile paths exist | ✅ | `characters/서진우.md`, `characters/서도현.md`, `locations/북문서가-본가.md`, and `world/혈맥계약과-약그릇.md` were read OK this turn; Staging is `none`, so no staging file is required. |
| Episode List plot | ✅ | `series.md` Summary says Jinwoo rejects departure by proposing inner-hall control conditions; Hook says Dohyun accepts and takes the sword. Scenes 1–3 execute both clauses. |
| Hook evidence strength (internal) | ✅ | Series Hook “칼집 안쪽에 흑풍루의 두 번째 협박문이 숨겨져 있다” matches Summary, Out, Seed, Scene 3 Turn, and Transition: physical hidden letter is sensed, not read. |
| Hook scope (no Out creep) | ✅ | The closing adds no chase, faction arrival, or second independent reveal; only the hidden letter is carried forward. |
| No design-paste / meta-only scenes | ✅ | Each scene changes control: departure order → inner-hall condition → sword seizure; no scene is meta-only. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | Physical standoff becomes evidence negotiation, then authority exchange, then immediate cost. |
| Scene transitions chain | ✅ | Scene 1 sends the parties to the inner-hall threshold; Scene 2’s accepted authority leads to the main-house threshold; Scene 3 closes on the seized sword and sensed paper. |
| Scene sections complete | ✅ | All three Scene Index rows have full Key Events and forecasts. |
| Generation Readiness | ✅ | All Schema rows pass; no Pending adjudication item remains. |
| Beat concreteness | ✅ | Every Beat names observable acts: seal, ledger, key, authority token, sword seizure. |
| Est. length budget | ✅ | Recomputed 7,100 equals header and stays in 4,000–8,000. |
| Prose forecast quality | ✅ | Unit types map to dialogue, physical control, sensory anchors, POV deductions, and transitions. |
| Episode List scope aligned | ✅ | No later Episode 020 content is resolved; the second letter is only planted. |
| Prior hook addressed | ✅ | Scene 1 directly resumes the prior blade-between-hand-and-corpse confrontation. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A, Phase B, continuity, facet preflight, and additive Extend are recorded. |
| Series / overview tone & arc honored | ✅ | Dry event-first control negotiation and ambiguous father conflict match P1. |
| Episode List Summary / Hook honored | ✅ | Summary clauses and Hook are quoted in the design Gate Evidence and mapped to Scenes 1–3. |
| Hook internal consistency | ✅ | Body surfaces consistently show one hidden second threat letter, with no contents revealed. |
| Characters from architecture; profiles not redefined | ✅ | Character profiles are cited for behavior and voice; no core drive or identity changes. |
| Profile-backed knowledge / recognition | ✅ | Jinwoo infers from Dohyun’s observed hands, seal, and prior actions; no unsupported recognition claim is made. |
| Locations from architecture; profiles not redefined | ✅ | One Key Location is used with exact profile path. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` read OK this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | All three exact anchors are present; `내당 문앞` was additively extended before design. |
| Stagings from episode design; blocking not redefined | ✅ | All scenes are separate situations with `Staging: none`; no fixed multi-scene blocking is smuggled in. |
| World rules / history consistent with bible | ✅ | Authority and evidence custody are social/legal actions; contract mechanics remain bounded. |
| No improvised entities or silent lore | ✅ | No new named cast, faction, world rule, or uncatalogued place is introduced. |
| Continuity files used | ✅ | Immediate prior summary and story-so-far are cited and reflected. |
| Character/location state vs `story-so-far` | ✅ | Jinwoo, Dohyun, and the main-house corridor state continue from Episode 018 without reset. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-027/028 pick up; TH-025/026/006 advance; second letter plants; command chain and contract hold. |
| No contradiction of released continuity | ✅ | Episodes 001–018 remain locked; Episode 018’s G8 continuity is used as authority. |
| Conflicts section empty or escalated | ✅ | Only the necessary additive `내당 문앞` facet is documented; no unresolved conflict is ignored. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — selective Phase A → Phase B load complete.
- Locations: ✅ — all scenes map to `북문서가-본가`; exact anchors are `가주전-회랑 접속부`, `내당 문앞`, `가주전 문앞`.
- Length / Prose forecast: ✅ — written/recomputed products are 2,490/2,490, 2,490/2,490, and 2,350/2,350; scene/header sum is 7,100.
- Episode List Summary: ✅ — inner-hall control condition → Scenes 1–2; Dohyun accepts and takes sword → Scene 3.
- Hook to Next / Closing: ✅ — series Hook「칼집 안쪽에 흑풍루의 두 번째 협박문이 숨겨져 있다」; Out「칼집 안쪽에 흑풍루의 두 번째 협박문이 숨겨져 있다」; Scene 3 Turn/Transition senses but does not open it.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | “Why is Dohyun sending Jinwoo out, and how can Jinwoo reverse it while keeping evidence?” |
| Personal stake present | ✅ | Leaving means losing the childhood-surveillance evidence and custody of the corpse. |
| Episode Out hook | ✅ | Hidden second threat letter is concrete, physical, and directly connected to the seized sword. |
| Exposition budget respected | ✅ | Authority appears through objects and actions; no contract lecture. |
| Seed discipline | ✅ | One plot Plant (inner-hall control) and one Out Plant (second letter), with contents held. |
| Scene-first Key Events | ✅ | All three scenes contain causal action and turns, not summaries only. |
| Sensory-emotional on every scene | ✅ | Light/hands, threshold/medicine smell, and paper resistance each drive POV interpretation. |
| Motifs planned across scenes | ✅ | Hand/authority in Scenes 1–2; seized sword in Scene 3. |
| Overview signature line | N/A | `overview.md` has no required signature dialogue line. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | New information is limited to usable authority and a physical letter hint under confrontation. |
| Sensory-emotional pairing | ✅ | Light, threshold smell, and paper resistance each alter Jinwoo’s reading. |
| Dialogue voices + Dialogue intent | ✅ | Jinwoo is clipped and conditional; Dohyun is low, slow, and withholding. |
| Reader-discovered meaning | ✅ | The design asks the reader to weigh protection versus control; it forbids a protective verdict. |
| Antagonist plausibility | ✅ | Dohyun’s departure order and sword seizure are strategically coherent without exposition of motive. |
| Closing image specified | ✅ | Dohyun holds Jinwoo’s sword while a thin sealed paper is sensed in the empty scabbard. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | The episode resumes on the blade-at-the-corpse standoff rather than recapping. |
| Personal stake matches what this reader came for | ✅ | Tactical control of the sect and father conflict are fused in a concrete bargain. |
| Pacing / density fits platform expectations | ✅ | Three escalating scenes, compact evidence negotiation, and 7,100-character forecast fit serial web-novel density. |
| Out hook makes this reader want the next episode | ✅ | A hidden second Black Wind Tower threat letter inside the protagonist’s stolen sword is a direct mystery and danger. |
| No alienation of core audience without overview intent | ✅ | No romance pivot, training filler, or abstract lore lecture is introduced. |

## Design Critique (required personas)

#### Target Reader
- **Stance:** Adult male regression-martial-arts reader seeking tactical reversal, sect power, revenge pressure, and father mystery.
- **Strengths:** Turns the departure order into a power bargain, gives Jinwoo an operational win, then charges him an immediate cost when the sword is taken.
- **Defects:** —
- **Reader impact:** The reader receives control-progression and a sharper physical mystery without resolving Dohyun too early.

#### Genre Critic
- **Stance:** Tests whether the martial-arts web-novel promise produces concrete advantage rather than only family dialogue.
- **Strengths:** The “win” is custody of ledgers, keys, and seals; the sword seizure preserves danger and prevents effortless domination.
- **Defects:** —
- **Reader impact:** Sect-management satisfaction and father-versus-son tension remain in the same action contract.

#### Plot Expert
- **Stance:** Checks causality, Hook alignment, Hook scope, and whether each turn earns the next.
- **Strengths:** Scene 1 converts the prior cliffhanger into negotiation; Scene 2 earns the authority token; Scene 3 makes the accepted condition costly and carries exactly the scheduled threat-letter hook.
- **Defects:** — — body Summary/Out/Turn/Seed surfaces match the series Hook at one strength; no Out creep.
- **Reader impact:** The next episode has a clean question: who sent the letter and why was it hidden in Jinwoo’s sword?

#### Reader-Editor
- **Stance:** Checks mobile rhythm, information load, scene separation, and serial sellability.
- **Strengths:** Each scene has a distinct object and action pivot: corpse/record, ledger/key, sword/paper.
- **Defects:** —
- **Reader impact:** The forecast supports three readable escalation blocks without a procedural stall or crowded closing.

#### Literary Critic
- **Stance:** Checks motif, image, restraint, and whether the ambiguity is dramatized rather than explained.
- **Strengths:** Hands and thresholds turn authority into physical grammar; the empty scabbard lets the letter remain a sensed absence.
- **Defects:** —
- **Reader impact:** The family ambiguity gains pressure without a thematic coda or exoneration.

#### Character Critic
- **Stance:** Checks motivation, profile-backed knowledge, voice, and relationship pressure.
- **Strengths:** Jinwoo’s conditions arise from his established evidence-first behavior; Dohyun’s low commands, scarred finger, and sword control match his profile.
- **Defects:** — — no unsupported recognition or identity knowledge is assigned.
- **Reader impact:** Both characters act now for visible reasons, while Dohyun’s true motive remains a productive mystery.

#### Setting/Lore Expert
- **Stance:** Checks citeable location facets, institutional logic, and contract/trace restraint.
- **Strengths:** `내당 문앞` is additively promoted to a citeable anchor before use; evidence custody is social procedure, not a new supernatural rule.
- **Defects:** —
- **Reader impact:** The sect feels spatially controlled and legible; the hidden letter remains an earned object mystery.

## Design Adjudication
| # | Finding | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| — | No High or Med design defect. Schema, continuity, location facet, forecast arithmetic, Hook scope, required personas, and target-reader checks pass. | — | No | No | The locked reader gets a concrete power reversal, an earned cost, and a single actionable next hook. | — | Skip |

**Adjudication result:** No Pending or Carry-⑥ items remain.

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5 (Design Eval Approval)
- **Status:** ✅ Architect-approved
- **Evidence:** Full schema/path/facet checks, independent forecast recomputation, Hook body quotes, continuity compliance, required personas, and target-reader checks pass.
- **Next:** Stage ⑥ manuscript generation.
