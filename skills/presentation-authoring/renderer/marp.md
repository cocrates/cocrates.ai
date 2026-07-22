# Marp Renderer (Phase B — PSM)

Read this file when **translating** approved Phase A page models (`slides/slide-*.md`) into a Marp deck.

Phase B is **renderer-specific authoring**: PIM → PSM. Do not redesign narrative or on-slide copy here.

- Phase A (PIM): main [`SKILL.md`](../SKILL.md) — define → architecture → detail design → `slides/slide-*.md`
- Other renderers: [`README.md`](README.md)

Match the user's language in artifact content.

## When to Use

| Situation | Action |
|-----------|--------|
| User asks to generate / render / export a Marp deck | Follow this file |
| `slides/slide-*.md` approved but no `marps/` yet | Create `marps/` then follow the pipeline |
| User asks for PDF/PPTX/HTML from Marp | Compile `{title-slug}.md`, then Marp CLI export |
| User asks for image→PPTX or other renderers | See [`README.md`](README.md) — use the matching spec |

**Prerequisite:** Phase A (PIM) complete — every slide has approved `slides/slide-{nn}.md`, templates in use are approved, and optional `scripts/` may exist.

## Input (PIM)

| Artifact | Role |
|----------|------|
| `slides/slide-{nn}.md` | Renderer-independent page model — layout, copy, spatial zones (**canonical on-slide content**) |
| `slides.md` | Detail design; per-slide `Template` assignment |
| `templates/template-{nn}.md` | Shared visual foundation (PIM, renderer-agnostic) |
| `scripts/script-{nn}.md` | Optional; when present, `## Script` section is the sole source of presenter notes |

## Output Layout

| Path | Role |
|------|------|
| `marps/template-{nn}.md` | PSM template — Marp chrome + CSS for one PIM template |
| `marps/marp-{nn}.md` | PSM per-slide preview deck (template chrome + slide body + script notes) |
| `{title-slug}.md` | PSM compiled deck (all templates + all slides + scripts) |

## Core Translation Model

Each slide is produced from **three inputs**:

| Input | Supplies |
|-------|----------|
| **Template** (`marps/template-{nn}.md`) | Page chrome — regions, background, typography, palette, shared CSS |
| **Slide** (`slides/slide-{nn}.md`) | On-slide content — key phrases, spatial zones, content layout, density |
| **Script** (`scripts/script-{nn}.md`) | Presenter notes only — exact `## Script` section in an HTML comment; never on-slide body |

```
templates/template-{nn}.md  ──►  marps/template-{nn}.md
                                        │
slides/slide-{nn}.md  ──read Template──┤
                                        ├──►  marps/marp-{nn}.md
scripts/script-{nn}.md  ────────────────┘
```

**Per-slide authoring rule:**

1. Read `slides/slide-{nn}.md`.
2. Resolve **`Template:`** (e.g. `template-02`) — must match `slides.md`.
3. Read `marps/template-{nn}.md` for that template (create it first if missing).
4. Read `scripts/script-{nn}.md` when scripts exist — extract the **`## Script`** section (see **Presenter notes from `## Script`**).
5. Write `marps/marp-{nn}.md` = **template chrome + slide content + script notes**.

- **Page structure / regions / style** → from the template.
- **Visible copy and layout of elements** → from the slide (faithful to Content Elements and Visual Design).
- **Presenter notes** → exact body of `## Script` from `scripts/script-{nn}.md` only.

## Pipeline (PIM → PSM)

| Step | Artifact | Role |
|------|----------|------|
| 0 | PIM inputs | Approved `slides/`, `templates/`, optional `scripts/` |
| 1 | `marps/template-{nn}.md` | Translate each used PIM template into Marp CSS/chrome |
| 2 | `marps/marp-{nn}.md` | Per-slide preview: template + slide + script |
| 3 | `{title-slug}.md` | Compiled deck: all templates, then each slide + script |
| 4 | HTML / PDF / PPTX | Optional Marp CLI export |

