# Stage ② — Design (World/Characters/Locations + Series/Episode/Cut Design)

**Prerequisites:** Approved `overview.md`

**Gate artifacts (design layer):**
- `{project-root}/world-bible.md`
- `{project-root}/characters.md` + `characters/{character-slug}.md`
- `{project-root}/locations.md` + `locations/{location-slug}.md`
- `{project-root}/stagings.md` + `stagings/{staging-slug}.md` (when ensemble blocking is needed)
- `{project-root}/series.md`
- `episodes/{nnn}-{episode-slug}.md` (episode files include **page → cut** design)

**Next stage:** `03-evaluate.md` (after user approves design gates)

---

## Procedure

### 2.0 Design constraints (global)

- **Architecture-first**: Image generation only in Stage ④.
- **Consistency sources**:
  - Characters: `characters/{character-slug}.md` (body identity + lasting change + equipment states)
  - Locations: `locations/{location-slug}.md` (set/stage; `state` for lasting set change; `position`/`view` framing anchors)
  - Stagings: `stagings/{staging-slug}.md` (who sits/stands where across a multi-cut span)
  - Story / direction: `episodes/{nnn}-{episode-slug}.md` (per page/cut: art / balloons / dialogue / narration / gutters)
- **Design-first fixes**: If Evaluate finds issues, fix Design files first — do not patch via image prompts alone.
- **Canonical reference rules:** `workflow/reference-models.md` (required reading for catalogs + cut guides).
- **Unlike picture books**:
  - Page height is **not fixed** — varies with cut count, gutters, balloon load
  - A page stacks **multiple cuts** top→bottom; **gutters** carry rhythm and scene breaks
  - Spoken lines use **speech balloons**; optional **captions/narration boxes** for exposition
  - Use **scroll hooks** (not page-turn hooks) to pull the reader down

#### What is a Reference Model?

**Reference model = durable identity or spatial anchor that page images must obey.**
Stage ④ generates matching **reference images**; later page images stay consistent against them.

Full rules: `workflow/reference-models.md`. Summary:

| Layer | **Included** (= reference image) | **Not included** (= cut direction) |
|------|------|------|
| **Character** | Face/body/silhouette; lasting body change (growth, accident aftermath); **equipment identity** (clothes, accessories, weapons, shields — e.g. the same sword across fight cuts) | Expression, mood, transient pose/action |
| **Location** | The **set/stage**: physical structure, fixed set dressing; lasting damage/rebuild | Time of day, season, weather, lighting mood, one-off camera, passersby |
| **Staging** | **Who is where** relative to whom for a **continuing situation** (café L/R, OR stations, meeting seats, standoffs, …) — one staging per situation; **2–3** ensemble reference views | Expression/mood; camera tightness as long as blocking stays recognizable |

**New reference image needed when:**
- Character: outfit/identity gear changes, or lasting body change → new state
- Location: set permanently changes (new building, fire damage that sticks) → new state
- Staging: reseat / leave / reform for the rest of the span → new or updated staging

**Handle in cut prompts (no new identity reference) when:**
- Expression, emotion, mid-action pose
- Time / season / weather
- Temporary/reversible set bits (curtains, door open for one beat)
- Transient elements (passersby, VFX)

**Staging is required when** a situation continues across cuts and relative placement must not flip (café L/R, OR doctor/nurse stations, meeting seats, etc.). Character + empty location refs alone do not lock positions. **One staging per continuing situation.**

---

### 2.1 World Design

1. Write `world-bible.md`:

```markdown
# World

## Name
{name}

## Description
{rules, atmosphere, key features}

## Core rules
- {rule 1}
- {rule 2}

## History / background
{background}
```

---

### 2.2 Character Catalog Design

1. Create `characters.md` as an index (table only).
2. For each character, create `characters/{character-slug}.md`:

```markdown
# {Character Name}

## Basics
- Role (lead / support / antagonist / …): {role}
- Core Drive: {what they want most}
- Central Conflict: {what blocks it}

## Appearance & behavior
- Fixed face/body/silhouette points (must not drift): {detail}
- Signature gestures / habits: {…}
- Recurring expression / speech feel: {…}  # performance notes — NOT reference states

## Relationships
- {other-character-slug}: {type, power dynamic, conflict seed}
- {other-character-slug}: ...

## Reference models (appearance + equipment states)
{Each state → one reference image in Stage ④ Phase 0}
{Locks identity gear: clothes, accessories, weapons, shields — do not silently swap across cuts}

- base: {default outfit + identity gear (e.g. named sword, shield, accessories)}
- {state-slug}: {what changes: outfit / gear / lasting body change; face/body identity rules unchanged}
- {state-slug}: ...

Rules:
- state = lasting physical/equipment identity only
- Same fight gear across cuts unless a designed state change
- Expression / mood / transient pose are NOT states → cut illustration guide
```

