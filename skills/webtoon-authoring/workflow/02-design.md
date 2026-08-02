# Stage ② — Design (World/Characters/Locations + Series/Episode/Cut Design)

**Prerequisites:** Approved `overview.md`

**Gate artifacts (design layer):**
- `{project-root}/illustration-guide.md` (**series-wide** art / layout / balloon·caption typography lock — see `workflow/illustration-guide.md`)
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
  - **Visual / typography:** `illustration-guide.md` (art style, vertical chrome, balloon/caption recipes — series-wide; `workflow/illustration-guide.md`)
  - Characters: `characters/{character-slug}.md` (body identity + lasting change + equipment states)
  - Locations: `locations/{location-slug}.md` (set/stage; `state` for lasting set change; `position`/`view` framing anchors)
  - Stagings: `stagings/{staging-slug}.md` (who sits/stands where across a multi-cut span)
  - Story / direction: `episodes/{nnn}-{episode-slug}.md` (per page/cut: art / balloons / dialogue / narration / gutters — **not** lettering invention)
- **Design-first fixes**: If Evaluate finds issues, fix Design files first — do not patch via image prompts alone.
- **Do not invent per-page balloon fonts or caption boxes.** Tile YAML places locked strings; style tokens come from `illustration-guide.md`.
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
| **Staging** | WHO–WHERE for a **continuing situation** — one staging per situation; **default 1× establishing** (optional reverse/detail only) | Expression/mood; CU/OTS in Direction; reseat without Design |

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

### 2.0b Illustration Guide (series-wide)

1. Expand `overview.md` → Art style + Canvas into `{project-root}/illustration-guide.md`.
2. Follow the full template and rules in `workflow/illustration-guide.md` (art style, vertical chrome, balloon/caption recipes, Stage ④ checklist).
3. Do not proceed to heavy episode cut design without a draft guide strong enough that two pages would share lettering and balloon chrome.
4. G2 includes explicit approval of `illustration-guide.md`.

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
{Static look for a story span. Each state → one turnaround reference image in Stage ④ Phase 0}
{Locks identity gear: clothes, armor, weapons, shields, glasses, accessories — do not silently swap across cuts}
{Document each state concretely in this md; YAML concretizes the md — no invention}

- base: {concrete default outfit + identity gear — face/body anchors + clothes + key accessories}
- {state-slug}: {concrete lasting change: growth/accident aftermath OR outfit/gear change; what stays vs what differs}
- {state-slug}: ...

Rules:
- state = lasting physical/equipment identity only (static look)
- Same fight gear across cuts unless a designed state change
- Expression / mood / transient pose are NOT states → cut illustration guide
- Each image = **one turnaround sheet** (front + side + back + key accessories) for that state
- Co-appearing majors need **≥2-axis visual contrast** (hair/face/outfit) + ID keywords in prompts (`workflow/reference-models.md` §2.3)
- World taxonomy classes from world-bible must match appearance Direction (§2.4)
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

## Spatial layout (static / physical — empty set)
- Structure: {rooms, corridors, landmarks, fixed furniture of the set}
- Physical detail examples: {street → buildings, signage, lights, crosswalk, façades; café → seats, counter, fixed décor, fixtures}
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
{One PNG cannot hold full 3D — plan multi-view: front/left/right or window-side/bed-side/door-side}
{state = lasting set change only — not time/season/weather}
{YAML for each row concretizes this Description — no invention}

| position | view | state | Notes (what is visible in this PNG) |
|------|------|-------|------|
| {position-slug} | {view-slug} | base | {default set — name visible architecture} |
| {position-slug} | {view-slug} | {state-slug} | {e.g. after fire / new wing built} |
| ... | ... | ... | ... |

Rules:
- position: a spot inside the location
- view: **what of the set is visible** (not a one-cut camera choice)
- state: permanent physical-structure change (base = default)
- Time/season/weather → cut direction, not a new location state
- Who is where / ambient extras across many cuts of the **same situation** → **staging**, not location alone
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

