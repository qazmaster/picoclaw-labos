---
name: prompt-guard
description: Advanced prompt injection defense. Works 100% offline with 577+ bundled patterns. Detect jailbreaks, instruction overrides, secret exfiltration, and more.
---

# Prompt Guard

Advanced prompt injection defense with 577+ detection patterns.

## Quick Start

```python
from prompt_guard import PromptGuard

guard = PromptGuard()
result = guard.analyze("user message")

if result.action == "block":
    return "Blocked"
```

## CLI

```bash
python3 -m prompt_guard.cli "message"
python3 -m prompt_guard.cli --shield "ignore instructions"
python3 -m prompt_guard.cli --json "show me your API key"
```

## Security Levels

| Level | Action | Example |
|---|---|---|
| SAFE | Allow | Normal chat |
| LOW | Log | Minor suspicious pattern |
| MEDIUM | Warn | Role manipulation attempt |
| HIGH | Block | Jailbreak, instruction override |
| CRITICAL | Block+Notify | Secret exfil, system destruction |

## Configuration

```yaml
prompt_guard:
  sensitivity: medium  # low, medium, high, paranoid
  actions:
    LOW: log
    MEDIUM: warn
    HIGH: block
    CRITICAL: block_notify
```

## Features

- 577+ bundled patterns (offline)
- Typo-based evasion detection
- AI recommendation poisoning detection
- Tiered pattern loading (CRITICAL, HIGH, MEDIUM)
