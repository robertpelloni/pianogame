package llmapi

import "fmt"

// KeyboardLayout holds geometry and note metadata for a given piano size.
type KeyboardLayout struct {
	Size          int    `json:"size"`
	TotalKeys     int    `json:"totalKeys"`
	WhiteKeyCount int    `json:"whiteKeyCount"`
	BlackKeyCount int    `json:"blackKeyCount"`
	StartMidiNote int    `json:"startMidiNote"`
	EndMidiNote   int    `json:"endMidiNote"`
	StartOctave   int    `json:"startOctave"`
	StartNoteName string `json:"startNoteName"`
	EndNoteName   string `json:"endNoteName"`
}

var noteNames = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

// MidiToNoteName converts a MIDI note number into standard scientific pitch notation.
func MidiToNoteName(midi int) string {
	note := noteNames[midi%12]
	octave := (midi / 12) - 1
	return fmt.Sprintf("%s%d", note, octave)
}

// GetKeyboardLayout returns the keyboard geometry for a given size.
func GetKeyboardLayout(size int) KeyboardLayout {
	switch size {
	case 37:
		return KeyboardLayout{
			Size:          37,
			TotalKeys:     37,
			WhiteKeyCount: 22,
			BlackKeyCount: 15,
			StartMidiNote: 41,
			EndMidiNote:   77,
			StartOctave:   2,
			StartNoteName: "F2",
			EndNoteName:   "F5",
		}
	case 49:
		return KeyboardLayout{
			Size:          49,
			TotalKeys:     49,
			WhiteKeyCount: 29,
			BlackKeyCount: 20,
			StartMidiNote: 36,
			EndMidiNote:   84,
			StartOctave:   2,
			StartNoteName: "C2",
			EndNoteName:   "C6",
		}
	case 61:
		return KeyboardLayout{
			Size:          61,
			TotalKeys:     61,
			WhiteKeyCount: 36,
			BlackKeyCount: 25,
			StartMidiNote: 36,
			EndMidiNote:   96,
			StartOctave:   2,
			StartNoteName: "C2",
			EndNoteName:   "C7",
		}
	case 76:
		return KeyboardLayout{
			Size:          76,
			TotalKeys:     76,
			WhiteKeyCount: 45,
			BlackKeyCount: 31,
			StartMidiNote: 28,
			EndMidiNote:   103,
			StartOctave:   1,
			StartNoteName: "E1",
			EndNoteName:   "G7",
		}
	case 88:
		fallthrough
	default:
		return KeyboardLayout{
			Size:          88,
			TotalKeys:     88,
			WhiteKeyCount: 52,
			BlackKeyCount: 36,
			StartMidiNote: 21,
			EndMidiNote:   108,
			StartOctave:   0,
			StartNoteName: "A0",
			EndNoteName:   "C8",
		}
	}
}
