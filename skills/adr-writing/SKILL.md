---
name: adr-writing
description: >-
  Creates Any Decision Records (ADR) for the user to analyze, review, and
  approve among alternatives for a specific concern before committing. Select
  when the user asks for options, tradeoffs, a recommended direction, or needs
  to decide among valid approaches — architectural or otherwise. Explicitly
  identifies and guides the user through downstream/derived concerns that
  arise from a chosen solution.
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

## Project Folder Layout

ADRs live under the project's **`adr/`** directory:

```
{project-slug}/
├── spec/
│   ├── ASR.md
│   ├── PRD.md
│   └── {requirement-slug}.md
├── adr/
│   └── {concern-slug}.md
├── verification/
└── …
```

- **`{project-slug}`:** English **kebab-case** project folder (same as used by `spec-driven-generation` / `spec-writing`). Resolve from the user's request or an existing project folder; ask if ambiguous.
- **ADR root:** `{project-slug}/adr/` (create if missing)
- **ASR registry:** `{project-slug}/spec/ASR.md` — keep Related ASRs / Related ADRs in sync when creating or approving ADRs
- **Filename:** English **kebab-case** for the concern (e.g. `item-storage-database.md`, `report-intro-structure.md`)
- One file = one concern
- **Do not** place `adr/` at the repository root — always nest under `{project-slug}/`.

### ASR ↔ ADR (Many-to-Many)

- **One ASR → many ADRs:** A single ASR may need several design reviews. Open a separate ADR per distinct concern; list all of them under that ASR in `ASR.md`.
- **One ADR → many ASRs:** When one concern spans multiple ASRs, list every related ASR ID in the ADR's **Related ASRs** section and mirror the ADR path under each ASR in `ASR.md`.
- Design review flow is always **ASR → ADR** for complex trade-offs; Direct Input may skip ADR but must still update `ASR.md`.

### Before Creating a New ADR

Before saving, check the `{project-slug}/adr/` directory.

- Search for files with the **same or similar concern** (compare filename, `## Concern`, `## Tags`, Related ASRs)
- If found, **do not create a new file** — **supplement and merge** into the existing ADR, or set `Status: superseded` and link to a replacing ADR when the concern is reframed
- If not found, create a new `{concern-slug}.md`

---

## Workflow

### 1. Identify the Concern & Related ASRs

Clarify the **concern** — the specific question the user must resolve — and why it matters now.
Capture briefly: **Concern**, **Context**, **Stakes**, and **Related ASR ID(s)** from `{project-slug}/spec/ASR.md` (create or amend ASR entries if the concern reveals a new ASR).

### 2. Enumerate Viable Options

Present **at least 2–3 valid options** using concise **bullet points** (name + 2–4 bullets covering what it is, pro, and con). Options must be real, non-strawman alternatives.

### 3. Compare Tradeoffs

Add a short, scannable **Tradeoffs** section (comparison table or bullet pairs) and an optional agent recommendation.

### 4. Save ADR File & Sync ASR.md

Write the ADR using the template below. Set `Status: proposed` and save to `{project-slug}/adr/{concern-slug}.md`.
Then update `{project-slug}/spec/ASR.md`:
- Set each related ASR to `reviewing` (if not already).
- Append this ADR path under **Related ADRs** for every linked ASR.
- Ensure the ADR lists those ASR IDs under **Related ASRs**.

### 5. User Review Gate

Ask the user to review and approve via structured Socratic guidance. Do not mark as approved without explicit user confirmation.

### 6. Record Approval & Identify Downstream Concerns

When the user approves an option:
1. Set `Status: approved`.
2. Fill **Decision** with the chosen option and a one-line rationale.
3. Record `Approved: YYYY-MM-DD`.
4. **Sync ASR.md:** For each Related ASR, update **Related ADRs** status to approved; if all ADRs needed for that ASR are approved (or the ASR is fully resolved by this ADR), set ASR status to `designed` and fill **Resolution** with the concrete outcome.
5. **Identify Downstream Concerns (Crucial):** Analyze the approved solution and proactively identify **derived questions or operational details** that must be answered next.
   - *Example:* If the user approves **"Introducing Redis Cache"** for performance, the agent must immediately identify downstream concerns such as **Cache Eviction Policy (TTL/LRU)** and **Cache Size/Memory Limit**.
   - Register new derived items as new ASRs in `ASR.md` when they are architecturally significant; otherwise keep them as Downstream Concerns that may spawn follow-up ADRs.
