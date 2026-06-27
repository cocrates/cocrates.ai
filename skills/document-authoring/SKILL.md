---
name: document-authoring
description: >-
  Select when the user requests a substantial general Markdown document that
  benefits from progressive elaboration (e.g., multi-section reports, proposals,
  analyses, overviews, reference notes, etc.). Do NOT select for simple
  summaries, information lookups, brief explanations, or raw content dumps —
  those should be written directly as Markdown files without this skill's
  workflow.
compatibility: opencode
metadata:
  agent: cocrates
---

# Document Authoring — General Document Writing Skill

## Core Principles

1. **Snowflake progressive elaboration:** Complete the document in order: outline → structure → section drafting → integration → final review.
2. **User approval gates:** Proceed to the next phase (P1–P5) only after the user has reviewed and approved the current one.
3. **Tailored document structure:** Agree with the user on the best structure for the document's purpose, audience, and message.
   - Persuasion / proposal → introduction–body–conclusion
   - Story / case-driven → narrative arc (setup, development, turn, conclusion)
   - Problem solving → problem–cause–solution–proposal
   - Comparison / analysis → criteria–alternatives–evaluation–conclusion
   - Procedure / how-to → step-by-step sequence
   - Other user-defined structures
4. **Section-by-section drafting and review:** Write and review each section sequentially to ensure quality.
5. **Clarify document purpose:** Before starting, make the document's purpose, audience, and core message explicit.

## Working Directory

Save all intermediate artifacts and the final document under:

```text
docs/{slug}/
  ├── outline.md              # P1: Outline (metadata + table of contents)
  ├── sections.md             # P2: Structure design (key points per section + structure choice)
  ├── sections/               # P3: Section bodies
  │   ├── 01-introduction.md
  │   ├── 02-main-point-1.md
  │   └── ...
  └── {slug}.md               # P4: Integrated final document
```

`{slug}` is a lowercase kebab-case identifier that describes the document topic. If the user does not specify one, Cocrates proposes it and the user approves it.

## Intermediate Artifact Definitions

### outline.md

The document's overall outline. Include:

```markdown
# Document Outline

## Metadata
- Title: (document title)
- Date: (date)
- Version: 0.1
- Audience: (target readers)
- Purpose: (why this document is being written)
- Core message: (what readers should take away)

## Table of Contents
1. (Section 1 title)
2. (Section 2 title)
3. (Section 3 title)
   ...
```

### sections.md

Defines key points per section and the document's structural approach.

```markdown
# Structure Design

## Structural Approach
(Chosen structure and rationale)

## Key Points by Section
### 1. (Section title)
- Key points:
- Content to include:
- Estimated length:

### 2. (Section title)
...
```

### sections/NN-{name}.md

Body content for each section. Include the section number (NN) and name.

```markdown
# (Section title)

(Body content in Markdown)
```

### {slug}.md (final document)

The complete document with all sections merged. Include metadata, table of contents, and full body.

## Phase-by-Phase Procedure

### P1: Outline Design

**Purpose:** Define the document's purpose, scope, audience, core message, and table of contents.

**Procedure:**
1. Confirm the document's purpose and audience from the user's request.
2. Use Socratic questions when needed to clarify ambiguities.
3. Draft `outline.md` and present it to the user.
4. After the user reviews and approves, proceed to P2.
5. If the user requests changes, apply them and present again.

**User review questions:**
- Does the document's purpose match the request?
- Are any sections missing or unnecessary in the table of contents?
- Does the core message reflect the document's purpose well?

**Approval condition:** The user explicitly agrees with phrases such as "approved", "looks good", or "next step".

---

### P2: Structure Design

**Purpose:** Design key points per section and the document's structural approach.

**Procedure:**
1. Agree with the user on the structure best suited to the document's purpose and content.
   - Examples: introduction–body–conclusion, narrative arc, problem–cause–solution–proposal, etc.
   - If the user has no clear preference, Cocrates recommends options and the user chooses.
