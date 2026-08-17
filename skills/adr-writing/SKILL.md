---
name: adr-writing
description: >-
  Creates Any Decision Records (ADR) for the user to analyze, review, and
  approve among alternatives for a specific concern before committing. Select
  when the user asks for options, tradeoffs, a recommended direction, or needs
  to decide among valid approaches — architectural or otherwise. Explicitly
  identifies and guides the user through downstream/derived concerns that
  arise from a chosen solution.
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

Match the user's language (Korean, English, etc.) in user-facing messages. ADR file content may use the language of the decision context; ADR concern filenames stay English kebab-case.

## Resolve Project Root

Before writing ADRs, resolve **`{project-root}`** — the folder where decision records for this topic should live.

- **Default (Cocrates workspace):** `{types}/{project-slug}/` per the Cocrates Workspace Convention (read via `read_agents: principles/02-workspace-conventions.md`). Pick the `{types}` folder that matches the deliverable context (e.g. `softwares/`, `documents/`, `novels/`).
- **Ad hoc:** When no project folder exists yet, use a lightweight root the user confirms (a task folder or existing deliverable root) and place `adr/` there.

**Rules:** Reuse an existing project folder wherever it is; **confirm location and name with the user before creating** a new one.

## Project Folder Layout

ADRs live under `{project-root}/adr/`:

```
{project-root}/
├── adr/
│   └── {concern-slug}.md
└── …
```

- **`{project-slug}`:** URL-friendly kebab-case slug in the user's language (Korean allowed). Used as the task folder name when applicable.
- **ADR root:** `{project-root}/adr/` (create if missing)
- **Filename:** English **kebab-case** for the concern (e.g. `item-storage-database.md`, `report-intro-structure.md`)
- One file = one concern
- Place `adr/` under `{project-root}` only.

### ADR Chaining (Many-to-Many via Links)

- **One parent decision → many ADRs:** A broad choice may surface several follow-up questions. Open a separate ADR per distinct concern; link them under **Related** in each file.
- **One ADR → many prior decisions:** When one concern depends on earlier choices, list parent or sibling ADR paths under **Related** so lineage stays traceable.
- ADR is the default review vehicle whenever a concern has **non-trivial tradeoffs** and the user must explicitly approve before proceeding.

### Before Creating a New ADR

Before saving, check the `{project-root}/adr/` directory.

- Search for files with the **same or similar concern** (compare filename, `## Concern`, `## Tags`, **Related** links)
- If found, **do not create a new file** — **supplement and merge** into the existing ADR, or set `Status: superseded` and link to a replacing ADR when the concern is reframed
- If not found, create a new `{concern-slug}.md`

---

## Workflow

### 1. Identify the Concern

Clarify the **concern** — the specific question the user must resolve — and why it matters now.
Capture briefly: **Concern**, **Context**, and **Stakes**.

### 2. Enumerate Viable Options

Present **at least 2–3 valid options** using concise **bullet points** (name + 2–4 bullets covering what it is, pro, and con). Options must be real, non-strawman alternatives.

### 3. Compare Tradeoffs

Add a short, scannable **Tradeoffs** section (comparison table or bullet pairs) and an optional agent recommendation.

### 4. Save ADR File

Write the ADR using the template below. Set `Status: proposed` and save to `{project-root}/adr/{concern-slug}.md`.

### 5. User Review Gate

Ask the user to review and approve via structured Socratic guidance. Do not mark as approved without explicit user confirmation.

### 6. Record Approval & Identify Downstream Concerns

When the user approves an option:
1. Set `Status: approved`.
2. Fill **Decision** with the chosen option and a one-line rationale.
3. Record `Approved: YYYY-MM-DD`.
4. **Identify Downstream Concerns (Crucial):** Analyze the approved solution and proactively identify **derived questions or operational details** that must be answered next.
   - *Example:* If the user approves **"Introducing Redis Cache"** for performance, the agent must immediately identify downstream concerns such as **Cache Eviction Policy (TTL/LRU)** and **Cache Size/Memory Limit**.
   - Record each item under **`## Downstream Concerns`**; open a follow-up ADR when a downstream item itself needs option comparison and user approval.
