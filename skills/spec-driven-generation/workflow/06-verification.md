# Step ⑥ — Verification

## Purpose

Check generated artifacts against the Spec — item by item, with evidence. Also discover **undocumented ASRs** (Architecturally Significant Requirements) embodied in the deliverable that the Spec never addressed. This step triggers the **reverse traceability loop** back into the pipeline.

## Prerequisites

- Step ⑤ complete: artifact generated.
- `{project-root}/spec/PRD.md` and `{project-root}/spec/{requirement-slug}.md` available.

## Procedure

### 1. Determine Verification Scope

Clarify what to verify and which Specs govern it:

| Request pattern | Action |
|-----------------|--------|
| Named file or slug | Verify against `PRD.md` AND that specific Spec |
| "All specs", "everything" | Verify against all `{project-root}/spec/*.md` including `PRD.md` |
| Unspecified | Ask: which project and single deliverable vs multiple — then use the matching set |

### 2. Load Spec(s) & Extract Checklist

Read `{project-root}/spec/PRD.md` and the target `{project-root}/spec/*.md`. Extract every verifiable item:
- **PRD Level:** High-level goals, target audience, core deliverable form.
- **Spec Level:** Explicit **Decisions**, **Requirements** (must-haves), **Constraints**, **Out of Scope** items.

### 3. Load Artifact(s)

Read the full generated deliverable(s) at the identified paths. Do not skim. Verification requires reading enough to judge compliance and scan for unstated ASRs.

### 4. Build the Verification Inventory

List every Spec item used for this verification explicitly before checking them. One Spec bullet = one inventory row.

### 5. Verify Item by Item

Check each inventory item individually. Record:

| Field | Content |
|-------|---------|
| **Spec item** | The exact Spec bullet or PRD goal |
| **Status** | `pass` \| `fail` \| `partial` \| `not-verifiable` |
| **Evidence** | Path, section, line number in the artifact |
| **Notes** | Explanation of status |

- **Out of Scope** items: Presence of out-of-scope material must be marked as a **fail**.

### 6. Find Spec Deviations

Consolidate all items with status `fail` or `partial`. Classify severity:

| Severity | Definition |
|----------|-----------|
| **Critical** | Violates a must-have PRD/Spec **Requirement** or **Decision** |
| **Major** | Violates a **Constraint** or partially meets a must-have |
| **Minor** | Cosmetic or non-functional difference not explicitly forbidden |

### 7. Find Undocumented ASRs

Scan the artifact for structural, architectural, or qualitative choices made during production that bypassed the specification phase. Use the **ASR gap scan protocol**:

| Field | Content |
|-------|---------|
| **What was decided** | The silent choice the generation made |
| **Where** | Location in the artifact |
| **Category** | From universal ASR categories (Purpose & audience, Deliverable form, Scope boundary, Quality bar, Constraints, Structure & organization, Integration & dependencies) |
| **Gap** | What the Spec should have stated but didn't |
| **Risk** | What could go wrong if this default is wrong |
| **Recommendation** | Whether to accept, update Spec, or open an ADR |

**Filter by Significance:** Only keep findings that are truly architecturally significant — where a wrong silent default would degrade project quality. Minor implementation details should be ignored.

### 8. Write Verification Report

Save the full report to **`{project-root}/verification/{requirement-slug}.md`** (create directory if missing). If the file exists, **read it first** to preserve any previous User Review notes.

### 9. User Review and Loop Activation

After writing the report, present it and invite user review. Based on the user's response, execute the **Reverse Traceability Loop**:

```
[Verification Report] ──> [User Review] ──> [Dynamic Routing]
                                                   │
     ┌─────────────────────────────────────────────┴─────────────────────────────┐
     ▼                                              ▼                            ▼
 [Fix Artifact]                              [Update Spec]                  [Run ADR Analysis]
  - Re-run Step ⑤                            - Update Spec via             - Register ASR in ASR.md
  - Targeted patch or rebuild                  Step ④ procedure            - Open ADR via adr-writing
     │                                              │                            │
     └──────────────────────────────────────────────┴────────────────────────────┘
                                                   │
                                                   ▼
                                      [Re-generate (⑤) & Re-verify (⑥)]
```

| User decision | Action |
|---------------|--------|
| Fix deviations | Re-run Step ⑤ scoped strictly to failed items, then re-verify |
| Accept a deviation or confirm undocumented ASR | Register/reopen ASR in `ASR.md`, update Spec via Step ④, mark ASR `approved`, then re-generate |
| Undocumented ASR involves complex trade-offs | Register in `ASR.md` as `reviewing`, open ADR via `adr-writing`, approve, update Spec, re-generate |

## Gate

- Verification inventory complete ✓
- Every item individually checked with status and evidence ✓
- Deviations and Undocumented ASRs categorized and recorded ✓
- Report saved to `{project-root}/verification/{requirement-slug}.md` ✓
- User reviewed and routing executed ✓

---

## Verification Report Template

```markdown
# Verification: {Requirement Title}

**Spec:** {project-root}/spec/{requirement-slug}.md (Aligned with {project-root}/spec/PRD.md)
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
{each unstated ASR — decision made, location, category, gap, risk, recommendation}

## Recommended Next Steps
- Follow the loop: Update Spec/PRD → Re-approve → Re-generate.

## User Review
{Leave empty for user notes}
```
