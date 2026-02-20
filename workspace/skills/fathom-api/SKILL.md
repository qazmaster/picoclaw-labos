---
name: fathom-api
description: Fathom API integration with managed OAuth. Access meeting recordings, transcripts, summaries, action items, and manage webhooks.
metadata: {"clawdbot":{"emoji":"🎥","requires":{"env":["MATON_API_KEY"]}}}
---

# Fathom API

Access meeting recordings, transcripts, summaries, and action items via Maton gateway.

## Setup

```bash
export MATON_API_KEY="YOUR_API_KEY"
# Get key: maton.ai/settings
# Manage connections: ctrl.maton.ai
```

## Quick Start

```python
# List recent meetings
import urllib.request, os, json
req = urllib.request.Request('https://gateway.maton.ai/fathom/external/v1/meetings')
req.add_header('Authorization', f'Bearer {os.environ["MATON_API_KEY"]}')
print(json.dumps(json.load(urllib.request.urlopen(req)), indent=2))
```

## API Endpoints

| Action | Method | Path |
|---|---|---|
| List meetings | GET | `/fathom/external/v1/meetings` |
| Get summary | GET | `/fathom/external/v1/recordings/{id}/summary` |
| Get transcript | GET | `/fathom/external/v1/recordings/{id}/transcript` |
| List teams | GET | `/fathom/external/v1/teams` |
| List members | GET | `/fathom/external/v1/team_members` |
| Create webhook | POST | `/fathom/external/v1/webhooks` |
| Delete webhook | DELETE | `/fathom/external/v1/webhooks/{id}` |

## Filters

- `created_after` / `created_before` — ISO 8601 timestamps
- `teams[]` — filter by team names
- `recorded_by[]` — filter by recorder email
- `cursor` — pagination

## Notes

- Auth: `Authorization: Bearer $MATON_API_KEY`
- Manage OAuth at `ctrl.maton.ai`
- Webhooks: set `triggered_for` to `my_recordings`, `shared_team_recordings`, etc.
- Use `destination_url` param for async transcript/summary delivery
