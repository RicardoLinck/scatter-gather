package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RicardoLinck/scatter-gather/fetch"
	"github.com/RicardoLinck/scatter-gather/nameservice"
	"golang.org/x/sync/errgroup"
)

func main() {
	go nameservice.StartServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	fs := configureFetchers(ctx)
	results := make(chan fetch.Result, len(fs))
	for _, f := range fs {
		f := f
		g.Go(func() error { return f.Fetch(ctx, "test@test.com", results) })
	}

	err := g.Wait()
	close(results)

	for result := range results {
		r := result
		fmt.Println(r)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func configureFetchers(ctx context.Context) []fetch.Fetcher {
	return []fetch.Fetcher{
		fetch.NewCancellationFetcher(fetch.NewNameFetcher("http://localhost:8787")),
		fetch.NewCancellationFetcher(fetch.NewPartnerFetcher()),
	}
}
