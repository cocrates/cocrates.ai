# Stage ③ — Evaluate (Criteria + Craft + Visual Reference Integrity + Story Lock)

**Prerequisites:** Approved Stage ② design artifacts (`illustration-guide.md`, `world-bible.md`, `characters*`, `locations*`, `series.md`, `episodes*`).

**Gate artifacts:**
- `evaluations/{nnn}-{episode-slug}.md` (story lock decision per episode)

**Next stage:** `04-generate.md` (after user approves story lock gates)

---

## Procedure

### 3.1 Evaluate per episode

For each `episodes/{nnn}-{episode-slug}.md`:

1. Load `overview.md` and extract **Validation Criteria**.

2. Check **Schema / structural integrity first** (mechanical — any ❌ blocks G3 story lock):
   - **`illustration-guide.md` exists** with art style + layout + typography/text-role sections (missing or empty → ❌)
   - **Flat page schema only:** zero matches for `#### Page story`, `#### Illustration guide`, `#### Rendering text`, `#### Text–image split`, `#### Page-turn hook`
   - Every page field uses `- **Field:**` with colon **inside** bold
   - Every `### Page` has required fields: Page story, Staging, Characters, Location, Direction, Rendering text, Text carries, Picture carries, Page-turn hook (values may be `none`)
   - **`### Page 0` exists** as the first page heading (cover) — missing Page 0 → ❌
   - **Cast roster** ↔ page Characters (ghost / missing on-page → ❌; mention-only tagged)
   - Craft Notes page count == measured `### Page` headings (incl. Page 0); `series.md` matches or exception noted

3. Check **episode/page craft completeness** (design correctness):
   - Page 0 / cover rule followed (see cover criteria below)
   - Page-turn hooks exist (Page 0 = opening curiosity; final = resolution/afterimage)
   - Text–image split is not redundant captioning
   - Rendering text fits target age (read-aloud rhythm + word density; Page 0 = title display OK)
   - No didactic closing monologue (theme emerges via final scene/image)
   - **Beat chronology / causality:** page order must not contradict physical sequence (e.g. already dressed → then towel-dry; food eaten → then first bite). If order is wrong → ❌ Design fix (swap/reorder pages), not Generate patch

   **Page 0 cover criteria (all should pass):**
   - Heading is `### Page 0` (optional `(Cover)` label)
   - Page story identifies cover role (invite to open; no climax spoiler)
   - Rendering text is **title** (subtitle/credit optional) — not story-body prose
   - Text carries = title; Picture carries = world / protagonist / mood
   - Direction notes title safe zone (or equivalent space for title)
   - Page-turn hook invites into Page 1

4. Check **Reference Model Integrity** (core of picture-book scene consistency — `workflow/reference-models.md`):
   - Character: state in catalogs; identity gear stable; co-cast **visual contrast**; taxonomy; expression only in Direction
   - Location: every page Location is a **catalog path** (no free text) ∈ Scene List; **neutral** set (no weather/time in location PNG); visibility coverage §3.1
   - Staging: mandatory for multi-page ensembles; header binds catalog states + location path; **establishing** default (no OTS/CU staging refs); props/ambient/situation environment when needed
   - Same staging span: seats/L-R + props stay recognizable

5. Check **Scene Continuity & Visual Consistency**:
   - Spans treated as “the same scene” keep the same character state / location scene (position+view) / state
   - When change is needed, the change type is explicit (character: outfit/gear state; location: physical-structure state; staging: reseat / new situation / documented progressive props)
   - `position` / `view` are defined by “what is visible”
   - **Same staging span:** seats/L-R **and** situation props (e.g. meal layout) stay recognizable page-to-page — not reinvented

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
| `illustration-guide.md` present (art + layout + typography roles) | ✅ / ⚠️ / ❌ | {file} |
| Flat page schema only (no nested `####` page subsections) | ✅ / ⚠️ / ❌ | {file / line} |
| Field notation `- **Field:**` (colon inside bold) | ✅ / ⚠️ / ❌ | {page} |
| Required page fields present (value `none` OK) | ✅ / ⚠️ / ❌ | {page} |
| `### Page 0` cover heading present (first page) | ✅ / ⚠️ / ❌ | {heading} |
| Cast roster ↔ page Characters (no ghosts; mention-only tagged) | ✅ / ⚠️ / ❌ | {roster vs pages} |
| Craft Notes page count == measured `### Page` count (incl. Page 0) | ✅ / ⚠️ / ❌ | {n vs n} |
| series.md pages match measured (or exception noted) | ✅ / ⚠️ / ❌ | {series row} |

