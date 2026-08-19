# Step ② — ASR Identification

## Purpose

Discover every **Architecturally Significant Requirement** from the approved PRD, user context, and universal ASR categories. Register each in `{project-root}/spec/ASR.md` with status `identified` and establish the dependency order for review.

## Prerequisites

- Step ① complete: `PRD.md` approved.
- `{project-root}/spec/ASR.md` exists (initialized in Step ①).

## Procedure

### 1. Analyze the PRD

Read `{project-root}/spec/PRD.md`. For each section, ask: *What decisions must be made before generation can proceed without silent defaults?*

### 2. Scan Universal ASR Categories

Use the universal categories from SKILL.md § ASR:

| Category | Probe question |
|----------|---------------|
| Purpose & audience | Is the target reader/user clearly defined? |
| Deliverable form | Is the artifact type, format, and output path locked? |
| Scope boundary | Are in-scope and out-of-scope items explicit? |
| Quality bar | Is "done" and "good enough" defined? |
| Constraints | Are hard limits (language, platform, budget) stated? |
| Structure & organization | Is the content/function arrangement decided? |
| Integration & dependencies | Are connections to existing work specified? |
| Process & staging | Are approval gates or iterative delivery defined? |

### 3. Register ASRs in ASR.md

For each discovered ASR, add to the **ASR Detail** section in `{project-root}/spec/ASR.md`:

```markdown
### ASR-00N — {Title}

- **Category:** {from universal categories}
- **Status:** identified
- **Statement:** {What must be decided — one clear requirement/concern}
- **Why it matters:** {Impact if left to silent default}
- **Depends on:** ASR-00x (optional)
- **Related ADRs:** —
- **Resolution path:** direct-input | adr | mixed
- **Resolution:**
- **Spec:**
- **Notes:** {optional}
```

**Rules:**
- Assign stable IDs (`ASR-001`, `ASR-002`, …). Never reuse IDs.
- Update **Summary** table and **ASR Detail** together — do not leave the table stale.

### 4. Fill Dependency Order

Analyze structural dependencies between ASRs. Record the recommended review path under **Dependency Order** in `ASR.md`:

```markdown
## Dependency Order (recommended review path)

1. ASR-00x → ASR-00y → ASR-00z
{List IDs in the order structural dependencies require.}
```

Present this to the user for confirmation.

### 5. Present for Review

Show the user the full ASR list with:
- ID, Title, Category, Status
- Dependency Order
- Brief rationale for each

Wait for acknowledgment before proceeding.

## Gate

- All blocking ASRs registered in `ASR.md` (at least `identified`) ✓
- **Dependency Order** filled and presented ✓

**Next → Step ③ (ASR Review).**
