# CHANGELOG

See `VERSION.md` for the current active version string.

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
