package com.harness;

import java.util.List;
import java.util.Objects;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.stream.Collectors;

public class Orchestrator {

    /**
     * Consensus Loop: Dispatches requests to all provided routers concurrently,
     * waits for all of them to resolve using CompletableFuture.allOf, and returns a unified array of responses.
     */
    public List<LlmApi.LlmResponse> consensus(List<LlmRouter> routers, LlmApi.LlmRequest request) throws Exception {

        List<CompletableFuture<LlmApi.LlmResponse>> futures = routers.stream().map(router ->
            CompletableFuture.supplyAsync(() -> {
                try {
                    LlmRouter r = new LlmRouter(router.getProvider(), router.getApiKey());
                    return r.sendRequest(request);
                } catch (Exception e) {
                    throw new RuntimeException(e);
                }
            })
        ).collect(Collectors.toList());

        CompletableFuture<Void> allFutures = CompletableFuture.allOf(futures.toArray(new CompletableFuture[0]));

        try {
            allFutures.join();
        } catch (Exception e) {
            // In a production setup, we could filter out failures instead of throwing immediately.
            throw new Exception("One or more consensus requests failed", e);
        }

        return futures.stream()
                .map(CompletableFuture::join)
                .filter(Objects::nonNull)
                .collect(Collectors.toList());
    }

    /**
     * Race Loop: Dispatches requests to all routers concurrently,
     * and returns the very first *successful* response it receives.
     */
    public LlmApi.LlmResponse race(List<LlmRouter> routers, LlmApi.LlmRequest request) throws Exception {
        if (routers == null || routers.isEmpty()) {
            throw new IllegalArgumentException("No routers provided for race loop");
        }

        CompletableFuture<LlmApi.LlmResponse> firstSuccess = new CompletableFuture<>();
        AtomicInteger failures = new AtomicInteger(0);
        int totalRouters = routers.size();

        List<CompletableFuture<Void>> futures = routers.stream().map(router ->
            CompletableFuture.runAsync(() -> {
                try {
                    LlmRouter r = new LlmRouter(router.getProvider(), router.getApiKey());
                    LlmApi.LlmResponse resp = r.sendRequest(request);
                    // Complete the primary future on the first success
                    firstSuccess.complete(resp);
                } catch (Exception e) {
                    // Count failures. If ALL fail, we complete the future exceptionally.
                    if (failures.incrementAndGet() == totalRouters) {
                        firstSuccess.completeExceptionally(new Exception("All race requests failed"));
                    }
                }
            })
        ).collect(Collectors.toList());

        try {
            LlmApi.LlmResponse result = firstSuccess.join();

            // Attempt to aggressively cancel remaining futures to save resources
            for (CompletableFuture<Void> future : futures) {
                if (!future.isDone()) {
                    future.cancel(true);
                }
            }

            return result;
        } catch (Exception e) {
             throw new Exception("All race requests failed", e);
        }
    }
}
