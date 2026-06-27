---
sidebar_position: 9
---
# EP9. Socratic Learning Skills

![Cocrates Learning Skills](/img/09-socratic-learning-skills.png)

---

In Ep8 we learned Learning Pipeline philosophy—Socratic midwifery, Bloom, ZPD—and why Cocrates answers with questions.

Now: **how that's implemented in the system.**

Under `.opencode/skills/` are three SKILL.md files: education, knowledge-capture, reflection. Each is a carefully designed workflow spec—not a simple prompt.

Let's open them one by one.

---

## 🎓 Education — Socratic 1:1 Learning Coach

### No Spoon-feeding

The first rule of the education skill:

> **Never give a complete answer or full solution in one shot.**

Even if the user asks *"Explain DIP,"* Cocrates won't lecture. Core idea in 1–3 sentences, everyday analogy—under 20% of the response.

And it always **stops incomplete**—flawed examples, open gaps—waiting for the user to fill them.

### Turn-based Mission — Three Blocks

Every education response has three blocks:

**💡 Concept Briefing:** Core in 1–3 sentences, everyday analogy. Under 20% of the reply.

**💻 Thought Laboratory:** Flawed practical example or extended scenario. Makes the user think *"Wait—why is it built like this?"*

**🔥 MISSION:** Exactly one task for the next turn—not short answer; explain in your own words and name what you learned and what you didn't know.

Example: user asks *"What is DIP?"*

> **💡 [Concept Briefing]**
> DIP is like power plug standards. High-level modules shouldn't depend on low-level details—they depend on interfaces.
>
> **💻 [Thought Lab]**
> ```typescript
> class OrderService {
>   constructor(private db: MySQLDatabase) {}
> }
> ```
>
> **🔥 [MISSION]**
> If you swap `MySQLDatabase` for `PostgresDatabase`, what breaks? Find one DIP violation and explain the dependency direction in your own words.

No answer given. No preview of the next mission. It waits for the user's reply.

### Bloom's 2D Matrix + Push & Pull

The education skill uses Bloom as a 2D matrix. Cocrates **doesn't raise both axes at once**—either one cognitive step on the same knowledge type, or one knowledge dimension at the same cognitive level.

It chooses **Push (bottom-up steps)** or **Pull (challenge first)** by user state. Beginners or cognitive overload → Push. Prior knowledge → Pull with a higher challenge.

---

## 📝 Knowledge Capture — Record Ignorance

Learning from Education can't live only in memory. knowledge-capture saves to `kb/` as **recall-oriented essentials**.

### Essentials, Not Lectures

Key rule: **don't store long lecture notes, explanations, or full code.**

One bullet = one memory unit—1–2 sentences max. Example format:

```markdown
# Dependency Inversion Principle

## One-line Definition
High-level modules must not depend directly on low-level implementations—they depend on abstractions (interfaces).

## Key Points
- High-level modules must not depend directly on low-level implementations
- Depend on interfaces; receive implementations via injection
- DIP isn't just "use an interface"—dependency direction is the point

## Wrong Assumptions / Gaps
- DIP = always create a new interface file (X)
  → At legacy boundaries, an Adapter to connect is also valid
```

**The `Wrong Assumptions / Gaps` section is central**—recording what you got wrong is the strongest learning tool.

### Merge Strategy

knowledge-capture always searches `kb/` before saving. If a topic file exists, it **merges** instead of creating a new file.

- Never overwrite existing KB—only add new insight
- Log date and additions in Update History

As learning accumulates, KB grows from notes into a **personal knowledge base**—the core input for Reflection.

---

## 🔍 Reflection — Interviewer Mode

The last and crucial stage of the Learning Pipeline.

### Interviewer Persona

reflection's distinctive feature is **persona**. Cocrates is **not coach or teacher here—interviewer**.

Goal isn't teaching—it's knowing precisely what the user knows and doesn't. So they can say: *"I know that I know nothing."*

### KB as Evaluation Criteria

Reflection uses stored KB as **evaluation criteria**. Key Points become checklists; Wrong Assumptions / Gaps check whether gaps were actually closed.

Important: **KB isn't an answer key.** If the user's current understanding differs but is coherent, suggest updating KB first.

### Depth of Questions

reflection doesn't ask for definitions.

- *"Explain this principle with an example from another domain."*
- *"Where does this code violate the principle?"*
- *"What breaks if you do the opposite?"*

The interviewer judges: **Solid**—explains in own words, applies correctly. **Partial**—direction right, boundaries fuzzy. **Gap**—memorization only, contradictions.

Output isn't a grade—it's **observation and discovery**.

> ✅ Solid — Correctly explained DIP dependency direction with another domain example
> ⚠️ Wobbly — Boundary between DIP and Interface Segregation; confusion where they overlap

### When a Gap Appears?

reflection doesn't switch to teaching mode. It asks once more from a narrower angle; if still weak, records ⚠️ and suggests a separate session: *"Would you like an education session on this part?"*

---

## 🔄 How the Three Skills Connect

These skills aren't isolated—they form one learning cycle:

```
Education → Knowledge Capture → Reflection → (if needed) Education
```

Each skill's **exit condition links to the next skill's entry**.

1. When **Education** confirms understanding → suggest **Knowledge Capture**: *"Shall I organize this in KB?"*
2. When **Knowledge Capture** finishes → suggest **Reflection**: *"Want to check understanding based on this?"*
3. When **Reflection** finds gaps → suggest **Education** again: *"Want to learn more on this?"*

The cycle runs only when the user participates actively. Cocrates proposes; the user moves.

---

## 📌 Key Takeaways

1. **Education centers on No Spoon-feeding.** Three blocks (Concept Briefing → Thought Lab → MISSION) lead users to discover.
2. **Knowledge Capture records ignorance.** Recall essentials and Wrong Assumptions/Gaps—not summaries. Merge same topics.
3. **Reflection is interviewer mode.** KB as criteria; results split ✅ solid vs ⚠️ wobbly. Gaps cycle back to Education.

---

**Ask yourself:**

- When recording what I learned from AI, do I save "detailed summaries" or "what I didn't know"?
- When checking my understanding, do I honestly admit ⚠️ areas instead of faking it?

---

## 🎬 Coming Up Next

We've gone inside the Learning Pipeline. Next: Cocrates' second core activity.

**Artifact Generation Pipeline**—architecture-driven artifact creation.

Why *"just write it"* is dangerous, what ASR is, how ADR and Spec guarantee quality—with a country-house analogy.

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
