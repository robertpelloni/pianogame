package llmapi

import "math"

// TrackMode represents the playback mode for a track.
type TrackMode int

const (
	ModePlayedAutomatically TrackMode = 0
	ModeYouPlay             TrackMode = 1
	ModePlayedButHidden     TrackMode = 2
	ModeNotPlayed           TrackMode = 3
)

// PerformanceStats holds accuracy, combo, score, and grade metrics.
type PerformanceStats struct {
	TotalNotes      int     `json:"totalNotes"`
	PerfectHits     int     `json:"perfectHits"`
	GoodHits        int     `json:"goodHits"`
	MissedNotes     int     `json:"missedNotes"`
	MaxCombo        int     `json:"maxCombo"`
	AccuracyPercent float64 `json:"accuracyPercent"`
	FinalScore      int     `json:"finalScore"`
	Grade           string  `json:"grade"`
}

// EvaluateScore calculates accuracy, combo bonus, and performance grade.
func EvaluateScore(perfectHits, goodHits, missedNotes, maxCombo int) PerformanceStats {
	total := perfectHits + goodHits + missedNotes
	if total == 0 {
		return PerformanceStats{
			TotalNotes:      0,
			PerfectHits:     0,
			GoodHits:        0,
			MissedNotes:     0,
			MaxCombo:        0,
			AccuracyPercent: 100.0,
			FinalScore:      0,
			Grade:           "S",
		}
	}

	weighted := float64(perfectHits)*1.0 + float64(goodHits)*0.7
	rawAccuracy := (weighted / float64(total)) * 100.0
	accuracy := math.Round(rawAccuracy*10) / 10

	baseScore := perfectHits*1000 + goodHits*500
	comboBonus := maxCombo * 150
	finalScore := baseScore + comboBonus

	grade := "F"
	if accuracy >= 98.0 {
		grade = "S"
	} else if accuracy >= 90.0 {
		grade = "A"
	} else if accuracy >= 80.0 {
		grade = "B"
	} else if accuracy >= 70.0 {
		grade = "C"
	}

	return PerformanceStats{
		TotalNotes:      total,
		PerfectHits:     perfectHits,
		GoodHits:        goodHits,
		MissedNotes:     missedNotes,
		MaxCombo:        maxCombo,
		AccuracyPercent: accuracy,
		FinalScore:      finalScore,
		Grade:           grade,
	}
}
