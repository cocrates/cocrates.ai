# Stage ④ — Chapter Design

**Prerequisites:**
- Stage ③ architecture approved
- Parent `chapters/{nnn}-{chapter-slug}.md` chapter catalog approved
- Detail profiles exist for all characters/locations in this chapter
- **Chapter 002+:** Prior chapter released (⑧); continuity files loaded

**Gate artifact:** `{project-root}/chapters/{nnn}-{chapter-slug}.md` — **one file** holding chapter-level design **and** all episode/scene design

**Next stage:** `06-evaluation.md` (design evaluation, recommended) → `05-generation.md` (after user approves design)

**Also used by:** stages ⑤ (generation), ⑥ (evaluation), ⑦ (revision)

**Cross-stage sync:** [`consistency.md`](consistency.md) — load before drafting; cascade when higher artifacts change.

---

## Core Concept

Chapter design is **composition, not invention**. It arranges approved characters, locations, world rules, and continuity state into a **chapter arc** and **episode → scene sequences** — all in **one chapter design file**.

**Faithfulness (hard rule):** Chapter design is bound to already-approved higher-level artifacts. Before writing or expanding episode/scene content, **load and use** series/part plans, world bible, character and location profiles, and continuity (ch 002+). Do not redesign the work from memory. If the emerging chapter design **conflicts** with those sources, stop — fix by updating the higher-level artifact (with user approval) or revise the chapter plan to fit; never silently override.

**Manuscript-ready design (hard rule):** By the end of stage ④, `chapters/{nnn}-{chapter-slug}.md` must contain everything stage ⑤ needs to write prose **without inventing plot beats, scene order, POV, location, seeds, or dialogue intent**. If a writer (or agent) would have to guess what happens next, the design is incomplete — do not proceed to generation.

**Chapter is the publication unit.** Episodes and scenes are design subdivisions only — not separately published, evaluated as manuscripts, or released. **Scenes appear only under episode Key Events.**

**Single-file rule:** Do **not** create `chapters/{nnn}-{chapter-slug}/` nested episode files. Episode and scene design live as sections inside `chapters/{nnn}-{chapter-slug}.md`.

**Anti-duplication:** Chapter-level sections own *chapter-wide* plans (arc, seeds table, motifs, hooks). Episode sections own *execution* (purpose, craft, Key Events/scenes). Do **not** add an "Episode Detail" / prose summary block that repeats what Key Events already specify. Episode Index is TOC only.

**Design is not the manuscript.** Key Events specify **concrete dramatic flow** (who, where, when, what changes, sensory/emotional cue, dialogue *intent*) — not quoted speech or publishable prose. Stage ⑤ writes the actual lines.

**Episode count is decided here.** Based on this chapter's story needs — not predetermined in `series.md`, part files, or architecture catalogs.

Load Structure Mode from `overview.md` / `series.md`. Chapter List source for Architecture References:
- **Short:** `series.md`
- **Series:** `parts/{nnn}-{part-slug}.md`

### Design Procedure

1. **Load prior design** — complete the Pre-Design Load Checklist (below). Do not skip; do not rely on chat memory alone
2. Expand `chapters/{nnn}-{chapter-slug}.md` from architecture catalog into full chapter-level sections, including **Prior Design Alignment**
3. Draft Purpose / Chapter Arc / Episode Index so they **match** Chapter List Role, Hook to Next, and (series mode) part arc — record any needed higher-level change as a stop/ask, not a silent drift
4. Under `## Episodes`, write one **complete** `### Episode {nnn}: {Title}` section per Index row — including every scene under Key Events with the mandatory scene fields below
5. Walk episode hooks and scene transitions top-to-bottom: no orphan scenes, no unexplained jumps
6. Run **Prior-Design Consistency Gate**, then **Manuscript Readiness** (below)
7. Present the **single** chapter design file for user approval only if both gates pass — include a one-line status: *"Load ✅ · Consistency ✅ · Readiness ✅"*
8. **Recommend design evaluation** (`06-evaluation.md`, design mode) before manuscript generation — eval must re-check prior-design consistency

**Legacy note:** Older projects may still have `chapters/{slug}/*.md` episode files. Prefer migrating into the parent chapter.md; until then, treat nested files as authoritative for Key Events and strip duplicated Episode Detail from chapter.md. New work must use the single-file form.


---

## Chapter File Template

