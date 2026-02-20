---
name: automation-workflows
description: A playbook for solopreneurs to identify, design, and implement no-code automation workflows using tools like Zapier, Make, and n8n.
---

# Automation Workflows

Identify, design, and implement no-code automations. Automate anything done more than twice a week that doesn't require creative thinking.

## Step 1: Identify What to Automate

**Automation audit**: Track tasks for a week, note time/frequency/repetitiveness.

`Time Cost = (Minutes per task × Frequency per month) / 60`

**Good candidates**: Repetitive, rule-based, high-frequency, time-consuming (10+ min).

## Step 2: Choose Tool

| Tool | Best For | Price |
|---|---|---|
| Zapier | Simple 2-app connections | Free (100 tasks/mo) |
| Make | Visual multi-step workflows | Free (1000 ops/mo) |
| n8n | Self-hosted, complex logic | Free (self-hosted) |

## Step 3: Design Pattern

```
Trigger → Filter → Action → Notification
```

Common patterns: Data sync, lead capture, content distribution, invoice automation.

## Step 4: Build & Test

1. Start with the trigger
2. Add one action at a time
3. Test with real data
4. Add error handling
5. Monitor for 1 week before trusting fully
