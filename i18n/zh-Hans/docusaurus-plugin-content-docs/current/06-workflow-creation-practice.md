---
sidebar_position: 6
---
# EP6. 架构驱动 Skill 创建实践

![教 AI 创建报告撰写 Skill](/img/06-workflow-creation-practice.png)

---

在 EP5 中我们看了向 Cocrates 请求 "build jsondb" 时会发生什么——ADR 与 Spec 在架构驱动生成之前。我们学到不能盲目信任 AI，要带着审查去设计。

今天我们再进一步。EP5 是*"如何让 AI 做代码"*。EP6 是**"如何教 AI 如何工作"**。

副标题说的：教 AI *"如何写报告"*。超越一次性请求——设计 AI 的行为本身。

---

## 🎣 "写一份报告" vs "如何写报告"

比较两种方式。

**典型："写一份报告"**

每次请求质量不同。今天还行，明天差——同一主题结构也各异。重复疲劳。这就是**一次性请求**的极限。

**Cocrates："创建一个写报告的 skill"**

建一次，永久复用。同一请求类型 → 同一工作流 → 一致质量。

像每次说*"给我抓条鱼"* vs 教*"如何钓鱼"*。

**Skill 是 AI 的行为标准——可重复的工作流。** 规则让 AI 一致处理同一类请求。

今天我们构建 **`document-authoring`**——报告与说明文档的工作流。

---

## 🎯 Kernel — 核心目的一句话

每个 skill 从 **Kernel** 开始——一句话定义 skill 做什么。

Cocrates 问了三个问题：

1. *"这个 skill 生成什么？"*
2. *"谁使用？"*
3. *"输出是什么形式？"*

从答案得出 Kernel：

> *"This skill helps generate Markdown explanations or reports through reviewable stages in typical situations."*

一句话定义 skill。Kernel 指引之后一切。

---

## 🏗️ Frame — 构建骨架

Kernel 定义后，**Frame** 设计文档 hierarchy 与生成阶段。

文档 hierarchy：
```
Meta → Outline → Body → Conclusion → Appendix
```

生成阶段（P1–P5）：
```
P1: Requirements analysis → P2: Outline → P3: Section content → P4: Integration & review → P5: Final formatting
```

每阶段有清晰的输入、输出与批准条件。

---

## 🎯 Outline — 戏剧性时刻

**最重要的事发生在 Outline。**

Cocrates 提议 *"Introduction → Body → Conclusion"*——典型且安全。

用户反驳：

> *"不总是 引言→正文→结论。有些报告不需要引言；结论先行可能更好。"*

这打破了 AI 的 **Silent Default**。Cocrates 设计了**六种可定制文档结构**：

1. **Standard：** Introduction → Body → Conclusion
2. **Conclusion-first：** Conclusion → Body → Summary
3. **Problem-solution：** Problem → Cause → Solution
4. **Chronological：** Background → Development → Result
5. **Comparative：** Topic A → Topic B → Synthesis
6. **Freeform：** 用户自定义结构

**💡 教训：** AI 初稿并不总对。AI 提供最常见模式。你的上下文与洞察构建更好结构。*"不总是 引言→正文→结论"* 打破固定观念，演进为更灵活的架构。

---

## 📋 Spec — 完成设计

Outline 定稿后，Cocrates 将决策合并为 **Spec**。

**五条核心原则：**
1. Stage approval gates——每步用户审查
2. Snowflake gradual expansion——从小开始，逐步 refine
3. Self-containment——仅凭 Spec 必须足够生成
4. Verifiability——每项 pass/fail 清晰
5. User-led——用户做最终决定，非 AI

**六种文档结构：** 如上

**七条禁止：**
1. No verbose prose
2. No advancing stages without user confirmation
3. No content outside the Spec
4. ...

**五条完成标准：**
1. User approval on all stages
2. Each section matches Spec
3. ...

---

## 📄 Skill 创建 — 与 Verification

Spec 之后，Cocrates 创建 `.opencode/skills/document-authoring/SKILL.md`。

用 **7 项 checklist** 验证 skill。

**结果：全部 PASS ✅**

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

## 🔄 Skill 之后 — 什么变了

有了 skill，体验完全改变。

**之前：** *"写一份报告"* → AI  wildly 生成 → 用户被动接受

**之后：** *"写一份项目回顾"* → `document-authoring` 激活 → P1 起各阶段与用户协作 → 每步审查批准 → 一致质量

你不再被动等待 AI 输出。你成为**主动设计者**，塑造并控制过程。

---

## ❄️ 共享 Snowflake 模式

EP4（学习）、EP5（生成）、EP6（skill 创建）——共享 **Snowflake 模式**。

| Stage | EP4 (Learning) | EP5 (Generation) | EP6 (Skill) |
|-------|----------------|------------------|-------------|
| Start small | Concept briefing | Clarify requirements | Define Kernel |
| Expand gradually | 3 missions | 3 ADRs | Frame→Outline→Spec |
| Review & approve | After each mission | Each ADR approved | Each stage approved |
| Save to file | `kb/bloom-taxonomy.md` | `spec/jsondb.md` | `skills/document-authoring/SKILL.md` |

这一模式在 Cocrates activities 中一致适用。

---

## 📌 要点回顾

1. **Skills 设计 AI 如何工作。** 超越一次性——定义 AI 行为。建一次，作为 lasting asset 复用。
2. **不要不经审查接受 AI 提案。** AI 倾向简单、通用选择。用 U→A→E→A 得到更好 fit。
3. **Snowflake 是共享模式。** 从小开始、扩展、审查批准——学习、生成、skills。
4. **有了 skill，你成为主动设计者。** 不再被动等待——你设计并控制。

---

**问自己：**

- 你能解释一次性请求 vs skills 吗？
- 你能说出 skill 创建的六个阶段吗？
- Snowflake 在 EP4、EP5、EP6 如何适用？

---

## 🎬 下期预告

三集实践覆盖了 Cocrates 的核心 activities。

EP4：**学习** — Education、Knowledge Capture、Reflection。
EP5：**生成** — ADR、Spec、Generation、Verification。
EP6：**Skill 创建** — Kernel、Frame、Outline、Spec、Skill、Verification。

你已体验全部三种。

接下来 EP7：**Cocrates 框架结构与原则——架构深潜。**

你已学会**使用** Cocrates；EP7 打开 hood——Agent 与 Skills、harness 设计、Intent-To-Skill Routing。

> **"你会向 AI 要鱼——还是教它钓鱼？"**

无论你选什么，Cocrates 与你同行。

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
