---
sidebar_position: 6
---
# EP6. Architecture-Driven Skill Creation in Practice

![Teaching AI to Create a Report-Authoring Skill](/img/06-workflow-creation-practice.png)

---

In Ep5 we saw what happens when you ask Cocrates to "build jsondb"—ADR and Spec before architecture-driven generation. We learned not to trust AI blindly and to design with review.

Today we go one step further. Ep5 was *"how to have AI make code."* Ep6 is **"how to teach AI how to work."**

The subtitle says it: teach AI *"how to write reports."* Beyond one-off requests—design AI's behavior itself.

---

## 🎣 "Write a Report" vs "How to Write Reports"

Compare two approaches.

**Typical: "Write a report"**

Every request yields different quality. Fine today, bad tomorrow—structure varies even on the same topic. Repetition fatigue. That's the limit of **one-off requests**.

**Cocrates: "Create a skill for writing reports"**

Build once, reuse forever. Same request type → same workflow → consistent quality.

Like asking *"catch me a fish"* every time vs teaching *"how to fish."*

**A skill is AI's behavioral standard—a repeatable workflow.** Rules so AI handles the same request type consistently.

Today we build **`document-authoring`**—a workflow for reports and explanatory documents.

---

## 🎯 Kernel — One Sentence for Core Purpose

Every skill starts with a **Kernel**—one sentence defining what the skill does.

Cocrates asked three questions:

1. *"What does this skill generate?"*
2. *"Who uses it?"*
3. *"What form does the output take?"*

From the answers came the Kernel:

> *"This skill helps generate Markdown explanations or reports through reviewable stages in typical situations."*

One sentence defines the skill. The Kernel guides everything after.

---

## 🏗️ Frame — Building the Skeleton

With Kernel defined, **Frame** designs document hierarchy and generation stages.

Document hierarchy:
```
Meta → Outline → Body → Conclusion → Appendix
```

Generation stages (P1–P5):
```
P1: Requirements analysis → P2: Outline → P3: Section content → P4: Integration & review → P5: Final formatting
```

Each stage has clear inputs, outputs, and approval conditions.

---

## 🎯 Outline — The Dramatic Moment

**The most important thing happened in Outline.**

Cocrates proposed *"Introduction → Body → Conclusion"*—typical and safe.

The user pushed back:

> *"It's not always intro → body → conclusion. Some reports don't need an intro; leading with the conclusion can work better."*

That broke AI's **Silent Default**. Cocrates designed **six customizable document structures**:

1. **Standard:** Introduction → Body → Conclusion
2. **Conclusion-first:** Conclusion → Body → Summary
3. **Problem-solution:** Problem → Cause → Solution
4. **Chronological:** Background → Development → Result
5. **Comparative:** Topic A → Topic B → Synthesis
6. **Freeform:** User defines structure

**💡 Lesson:** AI's first proposal isn't always right. AI offers the most common pattern. Your context and insight build better structure. *"It's not always intro→body→conclusion"* broke a fixed idea and evolved a more flexible architecture.

---

## 📋 Spec — Completing the Design

After Outline was fixed, Cocrates merged decisions into a **Spec**.

**Five core principles:**
1. Stage approval gates—user review at each step
2. Snowflake gradual expansion—start small, refine incrementally
3. Self-containment—Spec alone must be enough to generate
4. Verifiability—each item pass/fail
5. User-led—user makes final calls, not AI

**Six document structures:** as above

**Seven prohibitions:**
1. No verbose prose
2. No advancing stages without user confirmation
3. No content outside the Spec
4. ...

**Five completion criteria:**
1. User approval on all stages
2. Each section matches Spec
3. ...

---

## 📄 Skill Creation — And Verification

After the Spec, Cocrates created `.opencode/skills/document-authoring/SKILL.md`.

It verified the skill with **7 checklist items**.

**Result: all PASS ✅**

| Check | Result |
|-------|--------|
| Kernel clearly defined? | ✅ |
| Frame designed? | ✅ |
| Outline detailed? | ✅ |
| Spec self-contained? | ✅ |
| Approval gates at each stage? | ✅ |
| Prohibitions clear? | ✅ |
| Completion criteria verifiable? | ✅ |

---

## 🔄 After the Skill — What Changes

Once the skill exists, the experience changes completely.

**Before:** *"Write a report"* → AI generates wildly → user passively accepts

**After:** *"Write a project retrospective"* → `document-authoring` activates → P1 through stages with user collaboration → review and approve each step → consistent quality

You're no longer waiting passively for AI output. You become an **active designer** who shapes and controls the process.

---

## ❄️ Shared Snowflake Pattern

Ep4 (learning), Ep5 (generation), Ep6 (skill creation)—all share the **Snowflake pattern**.

| Stage | Ep4 (Learning) | Ep5 (Generation) | Ep6 (Skill) |
|-------|----------------|------------------|-------------|
| Start small | Concept briefing | Clarify requirements | Define Kernel |
| Expand gradually | 3 missions | 3 ADRs | Frame→Outline→Spec |
| Review & approve | After each mission | Each ADR approved | Each stage approved |
| Save to file | `kb/bloom-taxonomy.md` | `spec/jsondb.md` | `skills/document-authoring/SKILL.md` |

This pattern applies consistently across Cocrates activities.

---

## 📌 Key Takeaways

1. **Skills design how AI works.** Beyond one-offs—define AI behavior. Build once, reuse as lasting asset.
2. **Don't accept AI proposals without review.** AI tends toward easy, generic choices. Use U→A→E→A for better fits.
3. **Snowflake is the shared pattern.** Start small, expand, review and approve—learning, generation, skills.
4. **With a skill, you become an active designer.** No more passively waiting—you design and control.

---

**Ask yourself:**

- Can you explain one-off requests vs skills?
- Can you name the six stages of skill creation?
- How does Snowflake apply in Ep4, Ep5, and Ep6?

---

## 🎬 Coming Up Next

Three practice episodes covered Cocrates' core activities.

Ep4: **Learning** — Education, Knowledge Capture, Reflection.
Ep5: **Generation** — ADR, Spec, Generation, Verification.
Ep6: **Skill creation** — Kernel, Frame, Outline, Spec, Skill, Verification.

You've experienced all three.

Next, Ep7: **Cocrates framework structure and principles—architecture deep dive.**

You've learned to **use** Cocrates; Ep7 opens the hood—Agent and Skills, harness design, Intent-To-Skill Routing.

> **"Will you ask AI for a fish—or teach it how to fish?"**

Whatever you choose, Cocrates walks with you.

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
