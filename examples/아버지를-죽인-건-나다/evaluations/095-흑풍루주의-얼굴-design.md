# Design Evaluation: Episode 095 — 흑풍루주의 얼굴

## Evaluation Scope
- **Target Reader:** 회귀·빙의·환생 무협, 문파 장악물, 가족 반전과 복수형 사이다를 선호하는 성인 남성향 웹소설 독자.
- **Design artifact:** `episodes/095-흑풍루주의-얼굴.md` — canonical path read in this turn.
- **Gate context:** G4 Architect Approved. This evaluation is mandatory before Stage ⑥.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 / series-opening criterion; not this episode’s design scope. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Early-arc placement criterion; already outside Episode 095 scope. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 금고 대면·기록 검증·탈출 대립이 Scenes 1–4의 실질 사건이며, 도현의 선택 기록/주장이 Episode 096 후크로 남는다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1 maintenance criterion; Episode 095 is P3 and does not reframe the P1 section. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series-level late payoff criterion; this episode uses 가환 as a witness but is not the final recovery. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 회귀 전 죽음의 기억을 흑풍루주의 증거와 대조하지만, 기억을 진실의 자동 증명으로 쓰지 않고 주장을 보류한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Endgame criterion; this episode escalates rather than resolves the family conflict. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript Stage ⑥/⑦ criterion; forecast is checked in Schema below. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | Four unique `### Scene n` sections use the required meta lines and flat bullet fields. |
| No skill/workflow dump after the design | ✅ | Episode body contains narrative design only; no copied workflow procedure. |
| Unique Scene headings; no pasted twin scenes | ✅ | Scenes 1–4 have distinct functions: face, relation, claim, escape. |
| Canonical episode path | ✅ | `episodes/095-흑풍루주의-얼굴.md` exactly matches the Episode List title slug. |
| Field notation | ✅ | Required fields use `**Field:**` and `- **Field:**` notation. |
| Every scene has required fields | ✅ | POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, outline, Unit budget, and one Est. are present in every scene. |
| Characters Appearing ↔ On stage union | ✅ | Appearing is exactly the union of the five names on stage across Scenes 1–4; no ghost cast. |
| On stage includes speakers | ✅ | 흑풍루주, 진우, 혁, 가환, and the 제자 are explicitly on stage wherever their dialogue intent or named action appears. |
| Characters ⊆ `characters.md` | ✅ | All five names have catalog rows and this-turn readable profile paths. |
| Summary/Hooks cast alignment | ✅ | Summary, seeds, closing, and hooks use only the five appearing characters or 도현 as a record subject, not an on-stage speaker. |
| No later-list cast debut | ✅ | No cast first introduced after Episode 095 is added. |
| Locations ⊆ Key Locations | ✅ | `흑풍루-본거지` is a Key Locations slug in all four scenes. |
| Location facets ⊆ Multi-facet anchors | ✅ | `금고`, `의식장`, and `탈출 수로` exactly match `locations/흑풍루-본거지.md` anchors. |
| Nested scene files absent | ✅ | One canonical episode file only. |
| No template residue | ✅ | No unresolved instruction braces remain. |
| Prose forecast present | ✅ | All four scenes use five allowed unit types with integer products. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 3×250+2×180+2×120+2×140+1×80=1,630; Sc2 same=1,630; Sc3 4×250+3×180+2×120+2×140+1×80=2,050; Sc4 3×250+3×180+2×120+2×140+1×80=1,800. Written products equal recomputation; Est. 1,600+1,600+2,000+1,800=7,000 and each is within the outline density band. |
| Dialogue intent vs outline speech | ✅ | Each scene’s outline includes speech only where Dialogue intent assigns an on-stage speaker. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields: 1,600 + 1,600 + 2,000 + 1,800 = 7,000. Header addends: 1,600 + 1,600 + 2,000 + 1,800 = 7,000. |
| Est. length sum ≥ Scale min | ✅ | 7,000 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 7,000 ≤ 8,000; central-band estimate leaves generation headroom. |
| Cited staging/profile paths exist | ✅ | This-turn `read_files` succeeded for `characters/서진우.md`, `남궁혁.md`, `가환.md`, `진우의-스승의-제자.md`, `흑풍루주.md`, `locations/흑풍루-본거지.md`, `world/혈맥계약과-약그릇.md`, and `stagings/095-금고-대면.md`. |
| Episode List plot | ✅ | Series Summary clauses map respectively to Scenes 1, 2, 3, and 4 without stealing Episode 096’s full payoff. |
| Hook evidence strength (internal) | ✅ | Series Hook「회귀를 만든 것은 흑풍루주의 술법이 아니라 도현의 선택」; Summary/Out preserve the same choice claim as a record and antagonist assertion; Scene 3 Turn names 도현 in the 선택자란; Scene 4 leaves the record behind while refusing full causal confirmation. |
| Hook scope | ✅ | Out has one central obligation (도현 선택 기록/주장) plus the supporting escape consequence; no unrelated faction arrival or second reveal. |
| No design-paste / meta-only scenes | ✅ | Every scene contains a concrete action and turn; no scene exists only to announce the next episode. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | Face → relationship proof → experiment claim → evidence-preserving escape is causal and escalating. |
| Scene transitions chain | ✅ | Scene 1 opens the door to Scene 2’s record comparison; Scene 2 yields the contract fragment; Scene 3 yields the lock gap; Scene 4 resolves into the waterway. |
| Scene sections complete | ✅ | All four Scene Index rows have full generation briefs. |
| Generation Readiness | ✅ | No Schema or arc blocker; G4 evidence and required staging/path checks are complete. |
| Beat concreteness | ✅ | Named actions include entering, comparing seals, presenting a contract fragment, opening the lock gap, and escaping. |
| Est. length budget | ✅ | Independently recomputed 7,000-character forecast is within 4,000–8,000. |
| Prose forecast quality | ✅ | Unit types correspond to planned dialogue, movement, sensory detail, POV separation, and transitions. |
| Episode List scope aligned | ✅ | Face, same master, experiment claim, and escape all execute this row; full regression cost remains held. |
| Prior hook addressed | ✅ | Scene 1 directly resolves the prior call from outside the sealed vault. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected | ✅ | Phase A, Phase B, staging, and the immediate prior continuity set are explicitly recorded. |
| Series / overview tone & arc honored | ✅ | Closed-space investigation, family stakes, and antagonist pressure fit P3 and the locked audience. |
| Episode List Summary / Hook honored | ✅ | Summary and Hook are quoted in the design’s Gate Evidence and each clause has a named scene. |
| Hook internal consistency | ✅ | Summary, Arc close, Out, Seed, Scene 3 Turn, and Scene 4 transition consistently preserve the 도현 선택 record/claim without resolving its full cost. |
| Characters from architecture | ✅ | Profiles are cited, read, and not redefined. |
| Profile-backed knowledge / recognition | ✅ | 흑풍루주의 recognition of 진우 and 도현 is supported by the catalog relationships and the established contract/experiment role; the design still treats his causal interpretation as contested. |
| Locations from architecture | ✅ | Only the catalogued 흑풍루 본거지 and exact anchors are used. |
| Location profile paths readable | ✅ | Exact location and staging paths were read successfully this turn. |
| Location facets ⊆ anchors | ✅ | All three used facets are listed under Multi-facet anchors. |
| Stagings from episode design | ✅ | New situation staging is authored at Stage ④ and binds existing character states without inventing wardrobe or gear. |
| World rules / history consistent | ✅ | The design treats 회귀 as costly and leaves its actual memory/lifespan payment unresolved. |
| No improvised entities or silent lore | ✅ | No new named entity, faction, rule, or uncatalogued set is introduced. |
| Continuity files used | ✅ | Immediate prior summary and story-so-far are the only continuity authorities. |
| Character/location state vs story-so-far | ✅ | Jinwoo retains elixir, seal proof, and 봉인-해제 state; the vault’s prior closure and staging carry forward. |
| Unresolved threads handled | ✅ | TH-126/135/136/137/138/139/140 are picked up or advanced; regression cost and elixir result are explicitly held. |
| No contradiction of released continuity | ✅ | The episode does not claim elixir administration or mother contact before release. |
| Conflicts section empty or escalated | ✅ | No source conflict was found; Episode 096’s full cost explanation is explicitly held. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — selective Phase A/Phase B load complete; no later continuity used.
- Locations: ✅ — index `흑풍루-본거지`; path `locations/흑풍루-본거지.md` and staging path read OK; facets `금고`, `의식장`, `탈출 수로` ⊆ anchors.
- Length / Prose forecast: ✅ — scene fields and header both recompute to 7,000; every written Unit product equals independent arithmetic.
- Episode List Summary: ✅ — face → Scene 1; same master → Scene 2; experiment claim → Scene 3; elixir-preserving escape → Scene 4.
- Hook to Next / Closing: ✅ — series Hook「진우의 회귀를 만든 것은 흑풍루주의 술법이 아니라 도현의 선택이었다」; design Out「도현의 선택이었다는 기록과 주장」; Scene 3 Turn「선택자란에는 도현의 이름」; closing image leaves the record behind without adding a second reveal.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | Why does the man outside know Jinwoo’s name and what price authorized Dohyun’s extraction? |
| Personal stake present | ✅ | The elixir is Dohyun’s remaining chance; accepting the Lord’s explanation would also relocate Jinwoo’s guilt. |
| Episode Out hook | ✅ | The record and claim force the reader into Episode 096’s regression-cost verification. |
| Exposition budget respected | ✅ | Relationship and lore arrive through seals, a fragment, and conflict; full regression mechanics are held. |
| Seed discipline | ✅ | One Plant (choice record) and one Hint (lock gap) are concrete and limited. |
| Scene-first Key Events | ✅ | All scenes carry complete required fields. |
| Sensory-emotional on every scene | ✅ | Each scene pairs a physical vault/ritual/water detail with Jinwoo’s POV reaction. |
| Motifs planned | ✅ | Face/name and three nail taps recur with scene placements. |
| Overview signature line | N/A | Overview contains no mandated signature dialogue line. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | Information is earned through an active standoff and escape deadline rather than a lecture. |
| Sensory-emotional pairing | ✅ | Seals, iron smell, red light, and water silence are tied to Jinwoo’s choices. |
| Dialogue voices + intent | ✅ | Lord’s polite educator tone, Jinwoo’s clipped verification, Hyuk’s direct pressure, and G환’s limited formal reading are distinct. |
| Reader-discovered meaning | ✅ | The reader sees a protective record being reframed as ownership/experiment; no thematic explanation is assigned to the ending. |
| Antagonist plausibility | ✅ | Lord’s worldview treats people as testable contract outcomes, consistent with his profile and institutional role. |
| Closing image specified | ✅ | Red reflection of the elixir on black water and the closing vault door. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | The awaited antagonist appears on page-one pressure and demands the immediate objective. |
| Personal stake matches reader expectation | ✅ | Father’s survival and the meaning of Jinwoo’s murder are active simultaneously. |
| Pacing / density fits platform expectations | ✅ | Four compact, escalating scenes and 7,000-character headroom favor event density over explanation. |
| Out hook makes this reader want the next episode | ✅ | A concrete name in the choice record challenges the reader to learn whether Dohyun paid the regression cost. |
| No alienation of core audience | ✅ | The design preserves revenge, tactical escape, family reversal, and a direct antagonist confrontation. |

