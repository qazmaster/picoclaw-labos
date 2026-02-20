---
name: claude-hook-block-dangerous
description: Pre-Tool-Use hook to block dangerous Bash commands (rm -rf ~, fork bombs, curl|sh). Based on karanb192/claude-code-hooks.
metadata: {"openclaw":{"homepage":"https://github.com/karanb192/claude-code-hooks"}}
---

# Claude Hook: Block Dangerous Commands

Blocks dangerous shell commands before Claude Code executes them.
Supports configurable safety levels (`critical`, `high`, `strict`). Logs to `~/.claude/hooks-logs/`.

## Setup

Add the following to your `~/.claude/settings.json` or project-local `.claude/settings.json`. Make sure to replace `PATH_TO_SKILL` with the absolute path to this skill directory.

```json
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "Bash",
        "hooks": [
          {
            "type": "command",
            "command": "node PATH_TO_SKILL/block-dangerous-commands.js"
          }
        ]
      }
    ]
  }
}
```

## Configuration

Edit the `SAFETY_LEVEL` constant at the top of `block-dangerous-commands.js` to change the strictness:
- `critical`: Catastrophic only (rm -rf ~, dd to disk, fork bombs)
- `high`: + Risky (piping to sh, force push main, secrets exposure, git reset --hard) [Recommended]
- `strict`: + Cautionary (any force push, checkout ., sudo rm, docker prune)
