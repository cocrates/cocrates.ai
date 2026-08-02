# Cut Composition Guide

**Not a numbered stage.** Read when designing or evaluating episode cuts (Stage ② / ③), and when sizing page images (Stage ④).

Cuts are not mere “picture boxes.” **Cut height and gutter distance control the reader’s eye speed, tension, and breathing.**

---

## 1. Canvas baseline (before dividing cuts)

| Spec | Guideline | Notes |
|------|-----------|--------|
| Canvas width | **768px** default (690–800 common) | Lock exact width in `overview.md`. Some portals also accept ~1080px — only if the user/portal requires it |
| Cut horizontal fit | **Full bleed** or **30–50px side margins** | Fill the strip width, or keep left/right margins even |
| Cuts per episode (weekly-serial average) | **Emergent from ~10–15 page design** | Do not budget cuts in `series.md`; pack by dwell (`§4b`) |

Generation default for phone-width page strips: **`gemini-3.1-flash-image` + 1K + 9:16**, typically **2 tiles per design page**, then stitch (`04-generate.md`). Optional single tall frame: **Pro + 2K + 1:8** only. Individual cuts inside a page still follow the height classes below.

---

## 2. Directing with cut height + gutter

Vertical scroll means **cut vertical length** and **gutter (space between cuts)** are the primary directing tools.

```text
1) Everyday / dialogue          2) Climax / immersion         3) Speed / action
   ┌──────────┐                    ┌──────────┐                  ┌──────────┐
   │  Cut 1   │                    │          │                  │ Cut1 /  /│
   └──────────┘                    │  Cut 1   │                  │    /  / │
    (normal gutter)                │ (tall)   │                  └─/──/────┘
   ┌──────────┐                    │          │                   / Cut2 /
   │  Cut 2   │                    └──────────┘                  /─────/
   └──────────┘
```

### A. Dialogue / everyday (standard cut)

| Spec | Guideline |
|------|-----------|
| Cut height | **~400–600px** (about half a phone screen) |
| Gutter after | **~150–300px** |
| Use when | Calm rhythm; one cut + balloons fit naturally in one viewport beat |

### B. Climax / spectacle (one-cut / long cut)

| Spec | Guideline |
|------|-----------|
| Cut height | **~1,200–2,000px+** (taller than one phone screen) |
| Use when | Grand establishing shots, full-body reveals, peak emotion (tears, shock). Reader scrolls *through* the cut for overwhelm |

### C. Tension / time passing (gutter directing)

| Spec | Guideline |
|------|-----------|
| Gutter after | **~500–1,000px+** (intentional empty vertical space) |
| Use when | Silence, stillness, time skip, emotional aftertaste — emptiness *is* the beat |

### D. Action / fast breath (diagonal / open cut)

| Spec | Guideline |
|------|-----------|
| Cut shape | Diagonal frames and/or **open cuts** (no box border; art/effects break out of the frame) |
| Gutter after | **Short** (tight stacking) |
| Use when | Speed, impact, chase — break the rectangle; keep gaps small |

---

## 3. Cut purpose → size class (pick one per cut)

Assign every cut a **size class** in episode design:

| Size class | Height (guide) | Typical gutter after | Intent |
|------------|----------------|----------------------|--------|
| `standard` | 400–600px | 150–300px | Dialogue, daily life, exposition |
| `tall` / `long` | 1,200–2,000px+ | normal or wide | Climax, spectacle, scroll-through reveal |
| `compact` | <400px (sparingly) | short | Quick reaction, insert beat |
| `open` / `diagonal` | as needed | short | Action, speed, break-frame energy |

Gutter overrides by intent (may exceed the class default):

| Gutter class | Distance (guide) | Intent |
|--------------|------------------|--------|
| `tight` | <150px | Fast rhythm, action chain |
| `normal` | 150–300px | Default dialogue/everyday |
| `wide` | 300–500px | Soft scene shift |
| `pause` | 500–1,000px+ | Silence / time / aftertaste |

---

## 4. Mobile viewport discipline

When designing or evaluating a cut sequence, always ask:

1. **What fits in one phone viewport** (roughly 16:9 or 19.5:9 height at canvas width)?
2. Is that viewport showing **one clear beat** (or an intentional partial reveal of a tall cut)?
3. Do balloons + face/action remain readable without pinch-zoom?

Do not stack so many standard cuts that a single screen becomes a crowded strip of unreadably small figures.

---

## 4b. Page packing by dwell time (reader breath)

A **page** is not “as many cuts as fit.” It is a **breath unit**: how long the reader should stay in this vertical segment before the next page/scroll chunk.

### Decide cuts-per-page from dwell need

| Dwell need | Typical packing | Why |
|------------|-----------------|-----|
| **High** — dense dialogue/balloons; first look at a new place/cast; reader must think or feel; spectacle scroll-through | **1 cut per page** (sometimes 1 tall cut alone) | Gives eye time; avoids competing beats in one viewport chain |
| **Medium** — clear beat + short reaction, or two linked dialogue beats | **2–3 cuts** stacked top→bottom | One breath, still readable |
| **Low** — quick reactions, punchline chain, action stutter | **n cuts** (tight gutters) | Speed; still keep faces/balloons readable per viewport |

**Default stack direction:** top → bottom only (no side-by-side panels unless the series explicitly locks a rare exception).

### Heuristics (apply in Stage ② when grouping cuts into pages)

1. **Estimate dwell** per cut: skim / read balloons / absorb new info / feel / scroll-through tall.
2. **High-dwell cut → own page** (or own page + only a tiny reaction cut if the hook needs it).
3. **Do not pack** a first-time establishing + heavy dialogue + emotional turn into the same page.
4. **Page estimated height** ≈ sum of cut heights + gutters; if that exceeds one comfortable generation frame, prefer **fewer cuts per page** over cramming (see Stage ④ aspect guidance in `04-generate.md`).
5. **Viewport plan** must name the intended breath: e.g. `1 cut — linger`, `2 cuts — quick exchange`, `tall alone — scroll-through`.

Breath > filling the frame. Empty vertical space (`pause` gutter) or a single-cut page is valid directing.

---

## 5. Rhythm checklist (per scene / episode stretch)

- [ ] Size class matches dramatic intent (not all cuts the same height)
- [ ] Climax uses tall/long cuts where overwhelm is wanted
- [ ] Silence/time uses `pause` gutters, not empty “forgotten” gaps
- [ ] Action uses open/diagonal + tight gutters
- [ ] Everyday dialogue stays in standard height + normal gutters
- [ ] **Cuts-per-page matches dwell/breath** (high-dwell → 1 cut/page allowed; not max-packing by default)
- [ ] Episode cut count is recorded in Craft Notes (measured) — not managed in `series.md`
- [ ] Episode page volume is ~10–15 (or short-form exception noted)
- [ ] Side margins are full-bleed or consistently 30–50px
