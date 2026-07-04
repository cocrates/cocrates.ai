---
name: presentation-authoring
description: >-
  Creates slide-based presentations — slide decks, talk decks, lecture slides,
  pitch decks, conference presentations, and any other presentation content.
  Select when the user asks to create, draft, compose, or author a presentation,
  slide deck, talk, pitch deck, lecture slides, or any slide-based deliverable.
  Phase A produces a renderer-independent model (overview → outline → slides.md
  → scripts → templates → slides/slide-*.md, plus Q&A and validation).
  Phase B translates that model into a renderer-specific deck (e.g. Marp).
  When the user requests a rendered deck, read renderer/ in this skill directory.
compatibility: opencode
metadata:
  agent: cocrates
---

# Presentation Authoring

Structured pipeline for slide-based presentations. **Slides** are the unit of composition. All Phase A artifacts are **markdown files**.

## PIM / PSM Model

This skill follows the same separation as **PIM** (Platform Independent Model) and **PSM** (Platform Specific Model):

| Phase | Role | Analogy | Scope |
|-------|------|---------|-------|
| **Phase A** (this file) | **Renderer-independent** presentation authoring | **PIM** | Define → architecture → detail design → generate platform-independent slide model |
| **Phase B** (`renderer/`) | **Renderer-specific** presentation authoring | **PSM** | Translate the PIM into a concrete renderer (Marp, etc.) |

Phase A never embeds Marp syntax, PPTX layout, or other renderer mechanics. Phase B never redesigns narrative or on-slide copy — it **translates** approved `slides/slide-*.md` (and templates) into renderer artifacts.

| Phase | Artifacts (order) | Design level |
|-------|-------------------|--------------|
| **A — PIM** | `overview.md` | define |
| | `outline.md` | architecture design (title-level structure) |
| | `slides.md` | detail design (key message, components) |
| | `scripts/` → `templates/` → `slides/slide-*.md` | generation path (speech → chrome → on-slide layout) |
| **B — PSM** *(on user request)* | PIM page models → per-slide PSM → compiled deck | translate PIM into renderer-specific deck (paths in `renderer/`) |

Match the user's language in messages and artifact content.

## Core Design Ladder (Phase A)

Three artifacts form the design ladder. Do not collapse them:

| Artifact | Design level | Role | Review question |
|----------|--------------|------|-----------------|
| **`outline.md`** | **Architecture design** | Review the **overall structure at title level** — sections and slide titles only | Does the title sequence form a coherent whole? |
| **`slides.md`** | **Detail design** | Design each slide’s **content** — key message, components, purpose | Is the content design sound? |
| **`slides/slide-{nn}.md`** | **Generation (PIM)** | Specify **concrete layout and expression** — what sits where, key phrases, figures/images | Can a renderer implement this page without inventing content? |

### Content generation path

Default order after detail design is approved:

| Order | Artifact | Role |
|-------|----------|------|
| 1 | `slides.md` | Content design (key message, components) |
| 2 | `scripts/script-{nn}.md` | Spoken narrative that delivers those messages and components *(default; ask at entry)* |
| 3 | `templates/` + `Template` on `slides.md` | Shared background and style chrome |
| 4 | `slides/slide-{nn}.md` | On-slide visuals the audience sees |

- **Scripts** expand `slides.md` for the presenter: explain components and key messages fully in speech.
- **`slide-{nn}.md`** is **not** a dump of `slides.md`. Put only what the audience should **see** — **key phrases**, positions, figures/images. Prefer concise on-slide wording over long explanation (that belongs in the script).
- **Templates** are shared across slides (background, style, regions). They are not per-slide content.

## Core Principles

- **Purpose-first** — structure follows goal, audience, and context
- **Time over slide count** — lock **duration** in `overview.md`; slide count is provisional
- **Progressive refinement** — architecture (titles) → content design → scripts → templates → page generation
- **Title-level architecture** — judge the whole deck from `outline.md` titles before detailing content
- **One message per slide** — key phrases on-slide; spoken detail in scripts when used
- **Spatial clarity** — every on-slide element has position and exact visible copy
- **Visual consistency** — shared background and style via **templates**
- **Renderer independence (PIM)** — Phase A describes *what* to show; Phase B decides *how* for a platform

Presenter **scripts** are generated **by default**. At stage ④ entry, ask whether to write scripts or skip to templates and page generation. Q&A and validation use **`slides/`** and **`scripts/`** when scripts exist.

When Phase A is complete and the user asks to **generate, render, or export** a deck, enter **Phase B**: read [`renderer/README.md`](renderer/README.md) and follow the matching renderer. Do not fold PSM translation into Phase A stages.

## Working Location

Default root: `presentations/{title-slug}/` (or a path the user specifies).

