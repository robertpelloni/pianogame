# Song Library & File Selector Protocol (v1.0.0)

## Overview
Connects legacy C++ file dialog and filename sanitization (`legacy/src/file_selector.h`, `file_selector.cpp`) with the modern Agentic Coding Harness. Standardizes directory scanning for MIDI tracks, filename title cleaning, metadata extraction, and recent song history tracking.

## Song Item Schema
```json
{
  "path": "legacy/testing/pitch_bend.mid",
  "filename": "pitch_bend.mid",
  "title": "pitch_bend",
  "sizeBytes": 2707,
  "lastPlayedTimestamp": 1731034200
}
```

## Protocol Operations
* `ScanLibrary(dirPath)`: Scans filesystem directory for `.mid` and `.midi` files, extracting file sizes and cleaned song titles.
* `TrimFilename(filePath)`: Strips path segments and extensions to produce human-readable song display titles.
* `RecordRecentSong(filePath)`: Appends played track to recent playback history with LRU eviction.
* `GetRecentSongs()`: Returns list of recently opened songs.
