---
name: image-generation
description: >-
  Select when the user asks to generate, create, draw, illustrate, or design an
  image, picture, illustration, icon, thumbnail, concept art, or other still
  visual asset — or needs a Gemini-compatible image YAML. Uses Gemini image
  models via cocrates-google-genai. Do not select for architecture/flow diagrams
  (diagram-generation), video clips (video-generation), or TTS/speech
  (speech-generation).
metadata:
  agent: cocrates
---

# Image Generation

Produce images that **deliver an approved message** — not decorative guesses.

| Layer | Owns | Must not |
|-------|------|----------|
| **YAML** | `title`, optional `summary`, `message`, `design`, MCP request | Call MCP before YAML approval (except Express) |
| **MCP** `cocrates-google-genai` | Pixels from approved YAML | Invent meaning absent from YAML |

Pipeline: **`message`** (story/beats to convey) → **`design`** (how to visualize it) → **`params.prompt`** (English realization). Not an interpretation of the finished image. (Diagrams use `explanation` — a pointing/interpretation script — instead.)

## Core Principles

- **Message-first** — the image must make `message` (story/beats) graspable at a glance.
- **Design shapes the message** — `design` materializes the story; it is not a separate decorative brief.
- **Chat then YAML** — gather what the YAML needs in chat; write when enough is known. If already enough, write YAML immediately.
- **YAML gate before generate** — user reviews YAML (`design` primary), then MCP `generate`.
- **Consistency** — requirement fields aligned; `design` ↔ `params.prompt` ↔ `params.images` aligned.
- **Intent fidelity** — do not invent unrequested subjects, props, text, or styles.

Human YAML fields use the **user's language**. `params.prompt` is **English only**.

## Workflow

```
1 Discover (chat)   requirement + visual design
2 Write YAML        title / [summary] / message / design / MCP block  → review
3 Approve → Generate   cocrates-google-genai                              → image
4 Revise            update YAML; keep fields consistent
```

**Enough already:** Write complete YAML immediately; stop for approval (unless Express).

**Express** (e.g. *generate now*, *express*): YAML + generate; revise from the image.

Mark only the MCP block with `# --- cocrates-google-genai ---`.

## Paths

Default: `images/{slug}.yaml`, `output: "./{slug}.png"` beside it. Relative paths resolve against **that YAML's directory**. Optional series layout: `characters/`, `locations/`.

---

## Phase 1 — Discover (chat)

Do not write partial YAML here. Work **requirement before visual design**.

### 1.1 Requirement

| Field | Role | Required |
|-------|------|----------|
| `title` | Short image name | Yes |
| `summary` | One-sentence claim | No |
| `message` | Communicative story / beats — what the image should *mean* | Yes |

**`message` (story / beats):** The takeaway narrative the viewer should receive — claims, emotion, relationship, plot beats, or on-image text *as meaning*. Write it so someone could retell the intended story without sketching the picture.

| `message` owns | Belongs in `design` / prompt instead |
|----------------|--------------------------------------|
| What to convey (story, claim, mood takeaway) | Who looks like what, pose, props, wardrobe |
| Ordered meaning beats | Framing, camera, lighting, style, format |
| Why this image exists for the viewer | Refs, keep/change, model, size, aspect |

**Anti-pattern:** A `message` that reads like a shot list or duplicates `design`. Example: message = “the user leads AI” → design = person directing an AI robot (wardrobe, gesture, setting, style…).

**Consistency:** Present fields stay one story (`title` ↔ [`summary`] ↔ `message`). Realign all present fields on edits.

**Fit test:** Could a viewer name the intended story/beats after seeing a picture that matches — without needing the YAML? Could you write a different `design` that still serves the same `message`?

### 1.2 Visual design (maps to `design` + prompt)

Agree in chat (then encode in YAML `design`) — **shape the approved message into a picture** (realization layer; prompt executes this):

| Topic | Capture |
|-------|---------|
| Subject | Who/what, appearance, distinctive features |
| Action / pose | Body position, gesture, gaze |
| Setting | Place, time, props in frame |
| Composition | Framing, angle, lens feel, focal point |
| Style / lighting / mood | Medium, light quality, tone |
| On-image text | Exact strings, placement |
| Format | Model (Flash/Pro), size, aspect |
| Refs | Source YAML or files; what to keep vs change |

**Fit test:** From the design brief alone, could you sketch the picture? Does every design choice serve `message`? Thin briefs (subject + style only) are not enough.

When coherent → Phase 2.

---

## Phase 2 — Write YAML

