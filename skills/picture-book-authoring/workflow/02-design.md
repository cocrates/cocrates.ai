# Stage ② — Design (World/Characters/Locations + Series/Episode Design)

**Prerequisites:** Approved `overview.md`

**Gate artifacts (design layer):**
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
  - Character consistency: `characters/{character-slug}.md` (base appearance / face / silhouette + reference model list)
  - Location consistency: `locations/{location-slug}.md` (reference models center on `state` change; `position` · `view` are framing anchors)
  - Story / direction consistency: `episodes/{nnn}-{episode-slug}.md` (per-page rendering text + illustration guide)
- **Design-first fixes**: If Evaluate flags a problem, do not fix by “image prompt only” — update Design files first, then re-evaluate.

#### What is a Reference Model?

**Canonical rules:** `workflow/reference-models.md` (same three layers as webtoon / video / novel).

**Reference model = durable physical appearance/structure across the story + locked relative placement for continuing situations.**
In Stage ④ these become “reference images” via the image-generation skill; every later page image builds consistency from them.

| Kind | **Included** in reference model (= generate as reference images) | **Excluded** (= direct in page illustration guide) |
|------|------|------|
| **Character** | Face / body / silhouette, **lasting body change**, **equipment identity** (clothes, accessories, weapons, shields — do not swap to a different sword page-to-page) | Expression, pose, action, emotion |
| **Location** | Physical structure / layout / fixed set dressing of the **set / stage** | Lighting, time of day, season, weather, mood, camera direction / framing / angle, transient elements |
| **Staging** | **Who is where** in a continuing situation (café L/R, meeting seats, OR stations, …). One per situation; **2–3** reference images | Expression / mood; camera tightness while seats stay locked |

**When a reference model change is required (= new reference image):**
- Character: outfit/gear changes, or lasting body change → add a new state slug
- Location: physical structure changes permanently (e.g. new building outside the window, wall destroyed) → add a new state slug
- Staging: seats/formation change or a new situation starts → new staging

**When a reference model change is not required (= handle in page prompt):**
- Expression, emotion, transient pose
- Temporary / reversible changes (curtains open/closed, door open/shut)
- Lighting / time / weather / mood shifts
- Transient elements (passerby, bird flying by)

**When staging is required:** The same situation spans multiple pages and relative positions must not flip (café facing seats, OR, meeting, …). Write `stagings.md` + `stagings/{slug}.md`, then generate 2–3 ensemble refs in Phase 0. Character + empty location alone do not lock seats / L-R.

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

| position | view | state | Description |
|------|------|-------|------|
| {position-slug} | {view-slug} | base | {base-state description} |
| {position-slug} | {view-slug} | {state-slug} | {changed state — fire / destruction / new build / …} |
| ... | ... | ... | ... |

Reference model rules:
- position: a specific spot inside `location` (e.g. living-room-window-side, living-room-door-side, bed-headboard-side)
- view: the **visible scene** from that position (e.g. window-facing view, bed-facing view, bookshelf-facing view). Because a full 3D space is not managed, define views by “what is visible.” Camera direction / framing / angle are not view — those are decided at page image generation.
- state: lasting change to the set’s physical structure (base = default). Reference models update only via state change.
  - Included examples: new building permanently changes the outside view, wall destroyed, furniture rearranged for good
  - Excluded examples: curtains open/closed, door open/shut, lighting, weather/time, transient cast/dynamic props → page prompt
- Each row (position × view × state) → one reference image. Physical-identity updates are judged by `state` change only.
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

### 2.3b Staging Catalog Design (continuing-situation blocking)

**When:** The same situation spans multiple pages and relative positions must not flip — café dialogue, OR, meeting, dinner table, standoff, etc. **One staging per situation.** Canonical: `workflow/reference-models.md`.

1. Write `stagings.md` index
2. Write `stagings/{staging-slug}.md` (seats / stations / L-R + facing)
3. Plan 2–3 reference views (establishing / reverse / detail)
4. Confirm cast character states and location position/view/state exist in catalogs

If a continuing situation has no staging → stop, add staging, then resume.

---

### 2.4 Series (Episode List) Design

1. Create `series.md`:
   - `Episode List` (number / title / summary / expected page count)
   - Overall structure (emotional peaks + page-turn tension flow)

```markdown
# Episode List

| # | Title | Summary | Pages (expected) |
|------|------|------|------------------|
| 1 | {title} | {summary} | {pages} |

## Overall Structure
{story arc overview — emotional peaks, where page-turn tension builds}
```

---

### 2.5 Episode Page Design

For each `episodes/{nnn}-{episode-slug}.md`, design **all pages**:

#### Required page fields
For each `Page {idx}`:
- `Page 0` is fixed as the cover
- `Page story`: the story the scene delivers (setting / situation / emotion / beat)
- `Illustration guide`: concise “must-include” visual components
  - Character / location **slugs**
  - Character **state** (outfit / gear) or `base`
  - Location **position + view** (reference-model axes) and **state** (when physical structure changed)
  - Camera direction / framing / angle (page direction — not reference model)
