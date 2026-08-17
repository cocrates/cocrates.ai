# Design Evaluation: Episode 001 — 죽기 전날

> 평가 대상: `episodes/001-죽기-전날.md`
> 평가 범위: Stage ⑤ 설계 재평가
> 재평가 사유: Scene 2 산술 불일치 및 북문서가 본가 내부 facet 근거 보강
> 상태: 재평가 완료 — Gate G4(에피소드 설계 승인)는 별도 사용자 승인이 필요함

## Criteria Check (from overview.md)

| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | ✅ | Scene 1의 살인·죽음, Scene 2의 배신 전날 회귀, Scene 3의 약 그릇 이상과 명령 흔적이 실행된다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | ✅ | Episode 001 Scene 1에서 두 사람이 전생 마지막 밤에 대치한다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 3에서 약 그릇을 명령 계통 추적 단서로 판독하고 전달자를 첫 표적으로 정한다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | 도현은 복수 대상이지만 칼을 피하지 않고 뒤의 살기를 막는 행동으로 해석의 균열을 남긴다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반 회수 기준은 이 회차의 범위를 벗어난다. 약 그릇은 Plant로 제시되며 후반 회수를 완료했다고 보지 않는다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 신체·날짜·등불을 검증하고 변화한 현재를 전제로 첫 조사를 선택한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 101~120회 후반 아크 기준이며 001회 범위 밖이다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | ✅ | Scene Est.는 2,200 + 1,460 + 1,100 = 4,760자로 Scale 4,000~8,000 안이며, 각 Unit subtotal과 일치한다. |

## Schema / Structural Integrity

| Check | Result | Evidence |
|---|---|---|
| Canonical scene schema only | ✅ | 세 장면 모두 표준 메타 라인과 flat bullet 필드를 사용한다. |
| No skill/workflow dump after the design | ✅ | 설계 뒤에 workflow 전문이 없다. |
| Unique Scene headings; no pasted twin scenes | ✅ | Scene 1~3의 사건·전환·outline·budget이 구별된다. |
| Canonical episode path | ✅ | 실제 경로는 `episodes/001-죽기-전날.md`다. |
| Field notation | ✅ | `**Field:**` 및 `- **Field:**` 표기를 준수한다. |
| Every scene has required fields | ✅ | POV, Location, When, On stage, Staging, Situation~Est. length가 모두 있다. |
| Characters Appearing ↔ On stage union | ✅ | Appearing은 서진우·서도현이고, On stage union과 일치한다. |
| On stage includes speakers | ✅ | Scene 1의 두 발화 주체가 모두 On stage이며 Scene 2~3은 대화 intent가 none이다. |
| Characters ⊆ `characters.md` | ✅ | 두 인물 모두 index와 profile에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hook·Seeds·Closing의 인물은 Appearing 및 catalog에 있다. |
| No later-list cast debut | ✅ | 후발 Episode List 인물을 새로 등장시키지 않았다. |
| Locations ⊆ `locations.md` Key Locations / facets | ✅ | 세 장면 모두 `북문서가 본가`의 가주전 문앞·회랑·약방 facet을 사용하며, profile에 회랑 접속부와 약방 내부가 명시됐다. |
| Nested scene files absent | ✅ | 단일 episode 파일 구조다. |
| No template residue | ✅ | 미완성 placeholder가 없다. |
| Prose forecast present | ✅ | 각 장면에 5개 Unit 유형 기반 정수 formula와 paragraph outline이 있다. |
| Forecast ↔ Est. cross-check | ✅ | Scene 1: 2,200, Scene 2: `3×200 + 2×130 + 3×170 + 1×90 = 1,460`, Scene 3: 1,100. 각 Est.가 subtotal과 일치한다. Outline density도 각각 1,400~2,450 / 1,200~2,100 / 1,000~1,750 안이다. |
| Dialogue intent vs outline speech | ✅ | Scene 1만 대화 intent가 있으며 Scene 2~3에는 직접 발화가 없다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: `2,200 + 1,460 + 1,100 = 4,760`; header addends: `2,200 + 1,460 + 1,100 = 4,760`. |
| Est. sum ≥ Scale min | ✅ | `4,760 ≥ 4,000`. |
| Est. sum ≤ Scale max | ✅ | `4,760 ≤ 8,000`. 중앙 목표보다 낮지만 hard Scale 안이며 padding은 요구하지 않는다. |
| Cited staging/profile paths exist | N/A | 모든 장면이 `Staging: none`이다. |
| Episode List plot | ✅ | Summary의 죽음·회귀·약 그릇 이상·첫 표적이 Scene 1~3의 Beat/Turn에 대응한다. |
| Hook evidence strength internal | ✅ | Series Hook의 “독이 아니라 명령을 증명하는 흔적”이 Summary·Out·Seeds·Scene 3 Turn에 같은 강도로 유지된다. |
| Hook scope | ✅ | 마지막 Out은 명령 흔적과 첫 표적 설정만 남기며 추격·새 세력·추가 폭로를 만들지 않는다. |
| No design-paste / meta-only scenes | ✅ | 세 장면 모두 관찰 가능한 극적 사건과 선택을 가진다. |

