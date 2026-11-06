# CHANGELOG

See `VERSION.md` for the current active version string.

## [v5.82.0]

### Added
* Defined Cognitive RAG Vector Indexing Protocol in `schema/rag_protocol.md`.
* Implemented in-memory RAG Vector Indexer and cosine similarity semantic search in TypeScript (`typescript/src/rag_indexer.ts`) and Go (`go/llmapi/rag_indexer.go`).
* Integrated RAG Vector Search card with interactive semantic queries into the Unified WebUI Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.82.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.81.0]

### Added
* Implemented Cross-Engine Stability & Performance Health Monitor in TypeScript (`typescript/src/monitor.ts`) and Go (`go/supervisor/monitor.go`) providing latency tracking, CPU/memory telemetry, uptime metrics, and performance health benchmarking.
* Integrated Stability & Performance Monitor card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`) with live status indicators and interactive benchmark controls.

### Changed
* Updated `VERSION.md` to `v5.81.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.80.0]

### Added
* Implemented Persistent Key-Value State Store and Memory Manager in TypeScript (`typescript/src/state_store.ts`) and Go (`go/llmapi/state_store.go`).
* Added Agent Memory Store card with interactive key-value storage and live refresh controls into the Unified WebUI Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.80.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).
