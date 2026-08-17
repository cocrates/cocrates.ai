# Design Evaluation: Episode 003 — 아버지의 칼끝

> Stage ⑤ 재평가 — 수정된 Stage ④ 설계를 기준으로 평가함. G4(사용자 설계 승인)는 아직 열려 있다.

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 실행 범위이며 Episode 003에서 재실행할 기준이 아니다. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | ✅ | Scene 2에서 두 사람이 처음 대면해 장부·인장·칼끝을 둘러싼 직접 대치를 실행한다. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Scene 2의 첫 부자 대립과 칼의 유보가 실질적 사건이며, 밤의 비밀 회합 초대가 Episode 004 후크다. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | 진우는 도현을 복수 대상으로 보지만, 도현은 죽일 수 있는 순간 칼을 거두고 회합을 명령한다. 행동의 해석을 열어 둔다. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | 후반 시리즈 페이오프 기준이다. 이 회차는 해당 인과를 회수하지 않고 관련 스레드를 간접 전진시킨다. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | 진우는 장부와 인장이라는 선점 정보를 미끼로 쓰되, 변동한 현재에서 도현의 반응을 직접 시험한다. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | 결말·후반 아크 기준이며 이번 회차는 초기 부자 충돌만 실행한다. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | 원고 단계 기준이다. 설계의 분량·forecast는 Schema에서 별도로 검증한다. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | 세 개의 flat `### Scene`이 표준 메타 필드와 bullet 필드를 사용한다. |
| No skill/workflow dump after the design | ✅ | 설계 뒤에 workflow 절차나 스킬 본문이 붙어 있지 않다. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Scene 1–3의 제목·기능·Beat가 서로 다르다. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | 실제 경로는 `episodes/003-아버지의-칼끝.md`다. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | 메타·bullet 필드 모두 표준 표기다. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | 세 장면 모두 POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, Est. length를 가진다. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing과 모든 scene On stage의 합집합이 서진우·서도현으로 일치한다. |
| On stage includes speakers | ✅ | 대화 의도가 있는 Scene 2의 두 발화 주체가 모두 On stage다. Scene 1·3은 `Dialogue intent: none`이다. |
| Characters ⊆ `characters.md` | ✅ | 두 인물 모두 index와 프로필에 존재한다. |
| Summary/Hooks cast alignment | ✅ | Summary, Out, Seeds, Closing과 각 장면이 두 인물만 사용한다. |
| No later-list cast debut | ✅ | 두 인물 모두 Episode 001부터 series presence가 있다. |
| Locations ⊆ `locations.md` Key Locations | ✅ | `북문서가-본가`가 Key Location이며 세 facet이 모두 profile의 exact Multi-facet anchors다. |
| Nested `episodes/{slug}/` scene files absent | ✅ | 장면은 하나의 episode 파일 안에 있다. |
| No template residue | ✅ | 미완성 placeholder나 template brace가 없다. |
| Prose forecast present (outline + typed units) | ✅ | 각 장면에 paragraph outline과 허용된 다섯 유형의 정수식이 있다. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Sc1 `3×180+2×130+3×150+1×80=1,580`, Sc2 `4×300+4×200+2×140+2×160+1×100=2,700`, Sc3 `2×180+2×120+3×150+1×80=1,380`으로 재계산된다. Est.는 각각 1,600/2,700/1,400이며 outline density band에도 들어간다. |
| Dialogue intent vs outline speech | ✅ | Scene 1·3은 speech 없이 관찰·판단으로 설계되고, Scene 2는 dialogue units와 의도가 일치한다. |
| Recorded Estimated Length = scene Est. sum | ✅ | `scene Est. fields: 1,600 + 2,700 + 1,400 = 5,700`; `header addends: 1,600 + 2,700 + 1,400 = 5,700`. |
| Est. length sum ≥ Scale min (hard) | ✅ | 재계산 합계 5,700은 최소 4,000 이상이다. |
| Est. length sum ≤ Scale max (hard) | ✅ | 재계산 합계 5,700은 최대 8,000 이하이며 중심 목표대에도 들어간다. |
| Cited staging/profile paths exist | ✅ | 모든 장면이 `Staging: none`이며 인용한 인물·장소 프로필은 디스크에 존재한다. |
| Episode List plot (not a different story) | ✅ | Series Summary의 장부 미끼·첫 대치·칼의 유보·비정상적 지식 감지가 각각 Scene 1–2의 Beat/Turn에 대응한다. |
| Hook evidence strength (internal) | ✅ | Series Hook, Summary, Arc close, Out, Seeds, Scene 2 Turn, Scene 3 Transition이 모두 ‘도현이 진우를 죽이지 않고 밤 회합에 부른다’는 동일한 강도로 유지된다. |
| Hook scope (no Out creep) | ✅ | 마지막 전진 의무는 밤 회합 참석뿐이며 추격·새 세력 도착·두 번째 폭로를 추가하지 않는다. |
| No design-paste / meta-only scenes | ✅ | Scene 1은 시험 준비, Scene 2는 대치 실행, Scene 3은 초대 후 전략 선택이라는 독립된 사건을 가진다. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | 장부 미끼 배치 → 칼끝 대치와 유보 → 회합 참석 선택의 상승·완화 구조가 명확하다. |
| Scene transitions chain | ✅ | 문이 열려 Scene 2 대치로 이어지고, 도현의 복귀와 초대가 Scene 3의 닫힌 문·선택으로 이어진다. 모든 When은 배신 전날 아침으로 고정된다. |
| Scene sections complete | ✅ | Scene Index의 세 행 모두 대응하는 완전한 Scene 섹션을 가진다. |
| Generation Readiness | ✅ | Schema의 길이·forecast·cast·path·facet·Hook-body가 모두 통과하고, 수정된 Scene 3의 보조 미끼가 출처와 기능까지 구체화되어 생성 중 추측을 요구하지 않는다. |
| Beat concreteness | ✅ | 봉인 확인, 손 위치 조정, 인장 노출, 칼의 발도·정지·거둠, 운송표 조각을 별도 미끼로 챙기는 행동이 관찰 가능하게 지정되어 있다. |
| Est. length budget | ✅ | 장면 합계와 header가 모두 5,700이며 Scale 4,000–8,000을 충족한다. |
| Prose forecast quality | ✅ | Dialogue, action, sensory, POV, transition의 typed units가 각 outline과 대응한다. |
| Episode List scope aligned | ✅ | Summary의 모든 절과 Hook을 실행하며 Out에서 새로운 사건을 덧붙이지 않는다. |
| Prior hook addressed (ep 002+) | ✅ | Episode 002에서 펼치지 않은 장부와 인장을 Scene 1에서 즉시 대치의 미끼로 전환한다. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | overview, series, appearing profiles, used location, world bible, Staging N/A, Episode 002 continuity가 기록되어 있다. |
| Series / overview tone & arc honored | ✅ | 냉정한 사건 중심 톤과 P1의 첫 부자 충돌을 유지한다. |
| Episode List Summary / Hook to Next honored | ✅ | Summary와 Hook이 장면의 구체적 사건·Turn으로 추적된다. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Arc close, Out, Plant, Scene 2 Turn이 동일한 초대 claim을 유지한다. |
| Characters from architecture; profiles not redefined | ✅ | 서진우의 인장 확인 습관과 서도현의 장부·칼 판단, 청강검·화상 단서를 기존 프로필 안에서 사용한다. |
| Profile-backed knowledge / recognition | ✅ | 진우의 인장 인식은 기존 반복 단서에, 도현의 진우 지식 감지는 부자 관계·현재 사건의 압력에 근거한다. 새로운 정체 폭로를 단정하지 않는다. |
| Locations from architecture; profiles not redefined | ✅ | 두 facet 모두 `locations/북문서가-본가.md`의 Multi-facet anchors에 정확히 있다. |
| Stagings from episode design; blocking not redefined | N/A | 모든 장면이 독립 상황이며 `Staging: none`이다. |
| World rules / history consistent with bible | ✅ | 문서·인장의 추적 가능성, 변동한 현재, 회귀 지식의 불완전성을 사용하며 새 초능력 규칙을 만들지 않는다. |
| No improvised entities or silent lore | ✅ | 북항 운송표 조각은 Episode 002 continuity에 이미 존재하는 소품이며, 새 인물·장소·세력·규칙을 추가하지 않는다. |
| Continuity files used (ep 002+) | ✅ | Episode 002 summary, story-so-far, unresolved threads를 Prior Hook와 thread table에 반영한다. |
| Character/location state vs `story-so-far` | ✅ | 진우의 팔·손바닥 부상과 장부 선점을 이어가며, 본가의 문서·운송망 압력을 유지한다. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-001~004가 Picks up/Advances/Plants/Holds에 명시적으로 배치된다. |
| No contradiction of released continuity | ✅ | Episode 001–002의 사건을 되돌리지 않고, Episode 002가 예고한 첫 대치를 실행한다. |
| Conflicts section empty or escalated (not ignored) | ✅ | 수정 후 남은 higher-layer conflict가 없다. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Loaded required artifacts | ✅ | 이 회차에 필요한 overview, series, world bible, index와 서진우·서도현·북문서가 본가 프로필, Episode 002 continuity를 대조했다. |
| Locations | ✅ | Sc1·3의 `가주전 문앞`, Sc2의 `가주전-회랑 접속부`가 모두 exact anchor다. |
| Length / Prose forecast | ✅ | Scene written/recomputed는 1,580/1,580, 2,700/2,700, 1,380/1,380이며 Est. header 합계는 5,700이다. |
| Episode List Summary | ✅ | 장부 미끼→Scene 1–2, 첫 대치→Scene 2, 칼을 거둠→Scene 2 Turn, 비정상적 지식 감지→Scene 2 Turn으로 대응한다. |
| Hook to Next / Closing | ✅ | Series Hook과 Episode Out이 같고 Scene 2 Turn이 실행하며 Scene 3이 회합 참석 선택으로 인계한다. |
| Hook scope | ✅ | 회합 초대 외에 추격·추가 폭로·새 세력 도착을 넣지 않는다. |
| Hook internal consistency | ✅ | Summary/Arc/Out/Seeds/Turn의 body surfaces가 동일한 강도다. |
| Overview signature lines | N/A | overview에 이 회차에 반드시 배치해야 할 별도 signature line이 없다. |
| Profile-backed knowledge | ✅ | 인장 인식과 부자 간 지식 시험에 관계·습관·drive 근거가 있다. |
| Opening honors prior | ✅ | Episode 002의 장부·인장 발견 직후를 시작점으로 삼는다. |
| Continuity states | ✅ | 부상, 장부 소유, 변동한 공격 시점, 도현의 미대치 상태를 유지한다. |
| Unresolved threads | ✅ | 네 스레드 모두 pick up/advance/plant/hold 처리되어 누락되지 않는다. |
| Characters | ✅ | Appearing과 On stage가 일치하고 의도된 발화 주체가 모두 등록되어 있다. |
| Locations | ✅ | 모든 장소가 카탈로그 key location과 exact facet을 사용한다. |
| When / timeline | ✅ | 모든 장면이 배신 전날 아침이며 인접 장면의 관계가 명시된다. |
| Staging | ✅ | 독립 장면의 `none` 처리로 staging 파일을 잘못 인용하지 않는다. |
| World | ✅ | 인장·장부 권위와 회귀 후 변동성만 사용한다. |
| Length / Prose forecast | ✅ | 정수식·재계산·Est.·outline density·회차 합계가 일치한다. |
| Series arc | ✅ | P1의 첫 부자 충돌과 Episode 004 회합 후크를 실행한다. |
| Tone / voice | ✅ | 낮고 건조한 긴장, 짧은 응답, 도현 내면 보류라는 제약을 지킨다. |
| No silent lore invention | ✅ | 북항 운송표 조각까지 모두 기존 continuity 또는 승인 아키텍처에 근거한다. |
| Cast / hook alignment | ✅ | Summary, Hooks, Seeds, Closing의 인물이 Appearing과 일치한다. |
| No template residue | ✅ | raw template 표식이 없다. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | 인장이 왜 찍혔고 들킨 순간 살아남을 수 있는가라는 질문이 명시되어 있다. |
| Personal stake present | ✅ | 장부를 잃으면 물증을 잃고, 인장을 내보이면 비정상적 지식을 들킬 위험이 있다. |
| Episode Out hook | ✅ | 죽일 수 있었던 도현이 밤 회합으로 부른다는 구체적 후크다. |
| Exposition budget respected | ✅ | 새 개념 설명 없이 장부·인장·회귀 지식이 행동과 압박으로만 진행된다. |
| Seed discipline | ✅ | 인장 사용 목적은 Hint, 밤 회합은 Plant이며 장부 내용과 최종 의미는 Hold다. |
| Scene-first Key Events (all required fields) | ✅ | 세 장면 모두 사건·전환·감각·outline·forecast를 갖춘다. |
| Sensory-emotional on every scene | ✅ | 먹물·기름 냄새, 금속음·화상 자국, 닫힌 문·장부 모서리가 진우의 감정·판단과 연결된다. |
| Motifs planned across scenes | ✅ | 칼끝·인장과 펼치지 않은 장부가 각 Scene의 Motif touch에 반복·변주된다. |
| Overview signature line | N/A | 별도 배치 대상이 없다. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | 장부 내용을 보류하고 칼끝의 즉시 위험으로 정보를 압축한다. |
| Sensory-emotional pairing | ✅ | 장부·인장·검·문이 목록이 아니라 진우의 관찰과 선택을 촉발한다. |
| Dialogue voices + Dialogue intent | ✅ | 도현은 낮은 추궁, 진우는 짧은 진단적 압박으로 대비되며 실제 대사는 원고 단계에 남긴다. |
| Reader-discovered meaning | ✅ | 도현이 죽이지 않은 선택의 의미를 독자가 추론하고 보호자라는 결론은 보류한다. |
| Antagonist plausibility | ✅ | 도현은 정보 또는 미확정 목적을 지키기 위해 칼을 뽑고도 통제된 선택을 한다. |
| Closing image specified | ✅ | 접힌 장부 모서리와 밤을 가리키는 손짓을 겹친다. |

