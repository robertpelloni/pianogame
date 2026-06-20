export type Role = "system" | "user" | "assistant" | "tool";

export interface FunctionCall {
  name: string;
  arguments: string;
}

export interface ToolCall {
  id: string;
  type: "function";
  function: FunctionCall;
}

export interface Message {
  role: Role;
  content: string;
  tool_calls?: ToolCall[];
}

export interface LlmRequest {
  model: string;
  messages: Message[];
  temperature?: number;
  stream?: boolean;
}

export interface Choice {
  message: Message;
  finish_reason?: string;
}

export interface Usage {
  prompt_tokens?: number;
  completion_tokens?: number;
  total_tokens?: number;
}

export interface LlmResponse {
  id: string;
  model_used?: string;
  choices: Choice[];
  usage?: Usage;
}

export function createLlmRequest(model: string, messages: Message[]): LlmRequest {
  return {
    model,
    messages,
    temperature: 0.7,
    stream: true,
  };
}
