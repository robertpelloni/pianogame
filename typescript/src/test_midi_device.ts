import assert from "assert";
import { MidiDeviceManager } from "./midi_device.js";

// Test 1: List devices
const devices = MidiDeviceManager.listDevices();
assert.strictEqual(devices.inputs.length, 2);
assert.strictEqual(devices.outputs.length, 2);
assert.strictEqual(devices.activeInputId, 0);
assert.strictEqual(devices.activeOutputId, 0);

// Test 2: Select input
const inDev = MidiDeviceManager.selectInput(1);
assert.strictEqual(inDev.id, 1);
assert.strictEqual(MidiDeviceManager.listDevices().activeInputId, 1);

// Test 3: Select output
const outDev = MidiDeviceManager.selectOutput(0);
assert.strictEqual(outDev.id, 0);

// Test 4: Invalid device handling
assert.throws(() => {
  MidiDeviceManager.selectInput(999);
}, /not found/);

console.log("TypeScript MidiDeviceManager tests passed cleanly!");
