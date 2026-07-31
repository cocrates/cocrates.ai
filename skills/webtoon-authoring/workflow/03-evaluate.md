# Stage ③ — Evaluate (Criteria + Craft + Visual Reference Integrity + Story Lock)

**Prerequisites:** Approved Stage ② design artifacts (`world-bible.md`, `characters*`, `locations*`, `stagings*` when needed, `series.md`, `episodes*`).

**Gate artifacts:**
- `evaluations/{nnn}-{episode-slug}.md` (story lock decision per episode)

**Next stage:** `04-generate.md` (after user approves story lock gates)

---

## Procedure

### 3.1 Evaluate per episode

For each `episodes/{nnn}-{episode-slug}.md`:

1. Load `overview.md` and extract **Validation Criteria** + canvas specs.

2. Check **Schema / structural integrity first** (mechanical — any ❌ blocks G3 story lock):
   - **Flat cut schema only:** zero matches for `###### Cut story`, `###### Composition`, `###### Illustration guide`, `###### Balloons`, `###### Captions`, `###### Gutter` (and hybrids that nest Balloons/Gutter under Illustration guide)
   - Every cut field uses `- **Field:**` with colon **inside** bold (reject `**Field**:` and bare `- Location:`)
   - Every `##### Cut` has required fields present: Cut story, Purpose/Size class/Height/Shape, Staging, Characters, Location, Direction, Balloons, Captions, Gutter (values may be `없음`)
   - Balloon/caption tables (when not `없음`) use 2-space indent under the field
   - **Size class ↔ Height:** standard ≈400–600px; tall/long ≥1,200px; compact &lt;400px sparingly — contradiction is ❌ (`workflow/cut-composition.md`)
   - **Cast roster:** episode `**Characters:**` list ↔ union of cut `- **Characters:**` (ghost = roster without on-cut and without mention tag → ❌; on-cut missing from roster → ❌). Mention-only must be tagged `(언급)` (or equivalent)
   - **Cut count:** Craft Notes cut count == number of `##### Cut` headings; `series.md` row matches measured pages/cuts or Craft Notes documents `exception noted`

3. Check **episode/page/cut craft completeness** (design correctness):
   - Page 0 / cover-thumbnail rule followed (if agreed)
   - Every `Page {N}` has `page story`, `estimated height`, `viewport plan`, `scroll hook` (non-final), `cut list`
   - Spoken lines have speaker + balloon type + placement
   - Scroll hooks on non-final pages (or intentional end-of-beat gutters)
   - Balloon/caption text is not mere caption-of-the-drawing
   - No didactic closing monologue caption (theme via final scene)
   - Vertical readability: key face/action + balloons plausible within a phone viewport beat
   - Page heights treated as **variable** (sum of cuts + gutters)
   - Cut composition follows `workflow/cut-composition.md`:
     - dialogue/everyday → roughly standard height + normal gutter
     - climax/spectacle → tall/long where overwhelm is intended
     - silence/time → pause gutters, not accidental empty gaps
     - action → open/diagonal + tight gutters
     - size classes vary with drama (not flat unless intentional)

4. Check **Reference Model Integrity** (see `workflow/reference-models.md`):
   - Characters: each cut’s state (or base) is in catalogs; **identity gear** (weapon/outfit/accessories) does not silently swap; expression/mood/pose only in Direction
   - Locations: each cut’s location + position + view (+ state) is in catalogs; state only for lasting **set** change; time/season/weather are cut direction
   - Stagings: every multi-cut ensemble with fixed seating/formation cites a staging; cast states + location anchor match the staging file; planned ref views listed
   - Cuts in a staging span do not reseat characters without a Design update

5. Check **Scene Continuity & Visual Consistency**:
   - Continuous cut/page runs keep the same character state / location scene / state / staging blocking
   - Change points name the kind of change (character state vs location state vs staging reseat vs transient)
   - Gutters match scene transitions / silence / speed

6. Check **Balloon–Art Collaboration**:
   - Cut story core is carried by art + balloons (+ captions)
   - Balloon placement does not cover face/key action
   - Reading order is top→bottom (left→right within a cut if needed)
   - Scroll hooks set up the next cut/page

