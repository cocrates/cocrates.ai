# Reference Models (Identity + Staging Continuity)

**Not a numbered stage.** Canonical rules for what stays locked as reference images vs what is directed per page. Read in Stage ② (catalog design), Stage ③ (integrity checks), and Stage ④ Phase 0 (reference generation).

Reference models raise **visual consistency**. They are not the whole illustration — they are the durable identity and spatial anchors that page generation must obey.

**Shared model:** Same three layers as `webtoon-authoring`, `video-authoring`, and `novel-authoring` (character, location, staging). Continuity unit here = **page**.

---

## 1. Three layers

| Layer | What it locks | Artifact | Reference images |
|-------|---------------|----------|------------------|
| **Character** | Who they look like + lasting body change + **equipment identity** | `characters/{slug}.md` states | `images/characters/*` |
| **Location** | The **set / stage** (physical space) | `locations/{slug}.md` position×view×state | `images/locations/*` |
| **Staging** | **Who is where** + **situation props** (table dressing, shared objects) in a continuing multi-page scene | `stagings/{slug}.md` | `images/stagings/*` (typically **2–3** views) |

Expression, mood, transient pose, time of day, season, weather, and one-off camera choices are **not** reference models — direct them in the page illustration guide.

---

## 2. Character reference model

**Included (must stay consistent until a new state is defined):**

- Face / body / silhouette identity
- **Lasting physical change**: growth, aging beats, scars, injury aftermath, permanent deformity
- **Equipment identity**: clothes, accessories, weapons, shields, signature carried items  
  - Example: if they fight with a specific sword, **do not swap to a different sword across pages** without a designed state change (new weapon = new state)

**Excluded (per-cut direction only):**

- Facial expression
- Mood / atmosphere of the performance
- Transient pose and action (sitting lean, mid-swing, pointing)
- Temporary props that are not identity gear (a random coffee cup for one gag beat — unless it becomes recurring identity)

**New character state / reference image when:**

- Outfit or identity gear changes
- Lasting body change occurs (growth, accident aftermath, etc.)

**Do not create a new state for:** expression, emotion, or a one-off pose.

---

## 3. Location reference model (= set / stage)

**Included:**

- Physical structure and layout (walls, furniture placement of the set, landmarks)
- Fixed set dressing that defines the place
- Permanent set damage or rebuild (fire gutting, new building next door that stays)

**Excluded (per-page direction):**

- Time of day, season, weather
- Lighting mood (unless the location catalog documents a permanent fixture only)
- Transient occupancy (someone walking by)
- Camera tightness for a single page **within a view that already shows the needed architecture**

**New location state / reference image when:**

- The set permanently changes (new construction, destruction, major redecoration that sticks)

**Do not create a new location state for:** noon→night, rain, autumn leaves, temporary open curtains alone.

`position` / `view` remain framing anchors for **what part of the set is visible**; they are not a substitute for staging (who sits where).

### 3.1 Visibility coverage (mandatory — reference images are not 3D)

A reference PNG only shows what is **in frame**. Architecture, fixtures, and landmarks that are **off-camera in that PNG cannot be invented** in a later page prompt.

**Failure mode (must prevent):** Page needs “bathroom door from the hallway,” but the only location ref is an interior `bathroom.png` with **no door**. Forcing “add a door” in the page YAML invents geometry → broken / inappropriate images. **Fix:** add a location Scene List row (e.g. `hallway / door-looking-in / base`) and generate that reference image first.

| Rule | Do | Don’t |
|------|----|-------|
| Page needs element X visible | Cite a `position × view` whose Description / PNG **already shows X** | Ask page generate to invent X missing from the cited ref |
| Same room, different facing | Add another Scene List row + reference PNG | Stretch one interior ref to cover opposite walls / door / hallway |
| Character relative to fixtures (sink, faucet, tub, door) | Place characters only where the **cited view + staging** allow (e.g. beside tub, not behind faucet unless the view/staging says so) | Improvise “behind the faucet” when the staging/view never defined that spot |

