---
name: diagram-generation
description: >-
  Select when the user asks to draw, diagram, visualize, map, sketch, or explain
  a system, architecture, flow, or structure as a diagram — or to analyze,
  explain, reverse-engineer, or recover a design YAML from an existing diagram
  file (e.g. .excalidraw, .drawio, .mmd). Produces explanatory diagrams
  (Excalidraw default; draw.io and Mermaid flowchart also supported). Do not
  select for photoreal or illustrative AI images, video, or speech.
metadata:
  agent: cocrates
---

# Diagram Generation

Produce diagrams that **make an approved explanation script deliverable** — not decorative drawings.

| Layer | Owns | Must not |
|-------|------|----------|
| **YAML (what)** | `title`, optional `summary`, `explanation`, `diagram.elements` / `relationships` | Layout geometry, tool-specific formats |
| **Generator (how)** | [`generator/`](generator/README.md) — layout + file emit | Invent meaning, elements, or relationships |

`explanation` is a **pointing / interpretation script** (what you say while walking the picture). Image/video use `message` (story/beats to convey) instead — do not rename diagram fields to match.

Semantic design is **`diagram.elements` + `diagram.relationships`** (there is no separate `design` key). `summary` is optional; a complete YAML needs `title` + `explanation` + `diagram`.

## Core Principles

- **Explanation-first** — the picture must make `explanation` easy to speak while pointing at it.
- **Chat then YAML** — gather requirement and graph design in chat; write YAML when enough is known. If the request already has enough, write YAML immediately.
- **What before how** — YAML defines *what*; generators define *how* (including layout).
- **YAML gate before generate** — user reviews/approves the YAML (especially `diagram`), then generate. Express only on explicit user request.
- **Consistency** — keep present `title` / `summary` / `explanation` aligned; the graph must cover `explanation`.
- **Containment is structural** — use `parentId`, never a containment-only arrow.

Chat and human-facing YAML fields (`title`, `summary`, **`explanation`**, `label`, `notes`) use the **user's language** — never default them to English when the user writes in another language. Machine ids are English kebab-case.

## Workflow

```
1 Discover (chat)   requirement + element/relationship design
2 Write YAML        title / [summary] / explanation / diagram   → user reviews
3 Generate          layout + implement (generator/)             → backend file
4 Revise            fix owning layer; keep fields consistent
```

| Phase | Artifact | User reviews |
|-------|----------|--------------|
| Discover | Chat | Explanation beats; proposed graph summary when non-trivial |
| Write YAML | `{slug}.yaml` | Full YAML — **`diagram` is primary** |
| Generate | `.excalidraw` / `.drawio` / `.mmd` | Rendered picture |

**Enough already:** Skip lengthy Discover; write complete YAML; stop for approval (unless Express).

**Express Mode** (e.g. *draw it all*, *generate now*, *express*): write complete YAML and generate; user revises from the render. Never auto-enable from graph size.

## Paths

Default: `diagrams/{slug}.yaml` and sibling output (e.g. `{slug}.excalidraw`). Optional: `backend: excalidraw|drawio|mermaid`, `output`, `legend`.

---

## Phase 1 — Discover (chat)

Do not write YAML until Discover is sufficient (or already was). Work **requirement before graph**: do not invent a box inventory until `explanation` is coherent.

### 1.1 Requirement

| Field | Role | Required |
|-------|------|----------|
| `title` | Short diagram name | Yes |
| `summary` | One-sentence claim the viewer could verify by looking | No |
| `explanation` | Viewing-order walkthrough — what you would say while pointing at the picture | Yes |

**`explanation`:** Ordered beats the eye can follow, written in the **user's language**. Name concepts that will become elements and edges. Do not merely restate title/summary.

**Consistency:** Present fields must stay one story. Without `summary`, align `title` ↔ `explanation`. On edits, realign all present fields before designing the graph.

**Fit test:** Could someone speak `explanation` while pointing at a future diagram and feel complete?

Confirm requirement in chat when ambiguous, then proceed to graph design.

### 1.2 Graph design (maps to `diagram`)

Map approved `explanation` to elements and relationships — **semantic only** (no coordinates/sizes/layout).

In chat when the graph is non-trivial, present: proposed element list (id, label, kind, parent), relationship list (from/to/type/direction/label), and a short containment tree.

**Coverage:** every explanation beat → element, relation, or explicit out-of-scope note; no orphan elements; implied nesting uses `parentId`.

**Direction:** `directed` (default), `undirected`, or `bi-directed`. Prefer one `bi-directed` over two opposite `directed` rows for the same meaning.

Independent stories → prefer separate YAML files. When coherent → Phase 2.

---

## Phase 2 — Write YAML

Write the **complete** design YAML. No layout coordinates.

