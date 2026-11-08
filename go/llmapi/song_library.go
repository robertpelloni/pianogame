package llmapi

import (
	"os"
	"path/filepath"
	"strings"
)

// SongMetadata holds file metadata and clean title.
type SongMetadata struct {
	Path      string `json:"path"`
	Filename  string `json:"filename"`
	Title     string `json:"title"`
	SizeBytes int64  `json:"sizeBytes"`
}

var recentSongs []SongMetadata
const maxRecent = 10

// TrimFilename removes directory prefixes and .mid/.midi extensions.
func TrimFilename(pathStr string) string {
	base := filepath.Base(pathStr)
	ext := strings.ToLower(filepath.Ext(base))
	if ext == ".mid" || ext == ".midi" {
		return strings.TrimSuffix(base, filepath.Ext(base))
	}
	return base
}

// ScanSongDirectory searches a directory for MIDI files.
func ScanSongDirectory(dirPath string) []SongMetadata {
	var songs []SongMetadata
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return songs
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			name := entry.Name()
			ext := strings.ToLower(filepath.Ext(name))
			if ext == ".mid" || ext == ".midi" {
				info, err := entry.Info()
				size := int64(0)
				if err == nil {
					size = info.Size()
				}
				songs = append(songs, SongMetadata{
					Path:      filepath.Join(dirPath, name),
					Filename:  name,
					Title:     TrimFilename(name),
					SizeBytes: size,
				})
			}
		}
	}
	return songs
}

// RecordRecentSong adds a song to the recent list with de-duplication.
func RecordRecentSong(song SongMetadata) {
	var updated []SongMetadata
	for _, s := range recentSongs {
		if s.Path != song.Path {
			updated = append(updated, s)
		}
	}
	recentSongs = append([]SongMetadata{song}, updated...)
	if len(recentSongs) > maxRecent {
		recentSongs = recentSongs[:maxRecent]
	}
}

// GetRecentSongs returns recent tracks.
func GetRecentSongs() []SongMetadata {
	return recentSongs
}

// ClearRecentSongs clears recent history.
func ClearRecentSongs() {
	recentSongs = nil
}