| Path | Stage | Role |
|------|-------|------|
| `overview.md` | ① | define |
| `outline.md` | ② | architecture — title-level structure |
| `docs/{topic-slug}.md` | any | research (optional) |
| `slides.md` | ③ / ⑤ | detail design (key message, components); `Template` from ⑤ |
| `scripts/script-{nn}.md` | ④ | spoken delivery of ③ (default; skip only if user declines) |
| `templates/template-{nn}.md` | ⑤ | shared background, style, regions (reusable) |
| `slides/slide-{nn}.md` | ⑥ | concrete layout and on-slide expression (**required**) |
| `q-and-a.md` | ⑦ | expected questions |
| `validation.md` | ⑦ | PIM quality gate |

Phase B adds renderer-specific paths — defined in `renderer/{name}.md`.

**Conventions:**
- Sections live in `outline.md` / `slides.md`, not subfolders — flat numbering in `slides/` and `scripts/`
- Slide numbers: `01`–`99` for the main deck; `B01`, `B02`, … for **backup** slides
- Template numbers: `01`–`99` within `templates/`; `01` is primary/default
- `{title-slug}`: URL-friendly slug from presentation title

**Backup slides (`B01`, …):**
- Optional; list under a Backup section in `outline.md` when useful for Q&A or overflow
- May have `slides/slide-B{nn}.md`, template assignment, and optional scripts
- **Exclude** from duration totals and from Duration fit in validation
- **Not required** for Phase A completeness, scripts (when scripts are on), or Q&A coverage
- May appear as **Related slide** in `q-and-a.md`

## Research

Research is **not a fixed stage** — it can run at any time (before or during the pipeline). Record findings in `docs/{topic-slug}.md`, review with the user, then update affected artifacts.

## Stages

| Stage | Gate artifact | Design level | Required |
|-------|---------------|--------------|----------|
| **① define** | `overview.md` | Requirements | Yes |
| **② architecture design** | `outline.md` | Structure (section name, role, duration + titles) | Yes |
| **③ detail design** | `slides.md` (no `Template` yet) | Key message + components | Yes |
| **④ scripts** | `scripts/script-*.md` | Spoken delivery of ③ | **Default** — ask at entry; skip → ⑤ |
| **⑤ template design** | `templates/template-*.md` + `slides.md` (`Template` per slide) | Shared background / style / regions | Yes |
| **⑥ PIM generation** | `slides/slide-*.md` | Concrete layout + on-slide expression | Yes |
| **⑦ validation** | `q-and-a.md` + `validation.md` | PIM quality gate | Yes |

Each stage ends when its artifact is **reviewed and explicitly approved**. Stage ④ is skipped only when the user explicitly declines scripts at stage entry — do not pre-decide in `overview.md`.

Short decks may combine adjacent stages **with approval**, but never skip architecture (`outline.md`) before detail design, or detail design before PIM generation.

## Pipeline Rules

### Design ladder rules

| Change type | Update first | Then |
|-------------|--------------|------|
| Sectioning, slide order, titles | `outline.md` | `slides.md` → `scripts/` (if any) → `slides/` |
| Key message, components, purpose | `slides.md` | `scripts/` (if any) → `slides/` |
| Spoken narrative | `scripts/script-*.md` | `slides/` if on-slide phrases must change |
| Exact on-slide layout, key phrases, figures | `slides/slide-*.md` | Phase B artifacts if they exist |
| Shared background / style (templates) | `templates/` | Review assigned `slides/` |

### `slides.md` lifecycle

| When | `slides.md` contains |
|------|----------------------|
| Stage ③ | Purpose, **key message**, **components** (labels), optional layout hints, density — **no `Template`**, no final on-slide phrases |
| Stage ⑤ | Same content **plus** per-slide **`Template`** assignment |

### Templates

Templates define **common slide chrome** — what stays consistent across many slides. They are part of the **PIM** (renderer-agnostic plain language) and are **reused**, so they must not contain content for any specific slide.

**What a template defines:**

| Concern | Examples |
|---------|----------|
| **Background** | Color, gradient, image, graphic treatment |
| **Style** | Palette, typography, spacing defaults |
| **Regions** | **Title**, **sub-title**, **header**, **footer**, **body** — positions and default styling |

Page generation (`slide-{nn}.md`) fills **body** (and title text) per slide; the template only defines the shared frame.

**When to define a separate template:**

| Situation | Action |
|-----------|--------|
| Same background and style; same title/header/footer/body structure | Reuse one template |
| Different background, palette, or typography | **New template** (or a derived variant) |
| Different region structure (e.g. no title bar, centered cover title, extra header band) | **New template** |

**Typical set:**

| Template | Role | Usage |
|----------|------|-------|
| **`template-01`** | Cover | Opening — often full-slide title/sub-title, minimal header/footer |
| **`template-02`** | Content | **Default for most slides** — derives from `template-01`; title + body (+ optional header/footer) |
| **`template-03`** | Section divider | Optional — different region emphasis or accent background |

**Invariants:**
- **No current-slide content** in `templates/template-*.md` — no slide numbers, slide titles, key messages, or per-slide copy
- Slide-to-template binding lives only in **`slides.md`** and **`slides/slide-{nn}.md`**
- Templates are **renderer-agnostic** — Phase B maps them to platform mechanisms
- **2–3 templates** cover most decks

