# Stage ② — Design (References + Sequence / Segment / Clip Messages)

**Prerequisites:** Approved `overview.md`

**Gate artifacts:**
- `{project-root}/references.md` (+ optional `references/**`, `context.md`)
- `{project-root}/sequence.md`
- Long only: `segments/{nnn}-{segment-slug}.md`

**Next stage:** `03-evaluate.md` (after G2)

---

## Design constraints (global)

- **Markdown owns story/message.** No media generation. No edit-spec YAML as design substitute.
- **Short**: there is one implicit segment. Put **all segment-level content** (clip messages, direction, track needs) **inside `sequence.md`**. Do not create `segments/`.
- **Long**: `sequence.md` = series-level arc + segment list; each `segments/*.md` = episode-level story + clips.
- **Composition, not invention**: clips only reference entities declared in catalogs (when catalogs exist). While designing story/clips, **co-design** staging + character states + location states — **reuse** when present, **add** when missing (`workflow/reference-models.md` §7). Missing entity → stop, update Design, resume.
- **Design-first fixes**: Evaluate findings update Markdown first, not Stage ④/⑤ YAML alone.
- **User language**: Design Markdown prose may use the user's language (e.g. Korean clip messages); keep structure/keys consistent.


### Reference models (conditional)

A **reference model** defines visual identity that must hold across clips. Stage ④ materializes it as reference images.

**Shared three-layer model** (when characters/locations apply): `workflow/reference-models.md` — **character** (body + equipment identity), **location** (set/stage), **staging** (who is where in a continuing multi-clip situation: café L/R, OR stations, meeting seats, …). One staging per continuing situation; typically 2–3 ensemble reference views.

Concrete the intent analysis from Define:

| Question | Result |
|----------|--------|
| What must look the same across clips? | Declare catalog kinds |
| Are references unnecessary? | State **none** in `references.md` → skip Phase 0 |
| Do multi-clip situations need fixed placement? | Declare kind **stagings** (+ character/location kinds) |

Example kinds: `characters`, `locations`, `stagings`, `slides`, `props`, `brand` — declare per project. Unlike picture books, characters/locations are not always required — but if they are used across clips with continuing situations, **stagings are required** for placement continuity.


---

## Procedure

### 2.1 Reference need + catalogs

1. Write `references.md` (required).
2. If kinds are needed, create catalogs under `references/{kind}/` (index + per-entity files as appropriate).
3. Optional: `context.md` for domain/world/style rules when useful (education scripts, fictional worlds, brand voice).

```markdown
<!-- references.md -->
# Reference model plan

## Intent summary
{from overview}

## Reference needed?
- Needed: {yes / no}
- Why: {why consistency matters or why not}

## Catalog kinds
| Kind | Consistency goal | Entity count (approx) |
|------|------------------|------------------------|
| {characters / locations / stagings / slides / … / none} | {…} | {n} |

## Stage ④ impact
- Phase 0 reference images: {required / skip}
```

Character/location/staging catalogs, when used, follow `workflow/reference-models.md` (state = lasting body + equipment identity; location = set; staging = continuing-situation blocking; expression/pose/camera/time/weather in clip direction). Slides/brand/props: define the stable visual identity fields appropriate to that kind.

---

### 2.2 Sequence design (`sequence.md`)

**Always create `sequence.md`.**

#### Long mode

```markdown
# Sequence: {Title}

## Overall story / message arc
{series-level narrative or argument arc}

## Segment List

| # | Slug | Summary | Clips (approx) | Length (approx) |
|---|------|---------|----------------|-----------------|
| 1 | {slug} | {summary} | {n} | {e.g. 40s} |

## Tone / rhythm notes
{pacing, peaks, callbacks}
```

Clip detail lives in `segments/*.md`.

#### Short mode (single segment)

`sequence.md` **must include everything a `segments/*.md` would contain**: segment story + all clips.

```markdown
# Sequence: {Title}

## Overall story / message
{arc — also serves as the single segment story}

## Craft Notes
- Genre rhythm: {…}
- Core message: {…}
- Climax clip: {id}

## Clips

### Clip {N} — {slug}

#### Clip message
{Story/info/emotion this clip must deliver — reading this alone should make the beat imaginable}
{Basis for Stage ④ image/video `message`}

#### Direction guide
- Visual: {still image / motion from keyframe / hold frame / …}
- References: {kind/slug/state or none; include staging/{slug}+view when placement must hold}
- Clip-specific direction (camera/action/expression/…): {…}
- On-screen text (if any): {exact copy in the target language, or none}

#### Required tracks
- image: {yes/no — clip key image}
- video: {yes/no — motion <~10s}
- voice (speech): {yes/no — script summary}
- bgm/sfx (music): {yes/no — role}

#### Hook to next clip
{curiosity/transition — final clip: resolution/aftertaste}
```

---

### 2.3 Segment design (Long only)

For each `segments/{nnn}-{segment-slug}.md`:

```markdown
# Segment {nnn}: {Title}

## Summary
{episode-level story / teaching beat / ad act}

## Craft Notes
- Role in the sequence arc: {…}
- Core message: {…}
- Climax clip: {id}

## Clips

### Clip {N} — {slug}

#### Clip message
{…}

#### Direction guide
- Visual: {…}
- References: {…}
- Clip-specific direction: {…}
- On-screen text: {…}

#### Required tracks
- image / video / voice / bgm-sfx: {as above}

#### Hook to next clip
{…}
```

---

### 2.4 Internal feedback loop (story ↔ catalogs)

**Story design co-locks catalogs** — see `workflow/reference-models.md` §7 (when visual catalogs are used).

While designing sequence / segments / clips, **together** design (when kinds are declared):

1. **Staging** for each continuing multi-clip situation (cite existing or **add** under `references/stagings/`)
2. **Character appearance states** (cite existing or **add** to character catalog)
3. **Location set states** (cite existing or **add** to location catalog)

**Reuse if present; add if missing.** Do not invent outfit, gear, set damage, seats, or L/R only inside clip direction.

If a clip needs an undeclared reference entity, state, staging, or track type:

1. Update `references.md` / catalogs / `context.md` first (user-visible Design update).
2. Re-check affected sequence/segment files.
3. Cite the new/existing refs in clip direction — do not invent mid-clip.

---

## Completeness Check (Stage ②)

- [ ] `references.md` exists (explicit none allowed)
- [ ] Declared catalogs exist when kinds ≠ none
- [ ] Story units cite staging + character states + location states; new refs were **added to catalogs** before citing (or existing reused) — `workflow/reference-models.md` §7
- [ ] `sequence.md` exists
- [ ] Short: `sequence.md` contains full clip designs (segment-level content)
- [ ] Long: every segment file has all clip messages + track needs
- [ ] Every clip message is imaginable standalone
- [ ] Track choices fit duration budgets (motion &lt;~10s; TTS-led ~30s ok)

---

## Gate G2

User approves the design set:

- Reference plan appropriate to intent? (not over-forcing characters/locations)
- Sequence arc clear?
- Short: clip section in `sequence.md` sufficient?
- Long: each segment’s clips sufficient?
- Messages strong enough to lock in Evaluate?

**Do not proceed to Stage ③ without approval.**
