# Reference Models (Identity + Staging Continuity)

**Not a numbered stage.** Canonical rules for what stays locked as reference images vs what is directed per clip. Read in Stage ② (catalog design), Stage ③ (integrity checks), and Stage ④ Phase 0 (reference generation).

**Shared model:** Same three layers as `webtoon-authoring`, `picture-book-authoring`, and `novel-authoring` (character, location, staging). Continuity unit here = **clip**.

**Conditional:** Video projects may declare `references.md` → **none** and skip Phase 0 (e.g. slides + narration only). When characters and/or locations are used across clips, apply this file. When a **continuing situation** spans multiple clips, also declare and use **stagings**.

Reference models raise **visual consistency**. They are durable identity and spatial anchors that clip generation must obey.

---

## 1. Three layers

| Layer | What it locks | Artifact | Reference images |
|-------|---------------|----------|------------------|
| **Character** | Who they look like + lasting body change + **equipment identity** | `references/characters/{slug}.md` (when kind declared) states | `images/references/characters/*` |
| **Location** | The **set / stage** (physical space) | `references/locations/{slug}.md` (when kind declared) position×view×state | `images/references/locations/*` |
| **Staging** | **Who is where** relative to whom in a continuing multi-clip scene | `references/stagings/{slug}.md` | `images/references/stagings/*` (typically **2–3** views) |

Expression, mood, transient pose, time of day, season, weather, and one-off camera choices are **not** reference models — direct them in the clip direction guide.

---

## 2. Character reference model

**Included (must stay consistent until a new state is defined):**

- Face / body / silhouette identity
- **Lasting physical change**: growth, aging beats, scars, injury aftermath, permanent deformity
- **Equipment identity**: clothes, accessories, weapons, shields, signature carried items  
  - Example: if they fight with a specific sword, **do not swap to a different sword across clips** without a designed state change (new weapon = new state)

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

**Excluded (per-cut / page direction):**

- Time of day, season, weather
- Lighting mood (unless the location catalog documents a permanent fixture only)
- Transient occupancy (someone walking by)
- Camera angle / framing for a single cut

**New location state / reference image when:**

- The set permanently changes (new construction, destruction, major redecoration that sticks)

**Do not create a new location state for:** noon→night, rain, autumn leaves, temporary open curtains alone.

`position` / `view` remain framing anchors for **what part of the set is visible**; they are not a substitute for staging (who sits where).

---

## 4. Staging reference model (continuing-situation blocking)

**Rule of thumb:** If a **situation continues across multiple cuts** and relative positions must not flip, make **one staging for that situation** and reuse it until the situation ends or the blocking intentionally changes.

Staging is not “meeting-room only.” It is the lock for **who is where relative to whom** (and to the set) while a scene keeps going.

### When required (default: yes for continuing situations)

Create a **staging** whenever multi-clip continuity would otherwise drift, for example:

| Situation | What must not flip clip-to-clip |
|-----------|-------------------------------|
| Meeting / briefing | Who sits in which seat |
| Café / restaurant dialogue | Who is left vs right across the table (facing pair) |
| Operating room / procedure | Surgeon / nurse / patient positions around the table |
| Classroom / courtroom | Desk/bench order |
| Dinner / bar counter | Seat order and facing |
| Confrontation / standoff | Who stands left / right / center |
| Vehicle / cockpit | Who sits where |
| Stage / performance lineup | Fixed formation |

**Two people opposite at a café:** left/right must stay stable across over-shoulder and reverse cuts — that is a staging, even with only two cast members.

**OR scene:** doctor/nurse stations must not teleport around the patient between cuts — that is a staging.

Character refs + empty location refs alone do **not** lock relative placement. Without staging, generators drift (seat swaps, L/R flips, role positions wandering).

**One staging per continuing situation.** A later café date, a different surgery, or a reseated meeting → new staging slug (or an explicit Design update), not silent prompt improvisation.

### What a staging is

A staging composes:

1. A **location** anchor (`location` + `position` + `view` + `state`)
2. A **cast list** with each character’s **appearance state** (incl. role gear: OR scrubs, café casual, etc. — via character states)
3. A **blocking map**: seat/spot/station + facing for each participant
4. Continuity rules: who may leave/stand/move, and when blocking resets

Then Stage ④ generates **2–3 reference images** of that ensemble (different useful framings of the **same** blocking), and later clips cite the staging + those images.

### What staging locks vs what cuts still vary

| Locked by staging | Still free per clip |
|-------------------|--------------------|
| Who occupies which seat / station / spot | Expression, mood |
| Relative left/right/center order (incl. facing pairs) | Gesture within the spot (lean, hand work, stand only if design allows) |
| Facing / orientation of the group layout | Camera tightness (close-up vs wide, OTS) **as long as L/R and stations stay recognizable** |
| Which character-state (outfit/gear) each person wears in this situation | Balloons, captions, color-grade for time/weather |

