# Design Evaluation: Episode 016 — 옛 전장의 물자

> **Target:** `episodes/016-옛-전장의-물자.md`
> **Target Reader:** 회귀·빙의·환생 무협과 문파 장악물, 가족 관계의 반전, 복수형 사이다를 선호하는 성인 남성향 웹소설 독자
> **Evaluation scope:** Stage ⑤ design validation after G4 Architect approval

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 criterion; this is Episode 016 design scope. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Early-arc criterion already belongs to Episodes 001–003. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scenes 1–3 execute a physical supply-line investigation, confirm older trade evidence, and reveal the `서진우` name inside the seal as the next hook. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | The design presents Dohyun’s seal as evidence and preserves the possibility of conflicting motive; it explicitly Holds a benevolent explanation and does not flatten him. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Late-series payoff criterion; this episode Holds the direct connection rather than claiming a payoff. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | Jinwoo uses observed material chronology, not omniscience; he limits conclusions when the seal’s purpose is unknown. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Endgame criterion; this episode advances the family mystery without resolving it. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript-stage quality criterion; forecast is checked under Schema below. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only | ✅ | One complete canonical section for each of Scenes 1–3; no nested alternate scene layout. |
| No skill/workflow dump after the design | ✅ | File contains episode design and its gate evidence only; no copied workflow procedure. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Exactly one heading each for Scenes 1, 2, and 3; each changes location facet and evidence state. |
| Canonical episode path | ✅ | Actual target is `episodes/016-옛-전장의-물자.md`. |
| Field notation | ✅ | Meta fields use `**Field:**`; execution fields use `- **Field:**`. |
| Every scene has required fields | ✅ | All scenes contain POV, Location, When, On stage, Staging, Situation through Est. length, and 7 outline lines. |
| Characters Appearing ↔ On stage union | ✅ | Appearing is exactly 서진우·남궁혁·백무진, the union of all three On stage lists. |
| On stage includes speakers | ✅ | All named speech/action roles in each scene are among the three listed characters. |
| Characters ⊆ catalog | ✅ | `characters.md` and the three exact profiles were read OK. |
| Summary/Hooks cast alignment | ✅ | Summary and hook name no cast beyond the Appearing roster; `도현` is a referenced prior-character name, not an acting speaker. |
| No later-list cast debut | ✅ | All appearing cast are established by Episode 016; no later-list character is introduced. |
| Locations ⊆ Key Locations | ✅ | `옛-전장` maps to the Key Locations row; each named set is a profile facet. |
| Location facets ⊆ Multi-facet anchors | ✅ | `성문`, `병기 창고`, and `봉인석` exactly match `locations/옛-전장.md` Multi-facet anchors. |
| Nested scene files absent | ✅ | Single canonical episode file; no nested episode scene directory. |
| No template residue | ✅ | No raw template placeholders remain. |
| Prose forecast present | ✅ | Each scene has typed five-category units and an integer subtotal. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 `3×250+3×180+2×120+2×140+1×80 = 1890`, Est 1900; Sc2 `4×250+3×180+2×120+2×140+1×80 = 2140`, Est 2100; Sc3 `4×250+2×180+2×120+2×140+1×80 = 1960`, Est 2000. Each product is exact and each Est is within ±20%; 7 outline lines × 200–350 also contains each Est. |
| Dialogue intent vs outline speech | ✅ | Every scene has dialogue intent matching its outline’s named exchanges. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields: `1900 + 2100 + 2000 = 6000`; header addends: `1900 + 2100 + 2000 = 6000`. |
| Est. length sum ≥ Scale min | ✅ | 6000 ≥ 4000. |
| Est. length sum ≤ Scale max | ✅ | 6000 ≤ 8000; central target band. |
| Cited staging/profile paths exist | ✅ | `characters/서진우.md`, `characters/남궁혁.md`, `characters/백무진.md`, and `locations/옛-전장.md` were read OK this turn; Staging is `none` for every scene, so no staging file is required. |
| Episode List plot | ✅ | Series Summary clauses map to Scene 2’s blackwind weapons/older trade evidence and Scene 3’s Dohyun seal; no different story is substituted. |
| Hook evidence strength | ✅ | Series Hook, Summary, Out, Arc close, Seed, and Scene 3 Turn all state that the name inside the seal is `서진우`, not Dohyun; no surface softens or hardens it. |
| Hook scope | ✅ | Out has one reveal obligation; no unplanned chase, faction arrival, or second independent reveal is added. |
| No design-paste / meta-only scenes | ✅ | Each scene performs a distinct event: arrival/route, weapon chronology, seal-name reveal. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | Arrival → material proof → personal seal reveal is a causal escalation. |
| Scene transitions chain | ✅ | Scene 1 sends the party to the weapon storehouse; Scene 2 sends them to the seal stone; Scene 3 exits into evidence preservation and the next pursuit. |
| Scene sections complete | ✅ | Every Scene Index row has one complete Key Events section. |
| Generation Readiness | ✅ | All Schema rows pass; no Pending adjudication item blocks prose generation. |
| Beat concreteness | ✅ | Wheels, drag marks, weapon wear, seals, ink, and stone layers are observable actions/evidence. |
| Est. length budget | ✅ | Independent products and 6000-character sum pass the 4000–8000 Scale. |
| Prose forecast quality | ✅ | Dialogue/action/sensory/POV/transition counts correspond to each outline and scene function. |
| Episode List scope aligned | ✅ | The design executes the approved Summary and exact Hook without Out creep. |
| Prior hook addressed | ✅ | Scene 1 immediately reaches the `옛 전장` destination and preserves the joint pursuit. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected | ✅ | Phase A, Phase B, continuity set, facet preflight, and no-staging decision are recorded in Prior Design Alignment. |
| Series / overview tone & arc honored | ✅ | Cold, event-driven P1 investigation with family-revenge pressure; no slow training or romance focus. |
| Episode List Summary / Hook honored | ✅ | Exact Summary and Hook clauses are mapped in the design gate evidence. |
| Hook internal consistency | ✅ | Summary, Arc close, Out, Seed, and Scene 3 Turn agree on the same name reveal. |
| Characters from architecture; profiles not redefined | ✅ | Existing drives, voices, and equipment states are cited rather than rewritten. |
| Profile-backed knowledge / recognition | ✅ | No unauthorized character recognition is asserted; Jinwoo observes resemblance in seal pressure but refuses identity certainty. |
| Locations from architecture; profiles not redefined | ✅ | Only `옛-전장` and its exact anchors are used. |
| Location profile paths readable | ✅ | `locations/옛-전장.md` read OK this turn. |
| Location facets ⊆ anchors | ✅ | All three scene facets are exact anchors from the loaded profile. |
| Stagings | ✅ | All scenes are separate facet investigations with `Staging: none`; no fixed multi-scene blocking is silently invented. |
| World rules / history consistent | ✅ | Material record comparison follows the established document/transport logic; no new bloodline or seal rule is asserted. |
| No improvised entities or silent lore | ✅ | No new character, faction, location, or world rule; the seal’s function remains Hold. |
| Continuity files used | ✅ | Immediate prior summary and `story-so-far.md` are the stated authority. |
| Character/location state vs story-so-far | ✅ | Jinwoo, Hyeok, and Mujin remain conditional joint trackers; 옛 전장은 newly reached, not previously occupied. |
| Unresolved threads mapped | ✅ | TH-020 and TH-021 are picked up; TH-018, TH-009, and TH-002 advance; family/poison threads are explicitly held. |
| No contradiction of released continuity | ✅ | Episode 015’s destination, alliance condition, and faint mark are preserved. |
| Conflicts section | ✅ | Empty; facet preflight passed and no higher-layer change is needed. |

