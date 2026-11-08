# CHANGELOG

See `VERSION.md` for the current active version string.

## [v5.96.0]

### Added
* Defined Cross-Engine Song Library & File Selector Protocol in `schema/song_library_protocol.md`.
* Implemented modern Song Library Engine with directory scanning, title sanitization, and recent song tracking in TypeScript (`typescript/src/song_library.ts`) and Go (`go/llmapi/song_library.go`) with tests.
* Integrated Song Library & Files card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) wired to `/api/songs/list` and `/api/songs/recent`.

### Changed
* Updated `VERSION.md` to `v5.96.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.95.0]

### Added
* Defined Cross-Engine MIDI Hardware Device Enumeration Protocol in `schema/midi_device_protocol.md`.
* Implemented modern MIDI Device Manager with input/output device discovery, latency profiling, and port selection in TypeScript (`typescript/src/midi_device.ts`) and Go (`go/llmapi/midi_device.go`) with tests.
* Integrated MIDI Hardware & Devices Studio card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) wired to `/api/midi/devices`.

### Changed
* Updated `VERSION.md` to `v5.95.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.94.0]

### Added
* Defined Cross-Engine User Preferences & Audio Settings Protocol in `schema/user_settings_protocol.md`.
* Implemented modern User Settings Engine managing audio buffers, sample rates, keyboard size preferences, and visual styles in TypeScript (`typescript/src/user_settings.ts`) and Go (`go/llmapi/user_settings.go`) with tests.
* Integrated Audio & Preferences Studio card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) wired to `/api/settings/get` and `/api/settings/update`.

### Changed
* Updated `VERSION.md` to `v5.94.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.93.0]

### Added
* Defined Cross-Engine Keyboard Geometry & Layout Protocol in `schema/keyboard_geometry_protocol.md`.
* Implemented modern Keyboard Geometry Engine supporting 37, 49, 61, 76, and 88-key piano layouts in TypeScript (`typescript/src/keyboard_geometry.ts`) and Go (`go/llmapi/keyboard_geometry.go`) with tests.
* Integrated Virtual Keyboard Geometry Studio card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) wired to `/api/keyboard/layout`.

### Changed
* Updated `VERSION.md` to `v5.93.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.92.0]

### Added
* Defined Cross-Engine Track Properties & Performance Scoring Protocol in `schema/track_scoring_protocol.md`.
* Implemented modern Track Scoring Engine with accuracy weighting, combo bonus, and grading in TypeScript (`typescript/src/track_scoring.ts`) and Go (`go/llmapi/track_scoring.go`) with tests.
* Integrated Performance & Scoring Studio card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) wired to `/api/scoring/evaluate`.

### Changed
* Updated `VERSION.md` to `v5.92.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.91.0]

### Added
* Defined Cross-Engine MIDI Event & Pitch Analysis Protocol in `schema/midi_analysis_protocol.md`.
* Implemented modern binary MIDI analyzer in TypeScript (`typescript/src/midi_analyzer.ts`) and Go (`go/llmapi/midi_analyzer.go`) with tests verifying against `legacy/testing/pitch_bend.mid`.
* Integrated MIDI & Audio Event Studio card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) wired to `/api/midi/analyze`.

### Changed
* Updated `VERSION.md` to `v5.91.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.90.0]

### Added
* Expanded Umbrella v6 architecture roadmap (`ROADMAP.md`) covering the consolidated single-page dashboard UI, system tray process supervision, and MCP tool discovery.
* Comprehensive deployment guide updates (`DEPLOY.md`) detailing unified dashboard launching, background tray supervision, and dynamic manifest synchronization.

### Changed
* Updated `VERSION.md` to `v5.90.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`VISION.md`, `ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.89.0]

### Added
* Defined Cross-Engine Prompt Template Engine Protocol in `schema/prompt_template_protocol.md`.
* Implemented Prompt Template Engine with parameter extraction, validation, and multi-turn message formatting in TypeScript (`typescript/src/prompt_template.ts`) and Go (`go/llmapi/prompt_template.go`).
* Added Prompt Template Studio interactive card to the Unified Mission Control Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) and wired `/api/prompt/render`.

### Changed
* Updated `VERSION.md` to `v5.89.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.88.0]

### Added
* Implemented Self-Balancing Adaptive LLM Router in TypeScript (`typescript/src/adaptive_router.ts`) and Go (`go/llmapi/adaptive_router.go`) evaluating provider latency, token costs, and error rates to balance requests across models.
* Integrated Adaptive LLM Router card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) with interactive strategy selection (`balanced`, `fastest`, `cheapest`).

### Changed
* Updated `VERSION.md` to `v5.88.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.87.0]

### Added
* Defined Desktop Webview Container & Native IPC Protocol in `schema/desktop_container.md`.
* Implemented Desktop Container configuration and IPC launcher in Rust (`rust/src/desktop_container.rs`) and Go (`go/supervisor/desktop_container.go`).
* Integrated Desktop Webview Container card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.87.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.86.0]

### Added
* Defined SQLite Vector & Persistent Storage Protocol in `schema/sqlite_store_protocol.md`.
* Implemented cross-language SQLite/Storage Provider in TypeScript (`typescript/src/sqlite_store.ts`) and Go (`go/llmapi/sqlite_store.go`) providing structured tables for agent key-value pairs, vector embeddings, and telemetry logs.
* Integrated SQLite Storage Provider card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.86.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).
