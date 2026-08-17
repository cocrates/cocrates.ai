# Design Evaluation: Episode 010 — 가환의 침묵

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001의 도입·장치 제시 기준이며 Episode 010 설계 평가 범위를 벗어난다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 003에서 이미 실행된 초반부 기준이다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Episode 010은 가환 미행과 도현의 약 교체 목격이라는 실질적 추적 사건을 실행하고, Scene 3에서 약상자 속 흑풍루 협박문 조각을 다음 회차 후크로 확보한다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | Scene 2에서 도현은 약을 직접 교체하지만 복용 대상과 목적은 드러나지 않는다. 진우의 의심은 유지되면서도 보호·살해 양면성이 열려 단순 악역 확정을 피한다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반 회수는 시리즈/후속 아크 범위다. Episode 010에서는 약 그릇·가환·협박문을 Plant/Hint로 전진시키며, 심은 것 자체를 회수로 판정하지 않는다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | N/A | 이번 회차는 회귀 지식으로 미래를 확정하는 설계보다 현장 관찰과 증거 판단을 다루며, 해당 기준의 직접 실행 범위가 아니다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 결말·후반 정서적 절정 기준이며 Episode 010 설계 범위를 벗어난다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 분량·사건 밀도는 Stage ⑥/⑦에서 판정한다. 설계의 forecast arithmetic은 아래 Schema에서 별도로 확인한다. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `episodes/010-가환의-침묵.md`는 Episode 메타 필드와 `### Scene 1`~`### Scene 3` 아래의 flat bullet 필드를 사용한다. 대체 장면 파일이나 중첩 scene 디렉터리는 없다. |
| No skill/workflow dump after the design | ✅ | 설계 파일 후반에 `04-episode-design.md` 등의 워크플로 전문이나 스킬 덤프가 없다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | `### Scene 1`, `### Scene 2`, `### Scene 3`이 각각 한 번씩 존재하며 Beat·Situation·Unit budget·outline이 서로 다르다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 평가 대상 경로는 `episodes/010-가환의-침묵.md`이며 회차 번호와 제목 slug가 일치한다. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | Scene 헤더의 `**POV:**`, `**Location:**`, `**When:**` 및 bullet의 `- **Situation:**`, `- **Beat:**` 등 canonical notation을 사용한다. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | 세 장면 모두 POV·Location·When·On stage·Staging과 Situation·Beat·Turn·Function·Sensory-emotional·Dialogue intent·Transition out·Paragraph outline·Unit budget·Est. length를 갖는다. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Characters Appearing은 서진우·가환·서도현이며 Scene 1–2 On stage에 세 명, Scene 3 On stage에 서진우만 있다. 모든 장면의 합집합이 일치한다. |
| On stage includes speakers | ✅ | Dialogue intent에서 언급된 도현·가환·진우는 해당 Scene 1–2의 On stage에 모두 있다. Scene 3은 `none`이며 대사가 없다. |
| Characters ⊆ `characters.md` | ✅ | `characters.md`의 Character Catalog와 `characters/서진우.md`, `characters/가환.md`, `characters/서도현.md` 프로필을 이번 턴에 읽었고 세 인물 모두 등록·프로필 존재가 확인된다. |
| Summary/Hooks cast alignment | ✅ | Summary·Arc·Hooks·Seeds에서 핵심 행동 주체로 언급되는 진우·가환·도현이 모두 Characters Appearing에 포함된다. |
| No later-list cast debut | ✅ | 세 인물은 모두 Episode 010 이전부터 Episode List/architecture에 존재한다. 가환은 009부터, 진우·도현은 001부터 등장한다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | 두 장소 표기는 `북문서가-본가`의 `가주전-회랑 접속부`·`가주전 문앞` facet으로 구성되며 `locations.md` Key Locations의 `북문서가 본가` row에 매핑된다. |
| Location facets ⊆ Multi-facet anchors | ✅ | `locations/북문서가-본가.md`를 이번 턴에 읽었고 Multi-facet anchors에 `가주전 문앞`과 `가주전-회랑 접속부`가 정확히 있다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | Episode 010은 단일 canonical 파일이며 장면별 nested 파일을 사용하지 않는다. |
| No template residue | ✅ | 원시 `{placeholder}`·`{TOC only}`·미완성 템플릿 표기가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 각 장면에 6~7개 Paragraph outline과 허용된 dialogue/action/sensory/POV/transition의 정수형 n×pick Unit budget이 있다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Scene 1: 3×250+3×180+2×120+2×140+1×80 = 1,890, Est 1,900. Scene 2: 4×250+3×180+2×120+2×140+1×80 = 2,140, Est 2,100. Scene 3: 3×250+3×180+2×120+2×140+1×80 = 1,890, Est 1,900. 각 Est는 outline density band(6줄×200–350=1,200–2,100; 7줄×200–350=1,400–2,450) 안이며, written product도 독립 산술과 일치한다. |
| Dialogue intent vs outline speech | ✅ | Scene 1–2는 Dialogue intent가 있으며 outline/Beat에 낮은 확인·회피·통제의 말이 포함된다. Scene 3은 outline과 Beat 모두 비언어 수색이고 Dialogue intent가 `none`이다. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: 1,900 + 2,100 + 1,900 = 5,900.\nheader addends:    1,900 + 2,100 + 1,900 = 5,900. |
| Est. length sum ≥ Scale min (hard) | ✅ | 독립 합계 5,900글자로 overview Scale min 4,000글자를 넘는다. |
| Est. length sum ≤ Scale max (hard) | ✅ | 독립 합계 5,900글자로 overview Scale max 8,000글자 이내이며 중심 목표 5,000–7,200에도 들어간다. |
| Cited staging/profile paths exist | ✅ | 이번 턴에 `characters/서진우.md`, `characters/가환.md`, `characters/서도현.md`, `locations/북문서가-본가.md`, `stagings/010-도현의-약상자.md`를 각각 `read_files`로 읽어 모두 read OK를 확인했다. |
| Episode List plot (not a different story) | ✅ | `series.md` Episode 010 Summary는 “가환을 미행해 도현이 약을 바꿔치기하는 장면을 본다. 그러나 바뀐 약이 누구를 살리기 위한 것인지는 확인하지 못한다”이다. Scene 1–2가 미행·약 대조·교체·목적 미확정을 그대로 실행한다. |
| Hook evidence strength (internal) | ✅ | `series.md` Hook은 “도현의 약상자 안에서 흑풍루의 협박문 조각이 나온다”이다. Episode Summary·Out·Scene 3 Turn·Seed touch 모두 약상자 이중 바닥에서 협박문 조각을 발견하되 약 목적은 확인하지 않는 동일한 관찰 강도를 유지한다. |
| Hook scope (no Out creep) | ✅ | 마지막 Transition은 협박문 조각과 다음 추적 방향만 남긴다. 추격·연락책 생포·암호 해독 같은 별도 의무를 Episode 010 안에서 선행하지 않으며 Out 의무도 1개다. |
| No design-paste / meta-only scenes | ✅ | 모든 장면에 관찰·약 교체·수색이라는 구체적 극적 사건이 있고, 장면별 Beat·outline·Unit budget이 반복 복사되지 않았다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 009의 열린 문에서 시작해 가환·도현의 약 교체를 목격하고, 진우가 약 대신 협박문 조각을 선택하는 사건 성과로 끝난다. |
| Scene transitions chain | ✅ | Scene 1의 두 봉지가 Scene 2의 교체로 이어지고, Scene 2의 닫힌 약상자·퇴장과 Scene 3의 단독 수색이 자연스럽게 연결된다. 모든 When은 같은 날 오전으로 명시된다. |
| Scene sections complete | ✅ | Scene Index의 3개 row가 모두 완전한 `### Scene` 섹션과 대응한다. |
| Generation Readiness | ✅ | Schema의 필수 구조·cast·path·facet·hook·forecast·length가 모두 통과하며, 아래 Adjudication에는 ④ 설계 수정이 필요한 Pending 항목이 없다. |
| Beat concreteness | ✅ | 숨기기, 대조하기, 봉지 풀기·교체하기, 이중 바닥을 들어 종이를 꺼내기 등 관찰 가능한 행동이 중심이다. |
| Est. length budget | ✅ | 독립 재계산 합계 5,900 = header addends 5,900 = scene Est fields 5,900이며 4,000–8,000 범위다. |
| Prose forecast quality | ✅ | Dialogue·action·sensory·POV·transition 단위가 장면의 대사·행동·감각·판단·장면 연결과 대응한다. |
| Episode List scope aligned | ✅ | Episode List Summary의 미행·약 교체·목적 미확정과 Hook의 협박문 조각 발견을 모두 실행하며, 마지막 Out에 범위 밖 사건을 추가하지 않는다. |
| Prior hook addressed | ✅ | Episode 009 Summary의 “가환이 도현의 방으로 들어가고 진우가 뒤따른다”가 Scene 1의 첫 상황과 When에 직접 이어진다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | `Prior Design Alignment`가 overview·series·profiles·location·world·staging·continuity load를 전제로 009 직후의 문·가환·약 그릇 조각·도현 방 상태를 명시한다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 중심 감시, 가족 복수와 보호의 충돌, P1의 도현 의심 유지가 Summary·Tone Notes·Scene 선택에 반영된다. |
| Episode List Summary / Hook to Next honored | ✅ | `series.md` Episode 010의 Summary와 Hook을 각각 Scene 1–2의 약 교체 및 Scene 3의 협박문 조각 발견으로 연결한다. |
| Hook internal consistency (design surfaces) | ✅ | Summary·Episode Arc closing·Out·Seeds·Scene 3 Turn이 모두 “협박문 조각 발견 + 약 목적 미확정”을 같은 강도로 유지한다. |
| Characters from architecture; profiles not redefined | ✅ | 세 인물은 `characters.md`와 각 프로필에 등록되어 있고, Episode 010은 기존 외형·행동 습관·관계 동기를 사용한다. 신규 states나 인물 재정의가 없다. |
| Profile-backed knowledge / recognition | ✅ | 진우의 약 그릇 판독 습관·검집 두드리기, 가환의 약재 얼룩·보호 동기, 도현의 화상·약상자 통제가 각 프로필의 Appearance/Behavior/Relationships에 근거한다. 약의 목적은 안다고 주장하지 않는다. |
| Locations from architecture; profiles not redefined | ✅ | `북문서가-본가`와 두 citeable facet은 `locations.md`와 `locations/북문서가-본가.md`에 등록되어 있으며, 약상자 상황은 위치 프로필을 수정하지 않고 staging에 둔다. |
| Location profile paths readable | ✅ | 정확한 경로 `locations/북문서가-본가.md`를 이번 턴에 읽어 read OK를 확인했다. |
| Location facets ⊆ Multi-facet anchors | ✅ | 두 장면 facet `가주전-회랑 접속부`, `가주전 문앞`이 profile line 12의 Multi-facet anchors에 정확히 존재한다. |
| Stagings from episode design; blocking not redefined | ✅ | `stagings/010-도현의-약상자.md`는 Episode 010 소유의 Stage ④ staging이며 Scene 1–2의 L/C/R, cast states, 약상자·봉지·조각을 바인딩한다. Episode design은 이를 cite하고 좌우를 새로 바꾸지 않는다. |
| World rules / history consistent with bible | ✅ | 약 그릇이 치료·살해 양면성을 가지며 초기에는 대상·명령 계통을 판독할 수 없다는 `world/혈맥계약과-약그릇.md`를 넘지 않는다. 혈맥계약 전체 구조를 새로 확정하지 않는다. |
| No improvised entities or silent lore | ✅ | 새 인물·장소·세력·무공 규칙이 없고, 흑풍루 협박문 조각은 기존 흑풍루 표식·문서 추적 체계에 속한다. |
| Continuity files used (ep 002+) | ✅ | `continuity/story-so-far.md`와 `continuity/009-불탄-기록-summary.md`를 이번 턴에 읽었으며, 009 직후·같은 날 오전·TH-012 상태를 설계에 반영했다. |
| Character/location state vs `story-so-far` | ✅ | 진우의 후계자·추적 상태, 가환의 약 그릇 조각과 침묵, 도현의 미확정 약상자 상태, 본가의 가주전 추적선을 유지한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-012를 약 대조·침묵으로 전진시키고, TH-001·004·002·006·011을 새 단서로 확장한다. 약 목적·가환 최종 충성·협박문 전체는 Hold한다. |
| No contradiction of released continuity | ✅ | Episode 009의 가환 진입·진우의 뒤따름·약 그릇 조각을 되돌리지 않고 즉시 이어간다. |
| Conflicts section empty or escalated (not ignored) | ✅ | `Conflicts / open questions`는 None으로 기록되지만 약의 목적과 가환의 동기를 의도적 미확정으로 명시하고, Continuity References/Threads에서 해당 질문을 관리한다. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Loaded required artifacts | ✅ | Episode design의 Load confirmation과 이번 턴의 overview·series·indexes·profiles·world aspects·staging·continuity read 결과가 일치한다. |
| Canonical path | ✅ | `episodes/010-가환의-침묵.md`가 실제 canonical path다. |
| Staging / location index / path / facets | ✅ | `stagings/010-도현의-약상자.md`가 read OK이며, `북문서가-본가`와 두 facet이 각각 index·profile anchor에 매핑된다. |
| Length / prose forecast | ✅ | Scene fields 1,900+2,100+1,900=5,900, header addends 1,900+2,100+1,900=5,900; Unit product는 모두 독립 recompute와 일치한다. |
| Episode List Summary / Hook | ✅ | Episode List의 미행·약 교체·목적 미확정·협박문 조각 발견이 본문 Beat/Turn/Out에 같은 강도로 반영된다. |
| Overall gate | ✅ | Schema / Structural Integrity에 ❌가 없고, ④ 설계 파일을 평가 단계에서 수정하지 않았다. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | “가환은 왜 약 그릇 조각을 들고 도현의 방에 들어갔으며, 도현은 누구를 위해 약을 바꾸는가?”가 명시되고 Scene 1–2에서 관찰 질문으로 작동한다. |
| Personal stake present | ✅ | 진우는 약이 전생의 아버지 죽음과 이어질 수 있지만 치료인지 독인지 몰라 칼을 뽑지 못한다. |
| Episode Out hook | ✅ | 약상자 이중 바닥의 흑풍루 협박문 조각이라는 구체적 단서가 Episode 011의 암호 해독·연락책 추적으로 연결된다. |
| Exposition budget respected | ✅ | 약 그릇의 양면성만 행동 속에서 한 번 상기하고 혈맥계약 전체나 협박 조건을 설명하지 않는다. |
| Seed discipline | ✅ | 바뀐 약의 목적과 협박문 조각은 Plant, 가환의 보호 약속·도현 인장과 약상자는 Hint로 구분되며 Hold 목록과 충돌하지 않는다. |
| Scene-first Key Events (all required fields) | ✅ | 모든 장면이 정보 요약이 아니라 POV가 현장에서 보고 선택하는 Key Events로 구성된다. |
| Sensory-emotional on every scene | ✅ | Scene 1의 식은 약 냄새·돌 냉기, Scene 2의 쓴 냄새·칼자루에서 손을 떼는 판단, Scene 3의 젖은 종이·먹물 냄새와 흔적 보존 선택이 있다. |
| Motifs planned across scenes | ✅ | `문틈`은 Scene 1–2, `약재 얼룩과 검은 먹물`은 Episode 표에서 Scene 1–3으로 계획되어 있다. Scene 3의 별도 Motif touch 필드가 없으므로 원고에서 약 냄새·먹물·흔적을 의도적으로 회수해야 한다(Carry-⑥). |
| Overview signature line | ✅ | 약 그릇·가족 복수·사건 중심이라는 overview 제약이 Dialogue intent, 관찰 행동, 개입 보류의 선택으로 배치되고 설명문으로 소비되지 않는다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 약 그릇 규칙은 1회 행동 상기로 제한되고, 대부분의 분량은 가려진 표기·손·잠금쇠·개입 보류의 긴장에 할당된다. |
| Sensory-emotional pairing | ✅ | 각 장면의 감각이 진우의 전생 기억, 독/치료 판별 불능, 흔적 보존 선택과 직접 결합한다. |
| Dialogue voices + Dialogue intent | ✅ | 도현은 낮은 통제, 가환은 주어를 생략하는 회피, 진우는 질문을 삼키는 관찰로 역할별 말의 목적을 구분한다. |
| Reader-discovered meaning | ✅ | 도현·가환의 보호 가능성은 손동작·침묵·약 교체로 독자가 추론하고, 진우가 선의를 확정하는 독백은 Hold한다. |
| Antagonist plausibility | ✅ | 도현은 적대적 설명 대신 약을 숨기는 선택을 하며 P1의 의심 대상성을 유지한다. 협박문 조각은 배후 압력의 물리적 흔적이 된다. |
| Closing image specified | ✅ | 젖은 종이의 흑풍루 표식과 손끝에 남는 약 냄새가 테마 설명 없이 회차를 닫는다. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | 회귀·복수 무협과 사건 중심 웹소설을 기대하는 성인 남성향 독자에게 009의 문을 연 즉시 약상자 감시로 진입해 지연 없이 사건을 제공한다. |
| Personal stake matches what this reader came for | ✅ | 아버지를 죽인 미래와 현재의 약 교체가 직접 연결되고, 진우가 칼을 뽑을지 말지 선택해야 한다. |
| Pacing / density fits platform expectations | ✅ | 3장면·5,900자 forecast로 수련·설명 없이 미행→교체→수색의 단일 사건선을 압축한다. |
| Out hook makes this reader want the next episode | ✅ | 약의 목적을 해결하지 않고도 흑풍루 협박문 조각이라는 새 조사 단서를 제시해 다음 회차의 암호·연락책 추적 욕구를 만든다. |
| No alienation of core audience without overview intent | ✅ | 로맨스·일상·장황한 치유 서사로 이탈하지 않고 복수, 문파 내부 증거, 흑풍루 단서에 집중한다. |

