package com.harness;

import com.fasterxml.jackson.databind.ObjectMapper;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.function.Consumer;

public class LlmRouter {

    private final String provider;
    private final String apiKey;
    private final String baseUrl;
    private final HttpClient httpClient;
    private final ObjectMapper objectMapper;

    public LlmRouter(String provider, String apiKey) {
        this.provider = provider;
        this.apiKey = apiKey;

        if ("anthropic".equalsIgnoreCase(provider)) {
            this.baseUrl = "https://api.anthropic.com/v1/messages";
        } else {
            this.baseUrl = "https://api.openai.com/v1/chat/completions";
        }

        this.httpClient = HttpClient.newHttpClient();
        this.objectMapper = new ObjectMapper();
    }

    public String getProvider() {
        return provider;
    }

    public String getApiKey() {
        return apiKey;
    }

    public LlmApi.LlmResponse sendRequest(LlmApi.LlmRequest request) throws Exception {
        String requestBody = objectMapper.writeValueAsString(request);

        HttpRequest.Builder requestBuilder = HttpRequest.newBuilder()
                .uri(URI.create(baseUrl))
                .POST(HttpRequest.BodyPublishers.ofString(requestBody))
                .header("Content-Type", "application/json");

        if ("anthropic".equalsIgnoreCase(provider)) {
            requestBuilder.header("x-api-key", apiKey);
            requestBuilder.header("anthropic-version", "2023-06-01");
        } else {
            requestBuilder.header("Authorization", "Bearer " + apiKey);
        }

        HttpRequest httpRequest = requestBuilder.build();

        HttpResponse<String> response = httpClient.send(httpRequest, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() < 200 || response.statusCode() >= 300) {
            throw new RuntimeException("API request failed with status: " + response.statusCode());
        }

        return objectMapper.readValue(response.body(), LlmApi.LlmResponse.class);
    }

    public void streamResponse(LlmApi.LlmRequest request, Consumer<String> onChunk) {
        // Placeholder for SSE parsing logic
    }

    public String handleToolCall(LlmApi.ToolCall toolCall) {
        // Placeholder for local function routing
        return "{}";
    }
}
