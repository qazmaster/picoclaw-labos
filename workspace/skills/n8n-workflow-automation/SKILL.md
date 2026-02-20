---
name: n8n-workflow-automation
description: Design and output n8n workflow JSON with robust triggers, idempotency, error handling, logging, retries, and human-in-the-loop review queues.
---

# n8n Workflow Automation

Design and output n8n workflow JSON with robust triggers, idempotency, error handling, logging, retries, and review queues.

## When to Use

- Build n8n workflows with schedules and email outputs
- Add error handling and retries with review queues for failures
- Create webhook workflows with logging and status tracking
- Make flows idempotent to avoid duplicate records on reruns
- Instrument workflows with audit logs and human approval steps

## Required Inputs

- **Workflow intent**: trigger type + schedule/timezone + success criteria
- **Targets**: where to write results (email/Drive/Sheet/DB) + required fields

## Optional Inputs

- Existing n8n workflow JSON to modify
- Sample payloads / example records
- Dedup key definitions

## Outputs

- **Default**: Workflow design spec (nodes, data contracts, failure modes)
- **If requested**: `workflow.json` (n8n importable) + `runbook.md`

## Success Criteria

Workflow is idempotent, logs every run, retries safely, and routes failures to a review queue.
