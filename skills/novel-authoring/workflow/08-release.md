# Stage ⑧ — Release

**Prerequisites:** Manuscript evaluation passed (stage ⑦, main gate); user approves release

**Gate:** Continuity files updated; episode locked

**Next stage:** `04-episode-design.md` for next episode

---

## Procedure

Mark episode complete. **Released episodes should not be revised** unless user explicitly requests republication.

1. Confirm user explicitly approves release
2. Update `evaluations/{nnn}-{episode-slug}-manuscript.md` → Release section with date
   - **Release date = calendar date of this approval** (session “today”). **Forbidden:** copy `TODO.md` `Updated:` / sprint header / backlog dates into Release.
3. **Update continuity files** (critical for next episode). Prefer extracting from **approved design Key Events + manuscript outcomes**, not free paraphrase of long prose:

Open threads live in `continuity/story-so-far.md` → **Current Unresolved Threads**.

### `continuity/{nnn}-{episode-slug}-summary.md`

```markdown
# Episode {nnn}: {Title} — Summary

## Scenes Covered
| Sc | Title | Key Events |
|----|-------|------------|

## Key Events (Episode Level)
## Character State Changes
## Location State Changes
## New Unresolved Threads
## Resolved Threads
## Narrative Hooks for Next Episode
```

### `continuity/story-so-far.md` (upsert)

**SSOT for open threads** — maintain **Current Unresolved Threads** here only (status + window + notes). When updating, anchor on the last existing row or section head and append the new row (not a full-file overwrite). Do not invent the original text that is only the new row.

```markdown
# Story So Far — {Title}

## Active Characters
| Character | Current State | Last Episode | Key Changes |

## Active Locations
| Location | Current State | Last Changed In | Notes |

## Major Events (Timeline)
| Episode | Sc | Story time (if known) | Key Event |

## Current Unresolved Threads
| Thread | Set Up In | Status | Expected Window | Notes |
|--------|-----------|--------|-----------------|-------|
| {id or short label} | Ep {nnn} | open / advanced / held / resolved | {window or —} | {carry note} |
```

On release: add/update/resolve rows in this table from the episode’s New/Resolved Threads + manuscript outcomes.

4. Provide **continuity briefing** to user (states, open threads, next-episode obligation from Out hook)

5. Proceed to stage ④ for next episode — load `04-episode-design.md`; then **Selective artifact load**: Phase A indexes → Phase B Appearing/used details + continuity (`story-so-far` + **this** episode’s new summary as the prior for N+1). Do not reload every prior episode design or the full catalog.

---

## Series-Level Review (Optional)

After multiple episodes or a complete arc:

- `evaluations/arc-{arc-slug}.md`
- `evaluations/series.md`

---

## Gate G8 (Release Approval)

Episode locked. Continuity synced. Ready for next episode design.
