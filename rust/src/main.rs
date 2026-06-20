pub mod llm_api;
pub mod llm_router;

use llm_api::{LlmRequest, Message, Role};
use llm_router::LlmRouter;

#[tokio::main]
async fn main() {
    println!("Ultimate Agentic Coding Harness - Rust Edition");

    // Quick test to ensure the structs can be instantiated
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

    // Initialize the new LLM Router
    let router = LlmRouter::new("openai", "dummy_key");
    println!("Successfully initialized LLM Router for provider: {}", router.provider);
}
