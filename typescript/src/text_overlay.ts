/**
 * Rich Text Overlay & Head-Up Display (HUD) Engine
 * Positions and formats multi-color spans and calculates bounding geometry.
 */

export interface TextSpan {
  text: string;
  color: string;
}

export interface TextOverlayOptions {
  x: number;
  y: number;
  alignment: "left" | "center" | "right";
  fontSize: number;
  fontFamily: string;
  dropShadow: boolean;
  spans: TextSpan[];
}

export interface TextOverlayLayout extends TextOverlayOptions {
  renderX: number;
  renderY: number;
  estimatedWidth: number;
  estimatedHeight: number;
}

export class TextOverlayEngine {
  /**
   * Calculates bounding metrics and origin anchors based on alignment.
   */
  public static calculateLayout(options: TextOverlayOptions): TextOverlayLayout {
    // Average character width estimate relative to font size (approx 0.6em for Trebuchet)
    const charWidthFactor = 0.6;
    const fullText = options.spans.map((s) => s.text).join("");
    const totalChars = fullText.length;
    const estimatedWidth = Math.round(totalChars * options.fontSize * charWidthFactor);
    const estimatedHeight = Math.round(options.fontSize * 1.3);

    let renderX = options.x;
    if (options.alignment === "center") {
      renderX = Math.round(options.x - estimatedWidth / 2);
    } else if (options.alignment === "right") {
      renderX = Math.round(options.x - estimatedWidth);
    }

    return {
      ...options,
      renderX,
      renderY: options.y,
      estimatedWidth,
      estimatedHeight,
    };
  }

  /**
   * Helper to construct a standard two-span HUD metric line.
   */
  public static formatHudMetric(
    label: string,
    value: string,
    valueColor: string = "#10b981",
    x: number = 0,
    y: number = 0,
    alignment: "left" | "center" | "right" = "left"
  ): TextOverlayLayout {
    return this.calculateLayout({
      x,
      y,
      alignment,
      fontSize: 14,
      fontFamily: "Trebuchet MS",
      dropShadow: true,
      spans: [
        { text: label, color: "#94a3b8" },
        { text: value, color: valueColor },
      ],
    });
  }
}
