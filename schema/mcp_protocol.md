# Model Context Protocol (MCP) Server Specification

## Overview
The Model Context Protocol (MCP) provides an open, standardized interface enabling AI harnesses to dynamically expose and consume tools, prompts, and context resources over JSON-RPC 2.0.

## Protocol Core

### 1. Initialization (`initialize`)
Client initiates connection specifying protocol version and client capabilities.
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "protocolVersion": "2024-11-05",
    "capabilities": {
      "roots": { "listChanged": true },
      "sampling": {}
    },
    "clientInfo": {
      "name": "UltimateAgenticCodingHarness",
      "version": "1.0.0"
    }
  }
}
```

Server Response:
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "protocolVersion": "2024-11-05",
    "capabilities": {
      "tools": { "listChanged": true },
      "resources": { "subscribe": true },
      "prompts": { "listChanged": true }
    },
    "serverInfo": {
      "name": "HarnessMCPServer",
      "version": "1.0.0"
    }
  }
}
```

### 2. Tools Discovery (`tools/list`)
Lists available executable tools provided by the harness (e.g. AST analysis, router queries, multi-agent runs).
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list"
}
```

### 3. Tool Invocation (`tools/call`)
Executes an MCP tool with parameterized arguments.
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "parse_ast",
    "arguments": {
      "code": "function hello(): void {}"
    }
  }
}
```

### 4. Resources (`resources/list` & `resources/read`)
Exposes project documentation, AST cognitive syntax maps, and harness telemetry streams as MCP resources.
