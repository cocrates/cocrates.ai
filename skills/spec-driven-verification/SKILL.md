---
name: spec-driven-verification
description: >-
  Verifies generated deliverables against Spec documents in
  {project-slug}/spec/ (including {project-slug}/spec/PRD.md). Select when
  the user asks to verify, validate, check, or confirm that an artifact
  matches the specification; after spec-driven-generation completes; or when
  reviewing whether output was produced according to the spec. Performs
  item-by-item comparison, surfaces spec deviations, and flags undocumented
  ASRs (Architecturally Significant Requirements) to drive the iterative
  specification and refinement loop.
compatibility: opencode
metadata:
  agent: cocrates
---

# Spec-Driven Verification — Artifact Compliance Skill

This skill **checks generated artifacts against Spec documents**. The Spec is the **sole source of truth** for what should have been built — not the original user prompt, not ADR files, and not assumptions made during generation.

Deliverables include **software, documents, images, presentations, blog series, reports**, and any other finished work the Spec describes.

## Project Folder Layout

All Spec / ADR / verification work for a deliverable lives under a **project folder** at the repository root:

```
{project-slug}/
├── spec/
│   ├── ASR.md
│   ├── PRD.md
│   └── {requirement-slug}.md
├── adr/
│   └── {concern-slug}.md
├── verification/
│   └── {requirement-slug}.md
└── … (generated artifacts and other project files)
```

- **`{project-slug}`:** English **kebab-case** name for the project. Resolve from the user's request or an existing project folder; ask if ambiguous.
- **`{project-slug}/spec/ASR.md`:** When undocumented ASRs are confirmed, register them here (`identified` / `reviewing`) before routing to Spec or ADR work.
- **Do not** place `spec/`, `adr/`, or `verification/` at the repository root — always nest them under `{project-slug}/`. Keep `ASR.md` only at `{project-slug}/spec/ASR.md`.

## Core Principle: Verification Beyond Compliance

Verification is not only checking whether the artifact matches explicit Spec bullets. It also discovers **undocumented ASRs (Architecturally Significant Requirements)** embodied in the deliverable — retrospective findings of *"we should have decided this before building."*

An **ASR** is a requirement or design decision that materially affects the structure, organization, or quality of the final deliverable. Unstated ASRs become visible only when reading the finished artifact. Surfacing undocumented ASRs is a primary output of this skill, alongside pass/fail against the Spec.

### Workflow Nature: Non-linear Feedback Loop

This skill does not simply output a static report; it **triggers the reverse optimization loop** of the development lifecycle:
- **Defect/Gap Traceability:** Deviations and Undocumented ASRs found here must be fed back into the design phase. 
- **Continuous Refinement:** The user reviews the report, the agent updates the Spec (and runs ADR if complex alternatives emerge), and then re-triggers `spec-driven-generation` to refine the artifact.

**Spec vs ADR vs verification:**

| | ADR | Spec | Verification (this skill) |
|---|-----|------|---------------------------|
| **Unit** | One concern (one question) | One deliverable or requirement scope | One or more artifacts per Spec |
| **Purpose** | Compare alternatives; record approval | State decided requirements self-contained | **Confirm** the artifact matches the Spec |
| **Reader needs** | Understand options and tradeoffs | Understand what to build | Know what passed, failed, or was decided during implementation |
| **Role in pipeline** | Decision audit trail | **Sole criterion** for verification | Audit |

Match the user's language (Korean, English, etc.) in user-facing messages.

---

## Spec Input Scope

Verification reads from **`{project-slug}/spec/`**.

```
{project-slug}/spec/
├── ASR.md
├── PRD.md
└── {requirement-slug}.md
```

Verification criteria come from `PRD.md` and `{requirement-slug}.md`. Use `ASR.md` only as the decision-status registry when registering undocumented ASRs — **do not** treat it as a verification Spec.
### Which Spec(s) to Use

Determine scope from the user's request:

