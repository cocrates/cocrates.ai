# Design Evaluation: Episode 011 — 협박문의 조각

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| 1화 안에 죽음, 회귀, 약 그릇의 핵심 장치를 제시한다. | N/A | Episode 001 scope; this episode does not retroactively satisfy the pilot criterion. |
| 초반 3회 안에 서진우와 서도현의 첫 대치를 배치한다. | N/A | Episode 001–003 scope; already established before this episode. |
| 각 회차에 문파 장악, 거래, 추적, 대립 중 하나 이상의 실질적 사건과 다음 회차 후크를 둔다. | ✅ | Episode design Scene 1 executes 협박문 조각의 암호·운송 동선 추적, Scene 2 executes 문상철 생포, Scene 3 executes 출생 당일 화재라는 후크. |
| P1에서는 서도현을 복수 대상으로 유지하되, 단순 악역으로 납작하게 만들지 않는다. | ✅ | Design Arc states that 진우 does not confirm 문상철의 도현 배신 주장을 and does not confirm 도현’s innocence; Scene 3 Turn keeps the conflict unresolved. |
| 약 그릇·서찰·가환·협박 조건이 후반에 인과적으로 회수된다. | N/A | Late-arc payoff criterion; Episode 011 may advance TH-012~014 but cannot satisfy the eventual recovery. |
| 회귀 지식은 주인공의 설계와 사이다로 작동하되 만능 예언처럼 남용하지 않는다. | ✅ | Scene 1 uses remembered 거래 시각 to gain a lead, while the changed present requires 조각·운송표 대조; Scene 2 obtains only a partial win and loses the original. |
| 복수와 효의 충돌이 결말까지 유지되며, 흑풍루주에 대한 응징과 부자 대면이 결말의 정서적 절정을 이룬다. | N/A | Series/endgame criterion, outside this episode’s design scope. |
| 회차별 원고는 목표 분량과 사건 밀도를 지키고, 수련·설명·반복으로 분량을 부풀리지 않는다. | N/A | Manuscript quality belongs to Stages ⑥–⑦; forecast is recorded in Schema below. |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ | `episodes/011-협박문의-조각.md` uses one `## Scenes` section with unique flat `### Scene 1–3` schemas. |
| No skill/workflow dump after the design | ✅ | File ends at Generation Readiness; no workflow section or pasted procedure appears. |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ | Three distinct functions: array decoding, capture, interrogation; Beats and transitions are not duplicated. |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ | Actual disk path read: `episodes/011-협박문의-조각.md`. |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ | Scene meta fields and bullet fields consistently use the required notation. |
| Every scene has required meta + bullet fields (`none` OK) | ✅ | Scenes 1–3 each contain POV, Location, When, On stage, Staging, Situation, Beat, Turn, Function, Sensory-emotional, Dialogue intent, Transition out, Paragraph outline, Unit budget, and exactly one Est. |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ | Appearing = 서진우·문상철 plus 서도현 explicitly marked mention-only; Scene 1 has 서진우, Scenes 2–3 have 서진우·문상철. |
| On stage includes speakers | ✅ | Scene 1 dialogue is only 진우’s command; Scenes 2–3 speakers are 진우·문상철, all listed On stage. Ambient 인부 are explicitly non-speaking in `stagings/011-협박문-추적.md`. |
| Characters ⊆ `characters.md` | ✅ | `characters.md` catalog rows and this-turn profile reads confirm `characters/서진우.md`, `characters/문상철.md`, and `characters/서도현.md`. |
| Summary/Hooks cast alignment | ✅ | Summary, Out, and Seeds name only cataloged Appearing or mention-only 서도현; no ghost cast. |
| No later-list cast debut | ✅ | 서진우 is present throughout; 문상철 debuts at 004 and is listed through 012; 서도현 is series-present from 001. |
| Locations ⊆ `locations.md` Key Locations | ✅ | `locations.md` maps 북문서가-본가 and 북항 to Key Locations; no free-text primary location is used. |
| Location facets ⊆ Multi-facet anchors | ✅ | `locations/북문서가-본가.md` read OK with `가주전-회랑 접속부`; `locations/북항.md` read OK with `창고 골목` and `선착장`, all exact anchors. |
| Nested `episodes/{slug}/` scene files absent | ✅ | Episode design is a single canonical file; no nested scene-file layout is used. |
| No template residue | ✅ | No raw placeholder braces or template instructions remain. |
| Prose forecast present (outline + typed units) | ✅ | Every scene has a seven-line outline and only the five allowed typed units: dialogue, action, sensory, POV, transition. |
| Forecast ↔ Est. cross-check (independent) | ✅ | Recomputed: S1 `1×250 + 3×180 + 2×120 + 3×140 + 1×80 = 1,530`, Est 1,500; S2 `4×260 + 4×180 + 2×120 + 2×140 + 1×80 = 2,360`, Est 2,300; S3 `4×250 + 2×180 + 2×120 + 2×140 + 1×80 = 1,960`, Est 1,900. Each Est is within the outline-density band. |
| Dialogue intent vs outline speech | ✅ | Each scene with dialogue has matching Dialogue intent; no scene declares `none` while planning speech. |
| Recorded Estimated Length = scene Est. sum | ✅ | scene Est. fields: `1,500 + 2,300 + 1,900 = 5,700`; header addends: `1,500 + 2,300 + 1,900 = 5,700`. |
| Est. length sum ≥ Scale min (hard) | ✅ | Recomputed 5,700 ≥ overview Scale min 4,000. |
| Est. length sum ≤ Scale max (hard) | ✅ | Recomputed 5,700 ≤ overview Scale max 8,000; it remains in the stated central target band. |
| Cited staging/profile paths exist | ✅ | This-turn reads succeeded for `characters/서진우.md`, `characters/문상철.md`, `characters/서도현.md`, `locations/북문서가-본가.md`, `locations/북항.md`, and `stagings/011-협박문-추적.md`. Scene 1 correctly uses `Staging: none`; Scenes 2–3 cite the staging profile. |
| Episode List plot (not a different story) | ✅ | `series.md` Episode 011 Summary says 진우 선점·문상철 생포·도현 배신자 주장·출생일 북문서가 화재; these map respectively to Scene 1 Beat, Scene 2 Beat, and Scene 3 Beat. |
| Hook evidence strength (internal) | ✅ | Body quotes: Summary “문상철은 … 진우가 태어난 날 북문서가에서 불이 났다고 말한다”; Out “연락책은 … 불이 났다고 말한다”; Scene 3 Turn “출생 당일 화재라는 조사 가능한 사건”; Seeds classifies it as Plant and Scene 3 Dialogue intent repeats the concrete claim. Strength is consistent. |
| Hook scope (no Out creep) | ✅ | Out contains only the listed next obligation: investigate the birth-day fire; no chase, faction arrival, or second reveal is added. |
| No design-paste / meta-only scenes | ✅ | Each scene has a concrete causal event, distinct evidence exchange, and forward transition; no identical outline or budget loop. |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ | Abstract fragment → actionable route → capture → constrained testimony → concrete historical target. |
| Scene transitions chain | ✅ | Scene 1 exits before the harbor movement; Scene 2 moves the captured contact to the pier; Scene 3 ends with the birth-fire records as the next target. `When` remains the same morning with no timeline jump. |
| Scene sections complete | ✅ | All three Scene Index rows have complete corresponding scene sections and required fields. |
| Generation Readiness | ✅ | All Schema rows pass; no structural, length, cast, path, facet, or body-hook failure remains. |
| Beat concreteness | ✅ | Every Beat contains an observable deduction, capture, evidence confrontation, or testimony. |
| Est. length budget | ✅ | Independent total is 5,700; it matches the header and remains within 4,000–8,000. |
| Prose forecast quality | ✅ | Typed units correspond to each scene’s dialogue, action, sensory, POV, and transition demands. |
| Episode List scope aligned | ✅ | The design executes the full Episode 011 Summary and preserves the exact Hook to Next without scope creep. |
| Prior hook addressed (ep 002+) | ✅ | The Episode 010 hook—trace the fragment’s sender and connect 가환’s silence to an action/evidence choice—is addressed by preserving the fragment and choosing sender-tracking over immediate confrontation. |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ | Design records overview, series, appearing/used profiles, world aspects, staging, and both required continuity files in Load confirmation. |
| Series / overview tone & arc honored | ✅ | Event-centered pursuit, evidence pressure, and restrained family suspicion match overview’s 회귀 무협·가족 복수극 tone. |
| Episode List Summary / Hook to Next honored | ✅ | Summary and Hook are quoted and mapped in Schema; Scene 3 provides the exact birth-day fire claim. |
| Hook internal consistency (design surfaces) | ✅ | Summary, Arc close, Out, Plant, and Scene 3 Turn/Dialogue intent all preserve the same observable hook. |
| Characters from architecture; profiles not redefined | ✅ | Catalog and full profiles read; the episode cites existing states and does not redefine appearance, drive, or relationships. |
| Profile-backed knowledge / recognition | ✅ | 문상철’s knowledge is supported by `characters/문상철.md` relationships to 도현 and 흑풍루; no unsupported recognition claim is introduced. |
| Locations from architecture; profiles not redefined | ✅ | Both scene locations map to `locations.md` and their profile paths were read successfully. |
| Location profile paths readable | ✅ | Exact reads succeeded for `locations/북문서가-본가.md` and `locations/북항.md`. |
| Location facets ⊆ Multi-facet anchors | ✅ | Exact anchors verified after profile reads: `가주전-회랑 접속부`, `창고 골목`, `선착장`. |
| Stagings from episode design; blocking not redefined | ✅ | `stagings/011-협박문-추적.md` exists and was read; it fixes states, blocking, props, ambient, and continuity for Scenes 2–3. |
| World rules / history consistent with bible | ✅ | The design treats the fragment as a partial contract-system clue and explicitly withholds the full condition, consistent with `world/혈맥계약과-약그릇.md`. |
| No improvised entities or silent lore | ✅ | No new named person, faction, location, system, or rule is added; the birth-day fire is a Plant tied to the existing North Gate Seo-ga history mystery. |
| Continuity files used (ep 002+) | ✅ | `continuity/story-so-far.md` and `continuity/010-가환의-침묵-summary.md` were loaded and reflected in Prior Design Alignment and Threads. |
| Character/location state vs `story-so-far` | ✅ | 진우’s evidence-preserving choice, 문상철’s harbor role, and the unresolved states of 도현 and the harbor are carried forward without contradiction. |
| Unresolved threads: pick up / advance / plant / hold | ✅ | TH-012 and TH-013 advance; TH-014 remains held; the birth-day fire is planted; full conditions and sender remain held. |
| No contradiction of released continuity | ✅ | The design continues after Episode 010’s fragment discovery and does not alter any released event. |
| Conflicts section empty or escalated (not ignored) | ✅ | Design records no conflicts and explains why 문상철’s use remains within his approved profile. |

