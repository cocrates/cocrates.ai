# Stage ② — Design (World/Characters/Locations + Series/Episode Design)

**Prerequisites:** Approved `overview.md`

**Gate artifacts (design layer):**
- `{project-root}/illustration-guide.md` (**series-wide** art / layout / typography lock — see `workflow/illustration-guide.md`)
- `{project-root}/world-bible.md`
- `{project-root}/characters.md` + `characters/{character-slug}.md`
- `{project-root}/locations.md` + `locations/{location-slug}.md`
- `{project-root}/series.md`
- `episodes/{nnn}-{episode-slug}.md` (episode files include page design)

**Next stage:** `03-evaluate.md` (after user approves design gates)

---

## Procedure

### 2.0 Design constraints (global)

- **Architecture-first**: Image generation only in Stage ④.
- **Consistency sources**:
  - **Visual / typography:** `illustration-guide.md` (art style, layout, fonts, text boxes — series-wide; `workflow/illustration-guide.md`)
  - Character consistency: `characters/{character-slug}.md` (base appearance / face / silhouette + reference model list)
  - Location consistency: `locations/{location-slug}.md` (reference models center on `state` change; `position` · `view` are framing anchors)
  - Story / direction consistency: `episodes/{nnn}-{episode-slug}.md` (per-page rendering text + Direction / Picture carries — **not** fonts or series style)
- **Design-first fixes**: If Evaluate flags a problem, do not fix by “image prompt only” — update Design files first, then re-evaluate.
- **Do not invent per-page fonts or text-box styles.** Page YAML may only place locked text; style tokens come from `illustration-guide.md`.

#### What is a Reference Model?

**Canonical rules:** `workflow/reference-models.md` (same three layers as webtoon / video / novel).

**Reference model = durable physical appearance/structure across the story + locked relative placement for continuing situations.**
In Stage ④ these become “reference images” via the image-generation skill; every later page image builds consistency from them.

| Kind | **Included** in reference model (= generate as reference images) | **Excluded** (= direct in page illustration guide) |
|------|------|------|
| **Character** | Face / body / silhouette, **lasting body change**, **equipment identity** (clothes, accessories, weapons, shields — do not swap to a different sword page-to-page) | Expression, pose, action, emotion |
| **Location** | Physical structure / layout / fixed set dressing of the **set / stage** | Lighting, time of day, season, weather, mood, camera direction / framing / angle, transient elements |
| **Staging** | **Who is where** + **situation props / table dressing** in a continuing situation (meal, café, picnic, play, …). One per situation; **2–3** reference images | Expression / mood; camera tightness while seats **and** locked props stay recognizable |

**When a reference model change is required (= new reference image):**
- Character: outfit/gear changes, or lasting body change → add a new state slug
- Location: physical structure changes permanently (e.g. new building outside the window, wall destroyed) → add a new state slug
- Staging: seats/formation change, **situation props swap** (new meal layout, cleared table), or a new situation starts → new staging

**When a reference model change is not required (= handle in page prompt):**
- Expression, emotion, transient pose
- Temporary / reversible changes (curtains open/closed, door open/shut)
- Lighting / time / weather / mood shifts
- Transient elements (passerby, bird flying by)

**When staging is required:** The same situation spans multiple pages and relative positions **or situation props** must not flip/teleport (café facing seats, **family dinner food layout**, picnic spread, play-mat toys, OR, meeting, …). Write `stagings.md` + `stagings/{slug}.md` **including Situation props / table dressing**, then generate 2–3 ensemble refs in Phase 0. Character + empty location alone do not lock seats / L-R / **tonight’s meal**.

---

### 2.0b Illustration Guide (series-wide)

1. Expand `overview.md` → Illustration Style into `{project-root}/illustration-guide.md`.
2. Follow the full template and rules in `workflow/illustration-guide.md` (art style, layout, typography roles, text boxes, Stage ④ checklist).
3. Do not proceed to episode page design without a draft guide strong enough that two pages would share fonts/boxes/line quality.
4. G2 includes explicit approval of `illustration-guide.md`.

---

### 2.1 World Design

1. Write `world-bible.md` using the following structure:

```markdown
# World Bible

## World Name
{name}

## World Description
{description of the world — rules, atmosphere, key features}

## Core Rules
- {rule 1}
- {rule 2}

## History / Background
{background story of the world}
```

