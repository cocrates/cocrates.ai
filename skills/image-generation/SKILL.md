---
name: image-generation
description: >-
  Generates AI image prompts and YAML specs for Google Gemini image models
  (gemini-3.1-flash-image default, gemini-3-pro-image optional), then — after
  user approval — creates the image via google-genai-mcp `generate` (sync).
  Select when the user asks to generate, create, draw, illustrate, or design an
  image, picture, illustration, icon, thumbnail, concept art, or visual asset —
  or when they need an image-generation prompt or Gemini-compatible image YAML.
  Captures user intent first, writes a picture-ready prompt, and only calls MCP
  after explicit approval.
compatibility: opencode
metadata:
  agent: cocrates
---

# Image Generation

Generate AI image prompts from user intent and output a YAML specification file compatible with Google Gemini native image generation ([Nano Banana / Gemini image docs](https://ai.google.dev/gemini-api/docs/image-generation)).

## Core Principles

- **Intent-First**: Understand *what the user actually wants pictured* before writing anything. Never invent a scene that conflicts with stated or clearly implied intent.
- **Prompt as Picture**: The prompt must describe the image so completely that a reader can *see* it in their mind — subject, place, action, composition, look — without guessing.
- **Prompt as Architecture**: Structured natural-language prose, not a keyword dump. Prefer hyper-specific visible facts; include purpose/context when it shapes the look.
- **YAML as Contract**: First deliverable is a `{image-slug}.yaml` request file (google-genai-mcp schema).
- **YAML review + approval gate**: After writing the YAML, **stop** for user review. Do **not** create the final media (image) until the user has reviewed the YAML and given **explicit approval**. The user may edit the YAML directly; treat their file as source of truth. Call MCP `generate` only after that approval.

## Workflow

```
Understand Intent → Choose Model → Confirm Brief → Generate Prompt → Output YAML
  → Review YAML → Ask to Generate → (on approval) MCP generate → (Revise if needed)
```

No multi-stage document gates. Ask only what is unclear, confirm the visual brief, write YAML, then **ask before generating**.

## Working Location

### Spec folder (YAML)

When the user does **not** name a location, write the YAML under the workspace kind folder:

```text
images/{image-slug}.yaml
```

If the user specifies a folder or path, use that instead. Prefer an existing `images/` tree over inventing a parallel layout.

### Generated artifacts (`output`)

| Rule | Path |
|------|------|
| **Default** | Same directory as the YAML file (e.g. `images/{slug}.png`) — **not** `./output/` |
| **User override** | Exact folder or file path the user named |

Reference images may live beside the YAML (e.g. `images/references/…`) or wherever the user points.

**Example layout (defaults):**

```text
images/
├── cafe-cat.yaml
├── cafe-cat.png
└── references/
    └── mood.jpg
```

---

## Step 1: Understand Intent

The quality of the image depends on intent fidelity. Do not skip this step when the request is ambiguous.

### 1.1 Extract what is already known

From the user's message (and any attached images or prior context), extract:

| Dimension | Capture |
|-----------|---------|
| **Purpose** | Why the image exists (thumbnail, blog hero, icon, concept art, social post, slide, product mock, logo, infographic, personal use) |
| **Subject** | Who/what is the focal element |
| **Action / state** | What the subject is doing or how it appears |
| **Setting** | Place, time, atmosphere |
| **Style** | Medium and look (photo, anime, watercolor, flat vector, 3D render, sticker, etc.) |
| **Mood** | Emotion or tone (calm, dramatic, playful, serious) |
| **Text in image** | Exact words/phrases to render, font feel, placement |
| **Constraints** | Must-include / must-avoid elements, brand colors, reference images |
| **Format** | Model, `size`, `aspectRatio` |

Treat explicit user wording as authoritative. Prefer the user's words over generic aesthetic defaults.

### 1.2 Infer carefully; do not over-invent

- Fill only what is **strongly implied** by purpose and context (e.g. YouTube thumbnail → bold subject, high contrast, readable at small size; app icon → simple silhouette, centered, clean background).
- Do **not** add decorative subjects, props, or styles the user did not ask for.
- If purpose is clear but subject/style is vague, ask — do not guess a full scene.

### 1.3 Choose model

Supported models (user may pick either; default is Flash):

| Model ID | Role | When |
|----------|------|------|
| `gemini-3.1-flash-image` | **Default** — versatile workhorse | Everyday generation/editing, speed, volume, solid text, multi-reference consistency, 0.5K–4K |
| `gemini-3-pro-image` | Premium / Pro | Complex or professional assets needing max world knowledge, brand fidelity, precise creative control, advanced text/layout |

**Default:** `gemini-3.1-flash-image`.

**Recommend Pro** (state why briefly; let the user accept or keep Flash) when the request involves any of:

- Professional print/marketing/brand assets with strict fidelity
- Dense or critical **text rendering** (magazines, menus, UI mock copy, multi-line infographics)
- Precise **composition + typography** layouts (weather cards, editorial pages, labeled diagrams)
- Strong **brand / logo / product** consistency across references
- Complex **multi-style or multi-character** scenes, comics/storyboards with readable lettering
- High-stakes commercial photography or packaging where small details matter

Do not force Pro for simple icons, casual scenes, or quick drafts — Flash is enough.

### 1.4 Ask only what is still unclear

Ask the minimum set of questions. Skip anything already answered.

| Check | Question (ask in the user's language) |
|-------|----------------------------------------|
| Subject | What is the main subject or scene? |
| Setting | What background or situation? |
| Style | What visual style? (photo, animation, watercolor, sticker, etc.) |
| Mood / purpose | What mood or where will this image be used? |
| Text | Any words that must appear in the image? Exact spelling? |
| References | Do you have reference images? |
| Model | Flash (default) or Pro? (Recommend Pro only when justified above.) |
| Size | Size? (`0.5K` Flash-only / `1K` / `2K` / `4K`, default: `1K`) |
| Aspect | Aspect ratio? (see table below; default: `1:1` unless purpose implies otherwise) |

Use Socratic clarification when the request is vague: restate your understanding, then ask one focused question at a time when possible.

### 1.5 Confirm the visual brief

Before generating the prompt, briefly restate the intended image in 2–4 sentences covering subject, setting, style, mood, **model**, and format. Proceed only when the user confirms — or when the original request was already fully specific and unambiguous.

**Good brief example:**
> A single orange tabby cat sitting on a wooden cafe table by a window, warm afternoon light, cinematic photography, cozy calm mood — `gemini-3.1-flash-image`, 1K, 16:9 — for a blog header.

---

## Step 2: Generate Prompt

Write a vivid English paragraph (or short step-by-step block for complex scenes) that could stand alone as a description of the finished image. A stranger reading only the prompt should be able to imagine the picture clearly.

Follow Google Gemini prompting strategies: be hyper-specific, state purpose/context, control the camera, use semantic positives, and break complex scenes into ordered steps when needed.

### 2.1 Structure (default)

Follow this order (natural prose, not a labeled list):

> **[Subject & appearance] + [Action / pose] + [Setting & background] + [Composition & camera] + [Style] + [Lighting] + [Texture, color, mood, finish] + [Exact on-image text, if any]**

| Layer | What to spell out | Thin ❌ | Picture-ready ✅ |
|-------|-------------------|--------|------------------|
| **Subject** | Who/what, type, clothing/materials, colors, distinctive features | `a cat` | `a fluffy orange tabby cat with white chest fur and green eyes` |
| **Action / pose** | Body position, gesture, gaze, interaction with objects | `sitting` | `sitting upright on a wooden table, front paws together, looking toward the window` |
| **Setting** | Place, time of day, weather, furniture, props in frame | `in a cafe` | `in a small cozy coffee shop with mismatched chairs, steam rising from a ceramic cup beside the cat, rain-streaked window glass` |
| **Composition** | Framing, angle, lens, depth, focal point | *(omitted)* | `medium shot, subject slightly left of center, shallow depth of field, soft bokeh of cafe lights behind` |
| **Style** | Medium and artistic look | `nice photo` | `cinematic photography, naturalistic color grade` |
| **Lighting** | Source, direction, quality, color temperature | `nice lighting` | `warm late-afternoon side light from the window, gentle rim light on fur` |
| **Details** | Materials, atmosphere, finish | `highly detailed` | `visible wood grain on the table, soft dust motes in the light beam, rich detail` |
| **Text** | Exact string, font feel, placement, what else is lettered | `with a title` | `large bold serif title "London" centered at top; no other text` |

### 2.2 Prompt patterns by use case

Pick the closest pattern and fill it with picture-ready detail. Adapt; do not paste empty brackets.

**Photorealistic scene**
> A photorealistic [shot type] of [subject] in [setting]. [Lighting]. Shot from [camera angle] with a [lens]. [Mood / purpose].

**Stylized illustration / sticker / icon**
> A [style] of [subject with accessories/actions] doing [activity]. The design features [outlines, shading, palette]. Background [color/treatment]. [No text / exact text].

**Accurate text in image** (prefer Pro for professional lettering)
> Create a [image type] for [brand/concept] with the text "[exact text]" in a [font style]. Design: [layout, shapes]. Color scheme: [colors]. [Placement and hierarchy].

**Product mockup / commercial photo**
> A high-resolution, studio-lit product photograph of [product] on [surface/background]. Lighting: [setup] to [effect]. Camera: [angle] to showcase [feature]. Ultra-realistic, sharp focus on [detail].

**Minimal / negative-space design**
> A minimalist [asset type] of [subject]. Large empty [color] negative space for [copy/UI]. Subject occupies [corner/region]; clean, uncluttered.

**Complex / multi-element (step-by-step)**
> First, [background]. Then, [midground subjects]. Finally, [foreground / text / focal prop]. Keep [style/lighting] consistent throughout.

**Grounded / real-time info** (when search-backed imagery is intended)
> Use search to find [facts]. Visualize [chart/scene] as [style]. Include [labels/dates] with clear hierarchy.

### 2.3 Official-style examples (tone to match)

**Magazine cover (rich scene + exact text):**
> A photo of a glossy magazine cover, the minimal blue cover has the large bold words Nano Banana. The text is in a serif font and fills the view. No other text. In front of the text there is a portrait of a person in a sleek and minimal dress. She is playfully holding the number 2, which is the focal point. Put the issue number and "Feb 2026" date in the corner along with a barcode. The magazine is on a shelf against an orange plastered wall, within a designer store.

**Photorealistic environment:**
> A photorealistic wide-angle shot of a vibrant coral reef teeming with tropical fish. Crystal-clear turquoise water with sunbeams filtering down from the surface, illuminating a sea turtle gliding gracefully over the coral. Shot from a low perspective with a wide-angle lens.

**Sticker:**
> A kawaii-style sticker of a happy red panda wearing a tiny bamboo hat. It's munching on a green bamboo leaf. The design features bold, clean outlines, simple cel-shading, and a vibrant color palette. The background must be white.

**Logo with text:**
> Create a modern, minimalist logo for a coffee shop called 'The Daily Grind'. The text should be in a clean, bold, sans-serif font. The color scheme is black and white. Put the logo in a circle. Use a coffee bean in a clever way.

**Product:**
> A high-resolution, studio-lit product photograph of a minimalist ceramic coffee mug in matte black, presented on a polished concrete surface. The lighting is a three-point softbox setup designed to create soft, diffused highlights and eliminate harsh shadows. The camera angle is a slightly elevated 45-degree shot to showcase its clean lines. Ultra-realistic, with sharp focus on the steam rising from the coffee.

### 2.4 Picture-ready test

Before writing YAML, ask: *If I had only this prompt, could I sketch the image?* If key elements (who, where, doing what, how it looks, how it is framed, exact text) are missing, add them.

**Too thin:**
> A fluffy orange cat in a cozy coffee shop, cinematic photography, golden hour light, highly detailed.

**Picture-ready:**
> A fluffy orange tabby cat with white chest fur and bright green eyes sits upright on a worn oak cafe table, front paws neatly together, gazing toward a rain-streaked window on the left. Beside the cat a ceramic cup of coffee releases thin steam. The small coffee shop behind is softly out of focus — mismatched wooden chairs, warm pendant lights, a chalkboard menu on the far wall. Medium shot, subject slightly left of center, shallow depth of field. Cinematic photography with a naturalistic color grade. Warm late-afternoon side light from the window wraps the cat's fur in a gentle rim glow; dust motes drift in the light beam. Cozy, quiet mood; rich detail and clear texture on fur and wood.

### 2.5 Reference image prompts

When a reference image is provided, **do not rely on names or implied identity**. Use **explicit element-level commands**.

**Single reference:**
> Using the provided image, keep [elements to preserve] exactly the same. Change only [elements to modify] to [new descriptions]. Preserve the original [style/lighting/composition].

**Multiple references:**
> Take [specific element] from image 1 and [combine/place/merge] it with [specific element] from image 2. [Additional instructions for the result].

Flash: up to ~10 high-fidelity object refs + up to 4 character refs (within 14 total). Pro: up to 6 object + 5 character + 3 style refs (within 14 total). Prefer Pro when brand/style/character fidelity is critical.

### 2.6 Prompt rules

1. **English only** in `prompt` — chat with the user in their language.
2. **Describe the picture, not a wish list** — concrete visible facts over empty adjectives (`dramatic`, `beautiful`, `epic` alone are not enough).
3. **Hyper-specific** — materials, colors, poses, spatial relations, camera language (`wide-angle`, `macro`, `low-angle`, `45° isometric`).
4. **State purpose when it shapes design** — e.g. "logo for a high-end minimalist skincare brand" beats "make a logo".
5. **Positive / semantic negatives** — prefer `an empty deserted street` over `no cars`. For lettering, naming exact text and `No other text` is allowed when needed (matches Gemini text-rendering practice).
6. **Important content first** — subject and action near the start.
7. **Length** — usually **100–250 words**. Shorter is fine for icons/logos if every visible element is named; use step-by-step blocks for dense scenes.
8. **Intent fidelity** — do not swap subject, add competing focal points, or change style/mood after confirmation.
9. **No keyword dumps** — flowing scene-describing prose.
10. **Reference images need element-level commands** — never rely on names alone.

---

## Step 3: Output YAML

Write `{image-slug}.yaml` (kebab-case slug from subject) under `images/` unless the user chose another folder:

```yaml
type: image
model: gemini-3.1-flash-image   # or gemini-3-pro-image

params:
  prompt: |
    {picture-ready English scene description}
  images:
    - path: "./references/{ref}.png"   # only if references exist; relative to this YAML
  size: 1K
  aspectRatio: "16:9"
  seed: null

output: "./{slug}.png"   # same folder as this YAML by default
```

Maps to Gemini `response_format`: `image_size` ← `size`, `aspect_ratio` ← `aspectRatio`.

### Parameter reference

| Parameter | Values | Default | Notes |
|-----------|--------|---------|-------|
| `model` | `gemini-3.1-flash-image`, `gemini-3-pro-image` | `gemini-3.1-flash-image` | Recommend Pro per §1.3 when justified |
| `size` | `0.5K`, `1K`, `2K`, `4K` | `1K` | `0.5K` (512px) is **Flash only**. Pro: `1K` / `2K` / `4K` |
| `aspectRatio` | See tables below | `1:1` | Default square unless purpose implies otherwise; with a sole reference and no ratio set, APIs often match the input |
| `seed` | `null` or integer | `null` | |
| `images` | Array of `{path: "..."}` | `[]` | Max 14 refs; fidelity mix differs by model |
| `output` | Path to `.png` (etc.) | `./{slug}.png` next to the YAML | User-specified folder overrides |

### Aspect ratio — choose by purpose

| Use | Prefer |
|-----|--------|
| Icon, avatar, logo | `1:1` |
| Photo print / blog still | `4:3` or `3:2` |
| Portrait photo / poster | `3:4` or `2:3` |
| Social portrait / stories | `4:5` or `9:16` |
| Landscape / video / hero | `16:9` or `3:2` |
| Ultrawide / cinematic banner | `21:9` |
| Tall / wide strip banners | Flash: `1:4`, `1:8`, `4:1`, `8:1` |

### Aspect ratios by model

**Flash (`gemini-3.1-flash-image`)** — all of:
`1:1`, `1:4`, `1:8`, `2:3`, `3:2`, `3:4`, `4:1`, `4:3`, `4:5`, `5:4`, `8:1`, `9:16`, `16:9`, `21:9`

**Pro (`gemini-3-pro-image`)** —
`1:1`, `2:3`, `3:2`, `3:4`, `4:3`, `4:5`, `5:4`, `9:16`, `16:9`, `21:9`

(Do not use Flash-only ultra-strip ratios with Pro.)

### Size guidance

| Size | When |
|------|------|
| `0.5K` | Fast drafts/previews (Flash only) |
| `1K` | Default everyday use |
| `2K` | Delivery assets, slides, web heroes |
| `4K` | Print, large posters, max detail (costlier) |

---

## Step 4: Review YAML & Ask to Generate

YAML alone does **not** create pixels. Actual files come from **google-genai-mcp**.

1. Show or point to the written `{slug}.yaml` and briefly confirm it matches the brief (model, size, aspect, prompt, refs, `output`).
2. Ask explicitly (user's language), e.g.:  
   > YAML 스펙 검토가 끝났습니다. google-genai-mcp로 실제 이미지를 생성할까요?
3. **Stop** until the user approves or declines. Do not call MCP on silence or assumed consent.
4. If declined: leave the YAML as the deliverable; revise only if they ask.
5. If approved: proceed to Step 5.

---

## Step 5: Generate via google-genai-mcp

| Resource | Path |
|----------|------|
| MCP server | `../google-genai-mcp/src/mcp/server.ts` |
| Spec | `../google-genai-mcp/spec/google-genai-mcp.md` |
| Examples | `../google-genai-mcp/examples/` (e.g. `kwon-su-a.yaml`, `kwon-su-a-cafe.yaml`) |

### Tools

| Tool | Role |
|------|------|
| `generate` | **Primary.** `filePath` → the image YAML. Wait for response with `files` filled |
| `continue_interaction` | Multi-turn edit: prior `interactionId` + new instruction text |
| `list_interactions` / `sync_interactions` | Discover / clean local interaction mappings |
| `cancel_interaction` / `delete_interaction` | Stop in-progress / remove unwanted jobs |

### Image generation call

1. Ensure the approved YAML is on disk (relative paths resolve against **that file's directory**).
2. Call MCP `generate` with `filePath`. Wait for `{ interactionId, files, … }` with `files` filled.
3. Report saved path(s) from `files` (respect YAML `output`, e.g. `./{slug}.png` next to the YAML).

### Edits after a file exists

- Prefer updating the YAML and a new `generate`, or `continue_interaction` with a short keep-the-rest instruction when iterating on a prior `interactionId`.
- Still ask before spending another generation call when cost/time matters.

---

## Step 6: Revise if Needed

If the user requests changes:

1. Identify which intent dimension changed (subject, setting, style, mood, text, model, format, constraints).
2. Update the brief mentally, then edit the YAML prompt and params to match.
3. Re-confirm the updated file, then **ask again** before calling `generate` / `continue_interaction`.

Prefer small targeted edits ("Keep everything the same, but…") over full rewrites unless subject or purpose changed.

---

## Prohibitions

- Using non-English text in the final `prompt` field
- Keyword-dump style (must be natural scene-describing prose)
- Thin prompts that name a subject and style but leave pose, setting, and composition to guesswork
- Writing YAML before the visual brief is confirmed (skip confirmation only when the request is fully specific)
- Calling `generate` / `continue_interaction` / `download` **without** user approval after YAML review
- Adding subjects, props, text, or styles not requested or clearly implied
- Ignoring attached reference images when the user provided them
- Silently forcing Pro for trivial requests, or silently skipping a Pro recommendation when text/brand/layout fidelity clearly needs it
- Using `0.5K` or Flash-only aspect ratios with `gemini-3-pro-image`
