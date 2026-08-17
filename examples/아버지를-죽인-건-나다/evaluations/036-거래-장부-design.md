# Design Evaluation: Episode 036 — 거래 장부

## Evaluation Scope
- Evaluated artifact: `episodes/036-거래-장부.md`
- Target Reader: 회귀·빙의·환생 무협과 문파 장악물, 가족 관계의 반전, 복수형 사이다를 선호하는 성인 남성향 웹소설 독자.
- Evaluation gate: G5 design evaluation after Architect G4.
- Loaded authority: `overview.md`, `series.md`, architecture indexes and cited profiles, `stagings/036-거래-장부-압박.md`, `continuity/story-so-far.md`, `continuity/035-떨리는-칼-summary.md`.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 criterion; not a requirement for Episode 036. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Early-arc criterion already executed; not this episode’s design scope. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 2 executes a concrete ledger seizure through pressure and public verification; Scene 3 ends with the erased birth-record hook. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | Episode 036 is in P2, and its design still preserves this principle but the criterion is explicitly P1-scoped. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series/late-arc payoff criterion; this episode advances the medicine trail without claiming payoff. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우 uses a discovered classification code and current ledger evidence, not an unearned future prediction; changed-present uncertainty remains. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Series-ending criterion; no ending payoff is claimed here. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript Stage ⑥/⑦ criterion; design forecast is checked in Schema / Structural Integrity instead. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only | ✅ | `episodes/036-거래-장부.md` uses one canonical episode file and three complete flat-field scene sections. |
| No skill/workflow dump after the design | ✅ | The file contains design-specific gate and readiness records only, not pasted workflow instructions. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–3 have unique titles and distinct causal functions. |
| Canonical episode path | ✅ | Exact path is `episodes/036-거래-장부.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | Meta and bullet fields use the required bold-colon notation. |
| Every scene has required meta + bullet fields | ✅ | All three scenes contain POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, outline, Unit budget, and one Est. length. |
| Characters Appearing ↔ On stage union | ✅ | Union is 서진우, 남궁혁, 문상철; 서도현 is explicitly mention-only and not claimed on stage. |
| On stage includes speakers | ✅ | Scene 1 speakers are 진우·혁; Scenes 2–3 speakers are 진우·혁·문상철; each is on stage. |
| Characters ⊆ `characters.md` | ✅ | All four named characters map to catalog rows and readable profiles; no crowd label is used as cast. |
| Summary/Hooks cast alignment | ✅ | Summary, hooks, seeds, and closing refer only to appearing cast or mention-only 서도현. |
| No later-list cast debut | ✅ | No cast is imported from a later episode; 문상철 is established and has prior profile presence. |
| Locations ⊆ `locations.md` Key Locations | ✅ | `북문서가-본가` and `북항` are exact Key Locations. |
| Location facets ⊆ Multi-facet anchors | ✅ | `지하 보관실`, `창고 골목`, and `상단 객잔 마당` are exact anchors in the loaded profiles. |
| Nested episode scene files absent | ✅ | No nested episode directory or scene files. |
| No template residue | ✅ | No unfilled template braces remain. |
| Prose forecast present | ✅ | Each scene has five-type integer unit formula and 5–7 paragraph intents. |
| Forecast ↔ Est. cross-check | ✅ | Sc1: 2×260 + 3×190 + 2×130 + 3×150 + 1×90 = 1,890; Est 1,900. Sc2 = 2,190; Est 2,200. Sc3 = 2,280; Est 2,300. Every written product equals independent arithmetic and lies within outline-density bands. |
| Dialogue intent vs outline speech | ✅ | All scenes with speech have non-`none` dialogue intent, and every speaker is on stage. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields 1,900 + 2,200 + 2,300 = 6,400; header repeats 1,900 + 2,200 + 2,300 = 6,400. |
| Est. length sum ≥ Scale min | ✅ | 6,400 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 6,400 ≤ 8,000; central target band is satisfied. |
| Cited staging/profile paths exist | ✅ | This-turn reads succeeded for `locations/북문서가-본가.md`, `locations/북항.md`, all appearing character profiles, `world/혈맥계약과-약그릇.md`, and `stagings/036-거래-장부-압박.md`. |
| Episode List plot | ✅ | Series Summary’s pressure on the merchant, ledger seizure, monthly medicine records, and child-name month each map to concrete Scene 2–3 Beats. |
| Hook evidence strength (internal) | ✅ | Series Hook「그 달에 진우의 출생 기록도 함께 사라졌다」; Summary says the child-name month includes the erased record; Out says the same; Scene 3 Turn finds the removed birth record and closing image retains the stamp trace. |
| Hook scope | ✅ | Out contains one next-episode obligation, the erased birth record, with the child-name month as its supporting evidence; no chase or second reveal is added. |
| No design-paste / meta-only scenes | ✅ | Each scene has a dramatic event: code localization, ledger seizure, and record discovery. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | The chain is code → merchant pressure → ledger seizure → monthly record → erased birth record. |
| Scene transitions chain | ✅ | Scene 1 sends the pair to the harbor; Scene 2 carries the ledger to the inn yard; Scene 3 sends the seized record back to the main house. |
| Scene sections complete | ✅ | All three Scene Index rows have full generation briefs. |
| Generation Readiness | ✅ | No Schema / path / facet / cast / Hook / forecast failure blocks Stage ⑥. |
| Beat concreteness | ✅ | Every Beat names documents, seals, numbering, hand movement, seizure, or physical paper evidence. |
| Est. length budget | ✅ | Independent products and 6,400 sum pass the 4,000–8,000 episode band. |
| Prose forecast quality | ✅ | Unit counts correspond to the planned dialogue, physical pressure, sensory harbor details, POV restraint, and transitions. |
| Episode List scope aligned | ✅ | The design executes the approved Summary without adding an unscheduled chase or faction arrival. |
| Prior hook addressed | ✅ | The half-prescription and classification code from 035 directly initiate Scene 1. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A/B, staging, continuity, and episode-only scope are explicitly recorded. |
| Series / overview tone & arc honored | ✅ | Evidence-first pressure and family-revenge uncertainty fit P2 and the locked male-oriented serial tone. |
| Episode List Summary / Hook honored | ✅ | Summary clauses and Hook are quoted and mapped to scenes above. |
| Hook internal consistency | ✅ | Summary, Arc, Out, Seeds, Scene 3 Turn, and Transition all preserve the same erased-birth-record claim. |
| Characters from architecture; profiles not redefined | ✅ | Profiles are cited for drive, voice, appearance states, and relationships; no new state is invented. |
| Profile-backed knowledge / recognition | ✅ | 진우 recognizes document patterns from the present investigation; 혁 verifies seals publicly; no unsupported identity recognition is claimed. |
| Locations from architecture; profiles not redefined | ✅ | Both locations and all three facets are catalogued. |
| Location profile paths readable | ✅ | Exact profile paths were read successfully this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | All three scene facets match loaded anchor labels exactly. |
| Stagings from episode design; blocking not redefined | ✅ | New staging is authored at ④, cites base states, and matches Scenes 2–3 blocking and props. |
| World rules / history consistent with bible | ✅ | The design treats the prescription code as a traceable document key and keeps blood-contract causality unresolved. |
| No improvised entities or silent lore | ✅ | No new faction, rule, character, or uncatalogued place is introduced. |
| Continuity files used | ✅ | Immediate prior summary and story-so-far are reflected in the Prior Hook, state changes, and thread ledger. |
| Character/location state vs story-so-far | ✅ | 진우 remains evidence-seeking power-holder, 혁 public verifier, 도현 absent and concealed; north-harbor facets remain consistent. |
| Unresolved threads pick up / advance / plant / hold | ✅ | TH-056–059 are categorized under Threads; birth-record and wider network are planted, held causes remain held. |
| No contradiction of released continuity | ✅ | Nothing reverses 035’s discovery or reveals 도현’s inner state. |
| Conflicts section empty or escalated | ✅ | Conflicts are explicitly `None`; no catalog or continuity conflict was found. |

## Design Consistency Gate
- Loaded required artifacts: ✅
- Locations index / path / facets: ✅ — exact Key Location rows, readable profile paths, and anchor labels are separately evidenced in Schema above.
- Length / Prose forecast: ✅ — written products equal independent recomputation; scene fields and header both total 6,400.
- Episode List Summary: ✅ — merchant pressure and ledger seizure → Scene 2; 도현’s monthly medicine records and child-name month → Scene 3.
- Hook to Next / Closing: ✅ — series Hook「그 달에 진우의 출생 기록도 함께 사라졌다」; episode Out「아이들의 이름이 적힌 달에 진우의 출생 기록도 함께 사라졌다」; Scene 3 Turn finds the removed birth-record row; closing image retains the birth-year stamp trace.
- Generation Readiness: ✅ — all structural checks pass and no Pending design-field adjudication remains.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | The episode asks whether the code points to treatment or a hidden crime, and Scene 1 delays the answer through document work. |
| Personal stake present | ✅ | The ledger may explain both 도현’s illness and 진우’s own altered birth record. |
| Episode Out hook | ✅ | The birth-record disappearance is a concrete, singular next question for the target reader. |
| Exposition budget respected | ✅ | Three new information layers are dramatized through objects and pressure, with complete prescription and child history held back. |
| Seed discipline | ✅ | Code and monthly shipments are Plants; the erased birth record is a controlled Hint/Hook; full causes remain Hold. |
| Scene-first Key Events | ✅ | All information is assigned to observable scene actions rather than episode-level exposition. |
| Sensory-emotional on every scene | ✅ | Cold medicine storage, wet harbor cargo, and torn ledger fibers each produce a POV reaction. |
| Motifs planned across scenes | ✅ | Hands/seals and wet paper are placed in distinct scenes and support responsibility/traces. |
| Overview signature line | N/A | `overview.md` has no required signature dialogue line for this episode. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | Document interpretation is paired with a timed seizure, physical blocking, and the final personal discovery. |
| Sensory-emotional pairing | ✅ | Each sensory cue changes how 진우 reads the evidence or controls his reaction. |
| Dialogue voices + Dialogue intent | ✅ | 진우 is clipped and imperative, 혁 is clear and public-minded, 문상철 is deferential and evasive. |
| Reader-discovered meaning | ✅ | The reader infers that medicine logistics and family records share a hidden system; no thematic conclusion is stated as dialogue. |
| Antagonist plausibility | ✅ | 문상철 protects his commercial survival and uses regulation rather than cartoon malice. |
| Closing image specified | ✅ | Torn paper edge and surviving birth-year stamp create a concrete, forward-facing close. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | The prior episode’s tremor immediately becomes a coded evidence chase, avoiding a reset. |
| Personal stake matches what this reader came for | ✅ | The ledger implicates the father while threatening the protagonist’s own birth history. |
| Pacing / density fits platform expectations | ✅ | Three escalating action-information scenes forecast 6,400 characters with no training or lecture block. |
| Out hook makes this reader want the next episode | ✅ | A ledger has erased the protagonist’s birth record and lists unknown children, creating both revenge and family mystery momentum. |
| No alienation of core audience without overview intent | ✅ | The episode remains a cold, tactical investigation with a clear seizure payoff and restrained emotional leak. |

## Design Critique (required personas)
#### Target Reader
- Stance: Adult male web-novel reader seeking tactical reversal, family suspicion, and a strong next-episode question.
- Strengths: The code gives immediate procedural momentum; the ledger seizure supplies a tangible win; the erased birth record personalizes the final hook.
- Defects: —
- Reader impact: The episode should retain the reader because the victory produces a larger, more personal mystery rather than a reset.

#### Genre Critic
- Stance: Tests the episode against regression martial-arts and revenge-serial promises.
- Strengths: Future knowledge is converted into present evidence and leverage; the merchant pressure delivers a clean tactical payoff.
- Defects: —
- Reader impact: The genre reader receives a concrete “information advantage → decisive action” beat without making 진우 omniscient.

#### Plot Expert
- Stance: Audits causality, escalation, Hook body alignment, and Out scope.
- Strengths: The 035 code causally identifies the ledger route; Scene 2’s seizure enables Scene 3’s discovery; all Hook surfaces state the same erased-birth-record obligation.
- Defects: —
- Reader impact: The next question is earned by an observable paper trace, not by an arbitrary cliffhanger.

#### Reader-Editor
- Stance: Audits serialization, exposition restraint, and closing density.
- Strengths: Three scenes each have a distinct job; the last scene has one primary hook with one supporting child-name reveal; dialogue intent prevents exposition dumping.
- Defects: —
- Reader impact: The close is legible at a fast mobile read and offers a reason to open Episode 037.

#### Literary Critic
- Stance: Tests motif, image, and emotional restraint at design scope.
- Strengths: Hands/seals and wet paper turn responsibility and erasure into recurring physical objects; the closing image avoids thematic explanation.
- Defects: —
- Reader impact: The family mystery gains tactile memory rather than existing only as plot information.

#### Character Critic
- Stance: Tests motivation, voice, and profile-backed knowledge.
- Strengths: 진우’s choice not to force 도현’s hand becomes action logic; 혁’s public-verification drive constrains him; 문상철’s evasions match his catalogued commercial survival motive.
- Defects: —
- Reader impact: The trio’s competing methods make the pressure scene feel interpersonal, not merely procedural.

#### Setting/Lore Expert
- Stance: Tests document logic, location anchors, and blood-contract restraint.
- Strengths: The code is used as a trade classification rather than unexplained magic; every harbor set is an exact anchor; the design does not claim that the ledger proves the blood-contract cause.
- Defects: —
- Reader impact: The world remains intelligible through objects and movement while preserving mystery.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| — | No High or Med design defects. Required personas independently found the design generation-ready; Low preferences do not justify altering the locked causal chain. | Low | No | no | Target Reader receives a concrete seizure payoff and a singular personal hook; extra changes would risk softening the 035→036 escalation. | — | Skip |

**Adjudication Status:** No Pending design-field findings. Carry-⑥ is not required; all episode-level motif placements already have scene-level cues.

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Approved:** 2025-02-14 — Schema, continuity, hook alignment, exact forecast arithmetic, required persona critique, and target-reader checks all pass; no Pending Apply row remains.
- **Next:** Stage ⑥ — manuscript generation