## Structure & Arc Checks

| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | 살인의 결과 → 회귀 검증 → 약 그릇 판독과 추적 목표 설정으로 인과가 상승한다. |
| Scene transitions chain | ✅ | Scene 1의 의식 단절이 Scene 2로, Scene 2의 약방 등불이 Scene 3으로 이어진다. When도 전생과 회귀 후를 구분한다. |
| Scene sections complete | ✅ | Scene Index 세 행 모두 완전한 section으로 존재한다. |
| Generation Readiness | ✅ | 필수 schema, cast, location facet, forecast, hook, path 조건을 통과한다. |
| Beat concreteness | ✅ | 칼·상처·화상·등불·문양·잔류물·은패를 통한 행동이 각 장면을 지탱한다. |
| Est. length budget | ✅ | 재계산 합계 4,760자가 header와 scene fields에 동일하게 기록됐다. |
| Prose forecast quality | ✅ | Unit 유형이 각 장면의 dialogue/action/sensory/POV/transition 의도와 대응한다. |
| Episode List scope aligned | ✅ | 후속 회차의 약 그릇·명령 계통 조사를 위한 범위 안에 머문다. |
| Prior hook addressed | N/A | Episode 001이다. |

## Architecture & Continuity Compliance

| Check | Result | Evidence |
|---|---|---|
| Pre-Design load reflected | ✅ | overview, series, appearing cast, used location, world bible/aspects를 Prior Design Alignment에 반영했고 Ep 001 Continuity는 unchecked/N/A다. |
| Series / overview tone & arc | ✅ | 냉정한 사건 중심 회귀 무협과 P1의 정보 선점 구조를 따른다. |
| Episode Summary / Hook honored | ✅ | 두 상위 문장의 핵심 주장이 장면 사건으로 구현된다. |
| Hook internal consistency | ✅ | Summary·Arc·Out·Seeds·closing Turn이 명령 흔적이라는 단일 주장을 공유한다. |
| Characters from architecture | ✅ | 진우·도현의 행동과 말투가 profile과 일치한다. |
| Profile-backed knowledge | ✅ | 약 그릇 판독은 world aspect와 진우의 경험에 한정되며 도현의 의도는 확정하지 않는다. |
| Locations from architecture | ✅ | `locations/북문서가-본가.md`에 가주전 문앞·회랑 접속부·약방 내부 facet이 존재한다. 이번 변경은 기존 장소의 additive facet 확장이다. |
| Stagings | N/A | 모든 장면이 `Staging: none`이다. |
| World rules / history | ✅ | 약 그릇의 독 억제·명령 계통 판독, 회귀 대가, 봉인 규칙과 충돌하지 않는다. |
| No improvised entities or silent lore | ✅ | 새 인물·세력·장소를 추가하지 않았고 장소 보강은 profile에 기록했다. |
| Continuity files | N/A | Episode 001이다. |
| Unresolved threads | ✅ | Picks up은 `—`, Plants와 Holds가 명시되어 있다. |
| Released continuity contradiction | N/A | 이전 출시 회차가 없다. |
| Conflicts section | ✅ | 충돌 없음과 후크의 비확정성을 기록했다. |

## Design Consistency Gate

