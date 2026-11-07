import assert from "assert";
import path from "path";
import { MidiAnalyzer } from "./midi_analyzer.js";

const testMidiPath = path.resolve(__dirname, "../../legacy/testing/pitch_bend.mid");
const report = MidiAnalyzer.analyzeFile(testMidiPath);

console.log("Analyzed pitch_bend.mid:", report);

assert.strictEqual(report.validHeader, true);
assert.strictEqual(report.format, 1);
assert.strictEqual(report.trackCount, 3);
assert.strictEqual(report.division, 120);
assert.strictEqual(report.totalEvents, 844);
assert.strictEqual(report.noteCount, 11);
assert.strictEqual(report.pitchBendEvents, 253);
assert.deepStrictEqual(report.channelsDetected, [0, 5]);

console.log("TypeScript MidiAnalyzer test passed cleanly!");
