package llmapi

import (
	"fmt"
	"strings"
	"time"
)

// CdpNavigationResult encapsulates browser navigation outcomes
type CdpNavigationResult struct {
	URL        string    `json:"url"`
	StatusCode int       `json:"status_code"`
	Title      string    `json:"title"`
	Timestamp  time.Time `json:"timestamp"`
}

// BrowserCdpClient simulates headless browser CDP session interactions
type BrowserCdpClient struct {
	activeURL string
	pageTitle string
	connected bool
}

// NewBrowserCdpClient constructs a new CDP client instance
func NewBrowserCdpClient() *BrowserCdpClient {
	return &BrowserCdpClient{
		activeURL: "about:blank",
		pageTitle: "Blank Page",
		connected: true,
	}
}

// Navigate simulates navigating a browser context to target URL
func (c *BrowserCdpClient) Navigate(url string) CdpNavigationResult {
	c.activeURL = url
	if strings.Contains(url, "localhost:8080") {
		c.pageTitle = "Ultimate Agentic Coding Harness — Control Surface"
	} else {
		c.pageTitle = fmt.Sprintf("Document at %s", url)
	}

	return CdpNavigationResult{
		URL:        c.activeURL,
		StatusCode: 200,
		Title:      c.pageTitle,
		Timestamp:  time.Now(),
	}
}

// EvaluateScript simulates evaluating JS expression in page context
func (c *BrowserCdpClient) EvaluateScript(expr string) map[string]interface{} {
	if expr == "document.title" {
		return map[string]interface{}{"value": c.pageTitle}
	}
	return map[string]interface{}{"value": "Evaluated: " + expr}
}

// GetStatus returns current connection and active page details
func (c *BrowserCdpClient) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"connected": c.connected,
		"active_url": c.activeURL,
		"page_title": c.pageTitle,
	}
}
