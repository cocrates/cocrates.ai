---
name: knowledge-capture
description: >-
  Persists growth logs, discovered gaps, and new insights that the user has actually 
  understood — from an education session — into the knowledge base. Select when 
  the user's intent is to record, organize, or retain their learning journey and 
  reflection for later recall. Writes recall-focused reflection essentials to 
  kebab-case markdown files under kb/; merges into existing files when context matches.
compatibility: opencode
metadata:
  agent: cocrates
---

# Knowledge Capture — Reflection & Learning Log Storage Skill

When the Cocrates agent infers that the user wants to **record or retain what they learned and realized** — focus on capturing the user's cognitive gaps, new insights, and future learning paths rather than just storing passive encyclopedic knowledge.

The goal is to save **reflection-focused essentials** that document the user's growth and areas for future study, creating a continuous learning log.

Match the user's language (Korean, English, etc.) in user-facing messages. KB file content may use the language the user learned in; filenames and slugs stay English kebab-case.

## Storage Location

```
kb/
└── {topic-slug}.md
```

-- **KB root:** `kb/` at the project root (create if missing)
- **Filename:** English **kebab-case** for the topic (e.g. `dependency-inversion-principle.md`)

---

## Workflow

### 1. Search & Context Merge
Before saving, check the `kb/` directory. If a file with a similar context exists, **supplement and merge** the new reflection and gaps into it. Do not overwrite or delete existing learning history—append it as a cumulative growth log.

### 2. Extract and Write for Reflection
Do not store detailed lecture notes or full code. Based on the entire `education` dialogue, analyze and extract:
- **Discovered Gaps**: What did the user misunderstand or struggle with during the missions?
- **Aha Moments**: What core principles did they suddenly realize or grasp?
- **Next Actions**: What remains unexplored based on the current target level cut-off?

### 3. Notify User
After saving, briefly report the path and suggest a future milestone: *"I've logged your growth and next learning steps. Whenever you're ready to tackle the next level, just let me know!"*

---

## File Template (Revised for Reflection)

```markdown
# {Topic Title}

## 🎯 Target Level Achieved
- [e.g., Understanding level validated via Adapter Mission]

## 🧠 Discovered Gaps & Misconceptions
- {What the user initially got wrong or struggled with during the dialogue}
- {The logical contradiction identified during the Socratic exchange}

## 💡 Aha Moments & New Insights
- {Core realization or analogy the user grasped}
- {Key takeaway expressed in a recall-oriented, concise bullet (1–2 sentences max)}

## 🚀 Future Learning Roadmap (Next Action)
- {What to learn next to level up this concept, e.g., "Advance from Apply to Analyze by reviewing system boundary exceptions"}

## Related
- {Related kb file or concept, kebab-case link}

## Update History
- YYYY-MM-DD: {Brief 1-line summary of today's learning milestone}
```

---

## Merge & Reflection Example

**Existing** `kb/dip-principle.md`:
```markdown
# DIP

## 🧠 Discovered Gaps & Misconceptions
- Thought DIP was just about creating interfaces everywhere.

## 💡 Aha Moments & New Insights
- High-level modules must receive implementations via injection, not direct instantiation.
```

**After merging a new session focused on the Adapter Pattern:**

```markdown
# DIP

## 🎯 Target Level Achieved
- Applied level validated via Legacy Code Migration Mission (2026-06-29)

## 🧠 Discovered Gaps & Misconceptions
- Thought DIP was just about creating interfaces everywhere.
- *[Added]* Mistakenly tried to refactor external legacy code directly instead of wrapping it.

## 💡 Aha Moments & New Insights
- High-level modules must receive implementations via injection, not direct instantiation.
- *[Added]* Discovered that an Adapter Pattern acts as a perfect bridge to enforce DIP at system boundaries.

## 🚀 Future Learning Roadmap (Next Action)
- Analyze complex enterprise frameworks to evaluate how they handle dependency graph lifecycles automatically (Inversion of Control containers).

## Update History
- 2026-06-29: Supplemented with Adapter Pattern application and boundary refactoring insight.
```

---

## Prohibitions

- Storing long explanations, full tutorials, or raw dialogue transcripts.
- Writing generic AI summaries that ignore the user's specific errors, gaps, or unique realizations during the turn-based session.
- Erasing past reflection history during a merge (always append or update chronologically).