```markdown
<!-- `characters.md` index template -->
# Characters

| Name | Role | Key trait |
|------|------|----------|
| {name} | {role} | {key trait} |

## Relationship map
{character relationships}
```

---

### 2.3 Location Catalog Design

1. Create `locations.md` as an index (table only).
2. For each location, create `locations/{location-slug}.md`:

```markdown
# {Location Name}

## Basics
- Type: {world / region / city / building / room / outdoor / ...}
- Narrative role: {why this place matters}
- Role as set/stage: {what physical structure must stay recognizable}

## Spatial layout
- Structure: {rooms, corridors, landmarks, fixed furniture of the set}
- Scale feel: {cramped / vast / intimate / ...}

## Sensory environment (defaults — may be overridden per cut)
- Lighting: {…}
- Temperature / humidity: {…}
- Smell: {…}
- Sound: {…}
- Texture: {…}

## Mood notes
{psychological impression}

## Scene list (location → position → view → state)
{Each row → one Stage ④ location reference image}
{state = lasting set change only — not time/season/weather}

| position | view | state | Notes |
|------|------|-------|------|
| {position-slug} | {view-slug} | base | {default set} |
| {position-slug} | {view-slug} | {state-slug} | {e.g. after fire / new wing built} |
| ... | ... | ... | ... |

Rules:
- position: a spot inside the location
- view: **what of the set is visible** (not a one-cut camera choice)
- state: permanent physical-structure change (base = default)
- Time/season/weather → cut direction, not a new location state
- Who is where across many cuts of the **same situation** → **staging**, not location alone (café L/R, OR stations, meeting seats, …)
```

```markdown
<!-- `locations.md` index template -->
# Locations

| Name | Type | Key trait |
|------|------|----------|
| {name} | {world/region/location/position} | {key trait} |

## Hierarchy
{location hierarchy}
```

---

### 2.3b Staging Catalog Design (continuing-situation blocking)

**When:** any multi-cut (or multi-page) **continuing situation** needs stable relative placement — café dialogues, OR procedures, meetings, meals, courtrooms, formations, vehicles, etc. Default: **one staging per situation**.

1. Create `stagings.md` as an index.
2. For each staging, create `stagings/{staging-slug}.md` using the template in `workflow/reference-models.md` §4 (name by situation, e.g. `cafe-a-b-first-date`, `or-appendectomy-team`).
3. Plan **2–3** reference views (e.g. establishing / reverse / detail) that preserve L/R and stations.
4. Ensure every cast member’s appearance state and the location position/view/state already exist in catalogs.

```markdown
<!-- `stagings.md` index template -->
# Stagings

| Slug | Situation | Location | Cast (short) | Continuity span | Ref views |
|------|-----------|----------|--------------|-----------------|-----------|
| {staging-slug} | {café / OR / meeting / …} | {location} | {A, B, C} | ep{nnn} p{a}.{c}–p{b}.{c} | 3 |
```

If episode design discovers a continuing situation without a staging → stop, add staging (+ catalog gaps), then resume. Do not invent seats, L/R, or OR stations only inside cut prompts.

---

### 2.4 Series (Episode List) Design

1. Create `series.md` with Episode List + overall arc (emotional peaks + **scroll tension**):

```markdown
# Episode List

| # | Title | Summary | Pages (est.) | Cuts (est.) |
|---|------|---------|--------------|-------------|
| 1 | {title} | {summary} | {pages} | {cuts} |

## Overall structure
{story arc — emotional peaks, where scroll tension builds; note where tall cuts / pause gutters / action stacks land}
```

---

### 2.5 Episode Page & Cut Design

For each `episodes/{nnn}-{episode-slug}.md`, design **all pages**, and inside each page design **all cuts** top→bottom.

**Required reading:** `workflow/cut-composition.md` — cut height, gutter distance, size class, and viewport rhythm are first-class design decisions (not afterthoughts).

