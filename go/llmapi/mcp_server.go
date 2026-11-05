package llmapi

// McpTool describes a tool exposed by the MCP server
type McpTool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"inputSchema"`
}

// McpRequest represents a JSON-RPC 2.0 request
type McpRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// McpResponse represents a JSON-RPC 2.0 response
type McpResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// McpServer handles JSON-RPC Model Context Protocol requests
type McpServer struct {
	tools []McpTool
}

// NewMcpServer constructs a new MCP server with default tools
func NewMcpServer() *McpServer {
	return &McpServer{
		tools: []McpTool{
			{
				Name:        "parse_ast",
				Description: "Extracts AST definitions and dependencies in Go harness",
				InputSchema: map[string]interface{}{"type": "object"},
			},
			{
				Name:        "route_query",
				Description: "Resolves optimal LLM model routing",
				InputSchema: map[string]interface{}{"type": "object"},
			},
		},
	}
}

// ListTools returns the registered MCP tools
func (s *McpServer) ListTools() []McpTool {
	return s.tools
}

// HandleRequest processes incoming MCP JSON-RPC commands
func (s *McpServer) HandleRequest(req McpRequest) McpResponse {
	if req.Method == "initialize" {
		return McpResponse{
			JsonRpc: "2.0",
			ID:      req.ID,
			Result: map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"serverInfo": map[string]string{
					"name":    "GoHarnessMCPServer",
					"version": "1.0.0",
				},
			},
		}
	}

	if req.Method == "tools/list" {
		return McpResponse{
			JsonRpc: "2.0",
			ID:      req.ID,
			Result: map[string]interface{}{
				"tools": s.tools,
			},
		}
	}

	return McpResponse{
		JsonRpc: "2.0",
		ID:      req.ID,
		Error: map[string]interface{}{
			"code":    -32601,
			"message": "Method not found",
		},
	}
}
