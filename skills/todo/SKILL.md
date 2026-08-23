---
name: todo
description: >-
  Manages deliverable-generation tasks in TODO.md — initialize, track progress,
  sync with workspace artifacts, and pick the next actionable task. For large or
  iterative work (series, spec-driven), also maintains scrum/product-backlog.md
  (story status), sprint-NN.md (plans), and scrum/TODO-NN.md (closed-sprint
  archives). Use when starting or resuming a deliverable session, progress or
  next-step questions, sprint/backlog planning, or alongside any deliverable
  workflow.
metadata:
  agent: cocrates
---

# TODO — Deliverable Task Tracking Skill

Match the user's language in prose and status reports. Do not use the agent's built-in `todowrite` tool — on-disk boards are the source of truth.

| File | Role |
|------|------|
| `{project-root}/TODO.md` | **Current sprint** progress SSOT (Snapshot, tasks, Completed sprints index) |
| `scrum/product-backlog.md` | **Product** story status SSOT (survives sprint rewrites) |
| `scrum/sprint-NN.md` | Sprint **plan** (Goal, Scope, DoD, Demo) — not a live status board |
| `scrum/TODO-NN.md` | **Archive** of `TODO.md` at close of sprint-NN |

**Churn:** Update `TODO.md` on every task advance. Update `product-backlog.md` when stories are planned, reshaped, completed, skipped, or blocked — usually at **Sprint Boundary**, not on each `T-xxx`. Write/edit `sprint-NN.md` only when planning. Write `TODO-NN.md` only at Sprint Boundary (then leave immutable).

`{project-root}` comes from the companion skill.

## When to Use

| Situation | Action |
|-----------|--------|
| New deliverable | **Initialize** |
| Resume / mid-session | **Orient** |
| Draft done or gate approved | **Execute** → **Sync** → recommend next |
| "where are we?" / "what's next?" | **Report** from `TODO.md` (+ backlog if product-wide) |
| Scope change within sprint | **Sync** |
| Backlog / sprint plan needed | **Plan** / **Plan Next Sprint** |
| Sprint Goal met / all sprint tasks closed | **Sprint Boundary** |

## Working Location

```text
{project-root}/
├── TODO.md
└── scrum/                    # scrum mode only
    ├── product-backlog.md
    ├── sprint-NN.md
    └── TODO-NN.md
```

---

## Scrum Mode

| Trigger | Examples |
|---------|----------|
| **Series deliverable** | novel, webtoon, picture-book — **default on** |
| **Spec-driven implementation** | Specs ready — iterative build |
| **User asks** | "sprint", "backlog", "scrum", "split it up" |
| **Scale** | Would exceed ~30 tasks, or unbounded episode loop |
| **Next plan missing** | At Sprint Boundary, no next `sprint-NN.md` |

**Plain mode** (no `scrum/`): small one-shot work unless the user asks for sprints.

### Profiles

**Series** (default on):

| Sprint | Scope |
|--------|--------|
| `sprint-00` | ① define → ② plan → ③ architecture |
| `sprint-01`… | One publication unit (episode / equivalent) |

Prefer pre-writing upcoming `sprint-NN.md` when the Episode List is known. Narrative SSOT remains `series.md` (etc.); backlog stories map to those rows.

**Spec-driven** (after Specs are sufficient):

| Phase | Board |
|-------|--------|
| PRD · ASR · ADR · Spec writing | Plain todo — not implementation sprints |
| Implementation | `scrum/` + `TODO.md`; each sprint = vertical slice |

Every backlog story cites its Spec (`Spec:` field).

---

## Task Model

### States

| Signal | Meaning |
|--------|---------|
| `- [ ]` | Open |
| `- [x]` | Closed (done, or `_(skipped)_`) |
| Snapshot **Current focus** | At most one in-progress open task |
| `_(blocked: …)_` | Open but blocked |
| `_(skipped)_` | Closed without doing the work |

Checkboxes + Current focus only — no `` `pending` `` / `` `in_progress` `` / `` `done` `` tokens on headlines. Same signals apply to **PB** stories on the backlog.

**IDs**

