---
name: document-authoring
description: >-
  Creates general documents — reports, proposals, papers, blog posts, memos,
  emails, resumes, and any other written content. Select when the user asks to
  write, draft, compose, or author a document, report, proposal, paper, essay,
  memo, email, blog post, article, or any other written deliverable. This is the
  most general document-authoring skill — use it for any artifact whose primary
  form is prose or structured text.
compatibility: opencode
metadata:
  agent: cocrates
---

# Document Authoring — General Document Generation Skill

This skill **produces complete documents** through a structured multi-stage pipeline with independent cross-cutting research. Each stage produces intermediate artifacts as files, and each stage requires user review and explicit approval before proceeding.

## Core Principles

### Document as Architecture

A well-written document is not a wall of text — it is an **architectural artifact** with:

- **Purpose**: a clear reason to exist, defined before a word is written
- **Audience**: a specific reader whose needs drive every structural decision
- **Structure**: a hierarchy of sections that serves the purpose and guides the reader
- **Content**: each section delivering on its design intent, no more and no less
- **Coherence**: the whole reads as one work, not a patchwork of parts

### Progressive Refinement (Snowflake)

Do not jump from blank page to final document. Refine progressively through locked stages:

```
                            research (independent, any stage)
                                   ↕        ↕        ↕
overview.md  →  outline.md  →  sections.md  →  sections/*.md  →  {title}.md
                                                                       ↓
                                    q-and-a.md  →  validation.md  ←  (revision loop)
```

Each arrow represents an approved gate — no stage begins until the previous artifact is reviewed and accepted.

### Purpose-First Structure

Structure follows purpose. Do not impose a generic template. Analyze the document's **goal, audience, and context**, then design a custom structure that serves them. The same request type (e.g., "write a proposal") may produce very different structures depending on audience and purpose.

Match the user's language (Korean, English, etc.) in all user-facing messages and artifact content.

## Resolve Project Root

Before writing artifacts, resolve **`{project-root}`** — the folder that holds this document.

Workspace layouts fall into three types. **Inspect the workspace first**, then match:

| Type | When | `{project-root}` |
|------|------|------------------|
| **1** | Workspace *is* the single document project | `.` (workspace root) |
| **2** | Workspace holds multiple peer projects | `{title-slug}/` |
| **3** | Workspace groups projects by kind | `documents/{title-slug}/` |

**Examples:**
- Type 1 → `overview.md` at workspace root
- Type 2 → `{title-slug}/overview.md`
- Type 3 → `documents/{title-slug}/overview.md`

**Rules:**
1. Infer the type from existing structure (e.g. `overview.md` at root, peer project folders, or a `docs/` kind folder). Prefer an **existing** matching folder over creating a new one.
2. **Before creating** a new project folder, confirm **location and name** with the user.
3. If the project folder already exists, use it — do not recreate or relocate silently.
4. **`{title-slug}`:** URL-friendly slug from the document title (e.g. `ai-education-report`). Type 1 has no slug folder.

## Working Location

Documents live under `{project-root}/`:

```
{project-root}/
├── overview.md                     # define — document definition
├── outline.md                      # plan — document structure overview
├── docs/                           # research — independent, any stage
│   ├── {topic-slug}.md            # per-topic research files
│   └── ...
├── sections.md                     # architecture design — section/sub-section design
├── sections/
│   ├── 01-{section-slug}.md       # detail design — section content
│   ├── 02-{section-slug}/
│   │   ├── 02-01-{sub-slug}.md   # sub-section files for large documents
│   │   └── ...
│   └── ...
├── {title-slug}.md                 # generation — final integrated document
├── q-and-a.md                      # validation — reader Q&A and supplementation notes
└── validation.md                   # validation — purpose-fit check against overview.md
```

**File structure rules by document scale:**

