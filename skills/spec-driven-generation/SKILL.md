---
name: spec-driven-generation
description: >-
  Produces deliverables — software, documents, images, presentations, and
  other artifacts — strictly from Spec documents in spec/. Select when the
  user asks to build, create, or generate from a spec, from all specs, or
  "according to the specification"; when spec-writing hands off to generation;
  or when generation is requested but no ready spec exists yet (then establish
  a PRD, identify ASRs, review them based on structural dependencies, and guide
  through adr-writing and spec-writing via Socratic dialogue).
compatibility: opencode
metadata:
  agent: cocrates
---

# Spec-Driven Generation — Artifact Production Skill

This skill **produces final artifacts from Spec documents**. The Spec is the **sole source of truth** for what to build — not the original user prompt, not ADR files, and not assumptions the agent fills in silently.

Deliverables include **software, documents, images, presentations, blog series, reports**, and any other finished work the Spec describes.

## Core Principle: Spec as Quality Contract

The Spec exists to capture every **ASR (Architecturally Significant Requirement)** — requirements and design decisions that **materially affect the structure, organization, or quality of the final deliverable**. Optimal generation requires the Spec to encode:

- **What** to produce — scope, format, audience, output location
- **How** it is organized — structure, composition, narrative or functional layout
- **What bar** applies — constraints, acceptance criteria, explicit quality expectations
- **What is excluded** — out-of-scope items that would otherwise be guessed at

Unstated ASRs force silent defaults — and defaults degrade quality. The readiness gate exists to surface those choices **before** generation, not to document them after the fact.

### Workflow Nature: Non-linear Lifecycle

The progression through **PRD/ASR Identification → ADR (if needed) → Spec Writing → Generation → Verification** is **not a rigid, one-way waterfall process**. 
- It is an agile, continuous feedback loop where any stage can trigger a return to a previous phase.
- Gaps discovered during *Generation* or failures during *Verification* instantly re-open the *Readiness Gate* or trigger *ADR* reviews.

**Spec vs ADR vs generation:**

| | ADR | Spec | Generation (this skill) |
|---|-----|------|-------------------------|
| **Unit** | One concern (one question) | One deliverable or requirement scope | One or more artifacts per Spec |
| **Purpose** | Compare alternatives; record approval | State decided requirements self-contained | **Produce** the artifact |
| **Reader needs** | Understand options and tradeoffs | Understand what to build | Receive the finished work |
| **Role in pipeline** | Decision audit trail | **Sole input** for generation | Execution |

Match the user's language (Korean, English, etc.) in user-facing messages. Artifact content follows the language implied by the Spec unless the Spec states otherwise.

---

## Spec Input Scope

Generation reads from the project **`spec/`** directory at the repository root.

```
spec/
├── PRD.md
└── {requirement-slug}.md
```

### Multi-Spec & Multi-ASR Strategy

In large-scale projects, specifications must be broken down into multiple, **independent spec files** based on domain or component boundaries to maintain clarity and modularity. 

When dealing with a complex, multi-spec environment, do not focus on the generation or writing order of the files themselves. Instead, **focus on the structural dependencies between ASRs (Architecturally Significant Requirements)**:

1. **Map ASR Dependencies:** Analyze which foundational decisions must happen first (e.g., *Storage Model* must be decided before *API Surface*, which must be decided before *Client UI*).
2. **Sequential Review Path:** Present a recommended, step-by-step evaluation path to the user based on these ASR dependencies. Guide the user to review and resolve one architectural concern at a time to prevent cognitive overload.
3. **Delegation to Spec-Writing:** The actual formatting and markdown documentation of these independent specs are delegated strictly to the **`spec-writing`** skill.

---

## ASR — Architecturally Significant Requirements

An **ASR** is a requirement or design decision that materially affects the structure, organization, or quality of the final deliverable. When an ASR is unstated, generation fills the gap with a silent default. Use this framework during the readiness gate when no Spec exists or when a Spec has material gaps.

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

### ASR Decision Pathways (Based on User Context)

When evaluating open ASRs, do not force an ADR for every single concern. Match the user's domain expertise and context:
1. **Direct Input Route (High Context):** If the problem is straightforward, has a natural default, or the user has clear preferences, ask direct questions. Let them dictate the choice and update the Spec immediately.
2. **ADR Analysis Route (Low Context / Complex Trade-offs):** If the user is unsure of the solution, or if the problem involves complex architectural trade-offs (e.g., competing tech stacks, structural patterns), trigger **`adr-writing`**. Present a structured options analysis for the user to review.

