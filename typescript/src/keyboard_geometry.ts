/**
 * Keyboard Geometry & Layout Engine
 * Calculates virtual keyboard key counts, octave offsets, and MIDI note mappings.
 */

export interface KeyboardLayout {
  size: number;
  totalKeys: number;
  whiteKeyCount: number;
  blackKeyCount: number;
  startMidiNote: number;
  endMidiNote: number;
  startOctave: number;
  startNoteName: string;
  endNoteName: string;
}

export class KeyboardGeometryEngine {
  private static noteNames = ["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"];

  public static midiToNoteName(midi: number): string {
    const note = this.noteNames[midi % 12];
    const octave = Math.floor(midi / 12) - 1;
    return `${note}${octave}`;
  }

  /**
   * Resolves keyboard layout specifications for standard piano sizes.
   */
  public static getLayout(size: number): KeyboardLayout {
    switch (size) {
      case 37:
        return {
          size: 37,
          totalKeys: 37,
          whiteKeyCount: 22,
          blackKeyCount: 15,
          startMidiNote: 41, // F2
          endMidiNote: 77,   // F5
          startOctave: 2,
          startNoteName: "F2",
          endNoteName: "F5",
        };
      case 49:
        return {
          size: 49,
          totalKeys: 49,
          whiteKeyCount: 29,
          blackKeyCount: 20,
          startMidiNote: 36, // C2
          endMidiNote: 84,   // C6
          startOctave: 2,
          startNoteName: "C2",
          endNoteName: "C6",
        };
      case 61:
        return {
          size: 61,
          totalKeys: 61,
          whiteKeyCount: 36,
          blackKeyCount: 25,
          startMidiNote: 36, // C2
          endMidiNote: 96,   // C7
          startOctave: 2,
          startNoteName: "C2",
          endNoteName: "C7",
        };
      case 76:
        return {
          size: 76,
          totalKeys: 76,
          whiteKeyCount: 45,
          blackKeyCount: 31,
          startMidiNote: 28, // E1
          endMidiNote: 103,  // G7
          startOctave: 1,
          startNoteName: "E1",
          endNoteName: "G7",
        };
      case 88:
      default:
        return {
          size: 88,
          totalKeys: 88,
          whiteKeyCount: 52,
          blackKeyCount: 36,
          startMidiNote: 21, // A0
          endMidiNote: 108,  // C8
          startOctave: 0,
          startNoteName: "A0",
          endNoteName: "C8",
        };
    }
  }
}
