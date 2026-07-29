# draw.io Generator (Stage 3 — How)

Turn approved YAML (**what**) into a diagrams.net / draw.io file (**how**), or **reverse-read** a `.drawio` / `.xml` file for analysis / YAML recovery.

- Forward & reverse workflow: [`SKILL.md`](../SKILL.md)
- Other generators: [`README.md`](README.md)
- **Layout engine:** MCP `cocrates-elk` (`validate_graph` → `layout_graph`) — then map geometry to draw.io

**Hard rules:** Do not invent story meaning, elements, or relationships. Do not write layout coordinates into YAML. Do not hand-place the story graph when the ELK MCP is available. Optimize emission for draw.io idioms (cells, parent-relative geometry, orthogonal edges).

**Status:** Supported.

## When to use

| Situation | Action |
|-----------|--------|
| Generate / export draw.io | Size → **ELK layout MCP** → map to draw.io → validate (this file) |
| Stage 2 YAML approved; user wants draw.io | Same (`backend: drawio`) |
| Analyze / explain / recover YAML from `.drawio` | **Reverse reading** + [`SKILL.md`](../SKILL.md) Analyze & recover |
| Another backend | [`README.md`](README.md) |

**Forward prerequisite:** Approved `title`, `summary`, `explanation`, `diagram.elements`, `diagram.relationships`.

## Input / output

| Direction | Input | Output |
|-----------|-------|--------|
| Forward | `{slug}.yaml` (+ optional `backend: drawio`, `output`, `legend`) | `{slug}.drawio` (same folder unless `output` overrides) |
| Reverse | `.drawio` or draw.io `.xml` | Explanation and/or Stage-2 YAML (what only) |

## Pipeline (forward)

```
1. Read approved YAML
2. Size leaves (text heuristics + peer unify + grid)
3. Build ELK graph from YAML → validate_graph → layout_graph (cocrates-elk MCP)
4. Map layout → draw.io vertices / edges (+ style + optional legend)
5. Post-process: convert absolute → parent-relative; snap to 20px grid; emit .drawio
6. Run Parts 1–5 validation → report pass/fail to user
7. On failure → revise XML (or re-layout with adjusted sizes/options) and re-validate
   (never write layout into YAML)
```

**Layout authority:** Node positions, container sizes, and edge routes come from **`layout_graph`** on MCP server **`cocrates-elk`** (configured in mcp.json as `cocrates-elk`; Cursor exposes `cocrates-elk`). Do **not** hand-place the story graph when this MCP is available.

---

## Constants

| Token | Value |
|-------|-------|
| **Layout MCP** | Server `cocrates-elk` — tools `validate_graph`, `layout_graph` (`list_algorithms` / `list_layout_options` as needed) |
| **Default algorithm** | `layered` |
| **Default edge routing** | `ORTHOGONAL` (`elk.edgeRouting`) |
| **Hierarchy** | `elk.hierarchyHandling: INCLUDE_CHILDREN` when any `parentId` exists |
| **Grid size** | **20px** — snap geometry after layout (parent-local and absolute tops) |
| Parent / child inset | ≥ **20px** (prefer **40px** via ELK padding) |
| Title band (container with label) | **≥60px** top padding in ELK (`[top=60,…]`) + `startSize=40` on swimlane containers — 40px band + ≥20px gap before first child |
| Peer gap | **40px** preferred → `elk.spacing.nodeNode: 40` |
| Layer gap | **60px** preferred → `elk.layered.spacing.nodeNodeBetweenLayers: 60` |
| Canvas outer margin | ≥ **40px** (add after layout if content hugs 0,0) |
| Text padding inside shape | ≥ **10–15px**; then snap leaf size up to grid **before** ELK |
| Standard leaf size | **Width 160** / **height 60** when label fits; grow in **20px** steps |
| Standard actor (ellipse) | **140×60** or **160×80**, snapped |
| Font sizes | Title **20** / label **16** / edge label **14** |
| Default edge routing (draw.io) | **Orthogonal** (`edgeStyle=orthogonalEdgeStyle`) |
| Default font | Helvetica / `fontFamily=Helvetica` |

