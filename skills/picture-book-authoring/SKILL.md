---
name: picture-book-authoring
description: >-
  Creates complete children's picture books (ages 0–8) — page-by-page story
  text plus matching illustrations. Select when the user asks to write, draft,
  compose, illustrate, design, or author a picture book, children's book,
  illustrated kids' story, 그림책, or 동화책 with images. Runs Define → Design
  → Evaluate → Generate with mandatory approval gates; designs world,
  characters, locations, stagings, and episodes/pages (with per-page design) before any images;
  locks approved story text before image generation; generates character, location, and staging reference images first, then page images via the image-generation
  skill for visual consistency. Picture book structure is **series → episode → page**
  location consistency uses **location → position → view → state**; continuing-situation placement uses **stagings** (same three-layer reference model as webtoon/video/novel).
  For long-form prose fiction without page illustrations, use
  novel-authoring; for general non-fiction prose, use document-authoring; for
  a standalone image (not a full book), use image-generation.
metadata:
  agent: cocrates
---

# Picture Book Authoring

Generate complete children's picture books (ages 0–8) — from story concept through design, evaluation, and image generation — producing a cohesive narrative with AI-generated illustrations for each page using the **image-generation** skill.

## Core Principles
> **The unexamined picture book is not worth generating.**

- **Design Before Generation**: No image is generated without complete world, character, location, and episode/page design — and explicit user approval through Evaluate (G3).
- **Architecture-First**: Characters, locations, and episode/page design are designed as structured documents before any visual assets are created.
- **Composition, Not Invention**: Episode and page design arrange approved world, characters, locations, and stagings. While designing story, **co-design** staging + character states + location states — **reuse** catalog entries when present, **add** new reference models when missing (`workflow/reference-models.md` §7). Missing entities → stop, ask user, update Design catalogs, then resume. Do not invent mid-page.
- **Story Lock Before Images**: Page stories, rendering text, and illustration guides are locked at G3. Image generation implements that lock; it does not rewrite the story.
- **Design-First Revision**: Structure, story, text, character, or location problems are fixed in Design (and re-evaluated) before regenerating images. Never patch story by only changing image prompts.
- **Reference-Based Consistency**: Character, location, and **staging** reference images are generated first; page images reference them. Canonical rules: `workflow/reference-models.md` (same three layers as webtoon/video/novel).
- **Criteria-Driven Evaluation**: Evaluate checks the book against Validation Criteria from `overview.md`, plus picture-book craft rules — before any MCP generate.
- **YAML Approval Per image-generation**: All image generation follows the image-generation skill's YAML review → approval → MCP generate workflow.

## Resolve Project Root

| Type | When | `{project-root}` |
|------|------|------------------|
| **1** | Workspace is the single project | `.` (workspace root) |
| **2** | Workspace holds multiple peer projects | `{book-slug}/` |
| **3** | Workspace groups projects by kind | `picture-books/{book-slug}/` |

**Require**: inspect workspace structure first; confirm location and name with the user before creating; reuse an existing folder when present.

## Working Location

All picture book artifacts are authored under `{project-root}/`:

```text
{project-root}/
├── overview.md
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
│       └── {00}.png
└── output/
    └── {book-slug}-final/
```

## Workflow Overview

```
① Define → [G1] → ② Design → [G2] → ③ Evaluate → [G3 Story Lock] → ④ Generate → [G4]
```

| Stage | Purpose | Approval |
|-------|---------|----------|
| **① Define** | Define what picture book to create; lock Validation Criteria | G1: Overview approval |
| **② Design** | Design world, characters, locations, episodes/pages (story + craft) | G2: Full design approval |
| **③ Evaluate** | Criteria + craft + **required critic personas** (Target Child first); revise design if needed | G3: Evaluation approval = **story lock** |
| **④ Generate** | Generate reference images and page images from locked design | G4: Final result approval |

**Boundary:** Stages ①–③ own story and design. Stage ④ owns image YAML and pixels only. Crossing that boundary for story/structure changes requires explicit rollback.

---

## How to Use This Skill

This file (`SKILL.md`) defines global pipeline rules, gates, and prohibitions.
For step-by-step procedures, read the stage workflow file at the start of each stage:

- Stage ① Define: `workflow/01-define.md` (gate artifact: `overview.md`)
- Stage ② Design: `workflow/02-design.md` (gate artifact set: `series.md`, `episodes/*.md`, `world-bible.md`, `characters/*.md`, `locations/*.md`, `stagings/*.md`)
- Stage ③ Evaluate: `workflow/03-evaluate.md` (gate artifact set: `evaluations/*.md` = story lock)
- Stage ④ Generate: `workflow/04-generate.md` (gate artifact: final output under `output/`)
- Reference models (cross-stage): `workflow/reference-models.md` — character / location / staging

## Stage 1: Define
For detailed procedure, see `workflow/01-define.md`.

- Create `overview.md` (episode/page-scale metadata + `Validation Criteria`)
- G1: Do not proceed to Stage ② until the user approves

---

## Stage 2: Design
For detailed procedure, see `workflow/02-design.md`.

- Author/approve `world-bible.md`, `characters/*.md`, `locations/*.md`, `series.md`, `episodes/*.md`
- G2: Do not proceed to Stage ③ until the user approves

---

## Stage 3: Evaluate
For detailed procedure, see `workflow/03-evaluate.md`.

- Create `evaluations/*.md` per episode
- Run **required personas** (Target Child, Caregiver, Genre/Age, Story Critic, Illustration specialist; Educator/Librarian when overview locks education use) + **Adjudication** (Target Child first)
- G3: Do not proceed to Stage ④ until the user approves (story lock)

---

## Stage 4: Generate
For detailed procedure, see `workflow/04-generate.md`.

- Reference images (characters/locations/stagings) → page images
- Generate each image only after explicit user approval of its YAML
- G4: Do not deliver the final result until the user approves

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
- **Auto-generating images in batch without stopping for approval per image**
- Proceeding past any Approval Gate without explicit user confirmation
- Skipping the Evaluate stage
- Creating page images before reference images (Phase 0 before Phase 1)
- Inventing characters, locations, or world rules inside episode/page design or image prompts
- Modifying story, rendering text, craft fields, characters, or locations during Generate without returning to Design → Evaluate → new G3
- Fixing plot, pacing, theme, or text problems by editing image prompts only
- Using non-English text in image-generation YAML prompts
- Calling MCP generate without image-generation skill's approval workflow
- Writing page rendering text that only captions the illustration, or a final-page moral monologue that states the theme instead of showing it
- Generating page images without including the locked rendering text as overlay in the image

## Completion Criteria

- `overview.md` exists with complete book definition, episode structure (episode count + pages per episode), **and Validation Criteria**
- `world-bible.md` exists with world design
- `characters.md` + `characters/{character-slug}.md` exist with complete character designs
- `locations.md` + `locations/{location-slug}.md` exist with complete location designs
- `series.md` + `episodes/{nnn}-{episode-slug}.md` exist with complete episode designs (including craft fields)
- All evaluation records exist in `evaluations/{nnn}-{episode-slug}.md` with Criteria Check + Craft Checks + required personas + Adjudication
- G3 story lock approved
- All reference images generated and approved with explicit user confirmation (Phase 0)
- All page images generated with text overlay and approved with explicit user confirmation (Phase 1)
- Consistency review completed (Phase 2)
- Final assembly approved by user (G4)

