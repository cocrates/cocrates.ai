# Design Evaluation: Episode 117 — 복수의 칼끝

> Target: `episodes/117-복수의-칼끝.md`
> Target Reader: 회귀 무협·문파 장악·가족 복수 반전을 선호하는 성인 남성향 웹소설 독자
> G4: Architect-approved before evaluation

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|---|---|---|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 criterion; outside this late-arc scope. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Early-series criterion; outside this episode. |
| 각 회차에 실질적 사건과 다음 회차 후크를 둔다. | ✅ | 도현의 공개 책임 전환·어머니의 퇴로 차단·처분 권리 후크가 Scenes 1–4에 있다. |
| P1에서는 서도현을 복수 대상으로 유지하되 단순 악역으로 만들지 않는다. | N/A | P1 criterion. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Series payoff criterion; prior episodes establish it. |
| 회귀 지식은 만능 예언처럼 남용하지 않는다. | ✅ | No new future knowledge; Jinwoo acts from current evidence and released continuity. |
| 복수와 효의 충돌이 결말까지 유지된다. | ✅ | Jinwoo does not kill Dohyun or absolve him; public responsibility remains active. |
| 원고 분량·사건 밀도를 지킨다. | N/A | Manuscript Stage ⑥–⑦; forecast is checked below. |

## Schema / Structural Integrity
| Check | Result | Evidence |
|---|---|---|
| Canonical schema / path / fields | ✅ | `episodes/117-복수의-칼끝.md`; four unique scenes with all required fields and one Est. each. |
| Cast and speakers | ✅ | Six Appearing characters equal every scene’s On stage union; unnamed soldiers are nonverbal. All six profiles were read this turn. |
| Locations / facets | ✅ | `흑풍루-본거지` is a Key Location; `의식장`, `침투 회랑`, `의식장 입구 처형대`, `탈출 수로` are exact anchors in `locations/흑풍루-본거지.md`. |
| Staging / paths | ✅ | `stagings/117-복수의-칼끝-심판.md` and all cited profiles/location paths read OK; mother’s `감옥-상태` is catalogued. |
| Forecast ↔ Est. | ✅ | Sc1 `3×250+3×180+2×120+2×140+1×80=1980`, Est 1900; Sc2 `4×250+3×180+2×120+2×140+1×80=2130`, Est 2000; Sc3 and Sc4 same 2130/2000. All products exact; outline density passes. |
| Length sum | ✅ | Scene fields `1900+2000+2000+2000=7900`; header addends same; 4,000 ≤ 7,900 ≤ 8,000. |
| Episode List plot | ✅ | `series.md` 117 Summary clauses map to Scenes 1–4: do not kill/public responsibility, mother blocks escape, right to kill. |
| Hook body alignment / scope | ✅ | Series Hook「어머니는 진우에게 흑풍루주를 죽일 권리를 넘긴다」 matches Summary, Out, Seed, Scene 4 Turn/Dialogue/Transition; no name-reading or second reveal. |
| Continuity / world | ✅ | 116 dagger, bound Lord, poison, lost sword sense, mother seal state, army boundary, TH-172/167 persist; no new rule or later payoff. |
| No paste / template residue | ✅ | Each scene has distinct causal action and no unresolved braces. |

## Structure & Arc Checks
| Check | Result | Evidence |
|---|---|---|
| Episode arc coherent | ✅ | Stopped dagger → public record → blocked escape → transferred kill-right hook. |
| Scene transitions / completeness | ✅ | Every transition names the next physical situation; all four index rows have complete scenes. |
| Generation Readiness | ✅ | No Schema, path, cast, facet, forecast, or Hook failure. |
| Beat concreteness | ✅ | Recordboard, dagger, waterway, mother’s body, and army boundary are observable actions. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|---|---|---|
| Prior Design Alignment / selective load | ✅ | Indexes, six profiles, location, world, new staging, story-so-far, and immediate 116 summary are recorded. |
| Profile-backed knowledge | ✅ | Mother blocks a route and refuses to command; this follows her Drive/Relationships, not an unsupported recognition reveal. |
| No silent lore | ✅ | No new character, faction, room, weapon, or contract law. |
| Threads | ✅ | TH-172 and TH-167 advance; 118 judgment is explicitly Held. |

## Engagement & Literary Craft
| Check | Result | Evidence |
|---|---|---|
| Opening question / personal stake | ✅ | The stopped dagger asks kill vs living record; Jinwoo’s responsibility is explicit. |
| Out hook | ✅ | One dominant physical hook: mother transfers the right to kill the Lord. |
| Exposition / seed discipline | ✅ | Existing record continuity is dramatized; two seeds and explicit Holds prevent over-explanation. |
| Motifs / closing image | ✅ | Downward-stopped knife and blocked door/open record motifs culminate in the raised blade and blocked waterway. |
| Antagonist plausibility | ✅ | Lord tries to reuse the old command structure and escape rather than becoming a passive villain. |

## Design Critique (required personas)

#### Target Reader
- **Stance:** Action-first adult male serial reader.
- **Strengths:** Refusing to kill Dohyun is not softened into reconciliation; the mother’s physical intervention supplies a strong new turn.
- **Defects:** —
- **Reader impact:** Concrete revenge delay plus transferred authority should pull the reader into 118.

#### Genre Critic
- **Stance:** Regression martial-arts / revenge contract.
- **Strengths:** Pays off the dagger cliff with a tactical moral reversal, while preserving a kill-right rather than removing revenge satisfaction.
- **Defects:** —
- **Reader impact:** The genre promise is deferred, not denied.

#### Plot Expert
- **Stance:** Causality, Hook alignment, scope.
- **Strengths:** Every Summary clause has a named scene; the Out has one obligation and does not steal 118’s judgment.
- **Defects:** —
- **Reader impact:** Clear causal escalation from 116.

#### Reader-Editor
- **Stance:** Installment density and click-through.
- **Strengths:** Four distinct turns, visible record mechanics, and a clean final image.
- **Defects:** Low — forecast is near the Scale ceiling; Stage ⑥ must not pad testimony or repeat the same responsibility claim.
- **Reader impact:** Carry as generation constraint only.

#### Character Critic
- **Stance:** Family pressure and voice.
- **Strengths:** Mother refuses to issue the old killing command; Dohyun retains agency by refusing deletion; Jinwoo owns the choice.
- **Defects:** —
- **Reader impact:** Avoids instant reunion sentimentality while making the mother’s return active.

#### Literary Critic
- **Stance:** Motif and closing restraint.
- **Strengths:** Door/record and stopped-knife motifs are spatial, not explained; closing image is placeable.
- **Defects:** —
- **Reader impact:** Adds emotional residue without leaving action genre.

## Soft-craft / Design Adjudication
- **Crowded Out:** no hit — the final Transition contains the transferred right and one supporting image only.
- **Pending design findings:** none.
- **Carry-⑥:** Low forecast-ceiling risk; do not pad, repeat testimony, or state the Hold conclusions.

| # | Finding | Severity | Apply? | Rationale | Status |
|---|---|---|---|---|---|
| 1 | Forecast near Scale ceiling; avoid repeated responsibility exposition | Low | yes | Target reader benefits from kinetic public confrontation; no design change is needed. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|---|---|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |

## Gate G5
- **Status:** Approved by Architect
- **Approved:** 2025-02-14 — Schema, independent arithmetic, paths/facets, new staging, continuity, all required personas, and Hook scope pass. One Low forecast-ceiling risk is Carry-⑥; no Pending finding remains.
- **Next:** Stage ⑥ — manuscript generation.
