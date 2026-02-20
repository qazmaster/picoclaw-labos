---
name: claude-hook-protect-secrets
description: Pre-Tool-Use hook to prevent reading, modifying, or exfiltrating sensitive files. Based on karanb192/claude-code-hooks.
metadata: {"openclaw":{"homepage":"https://github.com/karanb192/claude-code-hooks"}}
---

# Claude Hook: Protect Secrets

Prevents reading, modifying, or exfiltrating sensitive files via Bash, Read, Edit, or Write tools. 
Supports configurable safety levels (`critical`, `high`, `strict`). Logs to `~/.claude/hooks-logs/`.

## Setup

Add the following to your `~/.claude/settings.json` or project-local `.claude/settings.json`. Make sure to replace `PATH_TO_SKILL` with the absolute path to this skill directory.

```json
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "Read|Edit|Write|Bash",
        "hooks": [
          {
            "type": "command",
            "command": "node PATH_TO_SKILL/protect-secrets.js"
          }
        ]
      }
    ]
  }
}
```

## Configuration

Edit the `SAFETY_LEVEL` constant at the top of `protect-secrets.js` to change the strictness:
- `critical`: Blocks SSH keys, AWS credentials, `.env` files.
- `high`: + Blocks generalized secrets files, env dumps (`printenv`), and exfiltration attempts (e.g. `curl -d @.env`). [Recommended]
- `strict`: + Blocks database configs, `known_hosts`, `.gitconfig`, or anything that might contain secrets.
