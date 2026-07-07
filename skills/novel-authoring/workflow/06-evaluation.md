# Stage ⑥ — Evaluation (`evaluations/{nnn}-{chapter-slug}.md`)

**Prerequisites:** Depends on mode — see below

**Gate artifact:** `novels/{title-slug}/evaluations/{nnn}-{chapter-slug}.md`

**Next stage:** `07-revision.md` if issues found; `05-generation.md` (after design eval) or `08-release.md` (after manuscript eval)

**Engagement checks:** See checklist in `04-chapter-design.md`

---

## Two Evaluation Modes

| Mode | When | Inputs | Purpose |
|------|------|--------|---------|
| **Design Evaluation** | After stage ④ approval (recommended before ⑤) | `chapters/{nnn}-{chapter-slug}.md` + all files in `chapters/{nnn}-{chapter-slug}/` | Catch structural, pacing, seed, and engagement issues **before** writing prose |
| **Manuscript Evaluation** | After stage ⑤ approval | Above + `manuscripts/{nnn}-{chapter-slug}.md` | Validate prose against design and criteria |

**Design evaluation is strongly recommended.** Fixing structure at design stage is more efficient than revising finished prose.

Use a **single evaluation file** per chapter — append or update sections as evaluation progresses.

---

## Procedure

### Design Evaluation

Propose after chapter design approval:

> *"Chapter {nnn} design is approved. Shall we evaluate the design before writing the manuscript? We'll validate structure, seeds, and engagement against criteria."*

Load:
- `chapters/{nnn}-{chapter-slug}.md`
- Every file in `chapters/{nnn}-{chapter-slug}/`
- `overview.md`, architecture references, continuity (ch 002+)

### Manuscript Evaluation

Propose after chapter manuscript approval:

> *"Chapter {nnn} manuscript is ready. Shall we evaluate it? We'll start with validation against criteria. Would you also like a critique from a specific critic persona?"*

Load design files + `manuscripts/{nnn}-{chapter-slug}.md`.

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
| Episode List fully designed | | {All episode files present?} |
| Chapter/part scope aligned | | |
| Prior hook addressed (ch 002+) | | |

### Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Characters from architecture | ✅ / ❌ | |
| Locations from architecture | | |
| World rules consistent | | |
| No improvised entities | | |
| Continuity files loaded (ch 002+) | | |
| Character/location state consistent | | |
| Unresolved threads handled | | |

### Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ / ⚠️ / ❌ | |
| Personal stake present | | |
| Chapter closing hook | | |
| Exposition budget respected | | |
| Seed discipline | | |
| Scene-first Key Events | | |
| Motifs planned across episodes | | |

### Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | | {Ch 001: ~50:50?} |
| Sensory-emotional pairing | | |
| Dialogue voices distinct | | |
| Reader-discovered meaning | | {Theme in Hold, not closing monologue?} |
| Antagonist plausibility | | |
| Closing image specified | | |

### Design Critique
{Persona critique if requested}

### Design Revision Decisions
| # | Finding | Apply? | Action Taken | Status |
|---|---------|--------|-------------|--------|

### Design Verdict
| Dimension | Result |
|-----------|--------|
| Design quality | |
| Manuscript-ready | ✅ / ⬜ |

---

## 2. Manuscript Evaluation
{Perform after stage ⑤ — skip if evaluating design only}

### Consistency Checks
| Check | Result | Evidence |
|-------|--------|----------|

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

### Manuscript Critique
{Persona critique if requested}

### Manuscript Revision Decisions
| # | Finding | Apply? | Action Taken | Status |
|---|---------|--------|-------------|--------|

### Manuscript Verdict
| Dimension | Result |
|-----------|--------|
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
| **Plot Expert** | Pacing, causality, structure | **Recommended for design evaluation** |
| **Character Critic** | Arc, motivation, voice | Character-driven works |
| **Reader-Editor** | Engagement, exposition restraint | **Default for serialization** |
| **Setting/Lore Expert** | World consistency, info-dumping | **Required for ch 001** |

Critique may be requested at any stage (part, character web, series level) — produce appropriate `evaluations/{scope}.md`.

---

## Gate

User reviews evaluation, selects revision items, decides revise or proceed (to manuscript generation after design eval, or to release after manuscript eval).