```markdown
# Chapter {nnn}: {Title}

## Part
{Short: part id/title from `series.md` Part Composition | Series: `parts/{nnn}-{part-slug}.md`}

## Architecture References
| Type | Artifact | Usage in this chapter |
|------|----------|----------------------|
| Overview | `overview.md` | Genre/tone/validation — do not contradict |
| Series | `series.md` | Arc, Structure Mode, voice; **Short:** Chapter List Role + Hook to Next |
| Part / Chapter List | **Short:** `series.md`. **Series:** `parts/{nnn}-{part-slug}.md` | Role, conflict, hooks, part/work arc |
| Character | `characters.md` + `characters/{name}.md` | Appearing cast only — voice/drive/arc as approved |
| Location | `locations.md` + `locations/{name}.md` | Used places only — atmosphere/layout as approved |
| World | `world-bible.md` / `world/{aspect}.md` | Laws, history, systems touched this chapter |

## Prior Design Alignment
{Filled after Pre-Design Load. Required before episode drafting.}

### Load confirmation
- [ ] `overview.md`
- [ ] `series.md` (+ Structure Mode)
- [ ] Chapter List source (short: `series.md` / series: parent `parts/…`)
- [ ] `characters.md` + profiles for cast in this chapter
- [ ] `locations.md` + profiles for places in this chapter
- [ ] `world-bible.md` (+ relevant `world/` aspects)
- [ ] Continuity set (ch 002+: `story-so-far`, prior summary, `unresolved-threads`)

### Bound to higher design
| Source constraint | How this chapter honors it |
|-------------------|----------------------------|
| Chapter List Role / Hook to Next | |
| Part / series arc (if applicable) | |
| Character drives / arcs touched | |
| Location / world rules used | |
| Continuity (prior hook + states) | {Ch 001: what this establishes} |

### Conflicts / open questions
{None | list. Any real conflict → stop; escalate to user; update architecture/continuity source first if changing lore.}

## Continuity References
{Ch 001: what this establishes}
{Ch 002+: mandatory}

| Source | What carries forward |
|--------|---------------------|
| `continuity/{nnn-1}-{chapter-slug}-summary.md` | |
| `continuity/story-so-far.md` | |
| `continuity/unresolved-threads.md` | |

### Prior Hook
{Ch 002+ required — from prior chapter closing}

### Threads This Chapter
- **Picks up**:
- **Advances**:
- **Plants**:
- **Holds**:

## Purpose
{One sentence — what this chapter accomplishes}

## Reader Engagement (Chapter Level)
### Personal Stake
### Opening Question
{Question the reader carries past the chapter opening — not answered immediately}
### Reader Questions
### Stakes

## Chapter Arc
{How tension builds and resolves across all episodes in this chapter}

## Episode Index
{TOC only — do not paste scene summaries here.}

| Ep | Title | Function in Chapter Arc |
|----|-------|------------------------|
| 001 | {Title} | {one line} |
| 002 | ... | ... |

## Chapter Hooks
- **Opening hook** (episode 001):
- **Closing hook** (last episode):

## Exposition Budget & Seeds (Chapter Level)
### Budget
### Seeds
| Element | Class | Episode | How it appears |
### Hold (do NOT include)

## Literary Craft (Chapter Level)
### Motifs (chapter-wide)
| Motif | Meaning | Episodes / Scenes |
### Reader-Discovered Meaning
**What reader should conclude** (design only):
**Hold — do NOT write in manuscript:**
**Closing image** (chapter end — no thematic monologue):

## Characters Appearing
## Locations Used
## Tone Notes
## Estimated Length
{Total chapter word count target}

---

## Episodes

### Episode 001: {Title}

#### Purpose
{One sentence — this episode's role in the chapter arc}

#### Reader Engagement
##### Personal Stake
##### Opening Question (if episode opens chapter or new POV block)
##### Reader Questions
##### Stakes

#### Literary Craft
##### Motifs
| Motif | Meaning | Appearances (Scene #) |

##### POV Inserts (if any)
| Insert | After Scene # | Reader learns (POV doesn't) |

##### Dialogue Voices
| Character | Speech pattern | Inner/outer gap (if POV) |

##### Reader-Discovered Meaning
**What reader should conclude** (design only):
**Hold — do NOT write in manuscript:**
**Closing image:**

#### Exposition Budget & Seeds
##### Budget
##### Seeds
| Element | Class | Scene # | How it appears |
##### Hold (do NOT include)

#### Key Events
{Every scene for this episode — mandatory fields per scene. See Key Events rules.}

##### Scene 1 — {title}
**POV:** {character} | **Location:** {architecture location} | **When:** {story time or relative (“same evening”)}
**On stage:** {characters present}

- **Situation:** {starting state — one line}
- **Beat:** {what happens — causal flow, concrete actions; not script}
- **Turn:** {what changes — character, relationship, stakes, or knowledge}
- **Function:** {why this scene exists in the episode arc}
- **Sensory-emotional:** {one world/setting detail} → {POV reaction or feeling}
- **Dialogue intent:** {required if anyone speaks — goal + tone; no quotes. Else: none}
- **Transition out:** {how this scene hands off to the next — cut / time skip / follow character / hook line intent}
- **Est. length:** {approx. words or characters for this scene}

##### Scene 2 — {title}
{…same fields…}

#### POV
{Primary POV for the episode}

#### Hooks
- **In:** {from prior episode / chapter opening}
- **Out:** {into next episode / chapter close}

#### Estimated Length
{Episode total — sum of scene Est. length; must fit chapter Estimated Length}

#### Characters Appearing
#### Locations Used
#### Tone Notes

### Episode 002: {Title}
{Same subsection pattern — every episode fully specified before approval}
```

