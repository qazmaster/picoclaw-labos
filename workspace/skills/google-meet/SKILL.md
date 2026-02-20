---
name: google-meet
description: Access the Google Meet API with managed OAuth authentication. Create and manage meeting spaces, list conference records, and retrieve participant information.
metadata: {"clawdbot":{"emoji":"📹","requires":{"env":["MATON_API_KEY"]}}}
---

# Google Meet

Create and manage Google Meet spaces, conference records, participants, recordings, and transcripts via Maton gateway.

## Setup

```bash
export MATON_API_KEY="YOUR_API_KEY"
# Get key: maton.ai/settings
# Manage connections: ctrl.maton.ai
```

## Quick Start

```python
# Create a meeting space
import urllib.request, os, json
data = json.dumps({}).encode()
req = urllib.request.Request('https://gateway.maton.ai/google-meet/v2/spaces', data=data, method='POST')
req.add_header('Authorization', f'Bearer {os.environ["MATON_API_KEY"]}')
req.add_header('Content-Type', 'application/json')
print(json.dumps(json.load(urllib.request.urlopen(req)), indent=2))
```

## API Endpoints

| Action | Method | Path |
|---|---|---|
| Create space | POST | `/google-meet/v2/spaces` |
| Get space | GET | `/google-meet/v2/spaces/{spaceId}` |
| End call | POST | `/google-meet/v2/spaces/{spaceId}:endActiveConference` |
| List conferences | GET | `/google-meet/v2/conferenceRecords` |
| List participants | GET | `/google-meet/v2/conferenceRecords/{id}/participants` |
| List recordings | GET | `/google-meet/v2/conferenceRecords/{id}/recordings` |
| List transcripts | GET | `/google-meet/v2/conferenceRecords/{id}/transcripts` |

## Notes

- All requests need `Authorization: Bearer $MATON_API_KEY`
- Manage OAuth connections at `ctrl.maton.ai`
- Rate limit: 10 req/sec per account
