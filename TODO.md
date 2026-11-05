# TODO

## Immediate Tasks

- [x] Initialize core documentation files (VISION, MEMORY, DEPLOY, IDEAS, CHANGELOG, VERSION, ROADMAP, TODO, HANDOFF).
- [x] Audit the current C++ codebase to decide what legacy structures to archive or adapt.
- [x] Phase 1: Ingest and analyze submodules (Aider, Claude Code, Pi-Mono, Rowboat).
- [x] Phase 2: Create Protocol Abstraction Layer (`schema/llm_api_protocol.json`).
- [x] Phase 2: Implement Unified LLM API wrapper in Rust (`rust/`).
- [x] Phase 2: Implement Unified LLM API wrapper in Go (`go/`).
- [x] Phase 2: Implement Unified LLM API wrapper in C# (`csharp/`).
- [x] Phase 2: Implement Unified LLM API wrapper in Java (`java/`).
- [x] Phase 2: Implement Unified LLM API wrapper in TypeScript (`typescript/`).
- [x] Phase 3: Define LLM Routing Interface (`schema/router_interface.md`).
- [x] Phase 3: Implement LLM Router in Rust.
- [x] Phase 3: Implement LLM Router in Go.
- [x] Phase 3: Implement LLM Router in C#.
- [x] Phase 3: Implement LLM Router in Java.
- [x] Phase 3: Implement LLM Router in TypeScript.
- [x] Phase 4: Abstract the Multi-Agent Orchestrator schema (`schema/orchestrator.md`).
- [x] Phase 4: Implement Orchestrator Logic in Rust.
- [x] Phase 4: Implement Orchestrator Logic in Go.
- [x] Phase 4: Implement Orchestrator Logic in C#.
- [x] Phase 4: Implement Orchestrator Logic in Java.
- [x] Phase 4: Implement Orchestrator Logic in TypeScript.
- [x] Phase 5: Abstract the AST Parsing Interface (`schema/ast_parser_interface.md`).
- [x] Phase 5: Implement AST Parser in Rust.
- [x] Phase 5: Implement AST Parser in Go.
- [x] Phase 5: Implement AST Parser in C#.
- [x] Phase 5: Implement AST Parser in Java.
- [x] Phase 5: Implement AST Parser in TypeScript.
- [x] Phase 6: Document UI dashboard consolidation and OS system tray requirements.
- [x] Phase 6: Implement Unified Single-Page WebUI Dashboard connecting orchestrator, router, and AST parser.
- [x] Phase 6: Implement System Tray daemon / background server process supervisor with one-click quit functionality.
- [x] Implement Git Automation dynamic version hooks reading from `VERSION.md` (`scripts/version_hook.py`).
- [x] Implement real-time Server-Sent Events (SSE) telemetry and live log streaming in WebUI dashboard.
- [x] Define MCP server specification (`schema/mcp_protocol.md`) and implement MCP server in TypeScript and Go.
- [x] Implement MCP server handlers in Rust, C#, and Java (100% 5-language parity achieved).
- [ ] Connect desktop native tray wrappers (Tauri/Systray).