#### Webtoon craft rules

1. **Use vertical space for characters and dialogue** — e.g. dialogue lines, then the figure, then breathing room; do not force print-page packing.
2. **Direct with cut height** — standard (~400–600px) for dialogue/everyday; tall/long (~1,200–2,000px+) for climax/spectacle; open/diagonal + tight gutters for action.
3. **Direct with gutters** — normal (~150–300px) for everyday rhythm; `pause` (~500–1,000px+) for silence/time; tight for speed.
4. **Mobile viewport** — plan what one phone screen shows; one clear beat (or intentional partial reveal of a tall cut).
5. **Direct with color** — e.g. flashback via sepia/desaturation; color is a primary tool.
6. **Variable page height** — page height = sum of cut heights + gutters; do not force equal page ratios.
7. **Vary size classes across the episode** — do not make every cut the same height unless that flatness is intentional.

#### Required page fields

Each `Page {idx}`:
- `Page 0` = cover/thumbnail vertical segment (or agreed series thumbnail rule)
- `Page story`: what this vertical segment conveys
- `Estimated height`: sum of cut heights + gutters (px guide); note if taller than one 9:21 segment
- `Viewport plan`: what one phone screen should show in this segment (1 beat / partial tall reveal / …)
- `Outside-cut area` / side margins: follow overview unless page exception
- `Scroll hook`: pull to next page (non-final); final = resolution / aftertaste

#### Required cut fields

Each `Cut {n}` (top→bottom):
- `Cut story`: the beat this cut alone carries
- `Purpose` / directing intent: dialogue | everyday | climax | silence/time | action | …
- `Size class`: `standard` | `tall`/`long` | `compact` | `open`/`diagonal` (see cut-composition.md)
- `Height (guide)`: target px range (e.g. 500 / 1500 / …)
- `Cut shape`: rectangle | borderless open | diagonal | bleed / break-frame
- `Illustration guide`:
  - **Staging** (if this cut is in a continuity span): `{staging-slug}` + which ref view (`establishing` / `reverse` / `detail`)
  - Character/location **slugs** (must match staging cast/location when staging is cited)
  - Character **state** or base (identity gear locked — no silent sword/outfit swap)
  - Location **position + view** (+ **state** if needed) — set/stage only
  - Cut direction: lighting / time / season / weather / color grade / camera tightness / **expression** / gesture within seat
  - Do **not** invent seating here — seating comes from staging
- `Balloons` (0+):
  - speaker (character slug or narrator)
  - type: speech / thought / shout / whisper / …
  - text: display string **verbatim in the target language**
  - placement: approx. in-cut position (top/mid/bottom × left/center/right) — must not cover face/key action
- `Captions` (optional): exposition / time / place / inner thought + placement
- `Gutter to next cut`:
  - class: `tight` | `normal` | `wide` | `pause`
  - distance (guide): px range
  - intent: pace / transition / silence / time / aftertaste

```markdown
<!-- `episodes/{nnn}-{episode-slug}.md` template -->
# Episode {nnn}: {Title}

## Summary
{what happens, emotional arc}

## Craft Notes
- Scroll rhythm: {fast / normal / lingering sections — map to size/gutter classes}
- Theme delivery: {shown in scenes — not sermon captions}
- Climax page.cut: {page.cut} — expect tall/long size class
- Cut count (this episode): {n} (target range from overview)
- Canvas: width {690–800px locked}, height variable, color, side margins {full bleed / 30–50px}, outside-cut {white/black/theme}

## Pages

### Page {N}

#### Page story
{story this vertical segment conveys}
{Must be imaginable from text alone — basis for Stage ④ message.}
{Only catalog-registered characters/locations}

#### Estimated height
{sum of cut heights + gutters in px; note strip-split if needed}

#### Viewport plan
{what one phone screen shows here — one beat / partial tall reveal / …}

#### Scroll hook
{curiosity/tension to scroll down — final page: resolution/aftertaste}

#### Cuts (top → bottom)

##### Cut {n}

###### Cut story
{beat}

###### Composition
- Purpose: {dialogue / everyday / climax / silence-time / action / …}
- Size class: {standard | tall/long | compact | open/diagonal}
- Height (guide): {e.g. 500px / 1500px}
- Shape: {rectangle | open | diagonal | bleed}

###### Illustration guide
- Staging (if any): {staging-slug} — ref view: {establishing|reverse|detail}
- Character: {character-slug} (state: {state-slug or base})  # must match staging cast when staging is set
- Location: {location-slug} / {position-slug} / {view-slug} (state: {state-slug or base})
- Direction: {expression / gesture / camera tightness / time-season-weather grade}
{Seating/formation comes from staging — do not reseat in this cut unless Design updates staging}
{Detailed image prompt only in Stage ④ YAML}

###### Balloons
| # | speaker | type | text | placement |
|---|---------|------|------|-----------|
| 1 | {slug} | speech | {dialogue in target language} | {e.g. top-left} |

###### Captions
| # | text | placement |
|---|------|-----------|
| 1 | {optional, target language} | {e.g. top-center} |

###### Gutter to next cut
- Class: {tight | normal | wide | pause}
- Distance (guide): {e.g. 200px / 800px}
- Intent: {pace / transition / silence / time / aftertaste}
```

