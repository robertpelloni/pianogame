package com.harness;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;

public class LlmApi {

    public enum Role {
        @JsonProperty("system") SYSTEM,
        @JsonProperty("user") USER,
        @JsonProperty("assistant") ASSISTANT,
        @JsonProperty("tool") TOOL
    }

    public static class FunctionCall {
        @JsonProperty("name")
        public String name;

        @JsonProperty("arguments")
        public String arguments;
    }

    public static class ToolCall {
        @JsonProperty("id")
        public String id;

        @JsonProperty("type")
        public String type = "function";

        @JsonProperty("function")
        public FunctionCall function;
    }

    @JsonInclude(JsonInclude.Include.NON_NULL)
    public static class Message {
        @JsonProperty("role")
        public Role role;

        @JsonProperty("content")
        public String content;

        @JsonProperty("tool_calls")
        public List<ToolCall> toolCalls;
    }

    public static class LlmRequest {
        @JsonProperty("model")
        public String model;

        @JsonProperty("messages")
        public List<Message> messages;

        @JsonProperty("temperature")
        public double temperature = 0.7;

        @JsonProperty("stream")
        public boolean stream = true;
    }

    @JsonInclude(JsonInclude.Include.NON_NULL)
    public static class Choice {
        @JsonProperty("message")
        public Message message;

        @JsonProperty("finish_reason")
        public String finishReason;
    }

    @JsonInclude(JsonInclude.Include.NON_NULL)
    public static class Usage {
        @JsonProperty("prompt_tokens")
        public Integer promptTokens;

        @JsonProperty("completion_tokens")
        public Integer completionTokens;

        @JsonProperty("total_tokens")
        public Integer totalTokens;
    }

    @JsonInclude(JsonInclude.Include.NON_NULL)
    public static class LlmResponse {
        @JsonProperty("id")
        public String id;

        @JsonProperty("model_used")
        public String modelUsed;

        @JsonProperty("choices")
        public List<Choice> choices;

        @JsonProperty("usage")
        public Usage usage;
    }
}
