# Unified LLM Router Interface

To achieve identical routing behaviors across the 5 language silos (Rust, Go, C#, Java, TypeScript) when interfacing with external API endpoints (OpenAI, Anthropic, Gemini, OpenRouter), each language must implement a Router/Client adhering to this specification.

## Core Interface Requirements

The Router must accept our internal `LlmRequest` object (defined in `llm_api_protocol.json`) and output an `LlmResponse` object.

### Standard Functions

1.  **`Initialize(provider: String, apiKey: String)`**
    *   **Description:** Configures the client with the required base URL and authentication tokens based on the provider (e.g., "openai", "anthropic").

2.  **`SendRequest(request: LlmRequest) -> LlmResponse`**
    *   **Description:** Synchronously sends the request to the upstream provider.
    *   **Behavior:** Must map the internal `LlmRequest` to the provider's specific API schema format before transmission. Upon receiving a response, it must map the provider's payload back into our internal `LlmResponse` format.

3.  **`StreamResponse(request: LlmRequest, onChunk: Function(String))`**
    *   **Description:** Asynchronously streams tokens back to the caller.
    *   **Behavior:** Must parse Server-Sent Events (SSE) from the upstream provider and emit the delta strings.

4.  **`HandleToolCall(toolCall: ToolCall) -> String`**
    *   **Description:** An execution router that maps the LLM's requested `ToolCall` object to local registered functions (e.g., executing bash commands, reading files) and returns the JSON stringified result.

## Provider Mapping Notes
*   **OpenAI:** Nearly 1-to-1 mapping with our internal schema.
*   **Anthropic:** Requires mapping `role: system` messages into a top-level `system` string array, and mapping `user/assistant` interactions.
*   **Google Gemini:** Requires mapping to `contents` and `parts` arrays.