| Request pattern | Action |
|-----------------|--------|
| Named file or slug (e.g. `{project-slug}/spec/item-catalog-api.md`) | Verify against **`{project-slug}/spec/PRD.md` AND that specific Spec** |
| "All specs", "everything in {project-slug}/spec/" | Verify against **all** `{project-slug}/spec/*.md` files including `PRD.md` |
| Unspecified | Ask one Socratic question: which `{project-slug}` and single deliverable vs multiple — then use the matching set |

**Do not read `{project-slug}/adr/` files for verification.** If a Spec bullet references ADR content without stating the decided requirement, flag it as a **Spec gap** and recommend **`spec-writing`** to make the Spec self-contained before re-verifying.

---

## Output Location

Write every verification report to **`{project-slug}/verification/`** (create the directory if missing).

```
{project-slug}/verification/
└── {requirement-slug}.md
```

- **Filename:** Same **kebab-case slug** as the governing Spec (e.g. `{project-slug}/spec/item-catalog-api.md` → `{project-slug}/verification/item-catalog-api.md`).
- **Before Writing:** If the file exists, **read it first** to preserve any **User Review** notes or previous manual decisions.

---

## ASR — Architecturally Significant Requirements

This section provides a **retrospective framework** for finding ASRs embodied in the artifact that the Spec never addressed. Use in step 7 after item-by-item Spec verification[cite: 1, 2].

### Universal ASR categories (always scan)

| ASR category | What to look for in the artifact | Examples across artifact types |
|--------------|----------------------------------|--------------------------------|
| **Purpose & audience** | Implicit target reader or user assumed | Technical depth chosen for experts vs beginners; formality level |
| **Deliverable form** | Format, layout, or medium choices not in Spec | File structure; export format; slide vs document |
| **Scope boundary** | Material included or excluded without Spec guidance | Extra sections; features beyond stated scope; missing promised depth |
| **Quality bar** | Standards applied but not specified | Evidence rigor; visual polish; completeness threshold |
| **Constraints** | Limits respected or violated silently | Language; brand; length; deadline-driven shortcuts |
| **Structure & organization** | Arrangement chosen without Spec guidance | Section order; narrative arc; functional grouping |
| **Integration & dependencies** | Connections to other work assumed or invented | References to prior episodes; reuse of templates |

### ASR gap scan protocol

After item-by-item Spec verification (step 5):
1. **Filter by Significance:** Scan the artifact to identify embodied decisions the Spec never addressed. **Keep only findings that are truly architecturally significant** (where a wrong silent default would degrade project quality). Minor implementation details should be ignored.
2. **Record each finding** with: what was decided, where in the artifact, ASR category, spec gap, risk, and recommended action.
3. **Present to the user** for confirmation or correction — do not treat silent choices as approved.

---

## Workflow

### 1. Determine Verification Scope
Clarify what to verify, which `{project-slug}` applies, and which Specs govern it. Always pull `{project-slug}/spec/PRD.md` alongside the target spec to ensure high-level alignment.

### 2. Load Spec(s) & Extract Checklist
Read `{project-slug}/spec/PRD.md` and the target `{project-slug}/spec/*.md`. Extract every verifiable item into a working checklist:
- **PRD Level:** High-level goals, target audience, and core deliverable form.
- **Spec Level:** Explicit **Decisions**, **Requirements** (must-haves), **Constraints**, and **Out of Scope** items.

### 3. Load Artifact(s)
Read the full generated deliverable(s) at the identified paths. Do not skim. Verification requires reading enough to judge compliance and scan for unstated ASRs.

### 4. Build the Verification Inventory
List every Spec item used for this verification explicitly before checking them. One Spec bullet = one inventory row.

### 5. Verify Item by Item
Check each inventory item individually. Record: **Spec item**, **Status** (`pass` | `fail` | `partial` | `not-verifiable`), **Evidence** (path, section, line), and **Notes**.
- **Out of Scope** items: Presence of out-of-scope material must be marked as a **fail**.

