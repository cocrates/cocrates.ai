# Stage ③ — Evaluate (Criteria + Craft + Message Lock)

**Prerequisites:** Approved Stage ② design artifacts.

**Gate artifacts:** `evaluations/{nnn}-{segment-slug}.md` (Short: e.g. `001-main.md`)

**Next stage:** `04-generate-components.md` after **G3 per segment**

---

## Evaluation unit

- **Always evaluate per segment.**
- **Short:** treat the clip block inside `sequence.md` as one segment (`001-main` or title slug).
- **Long:** one evaluation file per `segments/{nnn}-*.md`.

Do not proceed to Stage ④ for a segment until that segment’s G3 is approved. Other segments may still be in Design/Evaluate.

---

## Procedure

For each segment:

1. Load `overview.md` **Validation Criteria** and lock **Target Viewer** (who watches, where, what success means in viewer terms).
2. Load `sequence.md` (+ segment file if Long) and `references.md`.
3. Run checks below; write `evaluations/{nnn}-{segment-slug}.md`.
4. Run the [Default persona set](#persona-reference) — do not skip for speed; fill Adjudication.
5. If failures: Design-first revision → re-evaluate → new G3.

### Checks

1. **Criteria Check** — each Validation Criterion against this segment’s clips.
2. **Schema / structural integrity first** (mechanical — any ❌ blocks G3):
   - **Flat clip schema only:** zero `#### Clip message`, `#### Direction guide`, `#### Required tracks`, `#### Hook to next clip`
   - Every field uses `- **Field:**` (colon inside bold)
   - Every `### Clip` has required fields (Clip message, Visual, References, Direction, On-screen text, four Tracks lines, Hook to next) — values may be `none`/`no`
   - Cast roster ↔ References/on-clip cast when catalogs exist (no ghosts)
   - Craft Notes clip count == measured `### Clip` headings; Segment List synced or exception noted
3. **Craft / completeness**
   - Short: clip section lives in `sequence.md`
   - Messages alone make the beat imaginable
4. **Message–composition fit**
   - Track choices match the message (e.g. long narration → still/hold + speech, not forced 10s motion)
   - Motion clips planned &lt;~10s; TTS-led clips may be ~30s
5. **Reference integrity** (if catalogs exist) — see `workflow/reference-models.md`
   - Character identity gear does not silently swap; expression is clip direction; co-cast contrast when needed
   - Location cites are catalog paths only; location refs neutral (not time/weather); lasting set changes only as states
   - Continuing multi-clip ensembles cite **stagings** (mandatory rule); establishing default; no L/R drift; no OTS/CU as staging refs
   - Explicit approval before every generate; design↔prompt fidelity
   - Clips only use declared entities/states
   - Expression/pose/camera are clip direction, not reference-state abuse
   - If references.md says none: confirm no hidden identity-drift risk (or escalate to add a catalog)
6. **Sequence continuity**
   - Hooks connect; arc role of segment is clear
7. **Persona / audience lens** — see Persona Reference (required set)

---

## Evaluation record template

`evaluations/{nnn}-{segment-slug}.md` (prose may use the user's language):

```markdown
# Segment {nnn} evaluation — {title}

## 1. Criteria Check
| Criterion | Result | Evidence |
|-----------|--------|----------|
| {from overview} | ✅ / ⚠️ / ❌ | {clip / note} |

## 1b. Schema / Structural Integrity (any ❌ blocks G3)
| Check | Result | Evidence |
|-------|--------|----------|
| Flat clip schema only (no nested `####` clip subsections) | ✅ / ⚠️ / ❌ | |
| Field notation `- **Field:**` (colon inside bold) | ✅ / ⚠️ / ❌ | |
| Required clip fields present (`none`/`no` OK) | ✅ / ⚠️ / ❌ | |
| Cast roster ↔ on-clip refs (when catalogs exist) | ✅ / ⚠️ / ❌ / N/A | |
| Craft Notes clip count == measured `### Clip` count | ✅ / ⚠️ / ❌ | |
| Segment List clips match measured (or exception) | ✅ / ⚠️ / ❌ / N/A | |

## 2. Craft Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Clip message alone makes the beat imaginable | ✅ / ⚠️ / ❌ | |
| Track choice ↔ message / duration budget | ✅ / ⚠️ / ❌ | |
| Hooks / arc role | ✅ / ⚠️ / ❌ | |
| Short: sequence.md includes segment-level content | ✅ / ⚠️ / ❌ / N/A | |

## 3. Reference Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Matches references.md plan | ✅ / ⚠️ / ❌ | |
| Character/location/staging rules (`reference-models.md`) | ✅ / ⚠️ / ❌ / n/a | |
| Staging for continuing multi-clip situations | ✅ / ⚠️ / ❌ / n/a | |
| No undeclared entities | ✅ / ⚠️ / ❌ | |
| state vs clip direction separation | ✅ / ⚠️ / ❌ / N/A | |

## 4. Target Viewer Checks
| Check | Result | Evidence |
|-------|--------|----------|
| Opening earns locked viewer attention | ✅ / ⚠️ / ❌ | |
| Message clarity for that viewer | | |
| Desire to watch next clip / finish segment | | |
| Drop-risk (boredom, confusion, overload) named | | |

## 5. Persona Checks
{Active set = Core + format add-ons from overview Genre/form. Per persona: Stance / Strengths / Defects (severity + High/Med/Low + fix) / Viewer impact.}

### Target Viewer
- Who: {from overview — platform, length tolerance, why they press play: entertainment / learn / decide / …}
- Clarity: {rating}
- Engagement / finish (or learn-and-apply) intent: {rating}
- Feedback: {feedback}

### Format Critic
- Form promise vs delivery (film / animation / ad / education / explainer / …): {rating}
- Convention misuse for *this* form: {rating}
- Feedback: {feedback}

### Message Critic
- Segment message / learning beat completeness: {rating}
- Cold-viewer comprehension (without prior context): {rating}
- Show vs tell calibrated to form (drama ≠ lecture-dump; education may teach explicitly if clear): {rating}
- Feedback: {feedback}

### Pacing / Editing Critic
- Clip rhythm & hook (or checkpoint) chain: {rating}
- Duration budget realism: {rating}
- Cognitive load / denseness for the locked viewer: {rating}
- Feedback: {feedback}

### Craft Director (picture / motion / assembly fit)
- Track split matches message: {rating}
- Reference plan survivable into Stage ④–⑤: {rating}
- Feedback: {feedback}

### Learning Critic
{Required for education / explainer / tutorial / training — else omit or N/A}
- Objectives covered in order: {rating}
- Accuracy / no harmful misconception: {rating}
- Practice or check-for-understanding if overview requires: {rating}
- Feedback: {feedback}

### Stakeholder (client / teacher / parent / brand)
{Required for education/explainer/training, ad/brand, or when overview names a commissioner — else omit or N/A}
- Safety, accuracy, brand / curriculum fit: {rating}
- Feedback: {feedback}

### Audio / Voice Critic
{Required when speech/TTS/BGM/SFX-led — common for education; else omit or N/A}
- Speech/BGM/SFX necessity vs clutter: {rating}
- Narration density for ear: {rating}
- Feedback: {feedback}

## 6. Adjudication
{Record Apply decisions. Default tie-break: Target Viewer (education/brand exceptions in Persona Reference). Never silently drop a High finding.}

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Viewer first; education/brand exceptions per Persona Reference) | Design file | Status |
|---|-------------------|----------|-----------|--------|----------------------------------------------------------------------------------|-------------|--------|

## 7. Revisions (Design-First)
| # | Finding | Severity | Design file | Fix | Status |
|---|---------|----------|-------------|-----|--------|
| 1 | | High/Med/Low | | | todo/done |

## 8. Design Lock Readiness (G3)
- [ ] **Schema / Structural Integrity: ✅** (any ❌ → G3 blocked)
- [ ] Criteria / Craft / Reference checks acceptable
- [ ] Required personas + Adjudication filled
- [ ] Target-viewer readiness acceptable
- [ ] Clip messages and direction frozen for Stage ④–⑤
```

---

## Persona Reference

Critics **advise**. Fill the Adjudication table and apply chosen fixes before G3.

**Tie-break:** **Target Viewer** first. Exceptions from overview:
- **Education / explainer / training:** High **Learning Critic** (accuracy / misconception) and High **Stakeholder** (curriculum/safety) outrank “more entertaining but wrong/unsafe.”
- **Ad / brand:** High **Stakeholder** brand/safety may outrank pure entertainment taste.

Video forms in `overview.md` Genre/form include film, animation, ad, music-video, short-form, **education**, **explainer**, tutorial, training, and similar — persona sets must follow the locked form, not assume narrative entertainment.

### Core required (every segment, every form)

**Target Viewer** · **Format Critic** · **Message Critic** · **Pacing / Editing Critic** · **Craft Director**

### Format add-ons (also required when overview Genre/form matches)

| Form band (from overview) | Also required |
|---------------------------|---------------|
| **Education / explainer / tutorial / training** | **Learning Critic** · **Stakeholder** · **Audio / Voice Critic** (if any speech/TTS/narration — usual) |
| **Ad / brand / promo** | **Stakeholder** |
| **Film / animation / dramatic short / music-video (story-led)** | (core only; add Character notes under Message Critic if cast-driven) |
| **Any form with speech, TTS, heavy BGM/SFX** | **Audio / Voice Critic** |

If form is ambiguous, ask once to lock Genre/form before evaluation — do not default to film/animation personas only.

### Persona catalog

| Persona | Focus | Typical questions |
|---------|-------|-------------------|
| **Target Viewer** | Locked audience’s watch experience (fun *or* learn/apply) | Clear? Worth finishing? Where do they quit or get lost? |
| **Format Critic** | Form contracts (not only cinematic genre) | Does this feel like the promised *kind* of video (lesson, ad, film, …)? |
| **Message Critic** | Beat purpose — story turn *or* instructional point | Would a cold viewer get the point? Density right for the form? |
| **Pacing / Editing Critic** | Rhythm, hooks/checkpoints, duration, cognitive load | Too slow/fast/dense? Budgets realistic? |
| **Craft Director** | Track choice, refs, assemble-ability | Will Stage ④–⑤ survive this design? |
| **Learning Critic** | Objectives, accuracy, misconception risk, checks | Are learning goals covered without false claims? |
| **Stakeholder** | Teacher/client/parent/brand constraints | Shipable for the locked commissioner/curriculum? |
| **Audio / Voice Critic** | Speech, BGM, SFX load | Ear-friendly? Tracks fighting the picture? |

---

## Design-First Revision Loop

1. Edit Design Markdown (`sequence.md`, `segments/*`, `references/*`, `context.md`).
2. Update evaluation record (personas + Adjudication).
3. Seek new G3 for affected segment(s).
4. Only then enter Stage ④ for that segment.

---

## Gate G3 — Design Lock (per segment)

**Prerequisite:** Schema / Structural Integrity all ✅. Schema ❌ blocks lock — fix flat clip schema / roster / counts in Design first.

User confirms:

1. Schema / Structural Integrity — flat clip schema, fields, roster, clip counts ✅?
2. Criteria addressed for this segment?
3. Clip messages strong and lockable?
4. Reference plan intact?
5. Persona feedback + Adjudication — residual risk acceptable for the **Target Viewer**?
6. **Design lock:** clip messages, direction, and track *needs* are frozen for component generation and assembly.

**Do not proceed to Stage ④ for this segment without G3.**

After G3: any story/message change requires Stage ② → new G3 before regenerating components or re-assembling.
