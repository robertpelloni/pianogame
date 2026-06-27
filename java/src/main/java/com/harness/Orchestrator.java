package com.harness;

import java.util.List;
import java.util.Objects;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.stream.Collectors;

public class Orchestrator {

    public List<LlmApi.LlmResponse> consensus(List<LlmRouter> routers, LlmApi.LlmRequest request) throws Exception {

        List<CompletableFuture<LlmApi.LlmResponse>> futures = routers.stream().map(router -> {
            LlmRouter r = new LlmRouter(router.getProvider(), router.getApiKey());
            return r.sendRequestAsync(request);
        }).collect(Collectors.toList());

        CompletableFuture<Void> allFutures = CompletableFuture.allOf(futures.toArray(new CompletableFuture[0]));

        try {
            allFutures.join();
        } catch (Exception e) {
            throw new Exception("One or more consensus requests failed", e);
        }

        return futures.stream()
                .map(CompletableFuture::join)
                .filter(Objects::nonNull)
                .collect(Collectors.toList());
    }

    public LlmApi.LlmResponse race(List<LlmRouter> routers, LlmApi.LlmRequest request) throws Exception {
        if (routers == null || routers.isEmpty()) {
            throw new IllegalArgumentException("No routers provided for race loop");
        }

        CompletableFuture<LlmApi.LlmResponse> firstSuccess = new CompletableFuture<>();
        AtomicInteger failures = new AtomicInteger(0);
        int totalRouters = routers.size();

        List<CompletableFuture<LlmApi.LlmResponse>> futures = routers.stream().map(router -> {
            LlmRouter r = new LlmRouter(router.getProvider(), router.getApiKey());
            CompletableFuture<LlmApi.LlmResponse> future = r.sendRequestAsync(request);

            future.whenComplete((result, ex) -> {
                if (ex != null) {
                    if (failures.incrementAndGet() == totalRouters) {
                        firstSuccess.completeExceptionally(new Exception("All race requests failed"));
                    }
                } else {
                    firstSuccess.complete(result);
                }
            });
            return future;
        }).collect(Collectors.toList());

        try {
            LlmApi.LlmResponse result = firstSuccess.join();

            for (CompletableFuture<LlmApi.LlmResponse> future : futures) {
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