| Scale | Section files | Sub-section handling |
|-------|--------------|---------------------|
| Small (~5p) | `sections/01-intro.md` (flat) | Markdown headings (`###`, `####`) within section file |
| Medium (~20p) | `sections/02-bg.md` (flat) | Sub-section titles defined in `sections.md`; `###` headings inside |
| Large (20p+) | `sections/03-methodology/03-01-design.md` (directory) | One file per sub-section, numbered with `03-01-`, `03-02-` prefixes |

Prefix sections and sub-sections with two-digit numbers (`01`, `02` … `99`) to preserve order.

**Naming:**
- **`{title-slug}`**: URL-friendly slug from document title (e.g., `ai-education-report`)
- **`{section-slug}`**: URL-friendly slug for each section (e.g., `introduction`)
- **`{sub-slug}`**: URL-friendly slug for each sub-section (e.g., `research-design`)
- **`{title}.md`**: Final filename matching document title (e.g., `AI Education Report.md`)

## Research — Independent, Cross-Cutting Activity

Research is **not a fixed stage**. It is an independent activity that can be initiated at **any point** in the process — define, plan, architecture design, detail design, generation, or revision. Findings may feed back into any earlier artifact.

### When Research May Be Needed

| Stage | Situation |
|-------|-----------|
| **Define** | Background knowledge is insufficient — "I need to understand recent trends before defining this report." |
| **Plan** | Examples, case studies, or standard formats are needed to decide structure. |
| **Architecture design** | Specific data or evidence must be gathered to define what each section covers. |
| **Detail design** | While writing a section, additional facts or references become necessary. |
| **Validation/Revision** | Validation reveals weak evidence, purpose misalignment, or factual gaps that need reinforcement. |

### How Research Works

1. **Request**: The user asks for research ("Look into this topic"), or Cocrates proposes it ("This section needs recent data — shall I research it?").
2. **Perform**: Gather information via web search, user-provided materials, or internal sources.
3. **Record**: Per topic, write a separate file `{project-root}/docs/{topic-slug}.md` containing:
   - Summarized findings with source attribution
   - Key facts, data points, and terminology
   - Open questions requiring user input
4. **Review**: Present findings to the user for confirmation.
5. **Feed back**: Research findings may update any artifact:
   - Affects document definition → update `overview.md`
   - Affects structure → update `outline.md` or `sections.md`
   - Affects content → update the relevant `sections/{n}-{slug}.md`

### Standalone Research

The user may request research independently of document generation:

> *"Research AI education trends and save the results in docs/. I'll write the report later based on those materials."*

In this case, only `{project-root}/docs/*.md` files are created. When document generation is requested later, the Snowflake stages proceed using those materials.

### Ad-Hoc Research During Authoring

During any stage of document authoring, the user may request research or summarization on the fly:

> *"While I'm reviewing this, could you research the latest Korea education policy and save it?"*
> *"Can you summarize the key points from the case studies I mentioned and organize them into a reference doc?"*

When such a request is received:

1. **Determine context**: If a document session is active (`{project-root}/` exists), save research under that document's `docs/` directory: `{project-root}/docs/{topic-slug}.md`. If no document session is active, save at the project level: `docs/{topic-slug}.md`.
2. **Perform and record**: Same process as standard research — gather, synthesize, and save with source attribution.
3. **Notify**: Inform the user that the result was saved and can be referenced in the current or later stages.
4. **No gate required**: Unlike Snowflake stage artifacts, ad-hoc research files are informational and do not require explicit approval gates. The user can reference or ignore them as they see fit.

## Snowflake Stages

```text
define → plan → architecture design → detail design → generation → validation/revision
```

| Stage | Purpose | Gate |
|-------|---------|------|
| **① define** | Lock document purpose, audience, scope, constraints, validation criteria | `overview.md` approved |
| **② plan** | Establish overall structure and narrative direction | `outline.md` approved |
| **③ architecture design** | Detail-design each section and sub-section's role, content, and dependencies | `sections.md` approved |
| **④ detail design** | Write complete content for each section | Section files approved |
| **⑤ generation** | Integrate all sections into a single coherent document | `{title}.md` reviewed |
| **⑥ validation/revision** | Reader Q&A, validate purpose-fit, revise or add appendices | Q&A + Validation → revision → final approval |