---

### 2.x Episode Design Internal Feedback Loop (story ↔ catalogs)

**Story design co-locks catalogs** — see `workflow/reference-models.md` §7.

While designing episodes / pages / cuts, **together** design:

1. **Staging** for each continuing situation (cite existing or **add** `stagings/{slug}.md`)
2. **Character appearance states** (outfit/gear/lasting body) for everyone on stage (cite existing or **add** state to `characters/{slug}.md`)
3. **Location set states** (cite existing position/view/state or **add** to `locations/{slug}.md`)

**Reuse if present; add if missing.** Never invent outfit, gear, set damage, seats, or L/R only inside the episode file.

If a gap appears mid-design:
1. Stop the story unit
2. Update Design catalogs first (`characters/*`, `locations/*`, `stagings/*`, `world-bible.md`) with user visibility
3. Resume episode design and cite the new/existing refs
4. Re-check every affected cut

Also verify:
- Every cut’s character **state** exists in catalogs; identity gear does not silently swap
- Every cut’s location **position/view/state** exists in catalogs
- Multi-cut ensembles with fixed seating/formation cite a **staging**; staging has 2–3 planned ref views; staging cast/location anchors match catalogs
- Expression/mood/time/season/weather are cut direction, not new reference states
- Every spoken line has speaker + balloon type + placement
- Every cut has **size class + height guide + purpose** aligned with `cut-composition.md`
- Gutter class/distance matches scene change / silence / speed
- Episode cut count is near the overview target (or short-form exception is noted)
- Viewport plan is stated per page (or per dense stretch)

---

## Completeness Check (Stage ②)

- [ ] `world-bible.md`, `characters.md`, `locations.md`, `series.md` exist
- [ ] All `characters/{character-slug}.md`, `locations/{location-slug}.md` written
- [ ] `stagings.md` + needed `stagings/{staging-slug}.md` exist for ensemble continuity spans
- [ ] All `episodes/{nnn}-{episode-slug}.md` have page → cut design (art/balloons/dialogue/captions/composition/gutters/staging cites)
- [ ] Reference rules followed (`workflow/reference-models.md`)
- [ ] Cut size/gutter classes vary with dramatic intent (`workflow/cut-composition.md`)
- [ ] Variable page height + scroll hooks + viewport plans applied
- [ ] Character / location / staging catalogs ready for Stage ④ Phase 0
- [ ] Story co-lock: every cut’s staging/character-state/location-state is catalogued (reuse or added) — `workflow/reference-models.md` §7

---

## Gate G2 (User Approval)

Do not proceed to Stage ③ until the user approves:
- World rules/background: coherent and specific enough?
- Character catalog: lasting body + outfit/identity-gear states; expressions not treated as states?
- Location catalog: set/stage clear; time/weather not mistaken for location states?
- Staging catalog: each continuing situation (café / OR / meeting / …) has who-is-where + 2–3 ref views?
- Episode/page/cut design:
  - Cut stories carry emotion?
  - Staging cited wherever relative placement must not drift (incl. L/R pairs and OR stations)?
  - Cut **size classes** match intent (standard vs tall vs open/action)?
  - **Gutters** used intentionally (normal vs pause vs tight)?
  - Balloons/captions respect top→bottom reading and face clearance?
  - Viewport plans readable on a phone?
  - Scroll hooks present?
  - Reference-image readiness (characters / locations / stagings)?

**Do not proceed until G2 is approved.**
