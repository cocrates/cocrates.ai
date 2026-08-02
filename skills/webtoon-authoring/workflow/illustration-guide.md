# Illustration Guide — Series Visual & Typography Lock (Webtoon)

**Not a numbered stage.** Author `{project-root}/illustration-guide.md` in Stage ② (after G1). Read again at Stage ③ Evaluate and **every** Stage ④ tile/page YAML.

## Purpose

Lock **art style, vertical layout chrome, balloon/caption recipes, and fonts** once for the series so strip tiles do not invent different lettering or panel chrome each time.

| Layer | Owns | Does not own |
|-------|------|--------------|
| **`illustration-guide.md`** (series) | Medium, palette, line, outside-cut fill look, balloon shapes, caption boxes, lettering recipes | Per-cut story, exact balloon strings, one-off gesture |
| **Episode cut fields** | Cut story, balloons/captions tables, Direction, gutters | New balloon fonts, new caption box styles |
| **Stage ④ YAML** | Implements locked guide + locked cut fields | Style invention |

**Rule:** If two pages could look like different series (balloon typeface, caption box, line weight), the guide is too weak or YAML ignored it.

---

## When to write / update

1. **Create** early in Stage ② from `overview.md` → Art style + Canvas seed.
2. **Approve** with G2.
3. **Update** only via Design → re-G2 → re-Evaluate → regenerate. Never “fix lettering in one tile YAML only.”

---

## Project file template

Write `{project-root}/illustration-guide.md`:

```markdown
# Illustration Guide

Series-wide visual and typography lock for `{webtoon-title}`.
**Every** Stage ④ reference and page/tile YAML must follow this file.
Do not invent new balloon fonts, caption boxes, or art treatments per page.

## 1. Art style
- Medium: {digital color comic / ink+flat / painterly / …}
- Palette: {locked palette — key hues + grade notes for flashback if any}
- Line: {outline weight / cleanliness}
- Shading / screentone / flat: {…}
- Character proportions: {series silhouette rules}
- Background detail: {simple / medium / rich — keep consistent}
- Mood defaults: {…}
- References: {works — feel only}

## 2. Vertical layout & chrome
- Target width: {from overview, e.g. 768px}
- Cut horizontal fit: {full bleed / 30–50px even sides}
- Outside-cut fill: {white / black / theme — exact look}
- Panel / cut border: {hard box / soft / open/bleed rules}
- Gutter look: {empty fill matching outside-cut; no decorative junk unless locked}
- Phone readability: faces + balloons readable without pinch-zoom

## 3. Balloons (speech / thought / shout / …)
{Lock shape + lettering per balloon type used in the series.}

### 3.1 Speech
- Shape: {round / oval / soft rectangle}
- Tail: {simple pointer rules}
- Fill / stroke: {white fill, thin dark stroke, …}
- Lettering feel: {clean rounded sans / … — one recipe}
- Size: readable on phone; denser lines stay same typeface
- Placement: prefer locked cut placement; never cover faces/key action

### 3.2 Thought
- Shape: {cloud / …}
- Lettering: {same family or slightly softer — lock}

### 3.3 Shout / exclamation
- Shape: {jagged / spiky / …}
- Lettering: {bolder / larger — same family unless locked otherwise}

### 3.4 Other types
- {whisper / robot / phone — only if used; lock each}

## 4. Captions / narration boxes
- When used: {sparingly / establishing only / …}
- Shape: {soft rectangle / rounded bar}
- Fill / stroke: {…}
- Lettering: {smaller than speech or same family — lock}
- Placement defaults: {top of cut / bottom band}

## 5. Consistency rules
- Same art-style tokens from §1 in every tile/page YAML
- Same balloon/caption recipes (§3–4) across all pages and episodes
- Do **not** change lettering family, balloon chrome, or caption box page-to-page
- Page/cut-specific only: placement, balloon type from lock, exact strings
- Need a new style? Update this guide + re-approve — do not invent in YAML

## 6. Stage ④ checklist (required in every tile/page YAML design)
- [ ] Art style tokens from §1 applied
- [ ] Layout / outside-cut / border rules from §2 respected
- [ ] Each balloon uses a locked type recipe from §3
- [ ] Each caption uses §4 recipe (or omitted if none)
- [ ] No new font, balloon chrome, or caption box invented for this tile
- [ ] Locked balloon/caption text verbatim; faces not covered
- [ ] IMPORTANT block present (webtoon generate rules)
```

---

## Stage ④ application (mandatory)

Before writing any tile/page YAML:

1. Open `{project-root}/illustration-guide.md`.
2. Copy §1–2 tokens into `design` (with the mandatory IMPORTANT block from `04-generate.md`).
3. For each balloon/caption: use the locked type recipe; placement from cut table + §2/§3/§4.
4. Include the §6 checklist (or equivalent explicit statements) in `design`.
5. Guide vs cut conflict → stop; fix Design, do not override in YAML.

Reference-image YAMLs (Phase 0): apply §1; **no** balloons/captions on refs unless the guide explicitly allows a sheet type.
