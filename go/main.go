package main

import (
	"fmt"
	"harness/llmapi"
)

func main() {
	fmt.Println("Ultimate Agentic Coding Harness - Go Edition")

	// Quick test to ensure structs can be instantiated
	messages := []llmapi.Message{
		{
			Role:    llmapi.RoleSystem,
			Content: "You are a helpful coding harness.",
		},
	}

	req := llmapi.NewLlmRequest("gpt-4o", messages)

	fmt.Printf("Successfully initialized LLM Request for model: %s\n", req.Model)
}
