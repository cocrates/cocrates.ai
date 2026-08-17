# Design Evaluation: Episode 035 — 떨리는 칼

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | 1화 기준의 시리즈 도입 criterion이며 035 설계 범위가 아니다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | 초반 도입 criterion이며 035 설계 범위가 아니다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 도현의 은폐를 추적해 가환의 약상자를 찾고, 처방의 나머지와 흑풍루 거래 장부라는 다음 추적 목표를 남긴다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | N/A | 035는 P2 구간이다. 다만 설계상 진우의 의심과 도현의 보호 가능성을 동시에 유지한다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 시리즈·후반 회수 criterion이다. 이번 회차는 가환의 약상자와 처방 조각을 Plant/Hint로 전진시킨다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 034에서 본 떨림을 근거로 동선을 역산하지만 병명과 처방 전체는 알지 못하며, 혁의 검증을 받아 증거를 좁힌다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 결말 criterion이며 035에서는 부자 갈등을 심화하는 단서만 다룬다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | manuscript Stage ⑥/⑦ criterion. 설계의 Forecast는 Schema에서 별도 검증했다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | 세 장면 모두 표준 메타 라인과 flat bullet 필드를 사용한다. |
| No skill/workflow dump after the design | ✅ | 장면 설계와 게이트 상태만 있으며 workflow 본문을 복사하지 않았다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1~3의 장소·기능·Turn이 다르며 동일 장면 복제가 없다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 게이트 파일은 `episodes/035-떨리는-칼.md`다. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | 모든 메타·Key Event 필드가 canonical notation이다. |
| Every scene has required meta + bullet fields | ✅ | POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition, outline, Unit, Est.가 각 장면에 하나씩 있다. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing은 진우·도현·혁이며 Scene 1은 세 명, Scene 2~3은 진우·혁으로 union이 정확하다. |
| On stage includes speakers | ✅ | 각 장면의 Dialogue intent 발화자인 진우·도현·혁은 해당 On stage에 포함된다. |
| Characters ⊆ `characters.md` | ✅ | 세 인물 모두 `characters.md`와 프로필 파일에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hook·Closing에 새 인물이나 유령 cast가 없다. |
| No later-list cast debut | ✅ | 035 이전부터 등장한 세 인물만 사용했다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | 세 Location 모두 Key Location `북문서가-본가`에 매핑된다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `외당 무기고 입구`, `가주전 문앞`, `지하 보관실`이 `locations/북문서가-본가.md`의 anchors와 정확히 일치한다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 단일 episode 파일만 사용했다. |
| No template residue | ✅ | 미완성 placeholder와 지시용 중괄호가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 각 장면 7개 paragraph intents와 다섯 유형의 정수 Unit formula가 있다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 3×250+4×180+2×120+2×140+1×80=2,260; Sc2 2×250+5×180+2×120+2×140+1×80=2,440; Sc3 4×250+4×180+2×120+2×140+1×80=2,510. Est.는 각각 2,200/2,400/2,500으로 ±20% 이내이고 7-line outline density band 안이다. |
| Dialogue intent vs outline speech | ✅ | 세 장면 모두 speech가 Dialogue intent와 outline에 대응한다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene fields 2,200+2,400+2,500=7,100; header addends 2,200+2,400+2,500=7,100. |
| Est. length sum ≥ Scale min | ✅ | 7,100 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 7,100 ≤ 8,000. |
| Cited staging/profile paths exist | N/A | 모든 장면이 `Staging: none`; cited character/location/world profiles were read and exist. |
| Episode List plot (not a different story) | ✅ | Series Summary 「도현의 병세를 확인하려 하지만 도현은 몸을 숨긴다」와 「가환이 남긴 약상자에서 ... 처방의 일부를 찾는다」가 각각 Scene 1~2와 Scene 3의 Beat/Turn으로 실행된다. |
| Hook evidence strength (internal) | ✅ | Series Hook 「처방의 나머지는 흑풍루의 거래 장부에 있다」, Episode Out, Scene 3 Turn이 모두 ‘나머지가 거래 장부에 있다’는 동일한 확인 수준이다. |
| Hook scope (no Out creep) | ✅ | 다음 회차 의무는 거래 장부 추적 하나이며 추격·추가 폭로를 덧붙이지 않았다. |
| No design-paste / meta-only scenes | ✅ | 세 장면 모두 관찰·은폐·수색·발견이라는 물리적 사건을 가진다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 결투 후 확인 시도 → 도현의 은폐 → 약상자 수색 → 처방 조각 발견 → 장부 추적으로 인과가 이어진다. |
| Scene transitions chain | ✅ | Scene 1의 회랑 이동이 Scene 2의 가주전 문으로, Scene 2의 지하 흔적이 Scene 3의 보관실 수색으로 이어진다. |
| Scene sections complete | ✅ | Scene Index의 3개 행 모두 완전한 장면 섹션을 가진다. |
| Generation Readiness | ✅ | 모든 Schema·Continuity·Location·Hook·Forecast 행이 통과하며 Pending Apply finding이 없다. |
| Beat concreteness | ✅ | 손을 숨김, 걸쇠 잠금, 문턱 흠집, 약상자 인장, 이중판 처방 조각 등 관찰 가능한 행동이 중심이다. |
| Est. length budget | ✅ | Forecast subtotal 2,260/2,440/2,510과 Est. 2,200/2,400/2,500, 총 7,100이 Scale 안에 있다. |
| Prose forecast quality | ✅ | 대화·행동·감각·POV·전환 수가 각 Beat와 Dialogue intent에 대응한다. |
| Episode List scope aligned | ✅ | 035 Summary와 Hook의 사건 범위를 넘지 않는다. |
| Prior hook addressed (ep 002+) | ✅ | Scene 1이 034의 멈춘 칼과 떨리는 검지를 즉시 다룬다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A/B와 연속성 파일, 사용 프로필이 모두 기록되어 있다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 추적과 가족 갈등의 감정 누출을 유지한다. |
| Episode List Summary / Hook to Next honored | ✅ | Summary·Hook 원문 수준의 사건을 장면에 배치했다. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Arc, Seeds, Out, Scene 3 Turn이 처방의 나머지와 거래 장부라는 한 claim을 공유한다. |
| Characters from architecture; profiles not redefined | ✅ | 인물의 습관·말투·상태를 프로필 범위 안에서 사용했다. |
| Profile-backed knowledge / recognition | ✅ | 진우의 손 관찰 습관, 혁의 검증 성향, 도현의 오른손 감추는 습관은 프로필과 연속성에 근거한다. |
| Locations from architecture; profiles not redefined | ✅ | 본가와 세 anchors는 위치 프로필에 있다. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` read OK. |
| Location facets ⊆ Multi-facet anchors | ✅ | 세 facet 모두 프로필의 Multi-facet anchors에 있다. |
| Stagings from episode design; blocking not redefined | N/A | 고정 상황이 이어지는 장면이 없어 staging을 사용하지 않았다. |
| World rules / history consistent with bible | ✅ | 약과 혈맥계약의 양면성을 확정 해명으로 과장하지 않았다. |
| No improvised entities or silent lore | ✅ | 새로운 인물·세력·장소·규칙이 없다. |
| Continuity files used (ep 002+) | ✅ | 034 summary와 story-so-far를 Prior Design Alignment에 명시했다. |
| Character/location state vs story-so-far | ✅ | 도현의 병세 미확정 상태, 진우의 추적자 상태, 본가의 결투 후 상태를 유지한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-057·056·014·012를 Picks up/Advances/Plants/Holds로 분배했다. |
| No contradiction of released continuity | ✅ | 도현이 병을 치료받았다고 확정하지 않고, 떨림의 의미도 열어 둔다. |
| Conflicts section empty or escalated | ✅ | 충돌 없음으로 기록했다. |

## Design Consistency Gate
- Loaded required artifacts: ✅ — selective load and path checks complete.
- Locations: ✅ — index `북문서가-본가`; path `locations/북문서가-본가.md` read OK; facets `외당 무기고 입구`/`가주전 문앞`/`지하 보관실` ⊆ anchors.
- Length / Prose forecast: ✅ — Sc1 written=2,260; recomputed=2,260; Est=2,200 · Sc2 written=2,440; recomputed=2,440; Est=2,400 · Sc3 written=2,510; recomputed=2,510; Est=2,500 · scene fields/header=7,100.
- Episode List Summary: ✅ — 「도현의 병세를 확인하려 하지만 도현은 몸을 숨긴다」→ Scenes 1–2; 「가환이 남긴 약상자에서 ... 처방의 일부를 찾는다」→ Scene 3.
- Hook to Next / Closing: ✅ — Hook「처방의 나머지는 흑풍루의 거래 장부에 있다」; Out「처방의 나머지가 흑풍루 거래 장부에 있다는 새 추적 목표」; Scene 3 Turn「처방의 나머지가 장부에 있다는 결론」.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 도현이 왜 멈췄고 손을 왜 숨기는지 질문이 명시된다. |
| Personal stake present | ✅ | 진우가 회귀 전 자신이 죽인 아버지의 병세를 직접 확인하려 한다. |
| Episode Out hook | ✅ | 처방의 나머지와 거래 장부라는 단일하고 구체적인 후크다. |
| Exposition budget respected | ✅ | 계약 규칙과 약 그릇 양면성을 행동에 붙여 제한적으로 사용한다. |
| Seed discipline | ✅ | Plant 1개, Hint 1개, Hold 4개로 과잉 공개를 막았다. |
| Scene-first Key Events (all required fields) | ✅ | 모든 장면이 완전한 generation brief다. |
| Sensory-emotional on every scene | ✅ | 기름·약재·식은 냄새와 진우의 판단이 매 장면 연결된다. |
| Motifs planned across scenes | ✅ | 감춘 오른손과 반쪽 기록의 배치가 기록되어 있다. |
| Overview signature line | N/A | overview에 이번 회차에서 반드시 배치할 별도 signature line이 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 설명보다 문 잠금·흔적 수색·처방 발견이 중심이며 정보는 물건에 묶여 있다. |
| Sensory-emotional pairing | ✅ | 매 장면의 감각이 진우의 판단 변화로 이어진다. |
| Dialogue voices + Dialogue intent | ✅ | 진우의 짧은 압박, 도현의 낮은 회피, 혁의 또렷한 제동이 구분된다. |
| Reader-discovered meaning | ✅ | 도현의 칼이 살해만을 위한 것이 아니며 몸이 대가를 치른다는 결론을 독자가 조합하게 한다. |
| Antagonist plausibility | ✅ | 도현은 병세와 계약을 숨기는 행동을 하며 단순 악역으로 확정되지 않는다. |
| Closing image specified | ✅ | 처방 조각과 거래 장부 분류 부호의 겹침을 closing image로 고정했다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | 성인 남성향 회귀 무협 독자가 원하는 결투 직후의 즉시 후속 질문을 Scene 1 첫 행동으로 건다. |
| Personal stake matches what this reader came for | ✅ | 아버지의 떨리는 손을 확인하려는 진우의 복수·효 충돌이 사건 목표로 작동한다. |
| Pacing / density fits platform expectations | ✅ | 세 장면 모두 물리적 수색과 증거 전환이 있고, 7,100자 forecast가 플랫폼 범위 안이다. |
| Out hook makes this reader want the next episode | ✅ | 처방의 나머지와 흑풍루 거래 장부라는 구체적 추적 목표가 036의 행동을 약속한다. |
| No alienation of core audience without overview intent | ✅ | 로맨스·수련·장황한 가족 화해 없이 문서·약·칼 중심으로 진행한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 성인 남성향 회귀 무협·문파 장악물 독자의 연속 독서 욕구를 기준으로 판단한다.
- Strengths: 결투 직후의 손 떨림을 곧바로 추적 사건으로 바꾸고, 처방 조각을 거래 장부 후크로 연결해 다음 회차의 목표가 선명하다.
- Defects: —
- Reader impact: 진우가 아버지를 살리려는지 무너뜨리려는지 아직 결정하지 않은 상태가 유지되어 이탈 요인이 적다.

#### Genre Critic
- Stance: 회귀 무협의 선점 정보, 문파 내부 증거전, 부자 대립의 장르 약속을 점검한다.
- Strengths: 미래 지식이 병명 예언으로 쓰이지 않고 관찰·동선 역산의 우위로 쓰인다.
- Defects: —
- Reader impact: 칼의 액션이 문서·약재 추적의 사이다로 이어져 장르 리듬을 유지한다.

#### Plot Expert
- Stance: 원인-결과와 Hook body alignment, Out scope를 점검한다.
- Strengths: 034의 떨림 → 확인 시도 → 도현의 은폐 → 약상자 → 처방 조각 → 장부라는 인과가 끊기지 않는다. Summary·Out·Scene 3 Turn의 Hook 강도도 일치한다.
- Defects: —
- Reader impact: 다음 회차에서 왜 장부를 찾아야 하는지 독자가 질문 없이 이해한다.

#### Reader-Editor
- Stance: 회차 단위의 압축, 정보 노출, 마지막 전환을 점검한다.
- Strengths: 도현의 침묵을 설명으로 채우지 않고 문 잠금과 흔적으로 처리한다. Out 의무도 ‘거래 장부 추적’ 하나로 좁다.
- Defects: —
- Reader impact: 장면을 건너뛸 만한 설명 블록이 없고 마지막 물건 발견이 페이지 체류를 만든다.

#### Character Critic
- Stance: 진우·도현·혁의 행동 동기와 프로필-backed knowledge를 점검한다.
- Strengths: 진우는 상대의 손을 먼저 보는 습관대로 움직이고, 도현은 긴장할 때 오른손을 감싸는 프로필을 은폐 행동으로 확장하며, 혁은 명분과 검증을 중시하는 역할로 제동을 건다.
- Defects: —
- Reader impact: 세 인물의 기능이 겹치지 않아 짧은 대화만으로 관계 압력이 읽힌다.

#### Literary Critic
- Stance: 감각·모티프·Hold와 closing image의 실행 가능성을 점검한다.
- Strengths: 감춘 오른손과 반쪽 기록이 사건 장치이면서 관계의 불완전한 진실을 반영한다. 마지막 이미지는 설명 대신 물건의 겹침으로 끝난다.
- Defects: episode-level motif의 일부 배치는 Stage ⑥에서 문장 단위로 지켜야 한다.
- Reader impact: —

#### Setting/Lore Expert
- Stance: 본가 공간과 약 그릇·혈맥계약 규칙의 무리 없는 사용을 점검한다.
- Strengths: 모든 장소 facet이 catalog anchor이며, 치료 가능성은 조각의 판독 범위 안에서만 제시된다.
- Defects: —
- Reader impact: 폐쇄된 본가가 감시와 가족 비밀의 물리적 압박으로 기능한다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Episode-level motif placement should be carried into prose (Literary Critic) | Low | No | yes | 독자가 ‘감춘 손’과 ‘반쪽 기록’을 장면에서 체감해야 하지만 설계 구조를 다시 열 필요는 없다. | Stage ⑥ carry: 감춘 오른손은 Scene 1–2의 관찰/흔적으로, 반쪽 기록은 Scene 3의 물건·closing image로 구현 | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Approved:** 2025-02-14 — Schema·Consistency·Target Reader·required persona checks를 통과했다. 유일한 Carry-⑥ 모티프 지시는 생성 제약으로 기록되며 Pending design finding이 없다.
- **Next:** Stage ⑥ — manuscript generation
