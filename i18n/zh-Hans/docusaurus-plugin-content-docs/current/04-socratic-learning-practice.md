---
sidebar_position: 4
---
# EP4. 苏格拉底式学习实践

![Cocrates 学习：Bloom's Taxonomy](/img/04-socratic-learning-practice.png)

---

这一刻到了。

在 EP3 中我们安装了 Cocrates Harness。验证了 Plugin 与 Skill。准备好第一个问题。

于是我们问了：

> *"Tell me about Bloom's Taxonomy."*

一个再普通不过的问题。但 Cocrates 的回答……完全不像寻常 AI。

今天我们从那一刻起讲述故事。

---

## 🤷 "你知道 Bloom's Taxonomy 吗？"

首先，问你一个问题：

**"你知道 Bloom's Taxonomy 吗？"**

大多数人大概不知道。没关系——在这次对话之前，我也以为我*知道* Bloom's Taxonomy。

重要的是：

**不必为不知道而羞愧——问 AI。**

这就是 EP2 *"Never leave ignorance unattended."* 背后的实践。不知道？问。学。记录。验证。

---

## 💬 同样的问题，完全不同的回答

若向典型 AI 问 Bloom's Taxonomy，大概会得到：

> *"Bloom's Taxonomy 是 Benjamin Bloom 于 1956 年提出的框架，将认知领域分为六个层次。第一，Remember……第二，Understand……"*

大约三秒就会走神。信息一次性倾倒太多。

Cocrates 的回答完全不同：

> *"Bloom's Taxonomy 以六层金字塔闻名。但想象一位教师设计一门课，严格从底层（Remember）到顶层（Create）推进。这种做法有什么严重问题？"*

一两句核心概念，然后一个让你思考的问题。

对不了解 Bloom's Taxonomy 的人：它是六层金字塔。更高层次需要更高级的思维。

**Remember → Understand → Apply → Analyze → Evaluate → Create**

许多人把金字塔误当作**课程顺序**——*"必须完美掌握基础才能往上走。"*

Cocrates 的第一个问题正针对这一误解。

Cocrates 不是自动售货机。它是**让你思考的教练**。

---

## 🎯 Mission ① — 金字塔不是课程顺序

用户与 Cocrates 的问题角力——发现了 striking 的一点：

> **金字塔不是课程顺序，而是依赖结构。**

旧信念被打破：*"必须完美掌握基础才能继续。"*

实践中，反向做法往往更有效。

*"好——今天先建你自己的编程网站！（Create）"*

先抛高层任务，学生会撞墙、查语法（Remember）、琢磨代码为何那样运行（Understand）。

这就是 **Pull 策略**——高层任务**拉动**低层知识进入 play。

---

## 🎯 Mission ② — 二维矩阵

Cocrates 继续问：

> *"Bloom's Taxonomy 不是简单的 1D 金字塔。它实际是 2D 矩阵。Knowledge Dimension 与 Cognitive Process 的交叉定义每个学习目标。把你最近学的一个概念映射到这个矩阵上。"*

用户获得第二个洞察：

> **Bloom's Taxonomy 不是 1D——是 2D 矩阵。**

修订版 taxonomy 不只是六步金字塔。**Knowledge Dimension** × **Cognitive Process** 定义每个学习目标。

- Knowledge：Factual、Conceptual、Procedural、Metacognitive
- Cognitive Process：Remember、Understand、Apply、Analyze、Evaluate、Create

**4 × 6 = 24 个单元格**——这才是 Bloom 的真实形状。同一概念因强调的知识类型不同，需要不同的学习设计。

---

## 🎯 Mission ③ — Push 与 Pull

Cocrates 再问一个问题：

> *"解释 Push 策略、Pull 策略与 ZPD（Zone of Proximal Development，最近发展区）的关系，并总结何时使用各自。"*

第三个洞察：

> **Push（顺序）与 Pull（挑战优先）应根据学习者状态切换。**

- **Push：** 逐步引导。适合初学者。在 Remember 与 Understand 上有效。
- **Pull：** 高层任务拉动学习。基础已有则合适。在 Create 与 Evaluate 上有效。
- **ZPD：** 独自无法解决、但有帮助（scaffolding）就能解决的任务。Cocrates 将 mission 瞄准这一区间——不太易，不太难。

---

## 📝 Knowledge Capture — 记录缺口，而非摘要

三个 mission 之后，用户请求*"总结一下。"*

Cocrates 没有产出 Bloom's Taxonomy 的完美摘要。而是写入 `kb/bloom-taxonomy.md`：

1. **核心定义：** "Bloom's Taxonomy = 2D 矩阵（Knowledge × Cognitive Process）"
2. **错误假设（Gap）：** "我把金字塔当作课程顺序。它实际是依赖结构。"
3. **Push/Pull 策略：** "初学者 Push，有基础则 Pull。按 ZPD 混合。"

**记录你不知道的——而非摘要。** 这就是 Knowledge Capture 的核心。

---

## 🔍 Reflection — 面试官模式

最后一步。当用户请求*"评估我"*时，Cocrates 变成面试官。

四个 mission：

**Mission 1（Apply）：** "将 DIP（Dependency Inversion Principle，依赖倒置原则）映射到 Bloom 的知识维度。"

**Mission 2（Analyze）：** "解释 hierarchy 与 sequence 的区别，以及这对学习设计的影响。"

**Mission 3（Evaluate）：** "从 ZPD 视角，Push 与 Pull 哪个更有效？用理由与条件评估。"

**Mission 4（Create）：** "基于 Bloom's Taxonomy 设计 4 小时 Git 分支课程。"

Cocrates 返回清晰的报告：

| 领域 | 状态 |
|------|--------|
| Bloom 基础（定义、六层） | ✅ Solid |
| 2D 矩阵（Knowledge × Cognitive） | ✅ Solid |
| Push/Pull 与 ZPD | ⚠️ Wobbly——条件使用上有混淆 |
| 课程设计（Create 层） | ⚠️ Wobbly——Git 分支 hierarchy 需更清晰 |

**⚠️ Wobbly 领域** 才是真正的学习目标。聚焦那里，下一轮循环开始。

---

## 🔄 完整循环

整个过程是一个循环：

```
Education（学习）→ Knowledge Capture（记录）→ Reflection（验证）→（若需要）Education（再学）
```

与 Cocrates 学习不会在*"告诉我"*处结束。提问 → 洞察 → 记录 → 验证 →（若需要）再提问——这一循环完成真正的理解。

---

## 📌 要点回顾

1. **同样的问题，不同的回答。** Cocrates 不倾倒信息——它提出让你思考的问题。
2. **三个洞察。** 金字塔是依赖（Pull），非课程顺序。1D → 2D 矩阵（Knowledge × Cognitive）。Push vs Pull 取决于学习者状态。
3. **三步学习流水线。** Education（探究）→ Knowledge Capture（记录，尤其 Gaps）→ Reflection（验证，面试官模式）。

---

**问自己：**

- 你能解释 Bloom 2D 矩阵的知识轴与认知轴吗？
- 你能把最近的概念映射到该矩阵吗？
- 在你以为很懂的东西里，哪些实际是 ⚠️ wobbly？

---

## 🎬 下期预告

今天我们体验了**学习流水线**——Cocrates 的第一个核心 activity。

但 Cocrates 还有另一个核心 activity：**architecture-driven artifact generation（架构驱动产出生成）**。

接下来 EP5 这样开始：

> *"Develop jsondb in cocrates/jsondb."*

当你让 AI "构建" 时，Cocrates 如何回应？不只是代码生成—— remarkable 的过程展开。

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
