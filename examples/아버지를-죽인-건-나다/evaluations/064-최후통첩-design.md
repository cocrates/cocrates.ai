# Design Evaluation: Episode 064 — 최후통첩

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001의 초기 도입 기준이며 064 설계 범위를 벗어난다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003의 초기 아크 기준이며 064 설계 범위를 벗어난다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 교환망 두 지점을 역폭파하는 실질적 사건을 Scene 3에서 실행하고, Scene 4에서 흑풍루주 친필 지시 발견 후크를 남긴다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | P1 기준이며 064는 P2에 해당한다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 시리즈 후반 회수 기준이다. 064는 표식패·교환망·친필 지시를 전진시키며 완결 회수하지 않는다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 현재의 표식패·장부·운반 흔적을 조합해 역공을 설계하고, 미래 지식만으로 결과를 확정하지 않는다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 결말 기준이며 064는 부자 명령 충돌을 심화하는 중간 회차다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Stage ⑥/⑦ manuscript criterion. 설계의 Forecast는 Schema에서 별도 검사한다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `episodes/064-최후통첩.md`의 네 Scene이 canonical meta lines와 flat bullet fields를 사용한다. |
| No skill/workflow dump after the design | ✅ | Stage workflow 본문을 설계 파일에 붙여 넣지 않았다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–4 제목과 기능이 각각 다르며 동일한 Scene heading이나 Beat 복사가 없다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 경로 `episodes/064-최후통첩.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | POV/Location/When/On stage/Staging 및 모든 bullet이 요구 표기를 따른다. |
| Every scene has required meta + bullet fields | ✅ | 네 장면 모두 Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, Est. length를 갖는다. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing 네 명은 네 장면 중 최소 한 장면의 On stage에 있으며, Scene 1–4의 On stage union도 동일한 네 명이다. |
| On stage includes speakers | ✅ | 각 장면의 Dialogue intent에서 말하거나 의도적 행동을 하는 인물은 해당 On stage에 있다. 외부 전달자·군중 발화는 없다. |
| Characters ⊆ `characters.md` | ✅ | 서진우·서도현·남궁혁·봉인된 또 다른 아이 모두 Character Catalog와 프로필에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hooks·Seeds·Closing의 proper noun은 Appearing에 포함되고, Out은 흑풍루주를 사건의 주체로만 지칭하며 장면 On stage에 추가하지 않는다. |
| No later-list cast debut | ✅ | 네 인물 모두 064 이전에 등장·승인된 인물이다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | 모든 Scene Location이 Key Location slug `북문서가-본가`에 매핑된다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `지하 보관실`, `지하 통로 입구`, `가주전-회랑 접속부`가 `locations/북문서가-본가.md`의 Multi-facet anchors와 일치한다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 단일 episode 파일만 사용한다. |
| No template residue | ✅ | raw placeholder brace가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 네 장면 모두 5종 unit과 6–8개 paragraph outline을 갖는다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 1,350→1,400; Sc2 1,680→1,700; Sc3 2,320→2,300; Sc4 1,780→1,800으로 모두 ±20%이며 outline density band 안이다. |
| Dialogue intent vs outline speech | ✅ | 대화 unit이 각 outline와 Dialogue intent에 대응하고 `none` 불일치가 없다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene fields `1,400 + 1,700 + 2,300 + 1,800 = 7,200`; header addends `1,400 + 1,700 + 2,300 + 1,800 = 7,200`. |
| Est. length sum ≥ Scale min | ✅ | 7,200 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 7,200 ≤ 8,000; central target band의 상단이지만 800자의 여유가 있다. |
| Cited staging/profile paths exist | ✅ | `stagings/064-교환망-역폭파.md`, 네 캐릭터 프로필, `locations/북문서가-본가.md`, `world/혈맥계약과-약그릇.md`를 이번 평가 입력에서 읽었고 모두 존재한다. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary의 역폭파·도현 저지·진우의 명령 거부·친필 지시 발견이 각각 Scene 2–4의 Beat/Turn으로 실행된다. |
| Hook evidence strength (internal) | ✅ | Series Hook「폭파된 교환망에서 흑풍루주의 친필 지시가 발견된다」, Summary의 친필 지시 일부 발견, Out의 동일 발견, Scene 4 Turn의 금속함 발견, Transition의 첫 줄 노출이 같은 강도다. 상세 설계·서명은 보류한다. |
| Hook scope (no Out creep) | ✅ | Out은 친필 지시 발견 하나로 제한되며 추격·흑풍루주 직접 등장·도현 서명 등 독립 의무를 추가하지 않는다. |
| No design-paste / meta-only scenes | ✅ | 네 장면 모두 표식 대조, 명령 충돌, 동시 폭파, 잔해 조사라는 구체적 사건을 갖는다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | 표식패 해석 → 도현의 저지 → 아이 없는 역폭파 → 친필 증거 발견의 인과가 선형으로 이어진다. |
| Scene transitions chain | ✅ | Scene 1의 도현 등장 → Scene 2의 중단 명령 → Scene 3의 작전 실행 → Scene 4의 금속함 조사로 handoff가 명확하다. |
| Scene sections complete | ✅ | Scene Index 네 행 모두 완전한 Scene section을 가진다. |
| Generation Readiness | ✅ | Schema·Consistency의 Length, cast, path, facet, Hook body가 모두 ✅이며 Pending adjudication이 아직 없다. |
| Beat concreteness | ✅ | 표식패·장부·심지·가짜 상자·금속함의 observable action이 각 장면을 운반한다. |
| Est. length budget | ✅ | 독립 재계산 합계 7,200으로 Scale band를 통과한다. |
| Prose forecast quality | ✅ | dialogue/action/sensory/POV/transition 수가 각 장면 Beat와 outline에 대응한다. |
| Episode List scope aligned | ✅ | 아이를 미끼로 쓰지 않는 역공과 친필 지시 발견을 실행하고, 065의 상세 설계를 훔치지 않는다. |
| Prior hook addressed (ep 002+) | ✅ | Scene 1에서 063 표식패와 아이의 선택권을 즉시 이어 받는다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A/B와 immediate prior continuity가 모두 기록되어 있다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 추적, 가족 명령 충돌, P2 계약·거래망 축을 유지한다. |
| Episode List Summary / Hook to Next honored | ✅ | Summary와 Hook을 본문·Scene Turn에서 동일 강도로 실행한다. |
| Hook internal consistency (design surfaces) | ✅ | Summary / Arc close / Out / Seeds / Scene 4 Turn이 친필 지시 발견으로 일치한다. |
| Characters from architecture; profiles not redefined | ✅ | 네 프로필의 drive·voice·상태를 사용하고 핵심 성격을 바꾸지 않는다. |
| Profile-backed knowledge / recognition | ✅ | 도현은 기존 계약·혈맥계약 당사자이며, 진우는 표식·장부를 현재 사건에서 대조한다. 친필 필체는 기존 문서 추적 축 안에서 확인한다. |
| Locations from architecture; profiles not redefined | ✅ | 기존 본가의 세 citeable facet만 사용한다. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` read OK. |
| Location facets ⊆ Multi-facet anchors | ✅ | 세 facet 모두 profile의 Multi-facet anchors에 등재되어 있다. |
| Stagings from episode design; blocking not redefined | ✅ | Staging은 ④에서 새로 작성했고, 캐릭터 상태·blocking·소품을 episode scenes와 일치시켰다. |
| World rules / history consistent with bible | ✅ | 추적 표식·독 발작·인장·운송망의 기존 규칙만 사용하며 폭발도 지하 범위로 제한한다. |
| No improvised entities or silent lore | ✅ | 새 인물·장소·세력·규칙·고유 무공을 추가하지 않는다. |
| Continuity files used (ep 002+) | ✅ | 063 summary와 story-so-far를 Prior Design Alignment 및 Continuity References에 명시했다. |
| Character/location state vs `story-so-far` | ✅ | 도현은 병세가 드러난 상태, 아이는 보호·선택 주체, 혁은 혈통 가능성을 보류한 상태를 유지한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-101/102를 Picks up/Advances로 배치하고 친필 상세·서명·혈통을 Holds로 보존한다. |
| No contradiction of released continuity | ✅ | 063의 아이를 증거물로 넘기지 않겠다는 원칙을 역공 설계에 반영하며, 063의 사건을 되돌리지 않는다. |
| Conflicts section empty or escalated | ✅ | 충돌 없음. 도현의 중단 이유와 친필 상세는 명시적으로 보류했다. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — Phase A/B selective load와 immediate prior continuity가 완료됐다.
- Locations: ✅ — index `북문서가-본가`; path `locations/북문서가-본가.md` read OK; facets `지하 보관실`, `지하 통로 입구`, `가주전-회랑 접속부` 모두 anchors에 포함된다.
- Length / Prose forecast: ✅ — Sc1 written=1,350; recomputed=1,350; Est=1,400 · Sc2 written=1,680; recomputed=1,680; Est=1,700 · Sc3 written=2,320; recomputed=2,320; Est=2,300 · Sc4 written=1,780; recomputed=1,780; Est=1,800 · scene fields=7,200; header addends=7,200.
- Episode List Summary: ✅ — 「아이를 미끼로 쓰는 대신 교환망을 역으로 폭파」→ Scene 3 Beat; 「도현은 작전을 막으려 하지만 진우는 명령을 거부」→ Scene 2 Turn; 「친필 지시 발견」→ Scene 4 Turn.
- Hook to Next / Closing: ✅ — Hook「폭파된 교환망에서 흑풍루주의 친필 지시가 발견된다」; Out「폭파된 교환망의 금속함에서 친필 지시 일부가 발견된다」; Scene 4 Turn「흑풍루주 친필 지시가 발견」; Transition「친필의 첫 줄 노출」. 발견 강도는 동일하고 상세 해독은 보류된다.