---

## Workflow

### 1. Determine Generation Scope & Establish PRD

Before analyzing detailed technical specifications, clearly define the high-level objective and boundary of the product.

1. **Identify Goal:** Understand what artifact type, target audience, and core function the user wants to build.
2. **Document PRD:** Consolidate these high-level requirements into a **`spec/PRD.md` (Product Requirement Definition)** file at the repository root.
3. **The Anchor:** The `spec/PRD.md` serves as the ultimate anchor for the project, mapping out what deliverables are expected.

### 2. Spec Readiness Gate (Sufficiency Check)

The Readiness Gate evaluates whether the detailed specifications in `spec/*.md` are **sufficient** to successfully generate the deliverables defined in `spec/PRD.md`. Do not rely on rigid checklists or numbers of ASRs; assess fitness for purpose.

- **If Gaps are Found:**
  1. Identify missing architectural decisions or constraints that would trigger silent defaults.
  2. Route dynamically using the **ASR Decision Pathways** (Direct Input vs. ADR Analysis).
  3. **Real-time Specification Sync:** The moment a user confirms a decision (either directly or via an approved ADR), **instantly sync and update the corresponding `spec/*.md` file** via `spec-writing`. Do not wait to collect all decisions before updating documents.
  4. Do not proceed to generation until blocking gaps are resolved and encoded in the Spec.

- **If Spec exists but material gaps remain:**
  - Handle ambiguous requirements or missing ASRs by surfacing them Socratically, resolving them based on user context, and instantly updating the Spec before generating.

### 3. Load Spec(s)

Read the full content of each Spec in scope (including `spec/PRD.md` and related `spec/*.md` files).

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
3. Recommend **`spec-driven-verification`** — results are saved to `verification/{requirement-slug}.md` for review.

### 6. Post-Generation Feedback & Defect Handling Loop

If verification fails, or if the user identifies quality issues or regressions after generation, **never patch the artifact directly**. Follow the reverse traceability loop:

1. **Root Cause Analysis:** Identify which requirement or design decision was missing, flawed, or misunderstood.
2. **Review Alternatives:** If fixing the issue requires evaluating new trade-offs or shifting architecture, load **`adr-writing`** to approve the new solution. If it is a simple correction, get direct user confirmation.
3. **Update Contract:** Reflect the approved fix or changed scope into `spec/PRD.md` or the corresponding `spec/*.md` immediately.
4. **Regenerate / Refine:** Re-run this generation skill focusing strictly on the modified spec sections to patch or rebuild the artifact.

---

## Socratic Dialogue Rules (For Gatekeeping & Gaps)

Use when step 2 blocks on missing or unclear specification.

1. **State the Current Step:** Explicitly tell the user where we are in the lifecycle (e.g., ASR review, Readiness Gate, ADR writing, Spec writing, Generation).
2. **Reduce User Fatigue:** When asking about open ASRs, always propose a recommended default or best-practice option based on project context (e.g., *"We recommend X because of Y. Shall we proceed with this, or do you have another preference?"*). Do not ask open-ended questions without guidance.
3. **One Question at a Time:** Focus on one architectural concern or ASR dependency branch at a time to keep the conversation structured.
4. **Immediate Documentation:** Sync approved answers to the respective Spec file immediately before moving to the next question.

---

## Prohibitions

- Generating final artifacts **without** a ready Spec when blocking decisions remain unresolved.
- Reading **`adr/`** files as a substitute for Spec content during generation.
- Treating the original prompt as authoritative when it **contradicts** the Spec.
- Adding requirements, features, or scope not explicitly stated or implied by the Spec.
- Patching or editing generated artifacts directly without updating `spec/PRD.md` or `spec/*.md` first during a defect loop.
- Asking a flat checklist of open questions without providing recommendations or considering ASR dependencies.
- Silently merging conflicting Specs.

---

## Completion Criteria

- `spec/PRD.md` established at the start to define high-level product boundaries.
- All detailed `spec/*.md` files evaluated for **sufficiency** against the PRD.
- Gaps resolved incrementally (via Direct Input or ADR) and synced to Specs in real-time.
- Deliverable produced at agreed paths strictly following Spec **Requirements**, **Decisions**, and **Constraints**.
- User informed of outputs, mapped back to Spec bullets, and offered **`spec-driven-verification`**.

This skill owns production from Spec; **`spec-driven-verification`** owns checking the result against the Spec.
