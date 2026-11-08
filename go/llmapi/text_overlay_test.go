package llmapi

import (
	"testing"
)

func TestTextOverlay(t *testing.T) {
	opts := TextOverlayOptions{
		X:          100,
		Y:          200,
		Alignment:  "left",
		FontSize:   16,
		FontFamily: "Trebuchet MS",
		DropShadow: false,
		Spans:      []TextSpan{{Text: "Hello", Color: "#ffffff"}},
	}

	layout := CalculateTextLayout(opts)
	if layout.RenderX != 100 || layout.RenderY != 200 {
		t.Fatalf("unexpected origin: %d, %d", layout.RenderX, layout.RenderY)
	}

	hud := FormatHudMetric("Streak: ", "42", "#10b981", 640, 480, "center")
	if hud.Alignment != "center" {
		t.Fatalf("expected center alignment, got %s", hud.Alignment)
	}
	if hud.RenderX >= 640 {
		t.Fatalf("expected renderX to be left of center anchor, got %d", hud.RenderX)
	}
}
