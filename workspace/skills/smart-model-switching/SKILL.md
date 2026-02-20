---
name: smart-model-switching
description: Auto-route tasks to the cheapest model that works. Three-tier progression Haiku→Sonnet→Opus. Save 50-90% on API costs by starting cheap and escalating only when needed.
---

# Smart Model Switching

**Haiku → Sonnet → Opus**: Start cheap, escalate only when needed. Save 50-90% on API costs.

## Decision Tree

```
Greeting, lookup, status check, 1-2 sentence answer?
  YES → HAIKU
  NO ↓
Code, analysis, planning, writing, multi-step?
  YES → SONNET
  NO ↓
Architecture, deep reasoning, critical decision?
  YES → OPUS
  NO → SONNET (default workhorse)
```

## Cost Reality

| Model | Input | Output | Relative |
|---|---|---|---|
| Haiku | $0.25/M | $1.25/M | 1x |
| Sonnet | $3.00/M | $15.00/M | 12x |
| Opus | $15.00/M | $75.00/M | 60x |

## Golden Rule

> If a human needs >30 seconds of focused thinking → escalate.

## HAIKU (Default)

Factual Q&A, lookups, status checks, heartbeats, reminders, casual chat, simple file ops, 1-2 sentence tasks.

## SONNET (Workhorse)

Code >10 lines, debugging, analysis, comparisons, planning, reports, long writing, summarization, research synthesis.

## OPUS (Complex Only)

Architecture decisions, multi-file refactoring, complex debugging, strategic planning, deep research, nuanced judgment.
