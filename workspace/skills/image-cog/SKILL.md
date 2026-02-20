---
name: image-cog
description: AI image generation powered by CellCog. Create images, edit photos, consistent characters, product photography, reference-based images, sets of images, style transfer. Professional image creation with AI.
---

# Image Cog - AI Image Generation

Create professional images with AI — from single images to consistent character sets to product photography.

## Prerequisites

Requires the `cellcog` skill for SDK setup. Install: `clawhub install cellcog`

## Quick Usage

```python
result = client.create_chat(
    prompt="[your image request]",
    notify_session_key="agent:main:main",
    task_label="image-task",
    chat_mode="agent"  # "agent" for simple, "agent team" for complex
)
```

## Capabilities

- **Single images**: Scenes, portraits, products, abstract art
- **Image editing**: Style transfer, background removal, enhancement
- **Consistent characters**: Same character across multiple scenes
- **Product photography**: Hero shots, lifestyle, flat lays
- **Image sets**: Social media campaigns, website heroes, ad variations
- **Reference-based**: Style matching, composition reference

## Image Specs

| Aspect | Options |
|---|---|
| **Ratios** | 1:1, 16:9, 9:16, 4:3, 3:4, 3:2, 2:3, 21:9 |
| **Sizes** | 1K (~1024px), 2K (~2048px), 4K (~4096px) |
| **Styles** | Photorealistic, illustration, watercolor, oil painting, anime, digital art, vector |

## Tips

1. Be descriptive with subjects, lighting, mood
2. Specify style (photorealistic, illustration, etc.)
3. Include lighting description
4. For character series, describe character in detail first
5. Use `chat_mode="agent team"` for complex multi-image requests
