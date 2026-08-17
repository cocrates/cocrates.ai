# Design Evaluation: Episode 009 — 첫 번째 발췌

## Target Reader
`overview.md` 기준: 20–40대 한국어 문학형 웹소설 독자. 사건보다 문장·감정·물성·여운을 중시하고, 가족 사이의 오해와 뒤늦은 화해를 따라가는 성인 독자.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 001화 안에서 딸이 10년 동안 편지를 읽지 않은 개인적 이유와 현재의 개봉 행동이 드러난다. | N/A | 001화 기준의 완료 조건이며 이번 회차 설계 범위가 아님. |
| 제1통의 전문이 서간체로 읽히며, 어머니의 목소리는 편지 안에서만 존재한다. | N/A | 제1통 전문은 이전 회차 범위이며 이번 회차는 세 통의 발췌를 비교한다. 윤서영의 목소리는 편지 속 문장으로만 제한된다. |
| 편지의 종이·잉크·봉투와 현재의 계절 감각이 사건과 결합한다. | ✅ | Scene 1–2에서 날짜·종이 결·잉크 농도·마른 종이 냄새가 반복 단어의 계절별 차이를 판단하는 행동과 결합한다. |
| 다음 회차를 당기는 구체적 질문, 즉 편지의 날짜와 어머니가 남긴 120통의 의도가 남는다. | ✅ | 세 문장의 반복 단어가 어머니의 의도인지 서윤의 선택인지 미결로 남고, Scene 3에서 첫 열 통을 펼칠 비교 기준을 만든다. |
| 설명보다 감정과 물성이 앞서고, 마지막은 해설이 아닌 행동·이미지로 닫힌다. | ✅ | 의미 설명을 비워 둔 수첩과 역순 더미로 돌아가는 봉투가 마지막 이미지다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only | ✅ | `### Scene 1–3` 각각에 canonical meta lines와 고정 bullet fields가 있다. |
| No skill/workflow dump after the design | ✅ | episode-specific design만 있으며 workflow 전문을 복사하지 않았다. |
| Unique scene headings; no pasted twin scenes | ✅ | 세 장면의 장소 facet·기능·Turn·outline이 서로 다르다. |
| Canonical episode path | ✅ | 실제 gate artifact는 `episodes/009-첫-번째-발췌.md`다. |
| Field notation | ✅ | `**Field:**` 및 `- **Field:**` 표기만 사용한다. |
| Every scene has required fields | ✅ | POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, outline, Unit budget, Est. length가 각 장면에 있다. |
| Characters Appearing ↔ On stage union | ✅ | 현재 장면의 union은 한서윤이며 윤서영은 ‘편지 속 목소리’로 명시된 mention-only다. |
| On stage includes speakers | ✅ | Dialogue intent는 모두 `none`; named speech speaker를 발명하지 않았다. |
| Characters ⊆ `characters.md` | ✅ | 한서윤·윤서영 모두 catalog와 profile이 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hook·Seeds에 쓰인 인물은 두 catalog 인물뿐이다. |
| No later-list cast debut | ✅ | 새 인물 없음. |
| Locations ⊆ Key Locations | ✅ | Sc1–3 모두 `서윤의-아파트`이며 `locations.md` Key Locations에 있다. |
| Location facets ⊆ Multi-facet anchors | ✅ | 세 facet이 profile의 anchors인 `창가의 낮은 탁자`, `현관 쪽 바닥`, `벽 쪽 수납장`과 정확히 일치한다. |
| Nested scene files absent | ✅ | 단일 episode 파일만 사용한다. |
| No template residue | ✅ | 미완성 placeholder 없음. |
| Prose forecast present | ✅ | 세 장면 모두 다섯 유형의 정수형 `n×pick = subtotal`을 갖는다. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 1,520→1,600, Sc2 1,500→1,500, Sc3 1,680→1,700; 모두 ±20%이고 outline density 안이다. |
| Dialogue intent vs outline speech | ✅ | outline에도 발화가 없고 Dialogue intent는 모두 `none`이다. |
| Recorded Estimated Length = scene sum | ✅ | 1,600 + 1,500 + 1,700 = 4,800; header와 scene fields가 일치한다. |
| Est. sum ≥ Scale min | ✅ | 4,800 ≥ 4,000. |
| Est. sum ≤ Scale max | ✅ | 4,800 ≤ 8,000. |
| Cited profile paths exist | ✅ | `characters/한서윤.md`, `characters/윤서영.md`, `locations/서윤의-아파트.md`를 이 평가 턴에 읽었고 모두 read OK다. Staging은 전부 none이라 N/A다. |
| Episode List plot | ✅ | Summary의 세 문장 필사·반복 발견·계절별 의미 차이가 Scene 1–3의 concrete Beats로 실행된다. |
| Hook evidence strength | ✅ | series Hook·Summary·Arc close·Out·Seeds·Scene 3 Turn 모두 반복 단어의 계절별 의미 차이를 발견하되 의도는 확정하지 않는 동일한 강도다. |
| Hook scope | ✅ | Out은 비교 기준을 남기는 한 가지 보조 행동만 포함하며 새 추격·인물·폭로를 추가하지 않는다. |
| No design-paste / meta-only scenes | ✅ | 각 장면에 봉투 선택, 문맥 대조, 보류·배열이라는 독립적인 사건과 Turn이 있다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 의심의 행동화 → 계절별 대조 → 의미 보류와 기준 고정의 3단 상승이다. |
| Scene transitions chain | ✅ | Scene 1의 세 통 선택이 Scene 2의 펼침으로, Scene 2의 문장 대조가 Scene 3의 수첩 고정으로 이어진다. |
| Scene sections complete | ✅ | Index 세 행 모두 full scene section을 가진다. |
| Generation Readiness | ✅ | Stage ⑥이 plot·순서·POV·장소·감각·전환을 추측하지 않고 집필할 수 있다. |
| Beat concreteness | ✅ | 고르기·표시하기·기록하기·대조하기·되돌려놓기라는 관찰 가능한 행동이 있다. |
| Est. length budget | ✅ | 4,800자로 Scale 안이며 Unit product와 Est가 독립적으로 일치한다. |
| Prose forecast quality | ✅ | action/sensory/POV/transition 단위가 각 Beat와 outline에 대응한다. |
| Episode List scope aligned | ✅ | 009 Summary와 Hook을 확대하거나 앞당기지 않는다. |
| Prior hook addressed | ✅ | 제8통 마지막 문장의 의심을 세 문장 비교로 구체화한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected | ✅ | Phase A/B, continuity set, facet preflight가 design의 Prior Design Alignment에 기록됐다. |
| Series / overview tone & arc | ✅ | 현재 수기·물성·절제된 미스터리·P1의 조사 확장을 지킨다. |
| Profile-backed knowledge / recognition | N/A | 인물의 타인 인식·새 신분 폭로를 주장하지 않는다. |
| Locations from architecture | ✅ | 단일 Key Location과 세 citeable facet만 사용한다. |
| Stagings | ✅ | 모든 장면이 한서윤 단독이고 fixed multi-cast blocking이 없어 `none`이 적절하다. |
| World rules / history | ✅ | 읽기는 과거를 바꾸지 않고 현재의 감각과 기억을 바꾼다는 규칙을 지킨다. |
| No improvised entities or silent lore | ✅ | 새 인물·장소·질서가 없다. |
| Continuity files used | ✅ | immediate prior summary와 story-so-far만 사용했다. |
| Character/location state | ✅ | 제8통 역순 더미, 병원 편지 보류, 서윤의 역순 규칙을 되돌리지 않는다. |
| Unresolved threads | ✅ | 제8통 마지막 문장·같은 곳을 Advance하고 병원·희생·지워진 이름 등은 Hold한다. |
| No contradiction of released continuity | ✅ | 기존 사건을 수정하거나 확정하지 않는다. |

