using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace Harness.LlmApi
{
    public class LlmRouter
    {
        public string Provider { get; private set; }
        public string ApiKey { get; private set; }
        public string BaseUrl { get; private set; }
        private readonly HttpClient _httpClient;

        public LlmRouter(string provider, string apiKey)
        {
            Provider = provider;
            ApiKey = apiKey;
            BaseUrl = provider.ToLower() == "anthropic"
                ? "https://api.anthropic.com/v1/messages"
                : "https://api.openai.com/v1/chat/completions";

            _httpClient = new HttpClient();
        }

        public async Task<LlmResponse> SendRequestAsync(LlmRequest request)
        {
            var jsonRequest = JsonSerializer.Serialize(request);
            var content = new StringContent(jsonRequest, Encoding.UTF8, "application/json");

            var httpRequest = new HttpRequestMessage(HttpMethod.Post, BaseUrl)
            {
                Content = content
            };

            // Map provider-specific headers
            if (Provider.ToLower() == "anthropic")
            {
                httpRequest.Headers.Add("x-api-key", ApiKey);
                httpRequest.Headers.Add("anthropic-version", "2023-06-01");
            }
            else
            {
                httpRequest.Headers.Add("Authorization", $"Bearer {ApiKey}");
            }

            var response = await _httpClient.SendAsync(httpRequest);
            response.EnsureSuccessStatusCode();

            var responseBody = await response.Content.ReadAsStringAsync();
            var llmResponse = JsonSerializer.Deserialize<LlmResponse>(responseBody);

            return llmResponse ?? new LlmResponse();
        }

        public async Task StreamResponseAsync(LlmRequest request, Action<string> onChunk)
        {
            // Placeholder for SSE parsing logic
            await Task.CompletedTask;
        }

        public string HandleToolCall(ToolCall toolCall)
        {
            // Placeholder for local function routing
            return "{}";
        }
    }
}
