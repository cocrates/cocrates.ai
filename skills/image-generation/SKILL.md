---
name: image-generation
description: >-
  Select when the user asks to generate, create, draw, illustrate, or design an
  image, picture, illustration, icon, thumbnail, concept art, or other still
  visual asset — or needs a Gemini-compatible image YAML. Also select when they
  ask to interpret, analyze, evaluate, critique, describe, or ask questions
  about an existing still image (including using analysis to revise and
  regenerate via YAML). Uses Gemini image models and MCP analyze via
  cocrates-google-genai. Do not select for architecture/flow diagrams
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
| **MCP** `cocrates-google-genai` | `generate` → pixels; optional `analyze` → text | Invent meaning absent from YAML; auto-analyze |

Pipeline: **`message`** (story/beats to convey) → **`design`** (how to visualize it) → **`params.prompt`** (English realization). Not an interpretation of the finished image. (Diagrams use `explanation` — a pointing/interpretation script — instead.)

## Core Principles

- **Message-first** — the image must make `message` (story/beats) graspable at a glance.
- **Design shapes the message** — `design` materializes the story; it is not a separate decorative brief.
- **Chat then YAML** — gather what the YAML needs in chat; write when enough is known. If already enough, write YAML immediately.
- **YAML gate before generate** — user reviews YAML (`design` primary), then MCP `generate`.
- **User review by default** — after generate, the user evaluates the artifact; AI `analyze` only on explicit request.
- **Consistency** — requirement fields aligned; `design` ↔ `params.prompt` ↔ `params.references` aligned.
- **Intent fidelity** — do not invent unrequested subjects, props, text, or styles.

Human YAML fields (`title`, `summary`, **`message`**, `design`) use the **user's language** — never default them to English when the user writes in another language. `params.prompt` is **English only**.

## Workflow

```
A Forward
  1 Discover (chat)   requirement + visual design
  2 Write YAML        title / [summary] / message / design / MCP block  → review
  3 Approve → Generate   cocrates-google-genai                              → image
  4 Review            user evaluates (default); optional AI analyze on request
  5 Revise            update YAML; keep fields consistent; regenerate if needed

B From existing media (understand → YAML → regenerate)
  analyze image → draft message/design from findings → Write YAML (refs source)
  → review/approve → Generate → Review / Revise as in A
```

**Enough already:** Write complete YAML immediately; stop for approval (unless Express).

**Express** (e.g. *generate now*, *express*): YAML + generate; revise from the image.

Mark only the MCP block with `# --- cocrates-google-genai ---`.

## Paths

Default: `images/{slug}.yaml`, `output: "./{slug}.png"` beside it. Relative paths resolve against **that YAML's directory**. Optional series layout: `characters/`, `locations/`.

---

## Understand & revise from media

When the user asks to interpret / analyze / evaluate / Q&A an **existing** image (not only a just-generated one):

1. **Analyze** — MCP `analyze` with `inputs: [image path|URL]` and a prompt for the user’s question (description, critique, fidelity, Q&A). Present `{ text, interactionId }`. Follow up with `continue_interaction` as needed.
2. **Material for YAML** — from the analysis (and user intent), draft `message` + `design` (keep/change vs source). Confirm with the user before locking.
3. **Write YAML** — full request YAML; cite the source image in `design` and, when editing in place, in `params.references` as a ref. Prefer a new slug or explicit overwrite path.
4. **Review / approve** → **Generate** → **Review / Revise** (Phases 2–5).

Do not skip the YAML gate: analysis text alone is not a generation contract. Apply only user-agreed edits into YAML before regenerate.
---

## Phase 1 — Discover (chat)

Do not write partial YAML here. Work **requirement before visual design**.

### 1.1 Requirement

| Field | Role | Required |
|-------|------|----------|
| `title` | Short image name | Yes |
| `summary` | One-sentence claim | No |
| `message` | Communicative story / beats — what the image should *mean* | Yes |

**`message` (story / beats):** The takeaway narrative the viewer should receive — claims, emotion, relationship, plot beats, or on-image text *as meaning*. Write it in the **user's language** so someone could retell the intended story without sketching the picture.

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
# Human fields (title / message / design) = user's language; params.prompt = English
title: 카페 고양이
message: |
  일상의 고요한 쉼: 카페 안의 온기와 고요, 부드러운 동반감 —
  분주한 가게가 아니라 서두르지 않는 평온을 느끼게 한다.

design: |
  흰 가슴·초록 눈의 복슬복슬한 주황 태비. 낡은 오크 카페 테이블 위에
  앉고, 앞발을 모은 채 왼쪽 빗물 창을 바라봄. 옆에는 김이 나는 컵.
  부드러운 카페 보케. 미디엄 샷, 시네마틱 사진, 늦은 오후 측면광.
  텍스트 없음. Flash, 1K, 16:9, 블로그 헤더.

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
| `params.references` | Ordered `{path, type?}` refs (**image only**; omit `type` or set `image`) |
| `params.size` / `aspectRatio` / `seed` | Format |
| `output` | Default `./{slug}.png` |

**Do not use `params.images`** — removed; MCP returns `INVALID_INPUT` (`use params.references`).

Recommend Pro for brand/text/layout fidelity or complex multi-character work. Do not use Flash-only `0.5K` or strip ratios with Pro.

### Prompt & references

