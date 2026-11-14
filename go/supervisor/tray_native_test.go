package supervisor

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// The tray's Quit and Check-health items must work against the real Rust
// hub contract (GET /health, POST /shutdown from bobmani 0.2.15/0.2.16).
// These tests pin that contract without needing a desktop session.

func withURLs(t *testing.T, health, shutdown string) {
	t.Helper()
	oldH, oldS := rustHealthURL, rustShutdownURL
	rustHealthURL, rustShutdownURL = health, shutdown
	t.Cleanup(func() { rustHealthURL, rustShutdownURL = oldH, oldS })
}

func TestRustHealthyAndShutdown(t *testing.T) {
	shutdownHit := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/health" && r.Method == http.MethodGet:
			w.WriteHeader(http.StatusOK)
		case r.URL.Path == "/shutdown" && r.Method == http.MethodPost:
			shutdownHit = true
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()
	withURLs(t, srv.URL+"/health", srv.URL+"/shutdown")

	if !RustHealthy() {
		t.Fatal("RustHealthy should be true against a 200 /health")
	}
	if err := ShutdownRustServer(); err != nil {
		t.Fatalf("ShutdownRustServer: %v", err)
	}
	if !shutdownHit {
		t.Fatal("shutdown endpoint was not hit")
	}
}

func TestRustUnreachable(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	url := srv.URL
	srv.Close() // nothing listening now
	withURLs(t, url+"/health", url+"/shutdown")

	if RustHealthy() {
		t.Fatal("RustHealthy should be false when the hub is down")
	}
	if err := ShutdownRustServer(); err == nil {
		t.Fatal("ShutdownRustServer should error when the hub is down (tray logs it and continues)")
	}
}
