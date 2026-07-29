# Excalidraw Generator (Stage 3 — How)

Turn approved YAML (**what**) into Excalidraw JSON (**how**), or **reverse-read** an `.excalidraw` file for analysis / YAML recovery.

- Forward & reverse workflow: [`SKILL.md`](../SKILL.md)
- Other generators: [`README.md`](README.md)
- **Layout engine:** MCP `cocrates-elk` (`validate_graph` → `layout_graph`) — then map geometry to Excalidraw

**Hard rules:** Do not invent story meaning, elements, or relationships. Do not write layout coordinates into YAML. Do not hand-place the story graph when the ELK MCP is available.

## When to use

| Situation | Action |
|-----------|--------|
| Generate / export Excalidraw | Size → **ELK layout MCP** → map to Excalidraw → validate (this file) |
| Stage 2 YAML approved, no `.excalidraw` yet | Same |
| Analyze / explain / recover YAML from `.excalidraw` | **Reverse reading** + [`SKILL.md`](../SKILL.md) Analyze & recover |
| Another backend | [`README.md`](README.md) |

**Forward prerequisite:** Approved `title`, `summary`, `explanation`, `diagram.elements`, `diagram.relationships`.

## Input / output

| Direction | Input | Output |
|-----------|-------|--------|
| Forward | `{slug}.yaml` (+ optional `backend`, `output`, `legend`) | `{slug}.excalidraw` (same folder unless `output` overrides) |
| Reverse | `.excalidraw` | Explanation and/or Stage-2 YAML (what only) |

## Pipeline (forward)

```
1. Read approved YAML
2. Size leaves (text heuristics + peer unify + grid)
3. Build ELK graph from YAML → validate_graph → layout_graph (cocrates-elk MCP)
4. Map layout → Excalidraw shapes / BoundText / arrows (+ style + optional legend)
5. Post-process: snap to 20px grid; fix dock points; emit .excalidraw
6. Run Parts 1–5 validation → report pass/fail to user
7. On failure → revise JSON (or re-layout with adjusted sizes/options) and re-validate
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
| **Grid size** | **20px** — snap all shape `x` / `y` / `width` / `height` (and arrow bend points) after layout |
| Parent / child inset | ≥ **20px** (prefer **40px** via ELK padding) |
| Title band (parent with label) | **40px** top padding in ELK (`elk.padding` top) + BoundText in band |
| Peer gap | **40px** preferred → `elk.spacing.nodeNode: 40` |
| Layer gap | **60px** preferred → `elk.layered.spacing.nodeNodeBetweenLayers: 60` |
| Canvas outer margin | ≥ **40px** (add after layout if content hugs 0,0) |
| Text padding inside shape | ≥ **10–15px** per side; then snap leaf size up to grid **before** ELK |
| Standard leaf size | **Width 160** / **height 60** when label fits; grow in **20px** steps |
| Standard actor (ellipse) | **Width 140** / **height 60** (or 160×80), snapped |
| Arrow binding `gap` | **5** |
| Font sizes | Title **20** / label **16** / edge label **14** |
| Default `roughness` | **1** (uniform in file) |
| Default `fillStyle` | **`solid`** |

---

## Layout via cocrates ELK MCP (required)

### MCP workflow

Before calling tools, discover schemas with **GetMcpTools** for server `cocrates-elk`.

| Step | Tool | Purpose |
|------|------|---------|
| 0 (optional) | `list_algorithms` / `list_layout_options` | Pick algorithm or tune options when default is a poor fit |
| 1 | `validate_graph` | Catch duplicate ids, bad parents/endpoints, cycles, missing leaf sizes |
| 2 | `layout_graph` | Obtain absolute `x,y,width,height` per node and edge `sections` (start/end/bend points) |

If the server is `needsAuth`, call `mcp_auth` once, then retry. If layout MCP is unavailable or errors after retry, **stop and tell the user** — do not silently invent a full manual placement (offline fallback only if the user explicitly asks).

### YAML → ELK `graph`

Build a **flat** node list (prefer `parent` over nested `children` — do not mix both on one node).

| YAML | ELK graph field |
|------|-----------------|
| `diagram.elements[].id` | `nodes[].id` (keep kebab-case if the tool accepts it; otherwise sanitize and keep a map) |
| `label` | `nodes[].label`; for containers also `labels: [{ id, text, width, height }]` to reserve title space |
| `parentId` | `nodes[].parent` |
| Leaf size (precomputed) | `nodes[].width` / `height` (**required for leaves**) |
| Container size | **Omit** `width`/`height` on parents so ELK sizes them from children + padding |
| `diagram.relationships[]` | `edges[]` with `id`, `source`, `target` (skip `has-a` / `contains` — nesting only) |
| Relationship `label` | Optional `edges[].labels: [{ text, width, height }]` |

**Do not** put Excalidraw colors, arrowheads, or BoundText into the ELK graph — only structure + sizes + labels for layout.

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
| Soft-wrap | If longer, insert `\n` at word/phrase boundaries so each line’s `lineWidth` stays ≤ **200** (≈ shape content width before padding) |
| Never | Mid-word / mid-syllable breaks |

Use the wrapped label for **both** ELK sizing and the eventual BoundText `text`.

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
| Padding | Rely on `elk.padding` top **40** (title band); do not fake a huge empty leaf inside the container |

#### Edge labels (optional on ELK edges)

If the relationship has a `label`, you may pass `edges[].labels: [{ text, width, height }]` with **fontSize 14**:

```
width  = ceil((textWidth + 8) / 20) * 20
height = ceil((textHeight + 4) / 20) * 20   # often 20
```

If omitted, still draw the unbound Excalidraw edge label after layout.

#### Checklist before `validate_graph`

- [ ] Every **leaf** has positive `width` and `height` (grid-snapped).  
- [ ] No leaf uses the tool default ~80×40 “because sizing was skipped”.  
- [ ] Peer sets share identical W×H.  
- [ ] Parents have **no** fixed width/height.  
- [ ] Container `labels[]` present when the parent has a visible title.  
- [ ] Wrapped label text matches what will appear on the canvas.

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
    "elk.padding": "[top=40,left=20,bottom=20,right=20]",
    "elk.layered.nodePlacement.bk.fixedAlignment": "BALANCED"
  }
}
```

