---
name: music-generation
description: >-
  Select when the user asks to generate, compose, produce, or create music, a
  song, beat, soundtrack, BGM, loop, instrumental, or image-inspired track — or
  needs a Lyria-compatible music YAML. Also select when they ask to interpret,
  analyze, evaluate, or describe existing audio/music (or reuse speech/audio
  analysis) in order to draft a music YAML and regenerate a revised track.
  Uses Lyria 3 and MCP analyze via cocrates-google-genai. Do not select for
  TTS/speech as the primary deliverable (speech-generation), video clips
  (video-generation), still images, or diagrams.
metadata:
  agent: cocrates
---

# Music Generation

Produce music that **delivers an approved message** — not decorative noise.

Lyria 3 ([docs](https://ai.google.dev/gemini-api/docs/music-generation)): **44.1 kHz stereo** from text (and optionally images). Targets **Clip / Pro** — not Lyria RealTime streaming.

| Layer | Owns | Must not |
|-------|------|----------|
| **YAML** | `title`, optional `summary`, `message`, `design`, MCP request | Call MCP before YAML approval (except Express) |
| **MCP** `cocrates-google-genai` | `generate` → audio; optional `analyze` → text | Invent meaning absent from YAML; auto-analyze |

Pipeline: **`message`** (story / mood / purpose to convey) → **`design`** (how to shape it musically) → **`params.prompt`** (+ optional `lyrics`) (English realization). Not a post-hoc track critique.

## Core Principles

- **Message-first** — the track must make `message` (story/beats of feeling or purpose) graspable on listen.
- **Design shapes the message** — `design` materializes the musical story; it is not a separate decorative brief.
- **Chat then YAML** — gather what the YAML needs in chat; write when enough is known. If already enough, write YAML immediately.
- **YAML gate before generate** — user reviews YAML (`design` primary), then MCP `generate`.
- **User review by default** — after generate, the user evaluates the audio; AI `analyze` only on explicit request.
- **Consistency** — requirement fields aligned; `design` ↔ `params.prompt` ↔ `lyrics` / `references` aligned.
- **Separate direction from lyrics** — musical direction in `prompt`; custom words in `params.lyrics` (section-tagged).
- **Single-turn Lyria** — no in-model multi-turn edit; revise = new YAML + `generate`. Prefer Clip before Pro when exploring.
- **Intent fidelity** — do not invent unrequested genre, vocals, lyrics, or structure.

Human YAML fields (`title`, `summary`, **`message`**, `design`) use the **user's language** — never default them to English when the user writes in another language. `params.prompt` is **English** unless non-English lyrics require matching prompt language.

## Workflow

```
A Forward
  1 Discover (chat)   requirement + music design
  2 Write YAML        title / [summary] / message / design / MCP block  → review
  3 Approve → Generate   cocrates-google-genai                              → audio
  4 Review            user evaluates (default); optional AI analyze on request
  5 Revise            update YAML; keep fields consistent; regenerate if needed

B From existing media (understand → YAML → regenerate)
  analyze audio and/or its generation YAML (or reuse prior speech/audio analysis)
  → draft message/design (+ lyrics) → Write YAML → review/approve
  → Generate → Review / Revise as in A
```

**Enough already:** Write complete YAML immediately; stop for approval (unless Express).

**Express** (e.g. *generate now*, *express*): YAML + generate; revise from the audio.

Mark only the MCP block with `# --- cocrates-google-genai ---`.

## Paths

Default: `music/{slug}.yaml`, `output: "./{slug}.mp3"` beside it. Optional `lyricsOutput: "./{slug}.lyrics.txt"`. Relative paths resolve against **that YAML's directory**. Refs: e.g. `music/refs/…`.

**`params.references` rule (default YAML):** Mood refs are **images**. When a mood image was produced from a generation YAML, list that **`.yaml`** path — MCP resolves it to the YAML’s `output` image. Use a **media file** only when there is no generation YAML (raw photo/stock). Prefer YAML so generate and analyze share the same contract graph.

---

## Understand & revise from media

When the user asks to interpret / analyze / evaluate existing **music or audio** to make or revise a track ([audio understanding](https://ai.google.dev/gemini-api/docs/audio) — mood, instruments, non-verbal texture, structure, lyrics):

1. **Analyze** — MCP `analyze`:
   - Prefer `inputs: [generation YAML]` when revising a prior Lyria request (MCP loads that YAML’s `output` audio and inlines referenced YAMLs; `prompt` optional).
   - Else `inputs: [audio path|URL]` (or reuse findings from a prior speech-generation analysis) with a prompt for genre/mood/structure/lyrics as needed.
   Present `{ text, interactionId }`.
2. **Material for YAML** — draft `message` + `design` (+ `lyrics` if sung). Confirm keep/change vs source before locking.
3. **Write YAML** — full Lyria request; cite the source in `design` (Lyria has no audio-ref edit — encode the desired sound in prompt/lyrics; optional mood **references** — prefer image **YAML** when available, else media).
4. **Review / approve** → **Generate** → **Review / Revise**.

If the primary deliverable is revised **spoken TTS**, use **speech-generation** instead. Do not skip the YAML gate.
---

## Phase 1 — Discover (chat)

Do not write partial YAML here. Work **requirement before music design**.

### 1.1 Requirement

| Field | Role | Required |
|-------|------|----------|
| `title` | Short track name | Yes |
| `summary` | One-sentence claim | No |
| `message` | Communicative story / beats — what the music should *mean* or *do* for the listener | Yes |

**`message` (story / beats):** Purpose and emotional/narrative takeaway — why this music exists (BGM calm, night-city longing, game-loop energy…). Write it in the **user's language** so someone could retell the intent without naming instruments or BPM.

| `message` owns | Belongs in `design` / prompt instead |
|----------------|--------------------------------------|
| What to convey (mood arc, story, use-case feeling) | Genre, instruments, BPM, key, arrangement |
| Why the listener should care | Structure tags, timelines, Clip vs Pro, format |
| Vocal intent as meaning (sung story vs instrumental space) | Exact lyric lines, section tags, image paths |

**Anti-pattern:** A `message` that reads like a production sheet or duplicates `design`. Example: message = “quiet cafe warmth that doesn’t interrupt conversation” → design = lo-fi Rhodes, 85 BPM, instrumental Clip…

**Consistency:** Present fields stay one story (`title` ↔ [`summary`] ↔ `message`). Realign all present fields on edits.

**Fit test:** Could a listener name the intended story/purpose after hearing a track that matches — without the YAML? Could you write a different `design` that still serves the same `message`?

### 1.2 Music design (maps to `design` + prompt / lyrics)

Agree in chat (then encode in YAML `design`) — **shape the approved message into a track** (realization layer; prompt executes this):

| Topic | Capture |
|-------|---------|
| Model | `lyria-3-clip-preview` (always **30s**) vs `lyria-3-pro-preview` (full song; steer length in prompt) |
| Genre / blend | e.g. lo-fi hip hop, dreamy indie pop, cinematic |
| Instruments | Specific when possible (Rhodes, 808, fingerpicked acoustic…) |
| BPM / tempo | Exact or relative |
| Key / scale | When it matters |
| Vocals | Instrumental only vs sung; lyric language |
| Structure | `[Verse]` / `[Chorus]` / … and/or `[m:ss - m:ss]` (esp. Pro) |
| Lyrics | User-supplied (→ `params.lyrics`) vs model-written theme |
| Images | Up to **10** mood refs via `params.references` — prefer image **YAML**, else media; how they inspire the sound |
| Format | `mp3` (default) or `wav` (mainly Pro) |

**Defaults:** unclear / loop / preview → **Clip**; full song, multi-section custom lyrics, detailed timeline → **Pro**. Recommend Clip first to lock the brief, then upgrade.

**Fit test:** From the design brief alone, could a producer recreate the intended track? Does every design choice serve `message`? Thin briefs (genre + “nice” only) are not enough.

Do not invent copyrighted lyrics or “sound exactly like [famous artist]”.

When coherent → Phase 2.

---

## Phase 2 — Write YAML

```yaml
# Human fields (title / message / design) = user's language; params.prompt = English
title: 카페 BGM
message: |
  대화를 받쳐 주는 조용한 카페의 온기 — 존재하되 방해하지 않고,
  시선을 빼앗는 공연이 아니라 부드러운 쉼.

design: |
  Clip, 인스트루멘탈만. 로파이 힙합: 먼지 낀 바이닐, 부드러운 로즈,
  붐뱁 ~85 BPM, 재즈 업라이트 베이스. 보컬 없음. mp3.

# --- cocrates-google-genai ---
type: music
model: lyria-3-clip-preview
params:
  prompt: |
    A 30-second lofi hip hop beat with dusty vinyl crackle, mellow Rhodes piano
    chords, a slow boom-bap drum pattern at 85 BPM, and a jazzy upright bass
    line. Instrumental only, no vocals.
  outputFormat: mp3
output: "./cafe-bgm.mp3"
```

### `design`

User-facing music brief in the user's language — primary YAML review object. **Designs how to realize `message` musically**; `params.prompt` (and `lyrics`) realize that design. Cover the §1.2 topics; may be shorter than `params.prompt` but must not omit model, genre/instruments/groove, vocals intent, structure, refs, or format. Every musical claim in `params.prompt` must be grounded here. Do not restate the communicative story in place of production decisions — keep story in `message`.

### MCP request

| Field | Role |
|-------|------|
| `type` | `music` |
| `model` | `lyria-3-clip-preview` or `lyria-3-pro-preview` |
| `params.prompt` | English musical direction from `design` (+ timeline if Pro) |
| `params.lyrics` | Optional section-tagged custom lyrics |
| `params.references` | Ordered `{path}` mood image refs (max 10; **prefer YAML**, else image media) |
| `params.outputFormat` | `mp3` (default) or `wav` |
| `params.lyricsOutput` | Optional path for returned lyrics/structure text |
| `output` | Default `./{slug}.mp3` (or `.wav`) |
| `background` | Optional; prefer `true` for long Pro jobs |

**Do not use `params.images`** — removed; MCP returns `INVALID_INPUT` (`use params.references`).

### Prompt & lyrics

- Prompt: genre → instruments → BPM/key → mood → vocals intent → duration/structure (Pro) → image mood link → avoidances (`Instrumental only, no vocals.`).
- Custom lyrics → `params.lyrics` with `[Verse 1]`, `[Chorus]`, … — not unlabeled prose dumped into `prompt`.
- Non-English sung lyrics: write **prompt in that language** so pronunciation matches.
- Clip may say “30-second” but output is **always** 30s.
- Resolve ref paths against this YAML’s directory; for YAML refs, that YAML’s `output` image must exist; for raw media, the file must exist; search + user approve if moved.
- Music-ready test: could a producer recreate from the prompt alone?

### YAML gate

Present **`design`**; note Clip vs Pro, vocals/lyrics, references, format/output. Warn that Pro can take time. **Stop** (unless Express).

---

## Phase 3 — Generate

MCP **`cocrates-google-genai`** (GetMcpTools; `mcp_auth` if needed).

1. YAML on disk (`type: music`).
2. Preflight every `params.references[].path` against **this YAML's directory** (YAML refs → file + its `output` image; media refs → file); on fail, stop and ask — do not `generate`.
3. `generate` with `filePath` → report `files` (audio; plus lyrics file if `lyricsOutput` set). Do not treat empty `files` as success.

Lyria is **single-turn** — do not rely on `continue_interaction` to “fix the chorus”; revise the YAML instead.

Then Phase 4.

---

## Phase 4 — Review (user default; optional AI analyze)

**Default:** Present the generated audio path(s). The **user** reviews and evaluates. Do not call `analyze` unless asked.

**Optional AI analyze** — only on explicit request (e.g. *analyze*, *evaluate*, *AI review*, *평가해줘*). Never auto-run after generate.

1. Prefer analyzing via the **request YAML** (MCP loads `output` audio and recursively includes referenced YAMLs). Verify the YAML and its `output` on disk. Optionally add mood images if the user wants a comparison (1–10 `inputs`; at most one generation YAML).
2. Call MCP **`analyze`**:
   - `inputs`: `[this request YAML]` (preferred) — or `[artifact path]` when there is no YAML / raw-audio Q&A
   - `model`: omit unless the user overrides (default `gemini-3.5-flash`)
   - `prompt`: optional when a generation YAML is in `inputs`. When providing one, use the **user's language**, grounded in approved `message` + `design` (and `summary` if present). Ask for a structured report covering **all** of:
     1. **Intent fit** — does the track match the designed musical intent (`design` / genre / groove / vocals / structure)?
     2. **Message delivery** — is `message` (story / purpose / mood arc) clearly conveyed?
     3. **Functional** — wrong Clip/Pro length expectations, missing sections, lyric mishaps, ref misuse
     4. **Quality / completeness** — mix clarity, arrangement gaps, abrupt endings, polish issues
     5. **Improvements** — concrete, prioritized suggestions (what to change in `design` / prompt / lyrics)
3. Present `{ text, interactionId }` to the user. Do **not** silently edit YAML or regenerate from the report.
4. Follow-ups: `continue_interaction` with the same `interactionId` if the user asks more about the **analysis**. New audio → new `analyze` call. Musical fixes still go through Phase 5 + new `generate`.

---

## Phase 5 — Revise

Keep `design` ↔ prompt ↔ lyrics/references consistent; re-approve; **new** `generate`. If `message` changes, reshape `design` to match.

When improving from user feedback or an AI analyze report: apply only **user-agreed** changes; update owning YAML fields, then regenerate. Prefer Clip A/B before another Pro run when exploring.

---

## Full-song example (Pro + lyrics)

```yaml
# Human fields = user's language; params.prompt / lyrics language = as agreed for performance
title: 메아리
message: |
  밤도시의 그리움이 함께하는 밝음으로 바뀐다 — 네온 속을 나란히 걷고,
  외로움이 아니라 연결의 짧은 앤섬.

design: |
  Pro, 약 2분, 몽환 인디 팝, 미드템포, 부드러운 신스 + 어쿠스틱 기타.
  영어 verse/chorus 가사. mp3 + lyricsOutput.

# --- cocrates-google-genai ---
type: music
model: lyria-3-pro-preview
params:
  prompt: |
    Dreamy indie pop, mid-tempo, soft synths and acoustic guitar.
    Create a 2-minute song with verse / chorus / bridge.
    Use the provided lyrics.
  lyrics: |
    [Verse 1]
    Walking through the neon glow,
    city lights reflect below.

    [Chorus]
    We are the echoes in the night,
    burning brighter than the light.
  outputFormat: mp3
  lyricsOutput: "./echoes.lyrics.txt"
output: "./echoes.mp3"
```

---

## Image-inspired example (Pro)

```yaml
# Human fields = user's language; params.prompt = English
title: 사막 황혼 앰비언트
message: |
  넓고 고요한 황혼 — 열이 색으로 식어 가고, 춤곡이 아니라
  천천히 다가오는 스케일의 정적.

design: |
  Pro, 참조 기반 인스트루멘탈 앰비언트. 무드 참조: ./refs/desert-sunset.yaml
  (YAML 없으면 ./refs/desert-sunset.jpg). 타임라인: 패드 → 드문 퍼커션 →
  피크 → 피아노 페이드. mp3.

# --- cocrates-google-genai ---
type: music
model: lyria-3-pro-preview
params:
  prompt: |
    An atmospheric ambient track inspired by the mood and colors in the
    reference images. Instrumental only, no vocals.

    [0:00 - 0:20] Soft pads and distant piano
    [0:20 - 1:00] Add sparse percussion and rising strings
    [1:00 - 1:30] Peak, then fade to piano alone
  references:
    - path: "./refs/desert-sunset.yaml"   # or "./refs/desert-sunset.jpg" if no YAML
  outputFormat: mp3
output: "./desert-dusk.mp3"
```

---

## Limitations (do not promise around them)

- Safety filters block artist-voice cloning and copyrighted lyrics
- SynthID audio watermark on all outputs (inaudible)
- No multi-turn refine of a prior Lyria generation
- Clip = fixed 30s; Pro length approximate (prompt/timestamps)
- Non-deterministic — same prompt can vary
- Not Lyria RealTime unless the user asks for streaming jam
- Pro can take a long time; client timeouts → report failure

---

## Prohibitions

- MCP before YAML approval (except Express); auto Express; empty `files` as success
- Auto `analyze` without an explicit user request; applying analyze suggestions without user agreement
- `message` that duplicates `design` (instrument lists, BPM sheets, timeline dumps as the “story”)
- Defaulting `message` / `design` / `title` to English when the user writes in another language
- Vague prompts when a specific sound is wanted (“nice music” alone)
- Unlabeled custom lyrics mixed into direction; famous-artist voice / copyrighted lyrics
- Pro-only multi-minute expectations on Clip without warning it stays 30s
- Claiming in-place conversational edit of a generated file via Lyria follow-up
- Guessing ref paths; >10 reference images; inventing requirements that contradict the brief
- Preferring media paths in `params.references` when a generation YAML exists (use the YAML)
- Using `params.images` (removed — use `params.references`)
