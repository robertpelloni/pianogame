# User Preferences & Audio Settings Protocol (v1.0.0)

## Overview
Replaces legacy platform-dependent registry (`WIN32 RegOpenKeyEx`) and CFPreferences (`Mac OS X`) with a modern, cross-platform JSON configuration protocol for synthesizer latency, audio buffers, input/output MIDI devices, and visual themes.

## Settings Schema
```json
{
  "appName": "PianoGame",
  "version": "1.0.0",
  "audioBufferSize": 256,
  "audioSampleRate": 48000,
  "defaultMidiIn": "Generic USB MIDI",
  "defaultMidiOut": "Internal Software Synthesizer",
  "preferredKeyboardSize": 88,
  "notationStyle": "falling_notes",
  "showKeyLabels": true,
  "autoPlayUnselectedTracks": true
}
```

## Protocol Operations
* `GetSettings()`: Retrieves current active preferences with fallback defaults.
* `UpdateSettings(partial)`: Validates values and merges updates into the configuration state.
* `ResetToDefaults()`: Reverts user settings to factory defaults.