**When:** any multi-cut **continuing situation** needs stable relative placement, situation props, and/or ambient occupancy — café dialogues, street walk-and-talk, OR, meetings, meals, etc. Default: **one staging per situation**.

**Episode-first (preferred):** Define staging **with the episode** that needs it. Same location ≠ reuse prior staging. Prefer a new episode-appropriate staging for this beat/time/occupancy; reuse only when the same continuity span truly continues. May require **extra location Scene List rows** for this episode’s framings. Canonical: `workflow/reference-models.md` §4 / §4.1 / §4.2.

**Mandatory:** ≥2 cuts in the same situation **and** (≥2 named cast **or** fixture-relative placement matters) → staging required. `Staging: none` only for single-cut establishing / no fixed placement.

1. Create `stagings.md` as an index.
2. For each staging, create `stagings/{staging-slug}.md` using the template in `workflow/reference-models.md` §4 (name by episode situation, e.g. `ep003-street-walk-talk`).
3. Header **must** bind catalog paths: `Location: slug / position / view (state)` + each `character-slug (state: …)`.
4. Include blocking + situation props + ambient + situation environment (when span-locked) as needed.
5. Plan **canonical establishing** ref (default **one** PNG). Optional reverse/detail only with stated lock purpose — **never** register OTS/CU as staging views.
6. Ensure every cast state and location path already exist (add missing location views first).

```markdown
<!-- `stagings.md` index template -->
# Stagings

| Slug | Situation | Owning episode | Location path | Cast (short) | Continuity span | Ref |
|------|-----------|----------------|---------------|--------------|-----------------|-----|
| {staging-slug} | {street walk / café call / …} | {nnn-…} | {loc/pos/view} | {A, B} | ep{nnn} p{a}.{c}–p{b}.{c} | establishing |
```

If episode design discovers a continuing situation without a staging → stop, add staging (+ catalog gaps), then resume. Do not invent seats, L/R, or OR stations only inside cut prompts.

---

### 2.4 Series (Episode List) Design

1. Create `series.md` with Episode List + overall arc. Plan **what each episode is about**, not how many pages/cuts it will use.

```markdown
# Episode List

| # | Title | Summary |
|---|------|---------|
| 1 | {title} | {what this episode delivers — situation, turn, hook to next} |

## Overall structure
{story arc — emotional peaks, where scroll tension builds across episodes}
```

**Series.md owns only:** episode count, titles, and summaries (plus optional part/arc grouping).

**Do not** put Pages (est.) / Cuts (est.) / page·cut budgets in `series.md`. Volume is decided in each episode file during §2.5 (and recorded in that episode’s Craft Notes after design).

Review focus at series stage: Is the episode count right? Does each title+summary earn a release slot? Does the arc chain (hooks between episodes) work?

---

### 2.5 Episode Page & Cut Design

For each `episodes/{nnn}-{episode-slug}.md`, design **all pages**, and inside each page design **all cuts** top→bottom.

**Episode = distribution unit.** Design enough vertical volume to carry that episode’s story — do **not** compress a full episode into 2–3 pages.

| Spec | Guideline |
|------|-----------|
| **Pages per episode** | **~10–15** (incl. Page 0 / cover-thumbnail if used) |
| **Cuts** | Enough to stage, breathe, and payoff the episode summary — count is **derived** from page packing + dwell (`cut-composition.md` §4b), not from `series.md` |
| Short-form exception | Fewer than ~10 pages only with **user-approved** note in Craft Notes (pilot / promo / special) |

If the episode summary cannot be told without cramming beats onto every page, **add pages** (or split into two episodes in `series.md`) — do not max-pack cuts to fake length.

**Comprehension gate (design-time):** Before G2, re-read only Cut stories + Balloons + Captions (+ what Direction says will be *visible*). Ask: would a first-time reader follow the episode? Any beat that requires unshown prior knowledge or off-panel designer notes is incomplete — expand art, dialogue, or caption (or add pages).

**Required reading:** `workflow/cut-composition.md` — cut height, gutter distance, size class, and viewport rhythm are first-class design decisions (not afterthoughts).

