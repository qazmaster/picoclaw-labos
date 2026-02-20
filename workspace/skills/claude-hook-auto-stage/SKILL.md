---
name: claude-hook-auto-stage
description: Post-Tool-Use hook to automatically git stage files after Claude Code modifies them. Based on karanb192/claude-code-hooks.
metadata: {"openclaw":{"homepage":"https://github.com/karanb192/claude-code-hooks"}}
---

# Claude Hook: Auto Stage

Automatically stages files via `git add` after Claude Code uses the `Edit` or `Write` tools to modify them.
Logs to `~/.claude/hooks-logs/`.

## Benefits
- `git status` shows exactly what Claude modified
- Easy to review changes before committing
- No manual staging needed

*Note: This relies on `.gitignore` to exclude sensitive files.*

## Setup

Add the following to your `~/.claude/settings.json` or project-local `.claude/settings.json`. Make sure to replace `PATH_TO_SKILL` with the absolute path to this skill directory.

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Edit|Write",
        "hooks": [
          {
            "type": "command",
            "command": "node PATH_TO_SKILL/auto-stage.js"
          }
        ]
      }
    ]
  }
}
```
