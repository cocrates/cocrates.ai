# Stage ③ — Architecture Design

**Prerequisites:** Approved `overview.md` and `series.md`

**Gate artifacts:**
- Always: `characters.md`, `locations.md`, `world-bible.md`, `chapters/{nnn}-{chapter-slug}.md`
- When needed: `stagings.md` + `stagings/{slug}.md` (continuing-situation blocking)
- **Series mode only:** `parts/{nnn}-{part-slug}.md`

**Next stage:** `04-chapter-design.md` — only after architecture **and** parent chapter catalog are approved

---

## Procedure

Load `overview.md` and `series.md`. Respect **Structure Mode**.

| Mode | Chapter List source | `parts/` |
|------|---------------------|---------|
| **Short** | Already complete in `series.md` (stage ②) | **Do not create** |
| **Series** | Expand each Part Catalog row into `parts/{nnn}-{part-slug}.md` here | **Required** |

**Flexible counts:** Part and chapter counts from `overview.md` / `series.md` are planning targets. Adding, merging, or removing chapters is normal when story structure demands it — update the Chapter List source and `overview.md` Scale with user approval.

**Episodes are not planned here.** Do not specify episode counts. Episode Index and scenes are decided at **stage ④**.

Characters, locations, world-bible, and chapter catalogs may be designed **in parallel** once the Chapter List for the next chapters to design is known.

**“Early chapters”:** At minimum chapter 001 must have catalog + profiles for its cast/places. Expand catalogs and profiles for further chapters before designing them (typically a small upcoming window, e.g. next 1–3 chapters, or as the user requests).

---

## `characters.md` (Character Web)

```markdown
# Character Web: {Title}

## Factions & Groups
- {Faction/Group name}: {beliefs, goals, internal structure}

## Relationship Map
- {Character A} ↔ {Character B}: {relationship type, power dynamic, conflict seed}

## Character Roles
### {Character Name}
- **Role**: {protagonist / antagonist / ...}
- **Archetype**: {optional}
- **Core Drive**: {what they want most}
- **Central Conflict**: {what prevents them from getting it}
- **Arc Direction**: {positive / negative / flat — brief description}
```

Create `characters/{name-slug}.md` for each character who will appear in early chapters, using this structure:

```markdown
# {Character Name}

## Basic Info
- Role: {protagonist / antagonist / ...}
- Core Drive: {what they want most}
- Central Conflict: {what prevents them from getting it}
- Arc Direction: {positive / negative / flat — brief description}

## Appearance & Behavior Traits
- 외형(Appearance): {face/height/build/unique visual traits — fixed identity}
- 습관/거동(Behavior): {signature gestures, posture habits, movement style}
- 말투/표현(Voice): {speech patterns, recurring phrases, emotional tell}
- 반복되는 단서(Signature cues): {what readers notice consistently}

## Reference models (appearance + equipment states)
{Canonical: `workflow/reference-models.md` — static look for a story span; prose lock only (no PNG)}
- base: {concrete default outfit + identity gear — clothes, armor, weapons, shields, glasses, accessories}
- {state-slug}: {concrete lasting change: growth/accident aftermath OR outfit/gear; what stays vs differs}
Rules: document each state concretely; state = lasting physical/equipment identity only; expression/mood/transient pose are scene direction; do not silently swap identity gear across scenes.

## Key Relationships
- {Other Character A} ↔ {Character Name}: {relationship type, power dynamic, key tension seed}
- {Other Character B} ↔ {Character Name}: ...
```

---

## `locations.md` (Location Map)

```markdown
# Location Map: {Title}

## Regions & Territories
- {Region name}: {description, climate, culture, narrative role}

## Key Locations
- {Location name}: {description, significance, notable events}

## Spatial Hierarchy
{Tree or list showing containment relationships}
```

Create `locations/{name-slug}.md` for each location used in early chapters, using this structure:

```markdown
# {Location Name}

## Basic Info
- Type: {world / region / city / building / room / outdoor / ...}
- Narrative Role: {why this place matters to the story}
- Role as set/stage: {physical structure that must stay recognizable; lasting damage → new state — see `workflow/reference-models.md`}
- Key Events: {notable events that happen here}

## Spatial Composition (static / physical — empty set)
- Structure/Layout: {spatial layout — rooms, corridors, open areas, landmarks}
- Physical detail examples: {street → buildings, signs, traffic lights, crosswalks, exterior walls; café → seating, counter, fixed interior, props}
- Multi-facet anchors (when needed): {front/left/right or window-side/bed-side/door-side — do not pack the entire 3D space into one paragraph}
- Sense of Scale: {size impression — cramped / vast / intimate / etc.}

## Sensory Environment
- Lighting: {natural light / artificial / dim / harsh / flickering / ...}
- Temperature/Humidity: {cold / warm / humid / dry / seasonal shift}
- Smell: {dominant scent — dust, food, chemicals, nature, decay, ...}
- Sound/Noise: {ambient sound — silence, traffic, machinery, birdsong, echoes, ...}
- Texture/Touch: {what surfaces feel like if touched — rough stone, polished wood, ...}

## Atmosphere Notes
{psychological impression this place conveys — safety / threat / nostalgia / isolation / ...}
```

---

## `stagings.md` + `stagings/{slug}.md` (Continuing-situation blocking)

