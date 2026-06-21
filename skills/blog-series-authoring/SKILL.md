---
name: blog-series-authoring
description: >-
  Use when planning or writing a blog series, serialized posts, or multipart
  content step by step. Select when the user wants series deliverables such as
  overview.md (series overview), outline.md (episode structure),
  episodes.md (per-episode storyline and design), or episodes/NN-{slug}.md
  (publishable episode drafts) — or when the user wants to design or write
  blog content split across multiple posts rather than a single article.
  Refines work in Snowflake order: overview → outline → episodes.md →
  episode body; saves each intermediate artifact to a file and proceeds
  to the next step only after user review and approval.
compatibility: opencode
metadata:
  agent: cocrates
---

# Blog Series Authoring — Progressive Blog Series Writing Skill

Follow this skill when handling requests to create a **blog series, serialized posts, or multipart content**.

## Core Principles

> **The unexamined blog series are not worth of generating.**

Proceed to the next step only after the user has reviewed sufficiently and confirmed the intended direction.

- **Do not generate the full series in one pass.** Save an intermediate artifact to a file at each step and proceed to the next step only after **user review and approval**.
- **Architecture-driven**: Lock overview, series structure, and episode design before writing episode bodies.
- **Snowflake method**: Expand progressively from a one-line logline → series structure → episode design → episode body.

## Working Directory

Use a dedicated folder per series. Use a path agreed with the user; if unspecified:

```
blog/{series-slug}/
```

`{series-slug}`: Convert the series title to kebab-case (e.g. `rust-async-deep-dive`).

---

## Intermediate Artifacts

| Step | File | Content |
|------|------|---------|
| 1. Overview | `overview.md` | Series title, audience, logline, series arc |
| 2. Structure | `outline.md` | Episode composition, order, each episode's role |
| 3. Episode design | `episodes.md` | Storyline, hook, and CTA for all episodes |
| 4. Episode writing | `episodes/NN-{slug}.md` | Episode body draft (individual publication) |

### File Tree

```
blog/{series-slug}/
├── overview.md
├── outline.md
├── episodes.md
└── episodes/
    ├── 01-{slug}.md
    ├── 02-{slug}.md
    └── ...
```

### Naming Rules

- **Number prefix** (`01-`, `02-`, …): Must match episode order in `outline.md` and `episodes.md`
- **slug**: English kebab-case (e.g. `why-async`, `tokio-runtime`)
- **`episodes.md`**: **Design and storyline** for all episodes (not the body text)
- **`episodes/NN-{slug}.md`**: **Publishable body** for that episode

`outline.md` serves as the series **architecture** (structure).

---

## Step-by-Step Procedure

### Step 1 — Overview (`overview.md`)

**Purpose:** Lock the overall direction of the series and the value for readers.

Include the following in `overview.md`:

```markdown
# {Series Title}

## Audience
- Who reads this, prerequisite knowledge, reading goal

## Logline
- Express this series in one sentence

## Series Arc
- Opening → development → climax → close: what the reader gains by finishing the series

## Summary
- 3–5 sentences: core message, scope, expected outcome

## Constraints & Context (optional)
- Episode count, length, publishing cadence, platform, tone, references
```

**Gate:** Save the file, then ask the user to review. Do **not** proceed to Step 2 before approval.

Example prompt: *"Please review overview.md. Do the title, audience, logline, and series arc match your intent? Tell me what to change and I will update it."*

---

### Step 2 — Structure (`outline.md`)

**Purpose:** Design the episode composition and logical flow for the full series.

Include the following in `outline.md`:

```markdown
# Series Structure

## Structure Type
- (e.g. intro → deep dive → practice / problem → cause → solution → retrospective / concept → case study → application)

## Episode List

### 01. {Episode Title} (`01-{slug}`)
- **Role:** What this episode does in the series
- **Core claim:** One sentence
- **Reader takeaway:** What the reader knows after reading this installment
- **Links to previous/next episode:** cliffhanger, callback, foreshadowing

### 02. ...
```

**Gate:** Save `outline.md` → user review → proceed to Step 3 after approval.

---

### Step 3 — Episode Design (`episodes.md`)

**Purpose:** Flesh out the **storyline** for every episode from `outline.md` in one file. Lock logical flow, key points, hook, and CTA before writing bodies.

Include the following section per episode in `episodes.md`:

```markdown
# Episode Design

## 01. {Episode Title} (`01-{slug}`)

### Meta (optional)
- Expected length, publish order, SEO keywords

### Hook
- Question, situation, or claim that grabs the reader in the opening paragraph

### Storyline
- Logical flow for this episode (bullets)

### Core Message
- The one thing the reader should take away from this installment

### Include
- [ ] Point 1
- [ ] Point 2

### Exclude
- Out of scope or deferred to another episode

### CTA (Call to Action)
- Next-episode teaser, comment prompt, related links, etc.

### Tone & Style Notes (optional)

---

## 02. {Episode Title} (`02-{slug}`)
...
```

**How to proceed:**
- If there are many episodes, you may add only **1–2 episodes' worth** to `episodes.md` at a time, get review, then continue.
- Proceed to **Step 4** only after every episode from `outline.md` is designed in `episodes.md` and the user has approved.

**Gate:** Save `episodes.md` (or add a batch) → user review → proceed to Step 4 after approval.

---

### Step 4 — Episode Writing (`episodes/NN-{slug}.md`)

**Purpose:** Generate the episode **body** from the approved section in `episodes.md`.

- Write only the publishable body in `episodes/{nn}-{slug}.md` (keep design content in `episodes.md`).
- By default, generate and review **one episode at a time**.
- Do not add content that conflicts with `episodes.md`.
- Maintain terminology and narrative consistency with prior episodes.

**Gate:** Save per episode → user review → apply edits → next episode.

Example prompt: *"Please review episodes/01-why-async.md. Does it match the storyline? Point to any paragraphs you want changed."*

When all episode bodies are approved, the series is complete. Each file is an individual publication unit; do not create a separate combined edition.

---

## Reverting to Earlier Steps

When needed, revise an earlier artifact and redo downstream steps. Examples:
- If episode composition is wrong, revise `outline.md`, then rewrite `episodes.md` and affected episode drafts
- If series direction changes, re-review from `overview.md`
- If only one episode's design changes, revise that section in `episodes.md`, then rewrite the corresponding `episodes/NN-{slug}.md`

---

## Conversation Rules

1. **State the current step.** (e.g. *"Currently on Step 2 — writing outline.md."*)
2. After saving a file, share the **path and review points**.
3. **Do not create the next step's file without user approval.**
4. If structure or storyline is **ambiguous**, ask before writing `episodes.md`.
5. Even if the user says "just write it," do **not** generate all episode bodies without at least **overview + outline approval**.
6. When requesting review, give **specific checkpoints**. (e.g. episode order, series arc, hook/CTA)

## Prohibited Actions

- Batch-generating `episodes/NN-{slug}.md` without approval
- Writing episode bodies without `episodes.md`
- Creating a combined edition (`series.md`, etc.) — episodes are published individually
- Outputting intermediate artifacts only in chat without saving to files
- Skipping user review steps

## Completion Criteria

- `overview.md`, `outline.md`, and `episodes.md` are approved
- All `episodes/NN-{slug}.md` files are reviewed and revised as needed

When complete, offer a brief retrospective: *"What was the first design decision that changed during this series?"*
