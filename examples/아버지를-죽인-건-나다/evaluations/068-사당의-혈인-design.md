# Design Evaluation: Episode 068 — 사당의 혈인

## Evaluation Scope
- Target artifact: `episodes/068-사당의-혈인.md`
- Stage: ⑤ — design evaluation after Architect G4
- Target Reader: 성인 남성향 회귀·빙의·환생 무협과 문파 장악물, 가족 관계 반전과 복수형 사이다를 선호하는 독자
- Authority set: `overview.md`, `series.md`, `world-bible.md`, `world/혈맥계약과-약그릇.md`, `continuity/067-피로-쓴-계약-summary.md`, `continuity/story-so-far.md`

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001-only criterion; this design does not rewrite Episode 001. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 001–003 criterion, outside this episode’s design scope. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 보관단 개방·의식실 진입·봉인선 대조·신체 반응 검증·도현의 자발적 파기 시도가 Scenes 1–4의 실질적 사건이며, 도현의 자발적 봉인 파기 시도가 Out Hook이다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1-only criterion; Episode 068 is in P2. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series/late-arc payoff criterion; this episode advances TH-107 but does not claim the full later recovery. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | Jinwoo uses the current seal, recorded pressure marks, and bodily reactions; no future knowledge solves the shrine mechanism. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | ✅ | The episode makes Dohyun’s protection and control collide through a present choice while holding the endgame judgment and reconciliation question. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Published prose quality is Stage ⑥/⑦; forecast arithmetic is evaluated below under Schema. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | One canonical episode file contains four unique `### Scene` sections with required meta lines and flat bullets. |
| No skill/workflow dump after the design | ✅ | The file contains episode design, gate evidence, and G4 only; no pasted workflow procedure. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scenes progress from opening the storage, to physical proof of long-term sealing, to alternating bodily danger, to Dohyun’s attempted break; no identical Beat or Turn. |
| Canonical episode path | ✅ | Actual path is `episodes/068-사당의-혈인.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | Required bold-colon notation is used for all meta and bullet fields. |
| Every scene has required meta + bullet fields | ✅ | All four scenes include POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, 7-line outline, typed Unit budget, and exactly one Est. length. |
| Characters Appearing ↔ On stage union | ✅ | Appearing union is exactly 서진우·서도현·남궁혁·봉인된 또 다른 아이; all four are On stage in every scene. |
| On stage includes speakers | ✅ | All named speakers and intentional acts in Beat/Dialogue intent/outline are among the four On-stage characters; no anonymous voice or crowd is introduced. |
| Characters ⊆ `characters.md` | ✅ | All four map to the Character Catalog and the four exact profile paths were read this turn. |
| Summary/Hooks cast alignment | ✅ | Summary, In/Out, Seeds, Closing, and scene fields name only the four Appearing characters. |
| No later-list cast debut | ✅ | All four are established before Episode 068; no later Episode List character is introduced. |
| Locations ⊆ `locations.md` Key Locations | ✅ | `북문서가-사당` maps to the Key Locations row. |
| Location facets ⊆ Multi-facet anchors | ✅ | `사당 외실`, `혈인 보관단`, and `사당 아래 숨은 의식실` exactly occur in `locations/북문서가-사당.md` under Multi-facet anchors; the additive facet was approved before scene drafting. |
| Nested `episodes/{slug}/` scene files absent | ✅ | One flat canonical episode file only. |
| No template residue | ✅ | No raw `{placeholder}` or instructional brace remains. |
| Prose forecast present | ✅ | Every scene has 7 paragraph intents and five allowed unit types with integer `n×pick = subtotal` formulas. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 `3×240+2×180+2×120+2×140+1×80=1,680`; Sc2 `3×240+3×180+2×120+2×140+1×80=1,860`; Sc3 `3×250+3×180+2×120+2×140+1×80=1,890`; Sc4 `3×240+3×180+2×120+2×140+1×80=1,860`. Products recompute exactly; Est. values are within ±20% and outline-density bands. |
| Dialogue intent vs outline speech | ✅ | Each scene contains dialogue units and matching Dialogue intent; no scene declares `none` while outlining speech. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields: `1,700 + 1,800 + 1,900 + 1,800 = 7,200`; header addends: `1,700 + 1,800 + 1,900 + 1,800 = 7,200`. |
| Est. length sum ≥ Scale min | ✅ | Recomputed sum 7,200 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | Recomputed sum 7,200 ≤ 8,000; central target band is met. |
| Cited staging/profile paths exist | ✅ | This-turn `read_files` succeeded for `characters/서진우.md`, `characters/서도현.md`, `characters/남궁혁.md`, `characters/봉인된-또-다른-아이.md`, `locations/북문서가-사당.md`, `world/혈맥계약과-약그릇.md`, and `stagings/068-사당-의식실-개방.md`. |
| Episode List plot | ✅ | Series Summary 「사당 아래 숨은 의식실을 열고」→Scene 1; 「도현이 오랫동안 혈인을 봉인해 왔음」→Scene 2; 「봉인을 풀면 진우가 죽을 수 있지만, 그대로 두면 도현의 독이 발작」→Scene 3. |
| Hook evidence strength (internal) | ✅ | Body surfaces quote and preserve the same action: Series Hook 「도현이 스스로 봉인을 깨뜨리려 한다」; Summary says the choice is exposed; Out says 「도현이 스스로 봉인을 깨뜨리려 한다」; Scene 4 Turn says he begins the action without completing it. |
| Hook scope (no Out creep) | ✅ | Out contains one next-episode obligation—Dohyun’s self-breaking attempt; no faction arrival, chase, second reveal, or completed removal is added. |
| No design-paste / meta-only scenes | ✅ | Every scene performs a distinct physical evidence action and changes the available choice; none exists only to announce the next episode. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | Sequential escalation is storage lock → hidden room → long-term seal proof → alternating bodily cost → attempted self-break. |
| Scene transitions chain | ✅ | Scene 1’s revealed stone door leads to Scene 2; Scene 2’s seal-line endpoint leads to Scene 3’s minimum test; Scene 3’s unresolved Dohyun flare leads directly to Scene 4’s attempted self-break. |
| Scene sections complete | ✅ | All four Scene Index rows have complete generation briefs. |
| Generation Readiness | ✅ | All schema, path, cast, facet, length, and body Hook checks pass; no Pending adjudication remains. |
| Beat concreteness | ✅ | Beats name seal pressure, stone door, repeated finger-width marks, bodily reactions, and the attempted scratch of the central line. |
| Est. length budget | ✅ | Recomputed 7,200 equals header and scene fields; within 4,000–8,000 and central band. |
| Prose forecast quality | ✅ | Unit types correspond to opening dialogue, physical seal actions, sensory stone/line details, POV calculation, and transitions. |
| Episode List scope aligned | ✅ | The design executes the 068 Summary and preserves the 069 choice; complete removal is not stolen. |
| Prior hook addressed | ✅ | Scene 1 opens immediately on the closed blood-storage cabinet from Episode 067. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A indexes, four appearing profiles, shrine profile, world aspect, staging, story-so-far, and immediate 067 summary are listed. |
| Series / overview tone & arc honored | ✅ | The design remains cold, event-driven P2 investigation with revenge versus filial responsibility at its center. |
| Episode List Summary / Hook to Next honored | ✅ | Summary and Hook are quoted in the episode’s Gate Evidence at the same strength. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Arc, Out, Hint seed, and Scene 4 Turn all state Dohyun’s self-breaking attempt; completion is consistently held. |
| Characters from architecture; profiles not redefined | ✅ | Existing drives, voices, appearance states, and relationships are cited without core modification. |
| Profile-backed knowledge / recognition | ✅ | No unsupported identity revelation; the inference comes from physical pressure marks and Dohyun’s observable hand/vein response. |
| Locations from architecture; profiles not redefined | ✅ | Only the catalogued shrine is used; the hidden-room facet is an additive extension, not a silent invention. |
| Location profile paths readable | ✅ | Exact shrine profile and staging profile were read successfully this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | All three cited facets are exact labels in the successfully read shrine profile. |
| Stagings from episode design; blocking not redefined | ✅ | `068-사당-의식실-개방` is authored for Stage ④, cites existing character states, and keeps positions stable across all scenes. |
| World rules / history consistent with bible | ✅ | Blood as contract medium, violation as poison/vascular reaction, and the blood-seal linkage remain within the approved world aspect. |
| No improvised entities or silent lore | ✅ | No new faction, character, rule, or uncatalogued set is added; the only local gap was explicitly extended and approved. |
| Continuity files used (ep 002+) | ✅ | Immediate 067 summary and current story-so-far are the sole narrative authorities. |
| Character/location state vs `story-so-far` | ✅ | Jinwoo remains responsible and skeptical; Dohyun remains ill and withholding; Hyuk verifies; the child guards autonomy; the shrine’s storage was previously closed and is now opened as the next event. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-107 advances; Dohyun’s self-break is planted; TH-101/102 and complete removal remain held. |
| No contradiction of released continuity | ✅ | 067 only located the closed storage; 068 now opens it as the planned next unit and does not claim prior removal. |
| Conflicts section empty or escalated | ✅ | No unresolved source conflict remains after the additive facet extension. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — selective Phase A/B load completed; the new shrine facet and staging were approved and read.
- Locations (index): ✅ — `북문서가-사당` ∈ Key Locations.
- Locations (path): ✅ — `locations/북문서가-사당.md` and `stagings/068-사당-의식실-개방.md` read OK this turn; all four character and world paths also read OK.
- Locations (facets): ✅ — `사당 외실`, `혈인 보관단`, `사당 아래 숨은 의식실` ⊆ the shrine profile’s Multi-facet anchors.
- Length / Prose forecast: ✅ — Sc1 written=1,680; recomputed=1,680; Est=1,700 · Sc2 written=1,860; recomputed=1,860; Est=1,800 · Sc3 written=1,890; recomputed=1,890; Est=1,900 · Sc4 written=1,860; recomputed=1,860; Est=1,800 · scene fields and header addends both total 7,200.
- Episode List Summary: ✅ — 「숨은 의식실을 열고」→Scene 1; 「도현이 오랫동안 혈인을 봉인」→Scene 2; 「봉인을 풀면 진우가 죽을 수 있지만 그대로 두면 도현의 독이 발작」→Scene 3.
- Hook to Next / Closing: ✅ — Series Hook「도현이 스스로 봉인을 깨뜨리려 한다」; Episode Out「도현이 스스로 봉인을 깨뜨리려는 행동을 시작」; Scene 4 Turn「도현이 스스로 봉인을 깨뜨리려는 행동을 시작」(완전 파기는 보류).
- Hook scope: ✅ — one next-episode action, no extra obligation.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | The design asks whether the blood seal binds Jinwoo for protection or ownership and delays the moral answer. |
| Personal stake present | ✅ | The minimum test directly alternates Jinwoo’s death signs and Dohyun’s poison flare. |
| Episode Out hook | ✅ | Dohyun begins breaking the seal himself; the physical act is already in motion at the close. |
| Exposition budget respected | ✅ | Three rule fragments—opening order, repeated seal work, alternating reaction—are scene-bound; full contract and bloodline stay held. |
| Seed discipline | ✅ | Two Plants and one Hint are explicit; Holds are limited and actionable. |
| Scene-first Key Events (all required fields) | ✅ | Every scene has a concrete physical event, turn, transition, and generation forecast. |
| Sensory-emotional on every scene | ✅ | Wax/incense, wet stone/scratches, iron heat/low resonance, and red powder/two-way vibration each alter Jinwoo’s reading. |
| Motifs planned across scenes | ✅ | Overlapping red lines and first-moving hands recur with distinct scene functions. |
| Overview signature line | N/A | `overview.md` has no separate mandatory signature dialogue line for this episode. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | Each disclosure occurs through an access obstacle or bodily danger, not a lore lecture. |
| Sensory-emotional pairing | ✅ | Material details are paired with Jinwoo’s perception in every scene. |
| Dialogue voices + Dialogue intent | ✅ | Jinwoo presses and calculates; Dohyun warns and withholds; Hyuk classifies; the child blocks. |
| Reader-discovered meaning | ✅ | Protection and ownership remain simultaneous conclusions; no moral verdict is placed in the closing image. |
| Antagonist plausibility | ✅ | The unseen Black-Wind Tower remains an indirect contract structure; no villain exposition is added. |
| Closing image specified | ✅ | The two red lines, Dohyun’s hand, and Jinwoo’s grip create an image-led cut. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | The episode opens at the previously located blood seal and immediately makes access physical. |
| Personal stake matches what this reader came for | ✅ | The family-revenge reader receives a direct father-son survival conflict, not detached system exposition. |
| Pacing / density fits platform expectations | ✅ | Four distinct evidence turns fill a 7,200-character forecast without training filler or external detour. |
| Out hook makes this reader want the next episode | ✅ | Dohyun himself begins the dangerous act the reader has been waiting to see, while the result remains uncertain. |
| No alienation of core audience without overview intent | ✅ | No romance detour, reconciliation speech, or unexplained faction arrival is introduced. |

## Design Critique (required personas)
#### Target Reader
- Stance: Adult male-oriented regression-murim reader seeking bodily stakes, document/relic reversals, and father-son mistrust.
- Strengths: The episode pays the shrine-location promise by opening the actual storage and turns the blood seal into an immediate two-way survival dilemma.
- Defects: —
- Reader impact: The close gives a concrete “someone is acting now” pull without spending the 069 choice.

#### Genre Critic
- Stance: Tests the design against regression-murim, sect-control, and family-revenge serial contracts.
- Strengths: Jinwoo earns the reveal through present evidence; the hidden ritual room and bodily cost are genre-appropriate escalations.
- Defects: —
- Reader impact: The episode supplies a material breakthrough and preserves the next episode’s decisive confrontation.

#### Plot Expert
- Stance: Audits causality, Hook body alignment, and scope.
- Strengths: Closed cabinet → stone door → layered marks → alternating reaction → self-break attempt is causal; Summary, Out, and closing Turn preserve one Hook claim.
- Defects: —
- Reader impact: A cold reader can tell why the room opens, why Dohyun’s history is credible, and why the choice cannot resolve yet.

#### Reader-Editor
- Stance: Checks serial pacing, exposition restraint, and closing density.
- Strengths: Every scene changes access or bodily cost; the final Transition has one dominant obligation rather than a crowded cliffhanger.
- Defects: —
- Reader impact: Low skim risk; the procedural investigation is continuously converted into action.

#### Literary Critic
- Stance: Checks motif, sensory craft, Hold discipline, and image-led closure.
- Strengths: Repeated red lines turn protection into visible accumulated control; “first-moving hands” gives the father-son conflict a physical grammar.
- Defects: —
- Reader impact: The theme is discoverable through objects and hands without a didactic statement.

#### Character Critic
- Stance: Checks profile-backed motivation, voice, and relationship pressure.
- Strengths: Jinwoo tests rather than trusts; Dohyun’s long-term action is revealed through behavior; Hyuk and the child retain distinct functions and boundaries.
- Defects: —
- Reader impact: Each supporting character pressures the central choice without taking Jinwoo’s agency away.

#### Setting/Lore Expert
- Stance: Checks blood-contract rules, shrine geometry, and citeable facets.
- Strengths: The hidden room is explicitly added under Multi-facet anchors; the design uses only approved reaction rules and existing character states.
- Defects: —
- Reader impact: The reader can visualize the locked shrine-to-room route and understand the danger without receiving the whole ritual system.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| — | No High/Med design-field finding; all required schema, continuity, Hook, craft, and Target Reader checks pass. | — | No | no | The locked reader receives a physical breakthrough, a personal survival dilemma, and a single actionable cliffhanger. | — | Skip |
| 1 | Episode-level motif placement is carried through scene Sensory-emotional and Beat fields rather than separate `Motif touch` fields. (Literary Critic) | Low | No | yes | Separate metadata would not improve this reader’s clarity; Stage ⑥ needs a reminder to preserve the red-line/hand progression. | Carry the two motifs as a generation constraint; do not add redundant design fields. | Carry-⑥ |

**Adjudication:** No High or Med design-field finding remains. The sole Low item is permitted Carry-⑥ and does not block Generation-ready.

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Rationale:** Schema, independent arithmetic, exact path/facet checks, additive architecture, continuity, Hook alignment, target-reader checks, and all required persona critiques pass. The only Low item is a permitted Carry-⑥ motif constraint; no Pending revision remains.
- **Next:** Stage ⑥ — manuscript generation