```yaml
title: Cafe cat
message: |
  A quiet pause in everyday life: warmth, stillness, and gentle companionship
  in a cafe — the viewer should feel unhurried calm, not a busy shop.

design: |
  Fluffy orange tabby with white chest and green eyes, upright on a worn oak
  cafe table, paws together, looking toward a rain-streaked window on the left.
  Steaming cup beside the cat. Soft cafe bokeh. Medium shot, cinematic photo,
  late-afternoon side light. No text. Flash, 1K, 16:9, blog header.

# --- cocrates-google-genai ---
type: image
model: gemini-3.1-flash-image
params:
  prompt: |
    A fluffy orange tabby cat with white chest fur and bright green eyes sits
    upright on a worn oak cafe table, front paws neatly together, gazing toward
    a rain-streaked window on the left. Beside the cat a ceramic cup of coffee
    releases thin steam. Soft cafe bokeh. Medium shot, cinematic photography,
    warm late-afternoon side light. No text.
  size: 1K
  aspectRatio: "16:9"
  seed: null
output: "./cafe-cat.png"
```

### `design`

User-facing picture brief in the user's language — primary YAML review object. **Designs how to visualize `message`**; `params.prompt` realizes that design in English. Must be picture-complete (same topics as §1.2); may be shorter than `params.prompt` but must not omit subject, action, setting, composition, style/lighting, text, refs, or format. Every visible claim in `params.prompt` must be grounded here. Do not restate the communicative story in place of visual decisions — keep story in `message`.

### MCP request

| Field | Role |
|-------|------|
| `type` | `image` |
| `model` | `gemini-3.1-flash-image` (default) or `gemini-3-pro-image` |
| `params.prompt` | English prompt derived from `design` |
| `params.images` | Ordered `{path}` refs |
| `params.size` / `aspectRatio` / `seed` | Format |
| `output` | Default `./{slug}.png` |

Recommend Pro for brand/text/layout fidelity or complex multi-character work. Do not use Flash-only `0.5K` or strip ratios with Pro.

### Prompt & references

- Prompt: subject → action → setting → composition → style → lighting → details → on-image text; concrete English; ~100–250 words.
- Bind refs with `the provided image` or `Image 1` / `Image 2` … matching `params.images` order; names only after binding.
- Resolve YAML-sourced refs via source `output` → verify on disk → path relative to **this** YAML; if missing, search and get user approval before locking.
- Multi-ref: ref map → SCENE → optional TEXT OVERLAY. Max ~14 refs.

| Param | Notes |
|-------|-------|
| `size` | `0.5K` (Flash only), `1K` (default), `2K`, `4K` |
| `aspectRatio` | Default `1:1` unless purpose implies otherwise |
| `images` | Ordered `{path}`; order = `Image N` |

### YAML gate

Present **`design`** (and refs); note model/size/aspect/output. Ask approval; **stop** (unless Express).

---

## Phase 3 — Generate

MCP **`cocrates-google-genai`** (GetMcpTools; `mcp_auth` if needed).

1. YAML on disk.
2. Preflight every `params.images[].path` against **this YAML's directory**; on fail, stop and ask — do not `generate`.
3. `generate` with `filePath` → report `files` (or `download` if background).

---

## Phase 4 — Revise

Keep `design` ↔ prompt ↔ `images` consistent; re-approve; regenerate. Prefer small edits. If `message` changes, reshape `design` to match.

---

## Multi-ref example

```yaml
title: Page 04 — Sidekick asks for help
message: |
  Danger is coming: the sidekick seeks the hero’s help, and the hero must
  respond. Short on-image text carries the page’s spoken beat for young readers.

design: |
  Refs (params.images order):
  1. characters/hero-adventure.yaml → ./characters/hero-adventure.png — keep hero look/outfit
  2. characters/sidekick.yaml → ./characters/sidekick.png — keep sidekick look
  3. locations/plaza-peaceful.yaml → ./locations/plaza-peaceful.png — keep plaza
  Scene: sidekick runs to hero; hero surprised; dust/danger. Text overlays. 4:3, 1K.

# --- cocrates-google-genai ---
type: image
model: gemini-3.1-flash-image
params:
  prompt: |
    Using the reference images: Image 1 is the boy hero in adventure outfit —
    blue tunic, gold emblem, brown boots. Image 2 is the cute sidekick.
    Image 3 is the kingdom plaza.
    SCENE: Children's picture book page. The boy from image 1 meets the sidekick
    from image 2 in the plaza from image 3. Soft digital painting, pastel outlines.
    TEXT OVERLAY: short narration/dialogue without covering faces.
  images:
    - path: "./characters/hero-adventure.png"
    - path: "./characters/sidekick.png"
    - path: "./locations/plaza-peaceful.png"
  size: 1K
  aspectRatio: "4:3"
  seed: null
output: "./page04.png"
```

---

## Prohibitions

- MCP before YAML approval (except Express); auto Express
- `message` that duplicates `design` (shot lists, wardrobe, camera, lighting)
- Inconsistent `design` / prompt / `images`; non-English prompt
- Guessing ref paths; generate after failed preflight
- Inventing unrequested content; ignoring refs; Flash-only options on Pro
