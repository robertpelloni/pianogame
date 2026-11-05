# CHANGELOG

See `VERSION.md` for the current active version string.

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

## [v5.72.0]

### Added
* Consolidated build and test runner scripts in `scripts/build_and_verify_all.sh` and `scripts/run_all.sh` across all 5 language harness engines (Rust, Go, C#, Java, TypeScript).
* Created autonomous self-created multi-language verification tooling `/home/jules/self_created_tools/verify_all_harnesses.py`.
* Detailed dashboard layout consolidation and OS-level system tray integration specifications in `ROADMAP.md` and `TODO.md`.

### Changed
* Synchronized documentation governance files (`VERSION.md`, `CHANGELOG.md`, `ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`, `DEPLOY.md`, `IDEAS.md`, `VISION.md`).

## [v5.71.0]

### Added
* Implemented AST parsing capabilities across Rust, Go, C#, Java, and TypeScript for cognitive codebase mapping.
* Initialized core documentation files (`VISION.md`, `MEMORY.md`, `DEPLOY.md`, `IDEAS.md`, `CHANGELOG.md`, `VERSION.md`, `ROADMAP.md`, `TODO.md`, `HANDOFF.md`) to establish the new architecture directives.

### Changed
* Transitioned project scope from C++ PianoGame to the multi-language Ultimate Agentic Coding Harness.
