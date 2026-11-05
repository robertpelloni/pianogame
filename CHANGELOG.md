# CHANGELOG

See `VERSION.md` for the current active version string.

## [v5.77.0]

### Added
* Completed full multi-language MCP (Model Context Protocol) Server implementations across Rust (`rust/src/mcp_server.rs`), C# (`csharp/Harness/McpServer.cs`), and Java (`java/src/main/java/com/harness/McpServer.java`), achieving 100% parity with TypeScript and Go.
* Registered MCP tool capabilities (`parse_ast`, `route_query`) across all 5 language runtimes.

### Changed
* Updated `VERSION.md` to `v5.77.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.76.0]

### Added
* Defined Cross-Language Model Context Protocol (MCP) Server specification in `schema/mcp_protocol.md`.
* Implemented MCP Server in TypeScript (`typescript/src/mcp_server.ts`) and Go (`go/llmapi/mcp_server.go`) supporting tool listings and JSON-RPC dispatch.
* Integrated MCP Tool Discovery card into the Unified WebUI Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.76.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).
