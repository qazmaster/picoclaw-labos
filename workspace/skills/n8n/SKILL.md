---
name: n8n
description: Manage n8n workflows and automations via API. List, activate/deactivate, check execution status, trigger, and debug workflows.
metadata: {"openclaw":{"emoji":"⚙️","requires":{"env":["N8N_API_KEY","N8N_BASE_URL"]},"primaryEnv":"N8N_API_KEY"}}
---

# n8n Workflow Management

Manage n8n workflows via API: create, test, execute, monitor, and optimize.

## Setup

```bash
export N8N_API_KEY="your-api-key"
export N8N_BASE_URL="http://localhost:5678"
```

## Critical Rules

When creating workflows, ALWAYS:
- Generate COMPLETE workflows with all functional nodes
- Include actual HTTP Request nodes for API calls
- Add Code nodes for data transformation
- Create proper connections between all nodes
- Use real node types (`n8n-nodes-base.httpRequest`, `n8n-nodes-base.code`, `n8n-nodes-base.set`)

NEVER create placeholder/TODO nodes.

## Quick Reference

```bash
# List workflows
curl -s "$N8N_BASE_URL/api/v1/workflows" -H "X-N8N-API-KEY: $N8N_API_KEY"

# Get workflow
curl -s "$N8N_BASE_URL/api/v1/workflows/{id}" -H "X-N8N-API-KEY: $N8N_API_KEY"

# Activate
curl -s -X PATCH "$N8N_BASE_URL/api/v1/workflows/{id}" \
  -H "X-N8N-API-KEY: $N8N_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"active": true}'

# Execute
curl -s -X POST "$N8N_BASE_URL/api/v1/workflows/{id}/run" \
  -H "X-N8N-API-KEY: $N8N_API_KEY"

# List executions
curl -s "$N8N_BASE_URL/api/v1/executions?workflowId={id}" \
  -H "X-N8N-API-KEY: $N8N_API_KEY"
```
