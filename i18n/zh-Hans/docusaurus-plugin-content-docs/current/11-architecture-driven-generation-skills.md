---
sidebar_position: 11
---
# EP11. 架构驱动生成 Skills

![Cocrates 生成 Skills](/img/11-architecture-driven-generation-skills.png)

---

从 EP7 到 EP10 我们探索了 Cocrates 内部——Agent 路由、Skill 行为、架构驱动生成原则、Intent-To-Skill Routing。

我们理解了**为何**。

问题转向：

*"我懂原则了。Cocrates 实际有哪些 skills，各自何时运行？"*

今天来回答。EP11–EP13 剖析**架构驱动生成流水线的核心 skills**。今天：**四个 pillar skills** 与完整流程。

---

## 🎯 "Generate A" — Cocrates 的内部推理

当 Cocrates 听到 *"Generate A"*，内部发生什么？

**第一个问题：** *"Is there a dedicated skill for generating A?"*

*"Generate a blog series"* → `blog-series-authoring` 存在。用它。简单。

**第二个问题：** *"If no dedicated skill?"*

回退到 **spec-driven-generation**——架构驱动流水线的通用核心。

但 spec-driven-generation 不是 *"go build."* 它会问更细的问题。

**第三个问题：** *"Is the Spec clear?"*

若否——这里才真正开始。Cocrates 依次走过：

1. **ASR identification** — 对此产出，结构上什么重要？
2. **ADR** — 每个 ASR 的备选方案与决策
3. **Spec** — 将所有 ADR 决策合并成一份自包含文档
4. **Spec readiness** — Spec 是否足够具体？
5. **Only then generation** — Spec 完整之后

这就是为什么 Cocrates 从不允许*"先做了再说"*。

---

## 🏛️ 四个 Pillar Skills — 生成引擎

架构驱动流水线建立在四个 skills 之上：

### ① adr-writing — 设计决策

ADR（Architecture Decision Record）不是随意的决策日志。**One Concern = one ADR**——记录设计选择与审计轨迹。

核心规则：
- 至少比较 2–3 个真实备选方案
- 每个选项有清晰的利弊
- Status: `proposed` → approved → `approved`
- 也记录被拒绝的选项——*"为何没选这个"* 是未来资产

**为何重要：** 没有 ADR，理由会消失。后来问*"为何这样构建？"*——没有答案。

### ② spec-writing — Living Specification

Spec 将 ADR 决策整合成**自包含**文档。

核心规则：
- **No ADR links.** 仅凭 Spec 必须说明一切。*"See ADR 3"* 禁止。
- **Verifiable statements.** 不是*"must be fast"* 而是*"must handle 1000 requests per second"*——pass/fail 清晰。
- **Living document.** 需求变化 → Spec 变化。不是写一次就扔。

### ③ spec-driven-generation — Spec 是唯一权威

生成中一条铁律：**Spec is the only authority.**

不生成 Spec 中没有的。不从 AI 直觉添加。若 Spec 薄，加强 Spec——不即兴发挥。

Cocrates 运行 **ASR Readiness Gate**——八个通用类别确认 Spec 就绪。

**Spec > prompt.** 无论 prompt 多 clever，不能替代 Spec。Spec 可验证、可版本化、可复现。Prompts 会蒸发。

### ④ spec-driven-verification — 循环的开始，非结束

Verification 不只是测试。**两个目的：**

**Purpose 1 — Deviation:** 产出是否 match Spec？找不匹配。

**Purpose 2 — Undocumented ASR:** 产出中存在但 Spec 未记录的结构决策——不是失败；是完善 Spec 的机会。

Verification 之后仍有 **user review gate**。AI 验证完不算结束——你批准。

---

## 💡 今日洞察：Skills 赋予 AI 专门角色

Cocrates skill 系统最大洞察：

**Skills 给 AI 专门角色。**

- `adr-writing` → AI 作为**架构师**设计决策
- `spec-writing` → AI 作为**Spec 作者**文档化
- `spec-driven-generation` → AI 作为**工程师**按 Spec 实现
- `spec-driven-verification` → AI 作为**QA 工程师**验证

不是一个 AI 戴所有帽子——**角色专用 skills** 遵循既定程序。一致质量，可预期结果。

---

## 📌 要点回顾

1. **生成流水线有两层决策。** A 有专用 skill？→ 用它。否则 spec-driven-generation → Spec 清晰？→ 若否，先 ASR → ADR → Spec
2. **四个 pillar skills：** adr-writing、spec-writing、spec-driven-generation、spec-driven-verification
3. **Spec > prompt.** Spec 可验证、可版本化、可复现；prompts 不稳定

---

**问自己：**

- 你能说出 Cocrates 四个 pillar skills 吗？
- 你能解释 spec-driven-generation 回退时的内部流程吗？
- 为何 Spec 禁止 ADR links？

---

## 🎬 下期预告

今天我们概览了**哪些 skills** 构建架构驱动流水线。

还有一个问题：

*"若我需要的 skill 不存在呢？"*

 say 你需要*"regular code review report generator."* Cocrates 里没有。then？

**EP12：** *"How to create a skill when there is no skill."*

Cocrates 提供 **generating-skill-creation**——meta-skill。下次我们深入制作 skills 的 skill。

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
