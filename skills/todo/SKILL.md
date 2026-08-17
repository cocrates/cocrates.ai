---
name: todo
description: >-
  Manages deliverable-generation tasks in TODO.md — initialize, track progress,
  sync with workspace artifacts, and pick the next actionable task. For large or
  iterative work (series, spec-driven implementation), also plans Product Backlog
  and sprint-NN.md under {project-root}/scrum/; TODO.md stays the progress SSOT
  and is re-initialized from the next sprint plan when a sprint ends.
  Use when starting or resuming a deliverable session, progress/next-step questions,
  sprint/backlog planning, or alongside any deliverable-generation workflow.
metadata:
  agent: cocrates
---

# TODO — Deliverable Task Tracking Skill

`TODO.md` is the **progress board**. Match the user's language in prose and status reports. Do not substitute the agent's ephemeral todo tool.

**Two concerns in one skill (do not split into another skill):**

| Concern | Where |
|---------|--------|
| **Progress** | `{project-root}/TODO.md` — Snapshot, tasks, Current sprint, Completed sprints |
| **Plans** | `{project-root}/scrum/product-backlog.md` + `sprint-NN.md` — written when planning; not a live status board |

**Churn rule:** Prefer changing **only `TODO.md`** as work advances. Touch `scrum/*` when **plans** change (new sprint file, backlog reshape) — not on every task checkbox. **No** `scrum/README.md`.

## When to Use

| Situation | Action |
|-----------|--------|
| New deliverable | **Initialize** (plain or scrum-mode planning) |
| Resume / mid-session | **Orient** (read `TODO.md` first) |
| Draft done or gate approved | **Execute** → **Sync** → recommend next |
| "where are we?" / "what's next?" | Report from `TODO.md` |
| Scope change | **Sync** within current sprint board |
| Backlog / sprint **plan** needed | **Plan** / **Plan Next Sprint** (below) |
| Sprint Goal met / all sprint tasks closed | **Sprint boundary** |

## Working Location

```text
{project-root}/
├── TODO.md                 # progress SSOT
└── scrum/                  # plans only (scrum mode)
    ├── product-backlog.md
    ├── sprint-00.md        # series skeleton when applicable
    └── sprint-NN.md
```

`{project-root}` from the companion skill and Workspace Convention (`read_agents: principles/02-workspace-conventions.md`).

---

## Scrum mode — when to use sprint plans

| Trigger | Examples |
|---------|----------|
| **Series deliverable** | novel, webtoon, picture-book, … — **default on** |
| **Spec-driven implementation** | Spec readiness / Spec stage done — iterative implementation |
| **User asks** | sprint, backlog, scrum, "나눠서" |
| **Scale heuristic** | Would exceed ~30 tasks, or unbounded episode loop without sprint boundaries |
| **Next sprint plan missing** | At Sprint boundary, no next `sprint-NN.md` |

**Plain mode** (no `scrum/`): small one-shot work unless the user asks for sprints.

### Profiles (planning)

**Series** — default on:

| Sprint plan | Scope |
|-------------|--------|
| `sprint-00` | ① define → ② plan → ③ architecture |
| `sprint-01`… | One publication unit release (episode / equivalent) |

Backlog references Episode List rows; story SSOT stays in `series.md` (etc.). Prefer pre-writing upcoming `sprint-NN.md` when the list is known.

**Spec-driven implementation** — after Specs are sufficient:

| Phase | Board |
|-------|--------|
| PRD · ASR · ADR · Spec writing | Plain todo / design — not implementation sprints |
| Implementation | `scrum/` plans + `TODO.md` execution |

Epic/Story from Spec; each sprint = vertical slice. Live progress = `TODO.md` only.

---

## Task Model

| Signal | Meaning |
|--------|---------|
| `- [ ]` | Open |
| `- [x]` | Closed (`done`, or skipped with `_(skipped)_`) |
| Snapshot **Current focus** | Single in-progress open task (at most one) |
| `_(blocked: …)_` | Open but blocked |
| `_(skipped)_` | Closed without doing the work |

Checkbox + Current focus only — no `` `pending` `` / `` `in_progress` `` / `` `done` `` tokens on headlines.

**IDs:** `T-{nnn}` stable **within a board**; never renumber/reuse on that board. Append under `## Tasks`. Across sprint rewrites, IDs may restart at `T-001` (plans stay in `scrum/sprint-NN.md`).

```text
- [ ] **T-00n** — {short title}
  - Phase: {phase or "cross-cutting"}
  - Artifact: `{path}` or —
  - Depends: T-00m, T-00k
  - Notes: {optional}
```

