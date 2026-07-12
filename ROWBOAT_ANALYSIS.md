# Rowboat Submodule Analysis

## Overview
The `rowboat` repository is an open-source AI coworker that builds a local-first, long-lived knowledge graph from user data (emails, meetings, notes) and acts on it using an Obsidian-compatible Markdown vault.

## Extracted Core Features for Porting

1.  **Local-First Knowledge Graph (Markdown Vault):**
    -   *Feature:* Maintains an Obsidian-compatible vault of plain Markdown notes with backlinks acting as the agent's transparent "working memory."
    -   *Implementation Goal:* Port the markdown knowledge graph engine into the harness. The Rust/Go/C#/Java/TS implementations must be able to read, parse, interlink, and update local markdown files to act as a compounding persistent memory layer.

2.  **Live Notes & Context Updating:**
    -   *Feature:* Live notes that stay updated automatically by tracking specific topics, people, or projects across web/communications and writing back to the vault.
    -   *Implementation Goal:* Implement background daemon routines in the 5 architectures that continuously monitor integrations (e.g., email, calendars) and execute autonomous note-updating loops.

3.  **Extensive Integrations & MCP Support:**
    -   *Feature:* Connects natively to Gmail, Google Calendar, Fireflies, and supports external tools via MCP (Model Context Protocol) and Composio.
    -   *Implementation Goal:* Expand the MCP implementation we discovered in the Claude Code submodule to also ingest Google APIs and Composio toolchains natively.

4.  **Multi-Modal Output (Voice & Artifact Generation):**
    -   *Feature:* Uses ElevenLabs for voice output and can generate actual artifacts like PDFs and slide decks based on the knowledge graph context.
    -   *Implementation Goal:* Add artifact generators (PDF, voice synthesis endpoints) to the core harness capabilities.

## Implementation Strategy
To achieve feature parity with `rowboat`:
*   *Phase A:* Build the Obsidian-compatible Markdown Parser and Graph Linker across all 5 languages to establish the compounding memory system.
*   *Phase B:* Integrate OAuth and polling loops for Google Services (Gmail/Calendar).
*   *Phase C:* Reimplement the "Live Note" updating logic using the LLM routing capabilities extracted from previous submodules.