| Kind | Form | Stability |
|------|------|-----------|
| Task | `T-{nnn}` | Stable within one board; never renumber/reuse on that board. May restart at `T-001` when `TODO.md` is rewritten for a new sprint. |
| Story | `PB-{nnn}` | Stable for the deliverable; never renumber/reuse. Append new IDs only. |

### Task entry

```text
- [ ] **T-00n** — {short title}
  - Phase: {phase or "cross-cutting"}
  - Artifact: `{path}` or —
  - Depends: T-00m, T-00k
  - Notes: {optional; may note type: draft | gate | unit | research | sync}
```

| Type | Purpose |
|------|---------|
| `draft` | Write/generate an artifact |
| `gate` | Explicit user approval before proceeding |
| `unit` | Atomic slice within a phase |
| `research` | Investigation; may run while a gate waits |
| `sync` | Keep artifacts consistent |

**Story → tasks:** Each scoped `PB-*` becomes one or more `T-*` on the sprint board (drafts, units, gates). Prefer Notes or title to show the parent story (e.g. `PB-012`). Closing all tasks for a story does not auto-close the PB item — mark PB at Sprint Boundary (or sooner if the story is clearly done).

---

## TODO.md Structure

### Plain mode

```markdown
# TODO: {Title}

> **Project type:** `{type}`
> **Project root:** `{project-root}`
> **Updated:** {YYYY-MM-DD}

## Snapshot

| Done | Open | Blocked | Skipped |
|------|------|---------|---------|
| {n}  | {n}  | {n}     | {n}     |

**Current focus:** {T-id or "—"}
**Recommended next:** {T-id} — {one-line description}

## Tasks

{task entries}

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

{task entries — current sprint only}

## Completed sprints

| Sprint | Goal | Closed | Archive |
|--------|------|--------|---------|
| 00 | {one-liner} | {YYYY-MM-DD} | [`TODO-00.md`](./scrum/TODO-00.md) |

## Notes

{deferrals, carry-over}
```

**Completed sprints** accumulates across rewrites; never wipe it when starting the next sprint.

### Session status

- **Scrum:** header **Current sprint** `sprint-NN` → `S-NN` (e.g. `sprint-07` → `S-07`).
- **Plain:** Snapshot **Current focus** T-id.

---

## Plan File Templates (`scrum/`)

### product-backlog.md

```markdown
# Product Backlog — {Title}

> **Updated:** {YYYY-MM-DD}
> **Execution board:** [`../TODO.md`](../TODO.md)

## Stories

- [ ] **PB-001** — {short title}
  - Story: {concrete Done-when / who-wants-what}
  - Spec: `{spec/….md}` §{section}   # or —
  - Sprint: {sprint-NN or —}
  - Depends: —
  - Notes: …
```

| Field | Rule |
|-------|------|
| **Story** | Required. Acceptance-oriented; not a vague label. |
| **Spec** | **Spec-driven:** required cite of governing Spec path + section (extra refs in Notes). **Series:** `—`. |
| **Sprint** | Set when committed to a sprint; keep after done for history. |
| **Depends** | Other `PB-*` IDs, or `—`. |

**Series layout:** `PB-000` (skeleton / cross-cutting) + one story per publication unit. **Spec-driven:** Epic grouping only in Notes if useful — Spec citation is the link to design.

**Carry-over:** If a story misses DoD, leave `[ ]`, set **Sprint** to the next sprint (or `—` until replanned), and note carry-over in both boards' Notes.

### sprint-NN.md

Zero-pad (`sprint-00.md`, …). Do not edit Goal/DoD/Demo during execution to “mark done.”

```markdown
# Sprint {NN} — {short goal title}

> **Product Backlog:** [`product-backlog.md`](./product-backlog.md)
> **Execution board:** [`../TODO.md`](../TODO.md)

## Sprint Goal
{1–3 sentences}

## Scope

| ID | Story |
|----|-------|
| PB-xxx | {short title — match backlog headline} |

**Out of sprint:** {next slice or —}

## Definition of Done
1. …
2. …

## Demo Checklist
- …   # mirror as gate tasks on TODO.md
```

