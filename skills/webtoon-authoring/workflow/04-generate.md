# Stage ④ — Generate (Reference Images → Page Images → Episode Scroll)

**Prerequisites:** Stage ③ story lock approved (per episode).

**Gate artifacts:**
- reference images and page images under `images/characters`, `images/locations`, `images/stagings`, `images/{nnn}-{episode-slug}/`
- `images/{nnn}-{episode-slug}/{nnn}-{episode-slug}.png` (pages stitched top→bottom)
- final assembly confirmation in `output/{webtoon-slug}-final/` (user approval)

**Next stage:** none (this is the terminal stage for the skill)

---

## Hard Boundary

To preserve story lock:
- Allowed in ④:
  - image YAML prompts that implement locked cut/illustration guides **and** locked `illustration-guide.md` (art / chrome / balloon·caption typography)
  - re-rendering reference/page PNGs for quality/consistency (max 2 retries per image)
  - rendering locked balloon dialogue and narration/captions inside each page image
  - stitching approved page PNGs into the episode scroll **only on explicit user request** after page images are complete (and optional portal split-export)
- Not allowed in ④:
  - rewriting episode/page/cut story, balloon text, or captions
  - changing illustration guide meaning (only prompt phrasing/tightening)
  - inventing new balloon fonts, caption boxes, panel chrome, or art-medium treatments not in `illustration-guide.md`

If a problem is a **Design gap** (wrong cast/place/state/unclear guide/missing balloon / missing or weak `illustration-guide.md`):
- Stop generation
- Roll back to Stage ②
- Re-run Stage ③ for affected episode(s)

---

## Canvas & Generation Specs

Read from locked `overview.md` (defaults if unspecified). Cut heights/gutters come from locked episode design + `workflow/cut-composition.md`.

| Spec | Default |
|------|---------|
| Target width | **768px** default (lock exact px in overview; do not invent mid-generate) |
| Page height | **Variable** — sum of locked cut heights + gutters (page packing follows dwell/breath — `cut-composition.md` §4b) |
| **Default page generation** | **`gemini-3.1-flash-image` + `1K` + `9:16`**, split into **strip tiles** (typically **2** per design page), then stitch to `{page}.png` |
| Optional tall single frame | **`gemini-3-pro-image` + `2K` + `1:8`** only — do **not** use Flash at 2K/1:8 for balloon text (frequent text errors) |
| Color | Full color |
| Side margins | Full bleed or 30–50px even sides (per lock) |
| Outside-cut fill | Per overview / page note |

### Default: 9:16 / 1K strip tiles (Flash)

`gemini-3.1-flash-image` renders balloon text more reliably at **1K**. Prefer this path for all normal page work.

1. Partition the locked page’s cuts top→bottom into **tile A / tile B** (default **2** tiles). Prefer splits **on gutters between cuts**, not through the middle of a cut. A single tall cut may occupy one full tile (or more tiles if needed).
2. Generate each tile:
   - `images/{nnn}-{episode-slug}/{page}-a.yaml` → `{page}-a.png`
   - `images/{nnn}-{episode-slug}/{page}-b.yaml` → `{page}-b.png`
   - (If a page needs >2 tiles: `-c`, `-d`, … — still `1K` + `9:16` + Flash unless Pro path is locked.)
3. Stitch tiles top→bottom into `images/{nnn}-{episode-slug}/{page}.png` (page-level tile stitch only — do **not** auto-run episode-scroll stitch).
4. Each tile YAML gets its own approval + generate call (image-generation skill rules).

**Tile continuity:** Tile B’s `design` must state it continues the same page immediately below tile A (same width, style, cast/location refs, outside-cut fill). Do not re-introduce page chrome or restart cut numbering as on-image labels.

### Optional: 1:8 / 2K single frame (Pro only)

Use **only** when overview/user explicitly wants one tall generation:

| Field | Value |
|-------|--------|
| `model` | **`gemini-3-pro-image`** (required) |
| `size` | **`2K`** |
| `aspectRatio` | **`1:8`** |

**Forbidden:** `gemini-3.1-flash-image` with `2K` + `1:8` for pages with balloons/captions.

**Portal note:** Uploads may require splitting tall images (e.g. ~1280px height). Keep creative continuity as one scroll; split only at export if requested.

**Composition fidelity:** When writing tile/page YAML, preserve locked **size class**, approximate **cut height**, **shape** (open/diagonal vs box), and **gutter class/distance** between stacked cuts. A climax `tall` cut must feel scroll-through tall; a `pause` gutter must read as intentional empty vertical space, not a forgotten gap.

