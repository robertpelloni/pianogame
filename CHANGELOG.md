# CHANGELOG

See `VERSION.md` for the current active version string.

## [v5.75.0]

### Added
* Implemented real-time Server-Sent Events (SSE) `/api/events` telemetry stream in `typescript/src/dashboard.ts`.
* Added interactive streaming log terminal in `typescript/src/dashboard.html` to monitor router queries, multi-agent orchestrator cycles, AST cognitive parsing, and supervisor signals live.

### Changed
* Updated `VERSION.md` to `v5.75.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Updated documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.74.0]

### Added
* Implemented Git Automation Dynamic Version Hook (`scripts/version_hook.py`) to automatically read `VERSION.md` and propagate SemVer strings across `typescript/package.json`, `rust/Cargo.toml`, `java/pom.xml`, and `csharp/Harness/Harness.csproj`.
* Added version consistency checking flag (`--check`) to enforce package manifest synchronization during automated builds.

### Changed
* Updated `VERSION.md` to `v5.74.0`.
* Synchronized documentation files (`CHANGELOG.md`, `ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.73.0]

### Added
* Implemented Unified Single-Page Web Dashboard UI in `typescript/src/dashboard.ts` and `typescript/src/dashboard.html` with interactive forms, status badges, contextual tooltips, and API routes for LLM router, multi-agent orchestrator, cognitive AST parsing, and process supervision.
* Implemented System Tray Daemon & Process Supervisor in `go/supervisor/tray_supervisor.go` supporting cross-engine monitoring, OS interrupt signals, and one-click graceful termination (`QuitAll`).
* Added `AstParser.parseSource` to the TypeScript harness for cognitive code parsing from interactive UI payloads.

### Changed
* Updated `go/main.go` to initialize the Tray Supervisor and run background daemon monitoring.
* Synchronized documentation files (`VERSION.md`, `CHANGELOG.md`, `ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).
