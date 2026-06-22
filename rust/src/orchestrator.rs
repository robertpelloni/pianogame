use std::error::Error;
use futures_util::future::select_all;
use crate::llm_api::{LlmRequest, LlmResponse};
use crate::llm_router::LlmRouter;

pub struct Orchestrator;

// We need to ensure our Box<dyn Error> is Send so tokio::spawn can push it across threads
type AsyncError = Box<dyn Error + Send + Sync>;

impl Orchestrator {
    pub async fn consensus(routers: &[LlmRouter], request: &LlmRequest) -> Result<Vec<LlmResponse>, AsyncError> {
        let mut tasks = vec![];

        for router in routers {
            let provider = router.provider.clone();
            let api_key = router.api_key.clone();
            let req_clone = request.clone();

            tasks.push(tokio::spawn(async move {
                let r = LlmRouter::new(&provider, &api_key);
                r.send_request(&req_clone).await.map_err(|e| {
                    let boxed: AsyncError = e.to_string().into();
                    boxed
                })
            }));
        }

        let mut results = vec![];
        for task in tasks {
            if let Ok(Ok(response)) = task.await {
                results.push(response);
            }
        }

        Ok(results)
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

        let (result, _index, _remaining) = select_all(futures).await;

        result
    }
}