---

## Layout via cocrates ELK MCP (required)

### MCP workflow

Before calling tools, discover schemas with **GetMcpTools** for server `cocrates-elk`.

| Step | Tool | Purpose |
|------|------|---------|
| 0 (optional) | `list_algorithms` / `list_layout_options` | Pick algorithm or tune options when default is a poor fit |
| 1 | `validate_graph` | Catch duplicate ids, bad parents/endpoints, cycles, missing leaf sizes |
| 2 | `layout_graph` | Obtain absolute `x,y,width,height` per node and edge `sections` |

If the server is `needsAuth`, call `mcp_auth` once, then retry. If layout MCP is unavailable or errors after retry, **stop and tell the user** — do not silently invent a full manual placement (offline fallback only if the user explicitly asks).

### YAML → ELK `graph`

Build a **flat** node list (prefer `parent` over nested `children` — do not mix both on one node).

| YAML | ELK graph field |
|------|-----------------|
| `diagram.elements[].id` | `nodes[].id` |
| `label` | `nodes[].label`; containers also `labels: [{ id, text, width, height }]` for title space |
| `parentId` | `nodes[].parent` |
| Leaf size (precomputed) | `nodes[].width` / `height` (**required for leaves**) |
| Container size | **Omit** `width`/`height` on parents so ELK sizes them |
| `diagram.relationships[]` | `edges[]` with `id`, `source`, `target` (skip `has-a` / `contains`) |
| Relationship `label` | Optional `edges[].labels: [{ text, width, height }]` |

**Do not** put draw.io `style=` strings into the ELK graph — only structure + sizes + labels for layout.

### Pre-ELK node sizing (required — leaves must have width/height)

ELK lays out boxes of **known size**. Leaves without `width`/`height` get weak defaults (often ~80×40) and labels clip. **Compute leaf sizes here before `validate_graph` / `layout_graph`.** Containers omit size so ELK grows them.

#### Character width weights (`fontSize` 16 for leaf labels)

| Character class | Weight |
|-----------------|--------|
| ASCII letters / digits / punctuation | 0.55 |
| CJK (Hangul, Han, kana) / wide glyphs | 1.0 |
| Space | 0.35 |

```
lineWidth  = fontSize * sum(weights in that line)     # fontSize = 16 for leaves
textWidth  = max(lineWidth over lines) + 12
textHeight = fontSize * 1.35 * lineCount + 8
```

#### Soft wrap before measuring

| Rule | Value |
|------|-------|
| Prefer single line | If `textWidth + 24 ≤ 240` |
| Soft-wrap | If longer, insert line breaks at word/phrase boundaries so each line’s `lineWidth` stays ≤ **200** (encode as `&#xa;` in cell `value` later) |
| Never | Mid-word / mid-syllable breaks |

Use the wrapped label for **both** ELK sizing and the eventual cell `value`.

#### Leaf box formula

```
pad = 24                          # ≥10–15px per side
minW = textWidth + pad
minH = textHeight + pad

# Kind floors (before snap)
floorW, floorH =
  actor            → 140, 60   (if still tight after text: 160, 80)
  system/process/data/default → 160, 60
  custom leaf      → 160, 60

w0 = max(minW, floorW)
h0 = max(minH, floorH)

width  = ceil(w0 / 20) * 20
height = ceil(h0 / 20) * 20
```

#### Peer unify (same parent + same visual role)

1. Compute `(width, height)` per leaf as above.  
2. `W = max(widths)`, `H = max(heights)` in the peer set.  
3. Assign every peer `width=W`, `height=H` **before** calling ELK.  
4. Optional hero: one emphasized node may be +20 on one axis only if the script requires it.

#### Worked examples

| Label | Approx text | Leaf size sent to ELK |
|-------|-------------|------------------------|
| `API` (system) | ~40×30 | **160×60** (floor) |
| `API Gateway` (system) | ~110×30 | **160×60** |
| `인증 서비스` (system, CJK) | ~16*6+12 ≈ 108 → +24 | **160×60** |
| `Message Broker Service` | wrap → 2 lines or grow | e.g. **200×80** or **180×80** after snap |
| `Client` (actor) | short | **140×60** |

