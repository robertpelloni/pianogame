# CHANGELOG

See `VERSION.md` for the current active version string.

## [v5.76.0]

### Added
* Defined Cross-Language Model Context Protocol (MCP) Server specification in `schema/mcp_protocol.md`.
* Implemented MCP Server in TypeScript (`typescript/src/mcp_server.ts`) and Go (`go/llmapi/mcp_server.go`) supporting tool listings and JSON-RPC dispatch.
* Integrated MCP Tool Discovery card into the Unified WebUI Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.76.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

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
