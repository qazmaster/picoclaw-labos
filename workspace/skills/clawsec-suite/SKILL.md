---
name: clawsec-suite
description: Security suite manager with advisory-feed monitoring, cryptographic signature verification, approval-gated malicious-skill response, and guided setup.
---

# ClawSec Suite

Security suite for monitoring advisories, verifying skill integrity, and managing security protections.

## Features

- Monitor ClawSec advisory feed
- Track new advisories since last check
- Cross-reference against installed skills
- Recommend removal for malicious skills (requires user approval)
- Guarded skill installation with double confirmation
- Cryptographic signature verification

## Install

```bash
npx clawhub@latest install clawsec-suite
```

## Setup Advisory Hook

```bash
SUITE_DIR="${INSTALL_ROOT:-$HOME/.openclaw/skills}/clawsec-suite"
node "$SUITE_DIR/scripts/setup_advisory_hook.mjs"
```

## Guarded Skill Install

```bash
node "$SUITE_DIR/scripts/guarded_skill_install.mjs" --skill helper-plus --version 1.0.1
```

If advisory match found: prints details, exits. Requires `--confirm-advisory` for override.

## Security Notes

- Always verify checksums.json signatures
- Rate-limit advisory polling (5+ min intervals)
- Pin and verify public key fingerprints out-of-band
