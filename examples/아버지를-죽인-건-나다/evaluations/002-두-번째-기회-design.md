# Design Evaluation: Episode 002 — 두 번째 기회

> **Stage:** ⑤ Design Evaluation
> **Evaluated design:** `episodes/002-두-번째-기회.md`
> **Note:** 이 평가는 Stage ④ 설계의 재평가이며, Gate G4(사용자 설계 승인)를 대신하지 않는다.

## Criteria Check (from overview.md)

| Criterion | Result | Evidence |
|---|---:|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001의 범위이며 002 설계 평가의 실행 대상이 아니다. Episode 001 릴리스에서 이미 실행됨. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003의 회차 범위에 배치된 series-level 목표다. 002에서는 도현을 직접 등장시키지 않고 인장만 후크로 사용한다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 1–2에서 약 그릇 전달 경로와 운송 권한을 추적·대립하고, Scene 3–4에서 습격을 피하며 장부 위치를 선점한다. Out은 장부를 펼치기 전 서도현의 인장을 확인하는 다음 회차 후크를 남긴다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | 진우는 도현을 직접 선량하다고 확정하지 않고, 장부의 인장을 확인한 뒤에도 내용과 이유를 Hold한다. 도현은 현재 장면에 등장하지 않으며 인장이라는 양면적 증거만 제시된다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반·시리즈 회수 기준이다. 이 회차에서는 약 그릇과 가환만 인과를 진전시키고, 서찰·협박 조건은 아직 범위 밖이다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 전생의 동선과 습격 지점을 선점하지만 공격 시점이 달라져 팔을 다치고, 기억보다 현재 관찰·문서·수레 자국을 선택한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 시리즈 결말 기준이다. 002의 설계 범위를 넘어선다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 품질은 Stage ⑥–⑦에서 평가한다. 설계의 분량·예측 산술은 아래 Schema에서 검증한다. |

## Schema / Structural Integrity

