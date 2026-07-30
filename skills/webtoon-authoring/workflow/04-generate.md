# Stage ④ — Generate (Reference Images → Page Images → Episode Scroll)

**Prerequisites:** Stage ③ story lock approved (per episode).

**Gate artifacts:**
- reference images and page images under `images/characters`, `images/locations`, `images/stagings`, `images/{nnn}-{episode-slug}/`
- `images/{nnn}-{episode-slug}/episode-scroll.png` (pages stitched top→bottom)
- final assembly confirmation in `output/{webtoon-slug}-final/` (user approval)

**Next stage:** none (this is the terminal stage for the skill)

---

## Hard Boundary

To preserve story lock:
- Allowed in ④:
  - image YAML prompts that implement locked cut/illustration guides
  - re-rendering reference/page PNGs for quality/consistency (max 2 retries per image)
  - rendering locked balloon dialogue and narration/captions inside each page image
  - stitching approved page PNGs into the episode scroll (and optional portal split-export)
- Not allowed in ④:
  - rewriting episode/page/cut story, balloon text, or captions
  - changing illustration guide meaning (only prompt phrasing/tightening)

If a problem is a **Design gap** (wrong cast/place/state/unclear guide/missing balloon):
- Stop generation
- Roll back to Stage ②
- Re-run Stage ③ for affected episode(s)

---

## Canvas & Generation Specs

Read from locked `overview.md` (defaults if unspecified). Cut heights/gutters come from locked episode design + `workflow/cut-composition.md`.

| Spec | Default |
|------|---------|
| Target width | **690–800px** locked in overview (do not invent mid-generate) |
| Page height | **Variable** — sum of locked cut heights + gutters |
| Generation size | 1K unless user asks otherwise |
| Aspect ratio | **9:21** as the default vertical-strip frame (~672×1584). If a page needs more height, either (a) design/generate as multiple stacked strips within the page folder and stitch, or (b) generate taller then crop/extend per user approval — never silently change locked cut order or size classes |
| Color | Full color |
| Side margins | Full bleed or 30–50px even sides (per lock) |
| Outside-cut fill | Per overview / page note |

**Portal note:** Uploads may require splitting tall images (e.g. ~1280px height). Keep creative continuity as one scroll; split only at export if requested.

**Composition fidelity:** When writing page YAML, preserve locked **size class**, approximate **cut height**, **shape** (open/diagonal vs box), and **gutter class/distance** between stacked cuts. A climax `tall` cut must feel scroll-through tall; a `pause` gutter must read as intentional empty vertical space, not a forgotten gap.

---

## Phase 0: Reference Image Generation (Characters + Locations + Stagings)

Goal: Generate durable identity and spatial anchors so later page images stay consistent. Full rules: `workflow/reference-models.md`.

| Layer | Include in reference images | Exclude (cut direction) |
|-------|----------------------------|-------------------------|
| Character | Face/body; lasting body change; outfit + **identity gear** (weapons, shields, accessories) per state | Expression, mood, transient pose |
| Location | **Set/stage** structure per position/view/state | Time/season/weather; one-off camera |
| Staging | Continuing-situation **blocking** (café L/R, OR stations, meeting seats, …) — **2–3** views of the same map | Expression; reseating/L-R flip without Design change |

### 0.1 Generation order (per catalogs)

1. Character base: `images/characters/{character-slug}.yaml/.png`
2. Character state variants: `images/characters/{character-slug}-{state-slug}.yaml/.png`
3. Location base: `images/locations/{location-slug}.yaml/.png`
4. Location position/view: `images/locations/{location-slug}-{position-slug}-{view-slug}.yaml/.png`
5. Location state variants: `images/locations/{location-slug}-{position-slug}-{view-slug}-{state-slug}.yaml/.png`
6. **Staging ensemble refs** (after cast + location PNGs exist):  
   `images/stagings/{staging-slug}-{establishing|reverse|detail}.yaml/.png`  
   — typically **2–3** images; each must enforce the staging blocking map using character + location reference images

### 0.2.1 Sample: reference image YAML structure