## Design Consistency Gate
- **Locations index:** ✅ — all scenes use `옛-전장`, a Key Locations slug; facets are `성문`, `병기 창고`, `봉인석`.
- **Cited paths:** ✅ — `locations/옛-전장.md` and all three appearing character profiles were read OK this turn; no staging path is cited because every scene uses `none`.
- **Location facets:** ✅ — all three facets are listed under `locations/옛-전장.md` → `Multi-facet anchors`.
- **Length / Forecast:** ✅ — Sc1 written=1890; recomputed=1890; Est=1900 · Sc2 written=2140; recomputed=2140; Est=2100 · Sc3 written=1960; recomputed=1960; Est=2000 · scene fields and header both total 6000.
- **Episode List Summary:** ✅ — 「흑풍루의 병기와 도현이 남긴 봉인 흔적을 발견」→ Scenes 2–3; 「거래한 시점이 예상보다 오래됐음」→ Scene 2 Turn.
- **Hook:** ✅ — Series Hook「봉인 안쪽에서 도현의 이름이 아닌 진우의 이름이 나온다」; Summary/Out「봉인 안쪽에서 도현의 이름이 아닌 진우의 이름이 나온다」; Scene 3 Turn「봉인 안쪽에서 도현의 이름이 아닌 `서진우`라는 이름이 나타난다」. Same strength, one obligation.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | The design asks who moved the material, when, and why Dohyun’s seal is present. |
| Personal stake present | ✅ | Jinwoo’s revenge certainty is threatened by an older trace tied to his own name; Hyeok’s family-name recovery remains active. |
| Episode Out hook | ✅ | The personal name inside the seal is concrete, singular, and directly actionable. |
| Exposition budget respected | ✅ | Terrain and lore appear only while reading tracks, weapons, and seal layers. |
| Seed discipline | ✅ | One Plant, one Hint, and one Hook are bounded; seal function and transaction purpose remain Hold. |
| Scene-first Key Events | ✅ | Every scene begins with an observable situation and contains a causal Turn. |
| Sensory-emotional on every scene | ✅ | Wind/ore, rust/ink, and stone dust/pressure each trigger Jinwoo’s inference. |
| Motifs planned | ✅ | `녹과 새 먹` and `안쪽의 이름` have scene-level touch points. |
| Overview signature line | N/A | `overview.md` has no signature dialogue line requiring placement. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | Each discovery is contested by route safety, evidence preservation, or interpretive restraint; no catalog lecture. |
| Sensory-emotional pairing | ✅ | Material details directly alter Jinwoo’s chronology judgment. |
| Dialogue voices + intent | ✅ | Jinwoo limits claims, Hyeok presses for direct action, and Mujin reports practical evidence. |
| Reader-discovered meaning | ✅ | The design lets the reader infer that Jinwoo is inside the old record without explaining why. |
| Antagonist plausibility | ✅ | The absent transporter and mixed old/new supply chain imply an indirect operation without a caricatured villain. |
| Closing image specified | ✅ | The inner `서진우` lettering and old red pressure mark close the episode. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | The episode opens at the destination with a live supply trace, not travel exposition. |
| Personal stake matches what this reader came for | ✅ | Tactical evidence threatens Jinwoo’s revenge model and ties the mystery to his own name. |
| Pacing / density fits platform expectations | ✅ | Three escalating investigative scenes, 6000-character forecast, and no training detour fit serialized martial-arts rhythm. |
| Out hook makes this reader want the next episode | ✅ | A protagonist-name reveal inside Dohyun’s seal is a sharp, concrete click-forward. |
| No alienation of core audience without overview intent | ✅ | Action, rivalry, supply-line advantage, family mystery, and restrained cliffhanger remain central. |

