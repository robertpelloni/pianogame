package llmapi

import (
	"encoding/binary"
	"errors"
	"os"
	"sort"
)

// MidiAnalysisReport holds diagnostic metrics of a parsed MIDI file.
type MidiAnalysisReport struct {
	Format           uint16 `json:"format"`
	TrackCount       uint16 `json:"trackCount"`
	Division         uint16 `json:"division"`
	TotalEvents      int    `json:"totalEvents"`
	NoteCount        int    `json:"noteCount"`
	PitchBendEvents  int    `json:"pitchBendEvents"`
	ChannelsDetected []int  `json:"channelsDetected"`
	ValidHeader      bool   `json:"validHeader"`
}

// AnalyzeMidiBuffer parses raw MIDI bytes and returns structural metrics.
func AnalyzeMidiBuffer(data []byte) (*MidiAnalysisReport, error) {
	if len(data) < 14 {
		return nil, errors.New("data too short for MIDI header")
	}

	if string(data[0:4]) != "MThd" {
		return nil, errors.New("invalid MIDI header chunk")
	}

	format := binary.BigEndian.Uint16(data[8:10])
	trackCount := binary.BigEndian.Uint16(data[10:12])
	division := binary.BigEndian.Uint16(data[12:14])

	offset := 14
	totalEvents := 0
	noteCount := 0
	pitchBendEvents := 0
	channelsMap := make(map[int]bool)

	for offset < len(data) {
		if offset+8 > len(data) {
			break
		}
		chunkType := string(data[offset : offset+4])
		chunkSize := int(binary.BigEndian.Uint32(data[offset+4 : offset+8]))
		offset += 8

		if chunkType == "MTrk" {
			trackEnd := offset + chunkSize
			if trackEnd > len(data) {
				trackEnd = len(data)
			}
			trackPtr := offset
			runningStatus := byte(0)

			for trackPtr < trackEnd {
				// Read variable-length quantity for delta time
				b := byte(0)
				for {
					if trackPtr >= trackEnd {
						break
					}
					b = data[trackPtr]
					trackPtr++
					if b < 0x80 {
						break
					}
				}

				if trackPtr >= trackEnd {
					break
				}

				status := data[trackPtr]
				if status >= 0x80 {
					runningStatus = status
					trackPtr++
				} else {
					status = runningStatus
				}

				totalEvents++
				command := status & 0xf0
				channel := int(status & 0x0f)
				if status >= 0x80 && status < 0xf0 {
					channelsMap[channel] = true
				}

				if command == 0x90 || command == 0x80 {
					if trackPtr+1 < trackEnd {
						trackPtr++ // note
						vel := data[trackPtr]
						trackPtr++
						if command == 0x90 && vel > 0 {
							noteCount++
						}
					}
				} else if command == 0xe0 {
					// Pitch bend
					trackPtr += 2
					pitchBendEvents++
				} else if command == 0xc0 || command == 0xd0 {
					trackPtr += 1
				} else if command == 0xa0 || command == 0xb0 {
					trackPtr += 2
				} else if status == 0xff {
					// Meta event
					if trackPtr >= trackEnd {
						break
					}
					trackPtr++ // meta type
					metaLen := 0
					for {
						if trackPtr >= trackEnd {
							break
						}
						mb := data[trackPtr]
						trackPtr++
						metaLen = (metaLen << 7) | int(mb&0x7f)
						if mb < 0x80 {
							break
						}
					}
					trackPtr += metaLen
				} else if status == 0xf0 || status == 0xf7 {
					// Sysex
					sysLen := 0
					for {
						if trackPtr >= trackEnd {
							break
						}
						sb := data[trackPtr]
						trackPtr++
						sysLen = (sysLen << 7) | int(sb&0x7f)
						if sb < 0x80 {
							break
						}
					}
					trackPtr += sysLen
				}
			}
		}

		offset += chunkSize
	}

	var channels []int
	for ch := range channelsMap {
		channels = append(channels, ch)
	}
	sort.Ints(channels)

	return &MidiAnalysisReport{
		Format:           format,
		TrackCount:       trackCount,
		Division:         division,
		TotalEvents:      totalEvents,
		NoteCount:        noteCount,
		PitchBendEvents:  pitchBendEvents,
		ChannelsDetected: channels,
		ValidHeader:      true,
	}, nil
}

// AnalyzeMidiFile reads a MIDI file from disk and parses it.
func AnalyzeMidiFile(filePath string) (*MidiAnalysisReport, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return AnalyzeMidiBuffer(data)
}