## 2. Craft Checks (Design correctness)
| Check | Result | Evidence |
|-------|--------|----------|
| Page 0 is cover (title overlay; invite, no climax spoiler) | ✅ / ⚠️ / ❌ | {page} |
| Page 0 Rendering text = title (not story-body prose) | ✅ / ⚠️ / ❌ | {rendering text} |
| Page 0 Picture carries world / hero / mood; title safe zone | ✅ / ⚠️ / ❌ | {Direction / Picture carries} |
| Page-turn hook (Page 0 opening curiosity; final resolution/afterimage) | ✅ / ⚠️ / ❌ | {page} |
| Text–image split (no redundant caption fluff) | ✅ / ⚠️ / ❌ | {page} |
| Read-aloud rhythm / age density | ✅ / ⚠️ / ❌ | {page} |
| No didactic closing monologue | ✅ / ⚠️ / ❌ | {last page} |
| Characters/locations are catalog entities only | ✅ / ⚠️ / ❌ | {entities} |
| Beat chronology / causality (no dressed-before-dry, etc.) | ✅ / ⚠️ / ❌ | {page range} |

## 3. Reference Model Integrity Checks (Visual lock)

### 3.1 Character reference model (= state)
| Check | Result | Evidence |
|-------|--------|----------|
| Every page names character state (or base) | ✅ / ⚠️ / ❌ | {page + character-slug} |
| State reflects physical change only (outfit / gear / carried items) | ✅ / ⚠️ / ❌ | {page} |
| Co-appearing majors have ≥2-axis visual contrast | ✅ / ⚠️ / ❌ | {characters} |
| World taxonomy obeyed (when defined) | ✅ / ⚠️ / ❌ / n/a | {world-bible} |
| Expression / posture / action stay in page guide, not as state | ✅ / ⚠️ / ❌ | {page} |

### 3.2 Location reference model (= state) + scene definition
| Check | Result | Evidence |
|-------|--------|----------|
| Every page Location is catalog path (no free text) | ✅ / ⚠️ / ❌ | {page} |
| Path ∈ Scene List; Description names what is visible | ✅ / ⚠️ / ❌ | {location file} |
| **Visibility coverage:** must-see fixtures ⊆ cited position×view ref | ✅ / ⚠️ / ❌ | {page + ref} |
| Location refs neutral (no weather/time baked in) | ✅ / ⚠️ / ❌ | {location notes} |
| Span environment on staging; one-off on Direction | ✅ / ⚠️ / ❌ | {page / staging} |
| Continuous scene spans keep position / view / state consistent | ✅ / ⚠️ / ❌ | {page range} |

### 3.3 Staging reference model (= continuing-situation continuity)
| Check | Result | Evidence |
|-------|--------|----------|
| Mandatory staging rule satisfied (`none` only when allowed) | ✅ / ⚠️ / ❌ | {multi-page ensemble spans} |
| Header binds catalog character states + location path | ✅ / ⚠️ / ❌ | {staging file} |
| Props / ambient / situation environment maps when needed | ✅ / ⚠️ / ❌ | {staging file} |
| Canonical establishing only by default (no OTS/CU staging refs) | ✅ / ⚠️ / ❌ | {staging file} |
| Pages cite the staging | ✅ / ⚠️ / ❌ | {page} |
| No unauthorized L/R · seat · prop teleport | ✅ / ⚠️ / ❌ | {page range} |
| Progressive prop notes in Direction only if Continuity allows | ✅ / ⚠️ / ❌ | {page} |

### 3.4 Direction Completeness
| Check | Result | Evidence |
|-------|--------|----------|
| Direction / Characters / Location include needed refs | ✅ / ⚠️ / ❌ | {page} |
| Page story does not contradict Direction display elements | ✅ / ⚠️ / ❌ | {page} |

### 3.5 Series illustration guide + text consistency
| Check | Result | Evidence |
|-------|--------|----------|
| `illustration-guide.md` locks art style + layout + typography roles | ✅ / ⚠️ / ❌ | {file sections} |
| Guide recipes are specific enough to prevent page-to-page font/box drift | ✅ / ⚠️ / ❌ | {guide quotes} |
| Episode pages do not invent alternate fonts/boxes (placement only) | ✅ / ⚠️ / ❌ | {page} |
| Reading order / text zones align with guide §2 | ✅ / ⚠️ / ❌ | {page} |
| Each overlay line maps to a locked text role (title/narration/dialogue/emphasis) | ✅ / ⚠️ / ❌ | {page} |
| Anchor areas respect face / key-visual rule from guide | ✅ / ⚠️ / ❌ | {page} |
| Ages 8+ dense pages use guide §3.5 box rules (if applicable) | ✅ / ⚠️ / ❌ | {page} |

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
- **`illustration-guide.md` completeness** (art / layout / typography roles specific enough to prevent page drift): {rating}
- **Compliance plan for Stage ④:** YAML must apply guide tokens — do **not** invent per-episode fonts
- Page-specific placement notes only (anchors / reading order) — styles stay in the series guide:
  - Reading order / text-zone plan: {per guide §2 + any page exceptions}
  - Role map for this episode’s lines: {which lines = title / narration / dialogue / emphasis}
  - Guide gaps to fix in Design (if any): {missing recipes → update illustration-guide.md, not YAML}

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
- [ ] **Series `illustration-guide.md` lock + text consistency: ✅/⚠️/❌** (fonts/boxes from guide, not invented per page)
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