## Design Consistency Gate
- Locations: ✅ — index `서윤의-아파트` ∈ Key Locations; path `locations/서윤의-아파트.md` read OK; Sc1–3 facet 모두 Multi-facet anchors에 exact match.
- Length / Prose forecast: ✅ — scene fields `1,600+1,500+1,700=4,800`; header addends `1,600+1,500+1,700=4,800`; written/recomputed pairs are 1,520/1,520, 1,500/1,500, 1,680/1,680.
- Episode List Summary: ✅ — 「세 통의 짧은 문장을 나란히 옮겨 적고」→ Scene 1–2 Beat; 「같은 단어가 계절마다 다른 뜻」→ Scene 2 Turn·Scene 3 기록.
- Hook: ✅ — series Hook「같은 단어가 매번 다른 계절에 쓰였다」; Summary「같은 단어가 계절마다 다른 뜻으로 반복」; Out「같은 단어가 매번 다른 계절에 쓰였다는 사실」; Scene 3 Turn은 의도만 보류하고 반복·계절 차이는 유지한다.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 같은 단어가 동일 감정인지 다른 감정인지 묻는다. |
| Personal stake present | ✅ | 비교를 계속할수록 ‘버려졌다’는 단일 기억이 흔들린다. |
| Episode Out hook | ✅ | 추상적 의문이 아니라 반복 단어와 계절이라는 구체적 문서 단서다. |
| Exposition budget respected | ✅ | 새 설명 없이 세 편지의 물성과 현재 행동만 사용한다. |
| Seed discipline | ✅ | Plant 1개, Hint 1개이며 Hold를 명시했다. |
| Scene-first Key Events | ✅ | 모든 장면에 generation brief 필드가 완성되어 있다. |
| Sensory-emotional on every scene | ✅ | 빛·종이 냄새·유리 반사와 서윤의 인식 변화가 짝을 이룬다. |
| Motifs planned across scenes | ✅ | 빈칸과 계절의 종이를 1–3 장면에 배치했다. |
| Overview signature line | N/A | 001화에 배치된 시그니처 문장으로 이번 회차에 재사용할 필요가 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 편지의 정보보다 고르고 옮기고 보류하는 선택의 긴장이 앞선다. |
| Sensory-emotional pairing | ✅ | 각 장면에 물성 detail→POV reaction이 있다. |
| Dialogue voices + intent | ✅ | 발화를 직접 재현하지 않고 편지 속 목소리의 문장 선택만 다룬다. |
| Reader-discovered meaning | ✅ | 반복을 동일 감정으로 해설하지 않고 독자가 계절별 차이를 조합하게 한다. |
| Antagonist plausibility | N/A | 이번 회차에는 대립 인물이 없다. |
| Closing image specified | ✅ | 수첩과 역순 더미로 돌아가는 봉투로 닫힌다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | 제8통의 의심을 즉시 ‘세 통에서 한 줄씩 고르기’라는 물리적 과제로 전환한다. |
| Personal stake matches reader expectation | ✅ | 가족의 오해가 문장 비교를 통해 흔들리는 내적 비용이 있다. |
| Pacing / density fits platform expectations | ✅ | 3 scenes·4,800자 forecast로 짧은 조사 단위를 완결한다. |
| Out hook makes this reader want next | ✅ | 같은 단어의 계절별 변주와 필사한 한 장이 다음 비교를 구체적으로 예고한다. |
| No alienation of core audience | ✅ | 설명·자극·새 인물 없이 문장과 감정·물성 중심이다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 문장과 물성을 따라가며 다음 회차의 비교를 기다리는 성인 독자의 실제 읽기 기준.
- Strengths: 세 통을 고르는 구체적 과제가 즉시 제시되고, 반복 단어가 가족 기억의 균열과 연결된다.
- Defects: —
- Reader impact: 짧은 단서가 과잉 해설 없이 다음 회차의 행동을 약속한다.

