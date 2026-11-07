/**
 * Track Scoring & Performance Engine
 * Evaluates note accuracy, combo streaks, score multipliers, and grades.
 */

export enum TrackMode {
  ModePlayedAutomatically = 0,
  ModeYouPlay = 1,
  ModePlayedButHidden = 2,
  ModeNotPlayed = 3,
}

export interface PerformanceStats {
  totalNotes: number;
  perfectHits: number;
  goodHits: number;
  missedNotes: number;
  maxCombo: number;
  accuracyPercent: number;
  finalScore: number;
  grade: "S" | "A" | "B" | "C" | "F";
}

export class TrackScoringEngine {
  /**
   * Evaluates performance based on hit counts.
   */
  public static evaluateScore(
    perfectHits: number,
    goodHits: number,
    missedNotes: number,
    maxCombo: number
  ): PerformanceStats {
    const total = perfectHits + goodHits + missedNotes;
    if (total === 0) {
      return {
        totalNotes: 0,
        perfectHits: 0,
        goodHits: 0,
        missedNotes: 0,
        maxCombo: 0,
        accuracyPercent: 100,
        finalScore: 0,
        grade: "S",
      };
    }

    // Accuracy weighted: perfect = 100%, good = 70%, miss = 0%
    const weightedPoints = perfectHits * 1.0 + goodHits * 0.7;
    const accuracyPercent = Math.round((weightedPoints / total) * 1000) / 10;

    // Score calculation with combo bonus
    const baseScore = perfectHits * 1000 + goodHits * 500;
    const comboBonus = maxCombo * 150;
    const finalScore = baseScore + comboBonus;

    let grade: "S" | "A" | "B" | "C" | "F" = "F";
    if (accuracyPercent >= 98.0) {
      grade = "S";
    } else if (accuracyPercent >= 90.0) {
      grade = "A";
    } else if (accuracyPercent >= 80.0) {
      grade = "B";
    } else if (accuracyPercent >= 70.0) {
      grade = "C";
    }

    return {
      totalNotes: total,
      perfectHits,
      goodHits,
      missedNotes,
      maxCombo,
      accuracyPercent,
      finalScore,
      grade,
    };
  }
}
