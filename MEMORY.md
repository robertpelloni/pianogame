# MEMORY

## Internal Architectural Observations
* The workspace is transitioning from legacy C++ into the Ultimate Agentic Coding Harness.
* Submodule analysis has extracted patterns from `code` (Claude Code / Codex CLI proxy), `aider`, `pi-mono`, and `rowboat`.
* Current implementation spans 5 core languages with identical architectural components:
  - LLM API Protocol (`schema/llm_api_protocol.json`)
  - LLM Router (`schema/router_interface.md`)
  - Multi-Agent Orchestrator (`schema/orchestrator.md`)
  - Cognitive AST Codebase Parser (`schema/ast_parser_interface.md`)
* The TypeScript harness provides a Unified Single-Page Web Dashboard (`typescript/src/dashboard.ts`, `typescript/src/dashboard.html`) exposing interactive endpoints for all harness components.
* The Go harness hosts the System Tray Daemon and Process Supervisor (`go/supervisor/tray_supervisor.go`) providing process monitoring and `QuitAll` graceful termination.
* Scripts are consolidated in `scripts/build_and_verify_all.sh` and `scripts/run_all.sh`.
