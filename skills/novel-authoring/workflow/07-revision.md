# Stage ⑦ — Revision

**Prerequisites:** Evaluation completed (stage ⑥); user selected findings to apply

**Gate:** User approves revised artifact(s)

**Next stage:** `06-evaluation.md` (re-evaluate if needed), `05-generation.md` (after design revision), or `08-release.md`

**Chapter design rules:** `04-chapter-design.md`

---

## Core Rule

**Critique-driven revision updates the design before the manuscript** — not just patches prose.

For design evaluation findings: revise chapter file and episode files only.
For manuscript evaluation findings: update design first if structure/pacing/seeds affected, then rewrite manuscript.

---

## Revision Order

When evaluation suggests cutting, relocating, or rebalancing content:

1. **Update chapter design first:**
   - `chapters/{nnn}-{chapter-slug}.md` — chapter arc, Seeds, Hold, Literary Craft, Episode Index
   - `chapters/{nnn}-{chapter-slug}/{nnn}-{episode-slug}.md` — Key Events, episode-level craft
   - Move over-explained items from Plant → Hold
   - Adjust Exposition Budget and Seeds tables
   - Revise Key Events to scene-level beats
   - Add Personal Stake if missing
   - Update **Literary Craft**: Opening Question, Motifs, Sensory-Emotional Beats, Dialogue Voices, POV Inserts
2. **Then rewrite `manuscripts/{nnn}-{chapter-slug}.md`** following revised design (if manuscript exists)
3. Record changes in `evaluations/{nnn}-{chapter-slug}.md` → appropriate Revision Decisions section
4. Present revised design for approval (**design first**); manuscript second if both changed

---

## Standard Procedure

1. User selects findings from evaluation (design and/or manuscript section)
2. If structure/pacing/seeds affected → **revise chapter + episode designs first**
3. Update affected files
4. Record in evaluation Revision Decisions
5. Present for user approval

---

## Common Critique → Design Fixes

| Critique finding | Design fix | Manuscript fix |
|------------------|------------|----------------|
| Too much world-building in opening | Move lore to Hold; budget → Low | Scene opening; cut essay |
| Info-dump dialogue | Key Events: character questions | Compress dialogue |
| Too many mysteries explained | Reclassify as Hold; one Hint | Cut explanatory conversation |
| No personal connection | Add Personal Stake to chapter file | Add character-specific detail |
| Thematic summary ending | Chapter Hooks: closing image | Replace summary paragraph |
| Covers entire series themes | Distribute seeds in part/chapter list | Cut held content |
| Opening has action but no tension | Add **Opening Question** to chapter file | Rewrite opening with persistent question |
| Em-dash monotony / flat rhythm | Note rhythm target in Literary Craft | Vary sentence length; cut chains |
| World details are functional only | Add **Sensory-Emotional Beats** table | Pair each detail with POV reaction |
| Dialogue sounds same for all characters | Add **Dialogue Voices** table | Differentiate speech; add inner/outer gap |
| Motif appears once | Plan 2–3 placements across episodes | Thread motif across chapter |
| Too many POV inserts (logs) | Reduce POV Inserts to 1–2 per episode | Cut excess inserts |
| Info 70% / tension 30% | Lower exposition budget; add tension beats | Cut catalog blocks |
| Thematic closing monologue | Add to **Hold**; specify Closing image | Cut monologue; end on silence/image |
| Obvious villain dialogue | Dialogue Voices: self-justifying worldview | Rewrite antagonist as believer |
| POV emotion over-explained | Emotional indeterminacy in design | "angry or wronged?" — don't label |
| Violent scene over-described | One focal action in Key Events | Cut catalog; POV omission |
| Episode transitions jarring | Revise episode hooks in chapter file | Smooth transitions in prose |
| Chapter arc incoherent | Revise Chapter Arc and Episode Index | Restructure manuscript sections |

---

## Gate

User approves revised artifact(s). If design changed, design approval precedes manuscript approval.
