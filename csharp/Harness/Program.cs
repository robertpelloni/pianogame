using System;
using System.Collections.Generic;
using Harness.LlmApi;

namespace Harness
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Ultimate Agentic Coding Harness - C# Edition");

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

            var router = new LlmRouter("openai", "dummy_key");
            Console.WriteLine($"Successfully initialized LLM Router for provider: {router.Provider}");
        }
    }
}