**When:** a situation continues across multiple scenes and relative placement, props, and/or ambient must not flip (café L/R, street walk-and-talk, OR stations, meeting seats, …). **One staging per situation.** Full rules: `workflow/reference-models.md` §4 / §4.1.

**Chapter/episode-first (preferred):** Define staging with the chapter that needs it. Same location ≠ reuse prior staging.

**Mandatory:** ≥2 scenes in the same situation **and** (≥2 named cast **or** fixture-relative placement) → staging required. `Staging: none` only for single-scene / no fixed placement.

```markdown
# Stagings

| Slug | Situation | Owning chapter | Location path | Cast | Span |
|------|-----------|----------------|---------------|------|------|
| {staging-slug} | {street walk / café call / …} | {nnn-…} | {loc[+facet]} | {A, B} | ch/ep/scene range |
```

Per staging file: catalog-bound location + character states, seat/spot + facing, props/ambient/situation environment when locked, continuity rules. No PNG — prose blocking map only.

Gate artifacts for this stage include `stagings.md` when any continuing situation needs blocking.

---

## `world-bible.md` (World Bible)

```markdown
# World Bible: {Title}

## Core Premise
## Physics & Natural Laws
## Magic / Technology / Supernatural Systems (Overview)
## Society & Culture (Overview)
## History (Timeline)
## Thematic Landscape
```

Create `world/{aspect-slug}.md` for domains referenced in early chapters.

---

## `parts/{nnn}-{part-slug}.md` (Series mode only)

**Where chapter planning happens in series mode.** One file per row in `series.md` Part Catalog.

**Short mode: skip this entire section** — Chapter List already lives in `series.md`.

Prerequisite: Part appears in `series.md` Part Catalog.

Lists every planned chapter in the part. Chapter count may differ from the approximate count in Part Catalog — update Part Catalog and `overview.md` Scale when it does.

**No episode column** — episode structure is decided per chapter at stage ④.

**Chapter List is the flow map** — each row's Role and Hook to Next let you review momentum across the part before chapter detail design. Read the Hook column top-to-bottom: consecutive chapters should connect; gaps or dead ends need fixing here.

```markdown
# Part {nnn}: {Title}

## Role in Series Arc
## Core Conflict

## Chapter List
| Ch | Title | Role | Key POV | Hook to Next |
|----|-------|------|---------|--------------|
| 001 | {Title} | {one-line role} | {character} | {question / tension / unanswered beat that carries into Ch 002} |
| 002 | {Title} | ... | ... | {→ Ch 003} |
| {last} | {Title} | ... | ... | {→ next part, or closing beat if series end} |

## Part Arc
## Part Hooks
- **Opening hook** (chapter {first}):
- **Closing hook** (chapter {last}):

## Character Arcs Advanced
```

**Hook to Next vs Part Hooks / chapter catalog Closing hook:**
- **Hook to Next** (Chapter List column) — handoff *between* chapters; required for every row.
- **Part Hooks** — part entry/exit only (first chapter open, last chapter close).
- **Chapter catalog / stage ④ Closing hook** — refine the row's Hook to Next when designing that chapter; keep them aligned.

---

## `chapters/{nnn}-{chapter-slug}.md` (Chapter Catalog)

**Bridge between Chapter List and chapter detail design.** Catalog-level chapter info only — **no** Episode Index or Key Events here.

Prerequisite: Chapter appears in the mode's Chapter List source:
- **Short:** `series.md` → Chapter List
- **Series:** parent `parts/{nnn}-{part-slug}.md` → Chapter List

**Catalog vs design:** Stage ③ chapter catalog has Role / Hooks only — **no** Episode Index or Key Events. Stage ④ expands the same path into manuscript-ready chapter + episode + scene design. Do not leave stage ④ at Episode Index alone.


```markdown
# Chapter {nnn}: {Title}

## Part
{Short: part id/title from series.md Part Composition | Series: parts/{nnn}-{part-slug}.md}

## Role in Arc
## Core Conflict

## Chapter Arc
{Brief — how this chapter fits the part / work arc}

## Chapter Hooks
- **Opening hook:**
- **Closing hook:**

## Character Arcs Advanced
```

---

## Adjusting Chapter Counts

When architecture (or later revision) changes planned structure, follow [`consistency.md`](consistency.md):

1. Propose change to user with reason
2. Update Chapter List source first:
   - **Short:** `series.md` Chapter List (+ Part Composition counts)
   - **Series:** `parts/{nnn}-{part-slug}.md` Chapter List + `series.md` Part Catalog approximate counts
3. Update matching `chapters/{nnn}-*.md` catalogs (and full designs if they exist)
4. Update `overview.md` Scale if total shifts significantly
5. If short work grows past ~20 chapters or needs multiple major parts → migrate to **series** mode
6. Resume — estimates need not match the original numbers

---

## Completeness Check

- [ ] Every character in early chapters has `characters.md` entry + profile
- [ ] Every location in early chapters has `locations.md` entry + profile
- [ ] Chapter List source is complete with Role + Hook to Next; consecutive hooks connect
- [ ] Each early chapter has a catalog file under `chapters/`
- [ ] No Episode Index / episode sections or episode count in Chapter List or chapter catalog files (those are stage ④)
- [ ] **Short:** no `parts/` directory
- [ ] **Series:** each Part Catalog row has a `parts/` file; no chapter detail duplicated in `series.md`

---

## Gate

User approves architecture files (individually or as batch when requested).

**Do not begin chapter detail design until the relevant chapter catalog is approved.**
