# Design Evaluation: Episode 088 — 전쟁의 계곡

- **Evaluated artifact:** `episodes/088-전쟁의-계곡.md`
- **Stage:** ⑤ Design Evaluation
- **Target Reader:** 회귀·빙의·환생 무협, 문파 장악물, 가족 반전과 복수형 사이다를 선호하는 성인 남성향 웹소설 독자.
- **Evaluation date:** 2025-02-14

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | 1화 고정 기준이며 Episode 088 설계 범위가 아니다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | 초반부 완료 기준이며 이 회차의 설계 범위가 아니다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scenes 1–3에서 계곡 추적·민간인 피난·책임 전가 장치 차단이 실행되고 Scene 4에서 도현 인계 거래 대립과 다음 회차의 인질 선택 후크가 열린다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1 범위가 아닌 Episode 088이며, 도현의 자기 인계는 후반 책임 갈등의 설계다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 시리즈 후반 회수 기준이다. 이 회차는 가환의 지도 조각과 계약 인계 조건을 진전시키며 완전 회수하지 않는다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 미래 지식으로 계곡의 조작을 즉시 정답 처리하지 않고, 지도·수레·표식·신호를 현장에서 검증한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | ✅ | 도현을 살리는 일과 다시 인질로 내놓는 거래가 충돌하며, 흑풍루주 직접 등장은 Hold하고 부자 선택만 다음 회차로 밀어 둔다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 평가는 Stage ⑦에서 수행한다. 설계의 Forecast는 Schema에서 별도 검증한다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only | ✅ | `### Scene 1`–`### Scene 4`가 모두 canonical meta lines와 flat bullet fields를 사용한다. |
| No skill/workflow dump after the design | ✅ | Stage 절차를 복사한 본문이 없고, episode-specific 설계와 gate block만 있다. |
| Unique Scene headings; no pasted twin scenes | ✅ | 4개 제목·목적·Turn이 각각 다르다. |
| Canonical episode path | ✅ | 실제 경로 `episodes/088-전쟁의-계곡.md`. |
| Field notation | ✅ | `**POV:**`, `**Location:**`, `- **Beat:**` 등 canonical 표기다. |
| Required fields complete | ✅ | 각 Scene에 POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Outline, Unit budget, Est. length가 1회씩 있다. |
| Characters Appearing ↔ On stage union | ✅ | Appearing 5명은 Scene 1–4의 On stage 합집합과 일치한다. |
| On stage includes speakers | ✅ | 진우·혁·가환·도현·사자의 의도적 발화/행동이 각 해당 Scene On stage에 있다. 민간인 extras는 침묵·이동만 한다. |
| Characters ⊆ `characters.md` | ✅ | `서진우`, `남궁혁`, `가환`, `서도현`, `흑풍루의-사자` 모두 catalog row와 profile이 있다. |
| Summary/Hooks cast alignment | ✅ | Summary, Seeds, Closing의 고유명사는 Appearing roster 안에 있다. |
| No later-list cast debut | ✅ | 모두 088 이전 architecture와 series presence가 있는 인물이다. |
| Locations ⊆ Key Locations | ✅ | 4개 Scene 모두 `전쟁의-계곡`이며 `locations.md` Key Locations row에 있다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `진입로 / 마을터 / 수원 / 절벽길` 모두 `locations/전쟁의-계곡.md`의 exact anchors다. |
| Nested scene files absent | ✅ | 단일 canonical episode 파일이다. |
| No template residue | ✅ | raw placeholder brace가 없다. |
| Prose forecast present | ✅ | 5가지 unit type의 정수 `n×pick = subtotal`과 6–7개 outline이 모든 Scene에 있다. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 `3×240+4×180+2×120+1×140+1×80=1,900`, Est 1,800; Sc2 동일; Sc3 `3×250+4×190+2×120+2×140+1×80=2,110`, Est 2,000; Sc4 `5×250+3×190+2×120+2×150+1×80=2,440`, Est 2,300. 모두 ±20% 및 outline density 안이다. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields `1,800+1,800+2,000+2,300=7,900`; header addends `1,800+1,800+2,000+2,300=7,900`. |
| Est. length sum ≥ Scale min | ✅ | 7,900 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 7,900 ≤ 8,000; 상한 근접은 기록했으나 각 forecast가 mid-low이고 반복 padding이 없다. |
| Cited staging/profile paths exist | ✅ | This-turn reads succeeded for `characters/서진우.md`, `characters/남궁혁.md`, `characters/가환.md`, `characters/서도현.md`, `characters/흑풍루의-사자.md`, `locations/전쟁의-계곡.md`, `stagings/088-계곡-민간인-피난.md`, `stagings/088-계곡-인질-대치.md`. |
| Episode List plot | ✅ | `series.md` Summary의 민간인 피난·책임 전가 차단·도현 자기 인계가 각각 Scenes 1–3, 2–3, 4의 concrete Beat으로 실행된다. |
| Hook evidence strength | ✅ | Series Hook「진우는 아버지를 다시 인질로 내놓을 것인지 선택해야 한다」, Summary의 도현 자기 인계, Out의 선택 상태, Scene 4 Turn의 도현 전진이 같은 강도로 연결된다. |
| Hook scope | ✅ | Out은 ‘진우의 도현 인질 선택’ 하나이며 새 추격·새 세력 도착·두 번째 폭로를 추가하지 않는다. |
| No design-paste / meta-only scenes | ✅ | 각 Scene은 빈 마을 확인→물증 확보→신호 차단→인질 대치의 causal chain을 가진다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 지도 후크가 민간인 위기로 전환되고, 물증·신호의 물리적 차단이 도현 인계 선택으로 수렴한다. |
| Scene transitions chain | ✅ | 진입로의 피난 판단→마을터 물증→수원 신호→절벽길 사자 접근이 각 Transition out과 다음 Situation에 연결된다. |
| Scene sections complete | ✅ | Scene Index 4행 모두 완전한 Scene section을 가진다. |
| Generation Readiness | ✅ | Schema/Consistency에서 Length, cast, paths, facets, Hook body 모두 ✅이며 Generation-ready 조건을 충족한다. |
| Beat concreteness | ✅ | 수레·패찰·표식 돌·신호 심지·인장의 이동과 차단을 구체적으로 설계했다. |
| Est. length budget | ✅ | 독립 재계산 합계 7,900, Scale 통과, central band보다 높지만 상한 이내다. |
| Prose forecast quality | ✅ | Dialogue intent와 outline의 발화·행동·감각·POV 단위가 일치한다. |
| Episode List scope aligned | ✅ | Summary와 Hook을 확장하지 않고 사건화했다. |
| Prior hook addressed | ✅ | 087의 지도 조각·가환 구출·계곡 진입을 Scene 1에서 즉시 회수한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected | ✅ | Prior Design Alignment에 Phase A/B, staging, continuity 범위를 기록했다. |
| Series / overview tone & arc honored | ✅ | 사건 중심·성인 남성향 사이다의 현장 검증 리듬을 유지한다. |
| Episode List Summary / Hook honored | ✅ | Summary와 Hook을 직접 인용 가능한 Scene Turn으로 매핑했다. |
| Hook internal consistency | ✅ | Summary/Arc closing/Out/Scene 4 Turn이 모두 도현 자기 인계 뒤 진우의 선택을 가리킨다. |
| Characters from architecture | ✅ | 인물 profile을 재정의하지 않고 기존 states만 cite했다. |
| Profile-backed knowledge / recognition | ✅ | 도현-가환의 지도 관계, 진우-도현의 계약·혈맥 인식, 사자의 제한된 지식은 profiles에 근거한다. |
| Locations from architecture | ✅ | 단일 Key Location과 exact facet만 사용했다. |
| Location profile paths readable | ✅ | `locations/전쟁의-계곡.md` read OK. |
| Location facets ⊆ anchors | ✅ | profile의 anchors `진입로 / 마을터 / 수원 / 절벽길`과 exact 일치한다. |
| Stagings from episode design | ✅ | 두 상황을 분리해 cast/state/blocking을 고정했고 Scene 1–3과 4의 span이 혼합되지 않는다. |
| World rules / history consistent | ✅ | 봉인 추가 해제를 하지 않고, 계약은 독 발작·추적 신호의 관찰 가능한 위험으로만 사용한다. |
| No improvised entities or silent lore | ✅ | 새 인물·세력·규칙·장소를 만들지 않았다. 민간인과 경비는 speech 없는 extras로 제한했다. |
| Continuity files used | ✅ | immediate prior summary와 story-so-far만 사용했다. |
| Character/location state vs story-so-far | ✅ | 진우 `봉인-해제`, 혁 `전장-복장`, 가환 `심문-상태`, 도현 `병세-노출`을 유지하며 도현의 계곡 등장만 이번 회차의 상태 변화로 계획했다. |
| Unresolved threads pick up/advance/plant/hold | ✅ | TH-126/127/128 모두 Picks up·Advances·Plants·Holds에 명시되어 있다. |
| No contradiction of released continuity | ✅ | 어머니·흑풍루주 위치와 원액을 확정하지 않으며 087의 지도·구출 결과를 되돌리지 않는다. |
| Conflicts empty or escalated | ✅ | material conflict 없음; 도현의 계곡 도착 경로는 Hint/Hold로 남겨 조기 설명을 피했다. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — selective Phase A/B load and immediate prior continuity complete.
- Locations index: ✅ — all scenes map to `전쟁의-계곡` Key Location.
- Locations paths: ✅ — exact character, location, and staging paths were read OK this turn.
- Location facets: ✅ — all four exact anchors match the location profile.
- Length / Forecast: ✅ — Scene fields/header both sum to 7,900; all four written products equal independent recomputation.
- Episode List Summary: ✅ — 민간인 피난→Scenes 1–3; 책임 전가 계획 차단→Scenes 2–3; 도현 자기 인계→Scene 4.
- Hook / Closing: ✅ — Hook「진우는 아버지를 다시 인질로 내놓을 것인지 선택해야 한다」; Out「진우는 도현을 다시 인질로 내놓을지 아니면 사자와 계곡 안에 가둘지 선택해야 하는 상태」; Scene 4 Turn「도현이 청강검을 뽑지 않은 채 한 걸음 사자 쪽으로 나선다」. Same strength, no scope creep.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 지도 목적과 계곡의 함정 여부를 Opening Question으로 제시한다. |
| Personal stake present | ✅ | 민간인 생존과 도현 재인계의 반복 위험이 진우의 선택을 압박한다. |
| Episode Out hook | ✅ | 도현 인계라는 단일 선택이 구체적이고 089로 직결된다. |
| Exposition budget respected | ✅ | 지도·표식·신호를 현장 행동으로 판독하며 별도 강의 장면이 없다. |
| Seed discipline | ✅ | Plant 1, Hint 2, Hold 4로 과다 공개를 막았다. |
| Scene-first Key Events | ✅ | 모든 Scene이 canonical fields를 완성했다. |
| Sensory-emotional pairing | ✅ | 각 Scene에 계곡 바람·불씨·수원·봉랍의 감각과 POV 반응이 있다. |
| Motifs planned across scenes | ✅ | 검은 표식/붉은 실과 인장을 행동에 배치했다. |
| Overview signature line | N/A | overview.md에 이번 회차에 직접 배치해야 할 별도 signature dialogue가 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 정보는 물증을 확인하는 행동 안에 들어가며 Scene 4의 거래 압박이 설명을 앞선다. |
| Sensory-emotional pairing | ✅ | 차가운 바람, 곡물 냄새, 수원, 봉랍이 인물 반응과 결합한다. |
| Dialogue voices + intent | ✅ | 진우의 짧은 지시, 혁의 명분 중심 응답, 가환의 생략, 도현의 낮은 존대, 사자의 조건 반복이 구별된다. |
| Reader-discovered meaning | ✅ | 흑풍루가 ‘증거의 방향’을 배치한다는 의미를 독자가 표식·패찰·신호의 연결로 추론한다. |
| Antagonist plausibility | ✅ | 사자는 전쟁 전체를 설명하지 않고, 인계 조건과 증거 조작이라는 자신이 아는 범위만 반복한다. |
| Closing image specified | ✅ | 검은 표식 위 도현 인장과 그쪽으로 돌아가는 진우의 금 간 검 손잡이. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns locked reader attention | ✅ | 지도 후크가 곧바로 빈 마을·민간인 위기로 변해 설명 지연 없이 사건을 시작한다. |
| Personal stake matches reader promise | ✅ | 가족 복수와 아버지의 자기 인계가 계곡 전쟁의 명분 싸움과 결합한다. |
| Pacing / density fits platform expectations | ✅ | 4개 장면이 추적→검증→차단→대치로 직진하며 7,900자 forecast에 반복 padding이 없다. |
| Out hook makes this reader want next episode | ✅ | 도현을 인질로 넘길지 가둘지의 즉시 선택은 복수형 웹소설 독자의 다음 회차 기대를 직접 자극한다. |
| No alienation of core audience | ✅ | 민간인 피해를 과잉 감상화하지 않고, 검증·거래·부자 갈등 중심으로 유지한다. |