**Dependencies (derivation):** templates may depend on another template — “based on A, with these changes”:

| Pattern | Example |
|---------|---------|
| **Extend** | `template-02` inherits background/palette/typography from `template-01`; adds top title + body regions |
| **Variant** | `template-03` — same base as `template-02`, modified background or header for section breaks |
| **Cross-presentation** | `template-01` derives from another deck’s `templates/template-01.md` |

Always record **Base**, **Relationship**, and **Changes from base** (Inherited / Modified / Added / Removed). Specify bases before derivatives. When a base changes, review the derivation chain.

**Assignment values:**

| Value | Resolves to |
|-------|-------------|
| `template-01` | `{deck-root}/templates/template-01.md` |
| `{other-deck}/templates/template-01.md` | External reuse — copy locally only when creating a variant |

### Cross-artifact rules

1. **`overview.md` anchors** purpose, audience, validation criteria, visual style hints — not scripts preference
2. **One script per main-deck slide** when stage ④ is not skipped — scripts deliver `slides.md` key messages and components in speech
3. **`slide-{nn}.md` shows key phrases**, not the full detail design or full script
4. **Consistent terminology** from `docs/*.md`
5. **`slides/slide-{nn}.md` is the canonical PIM** for Phase B — never invent on-slide copy in the renderer

## Visual Design

Three layers — do not collapse them:

| Layer | Where | What |
|-------|-------|------|
| **Template** | `templates/template-*.md` | Shared **background**, **style**, **regions** (title / sub-title / header / footer / body) |
| **Content design** | `slides.md` | Key message and **components** (what the slide is about) |
| **Page expression** | `slides/slide-*.md` | **Concrete layout** — where each element sits, key phrases, figures/images |

### Principles

| Principle | Application |
|-----------|-------------|
| One focal message | One core takeaway per slide |
| Content before chrome | Design messages in `slides.md`; apply templates for consistency |
| Template consistency | Background, style, and regions from the assigned template |
| Spatial clarity | Every on-slide element has position and visible key phrase or figure |
| Prefer brevity on-slide | Key phrases only; long explanation in scripts when used |

### Content layout modes (body region)

Optional hints for how the **body** is arranged. Concrete placement lives in `slide-{nn}.md`.

| Mode | Structure | Use when |
|------|-----------|----------|
| `focal` | Single prominent element | Key sentence, stat, quote |
| `bullets` | Vertical list | Points, steps |
| `columns` | Side-by-side cells | 2–3 parallel blocks |
| `rows` / `stack` | Top-to-bottom blocks | Sequential sections on one slide |
| `grid` | 2×2, 2×3 card grid | Objectives, features, takeaways |
| `table` | Rows and columns | Data comparison |
| `diagram` | Arrows, synthesis, free spatial | Concept maps, A + B → C |
| `mixed` | Combination of modes | Complex slides — describe zones explicitly |

For contrast/comparison, describe **zones and axis** in page expression (horizontal, vertical, focal+periphery, before→after).

### Density

| Value | Meaning |
|-------|---------|
| `normal` | Default spacing and type size |
| `dense` | More items — tighter spacing or smaller type (renderer adjusts) |
| `compact` | Maximum items on one slide — use sparingly |

Many items are valid when **density** is set and the layout mode fits. Prefer split only when readability suffers.

### Spatial composition (page expression)

In `slides/slide-*.md` **Visual Design → Layout**, specify **concrete placement** — what content, figure, or image sits in which region:

```markdown
- **Layout:** {intent}
  - Title / sub-title: …
  - Header: … *(if used)*
  - Body — Top / Center / Bottom / Left / Right: …
  - Footer: … *(if used)*
  - Figures / images: {what and where}
```

This is the **canonical layout in the PIM**. Content layout mode is an optional hint; spatial zones and key phrases carry the detail.

### Checklist (per `slides/slide-*.md`)

1. Key message visible as a **key phrase** (not a paragraph)?
2. Every visible element has a **position**?
3. Figures/images described where used?
4. On-slide text is **shorter** than script / detail-design explanation?
5. **`Template`** matches `slides.md` and regions respect the template?
6. **`Density`** set when item count is high?
7. No slide-specific content leaked into `templates/`?

**Review flags (not automatic rejects):** 7+ items without `dense`/`compact`; elements without position; full script prose on-slide; page model contradicts template without note; multiple unrelated messages without intentional `mixed`.

## Stage Procedures

### ① Define (`overview.md`)

| Question | Purpose |
|----------|---------|
| Purpose, audience, takeaway? | Foundation |
| Format and **duration** (minutes)? | Time budget — not slide count |
| Topics to exclude? | Scope |
| Success criteria? | Validation |
| Branding, colors, fonts, reference decks? | Template hints |

Do **not** lock a target slide count in `overview.md`. Count emerges in stages ②–③. Optional rough range only if the user volunteers one — soft hint, not a gate.