## Engagement Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | 아이를 내놓지 않고 본가 안쪽 교환망을 끊을 수 있는가를 제시하며 Scene 1에서 즉시 행동으로 시작한다. |
| Personal stake present | ✅ | 아이의 생존·선택권과 진우의 회귀 전 이용 구조 반복 여부가 전술 결정에 직결된다. |
| Episode Out hook | ✅ | 친필 지시 발견은 065의 상세 설계 해독을 직접 강제한다. |
| Exposition budget respected | ✅ | 교환망 원리를 표식·장부·운반 흔적에 묶고 전체 계약 원리는 Hold한다. |
| Seed discipline | ✅ | Plant 2개와 Hint 1개로 제한하며 친필의 전체 의미를 설명하지 않는다. |
| Scene-first Key Events (all required fields) | ✅ | 네 장면 모두 사건·전환·감각·대화 의도·전환을 갖는다. |
| Sensory-emotional on every scene | ✅ | 냉기·기름·혈관·심지·폭발·그을린 종이와 POV 반응을 장면마다 배치한다. |
| Motifs planned across scenes | ✅ | 꺼지지 않는 심지와 두 홈의 배치를 각각 2개 이상 장면에 고정했다. |
| Overview signature line | N/A | overview.md에 064에서 반드시 배치해야 할 signature dialogue line이 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | 표식·인계망 설명은 각 장면의 설치·검증 행동에 포함되고, Scene 2–3의 명령·폭파가 정보보다 우세하다. |
| Sensory-emotional pairing | ✅ | 설계된 detail→POV reaction이 네 장면 모두 구체적이다. |
| Dialogue voices + Dialogue intent | ✅ | 진우의 짧은 거부, 도현의 낮은 명령, 혁의 검증, 아이의 반복적 거부가 분리된다. |
| Reader-discovered meaning | ✅ | ‘보호와 통제의 경계’를 Hold하고, 아이가 보호선에 남는 관찰로 의미를 독자가 결론 내리게 한다. |
| Antagonist plausibility | ✅ | 흑풍루는 직접 등장하지 않으며, 교환망과 표식이라는 조직적 작동 방식으로 위협을 설득한다. |
| Closing image specified | ✅ | 그을린 금속함과 붉은 실 봉랍에 눌린 첫 줄로 닫는다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|---|---|---|
| Opening earns the locked reader's attention | ✅ | 성인 남성향 회귀 무협 독자가 기대하는 표식 해독과 즉시 역공을 첫 장면에서 받는다. |
| Personal stake matches what this reader came for | ✅ | 아버지의 명령과 아들의 보호 원칙이 전술 사건 안에서 충돌한다. |
| Pacing / density fits platform expectations | ✅ | 4장면·7,200자 Forecast로 거래망 추적→부자 대치→폭파→후크의 상승 리듬을 만든다. |
| Out hook makes this reader want the next episode | ✅ | 흑풍루주 친필 지시는 065의 조작 설계 폭로를 직접 예고한다. |
| No alienation of core audience without overview intent | ✅ | 로맨스·일상·장기 수련으로 이탈하지 않고 사건·복수·가족 갈등을 유지한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 성인 남성향 회귀 무협·가족 반전 독자의 첫 페이지 유지와 다음 회차 클릭 욕구를 기준으로 판정한다.
- Strengths: 063의 선택 위기를 즉시 전술 사건으로 바꾸고, 아이를 미끼로 쓰지 않는 진우의 선택을 액션과 책임의 쾌감으로 만든다.
- Defects: —
- Reader impact: 도현의 중단 명령과 친필 지시 발견이 다음 회차 독서를 강하게 예약한다.

