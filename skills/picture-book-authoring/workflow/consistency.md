# Consistency — Cross-Stage Protocol (Picture Book)

**Not a numbered stage.** Read when starting Stage ②/③/④, when revising (within the stage’s loops), or when any approved artifact changes.

This file is the **single source of truth** for keeping:
- story / direction consistency
- character / location / staging / world-rule consistency
- image identity consistency (reference-based)

---

## Glossary (canonical terms)

| Term | Meaning |
|---|---|
| **Design Layer** | `world-bible.md`, `characters/*.md`, `locations/*.md`, `stagings/*.md`, `series.md`, `episodes/*.md` |
| **Canonical episode page schema** | Flat `- **Field:**` lists under `### Page {N}` only — no nested `####` page subsections; see `workflow/02-design.md` |
| **Evaluation Layer** | `evaluations/*.md` where **story lock** is approved per episode |
| **Reference Images** | `images/characters/*`, `images/locations/*`, `images/stagings/*` — identity + set + ensemble blocking |
| **Staging** | Continuing-situation who-is-where — `workflow/reference-models.md` |
| **Page Images** | `images/{nnn}-{episode-slug}/*` — rendered pages including locked `rendering text` overlay |
| **Story Lock** | G3 approval: page stories / rendering text / text–image split / page-turn hooks / illustration guides are frozen for image generation |

---

## Source-of-truth stack (top wins until updated)

```
overview.md
  → series.md
    → world-bible.md + characters/* + locations/* + stagings/*
      → episodes/*.md (page story + rendering text + illustration guide)
        → evaluations/*.md (story lock decision)
          → images/* (reference + page PNGs)
```

---

## Cascade protocol (what to do when something changes)

### When design catalogs change (characters/locations/stagings/world-bible)

1. User approves the higher-layer change first (Design layer).
2. Identify affected episodes (any page referencing changed cast/state/place/view).
3. Re-run Stage ③ for affected episodes and seek a new **G3 story lock** (if page story/text/guides must be adjusted).
4. Regenerate:
   - reference images (if identity/appearance/state changed)
   - page images for pages whose locked guides reference updated references

### When episode page design changes

1. Design change requires rollback to Stage ② for that episode.
2. Re-run Stage ③ for the episode and approve a new G3 story lock.
3. Regenerate only the affected episode’s page images (and reference images only if the episode requires new cast/place/state variants).
4. After design is final, sync `series.md` page estimates to **measured** `### Page` counts (or document an exception in Craft Notes).

### When only visual inconsistency is found during generation

- Treat it as a visual quality issue:
  - tighten prompts / re-render affected page image(s)
  - max 2 retries per image
- Do NOT alter locked rendering text or locked illustration guide meaning.

### When inconsistencies reveal a design gap

- Roll back to Stage ② (Design)
- Update catalogs/episode page design
- Re-run Stage ③ and then Stage ④ for affected assets

---

## Cascade checklist (copy/paste)

```markdown
## Consistency Cascade
- Change: {what changed}
- Highest updated file(s): {paths} — user approved: ✅/⬜
- [ ] overview.md (only if scale/age/validation changed)
- [ ] series.md
- [ ] world-bible.md / characters/* / locations/* / stagings/*
- [ ] episodes/{nnn}-{episode-slug}.md
- [ ] evaluations/{nnn}-{episode-slug}.md (story lock rerun)
- [ ] reference images to regenerate (if identity/state/view changed)
- [ ] page images to regenerate for affected episode pages
- Gates re-run: Story Lock (G3) for affected episodes
```

---

## Session resume

When resuming mid-work:
1. State the current stage and episode number(s).
2. Re-load disk artifacts for the active stage (do not rely only on chat summaries).
3. Confirm whether G3 story lock already exists for those episodes.
4. Follow that stage workflow file again (`workflow/0X-*.md`).
