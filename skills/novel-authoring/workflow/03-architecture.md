# Stage ③ — Architecture (World / Characters / Locations)

**Prerequisites:** Approved `series.md`

**Gate artifacts:**
- `{project-root}/world-bible.md` (+ `world/{aspect-slug}.md` as needed)
- `{project-root}/characters.md` + `characters/{character-slug}.md`
- `{project-root}/locations.md` + `locations/{location-slug}.md`

**Next stage:** `04-episode-design.md` (after user approves **series-level** architecture)

**See also:** SKILL.md § Consistency · § Reference Models (staging is Stage ④).

---

## Purpose

Build the **series-level architecture** so episode design happens **inside** a known whole — not a one-episode stub that is “finished” when ep 001’s props exist.

After Gate G3, a reader of the architecture + `series.md` should roughly know:

1. **What the novel is** — premise, systems, history, themes (`world-bible.md` + aspects), aligned with `series.md` Structure / Episode List
2. **Who appears** — major cast across the arc (`characters.md` + profiles), not only the first episode’s POV
3. **Where it happens** — regions and key places across the arc (`locations.md` + profiles)

**Staging is not Stage ③.** Situation stages (who wears what, situation props, ambient occupancy) are authored in **episode design** when scenes need them — see [`04-episode-design.md`](04-episode-design.md) and SKILL.md § Reference Models.

Episode design (④) **composes** this catalog. Gaps may still appear and are filled **additively** (SKILL.md § Consistency) — that is refinement, not a substitute for skipping series-wide architecture.

---

## Series-level scope (required at G3)

Scan the **entire** approved Episode List (and Structure), not only row 001.

| Layer | Required at Stage ③ (series-visible) | May deepen later at ④ |
|-------|--------------------------------------|------------------------|
| **World** | Core premise, laws, systems, society, history timeline, themes that the **full** arc uses or implies; aspect files for major domains (factions, magic tiers, institutions named across the List) | Fine rules, obscure history slices, one-off devices first needed in a late episode |
| **Characters** | Index + roles for **every series-significant named cast**; profile files with **concrete `base` outfit/gear** (clothes, accessories, weapons as identity). Series-significant people need enough wardrobe baseline that later stagings can cite states | Walk-ons; additional **states** when a new lasting look is needed (add state before staging cites it); deeper voice notes |
| **Locations** | Regions & territories + **Key Locations** that the List / Structure uses across the series; profile files for places that recur or carry arc weight | Room facets, transient sets, places invented only when a later episode needs them |

**Depth bar — test: could a cold reader answer these from architecture alone?**

1. **Premise:** What is the novel about? (world-bible Core Premise + Series Arc Anchors)
2. **Cast:** Who matters across the arc, and what does each want? (characters.md Relationship Map + Character Catalog roles + Core Drives)
3. **Setting:** Where does the story play, and why does each place matter? (locations.md Regions & Territories + Key Locations roles)
4. **Conflict:** What forces collide, and what is at stake? (world-bible systems + character Central Conflicts)

If any answer requires reading an episode design to respond, the architecture is incomplete for G3.

**Not required at G3:** every Episode List row’s crowd, every interior facet, or any staging (situation stage).

**Hard rule:** Do **not** mark architecture complete because “ep 001 only needs one character and one inn.” Episodes are written with series context in mind; Stage ③ must supply that context.

**Index + profile together:** Every slug listed in `characters.md` / `locations.md` must have a matching profile file on disk before G3. **Index-only rows are forbidden.** Characters need a **Character Catalog** row (Display + slug + Profile path); locations need a **Key Locations** row with the same three cells — and each Profile path file must exist.

Characters, locations, and world aspects may be authored **in parallel** after the Episode List is loaded. Prefer batching related writes.

**Dirs:** `world/`, `characters/`, `locations/` should already exist from Stage ① scaffold. If any are missing, create them **together in one step** — not in separate sequential rounds. Do not browse to “confirm empty” before writing profiles.

