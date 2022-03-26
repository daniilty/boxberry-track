package client

import "net/http"

// ClientOption - used for client DI.
type ClientOption func(*client)

// WithHTTPClient - set custom http client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *client) {
		c.httpClient = httpClient
	}
}
