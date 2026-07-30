# Stage ⑥ — Evaluation (`evaluations/{nnn}-{chapter-slug}.md`)

**Prerequisites:** Depends on mode — see below

**Gate artifact:** `{project-root}/evaluations/{nnn}-{chapter-slug}.md`

**Next stage:** `07-revision.md` if issues found; `05-generation.md` (after design eval) or `08-release.md` (after manuscript eval)

**Engagement / fidelity detail:** `04-chapter-design.md`, `05-generation.md`. **Layer sync:** [`consistency.md`](consistency.md).

---

## Two Evaluation Modes

| Mode | When | Inputs | Purpose |
|------|------|--------|---------|
| **Design Evaluation** | After stage ④ approval (recommended before ⑤) | `chapters/{nnn}-{chapter-slug}.md` (chapter + episode + scene Key Events) | Catch structural, pacing, seed, engagement, and **manuscript-readiness** issues **before** writing prose |
| **Manuscript Evaluation** | After stage ⑤ approval | Above + `manuscripts/{nnn}-{chapter-slug}.md` | Validate prose against design and criteria |

**Design evaluation is strongly recommended.** Fixing structure at design stage is more efficient than revising finished prose.

Use a **single evaluation file** per chapter — append or update sections as evaluation progresses.

---

## Procedure

### Design Evaluation

Propose after chapter design approval:

> *"Chapter {nnn} design is approved. Shall we evaluate the design before writing the manuscript? We'll validate structure, seeds, and engagement against criteria, then run the required critic personas."*

Load:
- `chapters/{nnn}-{chapter-slug}.md` (full file — chapter + episode sections + Prior Design Alignment)
- **Same prior-design set as stage ④ load:** `overview.md`, `series.md`, Chapter List source, `characters*`, `locations*`, `world-bible` (+ aspects), continuity (ch 002+)
- Re-run Prior-Design Consistency Gate from `04-chapter-design.md` against the finished design
- Lock **Target Reader** from `overview.md` (audience / platform / success-in-reader-terms) before persona critique
- **Literary Awards Juror:** also Chapter List source for work-level context — **Short:** `series.md`; **Series:** `parts/{nnn}-{part-slug}.md`; fill **Literary Awards Juror Checks (Design)** before persona critique

