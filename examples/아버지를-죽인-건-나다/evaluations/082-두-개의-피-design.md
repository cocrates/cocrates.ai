# Design Evaluation: Episode 082 — 두 개의 피

## Target Reader
`overview.md`의 성인 남성향 회귀·무협·문파 장악물 독자. 이 독자는 실질적 사건, 냉정한 설계의 쾌감, 가족 복수의 반전, 다음 회차를 강제하는 고백 후크를 기대한다.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 범위의 기준이며 이번 회차 설계 평가 범위가 아니다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | 초반 아크 완료 기준이며 이번 회차에서 재실행할 기준이 아니다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 2의 몰래 채취 발각과 Scene 4의 첫 살인 고백이 각각 실질적 대립과 다음 회차 후크를 실행한다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | Episode 082는 P3이며 해당 P1 유지 기준은 이 회차의 적용 범위가 아니다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반 시리즈 회수 기준이다. 이번 회차는 약 그릇과 조건의 현재 적용만 수행하며 회수 완료를 주장하지 않는다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 현재 기록과 의원의 제자 절차를 이용하지만 도현의 동의 없는 채취는 발각되며, 회귀 지식이 만능 해답이 되지 않는다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 결말 기준이다. 이번 회차는 부자 대면을 다음 회차로 확장하는 중간 압력으로 설계됐다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 Stage ⑥/⑦에서 판정한다. 설계의 Forecast는 Schema에서 별도 검증한다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | `episodes/082-두-개의-피.md`에 4개 Scene이 동일한 정규 메타·flat bullet 필드를 사용한다. |
| No skill/workflow dump after the design | ✅ | Episode body는 이야기 설계이며 workflow 설명을 복사하지 않았다. |
| Unique Scene headings; no pasted twin scenes | ✅ | Scene 1–4의 Situation/Beat/Turn/Function/outline가 서로 다른 인과 단계를 가진다. |
| Canonical episode path | ✅ | `episodes/082-두-개의-피.md` — 제목의 한국어 kebab slug와 3자리 번호가 일치한다. |
| Field notation | ✅ | `**POV:**`, `- **Situation:**` 등 정규 표기를 사용한다. |
| Every scene has required fields | ✅ | 네 장면 모두 POV, Location, When, On stage, Staging, 9개 Key Event 필드, outline, Unit, Est.를 갖춘다. |
| Characters Appearing ↔ On stage union | ✅ | Appearing은 세 장면의 union인 서진우·서도현·의원의-제자이며 Scene 4의 두 사람도 모두 포함된다. |
| On stage includes speakers | ✅ | 각 Scene의 Dialogue intent·Beat 주체가 모두 해당 On stage에 있다. |
| Characters ⊆ `characters.md` | ✅ | `characters/서진우.md`, `characters/서도현.md`, `characters/의원의-제자.md`를 이번 턴에 read OK. |
| Summary/Hooks cast alignment | ✅ | Summary·In·Out·Seeds·Closing에는 Appearing cast 외 이름이 없다. |
| No later-list cast debut | ✅ | 세 인물 모두 082 이전 architecture와 continuity에 존재한다. |
| Locations ⊆ Key Locations | ✅ | `북문서가-본가`가 `locations.md` Key Locations에 존재한다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `locations/북문서가-본가.md` read OK; `지하 보관실`, `가주전 문앞`이 Multi-facet anchors에 정확히 기재되어 있다. |
| Nested scene files absent | ✅ | 단일 episode gate 파일만 사용한다. |
| No template residue | ✅ | raw template braces가 없다. |
| Prose forecast present | ✅ | 모든 장면에 5종 typed integer formula와 5–6개 paragraph intents가 있다. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 1,480→1,500; Sc2 1,680→1,700; Sc3 1,680→1,700; Sc4 1,680→1,700. 각 written product는 재산술과 일치하고 outline density band 안이다. |
| Dialogue intent vs outline speech | ✅ | 대사가 있는 모든 장면에 dialogue units와 named speakers가 명시돼 있다. |
| Recorded Estimated Length = scene sum | ✅ | scene fields 1,500+1,700+1,700+1,700=6,600; header addends도 1,500+1,700+1,700+1,700=6,600이다. |
| Est. sum ≥ Scale min | ✅ | 6,600 ≥ 4,000. |
| Est. sum ≤ Scale max | ✅ | 6,600 ≤ 8,000; central band다. |
| Cited staging/profile paths exist | ✅ | `stagings/082-두-개의-피.md`, 세 character profiles, `locations/북문서가-본가.md`, `world/혈맥계약과-약그릇.md` 모두 이번 턴 read/write OK. |
| Episode List plot | ✅ | Series Summary의 몰래 채취·발각·불신 질문·첫 살인 고백이 각각 Scene 2 Turn, Scene 3 Turn, Scene 4 Turn으로 매핑된다. |
| Hook evidence strength | ✅ | Series Hook「진우는 처음으로 회귀 전 자신이 아버지를 죽였다고 말한다」; Summary/Out/Scene 4 Turn도 같은 ‘처음 말함’ 강도다. |
| Hook scope | ✅ | 마지막 Transition은 첫 고백 한 가지 의무만 남기며 추격·추가 폭로를 만들지 않는다. |
| No design-paste / meta-only scenes | ✅ | 각 장면에 채취 조건→발각→신뢰 조건→고백이라는 독립적 사건 변화가 있다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | 재료 필요 확인에서 비밀 채취, 발각, 신뢰 대면, 고백으로 압력이 상승한다. |
| Scene transitions chain | ✅ | 각 Transition out이 다음 Situation을 직접 만든다. 시간은 모두 같은 밤의 연속이다. |
| Scene sections complete | ✅ | Scene Index 4행과 완전한 Scene 4개가 일치한다. |
| Generation Readiness | ✅ | Schema·Consistency·Forecast·Hook·Path·Cast 모두 통과하며 Hold가 stage ⑥의 발명을 제한한다. |
| Beat concreteness | ✅ | 손목 잠금, 도구 회수, 문양 반응, 피 한 방울 직전 등 관찰 가능한 사건으로 구성된다. |
| Est. length budget | ✅ | 독립 재산술 합계 6,600으로 Scale 범위와 일치한다. |
| Prose forecast quality | ✅ | dialogue/action/sensory/POV/transition 단위가 각 scene Beat와 outline에 대응한다. |
| Episode List scope aligned | ✅ | Summary와 Hook을 넘는 추격·새 세력·치료 성공을 추가하지 않는다. |
| Prior hook addressed | ✅ | Scene 1이 Episode 081의 ‘원액과 도현의 피’ 후크를 즉시 집는다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected | ✅ | Phase A/B, staging, continuity가 Prior Design Alignment에 기록됐다. |
| Series / overview tone & arc honored | ✅ | 사건 중심, 냉정한 계산, 부자 관계의 제한적 감정 누출이 유지된다. |
| Episode List Summary / Hook honored | ✅ | Summary의 네 서명 사건과 Hook의 첫 고백이 body surfaces에 동일 강도로 반영됐다. |
| Hook internal consistency | ✅ | Summary·Arc close·Out·Seeds·Scene 4 Turn이 동일한 첫 고백을 말한다. |
| Characters from architecture | ✅ | 프로필을 재정의하지 않고 상태만 기존 `봉인-해제`·`병세-노출`·`base`로 인용했다. |
| Profile-backed knowledge | ✅ | 도현이 진우의 불신 이유를 묻는 관계 압력은 관계 맵과 현재 연속성에 근거한다. |
| Locations from architecture | ✅ | 두 장소는 `북문서가-본가`의 exact anchors다. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` read OK. |
| Location facets ⊆ anchors | ✅ | 두 facet 모두 profile의 Multi-facet anchors에 exact match한다. |
| Staging blocking | ✅ | `stagings/082-두-개의-피.md`가 세 cast state, blocking, props, ambient를 잠근다. |
| World rules / history consistent | ✅ | 피는 치료 확정물이 아니며 약 그릇은 분산만 한다. 새 규칙을 만들지 않았다. |
| No improvised entities or silent lore | ✅ | 새 인물·세력·장소·규칙이 없다. |
| Continuity files used | ✅ | story-so-far와 immediate prior summary만 사용했다. |
| Character/location state vs story-so-far | ✅ | 081 이후 상태를 유지하고 도현의 피를 아직 확보·치료 완료하지 않았다. |
| Unresolved threads pick up/advance/hold | ✅ | TH-121·TH-122·TH-115의 advance/hold가 Continuity References에 명시됐다. |
| No contradiction of released continuity | ✅ | 치료 결과와 원액 위치를 보류해 후속 권위를 침범하지 않는다. |
| Conflicts empty or escalated | ✅ | Conflicts는 None이다. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — required indexes, appearing/used profiles, staging, world aspect, and immediate continuity were loaded.
- Locations: ✅ — Index evidence `북문서가-본가` ∈ Key Locations; Path evidence `locations/북문서가-본가.md` read OK and `stagings/082-두-개의-피.md` read/write OK; Facet evidence `지하 보관실`, `가주전 문앞` ⊆ Multi-facet anchors.
- Length / Prose forecast: ✅ — Sc1 written=1,480; recomputed=1,480; Est=1,500 · Sc2 written=1,680; recomputed=1,680; Est=1,700 · Sc3 written=1,680; recomputed=1,680; Est=1,700 · Sc4 written=1,680; recomputed=1,680; Est=1,700 · scene fields/header sum=6,600.
- Episode List Summary: ✅ — 몰래 채취·발각→Scene 2; 피 대신 불신 이유 질문→Scene 3; 첫 살인 고백→Scene 4.
- Hook to Next / Closing: ✅ — Hook「처음으로 회귀 전 자신이 아버지를 죽였다고 말한다」; Out「처음으로 회귀 전 자신이 아버지를 죽였다고 말한다」; Scene 4 Turn「진우가 회귀 전 자신이 아버지를 죽였다고 처음 말한다」; same strength.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | 진우가 치료를 구하는지 도현을 증거로 묶는지 질문한다. |
| Personal stake present | ✅ | 독성 반응과 도현의 생존이 진우의 행동을 즉시 압박한다. |
| Episode Out hook | ✅ | 첫 살인 고백을 정확히 실행하며 다음 회차의 부자 대화를 요구한다. |
| Exposition budget respected | ✅ | 기존 규칙의 절차 적용만 설명하고 새 lore lecture를 보류한다. |
| Seed discipline | ✅ | Plant 1개, Hint 1개, Hold 목록으로 제한했다. |
| Scene-first Key Events | ✅ | 네 scene이 모두 concrete action과 transition을 갖는다. |
| Sensory-emotional on every scene | ✅ | 냄새·금속성·문양·등불이 각 POV 반응과 연결된다. |
| Motifs planned | ✅ | 두 개의 피와 손목/칼자루를 1–4에 배치했다. |
| Overview signature line | N/A | overview.md에 이 회차가 반드시 배치해야 할 고정 signature line은 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | 절차 설명은 Scene 1에 제한되고 Scene 2부터 물리적 발각과 관계 압력이 우세하다. |
| Sensory-emotional pairing | ✅ | 각 scene의 감각이 진우의 계산·긴장·회피로 이어진다. |
| Dialogue voices + intent | ✅ | 진우의 단문 회피, 도현의 낮은 존대, 제자의 빠른 기술 설명이 구분된다. |
| Reader-discovered meaning | ✅ | 진우가 도현을 재료로 취급한다는 의미를 행동으로 드러내며 해설을 Hold한다. |
| Antagonist plausibility | N/A | 이 회차는 흑풍루주와 직접 대치하지 않는 부자 대면 회차다. |
| Closing image specified | ✅ | 피가 떨어지기 직전 유리병과 풀리는 칼자루 손을 지정했다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | 첫 장면이 치료 재료의 부족과 도현의 피라는 즉시 목표로 시작한다. |
| Personal stake matches reader expectation | ✅ | 아버지의 피를 몰래 얻는 행위가 복수·효·생존을 한 번에 건드린다. |
| Pacing / density fits platform expectations | ✅ | 4 scenes, 6,600자 forecast, 발각→대면→고백의 명확한 상승 구조다. |
| Out hook makes this reader want the next episode | ✅ | 주인공이 아버지를 죽였다는 핵심 미공개 기억을 처음 언어화한다. |
| No alienation of core audience | ✅ | 감정 장면도 절차·행동·증거에 묶여 있어 치유물식 화해로 이탈하지 않는다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 성인 남성향 회귀·무협 독자의 즉시 사건성과 다음 회차 욕구를 기준으로 읽었다.
- Strengths: 도현의 피를 얻으려는 실무 목표가 곧 신뢰 갈등으로 전환되고, 마지막 고백이 강한 retention을 만든다.
- Defects: —
- Reader impact: 첫 장면부터 재료 부족과 부자 비밀을 동시에 제시해 이탈 지점을 만들지 않는다.

#### Genre Critic
- Stance: 회귀 무협의 선점 설계, 증거 대조, 사이다 직전의 대립을 기준으로 읽었다.
- Strengths: 회귀 지식이 정답이 아니라 실패하는 비밀 채취 계획으로 작동한다.
- Defects: —
- Reader impact: 주인공의 냉정함이 만능으로 보이지 않고 상대의 반격을 받아 장르 긴장이 유지된다.

#### Plot Expert
- Stance: Series Summary·Hook, 장면 인과, Out scope를 우선 점검했다.
- Strengths: `재료 필요→비밀 채취→발각→신뢰 조건→첫 고백`의 인과가 Hook을 정확히 향한다.
- Defects: —
- Reader impact: Episode 083의 고백 대화를 미루되 이번 회차 자체에는 완결된 발각과 조건 전환이 있다.

#### Reader-Editor
- Stance: 연재 단위의 opening, density, closing Out의 판매력을 점검했다.
- Strengths: 마지막 Transition이 첫 고백 하나만 남겨 과밀하지 않고, 6,600자 forecast가 사건 단위와 맞는다.
- Defects: —
- Reader impact: 설명을 길게 늘이지 않고 손목과 병을 통해 다음 회차를 당긴다.

#### Character Critic
- Stance: 프로필의 drive/voice와 도현의 직접 내면 공개 금지 제약을 점검했다.
- Strengths: 진우는 증거를 먼저 선택하고, 도현은 낮은 존대와 행동으로 피를 내줄 조건을 바꾼다. 도현의 내면 독백은 없다.
- Defects: —
- Reader impact: 두 인물의 신뢰 갈등이 고백의 감정적 비용을 만들며 관계 회복으로 성급히 봉합되지 않는다.

#### Literary Critic
- Stance: Motif, sensory-emotional pairing, closing image와 Hold의 절제를 점검했다.
- Strengths: ‘두 개의 피’가 재료와 책임의 이중 의미를 가지며, 칼자루를 찾았다가 푸는 손이 closing image로 기능한다.
- Defects: Episode-level motif touch가 일부 generation 단계에서 놓칠 수 있다.
- Reader impact: 독자가 설명을 듣기 전에 손과 피의 반복으로 진우의 죄책감을 감지한다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | Literary Critic: motif touch를 원고에서 놓치지 말 것 | Low | No | yes | 독자의 반복 이미지 인식을 강화하지만 현재 설계 Schema를 바꿀 정도의 결함은 아니다. | Stage ⑥ Carry: 두 개의 피와 손목/칼자루 이미지를 각 장면의 sensory/action에서 유지한다. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Decision:** 2025-02-14 — Schema, continuity, hook, forecast, target-reader, and required persona checks pass. The only craft reminder is a non-blocking Carry-⑥ motif constraint; no Pending finding remains.
- **Next:** Stage ⑥ — manuscript generation from the locked design.
