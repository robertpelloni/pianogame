package com.harness;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class McpServer {
    public static class McpTool {
        public String name;
        public String description;

        public McpTool(String name, String description) {
            this.name = name;
            this.description = description;
        }
    }

    private final List<McpTool> tools = new ArrayList<>();

    public McpServer() {
        tools.add(new McpTool("parse_ast", "Extracts AST definitions and dependencies in Java harness"));
        tools.add(new McpTool("route_query", "Resolves optimal LLM model routing in Java harness"));
    }

    public List<McpTool> listTools() {
        return Collections.unmodifiableList(tools);
    }

    public Map<String, Object> handleRequest(String method) {
        Map<String, Object> response = new HashMap<>();
        response.put("jsonrpc", "2.0");

        if ("initialize".equals(method)) {
            Map<String, String> serverInfo = new HashMap<>();
            serverInfo.put("name", "JavaHarnessMCPServer");
            serverInfo.put("version", "1.0.0");
            response.put("serverInfo", serverInfo);
            return response;
        }

        if ("tools/list".equals(method)) {
            response.put("tools", tools);
            return response;
        }

        response.put("error", "Method not supported: " + method);
        return response;
    }
}