- `Rendering text`: text that will appear on the page (in the locked rendering language) — storybook voice / read-aloud rhythm
- `Text–image split`: what text carries vs what the picture carries
- `Page-turn hook`: curiosity / tension that pulls the next page (omit on the final page)

#### Craft rules (apply)
- Page-turn hook: required except on the final page
- Text–image split: “no overlap” is **not** mandatory
  - Instead verify that **essential page-story information is sufficiently delivered by text + image together**
- Read-aloud rhythm: sentence / vocabulary density fit for target age
- No didactic close: final page ends as a scene, not a sermon / moral monologue
- Age density: tune word choice to the age range in `overview.md`

**Multilingual example (Korean edition rendering text only):**

```text
Rendering text (Korean):
“미나는 창밖을 보았다. 구름이 천천히 지나가고 있었다.”
```

YAML image prompts stay English; overlay strings keep the target-language wording verbatim.

```markdown
<!-- `episodes/{nnn}-{episode-slug}.md` template -->
# Episode {nnn}: {Title}

## Summary
{episode summary — what happens, emotional arc}

## Craft Notes
- Age density: {words/vocabulary fit for target age}
- Theme delivery: {how theme is shown in scenes — not stated as sermon}
- Climax page: {page number}

## Pages

### Page {N}

#### Page story
{Story this page delivers — setting, situation, emotion, beat}
{Important: reading page story alone must make this page (and the episode beat) clearly imaginable. This becomes the Stage ④ image-generation “message”.}
{Characters and places: catalog-registered names only}

#### Illustration guide
- Staging (when applicable): {staging-slug} — ref view: {establishing|reverse|detail}
- Characters: {character-slug} (state: {state-slug or base})  # keep equipment identity
- Location: {location-slug} / {position-slug} / {view-slug} (state: {state-slug or base})
- Page direction: {lighting / time / weather / mood, camera direction / framing / angle, expression / pose — page-specific, not reference model}
{Seats / L-R for continuing situations come from staging — do not invent on the page}
{Detailed image prompts are written in Stage ④ image-generation YAML}

#### Rendering text
{Exact page overlay text — storybook voice, read-aloud rhythm; in the locked rendering language}

#### Text–image split
- Text carries: {…}
- Picture carries: {…}
{Check: are essential page-story facts delivered by text+image together? (yes / fix)}

#### Page-turn hook
{Curiosity / tension that makes the reader turn — on the final page write “resolution / afterimage”}
```

---

### 2.x Episode Design Internal Feedback Loop (story ↔ catalogs)

**Story design co-locks catalogs** — see `workflow/reference-models.md` §7.

While designing episodes/pages, design **together**:

1. **Staging** — for each continuing situation (cite existing staging **or** **add** `stagings/{slug}.md`)
2. **Character appearance state** — outfit / gear / lasting body (cite existing state **or** **add** state to `characters/{slug}.md`)
3. **Location set state** — position / view / state (cite existing **or** **add** to `locations/{slug}.md`)

**Reuse if present; if missing, add to the catalog then cite.** Do not invent outfit, gear, set damage, seats, or L-R only inside the episode file.

When something is missing:
1. Pause the story unit
2. Update Design catalogs first (`characters/*`, `locations/*`, `stagings/*`, `world-bible.md`)
3. Resume episode design and cite
4. Re-check affected pages

Additional checks:
- Character **states** (outfit/gear) cited by pages exist in catalogs
- Location **position/view/state** cited by pages exist in catalogs
- Continuing situations cite a **staging**; staging cast states and location anchor match catalogs
- Expression / time / weather are page direction, not new reference states

---

## Completeness Check (Stage ②)

- [ ] `world-bible.md`, `characters.md`, `locations.md`, `series.md` exist
- [ ] Continuing-situation `stagings.md` + `stagings/{slug}.md` (when needed) + compliance with `workflow/reference-models.md`
- [ ] All `characters/{character-slug}.md`, `locations/{location-slug}.md` authored
- [ ] All `episodes/{nnn}-{episode-slug}.md` have per-page rendering text + illustration guides
- [ ] Text–image split (essential-info sufficiency) and page-turn hook rules applied
- [ ] Every page’s character state / location position-view-state / staging (when used) is catalogued so Stage ④ Phase 0 reference generation is possible
- [ ] During story design, new refs were **added to catalogs first** and existing ones **reused** (`workflow/reference-models.md` §7)

---

## Gate G2 (User Approval)

Do not proceed to Stage ③ until the user approves:
- World / rules / background: consistent and specific enough?
- Character catalog: outfit/gear per state correct; face/body consistency held?
- Location catalog: location → position → view → state hierarchy clear?
- Episode / page design:
  - Does page story deliver emotion well?
  - Can page story alone make the page beat clearly imaginable?
  - Does rendering text fit age and rhythm?
  - Are text–image split (essential-info sufficiency) and page-turn hooks present?
  - Does the final page end as a scene, not a sermon?
  - Does the illustration guide carry the needed scene elements?
  - Reference-image readiness (character states / location position-view-states defined in catalogs)?

**Do not proceed until G2 is approved.**