| Cue from `explanation` | `options.direction` |
|------------------------|---------------------|
| Left→right / request flow | `RIGHT` |
| Top→bottom / hierarchy stack | `DOWN` |
| Unclear | `RIGHT` if mostly flow edges; else `DOWN` |

**Algorithm choice (defaults):**

| Graph shape | Prefer |
|-------------|--------|
| Flow / pipeline / layered architecture (default) | `layered` |
| Pure tree from one root | `mrtree` |
| Dense undirected association mesh | `force` or `stress` (rare) |
| Unconnected boxes only | `box` / `rectpacking` |

When any compound nesting exists, keep **`layered` + `INCLUDE_CHILDREN`** unless the user asks otherwise.

### ELK layout → Excalidraw

Assume `absoluteCoordinates: true` (default).

**Nodes → shapes**

1. Map each laid-out node id → one rectangle or ellipse (from `kind` table).  
2. Use layout `x`, `y`, `width`, `height`.  
3. **Snap** each to multiples of 20 (`round` or `ceil` consistently; prefer values that preserve ≥20 inset).  
4. If content bbox is near the origin, translate so outer margin ≥40.  
5. Parents before children in `elements[]`. Nesting is geometric (and `frameId` if using frames).  
6. Add BoundText trinity from YAML `label` — **verticalAlign by role** (see BoundText trinity below).  
   Container title `x` = `container.x + (container.width - textWidth) / 2` for center alignment.  
7. Apply kind palette / stroke — styling is **not** from ELK.

**Edges → arrows**

Use **`layout_graph` edge output** — do not invent routes when `sections` is non-empty.

