---
name: webtoon-authoring
description: >-
  Creates complete vertical-scroll webtoons — episode-by-episode story
  with panel/cut layouts, speech balloons, narration, and matching color
  illustrations. Select when the user asks to write, draft, compose,
  illustrate, design, or author a webtoon, vertical scroll comic.
  Runs Define → Design → Evaluate → Generate with mandatory
  approval gates; designs world, characters, locations, stagings, and
  episodes/pages/cuts before any images; locks approved cut design (art +
  balloons + dialogue + narration) before image generation; generates character,
  location, and staging (ensemble blocking) reference images first, then page
  images via the image-generation skill; stitches pages into one long episode
  scroll. Webtoon structure is **series → episode → page → cut**; location
  consistency uses **location → position → view → state**; multi-cut
  seating/formation continuity uses **stagings**. Page width is fixed (690px);
  height varies by cut layout. For children's picture books with fixed-size
  pages, use picture-book-authoring; for long-form prose fiction, use
  novel-authoring; for a standalone image, use image-generation.
metadata:
  agent: cocrates
---

# Webtoon Authoring

Generate complete vertical-scroll webtoons — from concept through design, evaluation, and image generation — producing a cohesive color comic that reads top-to-bottom on a phone, using the **image-generation** skill for each page segment and assembling pages into one long episode image.

## What a Webtoon Is (vs Picture Book)

| | Picture book | Webtoon |
|---|---|---|
| Reading | Page turn | Vertical scroll |
| Episode output | Separate fixed-size pages | Pages stitched into **one long vertical image** |
| Page size | Uniform | **Width fixed (690px); height variable** |
| Unit of art | Full-page illustration | **Cuts (panels)** with gutters between them |
| Text | Narration/dialogue as page overlay | **Speech balloons** + optional **narration/caption** |
| Color | Style-dependent | **Color by default** (smartphone display) |

### Vertical-scroll craft (built-in)

- Place characters and dialogue to use vertical space (e.g. dialogue block, then figure, then breathing room) — do not cram like a print-page panel grid.
- **Cut height + gutter** set eye speed, tension, and breath — not decoration. Canonical sizes/intents: `workflow/cut-composition.md`.
- **Cuts-per-page follows dwell/breath** — high-dwell beats may be **1 cut per page**; otherwise stack n cuts top→bottom (`cut-composition.md` §4b). Do not max-pack a generation frame.
- **Episode volume ~10–15 pages** (distribution unit) — do not compress a full episode into 2–3 pages; page/cut counts live in episode Craft Notes, not in `series.md`.
- Keep phone-viewport readability: key face/action + balloons should read without pinch-zoom; prefer simpler art and shorter text than print comics when needed.
- Default to **full color**; recall/flashback via color grade (e.g. sepia), not only ink tones.
- Outside-cut area may be white, black, or a series theme color; cuts are **full bleed** or **30–50px** even side margins.
- Design and generate as one continuous scroll; split tall exports (e.g. ~1280px height per file) only when a portal requires it.

## Core Principles
> **The unexamined webtoon is not worth generating.**

- **Design Before Generation**: No image is generated without complete world, character, location, and episode/page/cut design — and explicit user approval through Evaluate (G3).
- **Architecture-First**: Characters, locations, and episode/page/cut design are structured documents before any visual assets.
- **Composition, Not Invention**: Episode and cut design arrange approved world, characters, locations, and stagings. While designing story, **co-design** staging + character states + location states — **reuse** catalog entries when present, **add** new reference models when missing (`workflow/reference-models.md` §7). Missing entities → stop, ask user, update Design catalogs, then resume.
- **Story Lock Before Images**: Page/cut stories, balloon text, narration, gutters, and illustration guides are locked at G3. Image generation implements that lock; it does not rewrite the story.
- **Design-First Revision**: Structure, story, dialogue, character, or location problems are fixed in Design (and re-evaluated) before regenerating images. Never patch story by only changing image prompts.
- **Reference-Based Consistency**: Character, location, and **staging** (ensemble blocking) reference images are generated first; page images reference them for visual and spatial consistency. Rules: `workflow/reference-models.md`.
- **Series Illustration Guide**: Art style, vertical chrome, balloon/caption fonts and boxes are locked in `illustration-guide.md`; every tile/page YAML must apply it — no per-page lettering invention (`workflow/illustration-guide.md`).
- **Scroll-Native Craft**: Design for vertical reading — cut order top→bottom; **cut height + gutter** as directing tools (`workflow/cut-composition.md`); balloon placement that uses vertical space.
- **On-Panel Comprehension**: Readers must understand the episode from **art + balloons + captions** alone; essential facts must not live only in designer notes.
- **Criteria-Driven Evaluation**: Evaluate checks against Validation Criteria from `overview.md`, plus webtoon craft rules — before any MCP generate.
- **YAML Approval Per image-generation**: All image generation follows the image-generation skill's YAML review → approval → MCP generate workflow.

