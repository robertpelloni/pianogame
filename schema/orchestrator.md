# Multi-Agent Orchestrator Schema

Based on the analysis of the `code` submodule (Claude Code fork), the Ultimate Agentic Coding Harness requires an orchestration layer capable of managing multiple asynchronous LLM operations simultaneously across all 5 target environments.

## Orchestration Patterns

The orchestrator must implement the following multi-agent execution loops:

### 1. `Consensus Loop` (e.g., `/plan`, `/code`)
*   **Description:** Dispatches identical prompts to multiple distinct model providers (e.g., OpenAI, Anthropic, Gemini) concurrently.
*   **Behavior:** The orchestrator waits for all models to return responses. It then evaluates the responses (either heuristically or via a secondary LLM "judge" prompt) and synthesizes a final, unified output representing the consensus of the swarm.

### 2. `Race Loop` (e.g., `/solve`)
*   **Description:** Dispatches identical, highly complex problem-solving prompts to multiple models concurrently.
*   **Behavior:** The orchestrator resolves the loop with the *first valid response* it receives. The other concurrent network requests are immediately cancelled/aborted to save bandwidth and compute cycles.

### 3. `Auto Drive Loop` (e.g., `/auto`)
*   **Description:** The core continuous execution loop.
*   **Behavior:** Hands off a multi-step task to an automated coordinator that:
    1. Triggers the `Consensus Loop` for planning.
    2. Sequentially executes code changes via the `LLM Router`.
    3. Triggers background linters/tests (Aider pattern).
    4. Automatically requests re-evaluations on test failures without user intervention.

## Target Implementations
To ensure universal parity, this logic must be implemented using native asynchronous/concurrent processing specific to each language:
*   **Rust:** `tokio::spawn` and `tokio::select!` (for Race logic).
*   **Go:** Goroutines and `sync.WaitGroup` / channels.
*   **C#:** `Task.WhenAll` and `Task.WhenAny`.
*   **Java:** `CompletableFuture.allOf` and `CompletableFuture.anyOf`.
*   **TypeScript:** `Promise.all` and `Promise.race`.
