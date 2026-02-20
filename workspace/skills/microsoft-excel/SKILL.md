---
name: microsoft-excel
description: Access the Microsoft Excel API (via Microsoft Graph) with managed OAuth authentication. Read and write workbooks, worksheets, ranges, tables, and charts.
metadata: {"clawdbot":{"emoji":"📊","requires":{"env":["MATON_API_KEY"]}}}
---

# Microsoft Excel

Read/write Excel workbooks via Microsoft Graph API through Maton gateway.

## Setup

```bash
export MATON_API_KEY="YOUR_API_KEY"
# Get key: maton.ai/settings
# Manage connections: ctrl.maton.ai
```

## Key Operations

```bash
# List worksheets
GET /microsoft-excel/v1.0/me/drive/root:/workbook.xlsx:/workbook/worksheets

# Get range
GET /microsoft-excel/v1.0/me/drive/root:/workbook.xlsx:/workbook/worksheets('Sheet1')/range(address='A1:B2')

# Update range
PATCH /microsoft-excel/v1.0/me/drive/root:/workbook.xlsx:/workbook/worksheets('Sheet1')/range(address='A1:B2')
# Body: {"values": [["A1","B1"],["A2","B2"]]}

# Create table
POST /microsoft-excel/v1.0/me/drive/root:/workbook.xlsx:/workbook/worksheets('Sheet1')/tables/add
# Body: {"address": "A1:D5", "hasHeaders": true}
```

## Notes

- Base URL: `https://gateway.maton.ai/microsoft-excel/{graph-api-path}`
- Auth: `Authorization: Bearer $MATON_API_KEY`
- Works with files in OneDrive or SharePoint