**Character base** (`images/characters/{character-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {webtoon-ready character description in English}
    A full-color webtoon / vertical comic character reference sheet...
    Include identity gear exactly (clothes, accessories, named weapons/shields). Neutral expression; no dramatic acting pose.
  size: 1K
  aspectRatio: "1:1"
  seed: null

output: "./images/characters/{character-slug}.png"
```

**Location base** (`images/locations/{location-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {webtoon-ready location description in English}
    A full-color webtoon set/stage background reference (empty of story action)...
  size: 1K
  aspectRatio: "9:16"
  seed: null

output: "./images/locations/{location-slug}.png"
```

**Staging establishing** (`images/stagings/{staging-slug}-establishing.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    Using the provided character and location reference images, compose a STAGING REFERENCE
    (continuing-situation blocking lock) for a webtoon scene (café / OR / meeting / formation / etc.).
    Lock who sits/stands where exactly as designed; preserve left/right and stations; neutral expressions; no balloons.
    Wide establishing view of the full group placement.
  images:
    - path: "./images/locations/...."
    - path: "./images/characters/...."
  size: 1K
  aspectRatio: "9:16"
  seed: null

output: "./images/stagings/{staging-slug}-establishing.png"
```

### 0.2 Per-image YAML approval (MANDATORY)

For each reference image:
1. Create image-generation YAML that reflects the approved Design notes.
   - `message`: identity/structure/blocking consistency only (no story/balloon detail dump)
   - `design`: visual implementation only (character identity gear, set structure, or staging seats; no on-image text)
2. Show the full YAML to the user and request explicit approval.
3. Only after explicit approval, call the MCP image generation.
4. Verify visual quality (do not treat it as story redesign).
5. For stagings: verify seating matches the blocking table across all 2–3 views.
---

## Phase 1: Page Image Generation (Cuts + Balloons + Captions)

For each episode:
- For each page index (`{00}`, `{01}`, ...), generate:
  - `images/{nnn}-{episode-slug}/{page-idx}.yaml`
  - `images/{nnn}-{episode-slug}/{page-idx}.png`

Each page image must depict **all locked cuts on that page** stacked top→bottom, with gutters, speech balloons, and captions as designed — not a single undivided picture-book illustration unless the locked design intentionally uses one full-bleed cut.

### 1.1 Per-page YAML approval (MANDATORY)

For each page image:
1. Create `{page}.yaml` that:
   - uses character/location/**staging** reference PNGs as locked
   - implements only the locked cut illustration guides (order preserved top→bottom)
   - when a cut cites a staging: keep **seating/formation** from that staging; vary expression/gesture/camera tightness only
   - does not silently swap identity gear or reseat characters
   - renders locked balloon `text` and caption `text` **verbatim** (target language; no rewrite)
   - places balloons/captions per locked placement (do not cover faces/key action)
   - uses outside-cut fill from overview/page notes
   - prefers `aspectRatio: "9:21"` unless the locked page height plan requires a different approved approach
   - YAML role boundary:
     - `title`: `{episode-slug} / Page {idx}`
     - `message`: page story meaning/emotion/beat only (1–2 sentences)
     - `design`: cut stack + balloon/caption map + gutters + color/mood + staging cites (from locked episode design)
2. Show YAML to user for explicit review and approval
3. On explicit approval only → call MCP generate
4. Verify:
   - cut order and gutters match lock
   - balloon/caption accuracy and readability
   - width/vertical-strip feel suitable for locked target width (note resize at export if generated at 1K)

### 1.2 Prompt pattern (guideline)

```
Using the provided reference images:
- Keep character appearance (including identity gear), set structure, and series art style consistent
- When staging refs are provided, keep WHO SITS/STANDS WHERE locked; do not reseat
- Compose a VERTICAL SCROLL WEBTOON PAGE (not a picture-book single illustration):
  stacked comic panels (cuts) from top to bottom with clear gutters between cuts
- Cut order and content must follow the locked design exactly
- Outside panel area: {white / black / theme color}

CUTS (top → bottom):
1. Cut {n}: size class {standard|tall|open|…}, height ~{px}, shape {box|open|diagonal}
   Art: {action / framing from locked illustration guide}
   Balloons:
   - {type} from {speaker}: "{verbatim text}" at {placement}
   Captions:
   - "{verbatim text}" at {placement}
   Gutter after: class {tight|normal|wide|pause}, ~{px} empty vertical space
2. ...

TEXT RULES:
- Balloon and caption strings must be exact; do not translate or paraphrase
- Reading order: top → bottom (then left → right within a cut if needed)
- Do not obscure faces or key action with balloons
- Full-color webtoon style suitable for smartphone vertical reading
```

**Must follow:**
- All locked balloon/caption lines appear in the image
- No text mutation of locked strings
- Cuts remain visually separated by gutters (unless design says full-bleed merge)
- Final look = finished webtoon strip segment, not a captioned picture-book page

---

## Phase 2: Stitch Episode Scroll

1. After all pages for the episode are approved, concatenate page PNGs **top → bottom** into:
   - `images/{nnn}-{episode-slug}/episode-scroll.png`
2. Optionally resize the scroll (or each page) to **target width** (690px or overview width) while preserving aspect.
3. If the user/portal needs split upload files, export additional slices (e.g. max height ~1280px) **without changing art** — document slice boundaries next to the files if useful.
4. Show the stitched scroll (or a clear multi-page preview path) to the user for review before G4.

If stitch reveals a **design** problem (wrong cut order, missing beat): rollback to Stage ② — do not “fix” in stitch.

---

## Phase 3: Visual Consistency Review

1. Compare generated page images and the episode scroll within each episode.
2. If **visual inconsistencies** (palette/style/identity) appear:
   - regenerate only the affected image(s)
   - do not change locked story/text
   - max 2 retries per image
   - re-stitch if pages changed
3. If inconsistencies indicate a **Design gap**:
   - rollback to Stage ②
   - re-evaluate and regenerate affected assets

---

## Approval Gate G4 — Final Result

Confirm with user:
1. All reference images — quality acceptable?
2. All page images — cuts, balloons, captions faithful to lock?
3. Episode scroll — continuous vertical read feels right?
4. Final result in `output/{webtoon-slug}-final/` — ready to deliver?

**Do not deliver until G4 is approved.**