---

### 2.2 Character Catalog Design

1. Create `characters.md` as an index (table only).
2. For each character, create `characters/{character-slug}.md` using this structure:

```markdown
# {Character Name}

## Basic Info
- Role (lead / supporting / antagonist / …): {role}
- Core Drive: {what they want most}
- Central Conflict: {what prevents them from getting it}

## Appearance & Behavior
- Face / body / silhouette “must-not-change” anchors: {describe in detail}
- Signature gestures / tics / movement patterns: {tics, posture habits, typical movement}
- Recurring expression / speech feel (character voice): {expression + speaking style notes}

## Relationships
- {other-character-slug}: {relationship type, power dynamic, conflict seed}
- {other-character-slug}: ...

## Reference Models (appearance change + state list)
{Each state maps to one reference image — generated via image-generation in Stage ④ Phase 0}

- base: {definition of default outfit / gear}
- {state-slug}: {what changes (outfit / gear / carried items) + face/body rules stay}
- {state-slug}: ...

Reference model rules:
- state = lasting body change + equipment identity (outfit / accessories / weapons / shields / carried items) — do not swap identity gear page-to-page
- Face / body / fixed anchors stay the same across all states
- Expression / pose / action / emotion are not states → direct in the page illustration guide
- If outfit/gear changes mid-story, add a new state slug
```

```markdown
<!-- `characters.md` index template -->
# Character List

| Name | Role | Key trait |
|------|------|----------|
| {name} | {role} | {key trait} |

## Relationship Map
{character relationships}
```

---

### 2.3 Location Catalog Design

1. Create `locations.md` as an index (table only).
2. For each location, create `locations/{location-slug}.md` using the template below.

```markdown
# {Location Name}

## Basic Info
- Type: {world / region / city / building / room / outdoor / ...}
- Narrative role: {why this place matters to the story}

## Spatial Layout
- Structure / layout: {spatial layout — rooms, corridors, open areas, landmarks}
- Scale feel: {size impression — cramped / vast / intimate / etc.}

## Sensory Environment
- Lighting: {natural light / artificial / dim / harsh / flickering / ...}
- Temperature / humidity: {cold / warm / humid / dry / seasonal shift}
- Smell: {dominant scent — dust, food, chemicals, nature, decay, ...}
- Sound / noise: {ambient sound — silence, traffic, machinery, birdsong, echoes, ...}
- Texture / touch: {what surfaces feel like if touched — rough stone, polished wood, ...}

## Mood Notes
{psychological impression of the place — safety / threat / nostalgia / isolation / ...}

## Scene List (location → position → view → state)
{This list is the scene-slug inventory used as-is for Stage ④ reference image generation.}
{Each row maps to one reference image.}
{**Visibility coverage:** if any page needs an architectural element (door, hallway, sink front, tub side), that element must appear in some row’s Description — add a row; do not invent it in page YAML. See `workflow/reference-models.md` §3.1.}

| position | view | state | Description (what is visible in this PNG) |
|------|------|-------|------|
| {position-slug} | {view-slug} | base | {e.g. hallway outside open bathroom door; door frame + tub + sink beyond} |
| {position-slug} | {view-slug} | base | {e.g. tub-side facing tub; faucet visible; door NOT in frame} |
| {position-slug} | {view-slug} | {state-slug} | {changed state — fire / destruction / new build / …} |
| ... | ... | ... | ... |

Reference model rules:
- position: a specific spot inside `location` (e.g. hallway-outside-door, tub-side, sink-front)
- view: the **visible scene** from that position (e.g. door-looking-in, toward-tub, toward-mirror). Because a full 3D space is not managed, define views by “what is visible.” A page may only show architecture that appears in the cited row’s reference PNG.
- state: lasting change to the set’s physical structure (base = default). Reference models update only via state change.
  - Included examples: new building permanently changes the outside view, wall destroyed, furniture rearranged for good
  - Excluded examples: curtains open/closed, door open/shut as a one-off beat (if the **door itself** must be visible, that is still a position×view row — open/closed may be Direction within that view), lighting, weather/time, transient cast/dynamic props → page prompt
- Each row (position × view × state) → one reference image. Physical-identity updates are judged by `state` change only.
- **Forbidden:** citing a ref that does not show required elements, then telling Stage ④ to “add the door / other wall.”
```

