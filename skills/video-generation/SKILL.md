---
name: video-generation
description: >-
  Select when the user asks to generate, create, animate, film, edit, or produce
  a video, clip, motion sequence, image-to-video, reference-to-video, or
  cinematic short — or needs a Gemini Omni-compatible video YAML. Uses Gemini
  Omni Flash via cocrates-google-genai. Do not select for still images
  (image-generation), architecture diagrams (diagram-generation), or TTS/speech
  (speech-generation).
metadata:
  agent: cocrates
---

# Video Generation

Produce videos that **deliver an approved message** — not decorative motion.

| Layer | Owns | Must not |
|-------|------|----------|
| **YAML** | `title`, optional `summary`, `message`, `design`, MCP request | Call MCP before YAML approval (except Express) |
| **MCP** `cocrates-google-genai` | Clip from approved YAML | Invent meaning absent from YAML |

Pipeline: **`message`** (story/beats over time) → **`design`** (how to visualize it in motion) → **`params.prompt`** (English realization). Not a post-hoc clip critique. (Diagrams use `explanation` — a pointing/interpretation script — instead.)

## Core Principles

- **Message-first** — the clip must make `message` (story/beats) graspable as short motion narrative.
- **Design shapes the message** — `design` materializes the story over time; it is not a separate decorative brief.
- **Chat then YAML** — gather what the YAML needs in chat; write when enough is known. If already enough, write YAML immediately.
- **YAML gate before generate** — user reviews YAML (`design` primary), then MCP `generate`.
- **Consistency** — requirement fields aligned; `design` ↔ `params.prompt` ↔ `params.images` aligned.
- **Motion, not a still** — prompt describes action over time, camera, environment change, and audio.
- **Intent fidelity** — do not invent unrequested cast, cuts, dialogue, or music.

Human YAML fields use the **user's language**. `params.prompt` is **English only**.

## Workflow

```
1 Discover (chat)   requirement + motion design
2 Write YAML        title / [summary] / message / design / MCP block  → review
3 Approve → Generate   cocrates-google-genai                              → video
4 Revise            update YAML; keep fields consistent
```

**Enough already:** Write complete YAML immediately; stop for approval (unless Express).

**Express** (e.g. *generate now*, *express*): YAML + generate; revise from the clip.

Mark only the MCP block with `# --- cocrates-google-genai ---`.

## Paths

Default: `videos/{slug}.yaml`, `output: "./{slug}.mp4"` beside it. Relative paths resolve against **that YAML's directory**.

---

## Phase 1 — Discover (chat)

Do not write partial YAML here. Work **requirement before motion design**.

### 1.1 Requirement

| Field | Role | Required |
|-------|------|----------|
| `title` | Short clip name | Yes |
| `summary` | One-sentence claim | No |
| `message` | Communicative story / beats — what the clip should *mean* over time | Yes |

**`message` (story / beats):** The narrative takeaway across the duration — claims, emotion, relationship, plot beats, spoken intent as meaning. Write it so someone could retell the story without directing the shot.

| `message` owns | Belongs in `design` / prompt instead |
|----------------|--------------------------------------|
| What to convey (story, claim, mood arc) | Cast look, gesture choreography, props |
| Meaning beats over time | Shot plan, camera, lighting, style, task/duration |
| Why this clip exists for the viewer | Refs/roles, exact dialogue lines as performance text* |

\*Exact spoken lines may live in `design` (performance) once the *point* of what is said is clear in `message`.

**Anti-pattern:** A `message` that is a motion/camera checklist or duplicates `design`. Example: message = “the user leads AI” → design = person directing an AI robot across the shot (motion, camera, audio…).

**Consistency:** Present fields stay one story (`title` ↔ [`summary`] ↔ `message`). Realign all present fields on edits.

**Fit test:** Could a viewer name the intended story/beats after watching — without the YAML? Could you write a different `design` that still serves the same `message`?

### 1.2 Motion design (maps to `design` + prompt)

Agree in chat (then encode in YAML `design`) — **shape the approved message into motion** (realization layer; prompt executes this):

| Topic | Capture |
|-------|---------|
| Subject & motion | Who/what; how action evolves across the clip |
| Shot plan | Single continuous shot vs multi-shot (Omni often defaults to multi — say so if single-take) |
| Camera | Angle, move (dolly, pan, handheld), framing |
| Setting / light / mood | Place, environment motion, lighting, tone |
| Audio | Music, ambience, dialogue, or `No dialogue` |
| Timing | Beats / cuts aligned with duration (e.g. `[0-3s]…`) |
| Refs & roles | First frame vs subject/style refs; source paths |
| Format | `task`, `durationSeconds`, `aspectRatio` (`16:9` / `9:16`) |

**Fit test:** From the design brief alone, could you storyboard the clip second-by-second (motion + camera + audio)? Does every design choice serve `message`?

