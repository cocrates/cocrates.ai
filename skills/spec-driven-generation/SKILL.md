---
name: spec-driven-generation
description: >-
  Produces deliverables — software, documents, images, presentations, and
  other artifacts — strictly from Spec documents in {project-slug}/spec/.
  Select when the user asks to build, create, or generate from a spec, from
  all specs, or "according to the specification"; when spec-writing hands off
  to generation; or when generation is requested but no ready spec exists yet
  (then create the {project-slug}/ project folder and spec/ASR.md registry, establish
  a PRD, identify ASRs, review them based on structural dependencies, and guide
  through adr-writing and spec-writing via Socratic dialogue).
compatibility: opencode
metadata:
  agent: cocrates
---

# Spec-Driven Generation — Artifact Production Skill

This skill **produces final artifacts from Spec documents**. The Spec is the **sole source of truth** for what to build — not the original user prompt, not ADR files, and not assumptions the agent fills in silently.

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

- **`{project-slug}`:** English **kebab-case** name for the project (e.g. `jsondb`, `item-catalog-api`). Derive from the product or deliverable name; confirm with the user if ambiguous.
- **`{project-slug}/spec/ASR.md`:** Living registry of Architecturally Significant Requirements — tracks identification, review, design, and approval. Owned and maintained by this skill (updated by `adr-writing` / `spec-writing` / `spec-driven-verification` when they touch ASRs). It is **not** a generation Spec — do not treat it as sole input for artifact production.
- **Do not** place `spec/`, `adr/`, or `verification/` at the repository root — always nest them under `{project-slug}/`. Keep `ASR.md` only at `{project-slug}/spec/ASR.md`.
- **On generation request:** If `{project-slug}/` does not exist yet, **create it** (and the `spec/` subdirectory as needed) before writing PRD or Spec files. This skill owns project-folder creation.

## Core Principle: Spec as Quality Contract

The Spec exists to capture every **ASR (Architecturally Significant Requirement)** — requirements and design decisions that **materially affect the structure, organization, or quality of the final deliverable**. Optimal generation requires the Spec to encode:

- **What** to produce — scope, format, audience, output location
- **How** it is organized — structure, composition, narrative or functional layout
- **What bar** applies — constraints, acceptance criteria, explicit quality expectations
- **What is excluded** — out-of-scope items that would otherwise be guessed at

Unstated ASRs force silent defaults — and defaults degrade quality. The readiness gate exists to surface those choices **before** generation, not to document them after the fact.

### Workflow Nature: Non-linear Lifecycle

The progression through **PRD → ASR Identification (`ASR.md`) → ADR (if needed) → Spec Writing → Generation → Verification** is **not a rigid, one-way waterfall process**. 
- It is an agile, continuous feedback loop where any stage can trigger a return to a previous phase.
- Gaps discovered during *Generation* or failures during *Verification* instantly re-open the *Readiness Gate* or trigger *ADR* reviews.

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

Generation reads from **`{project-slug}/spec/`**.

```
{project-slug}/spec/
├── ASR.md
├── PRD.md
└── {requirement-slug}.md
```

`ASR.md` is the decision-status registry. Generation input Specs are `PRD.md` and `{requirement-slug}.md` only — **exclude `ASR.md`** when loading Spec content for production.
### Multi-Spec & Multi-ASR Strategy

In large-scale projects, specifications must be broken down into multiple, **independent spec files** based on domain or component boundaries to maintain clarity and modularity. 

When dealing with a complex, multi-spec environment, do not focus on the generation or writing order of the files themselves. Instead, **focus on the structural dependencies between ASRs**, tracked in `{project-slug}/spec/ASR.md`:

1. **Map ASR Dependencies:** Analyze which foundational decisions must happen first (e.g., *Storage Model* must be decided before *API Surface*, which must be decided before *Client UI*). Record the order under **Dependency Order** in `ASR.md`.
2. **Sequential Review Path:** Present a recommended, step-by-step evaluation path to the user based on these ASR dependencies. Guide the user to review and resolve one ASR at a time; open ADRs as needed (many-to-many).
3. **Delegation to Spec-Writing:** The actual formatting and markdown documentation of independent specs are delegated strictly to the **`spec-writing`** skill; keep ASR statuses current in `ASR.md` throughout.

---

## ASR — Architecturally Significant Requirements

An **ASR** is a requirement or design decision that materially affects the structure, organization, or quality of the final deliverable. When an ASR is unstated, generation fills the gap with a silent default. Use this framework during the readiness gate when no Spec exists or when a Spec has material gaps.

