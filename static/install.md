# Install Cocrates Harness

This document describes how to install or upgrade **Cocrates Harness** on [OpenCode](https://opencode.ai). Use it when a user asks you to install Cocrates Harness: follow the steps below to configure the OpenCode plugin and copy the required skill files.

If Cocrates Harness is **already installed**, treat the request as an **upgrade** — do not skip version checks or overwrite skills without the user's consent.

## Prerequisites

Cocrates Harness is an **OpenCode plugin**, not a standalone program. OpenCode must be installed first.

- Install OpenCode from **https://opencode.ai/download**
- Supported clients: OpenCode terminal, OpenCode desktop, or the OpenCode extension for Cursor/VS Code
- Cocrates Harness works the same regardless of which client is used

## What you will do

1. Detect whether Cocrates Harness is already installed
2. Add or upgrade the `@cocrates/cocrates-harness` plugin in the OpenCode config
3. Copy or reconcile Cocrates skill files in the OpenCode skills directory
4. Ask the user to restart OpenCode and confirm the **Cocrates** agent is available

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
| Plugin cached | `~/.cache/opencode/packages/` contains a `@cocrates/cocrates-harness@*` directory |
| Skills present | `OPENCODE_CONFIG/skills/` contains one or more Cocrates skill subdirectories (e.g. `education/`, `spec-writing/`) |

If **any** of the above is true, proceed with the **upgrade flow** in Steps 1–2. Otherwise, perform a **fresh install**.

---

## Step 1: Configure or upgrade the OpenCode plugin

Edit `opencode.jsonc` in `OPENCODE_CONFIG`. Add `"@cocrates/cocrates-harness"` to the `plugin` array if it is missing. If other plugins are already listed, append this entry without removing them.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": ["@cocrates/cocrates-harness"]
}
```

Create the config directory and file if they do not exist.

For config details, see **https://opencode.ai/docs/config/**

### 1a. Fresh install

When the plugin entry was **not** present in Step 0:

1. Ensure `"@cocrates/cocrates-harness"` is in the `plugin` array.
2. Run:

   ```bash
   opencode plugin add @cocrates/cocrates-harness -f
   ```

3. Report the installed version (see **Version reporting** below).

### 1b. Upgrade (already installed)

When the plugin entry **was** present in Step 0:

1. Determine the **currently installed version** (try in order):
   - Read `version` from `package.json` inside `~/.cache/opencode/packages/@cocrates/cocrates-harness@*/` (use the most recently modified directory if several exist).
   - If a version pin exists in `opencode.jsonc` (e.g. `"@cocrates/cocrates-harness@0.1.2"`), use that version.
   - If neither is available, report `unknown`.

2. Determine the **latest published version**:

   ```bash
   npm view @cocrates/cocrates-harness version
   ```

3. Compare versions and report to the user **before** upgrading:

   | Situation | Action |
   |-----------|--------|
   | Installed `<` latest | Tell the user: `Upgrading plugin: {installed} → {latest}`. Then upgrade. |
   | Installed `=` latest | Tell the user: `Plugin is already at the latest version ({latest}). No plugin upgrade needed.` Skip the upgrade command. |
   | Installed `>` latest or comparison unclear | Explain what you found and ask the user how to proceed. |

4. To upgrade when a newer version is available:

   ```bash
   opencode plugin add @cocrates/cocrates-harness -f
   ```

   If the command does not refresh the cache, also remove the stale cache and retry:

   ```bash
   rm -rf ~/.cache/opencode/packages/@cocrates/cocrates-harness@latest
   opencode plugin add @cocrates/cocrates-harness -f
   ```

5. After upgrade, re-read the cached `package.json` and confirm the new version. Report: `Plugin upgraded: {old} → {new}` or `Plugin version unchanged: {version}`.

### Version reporting

Always show the user explicit version numbers, for example:

- `Plugin installed: 0.1.4`
- `Plugin upgraded: 0.1.2 → 0.1.4`
- `Plugin already up to date: 0.1.4`

---

## Step 2: Copy or reconcile skill files

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

### 2a. Fresh install

When no Cocrates skills were present in Step 0:

1. Clone or download the official `skills/` tree from the repository.
2. Copy every skill subdirectory into `OPENCODE_CONFIG/skills/`.
3. Report which skills were installed.

### 2b. Upgrade (skills already present)

Skills may contain user edits. **Never blindly overwrite** the entire skills directory.

1. **Fetch the official skills tree** into a temporary directory (shallow clone is fine):

   ```bash
   git clone --depth 1 https://github.com/cocrates/cocrates.ai /tmp/cocrates.ai
   ```

2. **Enumerate skills** — union of:
   - subdirectories under `/tmp/cocrates.ai/skills/`
   - subdirectories under `OPENCODE_CONFIG/skills/` that look like Cocrates skills (contain `SKILL.md` or match known skill names)

3. **For each skill**, compare the installed copy with the official copy:

   ```bash
   diff -rq OPENCODE_CONFIG/skills/{skill-name}/ /tmp/cocrates.ai/skills/{skill-name}/
   ```

   Classify each skill:

   | Status | Meaning |
   |--------|---------|
   | **identical** | No differences — no action needed |
   | **missing locally** | Exists in official repo but not installed — safe to add |
   | **missing upstream** | Exists locally but not in official repo — likely user-only; do not delete without asking |
   | **different** | Files differ — **requires user decision** |

4. **Report a summary table** to the user before changing anything:

   ```
   Skill reconciliation summary
   ────────────────────────────
   ✓ education                     — up to date
   + presentation-authoring        — new (will be added)
   ⚠ spec-writing                  — differs (3 files changed)
   ⚠ novel-authoring               — differs (12 files changed; local only edits suspected)
   ```

5. **For skills marked `different`**, show a concise diff summary so the user can choose wisely:
   - List changed, added, and removed file paths.
   - For text files, include a short excerpt of what changed (e.g. first few diff hunks or a one-line description per file: "workflow step reordered", "new evaluation criteria added").
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

## Step 3: Verify installation

After configuration and skill reconciliation are complete:

1. Tell the user to **restart OpenCode** (quit and reopen the terminal app, desktop app, or editor extension).
2. Ask them to open the agent picker and confirm **Cocrates** appears and can be selected.
3. If Cocrates is not listed, check that `opencode.jsonc` includes the plugin entry and that OpenCode was fully restarted.

Installation or upgrade is complete when the user can select the **Cocrates** agent in OpenCode.

---

## Example user requests

> Install Cocrates Harness using this document.

Perform Step 0, then Steps 1–3. On a machine that already has Cocrates Harness, upgrade the plugin and reconcile skills instead of overwriting blindly.

> Upgrade Cocrates Harness.

Same procedure — always run Step 0 first, then plugin version check (Step 1b) and skill reconciliation (Step 2b).
