# Design Evaluation: Episode 008 — 지하의 문

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope criterion. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003 scope criterion. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 2의 밀거래 물품 추적과 Scenes 3–4의 기록·경고 발견이 실질적 사건이며, Out은 복수의 시작을 의심하라는 후크를 남긴다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | 필체는 도현의 것으로 보이지만 진위와 의도는 확정하지 않아 복수 대상과 의심을 동시에 유지한다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반/시리즈 회수 기준이며, 이번 회차는 약 그릇 흔적과 전달망만 전진시킨다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 후계자 인장과 문서 판독 지식을 사용하지만 불탄 기록과 변화한 현재 때문에 확정적 예언은 하지 않는다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 시리즈/종결부 기준. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 품질은 Stage ⑥/⑦에서 평가하며, 설계 forecast는 Schema에서 검증한다. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | 네 장면 모두 canonical meta fields와 flat bullet fields를 사용한다. |
| No skill/workflow dump after the design | ✅ | 설계 파일에 workflow 본문이 없다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–4의 발견·Turn·전환·outline이 각각 다르다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 경로는 `episodes/008-지하의-문.md`이다. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | 모든 장면이 요구 표기를 따른다. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | 네 장면에 필수 필드, outline, typed Unit budget, 단일 Est.가 있다. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | On stage는 서진우이며 서도현은 필체 언급만 하는 mention-only다. |
| On stage includes speakers | ✅ | 모든 Dialogue intent가 none이고 익명 음성도 없다. |
| Characters ⊆ `characters.md` | ✅ | 서진우·서도현 모두 `characters.md`와 프로필에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary, Hooks, Seeds, Closing의 인물 참조가 Appearing과 일치한다. |
| No later-list cast debut | ✅ | 후발 데뷔 인물이 없다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | 모든 장면의 `북문서가-본가`가 Key Locations에 있다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `지하 통로 입구`, `지하 보관실`은 위치 프로필의 exact anchors다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 단일 episode 파일이다. |
| No template residue | ✅ | 미완성 brace나 template 지시가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 모든 장면이 다섯 허용 unit type 중 필요한 항목으로 정수 budget을 가진다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 `4×180+3×120+3×140+1×80=1,580`; Sc2 동일하게 `1,580`; Sc3 `4×180+3×120+4×140+1×80=1,720`; Sc4 `3×180+3×120+4×140+1×80=1,540`. 각 written product가 재계산값과 일치하고 Est.는 ±20% 안이다. |
| Dialogue intent vs outline speech | ✅ | 대화가 없는 관찰·판독 설계이며 Dialogue intent가 모두 none이다. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields `1,500+1,500+1,600+1,500=6,100`; header addends `1,500+1,500+1,600+1,500=6,100`. |
| Est. length sum ≥ Scale min (hard) | ✅ | 6,100 ≥ 4,000. |
| Est. length sum ≤ Scale max (hard) | ✅ | 6,100 ≤ 8,000. |
| Cited staging/profile paths exist | ✅ | `characters/서진우.md`, `characters/서도현.md`, `locations/북문서가-본가.md`, `world/무림-세력.md`, `world/혈맥계약과-약그릇.md`를 이 평가 턴에 읽었고 모두 성공했다. Staging은 전 장면 none이라 N/A다. |
| Episode List plot (not a different story) | ✅ | Series Summary의 밀거래 물품·어린 시절 기록·불탄 일부·도현 필체 경고가 각각 Scene 2–4에 구체적으로 대응한다. |
| Hook evidence strength (internal) | ✅ | Series Hook, Summary, Out, Seeds, Scene 4 Turn/Transition이 모두 복수의 시작을 의심하라는 동일한 강도로 정렬된다. |
| Hook scope (no Out creep) | ✅ | 마지막 Out은 하나의 경고 후크만 남기며 추격·추가 폭로를 덧붙이지 않는다. |
| No design-paste / meta-only scenes | ✅ | 네 장면 모두 물리적 발견과 선택의 사건을 진행한다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 인장 접근 → 거래품 → 불탄 기록 → 경고의 인과가 선명하다. |
| Scene transitions chain | ✅ | 금속음이 보관실로, 기록함이 기록 수색으로, 필압 자국이 필체 비교로 이어진다. |
| Scene sections complete | ✅ | Scene Index 네 행 모두 완전한 장면 섹션을 가진다. |
| Generation Readiness | ✅ | Schema의 Length, forecast, cast, path, facet, Hook-body가 모두 통과한다. |
| Beat concreteness | ✅ | 봉인, 포장끈, 발자국, 기록함, 불탄 장, 필압 등 관찰 가능한 행동과 증거가 있다. |
| Est. length budget | ✅ | 올바른 unit 산술과 outline density, 6,100 total이 Scale 안에서 일치한다. |
| Prose forecast quality | ✅ | Unit 수가 각 장면의 Beat·감각·POV·전환과 대응한다. |
| Episode List scope aligned | ✅ | Summary와 Hook을 실행하며 다음 회차의 기록 추적 내용을 미리 해결하지 않는다. |
| Prior hook addressed (ep 002+) | ✅ | Episode 007의 표식 입구와 첫 계단을 Scene 1에서 즉시 탐색한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Episode 007의 인장·입구·도현의 미공개 시험 목적과 continuity set이 기록되어 있다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 중심 추적과 제한된 부자 감정이 유지된다. |
| Episode List Summary / Hook to Next honored | ✅ | Summary의 네 핵심 요소와 Hook이 Scene 2–4 및 Out에 반영된다. |
| Hook internal consistency (design surfaces) | ✅ | Summary/Arc close/Out/Seeds/Scene 4 Turn·Transition이 같은 경고를 말한다. |
| Characters from architecture; profiles not redefined | ✅ | 진우와 도현의 기존 drive·행동 단서를 사용하고 프로필을 수정하지 않는다. |
| Profile-backed knowledge / recognition | ✅ | 진우의 문서·인장 판독 습관과 도현의 기존 필체·오른손 화상 단서에 기반하되 확정하지 않는다. |
| Locations from architecture; profiles not redefined | ✅ | 본가와 두 지하 facet을 기존 위치 프로필에서 사용한다. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` read OK. |
| Location facets ⊆ Multi-facet anchors | ✅ | 두 facet 모두 프로필의 Multi-facet anchors와 exact match다. |
| Stagings from episode design; blocking not redefined | N/A | 모든 장면이 단독 탐색이며 Staging: none이다. |
| World rules / history consistent with bible | ✅ | 인장·문서·운송망·약 그릇 흔적을 기존 규칙 안에서 판독하고 전체 계약 구조는 보류한다. |
| No improvised entities or silent lore | ✅ | 새 인물·세력·규칙·장소를 추가하지 않는다. |
| Continuity files used (ep 002+) | ✅ | `story-so-far.md`와 Episode 007 summary를 로드했다. |
| Character/location state vs `story-so-far` | ✅ | 진우의 후계자 인장과 본가의 미탐색 통로 상태가 보존된다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-006/009/010을 진전시키고 TH-011을 심으며 기록을 태운 사람 등은 Hold한다. |
| No contradiction of released continuity | ✅ | Episode 007 직후에서 시작하며 잠긴 사실을 되돌리지 않는다. |
| Conflicts section empty or escalated (not ignored) | ✅ | 충돌 없음과 확정하지 않는 범위를 명시한다. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Load / architecture references | ✅ | Appearing/used/cited paths만 로드했고 Staging은 none이다. |
| Episode List alignment | ✅ | Series Summary의 밀거래 물품·어린 시절 기록·불탄 부분·도현 필체 경고가 Scenes 2–4에 대응한다. |
| Continuity alignment | ✅ | 후계자 인장과 표식 통로의 첫 계단에서 이어지고 미해결 목적을 유지한다. |
| Location index / path / facets | ✅ | Index: `북문서가-본가` ∈ Key Locations; path: `locations/북문서가-본가.md` read OK; facets: `지하 통로 입구`, `지하 보관실` ⊆ anchors. |
| Length evidence | ✅ | Sc1 written=1,580; recomputed=1,580; Est=1,500 · Sc2 written=1,580; recomputed=1,580; Est=1,500 · Sc3 written=1,720; recomputed=1,720; Est=1,600 · Sc4 written=1,540; recomputed=1,540; Est=1,500 · scene fields/header sum=6,100. |
| Hook body alignment and scope | ✅ | Hook「경고문은 ‘진우가 믿는 복수의 시작’을 의심하라고 말한다」; Out 동일; Scene 4 Turn은 복수의 출발점 재검증으로 전환; Transition은 같은 경고를 남긴다. |
| Generation-ready gate | ✅ | Schema 전 항목 통과, Pending adjudication 없음. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 지하 문 안에 살해 흔적 또는 복수의 시작 기록이 있는지 묻고 Scene 1에서 즉시 시험한다. |
| Personal stake present | ✅ | 어린 시절 기록이 진우의 피해와 도현의 역할을 동시에 건드린다. |
| Episode Out hook | ✅ | 복수의 출발점 자체를 의심하게 하는 강한 단일 후크다. |
| Exposition budget respected | ✅ | 거래망, 선택적 소실, 필체 판단만 제시하고 동기·진위·전체망은 Hold한다. |
| Seed discipline | ✅ | 거래 보관망과 선택적 소실은 Plant, 필체는 Hint, 복수의 시작은 Hook으로 구분된다. |
| Scene-first Key Events (all required fields) | ✅ | 모든 장면이 물리적 발견과 판단의 완전한 generation brief다. |
| Sensory-emotional on every scene | ✅ | 냄새·그을음·냉기·소리가 진우의 계산과 감정 억제로 연결된다. |
| Motifs planned across scenes | ✅ | 인장/표식과 그을음/접힌 기록이 장면별 Motif touch로 배치된다. |
| Overview signature line | N/A | overview.md에 별도 서명 대사 제약이 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 정보가 단독 설명이 아니라 봉인·위험·물리적 판독 행동을 통해 나온다. |
| Sensory-emotional pairing | ✅ | 매 장면 감각 단서 뒤에 진우의 의심·계산·억제가 배치된다. |
| Dialogue voices + Dialogue intent | ✅ | 단독 탐색이라 대화는 없고, 기록의 문장은 인물의 발화로 처리하지 않는다. |
| Reader-discovered meaning | ✅ | 도현의 의도는 Hold하고 그을음 묻은 경고와 열린 계단으로 의미를 남긴다. |
| Antagonist plausibility | ✅ | 선택적 소실과 이동된 물품이 조작 가능성을 보여 주며 동기를 과잉 설명하지 않는다. |
| Closing image specified | ✅ | 그을음 묻은 손가락, 도현 필체로 보이는 경고, 아래의 어둠이 지정되어 있다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | 새로 얻은 후계자 인장이 비밀 통로를 여는 즉시 사건으로 시작한다. |
| Personal stake matches what this reader came for | ✅ | 물류 증거와 아버지의 흔적이 복수 독자의 기대를 함께 전진시킨다. |
| Pacing / density fits platform expectations | ✅ | 네 단계의 발견과 6,100자 forecast가 사건 중심 웹소설 범위에 맞는다. |
| Out hook makes *this* reader want the next episode | ✅ | 복수의 전제를 흔들지만 진실을 해결하지 않아 다음 회차 독서 동기를 만든다. |
| No alienation of core audience without overview intent | ✅ | 로맨스·수련·치유물이 아니라 침투·증거·가족 복수 사건이다. |

## Design Critique (required personas)

#### Target Reader
- Stance: 회귀 무협, 문파 장악, 물류 증거와 부자 반전을 기대하는 성인 남성향 독자.
- Strengths: 후계자 인장이 즉시 접근권으로 보상되고, 네 번의 물리적 발견이 명확한 독서 추적점을 만든다.
- Defects: 증거실 탐색이 정적으로 보이지 않도록 Stage ⑥에서 노출·위험·시간 압박 중 하나를 각 판독에 유지한다 → Low → Carry-⑥.
- Reader impact: 복수의 쾌감을 보존하면서 아버지에 대한 의심을 다음 회차로 넘긴다.

#### Genre Critic
- Stance: 회귀 무협 복수 연재의 사건 보상과 후크를 평가한다.
- Strengths: 이전 cliffhanger를 실행하고, 밀거래 물품과 가족 기록이라는 장르적 증거 보상을 제공한다.
- Defects: 물품 목록이 중립적 inventory가 되면 속도가 늦어질 수 있다 → Low → Stage ⑥에서 물리적 위협과 선택을 유지한다.
- Reader impact: 탐색 전환도 위협이 살아 있으면 장르 약속을 충족한다.

#### Plot Expert
- Stance: 인과, Summary/Hook 정합성, forecast 산술을 확인한다.
- Strengths: 인장 접근 → 거래품 → 기록 → 경고의 인과가 끊기지 않고 Hook이 body 전체에서 동일하다.
- Defects: 없음. 네 Unit budget의 written product를 독립 재계산해 모두 수정된 상태다.
- Reader impact: 생성 단계가 허위 분량 산술 없이 설계된 사건 밀도를 재현할 수 있다.

#### Reader-Editor
- Stance: 연재 가독성, 장면 전환, Out의 집중도를 평가한다.
- Strengths: 네 장면이 각기 다른 발견을 수행하고 마지막 Out이 단일 의무만 남긴다.
- Defects: Scene 4의 필체 비교·진위 보류·경고·증거 회수가 밀집한다 → Low → Stage ⑥에서 하나의 물리적 문서 행동과 반응으로 묶고 추가 폭로를 만들지 않는다.
- Reader impact: 마지막이 체크리스트가 아니라 날카로운 질문으로 읽히게 한다.

#### Character Critic
- Stance: 진우의 동기와 도현의 제한된 존재감을 프로필 기준으로 평가한다.
- Strengths: 진우의 문서 판독 습관과 도현의 필체·화상 단서가 근거이며, 진우는 곧바로 용서하지 않는다.
- Defects: 필체 유사성을 확정적 진실이나 도현의 의도로 서술하면 안 된다 → Low → Stage ⑥에서 conditional language와 mention-only를 유지한다.
- Reader impact: 도현을 단순 악역으로 확정하지 않으면서 복수 갈등을 보존한다.

#### Literary Critic
- Stance: 모티프와 독자 발견형 의미, 결말 이미지의 실행 가능성을 평가한다.
- Strengths: 인장/표식과 그을음/접힌 기록이 사물의 반복으로 설계되고, 결말이 설명 대신 열린 어둠을 사용한다.
- Defects: 모티프를 ‘역설’이나 ‘지워진 과거’ 같은 작가 설명으로 반복하지 않는다 → Low → Stage ⑥에서 사물·손동작·질감으로 구현한다.
- Reader impact: 권한이 조작된 과거로 이어졌다는 의미를 독자가 스스로 조립할 수 있다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | 증거실 탐색 중 물리적 위협과 선택을 유지한다 (Target Reader, Genre Critic) | Low | No | yes | 독자가 목록식 탐색에서 이탈하지 않도록 사건 압력을 보존한다. | Stage ⑥ Carry-⑥: 각 증거 판독에 노출·위험·시간 압박 또는 결과를 동반한다. | Carry-⑥ |
| 2 | 도현의 필체는 유사성으로만 처리하고 mention-only를 유지한다 (Character Critic) | Low | No | yes | 조기 무죄 판정은 복수 독자의 핵심 긴장을 약화시키므로 진위와 의도를 보류한다. | Stage ⑥ Carry-⑥: 도현의 내면·확정 동기·직접 발화를 추가하지 않는다. | Carry-⑥ |
| 3 | 모티프를 설명어가 아닌 사물과 행동으로 구현한다 (Literary Critic) | Low | No | yes | 설계된 모티프의 정서적 효과를 작가 메타 설명 없이 유지한다. | Stage ⑥ Carry-⑥: 그을음·접힘·인장·표식·손동작으로 재현한다. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

**Evaluation status:** Design Schema and all required persona checks pass. No Apply? = yes finding remains Pending; the three remaining items are generation-only Carry-⑥ constraints. Gate G4 design approval is still a separate user decision.

## Revision Decisions
- Unit budget 산술 오류 4건은 Stage ④에서 수정했다: Sc1/2 `1,500→1,580`, Sc3 `1,640→1,720`, Sc4 `1,460→1,540`.
- Est. scene sum은 `6,100`으로 유지되며 Scale `4,000–8,000` 안에 있다.
- 평가 후 추가 설계 수정은 없다. 원고 생성 시 Carry-⑥ 제약을 적용한다.