---

## Phase 0: Reference Image Generation (Characters + Locations + Stagings)

Goal: Generate durable identity and spatial anchors so later page images stay consistent. Full rules: `workflow/reference-models.md`.

**Required reading before any YAML:** `{project-root}/illustration-guide.md` (§1 Art style). Apply the same medium / palette / line / proportion tokens to every reference image. Full rules: `workflow/illustration-guide.md`.

| Layer | Include in reference images | Exclude (cut direction) |
|-------|----------------------------|-------------------------|
| Character | Static look + lasting body/equipment identity — **turnaround** sheet (front/side/back + key accessories) per state; YAML from character md | Expression, mood, transient pose |
| Location | Static/physical set — **multi-view** Scene List (front/left/right or fixture-sides); YAML from location md | Time/weather; ambient extras (→ staging) |
| Staging | Episode-first WHO–WHERE (+ props/ambient/situation env) — **default 1× establishing**; CU/OTS via Direction | Expression; reseat without Design; OTS/CU as staging refs |

### 0.1 Generation order (per catalogs)

1. Character base: `images/characters/{character-slug}.yaml/.png`
2. Character state variants: `images/characters/{character-slug}-{state-slug}.yaml/.png`
3. Location base: `images/locations/{location-slug}.yaml/.png`
4. Location position/view: `images/locations/{location-slug}-{position-slug}-{view-slug}.yaml/.png`
5. Location state variants: `images/locations/{location-slug}-{position-slug}-{view-slug}-{state-slug}.yaml/.png`
6. **Staging establishing** (after cast + location PNGs exist):  
   `images/stagings/{staging-slug}-establishing.yaml/.png`  
   — **required default (one image)**; optional reverse/detail only if staging md lists them with lock purpose. Must inject Character + Location PNGs in `params.references`. Lean prompt: WHO–WHERE–WHAT only (`workflow/reference-models.md` §4.2–4.3).

### 0.2.1 Sample: reference image YAML structure

**Character base** (`images/characters/{character-slug}.yaml`) — concretize `characters/{slug}.md` base state:

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {webtoon-ready character description in English — concretized from character md}
    A full-color webtoon / vertical comic CHARACTER TURNAROUND reference sheet on one image:
    front view, side view, and back view of the SAME outfit/look, plus key accessories/identity gear
    (clothes, armor, weapons, shields, glasses, etc.) clearly readable. Neutral expression; no dramatic acting pose.
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

**Staging establishing** (`images/stagings/{staging-slug}-establishing.yaml`) — lean WHO–WHERE–WHAT; refs required:

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    Using ONLY the provided character and location reference images, compose ONE canonical STAGING
    establishing lock (webtoon). WHO left/center/right, WHERE on the set, WHAT pose/facing.
    Include situation props/ambient/environment from the staging md. Do NOT re-describe face/hair/outfit/
    wall logos already in the refs. Neutral expressions; no balloons; no metadata text.
  references:
    - path: "./images/locations/...."
    - path: "./images/characters/...."
  size: 1K
  aspectRatio: "9:16"
  seed: null

