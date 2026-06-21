---
name: spec-writing
description: >-
  Creates and updates Spec documents — consolidating requests, decisions, and
  constraints for one deliverable. Select when the user wants to define or
  capture requirements, scope, constraints, or intent for a deliverable;
  consolidate approved ADR decisions into a single spec; refine an existing
  spec; or needs to clarify what they are asking for through Socratic dialogue
  before committing to build.
compatibility: opencode
metadata:
  agent: cocrates
---

# Spec Writing — Requirement & Decision Consolidation Skill

A Spec here is a **short, living document for one deliverable or requirement scope** — what the user asked for, what they decided, and what generation must respect. It is **not a design white paper** and must not restate long explanations the user did not request.

**Spec vs ADR:**

| | ADR | Spec |
|---|-----|------|
| **Unit** | One concern (one question) | One deliverable or requirement scope |
| **Purpose** | Compare alternatives; record why a choice was made | State what was decided — self-contained for generation |
| **Reader needs** | Understand options, tradeoffs, and selection rationale | Understand the agreed specification — **no ADR required** |
| **Role in pipeline** | Decision audit trail | **Sole input** for generation and verification |

One Spec file = one requirement scope.

**Self-containment rule:** A reader (or `spec-driven-generation`) must understand **all decided requirements and constraints** from the Spec alone. ADRs explain *how* a decision was reached — which alternatives existed and why one was chosen — but are **not** a dependency for building. When consolidating from approved ADRs, **copy the decided outcome into the Spec in full**; do not substitute ADR links for stated requirements.

Match the user's language (Korean, English, etc.) in user-facing messages. Spec file content may use the language of the requirement context; filenames and slugs stay English kebab-case.

## Storage Location

```
spec/
└── {requirement-slug}.md
```

- **Spec root:** `spec/` at the project root (create if missing)
- **Filename:** English **kebab-case** for the requirement scope (e.g. `item-catalog-api.md`, `q3-research-report.md`)
- One file = one deliverable or coherent requirement scope

### Before Creating a New Spec

Before saving, check the `spec/` directory.

- Search for files with the **same or similar scope** (compare filename, `## Requirement`, `## Tags`)
- If found, **do not create a new file** — **supplement and merge** into the existing Spec, or when the scope is reframed merge into a new `{requirement-slug}.md` and note the prior file under **Related**
- If not found, create a new `{requirement-slug}.md`

**Similar-scope examples:**
- `item-api.md` ↔ `item-catalog-api.md` → same scope, merge into one
- `report-v1.md` replaced by expanded scope → supersede with link to `q3-research-report.md`

---

## Workflow

### 1. Identify the Requirement Scope

Clarify **what deliverable or requirement scope** this Spec covers and why a Spec is needed now.

Capture briefly:

- **Requirement** — one clear statement of what the user wants produced or achieved (e.g. *"REST API for item catalog CRUD"*, *"Q3 AI education research report for executive audience"*)
- **Trigger** — new request, adr-writing handoff, or refinement of an existing Spec?
- **Stakes** — what downstream work depends on this Spec?

If multiple independent deliverables exist, create **separate Spec files** — one scope per file.

### 2. Resolve Ambiguity (Socratic)

Before drafting, surface gaps the user has not yet decided.

**Ask Socratic questions when:**

- The request mixes multiple deliverables or goals
- Scope, audience, format, quality bar, or constraints are unstated
- Terminology is vague (*"simple"*, *"modern"*, *"good enough"*)
- A choice was implied but never explicitly confirmed

**Dialogue rules:**

- One question at a time when possible; group related questions only when the user prefers speed
- Prefer questions that help the user discover their own answer over presenting a default
- Do not invent requirements — document only what the user states or confirms
- When the user defers a decision, record it under **Open Questions** and ask whether generation should wait

**Do not** fill the Spec with agent assumptions presented as facts.

### 3. Gather Inputs

Collect everything that belongs in the Spec:

| Source | What to transfer |
|--------|------------------|
| **User request** | Stated goals, audience, format, constraints, explicit non-goals |
| **Approved ADRs** | **The decided outcome only** — stated plainly in Spec bullets (what to build, not which option letter was chosen). One-line rationale optional. Do **not** rely on ADR links as substitutes for spec content. |
| **Prior Spec** | Still-valid requirements when refining an existing file |
| **Artifact-specific skills** | Structural constraints from skills like `document-authoring` when they apply |

From ADRs, transfer **decided specifications** — not option lists, tradeoff tables, or pointers that force the reader to open `adr/`. If a decision affects generation, it must appear as an explicit requirement or constraint in the Spec.

### 4. Write or Update the Spec (Concise)

Write or revise using the template below.

**When the user edited the file directly:** read the current file first; preserve their edits; merge agent updates without overwriting unstated user changes.

**Writing rules:**

