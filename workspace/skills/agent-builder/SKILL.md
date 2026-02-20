---
name: agent-builder
description: Build high-performing AI agents end-to-end. Design persona + operating rules and generate workspace files (IDENTITY.md, SOUL.md, AGENTS.md, USER.md, HEARTBEAT.md).
---

# Agent Builder

Design and generate complete AI agent workspaces with strong defaults.

## Workflow

### Phase 1 — Interview

Ask clarifying questions (keep tight, multiple short rounds):
- **Job statement**: Primary mission in one sentence?
- **Surfaces**: Which channels (Telegram/WhatsApp/Discord)?
- **Autonomy level**: Advisor, Operator, or Autopilot?
- **Hard prohibitions**: Actions the agent must never take?
- **Memory**: Should it keep curated MEMORY.md?
- **Tone**: Concise vs narrative; strict vs warm?
- **Tool posture**: Tool-first vs answer-first?

### Phase 2 — Generate Files

- `IDENTITY.md` — who the agent is
- `SOUL.md` — personality and values
- `AGENTS.md` — agent configuration
- `USER.md` — user preferences
- `HEARTBEAT.md` — periodic check rules
- Optional: `MEMORY.md`, `memory/YYYY-MM-DD.md`

### Phase 3 — Guardrails

- Ask-before-destructive rules
- Stop-on-CLI-usage-error
- Loop breaker guidance

### Phase 4 — Acceptance Tests

5–10 short scenario prompts to validate behavior.