## Design Critique (required personas)

#### Target Reader
- **Stance:** 성인 남성향 회귀 무협 독자의 다음 회차 지속 의사 기준으로 검토.
- **Strengths:** 087의 지도 후크가 빈 마을과 곧바로 연결되고, 민간인 피난이라는 사건이 가족 인질 선택으로 수렴한다.
- **Defects:** Low — 도현의 계곡 도착 경로가 즉시 해명되지 않는다.
- **Reader impact:** 미스터리로서 다음 회차 동력을 만들며, 현재 회차의 선택을 흐리지는 않는다.

#### Genre Critic
- **Stance:** 회귀 무협·문파 전쟁·사이다 사건의 장르 약속을 검토.
- **Strengths:** 주인공이 미래 지식으로 자동 승리하지 않고, 표식과 물증을 현장에서 뒤집는다. 아버지의 자기 인계는 장르의 관계 압박을 강화한다.
- **Defects:** —
- **Reader impact:** 전쟁 명분을 실물 증거로 다루어 ‘설계로 판을 뒤집는’ 쾌감을 유지한다.

#### Plot Expert
- **Stance:** 인과, 후크 정렬, Out scope를 검토.
- **Strengths:** 지도→병목→가짜 물증→신호→사자→도현 인계의 인과가 단선적이고, Hook의 선택 의무가 Summary/Out/Turn에 같은 강도로 반복된다.
- **Defects:** Low — 도현의 등장 경로는 이번 회차의 Hint로 남겨야 하며, 원고에서 설명을 덧붙이면 긴장이 약해질 수 있다.
- **Reader impact:** 원인 설명을 뒤로 미뤄도 현재 사건은 닫히므로 Out의 선택이 선명하다.

