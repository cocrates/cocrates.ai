---
sidebar_position: 7
---
# EP7. Cocrates Harness 结构

![Cocrates Harness 架构](/img/07-cocrates-harness-structure.png)

---

在 EP6 中我们学会了教 AI "如何写报告" 而不是 "写一份报告"。

但你可能会想：

*"Cocrates 是如何做到这些不同事情的？"*

EP4 学习、EP5 制作、EP6 skill 创建——该看看内部了。就像学会开车后打开引擎盖。

---

## 🔪 "一把刀能做完所有厨房活吗？"

想象一个厨房。切番茄、切面包、片鱼。

**一把刀**能全做吗？

也许——但番茄被压烂、面包撕破、鱼鳞还在。专业厨房分刀是有原因的。

Cocrates 也一样。

有人说 "教我 Bloom's Taxonomy"，有人说 "写报告"，有人说 "做幻灯片"。**一个 prompt** 能全处理吗？

**不能。**

每种交付物需要不同的结构方法。报告需要大纲与逻辑流；幻灯片每页需要核心信息；学习需要回合制 mission。一个巨大 prompt 不可维护。

这里显现 Cocrates 的设计哲学。

---

## 🏛️ Cocrates = Agent + Skills

Cocrates 不是一个巨大 prompt。它是**两层 harness**。

**Cocrates Agent** — 共享原则与控制。读取意图、选择 skills、管理流程。像指挥官部署合适单位。

**Skills** — 每项任务的具体程序。每个 skill 在自己的文件（`.opencode/skills/*/SKILL.md`），随时可增改。像专业团队。

关键是**分离**。

- 通用原则在 Agent；每项任务程序在 Skills
- Skills 独立——改一个不破坏其他
- 新请求类型？加 Skills。Agent 不变

---

## 🔍 Cocrates Agent — 六个 Section

`cocrates.md` 是一个含六个 section 的 prompt。

**1. Persona：** Agent 如何定义自己。*"将不确定性转化为系统化探究，并引导基于结构的设计、审查与批准，直到用户完全理解交付物。"* 这是 Cocrates 的身份。

**2. Principle：** 最重要——**Harness Ignorance（驾驭无知）**。不让用户接受不理解的输出；用问题揭示假设与缺口。

**3. Harness Architecture：** *"Cocrates 是 core agent 加 skills harness。"* Agent 处理原则、意图、skill 选择、任务管理、护栏；Skills 处理每项任务工作流。

**4. Request Handling：** 推断意图、加载合适 skill、运行其工作流。使用 **Intent-To-Skill Routing**——learning → education，capture → knowledge-capture，evaluation → reflection，decisions → adr-writing 等。八种 intent 映射八种 skills。

**5. Core Activities：** 两条流水线：

- **Generation Pipeline：** Design（ADR → Spec）→ Spec-based generation → Spec-based verification
- **Learning Pipeline：** Education → Knowledge Capture → Reflection

**6. Success Criteria：** 对话结束时——用户是否更清楚自己知道与不知道什么？能否解释交付物结构？是否每一步都主动参与？

---

## 💡 "Use → Understand → Evolve"

Cocrates 最重要的特质：**它不是成品 harness。**

**用户演进它。**

从 bundled skills 开始。使用并想*"这个工作流不适合我的情况。"* 编辑 skill。需要新任务类型？构建 skill。

这一循环是 Cocrates 的真正价值。

```
Use → Understand → Evolve
```

保持 Agent 原则与 persona；通过 Skills 扩展。一个巨大 prompt 起初看起来简单，但越来越难维护。Agent + Skills 从一开始就解决这一点。

---

## 📌 要点回顾

1. **一个 prompt 不能覆盖每种交付物。** 不同输出不同结构；一个巨大 prompt 不可维护。
2. **Cocrates = Cocrates Agent + Skills。** Agent 负责原则与控制；Skills 负责每项任务程序。分离是核心。
3. **Cocrates 不是成品——它演进。** Use → understand → evolve；用户 grow 系统。

---

**问自己：**

- 你的 AI 工具是一个 prompt 还是 Agent + Skills？
- 出现新请求类型时，扩展有多容易？

---

## 🎬 下期预告

今天我们概览了 Cocrates 结构。Agent + Skills，两条流水线——Generation 与 Learning。

接下来我们深入每条流水线，从**学习**开始。

为何 *"tell me"* 危险，为何 Cocrates 用问题回答——我们剖析哲学与机制。

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
