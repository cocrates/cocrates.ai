---
sidebar_position: 9
---
# EP9. 苏格拉底式学习 Skills

![Cocrates 学习 Skills](/img/09-socratic-learning-skills.png)

---

在 EP8 中我们学了 Learning Pipeline 哲学——Socratic midwifery、Bloom、ZPD——以及 Cocrates 为何用问题回答。

现在：**这在系统中如何实现。**

`.opencode/skills/` 下有三个 SKILL.md：education、knowledge-capture、reflection。每个都是精心设计的工作流 spec——不是简单 prompt。

我们逐个打开。

---

## 🎓 Education — Socratic 1:1 学习教练

### No Spoon-feeding

education skill 的第一条规则：

> **Never give a complete answer or full solution in one shot.**

即使用户问 *"Explain DIP"*，Cocrates 不会长篇讲授。核心 idea 1–3 句，日常 analogy——占回复不到 20%。

且总是**停在未完成状态**——有瑕疵的例子、开放的缺口——等待用户填补。

### Turn-based Mission — 三个 Block

每个 education 回复有三个 block：

**💡 Concept Briefing：** 核心 1–3 句，日常 analogy。占回复不到 20%。

**💻 Thought Laboratory：** 有瑕疵的实践例子或扩展场景。让用户想*"等等——为何这样构建？"*

**🔥 MISSION：** 下一回合恰好一个任务——不是短答；用自己的话解释并说出学到了什么、不知道什么。

示例：用户问 *"What is DIP?"*

> **💡 [Concept Briefing]**
> DIP 像电源插头标准。高层模块不应依赖低层细节——应依赖 interfaces。
>
> **💻 [Thought Lab]**
> ```typescript
> class OrderService {
>   constructor(private db: MySQLDatabase) {}
> }
> ```
>
> **🔥 [MISSION]**
> 若把 `MySQLDatabase` 换成 `PostgresDatabase`，什么会坏？找一处 DIP 违反，用自己的话解释 dependency direction。

不给答案。不预告下一 mission。等待用户回复。

### Bloom 2D 矩阵 + Push & Pull

education skill 用 Bloom 作为 2D 矩阵。Cocrates **不同时抬两轴**——要么同一 knowledge type 上抬一步 cognitive，要么同一 cognitive level 上抬一维 knowledge。

按用户状态选 **Push（bottom-up steps）** 或 **Pull（challenge first）**。初学者或认知过载 → Push。已有 prior knowledge → Pull 用更高 challenge。

---

## 📝 Knowledge Capture — 记录 Ignorance

Education 的学习不能只在记忆里。knowledge-capture 保存到 `kb/` 作为 **recall-oriented essentials**。

### Essentials，非 Lectures

关键规则：**不存长篇课堂笔记、解释或完整代码。**

一条 bullet = 一个 memory unit——最多 1–2 句。示例格式：

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

**`Wrong Assumptions / Gaps` section 是核心**——记录你搞错的是最强 learning tool。

### Merge Strategy

knowledge-capture 保存前总是搜索 `kb/`。若 topic 文件已存在，**merge** 而非新建。

- Never overwrite existing KB——只加 new insight
- 在 Update History 记录 date 与 additions

学习积累，KB 从 notes 成长为 **personal knowledge base**——Reflection 的核心输入。

---

## 🔍 Reflection — 面试官模式

Learning Pipeline 最后且 crucial 阶段。

### Interviewer Persona

reflection 的 distinctive feature 是 **persona**。Cocrates **此处不是 coach 或 teacher——是 interviewer**。

目标不是 teaching——是 precisely 知道用户知道与不知道什么。以便能说：*"I know that I know nothing."*

### KB 作为 Evaluation Criteria

Reflection 用 stored KB 作为 **evaluation criteria**。Key Points 成 checklist；Wrong Assumptions / Gaps 检查 gaps 是否真的 closed。

重要：**KB 不是 answer key。** 若用户 current understanding 不同但 coherent，建议先 update KB。

### 问题深度

reflection 不问定义。

- *"用另一领域的例子解释这一原则。"*
- *"这段代码哪里违反原则？"*
- *"若做相反会怎样？"*

Interviewer 判断：**Solid**——用自己的话解释，正确 apply。**Partial**——方向对，边界 fuzzy。**Gap**——只 memorization，有 contradictions。

输出不是 grade——是 **observation and discovery**。

> ✅ Solid — 用另一领域例子正确解释 DIP dependency direction
> ⚠️ Wobbly — DIP 与 Interface Segregation 边界；overlap 处混淆

### 出现 Gap 时？

reflection 不切换到 teaching mode。从更窄角度再问一次；仍 weak 则记录 ⚠️ 并建议单独 session：*"Would you like an education session on this part?"*

---

## 🔄 三个 Skill 如何连接

这些 skills 不是 isolated——它们形成 one learning cycle：

```
Education → Knowledge Capture → Reflection → (if needed) Education
```

每个 skill 的 **exit condition 链接下一 skill 的 entry**。

1. **Education** 确认理解 → 建议 **Knowledge Capture**：*"Shall I organize this in KB?"*
2. **Knowledge Capture** 完成 → 建议 **Reflection**：*"Want to check understanding based on this?"*
3. **Reflection** 发现 gaps → 再建议 **Education**：*"Want to learn more on this?"*

循环只在用户 actively participate 时运行。Cocrates 提议；用户移动。

---

## 📌 要点回顾

1. **Education 以 No Spoon-feeding 为中心。** 三 block（Concept Briefing → Thought Lab → MISSION）引导用户发现。
2. **Knowledge Capture 记录 ignorance。** Recall essentials 与 Wrong Assumptions/Gaps——不是 summaries。同 topic merge。
3. **Reflection 是 interviewer mode。** KB 作 criteria；结果分 ✅ solid vs ⚠️ wobbly。Gaps 循环回 Education。

---

**问自己：**

- 记录从 AI 学到的东西时，我存的是"详细摘要"还是"我不知道的"？
- 检查理解时，我 honestly 承认 ⚠️ 领域而不是装懂吗？

---

## 🎬 下期预告

我们深入了 Learning Pipeline。接下来：Cocrates 第二个核心 activity。

**Artifact Generation Pipeline**——architecture-driven artifact creation。

为何 *"just write it"* 危险、什么是 ASR、ADR 与 Spec 如何保证 quality——用 country-house analogy。

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
