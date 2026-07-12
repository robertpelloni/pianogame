# Aider Submodule Analysis

## Overview
Aider is an AI pair programming tool implemented in Python that integrates tightly with Git. It acts as an agent inside a local repository, allowing users to pair program with LLMs.

## Extracted Core Features for Porting

1.  **Codebase Mapping & AST Parsing:**
    -   *Feature:* Uses tree-sitter or similar to map the entire codebase to handle large projects.
    -   *Implementation Goal:* Port AST mapping logic into Rust, Go, C#, Java, and TypeScript versions.

2.  **Git Integration:**
    -   *Feature:* Automatically commits changes with AI-generated, sensible commit messages (`aider/git.py`).
    -   *Implementation Goal:* Implement native Git operations and commit-message generation loops in all 5 target architectures.

3.  **Multi-Model Support:**
    -   *Feature:* Supports OpenAI, Anthropic, DeepSeek, and OpenRouter natively (`aider/models.py`).
    -   *Implementation Goal:* Centralize API routing logic. Abstract the API calls so our multi-language harness can hot-swap models without altering the core prompt logic.

4.  **Auto Linting & Testing:**
    -   *Feature:* Automatically runs tests/linters after AI edits and feeds errors back into the LLM context for self-correction.
    -   *Implementation Goal:* Build a robust subprocess execution manager in each target language that parses test/lint outputs recursively.

5.  **Multi-Modal Inputs (Images, Voice):**
    -   *Feature:* Accepts images for visual context and supports voice-to-code inputs.
    -   *Implementation Goal:* The CLI/TUI/WebUI in all languages must support binary file ingestion and STT (Speech-To-Text) plugins.

## Implementation Strategy
To achieve the Ultimate Agentic Coding Harness vision, we will systematically ingest the core Aider logic:
*   *Phase A:* Copy the repository file mapping and Git interaction logic into the 5 target directories.
*   *Phase B:* Reimplement the LLM communication protocols (Litellm wrapper porting).
*   *Phase C:* Reimplement the "Search/Replace" and "Diff" applying mechanisms to ensure accurate code modification across all languages.
