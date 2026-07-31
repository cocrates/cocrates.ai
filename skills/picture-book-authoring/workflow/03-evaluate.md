# Stage ③ — Evaluate (Criteria + Craft + Visual Reference Integrity + Story Lock)

**Prerequisites:** Approved Stage ② design artifacts (`world-bible.md`, `characters*`, `locations*`, `series.md`, `episodes*`).

**Gate artifacts:**
- `evaluations/{nnn}-{episode-slug}.md` (story lock decision per episode)

**Next stage:** `04-generate.md` (after user approves story lock gates)

---

## Procedure

### 3.1 Evaluate per episode

For each `episodes/{nnn}-{episode-slug}.md`:

1. Load `overview.md` and extract **Validation Criteria**.

2. Check **Schema / structural integrity first** (mechanical — any ❌ blocks G3 story lock):
   - **Flat page schema only:** zero matches for `#### Page story`, `#### Illustration guide`, `#### Rendering text`, `#### Text–image split`, `#### Page-turn hook`
   - Every page field uses `- **Field:**` with colon **inside** bold
   - Every `### Page` has required fields: Page story, Staging, Characters, Location, Direction, Rendering text, Text carries, Picture carries, Page-turn hook (values may be `없음`)
   - **Cast roster** ↔ page Characters (ghost / missing on-page → ❌; mention-only tagged)
   - Craft Notes page count == measured `### Page` headings; `series.md` matches or exception noted

3. Check **episode/page craft completeness** (design correctness):
   - Page 0 / cover rule followed
   - Page-turn hooks exist (final = resolution/afterimage)
   - Text–image split is not redundant captioning
   - Rendering text fits target age (read-aloud rhythm + word density)
   - No didactic closing monologue (theme emerges via final scene/image)

4. Check **Reference Model Integrity** (core of picture-book scene consistency):
   - Character reference model (= state) integrity
     - Each page names a **state** (or base) defined in `characters/{character-slug}.md`
     - Expression / posture / action / emotion stay in Direction (page image direction), not as reference-model states
     - Character state updates only when the story has a lasting physical / outfit change
   - Location reference model (= state) integrity
     - Each page names a location with **scene axes (position + view)** and that scene’s **state (base or state-slug)** from `locations/{location-slug}.md`
     - Lighting / time / weather / mood / camera direction / framing / angle / transient elements are page direction, not reference-model state
     - “Outside-the-window view” treated as **transient** → page prompt, no state change
     - Permanent physical change (new building / wall destroyed) → new state

5. Check **Scene Continuity & Visual Consistency**:
   - Spans treated as “the same scene” keep the same character state / location scene (position+view) / state
   - When change is needed, the change type is explicit (character: outfit/gear state; location: physical-structure state)
   - `position` / `view` are defined by “what is visible”

6. Check **Text–Image Collaboration**:
   - Core page-story information is sufficiently delivered by `rendering text + what the image must show`
   - “No overlap” is not forced, but redundancy must not erase important information
   - Page-turn hooks naturally set up the next page

