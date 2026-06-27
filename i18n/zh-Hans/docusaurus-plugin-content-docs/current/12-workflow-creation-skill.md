---
sidebar_position: 12
---
# EP12. 架构驱动工作流创建 Skill

![Cocrates Skill 创建原则](/img/12-workflow-creation-skill.png)

---

至此我们学了 Cocrates 两种生成模式。

**第一：** 若 A 有 dedicated skill，用它。*"Generate a blog series"* → `blog-series-authoring`。快速简单。

**第二：** 若没有，回退到 **spec-driven-generation**。ASR → ADR → Spec → Generation → Verification。

但 EP11 之后你可能想：

*"若我需要的 skill 根本不存在呢？"*

 say 你需要*"generate a regular code review report."* 没有这样的 skill。你可以每次回退——或若会重复该模式，**建一次 skill**。

今天：**generating-skill-creation**——制作 skills 的 skill。

---

## 🔍 Meta-Skill：制作 Skills 的 Skill

generating-skill-creation 是 Cocrates 的 **meta-skill**。与普通 skill 产出 artifacts 不同，它**设计并创建 skills 本身**。

类比比较：

- Normal skill：*"Catch me a fish"* → 捕鱼技巧
- Meta-skill：*"Teach me to fish"* → 教捕鱼

generating-skill-creation 运行在 **五个 Artifact Components** 上：

**五个 Artifact Components：**

1. **Assignment & Constraints：** 这个 skill 解决什么问题与约束？
2. **Context & Rules：** 什么上下文与规则约束该 skill？
3. **Entities：** 该 skill 处理什么概念与对象？
4. **Space & Placement：** 输出存放在哪里、如何组织？
5. **Structure & Flow：** 各阶段的流程与结构是什么？

定义这五个是 skill 创建的第一步。

---

## ❄️ Snowflake 五阶段 — 逐步具体化

Components 识别后，Cocrates 通过 **Snowflake 五阶段** 细化 skill：

**Stage 1 — Define：** *"What does this skill do?"* 一句话定义 Kernel。

**Stage 2 — Plan：** *"In what order?"* 规划整体阶段与顺序。

**Stage 3 — Architecture Design：** *"How do components relate?"* 设计 Entities、Space、Flow。

**Stage 4 — Detail Design：** *"What exactly happens each stage?"* 输入、输出、批准条件、禁止项。

**Stage 5 — Generation：** *"Generate SKILL.md."* 全部设计完成之后。

一条关键规则：

---

## 🚫 绝对规则 — Detail Design 完成前禁止 Generation

Cocrates 生成原则中，这条执行最严：

> **"Do not move to Generation until Detail Design is fully settled."**

为何？EP2 原则：*"The unexamined code is not worth generating."* Skill 永久定义 AI 行为。坏的 skill 会反复产出坏结果。

因此在 Detail Design，Cocrates 严格检查：

- 每阶段输入与输出是否清晰？
- 每阶段是否有批准关卡？
- 禁止项是否具体？
- 异常如何处理？

每个问题都有答案之前，不写一行 SKILL.md。

---

## 🔄 Meta Snowflake 七阶段 — 更大图景

Snowflake 五阶段是基本 skill 设计流程。generating-skill-creation 描绘更大图景：**Meta Snowflake 七阶段**。

```
Kernel → Components → Frame → Outline → Spec → Skill → Verification
```

顺序：

**Kernel：** 一句话定义 skill 目的。

**Components：** 识别并定义五个 Artifact Components。

**Frame：** 整体结构与骨架——哪些阶段存在。

**Outline：** 细化 Frame——每阶段子项与流程。

**Spec：** 自包含 spec 整合所有决策。只读 Spec 就能理解完整行为。

**Skill：** 从 Spec 生成 SKILL.md。

**Verification：** Skill 是否 match Spec？缺什么？

不是纯顺序或并行——每阶段依赖前一阶段；verification 反馈可回到更早阶段。

---

## 🎯 每阶段 Refinement — Lazy Refinement Strategy

Meta Snowflake 七阶段中一条重要原则：

**Components 不会以同样速率具体化。**

有些早期就很具体；有些直到最后才清晰。Cocrates 称 **Lazy Refinement**。

示例：*Entities* 在 Kernel/Components 得粗略形状，在 Outline/Spec 细化。*Structure & Flow* 从 Frame 起就需要细节。

为何？**避免过早决策。** 早期锁定会使后期更好选项难以采纳。Lazy refinement 延迟决策以保持灵活性。

---

## 💡 今日洞察：Skills 是 AI 的 SOP

Skill 是给 AI 一条 **SOP（Standard Operating Procedure，标准作业程序）**。

新人入职，你教工作流：*"When this request comes, check A, review B, then do C."*

Skills 对 AI 做同样：*"For this request type, work this way."*

一旦定义，skill 是可复用资产——今天、下周、一年后，同样质量。

**一次性请求是消耗品。Skills 是资产。**

---

## 📌 要点回顾

1. **generating-skill-creation 是基于五个 Artifact Components 的 meta-skill。**
2. **Snowflake 五阶段：** Define → Plan → Architecture Design → Detail Design → Generation——阶段间有批准关卡。
3. **绝对规则：Detail Design 完成前禁止 Generation。** 坏的 skills 会放大坏产出。
4. **Meta Snowflake 七阶段：** Kernel → Components → Frame → Outline → Spec → Skill → Verification。
5. **Lazy refinement：** 不要一次全部具体化——只在需要时。

---

**问自己：**

- 你能解释 Snowflake 五阶段 vs Meta Snowflake 七阶段吗？
- 你能说出并解释五个 Artifact Components 吗？
- 为何 Lazy Refinement 重要？

---

## 🎬 下期预告

我们覆盖了 Cocrates 全部内容。

EP1–2 哲学与原则，EP3 安装，EP4–6 实践，EP7–10 内部，EP11–12 核心与 meta-skills。

**我们已准备好使用 Cocrates。**

但一个重要问题仍在：

*"这就结束了吗？'用好 Cocrates' 够了吗？"*

不。真正的部分现在开始。用好 Cocrates 不够。

**EP13：** 用户必须演进；Cocrates Harness 也必须演进——两者形成相互强化的反馈循环。

*"Using Cocrates isn't enough—you must evolve Cocrates."*

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
