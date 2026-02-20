---
name: edge-tts
description: |
  Text-to-speech conversion using node-edge-tts npm package for generating audio from text.
  Supports multiple voices, languages, speed adjustment, pitch control, and subtitle generation.
  Use when: (1) User requests audio/voice output with the "tts" trigger or keyword. (2) Content needs to be spoken rather than read (multitasking, accessibility, driving, cooking). (3) User wants a specific voice, speed, pitch, or format for TTS output.
---

# Edge-TTS Skill

Generate high-quality text-to-speech audio using Microsoft Edge's neural TTS service via node-edge-tts. No API key needed.

## Quick Start

```bash
cd scripts
npm install
node tts-converter.js "Your text" --voice en-US-AriaNeural --rate +10% --output output.mp3
```

## Options

- `--voice, -v`: Voice name (default: en-US-MichelleNeural)
- `--lang, -l`: Language code (e.g., en-US, es-ES)
- `--pitch`: Pitch adjustment (e.g., +10%, -20%)
- `--rate, -r`: Rate adjustment (e.g., +10%, -20%)
- `--volume`: Volume adjustment (e.g., +0%, -10%)
- `--save-subtitles, -s`: Save subtitles as JSON
- `--output, -f`: Output file path (default: tts_output.mp3)
- `--list-voices, -L`: List available voices

## Common Voices

**English:** `en-US-MichelleNeural` (female), `en-US-AriaNeural` (female), `en-US-GuyNeural` (male), `en-GB-SoniaNeural` (British female)

**Other:** `es-ES-ElviraNeural`, `fr-FR-DeniseNeural`, `de-DE-KatjaNeural`, `ja-JP-NanamiNeural`, `zh-CN-XiaoxiaoNeural`, `ar-SA-ZariyahNeural`

## Rate Guidelines

- `"-20%"` to `"-10%"`: Slow, clear (tutorials, accessibility)
- `"+10%"` to `"+20%"`: Slightly fast (summaries)
- `"+30%"` to `"+50%"`: Fast (news, efficiency)

## Configuration

```bash
node config-manager.js --set-voice en-US-AriaNeural
node config-manager.js --set-rate +10%
node config-manager.js --get
node config-manager.js --reset
```

## Notes

- No API key needed (free Microsoft Edge service)
- Requires internet connection
- Output is MP3 format by default
- Neural voices (ending in `Neural`) provide higher quality
- Test voices at: https://tts.travisvn.com/
