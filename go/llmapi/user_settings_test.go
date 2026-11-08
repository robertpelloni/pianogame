package llmapi

import (
	"testing"
)

func TestUserSettings(t *testing.T) {
	ResetUserSettings()
	defaults := GetUserSettings()
	if defaults.AppName != "PianoGame" || defaults.AudioBufferSize != 256 || defaults.PreferredKeyboardSize != 88 {
		t.Fatalf("unexpected defaults: %+v", defaults)
	}

	updated, err := UpdateUserSettings(512, 44100, 61, "sheet_music")
	if err != nil {
		t.Fatalf("unexpected update error: %v", err)
	}
	if updated.AudioBufferSize != 512 || updated.AudioSampleRate != 44100 || updated.PreferredKeyboardSize != 61 {
		t.Fatalf("unexpected updated settings: %+v", updated)
	}

	// Validation checks
	_, err = UpdateUserSettings(16, 0, 0, "")
	if err == nil {
		t.Fatalf("expected error for small buffer size")
	}

	_, err = UpdateUserSettings(0, 12345, 0, "")
	if err == nil {
		t.Fatalf("expected error for invalid sample rate")
	}
}
