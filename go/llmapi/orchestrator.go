package llmapi

import (
	"context"
	"errors"
	"sync"
)

type Orchestrator struct{}

// Consensus Loop: Dispatches requests to all provided routers concurrently,
// waits for all of them to resolve via WaitGroup, and returns a unified array of responses.
func (o *Orchestrator) Consensus(routers []*LlmRouter, request *LlmRequest) ([]*LlmResponse, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]*LlmResponse, 0, len(routers))
	errs := make([]error, 0)

	for _, router := range routers {
		wg.Add(1)

		// Capture loop variables for goroutine
		r := router
		req := request

		go func() {
			defer wg.Done()

			// In a real application, we would deep clone the request if it were mutated.
			resp, err := r.SendRequest(req)

			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errs = append(errs, err)
			} else {
				results = append(results, resp)
			}
		}()
	}

	wg.Wait()

	if len(results) == 0 && len(errs) > 0 {
		return nil, errors.New("all consensus requests failed")
	}

	return results, nil
}

// Race Loop: Dispatches requests to all routers concurrently,
// and returns the very first successful response it receives.
func (o *Orchestrator) Race(routers []*LlmRouter, request *LlmRequest) (*LlmResponse, error) {
	if len(routers) == 0 {
		return nil, errors.New("no routers provided for race loop")
	}

	// We use a channel to receive the first valid response
	resultChan := make(chan *LlmResponse, len(routers))
	errChan := make(chan error, len(routers))

	// In Go, context cancellation is used to abort the other requests once one succeeds.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, router := range routers {
		r := router
		req := request

		go func() {
			// A true implementation would pass `ctx` into `SendRequest` to abort HTTP calls mid-flight.
			// For this schema abstraction, we mock the behavior.
			resp, err := r.SendRequest(req)

			select {
			case <-ctx.Done():
				return // Already finished by another goroutine
			default:
				if err != nil {
					errChan <- err
				} else {
					resultChan <- resp
				}
			}
		}()
	}

	// Wait for the first valid result or until all return errors
	errs := 0
	for {
		select {
		case res := <-resultChan:
			// Cancel other goroutines immediately
			cancel()
			return res, nil
		case <-errChan:
			errs++
			if errs == len(routers) {
				return nil, errors.New("all race requests failed")
			}
		}
	}
}