### Canonical catalog slugs

**Same rule as episode titles.** Filename slugs are kebab-case of the **canonical display name** in the work’s language. Non-Latin names keep their script — **do not translate** into an English alias.

| Artifact | Slug source | Example (Korean work) |
|----------|-------------|------------------------|
| Character profile | Character’s canonical name | e.g. `characters/한서윤.md` — not `characters/한 서윤.md` (spaces) and not `characters/han-seo-yun.md` |
| Location profile | Location’s canonical name | e.g. `locations/옛-항구.md` — not `locations/옛 항구.md` (spaces) and not `locations/old-harbor.md` |
| World aspect | Aspect / domain name as used in the bible | e.g. `world/왕실-의례.md` — not `world/royal-rites.md` |
| Character / location **state** | State label as written in the profile | e.g. `여행-코트` under the character — not `travel-coat` |

**Display vs slug:** Prose and catalog **Display** cells may use spaces (e.g. `옛 항구`, `한 서윤`). The **profile path always uses kebab** (`locations/옛-항구.md`, `characters/한서윤.md` = spaces → `-`). **Never** load / cite a path built by pasting the Display string as a filename.

**Forbidden:** romanization-only aliases, English gloss filenames, bilingual dual files for the same entity, or treating a failed Display-with-spaces path as “profile missing.” If a wrong-language slug already exists: write the canonical path, update indexes/citations, stop using the alias.

Index rows, Architecture References, Appearing / On stage, and Key Events must cite the **same** slug as the profile filename — resolve Display → slug via the **Character Catalog** / **Key Locations** tables, never by guessing.

---

## Procedure

0. **Load (mandatory before any write):** Ensure approved `overview.md` and **full** `series.md` (Structure + entire Episode List) are in this turn’s context — load only if missing. Do not design from chat memory alone when those bodies are absent. Apply SKILL.md § Consistency and § Reference Models as needed.
   - **Initial series architecture (first pass):** writing new catalogs — load overview + series; open existing index/profile files only if already present and you are extending them.
   - **Later additive extend (during ④ gaps):** follow SKILL.md **Selective artifact load** — Phase A indexes (`characters.md`, `locations.md`, `world-bible.md`, `stagings.md`) → Phase B only the profile/aspect being extended. Do not re-dump the full catalog to add one walk-on or facet.

1. **Series pass (mandatory):** From Structure + Episode List, draft working inventories before writing catalogs:
   - **Cast inventory** — every named or role-defined figure the List implies (with first-appearance episode if known)
   - **Place inventory** — regions, institutions, recurring sites
   - **World systems inventory** — laws, factions, information/magic/tech the arc depends on  
   Present these inventories in the architecture files (indexes / bible sections), not as chat-only notes.

### 3.1 World Design

1. Write `world-bible.md` for the **series**, not a single-episode bootstrap boundary:

```markdown
# World Bible: {Title}

## Core Premise
## Physics & Natural Laws
## Magic / Technology / Supernatural Systems (Overview)
## Society & Culture (Overview)
## History (Timeline)
## Thematic Landscape
## Series Arc Anchors
{How P1 / P2 / P3 (or Structure beats) use the above — factions, institutions, mysteries held vs revealed}
```

2. Create `world/{aspect-slug}.md` for **major** domains named or implied across the List (e.g. factions, institutions, technology or magic systems). `{aspect-slug}` = that domain’s name in kebab-case in the work’s language (Non-Latin kept; no English alias). Prefer new aspect files later for truly local rules; do **not** leave series-critical factions undescribed until their debut episode.

### 3.2 Character Catalog Design

1. Create `characters.md` as the Character Web for the **series**:

