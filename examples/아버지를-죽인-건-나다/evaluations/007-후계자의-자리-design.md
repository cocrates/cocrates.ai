# Design Evaluation: Episode 007 — 후계자의 자리

> Evaluation stage: ⑤ Design Evaluation
> Evaluated design: `episodes/007-후계자의-자리.md`
> Status note: This evaluation was requested before explicit Gate G4 approval. G4 remains open.

## Criteria Check (from overview.md)

| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope; outside Episode 007 design-evaluation scope. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Already executed in Episodes 001–003; outside this episode’s scope. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scenes 1–3 execute institutional power seizure and procedural conflict; Scene 4 discovers the marked underground entrance promised by the Episode 007 Hook to Next. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | Dohyun observes, withholds the test’s purpose, recognizes Jinwoo as successor without celebration, and questions his observable choice rather than explaining his inner motive. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Late-arc / series-payoff criterion; Episode 007 only advances document, seal, and Black Wind Tower threads. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | Jinwoo compares documents, seals, and official responsibility. Scene 4 explicitly frames his action as an inference from “표결 문서와 시험 봉인의 방향,” not prior omniscience. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Endgame criterion; Episode 007 maintains father–son pressure but cannot satisfy the final resolution. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript Stage ⑥/⑦ criterion; forecast is checked under Schema. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)

| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | Four scenes use the required meta lines and flat bullet fields. |
| No skill/workflow dump after the design | ✅ | The file contains episode design only; no workflow sections are pasted into the artifact. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–4 headings are unique. Scene 2’s irreversible official-record outcome is distinct from Scene 1’s initial split declaration. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | Actual path is `episodes/007-후계자의-자리.md`. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | Required fields use canonical notation throughout. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | All four scenes contain POV, Location, When, On stage, Staging, Situation through Est. length. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Roster is exactly 서진우, 서도현, 장로-대표, 백무진, 윤태석; the union of scene On stage lists matches it. |
| On stage includes speakers | ✅ | Every intentional speaker and actor named in Beat or Dialogue intent is on stage in that scene; no anonymous crowd speech is used. |
| Characters ⊆ `characters.md` | ✅ | All five names are catalogued in `characters.md`; this turn read their exact profiles: `characters/서진우.md`, `characters/서도현.md`, `characters/장로-대표.md`, `characters/백무진.md`, `characters/윤태석.md`. |
| Summary/Hooks cast alignment | ✅ | Summary, hooks, seeds, and closing fields use only the five appearing characters. |
| No later-list cast debut | ✅ | All five are present by Episode 007 in the approved Episode List / catalog. |
| Locations ⊆ `locations.md` Key Locations | ✅ | Every scene maps to `북문서가-본가`, whose slug is present in the Key Locations table. |
| Location facets ⊆ Multi-facet anchors | ✅ | The exact cited facets `장로회당 표결단`, `중정 회화나무`, and `가주전-회랑 접속부` occur under `locations/북문서가-본가.md` Multi-facet anchors. The discovered passage is not cited as a new Location facet. |
| Nested `episodes/{slug}/` scene files absent | ✅ | The design is one canonical episode file. |
| No template residue | ✅ | No unresolved template braces or workflow placeholders remain. |
| Prose forecast present (outline + typed units) | ✅ | Every scene has a paragraph outline and five-type integer unit formula. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1: 3×250 + 2×180 + 2×120 + 2×140 + 1×80 = 1,710; Sc2 = 1,890; Sc3 = 2,140; Sc4 = 1,710. Est. values 1,700 + 1,900 + 2,100 + 1,700 are within tolerance and outline-density bands. |
| Dialogue intent vs outline speech | ✅ | Each scene with speech has a non-`none` Dialogue intent consistent with its outline and On stage roster. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields: 1,700 + 1,900 + 2,100 + 1,700 = 7,400. Header addends: 1,700 + 1,900 + 2,100 + 1,700 = 7,400. |
| Est. length sum ≥ Scale min (hard) | ✅ | Recomputed 7,400 ≥ overview minimum 4,000. |
| Est. length sum ≤ Scale max (hard) | ✅ | Recomputed 7,400 ≤ overview maximum 8,000; central target band is met. |
| Cited staging/profile paths exist | ✅ | This turn read successfully: `characters/서진우.md`, `characters/서도현.md`, `characters/장로-대표.md`, `characters/백무진.md`, `characters/윤태석.md`, `locations/북문서가-본가.md`, and `stagings/007-장로-표결-압박.md`. The cited indexes `characters.md`, `locations.md`, and `stagings.md` were also read successfully. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary says Jinwoo splits the elders’ votes and wins by twisting Dohyun’s succession test; Scenes 1–3 execute those clauses. Scene 4 keeps Dohyun’s purpose at question-level and preserves the Episode List Hook. |
| Hook evidence strength (internal) | ✅ | Body surfaces use discovery strength: Summary says the marked entrance is discovered; Out says the Black Wind Tower passage is found; Seeds limits the reveal to “표식과 입구”; Scene 4 Turn says the entrance exists while holding destination, structure, and purpose. |
| Hook scope (no Out creep) | ✅ | The closing adds one discovery obligation only. It does not confirm the house-to-network structure, destination, contents, purpose, or a second reveal. |
| No design-paste / meta-only scenes | ✅ | Each scene contains a distinct dramatic event: procedural conversion, irreversible record, rule reversal, and interrogation plus discovery. |