#### Containers (parents)

| Field | Rule |
|-------|------|
| `width` / `height` | **Omit** — ELK sizes from children + padding |
| `labels[]` | Include one label so title space is reserved: `text` = YAML `label`, `height` = **20** (or 24), `width` = `ceil((titleTextWidth + 8) / 20) * 20` with **fontSize 20** weights |
| Padding | Rely on `elk.padding` top **40** (title band / `startSize=40`); do not fake a huge empty leaf inside the container |

#### Edge labels (optional on ELK edges)

If the relationship has a `label`, you may pass `edges[].labels: [{ text, width, height }]` with **fontSize 14**:

```
width  = ceil((textWidth + 8) / 20) * 20
height = ceil((textHeight + 4) / 20) * 20   # often 20
```

If omitted, still set the edge cell `value` after layout.

#### Checklist before `validate_graph`

- [ ] Every **leaf** has positive `width` and `height` (grid-snapped).  
- [ ] No leaf uses the tool default ~80×40 “because sizing was skipped”.  
- [ ] Peer sets share identical W×H.  
- [ ] Parents have **no** fixed width/height.  
- [ ] Container `labels[]` present when the parent has a visible title.  
- [ ] Wrapped label text matches what will appear in cell `value`.

If `validate_graph` warns about missing leaf sizes, fix sizes and re-validate — do not call `layout_graph` yet.

### Default `options` for `layout_graph`

```json
{
  "algorithm": "layered",
  "direction": "RIGHT",
  "absoluteCoordinates": true,
  "layoutOptions": {
    "elk.hierarchyHandling": "INCLUDE_CHILDREN",
    "elk.edgeRouting": "ORTHOGONAL",
    "elk.spacing.nodeNode": 40,
    "elk.layered.spacing.nodeNodeBetweenLayers": 60,
    "elk.padding": "[top=60,left=20,bottom=20,right=20]",
    "elk.layered.nodePlacement.bk.fixedAlignment": "BALANCED"
  }
}
```

| Cue from `explanation` | `options.direction` |
|------------------------|---------------------|
| Left→right / request flow | `RIGHT` |
| Top→bottom / hierarchy stack | `DOWN` |
| Unclear | `RIGHT` if mostly flow edges; else `DOWN` |

| Graph shape | Prefer |
|-------------|--------|
| Flow / pipeline / layered architecture (default) | `layered` |
| Pure tree from one root | `mrtree` |
| Dense undirected association mesh | `force` or `stress` (rare) |
| Unconnected boxes only | `box` / `rectpacking` |

When any compound nesting exists, keep **`layered` + `INCLUDE_CHILDREN`** unless the user asks otherwise.

### ELK layout → draw.io

Assume `absoluteCoordinates: true` (default). ELK returns **absolute** page coordinates; draw.io nested cells need **parent-relative** geometry.

**Nodes → vertices**

1. Map each laid-out node id → one `mxCell` vertex (kind → `style`).  
2. Snap absolute `x,y,width,height` to multiples of 20.  
3. If content hugs the origin, translate so outer margin ≥40 (then re-snap).  
4. **Convert to parent-relative:**  
   - Top-level (`parent="1"`): `mxGeometry` = snapped absolute `x,y,width,height`.  
   - Nested (`parent="{containerId}"`):  
     `x' = child.x - parent.x`, `y' = child.y - parent.y` (after the same snap/translate).  
5. Emit containers before children; set child’s `parent` to the container cell id.  
6. `value` = YAML `label` (XML-escaped). Container styling by role — see **Vertex label alignment** below.  
7. Apply kind palette via `style=` — styling is **not** from ELK.

**Edges → connectors**

Use **`layout_graph` edge output** — do not invent routes when `sections` is non-empty.

| ELK field | Use |
|-----------|-----|
| `edges[].id` | Match YAML relationship `id` |
| `edges[].sources[0]` / `targets[0]` | Edge `source` / `target` cell ids |
| `edges[].sections[]` | Orthogonal waypoints — **required** when present |