## Design Critique (required personas)

#### Target Reader
- Stance: A serial reader seeking a direct villain reveal, tactical payoff, and a sharper father mystery.
- Strengths: The face appears immediately; the elixir remains physically at stake; the record gives the next episode a concrete question.
- Defects: —
- Reader impact: Strong continuation pressure without resolving the promised regression-cost mystery too early.

#### Genre Critic
- Stance: Checks the 회귀 무협 and revenge-webnovel contract.
- Strengths: Antagonist reveal, evidence reversal, confrontation, and escape all deliver genre movement.
- Defects: —
- Reader impact: The episode provides a satisfying confrontation while reserving the next escalation.

#### Plot Expert
- Stance: Checks causality, Hook body alignment, and Out scope.
- Strengths: The prior sealed-vault call directly causes the face reveal; every scene produces the next action; the Out stays within the Series Hook.
- Defects: —
- Reader impact: The reader can distinguish what is proven (face, relationship, record fragment) from what is claimed (experiment and full regression causality).

#### Reader-Editor
- Stance: Checks serialization, exposition restraint, and closing density.
- Strengths: The last scene has one main obligation—escape with the elixir—while the record remains the single forward hook.
- Defects: —
- Reader impact: No crowded closing stack or explanation loop should dilute the page-turn.