7. **Persona checks** — run the [Default persona set](#persona-reference); do not skip for speed. Lock **Target Reader** from `overview.md` first.
8. Record results into `evaluations/{nnn}-{episode-slug}.md`, including per-persona critique and **Adjudication**.

---

### 3.2 Evaluation Record Template (put on disk)

`evaluations/{nnn}-{episode-slug}.md`:

```markdown
# Episode {nnn} Evaluation

## 1. Criteria Check (from overview.md Validation Criteria)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| {criterion from overview} | ✅ / ⚠️ / ❌ | {page.cut / quote / note} |

## 1b. Schema / Structural Integrity (any ❌ blocks G3)
| Check | Result | Evidence |
|-------|--------|----------|
| Flat cut schema only (no `######` cut subsections / hybrid nesting) | ✅ / ⚠️ / ❌ | {file / line} |
| Field notation `- **Field:**` (colon inside bold) | ✅ / ⚠️ / ❌ | {page.cut} |
| Required cut fields present (incl. Characters/Captions/Shape; `없음` OK) | ✅ / ⚠️ / ❌ | {page.cut} |
| Balloon/caption tables 2-space indented when present | ✅ / ⚠️ / ❌ | {page.cut} |
| Size class ↔ Height ranges agree | ✅ / ⚠️ / ❌ | {page.cut} |
| Cast roster ↔ cut Characters (no ghosts; mention-only tagged) | ✅ / ⚠️ / ❌ | {roster vs cuts} |
| Craft Notes cut count == measured `##### Cut` count | ✅ / ⚠️ / ❌ | {n vs n} |
| series.md pages/cuts match measured (or exception noted) | ✅ / ⚠️ / ❌ | {series row} |

## 2. Craft Checks (Design correctness)
| Check | Result | Evidence |
|-------|--------|----------|
| All Page required fields present | ✅ / ⚠️ / ❌ | {page} |
| Page 0 cover/thumbnail rule | ✅ / ⚠️ / ❌ | {page} |
| Scroll hook (non-final pages) | ✅ / ⚠️ / ❌ | {page} |
| Balloon speaker/type/text/placement | ✅ / ⚠️ / ❌ | {page.cut} |
| Size class matches purpose (standard/tall/open/…) | ✅ / ⚠️ / ❌ | {page.cut} |
| Gutter class/distance/intent specified | ✅ / ⚠️ / ❌ | {page.cut} |
| Pause gutters used for silence/time (not accidental) | ✅ / ⚠️ / ❌ | {page.cut} |
| Viewport plan present / one-screen beat readable | ✅ / ⚠️ / ❌ | {page} |
| Variable page height (sum of cuts + gutters) | ✅ / ⚠️ / ❌ | {page} |
| No didactic closing monologue | ✅ / ⚠️ / ❌ | {last beat} |
| Characters/locations only from catalogs | ✅ / ⚠️ / ❌ | {entities} |

## 3. Reference Model Integrity Checks (Visual lock)

### 3.1 Character reference model (= state)
| Check | Result | Evidence |
|-------|--------|----------|
| Every cut names character state (or base) | ✅ / ⚠️ / ❌ | {page.cut + character-slug} |
| State is lasting body + outfit/identity gear only | ✅ / ⚠️ / ❌ | {page.cut} |
| Identity gear does not silently swap across cuts | ✅ / ⚠️ / ❌ | {page.cut range} |
| Expression/mood/pose only in cut guide | ✅ / ⚠️ / ❌ | {page.cut} |

### 3.2 Location reference model (= set/stage)
| Check | Result | Evidence |
|-------|--------|----------|
| Every cut names location + position + view | ✅ / ⚠️ / ❌ | {page.cut} |
| Location state only for lasting set change | ✅ / ⚠️ / ❌ | {page.cut} |
| Time/season/weather/camera mood are cut direction | ✅ / ⚠️ / ❌ | {page.cut} |
| Continuous scenes keep position/view/state consistent | ✅ / ⚠️ / ❌ | {page range} |

### 3.3 Staging reference model (= continuing-situation blocking)
| Check | Result | Evidence |
|-------|--------|----------|
| Each continuing situation with fixed placement has a staging | ✅ / ⚠️ / ❌ | {café/OR/meeting span + staging-slug} |
| Staging has 2–3 planned ref views | ✅ / ⚠️ / ❌ | {staging file} |
| Cuts cite staging + matching cast/location | ✅ / ⚠️ / ❌ | {page.cut} |
| No L/R flip, seat shuffle, or OR station drift without Design reseat | ✅ / ⚠️ / ❌ | {page.cut range} |

### 3.4 Direction Completeness
| Check | Result | Evidence |
|-------|--------|----------|
| Direction includes needed character / location / staging refs | ✅ / ⚠️ / ❌ | {page.cut} |
| Cut story does not contradict Direction | ✅ / ⚠️ / ❌ | {page.cut} |

### 3.5 Balloon / Caption Guide
| Check | Result | Evidence |
|-------|--------|----------|
| Reading order top→bottom (L→R within cut if needed) | ✅ / ⚠️ / ❌ | {page.cut} |
| Balloons do not cover face/key action | ✅ / ⚠️ / ❌ | {page.cut} |
| Dialogue vs caption distinction clear | ✅ / ⚠️ / ❌ | {page.cut} |
| Text volume does not hurt scroll readability | ✅ / ⚠️ / ❌ | {page.cut} |

## 4. Scene Continuity & Visual Consistency Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Continuity across adjacent cuts (incl. staging seats) | ✅ / ⚠️ / ❌ | {page.cut range} |
| Change kind distinguished (character state vs location state vs staging reseat vs transient) | ✅ / ⚠️ / ❌ | {page.cut} |
| Gutters match transition/psychology | ✅ / ⚠️ / ❌ | {page.cut} |

## 5. Balloon–Art Collaboration Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Cut story core carried by art+balloon(+caption) | ✅ / ⚠️ / ❌ | {page.cut} |
| Scroll hook creates natural tension/curiosity | ✅ / ⚠️ / ❌ | {page} |

## 6. Persona Checks
{Run every persona in the Default set. Per persona: Stance / Strengths / Defects (severity + High/Med/Low + fix) / Reader impact.}

### Target Reader (locked audience)
- Who: {from overview — platform, genre appetite, why they open this episode}
- Fun: {rating}
- Clarity / followability: {rating}
- Want to keep scrolling: {rating}
- First-viewport / early-cut retention risk: {…}
- Feedback: {feedback}

### Genre Critic
- Genre promise vs delivery: {rating}
- Trope use (earned vs empty): {rating}
- Feedback: {feedback}

### Plot / Pacing Critic
- Episode arc completeness: {rating}
- Cut/page rhythm & information timing: {rating}
- Scroll-hook chain: {rating}
- Feedback: {feedback}

### Character Critic
- Motivation readable in art+balloon: {rating}
- Voice distinctness in dialogue: {rating}
- Feedback: {feedback}

### Platform / editorial
- Width/color/upload-split feasibility: {rating}
- One-screen readability: {rating}
- Feedback: {feedback}

### Story Critic (webtoon craft)
- Story completeness without didacticism: {rating}
- Cut division / size-class variety vs drama: {rating}
- Balloon–art collaboration: {rating}
- Feedback: {feedback}

### Webtoon art specialist
- Reference-model consistency (character / location / staging): {rating}
- Vertical direction, cut heights & gutters: {rating}
- Size-class variety vs drama: {rating}
- Balloon guide (for Stage ④ YAML): {feedback}
  - Reading order: {top→bottom + balloon map}
  - Balloon list: {cut → speaker / type / text / placement / emphasis}
  - Caption style: {caption-box rules}
  - Composition map: {cut → size class / height / gutter class}
  - Staging cites: {cut → staging-slug + ref view}
  - Outside-cut fill / side margins: {full bleed or 30–50px}

## 7. Adjudication
{Record Apply decisions. Default tie-break: Target Reader. Never silently drop a High finding.}

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Where to apply | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|----------------|--------|

## 8. Revisions (Design-First)
| # | Finding | Severity | Where to apply | Proposed fix | Action Status |
|---|---------|-----------|----------------|--------------|----------------|
| 1 | {finding} | High/Med/Low | characters/... / locations/... / episodes/... | {edit} | {todo/done} |

## 9. Story Lock Readiness (G3)
- [ ] **Schema / Structural Integrity: ✅** (any ❌ → G3 blocked; fix Design first)
- [ ] Reference Model Integrity (character / location / staging): ✅/⚠️/❌
- [ ] Scene Continuity & Visual Consistency: ✅/⚠️/❌
- [ ] Direction Completeness: ✅/⚠️/❌
- [ ] Balloon/Caption Guide: ✅/⚠️/❌
- [ ] Craft Checks: ✅/⚠️/❌
- [ ] Required personas run + Adjudication filled: ✅/⚠️/❌
- [ ] Target-reader readiness acceptable: ✅/⚠️/❌
```

---

## Persona Reference

Critics **advise**. Fill the Adjudication table and apply chosen fixes before G3. Default tie-break: **Target Reader** from `overview.md`, unless overview locks another priority (e.g. platform compliance over spectacle).

### Default persona set (every episode evaluation)

**Required:** Target Reader · Genre Critic · Plot / Pacing Critic · Story Critic · Platform / editorial · Webtoon art specialist

**Also required when:** character-driven or dialogue-heavy episode → **Character Critic**; episode 001 / series premiere → emphasize Target Reader **first-viewport retention** and Genre promise.

### Persona catalog

| Persona | Focus | Typical questions |
|---------|-------|-------------------|
| **Target Reader** | Phone-scroll lived experience for the locked audience | Fun? Followable? Keep scrolling? Where do they bounce? |
| **Genre Critic** | Genre contracts and trope payoff | Does this episode pay the genre bill? |
| **Plot / Pacing Critic** | Arc, causality, hook chain, info timing | Dead mid-scroll? Hooks earned? |
| **Character Critic** | Motive, face-acting, dialogue voice | Do I care who speaks? |
| **Story Critic** | Completeness, cut rhetoric, no preachy caption ending | Is the story told in cuts, not lectures? |
| **Platform / editorial** | Width, color, split upload, one-screen beats | Shipable on the locked canvas? |
| **Webtoon art specialist** | Refs, gutters, heights, balloons, staging | Will Stage ④ stay consistent and readable? |

### Evaluation perspectives (rubric summary)

| Perspective | Focus | Typical questions |
|---|---|---|
| Target Reader | Understanding / fun / scroll | Is it fun? Followable? Want to keep scrolling? |
| Genre / Plot / Character / Story | Narrative craft | Promise, pace, people, completeness |
| Platform / editorial | Specs / readability | Readable at locked width? Split upload feasible? |
| Webtoon art | Visual consistency | Refs, balloons, cut heights, vertical space intact? |

---

## Design-First Revision Loop

If Evaluation finds issues:

1. Return to **Stage ② (Design)** and update the relevant Design files (`characters/*`, `locations/*`, `stagings/*`, `episodes/*`, …).
2. Re-run affected checks and update `evaluations/{nnn}-{episode-slug}.md`.
3. Only after user approves the story lock (G3), proceed to Stage ④.

---

## Approval Gate G3 — Story Lock (per episode)

**Prerequisite:** Section **1b Schema / Structural Integrity** must be all ✅. Schema ❌ blocks story lock — return to Stage ② and fix flat cut schema / roster / counts / size↔height before personas or lock.

User confirms:

1. Schema / Structural Integrity — flat cut schema, field notation, required fields, cast roster, cut counts all ✅?
2. Criteria Check — all items addressed?
3. Craft Checks — satisfactory?
4. Reference Model Integrity — character states / location set states / stagings consistent with catalogs (`workflow/reference-models.md`)?
5. Scene Continuity & Visual Consistency — continuity held (including seating); gutters/changes intentional?
6. Persona feedback + Adjudication — residual risk acceptable for the **Target Reader**?
7. **Story lock:** episode page/cut stories, balloon texts, captions, gutters, scroll hooks, staging cites, and Direction fields are now frozen for image generation.

**Do not proceed to Stage ④ until G3 is approved.**

After G3:
- Treat the episode’s design as locked.
- Any later change to story/text/craft/characters/locations/world rules requires rollback to Stage ② and a new G3 approval before regenerating images.
