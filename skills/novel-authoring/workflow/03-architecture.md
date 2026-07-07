# Stage ③ — Architecture Design

**Prerequisites:** Approved `overview.md` and `series.md`

**Gate artifacts:** `characters.md`, `locations.md`, `world-bible.md`, `parts/{nnn}-{part-slug}.md`, `chapters/{nnn}-{chapter-slug}.md`

**Next stage:** `04-chapter-design.md` — only after architecture **and** parent chapter catalog are approved

---

## Procedure

Load `overview.md` and `series.md`. Design the structural backbone. All files here represent the **whole** of their domain — individual details come in profiles (`characters/{name}.md`, etc.).

**Chapter planning lives in part files** — `series.md` provides part catalog only; expand each part into a chapter list here.

**Flexible counts:** Part and chapter counts from `overview.md` / `series.md` are planning targets. Adding, merging, or removing chapters during architecture is normal when story structure demands it — update affected files and `overview.md` Scale with user approval.

**Episodes are not planned here.** Do not specify episode counts in part chapter lists. Episode count and Episode List are decided at **stage ④ chapter design** based on each chapter's story needs.

Characters, locations, world-bible, parts, and chapter catalogs may be designed **in parallel** (after part files define chapter lists).

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

Create `characters/{name-slug}.md` for each character who will appear in early chapters.

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

Create `locations/{name-slug}.md` for each location used in early chapters.

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

## `parts/{nnn}-{part-slug}.md` (Part Design + Chapter Catalog)

**Where chapter planning happens.** One file per row in `series.md` Part Catalog.

Prerequisite: Part appears in `series.md` Part Catalog.

Lists every planned chapter in the part. Chapter count may differ from the approximate count in Part Catalog — update Part Catalog and `overview.md` Scale when it does.

**No episode column** — episode structure is decided per chapter at stage ④.

```markdown
# Part {nnn}: {Title}

## Role in Series Arc
## Core Conflict

## Chapter List
| Ch | Title | Role | Key POV |
|----|-------|------|---------|
| 001 | {Title} | {one-line role} | {character} |
| 002 | {Title} | ... | ... |

## Part Arc
## Part Hooks
- **Opening hook** (chapter {first}):
- **Closing hook** (chapter {last}):

## Character Arcs Advanced
```

---

## `chapters/{nnn}-{chapter-slug}.md` (Chapter Catalog)

**Bridge between part structure and chapter detail design.** Catalog-level chapter info only — **no Episode List here**.

Prerequisite: Chapter appears in parent `parts/{nnn}-{part-slug}.md` Chapter List.

Stage ④ expands this into full chapter design, creates the Episode List, and adds nested episode files based on story composition.

```markdown
# Chapter {nnn}: {Title}

## Part
{parts/{nnn}-{part-slug}.md}

## Role in Arc
## Core Conflict

## Chapter Arc
{Brief — how this chapter fits the part arc}

## Chapter Hooks
- **Opening hook:**
- **Closing hook:**

## Character Arcs Advanced
```

---

## Adjusting Chapter Counts

When architecture reveals a different part structure than planned:

1. Propose change to user with reason
2. Update `parts/{nnn}-{part-slug}.md` Chapter List
3. Update `series.md` Part Catalog approximate counts
4. Update `overview.md` Scale if total shifts significantly
5. Resume — no need to match original numbers exactly

---

## Completeness Check

- [ ] Every character in early chapters has `characters.md` entry + profile
- [ ] Every location in early chapters has `locations.md` entry + profile
- [ ] Each part file has a complete Chapter List for that part's story arc
- [ ] Each early chapter has a catalog file under `chapters/`
- [ ] No Episode List or episode count in part or chapter catalog files
- [ ] No nested episode files exist yet (those are stage ④)
- [ ] No chapter details duplicated in `series.md`

---

## Gate

User approves architecture files (individually or as batch when requested).

**Do not begin chapter detail design until the relevant chapter catalog is approved.**