**Section → waypoint algorithm**

For each edge, walk `sections` in order. For each section:

```
points_abs += [section.startPoint]
points_abs += section.bendPoints ?? []
points_abs += [section.endPoint]
```

De-duplicate consecutive identical coordinates. Then:

1. Snap intermediate points to the 20px grid.  
2. Put **intermediate** bends (not first/last) into `<Array as="points"><mxPoint x="…" y="…"/></Array>` in **absolute page coordinates** (draw.io convention).  
3. Keep `edgeStyle=orthogonalEdgeStyle`; `source` / `target` bind to vertex ids.  
4. Stroke + arrow markers from relationship `type` / `direction`; edge `value` = relationship `label` when present.  
5. After vertex snap, if endpoints drift, let draw.io re-route orthogonally from `source`/`target` — waypoints should still reflect ELK bend structure when multiple bends exist.

**When `sections` is empty**

| Cause | Action |
|-------|--------|
| Missing `elk.hierarchyHandling: INCLUDE_CHILDREN` on nested graphs | Fix options; **re-call** `layout_graph` |
| Cross-hierarchy edge without compound support | Fix hierarchy options or graph structure |
| Otherwise | Document in validation; omit `<Array as="points">` and rely on auto orthogonal routing |

**Do not** ignore non-empty `sections` and draw a conflicting hand layout.

**Legend:** Hand-composed after story layout (not sent to ELK). Place on grid to the right or bottom of the story bbox.

### Post-layout polish (still required)

| Rule | Detail |
|------|--------|
| **Snap to grid** | All vertex geometry (relative and absolute tops) and edge waypoints → multiples of 20 |
| **Inset check** | Children fully inside parents with ≥20px inset in absolute space after snap |
| **Title band** | Children below parent top + title/`startSize`; verify first child `rel y ≥ 60` for ≥20px gap below 40px band |
| **Uniform peers** | Leaf sizes unified **before** ELK; do not randomly resize after except snap |
| **No manual reshape** of ELK layering — fix sizes/options and **re-call** `layout_graph` |

### Composition (semantic rules unchanged)

- **Reading order:** Match `explanation` via ELK `direction` + edge flow.  
- **Containment:** `parentId` → ELK `parent` → draw.io cell `parent` + relative geometry. **No** containment edge.  
- **Kinds / types / labels / legend:** As in Visual distinction below.  
- **Edges:** One per relationship; prefer ELK orthogonal sections → waypoints.

### Peer box test (validation)

Skip if A is ancestor of B via `parentId` / cell `parent` chains:

```
peer overlap_area = 0  (absolute / page coordinates)
peer gap ≥ 20 (prefer 40)
all geometry x,y,width,height % 20 == 0 (parent-local and absolute tops)
child fully inside parent with ≥20px inset; below title band when parent labeled
layout positions originated from layout_graph (not freehand invent)
```

---

## Document format

Emit a single-page **uncompressed** `.drawio` (diagrams.net compatible):

```xml
<mxfile host="app.diagrams.net" modified="2026-01-01T00:00:00.000Z" agent="diagram-generation" version="22.0.0">
  <diagram id="diagram_1" name="{title}">
    <mxGraphModel dx="1200" dy="800" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="1169" pageHeight="827" math="0" shadow="0">
      <root>
        <mxCell id="0"/>
        <mxCell id="1" parent="0"/>
        <!-- vertices and edges with parent="1" or a container id -->
      </root>
    </mxGraphModel>
  </diagram>
</mxfile>
```

| Cell | Role |
|------|------|
| `id="0"` | Root |
| `id="1"` | Default layer |
| Vertex | `vertex="1"` + `mxGeometry` + `style` + `value` |
| Edge | `edge="1"` + `source` + `target` + `style` + optional `value` (label) |
| Nested child | `parent="{containerId}"`; geometry **relative to parent** |

**Ids:** Stable strings derived from YAML element/relationship ids when possible (`api-gateway`, `rel-https`). Legend cells: `legend-*` prefix. All ids unique.