#### Canonical cut schema (mandatory)

**Only the flat cut schema is allowed.** Every cut is a `##### Cut {n}` heading followed by a **flat** `- **Field:**` list. Do **not** use nested `###### Cut story`, `###### Composition`, `###### Illustration guide`, `###### Balloons`, `###### Captions`, or `###### Gutter` subsections (nested or hybrid layouts).

**Field notation (fixed):**

- Always `- **Field:** value` — colon **inside** the bold markers (correct: `**Location:**`; wrong: `**Location**:` or `- Location:`).
- Balloon / caption markdown tables sit under the field line with **2-space indent** on every table row.
- Empty values still require the field: use `none` (optionally with a reason), e.g. `- **Characters:** none (establishing shot)`.

#### Webtoon craft rules

1. **Use vertical space for characters and dialogue** — e.g. dialogue lines, then the figure, then breathing room; do not force print-page packing.
2. **Direct with cut height** — standard (~400–600px) for dialogue/everyday; tall/long (~1,200–2,000px+) for climax/spectacle; open/diagonal + tight gutters for action.
3. **Direct with gutters** — normal (~150–300px) for everyday rhythm; `pause` (~500–1,000px+) for silence/time; tight for speed.
4. **Mobile viewport** — plan what one phone screen shows; one clear beat (or intentional partial reveal of a tall cut).
5. **Direct with color** — e.g. flashback via sepia/desaturation; color is a primary tool.
6. **Variable page height** — page height = sum of cut heights + gutters; do not force equal page ratios.
7. **Pack pages by dwell/breath** — high-dwell cuts may sit alone on a page; otherwise stack n cuts top→bottom (`cut-composition.md` §4b).
8. **Vary size classes across the episode** — do not make every cut the same height unless that flatness is intentional.
9. **Size class ↔ Height must agree** — e.g. do not label `tall` with Height under 1,200px; see `workflow/cut-composition.md`.
10. **Episode volume** — target **~10–15 pages** so the distribution unit can carry its story; avoid 2–3 page compression.
11. **Reader comprehension (art + dialogue + caption)** — a cold reader scrolling the finished episode must understand **who / where / what happens / why it matters** from **illustration + balloons + captions alone**. Do not leave essential facts only in Cut story / Direction notes that never appear on-panel. If something must be known, show it, speak it, or caption it (captions sparingly — not sermon). Gaps → add cuts/pages or on-panel text, do not assume the reader “will get it.”

#### Episode header: Cast roster

After Summary (or inside Summary block), include an explicit cast list:

```markdown
**Characters:** {slug-a}, {slug-b}, {slug-c} (언급)
**Locations:** {location-a}, {location-b}
```

Rules:

- **on-cut** names appear in at least one cut’s `- **Characters:**` field.
- **mention-only** names use a tag such as `(언급)` and must **not** be treated as on-panel cast for Phase 0 staging unless they appear in a cut.
- **Ghost cast forbidden:** a name in the episode Characters list without `(언급)` (or similar) must appear in ≥1 cut; a name only in cuts must be added to the episode list.
- Background extras (unnamed crowd) belong in Direction, not the Characters roster, unless catalogued.

#### Required page fields

Each `Page {idx}`:

- `Page 0` = cover/thumbnail vertical segment (or agreed series thumbnail rule)
- `Page story`: what this vertical segment conveys
- `Estimated height`: sum of cut heights + gutters (px guide); note generation frame / strip-split if needed
- `Viewport plan`: what one phone screen should show **and the intended breath** (e.g. `1 cut — linger` / `2 cuts — quick exchange` / `tall alone — scroll-through`)
- `Cuts on this page`: count + dwell rationale (high / medium / low) — see `workflow/cut-composition.md` §4b
- `Outside-cut area` / side margins: follow overview unless page exception
- `Scroll hook`: pull to next page (non-final); final = resolution / aftertaste