## Literary Awards Juror Checks (Design)
Not required — overview.md has no prestige/awards criterion.

## Target Reader Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | 부상한 주인공, 물성 있는 장부, 임박한 부자 대치로 시작한다. |
| Personal stake matches what this reader came for | ✅ | 복수의 물증과 생존 위험을 동시에 건다. |
| Pacing / density fits platform expectations | ✅ | 세 장면·5,700자 forecast·무훈련/무설명 우회로가 연재 리듬에 맞는다. |
| Out hook makes this reader want the next episode | ✅ | 죽이지 않고 비밀 회합에 부른 이유를 확인하게 만든다. |
| No alienation of core audience without overview intent | ✅ | 전술·대립·가족 복수의 엔진을 유지하고 화해물로 전환하지 않는다. |

## Design Critique (required personas)

#### Target Reader
- Stance: 성인 남성향 회귀·무협 웹소설 독자의 연독성과 다음 클릭을 기준으로 본다.
- Strengths: 장부·인장·칼끝이 즉시 사건으로 연결되고, Scene 3의 보조 미끼도 `Episode 002에서 확보한 북항 운송표 조각`으로 구체화되었다.
- Defects: —
- Reader impact: 진우가 초대에 끌려가는 것이 아니라 원본과 보조 미끼를 분리해 다음 판을 설계하는 주인공으로 읽힌다.

