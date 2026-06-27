---
sidebar_position: 7
---
# EP7. Cocrates Harness Architecture

![Cocrates Harness Architecture](/img/07-cocrates-harness-structure.png)

---

In Ep6 we learned to teach AI "how to write reports" instead of "write a report."

But you might wonder:

*"How is Cocrates built so it can do all these different things?"*

Ep4 learning, Ep5 making, Ep6 skill creation—time to look inside. Like opening the hood after you've learned to drive.

---

## 🔪 "Can One Knife Do Every Kitchen Task?"

Picture a kitchen. Tomatoes to slice, bread to cut, fish to fillet.

Can **one knife** do it all?

Maybe—but tomatoes get crushed, bread tears, scales stay on the fish. A pro kitchen has separate knives for a reason.

Cocrates is the same.

Someone says "teach me Bloom's Taxonomy," someone says "write a report," someone says "make slides." Can **one prompt** handle all of that?

**No.**

Each deliverable needs a different structural approach. Reports need outlines and logic flow; slides need governing messages per page; learning needs turn-based missions. One giant prompt becomes unmaintainable.

Here Cocrates' design philosophy shows up.

---

## 🏛️ Cocrates = Agent + Skills

Cocrates isn't one huge prompt. It's a **harness** with two layers.

**Cocrates Agent** — shared principles and control. Reads intent, picks skills, manages flow. Like a commander deploying the right unit.

**Skills** — concrete procedures per task. Each skill lives in its own file (`.opencode/skills/*/SKILL.md`), addable and editable anytime. Like specialist teams.

The key is **separation**.

- Common principles on the Agent; per-task procedure on Skills
- Skills are independent—change one without breaking others
- New request types? Add Skills. Agent stays

---

## 🔍 Cocrates Agent — Six Sections

`cocrates.md` is a prompt with six sections.

**1. Persona:** How the agent defines itself. *"An entity that turns uncertainty into systematic inquiry and guides structure-based design, review, and approval until the user fully understands the deliverable."* That's Cocrates' identity.

**2. Principle:** Most important—**Harness Ignorance**. Don't let users accept output they don't understand; use questions to surface assumptions and gaps.

**3. Harness Architecture:** *"Cocrates is a core agent plus skills harness."* Agent handles principles, intent, skill choice, task management, guardrails; Skills handle per-task workflows.

**4. Request Handling:** Infer intent, load the right skill, run its workflow. Uses **Intent-To-Skill Routing**—learning → education, capture → knowledge-capture, evaluation → reflection, decisions → adr-writing, etc. Eight intents map to eight skills.

**5. Core Activities:** Two pipelines:

- **Generation Pipeline:** Design (ADR → Spec) → Spec-based generation → Spec-based verification
- **Learning Pipeline:** Education → Knowledge Capture → Reflection

**6. Success Criteria:** When the conversation ends—does the user know more clearly what they know and don't know? Can they explain the deliverable's structure? Did they participate actively at every step?

---

## 💡 "Use → Understand → Evolve"

Cocrates' most important trait: **it's not a finished harness.**

**Users evolve it.**

Start with bundled skills. Use them and think *"This workflow doesn't fit my case."* Edit the skill. Need a new task type? Build a skill.

That loop is Cocrates' real value.

```
Use → Understand → Evolve
```

Keep Agent principles and persona; expand through Skills. One giant prompt looks simple at first but gets harder to maintain. Agent + Skills solves that from the start.

---

## 📌 Key Takeaways

1. **One prompt can't cover every deliverable.** Different structures per output; one huge prompt is unmaintainable.
2. **Cocrates = Cocrates Agent + Skills.** Agent for principles and control; Skills for per-task procedure. Separation is the core.
3. **Cocrates isn't finished—it evolves.** Use → understand → evolve; users grow the system.

---

**Ask yourself:**

- Is your AI tool one prompt or Agent + Skills?
- When a new request type appears, how easily can you extend it?

---

## 🎬 Coming Up Next

Today we surveyed Cocrates structure. Agent + Skills, two pipelines—Generation and Learning.

Next we go deeper into each pipeline, starting with **Learning**.

Why *"tell me"* is dangerous, why Cocrates answers with questions—we unpack the philosophy and mechanics.

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
