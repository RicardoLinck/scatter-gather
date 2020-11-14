package fetch

import (
	"context"
	"time"
)

type PartnerFetcher struct {
	partners map[string]int
}

func NewPartnerFetcher() *PartnerFetcher {
	return &PartnerFetcher{
		partners: map[string]int{
			"test@test.com": 20,
		},
	}
}

func (p *PartnerFetcher) Fetch(ctx context.Context, email string, results chan<- Result) error {
	time.Sleep(time.Second)
	points, ok := p.partners[email]
	if ok {
		results <- Result{Key: "loyalty_points", Value: points}
	}
	return nil
}
