# Design Evaluation: Episode 066 — 서명

## Evaluation Scope
- Target artifact: `episodes/066-서명.md`
- Stage: ⑤ — design evaluation after Architect G4
- Target Reader: 회귀·빙의·환생 무협, 문파 장악물, 가족 반전과 복수형 사이다를 선호하는 성인 남성향 웹소설 독자 (`overview.md`)
- Authority set: `overview.md`, `series.md`, `world-bible.md`, `world/혈맥계약과-약그릇.md`, `continuity/065-흑풍루주의-친필-summary.md`, `continuity/story-so-far.md`

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 design-scope criterion; Episode 066 does not re-establish the series inciting event. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 001–003 early-arc criterion; already released and outside this episode’s scope. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 추적: 본가 인계 기록 역추적과 계약 원본 발견이 Scene 2–4의 concrete event다. 다음 회차 후크는 원본의 「진우의 피로 찍힌 조항」이다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1-only criterion; Episode 066 is in P2. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series-level later-payoff criterion; this episode may advance TH-105 but cannot claim the full recovery. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 회귀 기억을 정답으로 사용하지 않고 065의 서명·현재 인계 기록·봉랍을 직접 대조한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | ✅ | Episode 066은 부자 대면의 책임 충돌을 진우의 검증 기준으로 전진시키며, 결말의 흑풍루주 응징은 선취하지 않는다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript Stage ⑥/⑦ criterion; design forecast is checked separately under Schema. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | Four scenes use the required meta lines and flat bullet fields in the fixed order. |
| No skill/workflow dump after the design | ✅ | File contains episode design and short gate records only; no pasted workflow procedure. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Unique Scene 1–4 headings; each has a distinct location/function/turn. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | Actual path is `episodes/066-서명.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | Required bold-colon notation is used throughout. |
| Every scene has required meta + bullet fields | ✅ | All four scenes include POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, outline, Unit budget, and exactly one Est. length. |
| Characters Appearing ↔ On stage union | ✅ | Appearing roster is exactly the union of the four names present On stage in Scenes 1–4. |
| On stage includes speakers | ✅ | Named speech/intent is limited to the four On-stage characters in each scene; no off-stage or anonymous speaker is introduced. |
| Characters ⊆ `characters.md` | ✅ | `서진우`, `서도현`, `남궁혁`, `봉인된-또-다른-아이` each map to catalog rows and readable profiles. |
| Summary/Hooks cast alignment | ✅ | Summary, In/Out, Seeds, Closing mention only catalogued appearing cast; no orphan name. |
| No later-list cast debut | ✅ | All four characters are established before Episode 066. |
| Locations ⊆ `locations.md` Key Locations | ✅ | Every scene uses slug `북문서가-본가`, which maps to the Key Locations row. |
| Location facets ⊆ Multi-facet anchors | ✅ | `지하 보관실`, `가주전-회랑 접속부`, `지하 통로 입구` are exact labels in `locations/북문서가-본가.md` Multi-facet anchors. |
| Nested `episodes/{slug}/` scene files absent | ✅ | Single canonical episode file; no nested scene files. |
| No template residue | ✅ | No raw instructional braces or unfinished template fields remain. |
| Prose forecast present (outline + typed units) | ✅ | Every scene has 6–7 paragraph intents and five-type integer unit formulas. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 `3×240+2×180+2×120+2×140+1×80=1,680`; Sc2 `3×250+3×180+2×120+2×140+1×80=1,890`; Sc3 `4×240+3×180+2×120+2×140+1×80=2,100`; Sc4 `3×250+3×180+2×120+2×140+1×80=1,890`. Each written product matches recomputation; Est. values are within ±20% and outline-density bands. |
| Dialogue intent vs outline speech | ✅ | Each scene’s outline includes speech-bearing beats covered by non-`none` Dialogue intent and On-stage speakers. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene Est. fields: `1,700+1,900+2,100+1,900=7,600`; header addends: `1,700+1,900+2,100+1,900=7,600`. |
| Est. length sum ≥ Scale min | ✅ | Recomputed sum 7,600 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | Recomputed sum 7,600 ≤ 8,000. Central band is satisfied, with 400 characters of headroom. |
| Cited staging/profile paths exist | ✅ | This-turn reads succeeded for `characters/서진우.md`, `characters/서도현.md`, `characters/남궁혁.md`, `characters/봉인된-또-다른-아이.md`, `locations/북문서가-본가.md`, `world/혈맥계약과-약그릇.md`, and `stagings/066-계약-원본-추적.md`. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary clauses map to Scene 1 (signature demand and protection-contract claim), Scenes 2–4 (original contract search), and Scene 4 (blood-marked clause). |
| Hook evidence strength (internal) | ✅ | Series Hook: 「계약의 원본에는 진우의 피로 찍은 조항이 있다」; Summary: 「원본에는 진우의 피로 찍은 조항이 남아 있음을 확인」; Out: 「계약 원본에는 진우의 피로 찍힌 조항이 남아 있다」; Scene 4 Turn confirms the blood clause; Seed Plant and closing image show the same confirmation. |
| Hook scope (no Out creep) | ✅ | Out confirms one obligation—the blood-marked clause. The blood-seal removal method is a supporting next-step investigation, not a second reveal or faction arrival. |
| No design-paste / meta-only scenes | ✅ | Each scene has a concrete document action and changed evidence state; no identical Beat or Unit line is repeated. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | Demand → record reversal → responsibility test → original discovery is a causal four-step escalation. |
| Scene transitions chain | ✅ | Scene 1 sends the team to transfer records; Scene 2’s return mark sends them underground; Scene 3’s scratched line sends them to the hidden case; Scene 4 closes on the blood-seal trace. |
| Scene sections complete | ✅ | Every Scene Index row has a full generation brief. |
| Generation Readiness | ✅ | All schema, path, facet, length, cast, continuity, Hook, and staging checks pass; no unresolved design-field adjudication blocks handoff. |
| Beat concreteness | ✅ | Beats use observable acts: comparing strokes, reading transfer marks, finding a blank return field, opening the case, and exposing the blood clause. |
| Est. length budget | ✅ | 7,600 total inside 4,000–8,000; typed formulas independently recomputed. |
| Prose forecast quality | ✅ | Unit types track the actual action, dialogue, sensory, POV, and transition work in each outline. |
| Episode List scope aligned | ✅ | The episode confirms the signature’s limited meaning and blood clause without stealing 067’s blood-seal removal payoff. |
| Prior hook addressed | ✅ | Scene 1 begins directly with the 065 signature and burned-clause evidence. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A indexes, Phase B appearing/used profiles, 065 summary, story-so-far, and new staging are recorded. |
| Series / overview tone & arc honored | ✅ | Cold, event-centered investigation preserves P2’s revenge × filial-responsibility conflict. |
| Episode List Summary / Hook to Next honored | ✅ | Summary and Hook are quoted and mapped without strength drift. |
| Hook internal consistency | ✅ | Summary, Arc close, Out, Seed, Scene 4 Turn, Dialogue intent, and Transition all retain the blood-clause confirmation. |
| Characters from architecture; profiles not redefined | ✅ | Character behavior applies existing drives/voices; no new identity, outfit, or history is invented. |
| Profile-backed knowledge / recognition | ✅ | No unsupported identity reveal; characters inspect documents and physical marks available in the loaded profiles/world rules. |
| Locations from architecture; profiles not redefined | ✅ | Only the catalogued North Gate family compound is used. |
| Location profile paths readable | ✅ | Exact location and staging paths were read successfully this turn. |
| Location facets ⊆ Multi-facet anchors | ✅ | All three cited facets exactly match the location profile. |
| Stagings from episode design; blocking not redefined | ✅ | New staging is authored at Stage ④, cites existing character states, and fixes props/blocking across the four scenes. |
| World rules / history consistent with bible | ✅ | Blood as contract medium, blood-seal trace, and danger of contract violation stay within the world bible and blood-contract aspect. |
| No improvised entities or silent lore | ✅ | No new faction, place, named character, or rule is introduced; hidden storage is an additive situational discovery inside an approved location facet. |
| Continuity files used (ep 002+) | ✅ | Immediate 065 summary and story-so-far are explicitly cited and reflected. |
| Character/location state vs `story-so-far` | ✅ | Jinwoo remains responsible investigator, Dohyun remains the silent signatory with exposed blood illness, Hyuk remains procedural verifier, and the child remains guarded. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-105 advances; blood-seal method is planted; TH-101/102 and contract release are held. |
| No contradiction of released continuity | ✅ | 065’s signature and burned clause are not overwritten; the original is found as the next evidentiary layer. |
| Conflicts section empty or escalated | ✅ | No architecture or continuity conflict was found. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — selective Phase A/Phase B load completed; all cited paths read this turn.
- Locations (index): ✅ — every scene uses `북문서가-본가` from the Key Locations table.
- Locations (path): ✅ — exact paths `locations/북문서가-본가.md` and `stagings/066-계약-원본-추적.md` read OK.
- Locations (facets): ✅ — `지하 보관실`, `가주전-회랑 접속부`, `지하 통로 입구` ⊆ the location profile’s Multi-facet anchors.
- Length / Prose forecast: ✅ — written/recomputed products are `1,680/1,680`, `1,890/1,890`, `2,100/2,100`, `1,890/1,890`; scene fields and header both total `7,600`.
- Episode List Summary: ✅ — 「도현의 서명을 들이밀며 진실을 요구」→Scene 1; 「진우를 살리기 위한 계약」→Scenes 1–3; 「계약 원본」→Scenes 2–4; 「진우의 피로 찍힌 조항」→Scene 4.
- Hook to Next / Closing: ✅ — series Hook「계약의 원본에는 진우의 피로 찍은 조항이 있다」; Summary「원본에는 진우의 피로 찍은 조항이 남아 있음을 확인」; Out「계약 원본에는 진우의 피로 찍힌 조항이 남아 있다」; Scene 4 Turn confirms the clause at the same evidence strength.
- Hook scope: ✅ — no chase, faction arrival, or second reveal is added to the Out.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | 「도현은 왜 흑풍루주의 지시서에 직접 서명했는가?」 is stated and delayed until evidence accumulates. |
| Personal stake present | ✅ | Believing or rejecting the protection claim changes Jinwoo’s ability to understand his own survival and responsibility. |
| Episode Out hook | ✅ | The blood-marked clause is a concrete, reader-visible next investigation. |
| Exposition budget respected | ✅ | Contract rules are revealed only through the signature, transfer marks, and original; full conditions and release method remain held. |
| Seed discipline | ✅ | Two Plants and one Hint have explicit scenes; Holds are named and not executed. |
| Scene-first Key Events | ✅ | All scenes contain concrete situation, causal Beat, Turn, and transition. |
| Sensory-emotional on every scene | ✅ | Each scene pairs paper/hand, oil/seal, wet stone, or blood/iron with Jinwoo’s POV response. |
| Motifs planned across scenes | ✅ | Signature start-point and red seal thread through the evidence sequence. |
| Overview signature line | N/A | No separate overview signature dialogue line requires use in this episode. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | Every rule fragment is attached to a pressure action: confrontation, record reversal, blocked access, or hidden-case opening. |
| Sensory-emotional pairing | ✅ | The sensory cues always alter what Jinwoo notices or decides. |
| Dialogue voices + Dialogue intent | ✅ | Jinwoo is clipped and evidentiary, Dohyun low and withholding, Hyuk procedural, and the child boundary-focused. |
| Reader-discovered meaning | ✅ | The design asks the reader to hold protection and control together; it does not declare Dohyun innocent. |
| Antagonist plausibility | ✅ | The unseen Black-Wind Tower remains a system of contract and records rather than a convenient visible villain in this episode. |
| Closing image specified | ✅ | Red blood residue on the original clause under lamplight gives the episode a concrete ending image. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | It opens on the high-value 065 signature and forces an immediate father-son answer. |
| Personal stake matches what this reader came for | ✅ | The investigation directly concerns the protagonist’s murder, survival, and revenge target—not detached lore. |
| Pacing / density fits platform expectations | ✅ | Four short causal investigation stages escalate to a tangible original-document reveal within a 7,600-character forecast. |
| Out hook makes this reader want the next episode | ✅ | A blood-marked childhood clause creates a specific, actionable contract mystery for 067. |
| No alienation of core audience without overview intent | ✅ | No romance-centered detour, training filler, or moral absolution is introduced. |

## Design Critique (required personas)
#### Target Reader
- Stance: Adult male-oriented regression/murim reader seeking revenge momentum, document reversals, and father-son tension.
- Strengths: The 065 signature is paid forward immediately; the evidence chain turns a moral question into a concrete search; the blood clause is a strong serial Out.
- Defects: —
- Reader impact: Clear forward motion and a tangible next mystery support retention.

#### Genre Critic
- Stance: Tests the episode against regression-murim and family-revenge serial promises.
- Strengths: Jinwoo uses foreknowledge as investigative leverage rather than omniscience; each scene produces a new evidence state; the father’s limited admission preserves friction.
- Defects: —
- Reader impact: The episode supplies both a satisfying document reversal and the expected next-episode escalation without resolving the central contract too early.

#### Plot Expert
- Stance: Audits causality, Hook alignment, Hook scope, and reveal timing.
- Strengths: The causal chain is closed: signature → transfer ledger → return mark → hidden case → blood clause. Summary, Out, Seeds, and Scene 4 Turn use one Hook strength.
- Defects: —
- Reader impact: The reader can follow why the group moves from one place to the next and is not asked to accept an unexplained discovery.

#### Reader-Editor
- Stance: Checks serialization, opening pressure, exposition restraint, and Out density.
- Strengths: The opening has an immediate confrontation; the final Transition has one principal obligation, the blood clause; no crowded chase or faction arrival dilutes it.
- Defects: —
- Reader impact: Low skim risk because every record detail changes the search and the close makes the next click purposeful.

#### Literary Critic
- Stance: Checks motif, sensory pairing, responsibility theme, and non-didactic closure.
- Strengths: The signature start-point and red seal motif make the abstract protection/control conflict visible; the closing image avoids a thematic lecture.
- Defects: —
- Reader impact: The episode deepens the family conflict without asking the serial reader to accept reconciliation or innocence.

#### Character Critic
- Stance: Checks profile-backed motivation, voice, and relationship pressure.
- Strengths: Jinwoo’s habit of checking hands and seals drives the scene action; Dohyun’s withholding and Hyuk’s procedural skepticism match their profiles; the child’s exit-counting becomes an active boundary.
- Defects: —
- Reader impact: The four-way staging keeps the document scene relational rather than turning it into a detached exposition lecture.

#### Setting/Lore Expert
- Stance: Checks blood-contract rules, location facets, and whether the discovery invents unsupported infrastructure.
- Strengths: Blood as medium, seal traces, and violation danger match the world aspect; all named sets are citeable facets; the hidden case is a situational object within the approved underground storage room, not a new location.
- Defects: —
- Reader impact: The reader receives one usable rule implication—Jinwoo’s blood binds the contract—without a premature full-system explanation.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | Episode-level motifs do not have separate scene `Motif touch` fields. (Literary Critic) | Low | No | yes | The motifs are already concretely placeable in the scene sensory and action fields. For this target reader, adding extra metadata would not improve the evidence chase; generation should simply preserve the motif placements. | Carry the signature-start and red-seal motif placements into Stage ⑥ as generation constraints; no design-field change required. | Carry-⑥ |

**Adjudication:** No High or Med design-field finding remains. The only carried item is a generation constraint permitted by Stage ⑤; it does not block Generation-ready.

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Rationale:** Schema, independent forecast arithmetic, path/facet checks, continuity, Hook body alignment, target-reader checks, and all required persona critiques pass. The only low finding is a permitted Carry-⑥ motif constraint; no Pending design revision remains.
- **Next:** Stage ⑥ — manuscript generation
