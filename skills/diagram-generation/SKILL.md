---
name: diagram-generation
description: >-
  Designs explanatory diagrams via Snowflake stages (requirements → structure →
  detail → Excalidraw) with YAML artifacts and user approval gates at each
  design stage. Select when the user asks to draw, diagram, visualize, map,
  sketch, or explain a system, architecture, flow, process, hierarchy, ER/class
  model, concept map, or relationships as a diagram — especially Excalidraw.
  Prioritizes locking an explanation script (what you say while pointing at the
  diagram) plus meta description (goal, audience, scope), validates structure and
  layout against that script with the user, then emits valid Excalidraw JSON
  (shapes and text as separate elements). Each stage writes only its own YAML
  keys and stops for approval — no look-ahead structure or layout during Stage 1.
compatibility: opencode
metadata:
  agent: cocrates
---

# Diagram Generation

Produce diagrams that **make an approved explanation script deliverable** — not decorative drawings. Stage 1 locks `description` (meta) and `explanation` (walkthrough script); later stages design a picture that fits that script. Deliver a staged YAML design spec and valid **Excalidraw** JSON (`.excalidraw` / `.json`).

## Core Principles

> **The unexamined diagram is not worth generating.**

- **Explanation-First**: The primary deliverable of Stage 1 is an **`explanation` script** — the words someone would use while pointing at the finished diagram. The picture must make that script true and easy to follow. Meta fields (`description`) set goal/audience; they are not a substitute for the script.
- **Snowflake method**: Refine in order: requirements → structure design → detail design → generation. Do not skip ahead.
- **Per-stage artifacts + approval gates**: Each stage **writes or updates the YAML file with only that stage’s results**, then **stops for user review and explicit approval** before the next stage. The user may edit the YAML directly; treat their file as source of truth.
- **No look-ahead in the file**: Never pre-fill later-stage keys (`elements`, `relationships`, `diagram`, or Excalidraw) while still on an earlier stage. Grow the YAML **incrementally**.
- **Design before generation**: Do **not** emit Excalidraw JSON until detail design is approved.
- **Bi-directional Iteration**: If later stages expose gaps or contradictions, halt and roll back to the relevant upper stage; update the YAML baseline, re-approve, then continue.
- **Containment is structural**: If A **contains** B, B is drawn **inside** A’s shape — never as a peer box linked only by a containment arrow. Encode with `parentId` / `parentElementId` (and optional `children`); reserve arrows for non-nesting relations (`uses`, `flows-to`, …).
- **Minimal Valid Excalidraw**: Use the allowed element subset. Never invent properties (no `label` on shapes). Shapes and text are **separate** elements.
- **YAML as Contract**: Intermediate design lives in `{diagram-slug}.yaml`; final art defaults to `{diagram-slug}.excalidraw` in the **same folder**.

## Workflow

```
1 Requirements  →  (YAML + approve)
2 Structure     →  (YAML + validation report + approve)
3 Detail design →  (YAML + layout validation + approve)
4 Generate Excalidraw
5 Revise (roll back to the stage that owns the change)
```

**Hard gates:** Never start stage *N+1* until stage *N*’s YAML is saved and the user has approved (or edited and confirmed). Soft-skip only when the user explicitly waives a gate — or when **Express Mode** applies (below).

### Progressive YAML growth (mandatory)

Write **only** the keys unlocked by the current stage. Do not invent placeholders for future stages.

| After stage | YAML may contain | YAML must **not** yet contain |
|-------------|------------------|-------------------------------|
| **1 Requirements** | `title`, `description`, `explanation`, optional `output`, optional `mode` / `packaging` | `elements`, `relationships`, `diagram`, any Excalidraw file |
| **2 Structure** | Stage 1 keys **plus** `elements`, `relationships` | `diagram`, any Excalidraw file |
| **3 Detail design** | Stage 2 keys **plus** `diagram` | Excalidraw file (until Stage 4) |
| **4 Generation** | Full YAML **plus** `.excalidraw` sibling | — |

**Stop rule:** End the turn after saving the current stage’s YAML and requesting approval. Do not continue into the next stage in the same turn unless the user already approved in that turn **or Express Mode is active**.

### Express Mode (fast-track)

