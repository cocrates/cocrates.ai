# Design Evaluation: Episode 004 — 회합의 손님

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 범위이며 Episode 004 설계 평가 대상이 아니다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003에서 이미 실행된 기준이다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 2에서 문상철의 흑풍루 거래와 장로 대표의 대립을 실행하고, Scene 3에서 흑풍루 표식 전달 후크를 둔다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | 도현은 장부를 열지 않고 책임을 확정하지 않으며 진우를 시험한다. 거래의 궁극적 이유는 Hold로 유지된다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반/시리즈 회수 기준이며 Episode 004 범위를 넘어선다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 미래를 폭로하지 않고 날짜·수량·인장 대조로 발언 순서를 선점한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 시리즈 종결 기준이다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 Stage ⑥/⑦ 기준이며 분량 예측은 Schema에서 검증한다. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only | ✅ | 세 Scene 모두 표준 메타와 flat bullet 필드를 사용한다. |
| No skill/workflow dump after the design | ✅ | 워크플로 복사 구간이 없다. |
| Unique Scene headings; no pasted twin scenes | ✅ | 세 장면의 기능·사건·예산이 서로 다르다. |
| Canonical episode path | ✅ | 실제 경로 `episodes/004-회합의-손님.md`. |
| Field notation | ✅ | `**Field:**` 및 `- **Field:**` 형식이 일관된다. |
| Every scene has required meta and bullet fields | ✅ | POV, Location, When, On stage, Staging, 사건 필드, outline, typed Unit budget, Est.가 모두 있다. |
| Characters Appearing ↔ On stage union | ✅ | 네 인물이 세 장면의 On stage에 모두 있으며 익명의 문밖 손은 비언어 행위다. |
| On stage includes speakers | ✅ | Dialogue intent와 Beat의 발화자가 각 Scene On stage에 포함된다. |
| Characters ⊆ `characters.md` | ✅ | 서진우·서도현·문상철·장로-대표의 인덱스와 프로필이 모두 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary, Arc, Out, Seeds의 인물이 Appearing에 포함된다. |
| No later-list cast debut | ✅ | Episode 004의 문상철을 포함해 사용 인물의 등장 시점이 허용된다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | `locations.md`에 북문서가 본가가 있고, `locations/북문서가-본가.md`에 `가주전-회랑 접속부`가 정확한 Multi-facet anchor로 등재되어 있다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 단일 episode 설계 파일이며 nested scene 파일이 없다. |
| No template residue | ✅ | 미완성 템플릿 브레이스가 없다. |
| Prose forecast present | ✅ | 허용된 다섯 Unit type의 정수 `n×pick=subtotal`이 있다. |
| Forecast ↔ Est. cross-check | ✅ | S1 1,660, S2 2,180, S3 1,920으로 독립 재계산되며 Est. 1,700/2,200/1,900은 outline density band 안이다. |
| Dialogue intent vs outline speech | ✅ | 대화 의도와 outline/Beat의 발화가 일치한다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene fields `1,700 + 2,200 + 1,900 = 5,800`; header addends도 `5,800`이다. |
| Est. length sum ≥ Scale min | ✅ | 5,800자로 4,000자 이상이다. |
| Est. length sum ≤ Scale max | ✅ | 5,800자로 8,000자 이하이다. |
| Cited staging/profile paths exist | ✅ | `locations/북문서가-본가.md`와 `stagings/004-밤의-비밀-회합.md`, 네 인물 프로필이 모두 디스크에 존재한다. |
| Episode List plot | ✅ | Series Summary의 거래 당사자 확인과 장로 발언 차단이 Scene 2에서 실행된다. |
| Hook evidence strength (internal) | ✅ | Series Hook, Summary, Out, Scene 3 Turn/Transition이 표식 전달이라는 동일한 강도로 정렬된다. |
| Hook scope (no Out creep) | ✅ | 마지막 후크는 표식 전달이며 추격·추가 폭로를 만들지 않는다. |
| No design-paste / meta-only scenes | ✅ | 모든 장면에 증거 이동과 구체적 Turn이 있다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 입장과 검증 → 거래·권한 충돌 → 표식 전달의 인과가 선명하다. |
| Scene transitions chain | ✅ | 운송표 조각 공개가 서류 공개로, 해산이 문밖 표식으로 이어진다. |
| Scene sections complete | ✅ | Scene Index 1~3과 실제 장면이 일치한다. |
| Generation Readiness | ✅ | Schema의 필수 구조·분량·cast·path·facet·Hook 검사가 모두 통과한다. |
| Beat concreteness | ✅ | 장부 봉투, 운송표, 인장, 찢긴 서류, 검은 패가 행동 가능한 Beat로 배치된다. |
| Est. length budget | ✅ | 독립 재계산 5,800 = header = scene Est. 합계이며 Scale 범위 안이다. |
| Prose forecast quality | ✅ | typed units가 Dialogue, action, sensory, POV, transition과 대응한다. |
| Episode List scope aligned | ✅ | Summary와 Hook을 실행하며 Out scope creep가 없다. |
| Prior hook addressed | ✅ | Episode 003의 초대·장부 원본·운송표 조각·부상 상태를 Scene 1에서 이어받는다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected | ✅ | Prior Design Alignment와 Load confirmation에 필요한 파일이 기록되어 있다. |
| Series / overview tone & arc honored | ✅ | 사건 중심 회귀 무협, P1의 의심 유지, 4,000–8,000자 목표를 지킨다. |
| Episode List Summary / Hook honored | ✅ | 거래 당사자·장로 발언 차단·표식 전달이 모두 설계에 있다. |
| Hook internal consistency | ✅ | Summary·Arc·Seeds·Turn·Transition이 동일한 표식 전달 사건을 유지한다. |
| Characters from architecture | ✅ | 네 인물의 인덱스와 개별 프로필이 존재하며 성격을 재정의하지 않는다. |
| Profile-backed knowledge / recognition | ✅ | 진우는 표식의 의미를 안다고 주장하지 않으며, 각 인물의 행동은 프로필의 Drive·Relationships에 근거한다. |
| Locations from architecture | ✅ | `locations.md`의 Key Location, `locations/북문서가-본가.md`의 `가주전-회랑 접속부` anchor, staging의 location anchor가 일치한다. |
| Stagings from episode design | ✅ | `stagings/004-밤의-비밀-회합.md`가 존재하고 장면의 인물·소품·blocking과 일치한다. |
| World rules / history consistent | ✅ | 인장·장부·운송망의 추적성을 사용하고 표식의 기능은 Hold한다. |
| No improvised entities or silent lore | ✅ | 장소·facet·인물·세계 규칙 모두 승인된 아키텍처 또는 episode staging에 근거한다. |
| Continuity files used | ✅ | `story-so-far`, Episode 003 summary, `unresolved-threads`를 명시적으로 이어받는다. |
| Character/location state vs continuity | ✅ | 진우의 부상·미끼, 도현의 초대, 본가 회랑 지정 상태를 보존한다. |
| Unresolved threads handled | ✅ | TH-001~TH-005를 Picks up/Advances/Plants/Holds에 배치한다. |
| No contradiction of released continuity | ✅ | Episode 001~003의 공개 사건과 충돌하지 않는다. |
| Conflicts empty or escalated | ✅ | 설계의 Conflicts / open questions는 None이며 미해결 사항은 Hold와 thread 표에 기록되어 있다. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Required load and cited staging | ✅ | 인물·장소·세계·continuity·staging 파일이 존재하고 설계와 일치한다. |
| Canonical path and cast | ✅ | canonical episode path이며 네 인물은 catalog-backed다. |
| Locations / citeable facets | ✅ | `locations/북문서가-본가.md`의 Multi-facet anchors에 `가주전-회랑 접속부`가 정확히 있다. |
| Length / forecast | ✅ | written/recomputed pairs와 scene/header 합계가 일치한다. |
| Episode List Summary / Hook alignment | ✅ | Summary·Out·Turn·Transition이 series row와 일치한다. |
| Continuity / world / tone | ✅ | 출시 연속성·세계 규칙·P1 톤을 보존한다. |
| Gate result | ✅ | 위치 파일 및 facet 재확인으로 구조적 차단 사유가 없다. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 도현이 왜 진우를 회합에 불렀는지가 명시된다. |
| Personal stake present | ✅ | 잘못 움직이면 도현·장로회·흑풍루 거래망이 동시에 적이 된다. |
| Episode Out hook | ✅ | 표식 전달이 제한적 승리를 감시의 시작으로 뒤집는다. |
| Exposition budget respected | ✅ | 상단주·운송망·표식 세 개념을 사물과 발언으로 드러낸다. |
| Seed discipline | ✅ | 거래 서류는 Plant, 표식은 Hint이며 표식 의미는 확정하지 않는다. |
| Scene-first Key Events | ✅ | 세 장면의 사건·전환·기능·감각·대화·이동이 구체적이다. |
| Sensory-emotional on every scene | ✅ | 모든 장면에 감각 자극과 진우의 억제·판단 반응이 있다. |
| Motifs planned across scenes | ✅ | 인장과 손가락은 Scene 1·2, 닫히는 문은 Scene 1·3에 배치된다. 장면별 Motif-touch는 원고 생성 제약으로 전달한다. |
| Overview signature line | ✅ | 사건 중심성, 도현을 단순 악역으로 확정하지 않기, 분량 기준을 반영한다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 거래 정보가 장부·서류·인장 대조의 충돌 안에서 사건을 발생시킨다. |
| Sensory-emotional pairing | ✅ | 닫힌 문·등불·봉투·금속음이 진우의 기억 억제와 의심에 연결된다. |
| Dialogue voices + intent | ✅ | 도현·진우·문상철·장로 대표의 말투와 목적이 구별된다. |
| Reader-discovered meaning | ✅ | 도현의 보호 의도를 확정하지 않고 독자의 해석을 열어 둔다. |
| Antagonist plausibility | ✅ | 문상철·장로 대표·흑풍루가 서로 다른 이해로 움직인다. |
| Closing image specified | ✅ | 손바닥 위 검은 표식과 닫힌 문틈의 익명 손이 지정되어 있다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns attention | ✅ | 성인 남성향 회귀 무협 독자의 정보 우위와 부자 대치를 증거 싸움으로 즉시 전환한다. |
| Personal stake matches reader promise | ✅ | 아버지의 인장과 미래 배신의 증거가 진우의 복수 판단을 압박한다. |
| Pacing / density fits platform | ✅ | 3장면·5,800자 예측과 실질적 증거 이동이 있다. |
| Out hook drives next episode | ✅ | 표식이 초대인지 감시인지 모호해 Episode 005 추적을 요구한다. |
| No alienation of core audience | ✅ | 사건 중심 사이다와 가족 복수 축을 유지한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 성인 남성향 회귀 무협·문파 장악물 독자의 연속 독서 욕구를 기준으로 본다.
- Strengths: 회합이 즉시 거래 심문과 주도권 싸움으로 변하고, 표식 후크가 다음 질문을 남긴다.
- Defects: —
- Reader impact: 공간·증거·후크가 명료해 다음 회차 진입 동력이 충분하다.