#### Genre Critic
- Stance: 회귀 무협의 선점 정보, 즉시 대립, 사이다성 유보 후크를 검증한다.
- Strengths: 미래 지식이 만능 예언이 아니라 반응 시험으로 작동하며, 칼을 거두는 도현이 긴장을 끝내지 않고 다음 회합으로 연장한다.
- Defects: —
- Reader impact: 첫 부자 충돌의 보상과 미해결 미스터리를 동시에 얻는다.

#### Plot Expert
- Stance: 인과, Summary·Hook 충실성, body 표면의 후크 강도와 범위를 검증한다.
- Strengths: Episode 002의 장부·인장이 Scene 1 미끼, Scene 2 칼끝, Scene 3 회합 선택으로 연쇄된다. 수정된 운송표 조각은 기존 continuity에서 직접 이어져 인과가 닫혔다. Out creep도 없다.
- Defects: —
- Reader impact: “진우가 함정에 들어가는가, 역으로 미끼를 설계하는가”가 선명하다.

#### Reader-Editor
- Stance: 회차 판매성, 장면 밀도, closing transition의 과밀 여부를 본다.
- Strengths: 마지막 장면의 행동이 원본 장부 은닉과 기존 운송표 조각 준비로 분리되어 관찰 가능하다. 회합 초대라는 핵심 후크를 흐리지 않는다.
- Defects: —
- Reader impact: 모바일 연독에서 다음 회차 질문이 ‘밤 회합의 목적’으로 모인다.

