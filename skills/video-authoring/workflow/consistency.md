# Consistency — Cross-Stage Protocol (Video Authoring)

**Not a numbered stage.** Read when starting Design / Evaluate / Generate Components / Assemble, when revising, or when an approved artifact changes.

---

## Glossary

| Term | Meaning |
|------|---------|
| **Design Layer** | `overview.md`, `references.md`, catalogs (incl. stagings), `context.md`, `sequence.md`, `segments/*.md` |
| **Reference model rules** | `workflow/reference-models.md` when visual catalogs are used |
| **Evaluation Layer** | `evaluations/*.md` — **design lock** per segment |
| **Component Layer** | `images/`, `videos/`, `speech/`, `music/` — Stage ④ assets |
| **Assembly Layer** | `assembly/**/*.yaml` (edit-specs) + `output/**/*.mp4` via MCP `cocrates-video-edit` |
| **Design Lock (G3)** | Clip messages, direction, track *needs* frozen for ④–⑤ |
| **Component Lock (G4)** | Segment’s media assets approved for assemble |
| **Short** | Single implicit segment; clip design inside `sequence.md` |
| **Long** | `sequence → segment → clip`; per-segment assemble then final |

---

## Source-of-truth stack

```text
overview.md
  → references.md (+ catalogs / context)
    → sequence.md
      → segments/*.md          # Long; Short: clips in sequence.md
        → evaluations/*.md     # G3
          → images / videos / speech / music   # G4
            → assembly/segments/*.yaml → segment mp4   # G5a
              → assembly/final.yaml → output/{slug}.mp4  # G5b Long
```

Markdown story/message wins over YAML. Composition YAML must not redefine story.

---

## Cascade protocol

### References / catalogs change

1. Approve Design change.
2. Find affected segments/clips.
3. New G3 for affected segments if messages/direction change.
4. Regenerate affected reference images + dependent clip images/videos.
5. Re-assemble affected segments (and final if Long).

### Clip message / direction change

1. Stage ② edit → new G3.
2. Regenerate affected components (④).
3. Update edit-spec YAML and re-render via `cocrates-video-edit` (⑤).

### Component quality only

- Retry YAML/generate ≤2; do not change locked messages.
- Then re-assemble if binary outputs changed.

### Assembly timing / transition only

- Edit edit-spec YAML; re-approve; `validate_spec` → `render_video`.
- If “fix” actually changes story rhythm intent → Design + G3.

---

## Cascade checklist

```markdown
## Consistency Cascade
- Change: {what}
- Highest updated file(s): {paths} — approved: ✅/⬜
- [ ] overview.md
- [ ] references.md / catalogs / context.md
- [ ] sequence.md / segments/*
- [ ] evaluations/* (G3 rerun)
- [ ] images / videos / speech / music (G4)
- [ ] assembly/segments/* + output/segments/* (G5a)
- [ ] assembly/final.yaml + output/{slug}.mp4 (G5b, Long)
```

---

## Session resume

1. State stage + active segment id.
2. Reload disk artifacts (do not trust chat alone).
3. Confirm which gates exist (G3/G4/G5a) for that segment.
4. Continue the matching `workflow/0X-*.md`.