**Common gate rule:** At every stage, Cocrates presents the artifact → the user reviews it → explicit approval is required before proceeding. No stage begins until the previous artifact is accepted.

### Per-Stage Component Refinement

| Component | ① define | ② plan | ③ arch. design | ④ detail design | ⑤ generation | ⑥ validation |
|-----------|----------|--------|----------------|----------------|---------------|--------------|
| **Assignment & constraints** | ✅ Locked in `overview.md` | — | — | — | — | May revise |
| **Context & rules** | ◐ See Research section — may be initiated at any stage; findings update any artifact |
| **Structure & flow** | Type assessed | ✅ `outline.md` | ✅ `sections.md` — roles, hierarchy, dependencies | Micro-adjustments | Coherence check | Restructure if needed |
| **Content units** | Approach identified | Roles summarized | Templates decided; granularity assessed | ✅ Section files written | ✅ `{title}.md` integrated | Revise |
| **Review & validation** | Validation criteria set | Approach outlined | Checklist drafted | — | ✅ `q-and-a.md`, `validation.md` | ✅ Revisions applied |

## Stage-by-Stage Procedure

### Stage ① — Define (`overview.md`)

Surface the document's foundation through Socratic dialogue:

| Question | Purpose |
|----------|---------|
| *"What is the purpose of this document? Who will read it, and what should they be able to do after reading?"* | Purpose & audience |
| *"Are there any length or format constraints?"* | Constraints |
| *"Are there topics or content that should be excluded?"* | Out of scope |
| *"How will you confirm the document fulfills its purpose? What must be true for it to be complete?"* | Validation criteria |
| *"Are there reference materials or templates to follow?"* | Context & references |

Record answers into `{project-root}/overview.md`:

```
# Document Definition: {Title}

## Purpose
{Why this document exists}

## Audience
{Primary readers and what they should gain}

## Scope
{What is included}

## Out of Scope
{What is deliberately excluded}

## Validation Criteria
{Checks that confirm the document fulfills its purpose and scope — used in stage ⑥ validation}

## Constraints
{Length, format, deadline, taboos}

## References
{Templates, prior docs, related materials}
```

### Stage ② — Plan (`outline.md`)

Design the document's top-level structure.

1. Load `overview.md` to ground decisions in purpose and audience.
2. Analyze the document type, purpose, and audience to determine the best structure.

   **Structural heuristics (use as inspiration, not prescription):**

   | Document type | Typical structure |
   |---------------|-------------------|
   | **Report** | Executive summary → Background → Methodology → Findings → Analysis → Recommendations |
   | **Proposal** | Problem → Solution → Approach → Timeline → Budget → Qualifications |
   | **Paper / Essay** | Introduction → Thesis → Arguments → Counterarguments → Conclusion |
   | **Blog post** | Hook → Context → Main points → Takeaways → Call to action |
   | **Technical doc** | Overview → Prerequisites → Step-by-step → Troubleshooting → Reference |
   | **Memo** | Context → Key message → Details → Action items |
   | **Email** | Greeting → Context → Action requested → Details → Closing |

   **Always ask:**
   - *"Does this structure serve the document's purpose and audience effectively?"*
   - *"Is the reading order natural for the audience?"*
   - *"Is there a better structural alternative?"*

3. Draft `outline.md` with document title, top-level sections (with role descriptions), and narrative arc. Include an **appendix** (or appendices) in the outline when the document type or scope warrants supplementary material kept out of the main flow.
4. **Gate:** User approves the structural direction.

### Stage ③ — Architecture Design (`sections.md`)

Detail-design each section and its sub-sections.