#### Literary Critic
- Stance: 칼끝·인장·펼치지 않은 장부가 사건과 정서에 실제로 기여하는지 본다.
- Strengths: 권위가 폭력으로 바뀌었다가 유보되는 과정이 물성으로 반복되고, 닫힌 문과 접힌 장부가 설명 없는 여운을 만든다.
- Defects: —
- Reader impact: 도현의 보호 여부를 해설이 아닌 행동의 잔여로 판단하게 된다.

#### Character Critic
- Stance: 부자 양쪽의 profile-backed motivation, 행동 주체성, 지식 근거를 본다.
- Strengths: 진우의 인장 확인 습관과 도현의 장부·칼 중심 판단, 청강검·오른손 화상·멈추는 칼끝이 프로필과 맞물린다. 두 사람 모두 직접적인 진실 고백 없이 행동으로 압력을 만든다.
- Defects: —
- Reader impact: 진우는 수동적인 아들이 아니라 시험을 설계하는 후계자로, 도현은 단순 악역이 아니라 통제된 선택을 하는 상대자로 보인다.

## Design Adjudication

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| — | — | — | — | — | Stage ④ 수정으로 이전 Med finding이 해소되었고, 추가 적용 항목이 없다. | — | — |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate State
- Stage ⑤ 재평가: 통과.
- Stage ④의 이전 finding: `빈 표지와 다른 운송 단서`는 `Episode 002에서 확보한 북항 운송표 조각`으로 구체화되어 해소됨.
- G4 설계 승인: ⬜ 사용자 명시 승인 대기.
- 다음 가능 단계: 사용자가 설계를 승인하면 Stage ⑥ 원고 생성.
