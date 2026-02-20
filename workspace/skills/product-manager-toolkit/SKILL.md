---
name: product-manager-toolkit
description: Essential tools and frameworks for product management — RICE prioritization, customer interview analysis, PRD templates, discovery frameworks, and go-to-market strategies.
---

# Product Manager Toolkit

Tools and frameworks for modern product management, from discovery to delivery.

## Quick Start

```bash
# Feature Prioritization (RICE)
python scripts/rice_prioritizer.py sample
python scripts/rice_prioritizer.py sample_features.csv --capacity 15

# Interview Analysis
python scripts/customer_interview_analyzer.py interview_transcript.txt
```

## Core Workflows

### Feature Prioritization (RICE)

1. List candidate features with reach, impact, confidence, effort
2. Run `rice_prioritizer.py` with team capacity
3. Review portfolio analysis and roadmap output
4. Align with stakeholders on prioritized list

### Customer Discovery

1. Conduct interviews (use JTBD framework)
2. Run `customer_interview_analyzer.py` on transcripts
3. Review extracted pain points, feature requests, and patterns
4. Synthesize into opportunity areas

### PRD Development

1. Choose template from `references/prd_templates.md`
2. Fill sections based on discovery work
3. Review with engineering for feasibility
4. Version control in project management tool

## Tools

| Tool | Purpose |
|---|---|
| `rice_prioritizer.py` | RICE scoring with portfolio analysis and roadmap generation |
| `customer_interview_analyzer.py` | NLP extraction of pain points, feature requests, JTBD patterns |
| `references/prd_templates.md` | PRD templates for different product types |