**Design-time check (per page):** List every **must-see architectural / fixture element** in Picture carries / Direction (door frame, sink front, tub side, hallway…). Confirm each is covered by the cited location `position × view` row. If not → **stop**, add Scene List row + Phase 0 ref, then resume. Do not patch in Stage ④ alone.

**Scene List Description must name what is visible**, e.g.:

```text
| hallway-outside-door | door-looking-in | base | Open bathroom door from hallway; door frame + tub + sink visible beyond |
| tub-side | toward-tub | base | Standing beside tub; faucet and tub interior visible; door NOT in frame |
| sink-front | toward-mirror | base | Child-height sink and faucet front-on; mirror above; tub may be out of frame |
```

One generic `{location}.png` without position/view rows is **not enough** when episodes use multiple framings of the same room.

---

## 4. Staging reference model (continuing-situation continuity)

**Rule of thumb:** If a **situation continues across multiple pages** and the reader would notice if seats **or** the table/situation props suddenly changed, make **one staging for that situation** and reuse it until the situation ends or Design intentionally changes it.

Staging is not “meeting-room only.” It locks **who is where relative to whom** **and** the **situation’s key props / table dressing** while the scene keeps going.

**Picture-book failure mode (must prevent):** Dinner continues across pages, but the food on the table is a completely different meal each page → looks unfinished. Same for toys on a play mat, picnic spread, classroom desk objects that belong to that scene.

### When required (default: yes for continuing situations)

Create a **staging** whenever multi-page continuity would otherwise drift, for example:

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
| Stage / performance lineup | Fixed formation |

**Two people opposite at a café:** left/right must stay stable — and their cups/plates should not teleport or morph into unrelated dishes without Design.

**Meal across pages:** seating stays; the **same meal layout** stays (progressive eating/emptying is OK if continuity rules allow; a wholly new menu is not).

Character refs + empty location refs alone do **not** lock relative placement or situation props. Without staging, generators drift (seat swaps, L/R flips, **food/toys/props reinvented each page**).

**One staging per continuing situation.** A later café date, a different dinner, or a reseated meeting → new staging slug (or an explicit Design update), not silent prompt improvisation.

### What a staging is

A staging composes:

1. A **location** anchor (`location` + `position` + `view` + `state`)
2. A **cast list** with each character’s **appearance state** (incl. role gear — via character states)
3. A **blocking map**: seat/spot/station + facing for each participant
4. A **situation props / table dressing map**: key objects that belong to this situation (food layout, cups, toys, shared tools) — what they are and where they sit on the set
5. Continuity rules: who may leave/stand/move; which prop changes are allowed (e.g. bowl empties); when blocking or props reset

Then Stage ④ generates **2–3 reference images** of that ensemble (different useful framings of the **same** blocking **and same prop layout**), and later pages cite the staging + those images.

### What staging locks vs what pages still vary

| Locked by staging | Still free per page |
|-------------------|--------------------|
| Who occupies which seat / station / spot | Expression, mood |
| Relative left/right/center order (incl. facing pairs) | Gesture within the spot (lean, chew, lift spoon) |
| Facing / orientation of the group layout | Camera tightness (close-up vs wide) **as long as L/R, stations, and key props stay recognizable** |
| Which character-state (outfit/gear) each person wears in this situation | Lighting / time / weather grade (unless guide locks otherwise) |
| **Situation props / table dressing** visible in the scene (meal layout, picnic spread, toy arrangement, café cups/plates, tray contents) | **Progressive** prop change only if Continuity rules allow (e.g. soup bowl empties; crumbs appear) — not a wholesale swap |

If someone permanently changes place, leaves, or the table is cleared/reset for a new course that replaces the whole layout, either:

- update the staging (new staging slug or documented beat change), or  
- end the staging continuity and start a new staging for the new situation

### Design artifact

`stagings.md` (index) + `stagings/{staging-slug}.md`:

