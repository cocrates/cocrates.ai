---
name: generating-skill-creation
description: >-
  Select when the user asks to author or design an artifact-generation skill.
  Produces a SKILL.md that encodes a repeatable generation workflow — not the
  final artifact itself. Designs Snowflake stages
  (define → plan → architecture design → detail design → generation),
  per-component refinement, review points, and approval gates.
compatibility: opencode
metadata:
  agent: cocrates
---

# Generating Skill Creation — Artifact-Generation Skill Authoring

Follow this skill when the user requests **creating or designing an artifact generation skill**.

The goal is to design and write a **workflow skill** that produces a given artifact type in a repeatable, reviewable way — not to produce the final artifact itself.

## Core Principles

> **The unexamined artifact is not worth generating.**

- **Snowflake method:** Refine progressively in order: define → plan → architecture design → detail design → generation.
- **Per-stage artifacts:** Each stage leaves intermediate artifacts saved as files; proceed to the next stage only after user review and approval.
- **Design before generation:** Do not enter the generation stage until detail design is finalized.

## Working Location

```text
.opencode/skills/{skill-slug}/SKILL.md
```

---

## Snowflake: Progressive Refinement Stages

Design artifact-generation workflows as the five stages below.

```text
define → plan → architecture design → detail design → generation
```

| Stage | Purpose | Approval gate |
| ----- | ------- | ------------- |
| **define** | Lock assignment, scope, constraints, success criteria | Baseline for all later decisions |
| **plan** | Skeleton, direction, overview of key components | Approve overall rhythm and direction |
| **architecture design** | Per-component catalogs, hierarchy, cross-references | Approve structural consistency |
| **detail design** | Lock per-unit specs needed for generation | **Final gate before generation** |
| **generation** | Produce and assemble artifacts per locked specs | Per-unit and final quality review |

**Key rule:** Do not proceed to generation until detail design is fully finalized.

---

## Artifact Components

Decompose the final artifact into some or all of the dimensions below. Use domain-appropriate names; keep the roles the same.

| Dimension | Role |
| --------- | ---- |
| **Assignment & constraints** | Scope, quality bar, taboos, success conditions |
| **Context & rules** | Domain rules, worldbuilding, style guide |
| **Entities** | Reusable constituent elements |
| **Space & placement** | Locations, scenes, layout areas |
| **Structure & flow** | Hierarchical content organization |

When identifying components, confirm:

- Can the final artifact be split into independent files or units?
- Do components have cross-references or dependencies?
- Which components must be finalized first?
- Which components should be refined only when demand arises?

---

## Per-Stage Component Refinement

Define **which components are refined to what level** at each Snowflake stage. Fill in this table when designing a new skill to lock the workflow.

| Component | define | plan | architecture design | detail design | generation |
| --------- | ------ | ---- | ------------------- | ------------- | ---------- |
| **Assignment & constraints** | Final artifact form, scope, constraints, success criteria | — | — | — | — |
| **Context & rules** | Need and scope | Direction and tone summary | Rule-area catalog and relations | Per-area detailed rules and variants | — |
| **Entities** | Entity types | Core roles and relations summary | Catalog, relation map, hierarchy | Per-entity detailed spec, variants, state | Per-entity artifacts |
| **Space & placement** | Need and space types | Area and region overview | Catalog, hierarchy, links to structure | Per-space detailed spec, state, history | Space-related artifacts |
| **Structure & flow** | Top-level hierarchy | Upper-segment summary | Full hierarchy, unit catalog, cross-references | Per-unit detailed structure and content spec | Per-structure-unit artifacts, final assembly |

**Cross-references:** In architecture design, structure and flow reference entities and space. Reuse items already defined; if missing, register an empty slot and fill it in detail design.

**Refinement order:** assignment & constraints → context & rules → (entities ∥ space overview) → structure & flow → entity and space detail. When structure references entities or space, finalize catalogs first; fill detailed specs on demand.

---

## Meta Snowflake: Authoring Procedure for This Skill

Procedure for creating a **new generation skill** with this skill:

```text
Kernel → Components → Frame → Outline → Spec → Skill → Verification
```

| Step | Maps to | Output |
| ---- | ------- | ------ |
| **Kernel** | define | One-sentence definition of the generation target |
| **Components** | plan | Identify component dimensions |
| **Frame** | architecture design | Five-stage workflow, file structure, approval points |
| **Outline** | detail design | Per-stage file artifacts, inputs, completion criteria |
| **Spec → Skill → Verification** | generation | Write and verify `.opencode/skills/{skill-slug}/SKILL.md` |

### Outline: Per-File Artifacts

Define each intermediate and final artifact file as follows:

```markdown
### {file path}

- **Input**: Dependent artifacts
- **Creation activity**: Work performed to produce this file
- **Completion criteria**: When this file is considered done
- **Review questions**: Questions for user confirmation
- **Approval criteria**: Conditions to proceed to the next stage
```

### Rules for Authoring Generation Skills

**State in `description` which artifact-generation requests select the skill.** The Cocrates Agent finds and loads skills from the skill list using `description`. Do not add a `## When to Use` section in the body.

Include in `description`:

- **What** the skill does
- **Which artifact-generation requests** select it — artifact type, trigger terms
- Exclusion conditions when another skill is more appropriate (if needed)

### Generation Skill Body Structure

```markdown
# {Skill Title}

## Core Principles
## Working Location
## Component Definitions
## Snowflake Stages
## Stage-by-Stage Procedure
## Cross-Reference and Reuse Rules
## Dialogue Rules
## Prohibitions
## Completion Criteria
```

---

## Dialogue Rules

1. State the current step (Kernel, Components, Frame, Outline, Spec, Skill, Verification).
2. When designing Frame, present the **five Snowflake stages** and the **per-stage component refinement table** together.
3. Do not proceed to the next stage without user approval.
4. If components or Frame are unclear, ask questions before drafting the skill spec.

## Prohibitions

- Adding a `## When to Use` section in the body — selection criteria belong only in `description`
- Filling a `SKILL.md` template without component and per-stage refinement design
- Authoring a skill that proceeds to generation without detail design
- Leaving intermediate artifacts in chat only, without file-save rules
- Omitting `compatibility: opencode` and `metadata.agent: cocrates`

## Completion Criteria

- `.opencode/skills/{skill-slug}/SKILL.md` is created
- `description` states which artifact-generation requests select the skill (trigger terms)
- define → plan → architecture design → detail design → generation flow is specified
- Per-component, per-stage refinement content is specified
- Per-stage approval points and prohibitions are included
