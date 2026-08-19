# Step ③ — ASR Review → ADR Creation

## Purpose

Review blocking ASRs in **Dependency Order** and resolve each — either via **Direct Input** (user chooses) or **ADR Analysis** (complex trade-offs requiring `adr-writing`). This step advances ASR statuses from `identified` → `reviewing` → `designed`.

## Prerequisites

- Step ② complete: ASRs identified in `ASR.md` with Dependency Order.
- `adr-writing` skill available for complex trade-offs.

## Procedure

### 1. Iterate ASRs in Dependency Order

For each blocking ASR (one at a time, or one dependency branch at a time):

#### 1a. Set Status to `reviewing`

Update `ASR.md`: change the ASR's **Status** to `reviewing`.

#### 1b. Route via Decision Pathways

Evaluate the ASR and choose the appropriate pathway:

**Direct Input Route (High Context):**
- The problem is straightforward, has a natural default, or the user has clear preferences.
- Ask a direct question with a recommended default (e.g., *"We recommend X because of Y. Proceed?"*).
- User chooses → fill **Resolution** in `ASR.md` with the concrete outcome.
- Set status to `designed`.

**ADR Analysis Route (Low Context / Complex Trade-offs):**
- The user is unsure, or the problem involves competing alternatives (e.g., tech stacks, structural patterns).
- Set ASR status to `reviewing`.
- Trigger **`adr-writing`** with the ASR ID(s):
  - One ASR may spawn multiple ADRs.
  - One ADR may list multiple ASR IDs (many-to-many).
  - Keep **Related ADRs** on the ASR and **Related ASRs** on the ADR bidirectional.
- Leave ASR `reviewing` until linked ADRs are approved.

#### 1c. Record Resolution

After Direct Input or ADR approval:
- Fill **Resolution** in `ASR.md` with the concrete outcome (not ADR links alone).
- Update **Summary** table.
- If an ADR was used, ensure the ADR file's **Related ASRs** field lists the ASR ID(s).

### 2. Handle Remaining ADR Approvals

When all blocking ASRs have a resolution path:
- If any Related ADR is still `proposed`: **pause for user ADR approval** (do **not** invent approval).
- When all linked ADRs are `approved` (or no open ADRs remain): confirm affected ASRs are `designed` with concrete **Resolution**.

### 3. Prepare for Spec Writing

Before advancing, verify:
- Every blocking ASR has status `designed` or `deferred`.
- **Resolution** fields contain concrete decisions (not placeholder text).
- Bidirectional ADR ↔ ASR links are consistent.

## Gate

- No blocking ASR remains `identified` (each is `reviewing` or beyond, or explicitly `deferred`) ✓
- Every linked ADR is `approved` (or ASR resolved via Direct Input) ✓
- **Resolution** fields filled for all `designed` ASRs ✓

**Next → Step ④ (Spec Writing).**

---

## ASR.md Maintenance During This Step

| Event | Action |
|-------|--------|
| Starting review of an ASR | Status → `reviewing` |
| Opening an ADR for an ASR | Append ADR path under **Related ADRs** in `ASR.md` |
| ADR approved | Status → `designed`, fill **Resolution** |
| Direct Input chosen | Status → `designed`, fill **Resolution** |
| ASR deferred (user agreement) | Status → `deferred`, note reason |

Keep **Summary** table in sync with **ASR Detail** at all times.
