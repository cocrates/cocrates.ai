---
name: novel-authoring
description: >-
  Creates serialized narrative fiction — novels and web novels — episode by
  episode. Select when the user asks to write, draft, compose, or author a novel,
  story, fiction, series, web novel, or any narrative fiction work. Runs
  Define → Plan → Architecture → Episode Design → Evaluate/Revise → Generate →
  Evaluate/Revise → Release with mandatory approval gates; builds series-level
  world, characters, and locations at architecture, then incrementally expands
  (prefer add over modify) during episode design; locks approved episode design
  before manuscript generation. Structure is series → episode → scene; episode
  is the publication/release unit. Default workflow: architecture complete for
  the series, then one episode at a time. Manuscript evaluation is the main
  quality gate; design evaluation is recommended before prose. Prose-only — no
  illustration or image generation in this skill. Pipeline and shared protocols
  in SKILL.md; stage procedures in workflow/01–08. For general non-narrative
  documents, use document-authoring instead.
metadata:
  agent: cocrates
---

# Novel Authoring — Serialized Fiction Generation Skill

Produces **complete narrative fiction** — from a single short story to a multi-episode web novel — through a structured, stage-gated pipeline. Each stage produces artifact files; each requires **explicit user approval** before proceeding.

**Publication unit:** **Episode.** Scenes are design-only sections inside that episode’s design file (④ must be generation-ready before ⑥).

**No illustrations:** This skill generates **prose only**. Do not generate reference PNGs, image YAML, or in-novel illustrations. Covers or other still art require a **separate user request** via `image-generation`.

**Setup (once):** ① define → ② plan (`series.md`) → ③ **series-level** architecture → **episode loop** ④ design → ⑤ evaluate design (recommended) → ⑥ manuscript → ⑦ evaluate manuscript (main gate) → ⑧ release. Cross-layer sync: § Consistency.

---

## How to Use This Skill

**This file** — pipeline rules, approval gates, global constraints, and shared protocols (Consistency, Reference Models, Revision, Research). Orient here each session. It does **not** replace numbered stage workflows.

**Stage workflows** — `novel-authoring/workflow/01`–`08` only. Read the matching stage file **before** starting or resuming that stage. Do not design, evaluate, generate, revise, or release from memory or a prior-turn summary.

**Hard rule — no stage without its workflow:**
1. State the stage (e.g. *"Stage ④ — episode design."*).
2. Read the workflow that stage’s `workflow/0N-….md`.
3. Only then produce or edit artifacts for that stage.

**Selective artifact load (hard):** Never dump every profile, prior episode, or continuity file “for context.”

1. **Index pass (first):** `overview.md`, `series.md`, `characters.md`, `locations.md`, `world-bible.md`, `stagings.md` — enough to know what exists.
2. **Detail pass:** only paths this stage needs (Appearing cast, used locations, cited stagings, current episode/manuscript/eval target; ep 002+ → `continuity/story-so-far.md` + **immediate prior** summary only).

Stage Load sections in `04` / `05` / `06` / `07` implement this. Do not invent a fuller dump.

| Stage | Purpose | Workflow | Approval gate | Gate artifact |
|-------|---------|--------------------------|---------------|---------------|
| ① define | Scope, genre, criteria; **episode-scale only** | `01-define.md` | **G1** | `overview.md` |
| ② plan | Episode List + series arc | `02-plan.md` | **G2** | `series.md` |
| ③ architecture | World, cast, places — **series foundation** | `03-architecture.md` | **G3** | `world-bible.md`, `characters*`, `locations*` |
| ④ episode design | Episode + scenes + **situation stagings** | `04-episode-design.md` | **G4** | `episodes/{nnn}-{slug}.md`, `stagings*` as needed |
| ⑤ evaluate (design) | Design validation + critique — **recommended** | `05-evaluate-design.md` | **G5** | `evaluations/{nnn}-{slug}-design.md` |
| ⑥ generation | Prose locked to design + architecture + continuity | `06-generate.md` | **G6** | `manuscripts/{nnn}-{slug}.md` |
| ⑦ evaluate (manuscript) | **Main gate** — manuscript validation + critique | `07-evaluate-manuscript.md` | **G7** | `evaluations/{nnn}-{slug}-manuscript.md` |
| ⑧ release | Lock episode; update continuity | `08-release.md` | **G8** | Continuity files |

