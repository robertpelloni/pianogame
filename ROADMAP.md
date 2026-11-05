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

6. **Phase 6: Unified Dashboard UI & System Tray Integration**
   - Consolidate dashboard architecture into a singular, unified control surface without redundant subpages.
   - Native OS system tray daemon / tray icon integration across Win32 (`Shell_NotifyIcon`), macOS (`NSStatusBar`), and Linux (`libappindicator`/freedesktop tray status notifier).
   - In-app live status monitoring, server process supervision, and one-click graceful quit capabilities.
   - WebUI dashboard interface for the Ultimate Agentic Coding Harness providing visual orchestrator status, router metrics, and AST code graphs.
