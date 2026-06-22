pub mod llm_api;
pub mod llm_router;
pub mod orchestrator;

use llm_api::{LlmRequest, Message, Role};
use llm_router::LlmRouter;

#[tokio::main]
async fn main() {
    println!("Ultimate Agentic Coding Harness - Rust Edition");

    // 1. Initialize API structures
    let req = LlmRequest {
        model: "gpt-4o".to_string(),
        messages: vec![
            Message {
                role: Role::System,
                content: "You are a helpful coding harness.".to_string(),
                tool_calls: None,
            },
        ],
        temperature: 0.7,
        stream: true,
    };
    println!("Successfully initialized LLM Request for model: {}", req.model);

    // 2. Initialize Routers
    let openai_router = LlmRouter::new("openai", "dummy_key");
    let anthropic_router = LlmRouter::new("anthropic", "dummy_key");
    println!("Successfully initialized LLM Routers");

    // 3. Orchestration
    // Uncomment these to actually execute network logic if valid keys are provided.
    // let routers = vec![openai_router, anthropic_router];
    // let consensus_result = orchestrator::Orchestrator::consensus(&routers, &req).await;
    // let race_result = orchestrator::Orchestrator::race(&routers, &req).await;

    println!("Successfully initialized Multi-Agent Orchestrator");
}
