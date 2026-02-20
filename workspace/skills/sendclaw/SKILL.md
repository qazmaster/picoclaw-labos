---
name: sendclaw
description: Email for AI Agents — get an email address your_handle@sendclaw.com. Register, send, receive, and check emails via REST API.
---

# SendClaw — Email for AI Agents

Get a dedicated email address `your-handle@sendclaw.com` for the agent.

## Setup

```bash
# Register (once) — ask user for preferred handle
curl -X POST https://sendclaw.com/api/bots/register \
  -H "Content-Type: application/json" \
  -d '{"name":"MyBot","handle":"preferred-handle","senderName":"Bot Name"}'
# Returns API key
```

## API Reference

Base URL: `https://sendclaw.com/api`
Auth: `X-Api-Key: your-api-key` (or `Authorization: Bearer your-api-key`)

| Action | Method | Endpoint |
|---|---|---|
| Register | POST | `/api/bots/register` |
| Send email | POST | `/api/mail/send` |
| Check new | GET | `/api/mail/check` |
| Get unread | GET | `/api/mail/messages?unread=true` |
| Get all | GET | `/api/mail/messages` |

## Send Email

```bash
curl -X POST https://sendclaw.com/api/mail/send \
  -H "X-Api-Key: YOUR_KEY" \
  -H "Content-Type: application/json" \
  -d '{"to":"recipient@example.com","subject":"Hello","body":"Message body","cc":"optional@cc.com"}'
```

## Notes

- Ask user for preferred email handle before registering
- `GET /api/mail/messages?unread=true` auto-marks messages as read
- `GET /api/mail/check` returns `{unreadCount, quota}`
