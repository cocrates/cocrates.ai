# Reference Models (Identity + Staging Continuity)

**Not a numbered stage.** Canonical rules for what stays locked in Design vs what is free in scene execution. Read in Stage ③ Architecture, Stage ④ Chapter design, and evaluation/continuity checks.

**Shared model:** Same three layers as `webtoon-authoring`, `picture-book-authoring`, and `video-authoring` (character, location, staging). Continuity unit here = **scene** (within episode/chapter).

**Prose medium:** Novels do **not** generate reference PNGs. Reference models are **design locks** — appearance/equipment states, set structure, and blocking maps that prose and continuity must obey. Materialize them as catalog fields and scene cites, not images.

Reference models raise **descriptive and spatial consistency**. They are not the whole scene — they are durable identity and placement anchors that generation must obey.

---

## 1. Three layers

| Layer | What it locks | Artifact | Materialization |
|-------|---------------|----------|-----------------|
| **Character** | Who they look like + lasting body change + **equipment identity** | `characters/{slug}.md` states | Prose + continuity state — no PNG |
| **Location** | The **set / stage** (physical space) | `locations/{slug}.md` (+ optional position/view notes) | Prose description lock |
| **Staging** | **Who is where** relative to whom in a continuing multi-scene situation | `stagings/{slug}.md` | Scene fields cite staging; blocking map in Design |

Expression, mood, transient gesture, time of day, season, weather, and one-off camera-like focus are **not** reference models — direct them in the scene beat / sensory-emotional cues.

---

## 2. Character reference model

**Included (must stay consistent until a new state is defined):**

- Face / body / silhouette identity (as described to the reader)
- **Lasting physical change**: growth, aging beats, scars, injury aftermath, permanent deformity
- **Equipment identity**: clothes, accessories, weapons, shields, signature carried items  
  - Example: if they fight with a specific sword, **do not swap to a different sword across scenes** without a designed state change (new weapon = new state + continuity update)

**Excluded (per-scene direction only):**

- Facial expression in the moment
- Mood / atmosphere of the performance
- Transient pose and action
- Temporary props that are not identity gear

**New character state when:** outfit/identity gear changes, or lasting body change occurs.  
**Do not create a new state for:** expression, emotion, or a one-off pose.

Document states under each `characters/{slug}.md` (e.g. `base`, `armor`, `injured-aftermath`). Continuity (`story-so-far`) must track the active state.

---

## 3. Location reference model (= set / stage)

**Included:** Physical structure and layout; fixed set dressing; permanent set damage or rebuild.

**Excluded (per-scene direction):** Time of day, season, weather; lighting mood as transient; passersby; one-scene framing emphasis.

**New location state when:** the set permanently changes.  
**Do not create a new location state for:** noon→night, rain, autumn alone.

Who sits where across many scenes → **staging**, not location alone.

---

## 4. Staging reference model (continuing-situation blocking)

**Rule of thumb:** If a **situation continues across multiple scenes** (or dense beats) and relative positions must not flip in description, make **one staging for that situation**.

| Situation | What must not flip scene-to-scene |
|-----------|-----------------------------------|
| Meeting / briefing | Who sits in which seat |
| Café / restaurant dialogue | Who is left vs right across the table |
| Operating room / procedure | Surgeon / nurse / patient stations |
| Classroom / courtroom | Desk/bench order |
| Confrontation / standoff | Who stands left / right / center |
| Vehicle / cockpit | Who sits where |

Character profiles + empty location profiles do **not** lock relative placement. Without staging, prose drifts (L/R flips, seat swaps).

**One staging per continuing situation.**

### What a staging is

1. A **location** anchor  
2. A **cast list** with each character’s **appearance state**  
3. A **blocking map**: seat/spot/station + facing  
4. Continuity rules: who may leave/move; when blocking resets  

### What staging locks vs what scenes still vary

| Locked by staging | Still free per scene |
|-------------------|----------------------|
| Who occupies which seat / station / spot | Expression, mood |
| Relative left/right/center order | Gesture within the spot (if Design allows) |
| Facing / orientation of the group | Narrative focus / “camera” emphasis as long as L/R stays true |
| Active appearance state per character | Dialogue, sensory cues, weather/time grade |

### Design artifact

`stagings.md` + `stagings/{staging-slug}.md`:

```markdown
# {Staging title}

## Situation
- Kind: {café dialogue / OR / meeting / …}
- Continuity span: {chapter/episode/scene range}

## Location anchor
- Location: {location-slug}

## Cast & blocking
| character-slug | appearance state | seat/spot/station | facing | notes |
|----------------|------------------|-------------------|--------|-------|
| {slug} | {base\|…} | table-left / … | toward partner | … |

## Continuity rules
- Fixed until: {event}
- Allowed mid-situation moves: {…}
- Forbidden drift: {no L/R swap; no silent reseat}
```

Slug tip: `cafe-a-b-first-date`, `or-appendectomy-team`, `boardroom-budget-meeting`.

### Citing stagings in scene design

In chapter Key Events / scenes:

```text
**On stage:** {cast}
**Staging:** {staging-slug}   # required when relative placement must hold
**Location:** {location-slug}
```

Do not reseat or L/R-flip in prose without updating Design (staging) and continuity.

---

## 5. Decision cheat sheet

| Situation | Reference layer |
|-----------|-----------------|
| Same face/outfit/sword across chapters | Character state |
| New armor / different sword / scar after battle | New character state + continuity |
| Angry face / crying in one beat | Scene direction (not a ref) |
| Café as empty set | Location |
| Night rain outside the same room | Scene direction |
| Room burned; stays ruined | New location state |
| Café pair A left / B right across scenes | **Staging** |
| OR stations across procedure scenes | **Staging** |
| New situation later | **New staging** |

---

## 6. Completeness (before architecture / chapter approval)

- [ ] Every lasting outfit/gear/body change has a character state
- [ ] Identity gear does not silently swap across scenes
- [ ] Location states exist only for lasting set changes
- [ ] Every continuing situation with fixed relative placement has a staging
- [ ] Scenes in that span cite the staging; expressions stay in scene direction
- [ ] L/R, seats, and role stations do not flip without a Design staging update
- [ ] Story design co-locked catalogs: staging + character states + location states designed with the story; existing refs reused; new refs added to catalogs before cite (§7)

---

## 7. Story design co-locks catalogs (mandatory loop)

Designing the **story** (scenes / pages / cuts / clips) is not separate from designing **staging**, **character states**, and **location states**. They are designed together.

### While drafting story units

For each continuing situation and each story beat, explicitly decide:

1. **Staging** — If relative placement must hold across units → cite an existing staging **or** create `stagings/{slug}.md` (one per situation).
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

### Hard rules

- Do **not** invent outfit, gear, set damage, seats, or L/R only inside the story unit. Update catalogs first, then resume story design.
- Staging cast rows must use **existing** (or newly added) character states; staging location anchor must use **existing** (or newly added) location refs.
- Expression / mood / time / weather / one-off gesture stay in the story unit — they are **not** new reference states.
- After adding catalogs, re-check every story unit that should cite them.

### Mini checklist (per story stretch)

- [ ] Continuing situations have staging cites (or explicit “none — single beat / no fixed placement”)
- [ ] Every on-stage character has a catalogued appearance state
- [ ] Every place used has a catalogued set (and state if lasting change)
- [ ] New states/stagings were added to Design files before locking the story unit