| Check | Result | Evidence |
|---|---|---|
| Loaded required artifacts | ✅ | Appearing/used profiles only, Ep 001 Continuity N/A. |
| Canonical path | ✅ | `episodes/001-죽기-전날.md` |
| Catalog / staging slug language | ✅ | 한국어 canonical slug, `Staging: none`. |
| Episode List Summary | ✅ | 세 핵심 절이 Scene 1~3에 concrete Beat로 대응한다. |
| Hook to Next / Closing | ✅ | 명령 흔적의 강도가 Summary·Out·Turn·Seeds에서 유지된다. |
| Hook scope / internal consistency | ✅ | 추가 추격·세력 도착·제2 폭로 없이 하나의 후크로 닫힌다. |
| Overview signature lines | ✅ | 별도 고정 대사가 없다. |
| Profile-backed knowledge | ✅ | catalog/world 근거가 있다. |
| Characters | ✅ | 모든 발화 주체가 On stage이며 catalog cast다. |
| Locations | ✅ | 모든 scene Location이 Key Location의 승인 facet이다. |
| When / timeline | ✅ | 전생 마지막 밤과 회귀 직후 배신 전날을 명시한다. |
| Length / forecast | ✅ | Scene 2 수정 후 모든 formula·Est.·header 합계가 일치한다. |
| Series arc / tone | ✅ | P1의 정보 우위 전환과 사건 중심 톤을 수행한다. |
| No silent lore invention | ✅ | 위치 facet을 profile에 additive하게 기록했다. |
| No cross-scene paste / template residue | ✅ | 중복 및 placeholder가 없다. |

## Engagement Checks (Design)

| Check | Result | Evidence |
|---|---|---|
| Opening Question defined | ✅ | 칼끝의 선택 가능성을 질문으로 세우고 Scene 1에서 완결하지 않는다. |
| Personal stake present | ✅ | 아버지를 죽인 죄와 같은 죽음의 반복 위험이 중심 stake다. |
| Episode Out hook | ✅ | 독이 아닌 명령 흔적과 미공개 첫 표적이 002회를 요구한다. |
| Exposition budget | ✅ | 회귀 대가와 약 그릇 판독 원리만 행동 속에 둔다. |
| Seed discipline | ✅ | Plant 1개, Hint 1개와 후반 Hold가 구분된다. |
| Scene-first Key Events | ✅ | 모든 장면에 필수 execution field가 있다. |
| Sensory-emotional pairing | ✅ | 피 냄새·돌 냉기·먹물·약 냄새·문양이 POV 반응과 짝지어진다. |
| Motifs | ✅ | 칼끝과 약 그릇 문양의 배치가 명확하다. |
| Overview signature line | N/A | overview에 별도 signature line이 없다. |

## Literary Craft Checks (Design)

| Check | Result | Evidence |
|---|---|---|
| Info : tension balance | ✅ | 즉시 대치와 검증이 먼저 작동하고 설명은 제한적이다. |
| Sensory-emotional craft | ✅ | 감각이 진우의 흔들림·계산·선택을 촉발한다. |
| Dialogue voices | ✅ | 진우의 짧은 압박과 도현의 낮고 느린 회피를 분리했다. |
| Reader-discovered meaning | ✅ | 도현의 보호 의도는 Hold하고 독자가 행동으로 의심하게 한다. |
| Antagonist plausibility | ✅ | 도현을 초반부터 선량하다고 확정하지 않는다. |
| Closing image | ✅ | 등불을 받은 두 겹 문양으로 닫는다. |

## Literary Awards Juror Checks (Design)

Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)

**Locked Target Reader:** 회귀·빙의·환생 무협, 문파 장악물, 가족 반전, 복수형 사이다를 선호하는 성인 남성향 웹소설 독자.

| Check | Result | Evidence |
|---|---|---|
| Opening attention | ✅ | 아버지를 찌른 직후의 죽음 직전으로 시작한다. |
| Personal stake | ✅ | 살해 죄·배신 의심·회귀 기회가 즉시 결합한다. |
| Pacing / density | ✅ | 총 4,760자는 hard Scale 안이며 장면별 사건 밀도와 forecast가 일치한다. 중앙 목표보다 낮지만 padding 없이 설계된 범위다. |
| Out hook | ✅ | 약 그릇의 명령 흔적과 첫 표적 미공개가 다음 화 클릭 이유를 만든다. |
| No unintended alienation | ✅ | 수련·로맨스·장황한 설명 없이 복수와 설계 약속에 집중한다. |

## Design Critique (required personas)