---

## Key Events — Manuscript-Ready Scene Specs

Each scene under an episode is a **generation brief**. Stage ⑤ must be able to write the scene from these fields alone (plus Architecture References), without inventing missing beats.

### Required vs excluded

| Required | Exclude |
|----------|---------|
| POV, Location (from architecture), When, On stage | Full dialogue lines (`> Character: "..."`) |
| Situation → Beat → Turn → Function | Paragraph-length draft prose |
| Sensory-emotional pair (world detail → POV reaction) | Exact wording of speeches |
| Dialogue intent when speech occurs (goal + tone, no quotes) | Inner-monologue draft text |
| Transition out | Multi-line quoted exchanges |
| Est. length | Vague “they talk / something happens” |

### Format per scene

```markdown
##### Scene {n} — {title}
**POV:** {character} | **Location:** {where} | **When:** {time / relative}
**On stage:** {who is present}

- **Situation:** {starting state — one line}
- **Beat:** {concrete causal flow — what physically/dramatically happens}
- **Turn:** {what changes}
- **Function:** {why this scene exists in the episode arc}
- **Sensory-emotional:** {detail} → {POV reaction}
- **Dialogue intent:** {goal + tone, or none}
- **Transition out:** {handoff to next scene}
- **Est. length:** {approx. size}
```

Optional only when needed:

```markdown
- **Seed touch:** {Plant/Hint element id from episode Seeds — how it surfaces in this scene}
- **Motif touch:** {motif id — brief placement note}
```

### Completeness rules

1. **No empty Key Events.** Every episode section has ≥ 1 scene with all required fields filled.
2. **Causal chain.** Scene N's Transition out makes Scene N+1's Situation intelligible; episode Out hook matches the next episode's In (or chapter Closing hook).
3. **Concrete Beat.** A Beat must name actions or observable events — not only mood (“tension rises”).
4. **Still not prose.** If a Key Event could be pasted into `manuscripts/` as finished text, it is too dialogue-/prose-heavy — cut wording; keep intent and flow.
5. **Length budget.** Sum of scene Est. length ≈ episode Estimated Length ≈ share of chapter Estimated Length.

**Stage ⑤** writes the actual lines from Dialogue Voices, Sensory-emotional cues, and scene intent — not from pre-written quotes in Key Events.

---

## Pre-Design Load Checklist (mandatory — before composing)

Read artifacts from disk at the start of stage ④ (and again if resuming after a long gap). **Do not rely on prior chat alone.**

