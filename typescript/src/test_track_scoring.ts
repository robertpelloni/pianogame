import assert from "assert";
import { TrackScoringEngine } from "./track_scoring.js";

// Test 1: Perfect score (Grade S)
const perfect = TrackScoringEngine.evaluateScore(100, 0, 0, 100);
assert.strictEqual(perfect.totalNotes, 100);
assert.strictEqual(perfect.accuracyPercent, 100.0);
assert.strictEqual(perfect.grade, "S");
assert.strictEqual(perfect.finalScore, 100 * 1000 + 100 * 150);

// Test 2: Mixed score (Grade A)
const mixed = TrackScoringEngine.evaluateScore(90, 8, 2, 45);
assert.strictEqual(mixed.totalNotes, 100);
assert.strictEqual(mixed.accuracyPercent, 95.6);
assert.strictEqual(mixed.grade, "A");

// Test 3: Low score (Grade F)
const low = TrackScoringEngine.evaluateScore(20, 10, 70, 5);
assert.strictEqual(low.grade, "F");

console.log("TypeScript TrackScoringEngine tests passed cleanly!");