**Gate IDs match stages (G1–G8).** Do not reuse a gate number for a different stage.

### Default Episode Loop

Unless the user explicitly requests **design-ahead** (designs only, not manuscripts):

```
④ design episode N → G4
  → ⑤ evaluate design (recommended) → revise if needed → G5
    → ⑥ write manuscript → Design-Fidelity + Prose Quality Floor → G6
      → ⑦ evaluate manuscript (required) → revise if needed → G7
        → ⑧ release + continuity → G8
          → ④ design episode N+1 → ...
```

**Hard rules:**
- No **manuscript** for episode N+1 until N is **released** (G8) — **no exceptions**
- No episode N+1 **design** until N released (unless user requests design-ahead)
- Every artifact needs **explicit user approval**
- **⑤ recommended** — user may skip to ⑥ after G4
- **⑥** — pass Design-Fidelity + Prose Quality Floor before G6 (`06-generate.md`)
- **⑦ mandatory** before release (`07-evaluate-manuscript.md`)
- Architecture expands **additively** during ④ for local gaps (§ Consistency). Stage ③ is series foundation, not first-episode stubs.

### Approval Applies To

- `overview.md`, `series.md`, architecture, each episode design, each manuscript, evaluations, release
- Any **new** character, location, world rule, or staging (approve before citing)
- Any **modifying** architecture/staging change (user acknowledges impact)
- Consistency Cascade file set when a higher layer changes

### Agent Procedure

1. **Orient** — use `todo` skill; read `{project-root}/TODO.md` (scrum Initialize if series and `scrum/` missing)
2. State the current stage
3. Read the stage’s workflow — every time
4. Write artifacts to disk — never chat-only deliverables
5. Present and wait for explicit approval; update todo on draft/gate progress
6. After G8 → continuity briefing → sprint boundary → next ④ (reload continuity)
7. Mid-flight changes → § Consistency cascade before continuing

---

## Structure & Artifacts

**Plot hierarchy:** `Series → Episode → Scene (design)`

| Level | Artifact | Stage |
|-------|----------|-------|
| Constraints | `overview.md` | ① |
| Series | `series.md` | ② |
| World / cast / place | `world-bible.md` · `characters*` · `locations*` | ③ (+ additive ④) |
| Situation staging | `stagings.md` · `stagings/*` | ④ (with episode) |
| Episode design | `episodes/{nnn}-{slug}.md` | ④ |
| Prose | `manuscripts/{nnn}-{slug}.md` | ⑥ |
| Evaluation | `evaluations/*-design.md` · `*-manuscript.md` | ⑤ / ⑦ |
| Continuity | `continuity/*` | ⑧ write; ④/⑥/⑦ load |

## Resolve Project Root

Before writing artifacts, resolve **`{project-root}`** — the folder that holds this novel.

Workspace layouts fall into three types. **Inspect the workspace first**, then match:

| Type | When | `{project-root}` |
|------|------|------------------|
| **1** | Workspace *is* the single novel project | `.` (workspace root) |
| **2** | Workspace holds multiple peer projects | `{title-slug}/` |
| **3** | Workspace groups projects by kind | `novels/{title-slug}/` |

**Examples:**
- Type 1 → `overview.md` at workspace root
- Type 2 → `{title-slug}/overview.md`
- Type 3 → `novels/{title-slug}/overview.md`

**Rules:**
1. Infer the type from existing structure (e.g. `overview.md` at root, peer project folders, or a `novels/` kind folder). Prefer an **existing** matching folder over creating a new one.
2. **Before creating** a new project folder, confirm **location and name** with the user.
3. If the project folder already exists, use it — do not recreate or relocate silently.
4. **`{title-slug}`:** URL-friendly slug from the novel title. Type 1 has no slug folder.

