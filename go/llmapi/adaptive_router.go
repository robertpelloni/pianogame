package llmapi

import (
	"sort"
)

// ProviderMetrics represents runtime performance of an LLM provider
type ProviderMetrics struct {
	Provider        string  `json:"provider"`
	Model           string  `json:"model"`
	LatencyMs       float64 `json:"latency_ms"`
	CostPer1kTokens float64 `json:"cost_per_1k_tokens"`
	Healthy         bool    `json:"healthy"`
}

// AdaptiveRouteResolution holds the chosen endpoint
type AdaptiveRouteResolution struct {
	ChosenProvider string  `json:"chosen_provider"`
	ChosenModel    string  `json:"chosen_model"`
	LatencyMs      float64 `json:"latency_ms"`
	Cost           float64 `json:"cost"`
	Reason         string  `json:"reason"`
}

// AdaptiveRouter implements dynamic routing based on cost and latency
type AdaptiveRouter struct {
	providers []ProviderMetrics
}

// NewAdaptiveRouter initializes provider metrics
func NewAdaptiveRouter() *AdaptiveRouter {
	return &AdaptiveRouter{
		providers: []ProviderMetrics{
			{Provider: "openai", Model: "gpt-4o", LatencyMs: 220.0, CostPer1kTokens: 0.005, Healthy: true},
			{Provider: "anthropic", Model: "claude-3-5-sonnet", LatencyMs: 195.0, CostPer1kTokens: 0.003, Healthy: true},
			{Provider: "gemini", Model: "gemini-1-5-flash", LatencyMs: 140.0, CostPer1kTokens: 0.001, Healthy: true},
			{Provider: "ollama", Model: "llama3.2", LatencyMs: 85.0, CostPer1kTokens: 0.000, Healthy: true},
		},
	}
}

// Resolve returns the optimal provider for given strategy
func (r *AdaptiveRouter) Resolve(strategy string) AdaptiveRouteResolution {
	candidates := make([]ProviderMetrics, 0)
	for _, p := range r.providers {
		if p.Healthy {
			candidates = append(candidates, p)
		}
	}

	if strategy == "fastest" {
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].LatencyMs < candidates[j].LatencyMs
		})
		return AdaptiveRouteResolution{
			ChosenProvider: candidates[0].Provider,
			ChosenModel:    candidates[0].Model,
			LatencyMs:      candidates[0].LatencyMs,
			Cost:           candidates[0].CostPer1kTokens,
			Reason:         "Selected fastest endpoint",
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		scoreI := candidates[i].LatencyMs*0.5 + candidates[i].CostPer1kTokens*10000.0*0.5
		scoreJ := candidates[j].LatencyMs*0.5 + candidates[j].CostPer1kTokens*10000.0*0.5
		return scoreI < scoreJ
	})

	return AdaptiveRouteResolution{
		ChosenProvider: candidates[0].Provider,
		ChosenModel:    candidates[0].Model,
		LatencyMs:      candidates[0].LatencyMs,
		Cost:           candidates[0].CostPer1kTokens,
		Reason:         "Balanced cost and latency score",
	}
}

// ListProviders returns registered provider statuses
func (r *AdaptiveRouter) ListProviders() []ProviderMetrics {
	return r.providers
}
