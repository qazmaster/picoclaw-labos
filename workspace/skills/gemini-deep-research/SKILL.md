---
name: gemini-deep-research
description: Perform complex, long-running research tasks using Gemini Deep Research Agent. Multi-source synthesis, competitive analysis, market research, comprehensive technical investigations.
metadata: {"clawdbot":{"emoji":"🔬","requires":{"env":["GEMINI_API_KEY"]},"primaryEnv":"GEMINI_API_KEY"}}
---

# Gemini Deep Research

Use Gemini's Deep Research Agent for complex, long-running research synthesis.

## Prerequisites

- `GEMINI_API_KEY` environment variable (from Google AI Studio)
- Does NOT work with OAuth tokens — requires direct Gemini API key

## Usage

```bash
# Basic research
scripts/deep_research.py --query "Research the history of Google TPUs"

# Custom output format
scripts/deep_research.py --query "EV battery competitive landscape" \
  --format "1. Executive Summary\n2. Key Players\n3. Supply Chain Risks"

# Stream progress
scripts/deep_research.py --query "Your topic" --stream
```

## How It Works

1. Breaks complex queries into sub-questions
2. Searches the web systematically
3. Synthesizes findings into comprehensive reports
4. Provides streaming progress updates

## Output

Saves to timestamped files:
- `deep-research-YYYY-MM-DD-HH-MM-SS.md` — final report
- `deep-research-YYYY-MM-DD-HH-MM-SS.json` — full metadata

## Notes

- Agent: `deep-research-pro-preview-12-2025`
- Long-running tasks (minutes to hours)
- Get API key: https://aistudio.google.com/apikey