Scope lists committed `PB-*` only. Seed `TODO.md` tasks from Scope + companion workflow — do not maintain a second story table.

### TODO-NN.md

Exact copy of `TODO.md` at Sprint Boundary for sprint-NN. Immutable after a successful copy (fix copy errors only). Linked from **Completed sprints** → Archive.

---

## Core Workflows

### 1. Initialize

1. Resolve `{project-root}` and companion workflow.
2. **Scrum:** choose profile → create `scrum/` + backlog (open `[ ]`, Story/Spec rules) + first `sprint-NN.md` → seed `TODO.md` from Scope (empty Completed sprints) → set **Sprint:** on scoped PB items.
3. **Plain:** seed phase-level tasks; optional unused phases `[x]` + `_(skipped)_`.
4. Current focus = first open task; report (include `session` per principles).

### 2. Orient

1. Read `TODO.md` (missing → Initialize).
2. Optionally skim current `sprint-NN.md` Goal/DoD; read backlog when product-wide status matters.
3. Drift vs disk → Sync.
4. Report focus, counts, recommended next.

### 3. Execute

1. Exactly one Current focus.
2. Draft done → `[x]`; move to gate or next.
3. Gate → `[x]` only with explicit user approval (batch: flip matching headlines in place).
4. Refresh Snapshot + `Updated` → Sync.
5. Sprint complete → Sprint Boundary.

### 4. Sync

1. Compare disk artifacts to current `## Tasks`.
2. Append missing work; drop removed scope with `[x]` + `_(skipped)_`.
3. Recompute Snapshot.
4. Do not pull future episodes / unscoped Epics onto this board.

### 5. Plan / Plan Next Sprint

**Plan:** Companion artifacts → backlog stories → `sprint-NN.md` Scope → set PB **Sprint:** → seed/switch `TODO.md` only when that sprint is active.

**Plan Next:** If next `sprint-NN.md` is missing (or user pre-plans): pick open/carried Stories → create the plan file. Do not rewrite live `TODO.md` until that sprint starts (unless Initialize / Boundary).

### 6. Sprint Boundary

1. **Archive:** copy `TODO.md` → `scrum/TODO-{NN}.md` (zero-padded NN).
2. **Backlog:** for each Scope story that met DoD → `[x]`; intentionally dropped → `[x]` + `_(skipped)_`; incomplete → carry-over (above). Bump backlog **Updated:**.
3. Append **Completed sprints** row (Goal, Closed date, Archive link).
4. Next `sprint-*.md` **exists:** rewrite live `TODO.md` tasks/Snapshot/Current sprint from that plan; **keep** Completed sprints table; do not leave the closed task list on the live board.
5. Next plan **missing:** Plan Next Sprint, then rewrite as above.

### 7. Report Progress

```markdown
### Progress — {Title}

- **Sprint:** {sprint-NN or "—"} — {goal}
- **Phase:** {phase}
- **Done:** {done}/{total} ({percent}%)
- **Current focus:** {T-id} — {title}
- **Recommended next:** {T-id} — {title} ({why})
- **Backlog:** {done}/{total} stories closed   # scrum, when relevant
```

## Next-Task Selection

1. Current focus open and not blocked → continue.
2. Unblocked blocked task → clear suffix, select it.
3. First `[ ]` whose Depends are all `[x]`.
4. Prefer gate over new draft when draft closed and gate open.
5. Else ID order within phase; research may run while a gate waits (if asked).
6. No open tasks + scrum → Sprint Boundary (unless backlog idle → deliverable complete).

## Completion

**Plain:** All non-skipped `[x]`; final gate `[x]`; open = 0; focus `—`.

**Scrum:** No in-scope PB stories left open; final sprint board closed; focus `—`. Do not start the next phase draft while the previous gate is still open **in the same sprint**.

## Prohibitions

- Closed-sprint task lists on live `TODO.md` (use `TODO-NN.md`)
- PB left `[ ]` after DoD met
- Vague PB headlines without **Story:**
- Spec-driven PB without **Spec:**
- Renumbering / reusing `T-*` on a board or `PB-*` in a deliverable
- Editing `TODO-NN.md` or sprint DoD/Demo to fake completion
