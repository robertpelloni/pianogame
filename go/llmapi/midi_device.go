package llmapi

import "errors"

// MidiDevice represents an audio/MIDI interface endpoint.
type MidiDevice struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Connected bool    `json:"connected"`
	LatencyMs float64 `json:"latencyMs,omitempty"`
	Polyphony int     `json:"polyphony,omitempty"`
}

// MidiDeviceInventory holds available input and output devices.
type MidiDeviceInventory struct {
	Inputs         []MidiDevice `json:"inputs"`
	Outputs        []MidiDevice `json:"outputs"`
	ActiveInputID  int          `json:"activeInputId"`
	ActiveOutputID int          `json:"activeOutputId"`
}

var defaultInputs = []MidiDevice{
	{ID: 0, Name: "Generic USB MIDI Keyboard", Type: "hardware", Connected: true, LatencyMs: 1.2},
	{ID: 1, Name: "Virtual MIDI Loopback", Type: "virtual", Connected: true, LatencyMs: 0.3},
}

var defaultOutputs = []MidiDevice{
	{ID: 0, Name: "Internal Software Synthesizer", Type: "software", Connected: true, Polyphony: 128},
	{ID: 1, Name: "External MIDI Out", Type: "hardware", Connected: false, Polyphony: 16},
}

var activeInput = 0
var activeOutput = 0

// ListMidiDevices returns available MIDI input and output ports.
func ListMidiDevices() MidiDeviceInventory {
	return MidiDeviceInventory{
		Inputs:         defaultInputs,
		Outputs:        defaultOutputs,
		ActiveInputID:  activeInput,
		ActiveOutputID: activeOutput,
	}
}

// SelectMidiInput sets the active input device by ID.
func SelectMidiInput(id int) (*MidiDevice, error) {
	for _, dev := range defaultInputs {
		if dev.ID == id {
			activeInput = id
			return &dev, nil
		}
	}
	return nil, errors.New("input device not found")
}

// SelectMidiOutput sets the active output device by ID.
func SelectMidiOutput(id int) (*MidiDevice, error) {
	for _, dev := range defaultOutputs {
		if dev.ID == id {
			activeOutput = id
			return &dev, nil
		}
	}
	return nil, errors.New("output device not found")
}