| # | Artifact | Always / When | Confirm |
|---|----------|---------------|---------|
| 1 | `overview.md` | Always | Genre, tone, validation, Structure Mode, scale |
| 2 | `series.md` | Always | Series arc, voice; Part Composition or Part Catalog |
| 3 | Chapter List source | Always | **Short:** `series.md` Chapter List row for this ch. **Series:** `parts/{nnn}-{part-slug}.md` — Role, Key POV, Hook to Next, part arc |
| 4 | This chapter's catalog | Always | Existing `chapters/{nnn}-{chapter-slug}.md` from ③ — Role / Hooks / conflict |
| 5 | `characters.md` | Always | Cast allowed; relationship map |
| 6 | `characters/{name}.md` | Every character who will appear | Drive, voice, arc — no redefinition |
| 7 | `locations.md` | Always | Places allowed |
| 8 | `locations/{name}.md` | Every location used in scenes | Atmosphere, layout |
| 9 | `world-bible.md` | Always | Premise, laws, history, society — no contradiction |
| 10 | `world/{aspect}.md` | Any system/history slice this chapter uses | Already defined |
| 11 | `continuity/story-so-far.md` | Ch 002+ | Character/location states, timeline |
| 12 | `continuity/{prior}-summary.md` | Ch 002+ | Prior closing; events that must not be undone |
| 13 | `continuity/unresolved-threads.md` | Ch 002+ | Threads to pick up / advance / plant / hold |

**Chapter 001:** Skip continuity files; after design, document what this chapter **establishes** under Continuity References / Prior Design Alignment.

**Do not re-read prior manuscripts** — continuity files are authoritative for past plot.

Record Load confirmation under **Prior Design Alignment** in the chapter file before writing Key Events.

---

## Prior-Design Consistency Gate (mandatory — before approval)

Run after episode/scene design is drafted. Any ❌ blocks approval and stage ⑤.

| Check | Pass criteria |
|-------|----------------|
| Loaded all required artifacts | Pre-Design Load Checklist complete; Reflecting in Prior Design Alignment |
| Chapter List Role | Purpose + Chapter Arc execute the approved Role — not a different story |
| Hook to Next / Closing | Chapter Closing hook aligns with Chapter List Hook to Next (or next-part handoff); refine wording but do not drop the obligation |
| Opening honors prior | Ch 002+: Prior Hook addressed in ep 001 scenes; no soft-reset of prior chapter |
| Continuity states | Ch 002+: Character/location states in scenes match `story-so-far`; no resurrection/undo without design-approved continuity update |
| Unresolved threads | Each active thread is Picks up / Advances / Plants / Holds — none silently dropped |
| Characters | Every On-stage name exists in architecture; voice/drive matches profile; arc direction not reversed without architecture update |
| Locations | Every scene Location is in architecture; use matches profile |
| World | Beats do not break physics/laws/history in world-bible; exposition only for Plant/Hint (not Hold lore dumps that contradict Seed plan) |
| Series / part arc | Chapter advances the approved series/part role; no stealing later-part payoffs marked Hold elsewhere |
| Tone / voice | Matches `series.md` Narrative Voice & Style and overview Genre & Tone |
| No silent lore invention | New factions/places/rules → stop and update stage ③ first |

**On conflict:** Prefer correcting the chapter design to fit approved sources. If the chapter reveals a necessary lore/plot change, propose architecture or continuity update → user approves → then finish chapter design.

---

## Mandatory Loading (summary)

| Layer | Artifacts |
|-------|-----------|
| Scope / series | `overview.md`, `series.md` |
| Plot parent | Chapter List source + chapter catalog |
| World / cast / space | `world-bible.md`, `world/`, `characters*`, `locations*` |
| Continuity (002+) | `story-so-far`, prior summary, `unresolved-threads` |

### MAY / MUST NOT

**MAY:** Select approved characters/locations; apply world rules; plan scenes, hooks, engagement across chapter and episodes in the same file — within Chapter List Role and continuity.

**MUST NOT:**
- Start episode/scene drafting before Pre-Design Load Checklist is done
- Invent new named characters, locations, factions, or world rules
- Expand profiles beyond architecture files
- Contradict `story-so-far.md` or ignore active threads without Hold reason
- Override parent Chapter List scope (role/conflict/Hook to Next)
- Contradict series/part arc, world-bible, or character/location profiles without an explicit approved higher-level update
- Create nested `chapters/{nnn}-{chapter-slug}/` episode files, or a top-level `episodes/` directory
- Duplicate episode plot as chapter-level "Episode Detail" when Key Events already exist
- Approve or hand off to stage ⑤ when Prior-Design Consistency Gate or Manuscript Readiness fails
- Episode Index–only chapter designs (TOC without Key Events)

---

## New Character or Location Needed?

1. **Stop** chapter design
2. **Propose to user** with reason and role (recurring / one-scene)
3. **Wait** for decision
4. If approved → update `characters.md` / `locations.md` + profile → user approves → resume design

