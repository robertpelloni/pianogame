import { createLlmRequest, Message } from "./llmApi";
import { LlmRouter } from "./llmRouter";

console.log("Ultimate Agentic Coding Harness - TypeScript Edition");

const msg: Message = {
  role: "system",
  content: "You are a helpful coding harness."
};

const req = createLlmRequest("gpt-4o", [msg]);
console.log(`Successfully initialized LLM Request for model: ${req.model}`);

const router = new LlmRouter("openai", "dummy_key");
console.log(`Successfully initialized LLM Router for provider: ${router.provider}`);