When coherent → Phase 2.

---

## Phase 2 — Write YAML

```yaml
title: Garden walk
message: |
  A brief, wordless breath of calm: moving through nature should feel gentle
  and restorative — mood over plot; no spoken story.

design: |
  Single continuous shot. Woman in blue floral dress on a gravel path; dress
  drifts in breeze. Eye-level tracking, slight handheld. Warm morning light.
  Audio: breeze + distant birds; no dialogue. text_to_video, 8s, 16:9.

# --- cocrates-google-genai ---
type: video
model: gemini-omni-flash-preview
params:
  prompt: |
    In a single continuous shot. A woman wearing a blue floral dress walks
    through a sunlit garden along a gravel path; the dress moves lightly in
    the breeze. Smooth tracking shot at eye level. Warm morning light, soft
    leaf shadows. Sound design: gentle breeze, distant birds. No dialogue.
    No scene cuts.
  durationSeconds: 8
  aspectRatio: "16:9"
  task: text_to_video
  seed: null
output: "./garden-walk.mp4"
```

### `design`

User-facing motion brief in the user's language — primary YAML review object. **Designs how to visualize `message` over time**; `params.prompt` realizes that design in English. Cover the §1.2 topics; may be shorter than `params.prompt` but must not omit motion, shot plan, camera, audio, refs/roles, or format. Every motion claim in `params.prompt` must be grounded here. Do not restate the communicative story in place of motion decisions — keep story in `message`.

### MCP request

| Field | Role |
|-------|------|
| `type` | `video` |
| `model` | `gemini-omni-flash-preview` |
| `params.prompt` | English motion prompt from `design` |
| `params.images` | Ordered `{path}` refs |
| `params.durationSeconds` | e.g. `8` |
| `params.aspectRatio` | `"16:9"` or `"9:16"` |
| `params.task` | `text_to_video` / `image_to_video` / `reference_to_video` / `edit` |
| `params.seed` | `null` or int |
| `output` | Default `./{slug}.mp4` |

Set `task` when image roles are ambiguous. If omitted: 0 images → `text_to_video`, 1 → `image_to_video`, 2+ → `reference_to_video`.

### Prompt & references

- Prompt: subject → action over time → setting motion → camera → light/mood → **shot plan** → audio → timing/text → in-prompt avoidances (`No dialogue`). Say `single continuous shot` when needed (Omni often prefers multi-shot). Edits: short + `Keep everything else the same`.
- Resolve refs like image-generation (source YAML `output`, verify, search + approve if moved).
- Bind with `<FIRST_FRAME>`, `<IMAGE_REF_0>`… and/or `Image N` + explicit role. Max ~10 images.

### YAML gate

Present **`design`**; note duration/aspect/task/output; warn that video can take time. **Stop** (unless Express).

---

## Phase 3 — Generate

MCP **`cocrates-google-genai`**. Preflight `params.images` against this YAML's directory; on fail do not generate. Then `generate` / `download`. Do not treat empty `files` as success.

Do not promise: audio-ref upload, extension/interpolation, or a separate negative-prompt field.

---

## Phase 4 — Revise

Keep `design` ↔ prompt ↔ images/roles consistent; re-approve; regenerate or short `continue_interaction`. If `message` changes, reshape `design` to match.

---

## Image-to-video example

```yaml
title: Cafe call comes alive
message: |
  A warm phone call comes to life — connection and lightness without spoken
  words; the still should feel inhabited, not narrated.

design: |
  Ref: ./cafe-portrait.png as starting frame. Keep identity and framing.
  Motion: gestures, smile beats, latte drip. Single continuous medium shot.
  Lo-fi cafe bed; no dialogue. image_to_video, 8s, 16:9.

# --- cocrates-google-genai ---
type: video
model: gemini-omni-flash-preview
params:
  prompt: |
    In a single continuous shot, turn this into realistic footage using the
    image as the starting frame. <FIRST_FRAME> A young woman at a cafe table
    holds a phone to her ear, gesturing and laughing. Condensation drips on an
    iced latte. Warm window light; subtle handheld medium shot. Sound: soft
    lo-fi cafe music, distant cups. No dialogue.
  images:
    - path: "./cafe-portrait.png"
  durationSeconds: 8
  aspectRatio: "16:9"
  task: image_to_video
  seed: null
output: "./cafe-call.mp4"
```

---

## Prohibitions

- MCP before YAML approval (except Express); auto Express; empty `files` as success
- `message` that duplicates `design` (shot lists, camera, lighting, motion checklists)
- Still captions without motion/camera/audio; vague `make it move`
- Guessing ref paths; failed preflight; aspect ratios other than `16:9` / `9:16`
- Inventing unrequested cast/cuts/dialogue/music; mis-tagging first frame vs ref