### 6. Find Spec Deviations
Consolidate all items with status `fail` or `partial`. Classify severity:
- **Critical** — violates a must-have PRD/Spec **Requirement** or **Decision**.
- **Major** — violates a **Constraint** or partially meets a must-have.
- **Minor** — cosmetic or non-functional difference not explicitly forbidden.

### 7. Find Undocumented ASRs
Examine the artifact for structural, architectural, or qualitative choices made during production that bypassed the specification phase. Record them using the *ASR gap scan protocol* (What was decided, Where, Category, Gap, Risk, Recommendation).

### 8. Write Verification Report
Save the full report to **`{project-slug}/verification/{requirement-slug}.md`** using the standardized template.

```markdown
# Verification: {Requirement Title}

**Spec:** {project-slug}/spec/{requirement-slug}.md (Aligned with {project-slug}/spec/PRD.md)
**Artifact(s):** {paths}
**Verified:** {ISO date}
**Summary:** {N} pass, {N} fail, {N} partial, {N} not-verifiable

## Inventory & Results
| # | Spec item | Status | Evidence / Notes |
|---|-----------|--------|------------------|
| 1 | [PRD] High-level Goal | pass | ... |
| 2 | [Spec] Requirement A | fail | Missing implementation |

## Deviations (Non-compliance)
{each fail/partial — Spec item, actual artifact behavior, severity}

## Undocumented ASRs (Specification Gaps)
{each unstated ASR — decision made, location, category, risk, recommended action}

## Recommended Next Steps
- Follow the loop: Update Spec/PRD -> Re-approve -> Re-generate.

## User Review
{Leave empty for user notes}
```

### 9. User Review and Loop Activation

After writing the report, invite the user to review the file. When the user responds, execute the following Reverse Traceability Loop:

```
[Verification Report File] ──> [User Review] ──> [Dynamic Routing]
                                                       │
         ┌─────────────────────────────────────────────┴─────────────────────────────────────────────┐
         ▼                                             ▼                                             ▼
 [Fix Artifact]                                 [Update Spec]                                 [Run ADR Analysis]
  - Run 'spec-driven-generation'                 - Load 'spec-writing'                         - Load 'adr-writing' for complex gaps
  - Targeted patch or rebuild                    - Sync Deviations/ASRs to Spec                - Establish decisions, then update Spec
         │                                             │                                             │
         └─────────────────────────────────────────────┴─────────────────────────────────────────────┘
                                                       │
                                                       ▼
                                          [Re-generate & Re-verify]
```

- If the user wants to fix Deviations: Trigger `spec-driven-generation` scoped strictly to the failed items.
- If the user accepts a Deviation or confirms an Undocumented ASR: Register (or reopen) the ASR in `{project-slug}/spec/ASR.md`, then load `spec-writing` to update `{project-slug}/spec/PRD.md` or `{project-slug}/spec/*.md` and mark the ASR `approved` once encoded.
- If an Undocumented ASR involves complex trade-offs: Register it in `ASR.md` as `reviewing`, load `adr-writing` with Related ASR ID(s), approve a choice, update the Spec, and then trigger regeneration.

---

## Prohibitions

- Verifying against conversation history or ADR files instead of the Spec/PRD.
- Writing verification reports at the repository root instead of under `{project-slug}/verification/`.
- Cursory or summary-only verification without item-by-item artifact evidence.
- Ignoring Out of Scope violations or treating undocumented ASRs as automatically approved.
- Applying artifact fixes or patches silently before the user has reviewed the verification report and confirmed the next architectural step.
- Overwriting user edits in the User Review section.

---

## Completion Criteria

- Verification inventory mapped against both `{project-slug}/spec/PRD.md` and detailed `{project-slug}/spec/*.md`.
- Every item individually checked with explicit status and artifact evidence.
- Deviations and Undocumented ASRs clearly categorized and recorded.
- Report saved to `{project-slug}/verification/{requirement-slug}.md`.
- Dynamic routing executed based on user feedback (Artifact fixed, Spec updated, or ADR analysis initiated) to ensure a complete, closed-loop lifecycle.