# Step ① — PRD Creation

## Purpose

Define the product anchor: what to build, for whom, why now, and the global constraints that bind all downstream decisions. The PRD is the **top-level master document** — every ASR, ADR, and Spec traces back to it.

## Prerequisites

- Step 0 resolved `{project-root}`.
- `{project-root}/spec/` directory exists (create if missing).

## Procedure

### 1. Identify Goal

Clarify through Socratic dialogue:
- **Artifact type:** What kind of deliverable? (software, document, image, presentation, etc.)
- **Target audience:** Who will use or consume it?
- **Core function:** What job does it do?
- **Output location:** Where should the final artifact live?

### 2. Create Project Folder

Ensure `{project-root}` and `{project-root}/spec/` exist. If creating a new folder, confirm **location and name** with the user first.

### 3. Document PRD

Write or update **`{project-root}/spec/PRD.md`** using the template below.

### 4. Initialize ASR Registry

Create **`{project-root}/spec/ASR.md`** if missing (see SKILL.md § ASR.md File for template), even before ASRs are fully discovered.

### 5. Obtain PRD Approval

Present the PRD and get **explicit user approval**. Do not proceed without it.

## Gate

- `{project-root}/spec/PRD.md` exists ✓
- User has explicitly approved it ✓

**Next → Step ② (ASR Identification).**

---

## PRD Template

```markdown
# Product Requirement Definition

## Product Vision
{One-paragraph summary: what this is, who it's for, what problem it solves.}

## Target Audience
{Primary users/readers/consumers — be specific about their expertise level and expectations.}

## Core Deliverables
- {Primary artifact type and format}
- {Secondary artifacts if any}

## High-Level Goals
- {Goal 1: measurable or verifiable}
- {Goal 2}
- {Goal 3}

## Global Constraints
- {Hard limits: language, platform, deadline, brand, budget, dependencies}
- {Quality bar: minimum acceptance criteria}

## Out of Scope
- {Explicit non-goals to prevent feature creep}

## Success Criteria
- {How we know this is done and good enough}
```

Omit empty sections except **Product Vision** and **Core Deliverables**.
