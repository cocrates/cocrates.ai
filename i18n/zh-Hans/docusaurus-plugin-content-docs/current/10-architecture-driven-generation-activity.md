---
sidebar_position: 10
---
# EP10. 架构驱动生成原则

![Cocrates 生成原则](/img/10-architecture-driven-generation-activity.png)

---

在 EP8 与 EP9 中我们探索了 Learning Pipeline 的哲学与 skills。Education → Knowledge Capture → Reflection——提问、记录 ignorance、以面试官模式验证。

现在是 Cocrates 第二个核心 activity：**Artifact Generation Pipeline**——架构驱动的产出生成。

记住：**AI 干活。** 你是**AI原生型工程师（团队负责人）**，审查并决定 AI 的产出。在 AI 为你做一切的时代，你的价值在于如何**审查、判断、决定** AI 的工作。

若学习是"如何识别 ignorance 并增长知识"，生成就是**"你如何审查并批准 AI 提出的结构，再从该结构构建"**。

---

## 🚨 "Just Write It" 的陷阱

你向 AI 请求：*"Write a report."* 五秒三十页。你读——逻辑薄、各节深度不均。难要求重做；你勉强修补，浪费时间。

原因是什么？

**没有结构的生成。**

无结构产出有三个问题：

**第一，一致性与逻辑弱。** 没有单一主线——AI 按当下听起来合理的方向列举。深度不一；论断不对齐。

**第二，你无法理解或解释交付物。** *"看起来没问题"* 就继续——但*"为何这一节这样结构？"* 难住你。结构是 AI 定的，不是你。

**第三，不可审查的黑箱。** 审查需要标准。无结构？只有*"看起来合理吗？"*

那就是 *"just write it"*——实质上是 **"随便写"**。

---

## 🏗️ ASR — 什么是 fine china？

*"好——我得先设计结构。"*

但"结构"很抽象——像*"好好生活"*没有具体步骤。

进入 **ASR（Architecturally Significant Requirement，架构显著需求）**——**实质性地影响**最终交付物结构、组成与质量的需求或设计决策。来自软件架构，但适用于**每种交付物类型**——文档、幻灯片、博客系列。

ASR 重要，因为不审查 ASR，AI 会用 **Silent Defaults（静默默认值）** 填缺口——*"还算合理"* 的选择可能不符你的意图，你甚至不知道它们被选了。

---

## 🏠 乡村别墅类比 — Silent Defaults 的悲剧

你在搬家。要打包的箱子。

你告诉搬运工**"pack it well"**。他们按自己的方式打包。fine china 压在重书下面。

无恶意——只是他们的默认方式与你对"well packed"的理解不同。

那就是 **Silent Default**。你不明确指定，AI 用 defaults 填——可能与你的意图不符。

结构是决定什么放哪里、怎么放。**ASR 是识别什么是 fine china。**

想象建一座乡村别墅。

你告诉建筑师：*"Design the house that best fits our family's lifestyle."*

决策：一层还是两层？屋顶、露台、阁楼？

每个都是 **Concern** 与 **ASR**——对结构与可用性的真实影响。

**Case 1 — 用户清晰决定：**
> *"Build a second-floor attic room."*
决策已定。无需 ADR——记入 Spec。

**Case 2 — 用户委托：**
> *"Choose what best fits our lifestyle."*
**需要 ADR。** 建筑师比较选项；用户审查并决定。

---

## 🔄 四阶段 Pipeline

在 Cocrates 的 Artifact Generation Pipeline 中：

```
ASR identification → ADR → Spec → Generation & Verification
```

**阶段 1 — ASR identification：** AI 浮现 ASR。你的职责：**discernment（辨识力）**——哪些真正影响结构、哪些是噪声。没有这步，AI 要么把一切都当重要，要么独自决定一切。
> AI原生型技能：**辨识力**——区分重要 ASR 与次要需求

**阶段 2 — ADR（alternatives & decision）：** AI 按 Concern 分析选项。你的职责：**judgment（判断力）**——分析是否充分？还有其他选项吗？AI 遗漏了什么？ADR 审计*为何这样决定*。
> AI原生型技能：**判断力**——评估 AI 分析并做出最优决定

**阶段 3 — Spec（integrate decisions）：** AI 将已批准的 ADR 合并成一份文档。你的职责：**insight（洞察力）**——缺什么？从这份 Spec 最终产出会怎样？缺口送回 ADR。Spec 必须 **self-contained（自包含）**——一读就能说明一切。
> AI原生型技能：**洞察力**——验证 Spec 完整性并预判产出

**阶段 4 — Generation & Verification：** AI 从 Spec 生成。你的职责：**verification attitude（验证态度）**——不盲目信任。逐项对照 Spec 与产出。找出 **Undocumented ASR**——产出中存在但 Spec 未记录的结构决策。送回 ADR 审查。
> AI原生型技能：**验证态度**——你对交付物负责

---

## 💡 这条 Pipeline 给你什么

核心：**你在每个阶段都主动参与。**

AI 不替你审查。你决定、审查、批准。AI 辅助。

架构驱动生成从**承认结构是模糊的**开始。识别什么重要、分析选项、决定、整合进 Spec、按 Spec 生成并验证。

这就是为什么 Cocrates 让 "just create something" 看起来复杂——让你绝不接受不理解的产出。

---

## 📌 要点回顾

1. **AI 干活；你（AI原生型工程师）审查并决定。** 你的价值在于如何评估 AI 提案并做最终决定。
2. **ASR 阶段需要辨识力**——分离真正影响结构的需求与噪声。
3. **ADR 阶段需要判断力**——评估 AI 分析与提案是否最优；要求改进决策。
4. **Spec 阶段需要洞察力**——检查完整性；想象最终产出。
5. **Generation/Verification 需要验证态度**——不盲目信任；你负最终责任。

---

**问自己：**

- 向 AI 请求时，我是以团队负责人身份审查结构——还是接受初稿？
- AI 说"这个重要"时，是否同等重要？我是否区分什么真正重要？

---

## 🎬 下期预告

今天我们学了 Artifact Generation Pipeline 原则——ASR、ADR、Spec、Verification——以及每个阶段为何存在。

接下来：四个 skills——adr-writing、spec-writing、spec-driven-generation、spec-driven-verification——内部如何工作，像打开每个 SKILL.md。

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
