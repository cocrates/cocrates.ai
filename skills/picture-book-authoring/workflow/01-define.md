# Stage ① — Define (`overview.md`)

**Prerequisites:** None — first stage of a new picture book.

**Gate artifact:** `{project-root}/overview.md`

**Next stage:** `02-design.md` (after user approves `overview.md`)

---

## Publication Model (unit of generation)

- **Episode** is the generation unit: each `episodes/{nnn}-{episode-slug}.md` is evaluated and then generates page images into `images/{nnn}-{episode-slug}/`.
- **Page** design lives inside the episode file (story + rendering text + illustration guide).
- **Page 0 = cover** (mandatory): every episode starts with a cover page; body pages are Page 1….

---

## Procedure

1. Use Socratic dialogue to lock the book foundation (one question at a time).
2. After the user answers, write `overview.md` (do not leave it as chat-only).
3. Include **Validation Criteria** as concrete, testable checks used in Stage ③ Evaluate.
4. Do not design episodes or images in this stage.

### What to ask (one question at a time)

| Question | Purpose |
|---|---|
| What is the one-line logline (the story in one sentence)? | Core concept |
| Target age (e.g. 3–5) and read-aloud / illustration difficulty feel? | Age range + constraints |
| How many episodes? | Episode List scale |
| Rough pages per episode? (e.g. around 10 / 20 / 30 / 40, **incl. Page 0 cover**) | Episode size (approx only) — actual episode design may change; every episode includes a cover as Page 0 |
| Theme / genre and tone (warm / dreamy / comic, …)? | Craft direction |
| Illustration style (watercolor / digital painting / collage, …) and reference works? | Visual style constraints |
| Episode/book title for cover typography (if known)? | Page 0 rendering text seed |
| Taboos / out of scope? | Safety |
| What success criteria should Evaluate use (Validation Criteria)? | Evaluate criteria |

---

## `overview.md` Template

```markdown
# {Title}

## Basic Info
- Title: {title}
- Target age: {age range}
- Episode count: {episode count}
- Pages per episode (approx): {pages per episode} (e.g. around 10 / 20 / 30 / 40) — **includes Page 0 cover**; adjustable in actual episode design
- Theme / genre: {theme}
- Tone / mood: {tone}
- Rendering language: {e.g. English / Korean / bilingual} — page overlay text language (Page 0 = title in this language)

## Story Summary
{story summary — beginning, middle, end}

## Character Overview
- Protagonist: {name} — {one-line description}
- Supporting: {name} — {one-line description}

## Setting Overview
- World: {world name} — {one-line description}
- Key region: {region} — {one-line description}
- Key location: {location} — {one-line description}

## Illustration Style
{Seed only — expand into `illustration-guide.md` in Stage ②.}
- Medium: {watercolor / digital painting / collage / ...}
- Palette: {color palette description}
- Line: {outline style}
- Mood: {mood}
- References: {reference artists or styles if any}
- Typography intent (seed): {e.g. rounded friendly Hangul; soft narration; bold dialogue}

## Constraints
- Taboos: {taboos}
- Out of scope: {deliberately excluded}

## Validation Criteria
{Checks used in Stage ③ Evaluate — concrete, testable success conditions.}

| Criterion | How to verify |
|-----------|---------------|
| {e.g. Target age can follow the plot from pictures + short text} | {e.g. Each page readable aloud in ≤20 seconds} |
| {e.g. Every episode has Page 0 cover with title} | {e.g. `### Page 0` present; rendering text = title; cover invites opening} |
| {e.g. Series typography stays consistent across pages} | {e.g. Pages follow `illustration-guide.md` narration/dialogue/title recipes; no one-off fonts} |
| {e.g. Page-turn curiosity through climax} | {e.g. Pages 0–N-1 have a turn hook; Page 0 = open-the-book curiosity} |
| {e.g. Theme earned by scene, not closing sermon} | {e.g. Final page shows outcome; no moral monologue} |
| {e.g. Safe / taboo-compliant} | {e.g. No content from Taboos} |
```

**Multilingual note:** If `Rendering language` is not English, write locked page **rendering text** in that language (e.g. Korean for a Korean edition). Image-generation YAML **prompts** remain English; only the overlay strings use the target language.

---

## Completeness Check

- [ ] Title/age/scale fields are filled in episode/page terms (pages approx **incl. Page 0 cover**)
- [ ] Validation Criteria are concrete enough to be checked in Stage ③
- [ ] Story summary matches the user’s intent
- [ ] Illustration style and taboos are specific
- [ ] Rendering language is locked

---

## Gate

User approves `overview.md`.

**Do not proceed to Stage ② without approval.**
