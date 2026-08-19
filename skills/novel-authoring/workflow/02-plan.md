# Stage ② — Plan (`series.md`)

**Prerequisites:** Approved `overview.md`

**Gate artifact:** `{project-root}/series.md`

**Next stage:** `03-architecture.md` (after user approves `series.md`)

---

## Purpose

Design the **series-level blueprint** — what each episode delivers and how episodes chain. This stage plans **story rhythm**, not cast/place detail and not scene structure.

Architecture (world, characters, locations) follows in stage ③ as a **series-level** foundation informed by this Episode List — episode design then works inside that context.

---

## Procedure

1. Load approved `overview.md` — genre, scale, approximate episode count, validation criteria. load only if not already in this turn’s context (not chat memory alone when missing).
2. Create or update `series.md`:

```markdown
# Episode List: {Title}

## Structure
{story arc — emotional peaks, where tension builds across episodes; optional arc grouping}

## Episode List
| # | Title | Summary | Hook to Next |
|---|-------|---------|--------------|
| 001 | {Title} | {what this episode delivers — situation, turn, hook to next} | {tension / unanswered beat → Ep 002} |
| 002 | {Title} | ... | ... |
| {last} | {Title} | ... | {closing beat} |
```

### Rules

- **`series.md` owns only:** episode count, titles, summaries, and **Hook to Next** (required on every row)
- Episode count may differ from the approximate count in `overview.md` — update `overview.md` Scale when it does
- **Do not** put scene counts, Scene Index, word budgets, character lists, or location lists in `series.md`
- Scene volume is decided per episode at stage ④
- **Wording:** Use real plot vocabulary in Structure / Summary / Hook (e.g. reveal, expose, confront). Do **not** invent nonsense tokens or typos as plot terms — they propagate into later stages.
- **Seed timing:** Keep early-episode Summaries from stealing later payoffs. Signature beats belong in the episode that owns them; foreshadow with Hook / Hint language, not full payoff wording in ep 001 when the List schedules the payoff later.

### Review focus

- Is the episode count right for the scale in `overview.md`?
- Does each title + summary earn a release slot?
- Does the arc chain work? Read **Hook to Next** top-to-bottom — consecutive episodes should connect; gaps or dead ends need fixing here
- Do Summaries use clear verbs (no corrupted/nonsense plot words)?
- Does payoff timing match later rows (no early full reveal of a later episode’s climax)?

---

## Completeness Check

- [ ] Episode List complete with Summary + Hook to Next on every row
- [ ] Consecutive hooks connect
- [ ] No scene counts or per-episode scene planning
- [ ] Structure section gives arc context
- [ ] Episode count aligned with `overview.md` Scale (or Scale updated with user approval)
- [ ] Summaries use clear plot vocabulary (no nonsense/typo tokens as plot terms)
- [ ] Early-row Summaries do not fully execute later-row payoffs

---

## Gate G2 (Plan Approval)

User approves `series.md`. Do not proceed to stage ③ without approval.

When Episode List rows change later → cascade per SKILL.md § Consistency to open episode designs.
