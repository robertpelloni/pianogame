using System;
using System.Collections.Generic;
using Harness.LlmApi;

namespace Harness
{
    class Program
    {
        static async System.Threading.Tasks.Task Main(string[] args)
        {
            Console.WriteLine("Ultimate Agentic Coding Harness - C# Edition");

            // 1. Initialize API structures
            var request = new LlmRequest
            {
                Model = "gpt-4o",
                Messages = new List<Message>
                {
                    new Message
                    {
                        Role = "system",
                        Content = "You are a helpful coding harness."
                    }
                }
            };

            Console.WriteLine($"Successfully initialized LLM Request for model: {request.Model}");

            // 2. Initialize Routers
            var openaiRouter = new LlmRouter("openai", "dummy_key");
            var anthropicRouter = new LlmRouter("anthropic", "dummy_key");
            Console.WriteLine($"Successfully initialized LLM Routers");

            // 3. Orchestration
            var orchestrator = new Orchestrator();
            // var routers = new List<LlmRouter> { openaiRouter, anthropicRouter };
            // var consensusResult = await orchestrator.ConsensusAsync(routers, request);
            // var raceResult = await orchestrator.RaceAsync(routers, request);

            Console.WriteLine($"Successfully initialized Multi-Agent Orchestrator");
        }
    }
}