```markdown
# Presentation Definition: {Title}

## Purpose
## Audience
## Format & Duration
{Conference talk, lecture, pitch, … — total minutes}
## Visual Style
{Palette, fonts, mood — informs stage ⑤; reference decks if reusing templates}
## Narrative Goals
## Scope / Out of Scope
## Validation Criteria
## Constraints
{Duration, branding, taboos — not a fixed slide count}
## References
```

**Review gate (approve before stage ②):**

| Check | Fail if |
|-------|---------|
| Purpose | Empty or too vague to guide structure |
| Audience | Missing who they are or what they already know |
| Format & Duration | No total time budget in minutes |
| Narrative Goals | No measurable or checkable outcomes |
| Scope / Out of Scope | In-scope and out-of-scope not distinguished |

### ② Architecture design (`outline.md`)

Load `overview.md`. Produce a **title-level structural blueprint** — review the **overall composition** from section names and slide titles alone. Do not design slide content here.

**Required fields**

| Field | Level | Required | Notes |
|-------|-------|----------|-------|
| Document title | deck | Yes | `# Outline — {Title}` |
| Section name | section | Yes | Clear narrative unit (e.g. Opening, Problem, Solution) |
| Section role | section | Yes | One line: what this section does in the arc |
| Section duration | section | Recommended | Minutes or share of total time budget |
| Slide number | slide | Yes | `01`, `02`, … continuous across sections |
| Slide title | slide | Yes | Audience-facing title only — no purpose, layout, or copy |

**Do not include** in `outline.md`: purpose, key message, components, layout, density, template, speaker cues, or on-slide phrases. Those belong in `slides.md` / `scripts/` / `slides/`.

**Schema:**

```markdown
# Outline — {Title}

## Section 1: {Section name}
- **Role:** {one-line narrative role in the arc}
- **Duration:** {e.g. ~3 min / ~15% of total}

- **Slide 01:** {Title only}
- **Slide 02:** {Title only}

## Section 2: {Section name}
- **Role:** {…}
- **Duration:** {…}

- **Slide 03:** {Title only}

## Backup *(optional — excluded from duration)*
- **Slide B01:** {Title only}
```

**Architecture review (approve before stage ③)** — judge the deck **from titles alone**:

| Check | Fail if |
|-------|---------|
| Arc | No clear beginning → middle → end, or sections could be swapped without loss |
| Section boundaries | Overlapping topics, or a section that mixes unrelated goals |
| Title sequence | Titles do not read as a coherent story when listed alone |
| Coverage | `overview.md` Narrative Goals / in-scope topics missing from any section |
| Scope | Out-of-scope topics appear as slides |
| Time | Section durations (if set) cannot plausibly fit `Format & Duration` |
| Density of structure | One section dominates without justification, or too many one-slide sections |

Slide titles and counts are **drafts** — add, merge, or drop freely in stage ③, then sync `outline.md` when structure changes.

### ③ Detail design (`slides.md`)

Load approved `outline.md`. Design each slide’s **content** — **key message** and **components** — not title-level structure and not final on-slide layout.

Prefer fitting the **time budget** over hitting a slide total. Split crowded slides or merge thin ones as needed; update `outline.md` when titles or order change.

**Do not assign `Template` here** — added in stage ⑤.

**Depth rule:**
- Stage ③: **what** the slide must communicate (purpose, key message, components)
- Stage ④: **how** the presenter says it (script)
- Stage ⑥: **what the audience sees** (positions, key phrases, figures) — **not** everything from `slides.md`

**Components** are labels for content blocks (e.g. `headline stat`, `three failure modes`, `bottom callout`) — not final on-slide wording.

**Required fields**

| Field | Level | Required | Notes |
|-------|-------|----------|-------|
| Document title | deck | Yes | `# Slide Plan — {Title}` (or equivalent) |
| Section name | section | Yes | Matches `outline.md` |
| Section narrative role | section | Yes | How this section advances the arc |
| Section tone | section | Recommended | e.g. declarative, urgent |
| Section duration | section | Recommended | Should align with outline section duration |
| Slide number | slide | Yes | `01`, `02`, … (main deck); `B01`, … for backup |
| Slide title | slide | Yes | Matches outline (update both if changed) |
| Section (per slide) | slide | Yes | Owning section name |
| Purpose | slide | Yes | Why this slide exists |
| Key message | slide | Yes | One takeaway the slide must land |
| Components | slide | Yes | Content blocks as **labels** — not final phrases |
| Layout overview | slide | Optional | Light hint for generation — concrete layout is stage ⑥ |
| Content layout | slide | Optional | Mode hint: `focal`, `bullets`, `columns`, … |
| Density | slide | Recommended | `normal` \| `dense` \| `compact` when many items |
| Speaker cues | slide | Optional | Delivery hints — feed scripts, not on-slide copy |
| Estimated duration | slide | Yes for main deck | Seconds or minutes; **omit or mark N/A for backup** |
| Template | slide | No at ③ | Added in stage ⑤ |

**Schema:**