- **Bullets over prose** — one bullet = one requirement, decision, or constraint
- **User's words** when they stated something clearly; paraphrase only for clarity
- **Self-contained** — every decided item must be readable without opening `adr/`; generation must not need ADR files
- **No lengthy explanation** — if background is needed, one line under Context; link to reference materials
- **Testable statements** — each requirement should be verifiable during `spec-driven-verification`
- **Explicit out-of-scope** — record what the user said they do *not* want

**Bad (depends on ADR):**
> - Auth: Option B (see `adr/api-authentication.md`)
> - Storage: per approved database ADR

**Bad (too verbose):**
> The system should provide a comprehensive RESTful API architecture that leverages modern best practices for scalability and maintainability, ensuring that all endpoints follow consistent naming conventions and that the implementation will support future extensibility…

**Good (concise, self-contained bullets):**

> - REST API for item catalog: list, get, create, update, delete
> - Auth: API key per client; keys issued out of band; invalid key → 401
> - Storage: PostgreSQL; core fields relational; variable item attributes in JSONB columns
> - Not in scope: admin UI, bulk import, multi-region deployment

Save to `spec/{requirement-slug}.md`.

### 5. Notify User

After saving or updating, report the **file path** and a brief summary of what was captured.

Invite correction — the user may edit the file directly or ask for revisions:

- *"`spec/item-catalog-api.md` is updated. Edit the file directly or tell me what to change."*
- *"Open question: bulk import — defer or add to scope?"*

**Do not:**

- Add `Status`, `Approved`, or similar approval fields to the Spec
- Add requirements the user has not stated or confirmed
- Overwrite user edits in the file without reading the current content first

---

## File Template

```markdown
# {Requirement Title}

## Requirement
{One clear statement of what the user wants — deliverable, audience, success in one or two sentences.}

## Context
- {Why this is needed now — one line each}
- {Relevant constraints or prior work}

## Decisions
- {What was decided — stated fully, e.g. PostgreSQL with JSONB for variable attributes}
- {Optional one-line rationale — why this choice, in the user's words when possible}

## Requirements
- {Must-have: testable bullet}
- {Must-have: testable bullet}

## Constraints
- {Technical, legal, time, format, or quality constraints the user confirmed}

## Out of Scope
- {Explicit non-goals the user confirmed}

## Open Questions
- {Deferred item — and whether it blocks generation}

## Related
- adr/{concern-slug}.md — optional; audit trail for how a decision was reached, not required for generation
- kb/{topic-slug}.md
- spec/{other-spec}.md

## Tags
`tag-one`, `tag-two`
```

Omit empty sections except **Requirement** and **Requirements** (at least one confirmed item).

Do **not** include `Status`, `Approved`, or other approval-state sections.

---

## Example (abbreviated)

File: `spec/item-catalog-api.md`

```markdown
# Item Catalog API

## Requirement
REST API for internal teams to manage item catalog records (CRUD).

## Context
- Replacing manual spreadsheet updates
- Single-region deployment; moderate volume

## Decisions
- **PostgreSQL** for storage — core fields relational; variable item attributes in JSONB columns
- **API key** authentication — one key per client; invalid or missing key returns 401

## Requirements
- Endpoints: list items, get by id, create, update, delete
- Item fields: id, name, category (required); attributes (optional, schema varies by category)
- JSON request/response; errors return consistent `{ code, message }` shape
- OpenAPI 3 document shipped alongside implementation

## Constraints
- Team stack: TypeScript, existing Express monorepo
- No new infrastructure beyond current PostgreSQL instance

## Out of Scope
- Admin UI
- Bulk import/export
- Multi-region or read replicas

## Related
- adr/item-storage-database.md
- adr/api-authentication.md

## Tags
`api`, `item-catalog`, `rest`
```

---

## Dialogue Rules

1. **State the current step** (identify scope → resolve ambiguity → gather inputs → write or update → notify → generation handoff when requested).
2. After saving, report the **file path** and what was captured or changed.
3. Ask Socratic questions **before** drafting when anything material is unclear.
4. When entering from `adr-writing`, confirm the approved ADRs match the user's current intent before consolidating.

---

## Prohibitions

- Long narrative explanations instead of concise, user-confirmed bullets
- Requirements or decisions the user has not stated or confirmed
- Spec bullets that point to `adr/` instead of stating the decided specification
- Duplicating full ADR option analysis inside the Spec
- `Status`, `Approved`, or other approval-state fields in Spec files
- Duplicating the same scope across multiple Spec files
- Handing off to generation with ADR files as required reading

---

## Completion Criteria

- Requirement scope identified and stated clearly
- Material ambiguities resolved through dialogue or recorded as **Open Questions** with user awareness
- Spec saved or updated at `spec/{requirement-slug}.md` with user-confirmed requirements and decisions
- Spec is **self-contained** — decided specification understandable without reading `adr/` files
- User informed of the file path; may edit directly or request revisions
- When generation is next, user informed of **`spec-driven-generation`** or the relevant artifact-specific generation skill

After handoff, generation skills own production; this skill does not generate the final artifact.