| ELK field | Use |
|-----------|-----|
| `edges[].id` | Match YAML relationship `id` (same id sent in `graph.edges`) |
| `edges[].sources[0]` / `targets[0]` | `startBinding` / `endBinding` element ids |
| `edges[].sections[]` | Orthogonal polyline — **required** when present |

**Section → polyline algorithm**

For each edge, walk `sections` in order. For each section:

```
points_abs += [section.startPoint]
points_abs += section.bendPoints ?? []    # may be absent on straight segments
points_abs += [section.endPoint]
```

De-duplicate consecutive identical coordinates. Then:

1. Snap every point to the 20px grid.  
2. Set arrow `x` = first point `x`, `y` = first point `y`.  
3. Set `points` = `[[0,0], [p1.x-x, p1.y-y], …]` (relative; first entry always `[0,0]`).  
4. `startBinding` / `endBinding` → source/target shape ids; `gap: 5`. After snap, if the first/last polyline point drifted off the border, adjust only the **endpoint** to the nearest facing border point; keep **intermediate** bends from ELK.  
5. Stroke + arrowheads from relationship `type` / `direction`.  
6. Edge label: unbound text near the polyline midpoint, offset 6–10px.

**When `sections` is empty**

| Cause | Action |
|-------|--------|
| Missing `elk.hierarchyHandling: INCLUDE_CHILDREN` on nested graphs | Fix options; **re-call** `layout_graph` |
| Cross-hierarchy edge without compound support | Fix hierarchy options or graph structure |
| Otherwise | Document in validation; fall back to facing midpoints + orthogonal corridor; still bind |

**Do not** ignore non-empty `sections` and draw a hand-routed straight arrow instead.

**Legend:** Still hand-composed after the story layout (not sent to ELK). Place on grid to the right or bottom of the story bbox.

### Post-layout polish (still required)

| Rule | Detail |
|------|--------|
| **Snap to grid** | All shape geometry and arrow bends → multiples of 20 |
| **Inset check** | Children fully inside parents with ≥20px inset after snap; if violated, expand parent or nudge children and re-dock arrows |
| **Title band** | Children below parent top + title band; ELK top padding should already reserve this — verify visually |
| **Uniform peers** | Leaf sizes were unified **before** ELK; do not randomly resize after layout except snap |
| **No manual reshape** of the overall layering ELK chose — fix sizing/options and **re-call** `layout_graph` if the structure is wrong |

### Composition (semantic rules unchanged)

- **Reading order:** Match `explanation` via ELK `direction` + edge flow.  
- **Containment:** `parentId` → ELK `parent` → nested geometry. **No** containment arrow.  
- **Kinds / types / labels / legend:** As in Visual distinction below.  
- **Arrows:** One per relationship; prefer ELK orthogonal sections.

### Peer box test (validation)

Skip if A is ancestor of B via `parentId`:

```
peer overlap_area = 0
peer gap on separating axis ≥ 20 (prefer 40)
all x,y,width,height % 20 == 0
parent↔child: child inside parent; child.top ≥ parent.top + titleBand + 20 when parent has a label
layout positions originated from layout_graph (not freehand invent)
```

---

## Document envelope

```json
{
  "type": "excalidraw",
  "version": 2,
  "source": "https://excalidraw.com",
  "elements": [],
  "appState": { "viewBackgroundColor": "#ffffff" },
  "files": {}
}
```

## IDs, seed, versionNonce

| Field | Type | Rule |
|-------|------|------|
| `id` | string | Unique, ≥8 chars (e.g. `rect_a3f91c02`). Prefer `{role}_{hex}` |
| `seed` | positive 32-bit **integer** | `1 … 2_000_000_000` — **never a string** |
| `versionNonce` | positive 32-bit **integer** | Independent of `seed` |
| `version` | integer | Start at `1` |
| `updated` | integer (ms) | Unix epoch ms |

---

## Visual distinction

Viewers must distinguish YAML `kind` / `type` without relying only on labels. Same kind/type → same look everywhere.

### Element `kind` → shape & palette