## Design Critique (required personas)
#### Target Reader
- Stance: 회귀·복수 무협과 문파 장악, 가족 반전을 기대하는 성인 남성향 웹소설 독자의 실제 연속 독서 관점에서 평가한다.
- Strengths: 009의 열린 문을 즉시 회수하고, 약 교체라는 눈앞의 사건과 칼을 뽑지 못하는 개인적 딜레마를 함께 준다. 협박문 조각은 해결을 지연시키면서도 다음 행동 목표를 분명히 한다.
- Defects: —
- Reader impact: 오전 순찰 발걸음으로 장면의 시간·환경이 정렬되어 독자는 시간 모순에 멈추지 않고 약 교체와 협박문 조각에 집중할 수 있다.

#### Genre Critic
- Stance: 회귀 무협·가족 복수극의 장르 약속과 회차 단위 사건 보상을 확인한다.
- Strengths: 주인공이 미래 지식으로 자동 해결하지 않고 현장 증거를 선점하며, 도현을 적으로 의심하면서도 보호 가능성을 열어 장르의 가족 반전 축을 유지한다. 약 교체와 협박문 조각이라는 사건 보상도 명확하다.
- Defects: —
- Reader impact: 독자는 “이번 회차에서 무엇을 알아냈는가”를 약상자·협박문 조각으로 이해할 수 있다.