Escape XML in `value`: `&` → `&amp;`, `<` → `&lt;`, `>` → `&gt;`, `"` → `&quot;`. Newline in labels: `&#xa;`.

---

## Visual distinction

Same YAML `kind` / `type` → same look. Built-in rows are defaults, not a closed enum.

### Element `kind` → shape & style

| `kind` | draw.io shape cues | `fillColor` | `strokeColor` | Dashed | Notes |
|--------|--------------------|-------------|---------------|--------|--------|
| `actor` | `ellipse=1` | `#e0f2fe` | `#0284c7` | no | External role |
| `system` | rounded rect | `#dbeafe` | `#1d4ed8` | no | `rounded=1` |
| `process` | rounded rect | `#f3f4f6` | `#374151` | no | Step / activity |
| `data` | rounded rect | `#fef3c7` | `#d97706` | no | Store / payload |
| `boundary` | container / swimlane | `#fafafa` | `#9ca3af` | **yes** | `dashed=1`; holds children |
| `group` | container | `#f8fafc` | `#94a3b8` | **yes** | Logical group |
| `frame` | large container | `#ffffff` | `#9ca3af` | yes | Zone; title in `value` |
| *(missing kind)* | rect | `#ffffff` | `#1e1e1e` | no | Prefer explicit `kind` |

Example vertex style (system):

```
rounded=1;whiteSpace=wrap;html=1;fillColor=#dbeafe;strokeColor=#1d4ed8;strokeWidth=2;fontSize=16;fontColor=#1e1e1e;align=center;verticalAlign=middle;
```

Example container **with children** (swimlane — title in top band via `startSize`):

```
swimlane;startSize=40;fillColor=#fafafa;strokeColor=#9ca3af;strokeWidth=2;dashed=1;fontSize=20;fontColor=#1e1e1e;align=center;verticalAlign=middle;
```

Example container **without children** (plain box — label centered in full height):

```
rounded=1;whiteSpace=wrap;html=1;fillColor=#fafafa;strokeColor=#9ca3af;strokeWidth=2;dashed=1;fontSize=20;fontColor=#1e1e1e;align=center;verticalAlign=middle;
```

(`startSize` ≈ title band height — use **only** when the element has YAML children.)

### Vertex label alignment (required)

draw.io uses `verticalAlign=middle` in both cases, but **shape type** differs:

| Element role | Has YAML children? | Style approach |
|--------------|-------------------|----------------|
| Leaf | — | Normal vertex; `verticalAlign=middle`; `fontSize=16` |
| Container **with children** | **Yes** | **`swimlane;startSize=40`** — title sits in top band; children below with ≥20px gap. ELK top padding must be ≥60 (40 band + 20 gap) so children start at `rel y=60`, not `y=40`. |
| Container **without children** | **No** | **No swimlane** — dashed/rounded container rect; `verticalAlign=middle` centers label in full box |

Detect “has children” from YAML `diagram.elements[].parentId`, not from ELK output alone. This matches Excalidraw’s `verticalAlign: top` vs `middle` rule — different mechanism, same intent.

**Custom / unknown `kind`**

1. Prefer mapping to the closest built-in when the meaning matches (e.g. `service` → `system`, `db` → `data`, `zone` → `boundary`).  
2. Otherwise treat as **custom**: pick one stable style for that slug and reuse it for every element with the same `kind` in this diagram.  
3. Choose a palette that still differs from co-existing kinds (≥2 of shape / fill / stroke / dashed). Suggested custom pool (cycle if many customs):

| Slot | Shape cues | `fillColor` | `strokeColor` | Dashed |
|------|------------|-------------|---------------|--------|
| C1 | rounded rect | `#fce7f3` | `#db2777` | no |
| C2 | rounded rect | `#ede9fe` | `#7c3aed` | no |
| C3 | rounded rect | `#ccfbf1` | `#0d9488` | no |
| C4 | rounded rect | `#ffedd5` | `#ea580c` | no |
| C5 | `ellipse=1` | `#e0e7ff` | `#4f46e5` | no |

