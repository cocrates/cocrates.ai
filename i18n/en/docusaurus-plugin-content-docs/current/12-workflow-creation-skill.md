---
sidebar_position: 12
---
# EP12. Architecture-Driven Skill Creation Principles

![Cocrates Skill Creation Principles](/img/12-workflow-creation-skill.png)

---

So far we've learned two Cocrates generation modes.

**First:** If a dedicated skill exists for A, use it. *"Generate a blog series"* → `blog-series-authoring`. Fast and simple.

**Second:** If not, fallback to **spec-driven-generation**. ASR → ADR → Spec → Generation → Verification.

But after Ep11, you might wonder:

*"What if the skill I need doesn't exist at all?"*

Say you need *"generate a regular code review report."* No such skill. You could fallback every time—or **build the skill once** if you'll repeat the pattern.

Today: **generating-skill-creation**—the skill for making skills.

---

## 🔍 Meta-Skill: A Skill for Making Skills

generating-skill-creation is Cocrates' **meta-skill**. Unlike normal skills that produce artifacts, it **designs and creates skills themselves**.

Compare by analogy:

- Normal skill: *"Catch me a fish"* → fishing technique
- Meta-skill: *"Teach me to fish"* → teaching fishing

generating-skill-creation runs on **five Artifact Components**:

**Five Artifact Components:**

1. **Assignment & Constraints:** What problem and constraints does this skill solve?
2. **Context & Rules:** What context and rules govern the skill?
3. **Entities:** What concepts and objects does the skill handle?
4. **Space & Placement:** Where do outputs live and how are they arranged?
5. **Structure & Flow:** What are the stage-by-stage flow and structure?

Defining these five is the first step in skill creation.

---

## ❄️ Snowflake Five Stages — Gradual Concretion

Once Components are identified, Cocrates refines the skill through **Snowflake five stages**:

**Stage 1 — Define:** *"What does this skill do?"* Define Kernel in one sentence.

**Stage 2 — Plan:** *"In what order?"* Plan overall stages and sequence.

**Stage 3 — Architecture Design:** *"How do components relate?"* Design Entities, Space, Flow.

**Stage 4 — Detail Design:** *"What exactly happens each stage?"* Inputs, outputs, approval conditions, prohibitions.

**Stage 5 — Generation:** *"Generate SKILL.md."* Only after all design is complete.

One critical rule:

---

## 🚫 Absolute Rule — No Generation Before Detail Design Is Complete

Among Cocrates' generation principles, this is enforced most strictly:

> **"Do not move to Generation until Detail Design is fully settled."**

Why? Ep2's principle: *"The unexamined code is not worth generating."* A skill permanently defines AI behavior. A bad skill repeatedly produces bad results.

So in Detail Design, Cocrates rigorously checks:

- Are inputs and outputs clear per stage?
- Approval gates at every stage?
- Are prohibitions concrete?
- How are exceptions handled?

Not one line of SKILL.md until every question is answered.

---

## 🔄 Meta Snowflake Seven Stages — The Bigger Picture

Snowflake five stages are the basic skill design process. generating-skill-creation paints a bigger picture: **Meta Snowflake seven stages**.

```
Kernel → Components → Frame → Outline → Spec → Skill → Verification
```

In order:

**Kernel:** One sentence defining skill purpose.

**Components:** Identify and define the five Artifact Components.

**Frame:** Overall structure and skeleton—which stages exist.

**Outline:** Detail Frame—sub-items and flow per stage.

**Spec:** Self-contained spec integrating all decisions. Anyone reading only Spec understands the full behavior.

**Skill:** Generate SKILL.md from Spec.

**Verification:** Does the skill match Spec? Anything missing?

Not pure sequence or parallel—each stage depends on the prior; verification feedback can return to earlier stages.

---

## 🎯 Per-Stage Refinement — Lazy Refinement Strategy

One important principle in Meta Snowflake seven stages:

**Components don't concretize at the same rate.**

Some are very specific early; some clarify only at the end. Cocrates calls this **Lazy Refinement**.

Example: *Entities* get rough shape in Kernel/Components, detail in Outline/Spec. *Structure & Flow* need detail from Frame onward.

Why? **Avoid premature decisions.** Early lock-in makes later better options hard to adopt. Lazy refinement delays decisions to keep flexibility.

---

## 💡 Today's Insight: Skills Are AI SOPs

A skill is giving AI an **SOP (Standard Operating Procedure)**.

When a new hire joins, you teach workflow: *"When this request comes, check A, review B, then do C."*

Skills do that for AI: *"For this request type, work this way."*

Once defined, a skill is reusable asset—today, next week, a year from now, same quality.

**One-off requests are consumables. Skills are assets.**

---

## 📌 Key Takeaways

1. **generating-skill-creation is a meta-skill** based on five Artifact Components.
2. **Snowflake five stages:** Define → Plan → Architecture Design → Detail Design → Generation—with approval gates between stages.
3. **Absolute rule: no Generation before Detail Design is complete.** Bad skills multiply bad output.
4. **Meta Snowflake seven stages:** Kernel → Components → Frame → Outline → Spec → Skill → Verification.
5. **Lazy refinement:** don't concretize everything at once—only when needed.

---

**Ask yourself:**

- Can you explain Snowflake five vs Meta Snowflake seven?
- Can you name and explain the five Artifact Components?
- Why is Lazy Refinement important?

---

## 🎬 Coming Up Next

We've covered all of Cocrates.

Ep1–2 philosophy and principles, Ep3 install, Ep4–6 practice, Ep7–10 internals, Ep11–12 core and meta-skills.

**We're ready to use Cocrates.**

But an important question remains:

*"Is that the end? Is 'using Cocrates well' enough?"*

No. The real part starts now. Using Cocrates isn't enough.

**Ep13:** Users must evolve; Cocrates Harness must evolve too—and the two form a reinforcing feedback loop.

*"Using Cocrates isn't enough—you must evolve Cocrates."*

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
