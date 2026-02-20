---
name: security-audit-toolkit
description: Scan, detect, and fix security issues in codebases and infrastructure. Dependency scanning, secret detection, OWASP patterns, SSL/TLS verification.
---

# Security Audit Toolkit

Scan and fix security issues across code and infrastructure.

## Dependency Scanning

```bash
# Node.js
npm audit
npm audit --json | jq '.vulnerabilities | to_entries[] | {name: .key, severity: .value.severity}'

# Python
pip install pip-audit
pip-audit -r requirements.txt
```

## Secret Detection

```bash
# AWS keys
grep -rn 'AKIA[0-9A-Z]\{16\}' --include='*.{js,ts,py,go,env,yml}' .

# Private keys
grep -rn 'BEGIN.*PRIVATE KEY' .

# Generic high-entropy strings
grep -rn '[A-Za-z0-9+/]\{40,\}' --include='*.{env,yml,yaml,json}' .
```

## SSL/TLS Verification

```bash
openssl s_client -connect example.com:443 -servername example.com < /dev/null 2>/dev/null | \
  openssl x509 -noout -subject -issuer -dates -fingerprint
```

## Secure Coding Checklist

- [ ] No hardcoded secrets (use env vars)
- [ ] Input validation (type, length, format)
- [ ] Parameterized queries (no string concat)
- [ ] Dependencies up to date
- [ ] File permissions correct (600 for keys)
- [ ] HTTPS enforced
