# IDEAS

## Aggressive Ideas & Future Pivots

1. **Native OS Tray Daemon with Webview Overlay:**
   - Package a universal lightweight tray icon (using `tauri` or `wry`/`systray` in Rust or Go) that docks in the OS menu bar / notification tray.
   - Click to reveal a floating, glassmorphic HUD dashboard summarizing agent workloads, active background tasks, AST index updates, and token consumption.
   - Includes quick actions: Pause, Resume, Inspect Logs, and Graceful Shutdown of all background harness servers.

2. **Self-Balancing LLM Router:**
   - Incorporate live latency, cost, and rate-limit tracking into the LLM router across all 5 languages.
   - Automatically fall back from high-tier models to fast local models (Ollama/llama.cpp) if network conditions deteriorate or quotas are exhausted.

3. **Multi-Agent Coding Competitions (Worktree Racing):**
   - Allow the orchestrator to spin up separate git worktrees, tasking 3 different models with implementing a feature.
   - Run the unit test suite automatically against each worktree; the first model passing all tests with lowest token cost gets auto-merged into main.

4. **Bi-directional Live AST Synchronization:**
   - Watch the filesystem for modifications in real-time, incrementally recalculating syntax definitions and dependency maps so AI agents always have up-to-the-millisecond cognitive awareness without re-scanning.
