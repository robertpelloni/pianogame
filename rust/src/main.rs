pub mod llm_api;

fn main() {
    println!("Ultimate Agentic Coding Harness - Rust Edition");

    // Quick test to ensure the structs can be instantiated
    let req = llm_api::LlmRequest {
        model: "gpt-4o".to_string(),
        messages: vec![
            llm_api::Message {
                role: llm_api::Role::System,
                content: "You are a helpful coding harness.".to_string(),
                tool_calls: None,
            },
        ],
        temperature: 0.7,
        stream: true,
    };

    println!("Successfully initialized LLM Request for model: {}", req.model);
}