## Working Location

**Scaffold at ①:** when creating `{project-root}/`, also create `world/`, `characters/`, `locations/`, `stagings/`, `episodes/`, `manuscripts/`, `evaluations/`, `continuity/`, `scrum/` in **one** tool round. Details: `01-define.md` § Project scaffold.

```
{project-root}/
├── TODO.md · scrum/
├── overview.md · series.md
├── world-bible.md · world/
├── characters.md · characters/
├── locations.md · locations/
├── stagings.md · stagings/          # ④
├── episodes/ · manuscripts/ · evaluations/ · continuity/
└── docs/                            # optional research
```

### Task tracking (todo + scrum)

| Sprint | Scope |
|--------|--------|
| `sprint-00` | ① → ② → ③ |
| `sprint-01`… | One episode release (④→⑧, + additive architecture as needed) |

Story plan stays in `series.md`. Orient `TODO.md` first (use `todo` skill).

---

## Dialogue Rules

1. State stage; read the workflow before any stage action
2. Stage ①: one question at a time; episodes only (never scene planning)
3. After G8, prompt next episode and wait
4. Propose ⑤ after G4; **always** run ⑦ before release
5. Critique → **design** → prose; prefer **add** over **modify** on architecture gaps
6. Load gates at ④/⑥; Consistency Cascade on higher-layer edits

---

## Prohibitions

Global hard stops. Stage-specific MUST NOTs live in the matching workflow — this list is not a substitute for loading that workflow.

### Workflow & deliverables
- Enter any stage without reading that stage’s `01`–`08` workflow
- Chat-only design, manuscript, or evaluation (no gate artifact on disk)
- Paste skill/workflow text into `episodes/`, `manuscripts/`, or `evaluations/`

### Structure & order
- Scene planning at ①; scene counts in `series.md` / `overview.md`
- Episode design before approved ② + ③ series-level architecture
- Treating ③ as first-episode-only stubs
- Marking G4 done because the user asked only for ⑤
- Manuscript (⑥) before approved ④ (+ recommended ⑤)
- Episode N+1 **manuscript** before N released
- Nested `episodes/{slug}/` scene files; non-canonical episode filenames
- English-alias catalog/staging filenames when names are Non-Latin (e.g. `characters/han-seo-yun.md` instead of `characters/한서윤.md`)
- Image generation / reference PNGs / illustration YAML as part of this skill
- Authoring `stagings*` as Stage ③ gate deliverables

### Design fidelity & catalogs (summary)
- Invent catalog facts in Key Events or prose without architecture approval
- Cite missing Architecture paths; mark Load ✅ for missing files; index-only rows without profiles
- Free-text locations; cite `+ facet` not listed under **Multi-facet anchors**
- Ghost cast / Dialogue speakers absent from On stage; Appearing ≠ On stage union
- Soften/harden Hook evidence across body surfaces; Out creep beyond Summary + Hook to Next
- Present under-length, repeated, pasted, author-meta, or **installment-meta** manuscript (body cites this work’s own episode/part/화 numbers as time or place, e.g. “In episode 001 I didn’t see that scratch.”); skip Design-Fidelity or Prose Quality Floor
- Stage ⑤ editing `episodes/` (return to ④; Adjudication stays Pending until a later ⑤ turn)

**Full checklists, Gate Evidence, forecast math, Path existence, Floor rows:** `04`–`07` only.

---

## Completion Criteria

**Per episode:** ④ → [⑤] → ⑥ → ⑦ → revisions → ⑧ + continuity

**Work complete:** Episode List complete → all episodes designed and released → user final approval

---

# Shared protocols

## Consistency — Cross-Stage Protocol

**Not a numbered stage.** Apply when starting ②–④, revising (⑤/⑦ loops), discovering a catalog gap in ④, or changing any approved artifact.

### Glossary