#### Target Reader
- Stance: 성인 남성향 회귀 무협 독자의 첫 회차 체류와 다음 화 클릭을 기준으로 본다.
- Strengths: 죽음·회귀·가족 반전·조사 후크가 지체 없이 연결된다.
- Defects: —
- Reader impact: 공간 facet이 명확해져 약방 발견의 물리적 장면도 따라갈 수 있다.

#### Genre Critic
- Stance: 회귀 무협과 복수형 웹소설의 즉시 사건성과 정보 우위 쾌감을 검토한다.
- Strengths: 미래 지식을 검증 가능한 행동 계획으로 바꾼다.
- Defects: —
- Reader impact: 장르 독자가 기대하는 “이번에는 먼저 친다”는 약속을 훼손하지 않는다.

#### Plot Expert
- Stance: 인과, 전환, Summary/Hook alignment와 Out scope를 검토한다.
- Strengths: 죽음의 불완전한 기억이 회귀 검증과 약 그릇 조사로 자연스럽게 상승한다.
- Defects: —
- Reader impact: Scene 2의 분량 기준이 1,460자로 고정되어 원고 리듬의 기준이 하나다.

#### Reader-Editor
- Stance: 연재 단위의 판매성, 정보·긴장 배분, closing density를 본다.
- Strengths: 세 장면이 감정 부채·검증·실질 단서로 분담되며 마지막 후크도 과밀하지 않다.
- Defects: —
- Reader impact: 약방이 본가 내부의 승인 facet으로 잠겨 장소 점프의 혼란이 해소된다.

#### Literary Critic
- Stance: 모티프, 감각·정서 결합, 독자 발견성을 검토한다.
- Strengths: 칼끝과 두 겹 문양이 선택·조작·명령의 흔적을 물성으로 운반한다.
- Defects: —
- Reader impact: 주제 해설 없이 죄책감과 의심을 따라갈 수 있다.

#### Character Critic
- Stance: 진우·도현의 동기와 profile-backed knowledge를 점검한다.
- Strengths: 진우는 도현을 즉시 공격하지 않고 증거 순서를 재설정한다. 도현의 보호 가능성은 행동으로만 남는다.
- Defects: —
- Reader impact: 도현의 의도를 확정하지 않아 후반 반전의 발견성이 보존된다.

#### Setting/Lore Expert
- Stance: 북문서가 본가의 공간 잠금과 약 그릇·회귀 규칙의 일관성을 검토한다.
- Strengths: 회랑과 약방이 profile의 내부 facet으로 추가되어 가주전에서 약방까지의 동선이 잠겼다.
- Defects: —
- Reader impact: 약 그릇이 편의적으로 출현하지 않고 본가의 조제·보관 공간에서 발견된다.

## Design Adjudication

| # | Finding | Severity | Conflict? | Apply? | Rationale | Action Taken | Status |
|---|---|---|---|---|---|---|---|
| 1 | Plot Expert — Scene 2 Unit budget과 Est. 불일치 | High | No | yes | 생성 기준과 장면 분량을 하나로 고정해야 한다. | Scene 2를 1,460자로 수정하고 전체 합계를 4,760자로 재계산했다. | Applied |
| 2 | Setting/Lore Expert — 회랑·약방 facet 근거 부족 | High | No | yes | 약 그릇 발견 장면의 물리적 개연성을 잠가야 한다. | `locations/북문서가-본가.md`에 가주전-회랑 접속부와 약방 내부 facet을 additive하게 추가하고 episode references를 갱신했다. | Applied |
| 3 | Reader-Editor — 공간 표기 보강 | Med | No | yes | Finding 2와 함께 독자-facing 공간 혼란을 해소한다. | 위치 facet 보강과 episode Locations Used 문구 수정으로 처리했다. | Applied |
| 4 | Target Reader — 중심 목표 분량보다 낮음 | Low | No | no | 4,760자는 hard Scale 안이며 padding은 사건 밀도를 해칠 수 있다. | 추가 분량은 적용하지 않았다. | Skip |
| 5 | Character Critic — 도현 보호 의도의 직접 해설 위험 | Low | No | no | 설계상 이미 POV 관찰과 Hold로 제한되어 있어 추가 설계 수정이 필요 없다. | Stage ⑥에서 도현 내면을 직접 해설하지 않는 생성 제약으로 유지한다. | Skip |

**Adjudication Status:** Applied findings were re-verified after Stage ④ changes. No pending design revision remains. Gate G4 still requires explicit user approval.

## Design Verdict

| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |
