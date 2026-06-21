---
name: spec-driven-generation
description: >-
  Produces deliverables — software, documents, images, presentations, and
  other artifacts — strictly from Spec documents in spec/. Select when the
  user asks to build, create, or generate from a spec, from all specs, or
  "according to the specification"; when spec-writing hands off to generation;
  or when generation is requested but no ready spec exists yet (then identify
  ASRs (Architecturally Significant Requirements), review them with the user,
  and guide through adr-writing and spec-writing via Socratic dialogue).
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

When no Spec exists, the first job is not to build. It is to **identify ASRs**, **review them with the user**, and encode the outcomes in Spec (via ADR when real alternatives exist).

**Spec vs ADR vs generation:**

| | ADR | Spec | Generation (this skill) |
|---|-----|------|-------------------------|
| **Unit** | One concern (one question) | One deliverable or requirement scope | One or more artifacts per Spec |
| **Purpose** | Compare alternatives; record approval | State decided requirements self-contained | **Produce** the artifact |
| **Reader needs** | Understand options and tradeoffs | Understand what to build | Receive the finished work |
| **Role in pipeline** | Decision audit trail | **Sole input** for generation | Execution |

Match the user's language (Korean, English, etc.) in user-facing messages. Artifact content follows the language implied by the Spec unless the Spec states otherwise.

## Spec Input Scope

Generation reads from the project **`spec/`** directory at the repository root.

```
spec/
└── {requirement-slug}.md
```

### Which Spec(s) to Use

Determine scope from the user's request:

| Request pattern | Action |
|-----------------|--------|
| Named file or slug (e.g. `spec/item-catalog-api.md`, `item-catalog-api`) | Read **that Spec only** |
| "All specs", "everything in spec/", or cross-cutting build spanning multiple deliverables | Read **all** `spec/*.md` files |
| Unspecified | Ask one Socratic question: single deliverable vs multiple — then read the matching set |

When using **multiple Specs**:

- List each file loaded and its **Requirement** headline before generating
- Resolve conflicts between Specs through Socratic dialogue — do not silently merge contradictory requirements
- If Specs describe independent deliverables, generate each in a sensible order and report per deliverable

**Do not read `adr/` files for generation.** If a Spec points to ADRs instead of stating decided requirements, stop and load **`spec-writing`** to make the Spec self-contained first.

---

## ASR — Architecturally Significant Requirements

An **ASR (Architecturally Significant Requirement)** is a requirement or design decision that materially affects the structure, organization, or quality of the final deliverable. The term originates in software architecture but applies to **any artifact type** — documents, presentations, images, content series, and more.

ASRs are not minor implementation details. They are the choices that shape the deliverable's fitness for purpose: audience fit, organization, quality bar, constraints, integration, format. When an ASR is unstated, generation fills the gap with a silent default.

Use this framework during the readiness gate when no Spec exists or when a Spec has material gaps.

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

### Domain-specific ASRs (when relevant)

| Domain | ASRs to consider |
|--------|------------------|
| **Software** | Architecture, storage model, API surface, concurrency, error handling, deployment model |
| **Documents & reports** | Tone, depth, evidence standard, citation style, section structure, localization |
| **Presentations** | Narrative arc, slide density, visual style, speaker notes, duration |
| **Images & visual assets** | Style, dimensions, color palette, composition, brand alignment |
| **Content series** | Episode structure, continuity, progression, cross-references |
| **Data & schemas** | Data model, validation rules, serialization format |

Not every ASR category applies to every deliverable. **Select categories relevant to the stated request** — do not run a rigid checklist. Prioritize ASRs where a silent default would materially change the outcome.

### ASR review protocol

When no Spec exists, or when a Spec omits ASRs that would affect quality:

1. **Infer deliverable type** from the user request
2. **Scan universal ASR categories** — flag any that are unstated and architecturally significant
3. **Scan domain-specific ASRs** for the inferred type — same rule
4. **Present findings** to the user as a structured review (not a wall of questions):
   - List each relevant ASR with status: **decided** / **implied** / **open**
   - For **open** ASRs with real alternatives → **`adr-writing`**
   - For **open** ASRs with a natural default → propose briefly, ask for confirmation
5. **Do not generate** until blocking ASRs are resolved and encoded in Spec

When a Spec exists but an ASR is missing, treat it as **open** — surface it Socratically and update the Spec before generating.

---

## Workflow

### 1. Determine Generation Scope

Clarify what the user wants produced and **which Spec(s)** govern it.

Capture briefly:

- **Deliverable** — what artifact type(s) (code, document, image, presentation, …)
- **Spec scope** — one file, named slug, or all specs
- **Trigger** — explicit generation request, or handoff from `spec-writing`?

If the user names a deliverable but not a Spec, search `spec/` for a matching or similar scope (filename, `## Requirement`, `## Tags`) before assuming no Spec exists.

### 2. Spec Readiness Gate

Before generating, every Spec in scope must be **ready**.

**No Spec found:**

1. State that generation requires a Spec
2. Run the **ASR review protocol** (see ASR — Architecturally Significant Requirements):
   - Identify ASRs that would affect deliverable quality
   - Present them to the user as a structured review — decided / implied / open per ASR
   - Use **Socratic dialogue** to resolve **open** ASRs; do not silently default on architecturally significant choices
3. When the user must choose among **viable alternatives**, load and follow **`adr-writing`** — one concern per ADR, user approval required before proceeding
4. After blocking ADRs are **approved**, load and follow **`spec-writing`** to consolidate decisions and requirements into `spec/{requirement-slug}.md` — ensure every ASR reviewed in step 2 is reflected in the Spec
5. Return to step 3 (Load Spec) — do not generate until a Spec exists

