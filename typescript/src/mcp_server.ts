import { AstParser } from "./ast_parser";

export interface McpTool {
  name: string;
  description: string;
  inputSchema: object;
}

export interface McpRequest {
  jsonrpc: "2.0";
  id: number | string;
  method: string;
  params?: any;
}

export interface McpResponse {
  jsonrpc: "2.0";
  id: number | string;
  result?: any;
  error?: {
    code: number;
    message: string;
    data?: any;
  };
}

export class McpServer {
  private tools: Map<string, McpTool> = new Map();

  constructor() {
    this.registerDefaultTools();
  }

  private registerDefaultTools() {
    this.tools.set("parse_ast", {
      name: "parse_ast",
      description: "Extracts AST definitions and dependencies from source code",
      inputSchema: {
        type: "object",
        properties: {
          code: { type: "string", description: "Source code snippet" },
          filename: { type: "string", description: "Optional virtual filename" }
        },
        required: ["code"]
      }
    });

    this.tools.set("route_query", {
      name: "route_query",
      description: "Resolves optimal LLM provider and model for prompt execution",
      inputSchema: {
        type: "object",
        properties: {
          model: { type: "string", description: "Target model identifier" }
        },
        required: ["model"]
      }
    });
  }

  public listTools(): McpTool[] {
    return Array.from(this.tools.values());
  }

  public handleRequest(req: McpRequest): McpResponse {
    if (req.method === "initialize") {
      return {
        jsonrpc: "2.0",
        id: req.id,
        result: {
          protocolVersion: "2024-11-05",
          serverInfo: { name: "TypeScriptHarnessMCPServer", version: "1.0.0" },
          capabilities: { tools: { listChanged: true } }
        }
      };
    }

    if (req.method === "tools/list") {
      return {
        jsonrpc: "2.0",
        id: req.id,
        result: { tools: this.listTools() }
      };
    }

    if (req.method === "tools/call") {
      const toolName = req.params?.name;
      const args = req.params?.arguments || {};

      if (toolName === "parse_ast") {
        const fileMap = AstParser.parseSource(args.filename || "input.ts", args.code || "");
        return {
          jsonrpc: "2.0",
          id: req.id,
          result: {
            definitions: fileMap.definitions,
            dependencies: fileMap.dependencies
          }
        };
      }

      if (toolName === "route_query") {
        return {
          jsonrpc: "2.0",
          id: req.id,
          result: {
            provider: "openai",
            model: args.model || "gpt-4o",
            status: "ready"
          }
        };
      }

      return {
        jsonrpc: "2.0",
        id: req.id,
        error: { code: -32601, message: `Tool not found: ${toolName}` }
      };
    }

    return {
      jsonrpc: "2.0",
      id: req.id,
      error: { code: -32601, message: `Method not found: ${req.method}` }
    };
  }
}