```markdown
<!-- `locations.md` index template -->
# Location List

| Name | Type | Key trait |
|------|------|----------|
| {name} | {world/region/location/position} | {key trait} |

## Hierarchy
{location hierarchy}
```

---

### 2.3b Staging Catalog Design (continuing-situation continuity)

**When:** The same situation spans multiple pages and relative positions **or situation props** must not flip — café dialogue, **family dinner**, picnic, play mat, OR, meeting, etc. **One staging per situation.** Canonical: `workflow/reference-models.md` §4.

1. Write `stagings.md` index
2. Write `stagings/{staging-slug}.md`:
   - seats / stations / L-R + facing
   - **Situation props / table dressing** map (meal layout, cups, toys, shared objects — locked look + placement)
   - Continuity rules: allowed progressive change (bowl empties) vs forbidden wholesale swap
3. Plan 2–3 reference views (establishing / reverse / detail) that show **blocking + props**
4. Confirm cast character states and location position/view/state exist in catalogs

If a continuing situation has no staging → stop, add staging, then resume.  
If a meal/picnic/play scene has staging but **empty props map** → stop, add props map (empty table ≠ tonight’s dinner).

---

### 2.4 Series (Episode List) Design

1. Create `series.md`:
   - `Episode List` (number / title / summary / expected page count)
   - Overall structure (emotional peaks + page-turn tension flow)
   - **Pages (expected)** includes **Page 0 cover** (e.g. `11` = cover 1 + content 10)

```markdown
# Episode List

| # | Title | Summary | Pages (expected) |
|------|------|------|------------------|
| 1 | {title} | {summary} | {pages incl. Page 0 cover} |

## Overall Structure
{story arc overview — emotional peaks, where page-turn tension builds}
```

---

### 2.5 Episode Page Design

For each `episodes/{nnn}-{episode-slug}.md`, design **all pages**:

#### Canonical page schema (mandatory)

**Only the flat page schema is allowed.** Every page is a `### Page {N}` heading followed by a **flat** `- **Field:**` list. Do **not** use nested `#### Page story`, `#### Illustration guide`, `#### Rendering text`, `#### Text–image split`, or `#### Page-turn hook` subsections (nested or hybrid layouts).

**Field notation (fixed):**

- Always `- **Field:** value` — colon **inside** the bold markers (correct: `**Location:**`; wrong: `**Location**:` or `- Location:`).
- Empty values still require the field: use `없음` (optionally with a reason), e.g. `- **Staging:** 없음`.

#### Episode header: Cast roster

```markdown
**Characters:** {slug-a}, {slug-b}, {slug-c} (언급)
**Locations:** {location-a}, {location-b}
```

Rules: on-page ⊆ roster; ghost cast forbidden; mention-only tagged `(언급)` (or equivalent).

#### Required page fields (flat list, this order)

**Every episode must start with `### Page 0` (Cover).** Page 0 is mandatory — not optional. Story body pages are `### Page 1` … `### Page {N}`.

Each `### Page {N}` **must** include every field below (value may be `없음`):

| Field | Rule |
|-------|------|
| `- **Page story:**` | Beat this page alone carries (Stage ④ message basis) |
| `- **Staging:**` | `{slug} — ref view: …` or `없음` |
| `- **Characters:**` | slugs + `(state: …)` or `없음 (establishing / …)` |
| `- **Location:**` | `{location} / {position} / {view} (state: …)` |
| `- **Direction:**` | lighting / time / weather / mood / camera / expression / pose — page-specific, not reference model |
| `- **Rendering text:**` | Exact overlay text (locked rendering language) or `없음` |
| `- **Text carries:**` | What text delivers |
| `- **Picture carries:**` | What the image delivers |
| `- **Page-turn hook:**` | Curiosity/tension; final page: `resolution / afterimage` |

Do **not** invent seating **or situation props** in Direction — seating and table dressing come from staging. Detailed image prompts belong in Stage ④ YAML only. Progressive prop notes (e.g. “bowl half empty”) only if Continuity rules allow.

#### Page 0 (Cover) — field value rules

Page 0 uses the **same flat fields** as body pages, with cover-specific values:

| Field | Cover rule |
|-------|------------|
| `- **Page story:**` | Must read as a **cover**: invite the child to open the book — world + hero + adventure mood. Do **not** advance interior plot beats or spoil the climax. |
| `- **Staging:**` | Usually `없음` (hero establishing pose, not a continuing-situation map) |
| `- **Characters:**` | Protagonist (and cover-worthy companions only) with state |
| `- **Location:**` | Establishing location that sells the world |
| `- **Direction:**` | Inviting / heroic cover pose; **title safe zone** (space reserved for title); lighting/mood for shelf appeal |
| `- **Rendering text:**` | **Title** (subtitle optional; optional credit line). **Not** story-body prose or dialogue scenes |
| `- **Text carries:**` | Title (and optional subtitle / credit) |
| `- **Picture carries:**` | World introduction, protagonist presence, adventure mood |
| `- **Page-turn hook:**` | Opening curiosity (“Who is this? What adventure?”) — invite to turn into Page 1; no climax spoiler |

Heading form: `### Page 0` (optional label `(Cover)` after the number is fine). File order: Page 0 first, then Page 1….

#### Craft rules (apply)
- **Page 0 cover mandatory** on every episode
- Page-turn hook: required (Page 0 = opening curiosity; final page = resolution/afterimage wording)
- Text–image split: “no overlap” is **not** mandatory — essential page-story info must be delivered by text + image together
- Read-aloud rhythm: sentence / vocabulary density fit for target age (Page 0 title may be denser/display typography)
- No didactic close: final page ends as a scene, not a sermon / moral monologue
- Age density: tune word choice to the age range in `overview.md`
- Craft Notes **Page count** = measured `### Page` headings **including Page 0**; sync `series.md` or note exception (expected pages = cover + content)
- **Beat chronology:** before G2, walk the physical sequence (undress → bath → dry → dress; sit → eat → finish). Contradictions → reorder pages in Design
- **Visibility coverage:** each page’s Location cite must be a position×view whose ref shows every must-see fixture/architecture (`workflow/reference-models.md` §3.1)

**Multilingual example (Korean edition rendering text only):**

```text
Rendering text (Korean):
“미나는 창밖을 보았다. 구름이 천천히 지나가고 있었다.”
```

YAML image prompts stay English; overlay strings keep the target-language wording verbatim.

```markdown
<!-- `episodes/{nnn}-{episode-slug}.md` template — flat page schema only -->
# Episode {nnn}: {Title}

## Summary
{episode summary — what happens, emotional arc}

**Characters:** {slug-a}, {slug-b}, {slug-c} (언급)
**Locations:** {location-a}, {location-b}

## Craft Notes
- Age density: {words/vocabulary fit for target age}
- Theme delivery: {how theme is shown in scenes — not stated as sermon}
- Climax page: {page number}  <!-- body page, not Page 0 -->
- Page count (this episode): {n}  <!-- MUST equal count of ### Page headings, including Page 0 -->
- series.md sync: {updated to measured pages (cover + content) | exception noted: …}

## Pages

### Page 0
- **Page story:** Cover — {world + hero + adventure mood that invites opening the book; no interior plot spoiler}
- **Staging:** 없음
- **Characters:** {protagonist-slug} (state: {state-slug or base})  <!-- + cover companions if needed -->
- **Location:** {location-slug} / {position-slug} / {view-slug} (state: {state-slug or base})
- **Direction:** {heroic/inviting pose; title safe zone; lighting/mood for cover appeal}
- **Rendering text:** {exact title — locked language; optional subtitle / credit | not story-body prose}
- **Text carries:** Title (and optional subtitle / credit)
- **Picture carries:** World introduction, protagonist presence, adventure mood
- **Page-turn hook:** Opening curiosity — invite to turn into Page 1

### Page {N}
- **Page story:** {setting, situation, emotion, beat — imaginable standalone}
- **Staging:** {staging-slug — ref view: establishing|reverse|detail | 없음}
- **Characters:** {character-slug} (state: {state-slug or base})  <!-- or: 없음 (establishing) -->
- **Location:** {location-slug} / {position-slug} / {view-slug} (state: {state-slug or base})
- **Direction:** {lighting / time / weather / mood, camera, expression / pose}
- **Rendering text:** {exact overlay — storybook voice; locked language | 없음}
- **Text carries:** {…}
- **Picture carries:** {…}
- **Page-turn hook:** {curiosity / tension | final: resolution / afterimage}
```
---

### 2.x Episode Design Internal Feedback Loop (story ↔ catalogs)

