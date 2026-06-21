---
name: knowledge-capture
description: >-
  Persists knowledge the user has actually understood and confirmed — from an
  education session or equivalent learning dialogue — into the knowledge base.
  Select when the user's intent is to record, organize, or retain what they
  learned for later recall, not merely to get a summary. Skill selection follows
  inferred intent, not specific trigger words. Writes recall-focused essentials
  to kebab-case markdown files under kb/; merges into existing files when context
  matches.
compatibility: opencode
metadata:
  agent: cocrates
---

# Knowledge Capture — Learning Knowledge Storage Skill

When the Cocrates agent infers that the user wants to **record or retain what they learned** — not merely receive a summary — follow the instructions in this skill.

Persist **what the user actually understood and confirmed** through the `education` skill (or an equivalent learning dialogue and missions) into the knowledge base (KB).

The goal is not long-form explanation but **recall-focused essentials** the user can retrieve later.

Match the user's language (Korean, English, etc.) in user-facing messages. KB file content may use the language the user learned in; filenames and slugs stay English kebab-case.

## Storage Location

```
kb/
└── {topic-slug}.md
```

- **KB root:** `kb/` at the project root (create if missing)
- **Filename:** English **kebab-case** for the topic (e.g. `dependency-inversion-principle.md`, `bloom-taxonomy.md`)
- One file = one concept, topic, or context unit

## When to Use

Select this skill when the user's intent is to **capture or organize learned knowledge** for later recall — for example:

- Right after an `education` session ends (understanding confirmed, MISSION completed)
- When the user wants to preserve insights from a learning exercise they worked through
- When the user explicitly wants what they learned written down in the KB

**Before saving:** Save only when there is evidence the user understood the content (explained in their own words, completed a MISSION, explicit approval, etc.). Do not put AI-generated content alone into the KB without user understanding.

---

## Workflow

### 1. Search Existing KB

Before saving, check the `kb/` directory.

- Search for files with the **same or similar context** (compare filename, title, `## Tags`).
- If found, **do not create a new file** — **supplement and merge** into the existing file.
- If not found, create a new `{topic-slug}.md`.

**Similar-context examples:**
- `dip-principle.md` ↔ `dependency-inversion-principle.md` → same topic, merge into one
- Add a DIP section to `solid-principles.md` vs separate `dip-principle.md` → merge or cross-link depending on scope

### 2. Write for Recall

Do not store detailed lecture notes, long explanations, or full code. Keep only **essentials to recall later**.

Writing principles:
- One bullet = one memory unit (1–2 sentences max)
- Focus on analogies, one-line definitions, common mistakes, and gaps the user discovered
- Code: only **minimal snippets** needed for recall (optional)
- Do not store full education dialogue logs

### 3. Save File

When supplementing an existing file:
- Remove duplicate bullets
- Add new insights to the appropriate section
- Optionally record date and one-line summary under `## Update History`

### 4. Notify User

After saving, briefly report the path and what was added or merged.
And you may suggest a `reflection` session: *"Shall we run an understanding review based on this content?"*

---

## File Template

```markdown
# {Topic Title}

## One-line Definition
{Recall-oriented definition or analogy}

## Key Points
- {Point to remember 1}
- {Point to remember 2}
- {Point to remember 3}

## Wrong Assumptions / Gaps
- {Misconception clarified during learning}

## Related
- {Related kb file or concept, kebab-case link}

## Tags
`tag-one`, `tag-two`

<!-- Optional: update history -->
## Update History
- YYYY-MM-DD: {What was supplemented}
```

Omit empty sections. **Key Points** is required.

---

## Merge Example

**Existing** `kb/dip-principle.md`:
```markdown
# DIP

## Key Points
- High-level modules must not depend directly on low-level implementations
```

**Insight from this session:** Relationship to the adapter pattern

**After merge:**
```markdown
# DIP

## Key Points
- High-level modules must not depend directly on low-level implementations
- Depend on interfaces (abstractions); receive implementations via injection
- Gap with legacy code: bridge with an adapter at the boundary

## Wrong Assumptions / Gaps
- DIP = always create a new interface file first (X → adapter on existing code is valid)

## Update History
- 2026-06-15: Adapter pattern relationship, Gap supplement
```

---

## Prohibitions

- Saving AI-generated summaries without confirmed user understanding
- Creating duplicate files for the same context
- Storing long explanations, full tutorials, or complete dialogue transcripts
- Overwriting KB without user approval (on merge: supplement only, never delete existing content)

## Completion Criteria

- Recall-focused essentials saved in `kb/{topic-slug}.md`
- User informed whether the file was new or merged
