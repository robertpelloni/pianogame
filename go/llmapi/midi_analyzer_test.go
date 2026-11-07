package llmapi

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestAnalyzeMidiFile(t *testing.T) {
	midiPath := filepath.Join("..", "..", "legacy", "testing", "pitch_bend.mid")
	report, err := AnalyzeMidiFile(midiPath)
	if err != nil {
		t.Fatalf("failed to analyze MIDI file: %v", err)
	}

	if !report.ValidHeader {
		t.Fatalf("expected valid header")
	}
	if report.Format != 1 {
		t.Fatalf("expected format 1, got %d", report.Format)
	}
	if report.TrackCount != 3 {
		t.Fatalf("expected 3 tracks, got %d", report.TrackCount)
	}
	if report.Division != 120 {
		t.Fatalf("expected division 120, got %d", report.Division)
	}
	if report.TotalEvents != 844 {
		t.Fatalf("expected 844 total events, got %d", report.TotalEvents)
	}
	if report.PitchBendEvents != 253 {
		t.Fatalf("expected 253 pitch bend events, got %d", report.PitchBendEvents)
	}
	if !reflect.DeepEqual(report.ChannelsDetected, []int{0, 5}) {
		t.Fatalf("expected channels [0, 5], got %v", report.ChannelsDetected)
	}
}
