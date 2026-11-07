package llmapi

import (
	"testing"
)

func TestGetKeyboardLayout(t *testing.T) {
	kb88 := GetKeyboardLayout(88)
	if kb88.TotalKeys != 88 || kb88.WhiteKeyCount != 52 || kb88.BlackKeyCount != 36 {
		t.Fatalf("unexpected 88-key layout: %+v", kb88)
	}
	if kb88.StartNoteName != "A0" || kb88.EndNoteName != "C8" {
		t.Fatalf("unexpected 88-key note names: %s -> %s", kb88.StartNoteName, kb88.EndNoteName)
	}

	kb61 := GetKeyboardLayout(61)
	if kb61.TotalKeys != 61 || kb61.WhiteKeyCount != 36 || kb61.BlackKeyCount != 25 {
		t.Fatalf("unexpected 61-key layout: %+v", kb61)
	}

	if MidiToNoteName(60) != "C4" {
		t.Fatalf("expected 60 to be C4, got %s", MidiToNoteName(60))
	}
}
