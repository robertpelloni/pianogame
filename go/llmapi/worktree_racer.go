package llmapi

import (
	"fmt"
	"sort"
	"time"
)

// CandidateSolution tracks a competing model's worktree execution
type CandidateSolution struct {
	Model        string  `json:"model"`
	Branch       string  `json:"branch"`
	WorktreePath string  `json:"worktree_path"`
	TestPassRate float64 `json:"test_pass_rate"`
	DurationMs   int64   `json:"duration_ms"`
	Status       string  `json:"status"`
}

// WorktreeRaceResult aggregates race outcomes
type WorktreeRaceResult struct {
	RaceID      string              `json:"race_id"`
	Timestamp   time.Time           `json:"timestamp"`
	Candidates  []CandidateSolution `json:"candidates"`
	WinnerModel string              `json:"winner_model"`
	Status      string              `json:"status"`
}

// WorktreeRacer coordinates parallel worktree competition
type WorktreeRacer struct{}

// NewWorktreeRacer creates a racer instance
func NewWorktreeRacer() *WorktreeRacer {
	return &WorktreeRacer{}
}

// ExecuteRace runs a multi-agent race across simulated worktrees
func (w *WorktreeRacer) ExecuteRace(task string) WorktreeRaceResult {
	raceID := fmt.Sprintf("race_%d", time.Now().UnixMilli())

	candidates := []CandidateSolution{
		{
			Model:        "gpt-4o",
			Branch:       fmt.Sprintf("race/%s/gpt4o", raceID),
			WorktreePath: fmt.Sprintf(".worktrees/%s_gpt4o", raceID),
			TestPassRate: 1.0,
			DurationMs:   420,
			Status:       "passed",
		},
		{
			Model:        "claude-3-5-sonnet",
			Branch:       fmt.Sprintf("race/%s/claude", raceID),
			WorktreePath: fmt.Sprintf(".worktrees/%s_claude", raceID),
			TestPassRate: 1.0,
			DurationMs:   380,
			Status:       "passed",
		},
		{
			Model:        "gemini-1-5-pro",
			Branch:       fmt.Sprintf("race/%s/gemini", raceID),
			WorktreePath: fmt.Sprintf(".worktrees/%s_gemini", raceID),
			TestPassRate: 1.0,
			DurationMs:   490,
			Status:       "passed",
		},
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].DurationMs < candidates[j].DurationMs
	})

	winner := "none"
	if len(candidates) > 0 {
		winner = candidates[0].Model
	}

	return WorktreeRaceResult{
		RaceID:      raceID,
		Timestamp:   time.Now(),
		Candidates:  candidates,
		WinnerModel: winner,
		Status:      "completed",
	}
}
