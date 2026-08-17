# Design Evaluation: Episode 023 — 가장 오래된 종이

> Target: `episodes/023-가장-오래된-종이.md`
> Evaluator: Architect
> Gate: G5 design evaluation

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 001화 안에서 딸이 10년 동안 편지를 읽지 않은 개인적 이유와 현재의 개봉 행동이 드러난다. | N/A | 001화 전용 criterion; Episode 023 대신 120통 개봉의 현재 책임을 실행한다. |
| 제1통의 전문이 서간체로 읽히며, 어머니의 목소리는 편지 안에서만 존재한다. | N/A | 제1통/001화 전용 criterion; 이 회차는 제120통의 첫 문장만 다루며 어머니의 목소리는 편지 안에 한정된다. |
| 편지의 종이·잉크·봉투와 현재의 계절 감각이 사건과 결합한다. | ✅ | Scene 1의 종이·잉크·접힘 관찰과 Scene 2의 봉인 개봉이 현재 옛집의 녹나무 냄새·약한 빛·물방울과 결합한다. |
| 다음 회차를 당기는 구체적 질문, 즉 편지의 날짜와 어머니가 남긴 120통의 의도가 남는다. | ✅ | Episode Out와 Scene 3 Turn에서 첫 문장이 서윤의 현재 질문을 되돌리고, 120통의 전체 의도는 Hold로 남긴다. |
| 설명보다 감정과 물성이 앞서고, 마지막은 해설이 아닌 행동·이미지로 닫힌다. | ✅ | 물성 중심의 세 장면과 마지막 연필의 정지 이미지가 해설을 대체한다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | 3개 Scene 모두 정규 meta line과 flat bullet field를 사용한다. |
| No skill/workflow dump after the design | ✅ | 설계 본문은 episode-specific content와 짧은 Gate 블록만 포함한다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–3은 각각 장롱·마당·벽 틈의 고유 기능과 Turn을 가진다. |
| Canonical episode path | ✅ | 실제 경로 `episodes/023-가장-오래된-종이.md`. |
| Field notation | ✅ | 모든 필드가 `**Field:**` 또는 `- **Field:**` 형식이다. |
| Every scene has required fields | ✅ | 각 장면에 POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition, outline, budget, Est.가 있다. |
| Characters Appearing ↔ On stage union | ✅ | Appearing과 세 장면 On stage 모두 한서윤이다. |
| On stage includes speakers | ✅ | 현재 대화 없음; 120통의 문장은 편지 내부 텍스트이며 인물 Dialogue intent가 아니다. |
| Characters ⊆ `characters.md` | ✅ | `한서윤`은 `characters.md`와 `characters/한서윤.md`에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hooks·Seeds·Closing에 한서윤 외 미등록 인물이 없다. |
| No later-list cast debut | ✅ | 신규 인물 없음. |
| Locations ⊆ Key Locations | ✅ | 세 장면 모두 `어머니의-옛집`이며 `locations.md` Key Locations에 등록되어 있다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `나무 장롱 앞`, `마당의 수도가`, `안쪽 벽의 틈`이 profile의 exact anchors다. |
| Nested scene files absent | ✅ | 단일 episode 파일만 사용한다. |
| No template residue | ✅ | 미완성 placeholder 없음. |
| Prose forecast present | ✅ | 세 장면 모두 5종 typed integer formula와 outline을 갖는다. |
| Forecast ↔ Est. cross-check | ✅ | Sc1 `0+540+360+480+80=1,460`, Sc2 `0+720+360+640+80=1,800`, Sc3 `0+360+360+800+80=1,600`; Est는 각각 1,500/1,800/1,600으로 ±20% 및 outline density 안이다. |
| Dialogue intent vs outline speech | ✅ | 세 장면 모두 speech가 없고 Dialogue intent가 `none`이다. |
| Recorded Estimated Length = scene Est. sum | ✅ | `1,500+1,800+1,600=4,900`; header와 scene fields 일치. |
| Est. length sum ≥ Scale min | ✅ | 4,900 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 4,900 ≤ 8,000. |
| Cited staging/profile paths exist | ✅ | `characters/한서윤.md`와 `locations/어머니의-옛집.md`를 이 평가 턴에 읽었고 read OK; 모든 Staging은 none이라 staging profile은 N/A. |
| Episode List plot | ✅ | Summary의 물성 차이·120통 개봉·첫 문장 사건이 각각 Scene 1·2·3의 Beat/Turn에 대응한다. |
| Hook evidence strength | ✅ | Series Hook 「첫 문장이 현재의 질문을 되돌려 보낸다」와 Summary·Out·Scene 3 Turn·Transition이 같은 강도로 유지된다. |
| Hook scope | ✅ | 마지막에는 질문 되돌림과 연필의 정지만 있으며 추가 폭로·추격·제3의 의무가 없다. |
| No design-paste / meta-only scenes | ✅ | 각 장면에 손으로 확인·봉인 개봉·문장 앞 멈춤이라는 관찰 가능한 사건이 있다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | 물성 대조 → 통제된 개봉 → 첫 문장 앞의 정지로 긴장이 단계 상승한다. |
| Scene transitions chain | ✅ | Scene 1은 개봉 준비로 Scene 2에, Scene 2는 펼쳐진 첫 면으로 Scene 3에 인과적으로 연결된다. |
| Scene sections complete | ✅ | Index의 세 행 모두 완전한 Scene section을 가진다. |
| Generation Readiness | ✅ | Schema·Consistency·Length·Cast·Path·Facet·Hook body 모두 통과한다. |
| Beat concreteness | ✅ | 손끝, 봉인 가장자리, 접힘, 첫 문장, 멈춘 연필 등 행동이 명시됐다. |
| Est. length budget | ✅ | 재계산 합계 4,900, Scale 통과. |
| Prose forecast quality | ✅ | 물성 관찰·개봉·읽기 정지에 맞는 action/sensory/POV 단위다. |
| Episode List scope aligned | ✅ | 024화의 ‘첫 질문’ 전체 해명이나 025화의 120통 전문을 앞당기지 않는다. |
| Prior hook addressed | ✅ | 022화의 두 봉투와 120통 선개봉 약속을 Scene 1–2에서 직접 실행한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected | ✅ | Prior Design Alignment에 Phase A/B, continuity, staging N/A를 기록했다. |
| Series / overview tone & arc honored | ✅ | 서간체 미스터리, 물성, 절제된 현재 수기, P3의 시작을 유지한다. |
| Episode List Summary / Hook honored | ✅ | 설계 Gate Evidence가 Summary와 Hook의 각 문장을 장면에 매핑한다. |
| Hook internal consistency | ✅ | Summary / Arc close / Out / Seed / Scene 3 Turn이 같은 ‘현재 질문의 되돌림’을 말한다. |
| Characters from architecture; profiles not redefined | ✅ | 한서윤 profile의 관찰 습관·연필·수첩·우회적 감정 표현을 사용한다. |
| Profile-backed knowledge / recognition | N/A | 타인의 정체를 알아보거나 지식을 주장하지 않는다. |
| Locations from architecture; profiles not redefined | ✅ | 옛집과 세 exact facets만 사용한다. |
| Location profile paths readable | ✅ | `locations/어머니의-옛집.md` read OK. |
| Location facets ⊆ anchors | ✅ | 세 facet이 profile의 Multi-facet anchors에 정확히 있다. |
| Stagings from episode design | ✅ | 단독 조사 장면이므로 모두 `Staging: none`; staging을 불필요하게 만들지 않았다. |
| World rules / history consistent | ✅ | 문서 물성을 시간의 증거로 읽고 과거를 바꾸는 규칙을 만들지 않는다. |
| No improvised entities or silent lore | ✅ | 새 인물·장소·규칙·조직 없음. |
| Continuity files used | ✅ | 022 summary와 story-so-far를 직접 로드했다. |
| Character/location state vs story-so-far | ✅ | 옛집의 벽 속 봉투와 두 봉투 보존 상태를 유지한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | 두 봉투 순서를 advance, 첫 문장 질문을 plant, 벽 속 봉투와 전문을 hold로 명시했다. |
| No contradiction of released continuity | ✅ | 120통만 개봉하고 벽 속 봉투는 닫아 둔다. |
| Conflicts section empty or escalated | ✅ | `Conflicts / open questions: None`; 후속 공개는 Hold로 정리됐다. |

