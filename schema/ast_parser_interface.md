# AST Parser Interface Specification

Based on the feature analysis of the `aider` submodule, our agentic coding harness must possess a robust, native Abstract Syntax Tree (AST) parsing capability in each of the 5 target environments. This parser acts as the cognitive mapping system, allowing the agent to "see" the structure of the repository.

## Core Interface Requirements

The AST Parser must provide language-agnostic mapping capabilities by normalizing project structures into unified data models.

### Standard Functions

1.  **`Initialize(projectRoot: String)`**
    *   **Description:** Configures the parser to map the specified root directory. It must automatically respect `.gitignore` files to avoid parsing build artifacts and dependencies.

2.  **`ParseFile(filePath: String) -> FileSyntaxMap`**
    *   **Description:** Synchronously parses a target source file and returns a structured map of its internal tokens (e.g., classes, functions, interfaces, imports).
    *   **Behavior:** Should utilize Tree-Sitter bindings or native language parser libraries.

3.  **`ExtractDefinitions(fileMap: FileSyntaxMap) -> List<Definition>`**
    *   **Description:** Extracts high-level structural definitions (e.g., `pub fn main()`, `public class LlmRouter`) along with their byte/line offsets to construct a repository "skeleton."

4.  **`MapDependencies(fileMap: FileSyntaxMap) -> List<Dependency>`**
    *   **Description:** Identifies local imports/includes to build a directed acyclic graph (DAG) of the project's internal architecture, crucial for determining the ripple effects of an agentic code change.

## Target Implementations
To ensure universal parity, this AST engine must be implemented using standard tooling:
*   **Rust:** `syn` (for Rust parsing) or `tree-sitter`.
*   **Go:** `go/ast` and `go/parser`.
*   **C#:** `Microsoft.CodeAnalysis` (Roslyn).
*   **Java:** `JavaParser` or native `javax.lang.model`.
*   **TypeScript:** `typescript` compiler API.
