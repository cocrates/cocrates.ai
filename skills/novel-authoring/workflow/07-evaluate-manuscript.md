# Stage ⑦ — Evaluate & Revise — Manuscript (`evaluations/{nnn}-{episode-slug}-manuscript.md`)

**Prerequisites:** Approved manuscript (stage ⑥)

**Gate artifact:** `{project-root}/evaluations/{nnn}-{episode-slug}-manuscript.md`

**Next stage:** `08-release.md` (after evaluation approval) — or back to `04-episode-design.md` (structural) / `06-generate.md` (prose)

**Status:** **Main quality gate.** Every manuscript is evaluated before release — **not optional**. Design evaluation (Stage ⑤) is optional; manuscript evaluation is not.

**See also:** SKILL.md § Revision · § Consistency · [`06-generate.md`](06-generate.md).

---

## Procedure

Propose after manuscript approval (evaluation is **required**):

> *"Episode {nnn} manuscript is ready. We'll evaluate it against the episode design and prior architecture/continuity, then run the required critic personas before release."*

### Load

Follow SKILL.md **Selective artifact load** (Phase A indexes → Phase B needed details only).

- `manuscripts/{nnn}-{episode-slug}.md` — **once if not already in context**; keep this payload for **line count, character count, and every L cite**
- `episodes/{nnn}-{episode-slug}.md` (full design) — Est. sum, Exposition Budget, Dialogue intent, Seeds/Hold — **this episode only**, not prior episode designs
- **Phase A:** `overview.md`, `series.md`, `characters.md`, `locations.md`, `world-bible.md`, `stagings.md` (if missing)
- **Phase B:** Architecture + continuity set listed in the episode design (same set as Stage ⑥ Pre-Generation Load) — Appearing/used/cited only; ep 002+ → `story-so-far` + immediate prior summary only
- **Appearing cast profiles** (`characters/{slug}.md`) for every name in Characters Appearing — required before scoring **Uncatalogued proper nouns** (check **appearance states / base gear**, not indexes alone)
- Used **location** profiles via Architecture References **kebab paths** (never Display-with-spaces filenames). If place existence is disputed, follow the same path discipline as Stage ⑤ (`05-evaluate-design.md` § Path existence discipline)
- Re-run **Design-Fidelity Gate** from `06-generate.md`
- Lock **Target Reader** from `overview.md` before persona critique

**Do not** batch-read unused mid/late-arc profiles or every file under `episodes/` / `continuity/`.

**Mandatory measurements (same turn, before Floor scores):**

1. File **line count** → obtain the line count into the header  
2. File **character count** → obtain the character count (whole file, including H1) into the header **and** Length Evidence — **never leave “uncounted / recompute unverified”**; never invent from line×avg  
3. Design **Est. sum** from scene `Est. length` fields  
4. Scene spans from `---` / `* * *` (or equivalent) breaks  

If the manuscript is in context via load, those integers are authoritative. **Evaluator failure to copy them is not a manuscript defect** — it means the evaluation is **incomplete**. Do not mark Length ❌ for “could not obtain a count.” Copy tool metrics, then score.

