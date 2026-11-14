package supervisor

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"time"

	"fyne.io/systray"
)

//go:embed icon.ico
var trayIcon []byte

const dashboardURL = "http://127.0.0.1:3000/"
const rustHealthURL = "http://127.0.0.1:3000/health"
const rustShutdownURL = "http://127.0.0.1:3000/shutdown"

// ShutdownRustServer asks the Rust hub to stop gracefully (POST /shutdown,
// added in bobmani 0.2.15). Best-effort: the server may not be running.
func ShutdownRustServer() error {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Post(rustShutdownURL, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("rust shutdown: HTTP %d", resp.StatusCode)
	}
	return nil
}

// RustHealthy reports whether the Rust hub answers /health.
func RustHealthy() bool {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(rustHealthURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// RunNativeTray shows a real Windows system-tray icon and blocks until
// Quit. Menu: Open dashboard / Check health / Quit servers (Rust hub via
// /shutdown, then this supervisor). This is the native icon the dashboard
// previously had to report missing.
func RunNativeTray(sup *TraySupervisor) {
	launcher := NewDesktopTrayLauncher(sup)
	onReady := func() {
		systray.SetIcon(trayIcon)
		systray.SetTitle("bobmani")
		systray.SetTooltip("bobmani harness supervisor")
		mOpen := systray.AddMenuItem("Open dashboard", "Open the bobmani dashboard in your browser")
		mHealth := systray.AddMenuItem("Check health", "Check Rust hub and supervised engines")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit servers", "Stop the Rust hub and this supervisor, then remove the tray icon")
		go func() {
			for {
				select {
				case <-mOpen.ClickedCh:
					_ = launcher.OpenDashboard(dashboardURL)
				case <-mHealth.ClickedCh:
					engines := len(sup.GetStatus())
					if RustHealthy() {
						systray.SetTooltip(fmt.Sprintf("bobmani: Rust hub healthy, %d engines supervised", engines))
					} else {
						systray.SetTooltip(fmt.Sprintf("bobmani: Rust hub NOT reachable, %d engines supervised", engines))
					}
				case <-mQuit.ClickedCh:
					if err := ShutdownRustServer(); err != nil {
						fmt.Printf("[Tray] Rust shutdown request failed (may already be down): %v\n", err)
					}
					sup.QuitAll()
					systray.Quit()
					return
				}
			}
		}()
	}
	systray.Run(onReady, func() { fmt.Println("[Tray] icon removed") })
}
