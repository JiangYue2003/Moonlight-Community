package app

import (
	"context"
	"fmt"
	"sync"
)

// Component defines a managed sub-service in merged process.
type Component interface {
	Name() string
	Run(ctx context.Context) error
}

// Run starts all components and returns when:
// - parent ctx is canceled (nil), or
// - any component returns error (first error).
func Run(ctx context.Context, components []Component) error {
	if len(components) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, len(components))
	var wg sync.WaitGroup
	for _, comp := range components {
		c := comp
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := c.Run(ctx); err != nil {
				errCh <- fmt.Errorf("%s: %w", c.Name(), err)
				cancel()
			}
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		select {
		case err := <-errCh:
			return err
		default:
			return nil
		}
	case <-ctx.Done():
		<-done
		select {
		case err := <-errCh:
			return err
		default:
			return nil
		}
	}
}