```yaml
# Human fields (title / summary / explanation / label) = user's language
# Machine ids = English kebab-case
title: 인증 요청 경로
backend: excalidraw

# summary: |   # optional
#   클라이언트의 HTTPS 호출이 API 게이트웨이를 거쳐 인증 서비스에 도달한다.

explanation: |
  클라이언트에서 시작한다: HTTPS 요청을 API 게이트웨이로 보낸다.
  게이트웨이는 인증 서비스로 토큰을 검증한다.
  "토큰 검증" 화살표를 따라가며 설명한다.

output: "./auth-overview.excalidraw"

diagram:
  elements:
    - id: client
      label: 클라이언트
      kind: actor
    - id: vpc
      label: VPC
      kind: boundary
    - id: api-gateway
      label: API 게이트웨이
      kind: system
      parentId: vpc
    - id: auth-service
      label: 인증 서비스
      kind: system
      parentId: vpc

  relationships:
    - id: rel-https
      from: client
      to: api-gateway
      type: flows-to
      direction: directed
      label: HTTPS
    - id: rel-validates-token
      from: api-gateway
      to: auth-service
      type: uses
      direction: directed
      label: 토큰 검증
```

### Elements

| Field | Meaning |
|-------|---------|
| `id` | English kebab-case |
| `label` | On-diagram text — user's language |
| `kind` | Optional: `actor`, `system`, `data`, `process`, `group`, `boundary`, `frame`, … |
| `parentId` | Required when contained in another element |
| `notes` | Optional rationale — user's language |

### Relationships

| Field | Meaning |
|-------|---------|
| `id` | English kebab-case |
| `from` / `to` | Element ids |
| `type` | e.g. `is-a`, `uses`, `calls`, `flows-to`, `precedes`, `implements`, `associates` |
| `direction` | `directed` (default), `undirected`, `bi-directed` |
| `label` | Optional — user's language |

Prefer `parentId` over `contains` / `has-a` edges. Prefer one `bi-directed` over two opposite `directed` rows for the same meaning.

### Design validation (before asking approval)

| # | Check | Pass means |
|---|--------|------------|
| 1 | Explanation coverage | Every beat maps to element, relation, or explicit out-of-scope |
| 2 | No orphans | Every element serves the script |
| 3–4 | Containment | Nesting via `parentId`; valid, acyclic |
| 5 | Script fidelity | Graph read aloud delivers `explanation` |
| 6 | Relation honesty | Edges and `direction` match meaning |
| 7 | Summary alignment | If `summary` exists, graph supports it |
| 8 | Id discipline | English kebab-case ids |

### YAML gate

1. Point to the path; show validation report (+ containment tree).
2. User reviews YAML — **`diagram` is primary**.
3. **Stop** until approval (unless Express), then Generate.

---

## Phase 3 — Generate

1. Resolve `backend` (default: `excalidraw`).
2. Read [`generator/README.md`](generator/README.md) and the matching generator spec.
3. Layout + implement (Excalidraw / draw.io: MCP `cocrates-elk` `layout_graph` per generator spec).
4. Output validation; present; revise if needed.

| Backend | Spec |
|---------|------|
| Excalidraw | [`generator/excalidraw.md`](generator/excalidraw.md) |
| draw.io | [`generator/drawio.md`](generator/drawio.md) |
| Mermaid.js | [`generator/mermaid.md`](generator/mermaid.md) (flowchart) |

Do **not** write layout into YAML.

---

## Phase 4 — Revise

| Change | Action |
|--------|--------|
| Requirement fields | Realign title ↔ [summary] ↔ explanation → update `diagram` → re-approve → regenerate |
| `diagram` | Edit YAML graph → re-validate → re-approve → regenerate |
| Layout / tool output | Fix in generator; do not add layout to YAML |

Never fix layout by deleting explanatory relationships.

---

## Analyze & recover (reverse)

| Mode | Intent | Output |
|------|--------|--------|
| **A — Explain** | Understand a picture | title / [summary] / explanation (chat; optional save) |
| **B — Recover YAML** | Spec from file | Full YAML with `diagram` — what only, no layout |
| **C — Validate-by-explain** | Fidelity vs approved `explanation` | Gap report |

Procedure: read file → infer graph → write explanation → deliver by mode. On **B**, validate and **stop for approval** before regenerate. Do not invent nodes unsupported by the picture; ignore legend swatches as story elements.

---

## Prohibitions

- Generating without a complete approved YAML (unless Express)
- Layout geometry in YAML; inventing elements in a generator
- Containment as arrow only; non-English or non-kebab machine ids
- Defaulting `explanation` / `title` / `summary` / `label` to English when the user writes in another language
- Auto Express; hiding validation failures
- On reverse: inventing graph nodes; copying layout into recovered YAML