## Structure & Arc Checks

| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | Personal death record → initial vote split → sealed institutional consequence → test reversal → father–son interrogation → marked entrance discovery. |
| Scene transitions chain | ✅ | Scene 1’s test notice leads to Scene 2’s final record; Scene 2 moves to the test; Scene 3’s recognition leads to Scene 4; Scene 4 closes on the discovery hook. |
| Scene sections complete | ✅ | All four Scene Index rows have complete canonical scene sections. |
| Generation Readiness | ✅ | All Schema rows pass; no path, facet, cast, length, forecast, or Hook-body defect remains. |
| Beat concreteness | ✅ | Observable actions include withholding the register, comparing seals, sealing the vote record, cutting the judgment-stone cord, and exposing the marked entrance. |
| Est. length budget | ✅ | Independent total is 7,400; fields, header addends, Scale, and forecast agree. |
| Prose forecast quality | ✅ | Typed units correspond to dialogue, action, sensory, POV, and transition work in each outline. |
| Episode List scope aligned | ✅ | Summary and Hook are honored without confirming the held structural relation beneath the house. |
| Prior hook addressed (ep 002+) | ✅ | Scene 1 begins immediately after the Episode 006 disclosure that Jinwoo is the next victim. |

## Architecture & Continuity Compliance

| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected in Prior Design Alignment | ✅ | The design records the required overview, series, appearing profiles, used location, world bible, staging, and Episode 006 continuity set. |
| Series / overview tone & arc honored | ✅ | Cold, procedural, event-centered institutional pressure advances P1 control while preserving Dohyun’s ambiguity. |
| Episode List Summary / Hook to Next honored | ✅ | Summary and Hook are quoted and executed at the same discovery strength. |
| Hook internal consistency (design surfaces) | ✅ | Summary / Arc / Out / Seeds / Scene 4 Turn all describe discovery of a marked entrance while holding its explanation. |
| Characters from architecture; profiles not redefined | ✅ | All appearing characters are catalogued and their drives, voices, and states are used without redefinition. |
| Profile-backed knowledge / recognition | ✅ | The former unsupported “already knew the third seal location” claim is removed. Dohyun now questions an observable immediate movement, and Jinwoo attributes it to comparing documents and seal direction. |
| Locations from architecture; profiles not redefined | ✅ | All scene Locations map to the single approved Key Location and its exact anchors. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` was read OK this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | All three cited facets are exact anchors in the successfully read profile. |
| Stagings from episode design; blocking not redefined | ✅ | `stagings/007-장로-표결-압박.md` was read OK; Scene 1–2 use its fixed cast, positions, props, and continuity span. |
| World rules / history consistent with bible | ✅ | The design uses established document/insignia traceability and treats the new passage as a discovery, not a new supernatural or spatial rule. |
| No improvised entities or silent lore | ✅ | The prior confirmed test-site/Black Wind Tower connection is removed. The black mark and entrance are the approved Hook-level discovery, while structural relation and purpose are explicitly held. |
| Continuity files used (ep 002+) | ✅ | `story-so-far.md`, `continuity/006-죽음의-명부-summary.md`, and `unresolved-threads.md` were loaded and reflected. |
| Character/location state vs `story-so-far` | ✅ | Character states and the main-house state remain aligned; the unreleased passage is not prematurely added to continuity. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-006, TH-007, and TH-009 advance; TH-001–005 and TH-008 have explicit Holds. |
| No contradiction of released continuity | ✅ | Episodes 001–006 remain untouched; the episode does not claim the unreleased passage was previously known. |
| Conflicts section empty or escalated (not ignored) | ✅ | The design’s Conflicts section records the resolved treatment: discovery only, knowledge grounded in observable inference, and Scene 2’s distinct irreversible consequence. |

## Design Consistency Gate

| Check | Result | Evidence |
|---|---|---|
| Loaded required artifacts | ✅ | Required episode-specific indexes, profiles, staging, world bible, and continuity paths were read in this turn. |
| Canonical path | ✅ | `episodes/007-후계자의-자리.md`. |
| Locations index | ✅ | Sc1–4 use `북문서가-본가`, present in `locations.md` Key Locations. |
| Locations path | ✅ | Exact path `locations/북문서가-본가.md` read OK; exact staging path `stagings/007-장로-표결-압박.md` read OK. |
| Locations facets | ✅ | `장로회당 표결단`, `중정 회화나무`, and `가주전-회랑 접속부` are exact anchors. |
| Length / Prose forecast | ✅ | Sc1 written=1,710; recomputed=1,710; Est=1,700 · Sc2 written=1,890; recomputed=1,890; Est=1,900 · Sc3 written=2,140; recomputed=2,140; Est=2,100 · Sc4 written=1,710; recomputed=1,710; Est=1,700 · scene fields and header addends both total 7,400; Scale 4,000–8,000. |
| Episode List Summary | ✅ | 「명부를 후계 경쟁의 증거로 바꾸고 장로들의 표를 갈라놓는다」→ Scenes 1–2; 「계승 시험을 비틀어 승리한다」→ Scene 3; 「도현은 승리를 축하하지 않고 시험의 진짜 목적만 묻는다」→ Scene 4. |
| Hook to Next / Closing | ✅ | Hook「시험장 아래에서 흑풍루의 비밀 통로가 발견된다」; Out「흑풍루 표식이 있는 지하 통로 입구가 발견된다」; Turn「시험장 아래에 흑풍루 표식이 있는 통로 입구가 실제로 존재한다는 사실이 드러난다」. |
| Hook scope | ✅ | No chase, faction arrival, structural explanation, or second reveal is added. |
| Profile-backed knowledge | ✅ | Scene 4 uses observable seal-direction inference, not unsupported prior location knowledge. |
| Characters / speakers | ✅ | Catalog, On stage, and Dialogue intent align. |
| World / no silent lore | ✅ | Passage identity remains a discovery-level Hook and its unapproved structural relation is held. |
| No cross-scene paste | ✅ | Scene 2’s sealed official record is causally and materially distinct from Scene 1’s declaration. |

## Engagement Checks (Design)

| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | The design asks who created the death vote and whether Jinwoo can reverse it through succession. |
| Personal stake present | ✅ | Losing the successor seat exposes Jinwoo to both elder disposition and the death list. |
| Episode Out hook | ✅ | A concrete marked entrance is discovered, matching the locked Hook without explaining it. |
| Exposition budget respected | ✅ | The design limits new information to document linkage, test authority, and the existence of the marked entrance; destination and purpose are held. |
| Seed discipline | ✅ | Plant/Hints do not contradict the discovery-level Hook; the passage Plant explicitly holds structure and purpose. |
| Scene-first Key Events (all required fields) | ✅ | All four scenes contain complete causal fields. |
| Sensory-emotional on every scene | ✅ | Each scene pairs a setting detail with Jinwoo’s reaction. |
| Motifs planned across scenes | ✅ | Folded documents / reversed tablets and sheath tapping have episode-level and scene-level touches. |
| Overview signature line | N/A | `overview.md` has no designated signature dialogue line. |

## Literary Craft Checks (Design)

| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | Information arrives through contested records, votes, and a test rather than detached exposition. |
| Sensory-emotional pairing | ✅ | Every scene specifies a concrete sensory cue and POV reaction. |
| Dialogue voices + Dialogue intent | ✅ | Jinwoo, the elder representative, Mujin, Taeseok, and Dohyun have differentiated goals and speech registers. |
| Reader-discovered meaning | ✅ | The closing is an image of a marked entrance; thematic meaning remains in Hold rather than dialogue. |
| Antagonist plausibility | ✅ | The elder representative protects institutional authority and factional interest; Dohyun remains an ambiguous evaluator rather than a flat villain. |
| Closing image specified | ✅ | The black mark and first stair are specified without thematic explanation. |

## Literary Awards Juror Checks (Design)

Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)

Locked audience: adult male web-novel readers who prefer regression martial arts, sect-control power fantasies, family reversals, revenge, and cathartic strategic victories.

| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader’s attention | ✅ | The episode begins from the public reveal that Jinwoo himself is the next victim and immediately weaponizes it in a succession confrontation. |
| Personal stake matches what this reader came for | ✅ | The successor seat becomes a visible survival and status mechanism; the vote and test produce a concrete institutional win. |
| Pacing / density fits platform expectations | ✅ | Four distinct scenes and a 7,400-character forecast fit the 4,000–8,000 scale. Scene 2 now delivers a sealed, irreversible vote-record result rather than repeating Scene 1. |
| Out hook makes this reader want the next episode | ✅ | A Black Wind Tower-marked entrance appears beneath the controlled family space, creating a clear investigation target while preserving mystery. |
| No alienation of core audience without overview intent | ✅ | No romance detour, prolonged training, or explanatory reconciliation displaces the revenge / sect-control contract. |

## Design Critique (required personas)

#### Target Reader
- Stance: Adult male regression-martial-arts web-novel reader seeking strategic reversal, institutional power gain, and a concrete next-episode mystery.
- Strengths: The death-list threat becomes a succession advantage; the reader receives an irreversible status win, father–son pressure, and a marked underground hook.
- Defects: —
- Reader impact: The revised version gives a clear victory and a clean mystery without asking the reader to accept an unexplained architectural conclusion.

#### Genre Critic
- Stance: Tests the design against regression martial-arts and sect-power serial promises.
- Strengths: Jinwoo wins through prior information, procedure, and document logic rather than a power-up; the institutional victory has operational weight.
- Defects: Low — the succession test still introduces and resolves its rule in one scene, so Stage ⑥ should make the pre-existing record/authority logic visible before Jinwoo exploits it.
- Reader impact: A brief generation constraint preserves the feeling that Jinwoo is clever rather than authorially rescued by a new rule.

#### Plot Expert
- Stance: Checks causality, Episode List fidelity, Hook body alignment, and Hook scope.
- Strengths: The chain is now clean: public victim reveal → initial vote split → sealed official consequence → test manipulation → uncelebrated recognition → marked entrance discovery. Summary, Out, Seeds, and closing Turn share discovery strength.
- Defects: —
- Reader impact: The episode promises one next-step investigation instead of skipping directly to an answer about the passage’s origin.

#### Reader-Editor
- Stance: Checks serial readability, exposition restraint, scene economy, and closing density.
- Strengths: Scene 1 establishes the political split; Scene 2 fixes the result in an official record; Scene 3 converts that record into test leverage; Scene 4 changes emotional temperature and supplies one hook.
- Defects: —
- Reader impact: The middle now has a visible state change, reducing skim risk between the opening confrontation and the test.

#### Character Critic
- Stance: Checks motivation, relationship pressure, voice, and profile-backed knowledge.
- Strengths: Jinwoo’s document inference matches his profile and drive; Dohyun’s question tests a visible choice rather than asserting impossible foreknowledge; Mujin and Taeseok act from their catalogued survival and record pressures.
- Defects: Low — Stage ⑥ should keep Jinwoo’s answer uncertain and tactical, not turn the inference into a retrospective claim that he knew the location all along.
- Reader impact: Suspicion between father and son remains productive rather than confusing or omniscience-based.

#### Setting/Lore Expert
- Stance: Checks catalog compliance, location anchors, staging, and the credibility of the new underground discovery.
- Strengths: All surface locations and facets are citeable; the shared staging is properly separated from the static location; the passage is now an observed entrance with a Black Wind Tower mark, not a confirmed un-catalogued structural network.
- Defects: —
- Reader impact: The cliffhanger feels like a discovery that demands investigation, not a retroactive rewrite of the house map.

#### Literary Critic
- Stance: Checks motif, image, emotional restraint, and whether craft survives generation.
- Strengths: Folded documents and reversed tablets embody predetermined death versus present choice; the sheath taps restrain Jinwoo’s emotion; the closing stair is visual and non-expository.
- Defects: Low — do not narrate that the successor seal “covers” the death order; show the juxtaposition and let the reader infer the reversal.
- Reader impact: Concrete images reinforce the revenge/control premise without slowing the serial rhythm.

## Design Adjudication

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | Succession-test rule may feel newly convenient if the record/authority logic is not visible before Jinwoo exploits it (Genre Critic) | Low | No | yes | The design is generation-ready, but a brief visible setup will protect the reader’s sense of earned strategy. | Carry the existing record flow and sealed authority condition into prose before the exploit; do not add a new rule or plot beat. | Carry-⑥ |
| 2 | The “successor seal covers the death order” motif could become explanatory symbolism (Literary Critic) | Low | No | yes | The image works best when discovered by the reader rather than explained by narration. | Show document/insignia juxtaposition and Jinwoo’s restrained reaction; do not state the thematic meaning. | Carry-⑥ |

## Design Verdict

| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Required next action

The previous High/Med findings are resolved in the Stage ④ design:

1. The passage is limited to discovery of a Black Wind Tower-marked entrance; its structural relation, destination, and purpose remain held.
2. Dohyun’s question is grounded in Jinwoo’s observable comparison of the vote documents and test-seal direction; no unsupported prior knowledge is asserted.
3. Scene 2 now produces an irreversible sealed official vote record and authority condition distinct from Scene 1’s initial split.
4. Two Low `Carry-⑥` constraints remain for prose generation; they do not block Generation-ready.

G4 remains open because the user has not yet given explicit Episode 007 design approval. No manuscript generation is authorized until the design is explicitly approved.
