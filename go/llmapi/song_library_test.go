package llmapi

import (
	"path/filepath"
	"testing"
)

func TestSongLibrary(t *testing.T) {
	if TrimFilename("mozart_sonata.mid") != "mozart_sonata" {
		t.Fatalf("expected mozart_sonata, got %s", TrimFilename("mozart_sonata.mid"))
	}
	if TrimFilename("/path/to/chopin_nocturne.MIDI") != "chopin_nocturne" {
		t.Fatalf("expected chopin_nocturne, got %s", TrimFilename("/path/to/chopin_nocturne.MIDI"))
	}

	testDir := filepath.Join("..", "..", "legacy", "testing")
	songs := ScanSongDirectory(testDir)
	if len(songs) == 0 {
		t.Fatalf("expected at least one song in legacy/testing")
	}

	found := false
	for _, s := range songs {
		if s.Filename == "pitch_bend.mid" {
			found = true
			if s.Title != "pitch_bend" || s.SizeBytes != 2707 {
				t.Fatalf("unexpected pitch_bend metadata: %+v", s)
			}
		}
	}
	if !found {
		t.Fatalf("pitch_bend.mid not found in scanned songs")
	}

	ClearRecentSongs()
	RecordRecentSong(songs[0])
	if len(GetRecentSongs()) != 1 {
		t.Fatalf("expected 1 recent song, got %d", len(GetRecentSongs()))
	}
}
