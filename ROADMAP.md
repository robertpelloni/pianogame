# ROADMAP

## Long-term Structural Milestones

1. **Phase 1: Foundation and Ingestion**
   - Initialize strict documentation governance.
   - Begin submodule ingestion of target tools (Aider, Claude Code, Cursor/Codex CLIs, etc.).
   - Analyze ingested tools for feature extraction.

2. **Phase 2: Multi-Language Scaffolding**
   - Set up base project structures for TypeScript, Rust, Go, C#, and Java.
   - Establish inter-language testing and parity checks to guarantee 100% feature matching.

3. **Phase 3: Core Feature Porting**
   - Implement TUI, CLI, and WebUI interfaces for all 5 languages.
   - Implement session management, memory systems, and chat histories.

4. **Phase 4: Advanced Capabilities**
   - Develop the browser extension and MCP server functionalities.
   - Implement external integrations: email parsing, debug console scraping, and web search.

5. **Phase 5: Cognitive AST Codebase Mapping**
   - AST Parser implementations across all 5 languages (Rust, Go, C#, Java, TypeScript).
   - Code definition extraction and dependency graph building.

6. **Phase 6: Unified Dashboard UI & System Tray Integration (Umbrella v6)**
   - Consolidate dashboard architecture into a singular, unified control surface without redundant subpages or navigation fragmentation.
   - Cross-Engine Native OS system tray daemon (`go/supervisor/tray_supervisor.go`, `go/supervisor/desktop_tray.go`) with notification dispatch and graceful multi-engine termination (`QuitAll`).
   - Unified Mission Control Dashboard (`typescript/src/dashboard.ts`, `dashboard.html`) featuring:
     - Multi-Agent Orchestrator (consensus, race, worktree strategies).
     - Adaptive and self-balancing LLM routers dynamically tracking provider latency, token costs, and error rates.
     - Cross-Engine Prompt Template Studio (`schema/prompt_template_protocol.md`) with variable interpolation and validation.
     - Model Context Protocol (MCP) server integration (`schema/mcp_protocol.md`) across all 5 engines.
     - Real-time Server-Sent Events (SSE) telemetry and sanitized live log streaming.
     - Cognitive RAG semantic search and in-memory cosine similarity vector indexing.
     - Multi-agent Git worktree racing orchestrator.
     - Headless browser automation and Chrome DevTools Protocol (CDP) client.
     - Cross-engine SQLite vector and persistent key-value store diagnostics.
     - Remote WebSocket client bridges and desktop webview container launchers.
