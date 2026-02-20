---
name: gcloud
description: Manage Google Cloud Platform resources via gcloud CLI. Compute Engine, Cloud Run, Firebase, Storage, Secret Manager, Cloud SQL, and project management.
---

# Google Cloud Platform

Manage GCP resources using gcloud, gsutil, and firebase CLIs.

## Quick Reference

```bash
# Auth & Config
gcloud auth list
gcloud projects list
gcloud config set project PROJECT_ID

# Compute Engine
gcloud compute instances list
gcloud compute instances start/stop INSTANCE_NAME
gcloud compute ssh INSTANCE_NAME

# Cloud Run
gcloud run services list
gcloud run deploy SERVICE --source .
gcloud run services logs read SERVICE

# Firebase
firebase deploy --only hosting
firebase hosting:channel:deploy CHANNEL

# Cloud Storage
gsutil ls gs://BUCKET/
gsutil cp LOCAL_FILE gs://BUCKET/path/

# Secret Manager
gcloud secrets create SECRET --data-file=-
gcloud secrets versions access latest --secret=SECRET

# Cloud SQL
gcloud sql instances create INSTANCE
gcloud sql connect INSTANCE --user=USERNAME

# Build & Push Container
gcloud builds submit --tag REGION-docker.pkg.dev/PROJECT/REPO/IMAGE:TAG
```

## Troubleshooting

- "API not enabled": `gcloud services enable <service>.googleapis.com`
- "Permission denied": Check roles with `gcloud projects get-iam-policy`