Use when the user asks to finish in one go (`draw it all for me`, `do it in one pass`, `fast-track`, `express`, etc.) **or** the expected graph is tiny (roughly **≤5 elements**, single form, no nested zones).

| Behavior | Rule |
|----------|------|
| Still write YAML **progressively in logic** | Mentally (or in file history) respect stage order; do not invent meaning without an `explanation` |
| Gates | May combine Stages 1→4 in fewer turns; still save a complete YAML before/with Excalidraw |
| Safety | If Express would hide a contested explanation, **fall back to normal gates** and ask |
| Record | Set `mode: express` in Stage 1 YAML when used |

Normal gated flow remains the **default**.

Match the user's language in chat. Keep YAML field names and Excalidraw JSON in English keys; diagram *labels* may follow the user's language.

## Working Location

### Spec folder (YAML)

When the user does **not** name a location, write the design YAML under the workspace kind folder:

```text
diagrams/{diagram-slug}.yaml
```

If the user specifies a folder or path, use that instead. Prefer an existing `diagrams/` tree over inventing a parallel layout.

### Generated artifacts (Excalidraw)

| Rule | Path |
|------|------|
| **Default** | Same directory as the YAML file (e.g. `diagrams/{slug}.excalidraw`) — **not** `./output/` |
| **User override** | Exact folder or file path the user named |

Optional: record the path as `output: "./{slug}.excalidraw"` on the YAML when helpful for runtimes.

**Example layout (defaults):**

```text
diagrams/
├── auth-overview.yaml
└── auth-overview.excalidraw
```

---

## Stage 1 — Requirements

Lock **why** the diagram exists and **what will be said** while looking at it. This is the **most important** review: if `explanation` is wrong, every later stage is wasted.

Separate two fields:

| Field | Role |
|-------|------|
| **`description`** | Brief / meta: Goal, Audience, Scope, diagram kind, constraints, success criteria |
| **`explanation`** | **Spoken (or written) script** that walks through the finished diagram — the actual teaching narrative the picture must support |

The diagram is drawn **to fit `explanation`**, not merely to satisfy abstract goals in `description`.

### Capture checklist

**Into `description`:**

| Capture | Questions |
|---------|-----------|
| **Goal** | What should the viewer understand after looking? |
| **Audience** | Beginner overview vs expert detail? |
| **Scope** | In / out of frame; depth |
| **Diagram kind** | Flow, architecture, hierarchy, ER/class, swimlanes, concept map, … |
| **Constraints** | Must-include terms, reading direction, color, max boxes |
| **Success criteria** | One sentence: “After reading this diagram, the viewer can …” |
| **Packaging** | One file vs split files; if one file, whether to use **frames/zones** (see Choosing the Optimal Diagram Form) |

**Into `explanation`:**

Write a clear walkthrough as if narrating the slide/board aloud (or a caption sequence). Prefer ordered beats the eye can follow (e.g. “First… Then… Finally…”). Name the same concepts that will appear as boxes/edges. Do **not** dump only meta goals here — this is the **script**.

If the script naturally splits into independent stories (e.g. “deploy path” vs “data model”), decide in Stage 1: **separate YAML/Excalidraw files** or **one file with named frames** — do not cram unrelated narratives onto one undivided canvas.

Do **not** start listing boxes until both `description` and a usable `explanation` are clear.

### 1.1 Write requirements YAML (Stage 1 only)

Create `{diagram-slug}.yaml` with **exactly** these top-level keys (plus optional `output`, `mode`, packaging notes inside `description`). **Do not** add `elements`, `relationships`, or `diagram` yet.

```yaml
title: "{short title}"
# mode: express   # only when Express Mode / fast-track was requested or auto-eligible

description: |
  Goal: {what the viewer should understand}
  Audience: {who}
  Scope: {in / out}
  Diagram kind: {flow | architecture | …}
  Packaging: {single file | split files | frames/zones}
  Constraints: {optional}
  Success: {After viewing, the viewer can …}

explanation: |
  {Script for explaining the diagram — what you would say while pointing at it.
   Walk through the story in viewing order. This is what the picture must make visible.}

# optional:
# output: "./{slug}.excalidraw"
```

**Forbidden at Stage 1:** `elements`, `relationships`, `diagram`, generating `.excalidraw`, sketching box inventories in the YAML (Express Mode still drafts `explanation` first, then may continue only after that content is coherent).

