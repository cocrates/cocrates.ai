# Install Cocrates Harness

This document describes how to install or upgrade **Cocrates Harness** on [OpenCode](https://opencode.ai). Use it when a user asks you to install Cocrates Harness: follow the steps below to configure the OpenCode plugin, register the Google GenAI MCP server, and copy the required skill files.

If Cocrates Harness is **already installed**, treat the request as an **upgrade** — do not skip version checks or overwrite skills without the user's consent.

## Prerequisites

Cocrates Harness is an **OpenCode plugin**, not a standalone program. OpenCode must be installed first.

- Install OpenCode from **https://opencode.ai/download**
- Supported clients: OpenCode terminal, OpenCode desktop, or the OpenCode extension for Cursor/VS Code
- Cocrates Harness works the same regardless of which client is used
- Media-generation skills (image, video, speech, music) require the `cocrates-google-genai` MCP server and a working `npx` (Node.js) on the machine

## What you will do

1. Detect whether Cocrates Harness is already installed
2. Add or upgrade the `@cocrates/cocrates-harness` plugin in the OpenCode config
3. Ensure the `cocrates-google-genai` MCP server is configured in `opencode.jsonc`
4. Copy or reconcile Cocrates skill files in the OpenCode skills directory
5. Ask the user to restart OpenCode and confirm the **Cocrates** agent is available

---

## Step 0: Detect existing installation

OpenCode reads configuration from:

- **Default:** `~/.config/opencode/`
- **Override:** the directory set by the `OPENCODE_CONFIG_DIR` environment variable

Call this directory `OPENCODE_CONFIG` below.

Check for an existing installation:

| Signal | Location |
|--------|----------|
| Plugin configured | `OPENCODE_CONFIG/opencode.jsonc` contains `@cocrates/cocrates-harness` in the `plugin` array |
| Plugin cached | `~/.cache/opencode/packages/` contains a directory matching `*cocrates-harness*` |
| MCP configured | `OPENCODE_CONFIG/opencode.jsonc` contains `cocrates-google-genai` under `mcp` |
| Skills present | `OPENCODE_CONFIG/skills/` contains one or more Cocrates skill subdirectories (e.g. `education/`, `spec-writing/`) |

If **skills** are already present in Step 0, use skill **reconciliation** in Step 3b. Otherwise, use Step 3a. **Steps 1–2 are the same** for both fresh install and upgrade.

---

## Step 1: Install or upgrade the OpenCode plugin

Fresh install and upgrade use the **same command**. The only difference is whether a previous version exists in the cache for comparison.

The CLI command adds `"@cocrates/cocrates-harness"` to `opencode.jsonc` automatically (`-g`). You can also edit the config manually if needed. The **complete** target config (plugin + MCP) looks like this:

```jsonc
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": [
    "@cocrates/cocrates-harness"
  ],
  "mcp": {
    "cocrates-google-genai": {
      "type": "local",
      "command": [
        "npx",
        "-y",
        "@cocrates/google-genai-mcp"
      ]
    }
  }
}
```

Create the config directory and file if they do not exist. For config details, see **https://opencode.ai/docs/config/**

After the plugin CLI runs, continue to **Step 2** to ensure the `mcp` block is present — the plugin install command may not add MCP settings by itself.

1. Read the **current version** from the OpenCode cache — do **not** rely on `npm` being installed:

   ```bash
   find ~/.cache/opencode/packages -maxdepth 1 -type d -name '*cocrates-harness*'
   ```

   Read `version` from `package.json` inside the matching directory. Typical locations:

   ```
   ~/.cache/opencode/packages/@cocrates+cocrates-harness@0.1.4/package.json
   ~/.cache/opencode/packages/@cocrates/cocrates-harness@0.1.4/package.json
   ```

   Notes:
   - OpenCode caches npm plugins under `~/.cache/opencode/packages/`. The exact directory name may use `+` or `/` in the scope — use `find` rather than assuming one format.
   - If several cache directories exist, prefer the one referenced by your config or the most recently modified.
   - If a version pin exists in `opencode.jsonc` (e.g. `"@cocrates/cocrates-harness@0.1.2"`), record it as additional context.
   - If no cache entry exists, treat the current version as `unknown` and continue.

2. **Install or upgrade** with the OpenCode CLI (OpenCode fetches the latest release — no `npm` CLI required):

   ```bash
   opencode plugin @cocrates/cocrates-harness -g -f
   ```

   | Flag | Meaning |
   |------|---------|
   | `-g` / `--global` | Install into the global OpenCode config (`~/.config/opencode/`) |
   | `-f` / `--force` | Replace an existing plugin version and fetch the latest release |

3. Re-read `version` from `package.json` in `~/.cache/opencode/packages/` after the command completes.

4. Compare before and after, then report to the user:

   | Before | After | Report |
   |--------|-------|--------|
   | `unknown` | `{version}` | `Plugin installed: {version}` |
   | `{old}` | `{new}` (different) | `Plugin upgraded: {old} → {new}` |
   | `{version}` | `{version}` (same) | `Plugin already up to date: {version}` |
   | Could not determine | — | Report what you found in the cache and whether the command succeeded |

Always show explicit version numbers read from `~/.cache/opencode/packages/`. Do **not** use `npm view` or other npm CLI commands — `npm` may not be installed on the user's machine.

---

## Step 2: Configure the Google GenAI MCP server

Cocrates media skills call **`@cocrates/google-genai-mcp`** through OpenCode MCP. Plugin + skills alone are not enough — `opencode.jsonc` must include the `cocrates-google-genai` server.

1. Open `OPENCODE_CONFIG/opencode.jsonc`.
2. If `mcp.cocrates-google-genai` is missing or incomplete, **merge** the following into the existing config (do not remove other plugins, MCP servers, or settings):

   ```jsonc
   "mcp": {
     "cocrates-google-genai": {
       "type": "local",
       "command": [
         "npx",
         "-y",
         "@cocrates/google-genai-mcp"
       ]
     }
   }
   ```

3. Confirm the file still has both:
   - `"plugin"` entry for `@cocrates/cocrates-harness`
   - `"mcp"."cocrates-google-genai"` with `type: "local"` and the `npx -y @cocrates/google-genai-mcp` command array above

4. Report to the user:

   | Before | After | Report |
   |--------|-------|--------|
   | Missing | Present | `MCP configured: cocrates-google-genai` |
   | Present (same) | Present | `MCP already configured: cocrates-google-genai` |
   | Present (different command) | Updated to match | `MCP updated: cocrates-google-genai` |

Preserve any other entries under `mcp` when merging. OpenCode starts this server via `npx` on demand; no separate global install of `@cocrates/google-genai-mcp` is required.

---

## Step 3: Copy or reconcile skill files

The plugin registers the **Cocrates** agent, but the agent needs skill files to run workflows.

**Official source (always use this as the reference for upgrades):**

```
https://github.com/cocrates/cocrates.ai/tree/main/skills
```

**Destination:**

```
OPENCODE_CONFIG/skills/
```

Expected layout after copying:

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
    ├── generating-skill-creation/SKILL.md
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
| `generating-skill-creation` | Create skill to generate an artifact |

Copy **all** directories under `skills/` from the repository, not only the table above.

### 3a. Fresh install

When no Cocrates skills were present in Step 0:

1. Clone or download the official `skills/` tree from the repository.
2. Copy every skill subdirectory into `OPENCODE_CONFIG/skills/`.
3. Report which skills were installed.

### 3b. Upgrade (skills already present)

Skills may contain user edits. **Never blindly overwrite** the entire skills directory.

1. **Fetch the official skills tree** into a temporary directory (shallow clone is fine):

   ```bash
   git clone --depth 1 https://github.com/cocrates/cocrates.ai /tmp/cocrates.ai
   ```

2. **Enumerate skills** — union of:
   - subdirectories under `/tmp/cocrates.ai/skills/`
   - subdirectories under `OPENCODE_CONFIG/skills/` that look like Cocrates skills (contain `SKILL.md` or match known skill names)

3. **For each skill**, compare **every file** under the skill directory tree — not only `SKILL.md`. A skill may include supporting files and subdirectories (e.g. `workflow/01-define.md`, `renderer/marp.md`). Compare the full folder recursively:

   ```bash
   diff -rq OPENCODE_CONFIG/skills/{skill-name}/ /tmp/cocrates.ai/skills/{skill-name}/
   ```

   Also list all files explicitly when building the report:

   ```bash
   find OPENCODE_CONFIG/skills/{skill-name}/ -type f | sort
   find /tmp/cocrates.ai/skills/{skill-name}/ -type f | sort
   ```

   For each differing file, run a content diff to identify what changed:

   ```bash
   diff -u OPENCODE_CONFIG/skills/{skill-name}/{path} /tmp/cocrates.ai/skills/{skill-name}/{path}
   ```

   Classify each skill:

   | Status | Meaning |
   |--------|---------|
   | **identical** | All files match — no action needed |
   | **missing locally** | Exists in official repo but not installed — safe to add |
   | **missing upstream** | Exists locally but not in official repo — likely user-only; do not delete without asking |
   | **different** | One or more files added, removed, or modified — **requires user decision** |

   When a skill is **different**, record per-file changes:

   | Change type | Example |
   |-------------|---------|
   | **modified** | `SKILL.md`, `workflow/05-generation.md` |
   | **added** (official only) | `workflow/09-release.md` |
   | **removed** (local only) | `notes/local-tweaks.md` |

4. **Report a summary table** to the user before changing anything:

   ```
   Skill reconciliation summary
   ────────────────────────────
   ✓ education                     — up to date (4 files)
   + presentation-authoring        — new (will be added, 6 files)
   ⚠ spec-writing                  — differs (3 files changed)
       modified: SKILL.md, workflow/02-plan.md
       added:    workflow/09-release.md
   ⚠ novel-authoring               — differs (12 files changed; local edits suspected)
   ```

5. **For skills marked `different`**, show a concise per-file diff summary so the user can choose wisely:
   - List every **added**, **removed**, and **modified** file path under the skill folder.
   - For each modified text file, include a short description of what changed (e.g. first few diff hunks, or a one-line summary per file: "workflow step reordered", "new evaluation criteria added").
   - Do **not** dump entire files unless the user asks.

6. **Ask the user for each differing skill** (or offer a batch choice):

   > **{skill-name}** differs from the official version.
   >
   > Summary: {brief description of changes}
   >
   > - **Use official** — replace your local copy with the repository version
   > - **Keep local** — keep your current copy unchanged
   >
   > Which do you prefer?

   Accept answers like "official for spec-writing, keep local for novel-authoring" or "use official for all".

7. **Apply the user's decisions**:

   | Decision | Action |
   |----------|--------|
   | Use official | Copy `/tmp/cocrates.ai/skills/{skill-name}/` → `OPENCODE_CONFIG/skills/{skill-name}/` (replace) |
   | Keep local | Leave `OPENCODE_CONFIG/skills/{skill-name}/` unchanged |
   | New skill (missing locally) | Copy from official without asking, unless the user opted out of new skills |
   | Identical | Skip |
   | Local-only skill | Keep unless the user asks to remove it |

8. **Report the outcome**:

   ```
   Skills updated
   ──────────────
   Added:     presentation-authoring
   Updated:   spec-writing (official)
   Unchanged: education, reflection, adr-writing, …
   Kept local: novel-authoring
   ```

9. Remove the temporary clone when finished:

   ```bash
   rm -rf /tmp/cocrates.ai
   ```

You can clone the repository and copy folders, or download individual skill directories from GitHub. For upgrades, a shallow clone plus `diff` is preferred because it makes per-skill comparison reliable.

---

## Step 4: Verify installation

After plugin, MCP, and skill reconciliation are complete:

1. Tell the user to **restart OpenCode** (quit and reopen the terminal app, desktop app, or editor extension).
2. Ask them to open the agent picker and confirm **Cocrates** appears and can be selected.
3. If Cocrates is not listed, check that `opencode.jsonc` includes the plugin entry and that OpenCode was fully restarted.
4. Confirm `opencode.jsonc` still includes `mcp.cocrates-google-genai` after restart (needed for image / video / speech / music generation skills).

Installation or upgrade is complete when the user can select the **Cocrates** agent in OpenCode and `opencode.jsonc` has both the plugin and the `cocrates-google-genai` MCP entry.

---

## Example user requests

> Install Cocrates Harness using this document.

Perform Step 0, then Steps 1–4. On a machine that already has Cocrates Harness, upgrade the plugin, ensure MCP is configured, and reconcile skills instead of overwriting blindly.

> Upgrade Cocrates Harness.

Same procedure — always run Step 0 first, then Step 1 (plugin), Step 2 (MCP), and skill reconciliation (Step 3b) when skills already exist.
