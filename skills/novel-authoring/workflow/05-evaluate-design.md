# Stage ⑤ — Evaluate & Revise — Design (`evaluations/{nnn}-{episode-slug}-design.md`)

**Prerequisites:** Stage ④ design file on disk. **Preferred:** Gate G4 (explicit design approval) first. If the user requests evaluation on a draft before approving ④, evaluate anyway, keep G4 open, and do **not** mark design approval done.

**Gate artifact:** `{project-root}/evaluations/{nnn}-{episode-slug}-design.md`

**Next stage:** `06-generate.md` (after evaluation approval) — or back to `04-episode-design.md` for revisions

**Status:** **Recommended, not mandatory.** Catches structural, pacing, seed, engagement, and generation-readiness issues **before** prose. If the user declines, proceed to `06-generate.md` after ④ approval. **Manuscript evaluation (Stage ⑦) is the main quality gate** and runs for every episode.

**See also:** SKILL.md § Revision · § Consistency · [`04-episode-design.md`](04-episode-design.md).

---

## Procedure

Propose after episode design approval (or when the user asks to evaluate a draft — then say G4 is still open):

> *"Episode {nnn} design is approved. Shall we evaluate the design before writing the manuscript? We'll validate structure, seeds, and engagement against criteria, then run the required critic personas. It's optional — we can go straight to the manuscript if you prefer."*

Do **not** treat “evaluate the design” as design approval. Do not mark Gate G4 / design-approval TODO items done solely because ⑤ ran.

### Load

Follow SKILL.md **Selective artifact load** (Phase A indexes → Phase B needed details only).

- Resolve the **actual** design path from `episodes/` — do not assume the canonical name if the disk file is an alias
- That file in full (episode + scenes + Prior Design Alignment) — load **once if not already in this turn’s context**; do not re-read after writing the evaluation
- **Phase A:** `overview.md`, `series.md`, `characters.md`, `locations.md`, `world-bible.md`, `stagings.md` (if missing from context)
- **Phase B — same prior-design set as Stage ④:** character/location **profiles for Appearing/used only**, touched `world/` aspects, **each cited** `stagings/{slug}.md`, continuity ep 002+ (`story-so-far` + **immediate prior** summary only) — **not** the design file alone, and **not** unused mid/late-arc profiles or every continuity summary. Batch only paths still missing from context.
- **Path existence (hard — every eval):** Before Schema path rows, load **every** Architecture References character/location/staging profile path for this episode **this turn** (again after compaction if those bodies are gone). This is the existence check — **not** a forbidden “verify write” re-read (see `SKILL.md` Exception — eval path / existence). Index-only load is not enough for path ✅.
- **Do not** re-dump unused mid/late-arc profiles during eval. Build the Appearing/used list from the design file first, then load only those missing paths (+ Phase A indexes) — **except** cited Architecture Reference paths, which always need a this-turn body for path/facet checks. **Do not** batch-read all of `episodes/` or `continuity/`.
- Lore/cascade: SKILL.md § Consistency
- Re-run **Design Consistency Gate** from `04-episode-design.md` (including **Hook alignment** / Hook scope) against the loaded design text
- Lock **Target Reader** from `overview.md` before persona critique
- **Literary Awards Juror** (when in set): also load `series.md` Episode List + `overview.md`; fill **Literary Awards Juror Checks (Design)** before persona critique

**Wrong filename:** if the design is not `episodes/{nnn}-{episode-slug}.md`, Schema Path ❌. **Evaluated artifacts** must cite the **real** path. Return to ④ to rename — do not treat an alias as the gate file.

