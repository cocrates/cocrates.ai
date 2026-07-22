# Stage ④ — Chapter Design

**Prerequisites:**
- Stage ③ architecture approved
- Parent `chapters/{nnn}-{chapter-slug}.md` chapter catalog approved
- Detail profiles exist for all characters/locations in this chapter
- **Chapter 002+:** Prior chapter released (⑧); continuity files loaded

**Gate artifacts:**
- `{project-root}/chapters/{nnn}-{chapter-slug}.md` — chapter-level design
- `{project-root}/chapters/{nnn}-{chapter-slug}/{nnn}-{episode-slug}.md` — one file per episode (scene-level design)

**Next stage:** `06-evaluation.md` (design evaluation, recommended) → `05-generation.md` (after user approves design)

**Also used by:** stages ⑤ (generation), ⑥ (evaluation), ⑦ (revision)

---

## Core Concept

Chapter design is **composition, not invention**. It arranges approved characters, locations, world rules, and continuity state into a **chapter arc** (chapter file) and **scene sequences** (nested episode files).

**Chapter is the publication unit.** Episodes subdivide the chapter for design granularity only — they are not separately published, evaluated as manuscripts, or released.

**Design is not the manuscript.** Key Events describe **what happens and why** — not the exact words characters will say. Dialogue, prose rhythm, and literary craft belong to stage ⑤ (generation).

**Episode count is decided here.** Based on this chapter's story needs, determine how many episodes subdivide the chapter — not predetermined in part or architecture files.

### Design Procedure

1. Expand `chapters/{nnn}-{chapter-slug}.md` from architecture catalog into full chapter design
2. **Create the Episode List** in the chapter file — as many episodes as the story composition requires
3. Create one episode file per row in the Episode List under `chapters/{nnn}-{chapter-slug}/`
4. Ensure episode designs connect — hooks, threads, and motifs flow across episodes within the chapter
5. Present chapter file + all episode files for user approval
6. **Recommend design evaluation** (`06-evaluation.md`, design mode) before manuscript generation

---

## Chapter File — Chapter-Level Design

The chapter file governs the **whole publication unit**:

```markdown
# Chapter {nnn}: {Title}

## Part
{parts/{nnn}-{part-slug}.md}

## Architecture References
| Type | Artifact | Usage |
|------|----------|-------|
| Part | `parts/{nnn}-{part-slug}.md` | Chapter List (Role + Hook to Next); part arc |
| Character | `characters/{name}.md` | |
| Location | `locations/{name}.md` | |
| World | `world-bible.md` / `world/{aspect}.md` | |

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
| Ep | File | Title | Function in Chapter Arc |
|----|------|-------|------------------------|
| 001 | `{nnn}-{episode-slug}.md` | {Title} | {one line} |
| 002 | ... | ... | ... |

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
```

---

## Episode Files — Scene-Level Design

Each episode file lives at `chapters/{nnn}-{chapter-slug}/{nnn}-{episode-slug}.md`.

Episode numbers are **local to the chapter** (001, 002, …).

```markdown
# Episode {nnn}: {Title}

## Chapter
`chapters/{nnn}-{chapter-slug}.md`

## Purpose
{One sentence — this episode's role in the chapter arc}

## Reader Engagement
### Personal Stake
### Opening Question (if episode opens chapter or new POV block)
### Reader Questions
### Stakes

## Literary Craft
### Motifs
| Motif | Meaning | Appearances (Key Event #) |

### POV Inserts (if any)
| Insert | After Key Event # | Reader learns (POV doesn't) |

### Sensory-Emotional Beats
| World/setting detail | POV reaction or feeling |

### Dialogue Voices
| Character | Speech pattern | Inner/outer gap (if POV) |

### Reader-Discovered Meaning
**What reader should conclude** (design only):
**Hold — do NOT write in manuscript:**
**Closing image:**

## Exposition Budget & Seeds
### Budget
### Seeds
| Element | Class | How it appears |
### Hold (do NOT include)

## Key Events
{Scene flow summaries — NOT dialogue or prose. See below.}

### Scene 1 — {title}
**POV:** | **Location:**

- **Situation:**
- **Beat:**
- **Turn:**
- **Function:**
- **Dialogue intent:** (optional — one phrase, no quotes)

## POV
## Hooks
## Estimated Length
## Characters Appearing
## Locations Used
## Tone Notes
```

---

## Key Events — Flow Summary, Not Dialogue

Scene design must make the **dramatic flow** clear — not pre-write the episode.

### What belongs in Key Events