### Step 1 — Templates first

For every template ID referenced by any `slides/slide-*.md` (or `slides.md`):

1. Read `templates/template-{nn}.md` (PIM).
2. Write `marps/template-{nn}.md` (PSM) — Marp implementation of that chrome.
3. Use **merge-safe naming** (see **CSS naming for compilation**).

Do this **before** writing any `marps/marp-*.md`.

### Step 2 — Per-slide fragments

For each slide in `slides.md` order (including backup `B{nn}` when requested):

1. Read `slides/slide-{nn}.md` — resolve `Template`, content layout, density, Visual Design, Content Elements.
2. Read the matching `marps/template-{nn}.md`.
3. Read `scripts/script-{nn}.md` when present; copy the **`## Script`** section body verbatim into the fragment’s HTML comment.
4. Write `marps/marp-{nn}.md` as a **single-slide preview deck**.
5. Preview in Marp; fix layout until the slide uses the full frame without overflow or cramped corners.

### Step 3 — Compile

Merge into `{title-slug}.md` in this order:

1. Deck front matter (`marp: true`, default `theme`, `paginate`, `size`).
2. **All** `marps/template-*.md` style blocks (once each, merge-safe selectors).
3. For each slide in order: body from `marps/marp-{nn}.md` (no per-slide front matter) + presenter notes = exact `## Script` body from `scripts/script-{nn}.md`, separated by `---`.

Conceptual compiled shape:

```text
---
marp: true
theme: gaia
paginate: true
size: 16:9
---

<style>
/* from marps/template-01.md */
/* from marps/template-02.md */
</style>

<!-- slide 01 body (uses tpl-01) -->
<!-- exact ## Script body from scripts/script-01.md -->

---

<!-- slide 02 body (uses tpl-02) -->
<!-- exact ## Script body from scripts/script-02.md -->

---
…
```

## Built-in Themes

Use a built-in theme as the **base**, then layer template CSS on top.

| Theme | Alignment | Best for |
|-------|-------------|----------|
| **`gaia`** | Left-aligned content | Lectures, reports, technical talks, dense content |
| **`uncover`** | Center-aligned content | Keynotes, pitches, title slides, short punchy decks |

Record the deck default in `overview.md` **Visual Style** or per-template **Layout Defaults** when unclear. Default to **`gaia`**.

```yaml
---
marp: true
theme: gaia    # or uncover — see Templates (Marp mapping)
paginate: true
size: 16:9
---
```

Do **not** invent a separate custom `@theme` package file. Template chrome lives in `marps/template-{nn}.md` as **namespaced CSS** that compiles into the deck `<style>` block.

## CSS Naming for Compilation

Templates must merge into **one** deck stylesheet without selector clashes.

| Rule | Convention | Example |
|------|------------|---------|
| Template root class | `tpl-{nn}` | `tpl-01`, `tpl-02` |
| Region / element classes | `tpl-{nn}__{name}` (BEM-style) | `tpl-02__title`, `tpl-02__body`, `tpl-02__footer` |
| Shared tokens (optional) | `deck-{token}` only if identical across all templates | `deck-bg` — prefer per-template if values differ |
| Slide-only layout helpers | `s-{nn}__{name}` inside that slide fragment only | `s-11__cols` — do not put in template files |

**Apply the template class on the slide section:**

```markdown
<!-- _class: tpl-02 -->
```

**Selectors always scope under the template class:**

```css
section.tpl-02 {
  background: #0a0a0a;
  color: #ffffff;
  /* … */
}
section.tpl-02 h1 { /* title zone */ }
section.tpl-02 .tpl-02__body { /* content zone */ }
section.tpl-02 .tpl-02__footer { /* footer zone */ }
```

**Why:** compiled deck includes `template-01` + `template-02` + … in one `<style>` block, then slide bodies. Namespaced selectors keep chrome isolated and make compile a mechanical concat.