#### Reader-Editor
- **Stance:** 연재 단위의 흡인력, 정보량, 마지막 전환을 검토.
- **Strengths:** Scene 2의 물증과 Scene 3의 신호 차단이 중복되지 않고, Scene 4가 관계 갈등을 새 국면으로 연다.
- **Defects:** Med — 7,900자 forecast가 Scale 상한에 가깝다.
- **Reader impact:** 사건을 줄이지 않고도 허용되지만, Stage ⑥에서 각 장면을 ±20% 안에 쓰며 설명 반복을 금지해야 한다.

#### Literary Critic
- **Stance:** 반복 모티프와 감정의 물질화를 검토.
- **Strengths:** 검은 표식/붉은 실이 길의 통제권을, 인장이 보호와 인질의 경계를 시각화한다. 닫는 이미지가 주제 설명을 대신한다.
- **Defects:** —
- **Reader impact:** 사이다 사건의 속도를 늦추지 않고 부자 갈등을 시각적으로 각인한다.

#### Character Critic
- **Stance:** 인물 동기, 행동 가능성, profile-backed knowledge를 검토.
- **Strengths:** 진우는 추가 해제 대신 판단을 택하고, 혁은 명분 증거를 보존하며, 가환은 제한 증언을 넘지 않는다. 도현과 사자의 행동은 각 profile의 보호·거래 drive와 맞는다.
- **Defects:** Low — 도현의 계곡 도착은 감정적 편의를 피하도록 원고에서 이동 경로를 확정하지 말고, 인장·병세·자기 인계 행동으로만 설득해야 한다.
- **Reader impact:** 도현이 갑자기 모든 답을 아는 인물로 변하지 않아 기존의 의심과 복수 감정을 유지한다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Scale 상한에 가까운 forecast (Reader-Editor) | Med | No | no | 7,900자는 4,000–8,000 범위 안이며 4개 사건 전환을 삭제하면 민간인 피난·책임 전가 차단·자기 인계 중 하나가 약해진다. Target Reader에게는 압축보다 직진 사건 밀도가 유리하다. | Stage ⑥에 각 Scene Est. ±20%, 설명 반복 금지, mid-low 문장 밀도 Carry | Carry-⑥ |
| 2 | 도현의 계곡 도착 경로 미확정 (Target Reader, Plot Expert, Character Critic) | Low | No | no | 현재 회차의 핵심은 도현이 왜 왔는지의 해답이 아니라 자신을 거래물로 내놓는 행동이다. 경로를 확정하면 TH-127/후속 미스터리를 앞당긴다. | 원고에서 경로 설명·후속 정보 공개 금지; 인장·병세·자기 인계만 dramatize | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Architect G5 Approval
- **Status:** Approved by Architect (2025-02-14).
- **Rationale:** Schema, independent length arithmetic, path/facet checks, continuity, hook alignment, and required persona critiques all pass. The two Low/Med findings are generation constraints, not design-field defects; no Pending revision remains.
- **Next:** Stage ⑥ manuscript generation with Carry-⑥ constraints recorded above.