### 1.2 Approval gate (required unless Express Mode)

1. Point the user to the YAML path.  
2. Summarize in chat (user’s language): **`description` meta** briefly, then read or highlight the full **`explanation` script** — that is the main approval object.  
3. Ask them to **review and approve**, or **edit the YAML** (especially `explanation`) and say when ready.  
4. **Stop here** (default). Do **not** begin Stage 2 (and do **not** write structure/detail into the file) until approval — unless Express Mode is active.

**Fit test (for the agent):** Could a diagram exist such that speaking `explanation` while pointing at it feels natural and complete? If the script names concepts or steps the picture could not show, revise `explanation` (or scope in `description`) before the gate.

---

## Stage 2 — Structure Design

Turn the approved **`explanation`** into a **graph**: elements + relationships. Prefer the smallest set that still makes the script tellable. Proceed only from an **approved** Stage 1 YAML. Use `description` for audience/constraints; use **`explanation` as the coverage source of truth**.

### 2.1 Identify elements

Each element is a noun-like unit with a stable `id` and display `label`.

| Field | Meaning |
|-------|---------|
| `id` | Stable slug (`auth-service`, `user`) |
| `label` | Text shown on the diagram |
| `kind` | Optional: `actor`, `system`, `data`, `process`, `group`, `boundary`, `frame`, … |
| `parentId` | **Required** when this element is contained in another (see §2.1a). Parent’s `id`. |
| `notes` | Optional: why it exists for the explanation |

### 2.1a Containment model (mandatory when nesting applies)

Containment (`has-a` / `contains` / “inside”, VPC, cluster, module, package, swimlane cell, …) is a **first-class structure**, not a normal arrow.

| Rule | Detail |
|------|--------|
| **Detect** | From `explanation` and domain language: “in”, “inside”, “within”, “hosts”, “contains”, boundary names |
| **Encode** | Every contained element **must** set `parentId: <parent-id>` |
| **Parent kind** | Prefer `kind: boundary` or `group` for containers that exist mainly to hold others |
| **No peer layout** | Contained elements must **not** be treated as top-level peers of their parent in Stage 3 |
| **No containment arrow** | Do **not** draw an arrow whose only meaning is “contains”. Nesting **is** the visual |
| **Optional relation row** | You **may** keep `type: contains` / `has-a` in `relationships` for documentation, with `layout: nest` (or omit the row entirely). Never `layout: edge` for pure containment |
| **Multi-level** | Nested trees allowed (`pod` → `node` → `cluster`); each level has its own `parentId` |
| **Cycles forbidden** | `parentId` graph must be a forest (no cycles) |

**Sync check:** For every `relationships` entry with `type` in `{contains, has-a}` (and `layout` not forcing an edge), the **child** (`to` if parent→child, or the contained party) **must** have matching `parentId`. If `parentId` is set, Stage 3 **must** nest — even when no `contains` row exists.

### 2.2 Identify relationships

Edges between elements for **non-nesting** links. Name the relation type explicitly.

| Common types | Use |
|--------------|-----|
| `is-a` | Inheritance / specialization (usually an edge, not nesting) |
| `has-a` / `contains` | **Prefer `parentId` nesting**; relation row optional (`layout: nest`). Do not use as a decorative arrow |
| `uses` / `calls` | Dependency or invocation → **edge** |
| `flows-to` | Data or control flow → **edge** |
| `implements` | Contract vs realization → **edge** |
| `associates` | Loose link → **edge** |

| Field | Meaning |
|-------|---------|
| `id` | e.g. `rel_1` |
| `from` / `to` | Element ids |
| `type` | Relation kind |
| `label` | Optional edge label |
| `layout` | Optional: `nest` (containment → inside parent) \| `edge` (draw arrow). Default: `nest` for `contains`/`has-a`, else `edge` |

### 2.3 Structure validation (present clearly to the user)

Before asking for approval, run the checks below and **show the user a short validation report** in chat (pass/fail per item + one-sentence graph reading). Do not hide failures.

