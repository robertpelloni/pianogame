# Keyboard Geometry & Layout Protocol (v1.0.0)

## Overview
Defines standard virtual keyboard geometries, starting octave offsets, white/black key counts, and MIDI note mappings ported from the legacy C++ Synthesia/Pianogame engine (`KeyboardDisplay.h`, `KeyboardDisplay.cpp`).

## Supported Keyboard Configurations

| Size | Total Keys | White Keys | Black Keys | Starting Note | Starting Octave |
|------|------------|------------|------------|---------------|-----------------|
| 37   | 37         | 22         | 15         | F2 (41)       | 2               |
| 49   | 49         | 29         | 20         | C2 (36)       | 2               |
| 61   | 61         | 36         | 25         | C2 (36)       | 2               |
| 76   | 76         | 45         | 31         | E1 (28)       | 1               |
| 88   | 88         | 52         | 36         | A0 (21)       | 0               |

## Layout Schema
```json
{
  "size": 88,
  "totalKeys": 88,
  "whiteKeyCount": 52,
  "blackKeyCount": 36,
  "startMidiNote": 21,
  "endMidiNote": 108,
  "startOctave": 0,
  "startNoteName": "A0",
  "endNoteName": "C8"
}
```