| Check | Result | Evidence |
|---|---:|---|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `### Scene 1`–`### Scene 4` 아래에 표준 메타 필드와 표준 bullet 필드만 있다. |
| No skill/workflow dump after the design | ✅ | 설계 파일에 workflow 전문이나 `Pre-Design Load` 등 복제 섹션이 없다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–4 제목과 Beat·Outline·Unit이 서로 다르며 동일 장면 복제가 없다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 파일: `episodes/002-두-번째-기회.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | 메타와 bullet 필드가 콜론을 bold 내부에 둔 표기다. |
| Every scene has required meta + bullet fields | ✅ | 네 장면 모두 POV, Location, When, On stage, Staging 및 Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, Est.를 가진다. |
| Characters Appearing ↔ On stage union | ✅ | Appearing `{서진우, 가환, 장로-대표}` = 네 장면 On stage의 합집합. |
| On stage includes speakers | ✅ | Scene 1의 발화자는 서진우·가환, Scene 2의 발화자는 서진우·장로-대표이며 각 장면 On stage에 있다. Scene 3–4는 `none`이다. |
| Characters ⊆ `characters.md` | ✅ | 세 인물 모두 `characters.md`와 개별 프로필에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Arc·Seeds·Out에 등장하는 인물은 Appearing에 포함된다. 서도현은 장부 인장으로 언급되지만 이번 회차 장면의 발화·행동 인물로 제시되지 않는다. |
| No later-list cast debut | ✅ | 서진우·가환·장로-대표 모두 Episode 002 시점에 카탈로그상 사용 가능하다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | 북문서가 본가와 북항은 Key Locations이며, `가주전-회랑 접속부`, `장로회당 표결단`, `선착장`, `창고 골목`은 각 프로필의 정확한 Multi-facet anchors다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 장면은 단일 episode 파일 안에 있으며 중첩 장면 파일이 없다. |
| No template residue | ✅ | 원시 `{placeholder}`나 미완성 템플릿 잔여물이 없다. |
| Prose forecast present (outline + typed units) | ✅ | 다섯 허용 단위인 dialogue/action/sensory/POV/transition을 정수 `n×pick = subtotal`로 기록했다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | 독립 재계산: S1 `2×250+3×200+2×130+2×160+1×100=1,780`, Est 1,800; S2 `3×280+2×190+2×130+2×160+1×100=1,900`, Est 1,900; S3 `4×210+2×140+2×170+1×100=1,560`, Est 1,600; S4 `3×220+2×150+3×170+1×110=1,580`, Est 1,600. 각 Est는 outline line 수×200–350 범위 안이다. |
| Dialogue intent vs outline speech | ✅ | Scene 1–2는 대화 단위와 발화 의도가 있으며, Scene 3–4는 outline·Beat에도 말이 없고 Dialogue intent가 `none`이다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est fields: `1,800 + 1,900 + 1,600 + 1,600 = 6,900`; header addends: `1,800 + 1,900 + 1,600 + 1,600 = 6,900`. |
| Est. length sum ≥ Scale min | ✅ | 재합산 6,900자 ≥ 4,000자. |
| Est. length sum ≤ Scale max | ✅ | 재합산 6,900자 ≤ 8,000자. 중심 목표 5,000–7,200자 안이다. |
| Cited staging/profile paths exist | ✅ | 모든 Staging은 `none`이며 존재하지 않는 staging 파일을 인용하지 않는다. 인물·장소·세계 참조 경로가 디스크에 존재한다. |
| Episode List plot (not a different story) | ✅ | `series.md` Summary의 “과거의 동선을 바꾸고”, “내외당 인물을 역으로 추적”, “예정된 습격을 피함”, “장부 위치 선점”이 각각 Scene 1, 2, 3, 4의 Beat·Turn으로 실행된다. |
| Hook evidence strength (internal) | ✅ | `series.md` Hook의 “장부에는 흑풍루가 아니라 아버지의 인장이 찍혀 있다”와 Episode Out의 “표지 안쪽에 서도현의 인장이 찍혀 있음을 확인”, Seeds의 `서도현의 인장`, Scene 4 Turn의 같은 관찰이 모두 동일한 확인 강도다. 내용·이유는 모든 표면에서 Hold된다. |
| Hook scope (no Out creep) | ✅ | 마지막 Out은 장부 위치 선점과 서도현 인장 확인만 수행한다. 추격·새 세력 도착·두 번째 폭로·장부 내용 공개를 추가하지 않는다. |
| No design-paste / meta-only scenes | ✅ | 네 장면 모두 관찰·권한 대립·매복·장부 회수라는 구체적 사건이 있고, Beat·Outline·Unit budget이 중복되지 않는다. |

## Structure & Arc Checks

| Check | Result | Evidence |
|---|---:|---|
| Episode arc coherent | ✅ | 전달 경로 관찰 → 권한 기록 대립 → 변동한 습격 → 장부 선점·인장 확인으로 인과와 정보 상승이 이어진다. |
| Scene transitions chain | ✅ | S1 항만 번호 확인 → S2 권한 기록 → S3 북항 선착장 → S4 수레 자국이 끊기는 창고 골목. When은 회귀 직후에서 배신 전날 새벽·아침으로 순차적이다. |
| Scene sections complete | ✅ | Scene Index의 1–4가 모두 완전한 본문 장면으로 존재한다. |
| Generation Readiness | ✅ | Schema의 Length, forecast, cast, path, facet, Hook-body, scope 검사에 ❌가 없다. |
| Beat concreteness | ✅ | 모든 Beat가 손·소매·문서·빈칸·칼 궤적·수레 자국·봉인 끈 등 관찰 가능한 행동으로 작성됐다. |
| Est. length budget | ✅ | 정확한 재합산 6,900자이며 Scale과 중심 목표를 모두 만족한다. |
| Prose forecast quality | ✅ | Scene 1–2의 dialogue/단서 대립, Scene 3–4의 action·sensory·POV가 각 outline과 대응한다. |
| Episode List scope aligned | ✅ | Episode List Summary의 각 절을 실행하고 Out은 Hook 범위를 넘지 않는다. |
| Prior hook addressed (ep 002+) | ✅ | Episode 001의 약 그릇 확보와 명령 계통 추적을 Scene 1에서 직접 이어받는다. |

## Architecture & Continuity Compliance

| Check | Result | Evidence |
|---|---:|---|
| Pre-Design load reflected in Prior Design Alignment | ✅ | overview, series, Appearing profiles, used locations, world rules, continuity set을 모두 명시하고 각 제약의 반영을 기록했다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 중심 추적, P1 정보 선점, 복수 대상으로서의 도현 유지를 따른다. |
| Episode List Summary / Hook to Next honored | ✅ | Summary의 선점 사건과 Hook의 인장 확인을 각각 본문과 Out에서 같은 강도로 실행한다. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Episode Arc, Scene 4 Turn, Seeds, Out, Closing image가 모두 인장 확인 후 내용은 보류하는 구조다. |
| Characters from architecture; profiles not redefined | ✅ | 서진우·가환·장로-대표의 역할·습관·말투는 개별 프로필을 사용하며 새 핵심 사실을 재정의하지 않는다. |
| Profile-backed knowledge / recognition | ✅ | 진우의 문서·인장 판별은 `상대의 인장 확인` 습관 및 world의 문서·인장 추적 규칙에 근거한다. 도현의 의도·흑풍루 연결은 인장만으로 확정하지 않는다. |
| Locations from architecture; profiles not redefined | ✅ | 두 Key Location과 네 facet이 프로필의 anchors에 정확히 대응한다. |
| Stagings from episode design; blocking not redefined | ✅ | 모든 장면이 `Staging: none`이며 별도 staging을 허위로 로드하지 않는다. |
| World rules / history consistent with bible | ✅ | 미래 지식의 변동성, 문서·인장·운송망 추적, 약 그릇의 단서성, 회귀술의 미확정 대가를 지킨다. |
| No improvised entities or silent lore | ✅ | 새 인물·세력·장소·무공 규칙을 Key Events 안에서 조용히 발명하지 않는다. |
| Continuity files used (ep 002+) | ✅ | `story-so-far`, Episode 001 summary, `unresolved-threads`를 로드하고 TH-001을 Picks up, TH-003을 Advances, TH-002를 Holds로 처리한다. |
| Character/location state vs `story-so-far` | ✅ | 진우는 약 그릇을 확보한 회귀 직후 상태에서 출발하며 북문서가 본가의 새벽 상태를 유지한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-001 직접 진전, TH-003 변동성 확인, 장부·인장 관련 단서 Plant, TH-002 및 최종 명령자·장부 내용 Hold가 명확하다. |
| No contradiction of released continuity | ✅ | Episode 001의 릴리스 사실을 수정하거나 도현의 마지막 행동을 확정하지 않는다. |
| Conflicts section empty or escalated | ✅ | 명시적 `Conflicts / open questions`는 None이지만, 장부 내용·인장 이유·최종 명령자를 Hold로 남겨 후속 갈등을 보존한다. |

## Design Consistency Gate

| Check | Result | Evidence |
|---|---:|---|
| Canonical path and catalog references | ✅ | 실제 episode path와 모든 인물·장소·세계 경로가 존재한다. |
| Locations and citeable facets | ✅ | 네 장면의 facet이 각 장소 profile의 Multi-facet anchors에 있다. |
| Length / forecast | ✅ | Scene fields와 header가 모두 6,900이며 독립 Unit 재계산이 일치한다. |
| Episode List Summary / Hook | ✅ | Summary의 모든 핵심 절이 Beats에 있고, Hook body surfaces가 동일한 강도로 인장을 확인한다. |
| Continuity / architecture | ✅ | Ep 001 릴리스 상태·TH-001~003·프로필 기반 지식을 위반하지 않는다. |
| Generation readiness | ✅ | 구조·분량·캐스트·경로·facet·Hook-body에 차단성 ❌가 없다. |

## Engagement Checks (Design)

| Check | Result | Evidence |
|---|---:|---|
| Opening Question defined | ✅ | “전생에 자신을 감시하던 자는 약 그릇을 누구에게 넘겼고, 동선을 바꾸면 장부는 어디로 사라지는가?”가 명시되어 Scene 1에서 즉시 작동한다. |
| Personal stake present | ✅ | 실패 시 장부와 회귀 사실을 함께 잃고, 과도한 개입 시 표적이 드러난다. |
| Episode Out hook | ✅ | 장부를 펼치기 전 서도현의 인장을 확인해 다음 회차의 직접 대치로 연결한다. |
| Exposition budget respected | ✅ | 새 개념을 추가하지 않고 약 그릇 흔적·문서 권한·운송망 추적 세 항목만 사건 속에 배치한다. |
| Seed discipline | ✅ | 인장은 Hint로 제시되고 내용·이유는 Hold되며, Summary·Out의 확인 강도와 충돌하지 않는다. |
| Scene-first Key Events (all required fields) | ✅ | 각 장면에 Situation, Beat, Turn, Function, 감각·정서, Dialogue intent, Transition, Outline, Unit, Est가 있다. |
| Sensory-emotional on every scene | ✅ | 네 장면 모두 감각 단서와 진우의 정서·판단이 결합되어 있다. |
| Motifs planned across scenes | ✅ | 검집 두드리기는 Scene 1–3, 접힌 문서 모서리는 Scene 1–2·4의 Sensory-emotional/Beat에 구체적 실행 cue가 있다. |
| Overview signature line | ✅ | overview에 별도 고정 signature line이 없으므로 해당 없음으로 확인한다. |

## Literary Craft Checks (Design)

| Check | Result | Evidence |
|---|---:|---|
| Info : tension balance | ✅ | 정보는 문서·인장·운송망의 행동 단서로 전달되고, Scene 2 권한 대립과 Scene 3 매복이 즉시 긴장을 만든다. |
| Sensory-emotional pairing | ✅ | 약재 냄새·먹물·젖은 목재·피·봉인 끈이 진우의 계산·긴장·비용 인식과 함께 설계됐다. |
| Dialogue voices + Dialogue intent | ✅ | 가환은 느린 존대와 회피, 장로 대표는 격식·권한 용어, 진우는 짧은 질문과 공개 낭독 압박으로 구별된다. |
| Reader-discovered meaning | ✅ | “기억보다 관찰과 문서가 중요하다”는 의미를 선언하지 않고 동선·매복·단서 선택의 결과로 체험시킨다. |
| Antagonist plausibility | ✅ | 장로 대표의 운송표 회수·절차 위반 프레임·선열람 권한을 허가이자 함정으로 바꾸는 이해관계가 일관된다. |
| Closing image specified | ✅ | 어두운 창고에서 펼치지 않은 장부 안쪽의 서도현 인장만 등불을 받는 이미지가 구체적이다. |

## Literary Awards Juror Checks (Design)

{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)

> **Locked Target Reader:** 회귀·빙의·환생 무협과 문파 장악물, 가족 관계 반전, 복수형 사이다를 선호하는 성인 남성향 웹소설 독자.

| Check | Result | Evidence |
|---|---:|---|
| Opening earns the locked reader's attention | ✅ | Episode 001 직후 약 그릇의 전달자를 감시하는 즉시 추적과 전생 동선 뒤집기로 시작한다. |
| Personal stake matches what this reader came for | ✅ | 아버지를 향한 복수 전에 증거를 확보해야 한다는 주인공의 실질적 위험과 선점 쾌감이 있다. |
| Pacing / density fits platform expectations | ✅ | 4개 장면 모두 사건 중심이며 6,900자 forecast가 기본 Scale과 중심 목표를 만족한다. Scene 2 권한 대립, Scene 3 부상, Scene 4 장부 선점으로 상승한다. |
| Out hook makes this reader want the next episode | ✅ | 장부 내용이 아니라 아버지의 인장이 먼저 드러나며, 다음 회차의 첫 부자 대치 약속을 직접 당긴다. |
| No alienation of core audience without overview intent | ✅ | 도현의 내면을 공개하지 않고, 설명·수련·로맨스 없이 추적과 권한 대립을 유지한다. |

## Design Critique (required personas)

#### Target Reader
- **Stance:** 성인 남성향 회귀 무협 독자의 즉시 재미와 다음 회차 클릭 욕구를 기준으로 읽는다.
- **Strengths:** 전생 동선 역이용, 문서 권한을 뒤집는 장면, 예정된 습격의 선점, 아버지 인장 후크가 독자가 기대하는 정보 우위와 사이다를 순서대로 제공한다.
- **Defects:** —
- **Reader impact:** 진우가 기억만 믿지 않고 현재의 기록으로 승리하는 흐름이 주인공의 유능함과 위험을 동시에 보여줘 이탈 가능성을 낮춘다.

#### Genre Critic
- **Stance:** 회귀 무협·문파 장악물의 장르 약속과 복수형 사이다의 실질성을 검토한다.
- **Strengths:** 큰 선언 대신 동선·문서·권한을 선점하고, 승리마다 부상·절차 책임·불확실성을 비용으로 붙였다.
- **Defects:** —
- **Reader impact:** 만능 예언형 회귀가 아니라 변동한 현재를 다시 설계하는 장르 쾌감을 유지한다.

#### Plot Expert
- **Stance:** 인과, 정보 공개 순서, Episode List Summary/Hook의 body alignment와 Out scope를 검토한다.
- **Strengths:** Scene 1의 운송표 조각이 Scene 2 승인 기록, Scene 3 수레 자국, Scene 4 창고·장부로 이어진다. Summary·Out·Turn·Seeds가 모두 ‘장부 위치 선점 후 서도현 인장 확인’이라는 같은 강도이며, Out creep도 없다.
- **Defects:** —
- **Reader impact:** 독자는 단서를 따라가며 진우의 결론을 재구성할 수 있고, 다음 회차 질문도 “인장이 왜 찍혔는가”로 선명하다.

#### Reader-Editor
- **Stance:** 회차 단위 판매력, 정보와 긴장의 배치, closing beat의 밀도를 검토한다.
- **Strengths:** Scene 2의 공개 권한 대립이 중반 압력을 만들고, Scene 3의 매복과 Scene 4의 손바닥 상처가 선점을 공짜 승리로 만들지 않는다. 마지막 Transition은 단일 후크를 유지한다.
- **Defects:** —
- **Reader impact:** 6,900자 안에서 추적→대립→위험→선점의 상승이 뚜렷해 중간 이탈과 마지막 과밀을 모두 피한다.

#### Literary Critic
- **Stance:** 모티프, 감각, 의미의 독자 발견 가능성과 원고화 가능성을 검토한다.
- **Strengths:** 검집 두드리기와 접힌 문서 모서리가 감정 설명 대신 자기 통제·숨은 계층을 행동 cue로 만든다. 닫힌 장부와 떠오른 인장이 해설 없는 closing image를 제공한다.
- **Defects:** —
- **Reader impact:** 독자는 “기억보다 관찰”이라는 의미를 진우의 선택과 비용에서 발견하며, 후크가 주제와 사건 양쪽에서 작동한다.

#### Character Critic
- **Stance:** 인물의 동기, 행동 습관, 관계 압력, profile-backed knowledge를 검토한다.
- **Strengths:** 진우의 손·소매 우선 관찰과 검집 두드리기가 실제 추적·판단으로 쓰인다. 가환의 약재 얼룩·찻잔, 장로 대표의 금고리·권한 용어가 말과 행동을 분리한다. 진우는 도현의 의도를 인장만으로 확정하지 않는다.
- **Defects:** —
- **Reader impact:** 인물이 설정을 설명하는 대신 각자의 보호·권한·복수 이해관계로 움직여 다음 대치에 대한 신뢰를 만든다.

## Design Adjudication

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---:|---|---|---|---|---|---|---|
| 1 | Scene 3의 익명 매복 공격 주체를 원고에서 성급히 명명하지 말 것. 진우의 회귀 지식 한계와 불확실성을 유지해야 한다. (Plot Expert, Character Critic) | Low | No | yes | 독자는 이번 회차의 목표가 공격자 정체 확정이 아니라 변동한 시간표와 장부 위치 선점임을 유지할 때 다음 회차 후크에 집중할 수 있다. | Stage ⑥ 생성 제약으로 전달: 공격자의 정체·소속·대사는 설계에 없는 사실을 추가하지 않고, 칼 궤적·발소리·수레 자국만 관찰 가능한 범위에서 서술한다. | Carry-⑥ |
| 2 | Scene 4의 봉인 손상과 장부 회수 비용을 원고에서 감정적 독백으로 과설명하지 말 것. (Reader-Editor, Literary Critic) | Low | No | yes | 손바닥 상처·끊어진 끈·다시 닫는 행동이 이미 비용을 보여주므로 설명을 늘리면 냉정한 톤과 closing image가 약해진다. | Stage ⑥ 생성 제약으로 전달: 비용은 손·끈·호흡·행동으로 보여주고 도현 인장에 대한 해설성 결론은 Hold한다. | Carry-⑥ |

## Design Verdict

| Dimension | Result |
|---|---:|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate Note

- Schema / Structural Integrity: 모든 검사 ✅.
- Adjudication: `Carry-⑥` 2건만 있으며, ④ 설계 수정이 필요한 `Pending` 항목은 없다.
- 따라서 이 재평가 기준으로 Episode 002 설계는 Stage ⑥ 원고 생성에 구조적으로 준비되었다.
- 다만 Gate G4는 별도 사용자 설계 승인 게이트이므로, 사용자가 설계를 승인하기 전에는 원고 생성으로 이동하지 않는다.