#### Genre Critic
- Stance: 서간체 가족 미스터리의 단서 공정성과 정서적 절제를 점검.
- Strengths: 반복 단어를 암호로 과장하지 않고 종이·날짜·계절의 증거로 제한했다.
- Defects: —
- Reader impact: 장르 독자가 기대하는 문서 단서와 미해결 질문을 충족한다.

#### Plot Expert
- Stance: prior hook, causal scene chain, Hook scope와 body alignment를 점검.
- Strengths: 제8통의 문장 의심이 필사→대조→보류로 자연스럽게 진행된다. Summary·Out·Turn의 Hook 강도가 일치한다.
- Defects: —
- Reader impact: 다음 회차로 넘어갈 단서가 인과적으로 생긴다.

#### Reader-Editor
- Stance: 연재 단위의 시작·밀도·마지막 당김을 점검.
- Strengths: closing Transition이 과밀하지 않고, 단어와 수첩이라는 두 개의 시각적 손잡이가 있다.
- Defects: —
- Reader impact: 조용한 회차지만 장면마다 새 행동이 있어 정체감이 없다.

#### Literary Critic
- Stance: 반복·빈칸·계절 모티프가 설명이 아닌 장면으로 살아날 가능성을 점검.
- Strengths: 의미를 비워 둔 수첩과 계절별 종이가 독자의 추론을 남긴다.
- Defects: 생성 시 모티프가 추상적 해설로 바뀌지 않도록 장면별 물성 touch가 필요하다 → Med → generation constraint.
- Reader impact: 문장 중심 독자는 여운을 얻지만, 모티프를 해설하면 작품의 절제가 무너진다.

#### Character Critic
- Stance: 서윤의 행동 습관·방어와 선택의 변화를 점검.
- Strengths: 날짜를 먼저 적고 의미 설명을 비워 두는 행동이 profile의 정리 습관과 감정 우회 voice를 정확히 확장한다.
- Defects: —
- Reader impact: 서윤의 변화가 선언이 아니라 기록 방식의 변화로 읽힌다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | 반복·빈칸·계절 모티프가 생성에서 해설로 변질될 위험 (Literary Critic) | Med | No | yes | 이 독자는 물성과 여운을 중시하므로 모티프를 설명문으로 풀면 핵심 독자 경험이 약해진다. 다만 장면 구조를 바꿀 필요 없이 생성 제약으로 충분하다. | Stage ⑥ Carry-⑥: 반복 단어의 의미를 해설하지 말고 종이·간격·행동으로만 드러낼 것; Scene 3의 의미 보류를 유지할 것. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- Architect approval: ✅ (2026-06-01)
- Rationale: Schema·continuity·Hook·forecast가 모두 통과했고, 유일한 Med finding은 설계 변경이 아닌 Generation-time Carry-⑥로 adjudicate했다.
- Next: Stage ⑥ manuscript generation with Carry-⑥ constraints.