| `kind` | Shape | `backgroundColor` | `strokeColor` | `strokeStyle` | `strokeWidth` |
|--------|-------|-------------------|---------------|---------------|---------------|
| `actor` | **ellipse** | `#e0f2fe` | `#0284c7` | `solid` | 2 |
| `system` | rectangle | `#dbeafe` | `#1d4ed8` | `solid` | 2 |
| `process` | rectangle | `#f3f4f6` | `#374151` | `solid` | 2 |
| `data` | rectangle | `#fef3c7` | `#d97706` | `solid` | 2 |
| `boundary` | large rectangle | `#fafafa` | `#9ca3af` | **dashed** | 2 |
| `group` | large rectangle | `#f8fafc` | `#94a3b8` | **dashed** | 2 |
| `frame` | `type: "frame"` | — | — | — | — |
| *(built-in missing)* | rectangle | `#ffffff` | `#1e1e1e` | `solid` | 2 |

Built-in rows are **defaults**, not a closed enum. YAML may use any English kebab-case `kind`.

**Custom / unknown `kind`**

1. Prefer mapping to the closest built-in when the meaning matches (e.g. `service` → `system`, `db` → `data`, `zone` → `boundary`).  
2. Otherwise treat as **custom**: pick one stable style for that slug and reuse it for every element with the same `kind` in this diagram.  
3. Choose a palette that still differs from co-existing kinds (≥2 of shape / fill / stroke color / stroke style). Suggested custom pool (cycle if many customs):

| Slot | Shape | `backgroundColor` | `strokeColor` | `strokeStyle` |
|------|-------|-------------------|---------------|---------------|
| C1 | rectangle | `#fce7f3` | `#db2777` | `solid` |
| C2 | rectangle | `#ede9fe` | `#7c3aed` | `solid` |
| C3 | rectangle | `#ccfbf1` | `#0d9488` | `solid` |
| C4 | rectangle | `#ffedd5` | `#ea580c` | `solid` |
| C5 | ellipse | `#e0e7ff` | `#4f46e5` | `solid` |

4. If the custom kind is container-like (holds children via `parentId`), use **dashed** stroke like `boundary`/`group`.  
5. Note the mapping in chat (and legend, if shown): e.g. `kind: message-broker → custom C3`.  
6. Prefer setting an explicit `kind` in YAML over leaving it empty.

### Relationship `type` → connector style

Plus `direction` arrowheads (`directed` / `undirected` / `bi-directed`).

| `type` | `strokeColor` | `strokeStyle` | `strokeWidth` |
|--------|---------------|---------------|---------------|
| `flows-to` | `#2563eb` | `solid` | 2 |
| `requests` | `#7c3aed` | `solid` | 2 |
| `uses` / `calls` | `#0f766e` | `solid` | 2 |
| `precedes` / `follows` | `#4b5563` | **dashed** | 2 |
| `is-a` | `#b45309` | `solid` | 2 |
| `implements` | `#be185d` | **dashed** | 2 |
| `associates` | `#6b7280` | `solid` | **1** |
| `has-a` / `contains` | — | — | **Do not draw** — nest via `parentId` |

Built-in rows are **defaults**, not a closed enum. YAML may use any English kebab-case relationship `type`.

**Custom / unknown `type`**

1. Prefer mapping to the closest built-in when the meaning matches (e.g. `publishes` → `flows-to`, `depends-on` → `uses`, `then` → `precedes`).  
2. Otherwise treat as **custom**: assign one stable stroke for that slug and reuse it for every relationship with the same `type` in this diagram.  
3. Keep it visually distinct from co-existing types (color and/or dash and/or width). Suggested custom pool:

| Slot | `strokeColor` | `strokeStyle` | `strokeWidth` |
|------|---------------|---------------|---------------|
| R1 | `#0891b2` | `solid` | 2 |
| R2 | `#c026d3` | `solid` | 2 |
| R3 | `#65a30d` | **dashed** | 2 |
| R4 | `#e11d48` | `solid` | 2 |
| R5 | `#57534e` | **dotted** | 2 |

