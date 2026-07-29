# Stage ⑤ — Generation (`manuscripts/{nnn}-{chapter-slug}.md`)

**Prerequisites:**
- Approved chapter design (stage ④) — `chapters/{nnn}-{chapter-slug}.md`
- Design passed **Prior-Design Consistency Gate** and **Manuscript Readiness** (`04-chapter-design.md`)

**Gate artifact:** `{project-root}/manuscripts/{nnn}-{chapter-slug}.md`

**Next stage:** `06-evaluation.md` (manuscript evaluation, after user approves manuscript)

**References:** `04-chapter-design.md`; architecture + continuity paths in the chapter design; cross-stage sync [`consistency.md`](consistency.md)

---

## Core Rule — Design Fidelity

The manuscript **executes** the approved chapter design and prior-approved architecture/continuity. It does **not** redesign the chapter, cast, world, or timeline.

| Source of truth | Role in generation |
|-----------------|-------------------|
| `chapters/{nnn}-{chapter-slug}.md` | Scene order, beats, craft, seeds/Hold, hooks |
| Architecture refs in that file | Character voice/drive, location feel, world laws |
| Continuity (ch 002+) | Prior state, prior hook, open threads |
| `overview.md` / `series.md` | Tone, voice, genre — do not contradict |

**Do not invent plot.** Every scene comes from Key Events in order. If prose would need a beat, character, place, or lore fact that is not in design/architecture/continuity — **stop** and revise design (⑦) or architecture (③) first.

**Key Events are briefs — write dialogue and prose fresh**, guided by craft fields (not by pasting design into the manuscript as prose).

Do not start chapter N+1 until chapter N is released (unless user requests design-ahead).

---

## Pre-Generation Load Checklist (mandatory — before drafting)

Read from disk. **Do not rely on chat memory alone.**

| # | Artifact | Confirm |
|---|----------|---------|
| 1 | `chapters/{nnn}-{chapter-slug}.md` | Full file: Purpose, Arc, Seeds/Hold, Literary Craft, Prior Design Alignment, every Episode → every Scene Key Event |
| 2 | Architecture References listed in the chapter file | Open each listed path (`overview`, `series`, Chapter List source, `characters*`, `locations*`, `world*`) |
| 3 | Continuity References (ch 002+) | `story-so-far`, prior summary, `unresolved-threads` |
| 4 | Character profiles for speaking / POV cast | Voice, drive, inner/outer gap — match Dialogue Voices in design |
| 5 | Location profiles for scene Locations | Atmosphere/layout — match scene headers |
| 6 | World rules touched by Plant/Hint seeds | Consistent with bible; Hold stays out of prose |

**Abort drafting if:**
- Prior-Design Consistency or Manuscript Readiness would fail on the design file
- A scene is missing required Key Event fields
- Architecture Reference paths are missing or outdated vs On-stage / Location lists

**Do not re-read prior manuscripts** — continuity files are authoritative for past plot.

---

## Procedure

### Generation Procedure

1. Complete **Pre-Generation Load Checklist**
2. Optionally skim Prior Design Alignment in the chapter file — constraints still bind prose
3. Draft **scene-by-scene** in Key Events order within each episode, then next episode
4. For each scene, honor: POV | Location | When | On stage | Situation → Beat → Turn | Sensory-emotional | Dialogue intent | Transition out | Est. length
5. After draft, run **Design-Fidelity Gate** (below) before presenting to the user — status line: *"Load ✅ · Fidelity ✅"*
6. Present manuscript only if the gate passes

### Prose Generation Rules

Respect the design's **Seeds**, **Hold**, and **Literary Craft** (chapter-level and per-episode):

