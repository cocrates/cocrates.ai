# Design Evaluation: Episode 008 — 역순

> Target: `episodes/008-역순.md`
> Target Reader: 20–40대 한국어 문학형 웹소설 독자.

## Criteria Check
| Criterion | Result | Evidence |
|---|---|---|
| 001화 개인적 이유·개봉 행동 | N/A | 001화 전용. |
| 제1통 전문·어머니 목소리 | N/A | 001화 전용. |
| 종이·잉크·봉투와 계절 감각 | ✅ | 역순 배열과 봉투의 마른 냄새·잉크를 행동으로 연결한다. |
| 날짜와 120통 의도에 대한 질문 | ✅ | 제8통의 마지막 문장이 읽는 순서를 예상한 듯한 질문을 남긴다. |
| 설명보다 감정·물성, 행동·이미지 결말 | ✅ | 손으로 배열하고 마지막 봉투를 더미에 놓는 이미지로 닫힌다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical schema / unique scenes / no dump | ✅ | Scene 1–3 모두 고유한 행동과 표준 필드이며 workflow dump·placeholder가 없다. |
| Canonical path | ✅ | `episodes/008-역순.md`. |
| Cast / speakers | ✅ | On stage는 한서윤이며 Dialogue intent는 none; 편지 문장은 문서 텍스트로만 처리한다. |
| Locations index/path/facets | ✅ | 세 장면 모두 `서윤의-아파트`; `locations/서윤의-아파트.md`와 세 anchors가 일치한다. |
| Forecast arithmetic | ✅ | Sc1 1,640; Sc2 1,720; Sc3 1,640을 독립 재계산했다. Est 1,700+1,800+1,700=5,200이며 written 값과 일치한다. |
| Episode List plot | ✅ | 역순 배열은 Sc1, 제8통의 마지막 문장은 Sc2–3에서 실행된다. |
| Hook alignment / scope | ✅ | Summary·Out·Seed·Sc3 Turn이 모두 ‘역순 독자를 알고 있었던 듯한 문장’이라는 같은 강도를 유지하며, 어머니의 의도 확정은 Hold다. |
| Transitions / generation readiness | ✅ | 수납장→창가 탁자→현관 바닥의 이동과 모든 필드가 generation-ready다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load | ✅ | 007 summary와 story-so-far, appearing/used profiles를 Prior Design Alignment에 기록했다. |
| Tone / world rules | ✅ | 기록물의 시간성과 절제된 현재 수기를 유지한다. |
| Profile-backed knowledge | N/A | 신원 인식·타인 지식 주장이 없다. |
| Continuity / unresolved threads | ✅ | 해안 병원명·7월 이중 편지·날짜 공백을 유지하고 순서 질문만 Plant한다. |
| No silent lore | ✅ | 새 인물·장소·규칙이 없다. |

## Engagement / Literary / Target Reader Checks
| Check | Result | Evidence |
|---|---|---|
| Opening question / stake | ✅ | ‘순서를 정한 사람은 누구인가’가 서윤의 자기보호와 연결된다. |
| Seed / exposition discipline | ✅ | Plant 2, Hint 1, Hold 5로 제한했다. |
| Sensory-emotional / motifs | ✅ | 마른 나무 냄새·창가 빛·봉투 모서리가 역순과 마지막 문턱 모티프로 기능한다. |
| Closing image | ✅ | 제8통을 역순 더미에 놓는 손이다. |
| Target reader retention / Out density | ✅ | 지시인지 우연인지의 단일 질문으로 다음 회차를 당긴다. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Design Critique
#### Target Reader
- Stance: 물성과 순서의 균열을 따라가는 성인 독자 기준.
- Strengths: 읽는 방법 자체가 미스터리의 대상이 된다.
- Defects: — (opening, stake, Out 통과; probe D/E 예정 없음)
- Reader impact: 마지막 문장을 직접 확인하고 싶어진다.

#### Genre Critic
- Stance: 가족 미스터리의 단서 갱신과 폭로 지연을 검토했다.
- Strengths: 순서의 의심을 심되 어머니의 의도를 확정하지 않는다.
- Defects: — (장르 약속과 Scope 통과)
- Reader impact: 메타적이지 않은 문서 미스터리가 유지된다.

#### Plot Expert
- Stance: 007 Hook 회수, body Hook alignment, Out scope를 검토했다.
- Strengths: 병원명 이후 읽기 방식으로 자연스럽게 이동한다.
- Defects: — (Hook body alignment·scope 통과)
- Reader impact: 질문이 분산되지 않는다.

#### Reader-Editor
- Stance: 회차의 배열 행동과 마지막 전환 밀도를 검토했다.
- Strengths: 세 장면이 재배열→독해→유지의 선으로 이어진다.
- Defects: — (closing crowded Out 아님)
- Reader impact: 다음 편에서 문장의 주체를 확인할 이유가 남는다.

#### Literary Critic
- Stance: 뒤집힌 순서와 마지막 문턱의 모티프를 검토했다.
- Strengths: 결론을 말하지 않고 봉투의 위치로 의미를 남긴다.
- Defects: — (closing image가 구체적임)
- Reader impact: 읽기 규칙과 자기보호의 겹침을 독자가 발견한다.

#### Character Critic
- Stance: 서윤의 방어와 능동적 선택을 검토했다.
- Strengths: 순서를 의심하면서도 자신의 규칙을 유지하는 선택이 아크를 진전시킨다.
- Defects: — (profile-backed knowledge 문제 없음)
- Reader impact: 서윤이 수동적 수신자에서 읽기의 주체로 이동한다.

#### Setting/Lore Expert
- Stance: 아파트 facet과 문서 규칙을 검토했다.
- Strengths: 새 set이나 lore 없이 기존 봉투 배열을 확장한다.
- Defects: — (facet·silent lore 문제 없음)
- Reader impact: 현실적인 손의 동작이 문서 미스터리를 지탱한다.

## Design Adjudication
| # | Finding | Severity | Conflict? | Apply? | Rationale | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | — | Low | No | no | 모든 구조·산술·continuity·독자 기준 통과. | — | Skip |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- Architect approval: ✅ — 2026-06-01.
- Rationale: Schema, independent forecast, continuity, Hook evidence와 required persona critique가 통과했고 Pending이 없다.
