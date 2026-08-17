# Design Evaluation: Episode 030 — 폐위된 가주

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Series-level / Episode 001 criterion. Episode 030 design scope가 아님. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Early-arc criterion. Episode 030 design scope가 아님. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 1–2에서 공개 폐위와 실권 이양이라는 문파 장악 사건을 실행하고, Scene 3에서 문서 뒷면 문장을 확인해 다음 국면의 후크를 둔다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | 도현은 폐위안을 수락하지만 혐의까지 인정하지 않는다. 진우는 그를 용서하거나 선인으로 확정하지 않고, 수락의 이유를 미해결로 둔다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Late-arc / series payoff criterion. Episode 030에서는 무명주 기록을 Hold하고 후반 회수로 남긴다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 029 기록의 범위를 넘지 않고, 첫 명령도 미래 예언이 아니라 현재 문서의 공동 보관과 열람 제한에 근거한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Series-end criterion. Episode 030은 부자 충돌의 압력을 전진시키지만 결말을 충족시키는 회차가 아니다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript Stage ⑥/⑦ criterion. Design forecast는 별도 Schema에서 검증했다. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `episodes/030-폐위된-가주.md`; 3개 Scene 모두 canonical meta lines와 flat bullet fields를 사용한다. |
| No skill/workflow dump after the design | ✅ | Episode-specific design·gate·readiness만 있으며 workflow 절차 복사가 없다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–3의 제목·Beat·Turn·Function이 서로 다르다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 경로 `episodes/030-폐위된-가주.md`. |
| Field notation `**Field:**` / `- **Field:**` | ✅ | 모든 Scene required field가 canonical notation이다. |
| Every scene has required meta + bullet fields | ✅ | POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, Est. length가 각 Scene에 1회씩 있다. |
| Characters Appearing ↔ On stage union | ✅ | 세 Scene의 On stage 합집합이 서진우·서도현·남궁혁·의원의 제자·장로 대표이며 Characters Appearing과 동일하다. |
| On stage includes speakers | ✅ | Dialogue intent에 명시된 모든 인물은 해당 Scene On stage에 있다. |
| Characters ⊆ `characters.md` | ✅ | 다섯 인물 모두 `characters.md` 및 개별 profile에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary·Hooks·Seeds·Closing에 새 미등록 인물을 넣지 않았다. |
| No later-list cast debut | ✅ | 모두 030 이전부터 등장한 인물이다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | 세 Scene 모두 `북문서가-본가` Key Location을 사용한다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `장로회당 표결단`, `가주전 문앞`, `가주전-회랑 접속부`가 `locations/북문서가-본가.md`의 Multi-facet anchors와 정확히 일치한다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 단일 episode 파일만 사용했다. |
| No template residue | ✅ | raw `{placeholder}`가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 세 Scene 모두 7개 outline line과 5종 typed unit formula를 가진다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 `3×250+3×180+2×120+2×140+1×80=1,890`; Sc2 동일 1,890; Sc3 `3×250+3×180+2×120+3×140+1×80=2,030`. Est는 각각 1,900/1,900/2,000으로 ±20% 안이며 outline density band에도 들어간다. |
| Dialogue intent vs outline speech | ✅ | 모든 Scene의 speech intent가 outline의 대화 단위와 일치한다. |
| Recorded Estimated Length = scene Est. sum | ✅ | Scene fields `1,900+1,900+2,000=5,800`; header addends `1,900+1,900+2,000=5,800`. |
| Est. length sum ≥ Scale min | ✅ | 5,800 ≥ 4,000. |
| Est. length sum ≤ Scale max | ✅ | 5,800 ≤ 8,000. |
| Cited staging/profile paths exist | ✅ | 이번 평가 턴의 로드 경로: `characters/서진우.md`, `characters/서도현.md`, `characters/남궁혁.md`, `characters/의원의-제자.md`, `characters/장로-대표.md`, `locations/북문서가-본가.md`, `stagings/030-폐위-수락.md`; 모두 read OK 또는 이번 턴에 생성 성공했다. |
| Episode List plot (not a different story) | ✅ | Series Summary 「도현이 스스로 권좌를 내놓자 진우는 문파의 실권을 물려받는다」를 Scene 1–2가 실행하고, 「폐위 문서의 뒷면」을 Scene 3이 실행한다. |
| Hook evidence strength (internal) | ✅ | Series Hook의 문장 발견이 Summary·Out·Scene 3 Turn·Seed에 같은 강도로 반복된다. 어느 표면도 원액 위치나 추격을 추가하지 않는다. |
| Hook scope (no Out creep) | ✅ | Out 의무는 문서 뒷면 문장 발견 1개이며, 독립 의무가 2개를 넘지 않는다. |
| No design-paste / meta-only scenes | ✅ | 세 Scene 모두 공개 답변, 인장 이양, 문장 발견과 첫 명령이라는 관찰 가능한 사건이 있다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 답변 요구 → 수락·인장 이양 → 문장 발견·첫 명령의 인과가 선명하다. |
| Scene transitions chain | ✅ | Scene 1의 인장 이양 요구가 Scene 2 Situation으로 이어지고, Scene 2의 접힌 문서와 첫 명령 요구가 Scene 3으로 이어진다. |
| Scene sections complete | ✅ | Scene Index 3행 모두 완전한 Scene section을 가진다. |
| Generation Readiness | ✅ | Schema 전 항목 통과, Pending adjudication 없음. |
| Beat concreteness | ✅ | 재낭독, 인장 검증, 수락, 인장 내려놓기, 문장 발견, 장부 공동 보관 명령이 구체적이다. |
| Est. length budget | ✅ | 독립 재산술 합계 5,800, Scale 범위와 central band를 만족한다. |
| Prose forecast quality | ✅ | dialogue/action/sensory/POV/transition 수가 각 outline과 맞는다. |
| Episode List scope aligned | ✅ | Summary와 Hook을 확대하거나 축소하지 않았다. |
| Prior hook addressed | ✅ | 029의 ‘도현의 답변 대기’를 Scene 1에서 즉시 해결한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Phase A, Phase B, staging, continuity가 모두 기록되어 있다. |
| Series / overview tone & arc honored | ✅ | 냉정한 공개 절차와 사건 중심의 실권 장악을 유지한다. |
| Episode List Summary / Hook to Next honored | ✅ | 설계 Gate Evidence와 Schema Evidence가 각각 Summary/Hook을 직접 인용한다. |
| Hook internal consistency (design surfaces) | ✅ | Summary·Arc close·Out·Seed·Scene 3 Turn이 같은 문장 발견을 가리킨다. |
| Characters from architecture; profiles not redefined | ✅ | 기존 drive·voice·state를 사용하고 새 인물 설정을 추가하지 않았다. |
| Profile-backed knowledge / recognition | ✅ | 도현 필체 가능성은 진우의 문서 관찰과 기존 부자 관계에서 제시되며, 필체를 확정하지 않는다. |
| Locations from architecture; profiles not redefined | ✅ | 북문서가 본가의 세 citeable facet만 사용한다. |
| Location profile paths readable | ✅ | `locations/북문서가-본가.md` read OK. |
| Location facets ⊆ Multi-facet anchors | ✅ | 세 facet이 profile의 anchor 목록에 있다. |
| Stagings from episode design; blocking not redefined | ✅ | `030-폐위-수락`이 ④에서 작성되었고 cast state와 blocking을 고정한다. |
| World rules / history consistent with bible | ✅ | 공개 재판·인장·장부 권위와 장로회 구조를 따른다. |
| No improvised entities or silent lore | ✅ | 새 세력·인물·규칙·장소를 발명하지 않았다. |
| Continuity files used | ✅ | 029 summary와 story-so-far를 사용했다. |
| Character/location state vs story-so-far | ✅ | 진우는 칼 없는 회귀-직후 상태, 도현은 병세-노출 상태, 표결단은 직전 상태에서 이어진다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-050은 해결, TH-051은 진전, 문장은 Plant, 원액·동기는 Hold로 분리했다. |
| No contradiction of released continuity | ✅ | 029의 폐위안 낭독과 도현의 침묵을 되돌리지 않는다. |
| Conflicts section empty or escalated | ✅ | Conflicts: None. |

