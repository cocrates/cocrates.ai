# Consistency — Cross-Stage Protocol (Webtoon)

**Not a numbered stage.** Read when starting Stage ②/③/④, when revising (within the stage’s loops), or when any approved artifact changes.

This file is the **single source of truth** for keeping:
- story / direction consistency (page → cut → balloon)
- character / location / staging / world-rule consistency
- **series art / typography consistency** (`illustration-guide.md`)
- image identity consistency (reference-based)
- episode scroll continuity (stitched pages)

Canonical reference rules: `workflow/reference-models.md`.
Series visual/typography rules: `workflow/illustration-guide.md`.

---

## Glossary (canonical terms)

| Term | Meaning |
|---|---|
| **Design Layer** | `illustration-guide.md`, `world-bible.md`, `characters/*.md`, `locations/*.md`, `stagings/*.md`, `series.md`, `episodes/*.md` |
| **Canonical episode cut schema** | Flat `- **Field:**` lists under `##### Cut {n}` only — no nested `######` cut subsections; see `workflow/02-design.md` |
| **Series illustration guide** | `{project-root}/illustration-guide.md` — art style, chrome, balloon/caption typography; see `workflow/illustration-guide.md`. Stage ④ YAML must apply it. |
| **Evaluation Layer** | `evaluations/*.md` where **story lock** is approved per episode |
| **Reference Images** | `images/characters/*`, `images/locations/*`, `images/stagings/*` — identity + set + ensemble blocking assets |
| **Character ref** | Lasting body + outfit + **identity gear** (weapons/shields/accessories); not expression/mood |
| **Location ref** | **Set/stage** physical structure; not time/season/weather |
| **Staging ref** | WHO–WHERE for one continuing situation; **default 1× establishing** (optional reverse/detail only) |
| **Page Images** | `images/{nnn}-{episode-slug}/{page}.png` — stitched page strip; tiles `{page}-a.png`, `{page}-b.png`, … on the default Flash path |
| **Episode Scroll** | `images/{nnn}-{episode-slug}/episode-scroll.png` — pages concatenated top→bottom |
| **Cut** | Panel within a page; art + optional speech balloons + optional narration; has size class + height guide |
| **Gutter** | Vertical spacing between cuts (`tight` / `normal` / `wide` / `pause`) |
| **Size class** | `standard` / `tall`/`long` / `compact` / `open`/`diagonal` — see `workflow/cut-composition.md` |
| **Story Lock** | G3 approval: page/cut stories, balloon texts, captions, composition, staging cites, Direction fields frozen; series visual+typography lock stays in `illustration-guide.md` |

---

## Source-of-truth stack (top wins until updated)

```
overview.md (incl. canvas width / color / outside-cut fill)
  → illustration-guide.md (art / chrome / balloon·caption typography)
  → series.md
    → world-bible.md + characters/* + locations/* + stagings/*
      → episodes/*.md (page → cut + balloons + captions + gutters + staging cites)
        → evaluations/*.md (story lock decision)
          → images/* (refs + page PNGs — YAML must apply illustration-guide.md)
            → episode-scroll.png (+ optional portal slices)
```

---

## Cascade protocol (what to do when something changes)

### When `illustration-guide.md` changes

1. User approves the updated series guide (Design / G2).
2. Re-run Stage ③ if Evaluate must re-check lettering/chrome fitness.
3. Regenerate **all** page tiles that must match the new recipes (typically the whole episode) — do not leave old pages on the previous balloon font/box style. Mark `episode-scroll.png` stale if present; re-stitch **only on explicit user request** after page images are complete.
4. Reference images: regenerate only if §1 art style (medium/line/palette/proportions) changed.

### When design catalogs change (characters/locations/stagings/world-bible)

1. User approves the higher-layer change first (Design layer).
2. Identify affected episodes (any cut referencing changed cast/state/place/view/staging).
3. Re-run Stage ③ for affected episodes and seek a new **G3 story lock** (if cut story/text/guides must be adjusted).
4. Regenerate:
   - reference images (if identity/appearance/state/blocking changed)
   - page images for pages whose locked guides reference updated references
   - do **not** auto re-stitch; re-run episode-scroll stitch only on explicit user request after pages are done

### When episode page/cut design changes

1. Design change requires rollback to Stage ② for that episode.
2. Re-run Stage ③ for the episode and approve a new G3 story lock.
3. Regenerate only the affected episode’s page images (and reference images only if new cast/place/state/staging variants are required). Mark episode scroll stale; re-stitch only on explicit user request.
4. After design is final, record measured page/cut counts in the episode Craft Notes only — do **not** push page/cut budgets into `series.md`.

### When only visual inconsistency is found during generation

- Treat it as a visual quality issue:
  - tighten prompts / re-render affected page image(s)
  - max 2 retries per image
  - do **not** auto re-stitch; mark episode scroll stale if present; re-stitch only on explicit user request after pages are complete
- Do NOT alter locked balloon/caption text, locked cut Direction meaning, identity gear, or staging seats.
- Do NOT invent lettering/chrome outside `illustration-guide.md`; if typography is wrong, update the guide (or re-apply it in YAML), not one-off styles.

### When inconsistencies reveal a design gap

- Roll back to Stage ② (Design)
- Update catalogs/episode page/cut design (including stagings if seating drifted)
- Re-run Stage ③ and then Stage ④ for affected assets

### When canvas width / portal export rules change

1. Update `overview.md` and get user approval.
2. Regenerate or resize/export only as needed — do not rewrite locked dialogue.
3. Re-export portal slices if those were produced.

---

## Cascade checklist (copy/paste)

```markdown
## Consistency Cascade
- Change: {what changed}
- Highest updated file(s): {paths} — user approved: ✅/⬜
- [ ] overview.md (scale / audience / validation / canvas)
- [ ] illustration-guide.md (art / chrome / balloon·caption typography)
- [ ] series.md
- [ ] world-bible.md / characters/* / locations/* / stagings/*
- [ ] episodes/{nnn}-{episode-slug}.md
- [ ] evaluations/{nnn}-{episode-slug}.md (story lock rerun)
- [ ] reference images to regenerate (if identity/state/view/staging **or art-style §1** changed)
- [ ] page images to regenerate for affected episode pages (all pages if typography recipes changed)
- [ ] episode-scroll.png re-stitch (+ portal slices if any) — **only on explicit user request** after page images are complete
- Gates re-run: Story Lock (G3) for affected episodes
```

---

## Session resume

When resuming mid-work:
1. State the current stage and episode number(s).
2. Re-load disk artifacts for the active stage (do not rely only on chat summaries).
3. Confirm whether G3 story lock already exists for those episodes.
4. Follow that stage workflow file again (`workflow/0X-*.md`).
