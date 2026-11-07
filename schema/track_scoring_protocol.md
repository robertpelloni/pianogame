# Track Properties & Performance Scoring Protocol (v1.0.0)

## Overview
Standardizes track playback modes, note hit accuracy windows, combo streaks, and performance scoring ported from the legacy C++ Synthesia/Pianogame engine (`TrackProperties.h`, `State_Stats.cpp`).

## Track Modes
* `ModePlayedAutomatically` (0): Track audio is rendered automatically by the synth; notes fall passively.
* `ModeYouPlay` (1): Interactive track evaluated against user keystrokes / MIDI input.
* `ModePlayedButHidden` (2): Audio sounds, but visual note stems are hidden.
* `ModeNotPlayed` (3): Track is muted and excluded from rendering and scoring.

## Scoring Evaluation Schema
```json
{
  "totalNotes": 100,
  "perfectHits": 85,
  "goodHits": 12,
  "missedNotes": 3,
  "maxCombo": 64,
  "accuracyPercent": 94.0,
  "finalScore": 92500,
  "grade": "A"
}
```

## Grading Thresholds
* **S**: Accuracy >= 98.0%
* **A**: Accuracy >= 90.0%
* **B**: Accuracy >= 80.0%
* **C**: Accuracy >= 70.0%
* **F**: Accuracy < 70.0%