## Design Consistency Gate
- Loaded required artifacts: ✅
- Locations index / path / facets: ✅ — Key Location membership, exact profile paths, exact Multi-facet anchors를 각각 확인했다.
- Length / forecast: ✅ — 각 Scene written/recomputed 일치, Est·outline density·header 합계 일치.
- Episode List Summary: ✅ — 수락·실권 이양·뒷면 문장 발견이 각각 named Scene에 배치되었다.
- Hook to Next / Closing: ✅ — Series Hook의 `네가 나를 미워해야 살아남는다` 문장 발견이 Summary·Out·Scene 3 Turn·Seed에 같은 강도로 존재한다.

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | “도현은 왜 자신을 변호할 수 있는데도 폐위안을 받아들이는가?” |
| Personal stake present | ✅ | 진우가 원하던 몰락과 그 몰락의 보호 가능성이 충돌한다. |
| Episode Out hook | ✅ | 독자가 문장의 의미를 즉시 해석할 수 없게 하면서도 다음 전쟁의 방향을 명확히 연다. |
| Exposition budget respected | ✅ | 새 정보는 수락·인장·문장으로 제한하고 무명주 설명은 보류한다. |
| Seed discipline | ✅ | Plant 2개와 Hint 1개, Hold 목록이 분리되어 있다. |
| Scene-first Key Events | ✅ | 모든 Scene이 사건·전환·감각·대화 의도를 가진다. |
| Sensory-emotional on every scene | ✅ | 종이, 인장, 등불의 물리 감각이 진우의 관찰과 연결된다. |
| Motifs planned across scenes | ✅ | 인장과 손, 뒤집힌 종이를 Scene 배치와 함께 잠갔다. |
| Overview signature line | N/A | overview.md에 이 회차에 강제되는 별도 signature dialogue line이 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 절차 설명은 인장·문서 검증 행동 안에 묶이고, 장로 대표의 압박과 도현의 침묵이 장면을 추진한다. |
| Sensory-emotional pairing | ✅ | 매 Scene의 감각 단서가 진우의 판단 변화와 연결된다. |
| Dialogue voices + Dialogue intent | ✅ | 장로 대표의 격식·책임 언어, 혁의 명분 중심 확인, 제자의 기술적 제한, 진우의 짧은 명령, 도현의 낮은 응답이 분리된다. |
| Reader-discovered meaning | ✅ | 도현의 보호 가능성을 설명하지 않고, 독자가 수락과 문장을 함께 해석하게 한다. |
| Antagonist plausibility | ✅ | 장로 대표는 단순 악의가 아니라 장로회 권한을 보전하려는 제도적 목표로 압박한다. |
| Closing image specified | ✅ | 등불 아래 펼쳐진 문서 뒷면과 문장을 지정했다. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)
Target Reader: 회귀·빙의·환생 무협, 문파 장악물, 가족 반전과 복수형 사이다를 선호하는 성인 남성향 웹소설 독자.

| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | 직전 회차의 공개 폐위안에 즉시 답하고, 도현의 수락 여부를 첫 사건으로 해결한다. |
| Personal stake matches what this reader came for | ✅ | 아버지의 몰락과 권좌 탈취가 공개 절차와 부자 긴장으로 동시에 작동한다. |
| Pacing / density fits platform expectations | ✅ | 3 Scene, 5,800자 forecast, 각 Scene마다 명확한 Turn이 있어 설명 회차로 늘어지지 않는다. |
| Out hook makes this reader want the next episode | ✅ | ‘미워해야 살아남는다’는 보호인지 협박인지 판단할 수 없는 가족 후크가 다음 회차의 감금·거래 국면을 직접 연다. |
| No alienation of core audience without overview intent | ✅ | 도현을 초반부터 선인으로 확정하지 않고, 문파 장악 사이다를 유지한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 성인 남성향 회귀 무협 독자의 즉시 만족과 다음 회차 클릭 욕구를 기준으로 읽는다.
- Strengths: 029의 미결정 답변을 첫 Scene에서 바로 해결하고, 인장 이양이라는 시각적 권력 변화를 제공한다.
- Defects: —
- Reader impact: 도현이 몰락했지만 진우가 완전히 이긴 것은 아니라는 불안이 문장 후크로 남아 다음 회차를 견인한다.