**Every identified ASR must be registered in `{project-slug}/spec/ASR.md`.** Do not keep ASRs only in conversation memory.

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

Create or update **`{project-slug}/spec/ASR.md`** whenever ASRs are identified, reviewed, designed, or approved.

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
  - `{project-slug}/adr/{concern-slug}.md` — {proposed \| approved} — {how it relates}
- **Resolution path:** direct-input | adr | mixed
- **Resolution:** {Filled when designed/approved — concrete outcome, not ADR pointers alone}
- **Spec:** `{project-slug}/spec/{file}.md` — {section} (once encoded)
- **Notes:** {optional}
```

**Maintenance rules:**
1. Assign stable IDs (`ASR-001`, `ASR-002`, …). Never reuse IDs.
2. Update **Summary** and **ASR Detail** together — do not leave the table stale.
3. When opening an ADR for an ASR, set status to `reviewing` and append the ADR path under **Related ADRs**.
4. When an ADR is approved (or Direct Input settles the ASR), set status to `designed`, fill **Resolution**, then sync Spec via `spec-writing` and set `approved` after user confirmation.
5. One ASR may list multiple ADRs; one ADR may appear under multiple ASRs — keep both sides in sync with `adr-writing`.

### ASR Decision Pathways (Based on User Context)

When evaluating open ASRs, do not force an ADR for every single concern. Match the user's domain expertise and context:
1. **Direct Input Route (High Context):** If the problem is straightforward, has a natural default, or the user has clear preferences, ask direct questions. Let them dictate the choice, update `ASR.md` (`designed` → `approved`), and update the Spec immediately.
2. **ADR Analysis Route (Low Context / Complex Trade-offs):** If the user is unsure of the solution, or if the problem involves complex architectural trade-offs (e.g., competing tech stacks, structural patterns), set the ASR to `reviewing` and trigger **`adr-writing`**. Link every new ADR back to the relevant ASR ID(s). Multiple ADRs may be opened for one ASR; one ADR may cover multiple ASRs.

---

## Workflow

### 1. Determine Generation Scope, Create Project Folder & Establish PRD

Before analyzing detailed technical specifications, clearly define the high-level objective and boundary of the product.

1. **Identify Goal:** Understand what artifact type, target audience, and core function the user wants to build.
2. **Resolve `{project-slug}`:** Choose an English kebab-case project folder name from the product/deliverable. Confirm with the user if unclear.
3. **Create Project Folder:** If `{project-slug}/` does not exist, create it (and `{project-slug}/spec/` as needed) before writing any documents.
4. **Document PRD:** Consolidate high-level requirements into **`{project-slug}/spec/PRD.md` (Product Requirement Definition)**.
5. **Initialize ASR Registry:** Create **`{project-slug}/spec/ASR.md`** (if missing) using the template above, even if the first ASRs are still being discovered.
6. **The Anchor:** `{project-slug}/spec/PRD.md` serves as the ultimate product anchor; `{project-slug}/spec/ASR.md` serves as the decision-status board.

### 2. Identify ASRs & Spec Readiness Gate

1. **Identify ASRs** from the PRD, user context, and universal categories. Register each as `identified` in `ASR.md` with ID, title, category, and statement.
2. **Map Dependencies:** Fill **Dependency Order** in `ASR.md`. Present the recommended review path to the user.
3. **Readiness Gate:** Evaluate whether requirement Specs (`PRD.md` and `{requirement-slug}.md`, not `ASR.md`) are **sufficient** for the deliverables in `PRD.md`, using `spec/ASR.md` as the checklist of what must be settled. Do not rely on rigid ASR counts; assess fitness for purpose.

- **If Gaps are Found:**
  1. Ensure each missing decision is an ASR entry in `ASR.md` (add if absent).
  2. Set the active ASR to `reviewing`. Route via **ASR Decision Pathways** (Direct Input vs. ADR Analysis).
  3. **ADR route:** Hand off to `adr-writing` with the ASR ID(s). One ASR may spawn multiple ADRs; one ADR may list multiple ASR IDs. Keep Related ADRs / Related ASRs bidirectional.
  4. **Real-time sync:** When a decision is confirmed (Direct Input or approved ADR), update `ASR.md` to `designed`, sync Spec via `spec-writing`, then mark `approved` after user confirmation. Do not wait to collect all decisions before updating documents.
  5. Do not proceed to generation while any **blocking** ASR remains `identified` or `reviewing` (unless explicitly `deferred`).

- **If Spec exists but material gaps remain:**
  - Register or reopen ASRs in `ASR.md`, resolve them Socratically, and update Spec before generating.

### 3. Load Spec(s)

Read the full content of each Spec in scope (including `{project-slug}/spec/PRD.md` and related `{project-slug}/spec/*.md` files).

Extract and keep in working memory:
- **Requirement** — deliverable statement
- **Decisions** — stated outcomes (not ADR pointers)
- **Requirements** — must-haves (testable)
- **Constraints** — limits generation must respect
- **Out of Scope** — explicit non-goals

Treat the Spec as authoritative over the conversation history. If the user's latest message contradicts the Spec, ask which source to follow before generating.

### 4. Generate

Produce the artifact **strictly from the Spec**:
- Implement every **Requirements** bullet unless blocked — report any that could not be met.
- Honor **Decisions**, **Constraints**, and **Out of Scope**.
- Do not add features, sections, or dependencies the user did not specify in the Spec.
- Use existing project conventions when the Spec does not specify otherwise.

### 5. Report and Hand Off

After generation:
1. Report **output paths** and a concise summary of what was produced.
2. Map key **Requirements** bullets to what was delivered (brief checklist).
3. Recommend **`spec-driven-verification`** — results are saved to `{project-slug}/verification/{requirement-slug}.md` for review.

### 6. Post-Generation Feedback & Defect Handling Loop

If verification fails, or if the user identifies quality issues or regressions after generation, **never patch the artifact directly**. Follow the reverse traceability loop:

1. **Root Cause Analysis:** Identify which requirement or design decision was missing, flawed, or misunderstood — register or reopen the corresponding ASR in `ASR.md`.
2. **Review Alternatives:** If fixing the issue requires evaluating new trade-offs or shifting architecture, set ASR to `reviewing` and load **`adr-writing`** (link Related ASRs). If it is a simple correction, use Direct Input.
3. **Update Contract:** Reflect the approved fix into Spec via `spec-writing`, and advance the ASR to `designed` → `approved` in `ASR.md`.
4. **Regenerate / Refine:** Re-run this generation skill focusing strictly on the modified spec sections to patch or rebuild the artifact.

---

## Socratic Dialogue Rules (For Gatekeeping & Gaps)

Use when step 2 blocks on missing or unclear specification.

1. **State the Current Step:** Explicitly tell the user where we are in the lifecycle (e.g., ASR identification, ASR-00x review, ADR writing, Spec writing, Generation) and cite the ASR ID from `ASR.md`.
2. **Reduce User Fatigue:** When asking about open ASRs, always propose a recommended default or best-practice option based on project context (e.g., *"We recommend X because of Y. Shall we proceed with this, or do you have another preference?"*). Do not ask open-ended questions without guidance.
3. **One Question at a Time:** Focus on one ASR (or one dependency branch) at a time to keep the conversation structured.
4. **Immediate Documentation:** Update `ASR.md` status first, then sync approved outcomes to Spec before moving to the next ASR.

---

## Prohibitions

- Generating final artifacts **without** a ready Spec when blocking decisions remain unresolved.
- Identifying ASRs in conversation only — without registering them in `{project-slug}/spec/ASR.md`.
- Leaving ASR status stale after Direct Input, ADR open/approve, or Spec sync.
- Writing Spec / ADR / verification files at the repository root instead of under `{project-slug}/` (including placing `ASR.md` outside `{project-slug}/spec/`).
- Skipping creation of `{project-slug}/` when it does not yet exist on a generation request.
- Reading **`{project-slug}/adr/`** files as a substitute for Spec content during generation.
- Treating the original prompt as authoritative when it **contradicts** the Spec.
- Adding requirements, features, or scope not explicitly stated or implied by the Spec.
- Patching or editing generated artifacts directly without updating `{project-slug}/spec/PRD.md` or `{project-slug}/spec/*.md` first during a defect loop.
- Asking a flat checklist of open questions without providing recommendations or considering ASR dependencies.
- Silently merging conflicting Specs.

---

## Completion Criteria

- `{project-slug}/` project folder created (if missing), `{project-slug}/spec/PRD.md` established, and `{project-slug}/spec/ASR.md` initialized/maintained.
- Blocking ASRs reviewed (Direct Input or ADR), statuses advanced to `approved` (or explicitly `deferred`), and Related ADR links kept bidirectional.
- All detailed `{project-slug}/spec/*.md` files evaluated for **sufficiency** against the PRD and synced from approved ASR resolutions.
- Deliverable produced at agreed paths strictly following Spec **Requirements**, **Decisions**, and **Constraints**.
- User informed of outputs, mapped back to Spec bullets, and offered **`spec-driven-verification`**.

This skill owns production from Spec and the ASR registry; **`spec-driven-verification`** owns checking the result against the Spec.
