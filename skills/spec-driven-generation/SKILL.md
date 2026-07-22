---
name: spec-driven-generation
description: >-
  Produces deliverables — software, documents, images, presentations, and
  other artifacts — strictly from Spec documents in {project-root}/spec/.
  Select when the user asks to build, create, or generate from a spec, from
  all specs, or "according to the specification"; when spec-writing hands off
  to generation; or when generation is requested but no ready spec exists yet
  (then enter Step 0 Spec Readiness Gate: resolve/create {project-root} if needed,
  approve PRD, identify/review ASRs via ASR.md, approve ADRs, write Specs
  through adr-writing and spec-writing, then generate).
compatibility: opencode
metadata:
  agent: cocrates
---

# Spec-Driven Generation — Artifact Production Skill

This skill **produces final artifacts from Spec documents**. The Spec is the **sole source of truth** for what to build — not the original user prompt, not ADR files, and not assumptions the agent fills in silently.

Deliverables include **software, documents, images, presentations, blog series, reports**, and any other finished work the Spec describes.

## Resolve Project Root

Before writing Spec / ADR / verification files, resolve **`{project-root}`** — the folder that holds this deliverable.

Workspace layouts fall into three types. **Inspect the workspace first**, then match:

| Type | When | `{project-root}` |
|------|------|------------------|
| **1** | Workspace *is* the single project (no nested project folder) | `.` (workspace root) |
| **2** | Workspace holds multiple peer projects | `{project-slug}/` |
| **3** | Workspace groups projects by kind | `{kind}/{project-slug}/` (e.g. `apps/{project-slug}/`) |

**Examples:**
- Type 1 → `spec/PRD.md` at workspace root
- Type 2 → `{project-slug}/spec/PRD.md`
- Type 3 → `apps/{project-slug}/spec/PRD.md` (kind folder as used in the workspace)

**Rules:**
1. Infer the type from existing structure (e.g. `spec/` at root, peer project folders, or kind folders such as `apps/`, `services/`). Prefer an **existing** matching folder over creating a new one.
2. **Before creating** a new project folder, confirm **location and name** with the user.
3. If the project folder already exists, use it — do not recreate or relocate silently.
4. **`{project-slug}`:** English **kebab-case** (e.g. `jsondb`, `item-catalog-api`). Derive from the product or deliverable name; confirm if ambiguous. Type 1 has no slug folder.

This skill owns project-folder creation when generation starts and `{project-root}` is missing (after user confirmation).

## Project Folder Layout

All Spec / ADR / verification work for a deliverable lives under **`{project-root}`**:

