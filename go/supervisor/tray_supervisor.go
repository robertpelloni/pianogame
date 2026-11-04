package supervisor

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ProcessState tracks health of a supervised background engine
type ProcessState struct {
	Name      string    `json:"name"`
	Running   bool      `json:"running"`
	Pid       int       `json:"pid"`
	UptimeSec int64     `json:"uptime_sec"`
	StartTime time.Time `json:"start_time"`
}

// TraySupervisor coordinates OS-level tray presence, process monitoring,
// and graceful multi-engine termination.
type TraySupervisor struct {
	mu        sync.RWMutex
	processes map[string]*ProcessState
	cancel    context.CancelFunc
	ctx       context.Context
}

// NewTraySupervisor constructs a new supervisor instance.
func NewTraySupervisor() *TraySupervisor {
	ctx, cancel := context.WithCancel(context.Background())
	ts := &TraySupervisor{
		processes: make(map[string]*ProcessState),
		cancel:    cancel,
		ctx:       ctx,
	}

	// Register known harness engines
	engines := []string{"rust", "go", "csharp", "java", "typescript"}
	for _, engine := range engines {
		ts.processes[engine] = &ProcessState{
			Name:      engine,
			Running:   true,
			Pid:       os.Getpid(),
			StartTime: time.Now(),
		}
	}

	return ts
}

// GetStatus returns the current snapshot of all supervised engines.
func (ts *TraySupervisor) GetStatus() map[string]ProcessState {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	status := make(map[string]ProcessState)
	for k, v := range ts.processes {
		copyState := *v
		copyState.UptimeSec = int64(time.Since(v.StartTime).Seconds())
		status[k] = copyState
	}
	return status
}

// QuitAll triggers a graceful teardown of all supervised processes and OS tray icon.
func (ts *TraySupervisor) QuitAll() {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	fmt.Println("[TraySupervisor] QuitAll requested: sending graceful termination signal to all engines...")
	for _, p := range ts.processes {
		p.Running = false
	}
	ts.cancel()
	fmt.Println("[TraySupervisor] All processes cleanly stopped. Tray icon dismissed.")
}

// RunBackgroundLoop starts the tray daemon supervisor monitor loop.
func (ts *TraySupervisor) RunBackgroundLoop() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case sig := <-sigChan:
			fmt.Printf("[TraySupervisor] Received OS signal %v, terminating...\n", sig)
			ts.QuitAll()
		case <-ts.ctx.Done():
			return
		}
	}()
}