## Design Consistency Gate
| Check | Result | Evidence |
|-------|--------|----------|
| Required artifacts loaded | ✅ | Required design and reference set was read this turn, including every cited character/location/staging profile. |
| Canonical path | ✅ | `episodes/011-협박문의-조각.md`. |
| Locations index / paths / facets | ✅ | Index membership, exact profile reads, and exact Multi-facet anchors all pass separately. |
| Length / forecast | ✅ | Independent products and `scene Est. fields: 1,500 + 2,300 + 1,900 = 5,700` match header addends. |
| Episode List Summary and Hook | ✅ | Every Summary clause is executed, and Hook body surfaces remain at equal evidence strength. |
| Continuity / architecture / staging | ✅ | No catalog contradiction, released-continuity contradiction, or staging drift is found. |

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ | “협박문 조각은 누구에게 무엇을 요구했고, 왜 도현의 약상자 안에 숨겨져 있었는가?” |
| Personal stake present | ✅ | Revealing the fragment may expose 도현’s secret while giving the sender time to erase evidence; the “네 아들” wording also implicates 진우’s life. |
| Episode Out hook | ✅ | The contact’s concrete statement about the birth-day North Gate Seo-ga fire is a strong, single next-episode obligation. |
| Exposition budget respected | ✅ | Contract structure is limited to functional marks, delivery order, and partial wording; no full system lecture is planned. |
| Seed discipline | ✅ | The birth-day fire is clearly a Plant; original text and sender remain Hint/Hold rather than being falsely resolved. |
| Scene-first Key Events (all required fields) | ✅ | All scenes provide causal Key Events and complete generation fields. |
| Sensory-emotional on every scene | ✅ | Each scene pairs a concrete smell/sound/texture with 진우’s restrained reaction or decision. |
| Motifs planned across scenes | ✅ | Folded paper/empty box and wet ash/ember scent recur across Scenes 1–3. Scene-level Motif-touch is handed to Stage ⑥ as a generation constraint. |
| Overview signature line | N/A | Overview contains no episode-specific signature dialogue requirement; the design explicitly holds none. |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | ✅ | Information arrives through deduction, capture, and pressure rather than exposition; each reveal creates a narrower question. |
| Sensory-emotional pairing | ✅ | Ink/medicine smell, harbor noise, wet ash, and evidence texture are attached to 진우’s containment, focus, or redirection. |
| Dialogue voices + Dialogue intent | ✅ | 진우 uses short imperative/interrogative pressure; 문상철 uses low merchant honorifics and avoids direct liability, matching both profiles. |
| Reader-discovered meaning | ✅ | The design asks the reader to hold both possibilities: 도현 may be targeted and 문상철 may be deflecting blame; no closing monologue explains the theme. |
| Antagonist plausibility | ✅ | The unseen sender’s use of a fragmented message, delivery route, and emptied box creates operational plausibility without prematurely naming the mastermind. |
| Closing image specified | ✅ | Wet ash falls onto the fragment in 진우’s sleeve as he turns toward the birth records. |