#### Genre Critic
- Stance: 회귀 무협의 선점 정보, 역이용, 부자 대립, 강한 회차 후크를 검토한다.
- Strengths: 미래 지식이 아니라 현재 증거 조합으로 교환망을 역폭파하며 사이다를 만든다. 흑풍루주를 직접 등장시키지 않고도 조직의 손을 느끼게 한다.
- Defects: —
- Reader impact: 장르 약속을 충족하면서 065의 ‘상세 설계’라는 상위 보상을 남긴다.

#### Plot Expert
- Stance: Summary·Hook·장면 전환·인과와 후크 범위를 검증한다.
- Strengths: 표식패→인계망→두 심지→금속함의 causal chain이 끊기지 않는다. Out은 친필 지시 발견 하나로 제한된다.
- Defects: —
- Reader impact: 독자가 ‘왜 폭파가 가능한가’를 표식·장부·운반 흔적으로 따라갈 수 있다.

#### Reader-Editor
- Stance: 회차 상품성, 장면 밀도, 마지막 전환의 과밀 여부를 검토한다.
- Strengths: 1장의 해석, 2장의 관계 충돌, 3장의 실행, 4장의 증거 후크가 기능 중복 없이 분리된다. 마지막은 발견과 첫 줄 노출에 집중된다.
- Defects: —
- Reader impact: 설명 회차로 늘어지지 않고 매 장면 질문이 바뀌어 스크롤을 유지시킨다.