7. Persona checks — run the [Default persona set](#persona-reference); lock **Target Child Reader** (age band from overview) first.
8. Record results into `evaluations/{nnn}-{episode-slug}.md`, including per-persona critique and **Adjudication**.

---

### 3.2 Evaluation Record Template (put on disk)

`evaluations/{nnn}-{episode-slug}.md`:

```markdown
# Episode {nnn} Evaluation

## 1. Criteria Check (from overview.md Validation Criteria)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| {criterion from overview} | ✅ / ⚠️ / ❌ | {page / quote / note} |

## 1b. Schema / Structural Integrity (any ❌ blocks G3)
| Check | Result | Evidence |
|-------|--------|----------|
| Flat page schema only (no nested `####` page subsections) | ✅ / ⚠️ / ❌ | {file / line} |
| Field notation `- **Field:**` (colon inside bold) | ✅ / ⚠️ / ❌ | {page} |
| Required page fields present (value `없음` OK) | ✅ / ⚠️ / ❌ | {page} |
| Cast roster ↔ page Characters (no ghosts; mention-only tagged) | ✅ / ⚠️ / ❌ | {roster vs pages} |
| Craft Notes page count == measured `### Page` count | ✅ / ⚠️ / ❌ | {n vs n} |
| series.md pages match measured (or exception noted) | ✅ / ⚠️ / ❌ | {series row} |

## 2. Craft Checks (Design correctness)
| Check | Result | Evidence |
|-------|--------|----------|
| Page 0 is cover | ✅ / ⚠️ / ❌ | {page} |
| Page-turn hook (incl. final resolution/afterimage) | ✅ / ⚠️ / ❌ | {page} |
| Text–image split (no redundant caption fluff) | ✅ / ⚠️ / ❌ | {page} |
| Read-aloud rhythm / age density | ✅ / ⚠️ / ❌ | {page} |
| No didactic closing monologue | ✅ / ⚠️ / ❌ | {last page} |
| Characters/locations are catalog entities only | ✅ / ⚠️ / ❌ | {entities} |

## 3. Reference Model Integrity Checks (Visual lock)

### 3.1 Character reference model (= state)
| Check | Result | Evidence |
|-------|--------|----------|
| Every page names character state (or base) | ✅ / ⚠️ / ❌ | {page + character-slug} |
| State reflects physical change only (outfit / gear / carried items) | ✅ / ⚠️ / ❌ | {page} |
| Expression / posture / action stay in page guide, not as state | ✅ / ⚠️ / ❌ | {page} |
| State updates only when physical change occurs | ✅ / ⚠️ / ❌ | {page} |

### 3.2 Location reference model (= state) + scene definition
| Check | Result | Evidence |
|-------|--------|----------|
| Every page names location + position + view | ✅ / ⚠️ / ❌ | {page} |
| Lighting / time / weather / mood / camera / framing are not reference-model state (view = “what is visible”) | ✅ / ⚠️ / ❌ | {page} |
| Outside-window change rule is correct
|  - Transient (temporary / reversible) change is not a state | ✅ / ⚠️ / ❌ | {page + example} |
|  - Physical-structure change (new building / wall destroyed, …) → new state | ✅ / ⚠️ / ❌ | {page + example} |
| Continuous scene spans keep position / view / state consistent | ✅ / ⚠️ / ❌ | {page range} |


### 3.3 Staging reference model (= continuing-situation blocking)
| Check | Result | Evidence |
|-------|--------|----------|
| Continuing situations have a staging | ✅ / ⚠️ / ❌ | {span + staging-slug} |
| Staging plans 2–3 reference views | ✅ / ⚠️ / ❌ | {staging file} |
| Pages cite the staging | ✅ / ⚠️ / ❌ | {page} |
| No unauthorized L/R · seat · station swaps | ✅ / ⚠️ / ❌ | {page range} |

### 3.4 Direction Completeness
| Check | Result | Evidence |
|-------|--------|----------|
| Direction / Characters / Location include needed refs | ✅ / ⚠️ / ❌ | {page} |
| Page story does not contradict Direction display elements | ✅ / ⚠️ / ❌ | {page} |

### 3.5 Illustration text (visual effects) guide fitness
| Check | Result | Evidence |
|-------|--------|----------|
| Reading order for `rendering text` (left→right, top→bottom) is natural | ✅ / ⚠️ / ❌ | {page} |
| Dialogue / exclamation emphasis (largest + bold; glow/outline when useful) is clear | ✅ / ⚠️ / ❌ | {page} |
| Narration style (smaller + soft rounded serif + warm shadow) is consistent | ✅ / ⚠️ / ❌ | {page} |
| Key words/phrases are emphasized by size or color | ✅ / ⚠️ / ❌ | {page} |
| Anchor area suggested so text does not cover faces / key visuals | ✅ / ⚠️ / ❌ | {page} |
| Ages 8+ & text-heavy pages separate body copy into an in-image story section (panel/box) — recommended | ✅ / ⚠️ / ❌ | {page} |

## 4. Scene Continuity & Visual Consistency Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Continuous pages keep identity (character / place / structure) | ✅ / ⚠️ / ❌ | {page range} |
| Change moments distinguish change type (character state vs location state vs transient) | ✅ / ⚠️ / ❌ | {page} |
| position/view defined by “what is visible” so the scene is clear | ✅ / ⚠️ / ❌ | {page} |

## 5. Text–Image Collaboration Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Core page-story info is sufficiently delivered by text+image | ✅ / ⚠️ / ❌ | {page} |
| Redundancy / omission does not collapse next-page expectation | ✅ / ⚠️ / ❌ | {page} |
| Page-turn hook connects as natural tension / curiosity | ✅ / ⚠️ / ❌ | {page} |

## 6. Persona Checks
{Run the full required set. Per persona: Stance / Strengths / Defects(finding+High/Med/Low+fix) / Reader impact.}

### Target Child Reader (locked age band)
- Who: {overview age band · reading context — alone / read-aloud}
- Fun: {rating}
- Understanding: {rating}
- Want to turn the page: {rating}
- Confused / scary / boring stretches: {…}
- Feedback: {feedback}

### Caregiver / Parent (buyer · read-aloud)
- Value / share willingness: {rating}
- Safety (age fit): {rating}
- No sermon overload: {rating}
- Read-aloud rhythm: {rating}
- Feedback: {feedback}

### Genre / Age Critic (picture-book genre · development)
- Density / vocabulary vs age-band contract: {rating}
- Genre / form expectations (repetition / accumulation / humor, …): {rating}
- Feedback: {feedback}

### Story Critic (critique · story)
- Story completeness: {rating}
- Character appeal · motivation: {rating}
- Text–image collaboration · rhythm: {rating}
- Meaning without forced moral: {rating}
- Feedback: {feedback}

### Illustration specialist
- Reference-model consistency: {rating}
- Scene clarity (what is visible): {rating}
- Page framing lock compliance: {rating}
- Text visual-effects guide provided (font / emphasis / reading order / anchor area / panel split): {rating}
- Text overlay guide (for Stage ④ YAML): {feedback}
  - Reading order anchor plan: {left→right / top→bottom placement rules + text-area map}
  - TextBox list: {line-index → anchor area / size / font style / color / emphasis}
  - Dialogue/exclamation style: {largest+bold + (glow/outline, …) application}
  - Narration style: {smaller font + (serif/sans) + shadow application}
  - Key word emphasis: {which words/phrases and how}
  - Ages 8+ text-heavy case: {body panel/box split yes/no + position/size}

### Educator / Librarian (optional — required when overview locks school/library distribution)
- Fit to learning / socio-emotional goals: {rating}
- Classroom / library read-aloud fitness: {rating}
- Feedback: {feedback}

## 7. Adjudication
{Record Apply decisions. Default tie-break: **Target Child Reader**. Do not quietly drop High findings.}

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Child first) | Apply at | Status |
|---|-------------------|----------|-----------|--------|--------------------------------|-----------|--------|

