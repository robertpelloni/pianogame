import assert from "assert";
import { TextOverlayEngine } from "./text_overlay.js";

// Test 1: Left-aligned span
const left = TextOverlayEngine.calculateLayout({
  x: 100,
  y: 200,
  alignment: "left",
  fontSize: 16,
  fontFamily: "Trebuchet MS",
  dropShadow: false,
  spans: [{ text: "Hello", color: "#ffffff" }],
});
assert.strictEqual(left.renderX, 100);
assert.strictEqual(left.renderY, 200);
assert.strictEqual(left.estimatedWidth, Math.round(5 * 16 * 0.6));

// Test 2: Centered HUD metric
const centered = TextOverlayEngine.formatHudMetric("Accuracy: ", "98.5%", "#10b981", 500, 300, "center");
assert.strictEqual(centered.alignment, "center");
assert.ok(centered.renderX < 500, "Centered renderX should offset left of anchor");
assert.strictEqual(centered.spans.length, 2);

console.log("TypeScript TextOverlayEngine tests passed cleanly!");