```markdown
# Slide Plan — {Title}

## Section 1: {Section name}

- **Narrative role:** {how this section advances the arc}
- **Tone:** {…}
- **Duration:** {e.g. ~3 min}

### Slide 01: {Title}
- **Section:** {Section name}
- **Purpose:** {why this slide exists}
- **Key message:** {one takeaway}
- **Components:** {label}, {label}, {label}
- **Layout overview:** {optional hint}
- **Content layout:** {optional mode}
- **Density:** {normal | dense | compact}
- **Speaker cues:** {optional}
- **Estimated duration:** {e.g. 90 seconds}

### Slide 02: {Title}
- …

## Backup *(optional — excluded from duration; not required for Phase A)*

### Slide B01: {Title}
- **Section:** Backup
- **Purpose:** {…}
- **Key message:** {…}
- **Components:** {labels only}
- **Estimated duration:** N/A
```

### ④ Scripts (`scripts/script-{nn}.md`)

Default path: **`slides.md` → `script-{nn}.md` → (templates) → `slide-{nn}.md`**.

**Entry gate (required):** After `slides.md` is approved, ask before writing any scripts:

> *"Shall I write presenter scripts for each slide, or skip scripts and go on to templates and slide generation?"*

| User choice | Action |
|-------------|--------|
| **Write scripts** (default if unclear) | Produce `scripts/script-{nn}.md` for every **main-deck** slide; present for approval |
| **Skip scripts** | Do not create `scripts/`; proceed to stage ⑤ |

Do **not** record this choice in `overview.md`.

**Script role:** Deliver the **key message** and **components** from `slides.md` in spoken form — full explanation, examples, transitions. Scripts are for the presenter, not for on-slide display.

```markdown
# Script: Slide {nn} — {Title}
## Section
## Timing
## Script
{Spoken narrative that lands the key message and walks through components}
## Key Phrases
{Short phrases that may appear on-slide — candidates for stage ⑥}
## Transition From Previous / To Next
## Notes
```

### ⑤ Template design

Templates are **shared** background and style — not per-slide content. Analyze approved `slides.md` (and scripts if present). Group slides that share the same chrome; split templates when background/style or region structure differs (see **Templates**).

**5a — Assign templates in `slides.md`**

Default: **`template-01`** for cover, **`template-02`** for most body slides (usually **extends** `template-01`). Add **`Template`** to each slide entry.

External reuse: `- Template: {other-deck}/templates/template-01.md`

**5b — Write `templates/template-{nn}.md`**

For each local template ID, write a **renderer-agnostic**, **slide-independent** spec. Record derivation when a template depends on another. External-path templates need a local file only when creating a variant.

```markdown
# Template {nn}: {Name}

## Derivation
- **Base:** {none | template-01 | {other-deck}/templates/template-01.md}
- **Relationship:** {standalone | extends | variant-of}
- **Changes from base:**
  - **Inherited:** {…}
  - **Modified:** {…}
  - **Added:** {…}
  - **Removed:** {…}

## Background
- **Color / Gradient / Image:** {description — no slide-specific art}
- **Treatment:** {full-bleed, bordered, split, …}

## Typography
- **Font family:** {primary + fallbacks}
- **Title / Sub-title / Body / Caption:** {size, weight, color, line-height — plain language}

## Color Palette
- **Primary / Secondary / Accent / Background / Text:** {values or names}

## Regions
{Define only regions this template uses. No slide titles or messages.}

### Title
- **Position:** {top | center for cover}
- **Style:** {from Typography}
- **Max lines:** {1–2}

### Sub-title *(optional)*
- **Position:** {below title}
- **Style:** {…}

### Header *(optional)*
- **Use:** {section label, logo band — structural only}
- **Position:** {top}

### Body
- **Bounds:** {below title/header, above footer — plain language}
- **Default alignment:** {left | center}

### Footer *(optional)*
- **Use:** {slide number, brand, callout area — structural only}
- **Position:** {bottom}

## Layout Defaults
- **Content alignment:** {left | center}
- **Whitespace:** {generous | compact}

## Intended Use
{cover | content | section-divider — no slide references}
```

Start from `overview.md` Visual Style; derive variants instead of duplicating. Present templates for approval before stage ⑥.

### ⑥ PIM generation (`slides/slide-{nn}.md`)

**Generate** the page the audience **sees**: concrete **layout** and **expression** from `slides.md`, assigned **template**, and optional script (**Key Phrases**).

This is not a redesign of architecture or content intent, and **not** a paste of all of `slides.md` or the full script.

| Include on-slide | Leave to script / detail design |
|------------------|----------------------------------|
| Key phrases that land the message | Long explanations, examples, asides |
| Component labels resolved to short visible text | Speaker-only cues |
| Figures, icons, diagrams and **where** they sit | — |
| Positions within template regions (title, body, …) | — |

