# Stage ⑧ — Release

**Prerequisites:** Manuscript evaluation passed; user approves release

**Gate:** Continuity files updated; chapter locked

**Next stage:** `04-chapter-design.md` for next chapter

---

## Procedure

Mark chapter complete. **Released chapters should not be revised** unless user explicitly requests republication.

1. Confirm user explicitly approves release
2. Update `evaluations/{nnn}-{chapter-slug}.md` → Release section with date
3. **Update continuity files** (critical for next chapter design/generation):

### `continuity/{nnn}-{chapter-slug}-summary.md`

```markdown
# Chapter {nnn}: {Title} — Summary

## Episodes Covered
| Ep | Title | Key Events |
|----|-------|------------|

## Key Events (Chapter Level)
## Character State Changes
## Location State Changes
## New Unresolved Threads
## Resolved Threads
## Narrative Hooks for Next Chapter
```

### `continuity/story-so-far.md` (upsert)

```markdown
# Story So Far — {Title}

## Active Characters
| Character | Current State | Last Chapter | Key Changes |

## Active Locations
| Location | Current State | Last Changed In | Notes |

## Major Events (Timeline)
| Chapter | Ep | Key Event |

## Current Unresolved Threads
| Thread | Set Up In | Notes |
```

### `continuity/unresolved-threads.md` (upsert)

```markdown
# Unresolved Plot Threads

| Thread | Chapter Set Up | Current Status | Expected Resolution Window |
```

4. Provide **continuity briefing** to user:

   > *"Chapter 003 released. Character X is now in Y. Two hooks remain: … Next chapter should …"*

5. Proceed to stage ④ for next chapter — **read `04-chapter-design.md` and load continuity files**

---

## Series-Level Review (Optional)

After multiple chapters or a complete part:

- `evaluations/part-{nnn}-{part-slug}.md`
- `evaluations/series.md`

---

## Gate

Chapter locked. Continuity synced. Ready for next chapter design.