1. Load `overview.md` and `outline.md`. Research materials from `docs/*.md` may be referenced.
2. For each section, define:

   | Property | Description |
   |----------|-------------|
   | **Section title** | Final heading text |
   | **Purpose** | What this section accomplishes for the reader |
   | **Key message** | The single most important takeaway |
   | **Content elements** | What to include (narrative, data, examples, visuals, citations) |
   | **Sub-sections** | (Optional) List with title, purpose, and key message for each |
   | **Dependencies** | Which sections must be written before this one |
   | **Estimated length** | Rough word count or paragraph count |
   | **Tone / style notes** | Any section-specific voice adjustments |

   **Sub-section design example:**

   ```
   Section 2: Background and Related Work
   - Purpose: Provide academic and industry context
   - Key message: AI education is growing rapidly but lacks standardized approaches
   - Sub-sections:
     - 2-1. History of AI in Education — from early CAL to modern AI tutoring
     - 2-2. AI Education Policy — comparison of Korea, US, and EU approaches
     - 2-3. Key Research Trends — influential studies and meta-analyses (2023–2025)
   ```

3. Write `sections.md` with all section/sub-section designs and dependency graph. When an appendix is planned, design it here like any other section — title, purpose, key message, and content elements — even if its content is written later during validation/revision.
4. **Gate:** User approves section designs before content writing begins.

### Stage ④ — Detail Design (`sections/{n}-{slug}.md`)

Write each section's complete content, including sub-sections.

1. Follow the writing order from the dependency graph in `sections.md`.
2. Choose the file structure based on document scale (see Working Location for naming):

   **Simple (no sub-sections):** Single file `sections/01-{slug}.md`. Use `###`, `####` headings for internal organization.

   **Medium (with sub-sections):** Single file `sections/02-{slug}.md` with `###` headings per sub-section, following the design in `sections.md`.

   **Large (many sub-sections):** Directory `sections/03-{slug}/` with one file per sub-section (`03-01-{sub}.md`, `03-02-{sub}.md`, ...). Each sub-section file should be independently readable.

3. **Consistency checks per section/sub-section:**
   - Does content match the purpose and key message from `sections.md`?
   - Is terminology consistent across sections?
   - Are data and citations accurate and sourced?
   - Is tone consistent with `overview.md` guidelines?

4. **Gate:** User approves section content (individually or in batches).

### Stage ⑤ — Generation (`{title}.md`)

Integrate all sections into one coherent, polished document.

1. Load all approved section files.
2. Integrate:
   - Combine sections in order with smooth transitions
   - Add front matter (title page, table of contents if appropriate)
   - Refine introduction and conclusion for context and synthesis
   - Normalize formatting, heading levels, and terminology
   - Add references section if citations span multiple sections

3. **Coherence check:**
   - Does the document read as a single work, not assembled parts?
   - Is there a consistent authorial voice throughout?
   - Do all sections serve the purpose defined in `overview.md`?
   - Are there gaps or redundancies?

4. Write to `{project-root}/{title-slug}.md`.

### Stage ⑥ — Validation / Revision

Confirm the document fulfills its purpose and identify what still needs supplementation — then revise. This stage is **validation** (does the document meet its defined purpose?), not specification verification.

#### Step 6a — Q&A

Q&A is a **reader-perspective exercise**, not material assembled for the validation report. Its purpose is to surface what a reader might still wonder after reading, summarize answers, and note what should be added to the body or moved to an appendix.

1. Review `{title-slug}.md` from the reader's perspective: what additional questions, points of confusion, or objections might arise?
2. Write `q-and-a.md`:

   ```markdown
   ## Anticipated Questions

   ### Q1: {Question a reader might ask}
   **A:** {Concise answer — from the document, or note if not covered}
   **Supplementation:** {What to add to the main text, or flag for appendix — omit if fully covered}

   ### Q2: {Question}
   **A:** {Concise answer}
   **Supplementation:** {…}
   ```

3. Entries where the main text does not fully answer the question become candidates for revision or appendix content. Q&A does **not** serve as validation evidence — it feeds supplementation decisions that validation may later confirm.

#### Step 6b — Validation

Validation checks whether the document **fulfills the purpose and validation criteria in `overview.md`** and whether anything is still missing. It is an independent review; Q&A is one input among several, not the primary basis for the validation report.

