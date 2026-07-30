# Stage ⑤ — Assemble (Edit Spec → MCP Render → Segment / Final mp4)

**Prerequisites:** G4 component set approved for the segment being assembled.

**Gate:** G5 — per segment mp4; Long also final mp4.

**Next stage:** none (terminal). Long projects finish with `assembly/final.yaml`.

**MCP:** `cocrates-video-edit`

| Tool | Role |
|------|------|
| `validate_spec` | Validate edit-spec YAML/JSON (resolves duration defaults when given a file `path`) |
| `render_video` | Remotion render → output media path |

**Schema:** `workflow/edit-spec.md` in this skill (self-contained). Do not invent a parallel assemble format.

---

## Hard Boundary

**Allowed:** edit-spec YAML using **G4-approved** `src` paths; fields defined in `edit-spec.md`; MCP validate + render.  
**Not allowed:** generating new story media; changing locked messages; pointing `src` at unapproved files; ad-hoc ffmpeg/hand-stitch instead of this MCP; calling `render_video` before user approval of the edit-spec file.

Missing/bad asset → Stage ④. Message/timing intent wrong at story level → Stage ② + G3.

---

## Procedure

### 5.1 Write segment edit-spec YAML

- Long: `assembly/segments/{nnn}-{segment-slug}.yaml`
- Short: same path with `001-main` (or `assembly/{slug}.yaml` — pick one convention and keep it)

Map each designed clip (order preserved) to a `clips[]` entry with `tracks` pointing at G4 assets. Follow `edit-spec.md` for hold/trim/transition/duration rules. Resolve `src` relative to the YAML file’s directory (or paths the MCP can resolve).

Minimal pattern:

```yaml
output: ./output/segments/{nnn}-{segment-slug}.mp4
fps: 30
clips:
  - tracks:
      - type: video
        src: ./videos/.../clip.mp4
        startFrom: 0
        endAt: 5
      - type: audio
        src: ./speech/.../n1.mp3
        endAt: 5
```

More patterns (hold + narration, fades, image endcards, final stitch): see `edit-spec.md`.

### 5.2 Approve → validate → render (MCP)

1. Show the full edit-spec YAML to the user; get **explicit approval**.
2. Discover tools via MCP schema inspection for `cocrates-video-edit`.
3. Call **`validate_spec`** with `{ "path": "<absolute-or-workspace path to yaml>" }`.
   - On failure: fix YAML or assets using `edit-spec.md`; do not render.
4. On validation success, call **`render_video`** with the same `path` (optional `output` override if needed).
5. Confirm the returned media path exists (typically the spec’s `output`). Short may set `output` to `output/{slug}.mp4` directly.

**Never** call `render_video` before YAML approval and successful validate (unless the user explicitly requests Express-style assemble — still show the YAML).

### 5.3 Gate G5a — Segment approval

User reviews the rendered segment mp4:

- Timing/holds/transitions match locked rhythm?
- Audio/visual sync acceptable?
- Ready for final (Long) or deliver (Short)?

**Short:** G5a = final delivery gate (G5). Stop.

### 5.4 Long — Final stitch

After required segments are G5a-approved:

1. Write `assembly/final.yaml` as an edit-spec whose `clips[]` use `type: video` `src` pointing at approved **segment** mp4s (see final-stitch example in `edit-spec.md`).
2. User approves → `validate_spec` → `render_video` → `output/{slug}.mp4`.
3. **G5b:** user approves final.

---

## Segment-first workflow

```text
for each segment (user-driven order):
  Evaluate(G3) → Components(G4) → edit-spec + validate + render (G5a)

when all required segments G5a:
  assembly/final.yaml → validate + render → final mp4 (G5b)   # Long only
```

User may assemble segment A before segment B’s components exist.

---

## Completeness Check

- [ ] Every `src` is G4-approved (segment specs) or G5a-approved segment mp4 (final)
- [ ] Clip order matches locked Markdown design
- [ ] Edit-spec matches `edit-spec.md` and validates via `validate_spec`
- [ ] Render uses `render_video` only (no substitute stitcher)
- [ ] Short: final mp4 approved
- [ ] Long: each required segment mp4 + final approved

---

## Gate G5

- **Per segment:** approve segment mp4 before treating it as final input.
- **Long final:** approve `output/{slug}.mp4` before delivery.

**Do not deliver until G5 is approved.**