## Design Critique (required personas)

#### Target Reader
- **Stance:** Adult male-oriented regression martial-arts reader seeking tactical advantage, rivalry pressure, family mystery, and a concrete next click.
- **Strengths:** The destination is immediate, evidence is physical, and the final name reveal personalizes the old battlefield without explaining away the mystery.
- **Defects:** —
- **Reader impact:** The reader gets a tactical discovery and a stronger personal question rather than a travel interlude.

#### Genre Critic
- **Stance:** Tests action-investigation payoff, regression advantage, rivalry-to-alliance momentum, and wuxia serial promises.
- **Strengths:** Blackwind weapons, chronology proof, and a seal reveal provide escalating genre beats; Jinwoo’s prior knowledge is used as a method, not a prophecy.
- **Defects:** —
- **Reader impact:** The episode delivers a satisfying evidence reversal while preserving the revenge engine.

#### Plot Expert
- **Stance:** Checks causal sequence, Summary/Hook alignment, seed timing, and Out scope.
- **Strengths:** Each transition is causal; the exact Series Hook is executed in Scene 3; the Out adds no chase or unrelated faction arrival.
- **Defects:** —
- **Reader impact:** The reader can track what was found, what it proves, and what remains unknown without a continuity reset.

#### Reader-Editor
- **Stance:** Checks serialization rhythm, exposition restraint, closing sell, and scene density.
- **Strengths:** Arrival, evidence, and name reveal form a clean three-step ladder; each scene has a distinct facet and turn.
- **Defects:** —
- **Reader impact:** The episode should read quickly despite investigation because every material detail changes the pursuit.

#### Literary Critic
- **Stance:** Checks motifs, image-based closure, restraint, and reader-discovered meaning.
- **Strengths:** Rust/new ink creates a tactile time motif; the inside-name image gives the episode a non-expository ending.
- **Defects:** —
- **Reader impact:** The cliffhanger has emotional charge without a thematic monologue or premature explanation.

#### Character Critic
- **Stance:** Checks motivation, distinct voices, alliance pressure, and knowledge limits.
- **Strengths:** Jinwoo’s refusal to overclaim fits his evidence-driven behavior; Hyeok and Mujin have distinct practical roles; no one is granted knowledge unsupported by profile.
- **Defects:** —
- **Reader impact:** The alliance remains conditional and active rather than becoming instant trust.

#### Setting/Lore Expert
- **Stance:** Checks location facet integrity, material evidence, and unauthorized seal/world-rule invention.
- **Strengths:** All sets are exact `옛-전장` anchors; weapons and seal are learned through handling and comparison; the seal’s function stays held.
- **Defects:** —
- **Reader impact:** The world feels rule-bound and ancient without stopping for a system explanation.

## Design Adjudication
| # | Finding | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---------|----------|-----------|--------|----------------------------------|--------------|--------|
| — | No High or Med design finding. All required schema, continuity, hook, reader, and craft checks pass. | — | No | no | The locked reader receives an immediate material investigation, an older-trade reversal, and a single personal cliffhanger. | — | Skip |

**Adjudication result:** No Pending or Carry-⑥ items remain. The design is generation-ready as written.

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5 (Design Eval Approval)
- **Status:** ✅ Architect-approved
- **Evidence:** Schema, independent forecast arithmetic, location path/facet checks, continuity mapping, required persona critique, and hook-scope review all pass.
- **Next:** Stage ⑥ — generate the Episode 016 manuscript from the locked design.
