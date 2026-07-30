# Edit Spec (for MCP `cocrates-video-edit`)

Self-contained contract for Stage ⑤ Assemble. Write YAML/JSON matching this schema, then call MCP tools — no external repo required.

| MCP tool | Purpose |
|----------|---------|
| `validate_spec` | `{ path }` and/or `{ spec }` — validate; file `path` resolves duration defaults |
| `render_video` | `{ path }` and/or `{ spec }`, optional `{ output }`, `{ baseDir }` — Remotion render; returns output media path |

Discover the live server/tool schema before calling.

---

## Skeleton

```yaml
output: string              # optional. Default = ./{basename}.mp4 next to the spec file
width: number               # optional. Default 1920
height: number              # optional. Default 1080
fps: number                 # optional. Default 60
defaults:                   # optional
  duration: number          # seconds; fallback when duration cannot be inferred (recommend 4)
  transition: Transition | null
  track: object             # optional defaults merged into every track
  outputVolume: number      # master gain; default 1
clips: Clip[]               # required; ≥1; played in array order
audioTracks: AudioTrack[]   # optional; timeline-absolute audio beds
```

- Times in public fields are **seconds** (not frames). `fps` is for output/sampling only.
- `clips[]` are concatenated in order.
- Inside a clip, `tracks[0]` is the **top** visual layer; later tracks are behind.

---

## Clip

```yaml
duration: number            # optional seconds; if omitted, infer (see Duration)
tracks: Track[]             # required; ≥1
transition: Transition      # optional; at the **end** of this clip into the next
```

### Transition

```yaml
type: fade                  # MVP: fade only
duration: number            # seconds; default 0.5 if omitted when transition present
```

- Omit transition or `transition: null` → **hard cut**.
- Transition does **not** shorten clip `duration`; it **overlaps** the end of A with the start of B.
- Timeline length ≈ `sum(clip.duration) - sum(adjacent transition.duration)`.

---

## Track (common)

```yaml
type: video | audio | image
start: number               # optional; clip-local start (sec); default 0
end: number                 # optional; clip-local end (sec); default = clip duration
```

Valid after duration is known: `0 ≤ start < end ≤ duration`.

### `type: video`

| Field | Required | Notes |
|-------|----------|--------|
| `src` | yes | Video file path |
| `startFrom` | | Source in-point (sec). Default `0`. `-1` = last frame |
| `endAt` | | Source out-point (sec). Default = source end. `-1` = last frame |
| `volume` | | Source audio gain. Default `1`. **Ignored (muted) on hold** |
| `playbackRate` | | Default `1`. Ignored on hold. Pitch follows rate (no pitch-preserve) |
| `loop` | | Default `false`. Repeat source range to fill track window. **Ignored on hold** |

**Hold:** after normalize, `startFrom === endAt` → freeze that frame for the track window; source audio muted.  
Examples: first-frame hold → `endAt: 0` (startFrom defaults to 0); last-frame hold → `startFrom: -1`.

**Play:** after normalize, `startFrom < endAt`.

### `type: audio`

| Field | Required | Notes |
|-------|----------|--------|
| `src` | yes | Audio file path |
| `startFrom` / `endAt` | | Same sentinel rules as video (`-1` = end) |
| `volume` | | Default `1` (linear) |
| `playbackRate` | | Default `1`; pitch follows rate |
| `loop` | | Default `false` |

Audio with `startFrom === endAt` after normalize is an **error** (hold is video-only).

### `type: image`

| Field | Required | Notes |
|-------|----------|--------|
| `src` | yes | Image path |

Still for the track window. No intrinsic duration.

---

## Global `audioTracks`

Timeline-absolute beds across clip boundaries:

| Field | Required | Notes |
|-------|----------|--------|
| `src` | yes | Path |
| `start` | | Timeline start (sec). Default `0` |
| `end` | | Timeline end (sec). Default = timeline end |
| `startFrom` / `endAt` | | Source range; `-1` sentinel |
| `volume` / `playbackRate` / `loop` | | Same semantics as clip audio |

---

## Duration inference

When clip `duration` is **omitted**:

1. If any **playing** video exists (`startFrom !== endAt` after normalize): clip duration = **max** of those video play lengths `(endAt - startFrom) / playbackRate`. Longer audio is **trimmed**.
2. Else if any **audio** exists: duration = max audio lengths (hold video + narration pattern).
3. Else (image-only / hold-only with no audio): use **`defaults.duration`**, or **error**.

When `duration` is **set**, use it; do not re-infer.

### Media vs window `W`

| Case | Behavior |
|------|----------|
| Source one-shot `L` > `W` | Trim to `W` (no auto slow-mo) |
| `L` < `W`, `loop: false`, video | Play then **hold last frame** (source audio muted for hold tail) |
| `L` < `W`, `loop: false`, audio | Play then **silence** |
| `L` < `W`, `loop: true` | Repeat source range to fill `W` |
| Hold video | Freeze for whole `W`; `loop` ignored |
| Image | Hold for whole `W` |

Inference length is always one-shot `L` even if `loop: true`.

---

## Audio mix

At time `t`, linear sum of active sources × `defaults.outputVolume` (default 1).  
No auto-ducking / loudness normalization. `volume` is linear (0 = mute, 1 = unity).

---

## Validation failures (must fix before render)

- Empty `clips` or empty `tracks`
- `type` not in `video` \| `audio` \| `image`
- Missing/`src` unreadable
- After normalize, video/audio `startFrom > endAt`
- Audio with `startFrom === endAt`
- `playbackRate ≤ 0`
- Transition `type` other than `fade`
- Invalid `start`/`end` vs duration
- Non-positive `width` / `height` / `fps` when set

---

## Examples

### Cut + concat

```yaml
output: ./out/cut-concat.mp4
width: 1280
height: 720
fps: 30
clips:
  - tracks:
      - type: video
        src: ./a.mp4
        startFrom: 0
        endAt: 2
  - tracks:
      - type: video
        src: ./b.mp4
        startFrom: 0
        endAt: 2
    transition:
      type: fade
      duration: 0.5
```

### Hold + play + end hold + image

```yaml
output: ./out/hold-play.mp4
fps: 30
clips:
  - tracks:
      - type: video
        src: ./talk.mp4
        endAt: 0
      - type: audio
        src: ./n1.mp3
        endAt: 3
  - tracks:
      - type: video
        src: ./talk.mp4
        startFrom: 0
        endAt: 5
        volume: 0
      - type: audio
        src: ./n2.mp3
        startFrom: 0
        endAt: 8
    transition:
      type: fade
      duration: 0.5
  - tracks:
      - type: video
        src: ./talk.mp4
        startFrom: -1
      - type: audio
        src: ./n3.mp3
        endAt: 2
  - tracks:
      - type: image
        src: ./endcard.png
      - type: audio
        src: ./n4.mp3
        endAt: 4
```

### Partial track window + loop bed

```yaml
clips:
  - duration: 6
    tracks:
      - type: video
        src: ./a.mp4
        startFrom: 0
        endAt: 6
      - type: audio
        src: ./vo.mp3
        start: 1
        end: 5
        startFrom: 0
  - duration: 10
    tracks:
      - type: video
        src: ./bg-loop.mp4
        startFrom: 0
        endAt: 2
        loop: true
      - type: audio
        src: ./bed.mp3
        endAt: 4
        loop: true
        volume: 0.3
audioTracks:
  - src: ./bed-global.mp3
    start: 0
    loop: true
```

### Final stitch (segment mp4s)

```yaml
output: ./output/{slug}.mp4
fps: 30
clips:
  - tracks:
      - type: video
        src: ./output/segments/001-intro.mp4
  - tracks:
      - type: video
        src: ./output/segments/002-body.mp4
    transition:
      type: fade
      duration: 0.5
```

---

## Agent checklist

1. Write edit-spec under `assembly/` using only this schema.
2. User approves the YAML.
3. `validate_spec` with file `path` → fix until valid.
4. `render_video` with same `path` → confirm output file exists.
5. Do not use ad-hoc ffmpeg or invent fields outside this contract.
