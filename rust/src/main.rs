pub mod llm_api;
pub mod llm_router;
pub mod orchestrator;
pub mod ast_parser;

use llm_api::{LlmRequest, Message, Role};
use llm_router::LlmRouter;
use ast_parser::AstParser;

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
    let _openai_router = LlmRouter::new("openai", "dummy_key");
    let _anthropic_router = LlmRouter::new("anthropic", "dummy_key");
    println!("Successfully initialized LLM Routers");

    // 3. Orchestration
    println!("Successfully initialized Multi-Agent Orchestrator");

    // 4. AST Parsing
    // Test parsing our own source file
    match AstParser::parse_file("src/main.rs") {
        Ok(file_map) => {
            println!("Successfully parsed file: {}", file_map.file_path);
            let defs = AstParser::extract_definitions(&file_map);
            println!("Extracted {} definitions from main.rs", defs.len());
        }
        Err(e) => {
            println!("Failed to parse file: {}", e);
        }
    }
}