#### Plot Expert
- Stance: 인과, prior hook 회수, Episode List Summary/Hook의 본문 정렬과 Out scope를 검증한다.
- Strengths: 009의 문 열림→가환의 약상자 대조→도현의 봉지 교체→진우의 협박문 조각 확보가 단일 인과선이다. Summary·Out·Scene 3 Turn·Seeds가 Hook 강도를 유지하며, 마지막 Transition이 추격이나 해독으로 범위를 넘지 않는다.
- Defects: —
- Reader impact: 독자는 약의 답을 얻지 못한 것을 실패가 아니라 진우가 협박문 조각을 선택한 사건 성과로 받아들일 수 있다.

#### Reader-Editor
- Stance: 연재 회차의 읽힘, 정보-긴장 균형, 마지막 비트의 밀도를 검토한다.
- Strengths: 문틈·손·봉지·잠금쇠처럼 시각화 가능한 단서가 정보 설명을 대신한다. 마지막 장면은 협박문 조각 하나에 집중되어 다음 회차 의무가 과밀하지 않다.
- Defects: Scene 2에 교체·가림·장포 수납·개입 보류가 연속되므로 원고에서 동일한 “가렸다/확인하지 못했다” 문장을 반복하면 밀도가 떨어질 위험이 있다 → severity Low → Stage ⑥에서 각 제한을 서로 다른 행동·감각으로 표현하고 확인 불능을 한 번 이상 재진술하지 않는다.
- Reader impact: 반복되면 독자가 “또 못 봤다”고 느낄 수 있으나, 실행 단계에서 행동을 변주하면 긴장으로 유지된다.