```markdown
# Slide Content Specification — Slide {nn}

## Slide {nn}: {Title}

**Section:** {name}
**Type:** {optional free-form label — cover, quote, summary, …}
**Template:** {template-02}
**Content layout:** {optional mode hint}
**Density:** {normal | dense | compact}
**Time:** ~{duration}

### Visual Design
- **Background:** {from template — note overrides if any}
- **Layout:** {concrete placement}
  - Title / sub-title: {key phrase}
  - Header: {…}
  - Body: {zones — top / center / bottom / left / right}
  - Footer: {…}
  - Figures / images: {what and where}

### Content Elements
| Element | Content (key phrase) | Notes |
|---------|----------------------|-------|
| {component} | {short on-slide text} | {position, emphasis} |

### Visual Elements Description
- {hierarchy, figures, icons — consistent with template palette and regions}

### Transitions
- **In:** {…}
- **Out:** {…}
- **Notes:** {delivery — not on-slide}

### Reference
- Template: `templates/template-{nn}.md`
- Script: `scripts/script-{nn}.md` *(if in use)*
- Detail design: `slides.md` → Slide {nn}
```

**`Template` must match** `slides.md`. Every visible element needs position and key-phrase (or figure). Present per slide or in batches for approval.

### ⑦ Validation

Strict **PIM quality gate**. Read `overview.md`, `outline.md`, `slides.md`, `templates/`, `slides/`, and `scripts/` when present. Produce **`q-and-a.md`** and **`validation.md`** as files — not chat-only notes.

Phase A validation does **not** require Phase B output. Do not mark Phase A complete while any required check is **FAIL**, or while **Open issues** still need a fix (unless the user explicitly accepts a documented waiver).

**Re-validation:** After any fix from revision paths, update the **affected sections** in `validation.md` and recompute **Final verdict** / **Overall**. Do not leave stale PASS on a section that was changed. A full re-run of all sections is preferred when architecture or many slides change.

#### Q&A (`q-and-a.md`)

Anticipate real audience questions. Prefer **8–15** items for a typical talk; fewer only for very short decks.

**Required fields per question**

| Field | Required | Notes |
|-------|----------|-------|
| `Q{n}` | Yes | Natural phrasing an audience member would use |
| **Type** | Yes | `clarification` \| `challenge` \| `application` \| `scope` \| `logistics` |
| **A** | Yes | Answer the presenter can deliver in ~30–60 seconds |
| **Related slide** | Yes | Primary `slide-{nn}` (or `none` if purely meta) |
| **Deck change** | Yes | `none` \| slide to add/modify — if the question exposes a gap, record the fix |

**Coverage expectations** (include at least one of each where relevant):

| Type | Aim |
|------|-----|
| `clarification` | Core terms or claims a newcomer may not get |
| `challenge` | Pushback on the main thesis or a key slide |
| `application` | “How do I apply this tomorrow?” |
| `scope` | Out-of-scope asks — answer briefly and redirect |
| `logistics` | Time, materials, follow-up (if applicable) |

**Schema:**

```markdown
# Q&A — {Title}

Presenter reference for expected questions.

### Q1: {Question in audience voice}
- **Type:** {clarification | challenge | application | scope | logistics}
- **Related slide:** {slide-0N | none}
- **A:** {concise spoken answer}
- **Deck change:** {none | add/modify slide-0N — reason}

### Q2: …
```

If **Deck change** is not `none`, apply the change through the design ladder before closing validation (or list it under Open issues with owner/decision).

#### Validation (`validation.md`)

Run every section below. Section and row verdicts use **PASS** / **FAIL** / **WARN** only:

| Verdict | Meaning |
|---------|---------|
| **PASS** | Meets the criterion with evidence |
| **FAIL** | Must fix before Phase A complete |
| **WARN** | Does not fully meet the bar; Overall may still PASS **only if** the user explicitly accepts it and the finding is `accepted` under Open issues |

Cite **artifact + slide number** (or section) as evidence — not vague praise. Main-deck slides only for duration and required completeness; backup (`B*`) is optional (see Conventions).

**Required sections and checks**

| Section | Must verify | FAIL when |
|---------|-------------|-----------|
| **1. Purpose fit** | Purpose, audience, narrative goals, validation criteria from `overview.md` | Goals unmet; audience level mismatch; in-scope topics missing |
| **2. Architecture** | `outline.md` vs delivered deck; arc; section roles | Structure diverges from approved outline without update; broken arc |
| **3. Narrative & clarity** | One message per slide; on-slide **key phrases** (not essays); scripts carry explanation when present | Multiple unrelated messages; on-slide prose that belongs in script; unclear takeaway |
| **4. Detail design integrity** | `slides.md` key message/components reflected in scripts and pages; template assignment consistent | Drift; stage-③ **final phrases** instead of component labels; page dumps full `slides.md` text |
| **5. PIM completeness** | Every **main-deck** slide has `slide-{nn}.md`; positions + key phrases/figures for visible elements | Missing main-deck slides; elements without position or visible content |
| **6. Visual consistency** | Template assignment; background/style/regions; templates have **no slide-specific content**; derivation recorded | Assignment mismatch; style drift; slide content inside `templates/` |
| **7. Duration fit** | Sum of **main-deck** per-slide estimates vs `Format & Duration` (exclude `B*`) | Total **>110%** of budget |
| **8. Script quality** *(if scripts exist)* | One script per **main-deck** slide; delivers `slides.md` key message and components; fuller than on-slide text | Missing scripts; script ignores components; on-slide copies full script |
| **9. Q&A readiness** | `q-and-a.md` meets schema and coverage expectations | Thin Q&A; unanswered challenge/application gaps on core claims |
| **10. Completeness** | All required Phase A artifacts present and approved | Any required artifact missing |
| **11. Open issues** | Every finding listed with status | Unresolved FAIL-level issues; unaccepted WARN blocking Overall |
| **12. Final verdict** | Per-section and overall verdict | Overall PASS while any section is FAIL or has unaccepted WARN |

