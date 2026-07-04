# Image → PPTX Renderer

**Status: Planned** — not yet implemented. Do not start this pipeline until this spec is marked **Supported** in [`README.md`](README.md).

## Intended Purpose

Render each `slides/slide-{nn}.md` page model to a per-slide image, then assemble images into a PowerPoint file. Useful when pixel-perfect visual control or non-Marp tooling is required.

## Intended Pipeline (draft)

| Step | Artifact | Role |
|------|----------|------|
| 1 | `slides/slide-{nn}.md` | PIM page model (input) |
| 2 | `images/image-{nn}.png` (or `.svg`) | PSM per-slide image |
| 3 | `{title-slug}.pptx` | Assembled PPTX |

## Intended Output Layout (draft)

| Path | Role |
|------|------|
| `slides/slide-*.md` | PIM input (Phase A) |
| `images/image-{nn}.png` | Per-slide render |
| `{title-slug}.pptx` | Assembled deck |

## Open Design Questions

- Image generation tool (browser screenshot, design API, manual export workflow)
- Resolution and aspect ratio (16:9 default)
- Whether presenter notes live in PPTX speaker notes from `scripts/`
- Revision loop when a slide image must be regenerated

When implementing, flesh out: prerequisites, per-slide procedure, compile step, export, prohibitions, and completion criteria — following the structure in [`marp.md`](marp.md).
