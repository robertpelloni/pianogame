package llmapi

import (
	"context"
	"errors"
	"sync"
)

type Orchestrator struct{}

func (o *Orchestrator) Consensus(ctx context.Context, routers []*LlmRouter, request *LlmRequest) ([]*LlmResponse, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]*LlmResponse, 0, len(routers))
	errs := make([]error, 0)

	for _, router := range routers {
		wg.Add(1)

		r := router
		req := request

		go func() {
			defer wg.Done()

			resp, err := r.SendRequest(ctx, req)

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

func (o *Orchestrator) Race(ctx context.Context, routers []*LlmRouter, request *LlmRequest) (*LlmResponse, error) {
	if len(routers) == 0 {
		return nil, errors.New("no routers provided for race loop")
	}

	resultChan := make(chan *LlmResponse, len(routers))
	errChan := make(chan error, len(routers))

	raceCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, router := range routers {
		r := router
		req := request

		go func() {
			resp, err := r.SendRequest(raceCtx, req)

			select {
			case <-raceCtx.Done():
				return
			default:
				if err != nil {
					errChan <- err
				} else {
					resultChan <- resp
				}
			}
		}()
	}

	errs := 0
	for {
		select {
		case res := <-resultChan:
			cancel()
			return res, nil
		case <-errChan:
			errs++
			if errs == len(routers) {
				return nil, errors.New("all race requests failed")
			}
		case <-ctx.Done():
		    return nil, ctx.Err()
		}
	}
}
