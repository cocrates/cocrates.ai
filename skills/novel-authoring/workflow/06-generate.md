# Stage ⑥ — Generation (`manuscripts/{nnn}-{episode-slug}.md`)

**Prerequisites:**
- Approved episode design (stage ④) — `episodes/{nnn}-{episode-slug}.md`
- Design passed **Design Consistency Gate** and **Generation Readiness** (`04-episode-design.md`)
- [If design evaluation was run] Design evaluation approved (stage ⑤)

**Gate artifact:** `{project-root}/manuscripts/{nnn}-{episode-slug}.md`

**Next stage:** `07-evaluate-manuscript.md` (**main evaluation**, after user approves manuscript)

**See also:** [`04-episode-design.md`](04-episode-design.md); architecture + continuity in the episode design; SKILL.md § Consistency.

---

## Core rule — Design fidelity

The manuscript **executes** the approved episode design and prior-approved architecture/continuity. It does **not** redesign the episode, cast, world, or timeline.

| Source of truth | Role in generation |
|-----------------|-------------------|
| `episodes/{nnn}-{episode-slug}.md` | Scene order, beats, craft, seeds/Hold, hooks |
| Architecture refs in that file | Character voice/drive, location feel, world laws |
| Continuity (ep 002+) | Prior state, prior hook, open threads |
| `overview.md` / `series.md` | Tone, voice, genre — do not contradict |
| Design eval `Carry-⑥` rows (if ⑤ ran) | Generation-only constraints (POV limits, do-not-narrate Holds) — honor without redesign |

**Do not invent plot.** Every scene comes from Key Events in order. If prose would need a beat, character, place, or lore fact not in design/architecture/continuity — **stop** and revise design (④) or expand architecture (③) first.

**Key Events are briefs** — write dialogue and prose fresh, guided by craft fields (do not paste design into the manuscript as prose).

