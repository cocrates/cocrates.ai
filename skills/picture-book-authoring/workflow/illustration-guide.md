# Illustration Guide — Series Visual & Typography Lock (Picture Book)

**Not a numbered stage.** Author `{project-root}/illustration-guide.md` in Stage ② (after G1). Read again at Stage ③ Evaluate and **every** Stage ④ YAML.

## Purpose

Lock **art style, page layout, fonts, text boxes, and text-role recipes** once for the whole book so pages do not invent different typography or chrome each time.

| Layer | Owns | Does not own |
|-------|------|--------------|
| **`illustration-guide.md`** (series) | Medium, palette, line, layout safe areas, font/box recipes per text role | Per-page story, exact overlay strings, one-off pose |
| **Episode page fields** | Page story, rendering text, Direction, placement intent | New fonts, new box styles, new outline recipes |
| **Stage ④ YAML** | Implements locked guide + locked page fields | Style invention |

**Rule:** If two pages could look like different books (font, box, line quality), the guide is too weak or YAML ignored it.

---

## When to write / update

1. **Create** early in Stage ② from `overview.md` → Illustration Style seed.
2. **Approve** with G2 (same gate as other Design artifacts).
3. **Update** only via Design → re-G2 → re-Evaluate affected episodes → regenerate page images. Never “fix font in one YAML only.”

---

## Project file template

Write `{project-root}/illustration-guide.md`:

```markdown
# Illustration Guide

Series-wide visual and typography lock for `{book-title}`.
**Every** Stage ④ reference and page YAML must follow this file.
Do not invent new fonts, text-box styles, or art treatments per page.

## 1. Art style
- Medium: {watercolor / digital painting / collage / soft cel / …}
- Palette: {locked palette — key hues + forbidden clashes}
- Line: {outline weight / softness / roundness}
- Shading / texture: {flat / soft wash / grain / …}
- Character proportions: {age-appropriate silhouette rules}
- Background detail: {simple / medium / rich — keep consistent}
- Mood defaults: {warm / dreamy / comic / …}
- References: {artists / works — feel only, not to copy}

## 2. Page layout
- Canvas: {aspect e.g. 3:4 or 4:3; match overview if set}
- Margins / safe area: {e.g. keep faces and title out of outer ~8%}
- Default text zones: {e.g. top band narration/title; bottom dialogue; avoid center face}
- Reading order: left → right, top → bottom (default)
- Cover (Page 0): {title zone; hero center; optional credit micro-zone}
- Body pages: {where narration vs dialogue usually sit}
- Face / key-visual rule: text must not cover faces or primary action

## 3. Typography & text roles
{One locked recipe per role. YAML chooses role + placement only.}

### 3.1 Title (Page 0 cover)
- Feel: {rounded display / friendly bold / …}
- Size: largest on cover
- Color: {…}
- Effects: {soft glow / thick outline / none}
- Optional subtitle: {smaller; color; under title}
- Optional credit: {smallest; bottom; muted}

### 3.2 Narration (body)
- Feel: {soft rounded serif **or** rounded sans — pick one}
- Size: medium-small (read-aloud readable)
- Color: {…}
- Effects: {warm shadow **or** soft outline — pick one}
- Box: {none / soft rounded panel / translucent band — pick one}

### 3.3 Dialogue / exclamation
- Feel: {same family as narration or slightly bolder — lock choice}
- Size: larger than narration; exclamations largest + boldest among body text
- Color: {…}
- Effects: {glow / outline recipe — same every page}
- Box: {none / soft oval — pick one}

### 3.4 Emphasis (keywords)
- How: {larger size OR color change OR bold — allowed methods only}
- When: sparingly

### 3.5 Text boxes / panels (ages 8+ or text-heavy)
- When used: {never / only dense pages / always for narration}
- Shape / fill / border / padding: {lock each}

## 4. Consistency rules
- Same art-style tokens from §1 in every page YAML `design`
- Same typography recipe per role (§3) across all pages and episodes
- Do **not** change font family, box style, or outline recipe page-to-page
- Page-specific only: placement, which role applies, locked `rendering text` strings
- Need a new style? Update this guide + re-approve — do not invent in YAML

## 5. Stage ④ checklist (required in every page YAML design)
- [ ] Art style tokens from §1 applied
- [ ] Layout / safe area / text zones from §2 respected
- [ ] Each overlay line uses a locked role from §3
- [ ] No new font, box, or effect invented for this page
- [ ] Locked rendering text verbatim; faces not covered
```

---

## Stage ④ application (mandatory)

Before writing any page YAML:

1. Open `{project-root}/illustration-guide.md`.
2. Copy §1 art-style tokens into `design` (English prompt phrasing OK).
3. For each overlay line: assign a **role** from §3 and use that recipe only; set placement from §2 + page Direction.
4. Include the §5 checklist (or equivalent explicit statements) in `design`.
5. If the guide conflicts with a page field → stop; fix Design (guide or episode), do not silently override in YAML.

Reference-image YAMLs (Phase 0): apply §1 art style; **no** on-image story text unless the guide explicitly says otherwise for a sheet type.