4. If the custom kind is container-like (holds children via `parentId`), use **dashed** stroke like `boundary`/`group`.  
5. Note the mapping in chat (and legend, if shown).  
6. Prefer setting an explicit `kind` in YAML over leaving it empty.

### Relationship `type` → edge style

Plus `direction` arrow markers.

| `type` | `strokeColor` | Dashed | `strokeWidth` |
|--------|---------------|--------|---------------|
| `flows-to` | `#2563eb` | no | 2 |
| `requests` | `#7c3aed` | no | 2 |
| `uses` / `calls` | `#0f766e` | no | 2 |
| `precedes` / `follows` | `#4b5563` | **yes** | 2 |
| `is-a` | `#b45309` | no | 2 |
| `implements` | `#be185d` | **yes** | 2 |
| `associates` | `#6b7280` | no | **1** |
| `has-a` / `contains` | — | — | **Do not draw** — nest via `parent` |

Example edge style (`flows-to`, directed):

```
edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;strokeColor=#2563eb;strokeWidth=2;endArrow=block;endFill=1;startArrow=none;fontSize=14;fontColor=#1e1e1e;
```

| `direction` | Edge arrows |
|-------------|-------------|
| `directed` | `endArrow=block;endFill=1;startArrow=none` |
| `undirected` | `endArrow=none;startArrow=none` |
| `bi-directed` | `endArrow=block;endFill=1;startArrow=block;startFill=1` |

**Custom / unknown `type`**

1. Prefer mapping to the closest built-in when the meaning matches (e.g. `publishes` → `flows-to`, `depends-on` → `uses`, `then` → `precedes`).  
2. Otherwise treat as **custom**: assign one stable stroke for that slug and reuse it for every relationship with the same `type` in this diagram.  
3. Keep it visually distinct from co-existing types (color and/or dash and/or width). Suggested custom pool:

| Slot | `strokeColor` | Dashed | `strokeWidth` |
|------|---------------|--------|---------------|
| R1 | `#0891b2` | no | 2 |
| R2 | `#c026d3` | no | 2 |
| R3 | `#65a30d` | **yes** | 2 |
| R4 | `#e11d48` | no | 2 |
| R5 | `#57534e` | dotted (`dashed=1` + note) | 2 |

4. Still honor YAML `direction` for arrow markers.  
5. Prefer showing the edge `value` when the custom type name alone is unclear.  
6. Note the mapping in chat (and legend, if shown).  
7. Never invent a containment edge for a custom type that only means nesting — use cell `parent`.

### Font & global style

| Role | `fontSize` | `fontColor` |
|------|------------|-------------|
| Container title | 20 | `#1e1e1e` |
| Element label | 16 | `#1e1e1e` |
| Edge label | 14 | `#1e1e1e` |

Prefer `html=1;whiteSpace=wrap` on vertices. Keep stroke/fill language consistent; avoid monochrome-only or random rainbow outside the tables.

### Legends (optional)

Side panel keying **used** element kinds and relationship types. Not part of the YAML graph.

| Policy | Rule |
|--------|------|
| **Show** (judgment) | ≥2 distinct kinds **or** ≥2 distinct relationship types, or styles would be ambiguous |
| **Omit** (judgment) | Single kind and single relationship type; legend would be noise |
| **Force show** | User asks (*add legend*, *show key*) or `legend: true` |
| **Force hide** | User asks (*remove legend*, *no key*) or `legend: false` |

User chat overrides YAML. Note in chat why the legend was included or omitted.

**Content:** Only kinds/types present in the diagram. Elements block (mini swatches + kind name) then Relationships block (sample stroke + type name). Title: `Legend` (or user’s language equivalent).

**Draw:** Separate cells under layer `1` (not nested in story containers); ids `legend-*`; bottom-right or right; ≥40px (prefer) / ≥20px from story, grid-snapped; dashed light container; mini swatches + labels for used kinds/types only; sample edges must not `source`/`target` story cells.

**Remove:** Delete `legend-*` cells only. Never put legend geometry in YAML.

---

## Cell schema

### Vertex (element)