#### Literary Critic
- Stance: Checks motif and emotional design beyond mechanics.
- Strengths: Face/name and ownership motifs convert identity into a concrete record problem; the closing image avoids thematic speech.
- Defects: —
- Reader impact: The family reversal remains felt through the object and choice rather than narrated as a conclusion.

#### Character Critic
- Stance: Checks motivation, voice, and profile-backed recognition.
- Strengths: Jinwoo verifies instead of surrendering to the antagonist’s interpretation; the Lord’s educator voice and Dohyun-centered grievance are profile-consistent; all recognitions have relationship support.
- Defects: —
- Reader impact: The confrontation deepens both the father conflict and the villain’s personal menace without making Jinwoo passive.

#### Setting/Lore Expert
- Stance: Checks blood-contract rules, citeable facets, and exposition delivery.
- Strengths: `금고`, `의식장`, and `탈출 수로` are catalog-backed; blood-seal reactions remain costly rather than magical convenience; the new staging cites existing character states.
- Defects: —
- Reader impact: Late-arc lore arrives through a contested object reading, preserving clarity for readers who want action first.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| — | No High/Med design defects. All required schema, continuity, path, facet, forecast, cast, and Hook checks pass. | — | No | — | The locked reader receives the promised confrontation and a concrete Episode 096 question without premature regression-cost confirmation. | — | Applied-④ |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5 — Architect Approval
- **Status:** Approved by Architect — 2025-02-14.
- **Rationale:** Required design critics found no High/Med defect; the artifact passes schema, exact forecast arithmetic, continuity, catalog path/facet, staging, and Hook checks. The record/claim distinction preserves Episode 096’s regression-cost reveal without weakening the Episode 095 promise.
- **Next:** Stage ⑥ — manuscript generation, with Hold constraints preserved.