## Resolve Project Root

| Type | When | `{project-root}` |
|------|------|------------------|
| **1** | Workspace is the single project | `.` (workspace root) |
| **2** | Workspace holds multiple peer projects | `{webtoon-slug}/` |
| **3** | Workspace groups projects by kind | `webtoons/{webtoon-slug}/` |

**Require**: inspect workspace structure first; confirm location and name with the user before creating; reuse an existing folder when present.

## Working Location

All webtoon artifacts are authored under `{project-root}/`:

```text
{project-root}/
├── overview.md
├── illustration-guide.md          # series art / chrome / balloon·caption typography lock
├── world-bible.md
├── worlds/
│   └── {world-slug}.md
├── characters.md
├── characters/
│   └── {character-slug}.md
├── locations.md
├── locations/
│   └── {location-slug}.md
├── stagings.md
├── stagings/
│   └── {staging-slug}.md
├── series.md
├── episodes/
│   └── {nnn}-{episode-slug}.md
├── evaluations/
│   └── {nnn}-{episode-slug}.md
├── images/
│   ├── characters/
│   │   ├── {character-slug}.yaml
│   │   ├── {character-slug}.png
│   │   ├── {character-slug}-{state-slug}.yaml
│   │   └── {character-slug}-{state-slug}.png
│   ├── locations/
│   │   ├── {location-slug}.yaml
│   │   ├── {location-slug}.png
│   │   ├── {location-slug}-{position-slug}-{view-slug}.yaml
│   │   ├── {location-slug}-{position-slug}-{view-slug}.png
│   │   ├── {location-slug}-{position-slug}-{view-slug}-{state-slug}.yaml
│   │   └── {location-slug}-{position-slug}-{view-slug}-{state-slug}.png
│   ├── stagings/
│   │   ├── {staging-slug}-establishing.yaml
│   │   ├── {staging-slug}-establishing.png
│   │   ├── {staging-slug}-reverse.yaml
│   │   ├── {staging-slug}-reverse.png
│   │   ├── {staging-slug}-detail.yaml
│   │   └── {staging-slug}-detail.png
│   └── {nnn}-{episode-slug}/
│       ├── {00}.yaml
│       ├── {00}.png
│       ├── ...
│       └── episode-scroll.png          # pages stitched top→bottom
└── output/
    └── {webtoon-slug}-final/
```

## Publication Model

```text
series → episode → page → cut
```

| Unit | Role |
|------|------|
| **Series** | Episode list / overall arc (`series.md`) |
| **Episode** | Generation & evaluation unit; ultimately **one long vertical scroll image** |
| **Page** | Variable-height strip segment (fixed width 690px); several pages concatenate into the episode scroll |
| **Cut** | Panel inside a page; art + balloons + optional narration; gutters between cuts |

**Canvas defaults** (portal-oriented; confirm with user / target portal):

| Spec | Default | Notes |
|------|---------|--------|
| Width | **690–800px** | Lock exact width in `overview.md` (690 common; some portals ~760–800; ~1080 only if required) |
| Height | **Variable** | Sum of cut heights + gutters; not a fixed page size |
| Cut height classes | See `workflow/cut-composition.md` | standard ~400–600px; tall/long ~1,200–2,000px+; action = open/diagonal + tight gutters |
| Gutter classes | See `workflow/cut-composition.md` | normal ~150–300px; pause ~500–1,000px+ for silence/time |
| Pages / episode | **~10–15** | Distribution unit volume; short-form exception only with user note |
| Cuts / episode | **Derived in episode design** | From dwell packing — do **not** plan in `series.md` |
| Generation | **Default: Flash + 1K + 9:16**, typically **2 tiles per page** then stitch; optional Pro + 2K + 1:8 | See `04-generate.md` — never Flash at 2K/1:8 for balloon text |
| Color | **Full color** | Smartphone display default |
| Outside-cut / side margins | Theme fill; full bleed or 30–50px sides | Keep side margins even |

Upload portals often split tall files (e.g. max ~1280px height per file). Design and generate as one continuous scroll; split only at export if the portal requires it.