## Literary Awards Juror Checks (Design)
{Not required — overview.md has no prestige/awards criterion.}

## Target Reader Checks (Design)
**Locked audience:** 성인 남성향 웹소설 독자 who prefer regression/wuxia, faction control, family reversals, and revenge-driven catharsis.

| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ | The fragment becomes an immediately actionable code and pursuit rather than a static mystery. |
| Personal stake matches what this reader came for | ✅ | The investigation threatens both 진우’s control of the revenge plan and his interpretation of his father. |
| Pacing / density fits platform expectations | ✅ | Three event-bearing scenes, 5,700-character forecast, no training or daily-life padding. |
| Out hook makes *this* reader want the next episode | ✅ | A precise birth-date/location fire clue promises records, hidden parentage, and a new evidence hunt. |
| No alienation of core audience without overview intent | ✅ | The episode remains action- and evidence-centered; ambiguity complicates the revenge target without replacing the revenge engine. |

## Design Critique (required personas)
#### Target Reader
- Stance: A 성인 남성향 회귀 무협 독자로서 선점의 쾌감과 아버지에 대한 의심이 다음 회차 욕구를 만드는지 본다.
- Strengths: Scene 1’s code-to-location deduction gives immediate agency; Scene 2 delivers a clean capture; Scene 3 converts the capture into a concrete historical target.
- Defects: —
- Reader impact: The reader receives a partial win without losing the larger mystery, so continuation pressure remains high.

