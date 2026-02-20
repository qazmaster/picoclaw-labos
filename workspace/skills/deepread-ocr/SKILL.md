---
name: deepread-ocr
description: AI-native OCR platform that turns documents into high-accuracy structured data. Multi-model consensus achieves 97%+ accuracy with Human-in-the-Loop review for uncertain fields. Use for invoice processing, receipt OCR, contract analysis, form digitization.
metadata: {"openclaw":{"requires":{"env":["DEEPREAD_API_KEY"]},"primaryEnv":"DEEPREAD_API_KEY","homepage":"https://www.deepread.tech"}}
---

# DeepRead OCR

Production-grade OCR API with multi-model consensus (97%+ accuracy) and Human-in-the-Loop review.

## Setup

```bash
export DEEPREAD_API_KEY="sk_live_your_key_here"
# Get key: https://www.deepread.tech/dashboard
# Free tier: 2,000 pages/month
```

## Basic OCR

```bash
curl -X POST https://api.deepread.tech/v1/process \
  -H "X-API-Key: $DEEPREAD_API_KEY" \
  -F "file=@document.pdf" \
  -F "webhook_url=https://your-app.com/webhook"
```

## Structured Data Extraction

```bash
curl -X POST https://api.deepread.tech/v1/process \
  -H "X-API-Key: $DEEPREAD_API_KEY" \
  -F "file=@invoice.pdf" \
  -F 'schema={"type":"object","properties":{"vendor":{"type":"string"},"total":{"type":"number"},"invoice_date":{"type":"string"}}}'
```

Response includes `hil_flag` per field — `false` = confident, `true` = needs review.

## Check Results

```bash
curl https://api.deepread.tech/v1/jobs/JOB_ID \
  -H "X-API-Key: $DEEPREAD_API_KEY"
```

## Notes

- Processing: 2-5 min (async, use webhooks)
- Formats: PDF, JPG, JPEG, PNG (max 50MB)
- Free tier: 2,000 pages/month, 10 req/min
- HIL review interface: `preview.deepread.tech`
