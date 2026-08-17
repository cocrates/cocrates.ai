# Design Evaluation: Episode 006 — 죽음의 명부

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope; Episode 006 design-eval에서 재평가하지 않는다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 001~003 scope; 이번 회차의 설계 범위를 벗어난다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 1~3이 명부 판독·전달 경로 추적·장로 대표와의 공개 대립을 실행하고, Scene 3에서 「다음 희생자는 북문서가의 후계자 진우 자신」이라는 Episode 007 Hook을 공개한다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | 이번 회차는 서도현을 등장시키지 않으며, `Tone Notes`에서 그의 내면을 열지 않는다. 서도현을 단순 악역으로 확정하는 Beat도 없다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | P2~P4 후반 회수 기준이며, 이번 회차는 명부·표식·운송망의 중간 추적 범위다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 미래를 자동 예언하지 않고, 이미 확보한 사망 기록·검은 패·백무진의 실제 보고를 대조해 경로를 좁힌다. 작성자·최종 명령자·실제 사망 장소는 Hold한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 시리즈 결말 기준이며 Episode 006 설계 범위를 벗어난다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 분량·밀도는 Stage ⑥/⑦에서 판정한다. forecast 산술은 아래 Schema에서 별도 검증한다. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `episodes/006-죽음의-명부.md`는 단일 파일이며 각 장면이 표준 메타 라인과 flat bullet 필드를 사용한다. |
| No skill/workflow dump after the design | ✅ | 설계 뒤에 `04-episode-design.md` 지침 전문이나 작업용 절차 섹션이 붙지 않았다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | `### Scene 1`, `### Scene 2`, `### Scene 3`이 각각 한 번씩 존재하며 판독·경로 보고·다음 항목 공개의 Beat가 분리된다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 평가 대상은 `episodes/006-죽음의-명부.md`이며 Episode List 제목의 한국어 kebab slug와 일치한다. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | POV/Location/When/On stage/Staging 및 Situation~Est. length가 표준 표기다. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | 세 장면 모두 POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, Est. length를 갖는다. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing 네 명은 세 장면의 On stage 합집합과 일치한다: 서진우·윤태석·백무진·장로-대표. |
| On stage includes speakers | ✅ | 각 장면 Dialogue intent에 명시된 서진우·윤태석·장로-대표·백무진은 해당 장면 On stage에 있다. 군중의 귀속 가능한 발화도 없다. |
| Characters ⊆ `characters.md` | ✅ | `characters.md`와 프로필에서 서진우, 윤태석, 백무진, 장로-대표의 카탈로그 행과 프로필을 모두 확인했다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hooks·Seeds·Closing의 고유 인물은 모두 Appearing에 포함된다. |
| No later-list cast debut | ✅ | 네 인물 모두 Episode 006 이전 Episode List/아키텍처에 존재한다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | 두 장소 표기는 모두 Key Location `북문서가-본가`의 허용된 slug와 그 facet이다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `locations/북문서가-본가.md`의 `Multi-facet anchors`에서 `장로회당 표결단`, `외당 무기고 입구`를 각각 확인했다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 단일 episode 파일 구조이며 nested scene 파일이 없다. |
| No template residue | ✅ | `{placeholder}`와 미완성 템플릿 문구가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 모든 장면에 7개 Paragraph outline과 다섯 허용 단위 중 사용 단위의 정수식이 있다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 `3×260+3×180+2×130+2×150+1×80=1,960`, Sc2 `3×250+3×180+2×120+2×140+1×80=1,890`, Sc3 `3×260+3×180+2×130+2×150+1×90=1,970`; 세 written 값 모두 독립 재계산과 일치하며 Est. 2,000/1,900/2,000은 각각 ±20% 이내이고 outline density band 안이다. |
| Dialogue intent vs outline speech | ✅ | 세 장면 모두 대화 의도가 있고 outline/Beat의 발화와 일치한다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: `2,000 + 1,900 + 2,000 = 5,900`; header addends: `2,000 + 1,900 + 2,000 = 5,900`. |
| Est. length sum ≥ Scale min (hard) | ✅ | 독립 재합산 5,900 ≥ overview Scale min 4,000. |
| Est. length sum ≤ Scale max (hard) | ✅ | 독립 재합산 5,900 ≤ overview Scale max 8,000. |
| Cited staging/profile paths exist | ✅ | Character paths `characters/서진우.md`, `characters/윤태석.md`, `characters/백무진.md`, `characters/장로-대표.md`와 location path `locations/북문서가-본가.md`를 이 턴에 읽어 성공했다. 모든 장면이 `Staging: none`이므로 staging profile은 N/A다. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary 「명부를 통해 진우는 자신이 단순한 배신자가 아니라 흑풍루의 교체 가능한 도구였음을 확인한다」는 Scene 1 Turn으로, 「명부의 다음 이름을 구해 세력 판을 흔든다」는 Scene 2~3의 번호 공개와 Scene 3 Turn으로 실행된다. |
| Hook evidence strength (internal) | ✅ | `series.md` Hook 「다음 희생자의 이름은 북문서가의 후계자 진우 자신이다」, Episode Out 「명부의 다음 희생자는 북문서가의 후계자 진우 자신이다」, Episode Arc의 마지막 공개, Seed `진우 자신의 다음 희생자 항목`, Scene 3 Turn의 「명부의 다음 희생자 항목에 ‘북문서가 후계자 서진우’가 적혀 있음」이 모두 같은 강도로 물리적 공개를 말한다. |
| Hook scope (no Out creep) | ✅ | 마지막 Transition은 이름 공개와 호송 경로 실행·공개 답변 요구만 남긴다. 추격, 새 세력 도착, 두 번째 reveal은 추가하지 않으며 독립적 다음 회차 의무도 과밀하지 않다. |
| No design-paste / meta-only scenes | ✅ | 세 장면 모두 문서 판독·번호 대조·봉인 공개라는 관찰 가능한 사건과 Turn이 있다. 동일 Unit budget/outline을 붙여넣지 않았다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 개인 사망 기록의 의미 해석 → 외당·내당 전달 경로 좁히기 → 기록 공개로 세력 선택 강제의 인과가 선명하다. |
| Scene transitions chain | ✅ | Scene 1의 공개 보고 선언이 Scene 2 Situation으로 이어지고, Scene 2의 다음 봉인 확인 명령이 Scene 3 Situation으로 이어진다. 세 장면의 When도 같은 날 아침의 순서를 명시한다. |
| Scene sections complete | ✅ | Scene Index의 세 행 모두 완전한 Scene 섹션을 갖는다. |
| Generation Readiness | ✅ | Schema의 Length/forecast/cast/path/facet/Hook-body 항목에 ❌가 없고, 적용 보류가 필요한 구조 결함도 없다. |
| Beat concreteness | ✅ | 접힘·번호·표식·봉인·호송 경로·기록 공개처럼 Stage ⑥에서 행동으로 dramatize할 수 있는 동작이 지정되어 있다. |
| Est. length budget | ✅ | 독립 재계산 5,900 = header = scene Est. 합계이며 Scale 4,000~8,000 안이다. |
| Prose forecast quality | ✅ | 대화·행동·감각·POV·전환 단위가 각 장면의 7개 outline과 대응한다. |
| Episode List scope aligned | ✅ | Summary의 도구화 확인과 다음 이름 공개를 수행하되, 명부 작성자·최종 명령자 등 후속 회차의 답은 훔치지 않는다. |
| Prior hook addressed (ep 002+) | ✅ | Scene 1이 Episode 005의 미래 사망 기록과 검은 패를 즉시 사용한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | overview, series, Appearing profiles, used location, touched world aspects, staging N/A, Episode 005 continuity set을 설계의 Load confirmation과 Bound to higher design에 반영했다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 추적, P1 내부 장악, 서도현 내면 비공개가 유지된다. |
| Episode List Summary / Hook to Next honored | ✅ | `series.md`의 Episode 006 Summary와 Hook을 각각 Scene 1~3의 명부 해석 및 Scene 3 공개로 연결한다. |
| Hook internal consistency (design surfaces) | ✅ | Summary·Episode Arc close·Out·Seeds·Scene 3 Turn이 모두 ‘다음 희생자=현재의 서진우’라는 동일한 claim을 낸다. |
| Characters from architecture; profiles not redefined | ✅ | 네 인물은 인덱스와 개별 프로필에 존재하며, 진우의 계산 습관·윤태석의 내당식 존대·백무진의 거친 조건 제시·장로 대표의 절차 통제가 프로필과 맞는다. |
| Profile-backed knowledge / recognition | ✅ | 윤태석의 내당 처분 기록 지식, 백무진의 호송 실무 지식, 장로 대표의 절차·장부 통제가 각각 프로필의 Drive/Relationships/Role에 근거한다. 새 인물의 인식이나 무근거 지식은 없다. |
| Locations from architecture; profiles not redefined | ✅ | 두 facet은 `북문서가-본가` profile과 Key Locations 표에 근거한다. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md`를 이 턴에 읽어 성공했다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `장로회당 표결단`과 `외당 무기고 입구`가 profile의 exact anchor labels다. |
| Stagings from episode design; blocking not redefined | N/A | 모든 장면이 서로 다른 상황으로 `Staging: none`이며, 지속 상황의 좌석·배치·소품을 staging으로 인용하지 않는다. |
| World rules / history consistent with bible | ✅ | 명부는 혈맥계약 원문이나 회귀의 확정 해설이 아니라 문서·표식·운송망으로 추적되는 관리 단서로 제한된다. |
| No improvised entities or silent lore | ✅ | 새 캐릭터·장소·facet·규칙·staging을 만들지 않았다. |
| Continuity files used (ep 002+) | ✅ | `story-so-far.md`, `continuity/005-먼저-친-배신자-summary.md`, `unresolved-threads.md`를 읽고 Prior Hook/Threads This Episode에 반영했다. |
| Character/location state vs `story-so-far` | ✅ | 진우의 기록·표식 보유, 백무진의 첫 보고, 윤태석의 미확정 동생 기록, 본가의 공개 심문 후 상태를 유지한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-006/007/008을 Picks up·Advances하고, 작성자·최종 명령자·동생 생사·사망 장소/살해자는 Holds한다. |
| No contradiction of released continuity | ✅ | Episode 005가 확정한 기록 발견·백무진 보고 약속·윤태석의 동생 기록 탐색을 되돌리지 않는다. |
| Conflicts section empty or escalated (not ignored) | ✅ | `Conflicts / open questions`는 None이며, 후속 답변이 필요한 항목은 Hold로 명시되어 있다. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Loaded required artifacts | ✅ | 이 평가 턴에 overview, series, characters index와 Appearing 4 profiles, locations index와 used profile, world-bible, touched world aspects, Episode 005 continuity 3종을 읽었다. Staging은 전 장면 `none`이므로 N/A다. |
| Locations index | ✅ | Sc1/3 `북문서가-본가 + 장로회당 표결단`, Sc2 `북문서가-본가 + 외당 무기고 입구`; 모두 `locations.md` Key Locations의 `북문서가-본가`에 속한다. |
| Locations paths | ✅ | `locations/북문서가-본가.md` read OK; `characters/서진우.md`, `characters/윤태석.md`, `characters/백무진.md`, `characters/장로-대표.md` read OK. |
| Locations facets | ✅ | `locations/북문서가-본가.md`의 Multi-facet anchors에 두 facet이 exact match한다. |
| Length / Prose forecast | ✅ | Sc1 `written=1,960; recomputed=1,960; Est=2,000`; Sc2 `written=1,890; recomputed=1,890; Est=1,900`; Sc3 `written=1,970; recomputed=1,970; Est=2,000`; fields/header `2,000+1,900+2,000=5,900`. |
| Episode List Summary | ✅ | Summary의 「교체 가능한 도구였음을 확인」→ Scene 1 Turn, 「다음 이름을 구해 세력 판을 흔든다」→ Scene 2 Turn/Scene 3 Beat·Turn, 「다음 희생자는 진우 자신」→ Scene 3 Turn. |
| Hook to Next / Closing | ✅ | Hook 「다음 희생자의 이름은 북문서가의 후계자 진우 자신이다」; Out 「명부의 다음 희생자는 북문서가의 후계자 진우 자신이다」; Scene 3 Turn 「명부의 다음 희생자 항목에 ‘북문서가 후계자 서진우’가 적혀 있음」. 동일 강도다. |
| Hook internal consistency | ✅ | Summary·Arc close·Out·Seed·closing Turn 모두 동일한 물리적 공개를 가리킨다. |
| Overview signature lines | ✅ | 고정 서진우 POV, 사건 중심, P1 내부 장악, 핵심 장치 과잉 해설 금지에 맞는다. |
| Profile-backed knowledge | ✅ | 윤태석·백무진·장로 대표의 지식과 행동이 개별 프로필 근거를 가진다. |
| Opening honors prior | ✅ | Scene 1이 Episode 005의 미래 사망 기록과 검은 패에서 시작한다. |
| Continuity states | ✅ | released Episode 005의 본가·인물 상태를 유지한다. |
| Unresolved threads | ✅ | TH-006/007/008은 각각 advance되고, 미확정 사실은 Hold된다. |
| Characters / speakers | ✅ | On stage와 Dialogue intent의 화자가 일치하고 모두 카탈로그에 있다. |
| Locations / When | ✅ | Key Location·exact facet·같은 날 아침의 장면 순서가 명확하다. |
| Length / Generation Readiness | ✅ | 세 장면의 typed formula, exact product, Est., sum, outline density가 모두 통과한다. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 「진우의 사망 기록은 누가 쓴 것이 아니라, 어떤 명령망이 이미 그 죽음을 실행할 순서까지 정해 둔 것인가?」가 즉시 제시되며 답은 Hold된다. |
| Personal stake present | ✅ | 자신의 죽음이 증오가 아니라 교체 절차였다는 사실이 진우의 자존·복수 목표를 직접 위협한다. |
| Episode Out hook | ✅ | 장로회당 표결단에 진우의 이름이 펼쳐진 채 남아 Episode 007 후계자 자리 갈등을 직접 연다. |
| Exposition budget respected | ✅ | 명부 구조·교체 분류·표식/내당/외당 경로만 사건 속에서 제시하고 전체 계약·조직도·도현의 인장은 Hold한다. |
| Seed discipline | ✅ | Plant 1개, Hint 1개, Plant/Hook 1개로 관리하며 작성자·최종 명령자·동생 생사는 Hold한다. |
| Scene-first Key Events (all required fields) | ✅ | 모든 장면에 관찰 가능한 Situation, Beat, Turn, Function, sensory-emotional, dialogue intent, transition이 있다. |
| Sensory-emotional on every scene | ✅ | 기록지 섬유·검집, 목검 소리·기름 냄새, 봉인 종이 섬유·먹물 번짐이 각각 POV 반응과 연결된다. |
| Motifs planned across scenes | ✅ | `이름 옆의 빈칸`은 Scene 1·3, `봉인과 펼침`은 Scene 1~3에 배치되며, 장면별 별도 Motif touch가 없어 Stage ⑥ 전달 제약으로 기록한다. |
| Overview signature line | ✅ | overview의 명시적 대사가 아니라 POV·톤·사건 중심 제약이며 설계의 Tone Notes와 Scene fields에 반영된다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 정보가 문서 대조·보고·번호 확인의 행동 안에 들어가며, 각 장면에 회수하려는 압박과 봉인/압수 시도가 병행된다. |
| Sensory-emotional pairing | ✅ | 세 장면 모두 물질 감각에서 진우 또는 상대의 행동 반응으로 넘어간다. |
| Dialogue voices + Dialogue intent | ✅ | 진우의 짧은 통제형 질문, 윤태석의 내당식 존대, 백무진의 거친 조건 제시, 장로 대표의 절차 언어가 구분되어 있다. |
| Reader-discovered meaning | ✅ | 의미를 해설하지 않고 ‘이름 옆의 빈칸’과 마지막 명부 이미지로 독자가 도구화를 결론내리게 한다. |
| Antagonist plausibility | ✅ | 장로 대표는 단순 악의가 아니라 장부 용어·기밀·장로회 권한·가문 보존을 근거로 기록을 통제하려 한다. |
| Closing image specified | ✅ | 펼쳐진 명부 마지막 줄의 ‘다음 희생자’와 검은 먹물 속 진우의 이름이 지정되어 있다. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | 성인 남성향 회귀 무협 독자가 기대하는 ‘내 죽음의 기록’이라는 즉시성 있는 개인 미스터리로 시작한다. |
| Personal stake matches what this reader came for | ✅ | 진우의 미래 지식이 자신의 폐기 절차를 역으로 이용하는 사이다와 직결된다. |
| Pacing / density fits platform expectations | ✅ | 5,900자 forecast, 세 장면의 명확한 사건 전환, 장황한 계약 해설의 Hold가 연재 호흡에 맞는다. |
| Out hook makes *this* reader want the next episode | ✅ | 후계자 진우 자신의 이름이 공개되어 기록 추적이 곧 후계 경쟁/생존 대립으로 이어진다. |
| No alienation of core audience without overview intent | ✅ | 로맨스·수련물로 이탈하지 않고 추적·공개 압박·개인 생존 후크를 유지한다. |

## Design Critique (required personas)

#### Target Reader
- Stance: 성인 남성향 회귀 무협·문파 장악물 독자의 연속 독서 욕구와 사이다 회수를 우선해 읽었다.
- Strengths: 자신의 미래 사망 기록을 ‘교체 가능한 도구’로 재해석하고, 마지막에 자기 이름을 공개 무기로 쓰는 보상이 명확하다. 독자가 기억해야 할 증거가 기록·표식·번호로 좁혀져 있다.
- Defects: Scene 1의 문서 판독이 설명 대화로 길어질 위험이 있다 → severity Low → Stage ⑥에서 문서의 물리적 차이와 장로 대표의 회수 행동을 먼저 보여 주고, 색인 규칙 설명은 필요한 만큼만 대사화한다.
- Reader impact: 정보가 강압적 설명으로 바뀌면 사이다보다 장부 강의처럼 느껴질 수 있으나, 현재 설계의 행동·감각 단서가 이를 완화한다.

#### Genre Critic
- Stance: 회귀 무협의 선점 정보, 내부 장악, 다음 회차의 공개 후크가 장르 계약을 지키는지 점검했다.
- Strengths: 미래 기록을 만능 예언으로 쓰지 않고 현재의 보고와 대조하며, 진우가 증거를 숨기는 대신 공개해 판을 흔드는 장르적 쾌감이 있다.
- Defects: `교체 가능한 도구`라는 분류가 추상어로만 처리되면 익숙한 비밀조직 문서 후크와 차별성이 약해질 수 있다 → severity Low → Stage ⑥에서 분류가 실제로 사람의 이동·처분·대체 순서를 어떻게 바꾸는지 한 가지 관찰 가능한 항목으로 구현한다.
- Reader impact: 분류의 기능이 보이면 ‘내가 왜 죽는가’가 조직의 냉정한 시스템으로 구체화되어 다음 회차 기대를 높인다.

#### Plot Expert
- Stance: Episode List Summary/Hook, 장면 인과, closing scope와 body Hook 강도를 대조했다.
- Strengths: Scene 1의 판독이 Scene 2의 경로 대조를 만들고, Scene 2의 번호 연속성이 Scene 3의 다음 봉인 공개를 정당화한다. Hook은 Summary·Arc·Out·Seed·Turn에서 동일 강도이며 Out creep도 없다.
- Defects: Scene 2의 「다음 호송을 멈추지 말고 경로를 바꾸라」와 Scene 3의 공개 후 답변 요구가 각각 다음 회차 의무로 읽힐 수 있다 → severity Low → Stage ⑥에서 호송 경로 변경은 Scene 3의 진우 이름 공개를 위한 즉각적 실행으로 짧게 처리하고, 주된 closing obligation은 ‘다음 희생자=진우’로 유지한다.
- Reader impact: 마지막 후크가 두 갈래로 분산되면 이름 공개의 충격이 약해질 수 있으나, 설계상 두 번째 요소는 보조 행동으로 통제 가능하다.

#### Reader-Editor
- Stance: 회차 단위 판매력, 정보 밀도, 장면별 전환과 Out의 과밀 여부를 점검했다.
- Strengths: 세 장면이 판독→보고→공개로 판매 가능한 단일 질문을 확장하며, 마지막 Transition에 세력 도착·추격 같은 독립 신호를 쌓지 않는다.
- Defects: 네 인물이 세 장면 모두에 반복 등장해 표정·발화가 평평해질 위험이 있다 → severity Low → Stage ⑥에서 Scene 1은 윤태석의 지식과 장로 대표의 압수, Scene 2는 백무진의 출입구 계산과 조건, Scene 3은 세 증인의 공개 압박으로 각 장면의 인물 기능을 분리한다.
- Reader impact: 인물 기능이 분리되면 문서 추적의 반복이 아니라 매 장면 다른 권력 관계로 읽혀 이탈을 줄인다.

#### Character Critic
- Stance: 진우·윤태석·백무진·장로 대표의 프로필 근거, 즉시 행동 동기, 관계 압력을 검토했다.
- Strengths: 윤태석은 동생 기록 때문에 사실의 일부만 말하고, 백무진은 외당 생존 통로 때문에 조건을 먼저 제시하며, 장로 대표는 권한 보존을 위해 봉인을 통제한다. 진우의 검집 두드리기와 상대 손/소매 확인도 프로필에 근거한다.
- Defects: 진우가 ‘교체 가능한 도구’라는 결론을 받아들이는 순간이 내면 설명으로 평탄해질 위험이 있다 → severity Low → Stage ⑥에서 프로필의 검집 두드리기·곧아지는 자세·이름을 숨기지 않는 선택으로 감정을 행동에 남긴다.
- Reader impact: 진우의 상처가 해설이 아니라 통제된 행동의 균열로 보이면 복수 주인공의 냉정함과 개인적 공포가 함께 살아난다.

#### Setting/Lore Expert
- Stance: Episode 006에서 처음 본격화되는 명부·표식·내당/외당 전달망이 기존 세계 규칙을 넘지 않는지 확인했다.
- Strengths: `world-bible.md`의 문서·인장·운송망 추적 가능성, 흑풍루의 책임 분산, 내당·외당 구조를 사건의 증거 대조로 사용한다. 혈맥계약 원문이나 흑풍루 전체 조직도를 성급히 확정하지 않는다. Location facet도 `북문서가-본가`의 exact anchors와 일치한다.
- Defects: 명부의 ‘교체’ 분류가 혈맥계약의 확정 법칙처럼 서술되면 기존의 Hold 범위를 침범할 수 있다 → severity Low → Stage ⑥에서 ‘이번 기록이 보여 주는 관리 분류’로 한정하고, 계약 전체의 작동 원리나 최종 명령 체계로 일반화하지 않는다.
- Reader impact: 세계관의 미스터리를 유지하면서도 이번 회차에서 확실히 얻는 정보가 분리되어 혼란과 과잉 해설을 함께 피한다.

#### Literary Critic
- Stance: 이름·빈칸·봉인·펼침의 모티프가 장면 행동과 closing image로 살아날 가능성을 검토했다.
- Strengths: `이름 옆의 빈칸`과 `봉인과 펼침`이 문서의 물질성, 권력 통제, 마지막 이미지에 연결된다. 독자에게 의미를 선언하지 않고 펼쳐진 이름으로 결론을 발견하게 한다.
- Defects: episode-level Motifs 표에는 장면 배치가 있으나 개별 Scene의 `Motif touch`가 없어 Stage ⑥에서 반복 위치가 누락될 수 있다 → severity Low → generation constraint로 전달하며 Stage ⑥에서 Scene 1/3은 빈칸, Scene 1~3은 봉인과 펼침의 관찰 가능한 소리를 반드시 한 번씩 살린다.
- Reader impact: 모티프가 누락되면 회차의 문서 공포가 기능적 추적으로 축소되지만, 장면별 촉감·소리로 유지하면 마지막 이름의 잔상이 커진다.

## Design Adjudication

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Scene 1의 명부 판독과 `교체 가능한 도구` 설명이 강의처럼 길어질 위험 (Target Reader, Genre Critic) | Low | No | yes | 독자는 문서의 의미보다 즉각적인 위험과 진우의 사이다를 먼저 원한다. | Stage ⑥ 제약: 물리적 표식·회수 시도·행동을 먼저 쓰고, 기능 설명은 최소 대사로 제한한다. | Carry-⑥ |
| 2 | `교체` 분류를 혈맥계약 전체의 확정 법칙으로 확대하지 말 것 (Setting/Lore Expert) | Low | No | yes | 미스터리를 보존하면서 이번 회차의 확정 사실만 전달해야 한다. | Stage ⑥ 제약: 명부의 관리 분류로만 서술하고 작성자·최종 명령자·계약 전체 작동은 Hold한다. | Carry-⑥ |
| 3 | 마지막 후크의 우선순위가 호송 경로 변경·장로 대표의 답변 요구로 분산될 위험 (Plot Expert, Reader-Editor) | Low | No | yes | 다음 회차를 견인하는 핵심은 ‘다음 희생자=진우’이며 나머지는 보조 압력이어야 한다. | Stage ⑥ 제약: 이름 공개를 closing peak로 두고 보조 행동은 짧게 처리한다. | Carry-⑥ |
| 4 | 네 인물의 반복 등장으로 장면별 기능이 평평해질 위험 (Reader-Editor, Character Critic) | Low | No | yes | 같은 장소권과 인물군을 사용하므로 각 장면의 권력 변화가 분명해야 한다. | Stage ⑥ 제약: Scene 1=윤태석 지식/압수, Scene 2=백무진 조건/경로, Scene 3=세 증인의 공개 압박으로 발화 기능을 분리한다. | Carry-⑥ |
| 5 | Motif touch가 개별 Scene fields에 없어서 `이름 옆의 빈칸`과 `봉인과 펼침`이 누락될 위험 (Literary Critic) | Low | No | yes | 모티프는 설계상 유효하지만 현재 Schema 결함은 아니며, 원고에서 관찰 가능한 cue로 전달하면 된다. | Stage ⑥ 제약: 빈칸은 Scene 1·3, 봉인/펼침의 소리·동작은 Scene 1~3에 배치한다. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G3 — Evaluation Review
- Evaluation status: Complete.
- Schema blockers: None.
- Design/architecture revisions required before generation: None.
- Carry-⑥ constraints: 5 low-severity generation constraints listed in Design Adjudication.
- Next gate: User reviews this evaluation and explicitly approves proceeding to Stage ⑥, or requests a return to Stage ④/③.
