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

> *"Chapter {nnn} design is approved. Shall we evaluate the design before writing the manuscript? We'll validate structure, seeds, and engagement against criteria."*

Load:
- `chapters/{nnn}-{chapter-slug}.md` (full file — chapter + episode sections + Prior Design Alignment)
- **Same prior-design set as stage ④ load:** `overview.md`, `series.md`, Chapter List source, `characters*`, `locations*`, `world-bible` (+ aspects), continuity (ch 002+)
- Re-run Prior-Design Consistency Gate from `04-chapter-design.md` against the finished design
- **Literary Awards Juror:** also Chapter List source for work-level context — **Short:** `series.md`; **Series:** `parts/{nnn}-{part-slug}.md`; fill **Literary Awards Juror Checks (Design)** before persona critique

### Manuscript Evaluation

Propose after chapter manuscript approval:

> *"Chapter {nnn} manuscript is ready. Shall we evaluate it? We'll start with validation against the chapter design and prior architecture/continuity. Would you also like a critique from a specific critic persona?"*

Load:
- `manuscripts/{nnn}-{chapter-slug}.md`
- `chapters/{nnn}-{chapter-slug}.md` (full design)
- Architecture + continuity set listed in the chapter design (same as stage ⑤ Pre-Generation Load)
- Re-run **Design-Fidelity Gate** from `05-generation.md`

**Literary Awards Juror:** When this persona is requested, fill **Literary Awards Juror Checks** in the evaluation file (design and/or manuscript section) before the persona critique. Load Chapter List source (**Short:** `series.md`; **Series:** parent `parts/{nnn}-{part-slug}.md`) and `overview.md` for work-level context.

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
{Include when Literary Awards Juror persona is requested — see Persona Reference}

| Check | Result | Evidence |
|-------|--------|----------|
| Theme & vision earned (not asserted) | ✅ / ⚠️ / ❌ | |
| Original perspective or form planned | | |
| Human insight beyond plot mechanics | | |
| Artistic integrity — no planned didacticism or sentimentality | | |
| Chapter advances part/series literary arc | | |
| Competition readiness at design scope | | {What jurors would praise / hesitate over / reject} |

### Design Critique
{Persona critique if requested}

### Design Revision Decisions
| # | Finding | Apply? | Action Taken | Status |
|---|---------|--------|-------------|--------|

### Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ / ❌ |
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
{Include when Literary Awards Juror persona is requested — see Persona Reference}

| Check | Result | Evidence |
|-------|--------|----------|
| Theme & vision embodied in prose | ✅ / ⚠️ / ❌ | |
| Originality visible on the page | | |
| Human insight lands beyond plot | | |
| Artistic integrity — no sentimentality, didacticism, or ornamental prose | | |
| Chapter contributes to work-level coherence in sequence | | |
| Competition readiness | | {Praise / hesitate / reject — with revision priorities} |

### Manuscript Critique
{Persona critique if requested}

### Manuscript Revision Decisions
| # | Finding | Apply? | Action Taken | Status |
|---|---------|--------|-------------|--------|

### Manuscript Verdict
| Dimension | Result |
|-----------|--------|
| Design fidelity | ✅ / ❌ |
| Manuscript quality | |
| Next-chapter readiness | |
| Release ready | ✅ / ⬜ |

---

## 3. Release
- **Released**: ✅ / ⬜
- **Date**:
```

### Persona Reference

| Persona | Focus | When to suggest |
|---------|-------|-----------------|
| **Genre Critic** | Genre conventions, tropes | Always |
| **Literary Critic** | Prose rhythm, motif, sensory-emotional craft, dialogue voice | **Recommended for ch 001** and after major rewrites |
| **Literary Awards Juror** | Thematic depth, originality, artistic integrity, moral complexity, language as literary art, work-level coherence, lasting literary merit — criteria typical of major prizes and competition judging | **Recommended** when targeting awards, contests, or literary prestige; major part/manuscript milestones |
| **Plot Expert** | Pacing, causality, structure | **Recommended for design evaluation** |
| **Character Critic** | Arc, motivation, voice | Character-driven works |
| **Reader-Editor** | Engagement, exposition restraint | **Default for serialization** |
| **Setting/Lore Expert** | World consistency, info-dumping | **Required for ch 001** |

**Literary Critic vs Literary Awards Juror:** Critic = craft execution on the page. Juror = whether the work *as literature* earns lasting merit — theme, vision, insight, and form at work/part level, not only line-level polish.

**Literary Awards Juror critique should cover:**
- **Theme & vision** — Is the central concern substantial, coherent, and earned (not asserted)?
- **Originality** — What is genuinely new in perspective, form, or moral imagination?
- **Human insight** — Does the work illuminate condition, society, or consciousness beyond plot?
- **Artistic integrity** — Are ambition, structure, and language aligned — no sentimentality, didacticism, or ornamental prose?
- **Work-level coherence** — Do part/chapter arcs accumulate into a whole greater than serial entertainment?
- **Competition readiness** — What would jurors praise, hesitate over, or reject — with concrete revision priorities?

Critique may be requested at any stage (part, character web, series level) — produce appropriate `evaluations/{scope}.md`.

---

## Gate

User reviews evaluation, selects revision items, decides revise or proceed (to manuscript generation after design eval, or to release after manuscript eval).
