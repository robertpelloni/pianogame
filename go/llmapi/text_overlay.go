package llmapi

import "math"

// TextSpan represents a colored text chunk.
type TextSpan struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}

// TextOverlayOptions configures rendering placement and style.
type TextOverlayOptions struct {
	X          int        `json:"x"`
	Y          int        `json:"y"`
	Alignment  string     `json:"alignment"`
	FontSize   int        `json:"fontSize"`
	FontFamily string     `json:"fontFamily"`
	DropShadow bool       `json:"dropShadow"`
	Spans      []TextSpan `json:"spans"`
}

// TextOverlayLayout includes calculated bounding box and render anchors.
type TextOverlayLayout struct {
	TextOverlayOptions
	RenderX         int `json:"renderX"`
	RenderY         int `json:"renderY"`
	EstimatedWidth  int `json:"estimatedWidth"`
	EstimatedHeight int `json:"estimatedHeight"`
}

// CalculateTextLayout computes origin and dimensions for text spans.
func CalculateTextLayout(opts TextOverlayOptions) TextOverlayLayout {
	totalChars := 0
	for _, span := range opts.Spans {
		totalChars += len(span.Text)
	}

	charWidthFactor := 0.6
	estimatedWidth := int(math.Round(float64(totalChars) * float64(opts.FontSize) * charWidthFactor))
	estimatedHeight := int(math.Round(float64(opts.FontSize) * 1.3))

	renderX := opts.X
	if opts.Alignment == "center" {
		renderX = opts.X - (estimatedWidth / 2)
	} else if opts.Alignment == "right" {
		renderX = opts.X - estimatedWidth
	}

	return TextOverlayLayout{
		TextOverlayOptions: opts,
		RenderX:            renderX,
		RenderY:            opts.Y,
		EstimatedWidth:     estimatedWidth,
		EstimatedHeight:    estimatedHeight,
	}
}

// FormatHudMetric creates a standardized HUD metric.
func FormatHudMetric(label, value, valueColor string, x, y int, alignment string) TextOverlayLayout {
	return CalculateTextLayout(TextOverlayOptions{
		X:          x,
		Y:          y,
		Alignment:  alignment,
		FontSize:   14,
		FontFamily: "Trebuchet MS",
		DropShadow: true,
		Spans: []TextSpan{
			{Text: label, Color: "#94a3b8"},
			{Text: value, Color: valueColor},
		},
	})
}
