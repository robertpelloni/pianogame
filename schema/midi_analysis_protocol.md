# MIDI Event & Pitch Analysis Protocol (v1.0.0)

## Overview
Connects the legacy `libmidi` C++ parsing capabilities with the modern Agentic Coding Harness. Provides standardized serialization and diagnostic inspection of Standard MIDI Files (SMF 0/1), event delta timings, note-on/note-off distributions, and pitch bend sensitivity.

## Diagnostic Metrics Schema
```json
{
  "format": 1,
  "trackCount": 2,
  "division": 480,
  "totalEvents": 412,
  "noteCount": 184,
  "pitchBendEvents": 28,
  "maxPolyphony": 6,
  "pitchBendRangeSemis": 2.0,
  "channelsDetected": [0, 1]
}
```

## Agentic Operations
* **AnalyzeFile(path)**: Reads standard binary MIDI data (`.mid`), decodes `MThd` and `MTrk` chunks, validates header timing divisions, and tallies event types.
* **InspectPitchBendSensitivity()**: Analyzes RPN (Registered Parameter Number) pitch bend range sensitivity and raw 14-bit pitch wheel alterations.
* **ExtractNoteRanges()**: Computes min/max MIDI key values and polyphony density for autonomous audio synthesis.