| Include | Exclude |
|---------|---------|
| Scene purpose (dramatic function) | Full dialogue lines (`> Character: "..."`) |
| Who is present, where, POV | Paragraph-length prose |
| What **changes** in the scene | Exact wording of speeches |
| Turning point / emotional shift | Inner monologue draft text |
| Dialogue **intent** in one phrase | Multi-line quoted exchanges |

**Format per scene:**

```markdown
### Scene {n} — {title}
**POV:** {character} | **Location:** {where}

- **Situation:** {starting state — one line}
- **Beat:** {what happens — causal flow, not script}
- **Turn:** {what changes — character, relationship, stakes}
- **Function:** {why this scene exists in the episode arc}
```

Optional — when dialogue direction matters, use **intent only**:

```markdown
- **Dialogue intent:** Dong-hyeok requests a situation report in a professional tone (not casual friend speech)
```

### Rule

If a Key Event reads like it could be pasted into `manuscripts/`, it is **too detailed**. Cut dialogue and prose; keep causal flow and dramatic function.

**Stage ⑤** writes the actual lines — informed by `Dialogue Voices` and scene intent from design, not by pre-written quotes in Key Events.

---

## Mandatory Loading

### Architecture (always)

| Artifact | Verify |
|----------|--------|
| `chapters/{nnn}-{chapter-slug}.md` | Chapter catalog from ③; Episode List added in this stage |
| `parts/{nnn}-{part-slug}.md` | Chapter fits part arc |
| `characters.md` | Every appearing character listed |
| `characters/{name}.md` | Voice, motivation, arc — no redefinition |
| `locations.md` | Every used location listed |
| `locations/{name}.md` | Atmosphere, layout — no redefinition |
| `world-bible.md` | No contradiction with premise/laws |
| `world/{aspect}.md` | Referenced systems already defined |

### Continuity (chapter 002+)

| Artifact | Verify |
|----------|--------|
| `continuity/story-so-far.md` | Character/location states, timeline |
| `continuity/{prior-chapter-slug}-summary.md` | Prior hook addressed |
| `continuity/unresolved-threads.md` | Threads picked up, advanced, planted, or held |

**Chapter 001:** No continuity files — record what this chapter **establishes** for future continuity.

**Do not re-read prior manuscripts** — continuity files are authoritative.

### MAY / MUST NOT

**MAY:** Select approved characters/locations; apply world rules; plan scenes, hooks, engagement across chapter and episodes.

**MUST NOT:**
- Invent new named characters, locations, factions, or world rules
- Expand profiles beyond architecture files
- Contradict `story-so-far.md`
- Ignore active threads without Hold reason
- Override parent chapter catalog scope (role/conflict/Hook to Next in part Chapter List)
- Create top-level `episodes/` files — episodes nest under chapter folder only

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
- [ ] Scene-first Key Events in every episode file
- [ ] Architecture-bound
- [ ] Continuity-bound (ch 002+)
- [ ] No improvised entities
- [ ] **Motifs planned** — 1–2 motifs with placements across episodes
- [ ] **Sensory-emotional pairing** — world details filtered through POV reaction
- [ ] **Dialogue voices distinct**
- [ ] **POV inner/outer gap**
- [ ] **POV insert budget** — ≤ 2 per episode, placement planned
- [ ] **Reader-discovered meaning** — thematic conclusion in Hold
- [ ] **Antagonist plausibility**
- [ ] **Emotional indeterminacy**
- [ ] **Closing is image/silence**
- [ ] **Key Events are flow summaries** — no quoted dialogue
- [ ] **All episode files present** — one per Episode List row

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
- `chapters/{nnn}-{chapter-slug}.md`
- All files in `chapters/{nnn}-{chapter-slug}/`

Results go to `evaluations/{nnn}-{chapter-slug}.md`.

---

## Completeness Check

- [ ] Chapter file expanded from architecture catalog
- [ ] One episode file per Episode List row
- [ ] Architecture References complete (chapter file)
- [ ] Continuity References complete (ch 002+)
- [ ] Prior Hook documented (ch 002+)
- [ ] Threads This Chapter covers continuity
- [ ] States match `story-so-far.md`
- [ ] All entities in architecture — no improvisation
- [ ] Engagement Checklist passed
- [ ] Key Events are scene flow summaries in all episode files
- [ ] Episode hooks chain; chapter arc is coherent

---

## Gate

User approves chapter design (chapter file + all episode files). Do not proceed to `05-generation.md` without approval.

Design evaluation (`06-evaluation.md`) is strongly recommended between design approval and manuscript generation.
