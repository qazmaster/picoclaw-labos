---
name: linkedin-api
description: Access the LinkedIn API with managed OAuth. Share posts, manage ad campaigns, retrieve profiles and organization info, upload media, and access Ad Library.
metadata: {"clawdbot":{"emoji":"💼","requires":{"env":["MATON_API_KEY"]}}}
---

# LinkedIn API

Share posts, manage ads, retrieve profiles, upload media via Maton gateway.

## Setup

```bash
export MATON_API_KEY="YOUR_API_KEY"
# Get key: maton.ai/settings
# Manage connections: ctrl.maton.ai
```

## Quick Start

```python
# Get current user profile
import urllib.request, os, json
req = urllib.request.Request('https://gateway.maton.ai/linkedin/rest/me')
req.add_header('Authorization', f'Bearer {os.environ["MATON_API_KEY"]}')
req.add_header('LinkedIn-Version', '202506')
print(json.dumps(json.load(urllib.request.urlopen(req)), indent=2))
```

## Key Operations

| Action | Method | Path |
|---|---|---|
| Get profile | GET | `/linkedin/rest/me` |
| Create text post | POST | `/linkedin/rest/posts` |
| Init image upload | POST | `/linkedin/rest/images?action=initializeUpload` |
| List ad accounts | GET | `/linkedin/rest/adAccounts?q=search` |
| Search Ad Library | GET | `/linkedin/rest/adLibrary?q=criteria&keyword={kw}` |
| List orgs | GET | `/linkedin/rest/organizationAcls?q=roleAssignee` |

## Post Example

```json
{
  "author": "urn:li:person:{personId}",
  "lifecycleState": "PUBLISHED",
  "visibility": "PUBLIC",
  "commentary": "Hello LinkedIn!",
  "distribution": {"feedDistribution": "MAIN_FEED"}
}
```

## Notes

- Include `LinkedIn-Version: 202506` header on all requests
- Image/video uploads: 3-step process (init → upload binary → create post)
- Rate limit: 150 req/day per member