#### Literary Critic
- Stance: 보호·통제·선택의 모티프가 서술 지시가 아니라 장면 행동으로 살아날지를 본다.
- Strengths: 두 홈과 꺼지지 않는 심지가 선택을 시각적·감각적 행동으로 묶고, 아이의 시선과 도현의 손을 통해 의미를 독자에게 남긴다.
- Defects: —
- Reader impact: 사이다 사건 뒤에도 가족 갈등의 잔상이 남아 단순 폭파 에피소드가 되지 않는다.

#### Character Critic
- Stance: 네 인물의 행동 동기와 프로필 기반 관계 압력을 검토한다.
- Strengths: 진우는 보호 원칙 때문에 명령을 거부하고, 도현은 기존의 보호자·통제자 양면성을 행동으로 유지하며, 혁과 아이의 역할도 검증·선택권으로 분리된다.
- Defects: —
- Reader impact: 진우의 명령 거부가 감정적 반항이 아니라 아크 전환으로 읽힌다.

#### Setting/Lore Expert
- Stance: 기존 장소 facet, 혈맥계약, 인계망의 물리적·규칙적 일관성을 검토한다.
- Strengths: 모든 set은 `북문서가-본가`의 기존 anchors에 고정되고, 폭파 범위를 제한해 세계의 물리성을 지킨다. 새 규칙·장소·세력을 만들지 않는다.
- Defects: —
- Reader impact: 본가가 이미 흑풍루의 손 안쪽이었다는 공포가 표식과 통로로 구체화된다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| — | 필수 비평가 세트에서 High/Med 결함 없음 | — | No | no | 설계가 독자 약속·연속성·구조를 모두 충족한다. 추가 수정은 064의 실행 속도와 065 후크를 약화시킬 수 있다. | — | Skip |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Rationale:** Schema·Design Consistency·Generation Readiness가 모두 통과했고, 필수 비평가 세트에서 적용이 필요한 High/Med 결함이 없었다. 아이 없는 역폭파, 도현의 첫 명령 거부, 제한된 친필 지시 발견을 그대로 Stage ⑥에 전달한다.
- **Next:** Stage ⑥ — manuscript generation
