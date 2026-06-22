use std::error::Error;
use futures_util::future::{select_ok, join_all};
use crate::llm_api::{LlmRequest, LlmResponse};
use crate::llm_router::LlmRouter;

pub struct Orchestrator;

type AsyncError = Box<dyn Error + Send + Sync>;

impl Orchestrator {
    pub async fn consensus(routers: &[LlmRouter], request: &LlmRequest) -> Result<Vec<LlmResponse>, AsyncError> {
        let mut futures = vec![];

        for router in routers {
            let provider = router.provider.clone();
            let api_key = router.api_key.clone();
            let req_clone = request.clone();

            futures.push(Box::pin(async move {
                let r = LlmRouter::new(&provider, &api_key);
                r.send_request(&req_clone).await.map_err(|e| {
                    let boxed: AsyncError = e.to_string().into();
                    boxed
                })
            }));
        }

        let results = join_all(futures).await;

        let successful_responses: Vec<LlmResponse> = results.into_iter().filter_map(|r| r.ok()).collect();

        if successful_responses.is_empty() {
             return Err("All consensus requests failed".into());
        }

        Ok(successful_responses)
    }

    pub async fn race(routers: &[LlmRouter], request: &LlmRequest) -> Result<LlmResponse, AsyncError> {
        let mut futures = vec![];

        for router in routers {
            let provider = router.provider.clone();
            let api_key = router.api_key.clone();
            let req_clone = request.clone();

            futures.push(Box::pin(async move {
                let r = LlmRouter::new(&provider, &api_key);
                r.send_request(&req_clone).await.map_err(|e| {
                    let boxed: AsyncError = e.to_string().into();
                    boxed
                })
            }));
        }

        if futures.is_empty() {
            let err: AsyncError = "No routers provided for race loop".into();
            return Err(err);
        }

        // select_ok resolves with the first Ok result, discarding errors until all fail
        match select_ok(futures).await {
            Ok((result, _remaining_futures)) => Ok(result), // remaining futures are dropped, naturally cancelling them in Rust
            Err(e) => Err(e),
        }
    }
}
