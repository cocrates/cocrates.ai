# Consistency — Cross-Stage Protocol

**Not a numbered stage.** Read when starting ④ or ⑤, when revising (⑦), or when any approved artifact changes.

This file is the **single source of truth** for how plot/world/cast layers stay aligned. Stage workflows link here instead of redefining rules.

---

## Glossary (canonical terms)

| Term | Meaning |
|------|---------|
| **Structure Mode** | `short` or `series` — locked in `overview.md` / `series.md` |
| **Chapter List source** | **Short:** `series.md` Chapter List. **Series:** `parts/{nnn}-*.md` Chapter List |
| **Chapter catalog** | Stage ③ thin file: Role, conflict, hooks — no episodes |
| **Chapter design** | Stage ④ full file: same path, plus Alignment, Episode Index, episode/scene Key Events |
| **Episode Index** | TOC only inside chapter design (not "Episode List" elsewhere) |
| **Prior-Design Consistency Gate** | Stage ④ — chapter design vs higher artifacts (`04`) |
| **Manuscript Readiness** | Stage ④ — scene fields complete for generation (`04`) |
| **Canonical scene schema** | Meta `**Field:**` lines + flat `- **Field:**` bullets under `##### Scene` only — see `04-chapter-design.md` |
| **Design-Fidelity Gate** | Stage ⑤ — prose vs chapter design + architecture + continuity (`05`) |
| **Continuity** | Released-plot truth (`continuity/`) — not prior manuscript text |
| **Reference models** | Character / location / staging locks — [`reference-models.md`](reference-models.md) (prose; shared with webtoon/picture-book/video) |
| **Staging** | Who-is-where for a continuing situation across scenes |

---

## Source-of-truth stack (top wins until updated)

Higher layers constrain lower layers. Changing a lower layer **cannot** silently contradict a higher one.

```
overview.md (mode, tone, validation, scale estimates)
  → series.md (arc, voice; short: Chapter List / series: Part Catalog)
    → parts/*.md                    [series mode only]
      → chapters/{nnn}-*.md catalog (③) → full design (④)
        → manuscripts/{nnn}-*.md (⑤)
continuity/*                        [after ⑧; binds ④/⑤ for ch 002+]
characters* / locations* / stagings* / world*   [architecture; bind all ④/⑤]
```

| If you need to change… | Update first (user approve) | Then sync |
|--------------------------|----------------------------|-----------|
| Structure Mode, scale, validation | `overview.md` | `series.md`, Chapter List source |
| Series arc, voice, short Chapter List | `series.md` | chapter catalogs/designs that conflict |
| Part role / chapter rows (series) | `parts/` (+ Part Catalog counts) | chapter catalogs/designs |
| Character / location / world rule | architecture profile(s) | open chapter designs + future designs; manuscripts only after design update |
| Chapter Role / Hook to Next | Chapter List source, then chapter catalog/design Closing hook | Prior Design Alignment; regenerating manuscript if needed |
| Scene beats / seeds | chapter design (④) | manuscript (⑤) — never prose-only for plot |
| Released past events | **Don't** — released chapters are locked unless republication; next chapters use continuity | If user forces republication: revise design → manuscript → re-release → rewrite continuity |

---

## Agent / user workflow when anything changes

1. **Name the change** and which layer it belongs to (table above)
2. **Propose the highest-file edit first** — one approval at a time when layers differ
3. **Cascade downward** using the checklist below — mark each file touched
4. **Re-run the relevant gate** (Consistency / Readiness / Fidelity) before the next stage
5. Tell the user in one short status line, e.g. *"Updated series Chapter List Role for ch 004 → aligned chapter design Closing hook → manuscript not yet regenerated."*

### Cascade checklist (copy into chat or evaluation notes)

```markdown
## Consistency Cascade
- Change: {what}
- Highest file updated: {path} — user approved: ✅/⬜
- [ ] overview.md (if mode/scale/tone/validation)
- [ ] series.md
- [ ] parts/ (series mode)
- [ ] chapters/{nnn} catalog and/or full design
- [ ] characters* / locations* / world* (if cast/space/lore)
- [ ] manuscripts/{nnn} (only after design sync)
- [ ] continuity/* (only on ⑧ or approved republication)
- Gates re-run: Prior-Design Consistency / Manuscript Readiness / Design-Fidelity — {which}
```

---

## Session resume

When resuming mid-novel:

1. State current stage and chapter number
2. Read this file + the stage workflow
3. Confirm Structure Mode and Chapter List source path
4. For ④/⑤/⑥/⑦: reload disk artifacts — do not trust chat summary of lore or beats

---

## What not to duplicate

- Full Prior-Design / Design-Fidelity **tables** live only in `04` / `05`
- Craft anti-patterns live primarily in `04` (design) and `05` (prose rules)
- Structure Mode mechanics live in `SKILL.md` summary + `01`/`02`/`03` — do not invent a third mode