**Per-slide preview (`marps/marp-{nn}.md`):** include the **same** template CSS (from `marps/template-{nn}.md`) so the fragment previews correctly alone. On compile, emit each template’s CSS **once** at the top; strip duplicated template CSS from slide bodies.

## `marps/template-{nn}.md` — Template PSM

Translate `templates/template-{nn}.md` into Marp chrome. No slide-specific copy.

### Required contents

| Section | Source (PIM template) | Marp form |
|---------|----------------------|-----------|
| Theme choice | Layout Defaults (left / center) | Comment: `theme: gaia` or `uncover` for slides using this template |
| Root class | Template ID | `tpl-{nn}` |
| Background / palette | Background, Color Palette | `section.tpl-{nn} { … }` |
| Typography | Typography | `section.tpl-{nn} h1`, `p`, etc. |
| Regions | Title / content / footer zones | Classes `tpl-{nn}__title`, `tpl-{nn}__body`, `tpl-{nn}__footer` + placement rules |
| Layout defaults | Content alignment, whitespace | Flex/grid on `section`, padding, alignment |

### Skeleton

File shape for `marps/template-{nn}.md`:

~~~~markdown
# Marp Template {nn}: {Name from PIM}

<!-- theme hint: gaia | uncover -->
<!-- root class: tpl-{nn} -->

section.tpl-{nn} {
  /* background, color, padding, flex — fill the 16:9 frame */
}
section.tpl-{nn} h1 {
  /* title zone */
}
section.tpl-{nn} .tpl-{nn}__body {
  /* content zone */
}
section.tpl-{nn} .tpl-{nn}__footer {
  /* footer zone */
}

## Regions (Marp)

| Region | Markup / class |
|--------|----------------|
| Title | h1 or .tpl-{nn}__title |
| Content | .tpl-{nn}__body (or direct children under section) |
| Footer | .tpl-{nn}__footer |

## Notes

- Derivation from another template: inherit shared tokens, override only Changes from base.
- No slide titles, key messages, or per-slide copy.
~~~~

Wrap the CSS rules in a fenced `css` block (or a `<style>` excerpt) inside the file so compile can extract them mechanically. Keep prose sections (`## Regions`, `## Notes`) outside the extractable CSS.

### Theme choice (from PIM)

| Template spec | Marp approach |
|---------------|---------------|
| `Layout Defaults: left` / content template | Prefer **`gaia`** |
| `Layout Defaults: center` / cover | Prefer **`uncover`** |
| `Background`, `Color Palette` | Namespaced CSS on `section.tpl-{nn}` |
| `Common Elements` (header/footer) | Region classes in the template file |

### Region mapping

| Region (template spec) | Marp implementation |
|------------------------|---------------------|
| **Title zone** top | `h1` at top; styles under `section.tpl-{nn} h1` |
| **Title zone** center (cover) | Flex center on `section.tpl-{nn}`; `<!-- _paginate: false -->` on cover slides |
| **Content zone** | `.tpl-{nn}__body` — slide fills this with Content Elements |
| **Footer zone** | `.tpl-{nn}__footer` |

## `marps/marp-{nn}.md` — Per-Slide Fragment

Each file is a **single-slide preview deck** for review. It must implement **template + slide + script**.

### Authoring checklist (every slide)

1. Open `slides/slide-{nn}.md` — note `Template`, `Content layout`, `Density`, Visual Design, **Content Elements** table.
2. Open `marps/template-{nn}.md` for that template — copy its CSS into the fragment (or `@import` pattern is not used; **inline the template CSS** for standalone preview).
3. Open `scripts/script-{nn}.md` when present — embed the **`## Script`** section exactly (see below).
4. Place **every** Content Elements row in the correct spatial zone; use exact visible copy from the slide (key phrases only).
5. Apply content layout mode and density from the slide.
6. Check **full-frame layout** (see **Frame utilization**).

### Presenter notes from `## Script`

