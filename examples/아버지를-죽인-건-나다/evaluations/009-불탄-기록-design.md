# Design Evaluation: Episode 009 — 불탄 기록

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope criterion. Episode 009 does not retroactively execute the opening-device requirement. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episodes 001–003 scope criterion. Episode 009 is outside the first-confrontation window. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Episode 009 executes a physical trace in Scene 1, a confrontation with 가환 in Scenes 2–3, and the closing hook: “가환은 약 그릇을 확인한 뒤 도현의 방으로 사라진다.” |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | The design keeps 도현’s warning and involvement unresolved; 진우 does not declare 도현 innocent, while 가환’s silence and the 약 그릇 clue complicate the revenge reading. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | This is a later-arc payoff criterion. Episode 009 may advance 가환·약 그릇 threads, but planting or advancing them is not the eventual payoff. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | Scene 1 uses 진우’s investigative habit and prior knowledge to infer a route from physical traces; the design does not grant omniscient future knowledge or a guaranteed answer. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Series/endgame criterion. Episode 009 preserves the conflict but cannot satisfy the final confrontation or conclusion at design scope. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript quality belongs to Stages ⑥–⑦. Forecast arithmetic is recorded in Schema, not scored as a manuscript criterion here. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `episodes/009-불탄-기록.md` uses one episode header, Scene Index, and three canonical `### Scene` sections with flat scene fields. |
| No skill/workflow dump after the design | ✅ | No copied workflow sections such as `## Pre-Design Load`, `## Gate G4`, or `## Prose forecast` appear as a skill dump; the file contains its own design gate only. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Exactly one each of `### Scene 1`, `### Scene 2`, and `### Scene 3`; Beats, Turns, outlines, and budgets are distinct. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | Actual evaluated path: `episodes/009-불탄-기록.md`. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | Scene metadata and bullet fields use the required bold-label notation. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | Scenes 1–3 each include POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, Est. length, Seed touch, and Motif touch. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Episode Appearing is exactly 서진우·가환; Scene 1 has 서진우 and Scenes 2–3 have 서진우·가환. No ghost cast is named in the design surfaces. |
| On stage includes speakers | ✅ | Scene 1 Dialogue intent is `none`; Scenes 2–3 name only 서진우·가환, both present On stage. |
| Characters ⊆ `characters.md` | ✅ | `characters.md` contains catalog rows for 서진우 and 가환; this turn read `characters/서진우.md` and `characters/가환.md` successfully. |
| Summary/Hooks cast alignment | ✅ | Summary, In/Out hooks, Seeds, Closing, and scene fields name only the two Appearing characters. |
| No later-list cast debut | ✅ | 서진우 is present from Episode 001 and 가환’s catalog presence begins at Episode 009; no later-debut character is introduced. |
| Locations ⊆ `locations.md` Key Locations | ✅ | All three scenes use `북문서가-본가`, which maps to the Key Locations row and profile path `locations/북문서가-본가.md`. |
| Location facets ⊆ Multi-facet anchors | ✅ | After reading the profile, `지하 보관실`, `가주전-회랑 접속부`, and `가주전 문앞` are exact labels under `Multi-facet anchors`. |
| Nested `episodes/{slug}/` scene files absent | ✅ | The evaluated episode is a single canonical file; no nested scene-file layout is used. |
| No template residue | ✅ | No raw `{placeholder}` or unfilled template braces remain. |
| Prose forecast present (outline + typed units) | ✅ | Every scene has a numbered paragraph outline and typed dialogue/action/sensory/POV/transition units with integer picks. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Scene 1: `4×180=720 + 3×130=390 + 3×170=510 + 1×80=80 = 1,700`; Scene 2: `5×260=1,300 + 3×190=570 + 2×130=260 + 2×160=320 + 1×80=80 = 2,530`, Est. 2,500; Scene 3: `4×250=1,000 + 3×180=540 + 2×130=260 + 3×160=480 + 1×80=80 = 2,360`, Est. 2,400. Each Est. is within the outline-density band. |
| Dialogue intent vs outline speech | ✅ | Scene 1 is explicitly non-dialogue and its outline contains no speech; Scenes 2–3 contain dialogue units and named speaker intent. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: `1,700 + 2,500 + 2,400 = 6,600`; header addends: `1,700 + 2,500 + 2,400 = 6,600`. |
| Est. length sum ≥ Scale min (hard) | ✅ | Recomputed scene-field sum is 6,600, above overview Scale minimum 4,000. |
| Est. length sum ≤ Scale max (hard) | ✅ | Recomputed scene-field sum is 6,600, below overview Scale maximum 8,000 and within the stated central target band. |
| Cited staging/profile paths exist | ✅ | This turn read successfully: `characters/서진우.md`, `characters/가환.md`, `locations/북문서가-본가.md`, and `stagings/009-불탄-기록-추적.md`. |
| Episode List plot (not a different story) | ✅ | `series.md` Episode 009 Summary says 진우 tracks the person who burned the records, meets 가환, and finds that 가환 recognizes him but stays silent. Scene 1–2 execute those clauses; Scene 3 executes the stated next action. |
| Hook evidence strength (internal) | ✅ | Body quotes align: Episode Summary says “가환은 약 그릇 조각을 확인한 뒤 도현의 방으로 사라진다”; Out repeats “가환은 약 그릇을 확인한 뒤 도현의 방으로 사라진다”; Scene 3 Turn says he enters the inner 가주전 toward 도현’s room; Transition says the door closes after he disappears. The body does not soften or harden the hook. |
| Hook scope (no Out creep) | ✅ | The Out contains the disappearance and one supporting latch sound only; it adds no chase, faction arrival, or second reveal. |
| No design-paste / meta-only scenes | ✅ | Each scene advances a distinct event: physical inference, recognition/confrontation, then the object-check and disappearance. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | The causal chain is clear: ash and drag marks → narrowed suspect → 가환’s recognition and silence → 약 그릇 confirmation → 도현’s room as next pursuit line. |
| Scene transitions chain | ✅ | Scene 1 moves to the corridor junction; Scene 2 moves to the 가주전 door; Scene 3 closes with the door and latch, all in the same morning immediately after Episode 008 Scene 4. |
| Scene sections complete | ✅ | All three Scene Index rows have complete Key Events fields and generation-ready outlines. |
| Generation Readiness | ✅ | All required Schema rows pass, including cast, paths, facets, independent forecast arithmetic, length, and Hook body alignment. |
| Beat concreteness | ✅ | Beats specify physical traces, hand/cloth behavior, direct questions, gaze, object inspection, and movement rather than mood alone. |
| Est. length budget | ✅ | Independent sum is 6,600; it equals the header and remains within 4,000–8,000. |
| Prose forecast quality | ✅ | Typed units correspond to each scene’s action, sensory, POV, dialogue, and transition demands; no range-only or circular padding appears. |
| Episode List scope aligned | ✅ | The approved Summary and Hook to Next are honored without importing Episode 010’s 약 교체 observation. |
| Prior hook addressed | ✅ | Episode 008’s warning and burned childhood record become the direct investigative object in Scene 1; 진우 does not resolve the warning prematurely. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | `Prior Design Alignment` names Episode 008’s burned record and warning, preserves unresolved intent, and explicitly defers Episode 010’s following of 가환 and 약 교체. |
| Series / overview tone & arc honored | ✅ | The design uses cold, event-centered investigation and keeps the P1 revenge reading active while destabilizing its certainty. |
| Episode List Summary / Hook to Next honored | ✅ | The design quotes and operationalizes the Episode 009 Summary and the exact disappearance hook in Scenes 1–3. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Episode Arc close, Out, Scene 3 Turn, Transition out, and Seed touch retain the same observable claim. |
| Characters from architecture; profiles not redefined | ✅ | 서진우 and 가환 are catalogued and their profiles were loaded. The episode cites their established identity/behavior and does not create a new state or rewrite a profile. |
| Profile-backed knowledge / recognition | ✅ | 가환’s recognition and protective silence are supported by `characters/가환.md` Relationships and Drive; 진우’s hand/sleeve tracking is supported by `characters/서진우.md` Behavior. |
| Locations from architecture; profiles not redefined | ✅ | Each scene maps to the `북문서가-본가` Key Locations row and uses only its existing facets. |
| Location profile paths readable | ✅ | Exact path `locations/북문서가-본가.md` was read successfully this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | All three cited facets exactly match the profile’s Multi-facet anchor list. |
| Stagings from episode design; blocking not redefined | ✅ | `stagings/009-불탄-기록-추적.md` is episode-owned, exists, and binds both catalog states, blocking, props, and situation environment for Scenes 2–3. |
| World rules / history consistent with bible | ✅ | The 약 그릇 fragment is used as an ambiguous diagnostic clue, consistent with `world/혈맥계약과-약그릇.md`; no final command chain or treatment result is asserted. |
| No improvised entities or silent lore | ✅ | No new named cast, faction, location, world rule, or unciteable location facet is introduced. |
| Continuity files used (ep 002+) | ✅ | `continuity/story-so-far.md` and `continuity/008-지하의-문-summary.md` are explicitly loaded and cited in the design. |
| Character/location state vs `story-so-far` | ✅ | 진우 remains the 후계자 with the recovered record fragment; 가환 remains a possible internal link with motive unresolved; 본가’s discovered underground storage state is preserved. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-001, TH-002, TH-004, TH-006, TH-009, TH-010, and TH-011 are picked up/advanced; TH-008 is planted as a possible parallel; unresolved warning, command chain, and exact motive remain held. |
| No contradiction of released continuity | ✅ | Nothing reverses Episode 008’s discovery or the released state. The episode begins immediately afterward. |
| Conflicts section empty or escalated (not ignored) | ✅ | The design records `Conflicts / open questions: None` after its facet preflight and separately preserves all unresolved questions in its Hold and continuity tables; no architecture conflict is hidden. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Required artifacts loaded | ✅ | Overview, series Episode 009 row, Appearing profiles, used location profile, world bible/aspect, cited staging, and Episode 008 continuity set were read; this evaluation also forced exact cited profile reads. |
| Canonical path | ✅ | `episodes/009-불탄-기록.md` is the actual canonical episode path. |
| Locations index / paths / facets | ✅ | `북문서가-본가` is in `locations.md`; exact profile and staging paths were read successfully; all three facets are exact profile anchors. |
| Length / forecast evidence | ✅ | Written and independently recomputed unit products agree; scene-field and header totals both equal 6,600. |
| Episode List Summary and Hook | ✅ | Every Summary clause is executed and the body Hook surfaces retain the same evidence strength. |
| Architecture, continuity, cast, staging, and tone | ✅ | Catalog bindings, profile-backed knowledge, staging continuity, world ambiguity, and P1 arc direction all pass. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | “누가 진우의 어린 시절 기록에서 진실로 이어지는 장만 골라 태웠으며, 가환은 왜 진우를 알아보고도 모르는 사람처럼 침묵하는가?” is explicit. |
| Personal stake present | ✅ | The burned pages target the childhood evidence that shaped 진우’s belief that he had to kill 도현; suspecting 가환 threatens both revenge and remembered care. |
| Episode Out hook | ✅ | 가환’s checking of the fragment and disappearance into 도현’s room is concrete, narrow, and directly advances the next episode. |
| Exposition budget respected | ✅ | The design limits explanation of bloodline contracts and the bowl; clues remain physical and unresolved. |
| Seed discipline | ✅ | The burner is a Plant advanced through trace and silence; the bowl fragment is a Hint and is not converted into a final explanation. |
| Scene-first Key Events (all required fields) | ✅ | All three scenes have complete, event-driven fields and concrete paragraph outlines. |
| Sensory-emotional on every scene | ✅ | Ash/sweat/bitter medicine, lamp oil/medicine residue, and cold medicine odor/black line each pair sensory detail with 진우’s reaction. |
| Motifs planned across scenes | ✅ | `재와 약재 얼룩` touches Scenes 1–3; `닦이는 가장자리` touches Scenes 2–3 and is carried as a hand action rather than explained. |
| Overview signature line | N/A | `overview.md` defines no signature dialogue line requiring placement or Hold. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | Information is delivered as ash, drag marks, stains, and evasive behavior under immediate pursuit pressure; the design does not schedule a lore lecture. |
| Sensory-emotional pairing | ✅ | Each scene couples a tactile/olfactory clue to a specific interruption, memory, or decision by 진우. |
| Dialogue voices + Dialogue intent | ✅ | 진우’s short, declarative pressure contrasts with ��환’s slow honorifics and omitted subjects; intent is explicit in Scenes 2–3. |
| Reader-discovered meaning | ✅ | The possibility that silence is tied to a promise is staged through avoidance and object handling; the design explicitly forbids confirming protector status. |
| Antagonist plausibility | ✅ | 도현 is not present as a speaking antagonist, but his possible warning and the protective/complicit reading remain behaviorally plausible without premature absolution. |
| Closing image specified | ✅ | The bowl fragment’s briefly visible pattern in 가환’s palm and the closing 가주전 door are specified without a thematic monologue. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
Target Reader locked from `overview.md`: adult male-oriented web-novel readers who enjoy regression martial arts, sect power struggles, revenge-driven catharsis, and family reversals.

| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | The episode opens on a concrete forensic trace immediately after the underground discovery, rather than recap or training. |
| Personal stake matches what this reader came for | ✅ | The investigation threatens the protagonist’s revenge premise and turns a familiar caretaker into a suspect, combining action logic with family reversal. |
| Pacing / density fits platform expectations | ✅ | Three escalating scenes, 6,600-character forecast, short pressure dialogue, and no lore detour fit the locked serial format. |
| Out hook makes *this* reader want the next episode | ✅ | A recognizable caretaker silently enters the father’s room after checking a suspicious bowl fragment, creating a specific pursuit question for Episode 010. |
| No alienation of core audience without overview intent | ✅ | The design preserves revenge, investigation, and martial-family stakes; it complicates rather than replaces the genre promise. |

## Design Critique (required personas)

#### Target Reader
- Stance: Read as an adult male-oriented regression/revenge web-novel reader seeking concrete reversals, tactical inference, and a sharp next-episode pull.
- Strengths: The episode converts Episode 008’s discovery into an active trace, makes 가환’s recognition a personal reversal, and ends on a highly legible destination: 도현’s room.
- Defects: —
- Reader impact: The reader receives a clue-based win and a family-trust complication without losing the revenge question.

#### Genre Critic
- Stance: Test the episode against regression martial-arts and family-revenge genre contracts.
- Strengths: 진우’s prior knowledge functions as a tactical investigative advantage, while the unknown motive prevents a cheap “father was innocent” reversal. The episode has a concrete object hook and a suspect confrontation.
- Defects: Low — The design’s action is investigative rather than physically violent; generation should preserve pressure through blocking and the possibility of immediate pursuit so the episode does not read as a static interview. Proposed fix: carry the existing blocking and time pressure into Stage ⑥; no Stage ④ change required.
- Reader impact: A reader seeking cathartic domination still gets control of the route and interrogation, while the quieter scene earns its place by changing the meaning of the evidence.