If someone permanently changes place or leaves for the rest of the situation, either:

- update the staging (new staging slug or documented beat change), or  
- end the staging continuity and start a new staging for the new situation

### Design artifact

`references/stagings.md` (index) + `references/stagings/{staging-slug}.md` (or top-level `stagings/` if preferred — keep one convention per project):

```markdown
# {Staging title}

## Situation
- Kind: {café dialogue / OR procedure / meeting / standoff / …}
- Used by episode(s): {nnn-…}
- Continuity span: {from clip → to clip} (or “until X leaves / reseats / situation ends”)

## Location anchor
- Location: {location-slug} / {position-slug} / {view-slug} (state: {base|state-slug})

## Cast & blocking
| character-slug | appearance state | seat/spot/station | facing | notes |
|----------------|------------------|-------------------|--------|-------|
| {slug} | {base\|or-scrubs\|…} | table-left / surgeon-right / head-of-table / … | toward partner / toward patient | … |

## Continuity rules
- Fixed until: {event / end of situation}
- Allowed mid-situation moves: {e.g. nurse may step to tray then return to same station}
- Forbidden drift: {no L/R swap; no OR role teleport; no silent reseat}

## Reference image set

Plan the views needed to lock blocking (often establishing + reverse + detail — add or omit as the situation requires).

| image slug suffix | Purpose | Framing |
|-------------------|---------|---------|
| establishing | Whole situation with everyone placed | Wide |
| reverse | Opposite / reverse of the same blocking (keeps L/R identity) | Medium-wide |
| detail | Key flank, table side, or OR station cluster | Medium |

{Each row → one PNG under images/references/stagings/{staging-slug}-{suffix}.png}
```

Slug tip: name by situation, not only by room — e.g. `cafe-a-b-first-date`, `or-appendectomy-team`, `boardroom-budget-meeting`.

### Generation (Phase 0)

Order:

1. Character refs (needed states)
2. Location refs (needed position/view/state)
3. **Staging refs** — each image must use the character + location reference PNGs and enforce the blocking map; produce **one PNG per planned view** in the staging file (commonly establishing, reverse, and/or detail)

Filename pattern:

- `images/references/stagings/{staging-slug}-establishing.yaml/.png`
- `images/references/stagings/{staging-slug}-reverse.yaml/.png`
- `images/references/stagings/{staging-slug}-detail.yaml/.png`  
  (suffixes may vary; list every planned view in the staging file)

### Citing stagings in clip design

In episode clip direction guides, when the clip belongs to a continuing situation:

```text
Staging: {staging-slug} (use establishing | reverse | detail as appropriate)
Characters: (appearance states must match staging cast table)
Location: (must match staging location anchor)
Direction: {expression / gesture / camera tightness / time-weather grade}
```

Do not invent new seats, L/R flips, or OR station swaps in the clip prompt. If the story needs a reseat or formation change, update Design (staging) first.

---

## 5. Decision cheat sheet

| Situation | Reference layer |
|-----------|-----------------|
| Same face/outfit/sword across the episode | Character state |
| New armor / different sword / scar after battle | New character state |
| Angry face / crying / shouting pose | Cut direction (not a ref) |
| Café / OR / meeting room as empty set | Location |
| Night rain outside the same room | Cut direction (not a location state) |
| Room burned; stays ruined | New location state |
| Café pair: A left / B right across many clips | **Staging** (+ 2–3 ensemble refs) |
| OR: surgeon/nurse stations across procedure clips | **Staging** |
| Meeting seats across a briefing | **Staging** |
| Close-up on one speaker still in their place | Same staging; tighter camera; still cite staging |
| New situation later (different date / different surgery) | **New staging** |

---

## 6. Completeness (before G2 / G3)

- [ ] Every lasting outfit/gear/body change has a character state
- [ ] Identity gear does not silently swap across clips
- [ ] Location states exist only for lasting set changes
- [ ] Every **continuing situation** with fixed relative placement has its own staging + 2–3 planned reference views
- [ ] Clips in that span cite the staging; expressions stay in cut direction
- [ ] L/R, seats, and role stations do not flip without a Design staging update
- [ ] Story design co-locked catalogs: staging + character states + location states designed with the story; existing refs reused; new refs added to catalogs before cite (§7)

---

## 7. Story design co-locks catalogs (mandatory loop)

Designing the **story** (scenes / pages / cuts / clips) is not separate from designing **staging**, **character states**, and **location states**. They are designed together.

### While drafting story units

For each continuing situation and each story beat, explicitly decide:

1. **Staging** — If relative placement must hold across units → cite an existing staging **or** create `references/stagings/{slug}.md` (one per situation).
2. **Character appearance state** — Outfit/gear/lasting body for everyone on stage → cite an existing character state **or** add a new state to `references/characters/{slug}.md` (or project convention).
3. **Location set state** — Which set (position/view/state as applicable) → cite an existing location row **or** add position/view/state to `references/locations/{slug}.md`.

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
