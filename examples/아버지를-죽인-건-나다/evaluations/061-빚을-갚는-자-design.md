# Design Evaluation: Episode 061 — 빚을 갚는 자

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001의 설계·원고 범위이며 Episode 061 설계 평가 범위를 벗어난다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003의 초기 시리즈 기준이며 이번 회차의 설계 범위를 벗어난다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 거래·추적·대립이 모두 구체화된다. Scene 1은 출입패·봉랍·장부를 추적하고, Scene 2는 빚의 거래를 증언으로 확정하며, Scene 3은 동맹 해지를 공개 선언한다. Out은 측근의 비밀 명령과 혁의 동맹 해지를 다음 회차의 남궁가 공격 직전 상태로 연결한다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | Episode 061은 P2(031–070) 범위이며, 이번 설계에서는 도현의 보호를 면죄부로 확정하지 않고 거래와 책임의 양면을 유지한다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반 시리즈 회수 기준이다. 이번 회차는 혈맥계약의 최종 주관자와 명령 원문을 Hold한다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 미래 지식을 답으로 사용하지 않고 060에서 확보한 출입패·봉랍·기록을 재대조한다. 변한 현재의 증언과 동맹 반응이 사건을 결정한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 시리즈 결말 기준이며 이번 회차에서는 도현의 보호와 책임을 동시에 열어 두는 중간 진전만 수행한다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 실제 원고 분량·문장 밀도는 Stage ⑥/⑦ 범위다. 설계의 Forecast는 Schema에서 별도 검증한다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `episodes/061-빚을-갚는-자.md`의 세 Scene이 모두 표준 메타 라인과 고정 순서의 flat bullet fields를 사용한다. |
| No skill/workflow dump after the design | ✅ | 설계 본문 뒤에 워크플로 설명이나 절차 복사가 없고, 짧은 Gate G4/Readiness 기록만 있다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1은 물증 대조, Scene 2는 제한 증언, Scene 3은 공개 결별로 기능과 Turn이 분리된다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 경로: `episodes/061-빚을-갚는-자.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | POV·Location·When·On stage·Staging 및 모든 Key Event 필드가 표준 표기다. |
| Every scene has required meta + bullet fields | ✅ | 세 장면 모두 Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, 7-line Paragraph outline, Unit budget, 단일 Est. length를 가진다. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing 네 명은 세 장면 On stage의 합집합과 일치한다. |
| On stage includes speakers | ✅ | 각 장면의 진우·도현·혁·측근이 Dialogue intent와 Beat의 의도적 행위가 발생하는 장면에 모두 On stage다. |
| Characters ⊆ `characters.md` | ✅ | 네 인물 모두 `characters.md` 카탈로그와 프로필 경로에 존재하며, 이번 평가 턴에 네 프로필을 읽어 확인했다. |
| Summary/Hooks cast alignment | ✅ | Summary·In·Out·Seeds·Closing에 등장하는 인물은 모두 Appearing에 포함된다. |
| No later-list cast debut | ✅ | 네 인물 모두 061 이전부터 카탈로그와 시리즈에 등록되어 있다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | Scene 1은 `북문서가-본가`, Scenes 2–3은 `남궁가-본가`이며 두 slug 모두 Key Locations 표에 있다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `북문서가-본가.md`의 `장로회당 표결단`, `남궁가-본가.md`의 `객청`·`검루 계단`이 각각 Multi-facet anchors에 정확히 포함된다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 단일 canonical episode 파일만 사용한다. |
| No template residue | ✅ | 원시 placeholder나 지시용 중괄호가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 세 장면 모두 5개 허용 단위의 정수형 n×pick formula와 7개 paragraph intent를 가진다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | 세 장면 모두 `dialogue 4×250 + action 3×180 + sensory 2×120 + POV 2×140 + transition 1×80 = 2,140`; 독립 재계산도 1,000+540+240+280+80=2,140이며 Est. 2,100은 ±20%와 outline density 범위에 든다. |
| Dialogue intent vs outline speech | ✅ | 세 장면 모두 Dialogue intent가 있으며 outline은 해당 대화 목적을 실행하는 순서로 설계되어 있다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene fields: 2,100 + 2,100 + 2,100 = 6,300; header addends: 2,100 + 2,100 + 2,100 = 6,300. |
| Est. length sum ≥ Scale min | ✅ | 6,300 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 6,300 ≤ 8,000. |
| Cited staging/profile paths exist | ✅ | `characters/서진우.md`, `characters/서도현.md`, `characters/남궁혁.md`, `characters/남궁가의-측근.md`, `locations/북문서가-본가.md`, `locations/남궁가-본가.md`, `world/혈맥계약과-약그릇.md`를 이번 평가 턴에 모두 read OK. Staging은 모든 장면 `none`이므로 staging profile은 N/A다. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary의 “도현이 남궁가를 구한 대가로 비밀 명령을 받아왔음”은 Scenes 1–2의 물증·증언으로, “혁은 동맹을 해지하려 함”은 Scene 3의 선언으로 실행된다. |
| Hook evidence strength (internal) | ✅ | Series Hook의 비밀 명령·동맹 해지·다음 회차 공격을 Summary, Out, Seeds, Scene 3 Turn/Transition에서 같은 강도로 유지한다. 공격은 현재 회차에서 실행하지 않고 다음 회차의 위협으로만 전달된다. |
| Hook scope (no Out creep) | ✅ | Out은 비밀 명령의 확인과 동맹 해지, 다음 회차 공격의 진입 상태만 포함하며 새 추격·전투·추가 폭로를 실행하지 않는다. |
| No design-paste / meta-only scenes | ✅ | 세 장면 모두 물증 재대조→증언→공개 결별이라는 관찰 가능한 사건을 가진다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | 060의 출입패에서 시작해 빚의 증언을 거쳐 동맹 해지로 인과적으로 상승한다. |
| Scene transitions chain | ✅ | Scene 1의 남궁가 공개 증언 이동이 Scene 2의 객청을 열고, Scene 2의 기록을 든 혁의 이탈이 Scene 3의 검루 계단 결별로 이어진다. |
| Scene sections complete | ✅ | Scene Index의 세 행 모두 완전한 Scene section과 필수 필드를 가진다. |
| Generation Readiness | ✅ | 모든 Schema/Structural Integrity 행이 통과하고 Apply? = yes Pending 행이 없다. |
| Beat concreteness | ✅ | 봉랍 결, 장부 항목, 출입패, 증언 규칙, 동맹 문서, 봉화 신호 등 관찰 가능한 행동이 각 Beat를 지탱한다. |
| Est. length budget | ✅ | 독립 계산 6,300자로 4,000–8,000 범위의 중앙대에 있다. |
| Prose forecast quality | ✅ | dialogue·action·sensory·POV·transition 단위가 각 장면의 대화·추적·반응·handoff와 대응한다. |
| Episode List scope aligned | ✅ | 061의 빚과 동맹 균열을 완료하고 062의 본가 공격은 Hold한다. |
| Prior hook addressed (ep 002+) | ✅ | 060의 피의자·봉랍·빚을 Scene 1에서 즉시 재검증한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A/Phase B, continuity, path, staging 상태를 모두 기록했다. |
| Series / overview tone & arc honored | ✅ | 냉정한 물증 추적과 가족 책임의 감정 누출, P2 동맹 균열이 유지된다. |
| Episode List Summary / Hook to Next honored | ✅ | Summary와 Hook을 현재 회차의 증언·결별 및 다음 회차 진입 상태로 분리해 실행한다. |
| Hook internal consistency (design surfaces) | ✅ | Summary·Arc·Out·Seeds·Scene 3 Turn이 ‘비밀 명령 확인과 혁의 동맹 해지, 공격은 다음 회차’라는 동일한 증거 강도를 가진다. |
| Characters from architecture; profiles not redefined | ✅ | 네 프로필을 벗어난 외형·핵심 관계·새 인물을 추가하지 않았다. |
| Profile-backed knowledge / recognition | ✅ | 측근-도현의 과거 빚은 측근 프로필 관계와 060 연속성에 있고, 혁-도현 세대 동맹 불신은 혁 프로필에 근거한다. |
| Locations from architecture; profiles not redefined | ✅ | 두 Key Location과 세 citeable facet만 사용했다. |
| Location profile paths readable | ✅ | 두 location profile exact path read OK. |
| Location facets ⊆ Multi-facet anchors | ✅ | 세 facet 모두 해당 프로필의 Multi-facet anchors에 exact match한다. |
| Stagings from episode design; blocking not redefined | ✅ | 각 장면이 다른 장소·상황이므로 `Staging: none`; 좌석·고정 배치를 몰래 만들지 않았다. |
| World rules / history consistent with bible | ✅ | 문서·인장·운송망으로 정보 추적이 가능하다는 규칙만 사용하고, 혈맥계약의 전체 원문은 확정하지 않는다. |
| No improvised entities or silent lore | ✅ | 새 인물·세력·장소·규칙이 없고, 봉화 신호는 다음 위협을 감지하는 비특정 현장 단서로만 쓰인다. |
| Continuity files used (ep 002+) | ✅ | `continuity/story-so-far.md`와 `continuity/060-명단의-동맹-summary.md`만 권위 입력으로 사용했다. |
| Character/location state vs `story-so-far` | ✅ | 진우는 조사자, 도현은 봉랍과 빚을 숨기는 피고, 혁은 동맹을 지키는 후계자에서 결별 선언자로, 측근은 피의자에서 증언자로 이동한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-095·TH-096을 Picks up/Advances에 명시하고 명령 원문·공격 결과는 Plants/Holds로 분리했다. |
| No contradiction of released continuity | ✅ | 060에서 확정된 유출 사실과 도현 봉랍을 되돌리지 않고 동기를 추가로 밝힌다. |
| Conflicts section empty or escalated | ✅ | 설계의 Conflicts / open questions에 충돌 없음과 보류 범위를 명시했다. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — required indexes, appearing profiles, used locations, world aspect, and immediate prior continuity loaded.
- Locations: ✅ — index: `북문서가-본가`, `남궁가-본가` ∈ Key Locations; paths: `locations/북문서가-본가.md`, `locations/남궁가-본가.md` read OK; facets: `장로회당 표결단`, `객청`, `검루 계단` ⊆ respective Multi-facet anchors.
- Length / Prose forecast: ✅ — Scene 1 written=2,140; recomputed=2,140; Est=2,100 · Scene 2 written=2,140; recomputed=2,140; Est=2,100 · Scene 3 written=2,140; recomputed=2,140; Est=2,100 · scene fields 2,100+2,100+2,100=6,300; header addends 2,100+2,100+2,100=6,300.
- Episode List Summary: ✅ — `도현이 남궁가를 구한 대가로 비밀 명령을 받아왔음` → Scenes 1–2의 구호 항목·봉랍·측근 증언; `혁은 동맹을 해지하려 함` → Scene 3 Turn.
- Hook to Next / Closing: ✅ — Hook「측근은 도현이 남궁가를 구한 대가로 비밀 명령을 받아왔음을 알아낸다. 혁은 아버지들의 거래를 믿지 못하고 동맹을 해지하려 한다. 흑풍루가 남궁가의 본가를 먼저 공격한다」; Out「측근은 도현이 남궁가를 구한 대가로 비밀 명령을 수행해 왔다고 밝히고, 혁은 동맹을 해지하려 한다. 흑풍루의 남궁가 본가 공격은 다음 회차에서 시작된다」; Scene 3 Turn「혁은 남궁가와 북문서가의 공개 동맹을 해지한다고 선언」. 현재 회차는 공격을 실행하지 않는다.
- Hook scope / internal consistency: ✅ — 현재 회차의 공개 의무는 빚의 확인·동맹 해지이며, 공격은 다음 회차 진입 단서로만 남긴다.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | 측근의 행위가 배신인지 빚 상환인지 즉시 질문한다. |
| Personal stake present | ✅ | 동맹이 끊기면 진우의 정보·병력망과 도현 책임 추적이 동시에 흔들린다. |
| Episode Out hook | ✅ | 혁의 공개 동맹 해지와 남궁가를 향한 다음 공격 신호가 다음 회차 욕구를 만든다. |
| Exposition budget respected | ✅ | 과거 구원 사건은 장부·봉랍·제한 증언으로만 제시하고 전투 회상이나 설명문을 Hold한다. |
| Seed discipline | ✅ | Plant 2개와 Hint 1개, Hold 목록이 분리되어 있다. |
| Scene-first Key Events (all required fields) | ✅ | 세 장면 모두 표준 Key Events를 채웠다. |
| Sensory-emotional on every scene | ✅ | 봉랍 소리, 봉랍 조각의 빛, 운송표의 움직임이 각 POV 반응과 결합한다. |
| Motifs planned across scenes | ✅ | 손바닥과 접힌 운송표의 의미·배치가 표로 고정되어 있다. |
| Overview signature line | N/A | `overview.md`에 이번 회차에 반드시 배치할 고정 시그니처 대사가 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | 물증과 증언은 각 장면의 공개 압박·동맹 선택을 위해서만 제시되며, 도현의 거래 전체는 Hold한다. |
| Sensory-emotional pairing | ✅ | 모든 장면의 감각 단서가 진우의 해석 변화와 짝을 이룬다. |
| Dialogue voices + Dialogue intent | ✅ | 진우의 짧은 검증, 혁의 또렷한 명분, 측근의 짧은 실무 존대, 도현의 낮은 회피가 구분된다. |
| Reader-discovered meaning | ✅ | 구원과 부채의 양면성을 closing image로 남기며 도현의 면죄부를 선언하지 않는다. |
| Antagonist plausibility | ✅ | 측근은 남궁가 생존과 개인 의무 사이의 실무적 선택을 하며 단순 악인으로 처리되지 않는다. |
| Closing image specified | ✅ | 검루 계단 아래 접힌 운송표의 두 번째 면이 구체적으로 지정되어 있다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | 060의 봉랍 물증을 첫 장면에서 즉시 다시 움직여 회귀 무협 독자가 원하는 단서 회수와 심문을 제공한다. |
| Personal stake matches what this reader came for | ✅ | 성인 남성향 복수 독자가 기대하는 아버지의 숨은 거래와 동맹의 실질적 균열이 진우의 작전·세력 판에 직접 영향을 준다. |
| Pacing / density fits platform expectations | ✅ | 3개의 기능 분리된 장면, 6,300자 Forecast, 수련·설명 반복 없음으로 사건 밀도가 유지된다. |
| Out hook makes this reader want the next episode | ✅ | 동맹이 끊긴 직후 남궁가 공격 신호가 감지되어 062의 방어·역공을 기다리게 한다. |
| No alienation of core audience without overview intent | ✅ | 로맨스나 치유 중심으로 이탈하지 않고 복수·거래·대립 축을 유지한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 성인 남성향 회귀·무협·문파 장악물 독자의 즉시 연속 시청 욕구를 기준으로 판정한다.
- Strengths: 060에서 남긴 봉랍이 바로 물증 조사로 회수되고, 도현의 보호가 현재 동맹을 흔드는 대가로 전환된다. 혁의 동맹 해지는 다음 회차의 세력전 기대를 만든다.
- Defects: —
- Reader impact: 단서 회수와 세력 균열이 모두 직접적이어서 이탈 위험이 낮다.

#### Genre Critic
- Stance: 회귀 무협의 선점·사이다·세력 재편 약속을 점검한다.
- Strengths: 진우가 미래 지식이 아니라 두 작전의 후속 기록과 봉랍을 재조립해 진실을 압박한다. 동맹 해지가 외부 공격 직전의 실질적 불리함을 만든다.
- Defects: —
- Reader impact: 장르 독자가 기대하는 증거 역전과 세력 변화가 설명이 아니라 행동으로 발생한다.

#### Plot Expert
- Stance: 060 Hook의 인과, 061의 Summary 실행, 062로의 범위 통제를 점검한다.
- Strengths: 피의자 확정→빚의 물증→제한 증언→공개 동맹 해지의 상승이 명확하다. 공격은 Out에서 감지 단계로만 다루어 062의 사건을 훔치지 않는다.
- Defects: —
- Reader impact: 독자는 빚의 정체를 얻으면서도 명령 원문과 공격의 실체를 계속 추적할 수 있다.

#### Reader-Editor
- Stance: 회차 판매력, 장면 전환, Out의 밀도를 점검한다.
- Strengths: 각 장면의 질문이 다음 장소와 행동을 직접 만든다. Out의 현재 의무는 빚 확인과 동맹 해지로 제한되고, 공격은 보조적인 진입 신호다.
- Defects: —
- Reader impact: 마지막 공개 선언이 정서적 결산과 다음 화 클릭 이유를 동시에 제공한다.

#### Literary Critic
- Stance: 아버지의 보호·부채 모티프와 감각적 종결을 점검한다.
- Strengths: ‘뒤집히지 않는 손바닥’과 ‘접힌 운송표’가 거래의 비가시성을 행동으로 만든다. 도현의 침묵을 면죄부가 아닌 책임의 흔적으로 남긴다.
- Defects: —
- Reader impact: 사건 중심 독자의 속도를 해치지 않으면서 가족 복수극의 윤리적 긴장을 유지한다.

#### Character Critic
- Stance: 도현·혁·측근의 동기와 프로필-backed knowledge를 점검한다.
- Strengths: 측근의 행동은 프로필의 ‘남궁가 생존과 빚 상환’ 충돌을 그대로 실행한다. 혁은 프로필의 명분과 가문 보호를 따라 동맹을 해지하고, 도현은 직접 정당화하지 않는 기존 상태를 지킨다.
- Defects: —
- Reader impact: 누구도 편의적으로 악역·배신자·용서받은 아버지로 납작해지지 않아 갈등의 다음 선택을 기다리게 한다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| — | — | — | No | — | 필수 Schema·연속성·독자·장르·문학·인물 검사가 모두 통과했고 High/Med 결함이 없다. | — | — |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Rationale:** Stage ⑤의 Schema·Structural Integrity, 연속성·인물·장소 경로, Hook 범위, Target Reader 및 필수 비평 관점이 모두 통과했다. Apply? = yes Pending 항목이 없고, Generation-ready ✅이므로 Stage ⑥으로 진행한다.
- **Next:** Stage ⑥ — manuscript generation