When `scripts/script-{nn}.md` exists, presenter notes are **mandatory** and must come **only** from that file’s **`## Script`** section.

**Extract:**

1. Open `scripts/script-{nn}.md`.
2. Find the heading line `## Script` (exact).
3. Take every line **after** that heading until the next `## …` heading (e.g. `## Key Phrases`, `## Transition From Previous`, `## Notes`) or end of file.
4. Trim at most one leading and one trailing blank line from that span.
5. Place the result **verbatim** inside a single HTML comment at the **end** of `marps/marp-{nn}.md` (after all on-slide markup).

**Include as-is:**

- All wording, punctuation, emphasis (`**…**`), line breaks, and blank lines inside `## Script`
- Stage directions in parentheses, e.g. `(천천히, 무게를 실어서)`

**Do not include** (other sections of the script file):

| Exclude | Examples |
|---------|----------|
| File title / metadata | `# Script: Slide …`, `## Section`, `## Timing` |
| After-script sections | `## Key Phrases`, `## Transition From Previous`, `## To Next`, `## Notes` |
| Invented or paraphrased speech | Summaries, “key phrases only”, improvised notes |

**Do not:**

- Paraphrase, shorten, or “improve” the spoken text
- Pull notes from `slides/slide-{nn}.md` or memory
- Put script prose on the slide body
- Omit the comment when a script file exists
- Embed `## Script` as a visible heading inside the comment (comment body = section **content only**, not the `## Script` line itself)

If `scripts/script-{nn}.md` is missing (scripts skipped in Phase A), omit the presenter-notes comment.

**Shape:**

```markdown
<!--
{exact body of ## Script from scripts/script-{nn}.md}
-->
```

**Check:** the HTML comment body must be **byte-for-byte identical** (aside from a single optional leading/trailing newline) to the `## Script` section body in `scripts/script-{nn}.md`.

### Skeleton

```markdown
---
marp: true
theme: gaia
---

<!-- _class: tpl-02 -->
<!-- _paginate: true -->

<style>
/* === from marps/template-02.md (required for standalone preview) === */
section.tpl-02 { … }
section.tpl-02 h1 { … }
/* === slide-only layout helpers (optional) === */
section.tpl-02 .s-02__cols { display: grid; grid-template-columns: 1fr 1fr; gap: 24px; }
</style>

# {Title from slides/slide-{nn}.md Content Elements}

<div class="tpl-02__body">
{Body — implement Visual Design zones and Content Elements exactly}
</div>

<div class="tpl-02__footer">
{Footer element if specified in slide}
</div>

<!--
{exact ## Script section body from scripts/script-{nn}.md — verbatim}
-->
```

### Mapping inputs → fragment

| Source | What goes into `marp-{nn}.md` |
|--------|-------------------------------|
| `marps/template-{nn}.md` | `<!-- _class: tpl-{nn} -->`, template CSS, region classes |
| `slides/slide-{nn}.md` | All on-slide text, layout mode, density, spatial placement |
| `scripts/script-{nn}.md` | HTML comment = exact `## Script` section body only |

### Content layout modes (from slide)

| Mode (from slide spec) | Marp implementation |
|------------------------|---------------------|
| `focal` | Large heading or highlight block |
| `bullets` | Markdown list; apply **density** styles if needed |
| `columns` | Grid — `grid-template-columns` for 2–3 cols (`s-{nn}__cols`) |
| `rows` / `stack` | Block elements stacked vertically |
| `grid` | CSS grid — e.g. `repeat(2, 1fr)` |
| `table` | Markdown or HTML table |
| `diagram` | HTML flex/grid + arrows; free spatial per Visual Design zones |
| `mixed` | Combine modes per spatial zones in page spec |

**Contrast axis:**

| Axis | Marp approach |
|------|---------------|
| `horizontal` | Multi-column grid |
| `vertical` | Stacked sections (top block / bottom block) |
| `focal+periphery` | Large center element + side or surrounding cells |
| `before→after` | Row or arrow-flow layout |

