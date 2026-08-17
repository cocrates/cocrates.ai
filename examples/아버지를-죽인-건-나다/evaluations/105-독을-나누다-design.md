# Design Evaluation: Episode 105 — 독을 나누다

> Target: `episodes/105-독을-나누다.md`
> Evaluation scope: Stage ⑤ design validation after Architect G4
> Target Reader: 회귀·빙의·환생 무협과 문파 장악물, 가족 반전과 복수형 사이다를 선호하는 성인 남성향 웹소설 독자

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 105 design scope; 1화 완료 기준은 series-level early-arc criterion이다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 105 design scope; early-arc placement criterion이다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 흡수와 의식 역류라는 실질적 대립 사건, 도현의 첫 살인 기록이라는 다음 회차 증거 후크가 Scene 1–4에 구체적으로 배치되었다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1 criterion; Episode 105는 P4 범위이며 도현의 죄와 보호를 동시에 열어 둔다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series-level late-arc payoff criterion; 이번 회차는 흡수와 기록 후크를 전개한다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 미래 지식이 아니라 현장에서 혈맥선의 방향을 읽어 기존 의식을 역이용하며, 결과와 비용을 모두 감당한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Series-ending criterion; Episode 105는 도현의 기록을 판정하지 않고 다음 조사로 넘긴다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript Stage ⑥/⑦ criterion; design forecast is checked in Schema below. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | Four unique `### Scene n` sections use the required meta lines and flat bullet fields. |
| No skill/workflow dump after the design | ✅ | Body contains only episode design, gate evidence, and Architect gate record; no pasted workflow procedure. |
| Unique Scene headings; no pasted twin scenes | ✅ | Scene functions are distinct: cost → reversal → disruption → evidence hook. |
| Canonical episode path | ✅ | Actual path is `episodes/105-독을-나누다.md`. |
| Field notation | ✅ | Required fields use `**Field:**` and `- **Field:**` notation. |
| Every scene has required fields | ✅ | POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, outline, Unit budget, and one Est. appear in each scene. |
| Characters Appearing ↔ On stage union | ✅ | All four Appearing characters are On stage in all four scenes; no ghost cast. |
| On stage includes speakers | ✅ | Every named speaker in each Dialogue intent and Beat is one of the four On-stage characters. |
| Characters ⊆ `characters.md` | ✅ | `characters.md` and all four cited profiles were loaded; all names are catalogued. |
| Summary/Hooks cast alignment | ✅ | Summary, In/Out, Seeds, closing, and scenes name only the four appearing characters. |
| No later-list cast debut | ✅ | No new cast is introduced. |
| Locations ⊆ Key Locations | ✅ | All scenes cite `흑풍루-본거지`, a Key Locations slug. |
| Location facets ⊆ Multi-facet anchors | ✅ | `의식장` is an exact anchor in `locations/흑풍루-본거지.md` under Multi-facet anchors. |
| Nested scene files absent | ✅ | Single canonical episode file; no nested scene directory. |
| No template residue | ✅ | No unresolved template braces remain. |
| Prose forecast present | ✅ | Every scene has typed five-category integer budget and paragraph outline. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 `3×240+3×180+2×120+2×140+1×80=1900`; Sc2 `3×250+4×180+2×120+2×140+1×80=2070`; Sc3 `3×250+3×190+2×120+3×140+1×80=2060`; Sc4 `4×260+3×180+2×120+2×140+1×80=2180`; each written product equals independent arithmetic and each Est. is within ±20%. |
| Dialogue intent vs outline speech | ✅ | All scenes contain dialogue units and named speaking goals; no `none` mismatch. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields `1800+2000+1900+2000=7700`; header addends are identical. |
| Est. length sum ≥ Scale min | ✅ | 7,700 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 7,700 ≤ 8,000; upper-bound warning is recorded, not a failure. |
| Cited staging/profile paths exist | ✅ | The four character profiles, `locations/흑풍루-본거지.md`, `world/혈맥계약과-약그릇.md`, and `stagings/105-독을-나누다.md` were read or successfully written in this turn. |
| Episode List plot | ✅ | Series Summary clauses map to Scenes 1–3 absorption/disruption and Scene 4 whisper/record evidence. |
| Hook evidence strength internal | ✅ | Series Hook, Summary, Out, Scene 4 Beat/Turn all retain the same claim: the Lord presents Dohyun’s first murder record; no truth is resolved. |
| Hook scope | ✅ | One dominant evidence hook plus one supporting blank; no chase, faction arrival, or premature identity reveal. |
| No design-paste / meta-only scenes | ✅ | Every scene has a concrete physical event and causal transition. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | The causal chain moves from involuntary absorption to observed cost, controlled reversal, ritual damage, and record reveal. |
| Scene transitions chain | ✅ | Each Transition out supplies the next Situation: flow observation → first reversal → cracked gate → record evidence. |
| Scene sections complete | ✅ | All four Scene Index rows have complete canonical sections. |
| Generation Readiness | ✅ | All required design, cast, path, facet, hook, staging, and forecast rows pass. |
| Beat concreteness | ✅ | Beats specify bodily actions, direction changes, bloodline-line manipulation, and record handling. |
| Est. length budget | ✅ | Recomputed total 7,700 is inside 4,000–8,000; target-band warning is addressed as a generation constraint. |
| Prose forecast quality | ✅ | Unit categories correspond to the planned action, dialogue, sensory, POV, and transition beats. |
| Episode List scope aligned | ✅ | No later payoff is stolen; the first murder record is presented but its blank and truth remain held. |
| Prior hook addressed | ✅ | Scene 1 begins immediately with the absorption started in Episode 104. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A/Phase B, staging, and continuity loads are explicitly recorded. |
| Series / overview tone & arc honored | ✅ | Fast, physical, event-driven P4 escalation; no training or exposition-only scene. |
| Episode List Summary / Hook honored | ✅ | Both are quoted as mapped design obligations and executed by named scenes. |
| Hook internal consistency | ✅ | Summary, Arc, Out, Seed, and Scene 4 Turn maintain one evidence strength. |
| Characters from architecture; profiles not redefined | ✅ | Existing states are cited; no new identity or equipment state is invented. |
| Profile-backed knowledge / recognition | ✅ | No unsupported identity recognition is required; the Lord’s claim is an adversarial assertion, not accepted knowledge. |
| Locations from architecture; profiles not redefined | ✅ | Only the approved final-base ritual hall facet is used. |
| Location profile paths readable | ✅ | Exact kebab-case path is cited and loaded. |
| Location facets ⊆ anchors | ✅ | `의식장` matches the location profile anchor exactly. |
| Stagings from episode design; blocking not redefined | ✅ | New staging is authored for this episode and keeps cast states/positions fixed. |
| World rules / history consistent with bible | ✅ | Contract backlash and seal/poison interaction remain within established rules. |
| No improvised entities or silent lore | ✅ | No new named faction, place, rule, or character appears. |
| Continuity files used | ✅ | Immediate prior summary and story-so-far are recorded and honored. |
| Character/location state vs story-so-far | ✅ | First ring open, second incomplete, Dohyun collapsed, and Jinwoo absorbing remain intact. |
| Unresolved threads pick up / advance / plant / hold | ✅ | TH-153/154/155/156 are explicitly categorized; no active thread is silently dropped. |
| No contradiction of released continuity | ✅ | The episode extends rather than reverses Episode 104’s absorption and curse collapse. |
| Conflicts section empty or escalated | ✅ | No conflict found. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — required index/detail and continuity set is recorded in the design.
- Locations: ✅ — every scene maps to `흑풍루-본거지`; `+ 의식장` is an exact citeable anchor; all cited paths exist.
- Length / Prose forecast: ✅ — written and recomputed products match for all four scenes; scene fields and header both total 7,700; Scale passes.
- Episode List Summary: ✅ — absorption → Scenes 1–2; ritual disruption → Scenes 2–3; whisper and first murder record → Scene 4.
- Hook to Next / Closing: ✅ — series Hook「흑풍루주가 내민 증거는 도현의 첫 살인 기록이다」; Out「흑풍루주가…도현의 첫 살인 기록을 증거로 내민다」; Scene 4 Turn presents the same record while holding truth and blank for the next episode.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | The seal’s function—victory or replacement vessel—is stated and not immediately answered. |
| Personal stake present | ✅ | Breath, left-arm sensation, mother’s seal, and Dohyun’s collapse create immediate personal risk. |
| Episode Out hook | ✅ | Concrete first murder record, with a visible torn line, creates a next-episode investigation. |
| Exposition budget respected | ✅ | Rules are shown through breath, direction, and bloodline movement rather than lecture. |
| Seed discipline | ✅ | Two Plants and one Hint; held material is explicitly excluded. |
| Scene-first Key Events | ✅ | Each scene is a generation brief, not finished prose. |
| Sensory-emotional on every scene | ✅ | Each scene pairs a physical detail with Jinwoo’s immediate reaction. |
| Motifs planned across scenes | ✅ | Black current and cracked blood seal have observable placements. |
| Overview signature line | N/A | `overview.md` has no required signature dialogue line for this episode. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | Information is limited to actionable flow and bodily cost; every reveal is attached to danger. |
| Sensory-emotional pairing | ✅ | Cold current, metal scrape, broken ritual sound, and wet record all carry POV reaction. |
| Dialogue voices + intent | ✅ | Jinwoo’s terse commands, Dohyun’s low restraint, mother’s soft address, and Lord’s polite educator tone are separated. |
| Reader-discovered meaning | ✅ | The conclusion is held in the reversed flow and record blank, not stated as a moral. |
| Antagonist plausibility | ✅ | The Lord uses guilt and evidence as a control strategy consistent with his contract-centered worldview. |
| Closing image specified | ✅ | Wet torn record edge and one beat of Jinwoo’s left-arm bloodline. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | The episode begins at the prior cliffhanger with bodily danger, not recap. |
| Personal stake matches what this reader came for | ✅ | The father–son–mother conflict is converted into an immediate tactical choice and cost. |
| Pacing / density fits platform expectations | ✅ | Four short, escalating action-revelation scenes; upper-bound forecast is flagged for compression discipline. |
| Out hook makes this reader want the next episode | ✅ | A named first murder record with a torn blank promises concrete evidence investigation. |
| No alienation of core audience without overview intent | ✅ | No romance detour, training stall, or abstract family reconciliation; revenge and evidence remain central. |

