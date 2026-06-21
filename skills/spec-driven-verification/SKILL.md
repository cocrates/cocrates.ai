---
name: spec-driven-verification
description: >-
  Verifies generated deliverables against Spec documents in spec/. Select when
  the user asks to verify, validate, check, or confirm that an artifact matches
  the specification; after spec-driven-generation completes; or when reviewing
  whether output was produced according to the spec. Performs item-by-item
  comparison — not a cursory pass — and surfaces spec deviations plus
  undocumented ASRs (Architecturally Significant Requirements) that should
  have been reviewed during specification.
compatibility: opencode
metadata:
  agent: cocrates
---

# Spec-Driven Verification — Artifact Compliance Skill

This skill **checks generated artifacts against Spec documents**. The Spec is the **sole source of truth** for what should have been built — not the original user prompt, not ADR files, and not assumptions made during generation.

Deliverables include **software, documents, images, presentations, blog series, reports**, and any other finished work the Spec describes.

## Core Principle: Verification Beyond Compliance

Verification is not only checking whether the artifact matches explicit Spec bullets. It also discovers **undocumented ASRs (Architecturally Significant Requirements)** embodied in the deliverable — retrospective findings of *"we should have decided this before building."*

An **ASR** is a requirement or design decision that materially affects the structure, organization, or quality of the final deliverable. The term originates in software architecture but applies to **any artifact type**. Unstated ASRs become visible only when reading the finished artifact.

Surfacing undocumented ASRs is a primary output of this skill, alongside pass/fail against the Spec.

**Spec vs ADR vs verification:**

| | ADR | Spec | Verification (this skill) |
|---|-----|------|---------------------------|
| **Unit** | One concern (one question) | One deliverable or requirement scope | One or more artifacts per Spec |
| **Purpose** | Compare alternatives; record approval | State decided requirements self-contained | **Confirm** the artifact matches the Spec |
| **Reader needs** | Understand options and tradeoffs | Understand what to build | Know what passed, failed, or was decided during implementation |
| **Role in pipeline** | Decision audit trail | **Sole criterion** for verification | Audit |

Match the user's language (Korean, English, etc.) in user-facing messages.

## Spec Input Scope

Verification reads from the project **`spec/`** directory at the repository root.

```
spec/
└── {requirement-slug}.md
```

### Which Spec(s) to Use

Determine scope from the user's request:

| Request pattern | Action |
|-----------------|--------|
| Named file or slug (e.g. `spec/item-catalog-api.md`, `item-catalog-api`) | Verify against **that Spec only** |
| "All specs", "everything in spec/", or cross-cutting verification | Verify against **all** `spec/*.md` files |
| Unspecified | Ask one Socratic question: single deliverable vs multiple — then use the matching set |

When verifying **multiple Specs**:

- List each Spec and its corresponding artifact path(s) before verifying
- Verify per Spec — do not blend unrelated deliverables into one checklist
- If artifact paths are unclear, ask before assuming

**Do not read `adr/` files for verification.** If a Spec bullet references ADR content without stating the decided requirement, flag it as a **Spec gap** and recommend **`spec-writing`** to make the Spec self-contained before re-verifying.

## Output Location

Write every verification report to the project **`verification/`** directory at the repository root.

```
verification/
└── {requirement-slug}.md
```

- **Verification root:** `verification/` at the project root (create if missing)
- **Filename:** same **kebab-case slug** as the governing Spec (e.g. `spec/item-catalog-api.md` → `verification/item-catalog-api.md`)
- One file = one Spec / one deliverable scope

### Before Writing

- If `verification/{requirement-slug}.md` already exists, **read it first** — preserve any **User Review** notes the user added; update findings and evidence without discarding user decisions
- When re-verifying after artifact fixes, update the file in place and note what changed since the prior run

---

## ASR — Architecturally Significant Requirements

An **ASR (Architecturally Significant Requirement)** is a requirement or design decision that materially affects the structure, organization, or quality of the final deliverable. The term originates in software architecture but applies to **any artifact type** — documents, presentations, images, content series, and more.

This section provides a **retrospective framework** for finding ASRs embodied in the artifact that the Spec never addressed. Use in step 7 after item-by-item Spec verification.

These mirror the ASR categories from **`spec-driven-generation`** but read the artifact **backwards**: what did generation decide that the Spec never stated as an ASR?

### Universal ASR categories (always scan)

