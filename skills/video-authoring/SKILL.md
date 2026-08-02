---
name: video-authoring
description: >-
  Produces a complete multi-clip video project and a final assembled mp4 —
  film, animation, ads, music videos, short-form, educational/explainer videos,
  and similar — with designed story/message across sequence → segment → clip,
  optional visual reference assets, per-clip media components (images, motion
  clips, speech, BGM/SFX), and composition into segment then final video.
  Select when the deliverable is a full video (multiple clips and/or final
  stitch), not a single motion clip. For one standalone clip (~10s or less)
  or a single video YAML/generate, use video-generation; for stills only,
  image-generation; for TTS only, speech-generation; for picture books,
  picture-book-authoring.
metadata:
  agent: cocrates
---

# Video Authoring

Generate complete videos (film, animation, ads, music videos, short-form, educational/explainer videos, and similar) — from intent and story/message design through component generation and assembly — producing a cohesive **clip sequence** and a final assembled mp4.

## Core Principles

> **The unexamined video is not worth generating.**

- **Design Before Generation**: No media YAML or assets until Markdown story/message design is complete and Evaluate (G3) is approved.
- **Markdown vs YAML**: Markdown owns story/message review (`sequence` → `segment` → `clip`). YAML owns implementation (image/video/speech/music) and assembly **edit-specs** (`assembly/*.yaml`). Do not use edit-spec YAML as a substitute for story design.
- **Hierarchy**: Long = `sequence → segment → clip`. Short (single segment) = `sequence → clip`, with **segment-level clip design content living inside `sequence.md`** (no separate `segments/*.md`).
- **Message-First Clips**: Each clip has a designed message; composition (image/video/audio tracks) follows that message.
- **Conditional Reference Models**: Analyze intent/purpose and declare which visual identities need reference images. Not every project needs characters/locations (e.g. slides + narration). When characters/locations are used, apply the shared three-layer model in `workflow/reference-models.md` (character / location / **staging**). While designing story/clips, **co-design** staging + character states + location states — **reuse** catalog entries when present, **add** when missing (§7).
- **Components then Assemble**: Stage ④ generates clip components; Stage ⑤ writes an **edit-spec** YAML and renders segment/final mp4 via MCP **`cocrates-video-edit`** (`validate_spec` → `render_video`).
- **Segment as Approval Unit**: Evaluate, Generate Components, and Assemble are approved **per segment** (Short = the one implicit segment in `sequence.md`).
- **Duration Budgets**: AI motion video clips stay under ~10s. TTS-led clips (still image or held video + speech) may run ~30s.
- **Audio Roles**: Voice → `speech-generation`. BGM/SFX → `music-generation`.
- **YAML Approval**: Every generation YAML follows the subordinate skill’s review → approval → MCP generate workflow. Every **edit-spec** is approved, then validated and rendered via `cocrates-video-edit` — never render before approval.

## Resolve Project Root

| Type | When | `{project-root}` |
|------|------|------------------|
| **1** | Workspace is the single project | `.` (workspace root) |
| **2** | Workspace holds multiple peer projects | `{video-slug}/` |
| **3** | Workspace groups projects by kind | `videos/{video-slug}/` |

**Require**: inspect workspace structure first; confirm location and name with the user before creating; reuse an existing folder when present.

## Working Location

```text
{project-root}/
├── overview.md
├── context.md                         # optional
├── references.md                      # reference-need analysis + catalog kinds
├── references/                        # conditional catalogs (characters, locations, stagings, slides, …)
│   ├── characters/
│   ├── locations/
│   ├── slides/
│   ├── props/
│   └── brand/
├── sequence.md                        # series-level story; Short: includes segment/clip design
├── segments/                          # Long only
│   └── {nnn}-{segment-slug}.md
├── evaluations/
│   └── {nnn}-{segment-slug}.md        # Short: e.g. 001-main.md
├── images/
│   ├── references/
│   └── clips/
├── videos/                            # video-generation outputs
├── speech/                            # speech-generation outputs
├── music/                             # music-generation (BGM/SFX) outputs
├── assembly/
│   ├── segments/
│   │   └── {nnn}-{segment-slug}.yaml  # Short: 001-main.yaml (or assembly/{slug}.yaml)
│   └── final.yaml                     # Long only: stitch segment mp4s
└── output/
    ├── segments/
    └── {slug}.mp4
```

## Hierarchy (story units)

| Unit | Picture-book analogue | Role |
|------|----------------------|------|
| **sequence** | series | Whole-arc story/message |
| **segment** | episode | Mid-unit story; **approval & assembly unit** |
| **clip** | page | Atomic message unit |

```text
Long:   sequence → segment → clip
Short:  sequence → clip   (segment content embedded in sequence.md)
```

## Workflow Overview

```
① Define → [G1] → ② Design → [G2] → ③ Evaluate → [G3 Design Lock]
→ ④ Generate Components → [G4 per segment] → ⑤ Assemble → [G5 per segment; Long: final]
```

| Stage | Purpose | Approval |
|-------|---------|----------|
| **① Define** | Intent, genre, length, Short/Long, Validation Criteria, reference-need sketch | G1: `overview.md` |
| **② Design** | Markdown story/message: sequence → [segment →] clip; optional catalogs | G2: design set |
| **③ Evaluate** | Criteria + craft + reference integrity + **required critic personas** (Target Viewer first) | G3: **design lock** (per segment) |
| **④ Generate Components** | Reference/clip images, video, speech, music | G4: component set (per segment) |
| **⑤ Assemble** | Composition YAML → segment mp4 → (Long) final mp4 | G5: segment (+ final) |