#### Genre Critic
- Stance: 회귀 무협·문파 장악물의 정보 선점, 추적, 생포, 후크 계약을 점검한다.
- Strengths: Future knowledge produces a tactical advantage, not omniscience; the harbor capture preserves the genre’s practical “先手” pleasure.
- Defects: Low — The capture beat risks feeling familiar if execution becomes a generic overpowering. Proposed fix: Stage ⑥ should foreground the changed route, workers’ noise, and the red-jade-ring/rope evidence so the win is tactical rather than merely numerical.
- Reader impact: Keeping the advantage situational will make the protagonist’s competence feel earned instead of automatic.

#### Plot Expert
- Stance: Causality, thread movement, Episode List alignment, Hook body strength, and Hook scope are checked.
- Strengths: The causal chain is tight: fragment → transport code → preemption → capture → material interrogation → fire clue. Body surfaces preserve the exact Hook strength.
- Defects: —
- Reader impact: The reader can track why each scene follows the previous one and understands what question Episode 012 must answer.

#### Reader-Editor
- Stance: Serialization readability, exposition restraint, scene density, and closing sellability are checked.
- Strengths: The episode has one dominant investigation, distinct scene functions, and a single closing obligation; the final Transition is not crowded with multiple signals.
- Defects: Low — Scene 3 contains several disclosures (문상철’s blame, 약·운송 교체, 출생 암시, fire date/location). Proposed fix: Stage ⑥ should stage the disclosures as interrupted testimony and let the fire clue land after a pause, rather than delivering them as a continuous explanation.
- Reader impact: Controlled disclosure will prevent the final scene from reading like a summary while preserving the strong hook.

