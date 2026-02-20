---
name: stock-market-pro
description: Local-first market research toolkit. Real-time quotes, fundamentals, ASCII trends, publication-ready PNG charts with RSI/MACD/BB/VWAP/ATR indicators, and one-shot reports.
---

# Stock Market Pro

Local-first market research with charts and reports.

## Prerequisites

- `uv` ([install](https://github.com/astral-sh/uv))
- Optional: `pip3 install -U ddgs` (for news search)

## Commands

```bash
# Quotes
uv run --script scripts/yf.py price TSLA

# Fundamentals (Market Cap, PE, EPS, ROE)
uv run --script scripts/yf.py fundamentals NVDA

# ASCII trend
uv run --script scripts/yf.py history AAPL 6mo

# Pro chart (PNG) with indicators
uv run --script scripts/yf.py pro TSLA 6mo --rsi --macd --bb
uv run --script scripts/yf.py pro TSLA 6mo --vwap --atr

# One-shot report (text + chart PNG)
uv run --script scripts/yf.py report 000660.KS 6mo

# News search
python3 scripts/news.py NVDA --max 8
```

## Indicators

| Flag | Indicator |
|---|---|
| `--rsi` | RSI(14) |
| `--macd` | MACD(12,26,9) |
| `--bb` | Bollinger Bands(20,2) |
| `--vwap` | VWAP (cumulative) |
| `--atr` | ATR(14) |

## Tickers

- US: `AAPL`, `NVDA`, `TSLA`
- KR: `005930.KS`, `000660.KS`
- Crypto: `BTC-USD`, `ETH-KRW`
- FX: `USDKRW=X`
