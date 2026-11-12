package llmapi

import "errors"

// UserSettings represents the user configuration state.
type UserSettings struct {
	AppName                  string `json:"appName"`
	Version                  string `json:"version"`
	AudioBufferSize          int    `json:"audioBufferSize"`
	AudioSampleRate          int    `json:"audioSampleRate"`
	DefaultMidiIn            string `json:"defaultMidiIn"`
	DefaultMidiOut           string `json:"defaultMidiOut"`
	PreferredKeyboardSize    int    `json:"preferredKeyboardSize"`
	NotationStyle            string `json:"notationStyle"`
	ShowKeyLabels            bool   `json:"showKeyLabels"`
	AutoPlayUnselectedTracks bool   `json:"autoPlayUnselectedTracks"`
}

var defaultSettings = UserSettings{
	AppName:                  "PianoGame",
	Version:                  "1.0.0",
	AudioBufferSize:          256,
	AudioSampleRate:          48000,
	DefaultMidiIn:            "Generic USB MIDI",
	DefaultMidiOut:           "Internal Software Synthesizer",
	PreferredKeyboardSize:    88,
	NotationStyle:            "falling_notes",
	ShowKeyLabels:            true,
	AutoPlayUnselectedTracks: true,
}

var currentSettings = defaultSettings

// GetUserSettings returns current settings.
func GetUserSettings() UserSettings {
	return currentSettings
}

// UpdateUserSettings validates and updates configuration settings.
func UpdateUserSettings(bufSize, sampleRate, kbSize int, notationStyle string) (*UserSettings, error) {
	if bufSize != 0 {
		if bufSize < 64 || bufSize > 2048 {
			return nil, errors.New("invalid buffer size: must be 64-2048")
		}
		currentSettings.AudioBufferSize = bufSize
	}

	if sampleRate != 0 {
		if sampleRate != 44100 && sampleRate != 48000 && sampleRate != 96000 {
			return nil, errors.New("invalid sample rate: must be 44100, 48000, or 96000")
		}
		currentSettings.AudioSampleRate = sampleRate
	}

	if kbSize != 0 {
		if kbSize != 37 && kbSize != 49 && kbSize != 61 && kbSize != 76 && kbSize != 88 {
			return nil, errors.New("invalid keyboard size")
		}
		currentSettings.PreferredKeyboardSize = kbSize
	}

	if notationStyle != "" {
		currentSettings.NotationStyle = notationStyle
	}

	return &currentSettings, nil
}

// ResetUserSettings resets settings to factory defaults.
func ResetUserSettings() UserSettings {
	currentSettings = defaultSettings
	return currentSettings
}