| # | Check | Pass means |
|---|--------|------------|
| 1 | **Explanation coverage** | Every beat/concept named in `explanation` maps to an element, a relation, or an *explicit* out-of-scope note |
| 2 | **No orphans** | No element exists that is not needed to deliver the `explanation` script |
| 3 | **Containment declared** | Every containment implied by the script has `parentId` on the child; every `contains`/`has-a` row is synced to `parentId` |
| 4 | **Containment integrity** | `parentId` targets exist; no cycles; containers that only group children are present as elements |
| 5 | **Script fidelity** | Reading elements + relationships aloud can deliver `explanation` without inventing missing boxes/edges |
| 6 | **Relation honesty** | Non-nest edges match real dependencies; pure containment is **not** modeled only as an arrow |
| 7 | **Description alignment** | Structure respects Goal / Audience / Scope / Constraints / Packaging in `description` |

If any check fails, revise `elements` / `relationships` and re-run the report — **do not** jump to layout.

### 2.4 Update YAML + approval gate (required)

**Append** `elements` and `relationships` to the existing Stage 1 YAML. Keep `title` / `description` / `explanation` unless the user asked to change them. **Do not** add `diagram` yet.

```yaml
# ... existing title, description, explanation ...

elements:
  - id: vpc
    label: VPC
    kind: boundary
  - id: api-gateway
    label: API Gateway
    kind: system
    parentId: vpc          # REQUIRED — drawn inside VPC
  - id: auth-service
    label: Auth Service
    kind: system
    parentId: vpc          # REQUIRED — drawn inside VPC

relationships:
  # Optional documentary containment (layout: nest → no arrow in Stage 3/4)
  - id: rel_contains_gw
    from: vpc
    to: api-gateway
    type: contains
    layout: nest
  - id: rel_contains_auth
    from: vpc
    to: auth-service
    type: contains
    layout: nest
  # Real edge between siblings (still inside parent)
  - id: rel_1
    from: api-gateway
    to: auth-service
    type: uses
    label: validates token
    layout: edge
```

**Forbidden at Stage 2:** `diagram` section, Excalidraw generation, inventing layout coordinates, representing containment **only** as an edge with no `parentId`.

Then:

1. Show the validation report (§2.3), including a **containment tree** (parent → children) for the user.  
2. Invite review of `elements` / `relationships` (user may edit the YAML).  
3. **Stop here** (unless Express Mode). Do **not** begin Stage 3 until the user approves the structure.

---

## Stage 3 — Detail Design

Model **layout** for Excalidraw: position, size, shape, nesting, and arrow geometry — still in YAML under `diagram`, not yet final JSON. Proceed only from an **approved** Stage 2 YAML.

### 3.1 Layout rules

- **Reading order**: Left→right or top→bottom matching the explanation (e.g. request flow L→R).  
- **Containment → inside the parent (non-negotiable)**:  
  - Every Stage 2 `parentId` becomes `parentElementId` on the child node (same id).  
  - Lay out the **parent box first**, then place all children **strictly inside** it.  
  - Parent `width`/`height` must encompass all children with at least **20px** padding on every side (extra top padding if the parent has a BoundText title).  
  - Parent `children: [...]` list should match all nodes with that `parentElementId`.  
  - **Do not** emit an arrow for `layout: nest` / pure `contains` / `has-a`.  
- **Sibling edges inside a parent**: `uses` / `flows-to` between children are allowed; keep arrow geometry between child edges, preferably still visually within the parent.  
- **Frames / zones**: For multi-zone stories in one file, add `frame` nodes (see Stage 4); place related nodes inside with padding (`frameId` at generation).  
- **Spacing (peers only)**: Align peers; **min gap 24px** between non-nested sibling bounding boxes (not parent↔child).  
- **Shape vocabulary**:  
  - Process / service → `rectangle`  
  - Actor / concept emphasis → `ellipse` (sparingly)  
  - Groups / boundaries → larger `rectangle` (container; title as BoundText)  
  - Named zones → `frame` when packaging says so  
- **Labels**: Short; one idea per box.  
- **Arrows**: Only for `layout: edge` relations; dock on shape **edges** (Stage 4).

### 3.1a Containment layout algorithm (agent must follow)

For each parent P with children C1…Cn:

1. Choose child sizes from labels (with text heuristics).  
2. Arrange children inside P (row, grid, or stack) with ≥24px between siblings and ≥20px from P’s inner border.  
3. Set  
   `P.width ≥ max(Ci.x + Ci.width) - P.x + 20`  
   `P.height ≥ max(Ci.y + Ci.height) - P.y + 20`  
   (and ≥ title band if needed).  