1. **Write only Plant and Hint items** — never include Hold items
2. **Scene-first** — open on action/image/dialogue per opening scene design
3. **Opening Question** — first page leaves unanswered the question from chapter design
4. **Personal stake early** — within the first page, per design
5. **World through character** — honor each scene's Sensory-emotional pair; no catalog lists
6. **Dialogue ≠ exposition** — Dialogue intent + Dialogue Voices; POV inner/outer gap at least once where designed
7. **Trust the reader** — one mystery hint beats three explained
8. **Closing on image** — chapter Closing image from design; no thematic summary
9. **Motifs** — only at designed scene placements
10. **POV inserts** — only at planned placements; ≤ 2 per episode
11. **Reader discovers meaning** — never state theme/moral already marked Hold
12. **Emotion indeterminate** — where design calls for it
13. **Antagonists believe they're right** — per design / profile
14. **Omission over catalog** — intense scenes: one focal detail + POV looks away if designed
15. **Transitions** — honor Transition out and episode Hooks; Est. length as pacing guidance
16. **Fidelity** — cast, places, facts, and outcomes stay inside architecture + continuity + this chapter's Key Events

### Prose Rhythm

Vary sentence length and structure. Avoid monotonous patterns.

| ❌ Avoid | ✅ Prefer |
|---------|----------|
| `explanation — explanation` em-dash chains repeating | Short punch sentences amid longer ones |
| Same-length declarative sentences | Sensory fragments: "Machine smell. Monitor glow. Yujin was used to it." |
| Explaining everything | Omit; let reader infer |

**Rule:** If 3+ consecutive sentences use the same structure (especially em-dash apposition), rewrite for rhythm.

### Manuscript Format

```markdown
# Chapter {nnn}: {Title}

## Episode 001: {Title}

{Prose text}

---

## Episode 002: {Title}

{Prose text}

---
...
```

Episode headings are optional separators for author reference — adjust to platform style if user prefers seamless prose without headings.

---

## Design-Fidelity Gate (mandatory — before approval)

Any ❌ blocks presenting the manuscript and blocks stage ⑥ until fixed (revise prose or return to ⑦ design).

| Check | Pass criteria |
|-------|----------------|
| Scene coverage | Every Key Event scene is present in prose order; no extra plot scenes that lack a Key Event |
| Beat fidelity | Each scene realizes Situation → Beat → Turn (paraphrase OK; opposite/changed outcome ❌) |
| POV / stage | POV and On-stage cast match scene headers; no surprise named cast from outside architecture |
| Place / time | Location and When match scene headers and location profiles |
| Dialogue | Speech serves Dialogue intent; voices match Dialogue Voices / character profiles |
| Sensory-emotional | Designed pair appears (detail → POV reaction), not replaced by info-dump |
| Seeds / Hold | Only Plant/Hint; Hold absent from prose |
| Motifs / inserts | Only at designed placements; POV insert budget respected |
| Hooks | Prior hook addressed (002+); Opening/Closing hooks and Closing image match design |
| Continuity | No contradiction of `story-so-far` / prior summary; threads not casually resolved if Held |
| World / tone | No new laws, history beats, or tone breaks vs world-bible / series voice |
| No silent redesign | If prose “needs” a plot change → stop; update design first |

---

## Pre-Submission Checkpoint

Complete Design-Fidelity Gate, then:

- [ ] Pre-Generation Load Checklist done
- [ ] Design-Fidelity Gate all pass
- [ ] Chapter design file honored (chapter + episodes + scenes)
- [ ] No invented beats or entities
- [ ] Architecture + continuity honored
- [ ] **Opening Question** unanswered past first page
- [ ] **Motifs** at designed scenes only
- [ ] **POV inserts** ≤ 2 per episode, planned only
- [ ] Sensory-emotional cues realized — no catalog blocks
- [ ] Dialogue intent + voices; POV inner/outer gap where designed
- [ ] Prose rhythm varied — no em-dash explanation chains
- [ ] Transitions honor Transition out / episode Hooks
- [ ] Personal stake in first page
- [ ] **No thematic closing monologue** — Closing image/silence per design
- [ ] Antagonist self-justifies where applicable
- [ ] POV emotions indeterminate where designed

---

## Gate

User approves the chapter manuscript **only if Design-Fidelity Gate passes**. Do not proceed to manuscript evaluation without approval.