**Types:** draft · gate · unit · research · sync

---

## TODO.md Structure

### Plain mode

```markdown
# TODO: {Title}

> **Project type:** `{type}`
> **Project root:** `{project-root}`
> **Updated:** {YYYY-MM-DD}

## Snapshot
…

## Tasks
…

## Notes
…
```

### Scrum mode

```markdown
# TODO: {Title}

> **Project type:** `{type}`
> **Project root:** `{project-root}`
> **Updated:** {YYYY-MM-DD}
> **Current sprint:** sprint-{NN} — [{goal}](./scrum/sprint-{NN}.md)
> **Product Backlog:** [`scrum/product-backlog.md`](./scrum/product-backlog.md)

## Snapshot

| Done | Open | Blocked | Skipped |
|------|------|---------|---------|
| {n}  | {n}  | {n}     | {n}     |

**Sprint goal:** {one line}
**Current focus:** {T-id or "—"}
**Recommended next:** {T-id} — {one-line description}

## Tasks

- [ ] **T-001** — {title}
  - Phase: …
  - Artifact: …
  - Depends: —
  - Notes: …

## Completed sprints

| Sprint | Goal | Closed |
|--------|------|--------|
| 00 | {one-liner} | {YYYY-MM-DD} |

## Notes

{sprint-local — deferrals, carry-over}
```

**Hard rules:** Compact **Completed sprints** only (no full closed task dumps). No `scrum/README.md`. Execution-time edits ≈ **`TODO.md` only**.

**Identity:** `type` after Project type; title after `# TODO: `; workspace = Project root. **`session.status` (principles):** scrum 모드 → header **Current sprint**의 `sprint-NN`을 **`S-NN`**으로 투영 (예: `sprint-07` → `S-07`); plain 모드 → Snapshot **Current focus** T-id. Keep Current sprint / Current focus accurate so the session title stays correct.

**Edit discipline:** Short unique `edit_file` anchors; typical change = one headline + Snapshot lines. Append new tasks at end of `## Tasks`. **Exception:** intentional full-board rewrite at Sprint boundary.

**Inserting new tasks (avoid `old_string_not_found`):** `edit_file` cannot "insert beside" — it **replaces** an existing span. For a new task, `old_string` = the **last existing task block** (or the `## Tasks` heading + first task if appending after an empty section is wrong — use the last real task). `new_string` = that same block **plus** the new task lines. Never set `old_string` to only the new `- [ ] **T-0nn**` text (it is not on disk yet → `old_string_not_found`).

```text
# WRONG — new content as old_string
old_string: - [ ] **T-015** `pending` — Design ep 007
new_string: - [ ] **T-015** `pending` — Design ep 007
  - Phase: …

# RIGHT — expand from last existing task
old_string:
- [x] **T-014** `done` — Release ep 006
  - Phase: sprint-06
  - Artifact: `continuity/story-so-far.md`
  - Depends: T-013 ✓
  - Notes: released

new_string:
- [x] **T-014** `done` — Release ep 006
  - Phase: sprint-06
  - Artifact: `continuity/story-so-far.md`
  - Depends: T-013 ✓
  - Notes: released
- [ ] **T-015** `pending` — Design ep 007
  - Phase: sprint-07
  - Artifact: `episodes/007-….md`
  - Depends: T-014 ✓
  - Notes: —
```

Same pattern for Snapshot count bumps: `old_string` = current Snapshot lines that exist; `new_string` = updated counts (not a brand-new Snapshot invented as the anchor).

**Batch edits (same model step):** When Snapshot counts, Current focus, and task checkbox/Notes all change together — prefer **one** `edit_file` covering them, or **several non-overlapping `edit_file` calls in the same tool round** (parallel). Do **not** spend separate model rounds for "Snapshot only" then "task checkbox only" when both were already known. Same-file parallel edits only if each `old_string` remains valid after the others (disjoint regions); otherwise one combined edit / `write_file`.

**`old_string_not_found` with whitespace-only mismatch:** If `first_mismatch.expected` vs `.found` differ only by indent/spaces, copy **`found`** verbatim into the next `old_string`. After **two** consecutive failures on the same path in one turn, switch to **`write_file`** for the board (or the whole design artifact) instead of more micro-anchors.

---

## Plan file templates (`scrum/`)

### product-backlog.md

Update when **scope/plan** changes — not on each task completion.

