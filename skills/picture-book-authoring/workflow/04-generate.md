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
  - image YAML prompts that implement locked illustration guides
  - re-rendering reference/page PNGs for quality/consistency (max 2 retries per image)
  - rendering the locked `rendering text` inside each page image
- Not allowed in ④:
  - rewriting episode/page story or rendering text
  - changing illustration guide meaning (only prompt phrasing/tightening)

If a problem is a **Design gap** (wrong cast/place/state/unclear guide):
- Stop generation
- Roll back to Stage ②
- Re-run Stage ③ for affected episode(s)

---

## Phase 0: Reference Image Generation (Characters + Locations + Stagings)

Goal: Generate reference models (character · location · staging — `workflow/reference-models.md`) as reference images so every later page image has a consistency base.

**Included:** Character face/body/outfit·gear (per state), location set structure (per state), staging ensemble blocking (2–3 views).
**Excluded:** Expression/pose, lighting/time/weather/mood, one-off camera, transient elements — page direction.

### 0.1 Generation order (per catalogs)

1. Character base references: `images/characters/{character-slug}.yaml/.png`
2. Character state variants: `images/characters/{character-slug}-{state-slug}.yaml/.png`
3. Location base references: `images/locations/{location-slug}.yaml/.png`
4. Location position/view references: `images/locations/{location-slug}-{position-slug}-{view-slug}.yaml/.png`
5. Location state variants: `images/locations/{location-slug}-{position-slug}-{view-slug}-{state-slug}.yaml/.png`
6. Staging ensemble refs (2–3 views): `images/stagings/{staging-slug}-{establishing|reverse|detail}.yaml/.png` — see `workflow/reference-models.md`

### 0.2.1 Sample: reference image YAML structure

**Character base YAML structure** (`images/characters/{character-slug}.yaml`):

```yaml
type: image
model: gemini-3.1-flash-image

params:
  prompt: |
    {picture-ready character description in English}
    A children's book illustration style character reference sheet...
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

### 0.2 Per-image YAML approval (MANDATORY)

For each reference image:
1. Create image-generation YAML that reflects the approved Design notes.
   - YAML `title/message/design` role boundary:
     - `message`: only the reference image’s “meaning (identity / structure consistency)” briefly (no story / layout / font-detail duplication)
     - `design`: only the reference image’s “visual implementation” (character face/body/outfit state, location physical-structure state, framing anchors; no on-image text)
2. Show the full YAML to the user and request explicit approval.
3. Only after explicit approval, call the MCP image generation.
4. Verify visual quality (do not treat it as story redesign).

---

## Phase 1: Page Image Generation (Include text overlay)

For each episode:
- For each page index (e.g. `{00}`, `{01}`, ...), generate:
  - `images/{nnn}-{episode-slug}/{page-idx}.yaml`
  - `images/{nnn}-{episode-slug}/{page-idx}.png`

### 1.1 Per-page YAML approval (MANDATORY)

For each page image:
1. Create `{page}.yaml` that:
   - uses character/location/staging reference PNGs; keep seating/L-R from staging when cited
   - implements only the locked illustration guide action
   - renders the locked `rendering text` inside the image as part of the illustration
   - uses the locked `rendering text` **verbatim** (keep language as locked — no translation, paraphrase, summary, or word-order change)
   - (note) if `evaluations/{nnn}-{episode-slug}.md` includes an Illustration specialist text visual-effects guide, apply that guide to TextBox anchor areas / font / emphasis rules
   - YAML `title/message/design` respect role boundaries and avoid duplication:
     - `title`: short image name such as `{episode-slug} / Page {idx}` (no narrative / shot list)
     - `message`: 1–2 sentences from `Page story` on “what must be communicated (meaning / emotion / relationship / beat)” only (no shot list / layout / font duplication)
     - `design`: only “how to realize it in the picture” from `Illustration guide` + (TextBox / anchor area / font / color / emphasis / panel split) — grounded in episode design
2. Show YAML to user for explicit review and approval
3. On explicit approval only → call MCP generate
4. Verify:
   - visual fidelity to the locked guide
   - text readability and correct placement

### 1.2 Prompt pattern (guideline)

Using the provided reference images:
- Keep character appearance, setting, lighting, and style exactly the same
- Change only the character’s action described in the locked illustration guide
- Add extra elements only if the locked guide lists them

TEXT OVERLAY (Episode rendering text):
- Render every line from locked `rendering text` directly into the image
- Placement / reading order (important):
  - Text boxes should follow a left → right → top → bottom reading flow by default
  - For multiple lines in the same area, place lines so they continue naturally downward
- Dialogue / exclamations: largest + bold (thin glow/outline when useful)
- Narration: smaller + soft rounded serif (or similar) + warm shadow
- Key words/phrases: emphasize by (a) larger size or (b) color change (or bold)
- Do not obscure faces/key visuals with text

**(Reference) SKILL template prompt pattern**

```
Using the provided reference images:
- Keep the character's appearance, setting, lighting, and style exactly the same
- Change only the character's action: {action from locked illustration guide}
- [Add any new elements only if listed in the locked illustration guide]

TEXT OVERLAY — render these episode `rendering text` lines directly into the image as part of the illustration:
{For each line in rendering text, specify:}
1. {TextBox (ReadingOrder)}: {text} — {Anchor area (e.g. top-left / top-right / bottom-left / bottom-right), size, font style, color, emphasis level}
2. ...
{Place text in areas that do not obscure characters' faces}
{Key dialogue or exclamations should be largest and boldest}
{Narration lines should be smaller, in soft rounded serif/sans appropriate to the rendering language}
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
- Text must feel integrated into the illustration, not a plain overlay
- Text placement must respect reading order and avoid blocking key visuals
- Key dialogue / exclamations → largest, boldest, with glow
- Narration lines → smaller, soft rounded serif/sans appropriate to the rendering language, white with warm shadow
- Place text in areas that do not obscure characters' faces or key visual elements
- For ages 8+ & text-heavy pages:
  - Prefer separating body copy (in-image story) into its own text area (e.g. soft box/panel) for readability
  - Apply emphasis rules first to dialogue / key phrases inside that panel
- The final image should look like a finished picture book page where illustration and text work together

---

## Phase 2: Visual Consistency Review

1. Compare generated page images within each episode.
2. If **visual inconsistencies** (palette/style/identity) appear:
   - regenerate only the affected image(s)
   - do not change locked story/text
   - max 2 retries per image
3. If inconsistencies indicate a **Design gap** (wrong state/place/unclear guide):
   - rollback to Stage ②
   - re-evaluate and re-generate affected episode page images

---

## Approval Gate G4 — Final Result

Confirm with user:
1. All reference images — quality acceptable?
2. All page images — consistent, appealing, and faithful to locked guides?
3. Final result in `output/{book-slug}-final/` — ready to deliver?

**Do not deliver until G4 is approved.**
