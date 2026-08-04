# Chapter 002 Continuity Audit: 보약이라는 이름의 독

## Audit Scope

- **Audit mode:** Sequential republication audit, Chapter 002
- **Authority used:** `continuity/audit-part1/001-summary.md`, `continuity/audit-part1/story-so-far.md`, and `continuity/audit-part1/unresolved-threads.md` after Chapter 001
- **Explicit exclusion:** canonical `continuity/story-so-far.md` and canonical `continuity/unresolved-threads.md` were not used as authority
- **Reviewed:** Chapter 002 design, manuscript, architecture references, and workflows 04/05/06/07/08 plus `workflow/consistency.md`

## Handoff Checks

| Handoff | Initial Chapter 001 audit state | Chapter 002 result | Verdict |
|---|---|---|---|
| Sickroom state/staging | Same 취운각 sickroom; medicine table and remaining decoction are the next investigation; Taeui remains weak | Opens at the same medicine table, preserves the weak body, and adds Hwi as a temporary visitor to the locked staging | PASS after revision |
| Remaining medicine | Remaining decoction is the direct investigation object | Explicitly carried as the same medicine-table residue; new doses are distinguishable from the remaining residue | PASS after revision |
| Seolhyang knowledge | Loyal and uninformed about poison and perpetrator | She brings medicine, remains worried, and learns nothing about poison or Taeui's plan | PASS after revision |
| Namgung Hwi suspicion | Named as a clue only; perpetrator and motive unconfirmed | His visit and memory probe deepen suspicion; no proof, confession, or actual dosing scene is added | PASS after revision |
| Poison certainty level | Poison intent strongly suspected; name unknown | The medicine is narrowed to likely 혼백산 from evidence; culprit, procurement, and dosing remain unresolved | PASS after revision |
| Timing | Immediately after Chapter 001, same sickroom sequence | Morning investigation, same-morning visit, noon dose, evening concealment; no unexplained jump | PASS |
| Seeds/Holds | T-001 through T-005 active; restoration not executed | Poison identity advances, restoration remains a Hint for 003, procurement and actual dosing remain Hold | PASS after revision |
| Chapter 002 to 003 hook | Need to move from poison recognition to damaged-danjeon restoration | Taeui chooses the death-performance and calculates time to rebuild the body; 003 can begin restoration | PASS after revision |

## Findings Before Revision

Count: **6 findings**

| # | Severity | Finding | Layer | Resolution |
|---|---|---|---|---|
| 1 | High | The design cited the Chapter 001 design as its continuity authority instead of the audit-part1 handoff requested for this republication audit | Design/process | Replaced with the three audit-part1 continuity files and recorded the certainty boundary |
| 2 | High | Several design and manuscript lines collapsed strong poison-intent evidence into certainty that the death was deliberately designed by the suspected family member | Design + manuscript | Poison is identified as likely 혼백산; perpetrator and procurement remain unresolved |
| 3 | High | Hwi's visit was written as confirmation that he was the monitoring perpetrator, exceeding the handoff's clue-only status and violating the Hold on actual dosing | Design + manuscript | Visit now produces a stronger suspicion only; no dosing or confession is shown |
| 4 | Medium | Hwi appeared in the p1 sickroom staging although the staging cast lock did not include him | Architecture/staging | Added Hwi's temporary entry and blocking to `p1-taeui-sickroom.md` |
| 5 | Medium | The manuscript treated Seolhyang as if she could be implicated by the medicine and used “this hand made it” certainty, despite her uninformed state | Manuscript | Revised to show her as an uninformed carrier and Hwi as a possible connection, not a proven maker |
| 6 | Medium | The exact “half-year” timing was presented as settled inherited fact while Chapter 001 audit continuity marked survival time as undetermined | Design + manuscript | Retained the architecture-supported estimate but qualified it as “under the current dosage, around half a year” |

## Design-Fidelity Recheck

| Check | Result | Evidence |
|---|---|---|
| Every design scene remains in manuscript order | PASS | Ep001 S1-S2, Ep002 S3-S4, Ep003 S5-S6 preserved |
| Prior hook addressed | PASS | Same sickroom and remaining medicine open the chapter |
| Sickroom staging and cast | PASS | Hwi added to the staging lock as a temporary visitor; Taeui/Seolhyang positions remain unchanged |
| Situation -> Beat -> Turn | PASS | Poison investigation, suspicious visit, and active concealment all retain their designed turns |
| Seeds/Hold fidelity | PASS | Restoration is only hinted; procurement, dosing, and culprit proof remain absent |
| Continuity states | PASS | Taeui remains weak and poisoned; Seolhyang remains uninformed; Hwi remains unproven |
| Chapter 003 hook | PASS | Death-performance becomes the cover for rebuilding the body |
| Manuscript length | PASS | 4,5xx characters; verified after revision with `wc -m` |

## Verdict

**REPUBLICATION AUDIT PASS.** The six findings were resolved through design-first revision, followed by manuscript alignment. Chapter 002 now hands off a named-poison investigation, a stronger but unresolved Hwi suspicion, an uninformed Seolhyang, and an active restoration obligation to Chapter 003 without importing later canonical continuity.

## Release

- **Audit package released:** ✅
- **Canonical Chapter 002 release status:** unchanged by this audit
- **Audit date:** 2026-08-02
- **Continuity scope:** `continuity/audit-part1/` only
- **Verdict:** PASS
- **Finding count:** 6 identified, 6 resolved
- **Changed paths:**
  - `chapters/002-보약이라는-이름의-독.md`
  - `manuscripts/002-보약이라는-이름의-독.md`
  - `stagings/p1-taeui-sickroom.md`
  - `evaluations/002-continuity-audit.md`
  - `continuity/audit-part1/002-summary.md`
  - `continuity/audit-part1/story-so-far.md`
  - `continuity/audit-part1/unresolved-threads.md`
- **Protected paths:** `continuity/story-so-far.md`, `TODO.md`