```markdown
# {Staging title}

## Situation
- Kind: {café dialogue / family dinner / picnic / OR procedure / meeting / play / …}
- Used by episode(s): {nnn-…}
- Continuity span: {from page → to page} (or “until X leaves / reseats / table cleared / situation ends”)

## Location anchor
- Location: {location-slug} / {position-slug} / {view-slug} (state: {base|state-slug})

## Cast & blocking
| character-slug | appearance state | seat/spot/station | facing | notes |
|----------------|------------------|-------------------|--------|-------|
| {slug} | {base\|or-scrubs\|…} | table-left / head-of-table / … | toward partner / toward food | … |

## Situation props / table dressing (locked)
{Objects that belong to this situation — not permanent room furniture (that is location).}

| prop | placement | description (locked look) | notes |
|------|-----------|---------------------------|-------|
| {rice bowl} | in front of {slug} | white bowl, steam | … |
| {shared kimchi plate} | table center | red dish | … |
| {cup} | right of {slug} | blue mug | … |

## Continuity rules
- Fixed until: {event / end of situation / table cleared}
- Allowed mid-situation moves: {e.g. child may stand then return to same seat}
- Allowed progressive prop change: {e.g. bowls empty; bites taken; crumbs — same dishes, depleted}
- Forbidden drift: {no L/R swap; no silent reseat; **no wholesale food/menu/prop swap**; no teleporting dishes/toys}

## Reference image set

Plan the views needed to lock blocking **and** props (often establishing + reverse + detail).

| image slug suffix | Purpose | Framing |
|-------------------|---------|---------|
| establishing | Whole situation: everyone placed + full table/prop layout visible | Wide |
| reverse | Opposite / reverse of the same blocking (keeps L/R + props identity) | Medium-wide |
| detail | Key flank, table surface, or prop cluster | Medium |

{Each row → one PNG under images/stagings/{staging-slug}-{suffix}.png}
```

Slug tip: name by situation, not only by room — e.g. `dinner-family-tuesday`, `cafe-a-b-first-date`, `picnic-park-saturday`.

### Generation (Phase 0)

Order:

1. Character refs (needed states)
2. Location refs (needed position/view/state)
3. **Staging refs** — each image must use the character + location reference PNGs and enforce the **blocking map + situation props map**; produce **one PNG per planned view** in the staging file (commonly establishing, reverse, and/or detail)

Filename pattern:

- `images/stagings/{staging-slug}-establishing.yaml/.png`
- `images/stagings/{staging-slug}-reverse.yaml/.png`
- `images/stagings/{staging-slug}-detail.yaml/.png`  
  (suffixes may vary; list every planned view in the staging file)

**Staging ref prompt must show the locked props** (e.g. the actual meal layout), not an empty table with only people.

### Citing stagings in page design

In episode pages, when the page belongs to a continuing situation:

```text
- **Staging:** {staging-slug} — ref view: establishing|reverse|detail
- **Characters:** (appearance states must match staging cast table)
- **Location:** (must match staging location anchor)
- **Direction:** {expression / gesture / camera tightness / time-weather; progressive prop note if Continuity rules allow, e.g. “A’s bowl half empty”}
```

Do not invent new seats, L/R flips, **or a different meal/prop layout** in the page prompt. If the story needs a reseat, a cleared table, or a wholly new course spread, update Design (staging) first.

### Page image generation (Phase 1)

When a page cites a staging:

1. Attach the matching staging reference PNG(s) under `params.references` (not `params.images`; at least the cited ref view; establishing helps for prop lock).
2. YAML `design` must require: same seats/L-R **and** same situation props/table dressing as the staging ref (only Continuity-allowed progressive change).
3. Do not regenerate the table from scratch from Location alone — Location is the empty/permanent set; staging carries tonight’s meal / this scene’s toys.

---

## 5. Decision cheat sheet