**Spec exists but material gaps remain:**

| Condition | Action |
|-----------|--------|
| **Open Questions** that block generation | Ask Socratically; resolve, defer with user awareness, or wait — do not invent answers |
| Spec bullets reference `adr/` without stating the decision | Load **`spec-writing`** to consolidate; do not generate |
| Requirement ambiguous (scope, format, audience, quality bar) | Run ASR review on the gap; Socratic clarification; update Spec if the user confirms |
| ASR missing from Spec (structure, constraints, integration, …) | Surface via ASR review; resolve before generating; update Spec |
| User wants generation despite open blockers | Record explicit user override in dialogue; prefer updating **Open Questions** in the Spec with "proceeding with assumption: …" only after user confirms |

**Do not** skip ADR review when real alternatives exist just because the user said "just build it." Present tradeoffs briefly, recommend if helpful, and obtain approval — per `adr-writing`.

### 3. Load Spec(s)

Read the full content of each Spec in scope.

Extract and keep in working memory:

- **Requirement** — deliverable statement
- **Decisions** — stated outcomes (not ADR pointers)
- **Requirements** — must-haves (testable)
- **Constraints** — limits generation must respect
- **Out of Scope** — explicit non-goals
- **Open Questions** — only non-blocking items may be ignored

Treat the Spec as authoritative over the conversation history. If the user's latest message contradicts the Spec, ask which source to follow before generating.

### 4. Generate

Produce the artifact **strictly from the Spec**:

- Implement every **Requirements** bullet unless blocked — report any that could not be met
- Honor **Decisions**, **Constraints**, and **Out of Scope**
- Do not add features, sections, or dependencies the user did not specify in the Spec
- Do not contradict approved decisions embedded in the Spec
- Use existing project conventions when the Spec does not specify otherwise

When an artifact-specific skill applies, follow its staging and approval gates. The Spec still defines acceptance criteria for the final artifact.

### 5. Report and Hand Off

After generation:

1. Report **output paths** and a concise summary of what was produced
2. Map key **Requirements** bullets to what was delivered (brief checklist)
3. Note any **Open Questions** left unresolved or assumptions used (if user-approved)
4. Recommend **`spec-driven-verification`** — results are saved to `verification/{requirement-slug}.md` for review

Example prompts:

- *"Generated per `spec/item-catalog-api.md`: `src/api/items.ts`, `openapi.yaml`. Run verification? Results will be in `verification/item-catalog-api.md`."*
- *"`docs/ai-education-research-report/document.md` drafted per Spec. Verify against `spec/ai-education-report.md` when ready."*

---

## Multi-Spec Generation

When all specs in `spec/` are in scope:

1. Inventory files: list `spec/*.md` with one-line Requirement each
2. Identify dependencies (e.g. API spec before client spec) — ask if unclear
3. Generate per Spec in dependency order
4. Report per Spec — do not blend unrelated deliverables into one artifact unless a Spec explicitly requires it

If two Specs conflict, **stop** and resolve via dialogue or Spec update — do not pick a silent winner.

---

## Socratic Dialogue (When Spec Is Missing or Incomplete)

Use when step 2 blocks on missing or unclear specification. Complements the **ASR review protocol** — dialogue resolves individual **open** ASRs; the protocol ensures none are overlooked.

**Ask when a universal or domain-specific ASR is open:**

- Purpose or audience is unstated
- Deliverable form (type, format, output path) is unclear
- Scope boundaries are vague or the request mixes multiple independent deliverables
- Quality bar is undefined (*"simple"*, *"modern"*, *"good enough"*, *"professional"*)
- Structure or organization is unspecified and would materially affect the result
- Integration points or dependencies are unknown
- Constraints (language, brand, deadline, tooling) are unstated
- A real choice exists among viable approaches (→ **adr-writing**)

**Dialogue rules:**

1. **State the current step** (ASR review → readiness gate → ADR → spec → load spec → generate)
2. Present the structured ASR review first when no Spec exists — then drill into **open** ASRs
3. One question at a time when possible
4. Prefer questions that help the user discover their answer over defaulting silently
5. Do not invent requirements — document only what the user states or confirms
6. When alternatives matter, **`adr-writing`** before **`spec-writing`**, then return here
7. Confirm that resolved ASRs will be encoded in Spec before generation proceeds

---

## Prohibitions

- Generating final artifacts **without** a ready Spec when blocking decisions or open questions remain unresolved
- Reading **`adr/`** files as a substitute for Spec content during generation
- Treating the original prompt as authoritative when it **contradicts** the Spec
- Adding requirements, features, or scope not in the Spec
- Skipping **`adr-writing`** when the user must choose among real alternatives
- Silently merging conflicting Specs
- Marking generation complete without reporting output paths
- Replacing an artifact-specific skill's staged approval gates unless the user explicitly wants a one-shot draft
- Defaulting silently on ASRs (audience, structure, quality bar, constraints, …)
- Skipping ASR review when no Spec exists and proceeding straight to generation

---

## Completion Criteria

- When no Spec existed at start: ASRs reviewed with the user and encoded in Spec
- Spec scope determined (one file, named slug, or all specs)
- Every Spec in scope is **self-contained** and ready (or user explicitly approved proceeding with documented assumptions)
- Artifact-specific or generic generation path selected and followed
- Deliverable produced at agreed paths per Spec **Requirements**, **Decisions**, **Constraints**, and **Out of Scope**
- User informed of outputs and how Spec bullets were addressed
- User offered **`spec-driven-verification`** as the recommended next step

This skill owns production from Spec; **`spec-driven-verification`** owns checking the result against the Spec.