#### Literary Critic
- Stance: 모티프, 감각-정서 결합, Hold와 closing image의 독자 발견성을 평가한다.
- Strengths: 문틈은 접근 불가능한 진실의 거리로, 약재 얼룩과 검은 먹물은 보호·공모의 양면 흔적으로 기능한다. 결말을 협박문 표식과 약 냄새의 이미지로 닫아 주제 문장을 피한다.
- Defects: Episode-level 표의 “약재 얼룩과 검은 먹물”은 Scene 1–3으로 잡혀 있지만 Scene 3에 `Motif touch` 필드가 없다 → severity Low → generation constraint로 Scene 3의 젖은 종이·먹물·약 냄새를 이미지와 진우의 흔적 보존 선택에 결합한다.
- Reader impact: 원고에서 마지막 이미지가 모티프를 실제로 이어받으면 미스터리의 정서적 잔향이 강화된다.

#### Character Critic
- Stance: 진우·가환·도현의 profile-backed 행동, 관계 압력, 인물별 침묵의 차이를 검증한다.
- Strengths: 진우의 검집 두드리기와 손 먼저 보기, 가환의 약재 얼룩과 주어를 생략하는 회피, 도현의 화상 입은 검지와 낮은 통제가 각 프로필에 근거한다. 가환이 보호자인지 공범인지 열어 둔 상태도 프로필의 Central Conflict와 일치한다.
- Defects: —
- Reader impact: 세 인물은 같은 밀실에 있어도 행동 언어가 달라 독자가 누가 무엇을 숨기는지 추적할 수 있다.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Scene 2에서 동일한 확인 불능을 반복 표현할 위험 (Reader-Editor) | Low | No | yes | 연재 독자의 속도감을 지키려면 표기·복용자·냄새의 제한을 각각 다른 행동으로 보여주고 문장 반복을 피해야 한다. | Stage ⑥ generation constraint: 확인 불능을 재진술하지 말고 행동·감각·선택으로 분산한다. | Carry-⑥ |
| 2 | Scene 3의 episode-wide motif touch 누락 (Literary Critic) | Low | No | yes | 설계의 모티프 표가 Scene 3까지 지정하므로 원고에서 약재 얼룩·검은 먹물·약 냄새가 closing image에 실제로 이어져야 한다. | Stage ⑥ generation constraint: Scene 3에서 젖은 종이의 먹물·약 냄새를 진우의 흔적 보존 선택과 결합한다. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |
