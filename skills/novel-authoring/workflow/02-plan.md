# Stage ② — Plan (`series.md`)

**Prerequisites:** Approved `overview.md`

**Gate artifact:** `{project-root}/series.md`

**Next stage:** `03-architecture.md` (after user approves `series.md`)

---

## Procedure

Load approved `overview.md`. Design the **series-level blueprint**. Content depth depends on **Structure Mode**.

**Chapter counts are planning estimates** — useful for scale management, not fixed. When counts change, update `overview.md` Scale and the Chapter List source (`series.md` in short mode; `parts/` in series mode) with user approval.

**Episodes are never planned here** — episode count and Episode Index are decided at stage ④.


Confirm Structure Mode from `overview.md` (or set it now): **short** (~1 volume, under ~20 chapters) vs **series** (multi-part / long-form).

---

## Shared sections (both modes)

Write these in every `series.md`:

```markdown
# Series Blueprint: {Title}

## Structure Mode
{short | series} — locked from overview Scale

## World Concept
{One-paragraph essence of the world}

## Time & Space Framework
{Historical period, geographical scope, temporal span of the story}

## Core Characters
{3–5 key characters with their role in the series arc}

## Series Arc
{The story's overall trajectory — beginning → major turning points → ending}

## Structure Scale
- Part / Chapter / Episode(design) — see Structure Mode; episode count only at stage ④

## Narrative Voice & Style
{POV, tense, stylistic notes}
```

Then add the **mode-specific** block below.

---

## Short mode — parts + Chapter List in `series.md`

**Critical deliverable:** Answer *"What parts (if any) frame the work, and what is every chapter's role?"*

Include **Part composition** and a **full Chapter List**. Do **not** create `parts/`.

```markdown
## Part Composition

| Part | Title | Chapters (approx.) | Role in Arc |
|------|-------|-------------------|-------------|
| 001 | {Title} | ~{n} | {one-line role} |

{Optional second part row only if the short work still uses light major divisions — still no `parts/` files. If you need several substantial parts, switch Structure Mode to **series** instead.}


## Chapter List

{Full chapter catalog for the work. Flow map — Role + Hook to Next required.}

| Ch | Title | Part | Role | Key POV | Hook to Next |
|----|-------|------|------|---------|--------------|
| 001 | {Title} | 001 | {one-line role} | {character} | {tension / unanswered beat → Ch 002} |
| 002 | {Title} | 001 | ... | ... | {→ Ch 003} |
| {last} | {Title} | 001 | ... | ... | {closing beat} |

## Work Hooks
- **Opening hook** (chapter first):
- **Closing hook** (chapter last):
```

**Hook to Next:** Required on every row so chapter-to-chapter momentum can be reviewed in one table. Stage ④ Closing hook refines the matching row — keep them aligned.

---

## Series mode — Part Catalog only

**Critical deliverable:** Answer *"What parts exist, and what role does each play in the series?"*

**Do not include specific chapter plans in `series.md`.** Chapter titles, Role, and Hook to Next belong in `parts/{nnn}-{part-slug}.md` (stage ③).

```markdown
## Part Catalog

{Part-level skeleton only — no individual chapter titles or summaries.}

| Part | Title | Chapters (approx.) | Role in Series Arc |
|------|-------|-------------------|-------------------|
| 001 | {Title} | ~{n} | {one-line role} |
| 002 | {Title} | ~{n} | ... |
```

---

## Completeness Check

**Both modes:**
- [ ] Structure Mode is explicit (`short` or `series`)
- [ ] Series arc and world concept are defined
- [ ] No episode counts or episode plans

**Short mode:**
- [ ] Part Composition present
- [ ] Chapter List complete with Role + Hook to Next; consecutive hooks connect
- [ ] No `parts/` directory planned or created

**Series mode:**
- [ ] Every part listed in Part Catalog with approximate chapter count
- [ ] No individual chapter titles, summaries, or chapter lists in `series.md`

---

## Gate

User approves `series.md`. Do not proceed to stage ③ without approval.
