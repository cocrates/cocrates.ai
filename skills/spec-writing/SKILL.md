---
name: spec-writing
description: >-
  Creates and updates Spec documents — consolidating requests, decisions, and
  constraints for one deliverable or requirement scope. Select when the user
  wants to define or capture requirements, scope, constraints, or intent;
  consolidate approved ADR decisions into a single self-contained spec;
  establish a high-level {project-slug}/spec/PRD.md; or refine an existing
  spec based on post-generation verification feedback.
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

### Core Architectural Rules

1. **One File, One Scope Principle:** Every distinct component, feature, or independent deliverable must have its own dedicated Spec file to prevent document bloat and maintain modularity.
2. **Strict Self-Containment Rule:** A reader (or `spec-driven-generation`) must understand **all decided requirements and constraints** from the Spec alone. ADRs explain *how* a decision was reached, but are **not** a dependency for building. 
   - When consolidating from approved ADRs, **copy the decided outcome's exact parameters, technical values, and behaviors into the Spec in full**.
   - **Prohibited:** Never use ADR hyperlinks as a substitute for actual specification bullets (e.g., Do not write `- Auth: See {project-slug}/adr/auth.md`).

Match the user's language (Korean, English, etc.) in user-facing messages. Spec file content may use the language of the requirement context; filenames and slugs stay English kebab-case.

---

## Project Folder Layout

Specs are stored under the project's **`spec/`** directory:

```
{project-slug}/
├── spec/
│   ├── ASR.md
│   ├── PRD.md
│   └── {requirement-slug}.md
├── adr/
├── verification/
└── …
```

- **`{project-slug}`:** English **kebab-case** project folder (same as used by `spec-driven-generation` / `adr-writing`). Resolve from the user's request or an existing project folder; ask if ambiguous. If the folder does not exist yet, prefer routing through **`spec-driven-generation`** (which owns project-folder creation) or create `{project-slug}/spec/` when the user explicitly asks to write a Spec first.
- **`{project-slug}/spec/ASR.md`:** When encoding approved decisions, advance related ASRs to `approved` and record the Spec path under each ASR's **Spec** field.
- **`{project-slug}/spec/PRD.md` (Product Requirement Definition):** The top-level master anchor that defines the high-level product boundaries, overarching goals, and global constraints.
- **`{project-slug}/spec/{requirement-slug}.md`:** Modular, independent specifications for individual features or deliverables mapped to the PRD. Do not confuse with `ASR.md` (status registry) or `PRD.md` (product anchor).
- **Do not** place `spec/` at the repository root — always nest under `{project-slug}/`. Keep `ASR.md` only at `{project-slug}/spec/ASR.md`.

### Before Creating a New Spec

Before saving, check the `{project-slug}/spec/` directory.
- Search for files with the **same or similar scope** (compare filename, `## Requirement`, `## Tags`). Ignore `ASR.md` and `PRD.md` when matching feature Specs.
- If found, **do not create a new file** — **supplement and merge** into the existing Spec.
- For large projects, map out whether the request should expand an existing independent spec or requires a brand-new scoped spec file.

---

## Workflow

### 1. Identify the Requirement Scope & Hierarchy

Clarify **what deliverable or requirement scope** this Spec covers.
- **Initial Setup:** If this is the start of a project, create **`{project-slug}/spec/PRD.md`** first to capture the macro scope (ensure `{project-slug}/` exists).
- **Feature/Component Focus:** If the PRD is established, map the current requirement to a specific `{requirement-slug}.md` file under `{project-slug}/spec/`. Ensure it isolates a single, independent concern.

### 2. Resolve Ambiguity (Socratic)

Before drafting, surface gaps using Socratic dialogue. Ask questions when scope, target audience, technical format, or quality bars are unstated or vague. Do not fill the Spec with agent assumptions presented as facts.

### 3. Gather & Deep-Copy Inputs

Collect everything that belongs in the Spec:
- **User Requests & Direct Inputs:** Goals, explicit non-goals, and constraints.
- **Approved ADRs (Deep-Copy):** Transfer **the decided outcome only**, stated plainly in detailed Spec bullets. You must explicitly unpack the technical conclusions of the ADR. 
  - *Example:* If an ADR approves Redis Cache, pull the downstream decisions (e.g., TTL = 60s, LRU Eviction) directly into the Spec's `## Decisions` section.