output: "./images/stagings/{staging-slug}-establishing.png"
```

### 0.2 Per-image YAML approval (MANDATORY — never auto-generate)

**Hard gate:** Do **not** call image generate until the user has seen the full YAML and given an **explicit approval command** (e.g. “생성해줘”, “만들어줘”, “approve”). Never chain autonomous generate calls across assets.

For each reference image:
1. Create image-generation YAML that reflects the approved Design notes.
   - `message`: identity/structure/blocking consistency only
   - `design`: visual implementation only; **design ↔ prompt fidelity** — every proper noun / engraved logo / gear detail in `design` must appear in `params.prompt` (no genericizing)
   - Staging: lean WHO–WHERE–WHAT; Character+Location paths in `params.references` (required)
2. Show the full YAML to the user and request explicit approval.
3. Only after explicit approval, call the MCP image generation.
4. Verify visual quality (do not treat it as story redesign).
5. For stagings: verify L/C/R + props match the staging md on the establishing PNG.
---

## Phase 1: Page Image Generation (Cuts + Balloons + Captions)

**Required before every tile/page YAML:** open `{project-root}/illustration-guide.md` and apply art style (§1), vertical chrome (§2), and balloon/caption recipes (§3–4). Do not invent lettering. Procedure: `workflow/illustration-guide.md`.

For each episode:
- For each page index (`{00}`, `{01}`, ...):
  - **Default:** generate strip tiles `{page}-a`, `{page}-b` (…), stitch → `{page}.png`
  - **Pro optional path:** one `{page}.yaml` / `{page}.png` at `2K` + `1:8` with `gemini-3-pro-image` only

Each page (after tile stitch) must depict **all locked cuts on that page** stacked top→bottom, with gutters, speech balloons, and captions as designed — not a single undivided picture-book illustration unless the locked design intentionally uses one full-bleed cut.

### 1.1 Per-tile / per-page YAML approval (MANDATORY)

For each **strip tile** (default path) or single page frame (Pro path):
1. Create YAML that:
   - **cites and applies `{project-root}/illustration-guide.md`** (mandatory — include §6 checklist in `design`)
   - uses character/location/**staging** reference PNGs via `params.references` (image-generation skill; do not use `params.images`) as locked
   - implements only the locked cuts assigned to **this tile** (order preserved top→bottom within the page; tile B continues after tile A)
   - when a cut cites a staging: keep **seating/formation** from that staging; vary expression/gesture/camera tightness only
   - does not silently swap identity gear or reseat characters
   - renders locked balloon `text` and caption `text` **verbatim** (target language; no rewrite)
   - places balloons/captions per locked placement (do not cover faces/key action)
   - uses **balloon/caption chrome and lettering only from the series guide** (type from cut table → recipe from guide)
   - **when Staging is cited:** attach staging ref PNG(s) under `params.references`; keep **same seats/L-R and same situation props/table dressing** as the staging (Location alone ≠ tonight’s meal / this scene’s props)
   - uses outside-cut fill from overview/guide/page notes
   - **Default:** `model: gemini-3.1-flash-image`, `size: 1K`, `aspectRatio: "9:16"`
   - **Pro tall frame only:** `model: gemini-3-pro-image`, `size: 2K`, `aspectRatio: "1:8"`
   - embeds the mandatory **IMPORTANT** block below verbatim in `design` (or the generation prompt body) so the model cannot drop it
   - YAML role boundary:
     - `title`: `{episode-slug} / Page {idx} / tile {a|b|…}` (or `/ Page {idx}` for Pro single frame)
     - `message`: page/tile story meaning/emotion/beat only (1–2 sentences)
     - `design`: **IMPORTANT** block (required) + **illustration-guide tokens** + cuts in this tile + balloon/caption map + gutters + color/mood + staging cites + tile continuity note
2. Show YAML to user for explicit review and approval — **do not generate until explicit approval command**
3. On explicit approval only → call MCP generate (never auto-chain)
4. Verify:
   - cut order and gutters match lock (across tiles after stitch)
   - balloon/caption accuracy and readability (exact strings; no particle changes)
   - **typography/chrome matches `illustration-guide.md`** (not a one-off lettering style)
   - **no** cut labels / “CUT 1” / panel index text on the image
   - width/vertical-strip feel suitable for locked target width (resize/crop at export to overview width, default **768px**, as needed)
5. After all tiles for a page are approved: stitch `{page}-a.png` + `{page}-b.png` (+ …) → `{page}.png`

### 1.2 Prompt pattern (mandatory for every tile/page YAML)

Every tile/page `design` (prompt body) **must** include this **IMPORTANT** block **verbatim** (English), then the cuts **for this tile only**:

```
Using the provided reference images:
- Keep character appearance (including identity gear), set structure, and series art style consistent
- Series art style (from illustration-guide.md): {medium}, {palette}, {line}, {proportions}
- Balloon/caption recipes (from illustration-guide.md): speech={…}; thought={…}; shout={…}; caption={…}
- Outside-cut / panel chrome (from illustration-guide.md): {fill}, {border rules}
- When staging refs are provided, keep WHO SITS/STANDS WHERE locked and keep SITUATION PROPS / TABLE DRESSING the same (same meal/toys/cups — no wholesale swap); do not reseat
- Compose a VERTICAL SCROLL WEBTOON STRIP TILE (not a picture-book single illustration):
  stacked comic panels (cuts) from top to bottom with clear gutters between cuts
- This image is tile {a|b|…} of page {idx}; it continues the same page strip {above|below} the sibling tile — same width, style, and outside-cut fill
- Outside panel area: {white / black / theme color}

IMPORTANT:
- All balloon text must be rendered EXACTLY as written below, character by character. Do not add, omit, or alter any particle.
- Cut order from top to bottom must be preserved
- Clear gutters between cuts
- NO METADATA TEXT ON IMAGE. Do not draw episode titles, page numbers, cut numbers, panel labels, header/footer text, or watermarks (NO "Episode", NO "Page", NO "Cut", NO "Top Panel", NO "CUT 1"). Render ONLY the specified speech balloons and story captions.
- Suitable for smartphone vertical reading
- Do NOT invent new balloon fonts, caption boxes, or panel chrome not listed in the series illustration guide

