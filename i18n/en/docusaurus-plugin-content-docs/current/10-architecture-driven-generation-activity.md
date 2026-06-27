---
sidebar_position: 10
---
# EP10. Architecture-Driven Generation Principles

![Cocrates Generation Principles](/img/10-architecture-driven-generation-activity.png)

---

In Ep8 and Ep9 we explored Learning Pipeline philosophy and skills. Education → Knowledge Capture → Reflection—question, record ignorance, verify as interviewer.

Now Cocrates' second core activity: **Artifact Generation Pipeline**—architecture-driven artifact creation.

Remember: **AI does the work.** You are the **AI-native engineer (team lead)** who reviews and decides on what AI produces. In the age of AI doing everything for you, your value is how well you **review, judge, and decide** on AI's work.

If Learning was "how to recognize ignorance and grow knowledge," Generation is **"how you review and approve AI's proposed structure, then build from that structure."**

---

## 🚨 The Trap of "Just Write It"

You asked AI: *"Write a report."* Thirty pages in five seconds. You read it—thin logic, uneven depth per section. Hard to ask for a redo; you patch awkwardly and lose time.

What's the cause?

**Generation without structure.**

Unstructured output has three problems:

**First, weak consistency and logic.** No single thread—AI lists what sounds plausible moment to moment. Depth varies; claims don't align.

**Second, you can't understand or explain the deliverable.** *"Looks fine"* and move on—but *"Why is this section structured this way?"* stumps you. AI decided the structure, not you.

**Third, an unreviewable black box.** Review needs criteria. No structure? Only *"does it look plausible?"*

That's *"just write it"*—effectively **"write it any which way."**

---

## 🏗️ ASR — What Is the Fine China?

*"Okay—I need to design structure first."*

But "structure" is abstract—like *"live well"* without concrete steps.

Enter **ASR (Architecturally Significant Requirement)**—requirements or design decisions that **materially affect structure, composition, and quality** of the final deliverable. From software architecture, but applies to **every deliverable type**—documents, slides, blog series.

ASR matters because without reviewing ASR, AI fills gaps with **Silent Defaults**—*"reasonably okay"* choices that may not match your intent, and you may not even know they were chosen.

---

## 🏠 Country House Analogy — Tragedy of Silent Defaults

You're moving. Boxes to pack.

You tell the mover **"pack it well."** They pack their way. Fine china under heavy books.

No bad intent—just their default vs your idea of "well packed."

That's **Silent Default**. If you don't specify, AI fills with its defaults—and they may differ from your intent.

Structure is deciding what goes where and how. **ASR is identifying what's the fine china.**

Imagine building a country house.

You tell the architect: *"Design the house that best fits our family's lifestyle."*

Decisions: one story or two? Roof, rooftop terrace, attic room?

Each is a **Concern** and an **ASR**—real impact on structure and usability.

**Case 1 — User decides clearly:**
> *"Build a second-floor attic room."*
Decision made. No ADR needed—record in Spec.

**Case 2 — User delegates:**
> *"Choose what best fits our lifestyle."*
**ADR needed.** Architect compares options; user reviews and decides.

---

## 🔄 Four-Stage Pipeline

In Cocrates' Artifact Generation Pipeline:

```
ASR identification → ADR → Spec → Generation & Verification
```

**Stage 1 — ASR identification:** AI surfaces ASRs. Your job: **discernment**—which truly affect structure vs noise. Without this, AI treats everything as important or decides everything alone.
> AI-native skill: **discernment** between important ASRs and lesser requirements

**Stage 2 — ADR (alternatives & decision):** AI analyzes options per Concern. Your job: **judgment**—is this analysis enough? Other options? What did AI miss? ADR audits *why you decided this way*.
> AI-native skill: **judgment** to evaluate AI analysis and decide optimally

**Stage 3 — Spec (integrate decisions):** AI merges approved ADRs into one document. Your job: **insight**—anything missing? What will the final output look like from this Spec? Gaps send you back to ADR. Spec must be **self-contained**—one read explains everything.
> AI-native skill: **insight** to verify Spec completeness and anticipate output

**Stage 4 — Generation & Verification:** AI generates from Spec. Your job: **verification attitude**—don't trust blindly. Check each Spec item against output. Find **Undocumented ASRs**—structural choices in output not in Spec. Send those back to ADR for review.
> AI-native skill: **verification attitude**—you own the deliverable

---

## 💡 What This Pipeline Gives You

The core: **you participate actively at every stage.**

AI doesn't review for you. You decide, review, approve. AI assists.

Architecture-driven generation starts by **admitting structure is fuzzy**. Identify what matters, analyze options, decide, integrate into Spec, generate and verify against Spec.

That's why Cocrates makes "just create something" look complex—so you never accept output you don't understand.

---

## 📌 Key Takeaways

1. **AI works; you (AI-native engineer) review and decide.** Your value is how well you evaluate AI proposals and make final calls.
2. **ASR stage needs discernment**—separate truly structural requirements from noise.
3. **ADR stage needs judgment**—evaluate whether AI's analysis and proposal are optimal; ask to improve decisions.
4. **Spec stage needs insight**—check completeness; imagine resulting output.
5. **Generation/Verification needs verification attitude**—don't trust blindly; you are accountable.

---

**Ask yourself:**

- When requesting from AI, am I reviewing structure as team lead—or accepting the first proposal?
- When AI says "this matters," is it all equally important? Am I distinguishing what matters?

---

## 🎬 Coming Up Next

Today we learned Artifact Generation Pipeline principles—ASR, ADR, Spec, Verification—and why each stage exists.

Next: the four skills—adr-writing, spec-writing, spec-driven-generation, spec-driven-verification—how they work inside, as if opening each SKILL.md.

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
