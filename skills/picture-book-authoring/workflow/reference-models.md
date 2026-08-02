# Reference Models (Identity + Staging Continuity)

**Not a numbered stage.** Canonical rules for what stays locked as reference images vs what is directed per page. Read in Stage ② (catalog design), Stage ③ (integrity checks), and Stage ④ Phase 0 (reference generation).

Reference models raise **visual consistency**. They are not the whole illustration — they are the durable identity and spatial anchors that page generation must obey.

**Shared model:** Same three layers as `webtoon-authoring`, `video-authoring`, and `novel-authoring` (character, location, staging). Continuity unit here = **page**.

## 1. Three layers

| Layer | What it locks | Artifact | Reference images |
|-------|---------------|----------|------------------|
| **Character** | Static look + lasting body change + **equipment identity** | `characters/{slug}.md` states (concrete) | `images/characters/*` (turnaround per state) |
| **Location** | The **set / stage** (static/physical; multi-view 2D; **neutral** — no weather/time drama) | `locations/{slug}.md` position×view×state | `images/locations/*` |
| **Staging** | WHO–WHERE (+ props / ambient / **situation environment** when locked) — episode-first | `stagings/{slug}.md` | **Default: 1× establishing**; optional reverse/detail only (§4.2) |

Expression, mood, transient pose, and **one-off** camera/time/weather are **not** reference models — direct them in page Direction. Situation-spanning environment (rain through a café scene, etc.) may lock on **staging** (§3 / §4).

---

## 2. Character reference model

A character reference locks the **static look** the character must keep for a span of story time — face/body identity **and** the outfit/equipment identity of that span. It is not a performance pose sheet.

**Included (must stay consistent until a new state is defined):**

- Face / body / silhouette identity
- **Lasting physical change**: growth, aging beats, scars, injury aftermath, permanent deformity
- **Equipment identity** (outfit in the broad sense): clothes, armor, weapons, shields, glasses, accessories, signature carried items  
  - Example: if they fight with a specific sword, **do not swap to a different sword across pages** without a designed state change (new weapon = new state)

**Excluded (per-cut direction only):**

- Facial expression
- Mood / atmosphere of the performance
- Transient pose and action (sitting lean, mid-swing, pointing)
- Temporary props that are not identity gear (a random coffee cup for one gag beat — unless it becomes recurring identity)

### 2.1 States in the character md (concrete)

Document under `characters/{slug}.md`, in this order:

1. **`base`** — default outfit + identity gear (the look they maintain until a designed change)
2. **Additional states** — each lasting change gets its own slug and a **concrete written description** in the md (not a one-line label):
   - Growth / aging / accident aftermath / permanent body change
   - Outfit or equipment change (new armor, different weapon, glasses on/off as lasting identity, new accessories, etc.)

Each additional state must spell out what stays the same (face/body anchors) and what differs (outfit/gear/body). Vague “armor version” is not enough.

**New character state / reference image when:** outfit or identity gear changes, or lasting body change occurs.  
**Do not create a new state for:** expression, emotion, or a one-off pose.

### 2.2 md → YAML → reference image (Phase 0)

1. Lock the concrete appearance text in the character md (base + each state).
2. For every state that needs an image, write image-generation YAML by **concretizing that md text** (do not invent gear missing from the md).
3. Generate **one turnaround reference image per state**: a single sheet that shows **front, side, and back** of the same look, plus **key accessories / identity gear** clearly readable. Do not rely on a front-only portrait when later cuts need silhouette or gear from other angles.

Filename: `images/characters/{character-slug}.png` (base) / `images/characters/{character-slug}-{state-slug}.png`.

### 2.3 Visual contrast (mandatory when cast co-appears)

Major characters who share scenes must be distinguishable on **≥2 axes** among: hair style, hair color, face/silhouette shape, signature outfit. Document contrast in `characters.md` / profiles. In staging and cut prompts, append **identification keywords** per character (e.g. `brown curly bob`, `black short straight hair`) so generators do not fuse or swap faces.

### 2.4 World taxonomy (when the bible defines entity classes)

If `world-bible.md` defines species/tech classes (e.g. android vs humanoid, mage vs mundane), every character profile and page Direction must obey that taxonomy. Do not mix class-forbidden traits (e.g. human facial acting on a class that has no face).

### 2.5 Recurring brand / logo marks (optional catalog asset)

Marks that repeat on chests, walls, signs (hospital emblem, faction badge, …) are **not** reinvented per YAML. Create a standalone logo reference first (e.g. `images/locations/{brand}-logo.png` or `images/brand/…`), then inject it as the first `params.references` entry when generating any character/location asset that must show that mark.

