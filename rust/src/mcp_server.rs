use serde::{Deserialize, Serialize};
use serde_json::{json, Value};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct McpTool {
    pub name: String,
    pub description: String,
    pub input_schema: Value,
}

pub struct McpServer {
    tools: Vec<McpTool>,
}

impl McpServer {
    pub fn new() -> Self {
        McpServer {
            tools: vec![
                McpTool {
                    name: "parse_ast".to_string(),
                    description: "Extracts AST definitions and dependencies in Rust harness".to_string(),
                    input_schema: json!({
                        "type": "object",
                        "properties": {
                            "code": { "type": "string" }
                        },
                        "required": ["code"]
                    }),
                },
                McpTool {
                    name: "route_query".to_string(),
                    description: "Resolves optimal LLM model routing in Rust harness".to_string(),
                    input_schema: json!({
                        "type": "object",
                        "properties": {
                            "model": { "type": "string" }
                        },
                        "required": ["model"]
                    }),
                },
            ],
        }
    }

    pub fn list_tools(&self) -> Vec<McpTool> {
        self.tools.clone()
    }

    pub fn handle_request(&self, method: &str, params: Value) -> Value {
        match method {
            "initialize" => json!({
                "protocolVersion": "2024-11-05",
                "serverInfo": {
                    "name": "RustHarnessMCPServer",
                    "version": "1.0.0"
                },
                "capabilities": {
                    "tools": { "listChanged": true }
                }
            }),
            "tools/list" => json!({
                "tools": self.tools
            }),
            "tools/call" => {
                let tool_name = params.get("name").and_then(|v| v.as_str()).unwrap_or("");
                match tool_name {
                    "route_query" => json!({
                        "provider": "openai",
                        "status": "ready"
                    }),
                    "parse_ast" => json!({
                        "definitions": ["main"],
                        "dependencies": ["std"]
                    }),
                    _ => json!({ "error": format!("Tool not found: {}", tool_name) }),
                }
            }
            _ => json!({ "error": format!("Method not found: {}", method) }),
        }
    }
}
