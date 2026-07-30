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

1. Load `overview.md` **Validation Criteria**.
2. Load `sequence.md` (+ segment file if Long) and `references.md`.
3. Run checks below; write `evaluations/{nnn}-{segment-slug}.md`.
4. If failures: Design-first revision → re-evaluate → new G3.

### Checks

1. **Criteria Check** — each Validation Criterion against this segment’s clips.
2. **Craft / completeness**
   - Every clip has: message, direction guide, required tracks, hook (or resolution on the last clip)
   - Short: clip section lives in `sequence.md`
   - Messages alone make the beat imaginable
3. **Message–composition fit**
   - Track choices match the message (e.g. long narration → still/hold + speech, not forced 10s motion)
   - Motion clips planned &lt;~10s; TTS-led clips may be ~30s
4. **Reference integrity** (if catalogs exist)
   - Clips only use declared entities/states
   - Expression/pose/camera are clip direction, not reference-state abuse
   - If references.md says none: confirm no hidden identity-drift risk (or escalate to add a catalog)
5. **Sequence continuity**
   - Hooks connect; arc role of segment is clear
6. **Persona / audience lens** (genre-appropriate)
   - Viewer: clarity, engagement
   - Stakeholder (client/parent/teacher): safety, accuracy, brand
   - Craft: pacing, audio/visual split

---

## Evaluation record template

`evaluations/{nnn}-{segment-slug}.md` (prose may use the user's language):

```markdown
# Segment {nnn} evaluation — {title}

## 1. Criteria Check
| Criterion | Result | Evidence |
|-----------|--------|----------|
| {from overview} | ✅ / ⚠️ / ❌ | {clip / note} |

## 2. Craft Checks
| Check | Result | Evidence |
|-------|--------|----------|
| All required clip fields present | ✅ / ⚠️ / ❌ | |
| Clip message alone makes the beat imaginable | ✅ / ⚠️ / ❌ | |
| Track choice ↔ message / duration budget | ✅ / ⚠️ / ❌ | |
| Hooks / arc role | ✅ / ⚠️ / ❌ | |
| Short: sequence.md includes segment-level content | ✅ / ⚠️ / ❌ / N/A | |

## 3. Reference Integrity
| Check | Result | Evidence |
|-------|--------|----------|
| Matches references.md plan | ✅ / ⚠️ / ❌ | |
| No undeclared entities | ✅ / ⚠️ / ❌ | |
| state vs clip direction separation | ✅ / ⚠️ / ❌ / N/A | |

## 4. Audience / Craft Notes
- Viewer: {…}
- Stakeholder: {…}
- Rhythm / audio: {…}

## 5. Revisions (Design-First)
| # | Finding | Severity | Design file | Fix | Status |
|---|---------|----------|-------------|-----|--------|
| 1 | | High/Med/Low | | | todo/done |

## 6. Design Lock Readiness (G3)
- [ ] Criteria / Craft / Reference checks acceptable
- [ ] Clip messages and direction frozen for Stage ④–⑤
```

---

## Design-First Revision Loop

1. Edit Design Markdown (`sequence.md`, `segments/*`, `references/*`, `context.md`).
2. Update evaluation record.
3. Seek new G3 for affected segment(s).
4. Only then enter Stage ④ for that segment.

---

## Gate G3 — Design Lock (per segment)

User confirms:

1. Criteria addressed for this segment?
2. Clip messages strong and lockable?
3. Reference plan intact?
4. **Design lock:** clip messages, direction, and track *needs* are frozen for component generation and assembly.

**Do not proceed to Stage ④ for this segment without G3.**

After G3: any story/message change requires Stage ② → new G3 before regenerating components or re-assembling.