4. Verify each child: fully inside P (see §3.3 check 4).  
5. Recurse for nested parents (deepest children first when sizing upward, or size leaves then expand ancestors).

### 3.2 `diagram` section (design model)

Official nesting fields (required when Stage 2 has `parentId`):

| Field | Where | Meaning |
|-------|--------|---------|
| `parentElementId` | child `nodes[]` | **Must** equal Stage 2 `parentId` — child drawn inside this parent |
| `children` | parent `nodes[]` | Explicit list of child `elementId`s (keep in sync) |
| `shape: frame` | `nodes[]` | Zone frame for multi-diagram-in-one-file |

```yaml
diagram:
  canvas:
    viewBackgroundColor: "#ffffff"
  nodes:
    - elementId: vpc
      shape: rectangle
      x: 40
      y: 80
      width: 520
      height: 240
      children: [api-gateway, auth-service]   # required sync with parentElementId
    - elementId: api-gateway
      shape: rectangle
      parentElementId: vpc                   # inside VPC — not a peer
      x: 80
      y: 140
      width: 180
      height: 80
      style:
        strokeColor: "#1e1e1e"
        backgroundColor: "#a5d8ff"
        fillStyle: solid
        roughness: 1
    - elementId: auth-service
      shape: rectangle
      parentElementId: vpc
      x: 320
      y: 140
      width: 180
      height: 80
  edges:
    # only layout: edge relations — no contains arrow
    - relationshipId: rel_1
      fromNode: api-gateway
      toNode: auth-service
      endArrowhead: arrow
      label: validates token
  notes: []
```

### 3.3 Diagram validation (present clearly to the user)

Before asking for approval, validate that the layout **implements the approved structure** and **supports delivering the `explanation` script**. Show the user a short **layout validation report** (pass/fail + fixes). Do not proceed on silent assumptions.

| # | Check | Pass means |
|---|--------|------------|
| 1 | **Node completeness** | Every structure `element` has exactly one node (or a documented intentional omission) |
| 2 | **Position / reading order** | Node `x`/`y` follow the order implied by `explanation`; the script’s story is not scrambled |
| 3 | **Size adequacy** | `width`/`height` fit labels; parents large enough for **≥20px** inset around all children |
| 4 | **Containment geometry** | For every `parentElementId`: child fully inside parent — `px+20 ≤ cx`, `py+20 ≤ cy`, `cx+cw ≤ px+pw-20`, `cy+ch ≤ py+ph-20` (20px inset). Fail if child is only “near” the parent |
| 5 | **Containment sync** | Every Stage 2 `parentId` has matching `parentElementId`; every parent’s `children` lists those ids; no nest relation appears as an `edges[]` row |
| 6 | **Edge completeness** | Every `layout: edge` relationship has an edge; `fromNode`/`toNode` match |
| 7 | **Edge intent** | Arrow direction/label match meaning; **no** arrow for pure containment |
| 8 | **Collision / spacing** | Peer pairs (neither is ancestor of the other): overlap area **0** and gap ≥ **24px** (formula below). Parent↔child pairs **must** overlap (child inside) — do not apply peer gap rules to them |
| 9 | **Script walkthrough** | Pointing through the layout, one can speak `explanation` with containment visible as nesting |

**Check 8 — peer bounding-box test** (skip if A is ancestor of B or B of A via `parentElementId` chains):

```
overlap_x = max(0, min(Ax+Aw, Bx+Bw) - max(Ax, Bx))
overlap_y = max(0, min(Ay+Ah, By+Bh) - max(Ay, By))
overlap_area = overlap_x * overlap_y   # must be 0 for peers

gap_x = max(0, max(Ax, Bx) - min(Ax+Aw, Bx+Bw))
gap_y = max(0, max(Ay, By) - min(Ay+Ah, By+Bh))
# When x-ranges or y-ranges overlap (side-by-side / stacked), require the clear gap on the separating axis ≥ 24.
```

If any check fails, fix `diagram` (or roll back to Stage 2 if the graph itself is wrong) and re-run the report.

### 3.4 Update YAML + approval gate (required)

