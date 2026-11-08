package llmapi

import (
	"testing"
)

func TestMidiDeviceManager(t *testing.T) {
	devices := ListMidiDevices()
	if len(devices.Inputs) != 2 || len(devices.Outputs) != 2 {
		t.Fatalf("unexpected device list: %+v", devices)
	}

	inDev, err := SelectMidiInput(1)
	if err != nil || inDev.ID != 1 {
		t.Fatalf("failed to select input device 1: %v", err)
	}

	outDev, err := SelectMidiOutput(0)
	if err != nil || outDev.ID != 0 {
		t.Fatalf("failed to select output device 0: %v", err)
	}

	_, err = SelectMidiInput(999)
	if err == nil {
		t.Fatalf("expected error for non-existent input device")
	}
}