5. Update **Related** links to parent ADRs or sibling ADRs when applicable.

### 7. Chain Follow-Up ADRs

When **Downstream Concerns** are critical, ask the user if they want to immediately open a subsequent ADR to address the newly generated questions (e.g., *"Now that we decided to use a Cache, what should our Eviction Policy be?"*). Link the new ADR back to the parent under **Related**.

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

## Tradeoffs
| | Option A | Option B |
|---|----------|----------|
| {dimension} | … | … |

## Recommendation (optional)
- {Agent suggestion and why}

## Consequences
- {What follows if the approved option is chosen}
- {What we are explicitly not doing}

## Downstream Concerns
- {List new questions triggered by this approval that require future decisions}
- {e.g., If Cache approved -> [ ] Define Cache TTL policy, [ ] Determine Max Memory limit}

## Related
- {Links to parent ADRs, downstream ADRs, or other related artifacts}

## Tags
`tag-one`, `tag-two`

## Approved
- YYYY-MM-DD: {option chosen, by user confirmation}
```

---

## Example (Chained/Derived Concerns)

File: `{project-root}/adr/application-performance-cache.md`

```markdown
# Application Performance Cache

## Concern
How should we improve response latency for the item catalog API?

## Status
approved

## Context
- Database read operations are hitting high CPU limits during peak hours.
- Item catalog data changes infrequently (a few times a day).

## Decision
**Option A — In-Memory Redis Cache**
User-approved: Offload read traffic from the primary DB using a dedicated caching layer.

## Options
### Option A — In-Memory Redis Cache
- Deploy a standalone Redis instance in front of the DB.
- Pro: Sub-millisecond read latency, high throughput.
- Con: Introduces network hop and data stale risk.

### Option B — Database Read Replicas
- Scale horizontally by adding DB replica nodes.
- Pro: Familiar SQL pooling; zero application caching logic needed.
- Con: Higher infrastructure cost; eventual consistency lag.

## Downstream Concerns
- [ ] **Cache Eviction & TTL Policy:** How long should catalog data live? How do we handle manual invalidation? (Triggers next ADR)
- [ ] **Cache Size & Memory Constraints:** What happens when memory is full?

## Related
- Follow-up needed: `{project-root}/adr/cache-eviction-policy.md`

## Approved
- 2026-06-29: Option A, user confirmed
```

---

## Dialogue Rules

1. **State the current step** (identify concern → enumerate options → save → review → approve → identify downstream concerns → chain follow-up ADRs).
2. **Proactive Chaining Guidance:** When a user approves an option, congratulate the choice and **immediately present the next 1–2 critical sub-questions** that logically follow, offering to draft the next ADR right away.
   - *Example:* *"Great choice! You've decided to implement a Cache for performance optimization, and I have recorded this approval in `{project-root}/adr/application-performance-cache.md`. By introducing a cache, we now have downstream concerns that need attention: **1) Cache Eviction & TTL Policy**, and **2) Cache Size & Memory Constraints**. Would you like to proceed with drafting the next ADR for the Cache Eviction Policy right away?"*
3. When the user says "just pick one," still present alternatives with tradeoffs, recommend, and explicitly highlight what downstream obligations that choice brings.

---

## Prohibitions

- Fewer than **2 viable options** in an ADR.
- Writing ADR files outside `{project-root}/adr/`.
- Closing an ADR loop **without identifying or asking about downstream consequences/derived concerns** inherent to the chosen solution.
- Marking `approved` without explicit user confirmation.

---

## Completion Criteria

- Concern identified and stated as a clear question.
- ADR saved at `{project-root}/adr/{concern-slug}.md` with **≥ 2 viable options**.
- User has reviewed and **explicitly approved** the decision.
- **Downstream Concerns** triggered by the approved option are explicitly enumerated in the file.
- User guided on whether to chain into a follow-up ADR for any critical downstream concern.