```
{project-root}/
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

- **`{project-root}/spec/ASR.md`:** Living registry of Architecturally Significant Requirements — tracks identification, review, design, and approval. Owned and maintained by this skill (updated by `adr-writing` / `spec-writing` / `spec-driven-verification` when they touch ASRs). It is **not** a generation Spec — do not treat it as sole input for artifact production.
- Place `spec/`, `adr/`, and `verification/` under `{project-root}` only. Keep `ASR.md` only at `{project-root}/spec/ASR.md`.
- **On generation request:** After resolving (and confirming if creating) `{project-root}`, ensure `spec/` exists before writing PRD or Spec files.

## Core Principle: Spec as Quality Contract

The Spec exists to capture every **ASR (Architecturally Significant Requirement)** — requirements and design decisions that **materially affect the structure, organization, or quality of the final deliverable**. Optimal generation requires the Spec to encode:

- **What** to produce — scope, format, audience, output location
- **How** it is organized — structure, composition, narrative or functional layout
- **What bar** applies — constraints, acceptance criteria, explicit quality expectations
- **What is excluded** — out-of-scope items that would otherwise be guessed at

Unstated ASRs force silent defaults — and defaults degrade quality. **Step 0 (Spec Readiness Gate)** exists to surface those choices **before** generation, not to document them after the fact.

### Workflow Nature: Non-linear Lifecycle

Every generation request **starts at Step 0 (Spec Readiness Gate)** once to find the correct entry step. After that, steps advance **forward** (1 → 2 → 3 → ADR approval → 4 → 5) without re-checking earlier gates.
- The overall lifecycle (**PRD → ASR → ADR → Spec → Generation → Verification**) can still loop when gaps appear later.
- Gaps discovered during *Generation* or failures during *Verification* re-enter **Step 0** only to resume at the appropriate remediation step.

**ASR vs ADR vs Spec vs generation:**

| | ASR (`ASR.md`) | ADR | Spec | Generation (this skill) |
|---|-----|------|------|-------------------------|
| **Unit** | One architecturally significant requirement | One concern (one question) | One deliverable or requirement scope | One or more artifacts per Spec |
| **Purpose** | Track what must be decided and its lifecycle | Compare alternatives; record approval | State decided requirements self-contained | **Produce** the artifact |
| **Reader needs** | Know what is open, reviewing, designed, or approved | Understand options and tradeoffs | Understand what to build | Receive the finished work |
| **Role in pipeline** | Decision backlog & status board | Design review for open ASRs | **Sole input** for generation | Execution |

**ASR ↔ ADR cardinality (many-to-many):**
- **One ASR → many ADRs:** A single ASR may require several design reviews (e.g. *Storage Model* → partitioning ADR + indexing ADR + consistency ADR).
- **One ADR → many ASRs:** A single ADR may resolve or partially address multiple related ASRs when the concern spans them.
- Always record both sides: list Related ADRs on each ASR in `ASR.md`, and list Related ASRs on each ADR file.

Match the user's language (Korean, English, etc.) in user-facing messages. Artifact content follows the language implied by the Spec unless the Spec states otherwise.

---

## Spec Input Scope

Generation reads from **`{project-root}/spec/`**.

```
{project-root}/spec/
├── ASR.md
├── PRD.md
└── {requirement-slug}.md
```

`ASR.md` is the decision-status registry. Generation input Specs are `PRD.md` and `{requirement-slug}.md` only — **exclude `ASR.md`** when loading Spec content for production.
### Multi-Spec & Multi-ASR Strategy

In large-scale projects, specifications must be broken down into multiple, **independent spec files** based on domain or component boundaries to maintain clarity and modularity. 

When dealing with a complex, multi-spec environment, do not focus on the generation or writing order of the files themselves. Instead, **focus on the structural dependencies between ASRs**, tracked in `{project-root}/spec/ASR.md`:

1. **Map ASR Dependencies:** Analyze which foundational decisions must happen first (e.g., *Storage Model* must be decided before *API Surface*, which must be decided before *Client UI*). Record the order under **Dependency Order** in `ASR.md`.
2. **Sequential Review Path:** Present a recommended, step-by-step evaluation path to the user based on these ASR dependencies. Guide the user to review and resolve one ASR at a time; open ADRs as needed (many-to-many).
3. **Delegation to Spec-Writing:** The actual formatting and markdown documentation of independent specs are delegated strictly to the **`spec-writing`** skill; keep ASR statuses current in `ASR.md` throughout.

---

## ASR — Architecturally Significant Requirements

An **ASR** is a requirement or design decision that materially affects the structure, organization, or quality of the final deliverable. When an ASR is unstated, generation fills the gap with a silent default. Use this framework during **Step 0** when gates fail or Specs have material gaps.

**Every identified ASR must be registered in `{project-root}/spec/ASR.md`.** Do not keep ASRs only in conversation memory.

### Universal ASR categories (always consider)

| ASR category | What to surface | Examples across artifact types |
|--------------|-----------------|--------------------------------|
| **Purpose & audience** | Who is this for? What job does it do? | API consumers; report readers; presentation attendees; image viewers |
| **Deliverable form** | Artifact type(s), format, and output location | Go module; markdown report; slide deck; PNG; multi-file series |
| **Scope boundary** | What is in vs out? | MVP endpoints only; executive summary not appendices |
| **Quality bar** | What does "done" and "good enough" mean? | Test coverage; citation depth; visual fidelity; word count |
| **Constraints** | Hard limits generation must respect | Language, deadlines, brand guidelines, dependencies, budget |
| **Structure & organization** | How is content or function arranged? | API layers; document sections; episode order; slide narrative |
| **Integration & dependencies** | What must this connect to or build on? | Existing codebase; prior episodes; corporate template |
| **Process & staging** | Approval gates or iterative delivery? | Draft-review-publish; staged image approval |

### ASR Lifecycle Status

Track each ASR in `ASR.md` with exactly one status:

| Status | Meaning |
|--------|---------|
| `identified` | Discovered and registered; not yet under active review |
| `reviewing` | Under discussion — Direct Input or one or more ADRs in progress |
| `designed` | Solution chosen (via Direct Input and/or approved ADR(s)); Spec sync may still be pending |
| `approved` | User confirmed; outcomes encoded in Spec; blocking work for this ASR is done |
| `deferred` | Intentionally postponed; non-blocking for the current generation scope |

Status transitions: `identified` → `reviewing` → `designed` → `approved`. Use `deferred` only with explicit user agreement.

### ASR.md File

Create or update **`{project-root}/spec/ASR.md`** whenever ASRs are identified, reviewed, designed, or approved.

```markdown
# Architecturally Significant Requirements