**Required personas (design):** run every persona in the [Default persona sets](#default-persona-sets) for this chapter — do not skip for speed. Optional personas only when overview criteria or the user request them.

### Manuscript Evaluation

Propose after chapter manuscript approval:

> *"Chapter {nnn} manuscript is ready. Shall we evaluate it? We'll validate against the chapter design and prior architecture/continuity, then run the required critic personas."*

Load:
- `manuscripts/{nnn}-{chapter-slug}.md`
- `chapters/{nnn}-{chapter-slug}.md` (full design)
- Architecture + continuity set listed in the chapter design (same as stage ⑤ Pre-Generation Load)
- Re-run **Design-Fidelity Gate** from `05-generation.md`
- Same **Target Reader** lock as design eval

**Required personas (manuscript):** same [Default persona sets](#default-persona-sets). Additional named personas on request.

**Literary Awards Juror:** When this persona is in the set (or requested), fill **Literary Awards Juror Checks** in the evaluation file (design and/or manuscript section) before the persona critique. Load Chapter List source (**Short:** `series.md`; **Series:** parent `parts/{nnn}-{part-slug}.md`) and `overview.md` for work-level context.

---

## Evaluation Structure

```markdown
# Evaluation: Chapter {nnn} — {Title}

## 1. Design Evaluation
{Omit section if design eval not yet performed}

### Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|

### Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Chapter arc coherent | ✅ / ⚠️ / ❌ | |
| Episode hooks chain | | |
| Episode sections complete | | {Every Episode Index row has full Key Events?} |
| Manuscript Readiness | | {Every scene has required fields — see `04-chapter-design.md`} |
| Scene transitions coherent | | {Transition out → next Situation} |
| Beat concreteness | | {No mood-only Beats} |
| Est. length budget | | {Scenes → episode → chapter} |
| Chapter/part scope aligned | | |
| Prior hook addressed (ch 002+) | | {Also under Continuity below — one evidence cell is enough} |

### Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ / ❌ | |
| Series / overview tone & arc honored | | |
| Chapter List Role / Hook to Next honored | | |
| Part arc honored (series mode) | | |
| Characters from architecture; profiles not redefined | ✅ / ❌ | |
| Locations from architecture; profiles not redefined | | |
| World rules / history consistent with bible | | |
| No improvised entities or silent lore | | |
| Continuity files used (ch 002+) | | |
| Character/location state vs `story-so-far` | | |
| Unresolved threads: pick up / advance / plant / hold | | |
| No contradiction of released continuity | | |
| Conflicts section empty or escalated (not ignored) | | |

### Prior-Design Consistency Gate
{Copy pass/fail from `04-chapter-design.md` table — any ❌ → Design Verdict cannot be manuscript-ready}

### Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ / ⚠️ / ❌ | |
| Personal stake present | | |
| Chapter closing hook | | |
| Exposition budget respected | | |
| Seed discipline | | |
| Scene-first Key Events (all required fields) | | |
| Sensory-emotional on every scene | | |
| Motifs planned across episodes/scenes | | |

### Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | | {Ch 001: ~50:50?} |
| Sensory-emotional pairing | | |
| Dialogue voices + Dialogue intent | | |
| Reader-discovered meaning | | {Theme in Hold, not closing monologue?} |
| Antagonist plausibility | | |
| Closing image specified | | |

### Literary Awards Juror Checks (Design)
{Include when Literary Awards Juror is in the active set — see Persona Reference}

| Check | Result | Evidence |
|-------|--------|----------|
| Theme & vision earned (not asserted) | ✅ / ⚠️ / ❌ | |
| Original perspective or form planned | | |
| Human insight beyond plot mechanics | | |
| Artistic integrity — no planned didacticism or sentimentality | | |
| Chapter advances part/series literary arc | | |
| Competition readiness at design scope | | {What jurors would praise / hesitate over / reject} |

### Target Reader Checks (Design)
{Always — use audience locked in overview.md}

| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ / ⚠️ / ❌ | |
| Personal stake matches what this reader came for | | |
| Pacing / density fits platform expectations | | |
| Closing hook makes *this* reader want the next unit | | |
| No alienation of core audience without overview intent | | |

### Design Critique (required personas)
For **each** persona in the active set, write a block:

- `#### {Persona name}`
- Stance: {1–2 lines — what this persona values here}
- Strengths: {…}
- Defects: {finding → severity High/Med/Low → proposed fix}
- Reader impact: {how this affects the locked Target Reader}

### Design Adjudication
{Record Apply decisions. Default tie-break: Target Reader. Never silently drop a High finding.}

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|

### Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ / ❌ |
| Target-reader readiness | |
| Design quality | |
| Manuscript-ready | ✅ / ⬜ |

---

## 2. Manuscript Evaluation
{Perform after stage ⑤ — skip if evaluating design only}

### Consistency Checks (Design Fidelity)
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-generation sources loaded / used | ✅ / ⚠️ / ❌ | |
| Every Key Event scene present in order | | |
| No extra plot scenes without Key Event | | |
| Situation → Beat → Turn realized | | |
| POV / On stage / Location / When match | | |
| Dialogue intent + voices honored | | |
| Sensory-emotional realized (no catalog dump) | | |
| Seeds Plant/Hint only; Hold absent | | |
| Motifs / POV inserts at designed placements | | |
| Prior hook / closing hook / closing image | | |
| Continuity states & threads not contradicted | | |
| World / series tone not broken | | |
| Design-Fidelity Gate | | {All pass?} |

### Engagement Checks (Manuscript)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening hook effective | ✅ / ⚠️ / ❌ | |
| Opening question persists | | |
| Personal stake present | | |
| Closing hook / forward momentum | | |
| Exposition budget respected | | |
| Seed discipline | | |
| Scene-first prose | | |
| Dialogue naturalness | | |
| Episode transitions smooth | | |

### Literary Craft Checks (Manuscript)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | | |
| Sensory-emotional pairing | | |
| Prose rhythm varied | | |
| Dialogue voices distinct | | |
| POV inner/outer gap | | |
| Motifs threaded | | |
| POV insert discipline | | |
| World through character | | |
| Reader-discovered meaning | | |
| Emotion not over-labeled | | |
| Antagonist plausibility | | |
| Closing scene over statement | | |

### Literary Awards Juror Checks (Manuscript)
{Include when Literary Awards Juror is in the active set — see Persona Reference}

| Check | Result | Evidence |
|-------|--------|----------|
| Theme & vision embodied in prose | ✅ / ⚠️ / ❌ | |
| Originality visible on the page | | |
| Human insight lands beyond plot | | |
| Artistic integrity — no sentimentality, didacticism, or ornamental prose | | |
| Chapter contributes to work-level coherence in sequence | | |
| Competition readiness | | {Praise / hesitate / reject — with revision priorities} |

### Target Reader Checks (Manuscript)
{Always}

| Check | Result | Evidence |
|-------|--------|----------|
| Would the locked reader keep reading past the first page? | ✅ / ⚠️ / ❌ | |
| Emotional / curiosity payoff lands for that reader | | |
| Voice and density feel native to the platform/audience | | |
| Closing hook / next-chapter pull is concrete | | |
| Drop-risk moments (confusion, lecture, stall) identified | | |

### Manuscript Critique (required personas)
Same per-persona block as Design Critique for every persona in the active set.

### Manuscript Adjudication
Same table shape as Design Adjudication — Target Reader tie-break; never silently drop a High finding.

### Manuscript Verdict
| Dimension | Result |
|-----------|--------|
| Design fidelity | ✅ / ❌ |
| Target-reader readiness | |
| Manuscript quality | |
| Next-chapter readiness | |
| Release ready | ✅ / ⬜ |

---

## 3. Release
- **Released**: ✅ / ⬜
- **Date**:
```

---

## Persona Reference

Critics **advise**. They do not auto-apply changes — fill the Adjudication table and apply chosen fixes before proceeding. When personas conflict, **Target Reader** (from `overview.md`) is the default tie-break unless overview explicitly prioritizes prestige, niche, or another locked criterion.

### Default persona sets

| Context | **Required** (always run) | **Also required when** |
|---------|---------------------------|-------------------------|
| Design eval (any chapter) | **Target Reader**, **Genre Critic**, **Plot Expert**, **Reader-Editor** | Character-driven brief → **Character Critic**; ch 001 or major lore debut → **Setting/Lore Expert**; prestige/awards in overview → **Literary Awards Juror** |
| Manuscript eval (any chapter) | **Target Reader**, **Genre Critic**, **Reader-Editor**, **Literary Critic** | Same conditionals as design; after major rewrite → **Literary Critic** even mid-series |
| Ch 001 (design + manuscript) | Above + **Setting/Lore Expert** + **Literary Critic** (manuscript) | — |

Optional on request: any persona below; series/part-level critique → `evaluations/{scope}.md`.

### Persona catalog

| Persona | Focus | Typical questions |
|---------|-------|-------------------|
| **Target Reader** | The locked audience's lived read — fun, clarity, desire to continue, platform fit | Would *I* (this reader) keep going? Where do I skim, confuse, or quit? |
| **Reader-Editor** | Engagement craft, exposition restraint, serialization hooks | Is info dumped? Is the hook earned? Is the chapter sellable as a unit? |
| **Genre Critic** | Genre contracts, tropes, promise/payoff | Does it deliver what this genre's readers paid for — without empty cliché? |
| **Plot Expert** | Pacing, causality, structure, seed timing | Does every turn follow? Any stalled middle? Design ready to write? |
| **Character Critic** | Arc, motivation, distinct voice, relationship pressure | Why does this person act *now*? Voices interchangeable? |
| **Literary Critic** | Prose rhythm, motif, sensory-emotional craft, dialogue music | Does the page *feel*? Motifs earned? Emotion shown not labeled? |
| **Setting/Lore Expert** | World consistency, info-dumping, rule clarity | Is the world learned through character need — not catalog? |
| **Literary Awards Juror** | Theme, originality, insight, artistic integrity, work-level merit | Would this earn lasting literary respect — or only serial entertainment? |

**Literary Critic vs Literary Awards Juror:** Critic = craft on the page. Juror = whether the work *as literature* earns lasting merit at work/part level.

**Literary Awards Juror critique should cover:** theme & vision earned; originality; human insight; artistic integrity (no sentimentality/didacticism/ornament); work-level coherence; competition readiness with concrete revision priorities.

**Target Reader critique should cover:** who they are (from overview); what they came for; first-page / first-chapter retention; confusion or lecture risks; whether the closing hook pulls *them* specifically.

---

## Gate

After evaluation + adjudication (and any applied revisions): user reviews the evaluation, selects or confirms revision items as needed, and decides revise or proceed (to manuscript generation after design eval, or to release after manuscript eval).

## Reference-model integrity (architecture / design eval)

When evaluating design (⑥ after ④), also check `workflow/reference-models.md`:
- Character appearance/equipment states consistent; no silent gear swap
- Location used as set; lasting changes tracked
- Continuing situations cite stagings; L/R and stations do not flip without Design update
