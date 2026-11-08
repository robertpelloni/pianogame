# MIDI Hardware Device Enumeration Protocol (v1.0.0)

## Overview
Connects legacy C++ device discovery (`DeviceTile.h`, `MidiComm.h`) with the modern Agentic Coding Harness. Standardizes the detection, capability inspection, and lifecycle tracking of physical USB MIDI keyboards, hardware synthesizers, and virtual software loopback devices.

## Device Schema
```json
{
  "inputs": [
    {
      "id": 0,
      "name": "Generic USB MIDI Keyboard",
      "type": "hardware",
      "connected": true,
      "latencyMs": 1.2
    },
    {
      "id": 1,
      "name": "Virtual MIDI Loopback",
      "type": "virtual",
      "connected": true,
      "latencyMs": 0.3
    }
  ],
  "outputs": [
    {
      "id": 0,
      "name": "Internal Software Synthesizer (FluidSynth / CoreAudio)",
      "type": "software",
      "connected": true,
      "polyphony": 128
    },
    {
      "id": 1,
      "name": "External Hardware MIDI Out",
      "type": "hardware",
      "connected": false,
      "polyphony": 16
    }
  ]
}
```

## Protocol Operations
* `ListDevices()`: Returns available input and output MIDI interfaces.
* `SelectInputDevice(id)`: Connects and opens the designated input port.
* `SelectOutputDevice(id)`: Directs synthesized playback to the designated output port.
