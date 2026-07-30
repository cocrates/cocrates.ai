# Consistency — Cross-Stage Protocol (Webtoon)

**Not a numbered stage.** Read when starting Stage ②/③/④, when revising (within the stage’s loops), or when any approved artifact changes.

This file is the **single source of truth** for keeping:
- story / direction consistency (page → cut → balloon)
- character / location / staging / world-rule consistency
- image identity consistency (reference-based)
- episode scroll continuity (stitched pages)

Canonical reference rules: `workflow/reference-models.md`.

---

## Glossary (canonical terms)

| Term | Meaning |
|---|---|
| **Design Layer** | `world-bible.md`, `characters/*.md`, `locations/*.md`, `stagings/*.md`, `series.md`, `episodes/*.md` |
| **Evaluation Layer** | `evaluations/*.md` where **story lock** is approved per episode |
| **Reference Images** | `images/characters/*`, `images/locations/*`, `images/stagings/*` — identity + set + ensemble blocking assets |
| **Character ref** | Lasting body + outfit + **identity gear** (weapons/shields/accessories); not expression/mood |
| **Location ref** | **Set/stage** physical structure; not time/season/weather |
| **Staging ref** | **Who is where** for one continuing situation (café L/R, OR stations, meeting, …); typically 2–3 ensemble views |
| **Page Images** | `images/{nnn}-{episode-slug}/{page}.png` — variable-height strip segments with cuts + balloons + captions |
| **Episode Scroll** | `images/{nnn}-{episode-slug}/episode-scroll.png` — pages concatenated top→bottom |
| **Cut** | Panel within a page; art + optional speech balloons + optional narration; has size class + height guide |
| **Gutter** | Vertical spacing between cuts (`tight` / `normal` / `wide` / `pause`) |
| **Size class** | `standard` / `tall`/`long` / `compact` / `open`/`diagonal` — see `workflow/cut-composition.md` |
| **Story Lock** | G3 approval: page/cut stories, balloon texts, captions, composition, staging cites, illustration guides frozen |

---

## Source-of-truth stack (top wins until updated)

```
overview.md (incl. canvas width / color / outside-cut fill)
  → series.md
    → world-bible.md + characters/* + locations/* + stagings/*
      → episodes/*.md (page → cut + balloons + captions + gutters + staging cites)
        → evaluations/*.md (story lock decision)
          → images/* (character + location + staging refs, then page PNGs)
            → episode-scroll.png (+ optional portal slices)
```

---

## Cascade protocol (what to do when something changes)

### When design catalogs change (characters/locations/stagings/world-bible)

1. User approves the higher-layer change first (Design layer).
2. Identify affected episodes (any cut referencing changed cast/state/place/view/staging).
3. Re-run Stage ③ for affected episodes and seek a new **G3 story lock** (if cut story/text/guides must be adjusted).
4. Regenerate:
   - reference images (if identity/appearance/state/blocking changed)
   - page images for pages whose locked guides reference updated references
   - re-stitch episode scroll

### When episode page/cut design changes

1. Design change requires rollback to Stage ② for that episode.
2. Re-run Stage ③ for the episode and approve a new G3 story lock.
3. Regenerate only the affected episode’s page images (and reference images only if new cast/place/state/staging variants are required), then re-stitch.

### When only visual inconsistency is found during generation

- Treat it as a visual quality issue:
  - tighten prompts / re-render affected page image(s)
  - max 2 retries per image
  - re-stitch if needed
- Do NOT alter locked balloon/caption text, locked cut guide meaning, identity gear, or staging seats.

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
- [ ] series.md
- [ ] world-bible.md / characters/* / locations/* / stagings/*
- [ ] episodes/{nnn}-{episode-slug}.md
- [ ] evaluations/{nnn}-{episode-slug}.md (story lock rerun)
- [ ] reference images to regenerate (identity / set / staging)
- [ ] page images to regenerate for affected episode pages
- [ ] episode-scroll.png re-stitch (+ portal slices if any)
- Gates re-run: Story Lock (G3) for affected episodes
```

---

## Session resume

When resuming mid-work:
1. State the current stage and episode number(s).
2. Re-load disk artifacts for the active stage (do not rely only on chat summaries).
3. Confirm whether G3 story lock already exists for those episodes.
4. Follow that stage workflow file again (`workflow/0X-*.md`).
