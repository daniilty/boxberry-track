package client

import (
	"context"
	"net/http"
	"net/url"
)

type Client interface {
	Search(context.Context, string) ([]*SearchResult, error)
	Track(context.Context, string) (*TrackResult, error)
}

type client struct {
	httpClient *http.Client

	baseURL *url.URL
}

func NewClient(baseURL *url.URL, opts ...ClientOption) Client {
	c := &client{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
