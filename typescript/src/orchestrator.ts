import { LlmRequest, LlmResponse } from "./llmApi";
import { LlmRouter } from "./llmRouter";

export class Orchestrator {
  /**
   * Consensus Loop: Dispatches requests to all provided routers concurrently,
   * waits for all of them to resolve using Promise.all, and returns a unified array of responses.
   */
  public async consensus(routers: LlmRouter[], request: LlmRequest): Promise<LlmResponse[]> {
    const promises = routers.map((router) => {
      const r = new LlmRouter(router.provider, router.apiKey);
      return r.sendRequest(request);
    });

    try {
      const results = await Promise.all(promises);
      return results;
    } catch (e) {
      throw new Error(`One or more consensus requests failed: ${e}`);
    }
  }

  /**
   * Race Loop: Dispatches requests to all routers concurrently,
   * and returns the very first successful response it receives using Promise.any.
   * Unfinished requests are aggressively cancelled via AbortController.
   */
  public async race(routers: LlmRouter[], request: LlmRequest): Promise<LlmResponse> {
    if (!routers || routers.length === 0) {
      throw new Error("No routers provided for race loop");
    }

    const controller = new AbortController();

    const promises = routers.map((router) => {
      const r = new LlmRouter(router.provider, router.apiKey);

      // Pass the abort signal down to cancel competing fetch requests
      return r.sendRequest(request, controller.signal).then((res) => {
        // As soon as one succeeds, we send the abort signal to cancel the others
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
