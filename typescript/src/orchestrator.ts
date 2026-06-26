import { LlmRequest, LlmResponse } from "./llmApi";
import { LlmRouter } from "./llmRouter";

export class Orchestrator {
  /**
   * Consensus Loop: Dispatches requests to all provided routers concurrently,
   * waits for all of them to resolve using Promise.all, and returns a unified array of responses.
   */
  public async consensus(routers: LlmRouter[], request: LlmRequest): Promise<LlmResponse[]> {
    const promises = routers.map((router) => {
      // Re-initialize router to mock abstraction isolation
      const r = new LlmRouter(router.provider, router.apiKey);
      return r.sendRequest(request);
    });

    try {
      const results = await Promise.all(promises);
      return results;
    } catch (e) {
      // In a production setup, we could filter out failures using Promise.allSettled
      // instead of throwing immediately.
      throw new Error(`One or more consensus requests failed: ${e}`);
    }
  }

  /**
   * Race Loop: Dispatches requests to all routers concurrently,
   * and returns the very first successful response it receives using Promise.any.
   *
   * NOTE: Promise.any inherently returns the first *fulfilled* promise, skipping
   * rejected promises unless they *all* reject, which is the exact behavior we want
   * for an LLM Race loop.
   */
  public async race(routers: LlmRouter[], request: LlmRequest): Promise<LlmResponse> {
    if (!routers || routers.length === 0) {
      throw new Error("No routers provided for race loop");
    }

    // In a full implementation we would pass an AbortController to the fetch request
    // to actually sever the HTTP connection.
    const controller = new AbortController();

    const promises = routers.map((router) => {
      const r = new LlmRouter(router.provider, router.apiKey);

      // We pass the controller logic down ideally. For now we just run the request.
      return r.sendRequest(request).then((res) => {
        // As soon as one succeeds, we send the abort signal to the others
        controller.abort();
        return res;
      });
    });

    try {
      // Promise.any naturally ignores rejections unless all promises reject.
      return await Promise.any(promises);
    } catch (e) {
      throw new Error("All race requests failed");
    }
  }
}
