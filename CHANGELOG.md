# CHANGELOG

See `VERSION.md` for the current active version string.

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

## [v5.85.0]

### Added
* Defined External Email & Notification Dispatch Protocol in `schema/email_protocol.md`.
* Implemented native Email & Notification Dispatcher in TypeScript (`typescript/src/email_dispatcher.ts`) and Go (`go/llmapi/email_dispatcher.go`) supporting inbound task message parsing and outbound system notifications.
* Integrated Email & Notification Dispatch card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.85.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).
