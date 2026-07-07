# Stage ⑤ — Generation (`manuscripts/{nnn}-{chapter-slug}.md`)

**Prerequisites:** Approved chapter design (stage ④) — chapter file + all episode files in `chapters/{nnn}-{chapter-slug}/`

**Gate artifact:** `novels/{title-slug}/manuscripts/{nnn}-{chapter-slug}.md`

**Next stage:** `06-evaluation.md` (manuscript evaluation, after user approves manuscript)

**Engagement reference:** `04-chapter-design.md` — Scene-first, Seeds, Literary Craft

---

## Procedure

Write **chapter prose** strictly following the approved chapter design and all nested episode designs. The manuscript is a **single file** combining all episodes in order — this is what gets published.

**Key Events define flow and function — you write dialogue and prose fresh at this stage**, guided by `Dialogue Voices`, `Literary Craft`, and scene intent (not by copying pre-written quotes from design).

Do not start chapter N+1 until chapter N is released (unless user requests design-ahead).

### Prose Generation Rules

Respect the design's **Seeds**, **Hold**, and **Literary Craft** sections (chapter + episode files):

1. **Write only Plant and Hint items** — never include Hold items
2. **Scene-first** — open on action/image/dialogue (see `04-chapter-design.md`)
3. **Opening Question** — first page must leave the reader with an unanswered question from chapter design
4. **Personal stake early** — within the first page, reader feels why *this character* cares
5. **World through character** — every world detail paired with POV reaction; no catalog lists
6. **Dialogue ≠ exposition** — distinct voice per character; POV inner/outer gap at least once
7. **Trust the reader** — one mystery hint beats three explained
8. **Closing on image** — concrete moment, not thematic summary (chapter closing)
9. **Motifs** — place each planned motif at designed Key Events across episodes
10. **POV inserts** — only at planned placements; ≤ 2 per episode
11. **Reader discovers meaning** — never state theme/moral verdict narrator already showed
12. **Emotion indeterminate** — label feelings with possibilities, not certainties ("angry or wronged?")
13. **Closing = image** — no thematic monologue; fragment ("That was——") at most
14. **Antagonists believe they're right** — self-justification, not sneering villain
15. **Omission over catalog** — intense scenes: one focal detail + POV looks away
16. **Episode transitions** — smooth flow between episodes within chapter; respect episode hooks

### Prose Rhythm

Vary sentence length and structure. Avoid monotonous patterns.

| ❌ Avoid | ✅ Prefer |
|---------|----------|
| `explanation — explanation` em-dash chains repeating | Short punch sentences amid longer ones |
| Same-length declarative sentences | Sensory fragments: "Machine smell. Monitor glow. Yujin was used to it." |
| Explaining everything | Omit; let reader infer |

**Rule:** If 3+ consecutive sentences use the same structure (especially em-dash apposition), rewrite for rhythm.

### Generation Procedure

1. Load `chapters/{nnn}-{chapter-slug}.md` — chapter-level Seeds, Hold, Literary Craft, Engagement
2. Load all episode files in `chapters/{nnn}-{chapter-slug}/` in order
3. Load all architecture artifacts listed in Architecture References
4. **Chapter 002+:** Load continuity files per Continuity References
5. If adjacent chapters were designed ahead, read their designs for hook alignment
6. **Do not re-read prior manuscripts** — continuity files are authoritative
7. Write single chapter manuscript. Honor all designs. No improvised entities.

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

## Pre-Submission Checkpoint

- [ ] Chapter file + all episode designs honored?
- [ ] Architecture References honored — no improvised entities?
- [ ] Continuity References honored (ch 002+)?
- [ ] **Opening Question** — unanswered past first page?
- [ ] **Motifs** placed at designed points across episodes?
- [ ] **POV inserts** ≤ 2 per episode, at planned placements only?
- [ ] World details via POV reaction — no catalog blocks?
- [ ] Dialogue voices distinct; POV inner/outer gap present?
- [ ] Prose rhythm varied — no em-dash explanation chains?
- [ ] POV consistent with design?
- [ ] Events match Key Events in all episode files?
- [ ] Only Plant/Hint seeds — Hold excluded?
- [ ] Personal stake in first page?
- [ ] Character voice matches profile?
- [ ] Prior hook addressed (ch 002+)?
- [ ] Episode transitions smooth within chapter?
- [ ] **No thematic closing monologue** — reader concludes from scene?
- [ ] Chapter closing ends on **image/silence** per design?
- [ ] Antagonist self-justifies — not cartoon villain?
- [ ] POV emotions indeterminate where designed?

---

## Gate

User approves the chapter manuscript. Do not proceed to manuscript evaluation without approval.
