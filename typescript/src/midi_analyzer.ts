import fs from "fs";

export interface MidiAnalysisReport {
  format: number;
  trackCount: number;
  division: number;
  totalEvents: number;
  noteCount: number;
  pitchBendEvents: number;
  channelsDetected: number[];
  validHeader: boolean;
}

export class MidiAnalyzer {
  /**
   * Parses binary Standard MIDI file buffer according to SMF specification.
   */
  public static analyzeBuffer(buffer: Buffer): MidiAnalysisReport {
    if (buffer.length < 14) {
      throw new Error("Buffer too small for valid MIDI header");
    }

    const header = buffer.subarray(0, 4).toString("ascii");
    if (header !== "MThd") {
      throw new Error(`Invalid MIDI header: expected 'MThd', got '${header}'`);
    }

    const format = buffer.readUInt16BE(8);
    const trackCount = buffer.readUInt16BE(10);
    const division = buffer.readUInt16BE(12);

    let offset = 14;
    let totalEvents = 0;
    let noteCount = 0;
    let pitchBendEvents = 0;
    const channels = new Set<number>();

    while (offset < buffer.length) {
      if (offset + 8 > buffer.length) break;
      const chunkType = buffer.subarray(offset, offset + 4).toString("ascii");
      const chunkSize = buffer.readUInt32BE(offset + 4);
      offset += 8;

      if (chunkType === "MTrk") {
        const trackEnd = Math.min(offset + chunkSize, buffer.length);
        let trackPtr = offset;
        let runningStatus = 0;

        while (trackPtr < trackEnd) {
          // Read variable-length delta time
          let delta = 0;
          let byte = 0;
          do {
            if (trackPtr >= trackEnd) break;
            byte = buffer[trackPtr++];
            delta = (delta << 7) | (byte & 0x7f);
          } while (byte & 0x80);

          if (trackPtr >= trackEnd) break;
          let status = buffer[trackPtr];
          if (status >= 0x80) {
            runningStatus = status;
            trackPtr++;
          } else {
            status = runningStatus;
          }

          totalEvents++;
          const command = status & 0xf0;
          const channel = status & 0x0f;
          if (status >= 0x80 && status < 0xf0) {
            channels.add(channel);
          }

          if (command === 0x90 || command === 0x80) {
            // Note on / Note off
            const note = buffer[trackPtr++];
            const vel = buffer[trackPtr++];
            if (command === 0x90 && vel > 0) {
              noteCount++;
            }
          } else if (command === 0xe0) {
            // Pitch Bend
            trackPtr += 2;
            pitchBendEvents++;
          } else if (command === 0xc0 || command === 0xd0) {
            // Program change or Channel pressure (1 data byte)
            trackPtr += 1;
          } else if (command === 0xa0 || command === 0xb0) {
            // Polyphonic key pressure or Control change (2 data bytes)
            trackPtr += 2;
          } else if (status === 0xff) {
            // Meta event
            if (trackPtr >= trackEnd) break;
            const metaType = buffer[trackPtr++];
            let metaLen = 0;
            let mByte = 0;
            do {
              if (trackPtr >= trackEnd) break;
              mByte = buffer[trackPtr++];
              metaLen = (metaLen << 7) | (mByte & 0x7f);
            } while (mByte & 0x80);
            trackPtr += metaLen;
          } else if (status === 0xf0 || status === 0xf7) {
            // Sysex
            let sysLen = 0;
            let sByte = 0;
            do {
              if (trackPtr >= trackEnd) break;
              sByte = buffer[trackPtr++];
              sysLen = (sysLen << 7) | (sByte & 0x7f);
            } while (sByte & 0x80);
            trackPtr += sysLen;
          }
        }
      }

      offset += chunkSize;
    }

    return {
      format,
      trackCount,
      division,
      totalEvents,
      noteCount,
      pitchBendEvents,
      channelsDetected: Array.from(channels).sort((a, b) => a - b),
      validHeader: true
    };
  }

  /**
   * Analyzes a MIDI file from disk.
   */
  public static analyzeFile(filePath: string): MidiAnalysisReport {
    const buffer = fs.readFileSync(filePath);
    return this.analyzeBuffer(buffer);
  }
}