```xml
<mxCell id="api-gateway" value="API Gateway" style="rounded=1;whiteSpace=wrap;html=1;fillColor=#dbeafe;strokeColor=#1d4ed8;strokeWidth=2;fontSize=16;fontColor=#1e1e1e;align=center;verticalAlign=middle;" vertex="1" parent="vpc">
  <mxGeometry x="40" y="60" width="160" height="60" as="geometry"/>
</mxCell>
```

- `parent`: `"1"` for top-level, or container cell id when nested.  
- Geometry for nested cells is **relative to parent**.  
- Label is `value` on the same cell (no separate BoundText).

### Edge (relationship)

```xml
<mxCell id="rel-https" value="HTTPS" style="edgeStyle=orthogonalEdgeStyle;rounded=0;html=1;strokeColor=#2563eb;strokeWidth=2;endArrow=block;endFill=1;startArrow=none;fontSize=14;" edge="1" parent="1" source="client" target="api-gateway">
  <mxGeometry relative="1" as="geometry"/>
</mxCell>
```

Optional waypoints for clear orthogonal paths:

```xml
<mxGeometry relative="1" as="geometry">
  <Array as="points">
    <mxPoint x="200" y="100"/>
    <mxPoint x="200" y="180"/>
  </Array>
</mxGeometry>
```

Prefer ELK `sections` → `<Array as="points">` for orthogonal bends (see **Edges → connectors** under ELK layout). Omit points only when `sections` is empty and auto-route is acceptable. Do not invent a conflicting hand layout.

### Text sizing (Latin + CJK)

Same weights and formulas as **Pre-ELK node sizing** above (in this file). Use them for leaf sizes before ELK and for verifying cell `value` fits after mapping. Prefer slightly too wide over clipping; never break mid-word or mid-syllable.

---

## YAML → draw.io map

| YAML | Intermediate (ELK) | draw.io |
|------|--------------------|---------|
| `elements[]` | `graph.nodes` (+ leaf sizes) | Vertex `mxCell` + `value` = `label` |
| `kind` | — (style only) | Shape + `style` colors / dashed / ellipse |
| `parentId` | `nodes[].parent` | Child `parent="{parentCellId}"` + **relative** geometry |
| `relationships[]` | `graph.edges` | Edge `mxCell` from `sections` + type stroke + direction arrows |
| Containment | Hierarchy / padding | Parent cells only — no contains edge |
| — | `layout_graph` result | Absolute layout → relative cells + grid snap + style |

---

## Reverse reading

For explain / recover / validate-by-explain ([`SKILL.md`](../SKILL.md) Analyze & recover):

1. Parse `mxfile` → `mxCell` list (ignore missing/invisible).  
2. Vertices (`vertex="1"`) → elements; `value` → `label`; style → guess `kind`.  
3. Nesting: non-`1`/`0` `parent` that is a vertex → `parentId`.  
4. Edges (`edge="1"`) → relationships; `source`/`target`; arrows → `direction`; stroke → guess `type`; edge `value` → `label`.  
5. Skip `legend-*`.  
6. Draft `explanation` in visual/flow order.

| Visual cue | Infer |
|------------|--------|
| Ellipse | `kind: actor` |
| Blue fill rect | `kind: system` |
| Gray fill rect | `kind: process` |
| Amber fill | `kind: data` |
| Dashed container with children | `kind: boundary` / `group` |
| Teal / blue / gray-dashed edge | `uses`/`calls` / `flows-to` / `precedes`/`follows` |
| Both / none / end arrows | `bi-directed` / `undirected` / `directed` |

Slug ids from labels or existing cell ids when already kebab-case. Recovered YAML has **no** layout fields.

---

## Output validation

Report Parts 1–5 pass/fail in chat. Do not hide failures.

### Part 1 — Layout

