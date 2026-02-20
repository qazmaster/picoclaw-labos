---
name: playwright-mcp
description: Browser automation via Playwright MCP server. Navigate websites, click elements, fill forms, extract data, take screenshots.
metadata: {"openclaw":{"emoji":"🎭","os":["linux","darwin","win32"],"requires":{"bins":["npx"]},"install":[{"id":"npm-playwright-mcp","kind":"npm","package":"@playwright/mcp","bins":["playwright-mcp"],"label":"Install Playwright MCP"}]}}
---

# Playwright MCP

Browser automation powered by Playwright MCP server.

## Install

```bash
npm install -g @playwright/mcp
npx playwright install chromium
```

## Start Server

```bash
npx @playwright/mcp                    # Default
npx @playwright/mcp --headless         # Headless
npx @playwright/mcp --browser firefox  # Firefox
npx @playwright/mcp --viewport-size 1280x720
```

## MCP Tools

| Tool | Description |
|---|---|
| `browser_navigate` | Open URL |
| `browser_click` | Click element |
| `browser_type` | Type into input |
| `browser_select_option` | Select dropdown |
| `browser_get_text` | Extract text |
| `browser_evaluate` | Run JavaScript |
| `browser_snapshot` | Get page structure |
| `browser_press` | Press key |
| `browser_choose_file` | Upload file |
| `browser_close` | Close context |

## Example: Login Flow

```
browser_navigate: { url: "https://example.com/login" }
browser_type: { selector: "#username", text: "user" }
browser_type: { selector: "#password", text: "pass" }
browser_click: { selector: "#submit" }
browser_get_text: { selector: ".welcome-message" }
```

## Security

- Host validation prevents navigation to untrusted domains
- Use `--allowed-hosts` to whitelist domains
- Sandboxing enabled by default
