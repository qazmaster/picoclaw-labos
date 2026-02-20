---
name: hetzner-cloud
description: Hetzner Cloud CLI for infrastructure management. Servers, firewalls, networks, volumes, snapshots, and SSH keys.
---

# Hetzner Cloud CLI

Manage Hetzner Cloud infrastructure.

## ⚠️ Safety Rules

- **NEVER** execute delete commands
- **ALWAYS** ask confirmation before create/modify
- **ALWAYS** snapshot before modifications:
  `hcloud server create-image <server> --type snapshot --description "Backup"`

## Setup

1. Get API token: https://console.hetzner.cloud/
2. `hcloud context create <name>`
3. `hcloud context use <name>`

## Commands

```bash
# Servers
hcloud server list
hcloud server describe <name>
hcloud server create --name my-server --type cx22 --image ubuntu-24.04 --location fsn1
hcloud server poweron/poweroff/reboot <name>
hcloud server ssh <name>

# Firewalls
hcloud firewall create --name my-fw
hcloud firewall add-rule <fw> --direction in --port 80 --protocol tcp --source-ips 0.0.0.0/0

# Networks
hcloud network create --name my-net --ip-range 10.0.0.0/8
hcloud server attach-to-network <server> --network my-net

# Volumes
hcloud volume create --name my-vol --size 10 --server <server>

# Snapshots
hcloud server create-image <server> --type snapshot
hcloud image list --type snapshot

# Output formats
hcloud server list -o json
hcloud server list -o columns=id,name,status
```