**Prose Quality Floor (generation — hard):** Before presenting the manuscript, the draft must pass the same publishability bar as stage ⑦. Failures block Gate G6 — fix prose (or return to ④ if beats cannot be dramatized). See [Prose Quality Floor (generation)](#prose-quality-floor-generation).

**Do not start episode N+1 manuscript until episode N is released (⑧ complete).** No batching or design-ahead exception for prose.

---

## Pre-Generation Load (mandatory)

Load sources needed for drafting via **Selective artifact load** (SKILL.md). **load only paths not already in this turn’s context.** Do not re-read unchanged files or the manuscript you are about to write “to be sure.”

### Phase A — Indexes (if missing from context)

`overview.md`, `series.md`, `characters.md`, `locations.md`, `world-bible.md`, `stagings.md`

### Phase B — Episode-scoped details only

| # | Artifact | Confirm |
|---|----------|---------|
| 1 | `episodes/{nnn}-{episode-slug}.md` | Full file: Summary, Purpose, Arc, Seeds/Hold, Literary Craft, Prior Design Alignment, every Scene Key Event |
| 2 | Architecture References listed in the episode file | Open **each listed path**, including **profile files** — not indexes alone; **do not** open unlisted catalog entries |
| 3 | Continuity References (ep 002+) | `story-so-far` (incl. Current Unresolved Threads) + **immediate prior** summary only — not every `continuity/*` |
| 4 | Character profiles for speaking / POV cast | Voice, drive, inner/outer gap (subset of Appearing if listed) |
| 5 | Location profiles for scene Locations | Atmosphere/layout |
| 6 | World rules touched by Plant/Hint seeds | Consistent with bible; Hold stays out of prose |
| 7 | Cited `stagings/{slug}.md` | When Staging ≠ none |
| 8 | `evaluations/{nnn}-{slug}-design.md` (if Stage ⑤ ran) | Honor every Adjudication row with Status **Carry-⑥**; ignore Skip; Pending must be cleared before ⑥ |

**Abort drafting if:**
- Design Consistency or Generation Readiness would fail on the design file
- A scene is missing required Key Event fields (including Paragraph outline / Unit budget)
- Architecture Reference paths are missing or outdated vs On-stage / Location lists
- Prose forecast / Est. cross-check would fail

**Do not re-read prior manuscripts** or prior `episodes/*.md` — continuity files are authoritative for past plot. **Do not** batch-read all of `characters/`, `locations/`, `stagings/`, `episodes/`, or `continuity/`.

---

## Procedure

1. Complete **Pre-Generation Load** (Phase A indexes → Phase B listed profiles, not catalog dump)
2. Note `overview.md` **Scale** (default **4,000–8,000**) and the design’s scene Est. length list + sum (must already satisfy Scale min ≤ sum ≤ Scale max from ④/⑤). Est. sum is the **episode manuscript target**.
3. Draft **scene-by-scene** in Key Events order — prefer one scene per write/edit turn for long episodes
4. For each scene, honor: POV | Location | When | On stage | Situation → Beat → Turn | Sensory-emotional | Dialogue intent | Transition out | **Paragraph outline** (write in that order) | **Unit budget** / **Est. length** (**land near Est.** — ±20%; do not pad; do not inflate past Est. with confirmation loops)
5. **Dramatize** Beats: action, dialogue, sensory moment — never paste Summary / Key Events / plan-language / “implies that…” as narration. Each bullet or sentence in the **Paragraph outline** becomes **one paragraph** of prose; target **220–280 characters** per paragraph (density band); a paragraph exceeding **400 characters** without a line break or dialogue shift is too long — split or compress. Do not inflate a single outline bullet into multiple paragraphs.
6. After full draft, run **Design-Fidelity Gate** + **Prose Quality Floor (generation)** on the file on disk — status: *"Load ✅ · Fidelity ✅ · Prose Floor ✅"*
7. Present manuscript only if **both** gates pass
8. If honest dramatization lands **well under** Scale min / scene Est. because Beats were thin → **stop**; revise **④**, do not fill with repetition
9. If the draft lands **well over** scene Est. (+20%) or **above Scale max** → **compress before presenting** (cut repeated confirmation questions, restated facts, filler sensory stacks). Do **not** invent new plot. If Beats themselves force overshoot because Est. was packed to Scale max → return to **④** and slim the forecast

### Prose generation rules

Respect the design’s **Seeds**, **Hold**, and **Literary Craft**:

1. Write only Plant and Hint items — never include Hold
2. Scene-first — open on action/image/dialogue per opening scene design
3. Opening Question — first page leaves unanswered the question from episode design
4. Personal stake early — within the first page, per design
5. World through character — honor each scene’s Sensory-emotional pair; no catalog lists
6. Dialogue ≠ exposition — Dialogue intent + profile Voice; POV inner/outer gap where designed
7. Trust the reader — one mystery hint beats three explained
8. Closing on image — from design; no thematic summary; no author-meta “implies…”
9. Motifs — only at designed placements; each once, varied wording
10. POV inserts — only at planned placements; ≤ 2 per episode
11. Reader discovers meaning — never state theme/moral already marked Hold
12. Emotion indeterminate — where design calls for it
13. Antagonists believe they’re right — per design / profile
14. Omission over catalog — intense scenes: one focal detail + POV looks away if designed; no filler descriptor stacks
15. Transitions — honor Transition out; Est. length as **target** pacing (not a minimum to exceed)
16. Fidelity — cast, places, facts, outcomes stay inside architecture + continuity + Key Events
17. No duplicate prose — never reuse the same paragraph/sentence block (≥2×). If a beat must recur, rewrite with new action/dialogue
18. No design paste — Key Events / Summary / plan language stay off the page
19. Length — **aim for each scene ≈ Est. (±20%)**; episode total should land **≥ 70% of Scale min** and **≤ Scale max** (Floor still allows ≤ 130% of max as hard fail, but do not present drafts that need that slack). No padding. No confirmation-loop bloat. Under-length after honest dramatization → revise ④. Over-length after honest dramatization → compress at ⑥ first; if design Est. sum was already at/above Scale max with no headroom, slim ④. If design Est. sum was < Scale min, that is a **design** failure — return to ④/⑤ before rewriting prose for length alone
20. Self-scan before present — search for author-meta / plan-language / **installment-meta** patterns, identical repeated paragraphs, and **restated facts / repeated confirmation questions**; remove or rewrite
21. Staging vocabulary stays off the page — no `viewer-left` / `viewer-right` / camera or blocking jargon in narration
22. **No installment meta in the body.** Narration and dialogue must not cite this work’s own Part / Episode / Chapter / 화 numbers as in-world time or place. The file H1 (`# Episode {nnn}: {Title}`) is allowed.

### Author-meta / plan-language scan (locale-aware)

Treat as fail if **body** narration or dialogue uses **author stage direction**, **undramatized plan**, or **installment meta** instead of in-world showing. Scan the **prose after H1** only — do **not** fail the heading `# Episode {nnn}: {Title}`. Scan for language-appropriate patterns, including:

| Pattern class | English examples | Korean examples (when writing Korean prose) |
|---------------|------------------|---------------------------------------------|
| Author-meta | “implies that…”, “conveys…”, “confirms that…” as stage notes | `암시한다`, `전달한다`, `확인한다` as stage notes |
| Plan paste | “decided to declare…”, “planned to…” as undramatized tell | `계획을 세웠다` as undramatized tell |
| Installment meta | “In episode 001 I didn’t see that scratch.”; “back in Part 1”; “as Scene 2 established” | `001화에서 나는 그 흠집을 보지 못했다.`; `에피소드 1에서`; `파트 1에서`; `지난 화에서` as a time marker |

**Pass:** in-world time and memory — “I didn’t see that scratch then.” / `그때는 그 흠집을 보지 못했다.`
**Fail:** treating the serial’s table of contents as a place the narrator or characters inhabit.
**Exception:** in-universe media the story actually contains (a character watching a numbered broadcast), only if that media is in design/architecture. This work’s own episode/part/화 numbering is never spoken in the body.

### Prose rhythm

Vary sentence length and structure.

| ❌ Avoid | ✅ Prefer |
|---------|----------|
| Repeated explanation — explanation em-dash chains | Short punch sentences amid longer ones |
| Same-length declarative sentences | Sensory fragments mixed with longer lines |
| Explaining everything | Omit; let the reader infer |
| Same paragraph pasted 2+ times | One clear beat; next beat advances |
| Author stage direction | Show the glance, the silence, the unfinished motion |

**Rule:** If 3+ consecutive sentences use the same structure (especially em-dash apposition), rewrite. If the **same multi-sentence block** appears ≥2×, rewrite before Gate G6.

### Manuscript format

```markdown
# Episode {nnn}: {Title}

{Prose text}
```

- The whole file is **one episode** — no episode subheadings inside.
- Scene breaks are optional separators for author reference: blank line, `---`, or `* * *` at designed Transition-out points. Adjust to platform style if the user prefers seamless prose.
- The H1 heading is file metadata. **Installment numbers must not appear in the body** (see rule 22).

---

## Design-Fidelity Gate (mandatory — before approval)

Any ❌ blocks presenting the manuscript and blocks stage ⑦ until fixed (revise prose or return to ④).

| Check | Pass criteria |
|-------|----------------|
| Scene coverage | Every Key Event scene present in prose order; no extra plot scenes without a Key Event |
| Beat fidelity | Each scene realizes Situation → Beat → Turn (paraphrase OK; opposite/changed outcome ❌) |
| POV / stage | POV and On-stage cast match scene headers; no surprise named cast from outside architecture |
| Place / time | Location and When match scene headers and location profiles; facets ⊆ Multi-facet anchors when design cites them |
| Dialogue | Speech serves Dialogue intent; voices match character profiles; if intent is `none`, avoid attributable quoted speech unless a brief POV slip is clearly non-plot |
| Sensory-emotional | Designed pair appears (detail → POV reaction), not replaced by info-dump |
| Exposition Budget | Each Budget item from the episode design appears as dramatized sensation/action/dialogue (not omitted; not dumped as lecture) |
| Seeds / Hold | Only Plant/Hint; Hold absent from prose |
| Motifs / inserts | Only at designed placements; POV insert budget respected |
| Hooks | Prior hook addressed (ep 002+); In/Out hooks and closing image match design |
| Continuity | No contradiction of `story-so-far` / prior summary; threads not casually resolved if Held |
| World / tone | No new laws, history beats, or tone breaks vs world-bible / series voice |
| No silent redesign | If prose “needs” a plot change → stop; update design first |
| No uncatalogued proper nouns | Named weapons, rooms, factions, techniques ⊆ design / character states / locations / world — else stop → catalog |

---

## Prose Quality Floor (generation)

**Purpose:** Stop unpublishable drafts **before** user approval / stage ⑦ — same failure modes Stage ⑦ Floor catches (repetition, paste, under-/over-length, meta, installment meta in the body, catalog stacks). Stage ⑦ defines the full evidence rules; this gate uses the same pass/fail criteria on the draft.

**When:** After Design-Fidelity Gate, **before** presenting the manuscript. Run the checks against the manuscript text already in this turn (the write/update payload you produced). **Do not** re-load the manuscript solely to run this floor.

| Check | Pass | ❌ Fail (do not present) |
|-------|------|---------------------------|
| **No duplicate prose blocks** | Each paragraph advances | Same or near-same paragraph/sentence **≥2×** |
| **No design paste** | Beats are dramatized | Summary / Key Events / plan-language / tell-only decision as narration |
| **Scene dramatization** | Action, dialogue, sensory beats | Checklist / plan / diary-summary prose |
| **Episode length vs budget** | ≥ 70% of Scale **min**, **≤ Scale max** preferred, and ≤ 130% of max hard; scene totals near Est. (±20%) **without** padding or confirmation-loop bloat; record MS chars and (MS−Est)/Est % | Padding; severe under-length; **overshoot past Scale max** (or past Est.+20% without compress); invented char count; design Est. sum was < Scale min or > Scale max (fix design first) |
| **No degeneration loop** | Rhythm varies | Same phrase block **≥3×** in close proximity; repeated confirmation of the same fact |
| **Sensory/catalog discipline** | Details serve the moment | Filler descriptor stacks without new action |
| **No author-meta** | In-world narration only | Author stage-direction patterns (see scan table above); staging/camera jargon (`viewer-left`, etc.) |
| **No installment meta in body** | Body never cites this work’s own Part / Episode / Chapter / 화 numbers as time or place | Body contains “In episode 001 I didn’t see that scratch.” / `001화에서 나는 그 흠집을 보지 못했다.` / “in Part 1” / `지난 화에서`. H1 `# Episode {nnn}: {Title}` is **not** a fail |

**Carry-⑥ closing / soft-craft (hard when ⑤ left Carry or Applied closing constraints):** Before presenting Gate G6, scan the draft ending for the same soft-craft risks Stage ⑦ probes:

| Probe | Fail if | Fix before present |
|-------|---------|-------------------|
| **C — Closing image buried** | Designed closing image appears, then procedural/moral/summary dialogue or narration dilutes it | Compress post-reveal lines; end on the designed image/gesture |
| **Crowded Out (≥3 signals)** | Last beats stack rules disclosure + slip + reveal + next-action declaration at equal weight | Keep one dominant image; one short operational line max |

If any Carry-⑥ row mentions closing image / red-box / seal cord / final visual, probe **C** is **mandatory** this turn.

---

## Pre-Submission Checkpoint

Complete Design-Fidelity Gate **and** Prose Quality Floor (generation), then:

- [ ] Pre-Generation Load done (profiles opened)
- [ ] Design-Fidelity Gate all pass
- [ ] Prose Quality Floor (generation) all pass
- [ ] Character count vs Scale recorded (load `char_count` or equivalent whole-file length); design Est. sum noted; **(MS−Est)/Est %** noted; totals near Est. (±20%); ≤ Scale max; no pad-to-length; no confirmation-loop bloat
- [ ] Exposition Budget items dramatized; no uncatalogued proper nouns (gear/place/faction)
- [ ] No author-meta / plan-language / staging jargon left as narration
- [ ] No installment meta in the **body** — no “in episode 001” / `001화에서` / “in Part 1” as time or place (H1 heading is allowed)
- [ ] No duplicate multi-sentence blocks
- [ ] Episode design honored (episode + scenes)
- [ ] No invented beats or entities
- [ ] Architecture + continuity honored
- [ ] Opening Question unanswered past first page
- [ ] Motifs at designed scenes only; wording not copy-pasted across scenes
- [ ] POV inserts ≤ 2 per episode, planned only
- [ ] Sensory-emotional cues realized — no catalog blocks
- [ ] Dialogue intent + voices; POV inner/outer gap where designed
- [ ] Prose rhythm varied
- [ ] Transitions honor Transition out
- [ ] Personal stake in first page
- [ ] No thematic closing monologue — closing image/silence per design
- [ ] **Carry-⑥ closing honored** — soft-craft probe C (and crowded-Out compress) passed when applicable
- [ ] Antagonist self-justifies where applicable
- [ ] POV emotions indeterminate where designed

---

## Gate G6 (Manuscript Approval)

User approves the episode manuscript **only if Design-Fidelity Gate and Prose Quality Floor (generation) both pass**. Do not proceed to manuscript evaluation (`07-evaluate-manuscript.md`) while either fails.