## Workflow Overview

```
① Define → [G1] → ② Design → [G2] → ③ Evaluate → [G3 Story Lock] → ④ Generate → [G4]
```

| Stage | Purpose | Approval |
|-------|---------|----------|
| **① Define** | Define what webtoon to create; lock Validation Criteria + canvas specs | G1: Overview approval |
| **② Design** | World, characters, locations, episodes/pages/cuts (story + craft) | G2: Full design approval |
| **③ Evaluate** | Criteria + craft + **required critic personas** (Target Reader first); revise design if needed | G3: Evaluation approval = **story lock** |
| **④ Generate** | Reference images → page images → stitch episode scroll | G4: Final result approval |

**Boundary:** Stages ①–③ own story and design. Stage ④ owns image YAML, pixels, and stitch only. Crossing that boundary for story/structure changes requires explicit rollback.

---

## How to Use This Skill

This file (`SKILL.md`) defines global pipeline rules, gates, and prohibitions.
For step-by-step procedures, read the stage workflow file at the start of each stage:

- Stage ① Define: `workflow/01-define.md` (gate artifact: `overview.md`)
- Stage ② Design: `workflow/02-design.md` (gate artifact set: `illustration-guide.md`, `series.md`, `episodes/*.md`, `world-bible.md`, `characters/*.md`, `locations/*.md`, `stagings/*.md`)
- Stage ③ Evaluate: `workflow/03-evaluate.md` (gate artifact set: `evaluations/*.md` = story lock)
- Stage ④ Generate: `workflow/04-generate.md` (gate artifact: final output under `output/`)
- Cut composition (cross-stage): `workflow/cut-composition.md` — cut height, gutter, size class, viewport rhythm
- Reference models (cross-stage): `workflow/reference-models.md` — character / location / staging identity locks
- Illustration guide (cross-stage): `workflow/illustration-guide.md` — series art / chrome / balloon·caption typography lock
- Consistency: `workflow/consistency.md`

## Stage 1: Define
See `workflow/01-define.md`.

- Create `overview.md` (episode/page/cut-scale meta + canvas specs + `Validation Criteria`)
- G1: Do not proceed to Stage ② without user approval

---

## Stage 2: Design
See `workflow/02-design.md`.

- Author/approve `illustration-guide.md`, `world-bible.md`, `characters/*.md`, `locations/*.md`, `stagings/*.md`, `series.md`, `episodes/*.md`
- Episode files include **page → cut** design (art / balloons / dialogue / narration / gutters)
- Multi-cut ensembles with fixed seating/formation get a **staging** — **one per continuing situation** (café, OR, meeting, …); see `workflow/reference-models.md`
- G2: Do not proceed to Stage ③ without user approval

---

## Stage 3: Evaluate
See `workflow/03-evaluate.md`.

- Create `evaluations/*.md` per episode
- Run the **default persona set** (Target Reader, Genre, Plot/Pacing, Story, Platform/editorial, Webtoon art; Character when needed); fill **Adjudication** (Target Reader tie-break)
- G3: Do not proceed to Stage ④ without story-lock approval

---

## Stage 4: Generate
See `workflow/04-generate.md`.