6. Record these under **`## Downstream Concerns`** in the file template to maintain architectural lineage.

### 7. Chain or Hand Off

- **If Downstream Concerns are critical:** Ask the user if they want to immediately open a subsequent ADR to address the newly generated questions (e.g., *"Now that we decided to use a Cache, what should our Eviction Policy be?"*). Link any new ADR to the relevant ASR ID(s).
- **If all blocking concerns are resolved:** Load and follow the **`spec-writing`** skill to consolidate approved decisions into the self-contained Spec, then mark related ASRs `approved` in `ASR.md` after Spec sync and user confirmation.

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

## Related ASRs
- ASR-00x — {title} — {how this ADR addresses it}
- ASR-00y — {title} — {optional additional ASRs this ADR covers}

## Downstream Concerns
- {List new questions triggered by this approval that require future decisions}
- {e.g., If Cache approved -> [ ] Define Cache TTL policy, [ ] Determine Max Memory limit}

## Related
- {Links to parent ADRs, downstream ADRs, or specs}

## Tags
`tag-one`, `tag-two`

## Approved
- YYYY-MM-DD: {option chosen, by user confirmation}
```

Omit **Related ASRs** only when the ADR is exploratory and no ASR ID exists yet — then create the ASR in `ASR.md` in the same turn and link it.
---

## Example (Chained/Derived Concerns)

File: `{project-slug}/adr/application-performance-cache.md`

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

## Related ASRs
- ASR-003 — Catalog read latency — primary concern this ADR resolves
- ASR-001 — Deliverable form (API server) — constrains deployment options for the cache

## Downstream Concerns
- [ ] **Cache Eviction & TTL Policy:** How long should catalog data live? How do we handle manual invalidation? (Triggers next ADR)
- [ ] **Cache Size & Memory Constraints:** What happens when memory is full?

## Related
- Follow-up needed: `{project-slug}/adr/cache-eviction-policy.md`

## Approved
- 2026-06-29: Option A, user confirmed
```

---

## Dialogue Rules

1. **State the current step** (identify concern → enumerate options → save → review → approve → identify downstream concerns → chain/spec handoff).
2. **Proactive Chaining Guidance:** When a user approves an option, congratulate the choice and **immediately present the next 1–2 critical sub-questions** that logically follow, offering to draft the next ADR right away.
   - *Example:* *"Great choice! You've decided to implement a Cache for performance optimization, and I have recorded this approval in `{project-slug}/adr/application-performance-cache.md`. By introducing a cache, we now have downstream concerns that need attention: **1) Cache Eviction & TTL Policy**, and **2) Cache Size & Memory Constraints**. Would you like to proceed with drafting the next ADR for the Cache Eviction Policy right away?"*
3. When the user says "just pick one," still present alternatives with tradeoffs, recommend, and explicitly highlight what downstream obligations that choice brings.

---

## Prohibitions

- Fewer than **2 viable options** in an ADR.
- Writing ADR files at the repository root instead of under `{project-slug}/adr/`.
- Creating or approving an ADR **without** linking Related ASR ID(s) and syncing `{project-slug}/spec/ASR.md`.
- Closing an ADR loop **without identifying or asking about downstream consequences/derived concerns** inherent to the chosen solution.
- Marking `approved` without explicit user confirmation.

---

## Completion Criteria

- Concern identified and stated as a clear question, with Related ASR ID(s) recorded.
- ADR saved at `{project-slug}/adr/{concern-slug}.md` with **≥ 2 viable options**.
- `{project-slug}/spec/ASR.md` updated (status + Related ADRs) for every linked ASR — supporting one-to-many and many-to-one links.
- User has reviewed and **explicitly approved** the decision.
- **Downstream Concerns** triggered by the approved option are explicitly enumerated in the file.
- User guided on whether to chain into a follow-up ADR or hand off to spec-writing.    