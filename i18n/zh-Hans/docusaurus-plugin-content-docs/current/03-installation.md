---
sidebar_position: 3
---
# EP3. 安装 Cocrates Harness

![Cocrates 安装流程](/img/03-installation.png)

---

在 EP2 中我们宣告了一条强有力的原则：

> "The unexamined code is not worth generating."
> 未经审查的产出，不值得创造。

原则很清晰。AI 制作的一切都必须通过审查（U→A→E→A），你不能让无知 unattended。

但还有一个实际问题：

*"我懂原则了。实际怎么开始？"*

今天来回答。安装大约五分钟。但请记住——**安装之后**才是真正的开始。

---

## 🎭 Cocrates Harness 究竟是什么

首先，明确 Cocrates Harness 是什么。

Cocrates Harness 不是独立程序。它是运行在名为 **opencode** 的平台上的**插件**。

可以这样理解：opencode 是**舞台**；Cocrates Harness 是舞台上的**演员**。没有舞台，演员无处立足。

因此 Cocrates 设置分三步：

**一、** 安装 opencode。
**二、** 配置 Cocrates 插件。
**三、** 准备 skill 文件。

记住这三步。设置比你想的简单。

---

## 🖥️ 三种 UI 选项

使用 opencode 有三种方式。按你的环境选择。

### 1) opencode Terminal（TUI）
在终端中键盘操作为主。熟悉 CLI 的话很合适。

### 2) opencode Desktop（Beta）
安装的 GUI 应用。仍为 beta——功能与稳定性可能不完善。

### 3) VS Code 或 Cursor + OpenChamber 扩展 ⭐（推荐）
若你熟悉 VS Code 或 Cursor，可以像 Copilot 或 Cursor 一样自然使用 Cocrates。注意：VS Code + OpenChamber 仍需要先安装 opencode。

---

## ⚙️ 步骤 1 — 安装 opencode

### macOS / Linux

打开终端，运行一行：

```bash
curl -fsSL https://opencode.ai/install | bash
```

curl 安全下载安装脚本；bash 执行它。opencode 二进制会自动安装。

完成后验证：

```bash
opencode --version
```

若输出版本号，即成功。若出现 *"command not found"*，将 `~/.local/bin` 加入 PATH。

### Windows

Windows 内置 curl，但默认没有 bash。使用 Windows 包管理器 **Chocolatey**。

以**管理员**身份打开 PowerShell 并安装 Chocolatey：

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

安装 Chocolatey 后，一行安装 opencode：

```powershell
choco install opencode
```

然后用 `opencode --version` 检查。

### VS Code + OpenChamber 扩展

opencode CLI 完成。若终端 UI 不习惯，安装 VS Code + **OpenChamber** 扩展。

1. 若没有 VS Code，从 [code.visualstudio.com](https://code.visualstudio.com) 下载安装。
2. 打开 VS Code，点击扩展（Ctrl+Shift+X），搜索 **"OpenChamber"** 并安装。

有了 OpenChamber，可以在 VS Code 内像 Copilot 或 Cursor 一样使用 Cocrates。

![OpenChamber 扩展](/img/openchamber-extension.png)

---

## 🔧 步骤 2-3 — 插件配置 + Skill 文件

opencode 安装完成后，最简便的方式是让 AI 安装 Cocrates Harness。

运行 Cocrates 并请求：

> *"Install Cocrates Harness based on https://cocrates.ai/install.md."*

AI 会处理——从 GitHub 仓库下载 skills，将插件配置写入 opencode.jsonc，完成设置。

![Cocrates Harness 已安装](/img/cocrates-harness-installed.png)

**但是。那不是终点。**

若 AI 装完你就不管——仍是 AI辅助型。**审查并批准 AI 安装的内容**——才是 AI原生型工程师。

### 验证安装

Cocrates Harness 有两部分：**Plugin** 与 **Skill。** 两者都检查以确认设置。

**Plugin 检查。** 打开 `~/.config/opencode/opencode.jsonc`。是否有 `"plugin": ["@cocrates/cocrates-harness"]`？有则通过。

**Skill 检查。** 打开 `~/.config/opencode/skills/`。是否有 adr-writing、spec-writing、education 等 skill 文件？有则通过。

这两项检查足够。若 AI 做对了，你验证并批准。

若两者都正常，重启 opencode 并选择 Cocrates Agent。你已激活。

---

## 🦉 第一次对话——真正的部分现在开始

安装完成！你已就绪。该体验真东西了。

向 Cocrates 提出第一个问题：

> **"Tell me about Bloom's Taxonomy."**

一点提示：Cocrates 不会像典型 AI 那样表现。

多数 AI 会倾倒长篇解释。Cocrates 不会。它会 subtly 插入有瑕疵的例子，并交给你第一个 mission。

**安装不是终点。真正的部分现在开始。**

---

## 📌 要点回顾

1. **Cocrates 是 opencode 插件。** 运行在 opencode 上，由 Plugin + Skill 组成。
2. **推荐设置：opencode + VS Code + OpenChamber。** 在 VS Code 中 GUI 使用 Cocrates。
3. **让 AI 安装，然后验证。** 一行："Install based on install.md." 安装后确认 Plugin 与 Skill 就位。养成检查、审查 AI 工作的习惯。

---

**问自己：**

- opencode 安装了吗？
- Cocrates Agent 激活了吗？
- 准备好提第一个问题了吗？

若三者都是 yes，你已迈出 AI原生型工程师 的第一步。

---

## 🎬 下期预告

今天我们准备好了工具。我们安装了 Cocrates Harness，检查了 Plugin 与 Skill，准备好第一个问题。

但知道原则与践行不同。有工具与使用工具也不同。

接下来 EP4：与 Cocrates 的**第一次真实对话**。

当你问 "Tell me about Bloom's Taxonomy" 时，Cocrates 如何回应？我们将用真实对话走通三步学习流水线——Education → Knowledge Capture → Reflection。

与发人深省的 AI 的第一次相遇。值得期待。

---

*本系列介绍 Cocrates Harness 框架。Cocrates 是一个为苏格拉底式对话而设计的 agent harness，使用户保持主体性并持续成长。*
