# Install Cocrates Harness

This document describes how to install or upgrade **Cocrates Harness** on [OpenCode](https://opencode.ai). Use it when a user asks you to install Cocrates Harness: follow the steps below to configure the OpenCode plugin, install Node.js LTS and the required MCP packages globally, register the MCP servers (`cocrates-google-genai`, `cocrates-elk`, `cocrates-video-edit`) in `opencode.jsonc`, ensure `GEMINI_API_KEY` is available for `google-genai-mcp`, and copy the required skill files.

If Cocrates Harness is **already installed**, treat the request as an **upgrade** — do not skip version checks or overwrite skills without the user's consent.

## Prerequisites

Cocrates Harness is an **OpenCode plugin**, not a standalone program. OpenCode must be installed first.

- Install OpenCode from **https://opencode.ai/download**
- Supported clients: OpenCode terminal, OpenCode desktop, or the OpenCode extension for Cursor/VS Code
- Cocrates Harness works the same regardless of which client is used
- Media-generation skills (image, video, speech, music) require the `cocrates-google-genai` MCP server
- Diagram-generation skills (Excalidraw, draw.io) require the `cocrates-elk` MCP server
- Video-authoring (multi-clip assemble / final mp4) requires the `cocrates-video-edit` MCP server
- MCP servers require **Node.js LTS** and are installed as **global npm packages** (not via `npx`) — see Step 2
- `@cocrates/video-edit-mcp` requires Node.js **>= 20**; the other MCP packages require **>= 18**. Installing current Node.js LTS covers all three
- `cocrates-google-genai` (`google-genai-mcp`) requires a **Gemini API key** in `GEMINI_API_KEY` — see Step 2d

## What you will do

1. Detect whether Cocrates Harness is already installed
2. Add or upgrade the `@cocrates/cocrates-harness` plugin in the OpenCode config
3. Ensure Node.js LTS is installed; install or upgrade the three MCP packages globally; configure them in `opencode.jsonc`; ensure `GEMINI_API_KEY` is available for `google-genai-mcp`
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
| Node.js available | `node --version` and `npm --version` succeed (LTS; Node **>= 20** preferred) |
| MCP packages installed | `npm list -g --depth=0` shows `@cocrates/google-genai-mcp`, `@cocrates/elk-mcp`, and `@cocrates/video-edit-mcp` |
| MCP configured | `OPENCODE_CONFIG/opencode.jsonc` contains `cocrates-google-genai`, `cocrates-elk`, and `cocrates-video-edit` under `mcp`, each with a direct CLI command (not `npx`) |
| Gemini API key | `GEMINI_API_KEY` is set in the process environment, **or** under `mcp.cocrates-google-genai.environment` in `opencode.jsonc` |
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
        "google-genai-mcp"
      ]
    },
    "cocrates-elk": {
      "type": "local",
      "command": [
        "elk-mcp"
      ]
    },
    "cocrates-video-edit": {
      "type": "local",
      "command": [
        "video-edit-mcp"
      ]
    }
  }
}
```

Create the config directory and file if they do not exist. For config details, see **https://opencode.ai/docs/config/**

After the plugin CLI runs, continue to **Step 2** to install Node.js / MCP packages and ensure the `mcp` block is present — the plugin install command may not add MCP settings by itself.

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

## Step 2: Install Node.js, MCP packages, and configure MCP servers

Cocrates skills call MCP servers through OpenCode. Plugin + skills alone are not enough. You must:

1. Ensure **Node.js LTS** (and `npm`) are installed
2. Install or upgrade the three MCP packages as **global npm packages**
3. Point `opencode.jsonc` at the global CLI binaries (not `npx`)
4. Ensure `GEMINI_API_KEY` is available for `google-genai-mcp` (only configure it when missing — see Step 2d)

Do **not** use `npx -y @cocrates/...` in the MCP `command` arrays. Cold `npx` fetches often time out while OpenCode is starting MCP servers.

| MCP server key | npm package | Global CLI command | Used by |
|----------------|-------------|--------------------|---------|
| `cocrates-google-genai` | `@cocrates/google-genai-mcp` | `google-genai-mcp` | image, video, speech, music generation skills |
| `cocrates-elk` | `@cocrates/elk-mcp` | `elk-mcp` | diagram-generation (Excalidraw, draw.io layout) |
| `cocrates-video-edit` | `@cocrates/video-edit-mcp` | `video-edit-mcp` | video-authoring (edit-spec validate / render assemble) |

### 2a. Ensure Node.js LTS

1. Check whether Node.js and npm are available:

   ```bash
   node --version
   npm --version
   ```

2. If either command fails, or Node is older than **v20**, install **Node.js LTS** before continuing. Prefer the current LTS from **https://nodejs.org** (includes `npm`). Examples by platform — pick one that matches the user's OS and ask before changing system packages if needed:

   | Platform | Example |
   |----------|---------|
   | macOS (Homebrew) | `brew install node@lts` or `brew install node` |
   | Windows (winget) | `winget install OpenJS.NodeJS.LTS` |
   | Windows (Chocolatey) | `choco install nodejs-lts` |
   | Ubuntu / Debian | Install via [NodeSource](https://github.com/nodesource/distributions) or [nvm](https://github.com/nvm-sh/nvm) for the current LTS |
   | Any (nvm) | `nvm install --lts && nvm use --lts` |

3. Re-run `node --version` and `npm --version`. Confirm Node is **>= 20** (LTS). Report the installed versions to the user.

   | Before | After | Report |
   |--------|-------|--------|
   | Missing / too old | `{node}` / `{npm}` | `Node.js installed: {node} (npm {npm})` |
   | Already LTS (>= 20) | same | `Node.js already available: {node} (npm {npm})` |

### 2b. Install or upgrade MCP packages (global)

Fresh install and upgrade use the **same commands**. For **each** of the three packages, compare the locally installed global version with npm `latest`, then install or upgrade when needed — similar to the plugin version check in Step 1.

Packages:

- `@cocrates/google-genai-mcp`
- `@cocrates/elk-mcp`
- `@cocrates/video-edit-mcp`

For each package `{pkg}`:

1. Read the **current** global version (treat as `unknown` / missing if not installed):

   ```bash
   npm list -g {pkg} --depth=0
   ```

   Or parse JSON:

   ```bash
   npm list -g {pkg} --depth=0 --json
   ```

   Read `dependencies["{pkg}"].version` when present. If the package is absent, current version is `unknown`.

2. Read the **latest** version from the npm registry:

   ```bash
   npm view {pkg} version
   ```

3. Compare and act:

   | Current | Latest | Action |
   |---------|--------|--------|
   | `unknown` (not installed) | `{latest}` | `npm install -g {pkg}@latest` |
   | `{old}` | `{new}` (different) | `npm install -g {pkg}@latest` |
   | `{version}` | `{version}` (same) | Skip install — already up to date |

   Example (all three at once is fine when any need installing/upgrading):

   ```bash
   npm install -g @cocrates/google-genai-mcp@latest @cocrates/elk-mcp@latest @cocrates/video-edit-mcp@latest
   ```

4. After install/upgrade, confirm each CLI is on `PATH`:

   ```bash
   command -v google-genai-mcp
   command -v elk-mcp
   command -v video-edit-mcp
   ```

5. Re-read global versions with `npm list -g {pkg} --depth=0` and report **for each package**:

   | Before | After | Report |
   |--------|-------|--------|
   | `unknown` | `{version}` | `MCP package installed: {pkg}@{version}` |
   | `{old}` | `{new}` (different) | `MCP package upgraded: {pkg} {old} → {new}` |
   | `{version}` | `{version}` (same) | `MCP package already up to date: {pkg}@{version}` |
   | Could not determine | — | Report `npm list` / `npm view` output and whether install succeeded |

Always show explicit version numbers from `npm list -g` and `npm view`. On every install or upgrade request, **re-check** all three packages against npm latest — do not assume a previous global install is current.

### 2c. Configure `opencode.jsonc`

1. Open `OPENCODE_CONFIG/opencode.jsonc`.
2. If any server is missing, incomplete, or still uses `npx`, **merge** the following into the existing config (do not remove other plugins, MCP servers, or settings):

   ```jsonc
   "mcp": {
     "cocrates-google-genai": {
       "type": "local",
       "command": [
         "google-genai-mcp"
       ]
     },
     "cocrates-elk": {
       "type": "local",
       "command": [
         "elk-mcp"
       ]
     },
     "cocrates-video-edit": {
       "type": "local",
       "command": [
         "video-edit-mcp"
       ]
     }
   }
   ```

3. Confirm the file still has:
   - `"plugin"` entry for `@cocrates/cocrates-harness`
   - `"mcp"."cocrates-google-genai"` with `type: "local"` and `command: ["google-genai-mcp"]`
   - `"mcp"."cocrates-elk"` with `type: "local"` and `command: ["elk-mcp"]`
   - `"mcp"."cocrates-video-edit"` with `type: "local"` and `command: ["video-edit-mcp"]`

4. Report to the user **for each server**:

   | Before | After | Report |
   |--------|-------|--------|
   | Missing | Present | `MCP configured: {server-key}` |
   | Present (`npx` or other) | Updated to global CLI | `MCP updated: {server-key}` |
   | Present (same CLI command) | Present | `MCP already configured: {server-key}` |

   Use `{server-key}` = `cocrates-google-genai`, `cocrates-elk`, or `cocrates-video-edit`.

Preserve any other entries under `mcp` when merging. OpenCode starts these servers by invoking the global CLI on demand; the packages must already be installed via Step 2b.

### 2d. Ensure `GEMINI_API_KEY` for `google-genai-mcp`

Media-generation skills (image, video, speech, music) call `cocrates-google-genai`, which runs `google-genai-mcp`. That server needs a **Gemini API key** exposed as the environment variable **`GEMINI_API_KEY`**.

Create or manage a key at **https://aistudio.google.com/apikey** (Google AI Studio).

**Check first — configure only when missing.** Do not overwrite an existing key.

1. Check the **shell / process environment**:

   ```bash
   # Prints nothing if unset; do not echo the value in chat if it is set
   [ -n "${GEMINI_API_KEY:-}" ] && echo "GEMINI_API_KEY is set" || echo "GEMINI_API_KEY is unset"
   ```

2. Check **OpenCode MCP config** — in `OPENCODE_CONFIG/opencode.jsonc`, look for a non-empty value under:

   ```
   mcp.cocrates-google-genai.environment.GEMINI_API_KEY
   ```

   (OpenCode’s local MCP field is `environment`, not `env`.)

3. Decide:

   | Shell `GEMINI_API_KEY` | MCP `environment.GEMINI_API_KEY` | Action |
   |------------------------|----------------------------------|--------|
   | Set | — | Already available (MCP inherits the OpenCode process env). Report and skip. |
   | Unset | Set (non-empty) | Already configured in `opencode.jsonc`. Report and skip. |
   | Unset | Missing / empty | **Ask the user** for a Gemini API key, then set it (below). |

4. When the key is missing, ask the user for the value (do not invent one). Then set it in **one** of these ways — prefer the user’s choice; default to MCP `environment` if they have no preference:

   **Option A — MCP `environment` in `opencode.jsonc`** (scoped to this server; works even if the shell profile has no key):

   ```jsonc
   "cocrates-google-genai": {
     "type": "local",
     "command": [
       "google-genai-mcp"
     ],
     "environment": {
       "GEMINI_API_KEY": "{user-provided-key}"
     }
   }
   ```

   You may also reference a shell variable with OpenCode interpolation if the user prefers not to store the raw key in the file:

   ```jsonc
   "environment": {
     "GEMINI_API_KEY": "{env:GEMINI_API_KEY}"
   }
   ```

   In that case the shell variable must still be set where OpenCode is launched.

   **Option B — Shell / user environment** (e.g. `~/.bashrc`, `~/.zshrc`, or the OS user environment), so OpenCode and child MCP processes inherit it:

   ```bash
   export GEMINI_API_KEY="{user-provided-key}"
   ```

   Tell the user to restart the terminal / OpenCode after changing the profile.

5. Report to the user:

   | Before | After | Report |
   |--------|-------|--------|
   | Already in shell | Unchanged | `GEMINI_API_KEY already set in environment` |
   | Already in MCP `environment` | Unchanged | `GEMINI_API_KEY already set in mcp.cocrates-google-genai.environment` |
   | Missing | Set via MCP or shell | `GEMINI_API_KEY configured` (say which option was used) |
   | Missing and user declined | Still missing | Warn that media-generation skills will fail until the key is set |

Never print the full API key in chat or in install summaries. `elk-mcp` and `video-edit-mcp` do not require `GEMINI_API_KEY`.

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
4. Confirm Node.js LTS is available, the three MCP CLIs resolve on `PATH` (`google-genai-mcp`, `elk-mcp`, `video-edit-mcp`), and `opencode.jsonc` still includes `mcp.cocrates-google-genai`, `mcp.cocrates-elk`, and `mcp.cocrates-video-edit` with those direct commands (not `npx`).
5. Confirm `GEMINI_API_KEY` is available via the shell environment or `mcp.cocrates-google-genai.environment` (Step 2d). Without it, media-generation skills cannot call Gemini.

Installation or upgrade is complete when the user can select the **Cocrates** agent in OpenCode, `opencode.jsonc` has the plugin and all three MCP entries, the corresponding global MCP packages are at the npm latest versions checked in Step 2b, and `GEMINI_API_KEY` is configured for `google-genai-mcp`.

---

## Example user requests

> Install Cocrates Harness using this document.

Perform Step 0, then Steps 1–4. On a machine that already has Cocrates Harness, upgrade the plugin, re-check MCP package versions against npm latest (install/upgrade globals as needed), ensure MCP config uses the global CLI commands, ensure `GEMINI_API_KEY` is present only if missing, and reconcile skills instead of overwriting blindly.

> Upgrade Cocrates Harness.

Same procedure — always run Step 0 first, then Step 1 (plugin), Step 2 (Node.js + MCP package version check + config + `GEMINI_API_KEY` check), and skill reconciliation (Step 3b) when skills already exist.
