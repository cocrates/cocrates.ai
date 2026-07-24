---
name: video-generation
description: >-
  Generates AI video prompts and YAML specs for Gemini Omni Flash
  (gemini-omni-flash-preview), then — after user approval — creates the video
  via google-genai-mcp `generate` (sync). Select when the user asks to
  generate, create, animate, film, edit, or produce a video, clip, motion
  sequence, image-to-video, reference-to-video, or cinematic short — or when
  they need an Omni-compatible video YAML. Captures user intent first, writes a
  motion-ready prompt with camera, timing, and audio, and only calls MCP after
  explicit approval.
compatibility: opencode
metadata:
  agent: cocrates
---

# Video Generation

Generate AI video prompts from user intent and output a YAML specification file compatible with Gemini Omni Flash (`gemini-omni-flash-preview`) — [Gemini Omni docs](https://ai.google.dev/gemini-api/docs/omni).

## Core Principles

- **Intent-First**: Understand *what the user wants to see move* before writing anything. Never invent story, cuts, or audio that conflict with stated or clearly implied intent.
- **Prompt as Motion Picture**: The prompt must describe the finished clip so a reader can *play it in their mind* — subject, action over time, camera, setting, lighting, and sound — without guessing.
- **Cinematic Control**: Scene description + camera movement + lighting/mood + audio (+ timing when multi-beat). Not a static image caption.
- **YAML as Contract**: First deliverable is a `{video-slug}.yaml` request file (google-genai-mcp schema).
- **YAML review + approval gate**: After writing the YAML, **stop** for user review. Do **not** create the final media (video) until the user has reviewed the YAML and given **explicit approval**. The user may edit the YAML directly; treat their file as source of truth. Call MCP `generate` only after that approval.

## Workflow

```
Understand Intent → Infer Task → Confirm Brief → Generate Prompt → Output YAML
  → Review YAML → Ask to Generate → (on approval) MCP generate → (Revise if needed)
```

No multi-stage document gates. Ask only what is unclear, confirm the motion brief, write YAML, then **ask before generating**.

## Working Location

### Spec folder (YAML)

When the user does **not** name a location, write the YAML under the workspace kind folder:

```text
videos/{video-slug}.yaml
```

If the user specifies a folder or path, use that instead. Prefer an existing `videos/` tree over inventing a parallel layout.

### Generated artifacts (`output`)

| Rule | Path |
|------|------|
| **Default** | Same directory as the YAML file (e.g. `videos/{slug}.mp4`) — **not** `./output/` |
| **User override** | Exact folder or file path the user named |

Reference frames may live beside the YAML (e.g. `videos/references/…`) or wherever the user points.

**Example layout (defaults):**

```text
videos/
├── garden-walk.yaml
├── garden-walk.mp4
└── references/
    ├── dress.jpg
    └── woman.png
```

---

## Step 1: Understand Intent

Video quality depends on intent fidelity *and* motion specificity. Do not skip this when the request is ambiguous.

### 1.1 Extract what is already known

| Dimension | Capture |
|-----------|---------|
| **Purpose** | Why the clip exists (social reel, product demo, story beat, mood film, ad, tutorial B-roll, edit of an existing clip) |
| **Subject** | Who/what is on screen; appearance; consistency needs |
| **Action / motion** | What moves, how fast, how it evolves across the clip |
| **Shot plan** | Single continuous shot vs multi-shot narrative (Omni defaults to multiple shots unless told otherwise) |
| **Camera** | Angle, move (dolly, pan, handheld, orbit), lens feel |
| **Setting** | Place, time, weather, environment motion |
| **Style** | Photoreal, anime, illustration-to-live, cyberpunk, documentary, etc. |
| **Mood** | Tone and energy |
| **Audio** | Dialogue? Music genre/energy? Ambience? Silence / “No dialogue”? |
| **Timing** | Timed beats, cuts, text on screen (use natural language or `[0-3s]` blocks) |
| **References** | First frame, subject refs, style refs (images); prior video for edit |
| **Format** | `durationSeconds`, `aspectRatio`, `task` |
| **Constraints** | Must-keep / must-avoid (put simple avoidances *in* the prompt, e.g. `No dialogue`) |

Treat explicit user wording as authoritative.

### 1.2 Infer carefully; do not over-invent

- Fill only what purpose **strongly implies** (e.g. Instagram Story → `9:16`; YouTube bumper → `16:9`; product hero → slow camera + clean audio).
- Do **not** add characters, plot twists, dialogue, or scene cuts the user did not ask for.
- Vague “make it move” is not enough — push for (or infer from context) **subject motion + camera + environment**.

### 1.3 Infer `task`

| `task` | When |
|--------|------|
| `text_to_video` | Text only — no images |
| `image_to_video` | One image as starting frame / bring-a-still-to-life |
| `reference_to_video` | One or more images as subject/style/object references (not necessarily the first frame) |
| `edit` | Conversational refine of a previously generated Omni video |

If omitted in YAML, the runtime may infer from inputs — still set `task` when the role of images is ambiguous (first frame vs reference). Prefer omitting only when inference is obvious.

### 1.4 Ask only what is still unclear

| Check | Question (ask in the user's language) |
|-------|----------------------------------------|
| Subject & motion | Who/what is on screen, and what should they *do*? |
| Shot plan | One continuous shot, or several cuts? |
| Camera | Any preferred camera move or framing? |
| Setting / style | Where, and what look? |
| Audio | Music, ambience, dialogue, or silent / no dialogue? |
| Timing | Any timed beats (e.g. enter at 3s, cut every 2s)? |
| References | First-frame image? Subject/style reference images? |
| Duration | How long? (e.g. 8 seconds) |
| Aspect | `16:9` (default landscape) or `9:16` (portrait)? |

### 1.5 Confirm the motion brief

Restate in 2–4 sentences: subject, motion, shot plan, camera, audio, refs, **task**, duration, aspect. Proceed on confirm — or immediately if the request was already fully specific.

**Good brief example:**
> Continuous handheld shot of a woman in a blue floral dress walking through a sunlit garden; soft breeze, birds, no dialogue; dress + woman as references — `reference_to_video`, 8s, 16:9 — product mood film.

---

## Step 2: Generate Prompt

Write English prose that could stand alone as a director’s brief for the clip. A stranger reading only the prompt should imagine the *moving* picture and hear the soundtrack intent.

Follow the [Gemini Omni Flash prompt guide](https://ai.google.dev/gemini-api/docs/omni#gemini-omni-flash-prompt-guide): scene + camera + lighting/mood + audio; explicit single-shot wording when needed; timed events; simple edit language; image role tags when using refs.

### 2.1 Structure (text-to-video / generation)

Natural prose (not a labeled list), covering:

> **[Subject & appearance] + [Action over time] + [Setting & environment motion] + [Camera & lens] + [Lighting & mood] + [Shot plan] + [Audio] + [Timed beats / on-screen text if any] + [Simple avoidances if needed]**

| Layer | Spell out | Thin ❌ | Motion-ready ✅ |
|-------|-----------|--------|-----------------|
| **Subject** | Who/what, clothing, materials | `a woman` | `a woman wearing a blue floral dress` |
| **Action** | How motion unfolds | `walking` | `walks slowly along a gravel path, dress hem drifting in a light breeze` |
| **Setting** | Place + moving environment | `in a garden` | `sunlit garden, soft leaf shadows shifting, distant fountain spray` |
| **Camera** | Move + framing | *(omitted)* | `continuous smooth tracking shot at eye level, slight handheld drift` |
| **Light / mood** | Time, quality, tone | `nice light` | `warm late-morning sun, gentle lens flare through trees, calm mood` |
| **Shot plan** | Single vs multi | *(default multi — risky)* | `In a single continuous shot. No scene cuts.` |
| **Audio** | Music / SFX / speech | *(silent guess)* | `Sound design: gentle breeze, distant bird chirps. No dialogue.` |
| **Timing** | Beats on a clock | `later she turns` | `[0-3s] she walks toward camera; [3-6s] she pauses and turns; [6-8s] she walks out of frame right` |

### 2.2 Shot plan (critical)

By default Omni tries **several shots** and a small narrative. If the user wants one take, you **must** say so explicitly, e.g.:

- `In a single unbroken scene`
- `In a single continuous shot`
- `No scene cuts`
- `Continuous, unbroken handheld shot…`

### 2.3 Audio

Omni generates audio by default. Specify when it matters:

- `Include calm background music`
- `The video has a high energy techno beat`
- `The audio is a low tinny radio broadcast in the background, playing a song`
- `Sound design: Gentle breeze, distant bird chirps. No dialogue.`

Put simple avoidances **in the prompt** (Omni has no separate negative-prompt field): `No dialogue`, `No embellishments`, `No extra sound effects`.

### 2.4 Timing events

Natural language or timecode blocks — useful for cuts, rhythm, and rapid-fire:

- `After 3 seconds, a woman enters the scene.`
- `At 5s the chorus starts in the background audio.`
- `Every 2s cut to a new frame.`

```
[0-3s] A person is walking
[3-6s] They stop and turn around
[6-10s] They start running
```

Align beat structure with `durationSeconds` in YAML.

### 2.5 Prompt patterns by use case

**Text-to-video (cinematic)**
> [Shot plan]. [Subject] [action] in [setting]. [Camera move]. [Lighting/mood]. [Audio]. [Avoidances].

**Image-to-video (animate a still)**
> Turn this into [realistic footage / style], using the image as [starting frame / motion guide]. [Specific subject motion]. [Camera move]. [Environmental effects]. [Audio]. Do not show the still as a flat drawing if converting illustration → live action.

Vague `make it move` ❌ — prefer concrete motion + camera + environment ✅.

**Reference-to-video (subjects / style from images)**
> Use tags (see §2.7). Describe the action and how each ref appears. End with guiding instruction if roles could be confused.

**Product / commercial**
> Clean continuous shot of [product] [motion: rotate / pour / unbox]. Studio or lifestyle setting. Slow cinematic camera. Audio: subtle whooshes / soft music. No dialogue. Exact on-product text if any.

**Multi-shot narrative / rapid-fire**
> Timed blocks or `Every 1s…`. Upbeat music. Optional on-screen labels. Meta: rich micro-detail, natural timing.

**Edit (follow-up)**
> Keep prompts **short**. Over-describing causes unintended changes. Always add `Keep everything else the same` when changing one aspect.

Edit examples:
- `Make this video anime. Keep everything else the same.`
- `Add a cat that jumps onto his lap, he begins to pet it. Keep everything else the same.`
- `Make the phone invisible. Keep everything else the same.`
- `Change the text on the sign to say "Omni Flash". Keep everything else the same.`

Avoid long rewrite-the-whole-scene edit prompts.

### 2.6 Text in video

Specify exact readable strings for signs, titles, labels, plates:

- `One word on the screen at a time: "did, you, know…" Each word appears for 1s… No dialogue.`
- `There is a street sign that says: "…"`

### 2.7 Image role tags

When `images` are provided, bind roles in the prompt. Prefer **simple tags**:

| Tag | Role |
|-----|------|
| `<FIRST_FRAME>` | Starting frame of the video |
| `<IMAGE_REF_0>`, `<IMAGE_REF_1>`, … | Reference (subject/style/object); indices start at **0**, matching `images` order |

**Examples:**
> `<FIRST_FRAME> a woman is walking`
> `in the style of <IMAGE_REF_0> a woman <IMAGE_REF_1> is walking`
> `[0-3s] … woman <IMAGE_REF_0>, holding <IMAGE_REF_1> …`

**Complex cases — explicit declarations** (optional):
> `[# Sources <FIRST_FRAME>@Image1] [# References <IMAGE_REF_0>@Image2] a woman <IMAGE_REF_0> is walking. Use Image1 as the starting frame. Use Image2 as a reference for the video generation.`

Guiding suffixes when needed:
- `Use this image as the starting frame.`
- `Use the given image(s) as references for video generation. The images should not be used as literal initial frames.`

Never rely on character names alone — describe or tag the visual elements.

### 2.8 Official-style examples (tone to match)

**Single continuous shot + audio:**
> Continuous, unbroken handheld shot of a fluffy tabby cat sitting on a sunny windowsill, looking out into a leafy garden. The cat's tail twitches slowly, and its ears rotate slightly toward ambient noises. Sunbeams illuminate dust motes in the air. Sound design: Gentle breeze, distant bird chirps. No dialogue.

**Physics / motion clarity:**
> A marble rolling fast on a chain reaction style track, continuous smooth shot.

**Image-to-video:**
> turn this into realistic footage, using the drawing only as a guide for movement, do not show the drawing in the final video

**Subject references:**
> A cat playfully batting at a ball of yarn.

**Portrait cyberpunk (pair with `aspectRatio: "9:16"`):**
> A futuristic city with neon lights and flying cars, cyberpunk style

### 2.9 Motion-ready test

Before YAML: *If I had only this prompt, could I storyboard the clip second-by-second?* Missing motion, camera, audio, or shot plan → add them.

**Too thin:**
> A woman in a garden, cinematic, beautiful lighting.

**Motion-ready:**
> In a single continuous shot. A woman wearing a blue floral dress walks through a sunlit garden along a gravel path; the dress moves lightly in the breeze. Smooth tracking shot at eye level following beside her. Warm morning light, soft leaf shadows. Sound design: gentle breeze, distant birds. No dialogue.

### 2.10 Prompt rules

1. **English only** in `prompt` — chat with the user in their language.
2. **Describe motion, not a still** — verbs, camera moves, environment change over time.
3. **Hyper-specific** — appearance, action, camera language, materials, timing.
4. **State purpose** when it shapes pacing/framing (ad, reel, mood film).
5. **Shot plan explicit** when single-take is required.
6. **Audio explicit** when music, silence, or no-dialogue matters.
7. **Simple negatives in-prompt** — `No dialogue`, not a separate field.
8. **Edit prompts stay short** + `Keep everything else the same`.
9. **Length** — usually **80–220 words** for generation; **one short sentence** for many edits.
10. **Intent fidelity** — do not add cuts, cast, or soundtrack the user did not want.
11. **Refs need tags or clear role language** — first frame vs reference must be unambiguous.

---

## Step 3: Output YAML

Write `{video-slug}.yaml` (kebab-case slug from subject) under `videos/` unless the user chose another folder:

```yaml
type: video
model: gemini-omni-flash-preview

params:
  prompt: |
    A woman wearing a blue floral dress
    walks through a sunlit garden.
  images:
    - path: "./references/dress.jpg"
    - path: "./references/woman.png"
  durationSeconds: 8
  aspectRatio: "16:9"
  task: reference_to_video   # optional; inferred from images when omitted
  seed: null

output: "./{slug}.mp4"   # same folder as this YAML by default
```

Omit `images` when unused. Omit `task` only when inference is obvious; set it when image roles could be confused.

### Parameter reference

| Parameter | Values | Default | Notes |
|-----------|--------|---------|-------|
| `model` | `gemini-omni-flash-preview` | `gemini-omni-flash-preview` | Omni Flash (preview) |
| `durationSeconds` | Positive number (e.g. `8`) | Confirm with user; common short clips ~several seconds | Longer clips take more time to generate |
| `aspectRatio` | `"16:9"`, `"9:16"` | `"16:9"` | Landscape default; portrait for Stories/Reels/TikTok-style |
| `task` | `text_to_video`, `image_to_video`, `reference_to_video`, `edit` | Often inferred | Set when inputs are ambiguous |
| `seed` | `null` or integer | `null` | |
| `images` | Array of `{path: "..."}` | `[]` | Order matches `<IMAGE_REF_N>` / first-frame binding; high-res refs work best |
| `output` | Path to `.mp4` | `./{slug}.mp4` next to the YAML | User-specified folder overrides |

### Aspect ratio — choose by purpose

| Use | Prefer |
|-----|--------|
| YouTube, web hero, cinematic landscape | `16:9` |
| Stories, Reels, TikTok, mobile portrait | `9:16` |

### Task selection cheat sheet

| Inputs | Typical `task` |
|--------|----------------|
| Prompt only | `text_to_video` |
| One image + “animate / start from this frame” | `image_to_video` |
| One or more images as wardrobe/character/style/object refs | `reference_to_video` |
| Iterate on a prior Omni generation | `edit` |

---

## Step 4: Review YAML & Ask to Generate

YAML alone does **not** create a clip. Actual files come from **google-genai-mcp**.

1. Show or point to `{slug}.yaml` and confirm task, duration, aspect, refs, audio, `output`.
2. Ask explicitly (user's language), e.g.:  
   > YAML 스펙 검토가 끝났습니다. google-genai-mcp로 실제 영상을 생성할까요? (영상은 시간이 걸릴 수 있습니다.)
3. **Stop** until the user approves or declines.
4. If declined: leave the YAML as the deliverable.
5. If approved: proceed to Step 5.

---

## Step 5: Generate via google-genai-mcp

| Resource | Path |
|----------|------|
| MCP server | `../google-genai-mcp/src/mcp/server.ts` |
| Spec | `../google-genai-mcp/spec/google-genai-mcp.md` |
| Examples | `../google-genai-mcp/examples/` (e.g. `cafe-video.yaml`) |

### Tools

| Tool | Role |
|------|------|
| `generate` | **Primary.** `filePath` → the video YAML. Wait for response with `files` filled |
| `continue_interaction` | Conversational edit: `interactionId` + short instruction (+ `Keep everything else the same`) |
| `list_interactions` / `sync_interactions` | Discover / clean mappings |
| `cancel_interaction` / `delete_interaction` | Stop / remove jobs |

### Video generation call

1. Ensure the approved YAML is on disk (paths relative to that file's directory).
2. Call MCP `generate` with `filePath`. Wait for `{ interactionId, files, … }` with `files` filled (`.mp4`).
3. Report the saved path(s) from `files`.
4. Video can take a while — tell the user you are waiting; offer `cancel_interaction` if they want to abort and an `interactionId` is known.

### Edits

- Short `continue_interaction` text preferred for one-aspect changes.
- Full regen → new/updated YAML + ask again + `generate`.
- Ask before each paid/long generation when iterating.

---

## Step 6: Revise if Needed

1. Identify what changed (motion, camera, audio, timing, refs, duration, aspect, task).
2. For **generation** tweaks: update the full prompt + params as needed.
3. For **edit** turns: short keep-the-rest prompt via YAML `task: edit` or `continue_interaction` — not a new screenplay unless full regen.
4. Re-confirm YAML, then **ask again** before MCP calls.

---

## Limitations (do not promise)

Per Omni docs (preview): audio reference upload unsupported; multi-video reasoning unsupported; video extension/interpolation unsupported; voice editing unsupported; separate negative-prompt field unsupported (use in-prompt avoidances); recognizable-people / minors / regional edit restrictions may apply. Prefer English prompts for reliable results.

Video generation can take a long time. If the MCP client times out before completion, report the failure (an `interactionId` may still exist for recovery).

---

## Prohibitions

- Non-English text in the final `prompt` field
- Keyword dumps or still-image captions without motion/camera/audio
- Vague animate prompts (`make it move`) without motion detail
- Writing YAML before the motion brief is confirmed (unless the request is fully specific)
- Calling `generate` / `continue_interaction` / `download` **without** user approval after YAML review
- Treating empty `files` as a finished local file
- Adding cast, cuts, dialogue, or music the user did not request or clearly imply
- Ignoring reference images or mis-tagging first frame vs reference
- Overlong edit prompts that rewrite the whole scene when a simple change was asked
- Using aspect ratios other than `16:9` / `9:16` for this model