**Append** the `diagram` section to the existing Stage 2 YAML. Do not remove approved structure unless rolling back. **Do not** write the `.excalidraw` file yet.

Save the updated YAML, present the §3.3 report, and wait for **user approval** (or their direct YAML edits + confirm) — unless Express Mode.

**Stop here** (default). **Do not** enter Stage 4 until detail design is approved.

---

## Stage 4 — Generate Excalidraw JSON

Proceed only after **approved** detail design. Emit `{diagram-slug}.excalidraw` (or `.json`) **in the same folder as the YAML** by default (see Working Location). Output **only valid JSON** matching the minimal schema below.

### 4.1 Document envelope

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

### 4.2 Allowed element types (minimal schema)

**Rectangle / Ellipse**

- Required: `id` (string), `type` (`"rectangle"` | `"ellipse"`), `x`, `y`, `width`, `height` (numbers)  
- Optional: `strokeColor`, `backgroundColor` (hex), `fillStyle` (`"solid"` | `"hachure"` | `"cross-hatch"`), `roughness` (`0` | `1` | `2`), `boundElements` (`[{ "type": "text", "id": "text_id" }]`)  
- Safe defaults: `angle: 0`, `strokeWidth: 2`, `strokeStyle: "solid"`, `opacity: 100`, `groupIds: []`, `roundness: null` or `{ "type": 3 }`, `seed` (int), `version: 1`, `versionNonce` (int), `isDeleted: false`, `updated` (ms), `locked: false`, `link: null`, `frameId: null` (or parent frame id)

**Frame** (zones / multi-diagram-in-one-file)

- Required: `id`, `type: "frame"`, `x`, `y`, `width`, `height`, `name` (zone title)  
- Child shapes set `frameId` to this frame’s `id`  
- Size the frame so all children fit with ≥20px padding

**Text**

- Required: `id`, `type: "text"`, `x`, `y`, `width`, `height`, `text`, `fontSize`, `fontFamily` (`1` | `2` | `3`)  
- Also set `originalText` equal to `text`

**BoundText trinity** (all three required or the label drifts when the shape moves):

1. Shape: `boundElements: [{ "type": "text", "id": "text_id" }]`  
2. Text: `containerId: "shape_id"`  
3. Text: `textAlign: "center"`, `verticalAlign: "middle"`  

Plus: place text `(x, y)` / size so it sits in the shape’s content box (centered). Never use a fake `label` property on the shape.

**Text width / height (auto-size estimate)**

| Case | Width heuristic | Height heuristic |
|------|-----------------|------------------|
| Latin / single line | `fontSize * charCount * 0.6` | `fontSize * 1.5` |
| CJK / Korean / wide glyphs | `fontSize * charCount * 1.0` (or measure longest CJK run) | `fontSize * 1.6` |
| Multiline (`\n`) | Use **longest line**’s char width (same coeff as above) | `fontSize * 1.5 * lineCount` |
| Padding | Add **10–15px** to both width and height after the heuristic |

Prefer slightly **too wide** over clipping.

**Arrow**

- Required: `id`, `type: "arrow"`, `x`, `y`, `points` (e.g. `[[0,0],[dx,dy]]`), `endArrowhead: "arrow"` (or `null` / `"triangle"`)  
- Connections: `startBinding` / `endBinding`: `{ "elementId": "…", "focus": 0, "gap": 1 }`  
- `points` are **relative** to the arrow’s `x,y` (start at `[0,0]`, end at delta)

**Arrow edge docking (required)**  
`startBinding` / `endBinding` alone are not enough. Compute geometry so the arrow **touches shape borders**:

1. Arrow absolute start `(x, y)` lies on the **boundary edge** of the `from` shape (not floating, not deep inside).  
2. Absolute end `(x + dx, y + dy)` lies on the **boundary edge** of the `to` shape.  
3. Relative `points` are `[[0,0], [dx, dy]]` with that end delta.  
4. Prefer midpoints of facing sides (e.g. right edge of A → left edge of B for L→R flow).

**Containment rendering (required)**

1. Emit **parent shapes before children** in `elements[]` so children paint above the container.  
2. Child rectangles/ellipses use absolute canvas coordinates that lie **inside** the parent (same as Stage 3 YAML).  
3. **Do not** create arrows for `layout: nest` / pure `contains` / `has-a`.  
4. If the parent is an Excalidraw `frame`, set each child’s `frameId` to the frame id; for ordinary boundary rectangles, nesting is geometric only (coordinates inside parent) — still no containment arrow.  
5. Parent BoundText title should not overlap children (reserve a title band at the top inside the padding).