#### Plot Expert
- Stance: Check causality, episode scope, thread timing, and Hook body alignment/scope.
- Strengths: The causal chain is clean and the Episode 010 event is explicitly deferred. The body Summary, Out, Scene 3 Turn, Transition, and Seed touch all preserve the same disappearance-after-checking claim.
- Defects: —
- Reader impact: The reader can tell what was discovered, what remains unknown, and what action must follow.

#### Reader-Editor
- Stance: Judge serial readability, exposition restraint, scene escalation, and closing density.
- Strengths: The three-scene ladder escalates from evidence to face-to-face recognition to a locked door. The final beat carries one primary obligation—follow 가환 into 도현’s room—with only the latch sound as reinforcement, not a separate plot signal.
- Defects: Low — Scene 2 and Scene 3 both depend on verbal pressure; prose must vary the physical micro-actions so the middle does not become repetitive questioning. Proposed fix: use the designed cloth edge, exit glances, fragment handling, and left-side blocking as observable turns in Stage ⑥.
- Reader impact: Correct execution will keep a mobile serial reader moving through the confrontation instead of skimming a repeated interrogation.

#### Literary Critic
- Stance: Test whether motif and meaning are discoverable through concrete design rather than asserted interpretation.
- Strengths: Ash, medicine stains, and the wiped edge bind erasure to possible protection. The closing image is an object-and-door image, and the design explicitly withholds the explanatory verdict.
- Defects: Low — The motif meaning could become over-explicit if the manuscript names “protection” or “promise” as a conclusion. Proposed fix: carry the Hold instruction into generation; let the repeated hands, stains, and silence do the interpretive work.
- Reader impact: Readers can revise their judgment of 가환 and 도현 themselves, preserving the series’ moral tension.