4. Still honor YAML `direction` for arrowheads.  
5. Prefer showing the edge `label` when the custom type name alone is unclear.  
6. Note the mapping in chat (and legend, if shown).  
7. Never invent a containment arrow for a custom type that only means nesting — use `parentId`.

**Consistency:** Within one diagram, the kind/type → style map is fixed for that file. Do not recolor the same custom slug differently on regenerate unless the user asks.

### Font & global style

| Role | `fontSize` |
|------|------------|
| Container title | 20 |
| Element label | 16 |
| Edge label | 14 |

Text `strokeColor: "#1e1e1e"`, `fontFamily: 1`. Keep `roughness` and `fillStyle` uniform; avoid monochrome-only or random rainbow palettes outside the tables.

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

**Draw:** Place bottom-right or right; ≥40px (prefer) / ≥20px from story peers, grid-snapped; light/dashed container; swatches ~40×20 snapped; text 12–14px; ids `legend_*` / `legtxt_*` / `legarr_*`; optional `groupIds: ["legend"]`; sample arrows must **not** bind to story shapes.

**Remove:** Delete legend-prefixed/grouped elements only. Never write legend geometry into YAML.

---

## Element schema

### Rectangle / Ellipse

- Required: `id`, `type` (`rectangle` \| `ellipse`), `x`, `y`, `width`, `height`
- Style from kind table + `fillStyle`, `roughness`
- If labeled: `boundElements: [{ "type": "text", "id": "<text_id>" }]`
- Defaults: `angle: 0`, `strokeWidth` from table, `opacity: 100`, `groupIds: []`, `roundness: null` or `{ "type": 3 }`, integer `seed` / `versionNonce`, `version: 1`, `isDeleted: false`, `updated`, `locked: false`, `link: null`, `frameId: null`

**Never** put `label` / `title` on the shape.

### Frame

Required: `id`, `type: "frame"`, `x`, `y`, `width`, `height`, `name`. Children set `frameId`; ≥20px inner padding.

### BoundText trinity (required for shape labels)

| Side | Property | Rule |
|------|----------|------|
| Shape | `boundElements` | `[{ "type": "text", "id": "<text_id>" }]` |
| Text | `containerId` | Shape `id` |
| Text | `textAlign` | **`center`** |
| Text | `verticalAlign` | **By role** — see table below |
| Text | `autoResize` | **`true`** |
| Text | `text` / `originalText` | Identical strings |
| Text | `fontSize` | **16** for leaves; **20** for container titles |
| Text | geometry | Container titles reuse the same `width`/`height` computed for the ELK label (see **Containers** sizing in Pre-ELK section). Leaves sized per **Pre-ELK node sizing** formulas. ≥10–15px inset. |

#### `verticalAlign` by role (required)

Excalidraw defaults to `middle`, which centers text in the **full** shape — wrong for containers that hold children (title must sit in the top band). draw.io avoids this with swimlane `startSize`; Excalidraw must set `verticalAlign` explicitly.

| Element role | Has YAML children? (`parentId` → this id) | `verticalAlign` | Notes |
|--------------|-------------------------------------------|-----------------|-------|
| Leaf (actor, system, process, data, …) | — | **`middle`** | Label centered in the box |
| Container / boundary / group / frame | **Yes** | **`top`** | Title in title band; children start below band + inset |
| Container with label but **no** children | **No** | **`middle`** | Label centered — same as a leaf box |

Detect “has children” from YAML `diagram.elements[].parentId`, not from ELK output alone.

**Leaf example** (`verticalAlign: "middle"`):

```json
{
  "id": "text_b7e20d11",
  "type": "text",
  "containerId": "rect_a3f91c02",
  "text": "API Gateway",
  "originalText": "API Gateway",
  "textAlign": "center",
  "verticalAlign": "middle",
  "autoResize": true,
  "fontSize": 16,
  "fontFamily": 1
}
```

**Container with children** (`verticalAlign: "top"`, `fontSize: 20`):