### 4.3 Critical rules (hallucination guardrails)

1. **Never** put labels on shapes as `"label"` or `"title"`. Use the **BoundText trinity**.  
2. **Do not invent** properties outside this schema.  
3. Unique ids: `rect_1`, `text_1`, `arrow_1`, `frame_1`, …  
4. Arrow `points` use relative offsets, not absolute canvas coordinates.  
5. Size text with the CJK / multiline heuristics + 10–15px padding; do not rely on Latin-only `0.6` for Korean.  
6. Update both sides of bindings when connecting arrows.  
7. **Arrow point calculation**: Arrow `x,y` MUST start on the boundary edge of the `from` shape. Relative `points[[0,0],[dx,dy]]` MUST terminate exactly on the boundary edge of the `to` shape.  
8. **Containment**: Every `parentElementId` child is geometrically inside its parent; never replace nesting with a “contains” arrow; emit parents before children in the JSON array.  
9. Nested children set `frameId` when inside a `frame`; keep ids consistent with Stage 3 YAML.

### 4.4 Generation checklist

- [ ] Every structure `element` has a node (or intentional omission)  
- [ ] Every `parentId` → nested geometry inside parent (not a peer + arrow)  
- [ ] Every `layout: edge` relationship has an arrow; **no** arrow for containment  
- [ ] Parents appear before children in `elements[]`  
- [ ] Every caption uses BoundText trinity  
- [ ] Every arrow docks on from/to **edges**  
- [ ] Text boxes sized for CJK/multiline + padding  
- [ ] Peer collision check still passes; parent↔child correctly overlap (child inside)  
- [ ] JSON parses; opens in Excalidraw without repair  
- [ ] Speaking `explanation` while reading the canvas still works  

Also refresh YAML `diagram` if coordinates changed during generation so YAML remains the source of truth.

---

## Stage 5 — Revise

Classify each change request, then edit at the **owning stage** (bi-directional rollback). Re-save YAML and **re-run that stage’s approval gate** before regenerating when needed.

| Kind | Examples | Action |
|------|----------|--------|
| **Requirements** | Wrong topic, bad script, missing audience | Update `description` and/or `explanation`; re-approve Stage 1; then structure → detail → JSON |
| **Structure** | Add/remove box, change relation type | Update `elements` / `relationships`; re-run §2.3 report + approve; then detail → JSON |
| **Detail / visual** | Move, resize, recolor, straighten arrow | Update `diagram`; re-run §3.3 report + approve; then JSON |
| **Generation-only** | Invalid Excalidraw / binding bug | Fix JSON (and sync YAML coords if layout changed); no meaning change |

Never “fix layout” by deleting explanatory edges. Prefer smallest edit that satisfies the ask.

---

## YAML File Spec (grows by stage)

One file under `diagrams/` (unless the user chose another folder). Examples below show **what the file should look like after each stage** — never jump to the final shape during Stage 1.

### After Stage 1 (requirements only)

```yaml
title: Auth request path

description: |
  Goal: Show how a client call reaches Auth Service via the API Gateway for token validation.
  Audience: Backend onboarding.
  Scope: Happy-path request only; omit refresh tokens and admin APIs.
  Diagram kind: architecture / request flow (left to right).
  Constraints: Max ~5 boxes; HTTPS label on client→gateway.
  Success: After viewing, the viewer can name Client → API Gateway → Auth Service and say why the gateway calls Auth.

explanation: |
  Start at the Client on the left: it sends an HTTPS request to the API Gateway.
  The gateway does not decide authenticity alone — it uses the Auth Service to validate the token.
  Point along the arrow from Gateway to Auth Service labeled "validates token," then note that
  only after that check would the request continue deeper into the system (out of scope here).

output: "./auth-overview.excalidraw"
```

### After Stage 2 (structure added)

Same as Stage 1, **plus** (note `parentId` for containment — not a contains-only arrow):

