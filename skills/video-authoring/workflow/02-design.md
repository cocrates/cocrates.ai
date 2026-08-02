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

**Shared three-layer model** (when characters/locations apply): `workflow/reference-models.md` — character turnarounds + contrast/taxonomy; location **neutral** multi-view + **catalog-path cites only**; staging **segment-first**, mandatory for multi-clip ensembles, **default 1× establishing** (no OTS/CU as staging refs), lean WHO–WHERE–WHAT, Character+Location in refs. Explicit approval before every generate.

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

#### Canonical clip schema (mandatory)

**Only the flat clip schema is allowed.** Every clip is a `### Clip {N} — {slug}` heading followed by a **flat** `- **Field:**` list. Do **not** use nested `#### Clip message`, `#### Direction guide`, `#### Required tracks`, or `#### Hook to next clip` subsections (nested or hybrid layouts).

**Field notation (fixed):** Always `- **Field:** value` — colon **inside** bold (correct: `**Visual:**`; wrong: `**Visual**:` or `- Visual:`). Empty values still require the field (`none` / `no` as appropriate).

#### Segment / sequence cast roster (when catalogs exist)

```markdown
**Characters:** {slug-a}, {slug-b} (언급)
**Locations:** {location-a}
```

Rules: on-clip ⊆ roster; ghost cast forbidden; mention-only tagged. If `references.md` says none, omit roster.

#### Required clip fields (flat list, this order)

Each `### Clip {N}` **must** include every field below:

| Field | Rule |
|-------|------|
| `- **Clip message:**` | Beat this clip alone carries (Stage ④ message basis) |
| `- **Visual:**` | still / motion from keyframe / hold / … |
| `- **References:**` | kind/slug/state; staging/{slug}+view when placement must hold; or `none` |
| `- **Direction:**` | camera / action / expression — clip-specific |
| `- **On-screen text:**` | exact copy in target language, or `none` |
| `- **Tracks — image:**` | yes/no |
| `- **Tracks — video:**` | yes/no |
| `- **Tracks — voice:**` | yes/no (+ script summary when yes) |
| `- **Tracks — bgm/sfx:**` | yes/no (+ role when yes) |
| `- **Hook to next:**` | curiosity/transition; final: resolution/aftertaste |

Craft Notes **Clip count** = measured `### Clip` headings; Long Segment List approx clips sync to measured after design, or exception noted.

```markdown
# Sequence: {Title}

## Overall story / message
{arc — also serves as the single segment story}

**Characters:** {slug-a}, {slug-b} (언급)  <!-- omit if references: none -->
**Locations:** {location-a}

## Craft Notes
- Genre rhythm: {…}
- Core message: {…}
- Climax clip: {id}
- Clip count (this segment): {n}  <!-- MUST equal ### Clip headings -->
- Segment List sync: {updated to measured clips | exception noted | N/A short}

## Clips

### Clip {N} — {slug}
- **Clip message:** {story/info/emotion — imaginable standalone}
- **Visual:** {still image / motion from keyframe / hold frame / …}
- **References:** {kind/slug/state or staging/{slug}+view | none}
- **Direction:** {camera / action / expression / …}
- **On-screen text:** {exact copy in target language | none}
- **Tracks — image:** {yes/no}
- **Tracks — video:** {yes/no}
- **Tracks — voice:** {yes/no — script summary}
- **Tracks — bgm/sfx:** {yes/no — role}
- **Hook to next:** {curiosity/transition | final: resolution/aftertaste}
```

---

### 2.3 Segment design (Long only)

For each `segments/{nnn}-{segment-slug}.md` — same **flat clip schema** as Short:

```markdown
# Segment {nnn}: {Title}

## Summary
{episode-level story / teaching beat / ad act}

**Characters:** {…}  <!-- when catalogs exist -->
**Locations:** {…}

## Craft Notes
- Role in the sequence arc: {…}
- Core message: {…}
- Climax clip: {id}
- Clip count (this segment): {n}
- sequence.md Segment List sync: {updated | exception noted}

## Clips

### Clip {N} — {slug}
- **Clip message:** {…}
- **Visual:** {…}
- **References:** {…}
- **Direction:** {…}
- **On-screen text:** {…}
- **Tracks — image:** {yes/no}
- **Tracks — video:** {yes/no}
- **Tracks — voice:** {yes/no}
- **Tracks — bgm/sfx:** {yes/no}
- **Hook to next:** {…}
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
- [ ] Short: `sequence.md` contains full clip designs (segment-level content) using **flat clip schema**
- [ ] Long: every segment file has all clip messages + track needs using **flat clip schema** (no nested `####` clip subsections)
- [ ] Cast roster present when catalogs exist; no ghost cast
- [ ] Craft Notes clip count matches measured `### Clip` headings; Segment List synced or exception noted
- [ ] Every clip message is imaginable standalone
- [ ] Track choices fit duration budgets (motion &lt;~10s; TTS-led ~30s ok)

---

## Gate G2

User approves the design set:

- Reference plan appropriate to intent? (not over-forcing characters/locations)
- Sequence arc clear?
- Flat clip schema only throughout?
- Short: clip section in `sequence.md` sufficient?
- Long: each segment’s clips sufficient?
- Messages strong enough to lock in Evaluate?

**Do not proceed to Stage ③ without approval.**