```json
{
  "id": "text_vpc_title",
  "type": "text",
  "containerId": "rect_vpc",
  "text": "VPC",
  "originalText": "VPC",
  "textAlign": "center",
  "verticalAlign": "top",
  "autoResize": true,
  "fontSize": 20,
  "fontFamily": 1
}
```

Edge labels: unbound text (`fontSize: 14`), no `containerId`; offset 6–10px off the stroke so they do not cover the arrow or borders.

### Text size (Latin + CJK)

Same weights and formulas as **Pre-ELK node sizing** above. Use them for BoundText geometry checks after mapping layout → shapes. Prefer slightly too wide over clipping; never break mid-word or mid-syllable.

---

## Arrows

### Schema

- `type: "arrow"`, `x`, `y`, `points` (**relative** to `x,y`; first point `[0,0]`)
- Bindings: `{ "elementId": "…", "focus": 0, "gap": 5 }`
- Stroke from Relationship `type` table

| `direction` | `startArrowhead` | `endArrowhead` |
|-------------|------------------|----------------|
| `directed` | `null` | `"arrow"` |
| `undirected` | `null` | `null` |
| `bi-directed` | `"arrow"` | `"arrow"` |

### Docking & routing

1. **Primary source:** `layout_graph` → `edges[].sections[]` → absolute polyline → relative Excalidraw `points` (see **Edges → arrows** under ELK layout).  
2. After node snap, re-dock **endpoints only** onto shape borders; preserve ELK intermediate bends.  
3. Snap all polyline points to the 20px grid.  
4. **`sections` empty:** documented fallback only — fix ELK options and re-layout when nesting caused empty sections.  
5. Minimize crossings; flow matches `explanation` reading order.

---

## YAML → Excalidraw map

| YAML | Intermediate (ELK) | Excalidraw |
|------|--------------------|------------|
| `elements[]` | `graph.nodes` (+ leaf sizes) | Shape + BoundText from `label` |
| `kind` | — (style only) | Shape + palette row |
| `parentId` | `nodes[].parent` | Nested geometry; parents before children |
| `relationships[]` | `graph.edges` | Arrow from edge `sections` + type stroke + direction heads |
| Containment | Hierarchy / padding | Nesting only (no arrow) |
| — | `layout_graph` result | Positions, sizes, routes → then grid snap + style |

---

## Reverse reading

For explain / recover / validate-by-explain ([`SKILL.md`](../SKILL.md) Analyze & recover):

1. Parse `elements[]` (skip `isDeleted`).  
2. Shapes → bounds, colors, `strokeStyle`, `frameId`.  
3. BoundText (`containerId`) → labels.  
4. Nesting: child inside parent → `parentId`; `frameId` → zone.  
5. Arrows: bindings → from/to; heads → `direction`; stroke → guess `type`; nearby unbound text → edge `label`.  
6. Skip `legend_*` / legend group.  
7. Draft `explanation` in visual/flow order.

| Visual cue | Infer |
|------------|--------|
| Ellipse + sky fill | `kind: actor` |
| Blue rectangle | `kind: system` |
| Gray rectangle | `kind: process` |
| Amber fill | `kind: data` |
| Large dashed container | `kind: boundary` or `group` |
| Teal / blue / gray-dashed arrow | `uses`/`calls` / `flows-to` / `precedes`/`follows` |
| Arrowheads both / none / end only | `bi-directed` / `undirected` / `directed` |

Slug ids from labels (`API Gateway` → `api-gateway`). Note uncertainty. Recovered YAML has **no** layout fields.

---

## Output validation

After emit (or when validating an existing file), report Parts 1–5 pass/fail in chat. Do not hide failures. Fix critical fails before treating the job as done.

### Part 1 — Layout

