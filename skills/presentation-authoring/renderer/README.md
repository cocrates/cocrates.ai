# Presentation Renderers (Phase B — PSM)

Phase B is **renderer-specific presentation authoring** (PSM). It **translates** the approved **renderer-independent model** (PIM) from Phase A into a concrete platform.

| Concept | Phase A (PIM) | Phase B (PSM) |
|---------|---------------|---------------|
| Role | Define, architecture, detail design, generate page models | Translate page models into renderer artifacts |
| Primary input | — | `slides/slide-*.md` + `templates/` |
| Must not | Embed renderer syntax | Redesign narrative or on-slide copy |

Phase A (main [`SKILL.md`](../SKILL.md)) ends at `slides/` + Q&A + validation. Rendering is optional and separate.

**Prerequisite for every renderer:** Phase A complete — approved `slides/slide-*.md`, templates in use, optional `scripts/`.

## Renderer Registry

| Renderer | Spec file | PIM → PSM flow | Output | Status |
|----------|-----------|----------------|--------|--------|
| **Marp** (default) | [`marp.md`](marp.md) | `templates/` → `marps/template-*.md`; then `template` + `slides/` + `scripts/` → `marps/marp-*.md` → `{title-slug}.md` | HTML / PDF / PPTX | **Supported** |
| **Image → PPTX** | [`image-pptx.md`](image-pptx.md) | `slides/` → `images/` → `{title-slug}.pptx` | PPTX | **Planned** |
| **HTML → PPTX** | [`html-pptx.md`](html-pptx.md) | `slides/` → `html/` → `{title-slug}.pptx` | PPTX | **Planned** |

## Selecting a Renderer

1. **Default:** Marp — unless the user explicitly requests another supported renderer.
2. Record the choice in `overview.md` under **Output Renderer** when non-default or when the user cares about format:

   ```markdown
   ## Output Renderer
   marp | image-pptx | html-pptx
   ```

3. Read **only** the spec file for the chosen renderer. Do not mix pipelines (e.g., Marp `marps/` + image export).

## Shared Rules (All Renderers)

- **Input source (PIM):** `slides/slide-{nn}.md` + assigned `templates/template-*.md` + optional `scripts/script-{nn}.md`
- **Translation only:** implement spatial zones, copy, and template intent — do not invent new messages
- **Scripts:** optional `scripts/script-{nn}.md` — spoken narrative, not on-slide copy; Marp presenter notes must be the exact `## Script` section body
- **Marp:** translate templates to `marps/template-*.md` first (merge-safe CSS), then each `marps/marp-*.md` from template + slide + `## Script`; compile templates once, then slides in order
- **Revision:** layout/copy → Phase A `slides/` first; shared visual foundation → `templates/` then re-translate PSM artifacts
- **Independence:** Q&A and validation (Phase A) do not require any renderer output

## Adding a New Renderer

1. Add `{name}.md` in this directory with: when to use, prerequisites, output layout, PIM→PSM pipeline, per-slide artifacts, compile/export, prohibitions, completion criteria.
2. Register it in the table above with status **Planned** or **Supported**.
3. Do **not** change Phase A stages in `SKILL.md`.
