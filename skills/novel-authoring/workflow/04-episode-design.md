# Stage ④ — Episode Design (`episodes/{nnn}-{episode-slug}.md`)

**Prerequisites:** Approved `overview.md`, `series.md`, and **series-level** architecture (stage ③)

**Gate artifact:** `episodes/{nnn}-{episode-slug}.md` (includes **scene** design sections)

**Next stage:** `05-evaluate-design.md` (design evaluation, **recommended**) or `06-generate.md` (after user approves episode design)

**See also:** SKILL.md § Consistency · § Reference Models. load [`03-architecture.md`](03-architecture.md) when catalog gaps appear.

---

## Design constraints

Episode design **composes** approved `series.md` + series-level architecture + continuity into an **episode arc**, **scene sequence**, and **situation stagings**. Architecture (world / cast / place) is established at stage ③; **staging is authored here** when scenes need a situation stage. Local architecture gaps (walk-on, facet, outfit state) still grow additively — see SKILL.md § Consistency. Do not treat missing major cast/places as “normal bootstrap debt.”

| Rule | Meaning |
|------|---------|
| **Composition, not silent invention** | Before writing scene Key Events, load via **Selective artifact load** (indexes first → Appearing/used/cited details + continuity for ep 002+). Do not redesign from memory; do not dump unused catalogs. |
| **Faithfulness** | Bound to already-approved higher-level artifacts. On conflict: update the higher artifact (with user approval) or revise the episode plan — never silently override. |
| **Generation-ready** | By stage end, the episode file must contain everything stage ⑥ needs to write prose **without inventing** plot beats, scene order, POV, location, seeds, or dialogue intent. If a writer would have to guess, the design is incomplete. |
| **Episode = publication unit** | Scenes are design subdivisions only — not separately published, evaluated as manuscripts, or released. Scenes appear only under `## Scenes`. |
| **Single-file** | Do **not** create `episodes/{nnn}-{episode-slug}/` nested scene files. |
| **Anti-duplication** | Episode-level sections own episode-wide plans (arc, seeds, motifs, hooks). Scene sections own execution (purpose, craft, Key Events). Do not add a "Scene Detail" block that repeats Key Events. Scene Index is TOC only. |
| **Design ≠ manuscript** | Key Events specify concrete dramatic flow (who, where, when, what changes, sensory/emotional cue, dialogue *intent*) — not quoted speech or publishable prose. Stage ⑥ writes the lines. |
| **Catalog gaps** | New character, location, world aspect, or lasting appearance/set **state** → **stop** → expand architecture per [`03-architecture.md`](03-architecture.md) + SKILL.md § Consistency (prefer **add**) → user approves → resume and cite. **New staging** → author under this stage (below); do not treat staging as Stage ③. |
| **Scene count** | Decided here from this episode's story needs — not predetermined in `series.md` or `overview.md`. |

**Volume:** Design enough to carry the episode. Length budget comes from `overview.md` Scale (**default 4,000–8,000 characters per episode**), refined per scene. If the episode summary cannot be told without cramming every scene, **add scenes** (or split into two episodes in `series.md`) — do not compress.

---

## Procedure

