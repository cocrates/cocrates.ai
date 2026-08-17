# Design Evaluation: Episode 012 — 불이 난 날

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope; this design-evaluation does not re-evaluate the first-episode introduction. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003 scope; this episode preserves the already released sequence. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scenes 1–4 execute a record/transport investigation, institutional confrontation, and a concrete next target: the person inside the elders’ council holding the original. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | Scene 2 finds a possible child-rescue route while Scene 3 retains the forged-record suspicion aimed at Do-hyeon; his inner motive remains undisclosed. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series/late-arc payoff criterion; Episode 012 may advance the threads but cannot satisfy the later causal payoffs. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | N/A | No decisive future-memory payoff is designed in this episode; the active method is document and testimony verification. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Endgame criterion; this episode only intensifies the father/son conflict. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript quality and final prose density belong to Stages ⑥–⑦; forecast arithmetic is recorded below under Schema. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | The file uses one `## Scenes` section with four unique canonical `### Scene n` sections and the required scene fields. |
| No skill/workflow dump after the design | ✅ | No copied workflow sections such as `## Pre-Design Load`, `## Gate G4`, or prose-generation instructions appear after the design body. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Exactly four headings, Scene 1 through Scene 4; each Beat and outline has distinct events. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | Actual path read this turn: `episodes/012-불이-난-날.md`; it is the canonical Korean slug. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | Scene metadata and bullet fields consistently use the required bold-colon notation. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | All four scenes contain POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, and exactly one Est. length. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing set is exactly {서진우, 서도현, 문상철, 장로 대표}; the union of scene On stage sets is identical. |
| On stage includes speakers | ✅ | Scene 1 includes 진우·문상철; Scenes 2–3 include 진우·문상철·도현 / 진우·도현; Scene 4 includes 진우·장로 대표. Dialogue-intent speakers are on stage. |
| Characters ⊆ `characters.md` | ✅ | `characters.md` and the four profile paths were read this turn; all four are catalogued. |
| Summary/Hooks cast alignment | ✅ | Summary, In, Out, Seeds, and scene closing references use only the appearing cast or already named architecture roles. |
| No later-list cast debut | ✅ | No cast member is first introduced ahead of the Episode List; all four are present by Episode 012. |
| Locations ⊆ `locations.md` Key Locations | ✅ | Every scene maps to the Key Location `북문서가 본가`; the listed facets are internal facets of that location. |
| Location facets ⊆ Multi-facet anchors | ✅ | `회랑`, `가주전 문앞`, and `장로회당 표결단` were checked against `locations/북문서가-본가.md` and are exact citeable anchors. |
| Nested `episodes/{slug}/` scene files absent | ✅ | The actual episode is a flat file under `episodes/`; no nested scene-file structure is used. |
| No template residue | ✅ | No raw placeholder braces or incomplete template markers remain. |
| Prose forecast present (outline + typed units) | ✅ | Each scene has a six-line paragraph outline and integer typed units: dialogue, action, sensory, POV, transition. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Recomputed: S1 `3×250 + 2×180 + 1×120 + 2×140 + 1×80 = 1,590`, Est. 1,600; S2 `4×250 + 2×180 + 1×120 + 2×140 + 1×80 = 1,840`, Est. 1,800; S3 `3×250 + 3×180 + 1×120 + 2×140 + 1×80 = 1,770`, Est. 1,800; S4 same as S2 = 1,840, Est. 1,800. Each Est. is inside its outline-density band (6 lines × 200–350 = 1,200–2,100). |
| Dialogue intent vs outline speech | ✅ | Every scene has dialogue units and explicit dialogue intent; no scene claims `none` while its Beat/outline contains speech. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: `1,600 + 1,800 + 1,800 + 1,800 = 7,000`; header addends: `1,600 + 1,800 + 1,800 + 1,800 = 7,000`. |
| Est. length sum ≥ Scale min (hard) | ✅ | Recomputed scene-field sum is 7,000, above the 4,000 minimum. |
| Est. length sum ≤ Scale max (hard) | ✅ | Recomputed scene-field sum is 7,000, below the 8,000 maximum and inside the target central band. |
| Cited staging/profile paths exist | ✅ | This turn read `characters/서진우.md`, `characters/서도현.md`, `characters/문상철.md`, `characters/장로-대표.md`, `locations/북문서가-본가.md`, and `world/혈맥계약과-약그릇.md` successfully. All scene stagings are `none`, so no staging profile is claimed. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary says “자신의 출생 당일 기록을 조사” (Scenes 1–2), “도현이 아이를 구하기 위해 한밤중에 문파를 비웠다는 사실” (Scene 2), and “기록을 조작한 흔적은 도현에게 향한다” (Scene 3); each clause is executed in a Beat/Turn. |
| Hook evidence strength (internal) | ✅ | Body quotes align: Summary “장로회 기록 권한을 역으로 묶어…원본이 장로 회의 안에”; Episode Out “조작된 기록의 원본을 가진 사람은 장로 회의 안에 있다”; Scene 4 Turn “장로회 내부자가 있다는 증거”; Scene 4 Transition “원본을 가진 사람을 특정할 수 있는 다음 표결 기록”. |
| Hook scope (no Out creep) | ✅ | The closing adds no chase, new faction, or second reveal; the next action remains identifying the internal original-holder within the stated hook. |
| No design-paste / meta-only scenes | ✅ | All scenes contain a concrete investigation or confrontation turn, and no Beat/outline/unit-budget block is duplicated. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | Testimony fixes the date → records reveal a rescue route → seal evidence reopens suspicion → institutional pressure identifies the next investigative locus. |
| Scene transitions chain | ✅ | Scene 1 requests birth records; Scene 2 moves from time-grid comparison to seal inspection; Scene 3 moves to the authority governing seals/originals; Scene 4 ends with the next vote-record target. All `When` fields remain sequentially the same morning. |
| Scene sections complete | ✅ | All four Scene Index rows have matching full sections with every required field. |
| Generation Readiness | ✅ | All Schema rows are ✅, including length, cast, paths, facets, and body-level Hook evidence; no blocking adjudication is currently required. |
| Beat concreteness | ✅ | Each Beat specifies an observable document, gesture, question, discovery, or institutional response. |
| Est. length budget | ✅ | Recomputed total 7,000 = header total = scene-field total; 4,000–8,000 scale and central target band pass. |
| Prose forecast quality | ✅ | Typed units correspond to the declared dialogue, action, sensory, POV, and transition work in each scene. |
| Episode List scope aligned | ✅ | The design executes the Episode 012 Summary and preserves the single Hook to Next without Out creep. |
| Prior hook addressed (ep 002+) | ✅ | Scene 1 directly resumes from the captured 문상철 and his claim about the birth-night fire, as recorded in Episode 011 continuity. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | The design records Phase A indexes, Phase B profiles, the world aspect, and `story-so-far.md` plus the immediate Episode 011 summary. |
| Series / overview tone & arc honored | ✅ | The episode remains a cold, event-centred P1 investigation with emotion surfacing around 진우’s birth and his father. |
| Episode List Summary / Hook to Next honored | ✅ | The Summary and Hook are executed without resolving the held motive, arsonist, or custodian identity. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Episode Arc, Out, Scene 4 Turn/Transition, Seed, and closing image all converge on the elders’ council’s original-holder investigation. |
| Characters from architecture; profiles not redefined | ✅ | All appearing characters are catalogued and the read profiles support their drives, voices, and identifying behavior. |
| Profile-backed knowledge / recognition | ✅ | Jin-woo observes documents, seals, transport records, and bodily gestures; he does not claim access to Do-hyeon’s inner motive or the custodian’s identity. |
| Locations from architecture; profiles not redefined | ✅ | The sole set is the catalogued `북문서가 본가`; scene facets are cited rather than newly defined. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` was read successfully this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | `회랑` is an internal named route and `가주전 문앞` / `장로회당 표결단` are exact Multi-facet anchors in the profile. |
| Stagings from episode design; blocking not redefined | ✅ | Each scene is designed as a distinct single investigative beat and explicitly uses `Staging: none`; no reusable multi-scene blocking is claimed. |
| World rules / history consistent with bible | ✅ | Seals and stamps are used as evidentiary traces, not as proof of the full blood-contract structure; the touched world file confirms that the original and blood-mark are needed for full certainty. |
| No improvised entities or silent lore | ✅ | No new person, faction, place, rule, or uncatalogued location facet is introduced. |
| Continuity files used (ep 002+) | ✅ | Both `continuity/story-so-far.md` and `continuity/011-협박문의-조각-summary.md` were loaded and their immediate state is reflected. |
| Character/location state vs `story-so-far` | ✅ | 문상철 begins captured and is used as a witness; 도현’s altered medicine and motive remain unresolved; the 본가 remains the active investigation site. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-015 and TH-016 advance; TH-001, TH-002, TH-004, TH-009, TH-011, TH-013, and TH-014 are picked up or advanced; the fire cause, medicine purpose, letter source, and custodian identity are held. |
| No contradiction of released continuity | ✅ | The episode starts immediately after Episode 011’s capture and preserves the missing original letter and uncertain responsibility. |
| Conflicts section empty or escalated (not ignored) | ✅ | The design records `None` and explicitly explains why “서쪽 건물” remains a testimony target rather than an invented new location. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Required load and path existence | ✅ | All Architecture Reference profile paths were read successfully this turn; every staging is `none`. |
| Locations index / path / facets | ✅ | `북문서가 본가` is a Key Locations entry; `locations/북문서가-본가.md` read OK; all three used facets are citeable. |
| Cast and speaker alignment | ✅ | All four appearing characters are catalogued and all dialogue speakers are on stage in their scenes. |
| Forecast and length | ✅ | Exact unit products and independent recomputation pass; 7,000 is within 4,000–8,000. |
| Episode List plot and Hook body alignment | ✅ | Summary clauses and Hook strength are quoted and matched by the scenes, Out, Turn, and Seeds. |
| Continuity and no silent lore | ✅ | Episode 011’s capture and threads TH-015/TH-016 are carried forward without resolving held questions. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | “문상철이 말한 화재는 도현이 숨긴 죄의 현장인가, 아니면 도현이 진우를 살려낸 밤의 흔적인가?” directly links the prior hook to Jin-woo’s personal dilemma. |
| Personal stake present | ✅ | Jin-woo’s own birth record collides with the possibility that the father he wants to punish rescued him. |
| Episode Out hook | ✅ | The original-holder target is concrete, single, and consistent with the series Hook; it does not promise an unrelated escalation. |
| Exposition budget respected | ✅ | The design limits new information to fire date, absence/movement mark, resealing evidence, and record authority; full contract and medicine conditions remain held. |
| Seed discipline | ✅ | The fire record and Do-hyeon seal are Plants; the elders’ internal custodian is a Hint, matching the Hook without revealing identity. |
| Scene-first Key Events (all required fields) | ✅ | Each scene has an observable situation, beat, turn, function, sensory-emotional pairing, dialogue intent, and transition. |
| Sensory-emotional on every scene | ✅ | Each scene pairs smell/light/texture or gesture with Jin-woo’s controlled emotional response. |
| Motifs planned across scenes | ✅ | Soot/ink bleed spans Scenes 1–3; the hand checking seals spans Scenes 2–4. |
| Overview signature line | ✅ | The constraint dialogue is not required as a literal repeated line; the episode does use cold event-centred action and emotion only around father/son evidence. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | Each discovery is attached to an interrogation, gesture, or risk of losing the original; exposition is limited by the episode budget. |
| Sensory-emotional pairing | ✅ | Ink, soot, medicine smell, wax, and hands expose controlled shifts in Jin-woo’s certainty. |
| Dialogue voices + Dialogue intent | ✅ | Jin-woo is short and probing; Moon Sang-cheol calculates liability; Do-hyeon speaks slowly in restrained honorifics; the Elder Representative hides behind archival terms. |
| Reader-discovered meaning | ✅ | The design holds the moral conclusion and asks the reader to reconcile rescue evidence with forged-record suspicion. |
| Antagonist plausibility | ✅ | The institutional pressure comes from a representative protecting authority and procedure rather than from an arbitrary confession. |
| Closing image specified | ✅ | The representative’s hand slips from the gold ring and exposes the red document box’s seal cord beneath the ballot table. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
**Locked Target Reader:** Adult male web-novel readers who prefer regression/return martial-arts fiction, sect-power consolidation, cathartic revenge, and family reversals.

| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | A captured informant, an empty case/seal cord, and a verifiable birth-night fire give immediate investigative motion. |
| Personal stake matches what this reader came for | ✅ | The revenge target is reinterpreted through a concrete rescue trace, intensifying rather than replacing the father/son conflict. |
| Pacing / density fits platform expectations | ✅ | Four event-dense scenes, 7,000-character forecast, short interrogation beats, and no training or lore-only section fit the locked platform expectation. |
| Out hook makes *this* reader want the next episode | ✅ | The next target is an internal elder-council custodian, enabling factional pressure and a specific investigative continuation. |
| No alienation of core audience without overview intent | ✅ | The design does not turn into reconciliation or passive family drama; it preserves suspicion, institutional pressure, and practical evidence gathering. |

## Design Critique (required personas)

#### Target Reader
- Stance: Read as an adult male regression/martial-arts web-novel reader seeking a concrete reversal and a forceful next move.
- Strengths: The episode turns the prior capture into immediate evidence work; the rescue possibility complicates revenge without absolving Do-hyeon; the Elder Council target promises a satisfying power struggle.
- Defects: The closing could feel procedural if the red document box and the “internal custodian” conclusion are explained at equal length → severity Low → proposed fix: in Stage ⑥ make the exposed seal cord the dominant final image and keep the conclusion action-led.
- Reader impact: The reader is likely to continue because the father mystery deepens while Jin-woo gains a manipulable institutional target.

#### Genre Critic
- Stance: Test the episode against regression martial-arts and revenge-serial promises.
- Strengths: Future knowledge is not used as omniscience; transport records, seals, and authority are weaponized as genre-appropriate evidence; the episode ends with a factional target.
- Defects: —
- Reader impact: The design delivers investigation-as-power and a father-related reversal without pausing the revenge engine.

#### Plot Expert
- Stance: Audit causality, escalation, Hook body alignment, and Hook scope.
- Strengths: The causal chain is clean: testimony → dated transport record → birth/child-movement record → reseal discrepancy → archival authority pressure. Body surfaces consistently point to the Elder Council’s internal original-holder.
- Defects: The same-morning chain contains two adjacent record-inspection scenes at the same facet; if their action boundaries blur in prose, the middle may read as one extended examination → severity Low → proposed fix: Stage ⑥ should make Scene 2’s emotional discovery (possible rescue) and Scene 3’s forensic reversal (forged seal) visibly distinct.
- Reader impact: Clear scene-level reversals protect the serial rhythm; a blurred middle would reduce the sense of progress.

#### Reader-Editor
- Stance: Evaluate readability, exposition load, and the sellability of the closing beat.
- Strengths: The design limits lore, gives every scene a question/answer movement, and provides a concrete Out rather than a vague “more secrets” promise.
- Defects: The final Transition carries the rules disclosure, the representative’s slip, the red box reveal, and the next-record declaration—four signals that risk a crowded closing beat → severity Med → proposed fix: treat the red box/seal cord as the visual climax; subordinate the procedural declaration to one short line in prose.
- Reader impact: A clean final image will produce a sharper “next episode” click for mobile serial readers.

#### Literary Critic
- Stance: Assess planned motif, restraint, and whether meaning is left for the reader to discover.
- Strengths: Soot/ink bleed and seal-checking hands give material form to the conflict between protection and falsified authority; the design refuses a thematic monologue and keeps Do-hyeon’s motive withheld.
- Defects: —
- Reader impact: The moral discomfort survives beyond the plot clue: the same mark can indicate protection, concealment, or impersonated authority.

#### Character Critic
- Stance: Check motivation, voice, behavior continuity, and profile-backed recognition.
- Strengths: Jin-woo’s profile-backed habits—watching hands/seals and tapping the scabbard—turn investigation into character action. Do-hyeon’s right index-finger and restrained voice support his evasive presence; Moon Sang-cheol’s liability-calculating speech matches his profile; the Elder Representative’s ring and ledger jargon are functional rather than decorative.
- Defects: Jin-woo’s decision not to confront Do-hyeon risks appearing merely strategic unless the prose gives a visible personal cost to withholding the accusation → severity Low → proposed fix: carry to Stage ⑥ as an observable pause/gesture, not explanatory introspection.
- Reader impact: Showing the cost of restraint keeps Jin-woo emotionally legible while preserving the mystery.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Closing beat risks crowding the red-box image with procedural explanation (Reader-Editor) | Med | No | yes | The target reader needs one sharp visual Out more than several simultaneous explanations. | Carry the red document box/seal cord as the dominant final image; compress the next-record declaration during generation. | Carry-⑥ |
| 2 | Scene 2 and Scene 3 must remain distinct in prose despite the shared facet (Plot Expert) | Low | No | yes | A visible emotional discovery followed by a forensic reversal preserves the episode’s two-step hook. | Make the rescue possibility the Scene 2 endpoint and the reseal discrepancy the Scene 3 reversal; do not merge their functions. | Carry-⑥ |
| 3 | Jin-woo’s restraint before Do-hyeon needs an observable personal cost (Character Critic) | Low | No | yes | The reader should experience restraint as conflict, not as convenient plot management. | Use profile-backed gestures and a brief controlled pause; avoid adding inner exposition or a new plot beat. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5 (Design Eval Approval)
- **Status:** Approved by user
- **Approved:** 사용자가 선택지 1(수정된 Episode 012 설계 승인 후 Stage ⑤ 진행)을 선택해 본 설계 평가와 세 가지 Carry-⑥ 생성 제약을 승인했다.
- **Decision:** 설계 변경 없이 Stage ⑥ 원고 생성으로 진행한다.
- **Gate G4:** 사용자의 동일 승인으로 수정된 Episode 012 설계도 승인되었다.
- **Generation constraints to carry forward:**
  1. Make the red document box/seal cord the dominant closing image and compress procedural explanation.
  2. Keep Scene 2’s rescue possibility and Scene 3’s forensic seal reversal visibly distinct.
  3. Show Jin-woo’s personal cost of withholding confrontation through observable gesture/pause, without adding new lore or plot.
- **Next:** Stage ⑥ manuscript generation; run Design-Fidelity and Prose Quality Floor before Gate G6.
