---
name: claude-hook-notify-permission
description: Notification hook to send Slack alerts when Claude needs user input. Based on karanb192/claude-code-hooks.
metadata: {"openclaw":{"requires":{"env":["CCH_SLA_WEBHOOK"]},"primaryEnv":"CCH_SLA_WEBHOOK","homepage":"https://github.com/karanb192/claude-code-hooks"}}
---

# Claude Hook: Notify Permission

Sends Slack alerts when Claude needs user input (such as terminal permission, file write permissions, or general dialogue choices).
Logs to `~/.claude/hooks-logs/`.

## Setup 

1. Ensure you have the `CCH_SLA_WEBHOOK` environment variable exported with your Slack Webhook URL.

```bash
export CCH_SLA_WEBHOOK="https://hooks.slack.com/services/YOUR/WEBHOOK/URL"
```

2. Add the following to your `~/.claude/settings.json` or project-local `.claude/settings.json`. Make sure to replace `PATH_TO_SKILL` with the absolute path to this skill directory.

```json
{
  "hooks": {
    "Notification": [
      {
        "matcher": "permission_prompt|idle_prompt|elicitation_dialog",
        "hooks": [
          {
            "type": "command",
            "command": "node PATH_TO_SKILL/notify-permission.js"
          }
        ]
      }
    ]
  }
}
```