## Design Critique (required personas)

#### Target Reader
- Stance: 성인 남성향 회귀 무협 독자의 즉시성·사이다·다음 클릭 욕구 기준으로 읽었다.
- Strengths: 첫 장면이 104회의 흡수를 바로 이어가고, 진우가 비용을 감수해 의식을 역이용하며, 마지막에 기록이라는 물증을 얻는다.
- Defects: — (upper-bound forecast is a manageable generation constraint, not a design defect)
- Reader impact: 독자는 진우의 선택 비용과 아버지의 죄 의혹을 동시에 붙잡고 106회로 넘어갈 가능성이 높다.

#### Genre Critic
- Stance: 회귀 무협의 역전 설계와 가족 복수극의 증거 후크를 점검했다.
- Strengths: 미래 예언이 아니라 현장 판독으로 역전하고, 승리에도 왼팔 감각 상실이라는 값을 남긴다.
- Defects: —
- Reader impact: 주인공의 능력 상승이 만능화되지 않아 장르적 신뢰를 유지한다.

#### Plot Expert
- Stance: Episode List Summary/Hook, 인과, 장면 전환, Out scope를 점검했다.
- Strengths: 흡수 시작 → 방향 판독 → 부분 역류 → 혈인문 균열 → 첫 살인 기록이라는 단일 인과선이 명확하다.
- Defects: — (Hook scope and body-surface alignment checked; no drift found)
- Reader impact: 105회의 사건이 106회의 빈칸 추적으로 자연스럽게 이어진다.

