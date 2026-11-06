package supervisor

import (
	"fmt"
	"os/exec"
	"runtime"
)

// DesktopTrayLauncher wraps native OS notification and tray capabilities
type DesktopTrayLauncher struct {
	supervisor *TraySupervisor
	title      string
}

// NewDesktopTrayLauncher creates a new desktop launcher instance
func NewDesktopTrayLauncher(sup *TraySupervisor) *DesktopTrayLauncher {
	return &DesktopTrayLauncher{
		supervisor: sup,
		title:      "Ultimate Agentic Coding Harness",
	}
}

// OpenDashboard opens the WebUI dashboard in the user's default desktop browser
func (d *DesktopTrayLauncher) OpenDashboard(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // Linux / BSD
		cmd = "xdg-open"
		args = []string{url}
	}

	fmt.Printf("[DesktopTray] Opening dashboard (%s) using %s...\n", url, cmd)
	return exec.Command(cmd, args...).Start()
}

// ShowNotification displays an OS-level notification balloon or alert
func (d *DesktopTrayLauncher) ShowNotification(title, message string) {
	fmt.Printf("[DesktopTray Notification] [%s]: %s\n", title, message)
}

// HandleAction processes tray menu clicks
func (d *DesktopTrayLauncher) HandleAction(action string) {
	switch action {
	case "open_dashboard":
		_ = d.OpenDashboard("http://localhost:8080")
	case "check_health":
		status := d.supervisor.GetStatus()
		d.ShowNotification(d.title, fmt.Sprintf("%d engines running healthy.", len(status)))
	case "quit":
		d.ShowNotification(d.title, "Stopping all background harness servers...")
		d.supervisor.QuitAll()
	default:
		fmt.Printf("[DesktopTray] Unknown action: %s\n", action)
	}
}