2. Define key points and content to include for each section.
3. Draft `sections.md` and present it to the user.
4. After the user reviews and approves, proceed to P3.

**User review questions:**
- Is the chosen structure appropriate for the document's purpose?
- Do each section's key points carry the intended message?
- Is the section order and flow logical?
- Are there sections to add or remove?

**Approval condition:** The user explicitly agrees with phrases such as "approved", "looks good", or "next step".

---

### P3: Section Drafting

**Purpose:** Draft and review each section's body sequentially.

**Procedure:**
1. Start with the first section in the order defined in `sections.md`.
2. Draft the section body in Markdown and present it to the user.
3. The user reviews.
   - Approved: proceed to the next section.
   - Revision requested: apply changes and present again.
4. Repeat until all sections are complete.
5. Save each section to `docs/{slug}/sections/NN-{name}.md`.

**File naming convention:**
```
sections/01-introduction.md
sections/02-background.md
sections/03-analysis.md
sections/04-conclusion.md
```

**Section drafting guidelines:**
- Write each section so it stands alone and remains understandable.
- Use Markdown elements as needed: tables, code blocks, quotes, lists, etc.
- Consider flow and connections between sections.
- Adjust section length appropriately for the content (neither unnecessarily long nor short).

**User review questions:**
- Does this section convey its key points well?
- Anything to add, revise, or remove?
- May we proceed to the next section?

**Approval condition:** The user explicitly agrees for that section with phrases such as "approved", "looks good", or "next".

---

### P4: Document Integration

**Purpose:** Merge all sections into one complete document.

**Procedure:**
1. Read all section files in order and integrate them into a single document.
2. Add metadata (title, date, version) and a table of contents at the top.
3. Add transitional sentences where needed so section boundaries read naturally.
4. Save the integrated document to `docs/{slug}/{slug}.md`.
5. Notify the user that integration is complete and guide them to P5.

**Approval condition:** Automatic progression (no separate user approval; final review happens in P5).

---

### P5: Final Review

**Purpose:** Review the complete document and obtain final approval.

**Procedure:**
1. Summarize the integrated final document and present it to the user.
2. Request a full-document review.
3. When the user reviews and gives final approval, document creation is complete.
4. If revisions are requested, return to the relevant section(s) in P3 and continue from there.

**User review questions:**
- Is the document's purpose faithfully reflected in the final document?
- Is the overall flow and structure logical and consistent?
- Are there issues with wording, phrasing, or typos?
- Do you approve this document as final?

**Approval condition:** The user explicitly expresses final approval with phrases such as "approved", "done", or "looks good".

## Conversation Rules

1. Always state the current phase (P1–P5) so the user can track progress.
2. When presenting review questions, do not dump every question at once; start naturally with the one or two most important.
3. Even if the user says "just write it", clarify outline (P1) or structure (P2) first when they are unclear.
4. During section drafting (P3), if the user requests changes to a previous section, return to that section, revise, and present again.
5. Respect the user's preferences and experience; avoid overwhelming them with too many questions.
6. Write the document in the language the user requests (Korean, English, etc.).

## Prohibited Actions

- Starting body drafting (P3) without user approval of P1 (outline) or P2 (structure)
- Moving to the next section without user approval of the current section
- Marking document creation complete without user approval in P5 (final review)
- Leaving intermediate artifacts (`outline.md`, `sections.md`, `sections/*.md`) only in chat without saving them as files
- Suddenly generating software code or other deliverable types mid-document workflow
- Including content the user has not approved
- Using this skill for document types better served by dedicated skills (blog series, slide decks, etc.)

## Completion Criteria

- Final document file `docs/{slug}/{slug}.md` exists
- All phases P1–P5 completed with user final approval
- All intermediate artifacts (`outline.md`, `sections.md`, `sections/*.md`) saved as files
- Document purpose, audience, and core message are clearly reflected in the document
- The user can explain the document's content and structure