### Density (from slide)

| Density | Marp approach |
|---------|---------------|
| `normal` | Template default type sizes |
| `dense` | Slightly smaller `font-size`, tighter `line-height` and `gap` (slide-scoped) |
| `compact` | Smaller type, minimal padding — only when the slide requires many items |

Do not split slides at render time — if density is `dense` or `compact`, adjust typography in the fragment.

### Fidelity to `slides/slide-{nn}.md`

`slides/slide-{nn}.md` is the **canonical** on-slide source. `marps/marp-{nn}.md` must:

- Implement **every** row in **Content Elements** (text and position).
- Respect **Visual Design** zones and hierarchy (title / body / footer / callouts).
- Use **Content layout** and **Density** as specified.
- Use **Template** exactly as assigned (load that `marps/template-{nn}.md`).
- **Not invent** headlines, bullets, or figures absent from the slide.
- **Not** put script prose on the slide body — scripts are presenter notes only.
- When scripts exist, presenter notes **must** be the exact `## Script` section (see **Presenter notes from `## Script`**).
- Prefer Marp Markdown; add HTML only when layout requires it.

If the slide and template conflict, follow the slide’s Visual Design note; if still unclear, stop and ask — do not invent.

### Layout examples

**Horizontal columns (2-col contrast):**

```markdown
<style>
section.tpl-02 .s-11__cols { display: grid; grid-template-columns: 1fr 1fr; gap: 24px; }
</style>

# {Title from title zone}

<div class="tpl-02__body s-11__cols">
<div>…</div>
<div>…</div>
</div>
```

**Vertical stack (top / bottom):**

```markdown
# {Top section title}

- item
- item

> **"{Quote}"**
```

**Dense bullets:**

```markdown
<style>
section.tpl-02 { font-size: 0.9em; }
section.tpl-02 ul { line-height: 1.3; margin: 0.3em 0; }
</style>
```

### Cover slide (template-01)

```markdown
---
marp: true
theme: uncover
---

<!-- _class: tpl-01 -->
<!-- _paginate: false -->

<style>
/* from marps/template-01.md */
section.tpl-01 { … }
</style>

# {Title}
## *{Subtitle}*
```

## Frame Utilization (Full-Slide Layout)

Presentation slides must **use the full 16:9 frame**. Apply when writing both `marps/template-{nn}.md` and `marps/marp-{nn}.md`.

| Fail | Fix |
|------|-----|
| Tiny type clustered in one corner | Increase type scale; use template region padding; distribute content across title / body / footer |
| Large empty void with a small content island | Grow type, spacing, or zone sizes so primary content reads as the slide’s focus |
| Content overflows or clips past the slide edge | Reduce type, tighten gaps, switch to `dense`/`compact` only if the **slide** allows it; simplify layout helpers |
| Uneven margins (flush to one edge) | Honor template padding (e.g. ~8–10% horizontal); keep consistent inset |
| Title huge, body unreadable (or reverse) | Follow template typography scale and slide-specified sizes |

**Checks before marking a fragment done:**

1. Title zone, content zone, and footer (if any) are visibly occupied per the slide spec.
2. Primary message is readable at presentation distance (not micro-type).
3. No element is cut off in Marp preview at `size: 16:9`.
4. Whitespace is intentional (template “generous”) — not accidental underuse of the frame.

## Compilation — `{title-slug}.md`

