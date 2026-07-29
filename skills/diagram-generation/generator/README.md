# Diagram Generators (Stage 3 — How / PSM)

Each generator takes the approved **YAML (what)** and produces a **backend-specific artifact (how)**.

| Layer | Owns | Spec |
|-------|------|------|
| **YAML** | `title` → `summary` → `explanation` → `diagram.elements` / `diagram.relationships` (+ `parentId`) | [`SKILL.md`](../SKILL.md) |
| **Generator** | Layout design + tool implementation + output validation | `generator/{name}.md` |

**YAML never includes layout.** The same YAML may yield different layouts per generator when that tool’s idioms differ.

**Prerequisite:** Stages 1–2 complete — approved message + `diagram` graph.

## Generator Registry

| Generator | Spec | Input | Output | Status |
|-----------|------|-------|--------|--------|
| **Excalidraw** (default) | [`excalidraw.md`](excalidraw.md) | `{slug}.yaml` | `{slug}.excalidraw` | **Supported** (forward + reverse; layout via MCP `cocrates-elk`) |
| **draw.io** | [`drawio.md`](drawio.md) | `{slug}.yaml` | `{slug}.drawio` | **Supported** (forward + reverse; layout via MCP `cocrates-elk`) |
| **Mermaid.js** | [`mermaid.md`](mermaid.md) | `{slug}.yaml` | `{slug}.mmd` / `.md` | **Supported** (flowchart; forward + reverse) |

## Selecting a Generator

1. **Default:** Excalidraw unless the user requests another **Supported** generator.
2. Optional YAML: `backend: excalidraw | drawio | mermaid`
3. Read **only** that generator’s spec.

## Shared Rules (All Generators)

- **Input:** approved YAML only — do not invent elements, relationships, or story beats
- **Two internal steps:** (1) **layout** for this tool → (2) **implement** the output file. Excalidraw and draw.io: layout from MCP `cocrates-elk` (`layout_graph`), then map to the backend format.
- **Honor `parentId`:** nest / group / subgraph — **never** a containment-only arrow
- **Every relationship →** a visible link; honor `direction` (`directed` / `undirected` / `bi-directed`)
- **Visual encoding:** distinguish element `kind` and relationship `type` by shape/color/stroke (see generator tables); validate that distinctions are present
- **Deliver `explanation`:** reading order and visual clarity matter
- **Output validation:** run the generator’s validation checklist (Parts 1–5 in that generator’s file), present pass/fail to the user, fix failures before treating the job as done
- **Do not write layout back into YAML**
- **Colocate** output with YAML unless `output` / user overrides
- **Revision:** meaning/graph → Stages 1–2; layout/visual/tool bugs → regenerate in this generator
- **Reverse:** when analyzing an existing file, see `SKILL.md` Analyze & recover and that generator’s reverse-reading section

## Adding a New Generator

1. Add `{name}.md` with: when to use, layout approach for this tool, implementation schema, output validation, checklist, prohibitions.
2. Register above as **Planned** or **Supported**.
3. Do **not** add layout fields to `SKILL.md` YAML.
