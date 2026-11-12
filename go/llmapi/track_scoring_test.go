package llmapi

import (
	"testing"
)

func TestEvaluateScore(t *testing.T) {
	// Test Grade S
	sPerf := EvaluateScore(100, 0, 0, 100)
	if sPerf.Grade != "S" || sPerf.AccuracyPercent != 100.0 {
		t.Fatalf("expected Grade S with 100%%, got %s / %v", sPerf.Grade, sPerf.AccuracyPercent)
	}

	// Test Grade A
	aPerf := EvaluateScore(90, 8, 2, 45)
	if aPerf.Grade != "A" || aPerf.AccuracyPercent != 95.6 {
		t.Fatalf("expected Grade A with 95.6%%, got %s / %v", aPerf.Grade, aPerf.AccuracyPercent)
	}

	// Test Grade F
	fPerf := EvaluateScore(20, 10, 70, 5)
	if fPerf.Grade != "F" {
		t.Fatalf("expected Grade F, got %s", fPerf.Grade)
	}
}