0. When catalog gaps appear, load [`03-architecture.md`](03-architecture.md). Apply SKILL.md § Consistency / § Reference Models as needed.
0b. **Next-open episode routing:** If the user asks to design a **released** episode N while `TODO.md` / continuity point at **N+1**, offer next-open (N+1) as the default option; open republication of N only when the user explicitly chooses it.
1. Complete **Pre-Design Load** (below). Do not skip; do not rely on chat memory alone. **Two phases:** (A) indexes → (B) only detail paths this episode needs — never the entire architecture / `episodes/` / `continuity/` catalog.
2. Run **Summary-place facet preflight** (below) **before** drafting Scene Index / Key Events. Missing citeable facets → **stop** → Extend location profile → user approves → resume. **Do not** relocate onto an already-used facet to dodge Extend.
3. Create **`episodes/{nnn}-{episode-slug}.md`** from the `series.md` row — **that exact path only** (see [Canonical episode path](#canonical-episode-path)). **Size discipline:** Key Events stay compact bullets. **Default: one complete write** when the episode fits this turn’s the output limit. Skeleton → per-scene update only if a single write would clearly exceed that budget (or after a length-limit truncate).
4. Draft Purpose / Episode Arc / Scene Index so they **match** `series.md` Summary + Hook to Next. **Signature clauses** of the Summary must land as **concrete Beats in specific scenes** (not diluted across overlapping scenes). Record needed higher-level changes as stop/ask — not silent drift.
5. Under `## Scenes`, write one **complete** `### Scene {n} — {Title}` section per Index row (every mandatory field).
6. Walk scene transitions top-to-bottom: no orphan scenes, no unexplained jumps, **no duplicate Beat/outline sentences across or inside scenes**.
7. Run **Hook alignment check** (below) + **Unit product recompute** (below) + **Design Consistency Gate** (with **Gate Evidence**), then **Generation Readiness**.
8. Present the episode design for approval only if both gates pass — one-line status: *"Load ✅ · Consistency ✅ · Readiness ✅"*. Do **not** present while Locations / Length / Hook Evidence is missing or arithmetic fails.
9. **Recommend** design evaluation (`05-evaluate-design.md`) before manuscript generation — the user may decline and proceed to ⑥.

### File-edit discipline (mandatory)

- Tool name is **load** / **write** / **update** — never load.
- **Structural rewrite** (add/remove scenes, change cast, recalculate length, rewrite most Key Events): if the episode is already in this turn’s context, go straight to **write the full corrected file** (or replace one whole `### Scene` block). Load with load only when that path is missing from context. Do **not** chain 10+ tiny update calls that leave orphan fields.
- **Post-write polish:** if you notice multiple small defects in the text you just wrote (placeholder residue, typo, wrong facet, staging slug), fix them in **one** write or **≤2** update calls — never a drip of 3+ single-line edits for the same pass. **Do not** load the episode again to discover polish items.
- **Critique / Carry → design revision:** if ≥2 scenes change (or ≥3 Key Event fields across scenes), prefer **one** write of the full episode file. Then sync Scene Index · Locations Used · outlines · Closing · Gate G4. Do not leave Gate G4 as “pending re-approval” after the user has already re-approved in chat — update the status block in the same turn.
- **TODO.md:** Snapshot + task checkboxes that change together → **one** update (or one write), not two sequential edits in the same turn.
- **update:** copy the original text verbatim from the latest known file text (prior version or your own last content). Keep it short and unique (task/heading id + 1–3 lines; ≤2000 chars / ≤40 lines). Never paraphrase, invent, or reuse a pre-edit snapshot after a successful edit.
  - **Good:** `- [x] **T-044** — …` + Depends/Notes neighbors; `### Scene 2` + the next field line.
  - **Bad:** bare `- Notes:` / `- Phase:`; remembered paraphrase; garbled chars; a multi-paragraph block as anchor; an already-replaced sentence.
   - On update failure: **stop** → try a shorter unique anchor → retry. Do not overwrite as a failure fallback. Do not guess the next edit from memory.

- After removing a scene or changing Est.: use the text you just wrote (or a successful edit result) to check for **orphan** `Est. length`, Unit budget, and Scene Index rows that no longer match. **Exactly one** `- **Est. length:**` per scene.
- Before claiming gates pass: check header addends, each scene Est., outline density, and that Summary / Hooks / Appearing / On stage name the **same** cast — **and** that Hook evidence strength is identical across Summary / Out / Arc / Seeds / closing Scene Turn (see Hook alignment) — against the **in-context** episode text (the write/update payload you produced). **Do not** re-load solely to claim Load ✅ / Consistency ✅ / Readiness ✅.
- **Unit product recompute (mandatory, no file re-read):** for every scene, multiply each `n×pick` yourself and confirm the **written** Unit-budget `= {subtotal}` equals that product **exactly**. Then confirm Est. ≈ that correct subtotal (nearest-100 round OK). Wrong product in the formula line → Forecast ❌ — fix before presenting. This is arithmetic on the in-context Unit lines, not a verify load.

### Canonical episode path

**One file, one name.** `{nnn}` = Episode List `#` zero-padded to **3 digits**. `{episode-slug}` = that row’s **Title** in kebab-case. Non-Latin titles keep their script in the slug (do not translate the Title into an English alias).

| Episode List | Canonical path |
|--------------|----------------|
| `#` 001, Title `{Title}` | `episodes/001-{title-kebab}.md` |

**Forbidden (rename before approval / ⑤):** `episode-01-design.md`, `ep-1.md`, `001-design.md`, `design.md`, translated aliases that are not the Title slug, or any path other than `episodes/{nnn}-{episode-slug}.md`.

If a wrong-named file exists: write the canonical path, stop using the alias. Do not evaluate or present an alias as the gate artifact.

### Canonical catalog & staging slugs

Architecture profiles and stagings follow the **same Non-Latin rule** as episode titles (see [`03-architecture.md`](03-architecture.md) § Canonical catalog slugs):

| Artifact | Slug source |
|----------|-------------|
| `characters/{character-slug}.md` | Character’s canonical name |
| `locations/{location-slug}.md` | Location’s canonical name |
| `world/{aspect-slug}.md` | Aspect / domain name |
| Character / location **state** | State label as written in the profile |
| `stagings/{staging-slug}.md` | Staging’s short situation label (kebab-case of that label) |

Non-Latin labels keep their script — **do not** invent English aliases (e.g. romanized person names, glossed place names, or `meeting-001` when the work uses another script). Citations in Appearing / On stage / Architecture References / `**Staging:**` must match the filename slug.

---

## Episode file template

```markdown
# Episode {nnn}: {Title}

## Summary
{what happens, emotional arc — from `series.md` row; may be refined}

**Characters:** {slug-a}, {slug-b}, {slug-c} (mention-only if tagged)
**Locations:** {location-a}, {location-b}

## Architecture References
| Type | Artifact | Usage in this episode |
|------|----------|----------------------|
| Overview | `overview.md` | Genre/tone/validation — do not contradict |
| Series | `series.md` | Arc, voice; Episode List row — Summary + Hook to Next |
| Character | `characters.md` + `characters/{name}.md` | Appearing cast only |
| Location | `locations.md` + `locations/{name}.md` | Used places only |
| Staging | `stagings.md` + `stagings/{slug}.md` | Situation stage for this episode’s scenes (authored in ④ — not Stage ③) |
| World | `world-bible.md` / `world/{aspect}.md` | Laws, history, systems touched this episode |

## Prior Design Alignment
{Filled after Pre-Design Load. Required before scene drafting.}

### Load confirmation
- [ ] **Phase A (indexes):** `overview.md`, `series.md`, `characters.md`, `locations.md`, `world-bible.md`, `stagings.md`
- [ ] **Phase B (details):** profiles for **Appearing** cast + **used** places + touched `world/` aspects only
- [ ] Staging details: **N/A** if every scene is `Staging: none`; otherwise **each** cited `stagings/{slug}.md` on disk (index already in Phase A — do not check detail Load ✅ when Staging is none)
- [ ] Continuity set — **Ep 002+ only** (`story-so-far` + **immediate prior** summary only — not every `continuity/*`). **Ep 001: leave unchecked / write N/A — do not mark done.**
- [ ] `episodes/`: only this episode’s design path if editing/resuming — **not** prior episode designs (continuity is authoritative)

### Bound to higher design
| Source constraint | How this episode honors it |
|-------------------|----------------------------|
| Episode List Summary / Hook to Next | |
| Series arc (if applicable) | |
| Character drives / arcs touched | |
| Location / world rules used | |
| Continuity (prior hook + states) | {Ep 001: what this establishes} |

### Conflicts / open questions
{None | list. Any real conflict → stop; escalate; update architecture/continuity first if changing lore.}

## Continuity References
{Ep 001: what this establishes}
{Ep 002+: mandatory}

| Source | What carries forward |
|--------|---------------------|
| `continuity/{nnn-1}-{episode-slug}-summary.md` | |
| `continuity/story-so-far.md` | |

### Prior Hook
{Ep 002+ required — from prior episode closing}

### Threads This Episode
- **Picks up**: {Ep 002+ only — threads that already exist. Ep 001: write `—` / none}
- **Advances**:
- **Plants**: {Ep 001: what this episode establishes belongs here or under Continuity References}
- **Holds**:

## Purpose
{One sentence — what this episode accomplishes}

## Reader Engagement (Episode Level)
### Personal Stake
### Opening Question
{Question the reader carries past the episode opening — not answered immediately}
### Reader Questions
### Stakes

## Episode Arc
{How tension builds and resolves across all scenes}

## Scene Index
| Sc | Title | Function in Episode Arc |
|----|-------|------------------------|
| 1 | {Title} | {one line} |
| 2 | ... | ... |

## Episode Hooks
- **In:** {from prior episode / series opening}
- **Out:** {into next episode / episode close}

## Exposition Budget & Seeds (Episode Level)
### Budget
### Seeds
| Element | Class | Scene | How it appears |
### Hold (do NOT include)

## Literary Craft (Episode Level)
### Motifs (episode-wide)
| Motif | Meaning | Scenes |
### Reader-Discovered Meaning
**What reader should conclude** (design only):
**Hold — do NOT write in manuscript:**
**Closing image** (episode end — no thematic monologue):

## Characters Appearing
## Locations Used
## Tone Notes

---

## Scenes

### Scene 1 — {title}
**POV:** {character} | **Location:** {location-slug[+citeable-facet] — catalog path; facet only if in Multi-facet anchors} | **When:** {story time or relative}
**On stage:** {characters present}
**Staging:** {staging-slug or none — mandatory when ≥2 scenes + (≥2 cast or fixture placement); see SKILL.md § Reference Models}

- **Situation:** {starting state — one line}
- **Beat:** {what happens — causal flow, concrete actions; not script}
- **Turn:** {what changes}
- **Function:** {why this scene exists in the episode arc}
- **Sensory-emotional:** {one world/setting detail} → {POV reaction}
- **Dialogue intent:** {POV/on-stage speakers’ goal + tone; no quotes. Crowd speech → those speakers On stage. Ambient noise only → use Beat/Sensory, keep this `none` or POV-only. Else: none}
- **Transition out:** {handoff to next scene}
- **Paragraph outline:** {5–8 one-line paragraph intents — not finished prose}
- **Unit budget:** {typed n×band formula → subtotal}
- **Est. length:** {characters — must match Unit budget subtotal and outline density}

### Scene 2 — {title}
{…same fields…}
```
---

## Estimated Length
{**Episode target = sum of scene Est. length**. Each Est. is the **intended manuscript length** for that scene (not a floor for Stage ⑥ to exceed). Must be **≥ Scale min** and **≤ Scale max**. Prefer the **central band** of Scale (default 4,000–8,000 → aim ~5,000–7,200) so dramatization has headroom.}

**Length budget (recorded):** Scale min–max: {from overview} · addends: {s1}+{s2}+… · Scene Est. sum: {integer} · Forecast cross-check: ✅/❌ · Pass: Scale min ≤ sum ≤ Scale max ✅/❌ · Target band: central ✅/⚠️

One sum only — it **is** the sum of the scene `Est. length` fields. Do not record a second optimistic total.

---

## Scene schema (mandatory)

Each scene is a **generation brief**. Stage ⑥ must write from these fields + Architecture References alone.

### Canonical schema only

Under each `### Scene {n} — {title}`:

1. Two meta lines using `**Field:**` (colon **inside** bold; may join with `|`):
   - `**POV:**` · `**Location:**` · `**When:**`
   - `**On stage:**` · `**Staging:**` (`none` when not needed)
2. Flat bullets `- **Field:** value` for Situation → Est. length (fixed order below).

**Do not:** nest fields under `####` / extra `###`; omit required fields; use bare `- Location:` or `**Field**:` (colon outside bold); invent a parallel Key Events layout.

Empty optional meaning still needs the field: e.g. `- **Dialogue intent:** none`, `**Staging:** none`.

### Format

```markdown
### Scene {n} — {title}
**POV:** {character} | **Location:** {where} | **When:** {time / relative}
**On stage:** {who is present}
**Staging:** {staging-slug or none}

- **Situation:** {starting state — one line}
- **Beat:** {concrete causal flow}
- **Turn:** {what changes}
- **Function:** {why this scene exists}
- **Sensory-emotional:** {detail} → {POV reaction}
- **Dialogue intent:** {POV/on-stage speakers’ goal + tone, or none — see Cast roster}
- **Transition out:** {handoff}
- **Paragraph outline:**
  1. {one-line intent — opening image/action}
  2. {…}
  (5–8 lines for Est. ≥ ~1,500; fewer only if Est. is truly short)
- **Unit budget:** dialogue {n}×(~200–400) + action {n}×(~150–300) + sensory {n}×(~100–200) + POV {n}×(~100–250) + transition {n}×(~50–150) = {subtotal}
- **Est. length:** {subtotal, rounded; must agree with outline + unit budget}
```

Optional when needed (same `- **Field:**` notation):

```markdown
- **Seed touch:** {Plant/Hint element id — how it surfaces}
- **Motif touch:** {motif id — brief placement note}
```

### Cast roster

**Characters Appearing** = union of scene `**On stage:**` lists (plus explicit mention-only tags). Ghost cast (listed but never on stage without mention tag) is forbidden.

**Speakers ↔ On stage (hard):** Anyone who **speaks** or performs a **named intentional act** in Beat / Dialogue intent / outline **must** be on that scene’s `**On stage:**`. “At least the POV is listed” is not enough.

| Sound / presence | How to design |
|------------------|---------------|
| Named or role-labeled speech (e.g. guests order the host to clear a table; strangers demand names) | Those speakers → **On stage** (catalogued cast, or stop → add profile / mention-only rule). Do **not** put speech in Dialogue intent while On stage is POV-only. |
| Unnamed ambient noise (murmur, footsteps, off-stage voice with no attributable line) | Keep **out** of Dialogue intent as speech. Put in Beat / Sensory-emotional as **non-speech / ambient sound**. Dialogue intent stays `none` or only the POV’s speech goal. |
| Off-stage voice that is still a clear speaker for the scene | Either bring them **On stage**, or reduce to ambient noise (no speech intent), or mark **mention-only** in Appearing **and** state they are off-stage heard — still must not claim Dialogue intent as if they were On stage speakers without roster clarity. |
| Crowd labels (`guards`, “guests”, multilingual equivalents) as Appearing | Forbidden without `characters.md` + profile. Prefer unnamed ambient, or one catalogued representative. |

**Solo / `Staging: none` scenes:** Allowed for single-character focus, but ambient **speech** still triggers the speakers rule above. Do not use Dialogue intent to smuggle a crowd of speakers while claiming a one-person On stage.

### When (story time)

`**When:**` must make the **timeline direction** obvious to a cold reader — especially flashback, body/time shift, or any beat that rewinds or jumps relative to the prior scene.

| Prefer | Avoid |
|--------|--------|
| Clear before/after anchors across a timeline jump (e.g. death → earlier awakening) | Vague “same night, earlier” with no which-night / before-or-after cue |
| Absolute or clearly relative anchors tied to the episode Out hook | Vague “a little later” as the only cue across a timeline jump |

If Scene N ends at time T and Scene N+1 is earlier than T (flashback, rewind, alternate body), When on N+1 must name that jump — not merely “earlier the same night” without stating which timeline.

### Required vs excluded

| Required | Exclude |
|----------|---------|
| POV, Location (from architecture), When, On stage | Full dialogue lines |
| Situation → Beat → Turn → Function | Paragraph-length draft prose |
| Sensory-emotional pair | Exact wording of speeches |
| Dialogue intent when speech occurs (or `none`) | Inner-monologue draft text |
| Transition out | Multi-line quoted exchanges |
| Paragraph outline + Unit budget + Est. length | Invented Est. with no forecast; paste-ready outline lines |
| Staging (`none` or slug) | Alternate scene field layouts |

### Completeness rules

1. **No empty Key Events** — all required fields filled (canonical schema only).
2. **Causal chain** — Scene N Transition out makes Scene N+1 Situation intelligible; episode Out matches next In (or closing hook).
3. **Concrete Beat** — name actions or observable events, not mood-only (“tension rises”).
4. **Still not prose** — if a field (including outline lines) could paste into `manuscripts/` as finished text, cut wording; keep intent.
5. **Length budget** — see [Prose forecast](#prose-forecast) and [Length budget](#length-budget). **Scale min ≤ sum of scene Est. ≤ Scale max.** Est. derived from outline + Unit budget — not chosen first. Prefer central Scale band and mid-low unit multipliers.
6. **Cast roster** — Appearing ↔ On stage union.

Stage ⑥ dramatizes from Dialogue Voices, Sensory-emotional cues, Paragraph outline order, and scene intent — not from pre-written quotes. Stage ⑥ must **hit Est. (±20%)**, must **not** pad to Scale, and must **not** overrun Scale max with confirmation loops. Under-length after honest dramatization → **return to ④**; chronic overshoot with a sound design → compress at ⑥ (or slim forecast at ④ if Est. was packed to Scale max).

---

## Pre-Design Load (mandatory)

Load artifacts at the start of episode design via **Selective artifact load** (SKILL.md). **load only paths not already in this turn’s context** (and again only if a needed path was changed since last receive, or you are resuming after compaction left you without the body). Do not re-read unchanged files “to be sure.”

**Why two phases:** As episodes accumulate, character / location / world / staging catalogs grow. Reading every detail file burns input tokens without helping this episode. Indexes first → decide needs → open only those details.

### Phase A — Indexes (always first)

load these when missing from context (one batch is fine):

| Artifact | Confirm |
|----------|---------|
| `overview.md` | Genre, tone, validation, scale |
| `series.md` | Series arc, voice; **this** Episode List row — Summary + Hook to Next |
| `characters.md` | Cast allowed; relationship map |
| `locations.md` | Places allowed; Key Locations Display → slug → Profile path |
| `world-bible.md` | Premise, laws, history — no contradiction |
| `stagings.md` | Situation index (even if this episode later uses Staging: none — index is cheap; detail files are not) |

**Do not** open any `characters/`, `locations/`, `world/`, `stagings/`, `episodes/`, or `continuity/` detail file until Phase B.

### Phase B — Details (only what this episode needs)

From Phase A + the Episode List row, write the **Appearing slug list**, **used location slug list**, **world aspects touched**, and (if any) **staging slugs to cite**. Then load **only** those paths:

| Artifact | When | Confirm |
|----------|------|---------|
| `characters/{slug}.md` | **Appearing cast only** | Drive, voice, arc — no redefinition |
| `locations/{slug}.md` | **Locations used in this episode only** | Atmosphere, layout; Multi-facet anchors |
| `world/{aspect}.md` | Any system/history slice this episode uses | Already defined |
| `stagings/{slug}.md` | When Staging ≠ none and scene cites that slug | Blocking, props, span — file must exist on disk |
| `continuity/story-so-far.md` | Ep 002+ only | Character/location states, timeline, **Current Unresolved Threads** |
| `continuity/{prior}-summary.md` | Ep 002+ only — **immediate prior episode** | Prior closing; events that must not be undone |
| `episodes/{nnn}-{slug}.md` | Resuming / editing **this** episode only | Current draft body |

**Episode 001:** Skip continuity files entirely. Do **not** check Continuity set as loaded. Document what this episode **establishes** under Continuity References / Prior Design Alignment. Threads **Picks up** stays empty.

**Cited paths must exist:** Every path listed under Architecture References (including `stagings/{slug}.md`) must be present (readable via Load or already in context) before Load ✅. Citing a missing profile while checking Load as done is a gate failure.

**Do not re-read prior manuscripts** — continuity files are authoritative for past plot. **Do not** re-read prior `episodes/*.md` designs “for context.” **Do not** re-read the episode design you just wrote in this turn to tick Load / gates.

Record Load confirmation under **Prior Design Alignment** before writing scene Key Events.

### Load scope (hard — prevents catalog dump)

1. Finish **Phase A** before any profile load.
2. From the Episode List row + planned cast, write the **Appearing slug list** and **used location slug list** (2–6 names typical for ep 001). Prefer slugs from the Key Locations **slug** column — not Display strings.
3. Call load with **only** those Phase B paths (+ world aspects this episode touches + cited stagings). Prefer kebab profile paths from the index tables.
4. **Forbidden:** passing every `characters/*.md`, `locations/*.md`, `world/*.md`, or `stagings/*.md` “for context.” Mid/late-arc cast and unused Key Locations stay unread until they Appear or are used. **Forbidden:** load `locations/{Display with spaces}.md`.
5. **Forbidden:** batch-reading all of `episodes/` or all of `continuity/` — only this episode’s design (if needed) + (ep 002+) `story-so-far` + immediate prior summary.
6. Need a walk-on later → stop → add catalog → then load **that** profile. Do not pre-load the whole roster.

### Summary-place facet preflight (mandatory before Scene Index)

After Load, **before** drafting Scene Index / Key Events:

1. From `series.md` Episode List **Summary** + **Hook to Next** (and planned scene places), list every **place noun** that will appear as a scene Location facet or as a set where a signature device/object sits (e.g. device room, corridor, gate — not the Key Location slug alone).
2. For each noun, open the used `locations/{slug}.md` and check **Multi-facet anchors** (the only citeable facet list — see [Citeable facets](#citeable-facets-hard)).
3. **Layout / Physical detail / Internal facets / Sensory prose mentions do not count** as citeable facets. A word that only appears in Structure/Layout or Internal facets (e.g. “corridor”, “약방”) is **not** locked until it is listed under Multi-facet anchors.
4. Any missing noun → **stop** → additive Extend: add the facet under Multi-facet anchors (and enough physical/sensory detail to stage the beat) → user approves → then cite in scenes.
   - **Forbidden as default:** relocate the scene onto an already-used Multi-facet anchors label (especially the **previous** scene’s facet) solely to avoid Extend. That creates adjacent same-facet scenes and hides the catalog gap.
5. **Ep 001 / device places:** if Overview Validation or Summary names a core **device + set** (device→set, letter→archive, etc.), that set **must** be a citeable facet (or its own Key Location) before Key Events. Do not invent the room only inside Beat/Transition.

Record the preflight result in Prior Design Alignment Conflicts (None | list of facets added / still blocked). Do not draft scenes while any planned facet fails this check.

### Citeable facets (hard)

Scene `**Location:**` form: `{Key-Location-slug}` or `{Key-Location-slug} + {facet}`.

Resolve slug from `locations.md` **Key Locations** table (`Display` → `slug` → `Profile path`). **Never** build a profile path by pasting the Display cell (spaces) as a filename.

| Allowed `{facet}` | Not allowed |
|-------------------|-------------|
| Exact label from that profile’s **Multi-facet anchors** (or a facet just added there and approved this turn) | Free-text room/area invented in Key Events |
| | Word that only appears in Structure/Layout, Physical detail examples, **Internal facets**, Sensory, or Atmosphere |
| | Alias / paraphrase of an anchor (must match the anchor string) |
| | Relocating to a sibling already-used anchor to dodge Extend |

**Key Location alone** (no `+ facet`) is OK when the whole-place framing is enough and no interior angle is named in Location / Beat as a distinct set. As soon as the design names a specific interior/angle in `**Location:**` or relies on it as the scene’s set, that label must be a citeable facet.

Missing facet → **Location facets** ❌ **and** No silent lore invention ❌ (uncatalogued place-as-set). Path miss (wrong filename / unread profile) is a **separate** path check — do not record it as “not in Key Locations” when the index row exists. Fix via stop → Extend — do not self-pass Locations ✅.

### Hook alignment (mandatory before Consistency Gate)

Quote `series.md` **Hook to Next** (and Summary signature clauses). Then verify **one evidence strength** everywhere the design restates that obligation:

| Surface | Must match Hook strength |
|---------|--------------------------|
| Episode Summary | Same claim (not softer, not stronger with new plot) |
| Episode Out / last Transition out | Same claim |
| Episode Arc closing beat | Same claim |
| Seeds `How it appears` for the Hook element | Same claim |
| Closing scene Turn + Dialogue intent | Same claim — **observable** action/speech |

**Fails (do not self-pass Hook ✅):**
- Progressive softening: header says “knows the old name” but Scene Turn / Seeds say “first syllable / almost says / hint only”
- Progressive hardening or **scope creep:** Transition/Out adds chase, faction arrival, or second reveal **not** in Summary **or** Hook to Next (e.g. Hook = recognition only, but Out adds a pursuit beat absent from Summary+Hook)
- Crowded Out: **>2 independent** next-episode obligations stacked in the last Transition (name + ticket + chase + residual cost → pick the Hook obligation + at most one supporting plant)

**overview Constraints:** if overview lists a **signature line** (core dialogue), either place it in Dialogue intent for the owning scene **or** Hold it with an explicit reason. Do not leave it unused and unmentioned.

**Profile-backed claims:** if a character “knows / recognizes / reveals” a fact about another, that relationship or knowledge must appear in the character profile (Relationships / Drive / Arc) **or** Conflicts must escalate “profile lacks basis → stop/add architecture.” Do not plant identity threats with zero catalog support.

### MAY / MUST NOT

**MAY:** Select approved characters/locations; apply world rules; plan scenes, hooks, engagement in one file — within Episode List Summary and continuity.

**MUST NOT:**
- Start scene drafting before Pre-Design Load is done
- Start Scene Index / Key Events before **Summary-place facet preflight** passes (or Conflicts escalates a blocked facet)
- Cite a Location `+ facet` that is not an exact **Multi-facet anchors** label (layout/sensory mentions alone do not qualify)
- Mark Continuity set loaded on ep 001
- Mark staging Load ✅ when every scene is `Staging: none` (use N/A)
- Load every architecture profile “just in case” — Phase A indexes → Phase B Appearing / used / cited only
- Batch-load the entire `characters/`, `locations/`, `world/`, or `stagings/` directory for a single-episode design
- Batch-load all of `episodes/` or all of `continuity/` (this episode + story-so-far + immediate prior summary only)
- Cite `stagings/{slug}.md` (or any Architecture path) that is missing on disk
- Leave unfilled template braces in the gate file (`{TOC only…}`, `{Ep 001: what this establishes}`, other `{…}` placeholders as body text)
- Name anyone in Summary / Hooks / Seeds / Closing who is not in Characters Appearing (or tagged mention-only) and On stage where they act
- List crowd labels (`guards`, “guests”, “townsfolk”, multilingual equivalents) under Characters Appearing without a `characters.md` row + profile (stop → add catalog, or keep them unnamed extras with no Appearing entry and **no speech** in Dialogue intent — ambient only)
- Put speech goals for anonymous guests / off-stage voices in Dialogue intent while On stage lists only the POV
- Use vague When across a timeline jump (“same night, earlier”) without stating which timeline
- Invent named characters, locations, factions, or world rules (stop and add catalogs first)
- Expand profiles beyond architecture files
- Contradict `story-so-far.md` or drop active threads without a Hold reason
- Override Episode List Summary / Hook to Next
- Soften Hook evidence in Scene Turn/Seeds while Summary/Out keep the full Hook claim (or the reverse)
- Add Out/Transition beats that invent chase or reveal obligations absent from Summary **and** Hook to Next
- Plant “character knows X” with no profile relationship/knowledge basis and no Conflicts escalation
- Ignore overview signature lines without Dialogue intent placement or explicit Hold
- Steal a later Episode List payoff into this episode’s Beats (foreshadow/Hint only when the List schedules payoff later)
- Contradict series arc, world-bible, or profiles without an approved higher-level update
- Create nested scene files or a top-level `scenes/` directory
- Use a non-canonical episode filename
- Duplicate scene plot as episode-level "Scene Detail"
- Pull cast from a **later** Episode List row into an earlier episode without architecture stop/add
- Paste the same Beat / Situation / Turn / outline / Unit budget sentence across scenes (or 3+ times inside one scene)
- Split one Summary signature beat across two scenes that both claim the same reveal
- Approve a meta-only scene (no dramatic event — only “this leads to the next episode”)
- Hand off to stage ⑥ when Design Consistency Gate or Generation Readiness fails
- Ship Scene Index–only designs (TOC without Key Events)
- Claim Estimated Length Pass when header addends ≠ scene `Est. length` fields
- Record **two** `- **Est. length:**` lines in one scene, or leave Est./Unit/Index rows from a deleted scene
- Record Unit budget as a **range** (`1,650–2,150`) instead of a single `n×pick = subtotal` integer formula
- Write a Unit budget `= {subtotal}` that does not equal the arithmetic of its `n×pick` terms
- Mark Locations / Length / Hook Consistency ✅ without **Gate Evidence** lines (below)
- Mark Length ✅ when Evidence lacks per-scene **written vs recomputed** pairs, or when any written ≠ recomputed
- Write Hook Gate Evidence that softens/hardens relative to Episode Out / closing Turn / Seeds (checklist meta must match body quotes)
 - Chain many failed/partial update calls without retrying with shorter anchors (or guess/overwrite); call a non-existent load tool

 - Chain many failed/partial update calls without retrying with shorter anchors (or guess/overwrite)

- Re-load a just-written episode solely to claim gates pass
- Drip-fix post-write defects with 3+ tiny updates in one gate pass instead of one rewrite / ≤2 edits

---

## Prose forecast

Anticipate **what manuscript paragraphs will exist** before assigning Est. length. Evaluation and Generation Readiness judge the forecast — not a bare number.

**Est. = intended published length for that scene.** Stage ⑥ aims at Est., not “Est. and then some.” Fat forecasts (every unit at band-high, every scene at max density) cause manuscripts to overrun Scale max.

### A — Paragraph outline (mandatory per scene)

List **5–8 one-line paragraph intents** in manuscript order (shorter scenes with Est. ≪ 1,500 may use 3–4). Prefer the **fewest lines that still cover** Opening → Beat → Turn → Transition — do not pad outlines to inflate Est.

| Rule | Detail |
|------|--------|
| One line = one expected manuscript paragraph | Image, action, dialogue volley, or POV beat — not a chapter summary |
| **Not finished prose** | No quotable speeches; no paste-ready sentences |
| Covers the arc | Opening → Beat chain → Turn → Transition out appear in the outline |
| Density check | Rough guide: ~200–350 characters per outline line (prefer ~220–280). Est. 3,000 with only 3 outline lines → fail |

**Intent-only example (generic):**
1. Opening sensory detail that grounds place and body
2. Antagonist pressure + POV silence / glance
3. Brief memory flash that motivates the choice (short)
4. Hesitation at the decisive action (no thematic speech)
5. Exit motion that sets the Out hook location

### B — Unit budget (mandatory per scene)

Count expected **units by type**, multiply by the guidance band, sum = **Unit subtotal**. Set **Est. length ≈ subtotal** (round to nearest 100 OK).

**Product exactness (hard):** the integer after `=` on the Unit budget line **must equal** the arithmetic of the typed `n×pick` terms **exactly** (e.g. `3×200 + 2×130 + 3×170 + 1×90 = 1,460` — writing `= 1,480` fails even if Est. is “close”). Recompute before gate. **Nearest-100 rounding applies only to Est. length**, never to the formula’s stated product.

**±20% rule (Est only):** Est. may sit within ±20% of the **correct** Unit subtotal (and must sit inside the outline density band). ±20% does **not** excuse a wrong product written on the Unit line.

**Five types only.** Do not invent labels (`meditation`, `inner monologue`). Map inner thought → **POV beat**; map “Character speaks” → **Dialogue exchange**.

| Unit | What it is | Guidance characters / unit |
|------|------------|----------------------------|
| Dialogue exchange | One A↔B volley (or one charged line + beat) | 200–400 |
| Action beat | Observable physical/dramatic move | 150–300 |
| Sensory beat | World detail on-page (pairs with Sensory-emotional) | 100–200 |
| POV beat | Inner/outer gap, memory flash, decision **shown** | 100–250 |
| Transition beat | Handoff out of scene | 50–150 |

**Default picks:** use **mid-low** of each band (e.g. dialogue ~250, action ~180, sensory ~120, POV ~140, transition ~80) unless the Beat truly needs density. Do **not** default every multiplier to band-high — that stacks episode Est. near Scale max and leaves no room for natural dialogue rhythm at Stage ⑥.

**Required formula** (integers, not a pile of identical ~300s):

```text
dialogue {n}×{pick 200–400} + action {n}×{pick 150–300} + sensory {n}×{pick 100–200} + POV {n}×{pick 100–250} + transition {n}×{pick 50–150} = {subtotal}
```

Omit a type only if **n = 0** (write `dialogue 0`). **n** must be countable from Paragraph outline + Dialogue intent + Beat.

**Invalid:**
- Untyped padding (`~300+300+…`) with no type×count
- Type names outside the five
- **Range-only budgets** (`= 1,650–2,150` or `~1,800–2,000` as the Unit budget field) — must be one integer subtotal from picked multipliers
- **Wrong product:** written `= {subtotal}` ≠ sum of `n×pick` terms (even by 20 characters)
- `Dialogue intent: none` while outline/Beat contains speech
- Claiming Est. much higher than typed mid-band sum
- **Outline density fail:** Est. outside (outline line count × 200–350), e.g. 9 outline lines → band ≈ 1,800–3,150 but Est. 5,300 → fail even if Unit line equals 5,300
- **Duplicate Est. lines** in one scene, or an Est. that belongs to a removed scene still present in the file
- Packing every scene at band-high solely to “use” Scale max

**Worked example:** `dialogue 2×250 + action 3×180 + sensory 2×120 + POV 2×140 + transition 1×80 = 1,540` → Est. **1,500**. Same beats with all band-high picks claiming Est. ~2,400 without extra outline density → prefer the leaner Est. Five outline lines claiming Est. 5,000 without matching unit density → fail.

### C — Cross-check (independent, not circular)

1. **Unit subtotal** — multiply every `n×pick` yourself; **written product must match exactly** — **primary** Est. source
2. **Outline density band** = outline line count × 200–350 — range check only, not Est.
3. **Est. length** ≈ **correct** unit subtotal (round to 100 OK; ±20% of correct subtotal)
4. Est. must lie **inside** the outline density band

**Circular check (always fail):** setting Est. = outline lines × 300, then writing a unit line of identical ≈300 addends so it “matches.”

If they disagree → add real dialogue/action/sensory units (or split scenes) or **lower Est.** Never raise Est. alone. Never waive Scale min as “pilot / slightly under.” Never raise Est. to fill Scale max without new Beats. Never “fix” a wrong Unit product by changing only Est.

### D — Episode arithmetic

Write **integers** and **re-add** them in Estimated Length:

```text
Scene 1 Est. {n} + Scene 2 Est. {n} [+ …] = {sum}
Scale min {m} → sum ≥ m ? ✅/❌
Scale max {M} → sum ≤ M ? ✅/❌
```

Header addends **are** the scene Est. fields copied up, then re-added. Two different lists (header vs scene fields) → fail.

Episode **Estimated Length** = that **recomputed** sum after each scene passes A+B+C.

**Target band (default Scale 4,000–8,000):** prefer sum ≈ **5,000–7,200**. Sum at Scale max with no headroom → ⚠️ (still pass if ≤ max) — Stage ⑥ has little room and often overshoots; slim units/outline or accept a shorter central-band Est.

---

## Length budget

Scale is an **episode** band. Publishable volume is planned **per scene** via Prose forecast, then summed. Catch under-design **and over-packed forecasts** here, not only at manuscript Floor.

| Layer | Rule |
|-------|------|
| `overview.md` Scale | Episode character band (default **4,000–8,000**). **Min** and **max** bound Est. sum. No pilot exception. |
| Prose forecast | Paragraph outline + **typed** Unit budget **before** Est.; prefer mid-low multipliers |
| Per-scene Est. length | ≈ unit subtotal; **intended manuscript length**; must sit in outline density band |
| Sum of Est. | **Re-add integers.** **Scale min ≤ sum ≤ Scale max.** Prefer central band (~5,000–7,200 on default Scale). |
| Episode Estimated Length | Equals the recomputed sum (record addends + sum + pass/fail) |
| Beat justification | Thin Beat / thought-only scene + large Est. → add dialogue/action or split scenes |
| Stage ⑥ | Follow Paragraph outline order; dramatize to Est. (±20%); **never** pad; **never** overrun Scale max with repetition |

**Example:** Scale 4,000–8,000; Est. 1,900 + 1,800 = **3,700 < 4,000 → fail**. Three forecast-backed scenes 2,000 + 2,000 + 2,000 = **6,000 → pass** (central band). Est. 2,700×3 = **8,100 > 8,000 → fail** even if “almost max.”

**Do not:** invent Est. to hit Scale; circular outline×300; mis-add the sum; waive min for a pilot; stack band-high units until Est. touches Scale max “because we can.”

---

## Design Consistency Gate (mandatory — before approval)

Any ❌ blocks approval and handoff to stage ⑥.

### Gate Evidence (hard — anti rubber-stamp)

For the rows below, a bare ✅ with no Evidence line is **invalid** — treat as ❌. Write Evidence in the episode file under `## Design Consistency Gate` (one short line per check) or as inline notes next to those checks. **Do not** mark Consistency ✅ in chat or the status line while Evidence is missing.

| Check that needs Evidence | Evidence must show |
|---------------------------|-------------------|
| **Locations (index)** | Every scene place maps to a Key Locations **Display** or **slug** (prefer table). Format: `Sc1: {slug} ∈ Key Locations; …` — **not** path existence |
| **Locations (path)** | Every Architecture References `locations/{slug}.md` (and used character/staging paths) **exists** — quote path; at ④ after write, success counts; at ⑤ see eval path discipline |
| **Locations (facets)** | If `+ facet`: `Sc1: {slug}+{facet} ⊆ anchors; …` — only after profile body is known. Layout/sensory-only ≠ anchors |
| **Length / Prose forecast** | Per-scene **two-sided** product line (below). Est ≈ **correct** subtotal; header addends = scene Est sum; Scale min ≤ sum ≤ Scale max |
| **Hook to Next / Closing** + **Hook internal consistency** | Quotes from **canonical body surfaces only** (below) — same claim strength as `series.md` Hook |
| **Episode List Summary** | Quote Summary signature clause(s) → named Scene Beat that executes each |

#### Length Evidence — written vs recomputed (hard)

For **every** scene, Length Evidence must show both sides on one line:

```text
Sc{n}: written ={w}; recomputed ={r}; Est={e}
```

- `{w}` = the integer after `=` on that scene’s Unit budget line  
- `{r}` = your independent product of the same `n×pick` terms (**exact**)  
- Length / Forecast ✅ only if **every** scene has `{w} = {r}` **and** Est ≈ `{r}` (±20% / nearest 100)  

**Forbidden rubber-stamp:** listing only `{w}` (or copying the Unit line) and marking ✅ without a distinct `{r}`. If `{w} ≠ {r}`, fix the Unit line **before** claiming Consistency ✅ — do not publish Evidence that asserts a wrong product as verified.

#### Hook Evidence — body surfaces vs gate-meta (hard)

**Canonical Hook surfaces** (these decide Hook pass/fail):

1. `series.md` Hook to Next  
2. Episode Summary (closing obligation if restated)  
3. Episode Out  
4. Episode Arc closing beat  
5. Seeds `How it appears` for the Hook element  
6. Closing scene **Turn** / Dialogue intent / Transition out  

Gate Evidence for Hook **must quote those body surfaces** (and the series Hook). It must **not** invent a softer or harder paraphrase than the body.

| Situation | Result |
|-----------|--------|
| Body surfaces match series Hook at one strength | Hook ✅ — even if you still need to tidy checklist wording later |
| Body surfaces drift (soften/harden) vs Hook or each other | Hook ❌ |
| Body matches Hook, but the **checklist Evidence line** alone softens/hardens | Hook body still ✅; **do not** self-pass until the Evidence line is rewritten to match the body quotes — treat as Gate Evidence hygiene fail until fixed |

**Rule:** checklist meta never outranks the body. Do not write Gate Evidence that weakens an obligation the Out/Turn already state at full strength (and do not harden beyond Hook/Summary).

Other gate rows may stay checkbox-style, but Locations / Length / Hook / Summary **never** self-pass without Evidence.

| Check | Pass criteria |
|-------|----------------|
| Loaded required artifacts | Pre-Design Load complete (**Phase A indexes → Phase B Appearing/used/cited only** — not full catalog dump); **every Architecture Reference path exists on disk**; ep 001 Continuity unchecked; **Summary-place facet preflight** done |
| Canonical path | File is `episodes/{nnn}-{episode-slug}.md` — not an alias |
| Catalog / staging slug language | Character, location, world-aspect, state, and staging filenames match canonical names in the work’s language (Non-Latin kept — no English aliases) |
| Episode List Summary | Purpose + Episode Arc + scene Beats execute **every clause** of the approved Summary; signature beat sits in a **named scene** with concrete actions; **Evidence quotes Summary → Scene** |
| Hook to Next / Closing | Episode Out **and** closing Scene Turn/Dialogue intent/Seeds match `series.md` Hook to Next at the **same evidence strength** (no soften / harden drift); refine wording but do not drop the obligation; **Evidence quotes Hook + Out + Turn** |
| Hook scope (no creep) | Last Transition / Out does **not** invent chase, faction arrival, or second reveal absent from Summary **and** Hook; ≤2 independent Out obligations |
| Hook internal consistency | Summary / Episode Arc close / Out / Seeds How it appears / closing Turn all state the **same** Hook claim — disagree → fail (do not mark Hook ✅); **no ✅ without quotes in Evidence** |
| Overview signature lines | Constraints signature dialogue appears in Dialogue intent **or** explicit Hold — not silently omitted |
| Profile-backed knowledge | “Knows / recognizes X” claims have profile Relationships/Drive support **or** Conflicts escalated |
| Opening honors prior | Ep 002+: Prior Hook addressed in Scene 1; no soft-reset |
| Continuity states | Ep 002+: Character/location states match `story-so-far`; no undo without approved continuity update |
| Unresolved threads | Each active thread is Picks up / Advances / Plants / Holds — none silently dropped; ep 001 has no false Picks up |
| Characters | Every On-stage name exists in architecture; voice/drive matches profile; arc not reversed without architecture update; **every Dialogue intent / Beat speaker is On stage** (ambient noise ≠ speech) |
| Locations | **Three Evidence lines (do not merge):** (1) every scene Location ∈ Key Locations table Display/slug; (2) every cited `locations/{slug}.md` path exists (Architecture References — never build path from Display spaces); (3) every `+ facet` is an exact **Multi-facet anchors** label ([Citeable facets](#citeable-facets-hard)). Free-text places fail. Layout-only mentions ≠ anchors. Missing facet → also fail **No silent lore invention**. Path miss ≠ “not in Key Locations” when the index row exists |
| When / timeline | Scene **When** fields make timeline jumps (flashback, rewind, parallel timeline) explicit to a cold reader |
| Staging | Continuing multi-scene situations cite staging; **`stagings/{slug}.md` exists**; blocking matches catalog; cited character states exist |
| World | Beats do not break world-bible laws/history; exposition only for Plant/Hint |
| Length / Prose forecast | Typed **integer** unit formula per scene (no range-only); **written Unit product = arithmetic exactly**; **one** Est. per scene; Est ≈ correct subtotal (±20% / nearest 100); recomputed Est. sum satisfies **Scale min ≤ sum ≤ Scale max**; prefer central band; Estimated Length addends = scene Est. fields; Est. inside outline density band; A/B/C not circular; mid-low unit picks preferred; Dialogue intent matches outline speech; no pilot waiver; **Evidence required (per-scene products)** |
| Series arc | Episode advances the approved series arc; no stealing later payoffs marked Hold / later List rows |
| Tone / voice | Matches `series.md` voice and overview Genre & Tone |
| No silent lore invention | New factions/places/rules → stop and update architecture first; **unciteable Location facets count as place invention** |
| No cross-scene paste | Identical Beat / outline / Unit budget lines across scenes → fail |
| Cast / hook alignment | Every proper name in Summary, Hooks, Seeds, Closing appears in Characters Appearing (or mention-only) and is catalogued; crowd labels in Appearing require profiles |
| No template residue | No raw `{placeholder}` instruction text left in the gate file |

**On conflict:** Prefer correcting the episode design to fit approved sources. If a necessary lore/plot change appears, propose architecture or continuity update → user approves → then finish episode design.

**Episode-file Evidence block (required when claiming Consistency ✅):** under `## Design Consistency Gate`, include short Evidence lines for Locations, Length, Hook, Summary — not checkmarks alone. Example:

```markdown
## Design Consistency Gate
- Loaded required artifacts: ✅ — …
- Locations: ✅ — Evidence: Sc1 `{slug}+{facet}` ⊆ anchors; Sc2 … 
- Length / Prose forecast: ✅ — Evidence: Sc1 written=1,780; recomputed=1,780; Est=1,800 · Sc2 written=…; recomputed=…; Est=… · header sum=…
- Episode List Summary: ✅ — Evidence: 「{Summary clause}」→ Scene {n} Beat …
- Hook to Next / Closing: ✅ — Evidence: Hook「…」; Out「…」; Turn「…」 (same strength — quotes from body, not a softer checklist paraphrase)
```

Bare `- Locations: ✅` / Length Evidence without **written vs recomputed** pairs / Hook Evidence that softens the body → fail.

---

## Staging (situation stage — authored in this stage)

**Staging is episode design, not architecture.** For the three-layer model (Character / Location / Staging) and mandatory staging rule, see SKILL.md § Reference Models. Parallel catalogs vs situation:

| Layer | Owns | Example |
|-------|------|---------|
| Character (③ / additive states) | Wardrobe & equipment **states** (clothes, accessories, weapons) | A `base`, A `{travel-state}`; B `base` |
| Location (③) | Neutral empty set (meeting room, street, vehicle, …) | fixed layout, fixtures, landmarks |
| Staging (④) | This situation — **cites** states + situation props | `a (state: {travel-state})`, `b (state: base)`, seats L/R, shared props, ambient |

Staging **must not** invent outfits/accessories/weapons (see SKILL.md § Reference Models). Need a new look → add character **state** first → cite it. Across scenes sharing one staging, cited states stay fixed.

While designing scenes, for each continuing situation decide:

1. **Reuse** an existing `stagings/{slug}.md` only if the **same continuity span** still applies (same cast **states**, props, ambient, seats).
2. Otherwise **create** a new staging for this episode’s beat (preferred default — same location ≠ same staging).

**What the staging file locks** (layer rules: SKILL.md § Reference Models; procedures here):

- Each named character’s **cited appearance state** (must exist under `characters/{slug}.md`)
- Seat / spot / L-C-R blocking and facing
- Situation props (cups, plates, papers — not identity gear)
- Ambient occupancy when it must not reinvent itself
- Situation environment when it holds for the span

**When mandatory:** ≥2 scenes in the same situation **and** (≥2 named cast **or** fixture-relative placement) → staging required before Key Events. `Staging: none` only for single-scene / no fixed placement.

**Procedure:**

1. List continuing situations while drafting Scene Index / Key Events.
2. Resolve each cast look to a character **state** (add state on profile if missing).
3. Reuse or create `stagings.md` row + `stagings/{slug}.md` binding those states (index + profile together). `{slug}` = situation label kebab-case in the work’s language (Non-Latin kept — see [Canonical catalog & staging slugs](#canonical-catalog--staging-slugs)).
4. Cross-verify cast states and location paths exist on disk.
5. Cite `**Staging:** {slug}` (or `none`) on every scene; On stage must match staging cast when a slug is cited.
6. User approves staging files with the episode design (Gate G4), not at G3.

Do not invent outfits, weapons, drinks, seats, or ambient only inside Key Events — outfits/weapons on **character states**; drinks/seats/ambient on **staging** (see SKILL.md § Reference Models).

---

## New character, location, state, or staging

**Story design co-locks catalogs** — see SKILL.md § Reference Models + table below.

| Need | Already cataloged? | Action |
|------|-------------------|--------|
| Character / appearance state | Yes | Cite in **On stage** / staging |
| Character / appearance state | No | **Stop** → add architecture profile/state → approve → resume |
| Location / lasting set change | Yes | Cite **Location** |
| Location / lasting set change | No | **Stop** → add architecture location → approve → resume |
| Situation stage (who/where/props/ambient) | Yes, **same continuity span** | Cite existing **Staging** |
| Situation stage | No / different beat | **Create** staging in this stage (cast states + location must exist first) → approve with episode → cite |
| World rule | No | **Stop** → additive world bible / aspect — see SKILL.md § Consistency |

1. **Stop** when a gap appears  
2. **Propose** with reason and role (recurring / one-scene / one-situation)  
3. **Wait** for decision  
4. If approved → update catalogs → user approves → resume and **cite**

Do not invent outfit, gear, seats, or L/R only inside Key Events. Expression/mood/one-off weather stay in scene direction.

---

## Reader engagement & literary craft

### Principles

1. Open with tension, not exposition  
2. End with forward momentum (episode closing hook)  
3. Curiosity gaps — plant questions; delay answers  
4. Stakes escalation  
5. Emotional payoff  
6. Episode-to-episode contract — prior hook addressed  
7. Scene-to-scene flow — transitions connect  
8. Personal stake over world lecture  
9. Restraint earns attention — hold material back  

### Design rules (episode + scene)

| Concern | Design rule |
|---------|-------------|
| **Opening Question** | Write it in the episode file — separate from the opening image. Scene 1 inherits or refines it. Not answered immediately. |
| **Info : tension** | Ep 001 ≈ 50:50 (not briefing-heavy). Early episodes lean toward tension; mid-arc+ may shift toward earned revelation. |
| **World through character** | Pair each world detail with POV attitude in Key Events — not a separate exposition block. |
| **Motifs** | Plan 1–2 motifs with scene placements — not discovered during writing. Prefer an optional `- **Motif touch:**` on owning scenes (or Seed touch) so Stage ⑥ has an observable cue. Episode-level motif table alone is allowed, but then treat placement as a **generation constraint** (eval may **Carry-⑥**) rather than a Schema fail. |
| **Dialogue** | Distinct speech pattern per speaker (profile Voice). POV needs said/felt gap where designed. |
| **POV inserts** | ≤ 1–2 per episode; at turning points; list placement + what reader learns that POV does not. |
| **Reader-discovered meaning** | Design “what the reader should conclude” — not “what the narrator will say.” Thematic lines go to Hold; specify a closing **image**. |
| **Antagonist** | Plausible worldview the POV rejects — not moral caricature. |
| **Closing** | Image or action only — no thematic restatement. |

### Exposition budget

| Position | Budget | Guideline |
|----------|--------|-----------|
| Ep 001 | Very low | One world fact + one personal stake + one mystery hint |
| Ep 002–005 | Low | 1–2 new concepts per episode via scene |
| Mid-arc | Medium | Deeper exploration of planted seeds |
| Arc/climax episode | Higher | Payoffs earned by prior restraint |

### Seed discipline

| Class | Meaning |
|-------|---------|
| **Plant** | Show in scene this episode |
| **Hint** | Foreshadow only — reader notices, doesn’t understand yet |
| **Hold** | Do NOT include — reserved for later |

Pick 1–2 Plant, 1–2 Hint per episode; list Hold explicitly. If design lists 5+ “must perceive” elements, manuscript will over-explain.

### Over-seeded opener anti-patterns

| Problem | Fix |
|---------|-----|
| 5+ tension elements in ep 001 | Distribute across episodes |
| Trauma lore in dialogue (ep 001) | Hold; close with a data gap only |
| Political exposition dialogue block | One fact per exchange |
| Date + world essay opening | Specific image/action |
| No personal stake | POV relationship, job, or fear |
| Thematic summary ending | Concrete closing image |
| Key Events with full dialogue | Flow summary + dialogue intent only |

**Rule:** If improved manuscript is mostly *cuts*, the problem was in **design**. If mostly *rewording*, the problem was **literary craft** — update design and regenerate.

---

## Generation Readiness (required before stage ⑥)

Stage ⑥ must **not** invent plot. Confirm:

- [ ] Canonical scene schema only; field notation `**Field:**` / `- **Field:**`
- [ ] Every scene has: POV | Location | When | On stage | Staging (or none) | Situation | Beat | Turn | Function | Sensory-emotional | Dialogue intent (or none) | Transition out | Paragraph outline | Unit budget | Est. length (numeric)
- [ ] Length / Prose forecast: integer `n×pick = subtotal` per scene (not a range); **written product = arithmetic exactly**; **exactly one Est. per scene**; Est ≈ correct subtotal (±20% / nearest 100); Est. inside outline density band; recomputed sum satisfies **Scale min ≤ sum ≤ Scale max** (prefer central band); Estimated Length addends **equal** those scene Est. fields (header ≠ fields → fail); **Gate Evidence** for Length present
- [ ] Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) — **every** scene; every Appearing name has a catalog profile (no uncatalogued crowd labels)
- [ ] **Every speaker** named in Dialogue intent / Beat / outline is on that scene’s On stage (ambient noise ≠ speech — see Cast roster)
- [ ] Summary / Hooks / Seeds / Closing name only Appearing (or mention-only) cast — no orphan names (e.g. Hook Out names a character never On stage)
- [ ] Each scene **When** is timeline-clear (jumps / rewinds / flashbacks stated explicitly)
- [ ] Continuing situations cite staging; **staging profile file exists on disk**; no silent L/R or seat drift (SKILL.md § Reference Models); staging Load is N/A when Staging: none
- [ ] Location cites use Key Locations **slug** / kebab profile paths (not Display-with-spaces filenames); every `+ facet` is a **Multi-facet anchors** label; **Summary-place facet preflight** passed; **Locations Gate Evidence** has separate index / path / facet lines; every scene Location is in Architecture References; **all cited Architecture paths exist**
- [ ] Beats are concrete and dense enough to support each Est. / forecast; Summary signature beat is dramatizable from one scene without inventing
- [ ] No cross-scene or intra-scene paste of Beat / outline / Unit budget
- [ ] No unfilled `{template}` braces left in the file
- [ ] Scene Transition out → next Situation intelligible; last scene Out matches Episode Out / Hook to Next **at the same evidence strength**; no Hook-scope creep; Out not overcrowded (>2 independent obligations); **Hook Gate Evidence** quotes present
- [ ] Hook surfaces agree: Summary / Arc close / Out / Seeds / closing Turn state one claim
- [ ] Overview signature lines placed or Held; profile-backed knowledge claims OK or Conflicts escalated
- [ ] Seeds/Hold and motifs placeable from scene fields (or optional Seed/Motif touch); later-list payoffs not fully executed early
- [ ] Dialogue Voices cover every speaking character in On stage lists
- [ ] A cold reader could draft the episode from this file + architecture refs alone
- [ ] Ep 001: Continuity load not falsely checked; Picks up empty
- [ ] Pre-Design Load used Phase A indexes → Phase B Appearing/used paths only — not a full catalog / `episodes/` / `continuity/` dump

**Fails** if Length budget fails, Unit product ≠ arithmetic, duplicate/orphan Est. lines exist, Est. is outside outline density, any Scene Index row lacks a full scene section, Appearing ≠ On stage union, **Dialogue intent names speakers absent from On stage**, When is ambiguous across a timeline jump, cited paths are missing, **Location facets are not Multi-facet anchors**, Unit budgets are range-only, hook/cast names disagree, **Hook strength drifts across Summary/Out/Turn/Seeds**, **Out invents obligations absent from Summary+Hook**, **Gate Evidence missing for Locations / Length / Hook**, template residue remains, or Load dumped unused profiles. Do not present for Gate G4 while any of those fail.

### Pre-approval checklist

- [ ] Pre-Design Load completed; Prior Design Alignment filled
- [ ] **Summary-place facet preflight** passed (or Conflicts escalated)
- [ ] Design Consistency Gate all pass — **Gate Evidence** present for Locations / Length / Hook / Summary
- [ ] Unit products recomputed exactly; Generation Readiness passed
- [ ] Episode expanded from `series.md` row; no scene counts in `series.md` / `overview.md`
- [ ] Scene Index + one full scene section per Index row; no nested scene files
- [ ] Architecture + Continuity References complete (ep 002+)
- [ ] Engagement / craft rules applied (Opening Question, seeds, motifs, closing image, etc.)

---

## Design evaluation (recommended)

**Only after Gate G4** (explicit user approval of the episode design). Asking to “evaluate the design” is **not** G4 approval — do not mark the design gate done in `TODO.md` merely because evaluation was requested.

After the user approves an episode design, propose:

> *"Episode {nnn} design is approved. Shall we evaluate the design before writing the manuscript? Design evaluation catches structural issues more efficiently than post-prose revision — but it's optional. We can skip straight to the manuscript if you prefer."*

If the user asks for design evaluation **before** approving ④: run ⑤ against the current draft (Load + write evaluation), keep design gate **open**, and treat Schema ❌ / Pending adjudication as return-to-④ — still **not** G4 done.

If accepted → `05-evaluate-design.md`. If declined → `06-generate.md`.  
Results go to `evaluations/{nnn}-{episode-slug}-design.md`.

---

## Gate G4 (Episode Design Approval)

Requires **explicit** user approval (e.g. “approved”, “go ahead”, “design OK”; multilingual e.g. “승인”, “진행해”). Soft phrases like “evaluate the design” / “평가해줘” alone are **not** G4.

User may approve only if Design Consistency Gate and Generation Readiness both pass. Do not proceed to `06-generate.md` without approval. Do not present a header-only file as the gate artifact. Do not flip the design-approval task to done in `TODO.md` without that explicit approval.

Design evaluation (`05-evaluate-design.md`) is strongly recommended between design approval and manuscript generation. The user may decline and proceed to generation — or request evaluation on a **draft** before G4 (then keep G4 open).
