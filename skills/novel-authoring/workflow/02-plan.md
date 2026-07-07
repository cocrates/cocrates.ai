# Stage ② — Plan (`series.md`)

**Prerequisites:** Approved `overview.md`

**Gate artifact:** `novels/{title-slug}/series.md`

**Next stage:** `03-architecture.md` (after user approves `series.md`)

---

## Procedure

Load approved `overview.md`. Design the **series-level blueprint** — always produced, from short story to multi-volume epic.

**Critical deliverable:** `series.md` covers **series arc and part-level structure only**. It must answer: *"What parts exist, and what role does each play in the series?"*

**Do not include specific chapter plans in `series.md`.** Individual chapter titles and summaries belong in `parts/{nnn}-{part-slug}.md` (stage ③).

**Chapter counts are planning estimates** — useful for scale management, not fixed requirements. Part chapter counts may shift as the story is architected and written. When counts change, update `overview.md` Scale and affected part files with user approval.

Write `series.md` with:

```markdown
# Series Blueprint: {Title}

## World Concept
{One-paragraph essence of the world}

## Time & Space Framework
{Historical period, geographical scope, temporal span of the story}

## Core Characters
{3–5 key characters with their role in the series arc}

## Series Arc
{The story's overall trajectory — beginning → major turning points → ending}

## Structure Scale
{Which levels this work uses — check all that apply}
- [ ] Part (major divisions within the series — always present; single-part works use one part)
- [ ] Chapter (publication / serialization unit — always present)
- [ ] Episode (design subdivision within chapter — count decided at chapter design, stage ④)

## Part Catalog

{Part-level skeleton only — no individual chapter titles or summaries.}

| Part | Title | Chapters (approx.) | Role in Series Arc |
|------|-------|-------------------|-------------------|
| 001 | {Title} | ~{n} | {one-line role} |
| 002 | {Title} | ~{n} | ... |

## Narrative Voice & Style
{POV, tense, stylistic notes}
```

### Single-Part Works

Even a one-chapter flash story uses **one row** in Part Catalog (e.g., Part 001, ~1 chapter). Chapter details go in `parts/001-{part-slug}.md` at stage ③ — not here.

---

## Completeness Check

- [ ] Series arc and world concept are defined
- [ ] Every part is listed in Part Catalog with **approximate** chapter count
- [ ] No individual chapter titles, summaries, or episode plans in `series.md`
- [ ] Structural levels are explicit
- [ ] Part chapter counts are estimates — no strict lock to `overview.md` required

---

## Gate

User approves `series.md`. Do not proceed to stage ③ without approval.
