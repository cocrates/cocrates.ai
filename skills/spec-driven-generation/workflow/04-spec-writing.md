# Step ④ — Spec Writing

## Purpose

Consolidate all approved decisions — user requests, Direct Input outcomes, and ADR conclusions — into **self-contained Spec documents**. The Spec becomes the **sole input** for generation (⑤) and verification (⑥). A reader must understand all decided requirements from the Spec alone, without consulting ADR files or conversation history.

## Prerequisites

- Step ③ complete: ASRs reviewed, Resolutions filled.
- Approved ADRs available in `{project-root}/adr/` (if any).

## Procedure

### 1. Identify Requirement Scope & Hierarchy

Determine which Spec file(s) to create or update:

- **Initial Setup:** If no Spec exists yet, create `{project-root}/spec/{requirement-slug}.md` mapped to the PRD's primary deliverable.
- **Feature/Component Focus:** If the PRD covers multiple independent deliverables, create one `{requirement-slug}.md` per independent scope.
- **Check existing:** Before creating, search `{project-root}/spec/` for files with the **same or similar scope** (compare filename, `## Requirement`, `## Tags`). If found, **supplement and merge** into the existing Spec.

### 2. Resolve Ambiguity (Socratic)

Before drafting, surface gaps using Socratic dialogue:
- Ask questions when scope, target audience, technical format, or quality bars are unstated or vague.
- Do not fill the Spec with agent assumptions presented as facts.
- Propose recommended defaults where possible.

### 3. Gather & Deep-Copy Inputs

Collect everything that belongs in the Spec:

- **User Requests & Direct Inputs:** Goals, explicit non-goals, and constraints.
- **Approved ADRs (Deep-Copy):** Transfer **the decided outcome only**, stated plainly in detailed Spec bullets. Unpack the technical conclusions of the ADR.
  - *Example:* If an ADR approves Redis Cache, pull downstream decisions (e.g., TTL = 60s, LRU Eviction) directly into the Spec's `## Decisions` section.
- **Verification Feedback Loops:** If triggered by a ⑥ failure, extract **Deviations** or **Undocumented ASRs** from the verification report.

### 4. Write or Update the Spec

Use the template below. Key rules:

- **Self-Containment:** All decided outcomes must be fully stated in the Spec. Never use ADR hyperlinks as substitutes.
- **Incremental Edits:** When updating due to verification defects, target only the affected sections. Preserve existing user edits.
- **Formatting:** Use **bullets over prose**. Every line under `## Requirements` must be a **testable statement** for verification.
- **Exclusions:** Explicitly document `## Out of Scope` items.

Save to `{project-root}/spec/{requirement-slug}.md` or update `{project-root}/spec/PRD.md`.

### 5. Update ASR Registry

After saving the Spec and obtaining user confirmation:
- For each resolved ASR: set status to `approved` in `ASR.md`.
- Fill **Spec** field with the file path and section.
- Keep **Summary** table in sync.

### 6. Notify and Route

Report the **file path** and a brief summary:
- *"I have updated `{project-root}/spec/item-catalog-api.md` with the finalized cache requirements. The spec is now completely self-contained. ASR-003 is marked approved in `ASR.md`."*

**Next → Step ⑤ (Generation).**

## Gate

- Spec file(s) saved in `{project-root}/spec/` ✓
- All content deep-copied from ADRs (no ADR hyperlinks as substitutes) ✓
- Related ASRs marked `approved` in `ASR.md` ✓
- User confirmed the Spec ✓

---

## Spec File Template

```markdown
# {Requirement Title}

## Requirement
{One clear statement of what the user wants — deliverable, audience, success in one or two sentences.}

## Context
- {Why this is needed now — one line each}
- {Link to parent {project-root}/spec/PRD.md if applicable}

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
- {project-root}/adr/{concern-slug}.md — optional audit trail only; NOT required reading for generation
- {project-root}/spec/PRD.md

## Tags
`tag-one`, `tag-two`
```

Omit empty sections except **Requirement** and **Requirements**. Do **not** include Status or Approved fields.
