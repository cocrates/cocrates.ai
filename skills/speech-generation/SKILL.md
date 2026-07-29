---
name: speech-generation
description: >-
  Select when the user asks to generate, synthesize, narrate, voice-over, read
  aloud, or produce speech, TTS audio, podcast dialogue, audiobook lines, or
  other exact-text spoken audio — or needs a Gemini TTS YAML. Uses Gemini TTS
  via cocrates-google-genai (controllable style, accent, pace, tone, voices).
  Do not select for music (music-generation), video (video-generation), still
  images, diagrams, or interactive Live API conversation.
metadata:
  agent: cocrates
---

# Speech Generation

Exact-text TTS performances ([Speech generation](https://ai.google.dev/gemini-api/docs/speech-generation)) — not Live API conversation.

| Layer | Owns | Must not |
|-------|------|----------|
| **YAML** | `title`, `design`, MCP request | Call MCP before YAML approval (except Express) |
| **MCP** `cocrates-google-genai` | Audio from approved YAML | Paraphrase a locked script |

Spoken words come from an agreed **script** (user-supplied or drafted in chat) and are placed under `#### TRANSCRIPT` in `params.text`. There is no separate `message` / `explanation` field. `design` holds performance decisions and must stay consistent with preamble, voices, and TRANSCRIPT.

## Core Principles

- **Chat then YAML** — discover intent, script, and performance in chat; write YAML when enough is known. If already enough, write YAML immediately.
- **Script-first** — prefer a user-supplied script; if missing, draft and confirm in chat before YAML.
- **YAML gate before generate** — user reviews YAML (`design` primary; TRANSCRIPT matches agreed script), then MCP `generate`.
- **Director, not announcer** — `params.text` = direction preamble + labeled transcript.
- **Consistency** — `design` ↔ preamble ↔ voices/speakers ↔ TRANSCRIPT.
- **Audio tags always English** — map localized requests (e.g. whisper) to `[whispers]`; never non-English or closing tags.

`title` / `design` use the **user's language**. Transcript language is whatever was agreed; direction preamble is typically English.

## Workflow

```
1 Discover (chat)   intent + script + performance
2 Write YAML        title / design / MCP block  → review
3 Approve → Generate   cocrates-google-genai   → audio
4 Revise            update design + text/voices together
```

**Enough already:** Write complete YAML immediately; stop for approval (unless Express).

**Express** (e.g. *generate now*, *express*): YAML + generate; revise from audio.

Mark only the MCP block with `# --- cocrates-google-genai ---`.

## Paths

Default: `speeches/{slug}.yaml`, `output: "./{slug}.wav"` beside it. Relative paths resolve against **that YAML's directory**.

---

## Phase 1 — Discover (chat)

Work **script before performance direction**. Skip lengthy Discover when already clear.

### 1.1 Intent & script (requirement analogue)

| Topic | Capture |
|-------|---------|
| Purpose | Podcast, VO, dialogue, audiobook line, ad, … |
| Mode | Single narrator vs multi (max **2** speakers) |
| Language | Spoken language of the transcript |
| Script | Exact words to recite — user-supplied is authoritative; if missing, draft and **confirm** (invent dialogue only if asked to draft) |
| Model | Flash TTS default; Pro optional for longer dramatic reads |

**Fit test:** Is the exact spoken text locked (or clearly delegated for drafting)?

### 1.2 Performance design (maps to `design` + preamble)

| Topic | Capture |
|-------|---------|
| Speakers / roles | Names, relationship; how each should sound |
| Voices | Catalog voice(s) matched to persona |
| Style / tone | Emotion, energy, formality (prefer vivid notes over one vague adjective) |
| Accent | Specific when it matters (e.g. regional), not only “British” |
| Pace | Tempo and variation |
| Scene | Place/vibe affecting delivery |
| Audio tags | English open tags only (`[whispers]`, …); plan mid-line shifts if needed |
| Complexity | Simple one-liner vs directed performance (profile + notes + labeled TRANSCRIPT) |

**Fit test:** Could a voice actor perform this without asking what to say or how?

When coherent → Phase 2.

---

## Phase 2 — Write YAML

```yaml
title: Spooky Macbeth line

design: |
  Single narrator, firm low whisper, slow deliberate pace. Voice Kore.
  Horror reading for a short promo. No second speaker. Flash TTS, wav.

# --- cocrates-google-genai ---
type: speech
model: gemini-3.1-flash-tts-preview
params:
  text: |
    # SPEECH SYNTHESIS
    Synthesize speech for the character below. Do not read the directions aloud.
    Speak only the text under TRANSCRIPT.

    ### DIRECTOR'S NOTES
    Style: Spooky, intimate whisper; controlled dread
    Pace: Slow and deliberate
    Accent: Neutral dramatic English

    #### TRANSCRIPT
    [whispers] By the pricking of my thumbs...
    Something wicked this way comes
  voice: Kore
  outputFormat: wav
output: "./spooky-macbeth.wav"
```

### `design`

User-facing performance rationale (§1.2) — primary YAML review object. Do **not** paste the full script; state that TRANSCRIPT follows the agreed script. Must stay consistent with preamble, `voice` / `speakers`, and TRANSCRIPT.

### MCP request

| Field | Role |
|-------|------|
| `type` | `speech` (never `audio`) |
| `model` | `gemini-3.1-flash-tts-preview` (default); optional `gemini-2.5-flash-preview-tts`, `gemini-2.5-pro-preview-tts` |
| `params.text` | Preamble from `design` + `#### TRANSCRIPT` + agreed script |
| `params.voice` | Single-speaker voice name |
| `params.speakers` | Multi: `[{name, voice}]` (1–2); names match transcript labels |
| `params.outputFormat` | Typically `wav` |
| `output` | Default `./{slug}.wav` |

Use either `voice` or `speakers`, not a conflicting mix.

### `params.text`

Directed (recommended): synthesis preamble + director notes + `#### TRANSCRIPT` + exact script (optional English `[tags]`). Label the transcript so notes are not spoken. Simple shorts may use `Say in a spooky whisper:\n"…"`.

### Voices (exact catalog names)

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

Firm → Kore/Orus; upbeat → Puck/Zephyr; soft → Achernar/Enceladus; host → Charon/Sadaltager.

### Audio tags

English open tags only (`[whispers]`, `[laughs]`, `[excited]`, …). No localized tag names; no closing tags.

### YAML gate

Present **`design`**; confirm TRANSCRIPT matches the agreed script; note voice(s)/model/output. **Stop** (unless Express).

---

## Phase 3 — Generate

MCP **`cocrates-google-genai`**. `generate` with `filePath`; report `files`. Long TRANSCRIPT may be chunked server-side (~4000 UTF-8 bytes).

Limitations: text→audio only; long takes may drift — chunk; vague prompts may speak notes or hit `PROHIBITED_CONTENT` — use preamble + labeled transcript.

---

## Phase 4 — Revise

| Change | Action |
|--------|--------|
| Script | Update TRANSCRIPT (+ tags); adjust `design` if tone shifts; re-approve; regenerate |
| Performance | Update `design` and preamble/voices together; re-approve; regenerate |
| Small tweak | Light direction edit or `continue_interaction`; ask before another paid call |

---

## Multi-speaker example

```yaml
title: Friends phone brag

design: |
  Two Korean women, early 20s, casual phone call. Sua excited (Laomedeia);
  friend envious (Callirrhoe). Fast bouncy pace. Flash TTS, wav.
  TRANSCRIPT follows the agreed Korean dialogue.

# --- cocrates-google-genai ---
type: speech
model: gemini-3.1-flash-tts-preview
params:
  text: |
    # SPEECH SYNTHESIS
    Synthesize a casual phone conversation between two Korean female friends.
    Speak only the text under TRANSCRIPT.

    ### DIRECTOR'S NOTES
    Style: Natural Korean phone call; Sua excited, friend envious and playful
    Pace: Fast and bouncy
    Accent: Standard Korean, informal

    #### TRANSCRIPT
    수아: 야~ 나 여름 방학에 유럽 간다!
    친구: 뭐?! 진짜?! 나도 가고 싶은데ㅠㅠ
  speakers:
    - name: 수아
      voice: Laomedeia
    - name: 친구
      voice: Callirrhoe
  outputFormat: wav
output: "./friends-phone.wav"
```

---

## Prohibitions

- MCP before YAML approval (except Express); `type: audio`
- More than 2 speakers; mismatched speaker names; unlabeled TRANSCRIPT on directed prompts
- Non-English or closing audio tags; inventing script unless asked to draft
- Full script only in `design` instead of TRANSCRIPT; voices outside the catalog