| ASR category | What to look for in the artifact | Examples across artifact types |
|--------------|----------------------------------|--------------------------------|
| **Purpose & audience** | Implicit target reader or user assumed | Technical depth chosen for experts vs beginners; formality level |
| **Deliverable form** | Format, layout, or medium choices not in Spec | File structure; export format; slide vs document |
| **Scope boundary** | Material included or excluded without Spec guidance | Extra sections; features beyond stated scope; missing promised depth |
| **Quality bar** | Standards applied but not specified | Evidence rigor; visual polish; completeness threshold |
| **Constraints** | Limits respected or violated silently | Language; brand; length; deadline-driven shortcuts |
| **Structure & organization** | Arrangement chosen without Spec guidance | Section order; narrative arc; functional grouping |
| **Integration & dependencies** | Connections to other work assumed or invented | References to prior episodes; reuse of templates; coupling to external assets |
| **Process & staging** | Delivery assumptions not stated | Draft vs final polish; iterative placeholders left in |

### Domain-specific ASRs (when relevant)

| Domain | Undocumented ASRs to watch for |
|--------|-------------------------------|
| **Software** | Architecture, data model, interface shape, concurrency, error handling, deployment assumptions |
| **Documents & reports** | Tone, depth, evidence standard, citation style, section structure, localization |
| **Presentations** | Narrative arc, slide density, visual style, speaker notes, duration |
| **Images & visual assets** | Style, dimensions, color palette, composition, brand alignment |
| **Content series** | Episode structure, continuity, progression, cross-references |
| **Data & schemas** | Data model, validation rules, serialization format |

Not every ASR category applies to every deliverable. **Select categories relevant to the artifact type** — do not run a rigid checklist. Prioritize ASRs where disagreement would materially change quality or fitness for purpose.

### ASR gap scan protocol

After item-by-item Spec verification (step 5):

1. **Infer artifact type** from the deliverable and governing Spec
2. **Scan universal ASR categories** — identify embodied decisions the Spec never addressed
3. **Scan domain-specific ASRs** for the inferred type — same rule
4. **Filter** — keep only findings that are architecturally significant (a silent default would matter if wrong)
5. **Record each finding** with: what was decided, where in the artifact, which ASR category, spec gap, risk, recommended action
6. **Present to the user** for confirmation or correction — do not treat silent choices as approved

A finding is **not** a deviation (artifact contradicts Spec). It is an **ASR gap**: the artifact reveals an architecturally significant choice that should have been stated in the Spec.

---

## Workflow

### 1. Determine Verification Scope

Clarify what to verify and **which Spec(s)** govern it.

Capture briefly:

- **Artifact(s)** — file paths or deliverable identifiers to check
- **Spec scope** — one file, named slug, or all specs
- **Trigger** — explicit verification request, or handoff from `spec-driven-generation`?

If the user names an artifact but not a Spec, search `spec/` for a matching scope (filename, `## Requirement`, `## Tags`) before assuming no Spec exists.

If no artifact path is given, infer from recent generation output or ask the user.

### 2. Load Spec(s)

Read the full content of each Spec in scope.

Extract every verifiable item into a working checklist:

| Spec section | What to extract |
|--------------|-----------------|
| **Requirement** | Deliverable statement — overall scope and success criteria |
| **Decisions** | Each stated outcome (one checklist item per bullet) |
| **Requirements** | Each must-have bullet (testable) |
| **Constraints** | Each limit the artifact must respect |
| **Out of Scope** | Each explicit non-goal — verify the artifact does **not** include it |
| **Open Questions** | Note as non-verifiable unless generation documented an assumption |

Treat the Spec as authoritative. Do not verify against conversation history or ADR files.

### 3. Load Artifact(s)

Read the full generated deliverable(s) at the identified paths.

Read whatever the Spec's **Requirements** imply must exist — source files, document sections, slide decks, images, referenced assets, or other produced content.

Do not skim. Verification requires reading enough of the artifact to judge each Spec item and to scan for undocumented ASRs.

### 4. Build the Verification Inventory

**List every Spec item** used for this verification before checking any of them.

Present the inventory to working memory (and optionally to the user at the start of the report):

```
Verification inventory for spec/{requirement-slug}.md:
- [Requirement] {one-line summary}
- [Decision] {bullet text}
- [Requirement] {bullet text}
- [Constraint] {bullet text}
- [Out of Scope] {bullet text}
...
```

One Spec bullet = one inventory row. Do not merge or collapse items.

### 5. Verify Item by Item

**Check each inventory item individually.** Do not batch-pass sections or mark "looks good" without evidence.

For each item, record:

