---
name: himalaya
description: CLI email client using IMAP/SMTP backends. List, read, search, send, reply, forward emails from the terminal.
homepage: https://github.com/pimalaya/himalaya
metadata: {"clawdbot":{"emoji":"📧","requires":{"bins":["himalaya"]},"install":[{"id":"brew","kind":"brew","formula":"himalaya","bins":["himalaya"],"label":"Install himalaya (brew)"}]}}
---

# Himalaya Email CLI

Manage emails from the terminal using IMAP, SMTP, Notmuch, or Sendmail backends.

## Setup

Run the interactive wizard:
```bash
himalaya account configure
```

Or create `~/.config/himalaya/config.toml` manually with IMAP/SMTP credentials.

## Common Operations

```bash
# List folders
himalaya folder list

# List emails (default: INBOX)
himalaya envelope list
himalaya envelope list --folder "Sent"
himalaya envelope list --page 1 --page-size 20

# Search
himalaya envelope list from john@example.com subject meeting

# Read email by ID
himalaya message read 42

# Reply / Reply-all
himalaya message reply 42
himalaya message reply 42 --all

# Forward
himalaya message forward 42

# Compose new email (opens $EDITOR)
himalaya message write

# Send directly via template
cat <<'EOF' | himalaya template send
From: you@example.com
To: recipient@example.com
Subject: Test Message

Hello from Himalaya!
EOF

# Move/Copy/Delete
himalaya message move 42 "Archive"
himalaya message copy 42 "Important"
himalaya message delete 42

# Manage flags
himalaya flag add 42 --flag seen
himalaya flag remove 42 --flag seen

# Attachments
himalaya attachment download 42 --dir ~/Downloads

# Multiple accounts
himalaya account list
himalaya --account work envelope list

# JSON output
himalaya envelope list --output json
```

## Notes

- Message IDs are relative to the current folder; re-list after folder changes.
- Store passwords securely using `pass`, system keyring, or a command that outputs the password.
- Debug: `RUST_LOG=debug himalaya envelope list`