#### Genre Critic
- Stance: 회귀 무협·문파 장악·가족 복수의 장르 약속을 점검한다.
- Strengths: 공개 회의에서 실권을 빼앗고 첫 명령을 내리는 장르적 보상이 분명하다. 아버지를 즉시 선인으로 세탁하지 않는다.
- Defects: —
- Reader impact: 사이다와 가족 반전의 긴장을 동시에 유지한다.

#### Plot Expert
- Stance: 029의 정지, 030의 수락, 031의 거래 후크 사이 인과를 점검한다.
- Strengths: 답변 요구→수락→인장 이양→뒷면 문장→첫 명령의 순서가 자연스럽다. Hook body alignment와 Hook scope도 일치한다.
- Defects: —
- Reader impact: 독자가 “왜 수락했나”를 품은 채 사건은 앞으로 진행되므로 지연이 정체로 느껴지지 않는다.

#### Reader-Editor
- Stance: 웹소설 회차의 첫 페이지 유지와 마지막 Out의 밀도를 점검한다.
- Strengths: 절차 설명을 세 번의 물리적 문서 행동에 묶었고, 마지막 Transition에 문장 발견과 첫 명령만 둬 과밀하지 않다.
- Defects: —
- Reader impact: 장로회 회의가 정적 장면으로 늘어지지 않고, 인장과 문서가 계속 이동한다.

#### Literary Critic
- Stance: 반복 이미지, 감각, 설명되지 않은 의미의 잔향을 점검한다.
- Strengths: 손·인장·뒤집힌 종이가 권력의 공개성과 사적 의도를 함께 운반한다. 결말을 주제 해설이 아닌 펼쳐진 문서 이미지로 닫는다.
- Defects: —
- Reader impact: 독자는 도현의 수락을 보호·거래·패배 중 어느 하나로 고정하지 않고 다음 회차로 넘어간다.

#### Character Critic
- Stance: 부자 관계와 조연들의 행동 동기, profile-backed knowledge를 점검한다.
- Strengths: 진우는 도현의 죄를 대신 증언하지 않으면서 권력을 받는다. 도현은 내면 공개 없이 손과 인장으로 행동한다. 혁과 제자는 각자의 검증 역할을 수행한다.
- Defects: —
- Reader impact: 진우의 승리가 복수의 완결이 아니라 책임의 시작으로 바뀌며 부자 갈등을 보존한다.

#### Setting/Lore Expert
- Stance: 장로회 절차, 인장·장부 권위, 장소 facet과 staging을 점검한다.
- Strengths: 세 장면 모두 `북문서가-본가`의 실제 Multi-facet anchors를 사용하고, 030 staging이 cast state와 blocking을 고정한다. 무명주를 원액 위치로 과잉 해석하지 않는다.
- Defects: —
- Reader impact: 독자가 권력 이양을 추상적 선언이 아니라 문서·인장·자리 이동으로 이해한다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| — | — | — | — | no | Schema·Hook·인물·장소·독자 유지력에서 적용할 High/Med 결함이 없다. | — | Skip |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Decision:** 필수 비평가 7인과 Schema/Continuity/Target Reader 검사를 완료했다. 적용할 Pending finding이 없으므로 원고 생성으로 진행한다.
- **Next:** Stage ⑥ — manuscript generation
