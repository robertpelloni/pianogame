package llmapi

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"` // Must be "function"
	Function FunctionCall `json:"function"`
}

type Message struct {
	Role      Role       `json:"role"`
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

type LlmRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	Stream      bool      `json:"stream"`
}

func NewLlmRequest(model string, messages []Message) *LlmRequest {
	return &LlmRequest{
		Model:       model,
		Messages:    messages,
		Temperature: 0.7,
		Stream:      true,
	}
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason *string `json:"finish_reason,omitempty"`
}

type Usage struct {
	PromptTokens     *int `json:"prompt_tokens,omitempty"`
	CompletionTokens *int `json:"completion_tokens,omitempty"`
	TotalTokens      *int `json:"total_tokens,omitempty"`
}

type LlmResponse struct {
	ID        string  `json:"id"`
	ModelUsed *string `json:"model_used,omitempty"`
	Choices   []Choice `json:"choices"`
	Usage     *Usage   `json:"usage,omitempty"`
}
