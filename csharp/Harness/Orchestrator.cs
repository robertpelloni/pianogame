using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

namespace Harness.LlmApi
{
    public class Orchestrator
    {
        public async Task<List<LlmResponse>> ConsensusAsync(List<LlmRouter> routers, LlmRequest request, CancellationToken cancellationToken = default)
        {
            var tasks = routers.Select(router =>
            {
                var safeRouter = new LlmRouter(router.Provider, router.ApiKey);
                return safeRouter.SendRequestAsync(request, cancellationToken);
            }).ToList();

            try
            {
                var results = await Task.WhenAll(tasks);
                return results.ToList();
            }
            catch (Exception ex)
            {
                throw new Exception("One or more consensus requests failed", ex);
            }
        }

        public async Task<LlmResponse> RaceAsync(List<LlmRouter> routers, LlmRequest request, CancellationToken cancellationToken = default)
        {
            if (routers == null || !routers.Any())
            {
                throw new ArgumentException("No routers provided for race loop");
            }

            using var cts = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
            var exceptions = new List<Exception>();

            var tasks = routers.Select(async router =>
            {
                try
                {
                    var safeRouter = new LlmRouter(router.Provider, router.ApiKey);
                    var result = await safeRouter.SendRequestAsync(request, cts.Token);
                    return result;
                }
                catch (Exception e)
                {
                    lock (exceptions)
                    {
                        exceptions.Add(e);
                    }
                    throw;
                }
            }).ToList();

            while (tasks.Any())
            {
                var finishedTask = await Task.WhenAny(tasks);
                tasks.Remove(finishedTask);

                if (finishedTask.Status == TaskStatus.RanToCompletion)
                {
                    cts.Cancel();
                    return await finishedTask;
                }
            }

            throw new AggregateException("All race requests failed", exceptions);
        }
    }
}
