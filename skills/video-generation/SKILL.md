---
name: video-generation
description: >-
  Select when the user asks to generate, create, animate, film, edit, or produce
  a video, clip, motion sequence, image-to-video, reference-to-video, or
  cinematic short — or needs a Gemini Omni-compatible video YAML. Also select
  when they ask to interpret, analyze, evaluate, critique, describe, or ask
  questions about an existing video (including using analysis to revise and
  regenerate via YAML). Uses Gemini Omni Flash and MCP analyze via
  cocrates-google-genai. Do not select for still images (image-generation),
  architecture diagrams (diagram-generation), or TTS/speech (speech-generation).
metadata:
  agent: cocrates
---

# Video Generation

Produce videos that **deliver an approved message** — not decorative motion.

| Layer | Owns | Must not |
|-------|------|----------|
| **YAML** | `title`, optional `summary`, `message`, `design`, MCP request | Call MCP before YAML approval (except Express) |
| **MCP** `cocrates-google-genai` | `generate` → clip; optional `analyze` → text | Invent meaning absent from YAML; auto-analyze |

Pipeline: **`message`** (story/beats over time) → **`design`** (how to visualize it in motion) → **`params.prompt`** (English realization). Not a post-hoc clip critique. (Diagrams use `explanation` — a pointing/interpretation script — instead.)

## Core Principles

- **Message-first** — the clip must make `message` (story/beats) graspable as short motion narrative.
- **Design shapes the message** — `design` materializes the story over time; it is not a separate decorative brief.
- **Chat then YAML** — gather what the YAML needs in chat; write when enough is known. If already enough, write YAML immediately.
- **YAML gate before generate** — user reviews YAML (`design` primary), then MCP `generate`.
- **User review by default** — after generate, the user evaluates the clip; AI `analyze` only on explicit request.
- **Consistency** — requirement fields aligned; `design` ↔ `params.prompt` ↔ `params.references` aligned.
- **Motion, not a still** — prompt describes action over time, camera, environment change, and audio.
- **Intent fidelity** — do not invent unrequested cast, cuts, dialogue, or music.

Human YAML fields use the **user's language**. `params.prompt` is **English only**.

## Workflow

```
A Forward
  1 Discover (chat)   requirement + motion design
  2 Write YAML        title / [summary] / message / design / MCP block  → review
  3 Approve → Generate   cocrates-google-genai                              → video
  4 Review            user evaluates (default); optional AI analyze on request
  5 Revise            update YAML; keep fields consistent; regenerate if needed

B From existing media (understand → YAML → regenerate)
  analyze video → draft message/design from findings → Write YAML (refs source)
  → review/approve → Generate → Review / Revise as in A
```

**Enough already:** Write complete YAML immediately; stop for approval (unless Express).

**Express** (e.g. *generate now*, *express*): YAML + generate; revise from the clip.

Mark only the MCP block with `# --- cocrates-google-genai ---`.

## Paths

Default: `videos/{slug}.yaml`, `output: "./{slug}.mp4"` beside it. Relative paths resolve against **that YAML's directory**.

---

## Understand & revise from media

When the user asks to interpret / analyze / evaluate / Q&A an **existing** video:

1. **Analyze** — MCP `analyze` with `inputs: [video path|URL]` (YouTube/http allowed when public) and a prompt for the user’s question. Present `{ text, interactionId }`. Follow up with `continue_interaction` as needed.
2. **Material for YAML** — from the analysis (and user intent), draft `message` + `design` (motion, shot plan, audio, keep/change). Confirm before locking.
3. **Write YAML** — full request YAML; cite the source in `design` and in `params.references` when using it as first frame / style / edit ref.
4. **Review / approve** → **Generate** → **Review / Revise** (Phases 2–5).

Do not skip the YAML gate. Apply only user-agreed edits into YAML before regenerate.
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
| Meaning beats over time | Shot plan, camera, lighting, style, duration |
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
| Format | `durationSeconds`, `aspectRatio` (`16:9` / `9:16`); Omni `task` is inferred from refs (do not put `params.task` in YAML) |

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
  Audio: breeze + distant birds; no dialogue. 8s, 16:9.

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
| `params.references` | Ordered `{path, type?}` refs (`image` / `video` / `audio`) |
| `params.durationSeconds` | e.g. `8` |
| `params.aspectRatio` | `"16:9"` or `"9:16"` |
| `params.seed` | `null` or int |
| `output` | Default `./{slug}.mp4` |

