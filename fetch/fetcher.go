package fetch

import (
	"context"
	"log"
)

// Result represents the result of a fetcher
type Result struct {
	Key   string
	Value interface{}
}

// Fetcher fetches data for a given email and stores the result in results
type Fetcher interface {
	Fetch(ctx context.Context, email string, results chan<- Result) error
}

// CancellationFetcher provides a handy decorator context cancellation aware
type CancellationFetcher struct {
	Fetcher
}

// Fetch calls the underlying fetcher in a goroutine and selects between the result and the context cancellation
func (c *CancellationFetcher) Fetch(ctx context.Context, email string, results chan<- Result) error {
	r := make(chan Result)
	errCh := make(chan error)
	go func() {
		errCh <- c.Fetcher.Fetch(ctx, email, r)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case result := <-r:
		results <- result
	case err := <-errCh:
		log.Print(err)
	}
	return nil
}

// NewCancellationFetcher decorates the provided fetcher with CancellationFetcher
func NewCancellationFetcher(fetcher Fetcher) *CancellationFetcher {
	return &CancellationFetcher{fetcher}
}
