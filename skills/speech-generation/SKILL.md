---
name: speech-generation
description: >-
  Select when the user asks to generate, synthesize, narrate, voice-over, read
  aloud, or produce speech, TTS audio, podcast dialogue, audiobook lines, or
  other exact-text spoken audio — or needs a Gemini TTS YAML. Also select when
  they ask to interpret, analyze, evaluate, transcribe, describe, or ask
  questions about existing audio (speech and/or non-verbal sound), including
  recovering a script/performance brief for revised TTS — or when analysis of
  audio will feed a later music-generation YAML. Uses Gemini TTS and MCP
  analyze via cocrates-google-genai ([audio understanding](https://ai.google.dev/gemini-api/docs/audio)).
  Do not select for composing music as the primary deliverable (music-generation),
  video (video-generation), still images, diagrams, or Live API conversation.
metadata:
  agent: cocrates
---

# Speech Generation

Exact-text TTS performances ([Speech generation](https://ai.google.dev/gemini-api/docs/speech-generation)) — not Live API conversation.

| Layer | Owns | Must not |
|-------|------|----------|
| **YAML** | `title`, `design`, MCP request | Call MCP before YAML approval (except Express) |
| **MCP** `cocrates-google-genai` | `generate` → audio; optional `analyze` → text | Paraphrase a locked script; auto-analyze |

Spoken words come from an agreed **script** (user-supplied or drafted in chat) and are placed under `#### TRANSCRIPT` in `params.text`. There is no separate `message` / `explanation` field. `design` holds performance decisions and must stay consistent with preamble, voices, and TRANSCRIPT.

## Core Principles

- **Chat then YAML** — discover intent, script, and performance in chat; write YAML when enough is known. If already enough, write YAML immediately.
- **Script-first** — prefer a user-supplied script; if missing, draft and confirm in chat before YAML.
- **YAML gate before generate** — user reviews YAML (`design` primary; TRANSCRIPT matches agreed script), then MCP `generate`.
- **User review by default** — after generate, the user evaluates the audio; AI `analyze` only on explicit request.
- **Director, not announcer** — `params.text` = direction preamble + labeled transcript.
- **Consistency** — `design` ↔ preamble ↔ voices/speakers ↔ TRANSCRIPT.
- **Audio tags always English** — map localized requests (e.g. whisper) to `[whispers]`; never non-English or closing tags.

`title` / `design` use the **user's language**. Transcript language is whatever was agreed; direction preamble is typically English.

## Workflow

```
A Forward
  1 Discover (chat)   intent + script + performance
  2 Write YAML        title / design / MCP block  → review
  3 Approve → Generate   cocrates-google-genai   → audio
  4 Review            user evaluates (default); optional AI analyze on request
  5 Revise            update design + text/voices together; regenerate if needed

B From existing audio (understand → YAML → regenerate)
  analyze audio → transcript + non-verbal cues + performance notes
  → draft script/design (or hand off music brief) → Write YAML → review/approve
  → Generate → Review / Revise as in A
```

**Enough already:** Write complete YAML immediately; stop for approval (unless Express).

**Express** (e.g. *generate now*, *express*): YAML + generate; revise from audio.

Mark only the MCP block with `# --- cocrates-google-genai ---`.

## Paths

Default: `speeches/{slug}.yaml`, `output: "./{slug}.wav"` beside it. Relative paths resolve against **that YAML's directory**.

---

## Understand & revise from media

When the user asks to interpret / analyze / evaluate / transcribe / Q&A **existing** audio ([audio understanding](https://ai.google.dev/gemini-api/docs/audio)):

1. **Analyze** — MCP `analyze` with `inputs: [audio path|URL]` and a prompt that requests what the user needs. Typical recovery material:
   - **Speech transcript** (exact words; speaker labels if multi-speaker)
   - **Non-verbal sound** (laughter, sighs, ambience, SFX, music beds — Gemini can describe these)
   - **Performance cues** (pace, emotion, accent hints) for `design` / director notes
   - Timestamps (`MM:SS`) when segment-level detail matters
2. Present `{ text, interactionId }` to the user. Follow up with `continue_interaction` as needed.
3. **Material for YAML** — lock an agreed **script** (TRANSCRIPT) + `design` from the analysis and user edits. Confirm before locking.
4. **Branch by deliverable:**
   - **Revised TTS** — Write speech YAML (`type: speech`); put agreed words under `#### TRANSCRIPT`; encode performance in `design` / preamble. Review → generate.
   - **Music from audio understanding** — if the goal is a track/BGM/song, hand off to **music-generation** with a brief built from the analysis (mood, non-verbal texture, any lyric content). Do not force Lyria through this skill’s `type: speech`.
5. **Review / approve** → **Generate** → **Review / Revise**.

Do not skip the YAML gate. Analysis text alone is not a generation contract. Apply only user-agreed edits before regenerate.
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

Then Phase 4.

---

## Phase 4 — Review (user default; optional AI analyze)

**Default:** Present the generated audio path(s). The **user** reviews and evaluates. Do not call `analyze` unless asked.

**Optional AI analyze** — only on explicit request (e.g. *analyze*, *evaluate*, *AI review*, *평가해줘*). Never auto-run after generate.

1. Resolve the artifact path from YAML `output` (relative to the YAML directory); verify on disk. Prefer an absolute path (or a path valid against the MCP process CWD).
2. Call MCP **`analyze`**:
   - `inputs`: `[artifact path]` (1–10)
   - `model`: omit unless the user overrides (default `gemini-3.5-flash`)
   - `prompt`: evaluation brief in the **user's language**, grounded in agreed **script** (TRANSCRIPT) + `design` / director notes. Ask for a structured report covering **all** of:
     1. **Intent fit** — does the performance match the designed delivery (`design` / style / pace / accent / voices)?
     2. **Message / script delivery** — is the intended spoken content clear and faithful (no paraphrasing drift; notes not spoken)?
     3. **Functional** — wrong speaker, missing lines, tag mishaps, language/voice mismatches
     4. **Quality / completeness** — clarity, pacing, artifacts, abrupt cuts (chunking), polish gaps
     5. **Improvements** — concrete, prioritized suggestions (what to change in `design` / preamble / voices / tags)
3. Present `{ text, interactionId }` to the user. Do **not** silently edit YAML or regenerate from the report.
4. Follow-ups: `continue_interaction` with the same `interactionId` if the user asks more. New audio → new `analyze` call.

---

## Phase 5 — Revise

| Change | Action |
|--------|--------|
| Script | Update TRANSCRIPT (+ tags); adjust `design` if tone shifts; re-approve; regenerate |
| Performance | Update `design` and preamble/voices together; re-approve; regenerate |
| From analyze / user review | Apply only **user-agreed** changes; then regenerate |
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
- Auto `analyze` without an explicit user request; applying analyze suggestions without user agreement
- More than 2 speakers; mismatched speaker names; unlabeled TRANSCRIPT on directed prompts
- Non-English or closing audio tags; inventing script unless asked to draft
- Full script only in `design` instead of TRANSCRIPT; voices outside the catalog
