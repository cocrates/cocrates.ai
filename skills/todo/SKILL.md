---
name: todo
description: >-
  Manages deliverable-generation tasks in TODO.md — initialize task lists,
  track progress, sync with workspace artifacts, and determine the next actionable task.
  Use when starting or resuming a deliverable session, when the user asks about progress,
  next steps, or task planning, or alongside any deliverable-generation workflow.
compatibility: opencode
metadata:
  agent: cocrates
---

# TODO — Deliverable Task Tracking Skill

This skill maintains **`TODO.md`** as the single source of truth for deliverable-generation work. It **tracks** tasks implied by the deliverable workflow so progress is visible and the next action is always clear.

Match the user's language (Korean, English, etc.) in `TODO.md` prose and status reports.

## When to Use

| Situation | Action |
|-----------|--------|
| New deliverable session starts | Initialize `TODO.md` from the deliverable's workflow |
| User resumes work on an existing deliverable | **Read `TODO.md` first** — orient before acting |
| A task is completed (draft written or gate approved) | Update status, run **Sync**, recommend next task |
| User asks "where are we?", "what's next?", "progress?" | Read `TODO.md`, report summary, recommend next task |
| User changes scope or skips work | Sync backlog — add/skip tasks; never delete completed history |

## Working Location

`TODO.md` lives at the **deliverable workspace root**:

```text
docs/{title-slug}/TODO.md
presentations/{title-slug}/TODO.md
{workspace}/{title-slug}/TODO.md   # other artifact-generation skills
```

One `TODO.md` per deliverable session. Do not use the agent's ephemeral todo tool as a substitute — the file survives across sessions.

## Task Model

### Status Values

| Status | Meaning |
|--------|---------|
| `pending` | Ready or waiting; not started |
| `in_progress` | Currently being worked on — **at most one** per session |
| `blocked` | Cannot proceed; record reason in task notes |
| `done` | Completed; gate tasks need **explicit user approval** before `done` |
| `skipped` | Intentionally not done |

### Task ID Convention

- Format: `T-{nnn}` — three digits, zero-padded (`T-001`, `T-012`)
- IDs are **stable** — never renumber or reuse IDs after creation
- New tasks append the next available ID

### Task Fields

Every task line encodes:

```text
- [ ] **T-00n** `status` — {short title}
  - Phase: {workflow phase name or "cross-cutting"}
  - Artifact: `{relative path}` or —
  - Depends: T-00m ✓, T-00k
  - Notes: {optional — blockers, decisions, approval date}
```

Checkbox mirrors status: `[x]` = `done` or `skipped`; `[ ]` = everything else.

### Task Types

| Type | Description | Done when |
|------|-------------|-----------|
| **draft** | Produce or update an artifact file | File written; presented to user |
| **gate** | User review & explicit approval | User approves in dialogue |
| **unit** | Granular work inside a phase | Unit artifact approved or batch-approved |
| **research** | Cross-cutting reference material | Findings saved and presented |
| **sync** | Reconcile TODO with workspace reality | Backlog matches actual artifacts |

## TODO.md Structure

Use this skeleton when creating or rebuilding the file:

```markdown
# TODO: {Title}

> **Workspace:** `{path}`
> **Updated:** {YYYY-MM-DD}

## Snapshot

| Done | In progress | Pending | Blocked | Skipped |
|------|-------------|---------|---------|---------|
| {n}  | {n}         | {n}     | {n}     | {n}     |

**Current focus:** {T-id or "—"}
**Recommended next:** {T-id} — {one-line description}

## Active

{at most one in_progress task, expanded}

## Backlog

{grouped by workflow phase; pending and blocked tasks}

## Completed

{collapsed or abbreviated list of done/skipped — keep last 10 visible, rest summarized}

## Notes

{scope decisions, skipped phases, batch-approval agreements, open questions}
```

Keep **Snapshot** counts accurate — recompute whenever statuses change.

## Core Workflows

### 1. Initialize

Trigger: user starts a new deliverable; workspace folder created or about to be created.

1. Identify workspace path and the deliverable's workflow (stages, artifacts, gates).
2. Use [templates.md](templates.md) as a starting point; adapt phases and tasks to the specific deliverable.
3. Customize:
   - Mark optional phases `skipped` when the user has already decided to omit them
   - Add **unit tasks** when scope is known early; otherwise leave coarse phase-level tasks and expand during **Sync**
