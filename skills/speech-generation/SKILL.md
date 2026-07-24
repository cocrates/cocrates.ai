---
name: speech-generation
description: >-
  Generates controllable Gemini TTS prompts and YAML specs
  (gemini-3.1-flash-tts-preview default; 2.5 Flash/Pro TTS optional), then —
  after user approval — creates the audio via google-genai-mcp `generate`
  (sync). Select when the user asks to generate, synthesize, narrate,
  voice-over, read aloud, or produce speech, audio narration, podcast dialogue,
  audiobook lines, or single-/multi-speaker TTS — or when they need a
  Gemini-compatible audio YAML. Directs performance via style, accent, pace,
  tone, audio tags, and voice choice for exact text recitation; only calls MCP
  after explicit approval.
compatibility: opencode
metadata:
  agent: cocrates
---

# Speech Generation

Generate controllable Gemini text-to-speech (TTS) performances and output a YAML specification file — [Speech generation docs](https://ai.google.dev/gemini-api/docs/speech-generation).

TTS here is for **exact text recitation** with fine control over style, accent, pace, and tone (podcasts, audiobooks, VO, dialogue). It is not the Live API (interactive multimodal conversation).

## Core Principles

- **Intent-First**: Understand *who is speaking, what is said, and how it should sound* before writing YAML.
- **Director, not announcer**: Craft a performance brief (profile, scene, director’s notes) plus a clearly labeled **transcript** the model must speak — not a bare string unless the request is trivial.
- **Alignment**: Voice choice, character, writing style, and delivery instructions must agree (*who* / *what* / *how*).
- **YAML as Contract**: First deliverable is a `{speech-slug}.yaml` request file (google-genai-mcp schema).
- **YAML review + approval gate**: After writing the YAML, **stop** for user review. Do **not** create the final media (speech audio) until the user has reviewed the YAML and given **explicit approval**. The user may edit the YAML directly; treat their file as source of truth. Call MCP `generate` only after that approval.

## Workflow

```
Understand Intent → Choose Model & Voice(s) → Confirm Brief → Craft Text Prompt → Output YAML
  → Review YAML → Ask to Generate → (on approval) MCP generate → (Revise if needed)
```

No multi-stage document gates. Ask only what is unclear, confirm the performance brief, write YAML, then **ask before generating**.

## Working Location

### Spec folder (YAML)

When the user does **not** name a location, write the YAML under the workspace kind folder:

```text
speeches/{speech-slug}.yaml
```

If the user specifies a folder or path, use that instead. Prefer an existing `speeches/` tree over inventing a parallel layout.

### Generated artifacts (`output`)

| Rule | Path |
|------|------|
| **Default** | Same directory as the YAML file (e.g. `speeches/{slug}.wav`) — **not** `./output/` |
| **User override** | Exact folder or file path the user named |

**Example layout (defaults):**

```text
speeches/
├── spooky-whisper.yaml
└── spooky-whisper.wav
```

---

## Step 1: Understand Intent

### 1.1 Extract what is already known

| Dimension | Capture |
|-----------|---------|
| **Purpose** | Podcast, audiobook, ad VO, IVR, tutorial, spooky reading, dialogue scene, etc. |
| **Mode** | Single speaker vs multi-speaker (max **2**) |
| **Speakers** | Names, roles, relationship; how each should sound |
| **Transcript** | Exact words to recite (user-supplied or to be drafted) |
| **Language** | Spoken language of the transcript (auto-detected; many languages supported) |
| **Style / tone** | Emotion, energy, formality |
| **Accent** | Prefer specific (e.g. Brixton London, Laguna Beach valley) over vague “British” |
| **Pace** | Fast/slow, bouncing, liquid drift, etc. |
| **Audio tags** | Inline English only: `[whispers]`, `[laughs]`, mid-line shifts (never Korean/other-language tags) |
| **Voice(s)** | From the 30 prebuilt voices (match emotion/archetype) |
| **Format** | `outputFormat` (typically `wav`) |

### 1.2 Infer carefully

- Infer delivery from purpose (morning radio → high energy + vocal smile; bedtime story → soft, slow).
- Do **not** invent dialogue content the user did not ask for unless they requested script drafting.
- If they want a script *and* speech, draft the transcript first (may use a text model conceptually), then wrap it in a TTS-ready `text` field — the YAML still targets the TTS model only.

### 1.3 Choose model

| Model ID | Role |
|----------|------|
| `gemini-3.1-flash-tts-preview` | **Default** — Flash TTS preview; single + multi-speaker; streaming supported |
| `gemini-2.5-flash-preview-tts` | Alternate Flash preview TTS |
| `gemini-2.5-pro-preview-tts` | Pro preview TTS — consider for longer / higher-stakes performances |

**Default:** `gemini-3.1-flash-tts-preview`. Prefer Pro when the user wants maximum performance nuance on longer dramatic reads (still split long scripts — see Limitations).

### 1.4 Choose voice(s)

Match voice character to emotion and role. Examples of pairing:

| Need | Prefer voices like |
|------|---------------------|
| Firm / authoritative | Kore, Orus, Alnilam |
| Upbeat / lively | Puck, Zephyr, Laomedeia, Sadachbia |
| Informative / host | Charon, Rasalgethi, Sadaltager |
| Soft / gentle / whispery scenes | Achernar, Vindemiatrix, Enceladus (breathy) |
| Youthful / breezy | Leda, Aoede |
| Mature / warm | Gacrux, Sulafat |
| Gravelly / casual | Algenib, Zubenelgenubi |
| Excitable | Fenrir |

Full catalog (use exact names in YAML):

| Voice | Character | Voice | Character |
|-------|-----------|-------|-----------|
| Zephyr | Bright | Puck | Upbeat |
| Charon | Informative | Kore | Firm |
| Fenrir | Excitable | Leda | Youthful |
| Orus | Firm | Aoede | Breezy |
| Callirrhoe | Easy-going | Autonoe | Bright |
| Enceladus | Breathy | Iapetus | Clear |
| Umbriel | Easy-going | Algieba | Smooth |
| Despina | Smooth | Erinome | Clear |
| Algenib | Gravelly | Rasalgethi | Informative |
| Laomedeia | Upbeat | Achernar | Soft |
| Alnilam | Firm | Schedar | Even |
| Gacrux | Mature | Pulcherrima | Forward |
| Achird | Friendly | Zubenelgenubi | Casual |
| Vindemiatrix | Gentle | Sadachbia | Lively |
| Sadaltager | Knowledgeable | Sulafat | Warm |

**Multi-speaker:** At most **2** speakers. Speaker `name` in YAML **must match** names used in the transcript (`Joe:`, `Jane:`).

### 1.5 Ask only what is still unclear

| Check | Question (user's language) |
|-------|----------------------------|
| Mode | One narrator or a 2-person dialogue? |
| Script | Exact text to speak, or should we draft it? |
| Style | Tone, accent, pace? |
| Voices | Any preferred voice, or OK to pick for the mood? |
| Language | Which language should be spoken? |
| Output | `wav` OK? |

### 1.6 Confirm the performance brief

2–4 sentences: purpose, speaker(s), voice(s), style/accent/pace, language, model. Proceed on confirm or if already fully specific.

---

## Step 2: Craft the `text` Prompt

The YAML `text` field is the full TTS **input prompt**: direction + spoken words. Gemini TTS follows natural-language direction for style, accent, pace, and tone.

### 2.1 Complexity ladder

| Level | When | Shape |
|-------|------|--------|
| **Simple** | Short line, clear emotion | `Say in a spooky whisper:\n"…"` |
| **Styled dialogue** | 2 speakers, per-speaker mood | Guidance line + `Name: …` turns |
| **Directed performance** | Podcast, character VO, ads | Audio Profile + Scene + Director’s Notes + Sample Context + **labeled Transcript** (+ tags) |

Use the deepest level the request deserves. Do not overspecify — too many rigid rules hurt naturalness.

### 2.2 Directed performance structure (recommended for best results)

Think like a director casting virtual voice talent:

1. **Audio Profile** — Name, archetype, age/background if useful  
2. **Scene** — Place, vibe, what is happening around them  
3. **Director’s Notes** — Style, pacing, accent, breathing, articulation (skip other sections if needed; **prefer always including notes**)  
4. **Sample context** — Why this performance exists (optional but helpful)  
5. **Transcript** — Exact words to speak; writing style must fit the direction  
6. **Audio tags** — Inline `[whispers]`, `[excited]`, etc.

**Critical (Flash 3.1 classifier):** Vague prompts may be rejected or cause the model to **read the director’s notes aloud**. Always:

- Include a clear preamble that this is **speech synthesis / TTS** of the transcript only.  
- **Explicitly label** where the spoken transcript begins (e.g. `#### TRANSCRIPT` or `Speak only the following transcript:`).  
- Keep instructions and spoken text visually separated.

### 2.3 Audio tags

Inline modifiers for delivery and light non-verbals.

**MUST be English.** Audio tags are **always** English square-bracket tokens — even when the spoken transcript is Korean or another language. If the user asks in Korean (e.g. “강조”, “속삭이듯”, “신나게”), **map** to English tags; do **not** transliterate or invent Korean tag names.

| User intent (example) | Write in transcript |
|-----------------------|---------------------|
| 강조 / 힘주어 | `[serious]` or `[excited]` before the phrase (pick by tone) |
| 속삭임 | `[whispers]` |
| 웃음 | `[laughs]` / `[giggles]` |
| 한숨 | `[sighs]` |

**Forbidden tag forms**

- Non-English tags: `[강조]`, `[속삭임]`, `[신나게]`  
- XML-style open/close pairs: `[강조]…[/강조]`, `<강조>…</강조>`  
- Fake markup that is not a Gemini audio tag — tags are **open-only** cues; there is no closing tag

Common tags: `[amazed]`, `[crying]`, `[curious]`, `[excited]`, `[sighs]`, `[gasp]`, `[giggles]`, `[laughs]`, `[mischievously]`, `[panicked]`, `[sarcastic]`, `[serious]`, `[shouting]`, `[tired]`, `[trembling]`, `[whispers]`, plus pace like `[very fast]`, `[very slow]`.

Creative tags work (still English): `[like dracula]`, `[like a cartoon dog]`, `[sarcastically, one painfully slow word at a time]`.

Mid-line shifts:
> `[whispers] Hey there… [shouting] and I can say things in many different ways. [whispers] How can I help you today?`

Korean transcript example (tags English, words Korean):
> `서버 안에는 여러 [serious] 고루틴이 있어요.`

### 2.4 Patterns

**Single — simple style**
```text
Say in a spooky whisper:
"By the pricking of my thumbs...
Something wicked this way comes"
```

**Single — cheerful**
```text
Say cheerfully: Have a wonderful day!
```

**Multi — per-speaker guidance (max 2)**
```text
Make Speaker1 sound tired and bored, and Speaker2 sound excited and happy:

Speaker1: So... what's on the agenda today?
Speaker2: You're never going to guess!
```

**Multi — named turns (names must match `speakers`)**
```text
TTS the following conversation between Joe and Jane:
Joe: How's it going today, Jane?
Jane: Not too bad, how about you?
```

**Directed — skeleton**
```text
# SPEECH SYNTHESIS
Synthesize speech for the character below. Do not read the directions aloud.
Speak only the text under TRANSCRIPT.

# AUDIO PROFILE: {Name}
## "{Archetype}"

## THE SCENE: {Place}
{Environment + vibe + how it affects delivery.}

### DIRECTOR'S NOTES
Style: {specific, vivid — better than one vague adjective}
Pace: {tempo and variation}
Accent: {precise regional description}

### SAMPLE CONTEXT
{Optional: where this VO will be used.}

#### TRANSCRIPT
{Exact words. Optional [audio tags] inline.}
```

**Style notes tips**

- Prefer *“Infectious enthusiasm; listener feels part of a massive event”* over *“energetic”*.  
- Accent: *“British English as heard in Croydon”* / *“Brixton, London”* beats *“British”*.  
- Pace: simple (`as fast as possible`) to complex (`The Drift: incredibly slow and liquid…`).  
- Voiceover terms OK: “vocal smile”, proximity effect, punchy consonants.

### 2.5 Performance-ready test

Before YAML: *Could a voice actor perform this without asking what to say or how?* If direction and transcript are muddled, or transcript unlabeled, fix that.

### 2.6 Prompt rules

1. **`text` holds direction + transcript** — spoken words must be unambiguous.  
2. **Label the transcript** on directed prompts so notes are not spoken.  
3. **Align** script tone with director’s notes and selected voice(s).  
4. **Multi-speaker:** ≤2; names consistent across `text` and `speakers`.  
5. **Audio tags MUST be English** (map Korean requests like “강조” → `[serious]` / `[excited]`); never `[강조]`, `[/강조]`, or other non-English / closing tags. Transcript language as requested.  
6. **Don’t overspecify** — leave room for natural acting.  
7. **Long material:** split into multiple YAML jobs / chunks (quality drifts after a few minutes).  
8. **Exact recitation** — do not casually paraphrase user-provided lines unless asked.

---

## Step 3: Output YAML

### Single speaker

Write under `speeches/` unless the user chose another folder:

```yaml
type: speech
model: gemini-3.1-flash-tts-preview

params:
  text: |
    Say in a spooky whisper:
    "By the pricking of my thumbs...
    Something wicked this way comes"
  voice: Kore
  outputFormat: wav

output: "./{slug}.wav"   # same folder as this YAML by default
```

### Multi-speaker (max 2)

```yaml
type: speech
model: gemini-3.1-flash-tts-preview

params:
  text: |
    Joe: How's it going today, Jane?
    Jane: Not too bad, how about you?
  speakers:
    - name: Joe
      voice: Kore
    - name: Jane
      voice: Puck
  outputFormat: wav

output: "./{slug}.wav"   # same folder as this YAML by default
```

Use **either** top-level `voice` (single) **or** `speakers` (multi) — not both. Speaker names must appear in `text` the same way.

### Parameter reference

| Parameter | Values | Default | Notes |
|-----------|--------|---------|-------|
| `model` | `gemini-3.1-flash-tts-preview`, `gemini-2.5-flash-preview-tts`, `gemini-2.5-pro-preview-tts` | `gemini-3.1-flash-tts-preview` | |
| `text` | Prompt string (direction + transcript) | required | Controllable TTS input |
| `voice` | One of 30 voice names | required for single | e.g. `Kore` |
| `speakers` | Array of `{name, voice}` length 1–2 | required for multi | `name` matches transcript labels |
| `outputFormat` | `wav` (typical) | `wav` | PCM WAV at runtime (~24 kHz mono in API samples) |
| `output` | Path to `.wav` | `./{slug}.wav` next to the YAML | User-specified folder overrides |

Maps conceptually to API `speech_config`: single → `{voice}`; multi → `{speaker, voice}` entries.

---

## Step 4: Review YAML & Ask to Generate

YAML alone does **not** synthesize audio. Actual files come from **google-genai-mcp**.

1. Confirm YAML: model, voice(s)/speakers, `text` (direction + labeled transcript), `outputFormat`, `output`.
2. Ask explicitly (user's language), e.g.:  
   > YAML 스펙 검토가 끝났습니다. google-genai-mcp로 실제 음성 파일을 생성할까요?
3. **Stop** until the user approves or declines.
4. If declined: leave the YAML as the deliverable.
5. If approved: proceed to Step 5.

---

## Step 5: Generate via google-genai-mcp

| Resource | Path |
|----------|------|
| MCP server | `../google-genai-mcp/src/mcp/server.ts` |
| Spec | `../google-genai-mcp/spec/google-genai-mcp.md` |
| Examples | `../google-genai-mcp/examples/` (e.g. `conversation.yaml`) |

### Tools

| Tool | Role |
|------|------|
| `generate` | **Primary.** `filePath` → the speech YAML. Wait for response with `files` filled |
| `continue_interaction` | Follow-up turn with new direction/text |
| `list_interactions` / `sync_interactions` | Discover / clean mappings |
| `cancel_interaction` / `delete_interaction` | Stop / remove jobs |

### Speech generation call

1. Save YAML on disk (`type: speech` — never `audio`).
2. Call MCP `generate` with `filePath`. Wait for `{ interactionId, files, … }` with `files` filled.
3. Report the saved path(s) from `files`.
4. Relative `output` / refs resolve against the **YAML file's directory**.

Use either `params.voice` (single) or `params.speakers` (≤2) — not conflicting setups; MCP prefers `speakers` if both are present.

---

## Step 6: Revise if Needed

1. Identify change: wording, style, accent, pace, tags, voice, speaker count.  
2. Update `text` and/or `voice` / `speakers`.  
3. Re-confirm YAML, then **ask again** before `generate` / `continue_interaction`. Prefer small direction tweaks over rewriting the whole profile unless the character changed.

---

## Limitations (do not promise around them)

- Text in → audio out only (no image/video inputs to TTS).  
- Context window ~32k tokens.  
- Streaming: supported on `gemini-3.1-flash-tts-preview` (not assumed for older TTS variants).  
- Voice may not match prompts that fight the voice’s natural profile — align character with voice.  
- Long outputs (> a few minutes) may drift — chunk scripts.  
- Rare non-audio token / 500 responses — callers should retry.  
- Vague prompts → `PROHIBITED_CONTENT` or spoken stage directions — use synthesis preamble + labeled transcript.

---

## Prohibitions

- More than **2** speakers in one YAML  
- Mismatched speaker names between `text` and `speakers`  
- Leaving directed prompts without a clear **transcript boundary**  
- Choosing a voice that contradicts the written persona without warning the user  
- Keyword-only style (“happy sad fast”) without usable direction when the user wants a rich performance  
- Writing YAML before the brief is confirmed (unless the request is fully specific)  
- Calling `generate` / `continue_interaction` / `download` **without** user approval after YAML review  
- Using `type: audio` (removed — use `speech`)  
- Inventing script content the user did not request (unless they asked you to draft it)  
- Using voice names outside the 30-name catalog  
- Non-English audio tags or closing-tag markup (`[강조]`, `[/강조]`, etc.) — **always English open tags only**