CUTS (top → bottom) — this tile only:
1. Cut {n}: size class {standard|tall|open|…}, height ~{px}, shape {box|open|diagonal}
   Art: {action / framing from locked cut Direction}
   Balloons:
   - {type} from {speaker}: "{verbatim text}" at {placement} — apply locked recipe for {type}
   Captions:
   - "{verbatim text}" at {placement} — apply locked caption recipe   <!-- omit section if Captions: none; never invent caption text -->
   Gutter after: class {tight|normal|wide|pause}, ~{px} empty vertical space
2. ...

TEXT RULES:
- Balloon and caption strings must be exact; do not translate or paraphrase
- Reading order: top → bottom (then left → right within a cut if needed)
- Do not obscure faces or key action with balloons
- Full-color webtoon style suitable for smartphone vertical reading
- Lettering and balloon chrome must match the series illustration guide on every tile
```

**Must follow:**
- The **IMPORTANT** block is present in every tile/page YAML before generation
- **`illustration-guide.md` tokens** are present (art + balloon/caption recipes); no per-tile lettering invention
- Default path uses Flash + `1K` + `9:16` tiles (typically 2); Pro + `2K` + `1:8` only when explicitly chosen
- Never Flash + `2K` + `1:8` for balloon/caption pages
- All locked balloon/caption lines appear in the image **exactly** (character by character)
- No text mutation of locked strings (including particles)
- Cut order top→bottom preserved across tiles; cuts visually separated by clear gutters (unless design says full-bleed merge)
- **No** on-image labels: no “CUT 1/2/3”, panel numbers, index text, or extra captions beyond locked balloon/caption text
- Final look = finished webtoon strip segment for smartphone vertical reading, not a captioned picture-book page or a labeled storyboard

---

## Phase 2: Stitch Episode Scroll (user-requested only)

**Hard rule:** Do **not** auto-stitch (or re-stitch) the episode scroll when page images are created or updated. Page-level tile stitch (`{page}-a` + `{page}-b` → `{page}.png`) is part of Phase 1; episode-scroll stitch is a separate, explicit step.

1. Prerequisites: every page has a final `{page}.png` (from tile stitch on the default path, or the Pro single frame).
2. Run this phase **only on explicit user request** (e.g. “스크롤 스티치해줘”, “episode-scroll 만들어줘”) after page images for the episode are complete — never as a side effect of regenerating a page.
3. On request, concatenate page PNGs **top → bottom** into:
   - `images/{nnn}-{episode-slug}/{nnn}-{episode-slug}.png`
4. Optionally resize the scroll (or each page) to **target width** (**768px** default, or overview lock) while preserving aspect.
5. If the user/portal needs split upload files, export additional slices (e.g. max height ~1280px) **without changing art** — document slice boundaries next to the files if useful.
6. Show the stitched scroll to the user for review before G4.

If an existing `{nnn}-{episode-slug}.png` is stale after page edits: leave it, note that it is out of date, and re-stitch **only** when the user asks again.

If stitch reveals a **design** problem (wrong cut order, missing beat): rollback to Stage ② — do not “fix” in stitch.

---

## Phase 3: Visual Consistency Review

1. Compare generated page images and the episode scroll within each episode **and against `illustration-guide.md`**.
2. Check specifically: same art medium/line feel; same balloon/caption lettering and chrome; same outside-cut fill.
3. If **visual inconsistencies** (palette/style/identity/**typography drift**) appear:
   - regenerate only the affected image(s) with stronger guide tokens in YAML
   - do not change locked story/text
   - max 2 retries per image
   - do **not** auto re-stitch the episode scroll — mark it stale if it exists; re-run Phase 2 only on explicit user request
4. If inconsistencies indicate a **Design gap** (including incomplete `illustration-guide.md`):
   - rollback to Stage ②
   - re-evaluate and regenerate affected assets

---

## Approval Gate G4 — Final Result

Confirm with user:
1. All reference images — quality acceptable? Style matches `illustration-guide.md`?
2. All page images — tiles stitched; cuts, balloons, captions faithful to lock (exact text; no CUT/panel labels); **lettering/chrome consistent across pages**?
3. Episode scroll — continuous vertical read feels right? (only if user requested Phase 2 stitch)
4. Final result in `output/{webtoon-slug}-final/` — ready to deliver?

**Do not deliver until G4 is approved.**