**Do not set `params.task`.** Runtime infers Omni task from references: 0 → `text_to_video`; one image only → `image_to_video`; 2+ images or any video/audio → `reference_to_video`. Multi-turn edits use `continue_interaction`.

### Prompt & references

- Prompt: subject → action over time → setting motion → camera → light/mood → **shot plan** → audio → timing/text → in-prompt avoidances (`No dialogue`). Say `single continuous shot` when needed (Omni often prefers multi-shot). Edits: short + `Keep everything else the same`.
- Resolve refs like image-generation (source YAML `output`, verify, search + approve if moved).
- Bind with `<FIRST_FRAME>`, `<IMAGE_REF_0>`… and/or `Image N` + explicit role. Max ~10 refs via `params.references`.

### YAML gate

Present **`design`**; note duration/aspect/output (and inferred task from refs); warn that video can take time. **Stop** (unless Express).

---

## Phase 3 — Generate

MCP **`cocrates-google-genai`**. Preflight `params.references` against this YAML's directory; on fail do not generate. Then `generate` / `download`. Do not treat empty `files` as success.

Do not promise: audio-ref upload, extension/interpolation, or a separate negative-prompt field.

Then Phase 4.

---

## Phase 4 — Review (user default; optional AI analyze)

**Default:** Present the generated video path(s). The **user** reviews and evaluates. Do not call `analyze` unless asked.

**Optional AI analyze** — only on explicit request (e.g. *analyze*, *evaluate*, *AI review*, *평가해줘*). Never auto-run after generate.

1. Resolve the artifact path from YAML `output` (relative to the YAML directory); verify on disk. Prefer an absolute path (or a path valid against the MCP process CWD).
2. Call MCP **`analyze`**:
   - `inputs`: `[artifact path]` (1–10; may include first-frame/refs if the user wants a comparison)
   - `model`: omit unless the user overrides (default `gemini-3.5-flash`)
   - `prompt`: evaluation brief in the **user's language**, grounded in approved `message` + `design` (and `summary` if present). Ask for a structured report covering **all** of:
     1. **Intent fit** — does the clip match the designed motion intent (`design` / shot plan / camera / audio)?
     2. **Message delivery** — is `message` (story/beats over time) clearly conveyed?
     3. **Functional** — missing motion, wrong task/ref roles, dialogue/music errors, duration/aspect issues
     4. **Quality / completeness** — jitter, cuts, identity drift, audio sync, polish gaps
     5. **Improvements** — concrete, prioritized suggestions (what to change in `design` / prompt)
3. Present `{ text, interactionId }` to the user. Do **not** silently edit YAML or regenerate from the report.
4. Follow-ups: `continue_interaction` with the same `interactionId` if the user asks more. New media → new `analyze` call.

---

## Phase 5 — Revise

Keep `design` ↔ prompt ↔ references/roles consistent; re-approve; regenerate or short `continue_interaction`. If `message` changes, reshape `design` to match.

When improving from user feedback or an AI analyze report: apply only **user-agreed** changes; update owning YAML fields, then regenerate.

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
  Lo-fi cafe bed; no dialogue. 8s, 16:9.

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
  references:
    - path: "./cafe-portrait.png"
      type: image
  durationSeconds: 8
  aspectRatio: "16:9"
  seed: null
output: "./cafe-call.mp4"
```

---

## Prohibitions

- MCP before YAML approval (except Express); auto Express; empty `files` as success
- Auto `analyze` without an explicit user request; applying analyze suggestions without user agreement
- `message` that duplicates `design` (shot lists, camera, lighting, motion checklists)
- Still captions without motion/camera/audio; vague `make it move`
- Guessing ref paths; failed preflight; aspect ratios other than `16:9` / `9:16`
- Inventing unrequested cast/cuts/dialogue/music; mis-tagging first frame vs ref
- Putting `params.task` in YAML (ignored; task is inferred from references only)