**Story design co-locks catalogs** — see `workflow/reference-models.md` §7.

While designing episodes/pages, design **together**:

1. **Staging** — for each continuing situation (cite existing staging **or** **add** `stagings/{slug}.md` with blocking **+ props/table dressing**)
2. **Character appearance state** — outfit / gear / lasting body (cite existing state **or** **add** state to `characters/{slug}.md`)
3. **Location set state** — position / view / state (cite existing **or** **add** to `locations/{slug}.md`)

**Reuse if present; if missing, add to the catalog then cite.** Do not invent outfit, gear, set damage, seats, L/R, **or meal/toy/prop layouts** only inside the episode file.

When something is missing:
1. Pause the story unit
2. Update Design catalogs first (`characters/*`, `locations/*`, `stagings/*`, `world-bible.md`)
3. Resume episode design and cite
4. Re-check affected pages

Additional checks:
- Episode file uses the **flat page schema only** (no nested `####` page subsections)
- **`### Page 0` cover exists** and follows cover field-value rules (title overlay; invite, don’t spoil)
- Cast roster ↔ page Characters (no ghosts; mention-only tagged)
- Craft Notes page count == measured `### Page` headings (incl. Page 0); series.md synced or exception noted
- Character **states** (outfit/gear) cited by pages exist in catalogs
- Location **position/view/state** cited by pages exist in catalogs
- Continuing situations cite a **staging**; staging cast states and location anchor match catalogs; **props/table dressing map present** when the situation has shared objects
- Expression / time / weather are page direction, not new reference states; progressive prop depletion only per Continuity rules
- Every page has all required fields (value may be `없음`)

---

## Completeness Check (Stage ②)

- [ ] `illustration-guide.md` exists (art style + layout + typography / text-box roles locked) — `workflow/illustration-guide.md`
- [ ] `world-bible.md`, `characters.md`, `locations.md`, `series.md` exist
- [ ] Continuing-situation `stagings.md` + `stagings/{slug}.md` (when needed) with **blocking + situation props** — `workflow/reference-models.md` §4
- [ ] All `characters/{character-slug}.md`, `locations/{location-slug}.md` authored
- [ ] All `episodes/{nnn}-{episode-slug}.md` use **flat page fields** (no nested `####` page subsections)
- [ ] Every episode has **`### Page 0` cover** with title overlay + cover field-value rules
- [ ] Cast roster present; no ghost cast; mention-only tagged
- [ ] Craft Notes page count matches measured pages (incl. Page 0); series.md synced or exception noted
- [ ] Text–image split (essential-info sufficiency) and page-turn hook rules applied
- [ ] Every page’s character state / location position-view-state / staging (when used) is catalogued so Stage ④ Phase 0 reference generation is possible
- [ ] **Visibility coverage:** each page’s must-see fixtures/architecture appear in the cited location Scene List row (`workflow/reference-models.md` §3.1)
- [ ] **Beat chronology** walked (undress→bath→dry→dress etc.) with no contradictions
- [ ] During story design, new refs were **added to catalogs first** and existing ones **reused** (`workflow/reference-models.md` §7)

---

## Gate G2 (User Approval)

Do not proceed to Stage ③ until the user approves:
- **`illustration-guide.md`:** art style, layout, and typography/text-box recipes specific enough that pages won’t drift?
- World / rules / background: consistent and specific enough?
- Character catalog: outfit/gear per state correct; face/body consistency held?
- Location catalog: location → position → view → state hierarchy clear?
- Staging catalog (when needed): continuing situations have who-is-where **+ situation props/table dressing** (meal/picnic/toys…) + planned ref views?
- Episode / page design:
  - Flat page schema only throughout?
  - Does every episode open with **Page 0 cover** (title overlay; invite, don’t spoil)?
  - Cast roster clean (no ghosts)?
  - Does page story deliver emotion well?
  - Can page story alone make the page beat clearly imaginable?
  - Does rendering text fit age and rhythm?
  - Are text–image split (essential-info sufficiency) and page-turn hooks present?
  - Does the final page end as a scene, not a sermon?
  - Do Direction / Characters / Location / Staging fields carry the needed scene elements?
  - Continuing meal/café/play spans cite staging (not Location alone)?
  - Reference-image readiness (character states / location position-view-states / stagings defined in catalogs)?

**Do not proceed until G2 is approved.**