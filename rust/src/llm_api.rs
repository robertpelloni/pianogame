use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
#[serde(rename_all = "lowercase")]
pub enum Role {
    System,
    User,
    Assistant,
    Tool,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct FunctionCall {
    pub name: String,
    pub arguments: String,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct ToolCall {
    pub id: String,
    #[serde(rename = "type")]
    pub tool_type: String,
    pub function: FunctionCall,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Message {
    pub role: Role,
    pub content: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub tool_calls: Option<Vec<ToolCall>>,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct LlmRequest {
    pub model: String,
    pub messages: Vec<Message>,
    #[serde(default = "default_temperature")]
    pub temperature: f64,
    #[serde(default = "default_stream")]
    pub stream: bool,
}

fn default_temperature() -> f64 {
    0.7
}

fn default_stream() -> bool {
    true
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Choice {
    pub message: Message,
    pub finish_reason: Option<String>,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Usage {
    pub prompt_tokens: Option<u32>,
    pub completion_tokens: Option<u32>,
    pub total_tokens: Option<u32>,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct LlmResponse {
    pub id: String,
    pub model_used: Option<String>,
    pub choices: Vec<Choice>,
    pub usage: Option<Usage>,
}