| # | Pass means |
|---|------------|
| 1.0 | Positions/routes came from `cocrates-elk` `layout_graph` (or user-approved offline fallback) |
| 1.1 | Child inset ≥20px (absolute) from parent border |
| 1.2 | Title band (`startSize`) clear of children: first child `rel y ≥ 60` (40 band + ≥20 gap) |
| 1.3 | Multi-level nesting distinct |
| 1.4 | No containment-only edges |
| 1.5 | Peer gap ≥20px (prefer 40); no peer overlap |
| 1.6 | Layering / reading order matches `explanation` + ELK direction |
| 1.7 | Peer leaf sizes unified before ELK; nested geometry is parent-relative and correct |
| 1.8 | All vertex `x,y,width,height` are multiples of **20** (grid snap) |
| 1.9 | Edge routes follow ELK sections (or documented auto-route fallback) |
| 1.10 | ≥40px canvas margin; balanced; orthogonal waypoints prefer grid |


### Part 2 — Typography

| # | Pass means |
|---|------------|
| 2.1 | Labels fit inside vertices (no clip) |
| 2.2 | Shape sized with ≥10–15px text padding |
| 2.3 | CJK/mixed labels use weighted widths |
| 2.4 | `value` set; XML-escaped; readable |
| 2.4a | Leaf vertices: `verticalAlign=middle`, `fontSize=16` |
| 2.4b | Container **with children**: `swimlane;startSize=40` (title in top band) |
| 2.4c | Container **without children**: no swimlane; `verticalAlign=middle` in full box |
| 2.5 | Clean line breaks (`&#xa;`) |
| 2.6 | Edge labels do not obscure connectors/shapes |
| 2.7 | Font hierarchy 20 / 16 / 14 |

### Part 3 — Edges

| # | Pass means |
|---|------------|
| 3.1 | `source` / `target` resolve to vertices |
| 3.1a | When ELK returned `sections`, edge waypoints follow those bends (not ignored) |
| 3.2 | Orthogonal preferred; ELK `sections` → `<Array as="points">`; minimal crossings |
| 3.3 | Waypoints from ELK bends (or documented fallback); no body-through mess |
| 3.4 | Flow matches `explanation` |
| 3.5 | Arrow markers match `direction` |
| 3.6 | Every YAML relationship has an edge |

### Part 4 — Style & types

| # | Pass means |
|---|------------|
| 4.1 | Sufficient contrast |
| 4.2 | Containers dashed / distinct from leaves |
| 4.3 | Kind encoding consistent (built-in or stable custom) |
| 4.4 | Co-existing kinds visually distinct |
| 4.5 | Relationship type stroke encoding correct |
| 4.6 | Co-existing types visually distinct |
| 4.7 | No `has-a`/`contains` edges |
| 4.8 | Style language consistent across file |
| 4.9 | Legend policy respected; noted in chat |
| 4.10 | If shown: accurate used kinds/types; no story overlap |
| 4.11 | If shown: `legend-*` ids; sample edges unbound to story |

### Part 5 — Script & file

| # | Pass means |
|---|------------|
| 5.1 | `explanation` concepts on canvas |
| 5.2 | Natural pointing walkthrough |
| 5.3 | No junk (legend chrome OK) |
| 5.4 | No invented story edges/nodes |
| 5.5 | Well-formed XML; unique ids; opens in diagrams.net / draw.io |

---

## Prohibitions

- Hand-placing the story graph when `cocrates-elk` is available (must use `layout_graph`)  
- Skipping `validate_graph` when the graph is non-trivial (nesting, many edges)  
- Writing ELK coordinates or layout options back into YAML  
- Inventing story cells not in YAML  
- Containment-only edges instead of `parent` nesting  
- Absolute child geometry when `parent` is a container (must be parent-relative after ELK absolute → relative conversion)  
- Unescaped XML in `value`  
- Using `swimlane;startSize=40` on a container **without** children (label should center in full box)
- Ignoring non-empty ELK `sections` and hand-routing edges instead
- Title-band / inset / peer-spacing / **grid-snap** violations  
- Homogeneous styling that ignores `kind` / `type`  
- Bad legends; ignoring user legend show/hide  
- Writing layout or legend geometry into YAML  
- Shipping without Parts 1–5 validation report  
- Emitting a different backend format from this generator (use the matching file under [`README.md`](README.md))