#### Reader-Editor
- Stance: 연재 단위의 첫 페이지·중반 체류·마지막 클릭 후크를 점검했다.
- Strengths: 요약 회상 없이 시작하고, 매 장면의 Turn이 다음 행동을 만든다. 기록의 뜯긴 줄은 단순 폭로보다 조사 욕구를 만든다.
- Defects: — (closing is one dominant evidence hook plus one supporting blank, not crowded)
- Reader impact: 절정 뒤 설명으로 처지지 않고 증거물로 다음 회차의 독서 계약을 갱신한다.

#### Literary Critic
- Stance: 검은 물결·금 간 혈인·젖은 기록의 모티프와 의미 보류를 점검했다.
- Strengths: 보호와 실험의 경계를 물리적 흐름으로 보여주며, 마지막 의미를 기록의 빈칸에 남긴다.
- Defects: — (closing image and no thematic coda are explicitly designed)
- Reader impact: 독자가 도현의 죄를 즉시 면죄하거나 확정하지 않고 스스로 증거를 따라가게 한다.

#### Character Critic
- Stance: 네 인물의 현재 상태, 행동 동기, 목소리, 관계 압력을 점검했다.
- Strengths: 진우는 어머니를 자르지 않고 자신의 몸을 선택하며, 도현은 붕괴 중에도 막고, 흑풍루주는 죄의 해석권을 빼앗으려 한다. 어머니의 경고도 진우의 선택을 대신하지 않는다.
- Defects: — (no unsupported recognition or profile-backed knowledge leap)
- Reader impact: 가족 갈등이 설명이 아니라 서로의 손과 호흡을 놓고 벌이는 행동으로 읽힌다.

#### Setting/Lore Expert
- Stance: 혈맥계약·약 그릇·봉인 규칙과 최종 본거지 의식장 세트의 일관성을 점검했다.
- Strengths: 새 규칙을 추가하지 않고 기존 혈맥선과 혈인문을 역이용한다. `의식장`은 citeable facet이며 staging이 cast states와 위치를 잠근다.
- Defects: —
- Reader impact: 독자는 의식의 물리적 결과를 따라갈 수 있고, 설명 과잉 없이 다음 기록의 중요성을 이해한다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | Forecast total 7,700 is near Scale max and leaves little prose headroom (Target Reader / Reader-Editor) | Low | No | yes | Reader benefits from a tightly eventful climax, but generation must not pad or overshoot 8,000. The design already records the warning and the beats justify the total. | Carry forward as a Stage ⑥ compression constraint: avoid confirmation loops and land at or below the planned Est. sum. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Approved:** 2025-02-14 — Schema, continuity, cast, citeable facet, staging, hook scope, and independent forecast arithmetic pass. The sole low-risk upper-bound note is Carry-⑥ and does not block generation.
- **Next:** Stage ⑥ — manuscript generation.