**Boundaries**
- ①–③: story/message only (Markdown).
- ④: media YAML + assets only; no story rewrite.
- ⑤: edit-spec YAML + `cocrates-video-edit` render only; `src` must be G4-approved assets.

## How to Use This Skill

This file defines global pipeline rules, gates, and prohibitions.
For procedures, read the stage workflow at the start of each stage:

- Stage ① Define: `workflow/01-define.md`
- Stage ② Design: `workflow/02-design.md`
- Stage ③ Evaluate: `workflow/03-evaluate.md`
- Stage ④ Generate Components: `workflow/04-generate-components.md`
- Stage ⑤ Assemble: `workflow/05-assemble.md` (schema: `workflow/edit-spec.md`)
- Cross-stage: `workflow/consistency.md`
- Optional research: `workflow/research.md`

## Stage 1: Define

See `workflow/01-define.md` for the detailed procedure.

- Create `overview.md` (genre, Short/Long, Validation Criteria, reference-need sketch)
- Do not proceed to Stage ② before G1 approval

## Stage 2: Design

See `workflow/02-design.md` for the detailed procedure.

- Write `references.md` (+ catalogs if needed), `sequence.md`, and (Long) `segments/*.md`
- Short: include segment-level clip design (message, direction, tracks) inside `sequence.md`
- Do not proceed to Stage ③ before G2 approval

## Stage 3: Evaluate

See `workflow/03-evaluate.md` for the detailed procedure.

- Write and approve `evaluations/*.md` **per segment** (Short = one embedded segment)
- Run the **core persona set** (Target Viewer, Format, Message, Pacing/Editing, Craft Director) plus **format add-ons** from overview (education/explainer → Learning + Stakeholder + usually Audio; ad/brand → Stakeholder; speech-led → Audio); fill **Adjudication** (Target Viewer tie-break; Learning/Stakeholder High may win on education/brand)
- Do not proceed to Stage ④ before G3 design lock

## Stage 4: Generate Components

See `workflow/04-generate-components.md` for the detailed procedure.

- Reference images → clip key images → video / speech / music (generate only after each YAML is approved)
- G4: do not proceed to Stage ⑤ before the **per-segment** component set is approved

## Stage 5: Assemble

See `workflow/05-assemble.md` for the detailed procedure.

- Approve segment **edit-spec** YAML (`clips[]` + `tracks`) → MCP `validate_spec` → `render_video` → segment mp4
- Long: approve `assembly/final.yaml` → same MCP path → final mp4
- Do not deliver before G5

## Subordinate Skills & MCP

| Need | Skill / MCP |
|------|-------------|
| Still / reference / clip key images | `image-generation` |
| Motion clip from image (&lt;~10s) | `video-generation` |
| Voice (narration, dialogue) | `speech-generation` |
| BGM / SFX | `music-generation` |
| Integrate clips → segment / final mp4 | MCP **`cocrates-video-edit`**: `validate_spec`, `render_video` (schema: `workflow/edit-spec.md`) |

## Dialogue Rules

1. **State Current Stage**: Explicitly state Define / Design / Evaluate / Generate Components / Assemble and the active segment.
2. **Approval Before Progress**: Never pass a Gate without explicit user approval.
3. **YAML Approval**: Every media YAML and every **edit-spec** requires explicit approval before generate/render. No batch auto-run.
4. **Segment Unit**: Prefer evaluate / components / assemble loops **per segment**.
5. **Rollback Protocol**: Story/message → Design + Evaluate; component quality → Stage ④ retry; timing/transition only → Stage ⑤ edit-spec retry.
6. **Clarification First**: If intent, reference needs, or clip messages are unclear, ask before designing.

## Prohibitions

- Generating media or assembly edit-specs before G3 design lock for that segment
- Calling media MCP generate or `render_video` without user-approved YAML
- Auto-generating assets or assemblies in batch without per-YAML approval
- Assembling with ad-hoc ffmpeg/scripts instead of **`cocrates-video-edit`**
- Calling `render_video` without a prior successful `validate_spec` (unless validate already succeeded on the same approved file in this turn)
- Proceeding past any Gate without explicit user confirmation
- Skipping Evaluate
- Using edit-spec YAML as the primary story/message design artifact
- Inventing characters/locations/reference entities not declared in Design catalogs (when catalogs exist)
- Changing locked clip messages during Stages ④–⑤ without Design → Evaluate → new G3
- Using non-English text in image/video-generation `params.prompt` fields (follow those skills)
- Treating BGM/SFX as speech-generation (or voice as music-generation)
- Assembling with `src` paths that were not G4-approved (or G5a segment outputs for final)
- Delivering final mp4 before G5
- **Nested clip subsections** (`#### Clip message`, `#### Direction guide`, …) — canonical clip schema is flat `- **Field:**` lists under `### Clip {N}` only (see `workflow/02-design.md`)
- **Omitting required clip fields** when empty (use `none` / `no`)
- **Ghost cast** when catalogs exist — roster without on-clip refs or mention-only tag
- Writing Craft Notes / Segment List clip counts that disagree with measured `### Clip` headings without a documented exception
- Using field notation other than `- **Field:**` (colon must be inside bold)

## Completion Criteria

- `overview.md` with Validation Criteria and Short/Long mode
- `references.md` declaring reference needs (or explicitly none)
- `sequence.md` complete; Long has `segments/*.md`; Short has segment-level clip design inside `sequence.md`
- Clip designs match the **canonical flat clip schema**; Evaluate schema checks pass before G3
- `evaluations/*` with Schema Integrity + G3 design lock per segment (required personas + Adjudication)
- G4: all required component assets for each assembled segment approved
- G5: segment mp4(s) approved via `cocrates-video-edit` render; Long also has approved `assembly/final.yaml` and `output/{slug}.mp4`