Same for new world rules → `world-bible.md` or `world/{aspect}.md` first.

---

## Reader Engagement Design

### Principles

1. **Open with tension, not exposition**
2. **End with forward momentum** (chapter closing hook)
3. **Curiosity gaps** — plant questions; delay answers
4. **Stakes escalation**
5. **Emotional payoff**
6. **Chapter-to-chapter contract** — prior hook addressed
7. **Episode-to-episode flow** — hooks connect within chapter
8. **Personal stake over world lecture**
9. **Restraint earns attention** — hold material back

### Literary Craft (beyond engagement)

Engagement rules prevent info-dumps; **literary craft** prevents "safe but flat" prose. Design at chapter level (overall arc, closing image) and episode level (scene execution); execute in stage ⑤.

#### Opening Question (not just opening action)

Action-first is necessary but not sufficient. The chapter opening must plant a **question the reader carries past the first page** — one that is **not immediately answered**.

**Design rule:** Write the **Opening Question** in the chapter file — separate from the opening image. Episode 001 inherits or refines it.

#### Information vs Literary Tension

Especially chapter 001: world-building is unavoidable, but **info-heavy prose reads as briefing**. Target balance:

| Chapter position | Info : Tension (guideline) |
|------------------|---------------------------|
| **001** | ~50 : 50 (not 70 : 30) |
| **002–005** | ~40 : 60 |
| **Mid-part+** | Payoffs may shift toward revelation |

#### World Through Character

Every world-building beat must pass through **POV reaction** — one line does two jobs (world + character).

**Design rule:** In Key Events, pair each world detail with **POV attitude** — not a separate exposition block.

#### Motif Planning

Strong chapters thread **1–2 motifs** across episodes and scenes. Plan placements in chapter design — not discovered during writing.

#### Dialogue: Voice + Inner/Outer Gap

Each speaking character needs a **distinct speech pattern** in design. POV character needs **gap between said and felt**.

#### POV Insert Budget

Secondary viewpoints (robot logs, documents, news tickers) create **reader-only knowledge** — use sparingly.

| Guideline | Rationale |
|-----------|-----------|
| **1–2 inserts per episode** max | More breaks POV immersion |
| **Place at turning points** | After first meaningful contact, or closing |
| **Plan in design** | List each insert: placement, what reader learns that POV doesn't |

#### Reader-Discovered Meaning (show less, imply more)

The strongest fiction lets **the reader reach conclusions themselves**. When the narrator states theme, moral judgment, or emotional verdict, the reader has nothing left to do.

**Core rule:** In design, write **What the reader should conclude** — not **What the narrator will say**. If the design lists a thematic line for the closing, move it to Hold and specify a closing **image** instead.

#### Antagonist & Conflict — Avoid Moral Caricature

Antagonists should embody a **plausible worldview** the POV rejects — not confirm the reader's moral judgment.

#### Closing — Scene Over Statement

The chapter closing must **not** restate what the chapter already showed. **Design rule:** Closing Key Event = **image or action only**.

---

### Exposition Budget

| Position | Budget | Guideline |
|----------|--------|-----------|
| **Ch 001** | Very low | One world fact + one personal stake + one mystery hint |
| **Ch 002–005** | Low | 1–2 new concepts per chapter via scene |
| **Mid-part** | Medium | Deeper exploration of planted seeds |
| **Part/climax chapter** | Higher | Payoffs earned by prior restraint |

### Seed Discipline

| Class | Meaning |
|-------|---------|
| **Plant** | Show in scene this chapter |
| **Hint** | Foreshadow only — reader notices, doesn't understand yet |
| **Hold** | Do NOT include — reserved for later |

**Rule:** Pick 1–2 Plant, 1–2 Hint per chapter, explicitly list Hold. If design lists 5+ "must perceive" elements, manuscript will over-explain.

### Engagement Checklist