1. Verify every slide in `slides.md` has approved `slides/slide-{nn}.md` and `marps/marp-{nn}.md`; every referenced template has `templates/template-{nn}.md` and `marps/template-{nn}.md`.
2. Set deck YAML from `overview.md` (title, author); default `theme` from primary template **Layout Defaults**.
3. Emit **once**, in template-number order, the CSS from each `marps/template-{nn}.md` inside a single top-level `<style>` block (merge-safe `tpl-{nn}` selectors only).
4. For each slide in `slides.md` order:
   - Extract body from `marps/marp-{nn}.md`: strip per-slide front matter (`---` block) and **strip duplicated template CSS** (already at deck top). Keep slide-only helpers (`s-{nn}__*`) if needed — either inline in that slide’s section or hoist carefully without clashing.
   - Keep `<!-- _class: tpl-{nn} -->` (and `_paginate` overrides).
   - Append presenter notes in `<!-- -->` if scripts exist: use the **exact `## Script` section body** from `scripts/script-{nn}.md` (same extraction rules as fragments). Prefer the notes already in `marps/marp-{nn}.md` when they match that section; if they differ, **fix-copy from the script file** (script file wins).
   - Separate slides with `---`.
5. Do not embed `overview.md`, `slides.md`, or section metadata in the deck body.
6. Do not edit slide copy or script notes at compile time — fragments are the source of slide bodies; `## Script` is the source of presenter notes; templates are the source of shared chrome CSS.

## Visual Quality (Marp)

- One focal message per slide.
- Implement spatial zones from the page spec — do not flatten to a bare bullet wall.
- Use content layout mode (columns, grid, stack) when items are parallel or structured.
- Apply **density** styles when the slide says `dense` or `compact`.
- Use the full frame (see **Frame utilization**).
- Preview every `marps/marp-{nn}.md` before compile.
- Preview the compiled `{title-slug}.md` after compile.

## Revision

| Change type | Update |
|-------------|--------|
| Layout / copy intent | `slides/slide-{nn}.md` first, then `marps/marp-{nn}.md`, then recompile |
| Marp rendering only (placement, type scale) | `marps/marp-{nn}.md`, then recompile |
| Spoken narrative (`## Script`) | `scripts/script-{nn}.md` first, then re-copy `## Script` into `marps/marp-{nn}.md`, then recompile |
| Shared visual foundation (PIM) | `templates/template-{nn}.md` → `marps/template-{nn}.md` → all affected `marps/marp-*.md` → recompile |
| Template Marp CSS only | `marps/template-{nn}.md` → affected `marps/marp-*.md` → recompile |

Never edit `{title-slug}.md` slide bodies or template CSS directly.

## Prohibitions

- Skipping `marps/template-{nn}.md` and inlining ad-hoc chrome only inside slides (templates must exist as mergeable PSM artifacts)
- Non-namespaced template CSS that breaks when multiple templates compile together
- Compiling without approved `marps/marp-*.md` for every slide
- Compiling directly from `slides/slide-*.md` without `marps/` implementation
- Inventing on-slide copy not present in `slides/slide-{nn}.md`
- Putting script prose in the slide body — presenter notes in HTML comments only
- Paraphrasing, summarizing, or inventing presenter notes when `scripts/script-{nn}.md` exists
- Omitting presenter notes when `scripts/script-{nn}.md` exists
- Embedding `## Key Phrases`, `## Notes`, or other non-`## Script` sections as presenter notes
- Tiny corner layouts or overflowing content (see **Frame utilization**)
- Skipping Marp preview review before compile

## Export

```bash
npx @marp-team/marp-cli {project-root}/{title-slug}.md --pdf
npx @marp-team/marp-cli {project-root}/{title-slug}.md --pptx
```

Built-in themes require no `--theme-set` for `gaia` and `uncover`. Template CSS is embedded in the compiled markdown.

## Completion

- Every used template has `marps/template-{nn}.md` with merge-safe `tpl-{nn}` CSS
- Every slide has `marps/marp-{nn}.md` built from **template + slide + script**
- On-slide content matches `slides/slide-{nn}.md` Content Elements and Visual Design
- Presenter notes are the exact `## Script` section body from `scripts/script-{nn}.md` when scripts exist (no paraphrase)
- Fragments and compiled deck use the full frame without overflow or micro-type corners
- `{title-slug}.md` compiles templates once, then slides in order
- User approved final `{title-slug}.md` or export
