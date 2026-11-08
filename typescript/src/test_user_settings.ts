import assert from "assert";
import { UserSettingsEngine } from "./user_settings.js";

// Test 1: Defaults
UserSettingsEngine.reset();
const defaults = UserSettingsEngine.get();
assert.strictEqual(defaults.appName, "PianoGame");
assert.strictEqual(defaults.audioBufferSize, 256);
assert.strictEqual(defaults.audioSampleRate, 48000);
assert.strictEqual(defaults.preferredKeyboardSize, 88);

// Test 2: Valid update
const updated = UserSettingsEngine.update({
  audioBufferSize: 512,
  preferredKeyboardSize: 61,
  notationStyle: "hybrid",
});
assert.strictEqual(updated.audioBufferSize, 512);
assert.strictEqual(updated.preferredKeyboardSize, 61);
assert.strictEqual(updated.notationStyle, "hybrid");

// Test 3: Validation checks
assert.throws(() => {
  UserSettingsEngine.update({ audioBufferSize: 16 });
}, /Invalid audioBufferSize/);

assert.throws(() => {
  UserSettingsEngine.update({ audioSampleRate: 12345 });
}, /Invalid audioSampleRate/);

console.log("TypeScript UserSettingsEngine tests passed cleanly!");