```markdown
# Character Web: {Title}

## Factions & Groups
- {Faction/Group name}: {beliefs, goals, internal structure, arc role}

## Relationship Map
- {Character A} ↔ {Character B}: {relationship type, power dynamic, conflict seed}

## Character Catalog
| Display | slug | Profile path | Role |
|---------|------|--------------|------|
| {display name — spaces OK} | {kebab-slug} | `characters/{kebab-slug}.md` | {brief role} |

## Character Roles
### {Display name}
- **Role**: {protagonist / antagonist / ...}
- **Archetype**: {optional}
- **Core Drive**: {what they want most}
- **Central Conflict**: {what prevents them from getting it}
- **Arc Direction**: {positive / negative / flat — brief}
- **Series presence**: {arc span / first List row if known}
```

**Character Catalog table is mandatory.** Every series-significant character needs **Display**, **slug**, and **Profile path** in one row. Episode design and eval cite **`characters/{slug}.md`**, not the Display cell. **Character Roles** subsections keep arc detail; headings use Display (or slug) but Architecture References always use the table’s Profile path.

2. For **each series-significant** character, create `characters/{character-slug}.md` matching the table’s **slug** cell (`{character-slug}` = canonical name kebab-case; Non-Latin kept — see [Canonical catalog slugs](#canonical-catalog-slugs)):

```markdown
# {Character Name}

## Basic Info
- Role: {protagonist / antagonist / ...}
- Core Drive: {what they want most}
- Central Conflict: {what prevents them from getting it}
- Arc Direction: {positive / negative / flat — brief description}

## Appearance & Behavior Traits
- Appearance: {face/height/build/unique visual traits — fixed identity}
- Behavior: {signature gestures, posture habits, movement style}
- Voice: {speech patterns, recurring phrases, emotional tell}
- Signature cues: {what readers notice consistently}

## Reference models (appearance + equipment states)
{Canonical: SKILL.md § Reference Models — prose lock only; no PNG}
- base: {concrete default clothes + accessories + weapons/identity gear}
- {state-slug}: {concrete lasting change — what stays vs differs: outfit / accessory / weapon / body}
Rules: state = lasting physical/equipment identity only; expression/mood/transient pose are scene direction. **Staging cites these states** — it does not redefine wardrobe.

## Key Relationships
- {Other Character A} ↔ {Character Name}: {relationship type, power dynamic, key tension seed}
```

**Later at ④:** add walk-ons and new states additively. **Avoid** rewriting core drive/personality of an established character without Modify/Retcon cascade (SKILL.md § Consistency).

### 3.3 Location Catalog Design

1. Create `locations.md` as the Location Map for the **series**:

```markdown
# Location Map: {Title}

## Regions & Territories
- {Region name}: {description, climate, culture, narrative role in the arc}

## Key Locations
| Display | slug | Profile path | Role |
|---------|------|--------------|------|
| {display name — spaces OK} | {kebab-slug} | `locations/{kebab-slug}.md` | {significance; which arc beats use it} |

## Spatial Hierarchy
{Tree or list showing containment relationships}
```

**Key Locations table is mandatory.** Every series-significant place needs **Display**, **slug**, and **Profile path** in one row. Bullet-only lists without slug/path are incomplete for G3. Episode design and eval cite **`locations/{slug}.md`**, not the Display cell.

2. For each **series-significant** location, create `locations/{location-slug}.md` matching the table’s **slug** cell (`{location-slug}` = canonical name kebab-case; Non-Latin kept — see [Canonical catalog slugs](#canonical-catalog-slugs)):

```markdown
# {Location Name}

## Basic Info
- Type: {world / region / city / building / room / outdoor / ...}
- Narrative Role: {why this place matters to the story}
- Role as set/stage: {physical structure that must stay recognizable; lasting damage → new state}
- Key Events: {notable events that happen here across the series — high level}

## Spatial Composition (static / physical — empty set)
- Structure/Layout: {spatial layout}
- Physical detail examples: {landmarks, fixed interior, props}
- Multi-facet anchors (when needed): {front/left/right or window-side/bed-side/door-side — **only these labels are citeable as scene Location facets**; layout / Internal facets / sensory prose mentions are not citeable until listed here}
- Internal facets (optional prose): {named rooms/routes for narration — **not citeable** until promoted into Multi-facet anchors}
- Sense of Scale: {cramped / vast / intimate / ...}

## Sensory Environment
- Lighting / Temperature / Smell / Sound / Texture

## Atmosphere Notes
{psychological impression — safety / threat / nostalgia / ...}
```

**Later at ④:** add new places/facets additively under **Multi-facet anchors**. Lasting set damage → **new location state** slug, not silent drift. Missing citeable facet → **Extend anchors** — do not relocate scenes onto an already-used anchor to dodge the gap.

**Device / signature sets:** If Episode List Summary or overview Validation for early episodes (especially ep 001) names a **core device tied to a room/set** (e.g. signature device ↔ its room/set), prefer locking that set as a Multi-facet anchor (or its own Key Location) at G3 when already known — otherwise ④ **must** Extend before citing. Do not leave ep-001 device rooms as “bootstrap debt.”

---

## Expanding architecture during episode design (Stage ④)

④ designs episodes **within** the approved series architecture. When design discovers a **catalog gap** in world/cast/place (not a substitute for missing G3 series coverage), **stop Key Events** → expand architecture → user approves → resume ④.

| Gap | Preferred action (additive) |
|-----|---------------------------|
| New character | Add `characters.md` row + `characters/{slug}.md` |
| New location / region | Add `locations.md` row + profile |
| New magic / faction / history | Add `world/{aspect}.md` or extend bible with new section |
| New outfit / gear / scar | Add character **state** slug under existing profile (often while authoring a staging) |
| Lasting set change | Add location **state** or new facet |

**Staging** (situation stage: seats, shared props, ambient crowd, cited outfit states) is authored under Stage ④ — not here. See [`04-episode-design.md`](04-episode-design.md).

If the gap requires **modifying** an existing approved rule, classify per SKILL.md § Consistency **Modify / Retcon**.

---

## Completeness Check (series-level — Gate G3)

- [ ] `overview.md` + **full** `series.md` loaded this stage (from context or one load if missing — do not re-read unchanged)
- [ ] Consistency / Reference Models rules applied (SKILL.md)
- [ ] Series pass inventories (cast / place / systems) reflected in catalogs — not ep-001-only
- [ ] `world-bible.md` covers premise + systems + arc anchors for the **whole** Structure
- [ ] Series-significant factions/institutions have aspect coverage or clear bible sections (not deferred solely because they debut after ep 001)
- [ ] Every series-significant character has a **Character Catalog table row** (Display + slug + Profile path) + **profile file on disk at that path**
- [ ] No Character Catalog Display string is usable as a filename (spaces → kebab in **slug** / **Profile path** only)
- [ ] Every series-significant location has a **Key Locations table row** (Display + slug + Profile path) + **profile file on disk at that path**
- [ ] No Key Location Display string is usable as a filename (spaces → kebab in **slug** / **Profile path** only)
- [ ] Catalog / aspect / state filenames use the work’s language kebab-case (Non-Latin kept — no English alias filenames)
- [ ] Relationship map connects major cast used across the arc
- [ ] Location map covers regions/key sites used across the arc
- [ ] A cold reader of `series.md` + architecture can outline what the novel is, who matters, and where it plays — without reading episode designs
- [ ] No `stagings*` authored as Stage ③ deliverables (staging waits for episode design)
- [ ] No scene Key Events exist yet at this stage
- [ ] Architecture is **series-sufficient** for G3; further additive growth at ④ is expected for local detail only

---

## Gate G3 (Architecture Approval)

User approves the **series-level** architecture (world / characters / locations). Do not proceed to stage ④ episode design without approval.

Further architecture additions during ④ each require user approval before Key Events may cite them. Staging files are approved as part of episode design, not G3.
