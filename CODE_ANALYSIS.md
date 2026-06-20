# Claude Code / Codex CLI Submodule Analysis

## Overview
The `code` repository (fork of Codex CLI / proxy for Claude Code) is an advanced AI CLI tool featuring browser integration, multi-agent orchestration, and Model Context Protocol (MCP) support. It acts as an orchestrator for multiple AI backend models.

## Extracted Core Features for Porting

1.  **Multi-Agent Commands (`/plan`, `/solve`, `/code`):**
    -   *Feature:* Orchestrates multiple models (e.g., Claude, Gemini, GPT) to review tasks, race for the fastest solution, or create worktrees to implement optimal solutions.
    -   *Implementation Goal:* Port this orchestrator pattern. The Rust/Go/C#/Java/TS harnesses must be capable of spinning up multiple background API calls simultaneously, comparing their outputs, and executing the winning strategy.

2.  **Browser Integration (`/chrome`, `/browser`):**
    -   *Feature:* Connects to external Chrome instances via CDP or uses internal headless browsers to gain context.
    -   *Implementation Goal:* Implement a WebDriver/CDP abstraction layer across all 5 architectures to allow the agent to scrape, navigate, and verify frontend changes autonomously.

3.  **Model Context Protocol (MCP):**
    -   *Feature:* Standardized protocol for extending capabilities (file operations, DB connections, API integrations).
    -   *Implementation Goal:* The harness will natively host an MCP Server implementation in all 5 languages, allowing any standard MCP client to connect and utilize the harness's capabilities.

4.  **Auto Drive (`/auto`):**
    -   *Feature:* Hands off a multi-step task to an automated coordinator that runs agents and requests approvals.
    -   *Implementation Goal:* This perfectly aligns with our "Continuous Autonomous Execution" directive. The logic loop for `Auto Drive` will be the default operating mode of our new ultimate harness.

5.  **Project Context Ingestion (`AGENTS.md` / `CLAUDE.md`):**
    -   *Feature:* Automatically reads global context files to understand project structure.
    -   *Implementation Goal:* We have already implemented this via `VISION.md`, `MEMORY.md`, etc. We will port the parser that ingests these files and feeds them into the system prompt.

## Implementation Strategy
To achieve feature parity with `code`:
*   *Phase A:* Analyze the `.toml` configuration parsers and implement cross-language configuration syncing.
*   *Phase B:* Implement the MCP client/server specifications in Rust, Go, C#, Java, and TS.
*   *Phase C:* Build the Multi-Agent Orchestrator module.
