---
sidebar_position: 5
---
# EP5. 架构驱动生成实践

![架构驱动生成](/img/05-architecture-driven-generation-practice.png)

---

在 EP4 中我们体验了 Cocrates 的第一个核心 activity——**学习流水线**。Education → Knowledge Capture → Reflection。如何与 AI 学习。

今天我们实践第二个核心 activity：**architecture-driven artifact generation pipeline（架构驱动产出生成流水线）**。

Cocrates 做两件大事：**学习**与**制作**。EP4 是学习；今天是**制作**。

ADR → Spec → Generation → Verification。我们将用真实案例走通用 AI 构建软件的过程。

---

## 💬 "构建 jsondb"——但 AI 坚持先设计

想象你向 Cocrates 请求：

> *"Develop jsondb in examples/jsondb."*

典型 AI 可能说*"jsondb？好！"* 并倾倒数百行代码——无审查、无设计，代码先行。

Cocrates 回答不同：

> *"这个 jsondb 用来做什么？我先设计架构，得到你的审查与批准，再安全生成代码。"*

这就是 Cocrates 的核心价值。记得 EP2 吗？**"未经审查的产出不值得生成。"** Cocrates 是践行这一原则的 AI。无设计、无审查——无代码。

今天我们构建 jsondb——一个简单的本地 JSON 存储。即使这个小项目也需要真正的架构辩论。

---

## 🏛️ ADR ① — 存储模型辩论："通用"陷阱

**情境：** 如何存储数据？

**AI 建议：** *"像 MongoDB 那样的 Collection-Document 结构很常见、也容易。"*

AI 以*"通用"*打头——训练数据里最常见的模式。

但*"通用"*意味着*"通常这样做"*——不是**"适合我的情境"**。

用户反驳：

> *"不。我要基于 path 的寻址——`Set("episode/e1.json")` 应在磁盘上创建 `episode/e1.json`。"*

这是 **Path-Addressable** 存储——与 Collection-Document 不同。

他们决定采用 path 映射。

**💡 教训：** 通用并不总适合你的需求。AI 提供最常见模式。**只有你知道**真实上下文。

---

## 🏛️ ADR ② — 架构辩论："简单容易"陷阱

**情境：** 如何交付 jsondb？

**AI 建议：** *"Engine library + app 最简单——导入即用。"*

听起来方便。导入就行。

用户问：

> *"如果以后另一个 app 也要访问这个 jsondb 呢？"*

一切翻转。Library + app 只在进程内有效。多 app 需要 **Client-Server**。

他们从 Library + app 转向 Client-Server。*"简单"*路径撑不住未来需求。

**💡 教训：** 简单容易并不总能扩展。AI 偏好最快实现。你问未来，才能得到更好结构。

---

## 🏛️ ADR ③ — 并发辩论："质量"陷阱

**情境：** 并发请求——如何处理？

**AI 建议：** *"一个 DB 级 RWMutex 就够。简单。"*

技术上成立。

用户预见性能：

> *"写其中一个文件时，读其他文件是否都要等？"*

正是。一个 DB 级 RWMutex 在任一文件写入时阻塞所有读。100 或 1000 个文件时，那就是瓶颈。

他们改为 **Per-File RWMutex Map**——不同文件可并发访问。

**💡 教训：** AI 技术上没错——但**质量（性能、规模）**考虑不足。你质疑性能以提升质量。

---

## 🎯 ADR 教训——AI 建议失准的三个原因

三个 ADR，一条重要教训：

**AI 建议失准原因很多——不只是 hallucination。**

1. **Generic（通用）：** AI 提供最常见模式。通用 ≠ 适合你。（ADR ①：存储）
2. **Simple（简单）：** AI 偏好最易实现。可能不满足未来需求。（ADR ②：架构）
3. **Quality-blind（质量盲区）：** 技术上可行但性能或规模弱。（ADR ③：并发）

识别这些，有助于你更谨慎地审查 AI 的初稿。

---

## 📋 Spec — 整合决策

三个 ADR 批准后，Cocrates 将所有决策合并为**一份自包含文档**——**Spec**。

仅凭 Spec 就应能完全说明要构建什么。无需打开 ADR。

- Storage：Path-Addressable（key → 直接文件 path）
- Architecture：Client-Server（多进程）
- Concurrency：Per-File RWMutex Map

---

## ⚙️ Generation — 从 Spec 生成代码

Cocrates 从 Spec 生成代码：

- 7 个 Go library 文件
- REST API server
- CLI
- 全部测试通过

因生成遵循 Spec，代码与决策一致——无额外假设。

---

## 🔍 Verification — Spec 检查与意外发现

即使基于 Spec 的生成也可能偏离 Spec。Verification 使用 72 项 checklist：

- **71：pass**
- **1：fail** — CLI schema URL 中的双斜杠（`//`）

重要发现：**6 个 Undocumented ASR。**

Undocumented ASR 是产出中存在但 Spec 未陈述的结构决策——本例 6 个：

- URL 编码方式
- 缺少 SetPath 验证
- 其他四项隐式设计选择

**Undocumented ASR 不是失败。** 是 sharpen Spec 的机会。在设计中审查这六项并更新。

---

## 🔄 Living Cycle — 结构不是一次定死

许多人误以为是 **Waterfall（瀑布）**。

Waterfall：设计 → 实现 → 验证（结束）

实际是**循环**：

```
ADR → Spec → Generation → Verification
  ↑                            │
  └──────── feedback ←─────────┘
```

Verification 中的问题把你送回设计——重访 ADR，必要时修改。需求变了？同样循环：ADR → Spec → 再生成 → 再验证。

**ADR 是 living documents。** 今天的决策明天可能不成立。

**Spec 是 living documents。** ADR 变了，Spec 必须变。

**结构不是定一次就忘。** 结构驱动开发；验证与变化的需求改进结构。

> 架构驱动开发不是 Waterfall。**结构引领开发；开发经验改进结构。**

---

## 📌 要点回顾

1. **审查 AI 建议。** 不只是 hallucination——Generic、Simple 或 Quality-blind 失准都会发生。用户审查与批准必不可少。
2. **ADR 与 Spec 是 living documents。** 验证或需求变化把你送回设计。结构在演进。
3. **Verification 是质量。** Spec verification 不能 catch 一切——但没有 verification 你什么都 catch 不到。Undocumented ASR 是 Spec 改进机会。
4. **架构驱动开发是 Living Cycle。** 结构引领；经验 refine。不是 Waterfall。

---

**问自己：**

- 你能说出 Artifact Generation 流水线的四个阶段吗？
- 你能用例子解释 AI 建议失准的三个原因吗？
- 什么是 Undocumented ASR？
- 为何架构驱动开发不是 Waterfall？

---

## 🎬 下期预告

今天我们体验了**架构驱动生成**——ADR、Spec、Generation、Verification——在真实项目中。

接下来 EP6 更进一步。EP5 用了现有 skills；EP6 我们**与 Cocrates 一起构建新 skill**——制作 skill 的 skill。

不是*"写一份报告"*，而是教 AI *"如何写报告"*。设计并构建自动生报告与文档的 skill。

结尾一个问题：

> **"我是否曾不经思考就复制粘贴 AI 的初稿？"**

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