**Duration rules** (main deck only; backup excluded)

| Total estimated time vs budget | Verdict | Action |
|--------------------------------|---------|--------|
| ≤100% | **PASS** | — |
| 100–110% | **WARN** | List concrete cuts (which slides, how many seconds); Overall PASS only if user `accepted` |
| >110% | **FAIL** | Revise estimates and/or merge/split slides until ≤110%; a mitigation note alone does **not** clear FAIL |

**Overall PASS** when:

1. No section verdict is **FAIL**
2. Every **WARN** section/finding is listed under Open issues with status **`accepted`** (user acknowledgment)
3. No Open issues remain with severity FAIL and status `open`

**Schema:**

```markdown
# Validation — {Title}

## 1. Purpose fit

| Criterion (from overview) | Verdict | Evidence |
|---------------------------|---------|----------|
| Purpose | PASS / FAIL / WARN | {artifact, slides} |
| Audience | … | … |
| Narrative goals | … | … |
| Validation criteria | … | … |

## 2. Architecture

| Check | Verdict | Evidence |
|-------|---------|----------|
| outline.md matches delivered slide list | … | … |
| Arc (open → develop → close) | … | … |
| Section roles still valid | … | … |
| Scope / out-of-scope respected | … | … |

## 3. Narrative & clarity

| Check | Verdict | Evidence |
|-------|---------|----------|
| One message per slide | … | … |
| Key messages support narrative goals | … | … |
| On-slide uses key phrases (not full explanation) | … | … |
| Terminology consistent with docs/ overview | … | … |

## 4. Detail design integrity

| Check | Verdict | Evidence |
|-------|---------|----------|
| slides.md titles/order match outline + slides/ | … | … |
| Purpose / key message / components designed in slides.md | … | … |
| slide-*.md does not paste all of slides.md | … | … |
| Template assignment aligned slide-by-slide | … | … |

## 5. PIM completeness

| Check | Verdict | Evidence |
|-------|---------|----------|
| Every main-deck slide has slides/slide-{nn}.md | … | … |
| Concrete layout: position for every visible element | … | … |
| Key phrases and figures/images specified | … | … |
| Template field matches slides.md | … | … |
| Backup slides (if any) optional and excluded from duration | … | … |

## 6. Visual consistency

| Check | Verdict | Evidence |
|-------|---------|----------|
| Template assignment table (all slides) | … | … |
| Background / style / regions vs templates | … | … |
| Templates have no slide-specific content | … | … |
| Template derivation (base + changes) recorded when dependent | … | … |
| Separate templates used when chrome/structure differs | … | … |

### Template assignment

| Slide | Template | slides.md | slide-*.md | Match |
|-------|----------|-----------|------------|-------|
| 01 | template-0N | PASS/FAIL | PASS/FAIL | PASS/FAIL |

## 7. Duration fit

Main deck only — exclude `B*` backup slides.

| Section or slide | Estimated | Notes |
|------------------|-----------|-------|
| … | … | … |
| **Total** | **{mm:ss}** | Budget: {from overview} |

| Verdict | Mitigation |
|---------|------------|
| PASS / WARN / FAIL | {concrete cuts if WARN; N/A if PASS; must revise design if FAIL} |

## 8. Script quality *(omit section if scripts skipped)*

| Check | Verdict | Evidence |
|-------|---------|----------|
| script-{nn}.md for every main-deck slide | … | … |
| Script delivers slides.md key message and components | … | … |
| On-slide is shorter key phrases, not full script | … | … |
| Timing notes plausible vs estimates | … | … |

## 9. Q&A readiness

| Check | Verdict | Evidence |
|-------|---------|----------|
| Schema complete (type, slide, answer, deck change) | … | … |
| Coverage types present | … | … |
| Deck changes applied or tracked | … | … |

## 10. Completeness

| Stage | Artifact | Status |
|-------|----------|--------|
| ① | overview.md | PASS/FAIL |
| ② | outline.md | … |
| ③ | slides.md | … |
| ④ | scripts/ *(or skipped)* | … |
| ⑤ | templates/ + Template fields | … |
| ⑥ | slides/slide-*.md | … |
| ⑦ | q-and-a.md + validation.md | … |

## 11. Open issues

| # | Finding | Severity | Action | Status |
|---|---------|----------|--------|--------|
| 1 | {…} | FAIL / WARN | {fix path} | open / fixed / accepted |

## 12. Final verdict

| Section | Verdict |
|---------|---------|
| Purpose fit | PASS / FAIL / WARN |
| Architecture | PASS / FAIL / WARN |
| Narrative & clarity | PASS / FAIL / WARN |
| Detail design integrity | PASS / FAIL / WARN |
| PIM completeness | PASS / FAIL / WARN |
| Visual consistency | PASS / FAIL / WARN |
| Duration fit | PASS / FAIL / WARN |
| Script quality | PASS / FAIL / WARN / N/A |
| Q&A readiness | PASS / FAIL / WARN |
| Completeness | PASS / FAIL / WARN |
| **Overall** | **PASS / FAIL** |

**Overall PASS** when: no section is FAIL; every WARN is Open issues status `accepted`; no FAIL issues remain `open`.
```

