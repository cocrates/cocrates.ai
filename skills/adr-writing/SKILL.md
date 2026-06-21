---
name: adr-writing
description: >-
  Creates Any Decision Records (ADR) for the user to analyze, review, and
  approve among alternatives for a specific concern before committing. Select
  when the user asks for options, tradeoffs, a recommended direction, or needs
  to decide among valid approaches — architectural or otherwise.
compatibility: opencode
metadata:
  agent: cocrates
---

# ADR Writing — Any Decision Record Skill

An ADR here is an **Any Decision Record**: a short, reviewable document for **one concern** — a specific question the user must resolve — with **viable options** to compare, tradeoffs, and the user's approved choice. It is **not limited to software architecture** — technology choices, document structure, workflow design, naming, scope, and any other concern that affects a deliverable qualify.

**Concern vs options:**

| | Concern | Options |
|---|---------|---------|
| **What** | The question or topic under review | The valid alternatives (A, B, C, …) |
| **Example** | Which database to use for item storage | MySQL, MongoDB, PostgreSQL, … |
| **Example** | How to structure the report introduction | Chronological, problem-first, executive summary |

One ADR file = one concern. Options are the answer candidates for that concern — not separate files.

Match the user's language (Korean, English, etc.) in user-facing messages. ADR file content may use the language of the decision context; filenames and slugs stay English kebab-case.

## Storage Location

```
adr/
└── {concern-slug}.md
```

- **ADR root:** `adr/` at the project root (create if missing)
- **Filename:** English **kebab-case** for the concern (e.g. `item-storage-database.md`, `report-intro-structure.md`)
- One file = one concern

### Before Creating a New ADR

Before saving, check the `adr/` directory.

- Search for files with the **same or similar concern** (compare filename, `## Concern`, `## Tags`)
- If found, **do not create a new file** — **supplement and merge** into the existing ADR, or set `Status: superseded` and link to a replacing ADR when the concern is reframed
- If not found, create a new `{concern-slug}.md`

**Similar-concern examples:**
- `item-db.md` ↔ `item-storage-database.md` → same concern, merge into one
- `auth-method.md` updated because OAuth scope changed → supplement existing file or supersede with clear link

---

## Workflow

### 1. Identify the Concern

Clarify the **concern** — the specific question the user must resolve — and why it matters now.

Capture briefly:

- **Concern** — one clear question (e.g. *"Which database should we use for item storage?"*)
- **Context** — what constraints, goals, or prior decisions apply?
- **Stakes** — what does resolving this concern affect downstream?

If multiple independent concerns exist, create **separate ADR files** — one concern per file, each with its own options.

### 2. Enumerate Viable Options

Present **at least 2–3 valid options**. More is fine when each option is genuinely distinct and worth considering.

**Option writing rules:**

- Use **bullet points**, not long prose paragraphs
- Each option: **name + 2–4 bullets** covering what it is, key benefit, key cost or risk
- Options must be **real alternatives** the user could actually choose — not strawmen or a single disguised recommendation
- Include a "do nothing / defer" option only when it is truly viable

**Bad (too verbose):**
> Option A would involve implementing a comprehensive microservices architecture where each bounded context is deployed independently, which provides excellent scalability but requires significant operational overhead…

**Good (concise bullets):**

Concern: *Which database should we use for item storage?*

> **Option A — MySQL**
> - Relational store; mature ecosystem
> - Pro: strong consistency, familiar SQL tooling
> - Con: rigid schema for highly variable item attributes
>
> **Option B — MongoDB**
> - Document store; flexible per-item schema
> - Pro: fits heterogeneous item shapes
> - Con: fewer join patterns; consistency model differs from SQL

### 3. Compare Tradeoffs

After options, add a short **Tradeoffs** section:

- Comparison table or bullet pairs (option ↔ consequence)
- Non-obvious constraints, dependencies, or reversibility
- Agent recommendation (optional, clearly labeled) — never substitute for user approval

Keep this section scannable. The goal is decision support, not a white paper.

### 4. Save ADR File

Write the ADR using the template below. Set `Status: proposed`.

Save to `adr/{concern-slug}.md`.

### 5. User Review Gate

After saving, ask the user to review and approve.

Example prompts:

- *"Please review `adr/item-storage-database.md`. Is the concern framed correctly? Are the options complete and fairly stated? Which option do you approve?"*
- *"Option B is recommended because … — do you approve B, or prefer another option?"*

**Do not:**

- Mark a decision as approved without explicit user confirmation
- Proceed to `spec-writing` before approval
- Collapse multiple options into a single "obvious" choice without presenting alternatives

### 6. Record Approval

When the user approves:

1. Set `Status: approved`
2. Fill **Decision** with the chosen option and one-line rationale (user's words when possible)
3. Record `Approved: YYYY-MM-DD`
4. Notify the user of the updated path

If the user rejects all options or requests new ones, revise the ADR and return to step 2.

### 7. Hand Off to Spec Writing

After **all blocking ADRs for the current design scope** are approved, load and follow the **`spec-writing`** skill.

Transfer to spec-writing:

- Every approved decision as a **fully stated specification** (what to build — not a reference to an ADR)
- Constraints and rejected options worth documenting (briefly, in the Spec)
- Open questions deferred to spec or a later ADR

ADR files remain the **decision audit trail** — for understanding alternatives and selection rationale. The Spec becomes the **single, self-contained source of truth** for generation; downstream steps must not require reading ADRs.

---

## File Template

```markdown
# {Concern Title}

## Concern
{The question under review — e.g. Which database should we use for item storage?}

## Status
proposed | approved | rejected | superseded

## Context
- {Why this concern matters now}
- {Relevant constraints or prior decisions}

## Decision
{Leave empty while Status is proposed. After approval: chosen option + one-line rationale.}

## Options

### Option A — {Short Name}
- {What it is}
- Pro: {key benefit}
- Con: {key cost or risk}

### Option B — {Short Name}
- {What it is}
- Pro: {key benefit}
- Con: {key cost or risk}

### Option C — {Short Name}
- {What it is}
- Pro: {key benefit}
- Con: {key cost or risk}

## Tradeoffs
| | Option A | Option B | Option C |
|---|----------|----------|----------|
| {dimension} | … | … | … |

- {Additional note on reversibility, dependencies, or timing}

## Recommendation (optional)
- {Agent suggestion and why — not a substitute for user approval}

## Consequences
- {What follows if the approved option is chosen}
- {What we are explicitly not doing}

## Related
- {Link to other adr files, specs, or kb files}

## Tags
`tag-one`, `tag-two`

## Approved
- YYYY-MM-DD: {option chosen, by user confirmation}
```

Omit empty sections except **Concern**, **Options** (required, minimum 2), and **Status**.

---

## Example (abbreviated)

File: `adr/item-storage-database.md`

```markdown
# Item Storage Database

## Concern
Which database should we use for item storage?

## Status
approved

## Context
- Item attributes vary by category; some fields are optional
- Team has SQL experience; no existing NoSQL operations practice
- Expected volume: moderate; single-region deployment

## Decision
**Option B — PostgreSQL**
User-approved: relational integrity for core fields plus JSONB for variable attributes.

## Options

### Option A — MySQL
- Relational store; wide hosting support
- Pro: familiar SQL; strong ecosystem
- Con: less flexible for heterogeneous item attributes without schema churn

### Option B — PostgreSQL
- Relational store with JSONB for semi-structured fields
- Pro: schema where it matters; flexibility where needed
- Con: slightly more complex queries for mixed relational/document access

### Option C — MongoDB
- Document store; one document per item
- Pro: natural fit for varying item shapes
- Con: team learning curve; different consistency and query patterns

## Tradeoffs
| | MySQL | PostgreSQL | MongoDB |
|---|-------|------------|---------|
| Variable item fields | Migrations | JSONB | Native |
| Team familiarity | High | High | Low |
| Relational integrity | Strong | Strong | Weaker |

## Consequences
- Items stored in PostgreSQL; variable attributes in JSONB columns
- MongoDB driver and ops tooling not introduced

## Tags
`database`, `item-storage`, `persistence`

## Approved
- 2026-06-19: Option B, user confirmed
```

---

## Dialogue Rules

1. **State the current step** (identify concern → enumerate options → save → review → approve → spec handoff).
2. After saving, report the **file path** and what to review.
3. Ask Socratic questions when the concern is vague — before drafting options.
4. When the user says "just pick one," still present 2–3 viable options with tradeoffs; then recommend and ask for approval.
5. Multiple ADRs may run in sequence; track which concerns are still `proposed` vs `approved`.

---

## Prohibitions

- Fewer than **2 viable options** in an ADR
- Long narrative option descriptions instead of concise bullets
- Marking `approved` without explicit user confirmation
- Duplicating the same concern across multiple ADR files

---

## Completion Criteria

- Concern identified and stated as a clear question
- ADR saved at `adr/{concern-slug}.md` with **≥ 2 viable options** in bullet form
- User has reviewed and **explicitly approved** the decision (or requested revision)
- `Status`, **Decision**, and **Approved** sections updated accordingly
- User informed of next step: **`spec-writing`** to consolidate approved decisions into the spec

After handoff, the spec-writing skill owns consolidation; this skill does not write the spec itself.