```yaml
elements:
  - id: client
    label: Client
    kind: actor
  - id: vpc
    label: VPC
    kind: boundary
  - id: api-gateway
    label: API Gateway
    kind: system
    parentId: vpc
  - id: auth-service
    label: Auth Service
    kind: system
    parentId: vpc

relationships:
  - id: rel_0
    from: client
    to: api-gateway
    type: flows-to
    label: HTTPS
    layout: edge
  - id: rel_1
    from: api-gateway
    to: auth-service
    type: uses
    label: validates token
    layout: edge
```

Still **no** `diagram:` key. Contained services are nested via `parentId`, not via containment arrows.

### After Stage 3 (detail design added)

Same as Stage 2, **plus** nested geometry (`parentElementId` / `children` — children inside parent):

```yaml
diagram:
  canvas:
    viewBackgroundColor: "#ffffff"
  nodes:
    - elementId: client
      shape: ellipse
      x: 40
      y: 160
      width: 120
      height: 80
    - elementId: vpc
      shape: rectangle
      x: 200
      y: 80
      width: 480
      height: 240
      children: [api-gateway, auth-service]
    - elementId: api-gateway
      shape: rectangle
      parentElementId: vpc
      x: 240
      y: 160
      width: 180
      height: 80
    - elementId: auth-service
      shape: rectangle
      parentElementId: vpc
      x: 460
      y: 160
      width: 180
      height: 80
  edges:
    - relationshipId: rel_0
      fromNode: client
      toNode: api-gateway
      endArrowhead: arrow
      label: HTTPS
    - relationshipId: rel_1
      fromNode: api-gateway
      toNode: auth-service
      endArrowhead: arrow
      label: validates token
```

### After Stage 4

Emit `diagrams/auth-overview.excalidraw` (colocated unless `output` overridden). The picture must make the **`explanation` script** easy to deliver.

---

## Choosing the Optimal Diagram Form

| If the user needs to explain… | Prefer |
|-------------------------------|--------|
| Steps over time / pipeline | Left-to-right **flow** |
| System parts and calls | **Architecture** boxes + `uses` / `flows-to` |
| Taxonomy / inheritance | Top-down **hierarchy** (`is-a`) |
| Ownership / modules | Nested **containers** (`parentId` / `parentElementId`) |
| Data entities | Compact **ER-style** boxes + relation labels |
| “A then B then C” with actors | Swimlane-like stacked rows (rect groups) |
| Several loosely related stories | **Split files** *or* one file with multiple **`frame` zones** |

Pick **one** primary form per frame/file. Mixing ER + timeline + deployment on one undivided canvas usually fails the explanation test.

**Packaging rule (Stage 1):** If two explanations do not share a single walkthrough, prefer two YAML/Excalidraw pairs. If they share a board but need visual separation (e.g. “Control plane” vs “Data plane”), use Excalidraw **`frame`** elements as zones inside one file.

---

## Prohibitions

- Writing a **full** YAML (with `elements` / `relationships` / `diagram`) during Stage 1  
- Adding `diagram` during Stage 2, or emitting `.excalidraw` before Stage 4 approval (except Express Mode after a coherent full YAML)  
- Continuing to the next stage in the same turn without user approval (unless Express Mode or explicit waiver)  
- Emitting Excalidraw JSON before **detail design is finalized** (Express: finalize in-file before emit)  
- Skipping Stage 1–3 approval gates (unless Express Mode or explicit waiver)  
- Hiding §2.3 / §3.3 validation failures from the user  
- Inventing Excalidraw properties (`label` on rectangle, fake types)  
- Embedding caption text inside shape objects; incomplete BoundText trinity  
- Absolute coordinates inside arrow `points`  
- Arrow starts/ends floating in empty space or buried deep inside shapes  
- Latin-only text width math for CJK labels  
- Orphan decorative shapes that do not serve `explanation`
- Treating `description` alone as enough Stage 1 content (missing or empty `explanation` script)
- Drawing a picture that cannot support speaking the approved `explanation`  
- Representing containment **only** as an arrow / peer boxes outside the parent  
- Missing `parentId` / `parentElementId` when the script or `contains`/`has-a` implies nesting  
- Applying peer “no overlap” rules to parent↔child pairs (children **must** sit inside)  
- Silent scope changes during a “visual-only” revision  
- Giant undivided one-canvas dumps when frames or separate files would explain better