## Design Consistency Gate
- Loaded required artifacts: ✅
- Locations index / path / facets: ✅
- Length / Prose forecast: ✅ — all written products equal independent recomputation; header sum 4,900.
- Episode List Summary: ✅ — each signature clause has a named Scene Beat.
- Hook to Next / Closing: ✅ — body quotes in the design preserve one returned-question claim at one strength.
- Generation Readiness: ✅

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | “왜 제120통의 종이는 다른 편지들과 다르게 늙었는가?”가 Scene 1에서 즉시 답되지 않는다. |
| Personal stake present | ✅ | 첫 시작점에 손을 대는 책임과 다른 봉투의 목소리를 보류하는 부담이 있다. |
| Episode Out hook | ✅ | 질문이 문장 안으로 되돌아오고 연필이 멈춘다. |
| Exposition budget respected | ✅ | 물성·봉인·개봉 순서만 드러내고 전문·의도는 Hold다. |
| Seed discipline | ✅ | Plant 2개와 Hint 1개로 제한했다. |
| Scene-first Key Events | ✅ | 모든 장면이 관찰 가능한 행위로 설계됐다. |
| Sensory-emotional on every scene | ✅ | 녹나무·물방울·잉크와 서윤의 손/시선 반응이 짝을 이룬다. |
| Motifs planned | ✅ | 종이의 결과 봉인·손끝이 장면별로 배치됐다. |
| Overview signature line | N/A | overview의 signature line은 001화 전용 「이 편지가 네 손에 닿을 때쯤이면」이며 이 회차에 재사용하지 않는다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | 설명은 물성 확인에 한정되고, 개봉과 첫 문장 앞 정지가 긴장을 담당한다. |
| Sensory-emotional pairing | ✅ | 세 장면 모두 구체적 감각에서 서윤의 행동/멈춤으로 넘어간다. |
| Dialogue voices + intent | ✅ | 현재 대화 없이 편지 내부 문장만 제한적으로 사용하도록 설계했다. |
| Reader-discovered meaning | ✅ | 운명론적 해설은 Hold하고, 연필의 정지 이미지로 의미를 남긴다. |
| Antagonist plausibility | N/A | 대립 인물이 등장하지 않는 물성·개봉 문턱 회차다. |
| Closing image specified | ✅ | 첫 문장 위에 멈춘 연필을 명시했다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | 022화의 봉투 보류를 곧바로 종이의 차이라는 물성 질문으로 이어 간다. |
| Personal stake matches reader expectation | ✅ | 가족 미스터리의 핵심인 읽기의 책임과 원망의 흔들림을 개봉 행위에 결합한다. |
| Pacing / density fits platform expectations | ✅ | 3개의 짧고 분명한 행동 단위, 4,900자 forecast, 설명 지연이 균형을 이룬다. |
| Out hook makes this reader want next episode | ✅ | 첫 문장의 실제 내용과 ‘첫 질문’의 정체를 024화로 넘긴다. |
| No alienation of core audience | ✅ | 문학형 독자가 기대하는 물성·감정·여운을 유지하며 과잉 폭로를 피한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 20–40대 한국어 문학형 웹소설 독자의 연속 읽기와 정서적 잔향을 우선한다.
- Strengths: 022화의 두 봉투를 잊지 않고, 독자가 기다린 120통 개봉을 실제 행위로 제공한다. 마지막 질문은 024화로 이어질 구체적 궁금증이다.
- Defects: —
- Reader impact: 물성 관찰이 단순 지연이 아니라 봉인 개봉의 책임으로 전환되어 이탈 위험이 낮다.

