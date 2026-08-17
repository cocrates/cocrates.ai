# Design Evaluation: Episode 006 — 날짜의 가장자리

> Target: `episodes/006-날짜의-가장자리.md`
> Target Reader: 20–40대 한국어 문학형 웹소설 독자. 사건보다 문장·감정·물성·여운을 중시하는 성인 독자.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 001화 안에서 딸이 10년 동안 편지를 읽지 않은 개인적 이유와 현재의 개봉 행동이 드러난다. | N/A | 001화 전용 기준이며 006화는 이미 개봉 이후의 분류 회차다. |
| 제1통의 전문이 서간체로 읽히며, 어머니의 목소리는 편지 안에서만 존재한다. | N/A | 제1통 전용 기준이며 006화의 설계 범위는 제6통과 날짜 분류다. |
| 편지의 종이·잉크·봉투와 현재의 계절 감각이 사건과 결합한다. | ✅ | 세 장면 모두 종이·잉크·봉투의 상태가 날짜 판독과 분류 행동을 일으키고, 늦가을의 푸른 빛·찬 타일이 감정 반응과 연결된다. |
| 다음 회차를 당기는 구체적 질문, 즉 편지의 날짜와 어머니가 남긴 120통의 의도가 남는다. | ✅ | 2026년 7월의 두 봉투와 두 번째 날짜의 마지막 획이 구체적 질문으로 남는다. 120통의 최종 의도는 계획대로 Hold다. |
| 설명보다 감정과 물성이 앞서고, 마지막은 해설이 아닌 행동·이미지로 닫힌다. | ✅ | 분류·손 멈춤·연필의 위치가 감정을 운반하며, 마지막은 두 봉투 사이의 연필 이미지로 닫힌다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | 세 장면 모두 표준 메타 라인과 고정 bullet 필드를 사용한다. |
| No skill/workflow dump after the design | ✅ | 설계 본문에 워크플로 설명을 복사하지 않았다. |
| Unique Scene headings; no pasted twin scenes | ✅ | `### Scene 1`–`### Scene 3`이 고유하고 각 장면의 공간·행동·Turn이 다르다. |
| Canonical episode path | ✅ | 실제 경로 `episodes/006-날짜의-가장자리.md`가 제목 slug와 일치한다. |
| Field notation | ✅ | 모든 필드가 `**Field:**` 또는 `- **Field:**` 형식이다. |
| Every scene complete | ✅ | 각 장면에 POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition, outline, Unit budget, Est.가 1회씩 있다. |
| Characters Appearing ↔ On stage | ✅ | 세 장면 On stage의 합집합은 한서윤이며, 윤서영은 편지 속 목소리로만 명시된 등장이다. |
| On stage includes speakers | ✅ | 직접 발화가 없고 `Dialogue intent: none`; 편지 문장은 편지 전문 인용이 아니라 설계상 음성의 기능으로만 취급된다. |
| Characters ⊆ architecture | ✅ | `characters/한서윤.md`, `characters/윤서영.md`가 이번 턴에 읽혔고 카탈로그에 등록되어 있다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hook·Closing에 새 인물명이 없다. |
| Locations ⊆ Key Locations | ✅ | 세 장면의 `서윤의-아파트`는 `locations.md` Key Locations의 slug다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `창가의 낮은 탁자`, `현관 쪽 바닥`, `벽 쪽 수납장`은 `locations/서윤의-아파트.md`의 앵커와 문자 그대로 일치한다. |
| Nested scene files absent | ✅ | 단일 설계 파일이다. |
| No template residue | ✅ | 미완성 placeholder가 없다. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 `0×250+4×180+2×120+3×140+1×80=1,460`, Sc2 `0×250+5×180+2×120+3×140+1×80=1,640`, Sc3 동일 1,640. 각각 독립 재계산값과 written 값이 일치하며 Est 1,500/1,700/1,700은 ±20% 안이다. |
| Recorded Estimated Length | ✅ | scene fields `1,500+1,700+1,700=4,900`; header addends `1,500+1,700+1,700=4,900`. |
| Est. sum ≥ Scale min | ✅ | 4,900 ≥ 4,000. |
| Est. sum ≤ Scale max | ✅ | 4,900 ≤ 8,000. |
| Cited profile paths exist | ✅ | `characters/한서윤.md`, `characters/윤서영.md`, `locations/서윤의-아파트.md`를 이번 턴에 read OK. Staging은 전 장면 `none`이므로 N/A. |
| Episode List plot | ✅ | `series.md`의 006화 Summary인 “희미한 날짜를 ... 분류”가 Sc1–2에, “어떤 달에 ... 두 번”이 Sc3에 구체화된다. |
| Hook evidence strength | ✅ | Summary의 “2026년 7월 ... 두 번”, Out의 “두 봉투”, Sc3 Turn의 “한 달 한 통 전제 붕괴”, Seed의 이중 봉투가 같은 강도로 맞물린다. |
| Hook scope | ✅ | Out은 이중 작성 확인과 두 번째 날짜의 보류만 남기며 추가 추격·폭로를 만들지 않는다. |
| No design-paste / meta-only scenes | ✅ | 세 장면 모두 날짜를 기울이고, 옮기고, 멈추는 관찰 가능한 사건을 갖는다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | 흐린 숫자 판독 → 달력 대조 → 두 번째 7월 봉투 확인의 인과가 선명하다. |
| Scene transitions chain | ✅ | Sc1 탁자에서 바닥으로, Sc2 수납장 모서리 발견으로, Sc3 두 봉투의 병치로 자연스럽게 이어진다. |
| Scene sections complete | ✅ | Scene Index의 3개 행마다 완전한 Key Events 섹션이 있다. |
| Generation Readiness | ✅ | 모든 구조·길이·경로·후크 행이 통과하고 Pending 수정이 없다. |
| Beat concreteness | ✅ | 기울이기, 옮겨 적기, 바닥에 펼치기, 손을 거두기, 나란히 놓기 등 구체 행동이다. |
| Episode List scope aligned | ✅ | 후반 병원·희생·지워진 이름을 Hold로 지켜 범위를 넘지 않는다. |
| Prior hook addressed | ✅ | 005화에서 꺼낼 준비가 된 제6통을 첫 장면의 직접 대상로 삼는다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected | ✅ | Phase A/B와 직전 summary/story-so-far가 Prior Design Alignment에 기록되어 있다. |
| Series / overview tone & arc | ✅ | 현재 수기·물성·절제된 미스터리의 톤을 유지한다. |
| Hook internal consistency | ✅ | Summary, Arc closing, Out, Seed, closing Turn이 모두 7월 이중 작성 확인을 가리킨다. |
| Characters from architecture | ✅ | 기존 두 캐릭터만 사용하고 프로필을 재정의하지 않는다. |
| Profile-backed knowledge / recognition | N/A | 인물의 타인 인식·신원 폭로를 설계하지 않았다. |
| Locations from architecture | ✅ | 기존 아파트와 승인된 세 앵커만 사용한다. |
| Location profile paths readable | ✅ | `locations/서윤의-아파트.md` read OK. |
| Stagings | ✅ | 단독 장면만 있어 `Staging: none`이 적절하다. |
| World rules / history | ✅ | 기록의 물리적 흔적을 시간 증거로 쓰며 새 규칙을 만들지 않는다. |
| Continuity files used | ✅ | 005 summary와 story-so-far를 직접 반영했다. |
| Unresolved threads | ✅ | 빗물의 실제 상황·2016년 공백·미발송 이유를 Hold하고 날짜 균열을 Plant한다. |
| No contradiction of released continuity | ✅ | 제6통, 빗물 자국, 자기보호 인식을 모두 유지한다. |

## Design Consistency Gate
- Loaded required artifacts: ✅
- Locations index/path/facets: ✅ — index `서윤의-아파트`; path `locations/서윤의-아파트.md` read OK; 세 facet 모두 Multi-facet anchors에 있음.
- Length: ✅ — scene fields/header `1,500+1,700+1,700=4,900`; Scale 4,000–8,000.
- Hook: ✅ — body surfaces all preserve “2026년 7월에 편지를 두 번 썼다는 사실을 확인한다” at the same strength.
- Summary: ✅ — Summary clauses map to concrete Beats in Scenes 1–3.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | 흐린 날짜가 읽을 수 없는 것인지 빈자리인지 묻는다. |
| Personal stake present | ✅ | 서윤의 10년 달력과 기억 공백이 분류 행동에 걸려 있다. |
| Episode Out hook | ✅ | 두 봉투의 병치와 읽히지 않은 마지막 획이 다음 읽기를 요구한다. |
| Exposition budget respected | ✅ | 날짜·잉크·봉투·계절 감각만 필요한 장면 안에서 제시된다. |
| Seed discipline | ✅ | Plant 2개, Hint 1개, Hold 4개로 과밀하지 않다. |
| Scene-first Key Events | ✅ | 각 장면에 관찰 가능한 변화가 있다. |
| Sensory-emotional on every scene | ✅ | 빛·타일·종이 냄새가 각 장면의 POV 반응과 결합된다. |
| Motifs planned | ✅ | 가장자리와 옮겨 놓기가 장면 터치로 배치됐다. |
| Overview signature line | N/A | 001화에 배치된 시그니처 문장이다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | 설명은 분류 행동에 종속되고, 숫자 판독의 실패·발견이 장면을 끌고 간다. |
| Sensory-emotional pairing | ✅ | 모든 장면에 물성→서윤 반응이 있다. |
| Dialogue voices + intent | ✅ | 현재 대화는 없고, 편지 속 목소리도 직접 인용으로 과잉 전개하지 않는다. |
| Reader-discovered meaning | ✅ | ‘시간의 모양’은 Hold된 해석으로 남고, 행동이 의미를 운반한다. |
| Closing image specified | ✅ | 두 7월 봉투 사이의 연필과 가려지지 않은 마지막 획이다. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns attention | ✅ | 005화의 물 자국과 흐린 날짜가 즉시 이어진다. |
| Personal stake matches reader | ✅ | 가족 미스터리를 숫자와 기억의 불일치로 체험하게 한다. |
| Pacing / density fits platform | ✅ | 세 장면의 단일 행동선이 압축적이며 4,900자 forecast는 중앙대다. |
| Out hook makes this reader continue | ✅ | ‘왜 7월에 두 번인가’라는 검증 가능한 질문이 남는다. |
| No alienation | ✅ | 후반 진실을 설명하지 않고 현재의 읽기 경험에 집중한다. |

## Design Critique
#### Target Reader
- Stance: 물성과 감정의 작은 균열을 따라가는 성인 독자의 실제 읽기를 기준으로 평가했다.
- Strengths: 7월 두 봉투라는 질문이 추상적 미스터리가 아니라 손으로 셀 수 있는 물건으로 남는다.
- Defects: — (opening, personal stake, Out density 모두 통과)
- Reader impact: 다음 회차에서 두 번째 날짜를 확인하고 싶게 만든다.

#### Genre Critic
- Stance: 가족 미스터리·문학형 연재의 약속과 이 회차의 단서 기능을 점검했다.
- Strengths: 사건을 키우지 않고 날짜·잉크의 이상으로 긴장을 갱신한다.
- Defects: — (Hook scope와 genre promise 모두 통과)
- Reader impact: 과잉 폭로 없이도 조사 욕구가 유지된다.

#### Plot Expert
- Stance: 인과, Hook body alignment, Out scope를 검토했다.
- Strengths: 제6통의 희미한 날짜에서 분류, 이중 7월 발견까지 단계가 누적된다.
- Defects: — (Hook body alignment와 scope creep 없음)
- Reader impact: 독자가 서윤과 같은 순서로 증거를 조립한다.

#### Reader-Editor
- Stance: 회차의 판매력과 마지막 전환의 밀도를 검토했다.
- Strengths: 마지막 Transition이 이중 작성 확인과 날짜 보류로 통제되어 있다.
- Defects: — (closing signals가 2개 이하이며 crowded Out 아님)
- Reader impact: 다음 회차 클릭 이유가 단일하고 구체적이다.

#### Literary Critic
- Stance: 반복 모티프와 해설 억제를 설계 차원에서 검토했다.
- Strengths: ‘가장자리’와 ‘옮겨 놓기’가 손의 동작으로 구현될 수 있다.
- Defects: — (closing image와 reader-discovered meaning이 설계됨)
- Reader impact: 독자가 시간의 균열을 설명이 아닌 배열로 느낀다.

#### Character Critic
- Stance: 서윤의 회피와 분류 행동이 프로필의 성향에서 나오는지 검토했다.
- Strengths: 감정을 확정하지 않고 날짜·상태를 먼저 정리하는 행동이 기존의 관찰적 목소리와 맞다.
- Defects: — (profile-backed knowledge finding 없음)
- Reader impact: 서윤의 방어가 조사 방식으로 변하는 작은 진전을 납득한다.

#### Setting/Lore Expert
- Stance: 물성 규칙과 장소 facet의 근거를 검토했다.
- Strengths: 세 장소 facet이 실제 `Multi-facet anchors`에 있고, 종이·잉크·봉투가 기존 세계 규칙 안에서 기능한다.
- Defects: — (unciteable facet와 silent lore invention 없음)
- Reader impact: 현실적인 기록물의 감각이 미스터리의 신뢰도를 높인다.

## Design Adjudication
| # | Finding | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | — | Low | No | no | 모든 필수 구조·후크·독자 기준을 통과했으며 적용할 결함이 없다. | — | Skip |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- Architect approval: ✅ — 2026-06-01.
- Rationale: 설계 스키마·독립 산술·장소 facet·continuity·Hook evidence와 required persona critique를 모두 통과했다. Pending finding이 없어 Stage ⑥으로 진행한다.