```markdown
# Product Backlog — {Title}

> **Updated:** {YYYY-MM-DD}
> **Execution board:** [`../TODO.md`](../TODO.md)

## Stories

- **PB-001** — {title}
  - Band: {epic or sprint-00 / episode}
  - Artifact: `{path or —}`
  - Depends: —
  - Notes: …
```

No progress checkboxes. **Series:** `PB-000` + episode-mapped stories. **Spec-driven:** Band = Epic; cite `spec/….md §…`.

### sprint-NN.md

Zero-pad (`sprint-00.md`, …). Leave DoD/Demo unchanged during execution.

```markdown
# Sprint {NN} — {short goal title}

> **Product Backlog:** [`product-backlog.md`](./product-backlog.md)
> **Execution board:** [`../TODO.md`](../TODO.md)

## Sprint Goal
{1–3 sentences}

## Scope

| Band | Stories |
|------|---------|
| … | PB-xxx … |

**Out of sprint:** {next slice}

## Planned work

| ID | Story |
|----|-------|
| PB-xxx | … |

## Definition of Done
1. …
2. …

## Demo Checklist
- …   # plain bullets — mirror as gate tasks on TODO.md
```

---

## Core Workflows

### 1. Initialize

1. Resolve `{project-root}` and companion workflow.
2. If **scrum mode** triggers: choose profile → create `scrum/` + `product-backlog.md` + first `sprint-NN.md` → seed `TODO.md` from that plan (empty Completed sprints).
3. Else: seed phase-level tasks; optional phases `[x]` + `_(skipped)_`.
4. Current focus = first open task; tell the user (final turn includes `session` per principles).

### 2. Orient

1. Read `TODO.md` (missing → Initialize).
2. If Current sprint set: optionally skim that `sprint-NN.md` Goal/DoD — not the full backlog unless planning.
3. Drift vs disk → Sync.
4. Report focus, counts, recommended next.

### 3. Execute

1. Exactly one Current focus.
2. Draft done → `[x]`; focus gate or next.
3. Gate approved → `[x]` only with explicit user approval.
4. Batch approval → flip matching headlines in place.
5. Refresh Snapshot + `Updated` → **Sync**.
6. Scrum mode and sprint done → **Sprint boundary**.

### 4. Sync

1. Compare disk to current `## Tasks`.
2. Append missing; skip removed scope with `[x]` + `_(skipped)_`.
3. Recompute Snapshot.
4. Do not pull future episodes/Epics onto this board.

### 5. Plan / Plan Next Sprint (scrum mode)

**Plan:** Load companion artifacts → fill backlog → write/update `sprint-NN.md` → seed or switch `TODO.md` only when that sprint should be active.

**Plan Next:** When the next `sprint-NN.md` is missing (or user pre-plans): pick next Stories → **create** `sprint-NN.md`. Do not record completion in plan files — that is `TODO.md`.

### 6. Sprint boundary (scrum mode)

1. Append **Completed sprints** row; drop closed task list from the board.
2. Find next `scrum/sprint-*.md`.
3. Exists → rewrite `TODO.md` Tasks + Snapshot + Current sprint (keep Completed table).
4. Missing → **Plan Next Sprint** → then rewrite `TODO.md`.
5. Do not edit backlog/closed sprint files merely to mark done.

### 7. Report Progress

```markdown
### Progress — {Title}

- **Sprint:** {sprint-NN or "—"} — {goal}
- **Phase:** {phase}
- **Done:** {done}/{total} ({percent}%)
- **Current focus:** {T-id} — {title}
- **Recommended next:** {T-id} — {title} ({why})
```

## Next-Task Selection

1. Current focus open and not blocked → continue.
2. Unblocked blocked task → clear suffix, select it.
3. First `[ ]` whose Depends are all `[x]`.
4. Prefer gate over new draft when draft closed and gate open.
5. Else ID order within phase; research may run while a gate waits (if asked).
6. No open tasks + scrum mode → **Sprint boundary** (unless backlog idle → deliverable complete).

## Completion

**Plain:** All non-skipped `[x]`; final gate `[x]`; open = 0; focus `—`.

**Scrum:** No remaining in-scope backlog Stories and final sprint board closed; focus `—`. Phase incomplete until its gate is `[x]`. Do not start the next phase draft while the previous gate is open **in the same sprint**.

## Prohibitions

- Separate `scrum` skill / `read_skills` scrum — this file is complete
- `scrum/README.md` or progress under `scrum/`
- Full closed-sprint task lists on `TODO.md`
- Mixing Spec/ADR authoring into implementation sprint boards
- Inventing episode plot in backlog Notes (point to story artifacts)
