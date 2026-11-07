package supervisor

// WindowConfig encapsulates desktop webview window dimensions
type WindowConfig struct {
	Title       string `json:"title"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Resizable   bool   `json:"resizable"`
	TrayDocked  bool   `json:"tray_docked"`
}

// DesktopContainer handles webview configuration and lifecycle
type DesktopContainer struct {
	Window WindowConfig `json:"window"`
	URL    string       `json:"url"`
}

// NewDesktopContainer initializes container defaults
func NewDesktopContainer() *DesktopContainer {
	return &DesktopContainer{
		Window: WindowConfig{
			Title:      "Ultimate Agentic Coding Harness",
			Width:      1280,
			Height:     850,
			Resizable:  true,
			TrayDocked: true,
		},
		URL: "http://localhost:8080",
	}
}

// GetInfo returns a descriptive container summary
func (d *DesktopContainer) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"title":       d.Window.Title,
		"dimensions":  []int{d.Window.Width, d.Window.Height},
		"url":         d.URL,
		"tray_docked": d.Window.TrayDocked,
	}
}
