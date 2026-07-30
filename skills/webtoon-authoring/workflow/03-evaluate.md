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

2. Check **episode/page/cut craft completeness** (design correctness):
   - Page 0 / cover-thumbnail rule followed (if agreed)
   - Every `Page {N}` has `page story`, `estimated height`, `viewport plan`, `scroll hook` (non-final), `cut list`
   - Every `Cut {n}` has `cut story`, `purpose`, `size class`, `height guide`, `shape`, `illustration guide`, `balloons` (when needed), `captions` (when needed), `gutter class/distance/intent`
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
     - episode cut count near overview target (or short-form exception noted)

3. Check **Reference Model Integrity** (see `workflow/reference-models.md`):
   - Characters: each cut’s state (or base) is in catalogs; **identity gear** (weapon/outfit/accessories) does not silently swap; expression/mood/pose only in cut guide
   - Locations: each cut’s location + position + view (+ state) is in catalogs; state only for lasting **set** change; time/season/weather are cut direction
   - Stagings: every multi-cut ensemble with fixed seating/formation cites a staging; cast states + location anchor match the staging file; 2–3 ref views planned
   - Cuts in a staging span do not reseat characters without a Design update

4. Check **Scene Continuity & Visual Consistency**:
   - Continuous cut/page runs keep the same character state / location scene / state / staging blocking
   - Change points name the kind of change (character state vs location state vs staging reseat vs transient)
   - Gutters match scene transitions / silence / speed

5. Check **Balloon–Art Collaboration**:
   - Cut story core is carried by art + balloons (+ captions)
   - Balloon placement does not cover face/key action
   - Reading order is top→bottom (left→right within a cut if needed)
   - Scroll hooks set up the next cut/page

6. Persona checks:
   - Reader (scroll immersion / comprehension / desire to keep scrolling)
   - Platform/editorial (readability, width, color, upload-split feasibility)
   - Critic (quality, rhythm, cut division / size-class variety)
   - Webtoon art specialist (character/location/staging refs, balloons, gutters, cut heights, vertical direction)

7. Record results into `evaluations/{nnn}-{episode-slug}.md`.

---

### 3.2 Evaluation Record Template (put on disk)

`evaluations/{nnn}-{episode-slug}.md`:

```markdown
# Episode {nnn} Evaluation

## 1. Criteria Check (from overview.md Validation Criteria)
| Criterion | Result | Evidence |
|-----------|--------|----------|
| {criterion from overview} | ✅ / ⚠️ / ❌ | {page.cut / quote / note} |

## 2. Craft Checks (Design correctness)
| Check | Result | Evidence |
|-------|--------|----------|
| All Page/Cut required fields present | ✅ / ⚠️ / ❌ | {page.cut} |
| Page 0 cover/thumbnail rule | ✅ / ⚠️ / ❌ | {page} |
| Scroll hook (non-final pages) | ✅ / ⚠️ / ❌ | {page} |
| Balloon speaker/type/text/placement | ✅ / ⚠️ / ❌ | {page.cut} |
| Size class + height guide present | ✅ / ⚠️ / ❌ | {page.cut} |
| Size class matches purpose (standard/tall/open/…) | ✅ / ⚠️ / ❌ | {page.cut} |
| Gutter class/distance/intent specified | ✅ / ⚠️ / ❌ | {page.cut} |
| Pause gutters used for silence/time (not accidental) | ✅ / ⚠️ / ❌ | {page.cut} |
| Viewport plan present / one-screen beat readable | ✅ / ⚠️ / ❌ | {page} |
| Variable page height (sum of cuts + gutters) | ✅ / ⚠️ / ❌ | {page} |
| Cut-count near overview target (or exception noted) | ✅ / ⚠️ / ❌ | {episode} |
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

### 3.4 Cut Guide Completeness
| Check | Result | Evidence |
|-------|--------|----------|
| Illustration guide includes needed character / location / staging refs | ✅ / ⚠️ / ❌ | {page.cut} |
| Cut story does not contradict the guide | ✅ / ⚠️ / ❌ | {page.cut} |

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
### Reader
- Fun: {rating}
- Clarity: {rating}
- Want to keep scrolling: {rating}
- Feedback: {feedback}

### Platform / editorial
- Width/color/upload-split feasibility: {rating}
- One-screen readability: {rating}
- Feedback: {feedback}

### Critic
- Story completeness: {rating}
- Cut division / rhythm: {rating}
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

## 7. Revisions (Design-First)
| # | Finding | Severity | Where to apply | Proposed fix | Action Status |
|---|---------|-----------|----------------|--------------|----------------|
| 1 | {finding} | High/Med/Low | characters/... / locations/... / episodes/... | {edit} | {todo/done} |

## 8. Story Lock Readiness (G3)
- [ ] Reference Model Integrity (character / location / staging): ✅/⚠️/❌
- [ ] Scene Continuity & Visual Consistency: ✅/⚠️/❌
- [ ] Cut Guide Completeness: ✅/⚠️/❌
- [ ] Balloon/Caption Guide: ✅/⚠️/❌
- [ ] Craft Checks: ✅/⚠️/❌
```

---

### 3.3 Evaluation perspectives (rubric)
| Perspective | Focus | Typical questions |
|---|---|---|
| Reader | Understanding / fun / scroll | Is it fun? Followable? Want to keep scrolling? |
| Platform / editorial | Specs / readability | Readable at locked width? Color + split upload feasible? |
| Critic | Completeness / rhythm | Story complete? Cut division and size/gutter variety working? |
| Webtoon art | Visual consistency | Character/location/staging refs, balloons, cut heights, vertical space intact? |

---

## Design-First Revision Loop

If Evaluation finds issues:

1. Return to **Stage ② (Design)** and update the relevant Design files (`characters/*`, `locations/*`, `stagings/*`, `episodes/*`, …).
2. Re-run affected checks and update `evaluations/{nnn}-{episode-slug}.md`.
3. Only after user approves the story lock (G3), proceed to Stage ④.

---

## Approval Gate G3 — Story Lock (per episode)

User confirms:

1. Criteria Check — all items addressed?
2. Craft Checks — satisfactory?
3. Reference Model Integrity — character states / location set states / stagings consistent with catalogs (`workflow/reference-models.md`)?
4. Scene Continuity & Visual Consistency — continuity held (including seating); gutters/changes intentional?
5. Persona feedback — residual risk acceptable?
6. **Story lock:** episode page/cut stories, balloon texts, captions, gutters, scroll hooks, staging cites, and illustration guides are now frozen for image generation.

**Do not proceed to Stage ④ until G3 is approved.**

After G3:
- Treat the episode’s design as locked.
- Any later change to story/text/craft/characters/locations/world rules requires rollback to Stage ② and a new G3 approval before regenerating images.
