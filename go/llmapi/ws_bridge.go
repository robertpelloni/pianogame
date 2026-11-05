package llmapi

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"sync"
)

// WsBridge handles lightweight native WebSocket connections for remote agents
type WsBridge struct {
	mu      sync.RWMutex
	clients map[string]net.Conn
}

// NewWsBridge creates a new WebSocket bridge instance
func NewWsBridge() *WsBridge {
	return &WsBridge{
		clients: make(map[string]net.Conn),
	}
}

// HandleUpgrade performs WebSocket handshake without third-party dependencies
func (wb *WsBridge) HandleUpgrade(w http.ResponseWriter, r *http.Request) error {
	key := r.Header.Get("Sec-WebSocket-Key")
	if key == "" {
		http.Error(w, "Bad Request: Missing Sec-WebSocket-Key", http.StatusBadRequest)
		return fmt.Errorf("missing key")
	}

	h := sha1.New()
	h.Write([]byte(key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	accept := base64.StdEncoding.EncodeToString(h.Sum(nil))

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return fmt.Errorf("hijack not supported")
	}

	conn, _, err := hijacker.Hijack()
	if err != nil {
		return err
	}

	response := fmt.Sprintf(
		"HTTP/1.1 101 Switching Protocols\r\n"+
			"Upgrade: websocket\r\n"+
			"Connection: Upgrade\r\n"+
			"Sec-WebSocket-Accept: %s\r\n\r\n",
		accept,
	)

	_, err = conn.Write([]byte(response))
	if err != nil {
		conn.Close()
		return err
	}

	clientKey := conn.RemoteAddr().String()
	wb.mu.Lock()
	wb.clients[clientKey] = conn
	wb.mu.Unlock()

	// Read loop to detect disconnect
	go func() {
		defer func() {
			wb.mu.Lock()
			delete(wb.clients, clientKey)
			wb.mu.Unlock()
			conn.Close()
		}()
		buf := make([]byte, 1024)
		for {
			_, err := conn.Read(buf)
			if err != nil {
				break
			}
		}
	}()

	return nil
}

// ClientCount returns the number of active remote connections
func (wb *WsBridge) ClientCount() int {
	wb.mu.RLock()
	defer wb.mu.RUnlock()
	return len(wb.clients)
}
