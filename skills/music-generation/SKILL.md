---
name: music-generation
description: >-
  Generates Lyria 3 music prompts and YAML specs (lyria-3-clip-preview for
  30s clips; lyria-3-pro-preview for full songs), then — after user approval —
  creates the audio via google-genai-mcp `generate` (sync). Select when the
  user asks to generate, compose, produce, or create music, a song, beat,
  soundtrack, BGM, loop, instrumental, or image-inspired track — or when they
  need a Lyria-compatible music YAML. Crafts specific prompts (genre,
  instruments, BPM, key, mood, structure, lyrics, timelines) for optimal Lyria
  3 output; only calls MCP after explicit approval.
compatibility: opencode
metadata:
  agent: cocrates
---

# Music Generation

Generate music with Google Lyria 3 and output a YAML specification file — [Music generation docs](https://ai.google.dev/gemini-api/docs/music-generation).

Lyria 3 produces **44.1 kHz stereo** audio from text (and optionally images), with structural coherence: vocals, timed lyrics, and full arrangements. This skill targets **Lyria 3 Clip / Pro** (not Lyria RealTime streaming).

## Core Principles

- **Intent-First**: Understand genre, length, vocals vs instrumental, and use case before writing YAML.
- **Specificity Wins**: Vague prompts → generic music. Name instruments, BPM, key, mood, and structure.
- **Separate Direction from Lyrics**: Put musical direction in `prompt`; put custom words in `lyrics` (section-tagged). Do not bury lyrics inside unclear prose.
- **YAML as Contract**: First deliverable is a `{music-slug}.yaml` request file (google-genai-mcp schema).
- **YAML review + approval gate**: After writing the YAML, **stop** for user review. Do **not** create the final media (music audio) until the user has reviewed the YAML and given **explicit approval**. The user may edit the YAML directly; treat their file as source of truth. Call MCP `generate` only after that approval.

## Workflow

```
Understand Intent → Choose Model → Confirm Brief → Craft Prompt (+ Lyrics) → Output YAML
  → Review YAML → Ask to Generate → (on approval) MCP generate → (Revise = new generation)
```

No multi-turn in-model editing — Lyria 3 is **single-turn**. Revisions mean a new YAML / new generation. Prefer iterating on Clip before committing to Pro. Write YAML, then **ask before generating**.

## Working Location

### Spec folder (YAML)

When the user does **not** name a location, write the YAML under the workspace kind folder:

```text
music/{music-slug}.yaml
```

If the user specifies a folder or path, use that instead. Prefer an existing `music/` tree over inventing a parallel layout.

### Generated artifacts (`output`, `lyricsOutput`)

| Rule | Path |
|------|------|
| **Default** | Same directory as the YAML file (e.g. `music/{slug}.mp3`, `music/{slug}.lyrics.txt`) — **not** `./output/` |
| **User override** | Exact folder or file path the user named |

Reference images may live beside the YAML (e.g. `music/refs/…`) or wherever the user points.

**Example layout (defaults):**

```text
music/
├── echoes.yaml
├── echoes.mp3
├── echoes.lyrics.txt
└── refs/
    └── desert-sunset.jpg
```

---

## Step 1: Understand Intent

### 1.1 Extract what is already known

| Dimension | Capture |
|-----------|---------|
| **Purpose** | Song, loop, BGM, game cue, trailer, podcast bed, social clip, mood exploration |
| **Length** | ~30s clip vs full song (~1–3 min, prompt-controllable on Pro) |
| **Vocals** | Instrumental only vs sung vocals; lyric language |
| **Genre / blend** | e.g. lo-fi hip hop, dreamy indie pop, cinematic orchestral |
| **Instruments** | Specific when possible (Fender Rhodes, TR-808, fingerpicked acoustic) |
| **BPM / tempo** | Exact BPM or relative (slow ~70, mid-tempo, 120 BPM) |
| **Key / scale** | e.g. G major, D minor |
| **Mood** | nostalgic, aggressive, ethereal, warm, intimate… |
| **Structure** | Section tags and/or `[m:ss - m:ss]` timeline |
| **Lyrics** | User-supplied vs model-written; theme if auto |
| **Images** | Up to **10** mood/reference images |
| **Format** | `mp3` (default) or `wav` (Pro) |

### 1.2 Infer carefully

- Social/Reel bed → often Clip + instrumental + `9:16` video pairing elsewhere; music itself has no aspect ratio.
- Game UI / loop → Clip, instrumental, clear loopable energy.
- Full narrative song with verses → **Pro**.
- Do not invent copyrighted lyrics or “sound exactly like [famous artist]” — safety filters block artist-voice / copyrighted-lyric requests.

### 1.3 Choose model

| Model ID | Best for | Duration | Output |
|----------|----------|----------|--------|
| `lyria-3-clip-preview` | Short clips, loops, prompt previews | **Always 30 seconds** | MP3 |
| `lyria-3-pro-preview` | Full songs (verse / chorus / bridge), timelines, richer arrangements | A couple of minutes (steer via prompt / timestamps) | MP3 (WAV supported) |

**Defaults**

- Unclear length or “quick idea / loop / preview” → **Clip**.
- “Full song”, custom multi-section lyrics, detailed timeline, or image-inspired long ambient → **Pro**.
- **Recommend Clip first** to lock genre/mood/instruments, then upgrade the same brief to Pro.

State the recommendation briefly; let the user override.

### 1.4 Ask only what is still unclear

| Check | Question (user's language) |
|-------|----------------------------|
| Length | 30s clip or full song (~2 min)? |
| Vocals | Instrumental only, or with singing? |
| Genre / mood | Style and feeling? |
| Lyrics | Exact lyrics, or generate them? Language? |
| Structure | Verse/chorus, or a timed arrangement? |
| References | Mood images? (max 10) |
| Model | Clip vs Pro? |

### 1.5 Confirm the music brief

2–4 sentences: purpose, Clip vs Pro, genre/BPM/key/mood, vocals, structure, images, format. Proceed on confirm or if already fully specific.

**Good brief example:**
> Dreamy indie pop, mid-tempo, soft synths + acoustic guitar; custom English verse/chorus/bridge; ~2 min — `lyria-3-pro-preview`, mp3.

---

## Step 2: Craft the Prompt

Write a **music-ready** English (or lyric-language) prompt. Specificity and structure beat poetic vagueness.

### 2.1 Prompt ingredients (include what the brief needs)

| Element | Why | Example |
|---------|-----|---------|
| **Genre** | Core style | `dreamy indie pop`, `lo-fi hip hop`, `jazz fusion` |
| **Instruments** | Timbre / arrangement | `fingerpicked acoustic guitar`, `Fender Rhodes`, `808 bass` |
| **BPM** | Groove | `85 BPM`, `120 BPM`, `slow tempo around 70 BPM` |
| **Key / scale** | Tonality | `in G major`, `D minor` |
| **Mood** | Emotional color | `warm, intimate`, `dark, atmospheric`, `ethereal` |
| **Vocals** | Explicit | `gentle vocals…` or `Instrumental only, no vocals.` |
| **Duration** | Pro only | `Create a 2-minute song…` (Clip is always 30s) |
| **Structure** | Form | `[Verse]` / `[Chorus]` / `[Bridge]` or timestamps |
| **Theme** | If model writes lyrics | `about a summer road trip` |

### 2.2 Patterns

**Clip — instrumental bed**
> A short instrumental acoustic guitar piece. Warm, intimate, fingerpicked. No vocals.

**Clip — detailed beat**
> A 30-second lofi hip hop beat with dusty vinyl crackle, mellow Rhodes piano chords, a slow boom-bap drum pattern at 85 BPM, and a jazzy upright bass line. Instrumental only.

**Pro — full song (direction; lyrics in `lyrics` field)**
> Dreamy indie pop, mid-tempo, soft synths and acoustic guitar. Create a 2-minute song with verse / chorus / bridge.

**Pro — timeline structure**
```text
[0:00 - 0:10] Intro: soft lo-fi beat and muffled vinyl crackle.
[0:10 - 0:30] Verse 1: warm Fender Rhodes and gentle vocals about a rainy morning.
[0:30 - 0:50] Chorus: full band, upbeat drums, soaring synth leads; hopeful lyrics.
[0:50 - 1:00] Outro: fade with piano alone.
```

**Pro — image-inspired**
> An atmospheric ambient track inspired by the mood and colors in the reference images. Instrumental only.
>
> [0:00 - 0:20] Soft pads and distant piano
> [0:20 - 1:00] Add sparse percussion and rising strings
> [1:00 - 1:30] Peak, then fade to piano alone

**Chiptune / game**
> A bright chiptune melody in C Major, retro 8-bit video game style. Instrumental only, no vocals.

**Pop with auto lyrics (language = prompt language)**
> An upbeat, feel-good pop song in G major at 120 BPM with bright acoustic guitar strumming, claps, and warm vocal harmonies about a summer road trip.

**Non-English lyrics**: Write the **prompt in that language** (e.g. French) so lyrics and pronunciation match.

### 2.3 Custom lyrics

- Put user lyrics in YAML `params.lyrics`, not mixed into free prose without labels.
- Use section tags: `[Verse 1]`, `[Chorus]`, `[Bridge]`, `[Intro]`, `[Outro]`.
- Keep `prompt` as musical direction only; say that the song should use the provided lyrics when helpful.
- Never request copyrighted song lyrics or cloning a specific artist’s voice.

**Example `lyrics`:**
```text
[Verse 1]
Walking through the neon glow,
city lights reflect below.

[Chorus]
We are the echoes in the night,
burning brighter than the light.
```

### 2.4 Images

- Up to **10** images in `params.images`.
- Prompt should say how to use them (mood, color, atmosphere) — not “ignore the images.”
- Prefer Pro for image-inspired full pieces; Clip is fine for short visual-mood stingers if the user wants 30s.

### 2.5 Music-ready test

Before YAML: *Could a session musician / producer recreate the intended track from this prompt alone?* If genre, groove, instrumentation, and vocal intent are missing, add them.

**Too thin:** `A nice song.`  
**Music-ready:** `Dreamy indie pop at mid-tempo (~100 BPM) in A major; soft synth pads, fingerpicked acoustic guitar, light brushed drums; intimate male vocal; 2-minute verse/chorus/bridge. No harsh EDM drops.`

### 2.6 Prompt rules

1. Be **specific** — instruments, BPM, key, mood, structure.  
2. State **instrumental vs vocals** explicitly when it matters.  
3. **Separate** musical direction (`prompt`) from custom words (`lyrics`).  
4. Match **prompt language** to desired lyric language.  
5. Use **section tags** and/or **timestamps** for Pro structure/duration.  
6. Clip prompts may say “30-second” but Clip **always** outputs 30s.  
7. Do not promise multi-turn “make the chorus louder” edits — ship a revised full prompt instead.  
8. Prefer Clip iteration before expensive/long Pro runs when exploring.

---

## Step 3: Output YAML

### Text → clip (30s)

Write under `music/` unless the user chose another folder:

```yaml
type: music
model: lyria-3-clip-preview

params:
  prompt: |
    A short instrumental acoustic guitar piece.
    Warm, intimate, fingerpicked. No vocals.
  outputFormat: mp3

output: "./{slug}.mp3"   # same folder as this YAML by default
```

### Full song + custom lyrics (Pro)

```yaml
type: music
model: lyria-3-pro-preview

params:
  prompt: |
    Dreamy indie pop, mid-tempo, soft synths and acoustic guitar.
    Create a 2-minute song with verse / chorus / bridge.
  lyrics: |
    [Verse 1]
    Walking through the neon glow,
    city lights reflect below.

    [Chorus]
    We are the echoes in the night,
    burning brighter than the light.
  outputFormat: mp3
  lyricsOutput: "./{slug}.lyrics.txt"   # same folder as this YAML by default

output: "./{slug}.mp3"
```

### Image inspiration + timeline (Pro)

```yaml
type: music
model: lyria-3-pro-preview

params:
  prompt: |
    An atmospheric ambient track inspired by the mood and colors
    in the reference images. Instrumental only.

    [0:00 - 0:20] Soft pads and distant piano
    [0:20 - 1:00] Add sparse percussion and rising strings
    [1:00 - 1:30] Peak, then fade to piano alone
  images:
    - path: "./refs/desert-sunset.jpg"
    - path: "./refs/city-night.png"
  outputFormat: mp3

output: "./{slug}.mp3"
```

### Parameter reference

| Parameter | Values | Default | Notes |
|-----------|--------|---------|-------|
| `model` | `lyria-3-clip-preview`, `lyria-3-pro-preview` | Infer from length/structure | Clip = 30s; Pro = full song |
| `prompt` | Musical direction (+ timeline) | required | Specific genre/instruments/BPM/mood |
| `lyrics` | Section-tagged lyric text | omit | Custom vocals; keep separate from direction |
| `images` | Array of `{path}` (max 10) | `[]` | Mood / visual inspiration |
| `outputFormat` | `mp3`, `wav` | `mp3` | WAV primarily documented for Pro |
| `lyricsOutput` | Path for returned lyrics/structure text | `./{slug}.lyrics.txt` next to the YAML when used | User-specified folder overrides |
| `output` | Path to audio file | `./{slug}.mp3` (or `.wav`) next to the YAML | User-specified folder overrides |

### Model cheat sheet

| User want | Model | Tips |
|-----------|-------|------|
| Loop, stinger, prompt test | Clip | Fast iterate; always 30s |
| Full song, bridges, custom lyrics | Pro | Duration + sections in prompt/`lyrics` |
| Timed arrangement | Pro | `[0:00 - 0:20] …` blocks |
| Image mood track | Pro (or Clip if 30s OK) | ≤10 images; describe mood link in prompt |
| Instrumental BGM | Either | End with `Instrumental only, no vocals.` |

---

## Step 4: Review YAML & Ask to Generate

YAML alone does **not** produce audio. Actual files come from **google-genai-mcp**.

1. Confirm YAML: Clip vs Pro, prompt, lyrics, images, `outputFormat`, `lyricsOutput`, `output`.
2. Ask explicitly (user's language), e.g.:  
   > YAML 스펙 검토가 끝났습니다. google-genai-mcp로 실제 음악 파일을 생성할까요? (Pro는 시간이 걸릴 수 있습니다.)
3. **Stop** until the user approves or declines.
4. If declined: leave the YAML as the deliverable.
5. If approved: proceed to Step 5.

---

## Step 5: Generate via google-genai-mcp

| Resource | Path |
|----------|------|
| MCP server | `../google-genai-mcp/src/mcp/server.ts` |
| Spec | `../google-genai-mcp/spec/google-genai-mcp.md` |
| Examples | `../google-genai-mcp/examples/` (e.g. `cafe-bgm.yaml`) |

### Tools

| Tool | Role |
|------|------|
| `generate` | **Primary.** `filePath` → the music YAML. Wait for response with `files` filled |
| `continue_interaction` | Allowed by MCP but Lyria is single-turn — prefer a **new YAML + generate** |
| `list_interactions` / `sync_interactions` | Discover / clean mappings |
| `cancel_interaction` / `delete_interaction` | Stop / remove jobs |

### Music generation call

1. Save YAML on disk (`type: music`).
2. Call MCP `generate` with `filePath`. Wait for `{ interactionId, files, … }` with `files` filled (audio; plus lyrics file if `lyricsOutput` set).
3. Report the saved path(s) from `files`.
4. Tell the user you are waiting if Pro is slow.
5. Paths inside YAML are relative to the **YAML file's directory**.

One MCP `generate` = one request file. Multiple tracks → multiple calls.

---

## Step 6: Revise if Needed

Lyria 3 does **not** support conversational multi-turn edit of an existing clip. To change the track:

1. Adjust genre, instruments, BPM, structure, or lyrics in the brief.  
2. Write a new YAML (optionally keep Clip for A/B tests).  
3. Re-confirm, **ask again**, then `generate`.

---

## Limitations (do not promise around them)

- Safety filters block artist-voice cloning and copyrighted lyrics.  
- SynthID audio watermark on all outputs (inaudible).  
- No multi-turn refine of a prior generation.  
- Clip length fixed at 30s; Pro length approximate, steered by prompt/timestamps.  
- Non-deterministic — same prompt can vary between runs.  
- Not Lyria RealTime (streaming jam) — point users there only if they ask for real-time.  
- Pro jobs can take a long time. If the MCP client times out, report the failure.

---

## Prohibitions

- Vague prompts when the user wants a specific sound (“nice music”, “epic track” alone)  
- Mixing unlabeled custom lyrics into direction without `[Verse]` / `[Chorus]` (or a dedicated `lyrics` field)  
- Requesting famous-artist voice clones or copyrighted lyric reproduction  
- Using Pro-only expectations (multi-minute structure) on Clip without warning it will still be 30s  
- Claiming iterative in-place editing of a generated file via follow-up prompts  
- Calling `generate` / `download` **without** user approval after YAML review  
- Treating empty `files` as a finished local file  
- More than 10 reference images  
- Inventing musical requirements that contradict the user’s brief
