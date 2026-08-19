# Stage ① — Define (`overview.md`)

**Prerequisites:** None — first stage of a new novel.

**Gate artifact:** `{project-root}/overview.md`

**Next stage:** `02-plan.md` (after user approves `overview.md`)

---

## Resolve Project Root (before first write)

Per Cocrates Workspace Convention and this skill’s **Resolve Project Root**:

- Default folder: `novels/{title-slug}/`
- **`{title-slug}` = work title only** (kebab-case). Do **not** prepend/append author name, pen name, or source-file basename.
- Inspect existing `novels/` (or workspace root). **Confirm folder name and location with the user** before creating a new project. Prefer reuse; avoid create-then-rename.
- **Do not** browse another novel’s tree “for reference” when creating a new project — inspect folder **names** only, or reuse after user confirmation.

### Project scaffold (once, same model step)

After the project root exists (or in the **same** tool round as creating it), create the standard directories **in one step**. Do **not** create `world/`, `characters/`, `locations/` in separate sequential rounds, and do **not** check a missing path just to discover it should be created.

**Required dirs** (under `{project-root}/`):

| Dir | Used from |
|-----|-----------|
| `world/` | ③ |
| `characters/` | ③ |
| `locations/` | ③ |
| `stagings/` | ④ |
| `episodes/` | ④ |
| `manuscripts/` | ⑥ |
| `evaluations/` | ⑤ / ⑦ |
| `continuity/` | ⑧ write; later load |
| `scrum/` | todo Initialize (series) |

Directory creation creates intermediate paths automatically — one call per leaf dir is enough. Later stages **must not** re-create these if they already exist; only create a dir when a path is still missing (rare) and then batch any remaining missing dirs in **one** step.

---

## Publication Model (read first)

**Episode is the only publication and release unit.** One approved episode design → one episode manuscript → one release.

**Scene** is an internal design subdivision within an episode (beat grouping). It is **not** serialized, released, or planned at define stage. **Scene** design lives inside the episode design file (stage ④).

At stage ①, plan and ask about **episodes only** — never scenes.

---

## Procedure

Surface the novel's foundation through Socratic dialogue. Ask **one question at a time**.

| Question | Purpose |
|----------|---------|
| *"What is the novel about? What is its one-sentence logline?"* | Core concept |
| *"Who is the intended reader? What genre and tone?"* | Audience & genre |
| *"What is the target length — roughly how many episodes?"* | Scale (approximate episode count) |
| *"What does each episode deliver? (title + one-paragraph role)"* | Episode List planning — record in `overview.md`, formalized into `series.md` at ② |
| *"How will you publish — episode-by-episode serialization, batch, or complete volume?"* | Release model |
| *"Are there content boundaries, taboos, or mandatory elements?"* | Scope / out of scope |
| *"What will make this novel successful in your eyes?"* | Validation criteria |

**Default episode length:** **4,000–8,000 characters** per episode unless the user specifies otherwise. Record in `overview.md` Scale. At stage ④, each scene gets an **Est. length**; **sum of scene Est. must satisfy Scale min ≤ sum ≤ Scale max** (validated at ④ / ⑤) — prefer the central band (~5,000–7,200 on the default Scale). Do not rely on manuscript padding or confirmation-loop bloat to hit Scale.

**Episode count and rough Episode List** are planning targets — useful for scale management, not fixed contracts. When counts change, update `overview.md` Scale and `series.md` Episode List with user approval.

### Do NOT ask at define stage

| ❌ Do not ask | Why |
|--------------|-----|
| Total scene count | Scenes are design-internal; decided at stage ④ per episode |
| Scenes per episode | Decided at **episode design** (stage ④) per story needs |
| Words per scene | At ① plan **episode** Scale only; at ④ assign per-scene Est. whose **sum sits in Scale min–max** |
| Scene release schedule | Release is **episode**-by-episode |
| "How many scenes will you serialize?" | Wrong unit — ask **episodes** |

If the user mentions scenes, clarify: *"This workflow publishes by episode. Scenes are internal design blocks inside an episode — we'll decide how many when we design each episode. For now, roughly how many episodes do you plan?"*

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

- **Total episodes (approx.):** {n}
- **Characters per episode (default):** 4,000–8,000 characters — use this range unless the user sets a different target. Stage ④: scene Est. sum must satisfy **Scale min ≤ sum ≤ Scale max** (prefer ~5,000–7,200).
- **Total character count (optional estimate):**
- **Arc grouping (optional):** {rough grouping — e.g. Arc 1 ~10 ep, Arc 2 ~15 ep}

## Publication
- **Unit:** Episode — one manuscript file per episode; one release per episode
- **Cadence:** {e.g. weekly one episode, batch of 5, complete volume}

## Scope
{What is included}

## Out of Scope
{What is deliberately excluded}

## Validation Criteria
{Checks that confirm the novel fulfills its purpose — used in evaluation}

## Constraints
{Deadlines, platform requirements, taboos, episode-level serialization needs}

## References
{Influences, reference works, prior materials}
```

---

## Completeness Check

- [ ] Logline is one clear sentence
- [ ] Scale gives **approximate** episode count and **characters per episode (default 4,000–8,000 characters)** — no scene count
- [ ] Scale is framed as planning guidance, not a rigid contract
- [ ] Episode List (if filled) lists titles + one-line roles only — no scene counts
- [ ] Publication section states episode as release unit
- [ ] Validation criteria are concrete enough for stage ⑦ evaluation
- [ ] No scene-based planning in the artifact

---

## Gate G1 (Define Approval)

User approves `overview.md`. Do not proceed to stage ② without approval.
