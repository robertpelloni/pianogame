import { createLlmRequest, Message } from "./llmApi";
import { LlmRouter } from "./llmRouter";
import { Orchestrator } from "./orchestrator";

console.log("Ultimate Agentic Coding Harness - TypeScript Edition");

// 1. Initialize API structures
const msg: Message = {
  role: "system",
  content: "You are a helpful coding harness."
};

const req = createLlmRequest("gpt-4o", [msg]);
console.log(`Successfully initialized LLM Request for model: ${req.model}`);

// 2. Initialize Routers
const openaiRouter = new LlmRouter("openai", "dummy_key");
const anthropicRouter = new LlmRouter("anthropic", "dummy_key");
console.log(`Successfully initialized LLM Routers`);

// 3. Orchestration
const orchestrator = new Orchestrator();
console.log("Successfully initialized Multi-Agent Orchestrator");

// Example execution logic (commented out to prevent bad HTTP requests with dummy keys)
// async function run() {
//   const routers = [openaiRouter, anthropicRouter];
//   const consensusResult = await orchestrator.consensus(routers, req);
//   const raceResult = await orchestrator.race(routers, req);
// }