| Field | Content |
|-------|---------|
| **Spec item** | The exact bullet or statement being verified |
| **Status** | `pass` \| `fail` \| `partial` \| `not-verifiable` |
| **Evidence** | Where in the artifact this was confirmed or contradicted (path, section, line, behavior) |
| **Notes** | Only when status is not `pass` — what differs and why it matters |

**Verification rules:**

- **Pass** only when the artifact clearly satisfies the Spec item — cite evidence
- **Fail** when the artifact contradicts, omits, or materially differs from the Spec item
- **Partial** when some but not all of a compound requirement is met — state what is missing
- **Not-verifiable** when the Spec item is ambiguous, untestable, or blocked by **Open Questions** — do not silently pass

**Out of Scope** items: verify the artifact does **not** include the excluded feature or content. Presence of out-of-scope material is a **fail**.

**Do not** perform a cursory or summary-only check. Every inventory item must have an explicit status and evidence.

### 6. Find Spec Deviations

After item-by-item verification, consolidate all items with status `fail` or `partial`.

For each deviation:

1. Quote or paraphrase the **Spec item**
2. State what the **artifact actually does** instead
3. Classify severity:
   - **Critical** — violates a must-have **Requirement** or **Decision**
   - **Major** — violates a **Constraint** or partially meets a must-have
   - **Minor** — cosmetic or non-functional difference not explicitly forbidden

Deviations are the primary compliance output. Finding them is more important than reporting passes.

### 7. Find Undocumented ASRs

During generation, architecturally significant choices may be embodied in the artifact **without Spec guidance**. These bypassed the design stage — retrospective ASR gaps of *"we should have stated this in the Spec before building."*

Run the **ASR gap scan protocol** (see ASR — Architecturally Significant Requirements). Examine the artifact for ASRs the Spec never addressed.

**Look for decisions that:**

- Are architecturally significant — materially affect structure, organization, or quality
- Were not stated in **Decisions**, **Requirements**, **Constraints**, or user-approved **Open Questions**
- Would have warranted review during specification (or ADR when real alternatives existed)
- Are visible only after reading the finished artifact

**Examples across artifact types:**

- Structure or organization chosen without Spec guidance (section order, narrative arc, functional layout)
- Implicit audience or tone assumptions (depth, formality, technical level)
- Quality bar applied but not specified (evidence rigor, visual polish, completeness)
- Integration or dependency choices not mentioned in Spec (references, templates, external assets)
- Scope material added or omitted beyond what Spec states
- Format, style, or presentation choices invented during production

For each undocumented ASR:

| Field | Content |
|-------|---------|
| **What was decided** | Concrete description of the choice embodied in the artifact |
| **Where** | Location in the artifact (path, section, element) |
| **ASR category** | Which universal or domain-specific ASR category |
| **Spec gap** | Which Spec section lacks guidance (or which **Open Question** was silently resolved) |
| **Risk** | What could go wrong if the user disagrees — impact on quality or fitness for purpose |
| **Recommendation** | `confirm` (accept as-is) \| `update-spec` \| `adr-then-spec` \| `fix-artifact` |

Present these to the user for **confirmation or correction**. Do not treat silent production choices as approved.

**Distinguish from deviations:**

| | Deviation | Undocumented ASR |
|---|-----------|------------------|
| **Relationship to Spec** | Artifact contradicts or fails a stated Spec item | Artifact is consistent with Spec but reveals an unstated ASR |
| **Typical action** | Fix artifact to match Spec | Confirm, update Spec, or fix if user rejects the choice |

### 8. Write Verification Report

Save the full report to **`verification/{requirement-slug}.md`** using the template below. The chat summary is brief; the file is the **durable review artifact**.

After saving, report the **file path** and a one-line summary (pass/fail counts). Invite the user to review the file before acting on findings.

**File template:**

```markdown
# Verification: {Requirement Title}

**Spec:** spec/{requirement-slug}.md
**Artifact(s):** {paths}
**Verified:** {ISO date}
**Summary:** {N} pass, {N} fail, {N} partial, {N} not-verifiable

## Inventory
{numbered list of every Spec item checked}

## Results

| # | Spec item | Status | Evidence |
|---|-----------|--------|----------|
| 1 | ... | pass | ... |
| 2 | ... | fail | ... |

## Deviations

{each fail/partial — Spec item, what the artifact does instead, severity (Critical / Major / Minor)}

## Undocumented ASRs

{each unstated ASR — what was decided, where, ASR category, spec gap, risk, recommended action}

## Recommended Next Steps

- {fix artifact / update spec / adr-writing then spec-writing / re-generate / re-verify}

## User Review

{Leave empty on first write. User may edit this section directly to record decisions.}

<!--
  User: note accept/reject/fix decisions here, e.g.:
  - Deviation #2: fix artifact — scope item required
  - Undocumented ASR "narrative structure": confirm as-is
  - Deviation #5: update spec — additional section now in scope
-->
```