| Term | Meaning |
|------|---------|
| **Episode List source** | `series.md` Episode List (Summary + Hook to Next) |
| **Architecture layer** | `world-bible.md`, `world/*`, `characters*`, `locations*` (③; expand at ④) |
| **Episode design** | `episodes/{nnn}-{slug}.md` — Alignment, Scene Index, Key Events |
| **Scene Index** | TOC inside episode design (not a location facet list) |
| **Staging** | Situation stage (`stagings*`) — authored in **④** |
| **Design Consistency Gate** / **Generation Readiness** | Stage ④ checks — `04-episode-design.md` |
| **Design-Fidelity Gate** / **Prose Quality Floor** | Stage ⑥/⑦ — `06-generate.md` / `07-evaluate-manuscript.md` |
| **Additive / Modify / Retcon** | Prefer add; modify cascades; retcon needs republication |

### Source-of-truth stack (top wins)

```
overview.md
  → series.md
    → world-bible.md + world/* + characters* + locations*
      → stagings* + episodes/{nnn}-*.md
        → evaluations/*
          → manuscripts/*
continuity/*   [after ⑧; binds ④/⑥/⑦ for ep 002+]
```

| Change… | Update first | Then sync |
|---------|--------------|-----------|
| Scale, validation | `overview.md` | Episode List |
| Arc / Episode List row | `series.md` | conflicting designs |
| **New** cast / place / world / staging | architecture or `stagings*` (additive) | cite in design |
| **Modify** approved profile/rule | warn → approve → **cascade** | open designs/MS; released only via republication |
| Scene beats / seeds | episode design | manuscript — never prose-only for plot |
| Released past | don’t — unless republication | design → MS → re-release → continuity |

### Incremental architecture

- **③:** series-level foundation from the full Episode List (cold reader can outline the novel). **No stagings at ③.**
- **④:** compose approved catalogs; **author stagings**; fill local gaps with additive expands. Gap → stop → expand → approve → resume. Location `+ facet` cites require **Multi-facet anchors** labels (`04-episode-design.md`). Missing facet → **Extend** — do not relocate onto an already-used facet to dodge the gap.

### Additive-first (A/B/C/D)

| Type | Examples | Default |
|------|----------|---------|
| **A — Add** | New character, location, world aspect, character **state**, staging | Add → approve → cite |
| **B — Extend** | New Multi-facet anchors label; additive subsection | Append → approve → cite |
| **C — Modify** | Rename, rewrite core drive/rule/layout readers already saw | Stop → impact → cascade |
| **D — Retcon** | Contradict `continuity/*` or released MS | Republication path only |

When a gap appears: propose **A or B** first. Filenames: kebab-case in the work’s language; Non-Latin kept — no English alias files.

### Episode design ↔ architecture sync

1. Selective load (indexes → Appearing/used/cited)
2. Gap → classify A/B/C/D → architecture edit **before** more Key Events
3. User approves → record Prior Design Alignment + Architecture References
4. Re-run Design Consistency + Generation Readiness
5. C/D → Cascade checklist before continuing

### Cascade checklist

```markdown
## Consistency Cascade
- Change: {what}
- Type: Add / Extend / Modify / Retcon
- Prior episodes affected: {none | list + republication?}
- Highest file updated: {path} — user approved: ✅/⬜
- [ ] overview.md
- [ ] series.md
- [ ] world-bible.md / world/* / characters* / locations*
- [ ] stagings*
- [ ] episodes/{nnn}
- [ ] manuscripts/{nnn} (after design sync)
- [ ] continuity/* (⑧ or republication)
- Gates re-run: Design Consistency / Generation Readiness / Design-Fidelity — {which}
```

### What not to duplicate here

Full Design Consistency / Design-Fidelity / Floor **tables**, forecast math, Path existence, craft anti-patterns → stage workflows only.

---

## Reference Models (Identity + Staging Continuity)

**Not a numbered stage.** Prose design locks only — **no PNG / image YAML** in this skill.

### Three layers

