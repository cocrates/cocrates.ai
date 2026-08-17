# Design Evaluation: Episode 039 — 내부자의 거래

> Target: `episodes/039-내부자의-거래.md`
> Evaluated after Architect G4: 2025-02-14

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope; this episode does not retroactively satisfy the opening requirement. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003 scope; already established and not this episode’s design criterion. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 실질적 거래·인계망 대립을 Scenes 2–3에서 실행하고, Scene 3에서 남궁가 혼례 행렬 길목 후크를 남긴다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1 criterion; this is P2 and 도현은 직접 등장하지 않으며 최종 의도를 Hold한다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series/late-arc payoff; this episode advances 약 그릇·인계망 evidence but does not claim final recovery. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 038 이후의 표식과 현재 기록만으로 거래 범위를 좁히며, 미래 지식으로 혼례 결과를 단정하지 않는다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Series/endgame criterion; 도현의 의도와 부자 대면은 Hold한다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript Stage ⑦ criterion; design에서는 typed forecast를 Schema에서 별도 검증한다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only | ✅ | 3개 Scene Index row와 3개 완결 Scene section이 canonical fields로 대응한다. |
| No skill/workflow dump after the design | ✅ | 설계 본문 뒤에 절차 복사물이 없고 Gate G4 블록만 있다. |
| Unique Scene headings; no pasted twin scenes | ✅ | Scene 1은 기록 대조, Scene 2는 내부자 증언, Scene 3은 선택과 후크로 기능이 분리된다. |
| Canonical episode path | ✅ | `episodes/039-내부자의-거래.md`. |
| Field notation | ✅ | 모든 meta/bullet field가 `**Field:**` 표기다. |
| Every scene has required fields | ✅ | POV·Location·When·On stage·Staging·Situation·Beat·Turn·Function·Sensory·Dialogue intent·Transition·Outline·Unit·Est. 모두 존재한다. |
| Characters Appearing ↔ On stage union | ✅ | Appearing 4명은 세 장면 On stage의 합집합과 일치한다. |
| On stage includes speakers | ✅ | 내부자·혁·아이·진우의 발화 의도는 각 해당 장면 On stage에 있다. |
| Characters ⊆ `characters.md` | ✅ | 네 인물 모두 `characters.md` 및 개별 profile이 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hook·Seeds·Closing에 유령 인물이 없다. |
| No later-list cast debut | ✅ | 039의 허용 cast인 진우·혁·아이·내부자만 사용한다. 내부자는 이번 설계에서 additive profile로 승인했다. |
| Locations ⊆ Key Locations | ✅ | 세 Scene 모두 `북문서가-본가`의 Key Location에 속한다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `지하 보관실`, `지하 통로 입구`, `가주전-회랑 접속부`가 profile anchor와 정확히 일치한다. |
| No template residue | ✅ | raw placeholder 없음. |
| Prose forecast present | ✅ | 각 Scene에 5종 typed integer formula와 8/9/8-line outline이 있다. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 2,550→Est 2,500, Sc2 2,970→3,000, Sc3 2,540→2,500; 모두 ±20%이며 outline density 안이다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene fields 2,500+3,000+2,500=8,000; header addends 동일. |
| Est. length sum ≥ Scale min | ✅ | 8,000 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 8,000 ≤ 8,000; 상한선이므로 generation에서 반복·확인 루프를 금지한다. |
| Cited staging/profile paths exist | ✅ | `characters/흑풍루-내부자.md`, `stagings/039-내부자-거래.md`, `locations/북문서가-본가.md` 모두 이번 턴에 read OK. |
| Episode List plot | ✅ | Summary의 후계자 제작·아이 보호 대가·회합 장소가 각각 Scene 2 Beat, Scene 3 Beat/Turn에 대응한다. |
| Hook evidence strength | ✅ | Series Hook·Episode Summary·Out·Scene 3 Turn/Transition이 모두 “남궁가의 혼례 행렬이 지나가는 길목”을 같은 강도로 유지한다. |
| Hook scope | ✅ | 마지막에는 장소 후크만 남기고 혼례 습격 결과나 암살자를 선취하지 않는다. |
| No design-paste / meta-only scenes | ✅ | 모든 장면에 관찰 가능한 기록·진입·조건 제시·선택·퇴장이 있다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 038의 표식 확인 → 내부자 거래 제안 → 아이 보호와 좌표 확보의 인과가 선명하다. |
| Scene transitions chain | ✅ | 금속음이 Scene 1→2 진입을 만들고, 거래 쪽지가 Scene 2→3 선택을 만든다. |
| Scene sections complete | ✅ | Index 3행과 full Scene 3개가 일치한다. |
| Generation Readiness | ✅ | 모든 Schema/Consistency 항목이 통과하며 Pending design revision이 없다. |
| Beat concreteness | ✅ | 세 줄 표식, 인계패, 세 단계 설명, 쪽지 교환, 보호선 이동이 구체적이다. |
| Est. length budget | ✅ | 8,000으로 Scale 상한에 걸리지만 정확히 통과하며 압축 Carry를 기록한다. |
| Prose forecast quality | ✅ | dialogue/action/sensory/POV/transition 수가 outline과 대응한다. |
| Episode List scope aligned | ✅ | 040의 혼례 행렬 후크를 장소 정보로만 예고한다. |
| Prior hook addressed | ✅ | Scene 1–2에서 038의 검은 삼중선 인계 표식과 내부자를 직접 이어받는다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected | ✅ | Phase A/B, staging, continuity confirmation과 alignment table이 채워졌다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 중심, P2의 계약·쿠데타 추적, 가족 보호의 불신을 유지한다. |
| Hook internal consistency | ✅ | Summary/Arc/Out/Seeds/closing Turn이 동일한 장소 후크다. |
| Characters from architecture | ✅ | 네 profile을 읽었고 내부자는 additive approval 후 인용했다. |
| Profile-backed knowledge / recognition | ✅ | 내부자의 진우 관련 주장은 “원본 부재”를 함께 밝히는 제한된 증언이며, profile의 후계자 제작 관계에 근거한다. |
| Locations from architecture | ✅ | 장소와 세 facet 모두 profile-backed다. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` read OK. |
| Location facets anchors | ✅ | 세 facet 모두 Multi-facet anchors에 정확히 존재한다. |
| Staging blocking | ✅ | 네 cast의 state·이동·props가 staging과 Scene headers에 일치한다. |
| World rules / history | ✅ | 봉인흔·약물은 단서이며 혈통/기억의 확정으로 과장하지 않는다. |
| No improvised entities or silent lore | ✅ | 새 내부자는 profile을 추가했고, 새 장소·규칙은 만들지 않았다. |
| Continuity files used | ✅ | 038 summary와 story-so-far를 직접 반영했다. |
| Character/location state | ✅ | 아이는 탈출-상태, 진우·혁은 base, 장소는 038 직후의 보관실 상태에서 출발한다. |
| Unresolved threads | ✅ | TH-058~062를 Advances/Plants/Holds로 분리했다. |
| No contradiction of released continuity | ✅ | 아이를 다시 넘기지 않고 보호하는 선택은 038의 보호 상태를 강화한다. |

## Design Consistency Gate
- Loaded required artifacts: ✅
- Locations index/path/facets: ✅ — exact Key Location and three exact anchors verified.
- Length / Forecast: ✅ — Sc1 2,550/2,550/2,500 · Sc2 2,970/2,970/3,000 · Sc3 2,540/2,540/2,500; fields/header=8,000; Scale pass.
- Summary: ✅ — every signature clause maps to a concrete Scene 2 or 3 Beat.
- Hook: ✅ — Series Hook「회합 장소는 남궁가의 혼례 행렬이 지나가는 길목이다」; Out and closing Turn carry the same claim.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 내부자가 지금 거래하러 오는 이유를 Scene 1의 금속음과 Scene 2의 조건으로 지연한다. |
| Personal stake present | ✅ | 아이를 물건처럼 다시 넘길 위험이 첫 장면부터 걸려 있다. |
| Episode Out hook | ✅ | 장소가 적힌 쪽지라는 actionable hook이다. |
| Exposition budget respected | ✅ | 후계자 제작 설명은 Scene 2 거래 중 필요한 만큼만 공개한다. |
| Seed discipline | ✅ | Plant 2개·Hint 1개·Hold 목록이 분리되어 있다. |
| Scene-first Key Events | ✅ | 모든 scene이 행동과 선택 중심이다. |
| Sensory-emotional pairing | ✅ | 냉기·바람·종이 끝이 POV 반응과 결합된다. |
| Motifs planned | ✅ | 접힌 종이와 검은 장갑의 손가락이 장면별로 배치된다. |
| Overview signature line | N/A | overview.md에 별도 signature dialogue가 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 내부자 설명은 거래 압박 안에 배치되고, 진우의 선택이 즉시 뒤따른다. |
| Sensory-emotional pairing | ✅ | 세 장면 모두 detail→POV reaction이 명시된다. |
| Dialogue voices + intent | ✅ | 진우의 단정함, 혁의 명분, 내부자의 조건 언어가 분리된다. |
| Reader-discovered meaning | ✅ | 보호의 의미를 선언하지 않고 보호선·쪽지·붕대 이미지에 맡긴다. |
| Antagonist plausibility | ✅ | 내부자는 악의 고백이 아니라 자신의 생존 몫과 제한된 권한으로 거래한다. |
| Closing image specified | ✅ | 쪽지와 아이의 붕대 끝이 겹치는 이미지다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns locked reader attention | ✅ | 038의 표식 직후 금속음과 내부자 진입으로 첫 장면이 열린다. |
| Personal stake matches reader promise | ✅ | 가족 복수극 독자가 기대하는 아이·도현·흑풍루 구조의 충돌을 보호 선택으로 전환한다. |
| Pacing / density fits platform | ✅ | 3 scenes 모두 사건과 조건 변화가 있고 수련·일상 filler가 없다. |
| Out hook pulls this reader | ✅ | 남궁가 혼례 행렬이라는 기존 동맹 축을 다음 사건 좌표로 바꾼다. |
| No alienation | ✅ | 혈통·기억의 확정은 보류하고, 독자가 확인할 수 있는 물증과 거래만 전진한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 성인 남성향 회귀 무협 독자로서 038의 표식 후속과 즉시 사건을 우선 평가한다.
- Strengths: 아이 보호와 다음 회합 장소를 맞바꾸는 선택이 분명하고, 혼례 행렬 후크가 구체적이다.
- Defects: 상한선 분량 설계는 생성 단계에서 확인 대사가 늘어나면 늘어질 위험이 Low → 반복 확인 없이 조건 변화마다 행동을 붙인다.
- Reader impact: 진우가 보호와 선점을 동시에 잡으려는 장면이 다음 회차 클릭 동기를 만든다.

#### Genre Critic
- Stance: 회귀 무협의 물증 선점·흑막 거래·동맹 위기 계약을 본다.
- Strengths: 내부자가 새 악역으로 소비되지 않고 흑풍루의 분산 구조를 보여주는 실무자로 기능한다.
- Defects: 후계자 제작 3단계가 설명문으로 굳을 위험 Low → 인계패와 아이의 신체 반응을 사이에 둔 대화로만 드러낸다.
- Reader impact: 장르 약속인 ‘정보를 먼저 잡아 판을 바꾸는 주인공’이 유지된다.

#### Plot Expert
- Stance: 038 증거에서 040 후크로 가는 인과와 Hook scope를 검증한다.
- Strengths: 표식→내부자→조건 거래→혼례 행렬 길목의 causal chain이 단절되지 않는다. Out은 장소 하나로 통제된다.
- Defects: — (Hook body alignment와 scope를 확인했으며 불일치 없음)
- Reader impact: 후크가 추상적 조직 설명이 아니라 다음 행동 좌표로 변한다.

#### Reader-Editor
- Stance: 회차 단위 판매성, opening tension, closing density를 본다.
- Strengths: 첫 장면부터 금속음, 마지막은 쪽지와 붕대라는 물리적 마감이다.
- Defects: 상한선 forecast가 procedural repetition으로 번질 수 있음 Low → Carry-⑥: 각 확인은 새 권한/대가 변화가 있을 때만 쓴다.
- Reader impact: 중간 설명이 거래 조건을 바꾸는 동안에는 스킴 위험이 낮다.

#### Literary Critic
- Stance: 접힌 종이·장갑·붕대의 이미지가 설명을 대체하는지 본다.
- Strengths: 의미를 주제 문장으로 설명하지 않고 손과 물건의 배치로 남긴다.
- Defects: — (closing image와 Hold 분리가 설계에 있음)
- Reader impact: 가족 보호의 의미를 독자가 판단할 여백이 있다.

#### Character Critic
- Stance: 진우의 보호 선택, 혁의 명분, 내부자의 생존 동기를 본다.
- Strengths: 진우는 아이의 이름·혈통을 거래 대상에서 제외하고, 혁은 거래의 폭력을 공개적으로 거부한다. 내부자의 knowledge claim은 profile의 제한된 관계와 원본 부재 고지에 근거한다.
- Defects: — (주요 행동의 동기와 profile-backed knowledge가 충족됨)
- Reader impact: 진우의 변화가 선언이 아니라 보호선과 인계패 확보로 보인다.

#### Setting/Lore Expert
- Stance: 혈맥계약·약물·흑풍루 인계망의 규칙과 새 장소/인물 추가를 검증한다.
- Strengths: 새 인물 profile과 staging을 먼저 추가했고, 기존 facet만 사용한다. 봉인흔을 혈통으로 확정하지 않는다.
- Defects: — (unciteable facet·silent lore 없음)
- Reader impact: 새로운 후계자 제작 정보가 기존 약 그릇·봉인 단서에 붙어 보여 세계가 확장된다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | 상한선 forecast가 확인 대사 반복으로 늘어질 위험 (Target Reader, Reader-Editor) | Low | No | yes | 독자는 040 후크까지 빠른 사건 전환을 기대하므로 반복 확인을 막는다. 구조 변경이 아닌 생성 제약이다. | ⑥ 각 대화·확인은 인계패·권한·대가가 바뀔 때만 dramatize하고, 반복 질문을 삭제한다. | Carry-⑥ |
| 2 | 후계자 제작 3단계가 설명 블록으로 굳을 위험 (Genre Critic) | Low | No | yes | 구조 이해는 필요하지만 장르 독자는 설명보다 물증과 압박을 원한다. | ⑥ 인계패, 아이의 반응, 혁의 차단 행동 사이에 분산한다. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Approved:** 2025-02-14 — Schema·Consistency·Target Reader를 통과했고, 남은 두 Low finding은 Carry-⑥로 생성 제약화했다.
- **Next:** Stage ⑥ — Manuscript generation