**Missing cited path:** follow [Path existence discipline](#path-existence-discipline-hard) below. Do not mark Staging / path ✅ from the index alone. Do not mark path ❌ from memory, Display-with-spaces filenames, or a failed wrong path.

**This stage does not edit the episode design.** Missing Key Events, Length mismatches, etc. → Schema ❌ and **return to ④**. Do not edit `episodes/` while evaluating.

**Chat-only eval is not done.** After loading this workflow, tools must include ensuring the **actual** design is in context (read once if missing) and writing `evaluations/{nnn}-{episode-slug}-design.md`. Dumping Schema/Verdict in chat only is incomplete. If the user says the file is missing, write it **in that turn**. Do not re-read the evaluation you just wrote to “confirm.”

### Path existence discipline (hard)

Path ❌ means: **this turn’s load of the Architecture References path failed** — not “profile feels missing,” not “Key Location Display has spaces,” not “anchors were unread.”

| Step | Rule |
|------|------|
| 1. Cite exact path | Use the path string from Architecture References / Load (e.g. `locations/옛-항구.md`), never paste Key Locations **Display** (`옛 항구`) as a filename. |
| 2. Resolve Display → slug | If the design only shows a Display name, map via `locations.md` **Key Locations** table (`Display` → `slug` → `Profile path`). No table → still try kebab (`spaces` → `-`) once before ❌. |
| 3. Forced read | load that Profile path this turn. Success → path ✅ Evidence quotes the path + “read OK”. |
| 4. One recovery | On failure: (a) retry kebab form if Display had spaces, and/or (b) list the catalog dir (`locations/`, `characters/`, `stagings/`) and match the slug. Still missing → path ❌ with Evidence quoting the path attempted + tool failure. |
| 5. No tool, no ❌ | **Forbidden:** mark Cited paths ❌ without a failed load (or list) result **in this turn**. **Forbidden:** mark ✅ without quoting the exact path read. |

**Three checks — never merge into one Evidence cell:**

| Check id | Passes when | Does **not** mean |
|----------|-------------|-------------------|
| **Locations ⊆ Key Locations** | Scene location Display/slug maps to a Key Locations row | Profile file exists; facets OK |
| **Cited staging/profile paths exist** | This-turn load of Architecture References paths succeeded | Facets ⊆ anchors; index membership |
| **Location facets ⊆ Multi-facet anchors** | After a successful profile read, every scene `+ facet` is an exact Multi-facet anchors label | Path existence (already separate) |

A missing profile file → **Cited paths** ❌ (and facet row N/A or ❌ because anchors unread). Do **not** rewrite that as “Key Location missing from index” when the index row is present.

### Criteria Check discipline

Copy **every** overview Validation Criteria bullet. For each:

| Result | When |
|--------|------|
| ✅ | This episode’s Beats **execute** the criterion (not merely plant a seed for later) |
| ❌ | This episode’s **design-scope** includes the criterion and the design fails it |
| **N/A** | Criterion is **out of design-eval scope** (see map below) or series-/late-arc-level — Evidence must say *why* |

**Forbidden:** marking series-payoff criteria ✅ because “seeds were planted” or “direction looks right.” Planting ≠ satisfying a recovery/ending criterion — use **N/A**.

#### Criteria applicability map (design eval)

| Criterion shape (overview wording) | At Stage ⑤ design eval |
|------------------------------------|------------------------|
| “In this episode / ep 001, **present / place / execute** …” (inciting event, first confrontation, hook beat) | ✅ / ❌ — Beats must execute |
| “In P1, **maintain** …” / tone / antagonist flatness as **this episode’s** design choice | ✅ / ❌ if this episode’s Beats touch it |
| “**Pay off** later in the arc” / “hold until the ending” / endgame climax | **N/A** — Evidence: late-arc / series payoff |
| “Each episode **manuscript** hits target length / event density…” / published prose quality | **N/A** — Evidence: manuscript Stage ⑥/⑦. Forecast arithmetic belongs in **Schema** (Forecast ↔ Est.), **not** as Criteria ❌ |
| Platform Scale band for **design Est. sum** | Schema Length rows only — do not double-count as Criteria ❌ |

Wrong: marking Criteria ❌ for a Unit-budget multiply error. Right: Criteria N/A (manuscript quality) + Schema Forecast ❌.

### Schema gate (before craft/persona)

Run Schema / Structural Integrity on every scene. Any ❌ blocks design-eval pass and stage ⑥. **Recompute from the design body in context** (Load once if missing), not from a remembered header alone. Ensure `characters.md` (and needed profiles) are in context so Appearing ⊆ catalog is checked — do not mark Characters ✅ from the design text alone. Do not re-read after the eval write.

1. **n×band per scene** — multiply yourself (e.g. `4×200 = 800`). Do not trust a wrong product written in the design. If the Unit budget is only a **range** (`1,650–2,150`) with no integer `n×pick = subtotal` → Forecast ❌. If the **written** `= {subtotal}` ≠ your recomputed product → Forecast ❌ (**exact product**; ±20% does not apply to the formula line). If Est. ≠ **correct** subtotal beyond ~20% (nearest-100 round OK) → Forecast ❌.
2. **Outline density** — band = outline line count × 200–350. Est. outside that band → Forecast ❌ even when Unit subtotal equals Est.
3. **Single Est. per scene** — two `- **Est. length:**` lines, or Est. leftover from a deleted scene → Length ❌.
4. **Two-line episode Evidence is mandatory** in the Length rows (right-hand total = arithmetic of the left; never paste the header into that slot):

```text
scene Est. fields: {a} + {b} = {sum_fields}
header addends:    {c} + {d} = {sum_header}
```

If fields ≠ header → Length ❌ even when one side is ≥ Scale min. If recomputed sum < Scale min → Scale-min ❌. If recomputed sum > Scale max → Scale-max ❌. Prefer central band (default Scale 4,000–8,000 → ~5,000–7,200) — sum at Scale max with no headroom → note ⚠️ in Evidence (still ✅ on max row if ≤ max). Do not copy “cross-check ✅” from the design. Result cell must match Evidence (do not write ❌ when Evidence shows a pass, or ✅ when Evidence shows a fail).

5. **Hook alignment — judge body surfaces first (mandatory quotes — no quotes = ❌):**

   **Canonical surfaces (Schema Hook):** `series.md` Hook to Next; Episode Summary; Episode Out; Episode Arc close; Seeds How it appears (Hook element); closing Scene Turn / Dialogue intent / Transition out.

   - Evidence **must** quote those body surfaces (not only the design’s self-checklist Gate Evidence line)
   - Paraphrase-only Evidence **without quotes** → Hook evidence strength ❌ **and** Hook internal consistency ❌
   - **Same evidence strength** across **body** surfaces → Schema Hook ✅; soften/harden **in the body** → Episode List plot ❌ **and** Hook internal consistency ❌
   - If Summary/Out already match Hook but only Scene Turn softens → still Schema ❌ (return to ④ to align Turn/Seeds)
   - **Gate-meta hygiene (do not inflate to Schema Hook ❌):** if **body** surfaces already match Hook at one strength, but the design file’s Consistency Gate Evidence line alone uses a weaker/stronger paraphrase → Schema Hook stays ✅ (body wins); open a **Med** Adjudication finding “align Gate Evidence quotes with body” → Apply? yes → **Pending** (④ one-line Evidence fix) **or**, if body is already generation-ready and only checklist wording is wrong, still Pending until Evidence matches — never mark Hook Schema ❌ solely for checklist meta when body is aligned
   - Do **not** treat “Summary OK so plot passes” when a **body** surface softens

6. **Hook scope:** last Transition/Out invents chase/faction arrival/second reveal absent from Summary **and** Hook → Episode List scope ❌ (even if Out “feels exciting”)
7. **Locations — three separate checks** (see [Path existence discipline](#path-existence-discipline-hard)):
   - **Index:** scene place ∈ Key Locations (Display or slug via table) — do not fail this row for a missing file
   - **Path:** Architecture References profiles/stagings load'd this turn — Evidence quotes exact paths
   - **Facets:** only after path read OK — every `+ facet` ⊆ that profile’s **Multi-facet anchors** (layout/sensory-only ≠ anchors). Facet miss → Location facets ❌ **and** No improvised entities / silent lore ❌ (same finding — do not mark improvised ✅ while facets ❌)
8. **Length Evidence rubber-stamp:** if the design’s Length Gate Evidence lists products, **recompute independently**. Written Unit `=` ≠ recompute → Forecast ❌ even when Evidence “looks complete.” Prefer Evidence that already shows `written=` / `recomputed=` pairs; absence of pairs is not itself Schema ❌ if products on Unit lines are correct, but wrong pairs or wrong Unit lines are ❌
9. **Profile-backed knowledge:** character “knows X” with no Relationships/Drive support and Conflicts empty → Architecture Compliance ❌ or Character Critic finding (do not leave silent)
10. **overview signature lines:** Constraints core dialogue unused and not Held → Literary Craft / Engagement note (Med if the line is a series promise)
11. **Eval table hygiene:** do **not** duplicate the same Check row twice in Schema / Consistency / Design Consistency Gate tables. One row per check id.

**Also ❌ (no “pilot” ⚠️):** missing typed units; wrong Unit product; circular outline×300; Dialogue intent vs speech mismatch; Appearing ≠ On stage; **Dialogue intent / Beat names speakers (including anonymous guests / off-stage voices) absent from that scene’s On stage**; ambient crowd given speech intent while On stage is POV-only; **When** ambiguous across a timeline jump (flashback / rewind / parallel timeline); Location/cast not in architecture; **unciteable Location facets**; **uncatalogued crowd labels in Appearing**; **Summary/Hook names a character never Appearing/On stage**; missing staging profile file; staging Load checked while Staging: none; later-list cast debut; Episode List clause missing from Beats; **Hook strength drift across body Summary/Out/Turn/Seeds**; **Hook Evidence without body quotes**; **Out obligations beyond Summary+Hook**; design-paste / repeated outline loops; **raw `{template}` braces** left in the file; ep 001 Continuity falsely checked.

**Generation Readiness row (Structure & Arc):** Result must be **❌** (not ⬜, not ✅) if any Schema / Structural Integrity Length, forecast, cast, path, facet, or Hook-**body** row is ❌ — same bar as “would Gate G4 fail.” Use ⬜ only in the **Verdict** `Generation-ready` cell.

**Generation-ready (Verdict):** ✅ only when (1) **all** Schema / Structural Integrity rows are ✅, **and** (2) no Adjudication row has Apply? = yes with Status **Pending**. Status **Carry-⑥** does **not** block Generation-ready. Otherwise **⬜**. Chat must not say “fix then it becomes ready” while the Verdict already shows ✅, and must not offer skip-to-⑥ while Generation-ready is ⬜.

**Required personas:** run every persona in the [Default persona sets](#default-persona-sets). Do not skip for speed. Optional personas only when overview criteria or the user request them.

**Persona Defects floor (anti-empty-critique):** Plot Expert **must** check Hook alignment + Hook scope. Genre Critic or Reader-Editor **must** check Out density / crowded closing beat when the last Transition stacks ≥3 signals. Character Critic (when in set) **must** check profile-backed knowledge for identity/recognition hooks. Writing `—` for Defects is allowed only after those checks are actually negative.

---

## Evaluation record structure

Write `evaluations/{nnn}-{episode-slug}-design.md` (separate from manuscript eval at Stage ⑦).

```markdown
# Design Evaluation: Episode {nnn} — {Title}

## Criteria Check (from overview.md)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| {copy each bullet from overview **Validation Criteria** — do not invent genre-label rows} | ✅ / ❌ / N/A | {Use [Criteria applicability map](#criteria-check-discipline). **N/A** for late-arc payoffs **and** for “manuscript length/density” — put forecast arithmetic in Schema. **Never ✅ for “seeds planted toward” a late-arc payoff**} |

## Schema / Structural Integrity (any ❌ blocks design-eval pass / stage ⑥)
| Check | Result | Evidence |
|-------|--------|----------|
| Canonical scene schema only (no nested subsections / alternate layouts) | ✅ / ⚠️ / ❌ | |
| No skill/workflow dump after the design | ✅ / ❌ | {**Forbidden:** sections copied from 04 procedure such as `## Pre-Design Load`, `## Prose forecast` how-to, or pasted workflow checklists as body. **Allowed:** canonical gate status block `## Gate G4` (and short Status/Next lines) required by stage ④ — do **not** ❌ solely because `## Gate G4` exists} |
| Unique `### Scene n` headings; no pasted twin scenes | ✅ / ❌ | {duplicate Scene n, or Scene 1 Beat copied as Scene 2 → ❌} |
| Canonical episode path (`episodes/{nnn}-{slug}.md`) | ✅ / ❌ | {actual disk path; alias → ❌} |
| Field notation `**Field:**` / `- **Field:**` (colon inside bold) | ✅ / ⚠️ / ❌ | |
| Every scene has required meta + bullet fields (`none` OK) | ✅ / ⚠️ / ❌ | {includes Paragraph outline + Unit budget + **exactly one** Est.} |
| Characters Appearing ↔ On stage union (no ghosts; mention-only tagged) | ✅ / ❌ | {union of **every** scene On stage — not “all Appearing appear in every scene”} |
| On stage includes speakers | ✅ / ❌ | {**every** speaker in Dialogue intent / Beat / outline is on that scene’s On stage. Anonymous guests / off-stage voices that “speak” count. Ambient murmur/footsteps with no speech intent → OK without On stage. POV-only On stage + crowd speech intent → ❌} |
| Characters ⊆ `characters.md` | ✅ / ❌ | {verify on disk; crowd labels / improvised cast without profile → ❌} |
| Summary/Hooks cast alignment | ✅ / ❌ | {names in Summary / In / Out / Seeds / Closing must be Appearing or mention-only} |
| No later-list cast debut | ✅ / ❌ | {cast introduced in a later Episode List row but on stage now → ❌ unless user approved list/architecture change} |
| Locations ⊆ `locations.md` Key Locations | ✅ / ❌ | {**Index only:** scene Display/slug maps to a Key Locations row (prefer table slug). Free-text place → ❌. **Do not** fail this row because a profile load failed — that is Cited paths} |
| Location facets ⊆ Multi-facet anchors | ✅ / ❌ / N/A | {**After** successful profile read: every scene `+ facet` is an exact Multi-facet anchors label. **Evidence must quote** the profile’s Multi-facet anchors list (or the exact matching label). Layout / Internal facets / Sensory-only ≠ anchors. N/A if no `+ facet`. Path unread/missing → ❌ or N/A with Evidence “anchors unread” — do not invent anchors from index prose. **Two adjacent scenes sharing the same citeable facet** → Schema may stay ✅ only if Functions/Turns are visibly distinct; still raise a **Pending** (not Carry-⑥) design finding unless the design already separates them} |
| Nested `episodes/{slug}/` scene files absent | ✅ / ⚠️ / ❌ | |
| No template residue | ✅ / ❌ | {raw `{placeholder}` / `{TOC only…}` left in file → ❌} |
| Prose forecast present (outline + **typed** units) | ✅ / ❌ | {five-type n’s; **integer** `n×pick = subtotal`; reject ranges-only, untyped padding, non-allowed type labels} |
| Forecast ↔ Est. cross-check (independent) | ✅ / ❌ | {**recompute** n×band yourself; ignore design’s self-✅. **written product must equal arithmetic exactly**; Est. vs **correct** subtotal ±20% / nearest 100; Est. inside outline×200–350 band; circular outline×300 = Est. → ❌; range-only Unit budget → ❌; duplicate Est. lines → ❌. If design Length Evidence already shows written≠recomputed, still ❌} |
| Dialogue intent vs outline speech | ✅ / ❌ | {intent `none` but outline/Beat has speech → ❌} |
| Recorded Estimated Length = scene Est. sum | ✅ / ❌ | {**mandatory two-line Evidence** — scene fields vs header addends; fields ≠ header → ❌} |
| Est. length sum ≥ Scale min (hard) | ✅ / ❌ | {use sum from **scene fields**, not the header; Result must match Evidence arithmetic} |
| Est. length sum ≤ Scale max (hard) | ✅ / ❌ | {same sum; > Scale max → ❌. Default Scale 4,000–8,000. Central-band note optional in Evidence} |
| Cited staging/profile paths exist | ✅ / ❌ | {**Path Evidence mandatory:** quote each Architecture References path + this-turn load OK. No verified load this turn → cannot ✅. Path ❌ only after failed exact-path read **and** one kebab/list recovery ([Path existence discipline](#path-existence-discipline-hard)). Staging: none → N/A / ✅ without claiming a staging file exists. Display-with-spaces filename failure alone ≠ path ❌} |
| Episode List plot (not a different story) | ✅ / ❌ | {Evidence must **quote `series.md` Summary**; every clause maps to a Beat; speakers match who the Summary assigns} |
| Hook evidence strength (internal) | ✅ / ❌ | {**Quotes from body surfaces** — Hook + Summary + Out + closing Turn + Seeds. Soften/harden **in body** → ❌. Checklist Gate Evidence alone softer/harder while body matches Hook → Schema ✅ + Med Pending hygiene finding — **not** Schema ❌} |
| Hook scope (no Out creep) | ✅ / ❌ | {Last Transition/Out must not invent chase/faction/second reveal absent from Summary **and** Hook; >2 independent Out obligations → ❌ or Med finding} |
| No design-paste / meta-only scenes | ✅ / ❌ | {identical Beat or Unit budget lines across scenes; repeated outline loop in one scene; scene with no dramatic event → ❌} |

## Structure & Arc Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Episode arc coherent | ✅ / ⚠️ / ❌ | |
| Scene transitions chain | | {Transition out → next Situation; **When** clear across timeline jumps} |
| Scene sections complete | | {Every Scene Index row has full Key Events?} |
| Generation Readiness | ✅ / ❌ | {④ checklist. **❌ if any Schema Length/forecast/cast/path/facet/Hook-body ❌** — never ⬜ in this row; never ✅ while G4 would fail} |
| Beat concreteness | | {No mood-only Beats} |
| Est. length budget | | {Recomputed sum = Estimated Length header = scene Est. fields; Scale min ≤ sum ≤ Scale max; prefer central band} |
| Prose forecast quality | | {Typed units match Dialogue intent / Sensory / Beat} |
| Episode List scope aligned | | {Summary + Hook to Next honored — not a rewritten plot; **no Out creep** beyond Summary+Hook} |
| Prior hook addressed (ep 002+) | | |

## Architecture & Continuity Compliance
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-Design load reflected in Prior Design Alignment | ✅ / ❌ | |
| Series / overview tone & arc honored | | |
| Episode List Summary / Hook to Next honored | | {quote both; check Turn strength, not only Summary} |
| Hook internal consistency (design surfaces) | | {**Body** Summary / Arc / Out / Seeds / Turn agree. Checklist-only drift → Med hygiene, not automatic Compliance ❌} |
| Characters from architecture; profiles not redefined | ✅ / ❌ | {indexes-only load → cannot mark ✅} |
| Profile-backed knowledge / recognition | ✅ / ❌ / N/A | {identity or “knows X” hooks need Relationships/Drive — or Conflicts escalated; N/A if no such claim} |
| Locations from architecture; profiles not redefined | ✅ / ❌ | {**Index:** scene Location ∈ Key Locations table. Do not put path-missing here} |
| Location profile paths readable | ✅ / ❌ | {Same bar as Schema Cited paths — this-turn read Evidence} |
| Location facets ⊆ Multi-facet anchors | ✅ / ❌ / N/A | {Separate from path; anchors only after successful profile read} |
| Stagings from episode design; blocking not redefined | | {**profile file exists** when Staging ≠ none — this-turn read; index-only → ❌; authored in ④ not ③} |
| World rules / history consistent with bible | | |
| No improvised entities or silent lore | | {unciteable Location **facet** = place invention → ❌ **with** Location facets ❌ — do not mark this ✅ while facets fail; do not use this row for Display/slug path typos} |
| Continuity files used (ep 002+) | | |
| Character/location state vs `story-so-far` | | |
| Unresolved threads: pick up / advance / plant / hold | | |
| No contradiction of released continuity | | |
| Conflicts section empty or escalated (not ignored) | | |

## Design Consistency Gate
{Copy pass/fail from `04-episode-design.md` checks **once** — do not paste duplicate Locations/Length/Hook rows. Any Schema ❌ → Generation-ready ⬜. Also ❌ if Locations index/path/facet or Length lack Gate Evidence, or Length Evidence omits written/recomputed pairs when products are wrong. Hook Gate Evidence softer than body → Med hygiene finding if body already matches series Hook. Path Evidence must quote this-turn load paths — see Path existence discipline.}

## Engagement Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening Question defined | ✅ / ⚠️ / ❌ | |
| Personal stake present | | |
| Episode Out hook | | {matches Hook strength; not overcrowded; no scope creep sold as “extra excitement”} |
| Exposition budget respected | | |
| Seed discipline | | {Hint vs Plant wording must not contradict Summary/Out Hook claim} |
| Scene-first Key Events (all required fields) | | |
| Sensory-emotional on every scene | | |
| Motifs planned across scenes | | |
| Overview signature line | ✅ / ❌ / N/A | {Constraints dialogue in Dialogue intent or Held; N/A if overview has none} |

## Literary Craft Checks (Design)
| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | | {Ep 001: ~50:50?} |
| Sensory-emotional pairing | | |
| Dialogue voices + Dialogue intent | | |
| Reader-discovered meaning | | {Theme in Hold, not closing monologue?} |
| Antagonist plausibility | | |
| Closing image specified | | |

## Literary Awards Juror Checks (Design)
{Include **only** when Literary Awards Juror is in the active set (prestige/awards in overview, or user request). Otherwise write exactly: `{Not required — overview.md has no prestige/awards criterion.}` and **do not** fill the table with ✅ from design intent.}

| Check | Result | Evidence |
|-------|--------|----------|
| Theme & vision earned (not asserted) | ✅ / ⚠️ / ❌ | |
| Original perspective or form planned | | |
| Human insight beyond plot mechanics | | |
| Artistic integrity — no planned didacticism or sentimentality | | |
| Episode advances series literary arc | | |
| Competition readiness at design scope | | {What jurors would praise / hesitate over / reject} |

## Target Reader Checks (Design)
{Always — use audience locked in overview.md}

| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns the locked reader's attention | ✅ / ⚠️ / ❌ | |
| Personal stake matches what this reader came for | | |
| Pacing / density fits platform expectations | | |
| Out hook makes *this* reader want the next episode | | |
| No alienation of core audience without overview intent | | |

## Design Critique (required personas)
For **each** persona in the active set:

- `#### {Persona name}`
- Stance: {1–2 lines}
- Strengths: {…}
- Defects: {finding → severity High/Med/Low → proposed fix — or a single `—` if none}
- Reader impact: {how this affects the locked Target Reader — or `—`}

Do **not** duplicate empty Defect / Reader impact / Severity / Proposed fix blocks. One clean block per persona.

## Design Adjudication
{Record Apply decisions. Default tie-break: Target Reader. Never silently drop a High finding. If there are no High/Med findings, leave the table empty or one row `—` — do not invent filler rows. Schema ❌ (Length, plot, cast, path, forecast, facet) **are** findings — they must appear here. **Deduplicate:** same root cause from multiple personas → one adjudication row (cite personas).}

**⑤ does not apply revisions.** Status values:

| Status | Meaning |
|--------|---------|
| **Pending** | Apply? = yes — **Stage ④ (or ③)** must edit design/architecture (closing image field, crowded Out, adjacent same-facet separation, observable character-cost beat, body Hook, Unit-product); blocks Generation-ready |
| **Skip** | Apply? = no — rejected / not applying |
| **Carry-⑥** | Apply? = yes — **generation constraint only** (Hold, POV limits, “do not narrate X”, **episode-level motif placement without scene Motif-touch fields**). Design file need not change for handoff; does **not** block Generation-ready. **Do not** use Carry-⑥ for crowded closing, scene-function separation, or personal-cost beats that belong in Key Events |
| **Applied-④** | Former Pending/Carry already absorbed into the revised design body this loop — Evidence cites updated fields; does **not** block Generation-ready |

Writing “applied” / “all fixed” while `episodes/` is unchanged — or editing `episodes/` during ⑤ — is forbidden.

**Adjudication Status:** `Apply? = yes` + design/architecture / Unit-product / **body** Hook / closing / adjacent-facet / cost-beat fix → Status `Pending`. `Apply? = no` → **`Skip`**. Generation-time Holds / POV / motif-touch reminders / checklist-Evidence-only hygiene after body already matches Hook → prefer **`Carry-⑥`** or a **narrow Pending** one-line Evidence fix — **do not** escalate checklist-only Hook wording into Schema Hook ❌. Do not leave rejected findings as Pending. Do not put pure generation constraints as Pending.

**User “권고 반영” / apply recommendations:** Prefer Kind **B** (④) for Med+ design-field findings. After the design body changes, **re-run this entire Schema + Adjudication** on the revised file; set Status **Applied-④**; sync design `## Gate G4` with the user’s re-approval in the same turn. **Forbidden:** flip Gate G5 to Approved without rewriting the eval against the revised design.

**Generation-ready stays ⬜ while:** any Schema ❌, **or** any Apply? = yes is still **Pending**. **Carry-⑥** and **Applied-④** do **not** keep Generation-ready ⬜. Do not mark Generation-ready ✅ and then tell the user “fix these and it becomes ready.” Align chat with the Verdict cell.

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | {finding} | High/Med/Low | No | yes/no | {why} | {fix at ④ / note for ⑥ / —} | Pending \| Skip \| Carry-⑥ \| Applied-④ |

## Design Verdict
| Dimension | Result |
|-----------|--------|
| Prior-design consistent | ✅ / ❌ |
| Target-reader readiness | ✅ / ❌ |
| Design quality | ✅ / ❌ |
| Generation-ready | ✅ / ⬜ |
```

**Verdict dimension rules (outside the eval file template):**

| Dimension | How to set |
|-----------|------------|
| Prior-design consistent | ❌ if Architecture/Continuity/Locations/facet/**body** Hook-alignment Schema or Compliance has ❌ (checklist-only Hook wording ≠ automatic ❌ here) |
| Target-reader readiness | ❌ if Target Reader Checks have ❌, or High Apply **Pending** that harms retention; Schema-only arithmetic need not force ❌ here if Target Reader rows are fine |
| Design quality | From craft personas. Do **not** set ❌ solely because Forecast math / facet Schema failed when craft Defects are empty/`—` |
| Generation-ready | ✅ only if Schema all ✅ **and** no **Pending** Apply. **Carry-⑥** does not block. Else ⬜ |

**Verdict cells:** only the characters in the template. No ⚠️, no parentheticals in the cell. `Generation-ready` is **✅ or ⬜** (not ❌).

**Generation-ready ✅ requires all of:**
1. Schema / Structural Integrity all ✅ (including Length two-line match, integer Unit budgets with **exact products**, citeable facets, cited paths exist, Hook Evidence **with quotes**, no design-paste)
2. No Adjudication Apply? = yes still **Pending** (Carry-⑥ OK)

Otherwise Generation-ready **⬜**. Do not tell the user they may proceed to ⑥ while this is ⬜. Offering “skip fixes and go to Stage ⑥” while Generation-ready is ⬜ is forbidden.

---

## Design-first revision loop

If evaluation finds issues:

1. Return to **Stage ④** and/or **Stage ③** as needed — update episode design and/or architecture — following SKILL.md § Revision (prefer additive architecture). Med+ design-field findings (closing, adjacent same-facet, cost beat) → Kind **B** / Status **Pending**, not Carry-⑥ alone.
2. Re-run Schema / Structural Integrity, Design Consistency Gate, Generation Readiness; update the evaluation file — mark absorbed rows **Applied-④**
3. Present revised design for approval; sync Gate G4 + G5 + TODO; then proceed to `06-generate.md`

**Verdict rule:** Schema / Structural Integrity must be all ✅ **and** no Apply? **Pending** before **Generation-ready ✅** or stage ⑥. Schema ❌ or Pending Apply → Generation-ready ⬜; return to Stage ④ for Pending items. **Carry-⑥** alone does not require a ④ loop — hand the constraint to Stage ⑥.

### Common critique → design fixes

| Critique finding | Design fix |
|------------------|------------|
| Too much world-building in opening | Move lore to Hold; lower exposition budget |
| Info-dump dialogue | Key Events: intent only; cut lecture beats |
| Too many mysteries explained | Reclassify as Hold; one Hint |
| No personal connection | Add Personal Stake |
| Thematic summary ending | Closing image in Hold/craft |
| Opening action without tension | Add Opening Question |
| Episode vs Episode List Summary | Cascade C: `series.md` → design |
| Conflicts with world/continuity | Cascade D/E |

---

## Persona Reference

Critics **advise**. They do not auto-apply changes. When personas conflict, **Target Reader** (from `overview.md`) is the default tie-break unless overview explicitly prioritizes prestige, niche, or another locked criterion.

### Default persona sets

| Context | **Required** (always run) | **Also required when** |
|---------|---------------------------|-------------------------|
| Design eval (any episode) | **Target Reader**, **Genre Critic**, **Plot Expert**, **Reader-Editor**, **Literary Critic** | Character-driven episode → **Character Critic**; ep 001 or major lore debut → **Character Critic** + **Setting/Lore Expert**; prestige/awards in overview → **Literary Awards Juror** |
| Ep 001 (design) | Above + **Character Critic** + **Setting/Lore Expert** | — |

Optional on request: any persona below; series/arc-level critique → `evaluations/{scope}.md`.

### Persona catalog

| Persona | Focus | Typical questions |
|---------|-------|-------------------|
| **Target Reader** | Locked audience’s lived read — fun, clarity, desire to continue, platform fit | Would *I* keep going? Where do I skim, confuse, or quit? |
| **Reader-Editor** | Engagement craft, exposition restraint, serialization hooks, **Out density** | Is info dumped? Is the hook earned? Is the episode sellable as a unit? **Is the closing Transition overcrowded (≥3 independent signals)?** |
| **Genre Critic** | Genre contracts, tropes, promise/payoff | Does it deliver what this genre’s readers paid for — without empty cliché? |
| **Plot Expert** | Pacing, causality, structure, seed timing, **Hook body alignment + Hook scope** | Does every turn follow? **Do body Summary/Out/Turn agree on Hook strength?** Checklist-only soften ≠ Schema fail if body matches. **Does Out invent obligations beyond Summary+Hook?** |
| **Character Critic** | Arc, motivation, distinct voice, relationship pressure, **profile-backed knowledge** | Why does this person act *now*? Voices interchangeable? **Can they know/recognize X per profile?** |
| **Literary Critic** | Motif, sensory-emotional craft, Hold vs closing image, voice as designed | Will the planned craft survive generation — or is it only asserted in Hold? **Episode-level motifs without scene Motif-touch → Carry-⑥, not Schema ❌** |
| **Setting/Lore Expert** | World consistency, info-dumping, rule clarity, **citeable Location facets** | Is the world learned through character need — not catalog? **Is every scene `+ facet` under Multi-facet anchors?** |
| **Literary Awards Juror** | Theme, originality, insight, artistic integrity, work-level merit | Would this earn lasting literary respect — or only serial entertainment? |

**Literary Awards Juror** should cover: theme & vision earned; originality; human insight; artistic integrity; work-level coherence; competition readiness with concrete revision priorities.

**Target Reader** should cover: who they are (from overview); what they came for; first-page retention; confusion or lecture risks; whether the Out hook pulls *them* specifically.

---

## Gate G5 (Design Eval Approval)

After evaluation + adjudication (and any applied revisions): user reviews the evaluation, confirms revision items, and decides revise (→ ④ or ③) or proceed to `06-generate.md`.

Schema / Structural Integrity must be all ✅ **and** no Apply? **Pending** before **Generation-ready ✅** or stage ⑥. Carry-⑥ may remain.

If the user declines design evaluation entirely, skip this stage — proceed from **④** approval directly to `06-generate.md`.

## Reference-model integrity (design eval)

Also check SKILL.md § Reference Models:
- Character appearance/equipment states consistent; no silent gear swap
- Location used as set; lasting changes tracked; **scene Location facets ⊆ Multi-facet anchors**
- Continuing situations cite stagings; L/R and stations do not flip without Design update