### 9. User Review and Follow-Up

The verification file exists so the user can **review findings at their own pace** and then request changes to the final deliverable, the Spec, or both.

**After writing the file, invite review:**

- *"`verification/ai-education-report.md` is ready. Review the file, add notes under **User Review** if helpful, then tell me what to fix or update."*

**When the user responds after reviewing the file:**

1. Read `verification/{requirement-slug}.md` — including any **User Review** edits
2. Interpret their request:

| User intent | Action |
|-------------|--------|
| Fix specific deviations in the artifact | Apply fixes or run **`spec-driven-generation`** scoped to listed items |
| Accept a deviation; Spec should change | Load **`spec-writing`** to update the Spec, then re-verify |
| Reject an undocumented ASR | Fix the artifact to match Spec intent or run **`adr-writing`** then **`spec-writing`** if a real alternative exists |
| Confirm undocumented ASRs | Load **`spec-writing`** to record accepted ASRs in the Spec |
| Re-run verification after fixes | Return to step 3 (Load Artifact) and update the verification file |
| Verification looks good; no changes | Mark complete; no further action unless user asks |

**Do not** apply artifact changes silently before the user has had a chance to review the verification file — unless they explicitly asked to verify and fix in one pass.

Example prompts:

- *"`verification/ai-education-report.md` — 3 failures. Review the file, then tell me which deviations to fix."*
- *"You noted in User Review to fix the missing bibliography section. Updating the artifact now."*
- *"All items pass in `verification/report.md`. Deliverable verified."*

---

## Multi-Spec Verification

When all specs in `spec/` are in scope:

1. Inventory files: list `spec/*.md` with one-line Requirement each
2. Map each Spec to its artifact path(s) — ask if unclear
3. Run steps 2–9 **per Spec** — separate inventory and `verification/{requirement-slug}.md` per deliverable
4. Summarize across Specs in chat only after per-Spec files are written

---

## Socratic Dialogue (When Scope Is Unclear)

Use when artifact paths, Spec scope, or pass/fail criteria are ambiguous.

**Ask when:**

- Which artifact to verify is unstated
- Multiple Specs could apply to the same deliverable
- A Spec bullet is not testable against the artifact type
- User wants to accept a deviation — confirm whether Spec should be updated instead

**Dialogue rules:**

1. **State the current step** (scope → load spec → load artifact → inventory → verify → deviations → undocumented ASRs → write file → user review)
2. One question at a time when possible
3. Do not mark items `pass` to avoid uncomfortable findings
4. When the user accepts an undocumented ASR, recommend **`spec-writing`** to record it — do not leave the Spec stale

---

## Prohibitions

- Verifying against **`adr/`** or conversation history instead of the Spec
- Skipping the inventory — checking without listing every Spec item first
- Cursory or summary-only verification ("looks fine", "mostly matches") without per-item evidence
- Marking items `pass` without citing artifact evidence
- Ignoring **Out of Scope** violations
- Treating undocumented ASRs as approved without user review
- Skipping the ASR gap scan — reporting only Spec bullet pass/fail without examining architecturally significant choices
- Blending multiple Specs into one undifferentiated checklist
- Reporting verification complete while critical deviations or unreviewed undocumented ASRs remain
- Delivering only a chat summary without saving **`verification/{requirement-slug}.md`**
- Overwriting user edits in the **User Review** section without reading the current file first
- Applying artifact fixes before the user reviews the verification file (unless they asked for verify-and-fix in one pass)

---

## Completion Criteria

- Verification scope determined (artifact path(s), Spec file(s))
- Full Spec and artifact content loaded
- **Every** Spec item listed in the verification inventory
- **Each** inventory item checked individually with status and evidence
- All spec deviations (`fail`, `partial`) documented with severity
- ASR gap scan completed; all undocumented ASRs surfaced for user review
- Report saved to **`verification/{requirement-slug}.md`**
- User informed of the file path and invited to review before artifact or Spec changes
- On user follow-up: **User Review** section read and requested actions applied (fix artifact, update Spec, or re-verify)

This skill owns checking artifacts against Spec; **`spec-writing`** owns updating the Spec when gaps or accepted deviations are recorded; **`spec-driven-generation`** owns producing corrected artifacts.