1. Load `overview.md`, `{title-slug}.md`, and (for reference) `q-and-a.md`.
2. Validate against each validation criterion from `overview.md` and general document quality:

   | Dimension | Questions |
   |-----------|-----------|
   | **Purpose fit** | Does the document achieve the purpose and deliver what the audience needs, as defined in `overview.md`? |
   | **Clarity** | Is the main message immediately clear? Are complex ideas explained well? |
   | **Structure** | Does the organization guide the reader logically? Are transitions smooth? |
   | **Accuracy** | Are facts, data, and citations correct? Are arguments sound? |
   | **Completeness** | Does the document deliver on its stated purpose and scope? Are there gaps — including topics flagged in Q&A supplementation notes? |
   | **Conciseness** | Is every paragraph necessary? Is there redundancy? |
   | **Tone** | Is the voice appropriate for the audience and purpose? |

3. Write `validation.md` with:
   - Strengths
   - Weaknesses and gaps relative to `overview.md`
   - Specific revision suggestions by section — what to add, cut, rephrase, or restructure (including items surfaced in Q&A that validation confirms need supplementation)

#### Step 6c — Revision and Appendices

1. Present `q-and-a.md` and `validation.md` to the user.
2. Propose specific revisions based on validation findings. Explain how each revision addresses a purpose-fit gap or supplementation need.
3. **Appendix path:** When Q&A or validation identifies material that belongs outside the main flow — or the user explicitly requests an appendix:
   - Add an appendix entry to `outline.md` (if not already present) and detail it in `sections.md` (purpose, key message, content elements).
   - Write the appendix content as section file(s) under `sections/` (e.g., `sections/99-appendix-a-{slug}.md`), following the same numbering and scale rules as other sections.
   - Regenerate `{title-slug}.md` to include the appendix.
   - Appendices may be planned from stage ② or ③; they do not have to wait until validation.
4. Options: accept as-is, apply specific revisions, add appendices, or restructure (return to stage ③ or ④).
5. Re-run validation if significant changes were made.
6. **Final gate:** User confirms the document is complete.

## Cross-Reference and Reuse Rules

1. **`overview.md` is the anchor**: All later stages reference `overview.md` for purpose, audience, and validation criteria. If the user revises it mid-process, re-check downstream artifacts.
2. **`sections.md` governs content writing**: Section files must match `sections.md`. If writing or research reveals a need to change section design, update `sections.md` first.
3. **Terminology consistency**: When `docs/*.md` defines key terms, all sections must use them consistently. Update the research doc when new terms emerge.
4. **Dependency order**: If section B depends on section A, write A before B. The dependency graph in `sections.md` determines ordering.

## Dialogue Rules

1. **State the current stage** at each interaction (e.g., *"We are in stage ① (define). Let's establish the document's purpose and audience."*).
2. **One question at a time** when surfacing requirements.
3. **Use Socratic questioning**: prefer questions that help the user clarify intent over filling in blanks yourself.
4. **Present, don't dump**: when showing a draft, highlight what changed and what needs attention.
5. **Incorporate prior materials**: if the user has existing outlines, notes, or documents, ask if they should be incorporated before starting fresh.
6. **Adapt to document scale**: for very short documents (<1 page), propose combining stages (e.g., overview + outline + writing in one pass) with user approval.

## Prohibitions

- Generating the final document **without** approved section content
- Writing section content **before** section designs are approved in `sections.md`
- Using a fixed template without considering purpose and audience
- Adding content outside the scope defined in `overview.md`
- Silently fixing validation issues — present proposed revisions to the user
- Skipping validation/revision and declaring the document final
- Leaving intermediate artifacts in chat only — every stage writes to files

## Completion Criteria

- All Snowflake stage gates passed (see Snowflake Stages table)
- `{project-root}/{title}.md` exists as an integrated, coherent final document
- Validation findings and Q&A supplementation needs are addressed (unless the user explicitly accepts as-is)
- User has given explicit final approval
