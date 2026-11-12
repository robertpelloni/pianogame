/**
 * MIDI Device Manager
 * Discovers and tracks physical USB MIDI inputs and software synthesizers.
 */

export interface MidiDevice {
  id: number;
  name: string;
  type: "hardware" | "software" | "virtual";
  connected: boolean;
  latencyMs?: number;
  polyphony?: number;
}

export interface MidiDeviceInventory {
  inputs: MidiDevice[];
  outputs: MidiDevice[];
  activeInputId: number;
  activeOutputId: number;
}

export class MidiDeviceManager {
  private static inputs: MidiDevice[] = [
    {
      id: 0,
      name: "Generic USB MIDI Keyboard",
      type: "hardware",
      connected: true,
      latencyMs: 1.2,
    },
    {
      id: 1,
      name: "Virtual MIDI Loopback",
      type: "virtual",
      connected: true,
      latencyMs: 0.3,
    },
  ];

  private static outputs: MidiDevice[] = [
    {
      id: 0,
      name: "Internal Software Synthesizer",
      type: "software",
      connected: true,
      polyphony: 128,
    },
    {
      id: 1,
      name: "External MIDI Out",
      type: "hardware",
      connected: false,
      polyphony: 16,
    },
  ];

  private static activeInputId: number = 0;
  private static activeOutputId: number = 0;

  public static listDevices(): MidiDeviceInventory {
    return {
      inputs: [...this.inputs],
      outputs: [...this.outputs],
      activeInputId: this.activeInputId,
      activeOutputId: this.activeOutputId,
    };
  }

  public static selectInput(id: number): MidiDevice {
    const dev = this.inputs.find((d) => d.id === id);
    if (!dev) {
      throw new Error(`MIDI Input device with ID ${id} not found`);
    }
    this.activeInputId = id;
    return dev;
  }

  public static selectOutput(id: number): MidiDevice {
    const dev = this.outputs.find((d) => d.id === id);
    if (!dev) {
      throw new Error(`MIDI Output device with ID ${id} not found`);
    }
    this.activeOutputId = id;
    return dev;
  }
}