**Revision paths** (then **re-validate** affected sections + Final verdict):

| Problem | Update first | Then |
|---------|--------------|------|
| Structure (sections, titles, order) | `outline.md` | `slides.md` → `slides/` → re-validate |
| Detail design | `slides.md` | `slides/` → re-validate |
| PIM copy / layout | `slides/slide-{nn}.md` | re-validate |
| Template assignment | `slides.md` | affected `slides/` → re-validate |
| Shared background / style (templates) | `templates/` | review assigned slides → re-validate |
| Duration overrun (>110% FAIL) | `slides.md` estimates and/or merge/split | `outline.md` if structure changes → re-validate Duration + Overall |

## Phase B — Renderer-specific authoring (PSM)

Triggered when the user requests generate / render / export **after** Phase A (PIM) is approved.

Phase B **translates** the renderer-independent model into a platform-specific model. It does not redefine purpose, architecture, or on-slide messages.

### Conceptual flow

| Step | Artifact | Role |
|------|----------|------|
| Input (PIM) | `slides/slide-*.md` + `templates/` | Approved renderer-independent page models |
| Per-slide PSM | renderer-specific per-slide file | One platform fragment per slide |
| Compiled PSM | renderer-specific compiled deck | Full deck for the toolchain |
| Export | HTML / PDF / PPTX / … | Optional toolchain output |

Concrete paths (e.g. Marp: `marps/marp-{nn}.md`, `{title-slug}.md`) are defined only in `renderer/{name}.md`.

### Procedure

1. Confirm approved `slides/slide-*.md` and `templates/template-*.md` (or external template paths).
2. Read [`renderer/README.md`](renderer/README.md); follow the chosen renderer (default: Marp).
3. Map templates and spatial composition to **renderer mechanisms** only inside `renderer/{name}.md`.
4. If PIM copy or layout must change, update Phase A artifacts first, then re-translate.

## Dialogue Rules

1. State the current stage and design level (define / architecture / detail / PIM generation / PSM translation)
2. One question at a time when surfacing requirements
3. Socratic questioning — clarify intent
4. Present changes; don't dump files
5. Adapt to scale — short decks may combine adjacent stages with approval
6. Flag **duration** vs constraints — not slide count

## Prohibitions & Completion

**Do not:**
- Treat Phase A artifacts as renderer-specific (no Marp/PPTX syntax in PIM files)
- Redesign narrative or copy inside Phase B — translate only
- Lock a target slide count in `overview.md` or treat count as a success criterion
- Put detail-design fields (purpose, key message, components, template) in `outline.md`
- Put **final on-slide phrases** in `slides.md` — use **component labels** only at stage ③
- Dump all of `slides.md` or the full script into `slide-{nn}.md` — on-slide is **key phrases** and visuals only
- Skip concrete layout (positions, figures) in `slide-{nn}.md`
- Include backup (`B*`) slides in duration totals or treat them as required for Phase A complete
- Record scripts preference in `overview.md` — decide only at stage ④ entry
- Skip stage ④ without asking the user
- Assign `Template` in stage ③ — wait until stage ⑤
- Write `slides/` before templates and `slides.md` template assignments are approved
- Put slide numbers, titles, key messages, or any current-slide content in `templates/`
- Define a new template when chrome is identical — or reuse one template when background/style/regions differ
- Omit template **derivation** (base + changes) when a template depends on another
- Put full script prose on-slide
- Require rendered (PSM) output to complete Phase A validation
- Skip validation or leave artifacts in chat only
- Mark Phase A complete while `validation.md` Overall is FAIL or FAIL issues remain open
- Run renderer specs marked **Planned** in `renderer/README.md`

**Phase A (PIM) complete when:**
- Stages ①–③, ⑤, ⑥, ⑦ passed (④ completed, or skipped only after user declined at stage entry)
- `outline.md` approved as architecture; `slides.md` as detail design; `slides/slide-*.md` as PIM generation
- If scripts were written: every slide has approved `scripts/script-{nn}.md`
- Templates approved; every slide has `Template` in `slides.md` and `slides/slide-{nn}.md`
- `q-and-a.md` meets schema and coverage expectations
- `validation.md` **Overall PASS** (no FAIL sections; every WARN `accepted` under Open issues; no open FAIL issues)
- Backup slides optional — not required for completion
- User gives explicit approval

**Phase B (PSM)** — per `renderer/{name}.md`, when the user requests it.