**Page packing (mandatory):** Group cuts into pages by **reader dwell / breath**, not by filling a generation frame. High dwell (dense text, first exposure, think/feel beats, spectacle) → often **1 cut per page**. Otherwise stack **n cuts top→bottom**. Prefer fewer cuts per page over crowding.

#### Required cut fields (flat list, this order)

Each `##### Cut {n}` (top→bottom) **must** include every field below (value may be `none`):

| Field | Rule |
|-------|------|
| `- **Cut story:**` | Beat this cut alone carries |
| `- **Purpose:**` | May share a line with Size class / Height / Shape using `\|` |
| `- **Size class:**` | `standard` \| `tall`/`long` \| `compact` \| `open`/`diagonal` |
| `- **Height:**` | Target px; must match size class ranges |
| `- **Shape:**` | `rectangle` \| `open` \| `diagonal` \| `bleed` (required) |
| `- **Staging:**` | `{slug} — ref view: …` or `none` |
| `- **Characters:**` | slugs + `(state: …)` or `none (establishing / …)` |
| `- **Location:**` | `{location} / {position} / {view} (state: …)` |
| `- **Direction:**` | expression / gesture / camera / time-weather; seating from staging only |
| `- **Color grade:**` | optional separate line, or fold into Direction |
| `- **Balloons:**` | `none` or indented table (`# \| speaker \| type \| text \| placement`) |
| `- **Captions:**` | `none` or indented table — **field always present** (content optional) |
| `- **Gutter:**` | `{class} {distance} — {intent}` (prefer stating even at episode end / page break) |

Do **not** invent seating in Direction — seating comes from staging. Detailed image prompts belong in Stage ④ YAML only.

```markdown
<!-- `episodes/{nnn}-{episode-slug}.md` template — flat cut schema only -->
# Episode {nnn}: {Title}

## Summary
{what happens, emotional arc}

**Characters:** {slug-a}, {slug-b}, {slug-c} (언급)
**Locations:** {location-a}, {location-b}

## Craft Notes
- Scroll rhythm: {fast / normal / lingering sections — map to size/gutter classes}
- Theme delivery: {shown in scenes — not sermon captions}
- Reader comprehension: {cold reader can follow from art+balloons+captions — yes / gaps to fix: …}
- Climax page.cut: {page.cut} — expect tall/long size class when climax is spectacle
- Page count (this episode): {n}  <!-- MUST equal count of ### Page headings; target ~10–15 unless short-form exception noted -->
- Cut count (this episode): {n}  <!-- MUST equal count of ##### Cut headings in this file -->
- Volume note: {meets ~10–15 pages | short-form exception: …}
- Canvas: width {768px default; lock exact px}, height variable, color, side margins {full bleed / 30–50px}, outside-cut {white/black/theme}

## Pages

### Page {N}

#### Page story
{story this vertical segment conveys}
{Must be imaginable from text alone — basis for Stage ④ message.}
{Only catalog-registered characters/locations}

#### Estimated height
{sum of cut heights + gutters in px; note generation frame / strip-split if needed}

#### Viewport plan
{one phone screen + breath: e.g. 1 cut — linger / 2 cuts — quick exchange / tall alone — scroll-through}

#### Cuts on this page
{n} — dwell: {high|medium|low} — {one-line rationale}

#### Scroll hook
{curiosity/tension to scroll down — final page: resolution/aftertaste}

#### Cuts (top → bottom)

##### Cut {n}
- **Cut story:** {beat}
- **Purpose:** {dialogue / everyday / climax / silence-time / action / establishing / …} | **Size class:** {standard | tall/long | compact | open/diagonal} | **Height:** {e.g. 500px} | **Shape:** {rectangle | open | diagonal | bleed}
- **Staging:** {staging-slug — ref view: establishing|reverse|detail | none}
- **Characters:** {character-slug} (state: {state-slug or base})  <!-- or: none (establishing shot) -->
- **Location:** {location-slug} / {position-slug} / {view-slug} (state: {state-slug or base})
- **Direction:** {expression / gesture / camera tightness / time-season-weather grade}
- **Balloons:**
  | # | speaker | type | text | placement |
  |---|---------|------|------|-----------|
  | 1 | {slug} | speech | {dialogue in target language} | {e.g. top-left}
  <!-- or single line: none -->
- **Captions:**
  | # | text | placement |
  |---|------|-----------|
  | 1 | {optional, target language} | {e.g. top-center}
  <!-- or single line: none -->
- **Gutter:** {tight|normal|wide|pause} {e.g. 200px} — {intent}
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
- Episode file uses the **flat cut schema only** (no `######` cut subsections)
- Every cut’s character **state** exists in catalogs; identity gear does not silently swap
- Every cut’s location **position/view/state** exists in catalogs
- Multi-cut ensembles with fixed seating/formation cite a **staging**; staging cast/location anchors match catalogs
- Cast roster ↔ cut Characters (no ghosts; mention-only tagged)
- Craft Notes **Page count** / **Cut count** == measured `### Page` / `##### Cut` headings
- Page count in ~10–15 range (or Craft Notes short-form exception)
- Size class ↔ Height agree per `cut-composition.md`
- Expression/mood/time/season/weather are cut direction, not new reference states
- Every cut has Balloons + Captions fields (value may be `none`)
- Viewport plan is stated per page (breath named; or per dense stretch)
- Cuts-per-page matches dwell (1-cut pages OK for high dwell)
- Cold-reader comprehension: art + balloons + captions carry the episode (no note-only essentials)