Living registry for {project-slug}. Status of each ASR must stay current.

## Summary

| ID | Title | Category | Status | Related ADRs | Spec |
|----|-------|----------|--------|--------------|------|
| ASR-001 | {short title} | {category} | identified \| reviewing \| designed \| approved \| deferred | adr/{slug}.md, … | spec/{file}.md |
| ASR-002 | … | … | … | — | — |

## Dependency Order (recommended review path)

1. ASR-00x → ASR-00y → …
{List IDs in the order structural dependencies require.}

## ASR Detail

### ASR-001 — {Title}

- **Category:** {from universal categories}
- **Status:** identified | reviewing | designed | approved | deferred
- **Statement:** {What must be decided — one clear requirement/concern}
- **Why it matters:** {Impact if left to silent default}
- **Depends on:** ASR-00x (optional)
- **Related ADRs:** (many-to-many — zero or more)
  - `{project-root}/adr/{concern-slug}.md` — {proposed \| approved} — {how it relates}
- **Resolution path:** direct-input | adr | mixed
- **Resolution:** {Filled when designed/approved — concrete outcome, not ADR pointers alone}
- **Spec:** `{project-root}/spec/{file}.md` — {section} (once encoded)
- **Notes:** {optional}
```

**Maintenance rules:**
1. Assign stable IDs (`ASR-001`, `ASR-002`, …). Never reuse IDs.
2. Update **Summary** and **ASR Detail** together — do not leave the table stale.
3. When opening an ADR for an ASR, set status to `reviewing` and append the ADR path under **Related ADRs**.
4. When an ADR is approved (or Direct Input settles the ASR), set status to `designed`, fill **Resolution**. Spec sync and `approved` happen in **Step 4** after user confirmation of the Spec.
5. One ASR may list multiple ADRs; one ADR may appear under multiple ASRs — keep both sides in sync with `adr-writing`.

### ASR Decision Pathways (Based on User Context)

When evaluating open ASRs (Step 3), do not force an ADR for every single concern. Match the user's domain expertise and context:
1. **Direct Input Route (High Context):** If the problem is straightforward, has a natural default, or the user has clear preferences, ask direct questions. Let them dictate the choice, update `ASR.md` to `designed` (fill **Resolution**), then proceed — Spec sync is **Step 4**.
2. **ADR Analysis Route (Low Context / Complex Trade-offs):** If the user is unsure of the solution, or if the problem involves complex architectural trade-offs (e.g., competing tech stacks, structural patterns), set the ASR to `reviewing` and trigger **`adr-writing`**. Link every new ADR back to the relevant ASR ID(s). Multiple ADRs may be opened for one ASR; one ADR may cover multiple ASRs. After user approval → **Step 4** (do not invent approval).

---

## Workflow

**Begin at Step 0 once** to locate the first incomplete gate, then jump to that step. **After a step finishes, proceed to the next step** — do not re-enter Step 0 between Steps 1–5. Re-enter Step 0 only on a new generation request or when the post-generation defect loop needs to resume work.

```
Step 0 Spec Readiness Gate (entry only)
  ├─ Gate 1 fail → Step 1
  ├─ Gate 2 fail → Step 2
  ├─ Gate 3 fail → Step 3
  ├─ Gate 4 fail → ADR approval → Step 4
  ├─ Gate 5 fail → Step 4
  └─ All pass   → Step 5

