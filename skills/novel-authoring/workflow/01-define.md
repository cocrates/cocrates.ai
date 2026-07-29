# Stage ① — Define (`overview.md`)

**Prerequisites:** None — first stage of a new novel.

**Gate artifact:** `{project-root}/overview.md`

**Next stage:** `02-plan.md` (after user approves `overview.md`)

---

## Publication Model (read first)

**Chapter is the only publication and release unit.** One approved chapter design → one chapter manuscript → one release.

**Episode** is an internal design subdivision within a chapter (scene grouping). It is **not** serialized, released, or planned at define stage. **Scene** design lives inside episode sections of the chapter design file (stage ④).

At stage ①, plan and ask about **chapters only** — never episodes.

---

## Procedure

Surface the novel's foundation through Socratic dialogue. Ask **one question at a time**.

| Question | Purpose |
|----------|---------|
| *"What is the novel about? What is its one-sentence logline?"* | Core concept |
| *"Who is the intended reader? What genre and tone?"* | Audience & genre |
| *"What is the target length — roughly how many chapters, and how long is each chapter?"* | Scale (approximate) + Structure Mode |
| *"How will you publish — chapter-by-chapter serialization, batch, or complete volume?"* | Release model |
| *"Are there content boundaries, taboos, or mandatory elements?"* | Scope / out of scope |
| *"What will make this novel successful in your eyes?"* | Validation criteria |

### Structure Mode (from chapter scale)

After the scale answer, lock **Structure Mode** in `overview.md`:

| Mode | Guideline | Planning consequence |
|------|-----------|----------------------|
| **Short** | ~1 volume, under ~20 chapters | Stage ② puts **part + chapter lists** in `series.md`; **no** `parts/` |
| **Series** | Multi-part / long-form (typically ≥ ~20 chapters or multiple major parts) | Stage ② Part Catalog only; stage ③ expands each part under `parts/` |

If the user is unsure, propose from chapter count and confirm. **Borderline (~18–22 chapters or “maybe multi-part”):** ask explicitly which mode they want; prefer **series** if they expect several major parts.

When Mode later changes, update via [`consistency.md`](consistency.md) cascade (overview → series → parts/chapter lists).

### Do NOT ask at define stage

| ❌ Do not ask | Why |
|--------------|-----|
| Total episode count | Episodes are design-internal; decided at stage ④ per chapter |
| Episodes per chapter | Decided at **chapter design** (stage ④) per story needs |
| Words per episode | Plan words per **chapter** |
| Episode release schedule | Release is **chapter**-by-chapter |
| "How many episodes will you serialize?" | Wrong unit — ask **chapters** |

If the user mentions episodes, clarify: *"This workflow publishes by chapter. Episodes are internal design blocks inside a chapter — we'll decide how many when we design each chapter. For now, roughly how many chapters do you plan?"*

Record into `{project-root}/overview.md`:

```markdown
# Novel Definition: {Title}

## Logline
{One sentence}

## Genre & Tone
{Primary genre(s), mood, voice}

## Audience
{Target reader profile}

## Scale
{Planning targets for length management — **not fixed requirements**. Adjust as the story develops.}

- **Structure Mode:** {short | series}
- **Total chapters (approx.):** {n}
- **Words per chapter (approx.):** {target}
- **Total word count (optional estimate):**
- **Part breakdown (optional):** {rough part count and approximate chapters per part — e.g. Part 1 ~20 ch, Part 2 ~25 ch}

## Publication
- **Unit:** Chapter — one manuscript file per chapter; one release per chapter
- **Cadence:** {e.g. weekly one chapter, batch of 5, complete volume}

## Scope
{What is included}

## Out of Scope
{What is deliberately excluded}

## Validation Criteria
{Checks that confirm the novel fulfills its purpose — used in evaluation}

## Constraints
{Deadlines, platform requirements, taboos, chapter-level serialization needs}

## References
{Influences, reference works, prior materials}
```

---

## Completeness Check

- [ ] Logline is one clear sentence
- [ ] Structure Mode is set (`short` or `series`) consistent with approximate chapter count
- [ ] Scale gives **approximate** chapter count and words per chapter — no episode count
- [ ] Scale is framed as planning guidance, not a rigid contract
- [ ] Publication section states chapter as release unit
- [ ] Validation criteria are concrete enough for stage ⑥ evaluation
- [ ] No episode-based planning in the artifact

---

## Gate

User approves `overview.md`. Do not proceed to stage ② without approval.
