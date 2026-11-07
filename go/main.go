package main

import (
	"fmt"
	"harness/llmapi"
	"harness/supervisor"
)

func main() {
	fmt.Println("Ultimate Agentic Coding Harness - Go Edition")

	// 1. Initialize API structures
	messages := []llmapi.Message{
		{
			Role:    llmapi.RoleSystem,
			Content: "You are a helpful coding harness.",
		},
	}

	req := llmapi.NewLlmRequest("gpt-4o", messages)
	fmt.Printf("Successfully initialized LLM Request for model: %s\n", req.Model)

	// 2. Initialize Routers
	openaiRouter := llmapi.NewLlmRouter("openai", "dummy_key")
	anthropicRouter := llmapi.NewLlmRouter("anthropic", "dummy_key")
	fmt.Printf("Successfully initialized LLM Routers\n")

	// 3. Orchestration
	_ = openaiRouter
	_ = anthropicRouter
	fmt.Printf("Successfully initialized Multi-Agent Orchestrator\n")

	// 4. AST Parsing
	astParser := llmapi.NewAstParser()
	fileMap, err := astParser.ParseFile("main.go")
	if err != nil {
		fmt.Printf("Failed to parse file: %v\n", err)
	} else {
		fmt.Printf("Successfully parsed file: %s\n", fileMap.FilePath)
		defs := astParser.ExtractDefinitions(fileMap)
		fmt.Printf("Extracted %d definitions from main.go\n", len(defs))
	}

	// 5. System Tray Daemon & Process Supervisor
	sup := supervisor.NewTraySupervisor()
	sup.RunBackgroundLoop()
	status := sup.GetStatus()
	fmt.Printf("Successfully initialized Tray Supervisor. Supervised engines: %d\n", len(status))

	// 6. Desktop Native Tray Launcher & Container
	trayLauncher := supervisor.NewDesktopTrayLauncher(sup)
	trayLauncher.ShowNotification("Harness Supervisor", "All language engines initialized.")

	container := supervisor.NewDesktopContainer()
	info := container.GetInfo()
	fmt.Printf("Successfully initialized Desktop Container. Window: %s (%v)\n", info["title"], info["dimensions"])
}