- [ ] Chapter Opening Question — reader carries unanswered question past page 1
- [ ] Personal stake (chapter level)
- [ ] Chapter arc connects all episodes
- [ ] Episode hooks chain within chapter
- [ ] Chapter closing hook / forward momentum
- [ ] Prior hook addressed (ch 002+)
- [ ] Exposition budget respected (chapter level)
- [ ] Seed discipline (Plant/Hint/Hold)
- [ ] Scene-first Key Events under every episode section — **all mandatory scene fields filled**
- [ ] Architecture-bound — Prior Design Alignment filled; Consistency Gate green
- [ ] Continuity-bound (ch 002+) — states, prior hook, threads accounted for
- [ ] No improvised entities; no silent override of series/part/world/cast/place
- [ ] **Motifs planned** — 1–2 motifs with placements across episodes/scenes
- [ ] **Sensory-emotional** on every scene
- [ ] **Dialogue voices** distinct; Dialogue intent on scenes with speech
- [ ] **POV inner/outer gap**
- [ ] **POV insert budget** — ≤ 2 per episode, placement planned
- [ ] **Reader-discovered meaning** — thematic conclusion in Hold
- [ ] **Antagonist plausibility**
- [ ] **Emotional indeterminacy**
- [ ] **Closing is image/silence**
- [ ] **Key Events are briefs, not prose** — no quoted dialogue
- [ ] **Every Episode Index row has a full episode section** with Key Events
- [ ] **Scene transitions / episode hooks** chain without gaps
- [ ] **Est. length** filled per scene and summed at episode + chapter

### Over-Seeded Opener Anti-Pattern

| Problem | Fix |
|---------|-----|
| 5+ tension elements in ch 001 design | Distribute across chapters |
| Trauma lore in dialogue (ch 001) | Hold; close with data gap only |
| Political exposition dialogue block | One fact per exchange |
| Date + world essay opening | Specific image/action |
| No personal stake | POV family, job, fear |
| Thematic summary ending | Concrete closing image |
| Key Events with full dialogue | Flow summary + dialogue intent only |

**Rule:** If improved manuscript is mostly *cuts*, the problem was in **design**. If mostly *rewording*, the problem was **literary craft** — update design and regenerate.

---

## Design Evaluation (Recommended)

After user approves chapter design, **propose design evaluation** before manuscript generation:

> *"Chapter {nnn} design is approved. Shall we evaluate the design before writing the manuscript? Design evaluation catches structural issues more efficiently than post-prose revision."*

See `06-evaluation.md` — **Design Evaluation** mode. Evaluator reads:
- `chapters/{nnn}-{chapter-slug}.md` (chapter-level + episode sections + Prior Design Alignment)
- Prior-design set again (`overview`, `series`, Chapter List source, world, cast, places, continuity)

Results go to `evaluations/{nnn}-{chapter-slug}.md`.

---

## Completeness Check

- [ ] **Pre-Design Load Checklist** completed; Prior Design Alignment filled
- [ ] **Prior-Design Consistency Gate** all pass (no unresolved conflicts)
- [ ] Chapter file expanded from architecture catalog
- [ ] Episode Index + one full episode section per Index row (in the same file)
- [ ] No parallel episode summary / Episode Detail that duplicates Key Events
- [ ] No nested `chapters/{nnn}-{chapter-slug}/` episode files
- [ ] Architecture References complete (overview, series, Chapter List source, cast, places, world)
- [ ] Continuity References complete (ch 002+)
- [ ] Prior Hook documented (ch 002+)
- [ ] Threads This Chapter covers continuity
- [ ] States match `story-so-far.md`
- [ ] All entities in architecture — no improvisation
- [ ] Engagement Checklist passed
- [ ] **Manuscript Readiness** passed (below)
- [ ] Episode hooks chain; chapter arc is coherent

### Manuscript Readiness (required before stage ⑤)

Stage ⑤ must **not** invent plot. Confirm:

- [ ] Every episode has ≥ 1 scene under Key Events
- [ ] Every scene has: POV | Location | When | On stage | Situation | Beat | Turn | Function | Sensory-emotional | Dialogue intent (or none) | Transition out | Est. length
- [ ] Beats are concrete (actions/events), not mood-only
- [ ] Scene Transition out → next Situation is intelligible; last episode Out matches Chapter Closing hook
- [ ] Seeds/Hold and motifs are placeable from scene fields (or optional Seed/Motif touch)
- [ ] Dialogue Voices cover every speaking character in On stage lists
- [ ] A cold reader could draft the chapter from this file + architecture refs alone

---

## Gate

User approves the single chapter design file **only if Prior-Design Consistency Gate and Manuscript Readiness both pass**. Do not proceed to `05-generation.md` without approval.

Design evaluation (`06-evaluation.md`) is strongly recommended between design approval and manuscript generation — include prior-design consistency again.
