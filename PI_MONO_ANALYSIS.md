# Pi Agent Harness (pi-mono) Submodule Analysis

## Overview
The `pi-mono` repository is an interactive coding agent CLI and agent runtime with tool calling and state management, organized as a monorepo containing multiple packages (`ai`, `agent-core`, `coding-agent`, `tui`).

## Extracted Core Features for Porting

1.  **Unified Multi-Provider LLM API (`packages/ai`):**
    -   *Feature:* A consolidated library handling OpenAI, Anthropic, Google, and other models.
    -   *Implementation Goal:* Port this into a unified cross-language API wrapper so the harness can swap providers natively across all 5 target architectures.

2.  **Terminal UI Library with Differential Rendering (`packages/tui`):**
    -   *Feature:* Advanced TUI that only renders differences to avoid flickering in terminal environments.
    -   *Implementation Goal:* High-priority port for the TUI interface of our multi-language harness. Ensure Rust, Go, C#, Java, and TS implementations can natively differential-render to the console.

3.  **Agent Runtime and State Management (`packages/agent`):**
    -   *Feature:* Tool calling logic and state lifecycle management for the interactive coding loops.
    -   *Implementation Goal:* Extract the state machines used by Pi and implement identical FSMs (Finite State Machines) in the 5 architectures.

4.  **Security and Containerization Paradigms:**
    -   *Feature:* Supports running the agent inside local Linux micro-VMs (Gondolin), Plain Docker, or OpenShell sandboxes.
    -   *Implementation Goal:* The harness must support sandbox abstraction out-of-the-box. Integrate Docker/VM orchestration commands as core tool configurations.

5.  **Strict Project Governance (`AGENTS.md` / `CONTRIBUTING.md`):**
    -   *Feature:* Enforces strict rules on agentic behavior (e.g., no emojis, no `any`, mandatory check scripts).
    -   *Implementation Goal:* We have already adopted a variation of this with our `MEMORY.md`, `VISION.md`, and strict continuous execution rules.

## Implementation Strategy
To achieve feature parity with `pi-mono`:
*   *Phase A:* Port the TUI differential rendering engine into the target languages to serve as the baseline interactive client.
*   *Phase B:* Abstract the Unified LLM API patterns and integrate them with the LiteLLM/MCP routing discovered in previous submodule ingestions.
*   *Phase C:* Reimplement the robust runtime state management to orchestrate the continuous execution loops safely.
