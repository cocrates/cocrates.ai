# Stage ⑦ — Revision

**Prerequisites:** Evaluation completed (stage ⑥); user selected findings to apply

**Gate:** User approves revised artifact(s); Consistency Cascade complete if higher layers changed

**Next stage:** Re-evaluate (`06`) if needed → `05` (after design revision) or `08` (manuscript clean)

**Also read:** [`consistency.md`](consistency.md) · chapter scene rules in [`04-chapter-design.md`](04-chapter-design.md)

---

## Core Rule

1. **Design before prose** for anything that changes plot, pacing, seeds, cast, place, or world
2. **Higher before lower** — if the finding requires a Series/Part/architecture change, update that file first (user approves), then cascade per `consistency.md`
3. Tell the user which files will change **before** editing, then confirm cascade status when done

Prose-only polish is allowed **only** when Key Events, Seeds/Hold, Prior Design Alignment, and continuity remain true.

---

## Classify the finding

| Kind | Examples | First file to touch |
|------|----------|---------------------|
| A — Prose craft only | Rhythm, wording, minor clarity | `manuscripts/` only |
| B — Chapter design | Beats, hooks, seeds, motifs, scenes | `chapters/{nnn}-*.md` then manuscript if it exists |
| C — Higher plot/list | Wrong Role, Hook to Next, part arc | Chapter List source (`series.md` or `parts/`) → catalog/design → manuscript |
| D — Lore / cast / place | New rule, renamed place, personality change | `world*` / `characters*` / `locations*` → open designs → manuscripts as needed |
| E — Continuity error | Contradicts released ch | Prefer fix unreleased design/manuscript; if released must change → user must approve republication path |

---

## Procedure

1. User selects findings from evaluation
2. Classify A–E; propose cascade file list; get approval if C/D/E
3. Edit highest layer first → cascade downward (use checklist in `consistency.md`)
4. Re-run gates:
   - Design changed → Prior-Design Consistency + Manuscript Readiness (`04`)
   - Manuscript changed → Design-Fidelity (`05`)
5. Record in `evaluations/{nnn}-{chapter-slug}.md` → Revision Decisions (note cascade files)
6. Present for approval: design first if design changed; then manuscript

### When updating chapter design

- Keep Episode Index in sync with `## Episodes` sections
- Fill all required scene fields (see `04`)
- Refresh **Prior Design Alignment** if higher files changed
- Do not add Episode Detail summaries that duplicate Key Events

### When regenerating manuscript

- Reload design + Architecture/Continuity References (`05` Pre-Generation Load)
- Pass Design-Fidelity Gate before asking approval

---

## Common Critique → Design Fixes

| Critique finding | Design fix | Manuscript fix |
|------------------|------------|----------------|
| Too much world-building in opening | Move lore to Hold; lower budget | Scene opening; cut essay |
| Info-dump dialogue | Key Events: intent only; cut lecture beats | Compress dialogue |
| Too many mysteries explained | Reclassify as Hold; one Hint | Cut explanatory talk |
| No personal connection | Add Personal Stake | Character-specific detail |
| Thematic summary ending | Closing image in Hold/craft | End on image/silence |
| Covers entire series themes | Redistribute via Chapter List / seeds | Cut held content |
| Opening action without tension | Add Opening Question | Persist unanswered question |
| Flat rhythm | Note in Literary Craft | Vary sentences |
| World details functional only | Sensory-emotional per scene | Pair detail → POV reaction |
| Same voice for all | Dialogue Voices table | Differentiate + inner/outer gap |
| Motif once | Plan placements across scenes | Thread motif |
| Too many POV inserts | ≤ 2 per episode in design | Cut extras |
| Info-heavy opening | Lower exposition budget | Cut catalog blocks |
| Cartoon villain | Self-justifying worldview in Voices | Rewrite as believer |
| POV emotion over-labeled | Indeterminacy in design | Ambiguous feeling cues |
| Jarring episode joins | Fix episode Hooks + Transition out | Smooth handoffs |
| Chapter vs Chapter List Role | Cascade C: List → design | Restructure to match |
| Prose invents beat/cast | Do not “keep” invention — design first or cut | Align to Key Events |
| Conflicts with world/continuity | Cascade D/E | Regenerate after sync |

---

## Gate

User approves revised artifact(s). Design approval precedes manuscript approval when both changed. Cascade checklist must show no dangling higher/lower mismatch.