**Prose Quality Floor (before persona critique — mandatory):** run [Prose Quality Floor](#prose-quality-floor) on **`manuscripts/{nnn}-{episode-slug}.md` only**. Design files are **comparison** sources — quote them only when labeled as design, never as manuscript Evidence. **Content** Floor ❌ blocks manuscript-eval **pass**, **`Manuscript quality` ❌**, **Release ready ⬜**, and stage ⑧ — default revise at **`06-generate.md` (A)**; use ④ only if Key Events themselves cannot be dramatized. **Evidence integrity** ❌ always blocks **Release ready ⬜** until cites/counts are repaired — see [Verdict dimension rules](#verdict-dimension-rules). Do not claim evaluation passed, propose release, or soften Floor ❌ into ⚠️.

**Soft-craft probe (before persona Defects — mandatory):** fill the Soft-craft probe table in the eval template (rows A–E). Then run Literary Craft + Target Reader Checks using probe results. Required personas must address their assigned probes — bare Defects `—` without probe attestation is incomplete.

**Required personas:** run **every** persona in the [Default persona sets](#default-persona-sets). Always: Target Reader, Genre Critic, **Reader-Editor**, Literary Critic. Also ep 001 / lore debut: **Character Critic** + Setting/Lore Expert. Fill critique **even when Floor fails** — but do not treat personas as passing while Floor has ❌. If Literary Awards Juror is **not** in the set, write `{Not required}` and do not fill Juror Checks with design-based ✅.

**Literary Awards Juror:** When in the set (or requested), fill checks from the **manuscript page**, not from design intent. Load `series.md` Episode List + `overview.md` for work-level context.

---

## Evaluation record structure

Write `evaluations/{nnn}-{episode-slug}-manuscript.md` (separate from design eval at Stage ⑤).

**Writing rules:**
- **Create:** write the evaluation file.
 - **Update:** prefer updating with text copied verbatim from the latest known version. Habitual pre-update re-loading is forbidden when the file is already in context. On update failure, retry with a shorter unique anchor — do not guess from memory.

- **Overwrite** only when intentionally replacing the entire file.
- **Never chat-only:** do not claim Stage ⑦ complete, pass verdicts, or propose release until this file exists with all required sections filled.
- **Evidence (hard) — L + short quote:** every cell that cites a manuscript location must use `Lnn: 「…」` (or `Lnn–Lmm: 「opening words…」`) where the quoted fragment **appears on that line / range in the loaded manuscript**. **Take `nn` from the load content prefixes** (`329: …` → `L329`) — do not re-count lines from an unnumbered paste. **Range-only Evidence** (`L12–L19` with a paraphrase of content that is not on those lines) is **invalid** → Evidence integrity ❌ (blocks Release ready). Quoting `episodes/` Summary or Key Events **as if it were the manuscript** is invalid. For duplicates/loops: **every distinct repeated block**, each with **N matching the file** and **every line** + quote. Wrong N, wrong L, or L that does not contain the claimed content = incomplete Floor / fidelity.
- **L ≤ file length:** any Evidence `Lnn` with nn greater than the manuscript line count is fabricated — fix before Verdict. Header “Manuscript line count” must equal the line count of the loaded file.
- **Character count fidelity (hard):** Prefer integers from the manuscript’s latest file metrics: character count (whole file) and line count. Copy those into the eval header and Floor Length Evidence. If the manuscript was only written this turn and not re-read, use the actual character count the same way — do **not** estimate from line count × average. **Independent recompute** must match that integer (**exact**; tolerance 0). If written count ≠ tool/file recompute, Length Floor → **❌** for Evidence integrity. Do not invent a round number. Do not switch to hangul-only / whitespace-stripped / “after H1 only” counts unless Evidence explicitly labels that method **and** Scale comparison uses the same method — default is whole-file `char_count`.
- **Count sanity (hard):** before Length ❌ for Scale, check plausibility against the same payload:
  - Header **line count** must equal line count from the loaded file (or the same split-`\n` count on the in-context body). Header N ≠ file N → Evidence integrity ❌ (fix header; do not invent L>N).
  - `{n}` must equal `char_count` when that field is present. If claimed `{n}` ≈ **2×** recomputed `char_count`, or invents a number far from the tool field, the count is wrong → Evidence integrity ❌ — **use the tool integers**, then score Length. Do **not** open High “Scale overshoot” Adjudication on a failed count.
  - Rough band: for Korean prose, `char_count` is often on the order of **~15–40 × line_count** (wide); a claim like “14k chars / 340 lines” while `char_count` is ~7k is a failed count, not an over-long manuscript.
- **Forbidden Length Evidence:** `recompute unverified`, `could not count`, `line count only`, inventing MS chars without `char_count` from load (when the manuscript was read this turn), or any claim that MS length is “probably over Scale” **without** an integer `MS chars: {n}` that matches the tool field. Those phrases make the eval **incomplete** — use tool metrics, then score. Do **not** open a High Adjudication finding whose Action is “count the manuscript” (that is evaluator work, not a Stage ⑥ manuscript fix).
- **Length Floor after a real count:**
  - ❌ if `{n}` **< 70% of Scale min** or **> 130% of Scale max**
  - Prefer `{n}` ≤ Scale max and `|pct| ≤ 20%` vs Est.
  - If Scale band passes and `|pct| ≤ 20%` → Length **✅** (optional denser prose = Adjudication **Med** / Skip — **not** Length ❌)
  - If Scale passes but `|pct| > 20%` → Length still **✅** on Scale threshold; open **Med** Adjudication for compress/expand — **not** High “Scale violation”
  - Guessing over-length from **line count alone** → forbidden
  - **False Length ❌** (claimed over Scale while true `{n}` is inside Scale) → treat as Evidence integrity fail of the **eval**, not Manuscript quality ❌; fix the eval numbers before Adjudication High compress
- **Est distance (mandatory in Length Evidence):**
  ```text
  MS chars: {n}
  design Est. sum: {e}
  (MS − Est) / Est = {pct}%
  Scale min–max: {min}–{max}; % of min; % of max
  ```
- **L cite verification (before writing any Lnn):** copy the quote from the line you number. If `「…」` is not a substring of that line → **wrong cite** — fix the L number or the quote before Verdict. Wrong cites → Evidence integrity ❌. Do not leave knowingly mismatched L as support for Floor ✅ rows.
- **Scene boundary Evidence:** when claiming “Scene N = Lx–Ly”, the range must sit inside that scene’s manuscript span (delimited by `---` / `* * *` or equivalent). Citing Scene 1 lines as Scene 2 coverage → Design Fidelity ❌. **Continuous prose without `---` is allowed** — do not Floor ❌ or High Adjudicate solely for missing scene-break markers; use event landmarks in Evidence; optional Med/Low Skip for platform preference only.
- **Uncatalogued proper nouns (hard pre-check):** before ❌, open **Appearing** `characters/{slug}.md` (all **states**, esp. `base` gear) + used location profiles + episode design Key Events. Named gear that appears in a cited appearance state is **catalogued** → ✅ with Evidence `Lnn: 「…」; profile {path} state {id}`. ❌ only after that check fails. Indexes-only “I don’t remember gear” is not Evidence.
- **Motif / clue vs filler:** Designed motif recurrence or **advancing** clue confirmation (new fact, new place, new cost) is **not** Sensory/catalog Floor ❌. Floor ❌ only for **non-advancing** filler stacks. Repeated confirmation of the same already-established fact → Craft ⚠️ / Adjudication Med (compress), not automatic Floor ❌.
- **Absence is not a pass:** no quoted dialogue in the manuscript → Dialogue naturalness / Dialogue voices distinct = **❌** (not ✅ “no dialogue”). “Not verified” cannot be ✅ and cannot downgrade a Floor ❌ to Adjudication Medium.
- **Self-consistency:** if Floor is ✅ for duplicates/loops/catalog/meta/paste but Consistency, Craft, or any other cell quotes the **same** phenomenon, fix Floor to ❌ before Verdict. Search the manuscript for author-meta / plan-language patterns before scoring those rows ✅.
- **Persona vs table:** if Target Reader (or any required persona) says the reader would stall, skim, or quit, first-page / drop-risk / `Target-reader readiness` cannot all be ✅. Readiness follows the **harsher** of table vs persona. Drop-risk ✅ with Evidence “none noticed” is invalid when risks exist on the page. Persona Defects that cite `Lnn` must use the same **L + quote** rule.
- **Soft-craft probe (hard — before empty Defects):** run [Soft-craft probe](#soft-craft-probe-mandatory) on the manuscript **before** writing persona Strengths/Defects. A required persona may write Defects `—` **only** after its probe rows are filled with `no hit` **or** L-cited hits. Empty Defects with no probe attestation = **incomplete eval** (do not mark Release ready ✅). Hits are Craft / Adjudication findings by default — **not** automatic Floor ❌ unless they also match Floor rows (design paste, author-meta narration).
- **Adjudication Status:** see [Adjudication Status](#adjudication-status-hard) — `Apply? = no` → **Skip** (not Pending). Any **Pending** blocks Release ready ✅. Do not Pending “recount characters” as a manuscript Apply item. **Deduplicate** rows: same finding → one row (cite multiple personas); do not paste identical Skip rows twice.
- **Eval artifact only:** do not paste SKILL § Revision, Revision Loop, Floor skill prose, or this workflow’s checklist dump into the evaluation file — fill the template tables only.
- **Release / Gate footer must match Verdict:** if `Release ready` is **⬜**, do **not** write Gate G7 / footer language such as “no required fixes”, “optional polish only”, or “ready for Stage ⑧ after approval.” Those phrases are allowed **only** when Release ready is **✅**. While Floor Length ❌ or High **Pending**, the footer must say revise at ⑥ (or waive) — never invite ⑧.
- **Required personas:** ep **001** (and character-driven episodes) always include **Character Critic** + Setting/Lore Expert.

```markdown
# Manuscript Evaluation: Episode {nnn} — {Title}

> Target: `manuscripts/{nnn}-{slug}.md`
> Manuscript line count: {N = line count}
> Manuscript character count: {character count — whole file} — must match Length Evidence recompute

## Prose Quality Floor (any ❌ → Release ready ⬜; **content** ❌ → Manuscript quality ❌ — do not record Floor ❌ as ⚠️)
{Evidence: `Lnn: 「quote from that line」` — never range-only paraphrase; never unlabeled design-file quotes}

| Check | Result | Evidence |
|-------|--------|----------|
| No duplicate prose blocks (same or near-same paragraph/sentence ≥2×) | ✅ / ❌ | {each repeated block: Lnn: 「…」; N×; every L} |
| No design paste (Summary / Key Events / plan language / author-meta as undramatized narration) | ✅ / ❌ | {manuscript Lnn: 「…」; optional: “matches design Lx” labeled separately} |
| Scene dramatization (action, dialogue, sensory moment — not beat checklist prose) | ✅ / ❌ | {Lnn: 「…」 for dramatized beats} |
| Episode length vs budget (`overview.md` Scale; character count, not lines) | ✅ / ❌ | {MS chars; Est sum; (MS−Est)/Est %; Scale min–max; % of min; % of max — integer required; “uncounted” forbidden. Scale 70–130% pass → ✅; written≠recompute → ❌} |
| No degeneration loop (same phrase block repeated 3+× in close proximity) | ✅ / ❌ | {Lnn: 「…」; N×; L…} |
| Sensory/catalog discipline (no **non-advancing** filler descriptor stacks) | ✅ / ❌ | {filler stacks only — designed motif / advancing clue ≠ ❌; Lnn: 「…」} |
| No author-meta / stage-direction prose | ✅ / ❌ | {locale-aware scan; hit → ❌ with Lnn: 「…」} |
| No installment meta in body (“in episode 001” / `001화에서` as time/place — H1 heading allowed) | ✅ / ❌ | {body only; hit → ❌ with Lnn: 「…」} |
| Evidence integrity (char count + L quotes) | ✅ / ❌ | {header chars = Floor chars = recompute; every cited L contains its quote; file line count = header} |

## Consistency Checks (Design Fidelity)
| Check | Result | Evidence |
|-------|--------|----------|
| Pre-generation sources loaded / used | ✅ / ⚠️ / ❌ | |
| Every Key Event scene present in order | | {Scene 1 = La–Lb; Scene 2 = Lc–Ld — ranges must match `---` / scene breaks; Lnn: 「…」 for Turn} |
| No extra plot scenes without Key Event | | |
| Situation → Beat → Turn realized | | {per scene: Lnn: 「…」} |
| POV / On stage / Location / When match | | |
| Dialogue intent + voices honored | | {scene with Dialogue intent `none` but quoted speech → ⚠️ or ❌ + Lnn: 「…」} |
| Sensory-emotional realized (no catalog dump) | | |
| Exposition Budget items on page | ✅ / ⚠️ / ❌ | {each Budget item from design: present at Lnn: 「…」 / missing → ❌ or ⚠️} |
| Seeds Plant/Hint only; Hold absent | | |
| Uncatalogued proper nouns (gear / places / factions) | ✅ / ❌ | {**before ❌:** open Appearing `characters/{slug}.md` states + locations + design; ✅ Evidence: `Lnn: 「…」; profile {path} state {id}` if listed; ❌ only after miss} |
| Motifs / POV inserts at designed placements | | |
| Prior hook / Out hook / closing image | | {closing: Lnn: 「…」} |
| Continuity states & threads not contradicted | | |
| World / series tone not broken | | |
| Design-Fidelity Gate | | {All pass? Content Floor ❌ / uncatalogued after profile check → not Pass; Evidence integrity alone → Release ⬜ but Gate may still note cites to fix} |

## Engagement Checks (Manuscript)
| Check | Result | Evidence |
|-------|--------|----------|
| Opening hook effective | ✅ / ⚠️ / ❌ | {Lnn: 「…」 opening line} |
| Opening question persists | | {question only in design, not on the page → ❌} |
| Personal stake present | | |
| Out hook / forward momentum | | {Lnn: 「…」} |
| Exposition budget respected | | {align with Consistency Exposition Budget row} |
| Seed discipline | | |
| Scene-first prose | | |
| Dialogue naturalness | | {no quoted speech in manuscript → ❌} |
| Scene transitions smooth | | {boundary Lnn near `---`} |

## Literary Craft Checks (Manuscript)
{Run [Soft-craft probe](#soft-craft-probe-mandatory) first. Evidence for the rows below must cite **manuscript** `Lnn: 「…」` — design intent alone is not a pass.}

| Check | Result | Evidence |
|-------|--------|----------|
| Info : tension balance | | |
| Sensory-emotional pairing | | |
| Prose rhythm varied | | |
| Dialogue voices distinct | | {no dialogue in manuscript → ❌, not ✅} |
| POV inner/outer gap | | |
| Motifs threaded | | |
| POV insert discipline | | |
| World through character | | |
| Reader-discovered meaning | ✅ / ⚠️ / ❌ | {Compare design **Hold / “What reader should conclude”** to MS. Narrator or on-page speech **states that conclusion** → ❌ (or ⚠️ if one short slip). Pass = meaning left in image/action; Evidence: closing Lnn: 「…」 **and** “no restatement of Hold「…」”} |
| Emotion not over-labeled | ✅ / ⚠️ / ❌ | {Feeling nouns / “felt X” dumps without body cue → ⚠️/❌; pass needs Lnn body cue} |
| Antagonist plausibility | | |
| Closing scene over statement | ✅ / ⚠️ / ❌ | {Designed closing **image** missing → ❌. Image present **plus** thematic summary / moral coda after it → ❌ (or ⚠️ Low if one sentence). Evidence: final image Lnn: 「…」; if coda exists, Lnn: 「…」} |

## Literary Awards Juror Checks (Manuscript)
{Only if Juror is in the active set. Otherwise: `{Not required — overview.md has no prestige/awards criterion.}` Do not score Juror ✅ from design intent.}

| Check | Result | Evidence |
|-------|--------|----------|
| Theme & vision embodied in prose | ✅ / ⚠️ / ❌ | |
| Originality visible on the page | | |
| Human insight lands beyond plot | | |
| Artistic integrity — no sentimentality, didacticism, or ornamental prose | | |
| Episode contributes to series-level coherence in sequence | | |
| Competition readiness | | {Praise / hesitate / reject — with revision priorities} |

## Target Reader Checks (Manuscript)
{Always. Probe first. Do not ✅ all rows while Soft-craft or persona Defects cite skim/lecture/theme-tell on the same page.}

| Check | Result | Evidence |
|-------|--------|----------|
| Would the locked reader keep reading past the first page? | ✅ / ⚠️ / ❌ | {Lnn: 「…」 manuscript opening — not design hook} |
| Emotional / curiosity payoff lands for that reader | ✅ / ⚠️ / ❌ | {Payoff must be on-page action/object — not a stated moral. Lnn: 「…」} |
| Voice and density feel native to the platform/audience | ✅ / ⚠️ / ❌ | {Middle procedural / testimony chains: cite densest L-span if skim risk; “dense but event-driven” needs Lnn proof of a turn every ~screen} |
| Out hook / next-episode pull is concrete | ✅ / ⚠️ / ❌ | {Physical or actionable cliff — not thematic restatement. design-memo language ≠ concrete hook; Lnn: 「…」} |
| Drop-risk moments (confusion, lecture, stall) identified | ✅ / ⚠️ / ❌ | {✅ **only if** risks are listed with Lnn: 「…」 **or** Soft-craft probe rows all `no hit` and Evidence says `probe: no hit (rows A–E)`. “none noticed” without probe = invalid. Risks on page 1 → first-page ❌} |
| Opening question / curiosity handoff | ✅ / ⚠️ / ❌ | {Design Opening Question answered too early with no replacement curiosity → ⚠️/❌. Fully abandoned → ⚠️. Evidence: question status + replacement stake/hook Lnn} |

## Soft-craft probe (mandatory)
{Fill this table **before** persona Defects. Same turn as Floor. Hits → persona Defects + Adjudication; Floor ❌ only if also Floor-class (paste/meta narration).}

| # | Probe | Result | Evidence |
|---|-------|--------|----------|
| A | **Thematic / message coda** — last ~20 lines restate design Hold or “what reader should conclude” in narrator or speech | hit / no hit | {hit → Lnn: 「…」; quote design Hold/conclusion being echoed} |
| B | **Hold spoken aloud / plan-language dialogue** — character announces “not confirming X yet,” lists unresolved threads, or narrates Seeds/Hold as briefing | hit / no hit | {hit → Lnn: 「…」} |
| C | **Closing image buried** — designed closing image present but followed by moral/summary sentence that explains it | hit / no hit | {hit → image Lnn + coda Lnn} |
| D | **Middle skim / procedural stall** — testimony or procedure repeats without a new evidence/power shift for >1 screen-equivalent | hit / no hit | {hit → L-span: 「…」; note missing turn} |
| E | **Opening Q dead-end** — Opening Question resolved or dropped with no stronger replacement curiosity before midpoint | hit / no hit | {hit → where answered/dropped + whether Out hook compensates} |

## Manuscript Critique (required personas)
For **each** persona in the active set (do not replace with a single prose paragraph):

- `#### {Persona name}`
- Stance: {1–2 lines}
- Strengths: {…}
- Defects: {finding → severity High/Med/Low → proposed fix — cite Lnn: 「…」 when pointing at the page. **Forbidden:** blank `—` while any Soft-craft probe assigned to this persona is `hit`. If all assigned probes are `no hit`, write `— (probe A/B/… no hit)` — not bare `—`.}
- Reader impact: {how this affects the locked Target Reader}

## Manuscript Adjudication
{Record Apply decisions. Default tie-break: Target Reader. Never silently drop a High finding. Deduplicate same root cause. Use this table — not a bullet list. Soft-craft A/B/C hits default Apply? **yes** unless Target Reader explicitly prefers keep — see Soft-craft → Adjudication defaults.}

| # | Finding (persona) | Severity | Conflict? | Apply? | Rationale (Target Reader first) | Action Taken | Status |
|---|-------------------|----------|-----------|--------|---------------------------------|--------------|--------|
| 1 | {…} | High/Med/Low | | yes/no | | (A) ⑥ … / — | Pending \| Skip \| Carry-⑧ |

## Manuscript Verdict
| Dimension | Result |
|-----------|--------|
| Prose Quality Floor | ✅ / ❌ |
| Design fidelity | ✅ / ❌ |
| Target-reader readiness | ✅ / ❌ |
| Manuscript quality | ✅ / ❌ |
| Next-episode readiness | ✅ / ❌ |
| Release ready | ✅ / ⬜ |

## Release
- **Released**: ✅ / ⬜
- **Date**:
```

### Adjudication Status (hard)

| Status | When | Release ready |
|--------|------|----------------|
| **Pending** | `Apply? = yes` — manuscript (or design/architecture) must change before release | **Blocks** Release ready ✅ |
| **Skip** | `Apply? = no` — rejected / not applying this turn | Does not block |
| **Carry-⑧** | `Apply? = yes` — constraint only for release notes / next-ep continuity (rare); prose need not change | Does **not** block Release ready (record for ⑧) |

**Rules:**
- `Apply? = no` → Status **must be Skip** — never Pending.
- Do not write `Applied` in the same eval that first found the defect.
- Banner “Adjudication Status: Pending” while Verdict Release ready ✅ is **forbidden**.
- Align with Stage ⑤: generation-only Holds from design eval are **Carry-⑥** at ⑥; leftover release notes here may be **Carry-⑧**.

<a id="verdict-dimension-rules"></a>

**Verdict rules (hard) — do not soften / do not over-collapse:**

| Dimension | Set ❌ / ⬜ when | Do **not** auto-❌ solely because |
|-----------|------------------|-----------------------------------|
| **Prose Quality Floor** | Any Floor row ❌ (incl. Evidence integrity) | — |
| **Manuscript quality** | Any **content** Floor ❌ (duplicates, paste, dramatization, **true** Length after count, loops, sensory **filler**, author-meta, installment meta in body) | Evidence integrity alone (wrong L / unrepaired cites); optional Motif/clue Med Skip |
| **Design fidelity** | Consistency / Design-Fidelity Gate fails (paste, missing Key Events, uncatalogued after profile check, Budget miss) | Evaluator L typos; Length “uncounted”; motif recurrence that advances |
| **Target-reader readiness** | Target Reader Checks ❌ **or** persona says stall/quit | Floor Evidence integrity alone; Length Med compress suggestion while Scale passes |
| **Next-episode readiness** | Continuity / Out hook / Carry blockers | Same as Target-reader |
| **Release ready** | Floor any ❌ **or** any Adjudication **Pending** **or** Evidence integrity ❌ | Soften to ✅ while cites/counts broken |

- **Evidence integrity ❌** → `Prose Quality Floor` **❌** + `Release ready` **⬜** until header chars / L quotes are repaired in the **eval** (same turn). If all **content** Floor rows would ✅ under correct cites and a real Scale-passing count → `Manuscript quality` may stay **✅**; still **do not** release until Evidence integrity is ✅.
- Never write `Manuscript quality` ⚠️ or `Design fidelity` ⚠️ while a **content** Floor row is ❌.
- First-page Target Reader check ❌ **or** Target Reader persona says stall/quit → `Target-reader readiness` **❌**. No parentheticals in Verdict cells.
- Verdict cells are **only** `✅` / `❌` / `⬜`.
- Any **Adjudication High** finding with Status **Pending** → `Release ready` **⬜** until fixed or user explicitly waives in dialogue (record waiver in Adjudication).
- `Release ready` **✅** only when Floor all ✅ (incl. Evidence integrity) **and** no **Pending** rows **and** no open High Pending findings.
- `Design fidelity` and `Manuscript quality` are **✅ or ❌ only**. Partial craft issues that are not Floor/High stay in Adjudication as Skip or optional Pending.
- Floor / design-paste / under-length / over-length / undramatized beats → Adjudication **Action Taken = (A) `06-generate.md`** when design Est. sum is within Scale. Use **(B) episode design** when Key Events are wrong **or** design Est. sum is **outside Scale min–max** — not merely because the manuscript pasted a sound design.
- **Length Floor ❌ is High** only after a **sanity-checked** real `{n}` fails Scale 70%/130% (or written≠recompute). Write MS chars + Est distance in Floor Evidence in the same eval. Optional denser-than-Est prose inside Scale → Med Skip/Pending, not High “over Scale without count”.
- **Eval footer / Gate G7** must not contradict Verdict (no “ready for ⑧” / “no required fixes” while Release ready ⬜).
- Exposition Budget miss / uncatalogued proper noun (**after** profile check) → Design fidelity ❌ or ⚠️ with Adjudication (A) or (D) as appropriate.

---

## Prose Quality Floor

**Purpose:** Catch **basic publishability** problems on the page before treating the eval as a pass — repetition, design paste, under-/over-length, degeneration loops, author-meta, installment meta in the body, **false Evidence**. Manuscript counterpart of Stage ⑤ **Schema gate**.

**When:** After Load + Design-Fidelity re-check, **before** marking personas or verdicts as passing.

**How to run:** Ensure **`manuscripts/{nnn}-{episode-slug}.md`** is in this turn’s context (load once if missing). Copy **the line count** and **the character count**. Use only L numbers **≤ line_count**. Write `char_count` into Floor Evidence **and** the header — **never** score Length ❌ as “uncounted”; never invent from line×avg. Confirm the header integers match the tool fields before scoring Length ✅. For Uncatalogued nouns, open Appearing character **profiles** (states) before ❌. When writing each `Lnn: 「…」`, confirm the quote sits on that line. Scan for quoted dialogue. Scan for author-meta / plan-language / **installment meta in the body** (see Stage ⑥ locale-aware table — “In episode 001 I didn’t see that scratch.” / `001화에서 나는 그 흠집을 보지 못했다.`). Do **not** fail H1 `# Episode {nnn}: {Title}`. Scan the **full** prose for **every** repeated block. Map `---` (or equivalent) to Scene 1…N spans before claiming scene coverage. Do not invent L numbers from chat memory or use `episodes/` as a stand-in. Compare against the episode design and `overview.md` **Scale**. **Do not** re-read the manuscript after writing the evaluation just to confirm.

### Evidence rules (hard)

- **Source:** Floor and page-quality Evidence quote the **manuscript** as `Lnn: 「…」`. A quote that appears only in `episodes/` Summary is **not** manuscript Evidence. To show design paste, quote the **manuscript line** first, then optionally `matches episodes/… Lx` as a second labeled clause.
- **L + quote (mandatory):** the fragment after `「」` must appear on the cited line(s). Paraphrase-only or range-only Evidence is **invalid** → **Evidence integrity ❌**.
- **L in range:** if the manuscript has N lines, `L(N+1)` and above are invalid — correct Evidence from the same manuscript payload before Verdict (no verify re-read).
- **Duplicates / loops:** list **every distinct repeated block**, each with N and every L + quote. Citing one block when others repeat = incomplete Floor.
- **Length + Est distance:** use **`char_count`** (whole file from load). Evidence must include:
  - `MS chars: {n}` (must equal header and recompute)
  - `design Est. sum: {e}`
  - `(MS − Est) / Est = {pct}%`
  - Scale min–max and % of min / % of max  
  Never use line count as length; never skip the count; never invent n; never write “recompute unverified.” Compare to `overview.md` Scale (default **4,000–8,000**). Prefer |pct| ≤ 20% and MS near Est.
- **Threshold:** ❌ if actual **< 70% of Scale min** or **> 130% of Scale max**. ❌ if written MS chars ≠ recomputed chars. A short file is ❌, not ⚠️. If Scale band passes and Est |pct| ≤ 20% → Length ✅. Scale pass + Est |pct| > 20% → Length ✅ + Med Adjudication. Confirmation-loop bloat **inside** Scale → Floor loops/sensory rows and/or Med compress — not Length ❌ for “over 8000 without counting.”
- **Under-length routing:** if design Est. sum was **≥ Scale min** but manuscript is under → Adjudication **(A) `06-generate.md`**. If design Est. sum was **< Scale min** → also flag design Length budget fail → **(B) `04-episode-design.md`** (then regenerate). Never “fix” under-length by repeating paragraphs.
- **Over-length routing:** if manuscript **> Scale max** (or well over Est.+20%) and design Est. sum was within Scale → Adjudication **(A) `06-generate.md`** compress (cut repeated confirmation / restated facts). If design Est. sum was **> Scale max** → also **(B) `04-episode-design.md`**. Never “fix” over-length by deleting plot beats that the design requires — compress prose density first.
- **Cross-cell paste/meta:** quoting plan-language, author-meta, or installment meta in the **body** anywhere in the eval → corresponding Floor rows must be **❌**, not ✅. Do **not** fail H1 `# Episode {nnn}: {Title}`.
- Floor rows are **✅ or ❌ only**. Do not record a Floor fail as ⚠️.

| Check | Pass | ❌ Fail |
|-------|------|---------|
| **No duplicate prose blocks** | Each paragraph advances the scene | Same or near-same paragraph/sentence **≥2×** — count all hits |
| **No design paste** | Beats become **shown** scenes | Episode Summary / Key Events pasted as narration; undramatized plan / decision tell |
| **Scene dramatization** | Action, dialogue, concrete sensory beats | Long stretches of abstract plan, tell-only exposition, or checklist narration |
| **Episode length vs budget** | ≥ 70% of Scale **min** and ≤ 130% of Scale **max** (no pad); prefer ≤ Scale max and near design Est.; Est distance recorded | Line-count as length; invented / mismatched char count; under Scale min; overshoot with confirmation-loop bloat; Est distance omitted |
| **No degeneration loop** | Rhythm varies | Same **phrase block** **≥3×** within one scene or adjacent paragraphs; repeated confirmation of the same fact |
| **Sensory/catalog discipline** | Details serve the moment; designed motif / advancing clue OK | **Non-advancing** filler stacks without new action |
| **No author-meta / stage-direction prose** | Narration stays in-world | Design-memo sentences as author notes; staging/camera jargon (`viewer-left`, etc.) |
| **No installment meta in body** | Body never cites this work’s own Part / Episode / Chapter / 화 numbers as time or place | “In episode 001 I didn’t see that scratch.” / `001화에서 나는 그 흠집을 보지 못했다.` / “in Part 1” / `지난 화에서`. H1 `# Episode {nnn}: {Title}` is **not** a fail |
| **Evidence integrity** | Header chars = Floor chars = recompute; header line count = file N; every `Lnn` contains its 「quote」 and L≤N; scene L ranges match breaks or event landmarks | Fabricated / skipped / **implausible** char count (`{n}` > file length or ≈2× true body); wrong L; L>N; range-only Evidence; “uncounted” Length Evidence |

**Severity:** Any Floor ❌ → `Release ready` **⬜** (and `Prose Quality Floor` ❌). **Content** Floor ❌ → `Manuscript quality` **❌**. **Evidence integrity alone** → repair cites/count in the eval; `Manuscript quality` may stay ✅ if content rows ✅ — still no stage ⑧ until Evidence integrity ✅. Still write persona critique (for revision). Do not write Design-Fidelity Gate as overall Pass when content Floor/fidelity failed; do not put `Manuscript quality` ⚠️.

**Evolving floor:** This checklist is **living documentation**. When a session exposes a new baseline failure mode, **add a generic row here** and in the evaluation template above — do not rely on harness special-cases, and do not encode one project’s plot as a permanent rule.

---

## Design-Fidelity extras (manuscript)

Re-run Design-Fidelity from `06-generate.md`, plus:

| Extra check | Fail when |
|-------------|-----------|
| **Exposition Budget on page** | Design lists Budget items (e.g. “one cost of the premise”) that never appear as dramatized sensation/action/dialogue on the page |
| **Uncatalogued proper nouns** | Named weapon, room, faction, technique in MS **after** checking Appearing `characters/{slug}.md` **states** (incl. base gear), used locations, episode design, world — miss → ❌. Gear listed on a profile state → **not** uncatalogued. Indexes-only memory ≠ check |
| **Dialogue intent none** | Scene Key Events say `Dialogue intent: none` but MS has attributable quoted speech → note ⚠️ (or ❌ if it invents a speaker/plot beat) |
| **Scene span map** | Eval assigns wrong L span to Scene N relative to manuscript `---` breaks |

---

## Revision Loop

After adjudication, user selects findings to apply. Follow SKILL.md § Revision:

- **A — prose craft only** → revise `manuscripts/` (stage ⑥), re-run Design-Fidelity + Prose Quality Floor
- **B — episode design** (beats, hooks, seeds, motifs, scenes) → update episode design (stage ④), then regenerate manuscript
- **C — series/Episode List** (wrong Summary, Hook to Next, arc) → `series.md` → episode design → manuscript
- **D — lore / cast / place** → architecture catalogs (incl. Multi-facet anchors / gear states) → open designs → manuscripts as needed
- **D′ — modify approved lore** → Consistency C cascade (impact + republication if released)
- **E — continuity error** → prefer fix unreleased design/manuscript; if released must change → user must approve republication path
- **F — Carry-⑥** → generation constraint only; no ④ edit required
- **G — Skip** → `Apply? = no`; do not leave as Pending
- **Carry-⑧** → release-note only; does not require ⑥ rewrite

**Structural fixes require design-first updates; prose-only polish is allowed only when Key Events, Seeds/Hold, Prior Design Alignment, and continuity remain true.**

Update the evaluation file with revision decisions, then present the revised artifact for approval (design first if design changed).
---

## Persona Reference

Critics **advise**. They do not auto-apply changes. When personas conflict, **Target Reader** (from `overview.md`) is the default tie-break unless overview explicitly prioritizes prestige, niche, or another locked criterion.

### Default persona sets

| Context | **Required** (always run) | **Also required when** |
|---------|---------------------------|-------------------------|
| Manuscript eval (any episode) | **Target Reader**, **Genre Critic**, **Reader-Editor**, **Literary Critic** | Character-driven episode → **Character Critic**; ep 001 or major lore debut → **Setting/Lore Expert** + **Character Critic**; prestige/awards in overview → **Literary Awards Juror**; after major rewrite → **Literary Critic** even mid-series |
| Ep 001 (manuscript) | Above + **Setting/Lore Expert** + **Character Critic** | — |

Optional on request: any persona below; series/arc-level critique → `evaluations/{scope}.md`.

### Persona catalog

| Persona | Focus | Typical questions | **Mandatory probes (must address)** |
|---------|-------|-------------------|-------------------------------------|
| **Target Reader** | Locked audience’s lived read — fun, clarity, desire to continue, platform fit | Would *I* keep going? Where do I skim, confuse, or quit? | Soft-craft **D**, **E**; first-page retention; Out hook concreteness (not theme lecture); drop-risk L cites |
| **Reader-Editor** | Engagement craft, exposition restraint, serialization hooks | Is info dumped? Is the hook earned? Is the episode sellable as a unit? | Soft-craft **B**, **D**; closing sellable as click-through object; no design-memo dialogue |
| **Genre Critic** | Genre contracts, tropes, promise/payoff | Does it deliver what this genre’s readers paid for — without empty cliché? | Promise/payoff on page; empty victory lap vs earned reversal |
| **Character Critic** | Arc, motivation, distinct voice, relationship pressure | Why does this person act *now*? Voices interchangeable? | Soft-craft **B** if speech states Hold/plan; profile-backed knowledge vs on-page evidence leap |
| **Literary Critic** | Prose rhythm, motif, sensory-emotional craft, dialogue music, **restraint** | Does the page *feel*? Motifs earned? Emotion shown not labeled? | Soft-craft **A**, **C**; Reader-discovered meaning / Closing-over-statement rows; no thematic coda after image |
| **Setting/Lore Expert** | World consistency, info-dumping, rule clarity, **Exposition Budget on page**, **uncatalogued nouns** | Is the world learned through character need? **Is every Budget item dramatized? Any named gear/place not in catalog?** | Budget on page; Hold lore not lectured |
| **Literary Awards Juror** | Theme, originality, insight, artistic integrity, work-level merit | Would this earn lasting literary respect — or only serial entertainment? | Soft-craft **A**, **C** when Juror is in set; didacticism / sentimentality |

**Literary Critic vs Literary Awards Juror:** Critic = craft on the page. Juror = whether the work *as literature* earns lasting merit at series/arc level.

**Literary Awards Juror** should cover: theme & vision earned; originality; human insight; artistic integrity; work-level coherence; competition readiness with concrete revision priorities.

**Target Reader** should cover: who they are (from overview); what they came for; first-page retention; confusion or lecture risks; whether the Out hook pulls *them* specifically — and must not leave Soft-craft **D/E** hits as Defects `—`.

**Literary Critic** should cover: motif recurrence, sensory-emotional pairing, **and** whether the closing **shows** meaning without restating design Hold / reader-conclusion text. Soft-craft **A/C** hits → Defect with L cite (severity usually Med/Low; High only if the whole ending is a sermon).

### Soft-craft → Adjudication defaults

| Probe hit | Typical severity | Apply? default | Notes |
|-----------|------------------|----------------|-------|
| A thematic coda | Med (Low if one sentence after strong image) | **yes** unless Target Reader prefers keep | Action (A) ⑥ — cut coda; end on image |
| B Hold/plan dialogue | Med | **yes** | Action (A) ⑥ — turn into action/refusal/glance |
| C image + explaining coda | Med | **yes** | Same as A |
| D middle skim | Low–Med | optional | Skip OK if every exchange shifts evidence control |
| E Opening Q dead-end | Low | optional | Skip OK if Out hook clearly replaces curiosity |

Do **not** Skip A/B/C with rationale “Floor already ✅” — Floor does not cover thematic closing or Hold-as-dialogue.

---

## Gate G7 (Manuscript Eval Approval)

After evaluation + adjudication (and any applied revisions): user reviews the evaluation, confirms revision items, and decides revise (→ ④ or ⑥) or proceed to `08-release.md`.

**Release readiness** is determined here — **Prose Quality Floor** pass (incl. **Evidence integrity**) + Soft-craft probe table filled + no Adjudication **Pending** + manuscript verdict **Release ready** ✅ before stage ⑧. Stage ⑦ is not complete until the evaluation file exists with Floor, Soft-craft probe, persona critique, Adjudication, and Verdict filled. Do not propose ⑧ while header char count / L Evidence would fail a recompute, or while Soft-craft probe is missing / persona Defects are bare `—` without probe attestation.

## Reference-model integrity (manuscript eval)

Also check SKILL.md § Reference Models:
- Character appearance/equipment states consistent; no silent gear swap; **named weapons/gear in MS ⊆ character states / design**
- Location used as set; lasting changes tracked; **Location facets ⊆ Multi-facet anchors** when design cites `+ facet`
- Continuing situations cite stagings; L/R and stations do not flip without Design update