4. Write `TODO.md` with all tasks `pending` except the first task → `in_progress`.
5. Tell the user the task plan exists and state the first focus task.

Do **not** wait for every unit to be known upfront — start with phase-level tasks; expand during **Sync**.

### 2. Orient (session start)

**Always run before working on a deliverable**, even mid-conversation:

1. Read `{workspace}/TODO.md`. If missing, run **Initialize**.
2. Verify **Snapshot** against filesystem (quick check: does the `in_progress` artifact exist? are `done` gate artifacts present?).
3. If drift detected, run **Sync** before recommending next work.
4. Report to user:
   - Current phase and focus task
   - Progress counts
   - Recommended next task with rationale

### 3. Execute

While working on deliverable tasks:

1. **Start:** set exactly one task to `in_progress`; clear any stale `in_progress` on other tasks.
2. **Draft complete:** mark draft task `done`; create or activate the matching **gate** task as `in_progress` or next `pending`.
3. **Gate approved:** mark gate `done`.
4. **Unit batch:** when user approves a batch, mark all corresponding unit + gate tasks `done` in one update.
5. Update **Snapshot**, **Current focus**, **Recommended next**, and **Updated** date on every status change.
6. **Run Sync** (see below) before selecting the next task.

### 4. Sync

**Primary trigger:** a task is marked `done` — run Sync as the last step of **Execute**, before recommending the next task.

**Secondary triggers:** orient detects drift between `TODO.md` and workspace; user changes scope.

1. Review the just-completed task and its artifacts — does completion reveal new units or phases not yet in the backlog?
2. List artifacts on disk vs tasks in backlog.
3. **Add** missing unit or phase tasks with new IDs.
4. **Skip** removed scope — mark `skipped`, do not delete.
5. **Never** remove or rewrite `done` / `skipped` history.
6. Recompute **Snapshot**, **Current focus**, and **Recommended next**.
7. Record a `sync`-type task entry only when reconciliation itself was non-trivial; mark it `done` when finished.

### 5. Report Progress

When the user asks for status or at natural breakpoints (gate approved, phase complete):

```markdown
### Progress — {Title}

- **Phase:** {current workflow phase}
- **Done:** {done}/{total} ({percent}%)
- **In progress:** {T-id} — {title}
- **Recommended next:** {T-id} — {title} ({why it's unblocked})

{optional: blocked items, skipped phases, risks}
```

Match the user's language in report prose when they communicate in a language other than English.

## Next-Task Selection Rules

Apply in order; pick the **first** match:

1. If a task is `in_progress` and not blocked → **continue it** (do not switch without user consent).
2. Else if a `blocked` task has its blocker resolved → set `pending`, then select it.
3. Else first `pending` task where **all Depends are `done` or `skipped`**.
4. Prefer **gate** tasks over new **draft** tasks when a draft is done but gate is still `pending` — approval before new work.
5. Within a phase, respect numeric ID order unless dependencies say otherwise.
6. **Research** tasks are parallelizable — may run while a gate waits for user, if the user requested research.

When multiple valid next tasks exist, present the recommended one and mention alternatives.

## Gate Invariant

A phase is not complete in `TODO.md` until its **gate task** is `done`. Never mark a gate `done` without explicit user approval.

## Prohibitions

- Tracking deliverable work only in chat or the agent todo tool — **write `TODO.md`**
- Multiple `in_progress` tasks without user-approved parallel work
- Marking gate tasks `done` without explicit user approval
- Deleting completed tasks to "clean up"
- Renumbering task IDs
- Starting a new phase's draft while the previous phase's gate is still `pending`
- Skipping **Sync** after marking a task `done`

## Completion Criteria

The deliverable session's TODO is complete when:

- All non-skipped tasks are `done`
- The deliverable meets its defined completion criteria
- Final gate (user confirms deliverable complete) is recorded as `done`
- **Snapshot** shows 0 `pending`, 0 `in_progress`, 0 `blocked`

## Additional Resources

- Starter template and unit-expansion pattern: [templates.md](templates.md)
