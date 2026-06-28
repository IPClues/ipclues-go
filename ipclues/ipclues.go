// Package main provides a Go client for the IPClues IP intelligence API.
// It supports all tiers of the IPClues data platform: GeoI3-Lite,
// GeoI3-Country, and future NetI3-Org.
//
// The API key can be supplied via WithAPIKey or the IPCLUES_API_KEY
// environment variable. If neither is set, requests return ErrUnauthorized.
//
// Basic usage:
//
//	ctx := context.Background()
//	client := ipclues.New(ipclues.WithAPIKey("ipclues_sk_..."))
//	result, err := client.Lookup(ctx, "1.1.1.1")
package ipclues

import (
	"net/http"
	"os"
)

const version = "0.1.0"

// Client is the IPClues API client.
// Create one with New() and reuse it across your application;
// it is safe for concurrent use.
type Client struct {
	apiKey     string
	baseURL    string
	userAgent  string
	httpClient *http.Client
}

// New creates a new Client with the given options.
//
// The API key is resolved in this order:
//  1. WithAPIKey option
//  2. IPCLUES_API_KEY environment variable
//  3. Empty — requests will return ErrUnauthorized
//
// If no WithBaseURL option is provided, DefaultBaseURL is used.
// If no WithTimeout option is provided, DefaultTimeout is used.
func New(opts ...Option) *Client {
	c := &Client{
		apiKey:  os.Getenv("IPCLUES_API_KEY"),
		baseURL: DefaultBaseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// userAgentHeader returns the User-Agent header value for requests.
func (c *Client) userAgentHeader() string {
	base := "ipclues-go/" + version
	if c.userAgent != "" {
		return base + " " + c.userAgent
	}
	return base
}
