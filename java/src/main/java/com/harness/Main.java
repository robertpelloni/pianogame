package com.harness;

import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        System.out.println("Ultimate Agentic Coding Harness - Java Edition");

        LlmApi.Message msg = new LlmApi.Message();
        msg.role = LlmApi.Role.SYSTEM;
        msg.content = "You are a helpful coding harness.";

        LlmApi.LlmRequest request = new LlmApi.LlmRequest();
        request.model = "gpt-4o";
        request.messages = new ArrayList<>();
        request.messages.add(msg);

        System.out.println("Successfully initialized LLM Request for model: " + request.model);
    }
}