---

## 3. Location reference model (= set / stage)

A location reference locks **static / physical** characteristics of the place — the empty set — not who is acting there.

**Included (describe concretely in `locations/{slug}.md`):**

- Physical structure and layout (walls, furniture placement of the set, landmarks)
- Fixed set dressing that defines the place
- Permanent set damage or rebuild (fire gutting, new building next door that stays)

**Examples of what to write (not exhaustive):**

- **Street:** buildings along the block, storefronts/signage, traffic lights, crosswalk, pavement markings, façade materials, fixed street furniture
- **Café:** interior layout, seating types/placement, counter, fixed décor, permanent fixtures and set props that belong to the café as a place

**Excluded from location refs (never bake into location PNG):**

- Time of day, season, weather (rain, fog, sunset drama)
- Dramatic / situational lighting (sirens, red alert wash) — unless a **permanent fixture** of the set
- Transient occupancy (passersby) — **staging** ambient when continuity matters
- One-off camera for a single cut

Location PNGs are **neutral reusable sets** (typically clear daylight / neutral interior light) so many episodes can share them.

**Where dynamic environment goes:**

| Span | Put environment here |
|------|----------------------|
| Holds for the whole continuing situation (e.g. rain outside the café for the whole dialogue) | **Staging** (situation environment) |
| One cut / one beat only | Cut **Direction** |

**New location state / reference image when:** the set permanently changes (new construction, destruction, major redecoration that sticks).  
**Do not create a new location state for:** noon→night, rain, autumn leaves, temporary open curtains alone.

### 3.1 Multi-view 2D (mandatory — one image cannot hold a full 3D space)

A reference PNG only shows what is **in frame**. Plan **multiple 2D views** of the same location:

| Naming style | Examples |
|--------------|----------|
| Cardinal / facing | front, left, right, reverse |
| Fixture-side | window-side, bed-side, door-side, counter-side, hallway-outside-door |

`position` / `view` are framing anchors for **what part of the set is visible**; they are not a substitute for staging (who sits where).

Architecture/fixtures **off-camera in that PNG cannot be invented** in a later cut/page prompt.

**Failure mode:** Cut needs a door/hallway, but the only ref is an interior PNG without a door → inventing the door in YAML breaks images. **Fix:** add a Scene List row (position×view) that shows the door and generate that ref first.

| Rule | Do | Don’t |
|------|----|-------|
| Cut needs element X visible | Cite a `position × view` whose PNG already shows X | Prompt “add a door/wall” missing from the ref |
| Same room, different facing | Add another Scene List row + PNG | Stretch one interior ref to cover opposite walls |
| Character vs fixtures | Place only on stations allowed by staging + cited view | Put someone behind a faucet/sink the view never defined |

One generic `{location}.png` is **not enough** when the episode uses multiple framings of the same room.

### 3.2 md → YAML → reference image (Phase 0)

1. Lock concrete physical description + Scene List rows in the location md.
2. For each needed position×view×state, write YAML by **concretizing that row’s description** (no weather/time drama in location prompts).
3. Generate one PNG per Scene List row.
4. Do **not** prompt architecture/fixtures that are **outside this view’s frame** (avoids hallucinated walls/windows).

### 3.3 Strict Location path cite (hard rule)

Every page’s `- **Location:**` **must** be a catalog path — free-text place names are forbidden.

```text
<!-- Forbidden -->
- **Location:** 한국대학병원 / 별관 로봇 센터 / 외부 전경

<!-- Required -->
- **Location:** {location-slug} / {position-slug} / {view-slug} (state: {base|state-slug})
```

If the episode needs a new angle: **add a Scene List row** to `locations/{slug}.md` (and Phase 0 PNG) **before** citing it in the page. G2/G3 must verify: every page Location ⊆ some Scene List row.

---

## 4. Staging reference model (continuing-situation continuity)

Staging locks the **dynamic story situation** on top of approved character + location refs: named cast placement and action setup, **situation props**, and (when needed) **ambient occupancy** that must not reinvent itself page-to-page.

**Rule of thumb:** If a **situation continues across multiple cuts** and the reader would notice if seats, situation props, **or** the ambient crowd/background density suddenly changed, make **one staging for that situation** and keep it until the situation ends or Design intentionally changes it.

Staging is not “meeting-room only.” It locks **who is where relative to whom**, the **situation’s key props / table dressing**, and optional **ambient extras** while the scene keeps going.