#### Literary Critic
- Stance: Motif continuity, sensory-emotional craft, Hold discipline, and closing image are checked.
- Strengths: Folded paper/empty box embodies partial truth and absent evidence; wet ash links the harbor present to the birth-day past; the closing image is visual and non-didactic.
- Defects: Low — Episode-wide motifs are specified but not represented as explicit Motif-touch fields inside each scene schema. Proposed fix: carry the motif placements into Stage ⑥ as execution constraints: folded crease in Scene 1, emptied box in Scene 2, wet ash/returning paper image in Scene 3.
- Reader impact: Recurrence will give the clue hunt an emotional texture beyond mechanical decoding.

#### Character Critic
- Stance: Agency, profile-backed knowledge, motivation, and voice distinction are checked because the episode is driven by 진우·문상철’s confrontation.
- Strengths: 진우’s choice to preserve evidence instead of confronting 도현 expresses his arc; 문상철’s self-protective merchant logic is supported by his profile and gives the interrogation a distinct pressure dynamic.
- Defects: Low — Scene 2’s “붉은 옥 반지와 봉인끈을 먼저 제압” is vivid but should remain an evidence-focused tactical choice, not an unexplained combat flourish. Proposed fix: Stage ⑥ should make clear how control of the hand/rope prevents the concealed 호신도 and preserves the proof.
- Reader impact: Clear motivation will make both the capture and the later testimony feel character-driven rather than plot-required.

## Design Adjudication
| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | Episode-wide motif recurrence is not explicit in scene fields (Literary Critic) | Low | No | yes | Motif continuity improves the evidence hunt’s serial texture, but the existing design already identifies exact scene placements and does not require a structural redesign. | Carry folded paper/empty box and wet ash/ember-scent placements into manuscript generation. | Carry-⑥ |
| 2 | Scene 2 capture could read as generic overpowering (Genre Critic, Character Critic) | Low | No | yes | The target reader wants tactical competence; execution should show why the changed route, hand control, rope, and harbor noise produce the capture. | Preserve as a Stage ⑥ execution constraint; do not alter the approved Beat. | Carry-⑥ |
| 3 | Scene 3 has multiple disclosures that could become exposition (Reader-Editor) | Low | No | yes | The birth-fire Hook must remain sharp; generation should space the testimony and avoid a pasted explanation block. | Use interrupted testimony, pressure questions, and a pause before the exact date/location clue. | Carry-⑥ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ |
| Target-reader readiness | ✅ |
| Design quality | ✅ |
| Generation-ready | ✅ |
