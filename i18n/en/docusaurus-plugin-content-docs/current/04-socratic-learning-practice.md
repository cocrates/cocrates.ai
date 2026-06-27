---
sidebar_position: 4
---
# EP4. Socratic Learning in Practice

![Cocrates Learning: Bloom's Taxonomy](/img/04-socratic-learning-practice.png)

---

This moment has arrived.

In Ep3 we installed Cocrates Harness. We verified Plugin and Skill. We're ready for the first question.

So we asked:

> *"Tell me about Bloom's Taxonomy."*

A perfectly ordinary question. But Cocrates' answer was... anything but ordinary.

Today we tell the story from that moment on.

---

## 🤷 "Do You Know Bloom's Taxonomy?"

First, a question for you:

**"Do you know Bloom's Taxonomy?"**

Most people probably don't. That's fine—I thought I *knew* Bloom's Taxonomy before this conversation too.

Here's what matters:

**Don't be ashamed of not knowing—ask AI.**

That's the practice behind Ep2's *"Never leave ignorance unattended."* Don't know? Ask. Learn. Record. Verify.

---

## 💬 Same Question, Completely Different Answer

What happens if you ask a typical AI about Bloom's Taxonomy?

Probably something like:

> *"Bloom's Taxonomy is a framework proposed by Benjamin Bloom in 1956 that classifies the cognitive domain into six levels. First, Remember... Second, Understand..."*

You lose focus in about three seconds. Too much information dumped at once.

Cocrates answered completely differently:

> *"Bloom's Taxonomy is famous for its six-level pyramid. But imagine a teacher designing a class that goes strictly from the bottom (Remember) to the top (Create). What's seriously wrong with that approach?"*

One or two sentences on the core idea, then a question that makes you think.

For those who don't know Bloom's Taxonomy: it's a six-level pyramid. Higher levels demand more advanced thinking.

**Remember → Understand → Apply → Analyze → Evaluate → Create**

Many people mistake the pyramid for a **lesson sequence**—*"You must master the foundation perfectly before moving up."*

Cocrates' first question targets exactly that misunderstanding.

Cocrates isn't a vending machine. It's a **coach that makes you think**.

---

## 🎯 Mission ① — The Pyramid Isn't a Lesson Sequence

The user wrestled with Cocrates' question—and discovered something striking:

> **The pyramid isn't a lesson sequence; it's a dependency structure.**

The old belief breaks: *"You must master the basics perfectly before moving on."*

In practice, the reverse approach often works better.

*"Okay—today build your own programming website! (Create)"*

Throw a high-level task first, and students hit walls, search for syntax (Remember), and wonder why the code runs that way (Understand).

That's the **Pull strategy**—a high-level task **pulls** lower-level knowledge into play.

---

## 🎯 Mission ② — The 2D Matrix

Cocrates kept asking:

> *"Bloom's Taxonomy isn't a simple 1D pyramid. It's actually a 2D matrix. The intersection of Knowledge Dimension and Cognitive Process defines each learning objective. Map a concept you recently learned onto this matrix."*

The user gained a second insight:

> **Bloom's Taxonomy isn't 1D—it's a 2D matrix.**

The revised taxonomy isn't just a six-step pyramid. **Knowledge Dimension** × **Cognitive Process** defines each learning objective.

- Knowledge: Factual, Conceptual, Procedural, Metacognitive
- Cognitive Process: Remember, Understand, Apply, Analyze, Evaluate, Create

**4 × 6 = 24 cells**—that's Bloom's real shape. The same concept needs different learning design depending on which knowledge type you emphasize.

---

## 🎯 Mission ③ — Push and Pull

Cocrates asked one more question:

> *"Explain the relationship between Push strategy, Pull strategy, and ZPD (Zone of Proximal Development), and summarize when to use each."*

A third insight:

> **Push (sequential) and Pull (challenge-first) should switch based on the learner's state.**

- **Push:** Step-by-step guidance. Good for beginners. Effective at Remember and Understand.
- **Pull:** A high-level task pulls learning. Good when basics exist. Effective at Create and Evaluate.
- **ZPD:** Tasks you can't solve alone but can with help (scaffolding). Cocrates aims missions in that zone—not too easy, not too hard.

---

## 📝 Knowledge Capture — Record Gaps, Not Summaries

After three missions, the user asked to *"summarize it."*

Cocrates didn't produce a perfect summary of Bloom's Taxonomy. Instead it wrote to `kb/bloom-taxonomy.md`:

1. **Core definition:** "Bloom's Taxonomy = 2D matrix (Knowledge × Cognitive Process)"
2. **Wrong assumptions (Gap):** "I treated the pyramid as a lesson sequence. It's actually a dependency structure."
3. **Push/Pull strategy:** "Push for beginners, Pull when basics exist. Mix based on ZPD."

**Recording what you didn't know—not a summary.** That's the heart of Knowledge Capture.

---

## 🔍 Reflection — Interviewer Mode

Final step. When the user asked to *"evaluate me,"* Cocrates became an interviewer.

Four missions:

**Mission 1 (Apply):** "Map DIP (Dependency Inversion Principle) onto Bloom's knowledge dimension."

**Mission 2 (Analyze):** "Explain the difference between hierarchy and sequence, and how that affects learning design."

**Mission 3 (Evaluate):** "Which is more effective from a ZPD perspective—Push or Pull? Evaluate with reasons and conditions."

**Mission 4 (Create):** "Design a 4-hour Git branching curriculum based on Bloom's Taxonomy."

Cocrates returned a clear report card:

| Area | Status |
|------|--------|
| Bloom basics (definition, six levels) | ✅ Solid |
| 2D matrix (Knowledge × Cognitive) | ✅ Solid |
| Push/Pull and ZPD | ⚠️ Wobbly—confusion on conditional use |
| Curriculum design (Create level) | ⚠️ Wobbly—needs clearer Git branching hierarchy |

**⚠️ Wobbly areas** are the real learning targets. Focus there and the next cycle begins.

---

## 🔄 The Full Cycle

The whole process is one loop:

```
Education (learn) → Knowledge Capture (record) → Reflection (verify) → (if needed) Education (more learning)
```

Learning with Cocrates doesn't end at *"tell me."* Question → insight → record → verify → (if needed) question again—that cycle completes real understanding.

---

## 📌 Key Takeaways

1. **Same question, different answer.** Cocrates doesn't dump information—it asks questions that make you think.
2. **Three insights.** Pyramid as dependency (Pull), not lesson sequence. 1D → 2D matrix (Knowledge × Cognitive). Push vs Pull depends on learner state.
3. **Learning pipeline in three steps.** Education (inquiry) → Knowledge Capture (record, especially Gaps) → Reflection (verify, interviewer mode).

---

**Ask yourself:**

- Can you explain the knowledge and cognitive axes of Bloom's 2D matrix?
- Can you map a recent concept onto that matrix?
- Among things you think you know well, what's actually ⚠️ wobbly?

---

## 🎬 Coming Up Next

Today we experienced the **Learning pipeline**—Cocrates' first core activity.

But Cocrates has another core activity: **architecture-driven artifact generation**.

Next, Ep5 starts like this:

> *"Develop jsondb in cocrates/jsondb."*

When you ask AI to "build it," how does Cocrates respond? More than code generation—a remarkable process unfolds.

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