**Same location ≠ same staging.** A street at rush hour vs empty night, or a café during a phone call vs a packed lunch rush, need different stagings even when they cite the same location slug. Prefer defining staging **with the episode** that needs it (§4.1).

**Failure mode (must prevent):** Dinner continues across pages, but the food on the table is a completely different meal each page → looks unfinished. Same for toys on a play mat, picnic spread, classroom desk objects that belong to that scene.

### When staging is mandatory (hard rule)

**Require a staging** when **all** of the following hold:

1. The situation spans **≥2 pages** in the same place/situation, **and**
2. Either **≥2 named characters** appear, **or** relative placement vs set fixtures matters (bench, table, counter, vehicle seats, …)

**`Staging: none` allowed only for:** single-page establishing / empty set / no fixed relative placement needed. Long multi-page ensemble scenes with `Staging: none` are a **hard fail** at G2/G3.

Create a staging whenever multi-page continuity would otherwise drift, for example:

| Situation | What must not flip / teleport page-to-page |
|-----------|--------------------------------------------|
| Meeting / briefing | Who sits in which seat; shared papers/whiteboard state if visible |
| Café / restaurant dialogue | Who is left vs right; cups/plates/order on the table |
| **Dinner / meal** | Seat order **and** food layout (dishes, bowls, shared plates) |
| Picnic / snack on a blanket | Who sits where; picnic spread contents |
| Operating room / procedure | Surgeon / nurse / patient positions; instrument tray layout |
| Classroom / courtroom | Desk/bench order; books/props that define the beat |
| Play / toy scene | Who sits/stands where; toy arrangement on the floor/table |
| Confrontation / standoff | Who stands left / right / center |
| Vehicle / cockpit | Who sits where |
| Street walk-and-talk | Cast order + ambient pedestrians/cars when continuous |
| Stage / performance lineup | Fixed formation |

**Two people opposite at a café:** left/right must stay stable — and their cups/plates should not teleport or morph into unrelated dishes without Design.

**Meal across pages:** seating stays; the **same meal layout** stays (progressive eating/emptying is OK if continuity rules allow; a wholly new menu is not).

Character refs + empty location refs alone do **not** lock relative placement, situation props, or ambient extras. Without staging, generators drift (seat swaps, L/R flips, **food/toys/props reinvented**, crowd/cars morphing).

**One staging per continuing situation.** A later café date, a different dinner, or a reseated meeting → new staging slug (or an explicit Design update), not silent prompt improvisation.

### 4.1 Episode-first staging (preferred)

Prefer defining staging **as part of episode design**, not by default-reusing a catalog staging from another episode just because the location matches.

| Prefer | Avoid |
|--------|--------|
| Episode-appropriate new staging for this story beat / time of day / occupancy | Blind reuse of `cafe-a-b-first-date` for a different night’s phone-call scene |
| Explicit cite when the **same continuity span** truly continues | Assuming “same café location” preserves blocking, props, and ambient |

When drafting the episode:

1. List continuing situations in the episode md (or a linked `stagings/{slug}.md` created for this episode).
2. Write the staging md with blocking + props + ambient for **this** episode’s content.
3. If framings need architecture not yet in the location Scene List, **add location position×view rows** (and Phase 0 refs) before locking cuts — staging often reveals missing location views.

Catalog files still live under `stagings/` (indexed from `stagings.md`); ownership is **episode-driven**.

### What a staging is

A staging composes:

1. A **location** anchor — **catalog path only:** `{location-slug} / {position} / {view} (state: …)`  
2. A **cast list** with each character’s **appearance state** — `{character-slug} (state: …)` must exist in `characters/`  
3. A **blocking map**: seat/spot/station + facing (Viewer's Left / Center / Right)  
4. A **situation props / table dressing map** when continuity matters  
5. **Ambient occupancy** when continuity matters  
6. **Situation environment** when weather/time/dramatic light must hold for the whole span (rain, dusk wash, alert lighting) — **not** baked into location PNG  
7. Continuity rules  

**Cross-verify (hard):** every slug/state in the staging header exists in character/location catalogs. Unknown cast or free-text location → fail Design.

**Examples:** street walk-and-talk with passersby/cars; café phone call with mug + other patrons + optional rain-on-windows for the span.

### 4.2 Canonical staging reference image (default: one establishing)

Staging’s job is to lock **ensemble placement in space** — not to catalog every page camera.

| Allowed as staging ref | Forbidden as staging ref |
|------------------------|--------------------------|
| **`establishing`** (required default — one wide ensemble lock) | `over-shoulder`, `close-up`, `detail` used as a **page/camera angle** catalog |
| **`reverse`** (optional — only if reverse-shot L/R must be locked) | Any one-off cut framing |
| **`detail`** (optional — only to lock a **prop cluster**, not a CU face) | Registering every page’s camera as a staging view |

**Phase 0:** Generate **exactly one** `…-establishing.png` by default. Add reverse/detail only when the staging md lists them with a stated lock purpose.

**Phase 1 pages:** Attach Character + Location + Staging **establishing** under `params.references`. Derive CU / OTS / tighter cameras in cut **Direction** — do **not** invent new staging refs for those angles.

### What staging locks vs what cuts still vary

| Locked by staging | Still free per page |
|-------------------|--------------------|
| Who occupies which seat / station / spot | Expression, mood |
| Relative left/right/center order | Gesture within the spot |
| Facing / orientation of the group | Camera tightness (CU / OTS / wide) **as long as L/R, stations, props stay recognizable** |
| Character appearance state per person | Progressive prop change if Continuity allows |
| Situation props / ambient / situation environment (when locked) | One-off environment only if **not** locked on staging → Direction |

If someone permanently changes place, leaves, or the table is cleared/reset for a new course that replaces the whole layout, either:

- update the staging (new staging slug or documented beat change), or  
- end the staging continuity and start a new staging for the new situation

### Design artifact

`stagings.md` (index) + `stagings/{staging-slug}.md` — authored when the **episode** needs the situation (§4.1):

```markdown
# {Staging title}

## Situation
- Kind: {street walk-and-talk / café phone call / family dinner / …}
- Owning episode: {nnn-…}
- Continuity span: {from page → to page}

## Location anchor (catalog path only)
- Location: {location-slug} / {position-slug} / {view-slug} (state: {base|state-slug})
- Extra location views needed: {none | Scene List rows to add first}

## Cast (catalog binding — required)
- {character-slug-1} (state: {state-name})
- {character-slug-2} (state: {state-name})

## Cast & blocking
| character-slug | appearance state | seat/spot/station (L/C/R) | facing | notes |
|----------------|------------------|---------------------------|--------|-------|
| {slug} | {base\|…} | viewer-left / center / … | toward partner | … |

## Situation props / table dressing (locked)
| prop | placement | description | notes |
|------|-----------|-------------|-------|
| {mug} | right of {slug} | blue mug | … |

## Ambient occupancy (when locked)
| ambient | placement / density | description | notes |
|---------|---------------------|-------------|-------|
| {passersby} | mid-ground | 3–4 pedestrians | … |

## Situation environment (when locked for the span)
- {e.g. rain on windows + dim overcast interior — NOT in location PNG}

## Continuity rules
- Fixed until: {…}
- Allowed progressive prop change: {…}
- Forbidden drift: {no L/R swap; no wholesale prop swap; no ambient teleport}

## Canonical staging reference image
- establishing (required): {one-sentence ensemble lock — who L/C/R + key props/ambient}
- reverse (optional): {only if needed — purpose}
- detail (optional): {only prop-cluster lock — purpose}

{Default Phase 0 output: images/stagings/{staging-slug}-establishing.png only}
```

Slug tip: `ep003-street-walk-talk`, `ep005-cafe-phone-call`.

### 4.3 Lean staging prompts (WHO–WHERE–WHAT)

When writing staging YAML `design` / `prompt`:

- **Do:** WHO & WHERE (L/C/R), WHAT (pose/facing), added props, ambient, situation environment.
- **Do not:** re-describe face, hair, outfit, wall logos, chest emblems already in Character/Location refs (re-describing causes misplaced logos).
- **Do not:** fake negative controls like `NO WINDOWS` — omit off-frame elements instead.
- **Must:** put Character + Location PNGs in `params.references` — never generate staging from text alone.

### Generation (Phase 0)

Order:

1. Brand/logo refs (if any) → Character turnarounds → Location multi-views (incl. staging-required extras)  
2. **Staging establishing** (required) using those PNGs in `params.references`  
3. Optional reverse/detail only if listed in the staging md  

**Staging establishing prompt** shows locked props/ambient/environment and L/C/R blocking — not an empty set. YAML concretizes the staging md.

### Citing stagings in page design

```text
- **Staging:** {staging-slug} — ref view: establishing   # default; reverse|detail only if those PNGs exist
- **Characters:** {slug} (state: …)   # must match staging cast
- **Location:** {location-slug} / {position} / {view} (state: …)   # catalog path; must match staging anchor (or a listed extra view)
- **Direction:** {expression / CU|OTS|gesture / one-off weather if not on staging; progressive prop note}
```

Do not invent seats, L/R flips, or wholesale prop swaps in the page prompt. Update Design first.

### Page image generation (Phase 1)

When a page cites a staging:

1. Attach Character + Location + Staging **establishing** under `params.references` (plus reverse/detail only if cited).
2. YAML `design` keeps seats/L-R + props/ambient; camera tightness from Direction.
3. Do not rebuild the table from Location alone.
4. **design ↔ prompt fidelity:** every proper noun, engraved logo text, and gear detail in `design` must appear in `params.prompt` — no genericizing away.
5. **No metadata on image:** do not put episode title, page/cut numbers, or panel labels in the prompt as drawable text (see `04-generate.md` IMPORTANT block).

---

## 5. Decision cheat sheet

| Situation | Reference layer |
|-----------|-----------------|
| Same face/outfit/sword across the episode | Character state (turnaround) |
| New armor / scar after battle | New character state |
| Co-appearing cast look alike | Character **visual contrast** (§2.3) |
| Recurring hospital/faction emblem | Standalone **logo** ref (§2.5) |
| Angry face / CU / OTS | Page Direction (not a staging ref) |
| Café / street as empty physical set | Location (neutral multi-view Scene List) |
| Free-text location in cut | **Forbidden** — catalog path only (§3.3) |
| Rain for whole café dialogue | **Staging** situation environment |
| One-cut dusk grade | Page Direction |
| Café pair L/R across many cuts | **Staging** (+ establishing PNG) |
| Multi-cut ensemble with `Staging: none` | **Hard fail** (§4 mandatory) |
| Register OTS/CU as staging views | **Forbidden** (§4.2) |
| New episode same café different beat | **New episode staging** |

---

## 6. Completeness (before G2 / G3)

- [ ] Character states concrete; turnaround planned; co-cast **visual contrast** ≥2 axes
- [ ] World taxonomy obeyed in profiles + Direction (when defined)
- [ ] Recurring logos modularized when used
- [ ] Location = neutral static set; no weather/time in location PNG
- [ ] Every page `- **Location:**` is a catalog `slug / position / view (state)` path (§3.3)
- [ ] Scene List covers needed multi-views; no off-frame architecture in prompts
- [ ] Staging mandatory rule satisfied; headers bind catalog character states + location path
- [ ] Staging has establishing-only by default; no cut-angle ref list
- [ ] Props / ambient / situation environment maps when needed
- [ ] Lean staging prompts (WHO–WHERE–WHAT); Character+Location in staging `params.references`
- [ ] design ↔ prompt fidelity for proper nouns / engravings
- [ ] No metadata labels in page prompts (titles/page numbers as on-image chrome)
- [ ] Explicit user approval before every generate

---

## 7. Story design co-locks catalogs (mandatory loop)

Designing the **story** is not separate from designing **staging**, **character states**, and **location states** — **staging especially with the episode**.

### While drafting story units

1. **Staging** — If mandatory (§4) → create/cite episode-appropriate staging (prefer new; reuse only same continuity span). Bind catalog paths + states.
2. **Character appearance state** — cite or add with concrete description + contrast keywords.
3. **Location** — cite existing Scene List row **or add row first**; never free-text.

### Reuse vs add

| Need | Already in catalog? | Action |
|------|---------------------|--------|
| Character look / outfit / gear | Yes / No | Cite state / **Add** concrete state |
| Location path | Yes | Cite `slug / position / view (state)` |
| Location path | No | **Add Scene List row** (+ PNG plan), then cite |
| This episode’s continuing situation | Same span | Cite staging |
| This episode’s continuing situation | Only prior episode same location | **New episode staging** |
| Staging missing for multi-page ensemble | — | **Add** staging before locking cuts |

### Hard rules

- Do **not** invent outfit, seats, L/R, props, ambient, or free-text locations only inside the page. Update catalogs first.
- Staging cast/location must cross-verify to catalogs.
- Expression / one-off camera / one-off weather stay in Direction unless locked as situation environment on staging.
- **md → YAML:** Phase 0 YAML concretizes md; staging always references Character+Location PNGs.
- Never auto-generate without **explicit user approval**.

### Mini checklist (per story stretch)

- [ ] Location cites are catalog paths
- [ ] Mandatory stagings exist; `none` only when allowed
- [ ] Staging headers bind existing states/paths
- [ ] Establishing staging ref planned (not a cut-angle farm)
