# Design Evaluation: Episode 013 — 장로의 표

## Evaluation Scope & Gate State
- **Evaluated design:** `episodes/013-장로의-표.md` — this is the canonical disk path and was read in full this turn.
- **Stage:** ⑤ Design Evaluation.
- **Gate G4:** Approved by user before this evaluation; the design remains unchanged during Stage ⑤.
- **Gate G5:** Pending explicit user approval of this evaluation and its adjudication.
- **Target Reader:** 성인 남성향 웹소설 독자 — 회귀·빙의·환생 무협, 문파 장악물, 가족 반전, 복수형 사이다를 선호하는 독자.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 013 design-evaluation scope. This is an episode-001 criterion; prior episode continuity confirms those elements were established, but Stage ⑤ does not re-grade the released opening. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 013 design-evaluation scope. The criterion belongs to the opening arc and is already outside this episode's design scope. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | This episode executes both 추적 and 대립: Scene 1 converts the red document box into a procedural trap; Scene 2 exposes the custodian; Scene 3 reveals the coup's first order and confirms the ten-day-early start. Episode Out and the series Hook to Next both state that the coup date is ten days earlier than Jinwoo remembers. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | The design keeps the coup order and altered birth record as evidence against Dohyun, but explicitly Holds his actual intent and possible protective motive. Scene 3 has Jinwoo refuse to declare Dohyun innocent and only fixes the date discrepancy. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | This is a later-arc payoff criterion. Episode 013 may hold the medicine, letter, Gahwan, and threat-condition threads; planting or holding them is not the same as satisfying the eventual payoff. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | Jinwoo uses remembered procedure and the previous coup timing to design a trap, but the changed command date defeats simple prediction. The Out explicitly turns the mismatch into a new problem rather than treating memory as prophecy. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Series-/ending-level criterion. Episode 013 preserves the unresolved father question but does not and should not deliver the ending confrontation here. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript quality and final prose density belong to Stages ⑥–⑦. Forecast arithmetic is checked below under Schema, not used as a Criteria failure. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | The design uses one episode header, episode-level fields, `## Scenes`, and unique `### Scene 1–3` blocks with the required scene fields. |
| No skill/workflow dump after the design | ✅ | The body contains story design and a short Gate G4 block only; no copied workflow procedure or template instructions appear. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Exactly three unique headings: `### Scene 1 — 봉인 규칙의 미끼`, `### Scene 2 — 표결단의 손`, `### Scene 3 — 열흘 앞선 명령`. Their situations, turns, and functions differ. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | Exact disk path read: `episodes/013-장로의-표.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | Scene metadata and bullet fields consistently use bold field labels with colons. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | All scenes include POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, and exactly one Est. length. Scene 1 correctly uses `Staging: none` because it is a single-scene setup situation. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing set is 서진우, 장로-대표, 윤태석, 백무진. Scene 1 has 진우·백무진; Scenes 2–3 have all four. No named Appearing character is absent from the union and no scene introduces a ghost. |
| On stage includes speakers | ✅ | Scene 1 dialogue is 진우/백무진; Scenes 2–3 dialogue intentions assign speech to 진우, 장로 대표, 윤태석, 백무진, all present on stage in those scenes. |
| Characters ⊆ `characters.md` | ✅ | `characters.md` catalog maps all four names to profiles: `characters/서진우.md`, `characters/장로-대표.md`, `characters/윤태석.md`, `characters/백무진.md`; all four profile paths were read successfully this turn. |
| Summary/Hooks cast alignment | ✅ | Summary names 진우 and 윤태석 as the action center and the Characters line includes all four appearing characters. The Out names no new person; the closing action remains Jinwoo's. |
| No later-list cast debut | ✅ | All four characters are catalogued before Episode 013; the catalog gives 장로 대표 and 윤태석·백무진 series presence from earlier episodes. |
| Locations ⊆ `locations.md` Key Locations | ✅ | `locations.md` maps 북문서가 본가 / slug `북문서가-본가` to `locations/북문서가-본가.md`. Both cited facets are sub-facets of that catalog location. |
| Location facets ⊆ Multi-facet anchors | ✅ | Successful profile read at `locations/북문서가-본가.md` lists exact anchors `가주전-회랑 접속부` and `장로회당 표결단`; Scene 1 and Scenes 2–3 cite those exact labels. Adjacent Scenes 2–3 share the facet but have visibly distinct functions and turns: custodian identification, then original-order/date reveal. |
| Nested `episodes/{slug}/` scene files absent | ✅ | The episode uses inline scene sections; no nested scene-file layout is present or required. |
| No template residue | ✅ | No raw placeholder braces or unfinished template markers appear. |
| Prose forecast present (outline + typed units) | ✅ | Every scene has numbered Paragraph outline and typed units: dialogue, action, sensory, POV, transition. Each Unit budget has integer counts and picks. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Recomputed independently: Scene 1 `3×250 + 3×180 + 2×120 + 3×140 + 1×80 = 2,030`; Scene 2 `4×250 + 3×180 + 2×120 + 3×140 + 1×80 = 2,280`; Scene 3 has the same `2,280`. Est. values 2,000 / 2,300 / 2,300 are within the outline-density bands and within approximately 20% of their correct subtotals. |
| Dialogue intent vs outline speech | ✅ | Each scene has dialogue units and dialogue-bearing beats; no scene claims `none` while outlining speech. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: `2,000 + 2,300 + 2,300 = 6,600`\nheader addends: `2,000 + 2,300 + 2,300 = 6,600` |
| Est. length sum ≥ Scale min (hard) | ✅ | Recomputed scene-field sum is 6,600; overview Scale minimum is 4,000. |
| Est. length sum ≤ Scale max (hard) | ✅ | Recomputed scene-field sum is 6,600; overview Scale maximum is 8,000. It sits within the preferred central band of approximately 5,000–7,200. |
| Cited staging/profile paths exist | ✅ | This-turn `read_files` succeeded for `characters/서진우.md`, `characters/장로-대표.md`, `characters/윤태석.md`, `characters/백무진.md`, `locations/북문서가-본가.md`, and `stagings/013-붉은-문서함-표결-압박.md`. The staging file explicitly covers Scenes 2–3. |
| Episode List plot (not a different story) | ✅ | `series.md` Episode 013 Summary says Jinwoo makes elders suspect one another and the custodian reveals the first coup order. Scene 2 Beat/Turn performs the suspicion and custodian exposure; Scene 3 Beat performs the original order reveal. |
| Hook evidence strength (internal) | ✅ | Body surfaces agree at the same observable strength. Series Hook: “쿠데타의 개시일이 진우의 기억보다 열흘 앞당겨진다.” Episode Out: same ten-day discrepancy. Episode Arc close: date discrepancy is confirmed and forces the next move. Seed Hint: the order date differs from memory. Scene 3 Turn: “명령서의 쿠데타 개시일이 진우의 기억보다 열흘 앞당겨져 있음을 모두가 확인한다.” Closing Transition turns that confirmed fact into the decision to seize the armory first. |
| Hook scope (no Out creep) | ✅ | The Out has one primary obligation—ten-day-early coup date—and one directly consequent action—Jinwoo's next response. It does not introduce an unrelated chase, faction arrival, or second reveal. |
| No design-paste / meta-only scenes | ✅ | Each scene contains a distinct dramatic event: trap design, public identification, and document/date revelation. No repeated outline loop or copied Unit budget appears. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | The causal chain is clear: prior red-box hook → procedural trap → public pressure → keeper identification → voluntary opening → coup order/date reveal → next tactical response. |
| Scene transitions chain | ✅ | Scene 1 opens the hall and moves the box to the voting platform; Scene 2 ends with Taeseok choosing to open it; Scene 3 begins at the unfastened seal and ends with Jinwoo acting on the date. `When` is continuous same-morning time. |
| Scene sections complete | ✅ | Every Scene Index row has a matching full scene block with all required Key Event fields. |
| Generation Readiness | ✅ | All Schema rows pass, including independently recomputed length, cast, paths, facets, staging, and body-level Hook alignment. Adjudication contains no Pending design-field fix. |
| Beat concreteness | ✅ | Beats specify actions, evidence comparisons, procedural pressure, document handling, and observable decisions rather than mood alone. |
| Est. length budget | ✅ | Recomputed total 6,600 equals the header and stays within 4,000–8,000. Typed-unit subtotals and outline density are compatible. |
| Prose forecast quality | ✅ | Dialogue, action, sensory, POV, and transition units correspond to the listed beats and dialogue intentions. |
| Episode List scope aligned | ✅ | Summary and Hook to Next are executed without adding a separate plot obligation. |
| Prior hook addressed (ep 002+) | ✅ | Scene 1 immediately handles the Episode 012 red document box, seal cord, and internal archive trail. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | The design records Phase A indexes, appearing/used profiles, staging details, and the Episode 012 continuity pair; Prior Design Alignment explicitly expands the prior red-box hook into a procedural confrontation. |
| Series / overview tone & arc honored | ✅ | The design is cold, event-driven, procedural, and keeps the father question unresolved while advancing P1's internal power struggle. |
| Episode List Summary / Hook to Next honored | ✅ | Summary and Hook are quoted and executed in the Schema rows above; the date discrepancy is not weakened in the body. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Episode Arc close, Out, Seed, Scene 3 Turn, and Transition all preserve the same ten-day-early claim. |
| Characters from architecture; profiles not redefined | ✅ | Profile-backed identities and drives are used: Jinwoo reads hands/seals, Taeseok protects access to his brother's record, Mu-jin calculates exit safety, and the elder uses formal ledger language. No profile is rewritten. |
| Profile-backed knowledge / recognition | ✅ | Jinwoo's recognition of access marks is supported by his role and prior document-tracing arc; Taeseok's access motive is supported by his profile's brother-record drive. The design does not claim that Taeseok knows the command's ultimate author. |
| Locations from architecture; profiles not redefined | ✅ | `북문서가-본가` is a Key Location and the cited facets are exact profile anchors. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` was read successfully this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | Exact matching anchors were read from the location profile, as cited in Schema. |
| Stagings from episode design; blocking not redefined | ✅ | `stagings/013-붉은-문서함-표결-압박.md` was authored in Stage ④ and read this turn. It locks the Scene 2–3 cast states, positions, props, and continuity. Scene 2–3 do not silently swap blocking. |
| World rules / history consistent with bible | ✅ | The episode uses document, seal, ledger, and elder-council authority already defined in the world bible; it introduces no new supernatural rule or faction. |
| No improvised entities or silent lore | ✅ | No new named entity, place, power, or rule is required. The command's ultimate author and Dohyun's intent remain explicitly held rather than invented. |
| Continuity files used (ep 002+) | ✅ | `continuity/story-so-far.md` and `continuity/012-불이-난-날-summary.md` were loaded; the prior red-box hook, TH-015/016/017, and character states are reflected. |
| Character/location state vs `story-so-far` | ✅ | Jinwoo remains the pursuing successor; the elder remains the institutional pressure point; Taeseok and Mu-jin retain their open motives and conditional cooperation. The staging adds only this episode's situation state. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-017 advances from unknown custodian to Taeseok's present custodial role; TH-015 advances through the coup order; TH-014, threat-letter origin, Gahwan, and Dohyun's intent are explicitly held. |
| No contradiction of released continuity | ✅ | The episode continues from Episode 012's same-morning red-box discovery and does not alter any released event. |
| Conflicts section empty or escalated (not ignored) | ✅ | The design states no unresolved design conflict; it distinguishes what is revealed (current custodian / first order) from what remains unknown (higher authority and Dohyun's intent). |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical path and scene schema | ✅ | Canonical path, unique scene blocks, required fields, and no template residue pass. |
| Cast / speaker / architecture compliance | ✅ | Appearing and On stage sets match; all four characters are in the catalog and all cited profiles were read this turn. |
| Location / facet / staging compliance | ✅ | Key Location maps to `locations/북문서가-본가.md`; exact facets are profile anchors; the cited staging path exists and covers Scenes 2–3. |
| Length / forecast compliance | ✅ | Exact products recompute to 2,030 / 2,280 / 2,280; scene Est. sum and header addends both equal 6,600 within Scale 4,000–8,000. |
| Episode List / Hook alignment and scope | ✅ | Summary clauses map to beats; body surfaces consistently execute the ten-day-early date Hook without Out creep. |
| Continuity and world/profile compliance | ✅ | Episode 012 and story-so-far are carried forward; no released fact is retconned and no new world rule is invented. |
| **Gate G4 state** | ✅ | Design was explicitly approved by the user before Stage ⑤. |
| **Gate G5 state** | ⬜ | Evaluation is written; explicit user approval is still required before Stage ⑥. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | “붉은 문서함을 실제로 맡은 사람은 누구이며, 그 사람이 왜 도현의 몰락을 준비했는가?” is stated and Scene 1 immediately operationalizes it. |
| Personal stake present | ✅ | Public failure can turn Jinwoo into a record-forgery accomplice; success complicates the father he wants to condemn and risks the successor authority he has built. |
| Episode Out hook | ✅ | The ten-day-early coup date is concrete, legible, and directly actionable for the next episode. |
| Exposition budget respected | ✅ | New information is limited to seal procedure, current custodian role, and first coup order/date; rules are delivered through pressure and comparison rather than lecture. |
| Seed discipline | ✅ | TH-017 is planted/advanced in Scene 2; the date is a Hint in Scene 3. Neither seed claims the held final author or Dohyun motive. |
| Scene-first Key Events (all required fields) | ✅ | Each scene is built around a concrete action and turn, with typed forecast and transition. |
| Sensory-emotional on every scene | ✅ | Scene 1 seal fibers/recognition; Scene 2 ink smell and tightening cord/pressure; Scene 3 wax and divergent ink/acceptance of altered causality. |
| Motifs planned across scenes | ✅ | Seal-knot motif runs Scenes 1–3; voting hands motif runs Scenes 2–3. |
| Overview signature line | ✅ | The episode is event-centered and uses the series constraint of keeping Dohyun morally unresolved; no unused overview signature dialogue is specified in overview.md. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | Information arrives as contested evidence inside a live procedural confrontation; the design does not schedule a detached explanation scene. |
| Sensory-emotional pairing | ✅ | Each sensory cue produces a POV inference or emotional shift, not decorative atmosphere alone. |
| Dialogue voices + Dialogue intent | ✅ | Jinwoo is clipped and controlled; the elder is formal and ledger-bound; Taeseok is low and polite until truth shortens his sentences; Mu-jin is blunt and conditional. |
| Reader-discovered meaning | ✅ | The design asks the reader to infer that procedure distributes responsibility; it Holds the thematic conclusion that the elder representative is the ultimate coup leader. |
| Antagonist plausibility | ✅ | The elder's institutional self-protection, Taeseok's constrained access motive, and the hidden command create plausible competing incentives. |
| Closing image specified | ✅ | The red document box beside the order whose date is ten days earlier than Jinwoo's memory is specified without a thematic monologue. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | The prior episode's red box becomes an immediate procedural weapon; the reader receives a named suspect and a public confrontation rather than a pause for recap. |
| Personal stake matches what this reader came for | ✅ | The episode combines institutional power-grab satisfaction with the family-revenge complication: proving the order may strengthen Jinwoo while making Dohyun's guilt less simple. |
| Pacing / density fits platform expectations | ✅ | Three escalating scenes, 6,600 planned characters, and a concrete reveal/action Out fit the locked web-novel audience. |
| Out hook makes *this* reader want the next episode | ✅ | A remembered future is already obsolete by ten days, and the next tactical objective—seizing the armory—follows immediately. |
| No alienation of core audience without overview intent | ✅ | The procedural material is always attached to accusation, risk, and a tactical reversal; no romance, prolonged training, or healing detour is introduced. |

## Design Critique (required personas)

#### Target Reader
- **Stance:** Adult male-oriented regression/murim serial reader seeking tactical catharsis, institutional pressure, and a stronger next-episode pull.
- **Strengths:** The red box is paid off immediately; Jinwoo wins through remembered knowledge and procedural design rather than unexplained strength; the ten-day discrepancy creates a clean “the future has moved” shock.
- **Defects:** The emotional cost of seeing Dohyun's coup order is mostly stated at episode level rather than attached to one sharply observable Jinwoo reaction. Severity **Low** → Proposed fix: carry to Stage ⑥ as a POV constraint—show Jinwoo suppressing a conclusion about Dohyun while the date destabilizes his certainty, without adding a new beat.
- **Reader impact:** The reader will continue for the altered-future hook; a concrete restrained reaction will keep the reveal from feeling like document mechanics alone.

#### Genre Critic
- **Stance:** Tests the regression-murim and family-revenge contract: tactical reversal, institutional enemy, and morally unstable father figure.
- **Strengths:** The protagonist weaponizes future knowledge, the elders' procedural shield becomes a trap, and the father remains a target without becoming a flat villain.
- **Defects:** —
- **Reader impact:** Delivers the expected “先점한 정보로 판을 뒤집는” satisfaction while preserving the longer revenge question. **Severity:** —. **Proposed fix:** —.

#### Plot Expert
- **Stance:** Checks causality, escalation, Hook body alignment, and Out scope.
- **Strengths:** The prior red-box discovery causes Scene 1; the trap causes the public vote; the access evidence causes Taeseok's choice; opening the box causes the date reveal; the date causes the armory decision. Hook strength is identical across Summary, Out, Arc close, Seed, Turn, and Transition.
- **Defects:** —
- **Reader impact:** The episode feels like a chain of decisions rather than a static reveal. **Severity:** —. **Proposed fix:** —.

#### Reader-Editor
- **Stance:** Tests serial readability, exposition restraint, and closing density.
- **Strengths:** The Out has one main reveal plus its immediate tactical consequence; the procedural explanation is embedded in accusation and counter-accusation; Scenes 2 and 3 have distinct functions despite sharing a staging.
- **Defects:** —
- **Reader impact:** Low skim risk and a clear next-click reason. **Severity:** —. **Proposed fix:** —.

#### Literary Critic
- **Stance:** Tests whether motif, image, and meaning are designed to survive prose generation without authorial explanation.
- **Strengths:** Seal knots and voting hands turn abstract responsibility into recurring physical behavior; the date written beside the red box gives the ending a concrete image; the reader is invited to infer distributed culpability.
- **Defects:** The “hands” motif is explicit in Scenes 2–3 but only the seal-knot motif is carried in Scene 1. Severity **Low** → Proposed fix: Carry-⑥ a light image-level recurrence of hands/pressure in Scene 1, without inserting a new plot beat or changing the design.
- **Reader impact:** A subtle physical motif can make the institutional tension feel authored rather than merely procedural. **Severity:** Low. **Proposed fix:** generation-time motif reminder only.

#### Character Critic
- **Stance:** Checks profile-backed motivation, recognition, voice, and observable character pressure.
- **Strengths:** Jinwoo's habit of reading hands and seals is profile-backed; Taeseok's risk is tied to his brother's disposition record; Mu-jin's exit-counting and conditions match his survival drive; the elder's ledger vocabulary matches his profile.
- **Defects:** —
- **Reader impact:** The four-person pressure scene has differentiated motives and voices rather than interchangeable witnesses. **Severity:** —. **Proposed fix:** —.

#### Setting/Lore Expert
- **Stance:** Checks that the document/seal authority and the room's spatial use remain grounded in the approved world and location catalog.
- **Strengths:** The episode uses existing elder-council, seal, ledger, and archive rules; both location facets are exact Multi-facet anchors; the staging locks who stands where and which props remain in place.
- **Defects:** —
- **Reader impact:** The political mechanism remains legible because every piece of procedure has a visible object or person attached to it. **Severity:** —. **Proposed fix:** —.

## Design Adjudication

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Episode-level emotional cost at the date reveal is not yet specified as an observable POV reaction (Target Reader) | Low | No | yes | The tactical shock is strong, but one restrained reaction will keep the family-revenge promise present without slowing the serial pace. | Generation constraint: preserve Jinwoo's refusal to declare Dohyun innocent and render the destabilization through a concrete physical or clipped verbal reaction; do not add a new plot beat. | Carry-⑥ |
| 2 | Hands/voting-pressure motif could recur lightly in Scene 1 (Literary Critic) | Low | No | yes | A small visual echo strengthens coherence without requiring a design-field change or added exposition. | Generation constraint: echo hand pressure/inspection imagery in Scene 1 while preserving the existing seal-knot motif and unit budget. | Carry-⑥ |

No Schema or Design Consistency finding requires a Stage ④/③ revision. No finding is Pending. Stage ⑤ does not edit the episode design.

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5 — Design Evaluation Approval
- **Status:** Approved by user.
- **Approved:** 사용자가 `1`을 선택해 평가와 두 Carry-⑥ 생성 제약을 승인함.
- **Decision:** Stage ⑥ 원고 생성으로 진행. 설계 수정 없이 승인된 Episode 013 설계와 Carry-⑥ 제약을 적용한다.
- **Carry-⑥ constraints:** 날짜 공개 시 진우의 억제된 신체·언어 반응을 구체화하고, Scene 1에서 손·압력 이미지를 가볍게 반복한다.
- **Next:** Stage ⑥ manuscript generation.
