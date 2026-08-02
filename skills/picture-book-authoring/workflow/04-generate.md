# Stage ④ — Generate (Reference Images → Page Images)

**Prerequisites:** Stage ③ story lock approved (per episode).

**Gate artifacts:**
- reference images and page images generated into `images/characters`, `images/locations`, `images/{nnn}-{episode-slug}/`
- final assembly confirmation in `output/{book-slug}-final/` (user approval)

**Next stage:** none (this is the terminal stage for the skill)

---

## Hard Boundary

To preserve story lock:
- Allowed in ④:
  - image YAML prompts that implement locked illustration guides **and** locked `illustration-guide.md` (art / layout / typography)
  - re-rendering reference/page PNGs for quality/consistency (max 2 retries per image)
  - rendering the locked `rendering text` inside each page image
- Not allowed in ④:
  - rewriting episode/page story or rendering text
  - changing illustration guide meaning (only prompt phrasing/tightening)
  - inventing new fonts, text-box styles, outline recipes, or art-medium treatments not in `illustration-guide.md`
  - inventing new seats/L-R or **wholesale situation props** (different meal, toys, cups) when a staging is cited — update Design staging first
  - **inventing architecture/fixtures missing from the cited location position×view reference** (e.g. add a door not in the PNG) — add a Scene List row + Phase 0 ref first (`workflow/reference-models.md` §3.1)
  - placing characters in impossible spots relative to the cited view/staging (e.g. behind a faucet when the view/staging never defined that station)

If a problem is a **Design gap** (wrong cast/place/state/unclear guide / missing or weak `illustration-guide.md`):
- Stop generation
- Roll back to Stage ②
- Re-run Stage ③ for affected episode(s)

---

## Phase 0: Reference Image Generation (Characters + Locations + Stagings)

Goal: Generate reference models (character · location · staging — `workflow/reference-models.md`) as reference images so every later page image has a consistency base.

**Required reading before any YAML:** `{project-root}/illustration-guide.md` (§1 Art style). Apply the same medium / palette / line / proportion tokens to every reference image. Full rules: `workflow/illustration-guide.md`.

**Included:** Character turnarounds; neutral multi-view locations; staging **establishing** (default one; Character+Location refs required; lean WHO–WHERE–WHAT). YAML concretizes approved md; design↔prompt fidelity; explicit approval before every generate.
**Excluded:** Expression/pose; weather/time in location PNGs; OTS/CU as staging refs; metadata labels on page images.

### 0.1 Generation order (per catalogs)

1. Character base references: `images/characters/{character-slug}.yaml/.png`
2. Character state variants: `images/characters/{character-slug}-{state-slug}.yaml/.png`
3. Location base references: `images/locations/{location-slug}.yaml/.png`
4. Location position/view references: `images/locations/{location-slug}-{position-slug}-{view-slug}.yaml/.png`
5. Location state variants: `images/locations/{location-slug}-{position-slug}-{view-slug}-{state-slug}.yaml/.png`
6. Staging establishing (default one): `images/stagings/{staging-slug}-establishing.yaml/.png` — Character+Location in `params.references`; lean WHO–WHERE–WHAT; optional reverse/detail only if md lists them (`workflow/reference-models.md` §4.2–4.3)

### 0.2.1 Sample: reference image YAML structure

**Character base YAML structure** (`images/characters/{character-slug}.yaml`) — concretize `characters/{slug}.md` base state:

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {picture-ready character description in English — concretized from character md}
    A children's book illustration CHARACTER TURNAROUND reference sheet on one image:
    front view, side view, and back view of the SAME outfit/look, plus key accessories/identity gear
    clearly readable. Neutral expression; no dramatic acting pose.
  size: 1K
  aspectRatio: "1:1"
  seed: null