- Prompt: subject → action → setting → composition → style → lighting → details → on-image text; concrete English; ~100–250 words.
- Bind refs with `the provided image` or `Image 1` / `Image 2` … matching `params.references` order; names only after binding.
- Resolve YAML-sourced refs via source `output` → verify on disk → path relative to **this** YAML; if missing, search and get user approval before locking.
- Multi-ref: ref map → SCENE → optional TEXT OVERLAY. Max ~14 refs (Nano Banana 2: up to 19).

| Param | Notes |
|-------|-------|
| `size` | `0.5K` (Flash only), `1K` (default), `2K`, `4K` |
| `aspectRatio` | Default `1:1` unless purpose implies otherwise |
| `references` | Ordered `{path}`; order = `Image N`. **Not** `images` |

### YAML gate

Present **`design`** (and refs); note model/size/aspect/output. Ask approval; **stop** (unless Express).

---

## Phase 3 — Generate

MCP **`cocrates-google-genai`** (GetMcpTools; `mcp_auth` if needed).

1. YAML on disk.
2. Preflight every `params.references[].path` against **this YAML's directory**; on fail, stop and ask — do not `generate`.
3. `generate` with `filePath` → report `files` (or `download` if background).

Then Phase 4.

---

## Phase 4 — Review (user default; optional AI analyze)

**Default:** Present the generated image path(s). The **user** reviews and evaluates. Do not call `analyze` unless asked.

**Optional AI analyze** — only on explicit request (e.g. *analyze*, *evaluate*, *AI review*, *평가해줘*). Never auto-run after generate. *(Architect / automation agents that remapped this gate: always run analyze with the checklist below.)*

1. Resolve the artifact path from YAML `output` (relative to the YAML directory); verify on disk. Prefer an absolute path (or a path valid against the MCP process CWD).
2. Call MCP **`analyze`**:
   - `inputs`: `[artifact path]` plus **reference images used in the YAML** when present (location/character/staging) so geometry can be compared (1–10 total).
   - `model`: omit unless the user overrides (default `gemini-3.5-flash`)
   - `prompt`: evaluation brief in the **user's language**, grounded in approved `message` + `design` (and `summary` if present). **Quote** camera/POV, must-show, and must-hide/occlusion constraints from `design` — do not rely on a vague “evaluate this image.” Ask for a structured report covering **all** of:
     1. **Intent fit** — does the image match the designed intent (`design` / prompt realization)?
     2. **Message delivery** — is `message` (story/beats) clearly conveyed?
     3. **Spatial geometry (mandatory when the design implies place/camera)** — each item **PASS/FAIL + visual evidence**:
        - **Solid architecture** — walls/doors/floors that should exist are present (no vanished wall / open void)
        - **Occlusion / line-of-sight** — people or objects behind an opaque barrier from this POV are **not** visible (no X-ray / see-through wall; outdoor shot must not reveal indoor-only cast through missing architecture)
        - **Camera consistency** — visible cast matches what that camera can see
        - **Layout** — left/right, near/far, beside/behind, indoor vs outdoor match `design`
     4. **Functional** — missing/wrong subjects, text errors, ref fidelity, format issues
     5. **Quality / completeness** — artifacts, composition, clarity, polish gaps
     6. **Improvements** — concrete, prioritized suggestions (what to change in `design` / prompt), including **explicit negative constraints** for geometry Fails (e.g. “opaque wall must fully block caregiver; do not show interior figure”)
     7. **OVERALL** — `PASS` only if every High spatial/functional checkpoint is PASS; else `FAIL`
   - **Anti-vagueness:** Reject mood-only summaries. If a hide/show constraint was quoted and the report does not answer whether the hidden subject is visible → treat as incomplete and re-ask that checkpoint.
3. Present `{ text, interactionId }` to the user. Do **not** silently edit YAML or regenerate from the report *(unless an automation agent owns revise).*
4. Follow-ups: `continue_interaction` with the same `interactionId` if the user asks more questions. New media → new `analyze` call.

---

## Phase 5 — Revise

Keep `design` ↔ prompt ↔ `references` consistent; re-approve; regenerate. Prefer small edits. If `message` changes, reshape `design` to match.

When improving from user feedback or an AI analyze report: apply only **user-agreed** changes; update owning YAML fields, then regenerate.

---

## Multi-ref example

```yaml
# Human fields = user's language; params.prompt = English
title: 4쪽 — 조수가 도움을 청하다
message: |
  위험이 다가온다: 조수가 영웅에게 도움을 구하고, 영웅은 응답해야 한다.
  짧은 화면 속 글이 어린 독자를 위한 대사 비트를 나른다.

design: |
  참조 (params.references 순서):
  1. characters/hero-adventure.yaml → ./characters/hero-adventure.png — 영웅 외형/복장 유지
  2. characters/sidekick.yaml → ./characters/sidekick.png — 조수 외형 유지
  3. locations/plaza-peaceful.yaml → ./locations/plaza-peaceful.png — 광장 유지
  장면: 조수가 영웅에게 달려감; 영웅은 놀람; 먼지/위험. 텍스트 오버레이. 4:3, 1K.

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
  references:
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
- Auto `analyze` without an explicit user request; applying analyze suggestions without user agreement
- `message` that duplicates `design` (shot lists, wardrobe, camera, lighting)
- Defaulting `message` / `design` / `title` to English when the user writes in another language
- Inconsistent `design` / prompt / `references`; non-English prompt
- Using `params.images` (removed — use `params.references`)
- Guessing ref paths; generate after failed preflight
- Inventing unrequested content; ignoring refs; Flash-only options on Pro
