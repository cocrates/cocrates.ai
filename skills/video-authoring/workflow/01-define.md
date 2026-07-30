# Stage ① — Define (`overview.md`)

**Prerequisites:** None — first stage of a new video project.

**Gate artifact:** `{project-root}/overview.md`

**Next stage:** `02-design.md` (after user approves `overview.md`)

---

## Production Model

- **Segment** is the primary unit for Evaluate / Generate Components / Assemble.
- **Short**: one implicit segment — clip design will live inside `sequence.md` (no `segments/` folder).
- **Long**: multiple segments under `segments/`; final stitch via `assembly/final.yaml`.
- **Clip** is the atomic message unit (designed in Markdown; composed later in assembly YAML).

---

## Procedure

1. Use Socratic dialogue to lock the project foundation (one question at a time).
2. After answers, write `overview.md` (do not leave chat-only).
3. Include **Validation Criteria** as concrete checks for Stage ③.
4. Lock **Short vs Long** and a sketch of **reference needs** (what must stay visually consistent across clips).
5. Do not design clip messages or media YAML in this stage.

Ask in the **user's language**. Example prompts (English shown; localize when the user writes in another language, e.g. Korean: “한 줄로 이 영상의 목적/메시지는?”):

| Question | Purpose |
|---|---|
| “In one line, what is this video’s purpose/message?” | Intent |
| “What form/genre? (film, animation, ad, music video, short-form, education/explainer, …)” | Form |
| “Who is the audience, and what tone?” | Audience + tone |
| “Rough length, and Short (one segment) vs Long (multiple segments)?” | Scale + mode |
| “Any visuals that must stay consistent across clips? (characters, backgrounds, slides, product, logo — or none)” | Reference-need sketch |
| “Taboos / out of scope?” | Safety |
| “What would make this feel successful (Validation Criteria)?” | Evaluate criteria |

---

## `overview.md` Template

Human-facing field labels and prose may use the **user's language**. Structure below is shown in English:

```markdown
# {Title}

## Basic info
- Title: {title}
- Genre / form: {film / animation / ad / music-video / short-form / education / explainer / ...}
- Target audience: {audience}
- Tone / mood: {tone}
- Production mode: {Short | Long}
- Expected length: {e.g. ~45s / ~3min}
- (Long) segment count (approx): {n}
- (Short/Long) clip count (approx): {n}

## Intent / message summary
{what the video must communicate — beginning, middle, end or beat outline}

## Reference need sketch
- Needed: {yes / no}
- Candidate kinds: {characters / locations / slides / props / brand / none / ...}
- Consistency goal: {what must not drift across clips}

## Style direction
- Visual: {style notes}
- Audio: {voice / BGM / SFX intent at high level}
- References: {optional}

## Constraints
- Taboos: {taboos}
- Out of scope: {deliberately excluded}
- Technical budget: motion video clips < ~10s; TTS-led clips may be ~30s

## Validation Criteria

| Criterion | How to verify |
|-----------|---------------|
| {e.g. Core message clear without reading a script} | {e.g. Each clip message is graspable in its duration} |
| {e.g. Pacing fits genre} | {e.g. No motion clip over ~10s; narration clips within ~30s} |
| {e.g. Visual consistency where declared} | {e.g. Declared reference entities match across clips} |
| {e.g. Taboo-compliant} | {e.g. No content from Taboos} |
```

---

## Completeness Check

- [ ] Title, genre, audience, tone filled
- [ ] Short/Long locked
- [ ] Validation Criteria are testable in Stage ③
- [ ] Reference-need sketch present (including explicit “none”)
- [ ] Intent summary matches user request

---

## Gate G1

User approves `overview.md`.

**Do not proceed to Stage ② without approval.**
