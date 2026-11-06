# CHANGELOG

See `VERSION.md` for the current active version string.

## [v5.80.0]

### Added
* Implemented Persistent Key-Value State Store and Memory Manager in TypeScript (`typescript/src/state_store.ts`) and Go (`go/llmapi/state_store.go`).
* Added Agent Memory Store card with interactive key-value storage and live refresh controls into the Unified WebUI Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.80.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.79.0]

### Added
* Implemented cross-platform Desktop Native Tray Launcher in Go (`go/supervisor/desktop_tray.go`).
* Added `scripts/start_tray.sh` launcher script for background supervisor and notification dispatch.
* Integrated desktop browser opening (`xdg-open`, `open`, `start`) and notification capabilities into the supervisor.

### Changed
* Updated `VERSION.md` to `v5.79.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.78.0]

### Added
* Implemented cross-language native WebSocket client bridges in TypeScript (`typescript/src/ws_bridge.ts`) and Go (`go/llmapi/ws_bridge.go`) for remote agent connection and control.
* Integrated active WebSocket client tracking (`wsClientBadge`) into the Unified WebUI Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.78.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).
