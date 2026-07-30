# Stage ④ — Generate Components (Images → Video → Speech/Music)

**Prerequisites:** G3 design lock for the active segment.

**Gate:** G4 — **per segment** component set approved.

**Next stage:** `05-assemble.md` for that segment only after G4.

---

## Hard Boundary

**Allowed:** media YAML + assets that implement locked clip messages/direction.  
**Not allowed:** rewriting clip messages, inventing new story beats, changing reference identity meaning.

Design gap → stop → Stage ② → re-G3 → resume ④.  
Quality-only → tighten YAML / retry (max 2 per asset) without changing lock meaning.

---

## Unit of work

Generate components **per segment** (Short = the implicit segment in `sequence.md`).  
Shared reference images (characters/locations/…) may be generated once and reused across segments.

---

## Phase 0 — Reference images (if required)

**Skip** when `references.md` declares no catalogs.

Order (typical):

1. Entity base references under `images/references/{kind}/`
2. State / variant references as declared in catalogs

For each asset:

1. Write `image-generation` YAML (`message` = identity meaning; `design` = visual identity only).
2. User approves YAML.
3. MCP generate → PNG.
4. User accepts quality (or retry ≤2).

Follow `image-generation` skill rules (English `params.prompt`, etc.).

---

## Phase 1 — Clip key images

For each clip in the segment that needs an image (still, motion source, or hold source):

1. YAML under `images/clips/{segment}/{clip-slug}.yaml` (or flat naming `{nnn}-{clip}.yaml`).
2. Prefer referencing Phase 0 PNGs when catalogs exist.
3. `message` ← locked clip message (summary, no shot-list dump; may stay in the user's language).
4. `design` ← locked direction guide (visual implementation only; may stay in the user's language).
5. Approve → generate → accept.

---

## Phase 2 — Motion video (if needed)

When clip track plan includes video:

1. Use clip key image as reference / first frame per `video-generation` skill.
2. Keep duration **under ~10s**.
3. YAML in `videos/{segment}/{clip-slug}.yaml` → approve → generate → mp4.
4. Do not invent dialogue/casts beyond locked design.

---

## Phase 3 — Voice audio (if needed)

When clip needs narration/dialogue:

1. Confirm script from locked design (draft in chat only if missing, then lock into Design if story-affecting).
2. Use `speech-generation` YAML under `speech/{segment}/`.
3. Approve → generate → audio file.

---

## Phase 4 — BGM / SFX (if needed)

When clip/segment needs BGM or SFX:

1. Use `music-generation` under `music/{segment}/` (or shared `music/` for bed loops).
2. Approve → generate.
3. Never use music-generation for spoken words; never use speech-generation for BGM/SFX beds.

---

## Per-YAML approval (MANDATORY)

No batch auto-generation. Each YAML: show → explicit approve → generate.

---

## Segment component checklist (G4)

For segment `{nnn}`:

- [ ] Required reference images exist & approved (or N/A)
- [ ] Each clip’s required key image approved
- [ ] Each required video approved (&lt;~10s)
- [ ] Each required speech asset approved
- [ ] Each required music/SFX asset approved
- [ ] Asset paths documented for Assemble (`src` candidates)

---

## Gate G4 — Component set (per segment)

User confirms the segment’s component set is sufficient and approved to assemble.

**Do not enter Stage ⑤ for this segment without G4.**

Other segments may still be in ③/④.