| # | Pass means |
|---|------------|
| 1.0 | Positions/routes came from `cocrates-elk` `layout_graph` (or user-approved offline fallback) |
| 1.1 | Child inset ≥20px from parent |
| 1.2 | Title band clear of children |
| 1.3 | Multi-level nesting visually distinct |
| 1.4 | No containment-only arrows |
| 1.5 | Peer gap ≥20px (prefer 40); no peer overlap |
| 1.6 | Layering / reading order matches `explanation` + ELK direction |
| 1.7 | Peer leaf sizes were unified before ELK; containers sized by layout (then snapped) |
| 1.8 | All shape `x,y,width,height` are multiples of **20** (grid snap) |
| 1.9 | Arrow routes follow ELK sections (or documented midpoint fallback) |
| 1.10 | ≥40px canvas margin; balanced placement; orthogonal elbows prefer grid |

### Part 2 — Typography

| # | Pass means |
|---|------------|
| 2.1 | No text overflow |
| 2.2 | ≥10–15px text padding |
| 2.3 | CJK/mixed labels use weighted widths |
| 2.4 | Full BoundText trinity + `autoResize: true` |
| 2.4a | Leaf labels: `verticalAlign: "middle"` |
| 2.4b | Container **with children**: `verticalAlign: "top"`, `fontSize: 20` |
| 2.4c | Container **without children**: `verticalAlign: "middle"` |
| 2.5 | Clean line breaks |
| 2.6 | Edge labels clear of strokes/borders |
| 2.7 | Font hierarchy 20 / 16 / 14 |

### Part 3 — Arrows

| # | Pass means |
|---|------------|
| 3.1 | Docked on borders (`startBinding` / `endBinding`; `gap` ≈ 5) |
| 3.1a | When ELK returned `sections`, arrow polyline follows those bends (not a hand straight line) |
| 3.2 | Relative `points`; first `[0,0]` |
| 3.3 | Bindings match from/to; `gap` ≈ 5 |
| 3.4 | Minimal crossings; orthogonal when needed |
| 3.5 | Flow matches `explanation` |
| 3.6 | Arrowheads match `direction` |

### Part 4 — Style & types

| # | Pass means |
|---|------------|
| 4.1 | Sufficient contrast |
| 4.2 | Containers dashed / distinct from leaves |
| 4.3 | Kind encoding correct & consistent (built-in or stable custom mapping) |
| 4.4 | Co-existing kinds visually distinct (≥2 cues) |
| 4.5 | Relationship type stroke encoding correct (built-in or stable custom mapping) |
| 4.6 | Co-existing types visually distinct |
| 4.7 | No `has-a`/`contains` arrows |
| 4.8 | Uniform `roughness` |
| 4.9 | Consistent `fillStyle` |
| 4.10 | Legend show/omit matches policy; noted in chat |
| 4.11 | If shown: accurate used kinds/types only; no overlap with story |
| 4.12 | If shown: unbound samples; `legend_*` ids; removable safely |

### Part 5 — Script & file

| # | Pass means |
|---|------------|
| 5.1 | `explanation` concepts present on canvas |
| 5.2 | Natural pointing walkthrough |
| 5.3 | No junk decorations (legend chrome OK) |
| 5.4 | Every YAML relationship has a connector; no invented edges |
| 5.5 | Valid JSON; unique ids; integer seeds; opens in Excalidraw |

---

## Prohibitions

- Hand-placing the story graph when `cocrates-elk` is available (must use `layout_graph`)
- Skipping `validate_graph` when the graph is non-trivial (nesting, many edges)
- Writing ELK coordinates or layout options back into YAML
- Fake shape properties (`label` / `title` on shapes)
- Incomplete BoundText or missing `autoResize`
- String `seed` / `versionNonce`
- Absolute coords in arrow `points`; floating/buried tips; `gap: 0`
- Wrong `verticalAlign` on containers (must be `top` when the element has children; `middle` when it does not)
- Ignoring non-empty ELK `sections` and hand-routing arrows instead
- Title-band / inset / peer-spacing / **grid-snap** violations
- Latin-only width math for CJK; mid-word `\n`
- Homogeneous styling that ignores `kind` / `type`
- Bad legends (unused rows, overlap, bound into story); ignoring user legend show/hide
- Containment arrows; inventing story graph not in YAML
- Writing layout or legend geometry into YAML
- Shipping without a Parts 1–5 validation report
