---
name: market-news-analyst
description: Analyze major market-moving news from the past 10 days. Impact magnitude scoring, market reaction analysis, correlation assessment, and ranked reports covering US equities and commodities.
---

# Market News Analyst

Analyze recent market-moving news events and their impact on US equities and commodities.

## When to Use

- "Analyze the major market news from the past 10 days"
- "How did the latest FOMC decision impact the market?"
- "What were the most important market-moving events this week?"
- "Analyze recent geopolitical news and commodity price reactions"

## Workflow

### 1. News Collection
Use WebSearch/WebFetch to gather news from trusted sources (past 10 days).

### 2. Impact Magnitude Assessment
Score each event by market impact:

| Score | Level | Example |
|---|---|---|
| 9-10 | Extreme | Major policy shift, black swan |
| 7-8 | High | FOMC rate decision, mega-cap earnings miss |
| 5-6 | Moderate | Sector rotation, commodity disruption |
| 3-4 | Low | Minor data release, analyst upgrade |
| 1-2 | Minimal | Routine filing, minor guidance |

### 3. Market Reaction Analysis
- **Immediate** (0-24h): Initial price moves
- **Follow-through** (2-5 days): Trend continuation or reversal
- **Volume confirmation**: Was the move backed by volume?

### 4. Correlation Assessment
- Cross-asset correlations (equities ↔ bonds ↔ commodities)
- Sector rotation patterns
- Distinguish correlation from causation

### 5. Report Generation
Ranked by impact magnitude. Include:
- Event description + date
- Impact score + rationale
- Market reaction (immediate + follow-through)
- Forward implications
