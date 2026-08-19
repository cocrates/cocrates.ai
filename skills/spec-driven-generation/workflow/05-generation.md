# Step ⑤ — Generation

## Purpose

Produce the final artifact **strictly from the Spec**. The Spec is the sole source of truth — not the original user prompt, not ADR files, and not assumptions made during generation.

## Prerequisites

- Step ④ complete: Spec(s) saved and user-confirmed in `{project-root}/spec/`.
- All blocking ASRs are `approved` in `ASR.md`.

## Procedure

### 1. Load Spec(s)

Read the following files and keep in working memory:
- `{project-root}/spec/PRD.md` — high-level scope, goals, constraints.
- `{project-root}/spec/{requirement-slug}.md` — detailed Requirements, Decisions, Constraints, Out of Scope.

**Do NOT load:**
- `{project-root}/spec/ASR.md` — decision-status registry, not a generation input.
- `{project-root}/adr/` files — decision audit trail, not required reading.

### 2. Build Working Checklist

Extract from the loaded Specs:
- **Requirement:** What to produce.
- **Decisions:** How it is organized.
- **Requirements:** Testable must-haves.
- **Constraints:** Hard limits.
- **Out of Scope:** What must NOT appear.

Spec overrides conversation history. If the latest user message contradicts the Spec, ask which source to follow.

### 3. Generate

Produce the artifact **strictly from the Spec**:
- Every **Requirements** bullet must be addressed.
- Honor all **Decisions** and **Constraints**.
- Respect **Out of Scope** — no silent extras.
- Follow project conventions when the Spec is silent.
- Output paths must match the Spec's stated location.

### 4. Report and Hand Off

After generation:
- Output the file path(s) and a brief **Requirements checklist** mapping each Spec bullet to the generated output.
- Recommend **Step ⑥ (Verification)** → `{project-root}/verification/{requirement-slug}.md`.

## Gate

- Artifact produced at Spec-stated paths ✓
- Every Requirements bullet addressed (checklist provided) ✓
- No Out of Scope violations ✓
- User informed of outputs and offered verification ✓

---

## Generation Rules

| Rule | Detail |
|------|--------|
| **Spec is authoritative** | If Spec and conversation conflict, ask user which to follow |
| **No silent additions** | Never add features, scope, or requirements not in the Spec |
| **No ADR loading** | ADR files are never read during generation |
| **No patching** | If the artifact needs changes, update the Spec first, then regenerate |
| **Convention follows Spec** | When the Spec is silent on a detail, follow reasonable project conventions |
