import assert from "assert";
import path from "path";
import { SongLibraryEngine } from "./song_library.js";

// Test 1: Trim filename
assert.strictEqual(SongLibraryEngine.trimFilename("/songs/bach_fugue.mid"), "bach_fugue");
assert.strictEqual(SongLibraryEngine.trimFilename("beethoven_moonlight.midi"), "beethoven_moonlight");
assert.strictEqual(SongLibraryEngine.trimFilename("custom_track"), "custom_track");

// Test 2: Scan library directory
const testDir = path.resolve(__dirname, "../../legacy/testing");
const songs = SongLibraryEngine.scanDirectory(testDir);
assert.ok(songs.length >= 1, "Should find at least pitch_bend.mid");
const pitchBend = songs.find((s) => s.filename === "pitch_bend.mid");
assert.ok(pitchBend, "Should find pitch_bend.mid in legacy/testing");
assert.strictEqual(pitchBend!.title, "pitch_bend");
assert.strictEqual(pitchBend!.sizeBytes, 2707);

// Test 3: Recent song tracking
SongLibraryEngine.clearRecent();
SongLibraryEngine.recordRecent(pitchBend!);
assert.strictEqual(SongLibraryEngine.getRecent().length, 1);
assert.strictEqual(SongLibraryEngine.getRecent()[0].title, "pitch_bend");

console.log("TypeScript SongLibraryEngine tests passed cleanly!");
