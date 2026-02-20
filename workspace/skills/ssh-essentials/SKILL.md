---
name: ssh-essentials
description: SSH for remote access and secure file transfers. Key management, port forwarding, tunneling, config, SCP, SFTP, rsync.
---

# SSH Essentials

## Connection

```bash
ssh user@hostname                     # Basic
ssh user@hostname -p 2222             # Custom port
ssh -i ~/.ssh/id_rsa user@hostname    # Specific key
ssh user@hostname 'ls -la'            # Run command
ssh -A user@hostname                  # Forward agent
```

## Key Management

```bash
ssh-keygen -t ed25519 -C "email@example.com"   # Generate
ssh-copy-id user@hostname                       # Copy to server
ssh-add ~/.ssh/id_rsa                           # Add to agent
```

## Port Forwarding

```bash
ssh -L 8080:localhost:80 user@hostname     # Local forward
ssh -R 8080:localhost:3000 user@hostname   # Remote forward
ssh -D 1080 user@hostname                 # SOCKS proxy
ssh -f -N -L 8080:localhost:80 user@host  # Background tunnel
```

## Config (~/.ssh/config)

```
Host production
    HostName prod.example.com
    User deploy
    IdentityFile ~/.ssh/id_prod

Host internal
    HostName 10.0.0.5
    ProxyJump bastion
```

## File Transfers

```bash
scp file.txt user@host:/path/                       # SCP
rsync -avz --progress /local/ user@host:/remote/     # Rsync
sftp user@hostname                                   # SFTP
```

## Troubleshooting

- `ssh -vvv user@hostname` — max verbosity
- `chmod 700 ~/.ssh && chmod 600 ~/.ssh/id_rsa` — fix permissions
- `ssh-keygen -R hostname` — clear host key
