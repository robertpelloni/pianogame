import assert from "assert";
import { KeyboardGeometryEngine } from "./keyboard_geometry.js";

// Test 1: 88-key standard layout
const kb88 = KeyboardGeometryEngine.getLayout(88);
assert.strictEqual(kb88.totalKeys, 88);
assert.strictEqual(kb88.whiteKeyCount, 52);
assert.strictEqual(kb88.blackKeyCount, 36);
assert.strictEqual(kb88.startNoteName, "A0");
assert.strictEqual(kb88.endNoteName, "C8");
assert.strictEqual(kb88.startMidiNote, 21);
assert.strictEqual(kb88.endMidiNote, 108);

// Test 2: 61-key portable layout
const kb61 = KeyboardGeometryEngine.getLayout(61);
assert.strictEqual(kb61.totalKeys, 61);
assert.strictEqual(kb61.whiteKeyCount, 36);
assert.strictEqual(kb61.blackKeyCount, 25);
assert.strictEqual(kb61.startNoteName, "C2");
assert.strictEqual(kb61.endNoteName, "C7");

// Test 3: Note naming utility
assert.strictEqual(KeyboardGeometryEngine.midiToNoteName(60), "C4"); // Middle C
assert.strictEqual(KeyboardGeometryEngine.midiToNoteName(69), "A4"); // Concert Pitch 440Hz

console.log("TypeScript KeyboardGeometryEngine tests passed cleanly!");