| Layer | Locks | Artifact | When |
|-------|-------|----------|------|
| **Character** | Identity + wardrobe/equipment **states** | `characters/{slug}.md` | ③ (+ states at ④) |
| **Location** | Empty set (neutral, static); citeable **Multi-facet anchors** | `locations/{slug}.md` | ③ (+ facets at ④) |
| **Staging** | Who-is-where for one situation: cited states, seats, situation props, ambient, situation environment | `stagings/{slug}.md` | **④ only** |

Staging **cites** character states — it does not invent outfits/weapons. Same location ≠ same staging. Expression, one-off weather/camera → scene direction (or staging situation environment when the span must hold).

**Mandatory staging:** ≥2 scenes in the same situation **and** (≥2 named cast **or** fixture-relative placement). `Staging: none` only for single-scene / no fixed placement — fail **G4** if violated. Procedures, templates, situation examples: `04-episode-design.md` § Staging.

**Location cites:** catalog path only; every `+ facet` must be an exact **Multi-facet anchors** label (layout / Internal facets / Sensory prose alone is not citeable). Co-appearing major cast: distinguishable on ≥2 visual axes. Full state/facet rules: `03-architecture.md` + `04-episode-design.md`.

**Story design co-locks catalogs:** gap in cast/place/state/staging/facet → stop → add/extend architecture or staging → approve → resume. Prefer add over modify (§ Consistency). Do not “fix” a missing facet by parking the scene on a sibling facet already used by the previous scene.

---

## Revision — Cross-Stage Protocol

**Not a numbered stage.** After ⑤ or ⑦. Design before prose for plot/cast/place/world changes. Prefer additive architecture.

| Kind | First file |
|------|------------|
| **A** — Prose craft only | `manuscripts/` |
| **B** — Episode design | `episodes/{nnn}-*.md` (+ MS if exists) |
| **C** — Higher plot/list | `series.md` → design → MS |
| **D** — New lore/cast/place/facet/staging | Add architecture / staging → cite |
| **D′** — Modify lore/cast | Stop — Consistency C → cascade |
| **E** — Continuity / released | Prefer unreleased fix; else republication |
| **F** — Carry-⑥ only | No ④ edit — record for generation |
| **G** — Eval Skip | Status **Skip** — never leave Pending |

**Carry vs design (hard):**
- **Carry-⑥** = generation-time only (Hold, POV limits, do-not-narrate, motif-touch reminders). Does **not** rewrite design fields.
- Findings that need design-field changes (closing image/beat, crowded Out, adjacent scenes on the same Location facet, observable character cost, scene separation) → Status **Pending** (Kind **B**), not Carry-⑥.
- User says “apply recommendations / 권고 반영”: ask Kind **B** vs **F** in one line, **or** default Med+ design-field findings to **B**.
- After Kind **B**/**D** revision: **re-run** Schema + Consistency + Adjudication on the **revised** design body; mark applied rows **Applied-④**; sync design Gate **G4**, eval Gate **G5**, and `TODO.md` in the **same** approval turn. **Forbidden:** flip G4/G5 to Approved without rewriting the eval against the revised file.

**Cite-gap fix (hard):** Missing citeable Location `+ facet` → **stop → Extend Multi-facet anchors → approve → cite**. **Forbidden** as default: relocate a scene onto an already-used facet to avoid Extend — especially when that makes two **adjacent** scenes share the same facet.

**Procedure:** user selects findings → classify → highest layer first → cascade → re-run gates (design → Consistency + Readiness; MS → Design-Fidelity + Floor) → update the active evaluation file → present for approval.

Design updates: keep Scene Index ↔ `## Scenes` in sync; refresh Prior Design Alignment if higher files changed. Critique→fix mapping details: `05-evaluate-design.md` / `07-evaluate-manuscript.md` revision loops.

---

## Research — Independent Activity

**Not a fixed stage.** May run during ①–⑦. Save `{project-root}/docs/{topic-slug}.md`:

```markdown
# Research: {Topic}

## Summary
{Key findings}

## Sources
{Attribution}

## Application Notes
{Which artifact absorbs this — overview / series / world / character / location / episode design}
```

Applying research is not automatic: update the correct design/architecture file first, then cascade per § Consistency. Do not sneak new world rules only into manuscript prose.