#### Character Critic
- Stance: Check motivation, relationship pressure, voice distinction, and profile-backed recognition.
- Strengths: 가환’s recognition is grounded in the established relationship and protective drive; his silence is consistent with the profile’s shared secret with 도현. 진우’s scrutiny of hands and sleeves is directly profile-backed, and the two voices are distinct.
- Defects: Low — The manuscript must not let 가환’s silence become a generic “mysterious old servant” pose. Proposed fix: preserve the specific left ear, medicine stains, cloth-edge wiping, omitted subjects, and gaze-to-floor behavior as the evidence of a person choosing silence.
- Reader impact: The caretaker becomes a morally pressured participant rather than a disposable clue dispenser, increasing investment in the next encounter.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Preserve concrete blocking and physical pressure so the investigative confrontation does not flatten into a static interview (Genre Critic, Reader-Editor) | Low | No | yes | The locked reader accepts a quiet investigation when each exchange changes the tactical position; this is a generation constraint, not a design defect. | Carry cloth, exit glances, fragment handling, left/right blocking, and pursuit pressure into Stage ⑥. | Carry-⑥ |
| 2 | Do not state the protective interpretation or turn 가환’s silence into generic mystery (Literary Critic, Character Critic) | Low | No | yes | The core appeal is the unresolved revenge/family reversal; explaining it now would reduce the next-episode question. | Carry the Hold: no confirmation that 가환 burned the records to save 진우 and no confirmation that 도현 wrote the warning; preserve profile-specific behavior in prose. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |
