# Design Evaluation: Episode 010 — 열 통의 봉투

## Target Reader
20–40대 한국어 문학형 웹소설 독자. 사건보다 문장·감정·물성·여운을 중시하며 가족의 오해와 뒤늦은 화해를 따라가는 성인 독자.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 001화 안에서 딸이 10년 동안 편지를 읽지 않은 개인적 이유와 현재의 개봉 행동이 드러난다. | N/A | 001화 완료 조건이며 이번 회차 범위가 아니다. |
| 제1통의 전문이 서간체로 읽히며, 어머니의 목소리는 편지 안에서만 존재한다. | N/A | 001화 범위다. 이번 회차에서도 윤서영의 목소리는 문서 안에만 둔다. |
| 편지의 종이·잉크·봉투와 현재의 계절 감각이 사건과 결합한다. | ✅ | 세 장면에서 봉투 모서리·날짜 표면·종이빛·늦가을 빛이 수량 오류를 확인하는 행동과 결합한다. |
| 다음 회차를 당기는 구체적 질문, 즉 편지의 날짜와 어머니가 남긴 120통의 의도가 남는다. | ✅ | 열한 번째 무날짜 봉투가 편지인지 빈 봉투인지, 왜 월별 구조 밖에 있는지 남는다. |
| 설명보다 감정과 물성이 앞서고, 마지막은 해설이 아닌 행동·이미지로 닫힌다. | ✅ | 무날짜 봉투를 수납장 안쪽 빈 칸에 세우는 closing image다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical schema / path | ✅ | `episodes/010-열-통의-봉투.md`와 세 개의 canonical scene sections를 사용한다. |
| Required fields / no template residue | ✅ | 각 장면에 POV·Location·When·On stage·Staging 및 모든 bullet fields가 있고 placeholder가 없다. |
| Characters / speakers | ✅ | Appearing은 catalog된 한서윤·윤서영이며, 현재 On stage는 한서윤; Dialogue intent는 모두 `none`이다. |
| Locations / facets / paths | ✅ | `서윤의-아파트`가 Key Location이고 세 facet은 profile의 Multi-facet anchors다. profile path read OK. |
| Forecast ↔ Est. | ✅ | Sc1 1,520/1,520/1,500; Sc2 1,680/1,680/1,700; Sc3 1,680/1,680/1,700; sum 4,900. |
| Episode List plot | ✅ | 열 통 펼침→열한 장 확인→하나의 무날짜 봉투라는 Summary의 모든 절이 Scene 1–3에 concrete beat로 있다. |
| Hook strength / scope | ✅ | Summary·Out·Scene 3 Turn이 무날짜 봉투의 존재만 유지하고 내용·발신인을 과잉 폭로하지 않는다. |
| Continuity | ✅ | Ep 009의 필사 기준·역순 규칙을 이어가며 병원·희생·지워진 이름은 Hold한다. |
| Generation readiness | ✅ | 모든 scene transition, action, sensory-emotional cue, paragraph outline, numeric budget이 완성되어 있다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 범위 설정→반복 검증→이름 붙이기 보류의 상승이다. |
| Scene transitions chain | ✅ | 탁자에서 바닥으로, 바닥의 분리 봉투에서 수납장 격리로 이어진다. |
| Beat concreteness | ✅ | 세기·날짜 확인·봉투 분리·수납이라는 observable act가 있다. |
| Estimated length budget | ✅ | 1,500+1,700+1,700=4,900, Scale 4,000–8,000 안이다. |
| Episode List scope aligned | ✅ | 무날짜 봉투의 원인을 후속 회차로 넘긴다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected | ✅ | Phase A/B와 immediate prior continuity가 Prior Design Alignment에 기록됐다. |
| Profile-backed knowledge | N/A | 새 인물의 지식·인식 주장이 없다. |
| Staging | ✅ | 단독 인물·고정 다인 blocking이 없어 `none`이다. |
| World / tone / no silent lore | ✅ | 날짜·봉투·종이의 시간 증거만 사용한다. |
| Unresolved threads | ✅ | 이중 편지와 반복 단어를 진전시키고 원인은 Hold한다. |

## Design Consistency Gate
- Locations: ✅ — `서윤의-아파트` ∈ Key Locations; `locations/서윤의-아파트.md` read OK; 세 facet exact anchors.
- Length: ✅ — fields/header `1,500+1,700+1,700=4,900`; written/recomputed `1,520/1,520`, `1,680/1,680`, `1,680/1,680`.
- Summary: ✅ — 열 통·열한 장·날짜 없음이 각각 Scene 1, Scene 1–2, Scene 2–3에서 실행된다.
- Hook: ✅ — series Hook「여분의 봉투에는 날짜가 없다」와 Summary·Out·Scene 3 closing의 증거 강도가 일치한다.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question | ✅ | 열 통만 나올 것인지 묻고, Scene 1의 여분 발견으로 즉시 진행한다. |
| Personal stake | ✅ | 수량을 믿는 서윤의 자기보호가 하나 더 많은 봉투 앞에서 흔들린다. |
| Out hook | ✅ | 무날짜 봉투라는 actionable object가 다음 회차를 당긴다. |
| Exposition / seed discipline | ✅ | 수량·물성 중심이며 Plant 1, Hint 1, Hold 목록이 분리되어 있다. |
| Closing image | ✅ | 수납장 안쪽 빈 칸에 세워진 봉투다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Sensory-emotional pairing | ✅ | 빛·바닥 온도·종이빛이 세기와 격리의 심리 변화에 붙는다. |
| Reader-discovered meaning | ✅ | 여분 봉투를 편지/빈 봉투로 규정하지 않고 손동작으로 남긴다. |
| Motifs | ✅ | 열을 세는 손과 날짜 없는 표면이 장면별로 변주된다. |
| Closing overstatement risk | ✅ | 마지막은 봉투의 위치와 안정 여부 확인으로 끝나며 해설이 없다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening retention | ✅ | 독자가 수량을 함께 세게 하는 즉시 과제가 있다. |
| Personal stake | ✅ | 질서와 자기보호가 연결되어 단순 소품 발견을 피한다. |
| Platform density | ✅ | 3 scenes, 4,900자 forecast, 각 장면의 독립 Turn이다. |
| Next pull | ✅ | 날짜 없는 봉투의 정체가 구체적이고 미해결이다. |
| No alienation | ✅ | 새 인물·설명·자극 없이 물성과 감정 중심이다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 문학형 가족 미스터리 독자의 단서 추적과 다음 회차 욕구를 점검했다.
- Strengths: 열 통을 직접 세는 과제가 독자를 장면 안에 넣고 무날짜 봉투를 손에 남긴다.
- Defects: —
- Reader impact: 작은 수량 오류가 앞선 편지의 질서를 흔들어 계속 읽을 이유가 된다.

#### Genre Critic
- Stance: 문서 미스터리의 공정성과 폭로 속도를 점검했다.
- Strengths: 여분 봉투의 존재만 제시하고 정체는 보류해 장르 계약을 지킨다.
- Defects: —
- Reader impact: 후속 반전을 선취하지 않으면서도 새 물리적 단서를 얻는다.

#### Plot Expert
- Stance: 009 Hook에서 010 Hook으로의 인과와 Out scope를 점검했다.
- Strengths: 필사한 기준이 수량 검증으로 직접 사용되고, 마지막 행동은 격리 하나로 좁다.
- Defects: —
- Reader impact: 전 회차의 결과가 장식이 아니라 이번 조사 도구가 된다.

#### Reader-Editor
- Stance: 반복 세기 장면의 진행감과 closing object를 점검했다.
- Strengths: 날짜순·물리적 위치라는 두 세기 방식이 오류를 증명하는 단계가 된다.
- Defects: —
- Reader impact: 반복이 확인 루프가 아니라 수량의 의미를 바꾸는 선택으로 읽힌다.

#### Literary Critic
- Stance: 손·날짜·빈 표면 모티프의 장면화 가능성을 점검했다.
- Strengths: 날짜 없는 봉투가 해설 없이 표면과 위치로 의미를 얻는다.
- Defects: 생성에서 ‘질서의 불안정’을 직접 명명하지 말고 세는 손과 멈춤으로 보여 줄 것 → Med → generation constraint.
- Reader impact: 여운을 중시하는 독자가 결론을 스스로 조립할 수 있다.

#### Character Critic
- Stance: 서윤의 정리 습관과 불안 대응이 연속적으로 보이는지 점검했다.
- Strengths: 오류를 자기 실수로 먼저 돌리고, 마지막에는 봉투를 삭제하지도 확정하지도 않는 선택이 arc를 진전시킨다.
- Defects: —
- Reader impact: 서윤의 능동성이 ‘답을 아는 것’이 아니라 보류할 수 있는 힘으로 읽힌다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | ‘질서의 불안정’을 직접 설명할 위험 (Literary Critic) | Med | No | yes | 핵심 독자는 물성에서 의미를 발견하므로 직접 명명하면 여운이 줄어든다. 설계를 바꿀 필요 없이 집필 제약으로 통제한다. | Stage ⑥ Carry-⑥: 질서·불안을 해설하지 말고 세기, 재확인, 멈춤, 봉투의 위치로 드러낼 것. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- Architect approval: ✅ (2026-06-01)
- Rationale: Schema·forecast·continuity·Hook이 통과했고, Med finding은 설계 변경 없이 Carry-⑥로 adjudicate했다.
- Next: Stage ⑥ manuscript generation.
