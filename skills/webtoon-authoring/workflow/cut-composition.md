# Cut Composition Guide

**Not a numbered stage.** Read when designing or evaluating episode cuts (Stage ② / ③), and when sizing page images (Stage ④).

Cuts are not mere “picture boxes.” **Cut height and gutter distance control the reader’s eye speed, tension, and breathing.**

---

## 1. Canvas baseline (before dividing cuts)

| Spec | Guideline | Notes |
|------|-----------|--------|
| Canvas width | **690px–800px** | Common portal range; lock exact width in `overview.md`. Some portals also accept ~1080px — only if the user/portal requires it |
| Cut horizontal fit | **Full bleed** or **30–50px side margins** | Fill the strip width, or keep left/right margins even |
| Cuts per episode (weekly-serial average) | **~60–80 cuts** | Scale for short pilots; lock a target range in `overview.md` / `series.md` |

Generation default aspect for a phone-width strip remains **9:21** (~672×1584) when producing page segments; individual cuts inside a page still follow the height classes below.

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

## 5. Rhythm checklist (per scene / episode stretch)

- [ ] Size class matches dramatic intent (not all cuts the same height)
- [ ] Climax uses tall/long cuts where overwhelm is wanted
- [ ] Silence/time uses `pause` gutters, not empty “forgotten” gaps
- [ ] Action uses open/diagonal + tight gutters
- [ ] Everyday dialogue stays in standard height + normal gutters
- [ ] Episode cut count is in the agreed range (or short-form exception is noted)
- [ ] Side margins are full-bleed or consistently 30–50px
