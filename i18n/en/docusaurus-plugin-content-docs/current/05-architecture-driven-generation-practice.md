---
sidebar_position: 5
---
# EP5. Architecture-Driven Generation in Practice

![Arch-Driven Generation](/img/05-architecture-driven-generation-practice.png)

---

In Ep4 we experienced Cocrates' first core activity—the **Learning pipeline**. Education → Knowledge Capture → Reflection. How to learn with AI.

Today we practice the second core activity: the **architecture-driven artifact generation pipeline**.

Cocrates does two big things: **learning** and **making**. Ep4 was learning; today is **making**.

ADR → Spec → Generation → Verification. We'll walk through building software with AI using a real case.

---

## 💬 "Build jsondb" — But AI Insisted on Design First

Imagine you ask Cocrates:

> *"Develop jsondb in examples/jsondb."*

A typical AI might say *"jsondb? Got it!"* and dump hundreds of lines of code—no review, no design, code first.

Cocrates answers differently:

> *"What will you use this jsondb for? I'll design the architecture first, get your review and approval, then generate code safely."*

That's Cocrates' core value. Remember Ep2? **"Unexamined output isn't worth generating."** Cocrates is AI that honors that principle. No design, no review—no code.

Today we build jsondb—a simple local JSON store. Even this small project needed real architecture debates.

---

## 🏛️ ADR ① — Storage Model Debate: The "Generic" Trap

**Situation:** How do we store data?

**AI's suggestion:** *"A Collection-Document structure like MongoDB is common and easy."*

AI leads with *"generic."* The pattern it saw most in training.

But *"generic"* means *"usually done this way"*—not **"right for my situation."**

The user pushed back:

> *"No. I want path-based addressing—`Set("episode/e1.json")` should create `episode/e1.json` on disk."*

That's **Path-Addressable** storage—different from Collection-Document.

They decided on path mapping.

**💡 Lesson:** Generic isn't always right for your needs. AI offered the most common pattern. Only **you** know the real context.

---

## 🏛️ ADR ② — Architecture Debate: The "Easy and Simple" Trap

**Situation:** How do we ship jsondb?

**AI's suggestion:** *"Engine library + app is easiest—just import and use."*

Sounds convenient. Import and go.

The user asked:

> *"What if another app needs to access this jsondb later?"*

That flips everything. Library + app only works in-process. Multiple apps need **Client-Server**.

They moved from Library + app to Client-Server. The *"easy"* path didn't hold future needs.

**💡 Lesson:** Easy and simple doesn't always scale. AI favors the quickest implementation. You ask about the future to reach better structure.

---

## 🏛️ ADR ③ — Concurrency Debate: The "Quality" Trap

**Situation:** Concurrent requests—how do we handle them?

**AI's suggestion:** *"One DB-level RWMutex is enough. Simple."*

Technically valid.

The user anticipated performance:

> *"While writing one file, do all reads on other files have to wait?"*

Exactly. One DB-level RWMutex blocks every read while any file writes. With 100 or 1000 files, that's a bottleneck.

They switched to a **Per-File RWMutex Map**—concurrent access to different files.

**💡 Lesson:** AI wasn't wrong technically—but **quality (performance, scale)** wasn't fully considered. You question performance to improve quality.

---

## 🎯 Lessons from ADRs — Three Reasons AI Suggestions Miss

Three ADRs, one important lesson:

**AI suggestions can miss for many reasons—not just hallucination.**

1. **Generic:** AI offers the most common pattern. Generic ≠ right for you. (ADR ①: storage)
2. **Simple:** AI favors easiest implementation. May not meet future needs. (ADR ②: architecture)
3. **Quality-blind:** Technically fine but weak on performance or scale. (ADR ③: concurrency)

Recognizing these helps you review AI's first proposal more carefully.

---

## 📋 Spec — Integrating Decisions

After three approved ADRs, Cocrates merged every decision into **one self-contained document**—the **Spec**.

The Spec alone should fully explain what to build. No need to open ADRs.

- Storage: Path-Addressable (key → direct file path)
- Architecture: Client-Server (multi-process)
- Concurrency: Per-File RWMutex Map

---

## ⚙️ Generation — Code From Spec

Cocrates generated code from the Spec:

- 7 Go library files
- REST API server
- CLI
- All tests passing

Because generation followed the Spec, code matches decisions—no extra assumptions.

---

## 🔍 Verification — Spec Check and Unexpected Finds

Even Spec-based generation can diverge from Spec. Verification used 72 checklist items:

- **71: pass**
- **1: fail** — double slash (`//`) in CLI schema URL

Important finding: **6 Undocumented ASRs.**

Undocumented ASRs are structural decisions in the output not stated in the Spec—6 in this case:

- URL encoding approach
- Missing SetPath validation
- Four other implicit design choices

**Undocumented ASR isn't failure.** It's a chance to sharpen the Spec. Review those six in design and update.

---

## 🔄 Living Cycle — Structure Isn't Frozen

Many people mistake this for **Waterfall**.

Waterfall: design → implement → verify (done)

Reality is a **cycle**:

```
ADR → Spec → Generation → Verification
  ↑                            │
  └──────── feedback ←─────────┘
```

Problems in verification send you back to design—revisit ADRs, change them if needed. Requirements change? Same loop: ADR → Spec → regenerate → re-verify.

**ADRs are living documents.** Today's decision may not hold tomorrow.

**Specs are living documents.** When ADRs change, Specs must change.

**Structure isn't set once and forgotten.** Structure drives development; verification and changing needs improve structure.

> Architecture-driven development isn't Waterfall. **Structure leads development; development experience improves structure.**

---

## 📌 Key Takeaways

1. **Review AI suggestions.** Not just hallucination—Generic, Simple, or Quality-blind misses happen. User review and approval are essential.
2. **ADR and Spec are living documents.** Verification or requirement changes send you back to design. Structure evolves.
3. **Verification is quality.** Spec verification can't catch everything—but without verification you catch nothing. Undocumented ASR is a Spec improvement opportunity.
4. **Architecture-driven development is a Living Cycle.** Structure leads; experience refines. Not Waterfall.

---

**Ask yourself:**

- Can you name the four stages of the Artifact Generation pipeline?
- Can you explain the three reasons AI suggestions miss—with examples?
- What is Undocumented ASR?
- Why isn't architecture-driven development Waterfall?

---

## 🎬 Coming Up Next

Today we experienced **architecture-driven generation**—ADR, Spec, Generation, Verification—in a real project.

Next, Ep6 goes further. Ep5 used existing skills; Ep6 we **build a new skill with Cocrates**—a skill for making skills.

Instead of *"write a report,"* teach AI *"how to write reports."* Design and build a skill that auto-generates reports and documentation.

One closing question:

> **"Have I ever copy-pasted AI's first proposal without a second thought?"**

---

*This series introduces the Cocrates Harness framework. Cocrates is an agent harness designed for Socratic dialogue so users keep agency and grow.*
