# Chapter 001 Continuity Audit — Sequential Republication

## Scope and Authority

- Audit target: `chapters/001-죽은-천마-망나니로-눈을-뜨다.md` and `manuscripts/001-죽은-천마-망나니로-눈을-뜨다.md`.
- Primary truth: Chapter 001 design, the approved architecture profiles, `overview.md`, `series.md`, and Part 001.
- Historical reference only: `continuity/001~005-summary.md`. Later Chapter 002~005 events were not used to rewrite Chapter 001 truth.
- Canonical `continuity/story-so-far.md` and `TODO.md` were intentionally not modified.

## Findings

### F-01 — Design closing instructions conflicted (Medium, design)

The chapter-level closing image specified the window-side declaration, while Episode 003 specified no dialogue and a transition back to the remaining medicine. The manuscript reproduced both, creating a double ending and weakening the Chapter 002 handoff.

**Adjudication:** Apply. This is a chapter-design consistency defect, so design was corrected first. The closing image is now the sickroom window-to-medicine return; the unresolved poisoning question remains the forward hook.

### F-02 — Manuscript used a non-local place name (Medium, manuscript)

The scene is in Luoyang and the architecture names Luoyang as the visible city, but the manuscript referred to “장안의 지붕들.” This is an unsupported location substitution and can create a false geography for Chapter 002.

**Adjudication:** Apply. Replaced with “낙양의 지붕들.” No architecture change was necessary.

### F-03 — Manuscript turned a designed image-only resolve into spoken thematic lines (Medium, manuscript)

Episode 003’s design marked the resolve as `Dialogue intent: none`, and the chapter Hold prohibits overt propositions about the protagonist’s ideology. The manuscript added three separated spoken lines, including “정파를, 꼭지까지 쥐어보지,” before the final poison hook. This made the ending explanatory and duplicated the chapter’s intended reader-discovered meaning.

**Adjudication:** Apply. The design was clarified first: the resolve is inner narration/action and the closing image is the medicine bowl. The manuscript now keeps the resolve in close narration and ends on the designed poison question.

## Audit Matrix

| Area | Verdict | Evidence / inheritance rule |
|---|---|---|
| Plot beats | Pass after F-01/F-03 | Awakening, dual death memory, body diagnosis, poison intent, Seolhyang bond, second-brother clue, resolve, and Chapter 002 hook remain in order. |
| Cast | Pass | Only Namgung Taeui/Hyeok Muryun and Seolhyang appear; both are architectural cast and match their profiles. |
| Timing | Pass | One dawn-to-morning sickroom sequence; no later-chapter events were pulled backward. |
| World/lore | Pass after prose correction | Reincarnation synchronization, damaged danjeon, poison knowledge, and the held Golhyeoljaejo-gong reference stay within the world bible. No poison name, sourcing, or later faction lore is revealed. |
| Locations/staging | Pass after F-02 | All scenes use `namgung-manor/chwuiun-gak/interior` and `p1-taeui-sickroom`; no blocking drift. |
| Seeds/Holds | Pass after F-03 | Poison intent and Seolhyang’s loyalty are planted; Namgung Hwi’s guilt, poison name/source, repair execution, later family confrontations, and assassin plot remain held. |
| Opening hook | Pass | The chapter opens on the overlapping deaths and leaves the reincarnation/body question active. |
| Closing hook | Pass after F-01/F-03 | Remaining medicine and the conclusion that someone intends to kill him hand directly to Chapter 002’s poison investigation. |
| Chapter 002 inheritance | Pass | It must inherit: Taeui’s identity/state, sickroom staging, continuing poison suspicion without confirmation, Seolhyang’s ignorance and loyalty, and the unresolved “who/why” question. |

## Continuity Adjudication

Chapter 001 is continuity-safe after the applied revisions. The later 001~005 summary confirms the same broad direction, but its Chapter 002~005 facts are not promoted into Chapter 001 state. In particular, Chapter 001 does **not** establish the name 혼백산, a confirmed Namgung Hwi crime, Kim Chung, Golhyeoljaejo-gong execution, or any later family reaction.

## Changes Applied

1. Revised the Chapter 001 design’s chapter closing image and Episode 003 sensory/transition instructions to remove the internal contradiction and explicitly require an image-only resolve.
2. Corrected “장안의 지붕들” to “낙양의 지붕들.”
3. Reworked the manuscript’s spoken closing resolve into close narration and preserved the final medicine/poison hook.
4. Created the Chapter 001 audit-part continuity summary, story-so-far snapshot, and unresolved-thread register.

## Verdict

**PASS — continuity-safe for sequential republication and ready to hand to Chapter 002.**

**Findings:** 3 (2 applied design-first/prose cascade findings, with F-02 prose-only; no rejected findings).

**Changed paths:** 5 total: 1 chapter design, 1 manuscript, 1 audit report, 3 continuity audit-part files created. (`5` existing/new path entries if the report itself is excluded; `6` filesystem paths including this report.)

## Release — Sequential Audit Ledger
- **Released:** ✅
- **Date:** 2026-08-02
- **Ledger:** `continuity/audit-part1/001-summary.md`, `story-so-far.md`, `unresolved-threads.md`
- **Next chapter obligation:** Chapter 002 must begin from the unconfirmed poison suspicion and the remaining medicine at the sickroom.
