---
sidebar_position: 11
---
# EP11. Architecture-Driven Generation Skills

![Cocrates Generation Skills](/img/11-architecture-driven-generation-skills.png)

---

From Ep7 through Ep10 we explored Cocrates internals—Agent routing, Skill behavior, architecture-driven generation principles, Intent-To-Skill Routing.

We understand the **why**.

Now the question shifts:

*"I get the principles. What skills does Cocrates actually have, and when does each run?"*

Today we answer that. Ep11–Ep13 dissect **core skills of the architecture-driven generation pipeline**. Today: the **four pillar skills** and the full flow.

---

## 🎯 "Generate A" — Cocrates' Internal Reasoning

When Cocrates hears *"Generate A,"* what happens inside?

**First question:** *"Is there a dedicated skill for generating A?"*

*"Generate a blog series"* → `blog-series-authoring` exists. Use it. Simple.

**Second question:** *"If no dedicated skill?"*

Fallback to **spec-driven-generation**—the generic core of the architecture-driven pipeline.

But spec-driven-generation isn't *"go build."* It asks finer questions.

**Third question:** *"Is the Spec clear?"*

If not—here's where it really starts. Cocrates walks through:

1. **ASR identification** — what structurally matters for this artifact?
2. **ADR** — alternatives and decisions per ASR
3. **Spec** — merge all ADR decisions into one self-contained doc
4. **Spec readiness** — is Spec concrete enough?
5. **Only then generation** — after Spec is complete

That's why Cocrates never allows *"build first, see later."*

---

## 🏛️ Four Pillar Skills — The Generation Engine

The architecture-driven pipeline rests on four skills:

### ① adr-writing — Design the Decision

ADR (Architecture Decision Record) isn't a casual decision log. **One Concern = one ADR**—record design choices with audit trail.

Core rules:
- Compare at least 2–3 real alternatives
- Clear pros/cons per option
- Status: `proposed` → approved → `approved`
- Record rejected options too—*"why we didn't choose this"* is future asset

**Why it matters:** Without ADR, rationale disappears. Later: *"Why did we build it this way?"*—no answer.

### ② spec-writing — Living Specification

Spec integrates ADR decisions into a **self-contained** document.

Core rules:
- **No ADR links.** Spec alone must explain everything. *"See ADR 3"* is forbidden.
- **Verifiable statements.** Not *"must be fast"* but *"must handle 1000 requests per second"*—pass/fail clear.
- **Living document.** Requirements change → Spec changes. Not write-once trash.

### ③ spec-driven-generation — Spec Is the Only Source

One iron rule in generation: **Spec is the only authority.**

Don't generate what's not in Spec. Don't add from AI intuition. If Spec is thin, strengthen Spec—don't improvise.

Cocrates runs an **ASR Readiness Gate**—eight universal categories to confirm Spec is ready.

**Spec > prompt.** No matter how clever the prompt, it can't replace Spec. Spec is verifiable, versioned, reproducible. Prompts evaporate.

### ④ spec-driven-verification — Start of the Cycle, Not the End

Verification isn't just testing. **Two purposes:**

**Purpose 1 — Deviation:** Does output match Spec? Find mismatches.

**Purpose 2 — Undocumented ASR:** Structural decisions in output not in Spec—not failure; chance to perfect Spec.

After verification, a **user review gate** remains. AI verified isn't done—you approve.

---

## 💡 Today's Insight: Skills as Specialized AI Roles

The biggest insight from Cocrates' skill system:

**Skills give AI specialized roles.**

- `adr-writing` → AI as **architect** designing decisions
- `spec-writing` → AI as **spec author** documenting
- `spec-driven-generation` → AI as **engineer** implementing to spec
- `spec-driven-verification` → AI as **QA engineer** verifying

Not one AI wearing every hat—**role-specific skills** following defined procedures. Consistent quality, predictable outcomes.

---

## 📌 Key Takeaways

1. **Generation pipeline uses two decision layers.** Dedicated skill for A? → use it. Else spec-driven-generation → Spec clear? → if not, ASR → ADR → Spec first
2. **Four pillar skills:** adr-writing, spec-writing, spec-driven-generation, spec-driven-verification
3. **Spec > prompt.** Spec is verifiable, versioned, reproducible; prompts are volatile

---

**Ask yourself:**

- Can you name Cocrates' four pillar skills?
- Can you explain spec-driven-generation's internal flow when it fallbacks?
- Why does Spec forbid ADR links?

---

## 🎬 Coming Up Next

Today we surveyed **which skills** build the architecture-driven pipeline.

One question remains:

*"What if the skill I need doesn't exist?"*

Say you need a *"regular code review report generator."* No such skill in Cocrates. What then?

**Ep12:** *"How to create a skill when there is no skill."*

Cocrates offers **generating-skill-creation**—a meta-skill. Next time we go deeper into the skill that makes skills.

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