output: "./images/characters/{character-slug}.png"
```

**Location base YAML structure** (`images/locations/{location-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {picture-ready location description in English}
    A children's book illustration background reference (location base / establishing view)...
  size: 1K
  aspectRatio: "4:3"
  seed: null

output: "./images/locations/{location-slug}.png"
```

**Location position/view YAML structure** (`images/locations/{location-slug}-{position-slug}-{view-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {picture-ready location description in English}
    A children's book illustration background reference...
  size: 1K
  aspectRatio: "4:3"
  seed: null

output: "./images/locations/{location-slug}-{position-slug}-{view-slug}.png"
```

### 0.2 Per-image YAML approval (MANDATORY — never auto-generate)

**Hard gate:** Do **not** call image generate until the user has seen the full YAML and given an **explicit approval command**. Never chain autonomous generate calls.

For each reference image:
1. Create image-generation YAML that reflects the approved Design notes.
   - YAML `title/message/design` role boundary:
     - `message`: only the reference image’s “meaning (identity / structure consistency)” briefly
     - `design`: visual implementation; **design ↔ prompt fidelity** for proper nouns / logos / gear; staging = lean WHO–WHERE–WHAT + Character/Location in `params.references`
2. Show the full YAML to the user and request explicit approval.
3. Only after explicit approval, call the MCP image generation.
4. Verify visual quality (do not treat it as story redesign).
5. For stagings: verify L/C/R + props on the establishing PNG.

---

## Phase 1: Page Image Generation (Include text overlay)

**Required before every page YAML:** open `{project-root}/illustration-guide.md` and apply art style (§1), layout (§2), and typography roles (§3). Do not invent fonts or text boxes. Procedure: `workflow/illustration-guide.md`.

For each episode:
- Generate pages in order, **starting with Page 0 (cover)** → `00`, then body pages `01`, `02`, …
- For each page index (e.g. `{00}`, `{01}`, ...), generate:
  - `images/{nnn}-{episode-slug}/{page-idx}.yaml`
  - `images/{nnn}-{episode-slug}/{page-idx}.png`

### 1.0 Page 0 cover generation (mandatory)

Page 0 is the **episode cover**. Treat `00.yaml` / `00.png` as a finished picture-book cover, not a story-body page.

1. Create `00.yaml` that:
   - **applies locked `illustration-guide.md`** (art style + cover layout + title typography §3.1)
   - uses character/location reference PNGs via `params.references` for the cover cast and establishing place
   - implements only the locked Page 0 Direction / Picture carries
   - renders locked **title** (and optional subtitle / credit) as the primary on-image text — **not** story-body prose
   - reserves a clear **title safe zone**; title is the most prominent text after the hero figure
   - YAML role boundary:
     - `title`: e.g. `{episode-slug} / Page 00 (Cover)`
     - `message`: 1–2 sentences from Page 0 story — invite the child to open the book (world / hero / mood); no interior plot dump
     - `design`: cover composition + **guide typography/placement** + refs — grounded in locked Page 0 fields + `illustration-guide.md`
2. Show YAML → explicit user approval → MCP generate (same gate as every other page)
3. Verify: reads as a cover; title readable; faces/key hero not obscured; aspect matches locked page canvas from overview; **matches series guide** (not a one-off font)

**Cover prompt cues (add to design / prompt):** `Picture book COVER illustration`; title largest/boldest overlay per guide §3.1; inviting shelf appeal; do not advance Page 1+ plot.

### 1.1 Per-page YAML approval (MANDATORY)

For each page image (including Page 0):
1. Create `{page}.yaml` that:
   - **cites and applies `{project-root}/illustration-guide.md`** (mandatory — art style + layout + text-role recipes; include §5 checklist in `design`)
   - uses character/location/staging reference PNGs via `params.references` (image-generation skill; do not use `params.images`); keep seating/L-R from staging when cited
   - **when Staging is cited:** attach the cited staging ref PNG (and establishing if needed for prop lock) under `params.references` (not `params.images`); YAML must keep **same seats and same situation props/table dressing** as that staging — Location alone is the permanent set (empty dining room), not tonight’s meal
   - **visibility coverage:** cite only the location position×view PNG that already shows every must-see architectural/fixture element; if the page needs a door/hallway/other wall not in the current ref → stop and add that ref in Design/Phase 0 — do **not** prompt “add a door”
   - place characters only on stations allowed by staging + cited view (not behind fixtures the blocking never defined)
   - implements only the locked page Direction / Picture carries (page-specific action — not new style; progressive prop change only if staging Continuity rules allow)
   - renders the locked `rendering text` inside the image as part of the illustration
   - uses the locked `rendering text` **verbatim** (keep language as locked — no translation, paraphrase, summary, or word-order change)
   - assigns each overlay line a **text role** from the guide (title / narration / dialogue / emphasis) and uses **only** that role’s font/box/effect recipe
   - (note) if `evaluations/{nnn}-{episode-slug}.md` notes page-specific placement, apply placement only — do **not** override series typography from the guide
   - YAML `title/message/design` respect role boundaries and avoid duplication:
     - `title`: short image name such as `{episode-slug} / Page {idx}` (Page 0: include `(Cover)`; no narrative / shot list)
     - `message`: 1–2 sentences from `Page story` on “what must be communicated (meaning / emotion / relationship / beat)” only (no shot list / layout / font duplication)
     - `design`: “how to realize it in the picture” from page Direction + **series `illustration-guide.md` tokens** + (TextBox role / anchor area per guide) + **staging continuity (seats + props)** when cited — grounded in episode design
2. Show YAML to user for explicit review and approval
3. On explicit approval only → call MCP generate
4. Verify:
   - visual fidelity to the locked guide **and** series `illustration-guide.md`
   - text readability, correct placement, and **same typography recipe as other pages**
   - if staging cited: seats/L-R and situation props match staging refs (no wholesale food/prop swap)

### 1.2 Prompt pattern (guideline)

Using the provided reference images:
- Keep character appearance, setting, lighting, and style exactly the same
- Match series art style from locked `illustration-guide.md` §1
- When a staging reference is provided: keep WHO SITS WHERE and the SITUATION PROPS / TABLE DRESSING exactly as in that staging image (same meal layout, toys, cups — do not invent a new table). Progressive depletion only if Design Continuity rules allow
- Do **not** invent architecture missing from the location reference (no new doors, walls, or fixtures). If the scene needs them, a different location reference must be used
- Place characters only where the location view and staging allow (do not put a person behind a faucet/sink unless that station is designed)
- Change only the character’s action described in the locked page Direction
- Add extra elements only if the locked page fields list them

TEXT OVERLAY (Episode rendering text) — styles from `illustration-guide.md` §3:
- Render every line from locked `rendering text` directly into the image
- Assign each line a **role** (title / narration / dialogue / emphasis) and use **only** that role’s locked recipe (feel, size, color, effects, box)
- Placement / reading order:
  - Follow guide §2 text zones + page Direction
  - Text boxes should follow a left → right → top → bottom reading flow by default
  - For multiple lines in the same area, place lines so they continue naturally downward
- Do not invent a new font, box, glow, or outline for this page
- Do not obscure faces/key visuals with text

**(Reference) SKILL template prompt pattern**

```
Using the provided reference images:
- Keep the character's appearance, setting, lighting, and style exactly the same
- Series art style (from illustration-guide.md): {medium}, {palette}, {line}, {proportions}
- Staging continuity (if staging ref provided): same seats/L-R and same situation props/table dressing as Image {n}; do not replace the meal/toys/cups with a different layout
- Change only the character's action: {action from locked page Direction}
- [Add any new elements only if listed in the locked page fields]

SERIES TEXT RULES (from illustration-guide.md — do not invent):
- Narration: {locked feel / size / color / effects / box}
- Dialogue/exclamation: {locked feel / size / color / effects / box}
- Title (cover only): {locked recipe}
- Text zones: {locked defaults}; never cover faces

TEXT OVERLAY — render these episode `rendering text` lines directly into the image as part of the illustration:
{For each line in rendering text, specify:}
1. {TextBox (ReadingOrder)}: {text} — role: {title|narration|dialogue|emphasis}; Anchor {per guide + page}; apply locked recipe for that role only
2. ...
```

**Multilingual example (Korean overlay strings only; prompt English):**

```
TEXT OVERLAY — render these locked rendering-text lines verbatim:
1. TextBox: “미나는 창밖을 보았다.” — top-left, narration, smaller soft rounded serif, warm shadow
2. TextBox: “구름이 천천히 지나가고 있었다.” — below previous, narration, same style
```

**Text overlay design rules (must follow):**
- All rendering text from the locked episode design must appear in the image
- No text mutation: keep wording / spelling / punctuation of `rendering text` exactly
- **Typography and text boxes come from `illustration-guide.md` only** — same recipes across pages
- Text must feel integrated into the illustration, not a plain overlay
- Text placement must respect guide text zones / reading order and avoid blocking key visuals
- **Page 0 (cover):** title role from guide §3.1; treat as finished cover typography
- Place text in areas that do not obscure characters' faces or key visual elements
- For ages 8+ & text-heavy pages: use guide §3.5 box rules (do not invent a new panel style)
- The final image should look like a finished picture book page where illustration and text work together **and match the series guide**

---

## Phase 2: Visual Consistency Review

1. Compare generated page images within each episode **and against `illustration-guide.md`**.
2. Check specifically: same art medium/line feel; same narration/dialogue/title recipes; same text-box chrome (if any).
3. For each staging span: compare pages that cite the same staging — seats/L-R and **situation props** (meal layout, toys, cups) must stay recognizable; wholesale swaps → regenerate with staging refs.
4. If **visual inconsistencies** (palette/style/identity/**typography drift**/prop drift) appear:
   - regenerate only the affected image(s) with stronger guide + staging tokens in YAML
   - do not change locked story/text
   - max 2 retries per image
5. If inconsistencies indicate a **Design gap** (wrong state/place/unclear guide / incomplete `illustration-guide.md` / missing staging props map):
   - rollback to Stage ②
   - re-evaluate and re-generate affected episode page images

---

## Approval Gate G4 — Final Result

Confirm with user:
1. All reference images — quality acceptable? Style matches `illustration-guide.md`? Staging refs show seats **and** situation props?
2. Page 0 cover (`00`) — reads as a finished cover (title + hero/world invite) with **guide title typography**?
3. All page images — consistent art + **consistent fonts/text boxes**; staging spans keep seats + props; appealing, faithful to locks?
4. Final result in `output/{book-slug}-final/` — ready to deliver?

**Do not deliver until G4 is approved.**
