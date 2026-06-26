import { LlmRequest, LlmResponse, ToolCall } from "./llmApi";

export class LlmRouter {
  public provider: string;
  public apiKey: string;
  public baseUrl: string;

  constructor(provider: string, apiKey: string) {
    this.provider = provider;
    this.apiKey = apiKey;

    if (provider.toLowerCase() === "anthropic") {
      this.baseUrl = "https://api.anthropic.com/v1/messages";
    } else {
      this.baseUrl = "https://api.openai.com/v1/chat/completions";
    }
  }

  public async sendRequest(request: LlmRequest, signal?: AbortSignal): Promise<LlmResponse> {
    const headers: Record<string, string> = {
      "Content-Type": "application/json"
    };

    if (this.provider.toLowerCase() === "anthropic") {
      headers["x-api-key"] = this.apiKey;
      headers["anthropic-version"] = "2023-06-01";
    } else {
      headers["Authorization"] = `Bearer ${this.apiKey}`;
    }

    const response = await fetch(this.baseUrl, {
      method: "POST",
      headers,
      body: JSON.stringify(request),
      signal
    });

    if (!response.ok) {
      throw new Error(`API request failed with status: ${response.status}`);
    }

    // In a full implementation, provider-specific parsing happens here.
    return response.json() as Promise<LlmResponse>;
  }

  public streamResponse(request: LlmRequest, onChunk: (chunk: string) => void): void {
    // Placeholder for SSE parsing logic
  }

  public handleToolCall(toolCall: ToolCall): string {
    // Placeholder for local function routing
    return "{}";
  }
}