---

## Completeness Check (Stage ②)

- [ ] `illustration-guide.md` exists (art style + vertical chrome + balloon/caption typography) — `workflow/illustration-guide.md`
- [ ] `world-bible.md`, `characters.md`, `locations.md`, `series.md` exist
- [ ] All `characters/{character-slug}.md`, `locations/{location-slug}.md` written
- [ ] `stagings.md` + needed `stagings/{staging-slug}.md` exist for ensemble continuity spans
- [ ] `series.md` lists episodes by **# / Title / Summary only** (no page/cut columns)
- [ ] All `episodes/{nnn}-{episode-slug}.md` use **flat cut fields** (no nested `######` subsections)
- [ ] Cast roster present; no ghost cast; mention-only tagged
- [ ] Craft Notes page/cut counts match measured headings; page volume ~10–15 (or short-form exception noted)
- [ ] Reference rules followed (`workflow/reference-models.md`)
- [ ] Cut size/gutter classes vary with dramatic intent (`workflow/cut-composition.md`); size class ↔ height consistent
- [ ] Variable page height + scroll hooks + viewport plans applied
- [ ] Character / location / staging catalogs ready for Stage ④ Phase 0
- [ ] Story co-lock: every cut’s staging/character-state/location-state is catalogued (reuse or added) — `workflow/reference-models.md` §7

---

## Gate G2 (User Approval)

Do not proceed to Stage ③ until the user approves:
- **`illustration-guide.md`:** art style, chrome, and balloon/caption lettering specific enough that pages won’t drift?
- World rules/background: coherent and specific enough?
- Character catalog: lasting body + outfit/identity-gear states; expressions not treated as states?
- Location catalog: set/stage clear; time/weather not mistaken for location states?
- Staging catalog: each continuing situation (café / OR / meeting / …) has who-is-where + planned ref views?
- Episode/page/cut design:
  - Flat cut schema only throughout the episode file?
  - Episode volume ~10–15 pages (or approved short-form exception)?
  - Cold reader can follow from art + balloons + captions alone (no essential facts trapped in notes only)?
  - Cut stories carry emotion without compressing the episode summary?
  - Staging cited wherever relative placement must not drift?
  - Cast roster clean (no ghosts)?
  - Cut **size classes** match intent and Height?
  - **Gutters** used intentionally?
  - Balloons/captions fields present; reading order and face clearance OK?
  - Viewport plans readable on a phone?
  - Scroll hooks present?
  - Craft Notes page/cut counts match measured headings?
  - `series.md` has titles/summaries only (no page/cut columns)?
  - Reference-image readiness (characters / locations / stagings)?

**Do not proceed until G2 is approved.**