Forward path after entry:
  Step 1 → Step 2 → Step 3 → (ADR approval if needed) → Step 4 → Step 5
```

### Step 0. Spec Readiness Gate

Resolve `{project-root}` first (inspect workspace type; confirm slug/location if creating). If `{project-root}` is missing and is not the workspace root (Type 2/3), create it after user confirmation (and `spec/`) before evaluating gates that require files.

Evaluate gates **in order**; stop at the **first** failure and jump there. Do not re-run this gate sequence after each subsequent step.

| Gate | Check | Pass criteria | On fail |
|------|-------|---------------|---------|
| **Gate 1** | PRD approved? | `{project-root}/spec/PRD.md` exists and the user has explicitly approved it | → **Step 1** |
| **Gate 2** | ASRs identified for the PRD? | Blocking ASRs from the PRD, user context, and universal categories are registered in `ASR.md` (at least `identified`); **Dependency Order** is filled | → **Step 2** |
| **Gate 3** | ASRs all reviewed? | No blocking ASR remains `identified` (each is `reviewing` or beyond, or explicitly `deferred`) | → **Step 3** |
| **Gate 4** | ADRs all approved? | Every ADR linked from blocking ASRs is `approved` (or the ASR used Direct Input with no open ADR). No Related ADR remains `proposed` | → Pause for user ADR approval via `adr-writing` (do **not** invent approval). When approved → **Step 4** |
| **Gate 5** | Spec created? | Requirement Specs (`PRD.md` + needed `{requirement-slug}.md`, not `ASR.md`) encode approved resolutions and are **sufficient** for the deliverables in the PRD | → **Step 4** |

**All gates pass → Step 5.**

Do not rely on rigid ASR counts; assess fitness for purpose.

### Step 1. PRD Creation

1. **Identify Goal:** Artifact type, target audience, and core function.
2. **Create Project Folder:** Ensure `{project-root}` and `{project-root}/spec/` exist (confirm location/name with the user before creating a new folder).
3. **Document PRD:** Write or update **`{project-root}/spec/PRD.md`** (product anchor — high-level scope, goals, global constraints).
4. **Initialize ASR Registry:** Create **`{project-root}/spec/ASR.md`** if missing (template above), even before ASRs are fully discovered.
5. **Obtain PRD Approval:** Present the PRD and get explicit user approval.

**Next → Step 2.**

### Step 2. ASR Creation (Identification)

1. Identify ASRs from the approved PRD, user context, and universal categories.
2. Register each as `identified` in `ASR.md` (ID, title, category, statement, why it matters).
3. Fill **Dependency Order**; present the recommended review path to the user.

**Next → Step 3.**

### Step 3. ASR Review → ADR Creation

Review blocking ASRs in **Dependency Order**, one ASR (or dependency branch) at a time:

1. Set the active ASR to `reviewing`.
2. Route via **ASR Decision Pathways**:
   - **Direct Input:** User chooses; set ASR to `designed`; fill **Resolution**.
   - **ADR Analysis:** Hand off to `adr-writing` with the ASR ID(s). One ASR may spawn multiple ADRs; one ADR may list multiple ASR IDs. Keep Related ADRs / Related ASRs bidirectional. Leave ASR `reviewing` until linked ADRs are approved.
3. Document resolutions in `ASR.md` as they settle — do not wait to batch every ASR before recording.

When all blocking ASRs have a resolution path complete (Direct Input → `designed`, or ADRs created):
- If any Related ADR is still `proposed`: pause for **user ADR approval** (same as Gate 4). Do **not** invent approval.
- When all linked ADRs are `approved` (or no open ADRs remain): set affected ASRs to `designed` with concrete **Resolution**.

**Next → Step 4.**

### Step 4. Spec Creation

1. Hand off to **`spec-writing`** to encode `designed` / approved ASR resolutions into `{project-root}/spec/{requirement-slug}.md` (and update `PRD.md` if needed).
2. Specs must be self-contained — copy decided outcomes in full; never substitute ADR links for requirements.
3. After Spec sync and user confirmation, mark related ASRs `approved` in `ASR.md`.

**Next → Step 5.**

### Step 5. Generate from Spec

1. **Load Spec(s):** Read `PRD.md` and in-scope `{requirement-slug}.md` (**exclude `ASR.md`**). Keep in working memory: Requirement, Decisions, Requirements, Constraints, Out of Scope. Spec overrides conversation history; if the latest user message contradicts the Spec, ask which source to follow.
2. **Generate:** Produce the artifact **strictly from the Spec** — every Requirements bullet, honor Decisions / Constraints / Out of Scope, no silent extras, follow project conventions when Spec is silent.
3. **Report and Hand Off:** Output paths, brief Requirements checklist, recommend **`spec-driven-verification`** → `{project-root}/verification/{requirement-slug}.md`.

### Post-Generation Feedback & Defect Handling Loop

If verification fails, or the user finds quality issues after generation, **never patch the artifact directly**:

1. **Root Cause:** Register or reopen the ASR in `ASR.md`.
2. **Re-enter Step 0** once to resume at the correct remediation step (then advance forward through Steps 1–5 as needed).
3. **Regenerate:** Complete through Step 5 against the updated Spec only.

---

## Socratic Dialogue Rules (For Gatekeeping & Gaps)

Use when Step 0 routes into incomplete work, or during Steps 1–4 clarification.

1. **State the Current Step:** Tell the user the step (e.g., Step 3, ASR-00x review) and cite the ASR ID from `ASR.md`.
2. **Reduce User Fatigue:** When asking about open ASRs, always propose a recommended default or best-practice option based on project context (e.g., *"We recommend X because of Y. Shall we proceed with this, or do you have another preference?"*). Do not ask open-ended questions without guidance.
3. **One Question at a Time:** Focus on one ASR (or one dependency branch) at a time to keep the conversation structured.
4. **Immediate Documentation:** Update `ASR.md` as statuses change; sync Spec in Step 4 before advancing to Step 5.

---

## Prohibitions

- Generating final artifacts while any Step 0 gate still fails.
- Generating final artifacts **without** a ready Spec when blocking decisions remain unresolved.
- Identifying ASRs in conversation only — without registering them in `{project-root}/spec/ASR.md`.
- Leaving ASR status stale after Direct Input, ADR open/approve, or Spec sync.
- Writing Spec / ADR / verification files outside `{project-root}` (including placing `ASR.md` outside `{project-root}/spec/`).
- Skipping resolve/create of `{project-root}` when it does not yet exist on a generation request (or creating it without user confirmation of location and name).
- Reading **`{project-root}/adr/`** files as a substitute for Spec content during generation.
- Treating the original prompt as authoritative when it **contradicts** the Spec.
- Adding requirements, features, or scope not explicitly stated or implied by the Spec.
- Patching or editing generated artifacts directly without updating `{project-root}/spec/PRD.md` or `{project-root}/spec/*.md` first during a defect loop.
- Asking a flat checklist of open questions without providing recommendations or considering ASR dependencies.
- Silently merging conflicting Specs.

---

## Completion Criteria

- **Step 0** used once at entry (or defect resume): first failing gate selected; then Steps advance forward without re-gating.
- PRD approved → ASRs identified → blocking ASRs reviewed → linked ADRs approved (or Direct Input) → Specs synced → generation (Steps 1–5).
- `{project-root}` resolved (and created after confirmation when needed); `PRD.md` and `ASR.md` maintained; Related ADR ↔ ASR links bidirectional.
- Deliverable produced at agreed paths strictly following Spec **Requirements**, **Decisions**, and **Constraints** (Step 5).
- User informed of outputs, mapped back to Spec bullets, and offered **`spec-driven-verification`**.

This skill owns production from Spec and the ASR registry; **`spec-driven-verification`** owns checking the result against the Spec.
