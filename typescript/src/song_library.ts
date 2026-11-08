import fs from "fs";
import path from "path";

export interface SongMetadata {
  path: string;
  filename: string;
  title: string;
  sizeBytes: number;
}

export class SongLibraryEngine {
  private static recentSongs: SongMetadata[] = [];
  private static readonly maxRecentCount = 10;

  /**
   * Trims file path and extensions to yield a clean display title, matching legacy TrimFilename.
   */
  public static trimFilename(filePath: string): string {
    const base = path.basename(filePath);
    return base.replace(/\.(midi|mid)$/i, "");
  }

  /**
   * Scans a directory recursively or flat for MIDI files.
   */
  public static scanDirectory(dirPath: string): SongMetadata[] {
    const results: SongMetadata[] = [];
    if (!fs.existsSync(dirPath)) {
      return results;
    }

    const files = fs.readdirSync(dirPath);
    for (const f of files) {
      const fullPath = path.join(dirPath, f);
      try {
        const stat = fs.statSync(fullPath);
        if (stat.isFile() && /\.(mid|midi)$/i.test(f)) {
          results.push({
            path: fullPath,
            filename: f,
            title: this.trimFilename(f),
            sizeBytes: stat.size,
          });
        }
      } catch {}
    }

    return results;
  }

  /**
   * Records a song to recent playback history.
   */
  public static recordRecent(song: SongMetadata): void {
    this.recentSongs = this.recentSongs.filter((s) => s.path !== song.path);
    this.recentSongs.unshift(song);
    if (this.recentSongs.length > this.maxRecentCount) {
      this.recentSongs.pop();
    }
  }

  public static getRecent(): SongMetadata[] {
    return [...this.recentSongs];
  }

  public static clearRecent(): void {
    this.recentSongs = [];
  }
}
