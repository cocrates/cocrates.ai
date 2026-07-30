# Stage ① — Define (`overview.md`)

**Prerequisites:** None — first stage of a new webtoon.

**Gate artifact:** `{project-root}/overview.md`

**Next stage:** `02-design.md` (after user approves `overview.md`)

---

## Publication Model (unit of generation)

- **Episode** is the generation unit: each `episodes/{nnn}-{episode-slug}.md` is evaluated, then generates page images into `images/{nnn}-{episode-slug}/`, then stitches them into `episode-scroll.png`.
- **Page** = variable-height vertical strip segment (fixed width, usually 690px). Pages concatenate top→bottom into one episode scroll.
- **Cut** = panel inside a page; design lives inside the episode file (cut art + balloons + dialogue + narration + gutters).

Picture-book assumptions (uniform page size, page-turn hooks, read-aloud caption overlay) do **not** apply. Use scroll craft instead.

---

## Procedure

1. Use Socratic dialogue to lock the webtoon foundation (one question at a time).
2. After the user answers, write `overview.md` (do not leave it as chat-only).
3. Include **Validation Criteria** as concrete, testable checks used in Stage ③ Evaluate.
4. Lock **canvas specs** (width, generation aspect, color, outside-cut fill).
5. Do not design episodes, cuts, or images in this stage.

### What to ask (one question at a time)

| Question | Purpose |
|---|---|
| What is the one-line logline? | Core concept |
| Who is the target audience (age band / platform feel)? | Audience + tone constraints |
| How many episodes? | Episode List scale |
| Rough pages (vertical segments) per episode? | Episode length (approx) — actual design may change |
| Target cut count per episode? (weekly average ~60–80; short pilot may be lower) | Cut budget / pacing scale |
| Genre / tone (romance, action, comedy, thriller, …)? | Craft direction |
| Art style (digital color comic, line-heavy, painterly, …) and reference works? | Visual style |
| Target portal / canvas width? (690–800px common; ~1080 only if required) | Canvas width |
| Cut side margins: full bleed or ~30–50px even margins? | Horizontal fit |
| Outside-cut fill color (white / black / theme)? | Outside-cut fill |
| Taboos / out of scope? | Safety |
| What success criteria should Evaluate use (Validation Criteria)? | Evaluate criteria |

---

## `overview.md` Template

```markdown
# {Title}

## Basics
- Title: {title}
- Target audience: {audience / age band}
- Episode count: {episode count}
- Approx. pages per episode: {pages per episode} — adjustable in episode design
- Target cuts per episode: {e.g. 60–80 weekly / lower for short pilot}
- Genre: {genre}
- Tone / mood: {tone}

## Canvas / format
- Reading mode: vertical scroll
- Width: {690–800px; lock exact px} — ~1080 only if portal/user requires
- Height: variable (sum of cut heights + gutters; may differ per page)
- Cut horizontal fit: {full bleed / 30–50px even side margins}
- Generation aspect (1K): 9:21 (~672×1584) default — if a page is taller, plan strip splits or stitch extension
- Color: full color (default)
- Outside-cut area: {white / black / theme color}
- Cut composition: follow `workflow/cut-composition.md` (size classes + gutter classes)
- Target portal (optional): {name / self-publish}
- Upload split (optional): if portal-limited, export slices ~1280px tall — keep design/generate as one continuous scroll

## Story summary
{story summary — beginning, middle, end / season arc if multi-episode}

## Character overview
- Protagonist: {name} — {one-line description}
- Supporting: {name} — {one-line description}

## Setting overview
- World: {world name} — {one-line description}
- Main region: {region} — {one-line description}
- Main locations: {location} — {one-line description}

## Art style
- Medium: {digital color comic / ink+flat / painterly / ...}
- Palette: {color palette description}
- Line: {outline style}
- Balloon tone: {round / jagged / thought / shout conventions}
- Mood: {mood}
- References: {reference works or styles if any}

## Constraints
- Taboos: {taboos}
- Out of scope: {deliberately excluded}

## Validation Criteria
{Checks used in Stage ③ Evaluate — concrete, testable success conditions.}

| Criterion | How to verify |
|-----------|---------------|
| {e.g. One phone-screen beat is readable without pinch-zoom} | {e.g. Each cut’s key face/action + balloon fit in ~9:21 viewport planning} |
| {e.g. Scroll curiosity through climax} | {e.g. Non-final pages/cuts end with a scroll hook or intentional gutter beat} |
| {e.g. Dialogue lives in balloons; narration is optional captions} | {e.g. Every spoken line has a balloon owner + placement} |
| {e.g. Theme earned by scenes, not sermon caption} | {e.g. Final beats show outcome; no moral monologue box} |
| {e.g. Safe / taboo-compliant} | {e.g. No content from Taboos} |
```

---

## Completeness Check

- [ ] Title / audience / scale fields are filled in episode/page/cut terms
- [ ] Canvas width, cut horizontal fit, variable height, color, outside-cut fill are locked
- [ ] Target cuts-per-episode range is set (or short-form exception noted)
- [ ] Validation Criteria are concrete enough to be checked in Stage ③
- [ ] Story summary matches the user’s intent
- [ ] Art style and taboos are specific

---

## Gate

User approves `overview.md`.

**Do not proceed to Stage ② without approval.**
