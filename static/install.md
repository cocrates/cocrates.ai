# Install Cocrates Harness

This document describes how to install **Cocrates Harness** on [OpenCode](https://opencode.ai). Use it when a user asks you to install Cocrates Harness: follow the steps below to configure the OpenCode plugin and copy the required skill files.

## Prerequisites

Cocrates Harness is an **OpenCode plugin**, not a standalone program. OpenCode must be installed first.

- Install OpenCode from **https://opencode.ai/download**
- Supported clients: OpenCode terminal, OpenCode desktop, or the OpenCode extension for Cursor/VS Code
- Cocrates Harness works the same regardless of which client is used

## What you will do

1. Add the `@cocrates/cocrates-harness` plugin to the OpenCode config
2. Copy Cocrates skill files into the OpenCode skills directory
3. Ask the user to restart OpenCode and confirm the **Cocrates** agent is available

---

## Step 1: Configure the OpenCode plugin

OpenCode reads configuration from:

- **Default:** `~/.config/opencode/`
- **Override:** the directory set by the `OPENCODE_CONFIG_DIR` environment variable

Edit `opencode.jsonc` in that directory. Add `"@cocrates/cocrates-harness"` to the `plugin` array. If other plugins are already listed, append this entry without removing them.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": ["@cocrates/cocrates-harness"]
}
```

Create the config directory and file if they do not exist.

For config details, see **https://opencode.ai/docs/config/**

---

## Step 2: Copy skill files

The plugin registers the **Cocrates** agent, but the agent needs skill files to run workflows. Copy all skills from the Cocrates repository into OpenCode's skills directory.

**Source:**

```
https://github.com/cocrates/cocrates.ai/tree/main/skills
```

**Destination:**

```
~/.config/opencode/skills/
```

Copy each skill as a subdirectory containing its `SKILL.md` (and any supporting files in that skill folder). Expected layout after copying:

```
~/.config/opencode/
├── opencode.jsonc
└── skills/
    ├── education/SKILL.md
    ├── knowledge-capture/SKILL.md
    ├── reflection/SKILL.md
    ├── adr-writing/SKILL.md
    ├── spec-writing/SKILL.md
    ├── spec-driven-generation/SKILL.md
    ├── spec-driven-verification/SKILL.md
    └── ...
```

Core skills include:

| Skill | Purpose |
|-------|---------|
| `education` | Socratic 1:1 learning coach |
| `knowledge-capture` | Save learning outcomes to a knowledge base |
| `reflection` | Assess understanding |
| `adr-writing` | Record architectural decisions |
| `spec-writing` | Consolidate requirements into a spec |
| `spec-driven-generation` | Generate deliverables from a spec |
| `spec-driven-verification` | Verify deliverables against a spec |

Copy **all** directories under `skills/` from the repository, not only the table above.

You can clone the repository and copy the folder, or download individual skill directories from GitHub.

---

## Step 3: Verify installation

After configuration and skill copy are complete:

1. Tell the user to **restart OpenCode** (quit and reopen the terminal app, desktop app, or editor extension).
2. Ask them to open the agent picker and confirm **Cocrates** appears and can be selected.
3. If Cocrates is not listed, check that `opencode.jsonc` includes the plugin entry and that OpenCode was fully restarted.

Installation is complete when the user can select the **Cocrates** agent in OpenCode.

---

## Example user request

> Install Cocrates Harness using this document.

When you receive that request, perform Step 1 and Step 2 on the user's machine, then guide them through Step 3.