- **Verification Feedback Loops:** If triggered by a `spec-driven-verification` failure, extract the identified **Deviations** or **Undocumented ASRs** from the verification report.

### 4. Write or Update the Spec (Concise Bullets)

Write or revise using the template below.

- **Incremental Edits:** When updating due to verification defects, target only the affected sections. Preserve existing user edits and never overwrite unstated manual changes.
- **Formatting:** Use **bullets over prose**. Ensure every line under `## Requirements` is a **testable statement** that can be checked deterministically during verification.
- **Exclusions:** Explicitly document `## Out of Scope` items to prevent feature creep.

Save to `{project-slug}/spec/{requirement-slug}.md` or `{project-slug}/spec/PRD.md`.
After saving, update `{project-slug}/spec/ASR.md` for each resolved ASR: set status to `approved` (when the user confirms), fill **Resolution** / **Spec** fields, and keep Summary in sync.

### 5. Notify and Route

Report the **file path** and a brief summary of what was captured or modified.
- *Example:* *"I have updated `{project-slug}/spec/item-catalog-api.md` with the finalized cache requirements. The spec is now completely self-contained. ASR-003 is marked approved in `ASR.md`."*
- **Next Step Routing:** If the spec is ready and sufficient, guide the user to **`spec-driven-generation`** to build or refine the artifact.

---

## File Template

```markdown
# {Requirement Title}

## Requirement
{One clear statement of what the user wants — deliverable, audience, success in one or two sentences.}

## Context
- {Why this is needed now — one line each}
- {Link to parent {project-slug}/spec/PRD.md if applicable}

## Decisions
- {What was decided — stated fully with deep-copied concrete parameters}
- {e.g., Redis Cache layer with a 60-second TTL and LRU eviction policy}

## Requirements
- {Must-have: testable bullet for verification}
- {Must-have: testable bullet for verification}

## Constraints
- {Technical, legal, time, format, or quality constraints}

## Out of Scope
- {Explicit non-goals confirmed by the user to prevent creep}

## Open Questions
- {Deferred item — and whether it blocks generation}

## Related
- {project-slug}/adr/{concern-slug}.md — optional audit trail only; NOT required reading for generation
- {project-slug}/spec/PRD.md

## Tags
`tag-one`, `tag-two`
```

Omit empty sections except **Requirement** and **Requirements**. Do **not** include Status or Approved fields.

---

## Dialogue Rules

1. **State the Current Step**: (Identify scope → Resolve ambiguity → Gather inputs → Write/Update → Notify/Route).
2. **Enforce Clean Hand-offs**: When entering from `adr-writing`, explicitly tell the user: "Consolidating approved choices into the Spec. The Spec will act as the sole, self-contained input for generation; you won't need to check the ADR files again."
3. **Handle Feedback Loop Intelligently**: If the user arrives with a verification report, focus the dialogue strictly on deciding whether to update the Spec contract to match the deviation or to force an artifact fix.

---

## Prohibitions

- Long narrative explanations instead of concise, user-confirmed bullets.
- Spec bullets that point to `{project-slug}/adr/` or conversation history instead of stating the specification fully.
- Writing Spec files at the repository root instead of under `{project-slug}/spec/`.
- Handing off to generation with ADR files marked as required reading.
- Duplicating full alternative option analyses inside the Spec.
- Adding `Status`, `Approved`, or other lifecycle approval state fields in Spec files.
- Bundling multiple independent deliverables into a single Spec document (violating the Single Responsibility rule).

---

## Completion Criteria

- Requirement scope isolated, and `{project-slug}/spec/PRD.md` or functional `{project-slug}/spec/{requirement-slug}.md` target determined.
- All approved ADR parameters deep-copied into the Spec, achieving **100% self-containment**.
- Related ASRs in `{project-slug}/spec/ASR.md` advanced to `approved` with Spec path recorded.
- Requirements stated as **testable, itemized bullets** suitable for audit.
- Spec saved, user notified of the path, and seamlessly routed to `spec-driven-generation` for production.