- Reference images (characters/locations/**stagings**) → page images (cuts + balloons + narration) → episode vertical stitch — **every YAML applies `illustration-guide.md`**
- Each image YAML requires explicit user approval before generate
- G4: Do not deliver final output without user approval

---

## Dialogue Rules

1. **State Current Stage**: Always explicitly state which stage (Define/Design/Evaluate/Generate) and sub-step is active.
2. **Approval Before Progress**: Never proceed past a Gate without explicit user approval.
3. **YAML Approval Per Image**: Every image YAML must be shown to the user and receive explicit approval before calling MCP generate. No exceptions. No batch auto-generation.
4. **Rollback Protocol**: If issues are found in later stages, explicitly propose rollback to the relevant earlier stage. Story/structure → Design + Evaluate; visual-only → Generate retry.
5. **Clarification First**: If world, character, or story elements are unclear, ask before designing.
6. **Design-First Fixes**: When Evaluate or Generate surfaces a story/craft problem, update Design artifacts first, re-evaluate, then regenerate images.

## Prohibitions

- Generating images before G3 story lock
- Generating images without user-approved YAML
- **Calling MCP generate without explicit user approval for each YAML** — every image requires separate user confirmation
- **Auto-generating images in batch without stopping for approval for each YAML**
- Proceeding past any Approval Gate without explicit user confirmation
- Skipping the Evaluate stage
- Creating page images before reference images (Phase 0 before Phase 1)
- Skipping staging references for continuing situations that need fixed relative placement (café L/R flips, OR station drift, chair shuffle)
- Inventing characters, locations, world rules, or seating maps inside episode/page/cut design or image prompts
- Silently swapping identity gear (weapons, outfits, signature accessories) across cuts without a new character state
- Treating expression, mood, time-of-day, season, or weather as reference-model states
- Modifying story, dialogue, narration, craft fields, characters, locations, or stagings during Generate without returning to Design → Evaluate → new G3
- Fixing plot, pacing, theme, or text problems by editing image prompts only
- Using non-English text in image-generation YAML **prompts** (locked balloon/narration strings in the target language are rendered as on-image text; prompts themselves stay English)
- Calling MCP generate without image-generation skill's approval workflow
- Designing fixed equal-height pages as if this were a picture book (height must follow cut layout)
- Omitting cut/balloon/gutter design and treating a page as a single undivided illustration by default
- Generating page images without rendering locked balloon dialogue and narration into the image
- **Writing page/tile YAML without applying locked `illustration-guide.md`** (art style, chrome, balloon/caption typography)
- **Inventing per-page balloon fonts, caption boxes, or panel chrome** not defined in `illustration-guide.md`
- Omitting the mandatory **IMPORTANT** block from page-image YAML `design` (exact balloon text; cut order; clear gutters; no CUT/panel labels; smartphone vertical reading)
- Using **`gemini-3.1-flash-image` with `2K` + `1:8`** for pages with balloons/captions (use default Flash + `1K` + `9:16` tiles, or Pro + `2K` + `1:8` when a single tall frame is required)
- Delivering an episode without stitching approved pages into the episode scroll (unless user explicitly wants pages only)
- **Nested cut subsections (`###### Cut story`, `###### Composition`, `###### Illustration guide`, `###### Balloons`, …)** — canonical episode cut schema is flat `- **Field:**` lists under `##### Cut {n}` only (see `workflow/02-design.md`)
- **Omitting required cut fields** when empty (establishing shots still need `- **Characters:** 없음 …`; Captions/Balloons/Staging always present, value may be `없음`)
- **Ghost cast** — listing a character in the episode roster without on-cut appearance or an explicit mention-only tag
- Writing Craft Notes page/cut counts that disagree with measured `### Page` / `##### Cut` headings without a documented exception
- Compressing a distribution-unit episode into ~2–3 pages when the summary needs ~10–15 pages of breath
- Leaving essential story facts only in Cut story / Direction notes so a cold reader cannot follow from art + balloons + captions
- Putting page/cut estimates or budgets in `series.md` (series plans **# / Title / Summary** only)
- Using field notation other than `- **Field:**` (colon must be inside bold; reject `**Field**:` and bare `- Location:`)

## Completion Criteria

- `overview.md` exists with complete webtoon definition, episode structure, **canvas specs (width/aspect)**, **and Validation Criteria**
- `illustration-guide.md` exists with series art style, vertical chrome, and balloon/caption typography locked
- `world-bible.md` exists with world design
- `characters.md` + `characters/{character-slug}.md` exist with complete character designs
- `locations.md` + `locations/{location-slug}.md` exist with complete location designs
- `stagings.md` + `stagings/{staging-slug}.md` exist for every multi-cut ensemble that needs fixed blocking
- `series.md` lists episodes by **# / Title / Summary** only (no page/cut columns) + `episodes/{nnn}-{episode-slug}.md` with complete designs (~10–15 pages unless short-form exception)
- Every episode file matches the **canonical flat cut schema** (flat cut fields; cast roster; no `######` cut subsections) per `workflow/02-design.md`
- All evaluation records exist in `evaluations/{nnn}-{episode-slug}.md` with Criteria Check + **Schema / Structural Integrity** + Craft Checks + required personas + Adjudication
- Evaluate schema checks pass (flat cut schema, field notation, required fields, cast roster, cut counts, size class ↔ height) before G3
- G3 story lock approved
- All reference images generated and approved with explicit user confirmation (Phase 0: characters, locations, stagings)
- All page images generated with balloons/narration and approved with explicit user confirmation (Phase 1) — each YAML applying `illustration-guide.md`
- Episode scroll stitched and reviewed (Phase 2)
- Consistency review completed (Phase 3)
- Final assembly approved by user (G4)