| Situation | Reference layer |
|-----------|-----------------|
| Same face/outfit/sword across the episode | Character state |
| New armor / different sword / scar after battle | New character state |
| Angry face / crying / shouting pose | Page direction (not a ref) |
| Café / OR / meeting room as empty set | Location |
| Same bathroom: door-from-hallway vs tub-side vs sink-front | **Separate location position×view refs** (§3.1) — not one PNG |
| Night rain outside the same room | Page direction (not a location state) |
| Room burned; stays ruined | New location state |
| Café pair: A left / B right across many pages | **Staging** (+ 2–3 ensemble refs) |
| Dinner: same seats + same meal layout across pages | **Staging** (props map + progressive empty OK) |
| Picnic / play mat: same spread/toys across pages | **Staging** |
| OR: surgeon/nurse stations across procedure pages | **Staging** |
| Meeting seats across briefing pages | **Staging** |
| Close-up on one speaker still in their place | Same staging; tighter camera; still cite staging + keep recognizable props |
| New situation later (different dinner / different date) | **New staging** |
| Table cleared; dessert course replaces the whole layout | **New staging** (or documented staging update) — not silent YAML swap |

---

## 6. Completeness (before G2 / G3)

- [ ] Every lasting outfit/gear/body change has a character state
- [ ] Identity gear does not silently swap across pages
- [ ] Location states exist only for lasting set changes
- [ ] **Visibility coverage:** every page’s must-see architecture/fixtures are in the cited position×view PNG (§3.1) — no inventing doors/walls missing from the ref
- [ ] Multi-framing rooms have multiple Scene List rows (not a single generic location PNG)
- [ ] Every **continuing situation** with fixed relative placement **or situation props** has its own staging + 2–3 planned reference views
- [ ] Each staging has a **Situation props / table dressing** map when the situation has shared objects (meal, picnic, toys, café cups, …)
- [ ] Pages in that span cite the staging; expressions stay in page Direction
- [ ] L/R, seats, role stations, **and situation props** do not flip/swap without a Design staging update
- [ ] Story design co-locked catalogs: staging + character states + location states designed with the story; existing refs reused; new refs added to catalogs before cite (§7)

---

## 7. Story design co-locks catalogs (mandatory loop)

Designing the **story** (scenes / pages / cuts / clips) is not separate from designing **staging**, **character states**, and **location states**. They are designed together.

### While drafting story units

For each continuing situation and each story beat, explicitly decide:

1. **Staging** — If relative placement **or situation props** must hold across units → cite an existing staging **or** create `stagings/{slug}.md` (one per situation; include props/table dressing map).
2. **Character appearance state** — Outfit/gear/lasting body for everyone on stage → cite an existing character state **or** add a new state to `characters/{slug}.md`.
3. **Location set state** — Which set (position/view/state as applicable) → cite an existing location row **or** add position/view/state to `locations/{slug}.md`.

### Reuse vs add

| Need | Already in catalog? | Action |
|------|---------------------|--------|
| Character look / outfit / gear | Yes | Cite that **state** (or base) |
| Character look / outfit / gear | No | **Add** state to character catalog (user-visible Design update), then cite |
| Location set / lasting set change | Yes | Cite location (+ position/view/state) |
| Location set / lasting set change | No | **Add** to location catalog, then cite |
| Who-is-where for a continuing situation | Yes | Cite that **staging** |
| Who-is-where for a continuing situation | No | **Add** staging (+ ensure cast states & location anchor exist), then cite |
| Situation props (meal / picnic / toys / café cups) for a continuing situation | Yes | Cite that **staging** props map |
| Situation props for a continuing situation | No | **Add** props map to staging (or new staging), then cite |

### Hard rules

- Do **not** invent outfit, gear, set damage, seats, L/R, **or situation props (food layout, toys, cups)** only inside the story unit. Update catalogs first, then resume story design.
- Staging cast rows must use **existing** (or newly added) character states; staging location anchor must use **existing** (or newly added) location refs.
- Expression / mood / time / weather / one-off gesture stay in the story unit — they are **not** new reference states. Progressive prop depletion stays in Direction only if Continuity rules allow.
- After adding catalogs, re-check every story unit that should cite them.

### Mini checklist (per story stretch)

- [ ] Continuing situations have staging cites (or explicit “none — single beat / no fixed placement”)
- [ ] Stagings with shared objects include a **props / table dressing** map
- [ ] Every on-stage character has a catalogued appearance state
- [ ] Every place used has a catalogued set (and state if lasting change)
- [ ] New states/stagings were added to Design files before locking the story unit
