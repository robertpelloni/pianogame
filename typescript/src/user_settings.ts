/**
 * User Settings Engine
 * Manages audio buffer sizing, sample rates, MIDI device preferences, and visual styles.
 */

export interface UserSettings {
  appName: string;
  version: string;
  audioBufferSize: number;
  audioSampleRate: number;
  defaultMidiIn: string;
  defaultMidiOut: string;
  preferredKeyboardSize: number;
  notationStyle: "falling_notes" | "sheet_music" | "hybrid";
  showKeyLabels: boolean;
  autoPlayUnselectedTracks: boolean;
}

export class UserSettingsEngine {
  private static defaultSettings: UserSettings = {
    appName: "PianoGame",
    version: "1.0.0",
    audioBufferSize: 256,
    audioSampleRate: 48000,
    defaultMidiIn: "Generic USB MIDI",
    defaultMidiOut: "Internal Software Synthesizer",
    preferredKeyboardSize: 88,
    notationStyle: "falling_notes",
    showKeyLabels: true,
    autoPlayUnselectedTracks: true,
  };

  private static currentSettings: UserSettings = { ...UserSettingsEngine.defaultSettings };

  public static get(): UserSettings {
    return { ...this.currentSettings };
  }

  public static update(partial: Partial<UserSettings>): UserSettings {
    if (partial.audioBufferSize !== undefined) {
      if (partial.audioBufferSize < 64 || partial.audioBufferSize > 2048) {
        throw new Error("Invalid audioBufferSize: must be between 64 and 2048 samples");
      }
    }
    if (partial.audioSampleRate !== undefined) {
      if (![44100, 48000, 96000].includes(partial.audioSampleRate)) {
        throw new Error("Invalid audioSampleRate: must be 44100, 48000, or 96000 Hz");
      }
    }
    if (partial.preferredKeyboardSize !== undefined) {
      if (![37, 49, 61, 76, 88].includes(partial.preferredKeyboardSize)) {
        throw new Error("Invalid preferredKeyboardSize: must be 37, 49, 61, 76, or 88");
      }
    }

    this.currentSettings = {
      ...this.currentSettings,
      ...partial,
    };
    return { ...this.currentSettings };
  }

  public static reset(): UserSettings {
    this.currentSettings = { ...this.defaultSettings };
    return { ...this.currentSettings };
  }
}
