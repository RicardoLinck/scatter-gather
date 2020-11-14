package fetch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// NameFetcher fetches data from Name service
type NameFetcher struct {
	url string
}

func NewNameFetcher(url string) *NameFetcher {
	return &NameFetcher{url: url}
}

// NameServiceResponse represents the response from Name Service
type NameServiceResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (n *NameFetcher) Fetch(ctx context.Context, email string, results chan<- Result) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?email=%s", n.url, email), nil)
	if err != nil {
		return err
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return errors.New(r.Status)
	}

	defer r.Body.Close()
	nsr := NameServiceResponse{}
	err = json.NewDecoder(r.Body).Decode(&nsr)
	if err != nil {
		return err
	}

	results <- Result{Key: "name", Value: nsr}
	return nil
}