#### Genre Critic
- Stance: 회귀 무협의 정보 선점, 문파 권력 싸움, 부자 복수의 약속을 점검한다.
- Strengths: 진우의 미래 지식이 설명이 아니라 날짜·수량·인장 대조라는 사건으로 작동한다.
- Defects: —
- Reader impact: 장르 독자가 기대하는 설계의 실질적 효력이 드러난다.

#### Plot Expert
- Stance: Summary/Hook의 인과와 장면 전환을 검토한다.
- Strengths: 미끼 공개 → 서류 대조 → 표식 전달의 인과가 명확하고 Hook scope creep가 없다.
- Defects: —
- Reader impact: Episode 005로의 추진력이 자연스럽다.

#### Reader-Editor
- Stance: 회차 판매성, 정보 밀도, 마지막 전환의 과밀 여부를 본다.
- Strengths: 마지막 Transition은 표식 전달이라는 하나의 후크만 두며 5,800자 예측도 중앙 범위다.
- Defects: —
- Reader impact: 독자가 사건을 따라가면서도 결말의 여백을 느낄 수 있다.

#### Literary Critic
- Stance: 모티프와 감각·감정 연결이 원고에서 구현 가능한지 본다.
- Strengths: 인장은 권한과 책임, 문은 통제와 감시의 경계로 기능하며 결말은 사물 이미지로 닫힌다.
- Defects: episode-wide Motif를 장면별 `Motif-touch`로 분해할 필요 → severity Low → Stage ⑥에서 Scene 1/2에 인장·손가락, Scene 1/3에 닫히는 문을 실제 문장 단서로 구현한다.
- Reader impact: 원고 생성 제약으로 전달 가능하며 설계 재작성은 필요하지 않다.

#### Character Critic
- Stance: 인물 동기·목소리·진우의 지식 범위를 프로필과 대조한다.
- Strengths: 네 인물의 행동 동기가 프로필과 맞고 진우가 표식 의미를 과장하지 않는다.
- Defects: —
- Reader impact: 대사와 자리 배치가 권력 압박을 함께 전달한다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Episode-wide Motif를 장면별 `Motif-touch`로 명시할 필요 (Literary Critic) | Low | No | yes | 장면별 구현 단서는 원고의 반복·일관성을 높이지만 설계의 구조 결함은 아니다. | Scene 1/2의 인장·손가락, Scene 1/3의 닫히는 문을 Stage ⑥ 생성 제약으로 전달한다. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G3 — Design Evaluation Decision
- Evaluation status: Complete; Schema와 Design Consistency Gate를 통과했다.
- Location recheck: `locations/북문서가-본가.md` 존재 및 `가주전-회랑 접속부` Multi-facet anchor 확인 완료.
- Remaining handoff constraint: Motif-touch는 Stage ⑥에서 구현한다(Carry-⑥).
- User decision required: 설계 평가를 승인하면 Stage ⑥ 원고 생성으로 이동한다.