#### Genre Critic
- Stance: 가족 미스터리·서간체의 증거성과 지연된 공개 계약을 점검한다.
- Strengths: 종이·잉크·봉인이라는 장르의 증거를 사건화하고, 120통을 열되 전체 해답은 보류한다.
- Defects: —
- Reader impact: ‘오래된 종이’가 제목 장식이 아니라 미스터리의 새로운 증거로 기능한다.

#### Plot Expert
- Stance: 022화의 보존 선택에서 024화의 첫 질문으로 넘어가는 인과와 Hook 강도를 점검한다.
- Strengths: 선개봉 약속 → 물성 차이 → 통제된 개봉 → 현재 질문의 반환이라는 원인·결과가 명확하다. Out은 Summary와 Hook 밖의 의무를 추가하지 않는다.
- Defects: —
- Reader impact: 독자는 120통을 열었다는 보상과 아직 모르는 첫 질문을 동시에 얻는다.

#### Reader-Editor
- Stance: 회차 단위의 판매력, 중간 정체, 마지막 클릭 유인을 점검한다.
- Strengths: 각 장면의 장소와 행동이 달라 반복 감각을 피하고, Scene 3의 정지가 다음 회차의 문턱을 만든다.
- Defects: —
- Reader impact: 개봉을 지나치게 미루지 않으면서도 전문 전체를 소비시키지 않아 다음 회차 기대가 남는다.

#### Literary Critic
- Stance: 물성 모티프가 감정 설명을 대신하고 마지막 이미지가 해설을 피하는지 점검한다.
- Strengths: 종이 결·잉크·연필의 위치가 서윤의 망설임을 드러낸다. ‘어머니가 예견했다’는 의미를 Hold해 감상적 확정을 막았다.
- Defects: —
- Reader impact: 독자는 의미를 설명받기보다 손과 종이의 관계에서 감정을 읽게 된다.

#### Character Critic
- Stance: 서윤의 행동이 profile의 관찰 습관과 현재 arc에 맞는지 점검한다.
- Strengths: 날짜를 정리하려는 습관과 감정 앞에서 연필을 멈추는 행동이 이어진다. 120통을 여는 결정은 방어적 침묵에서 능동적 응답으로 가는 작은 행동이다.
- Defects: —
- Reader impact: 서윤의 변화가 선언이 아니라 개봉 순서와 손의 멈춤으로 보여 신뢰도가 높다.

#### Setting/Lore Expert
- Stance: 물성 증거와 옛집의 공간 사용이 기존 세계 규칙을 넘지 않는지 점검한다.
- Strengths: 기존 Multi-facet anchors만 사용하고, 종이·잉크·봉인을 시간의 흔적으로 읽는다. 벽 속 봉투와 120통을 혼동하지 않는다.
- Defects: —
- Reader impact: 새 설정을 추가하지 않고도 오래된 집과 편지의 시간성이 확장된다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| — | — | — | — | — | Required personas found no High/Med design defect; no silent finding dropped. | — | — |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- Architect approval: ✅ (2026-06-01)
- Rationale: Schema, independent forecast arithmetic, continuity, exact location facets, Hook alignment, and all required persona checks pass; no Pending findings.
- Next: Stage ⑥ manuscript generation.