## 8. Revisions (Design-First)
| # | Finding | Severity | Apply at (Design file / Stage ④ YAML overlay) | Improvement (what to change) | Action Status |
|---|---------|-----------|---------------------------|---------------------------|----------------|
| 1 | {finding} | High/Med/Low | characters/... / locations/... / episodes/... | {proposed edit} | {todo/done} |
| 2 | {finding} | ... | ... | ... | ... |

## 9. Story Lock Readiness (G3 check)
- [ ] **Schema / Structural Integrity: ✅** (any ❌ → G3 blocked; fix Design first)
- [ ] Reference Model Integrity: ✅/⚠️/❌ all items
- [ ] Scene Continuity & Visual Consistency: ✅/⚠️/❌ all items
- [ ] Direction Completeness: ✅/⚠️/❌ all items
- [ ] Illustration Text Styling/Effects Guide: ✅/⚠️/❌ all items
- [ ] Craft Checks: ✅/⚠️/❌ all items
- [ ] Required personas + Adjudication: ✅/⚠️/❌
- [ ] Target-child readiness: ✅/⚠️/❌
```

---

## Persona Reference

Critics **advise**. Fill the Adjudication table and apply chosen fixes before G3. Default tie-break: **Target Child Reader** (age band in `overview.md`). Caregiver safety findings that are High generally outrank “funnier but unsafe” unless overview explicitly allows.

### Default persona set (every episode evaluation)

**Required:** Target Child Reader · Caregiver / Parent · Genre / Age Critic · Story Critic · Illustration specialist

**Also required when:** overview locks school/library/education use → **Educator / Librarian**

### Persona catalog

| Persona | Focus | Typical questions |
|---------|-------|-------------------|
| **Target Child Reader** | Understanding · fun · page-turn desire for the locked age | Is it fun? Can I follow? Do I want the next page? |
| **Caregiver / Parent** | Safety · value · sermon · read-aloud | Want to share it? Too scary / preachy? |
| **Genre / Age Critic** | Age · form contract | Density / vocabulary right for this age of picture book? |
| **Story Critic** | Completeness · character · text–image collaboration | Do story and pictures share one breath? |
| **Illustration specialist** | Reference consistency · scene visibility · text overlay | Will Stage ④ hold together? |
| **Educator / Librarian** | Education · public read-aloud fitness | Usable in classroom / library? |

### Evaluation perspectives (rubric summary)

| Perspective | Focus | Typical questions |
|---|---|---|
| Target Child | Understanding / enjoyment | Fun? Followable? Want to turn? |
| Caregiver | Safety / value | Safe? Want to share? No sermon overload? |
| Genre/Age + Story | Finish quality / age fit | Density · collaboration · rhythm fit the age? |
| Illustration specialist | Visual consistency | Do reference links hold? Is the scene clearly “visible”? |

---

## Design-First Revision Loop

If Evaluation finds issues:

1. Return to **Stage ② (Design)** and update the relevant Design files:
   - `characters/{character-slug}.md` (character state)
   - `locations/{location-slug}.md` (location scene: position/view/state)
   - `episodes/{nnn}-{episode-slug}.md` (per-page illustration guide / rendering text)
2. Re-run affected checks and update `evaluations/{nnn}-{episode-slug}.md`.
3. Only after user approves the story lock (G3), proceed to Stage ④.

---

## Approval Gate G3 — Story Lock (per episode)

**Prerequisite:** Section **1b Schema / Structural Integrity** must be all ✅. Schema ❌ blocks story lock — return to Stage ② and fix flat page schema / roster / counts before personas or lock.

User confirms:

1. Schema / Structural Integrity — flat page schema, field notation, required fields, cast roster, page counts all ✅?
2. Criteria Check — all items addressed?
3. Craft Checks — satisfactory?
4. Reference Model Integrity — character states / location position-view-states defined without contradicting catalogs?
5. Scene Continuity & Visual Consistency — continuous scenes keep identity; changes only where needed?
6. Persona feedback + Adjudication — residual risk acceptable under **Target Child** priority?
7. **Story lock:** episode page stories, rendering text, text–image split, page-turn hooks, and Direction fields are now frozen for image generation.

**Do not proceed to Stage ④ until G3 is approved.**

After G3:
- Treat the episode’s design as locked.
- Any later change to story/text/craft/characters/locations/world rules requires rollback to Stage ② and a new G3 approval before regenerating images.
