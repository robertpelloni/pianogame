using System;
using System.Collections.Generic;

namespace Harness
{
    public class McpTool
    {
        public string Name { get; set; } = string.Empty;
        public string Description { get; set; } = string.Empty;
        public object InputSchema { get; set; } = new object();
    }

    public class McpServer
    {
        private readonly List<McpTool> _tools;

        public McpServer()
        {
            _tools = new List<McpTool>
            {
                new McpTool
                {
                    Name = "parse_ast",
                    Description = "Extracts AST definitions and dependencies in C# harness",
                    InputSchema = new { type = "object" }
                },
                new McpTool
                {
                    Name = "route_query",
                    Description = "Resolves optimal LLM model routing in C# harness",
                    InputSchema = new { type = "object" }
                }
            };
        }

        public IReadOnlyList<McpTool> ListTools() => _tools.AsReadOnly();

        public object HandleRequest(string method, Dictionary<string, object>? parameters = null)
        {
            if (method == "initialize")
            {
                return new
                {
                    protocolVersion = "2024-11-05",
                    serverInfo = new { name = "CSharpHarnessMCPServer", version = "1.0.0" },
                    capabilities = new { tools = new { listChanged = true } }
                };
            }

            if (method == "tools/list")
            {
                return new { tools = _tools };
            }

            return new { error = $"Method not supported: {method}" };
        }
    }
}
