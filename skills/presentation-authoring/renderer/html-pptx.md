# HTML → PPTX Renderer

**Status: Planned** — not yet implemented. Do not start this pipeline until this spec is marked **Supported** in [`README.md`](README.md).

## Intended Purpose

Render each `slides/slide-{nn}.md` page model as HTML (or a single HTML deck), then convert to PowerPoint. Useful when web layout/CSS is the authoring surface and PPTX is the delivery format.

## Intended Pipeline (draft)

| Step | Artifact | Role |
|------|----------|------|
| 1 | `slides/slide-{nn}.md` | PIM page model (input) |
| 2 | `html/slide-{nn}.html` | PSM per-slide HTML |
| 3 | `{title-slug}.pptx` | Compiled PPTX |

## Intended Output Layout (draft)

| Path | Role |
|------|------|
| `slides/slide-*.md` | PIM input (Phase A) |
| `html/slide-{nn}.html` | Per-slide HTML |
| `{title-slug}.pptx` | Converted deck |

## Open Design Questions

- HTML template system vs hand-authored per slide
- Conversion tool (python-pptx, LibreOffice, commercial API)
- Asset paths (fonts, images) and offline portability
- Mapping spatial zones from page models to CSS layout

When implementing, flesh out: prerequisites, per-slide procedure, compile step, export, prohibitions, and completion criteria — following the structure in [`marp.md`](marp.md).
