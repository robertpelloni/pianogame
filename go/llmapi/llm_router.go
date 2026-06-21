package llmapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LlmRouter struct {
	Provider string
	ApiKey   string
	BaseUrl  string
	client   *http.Client
}

func NewLlmRouter(provider string, apiKey string) *LlmRouter {
	baseUrl := "https://api.openai.com/v1/chat/completions" // Default
	if provider == "anthropic" {
		baseUrl = "https://api.anthropic.com/v1/messages"
	}

	return &LlmRouter{
		Provider: provider,
		ApiKey:   apiKey,
		BaseUrl:  baseUrl,
		client:   &http.Client{},
	}
}

func (r *LlmRouter) SendRequest(request *LlmRequest) (*LlmResponse, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", r.BaseUrl, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	// Map provider-specific headers
	if r.Provider == "anthropic" {
		httpReq.Header.Set("x-api-key", r.ApiKey)
		httpReq.Header.Set("anthropic-version", "2023-06-01")
		httpReq.Header.Set("Content-Type", "application/json")
	} else {
		httpReq.Header.Set("Authorization", "Bearer "+r.ApiKey)
		httpReq.Header.Set("Content-Type", "application/json")
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
	    return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// In a full implementation, provider-specific parsing happens here.
	var llmResponse LlmResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResponse); err != nil {
		return nil, err
	}

	return &llmResponse, nil
}

func (r *LlmRouter) StreamResponse(request *LlmRequest, onChunk func(string)) error {
	// Placeholder for SSE parsing logic
	return nil
}

func (r *LlmRouter) HandleToolCall(toolCall *ToolCall) string {
	// Placeholder for local function routing
	return "{}"
}
