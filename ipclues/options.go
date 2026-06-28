package ipclues

import (
	"fmt"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the production IPClues API endpoint.
	DefaultBaseURL = "https://api.ipclues.com/v1"
	// DefaultTimeout is the default HTTP request timeout.
	DefaultTimeout = 30 * time.Second
)

// Option is a functional option for configuring a Client.
// Pass options to New() to override defaults.
type Option func(*Client)

// WithAPIKey sets the API key used to authenticate requests.
// Takes precedence over the IPCLUES_API_KEY environment variable.
func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithBaseURL overrides the API base URL.
// Useful for staging environments or on-premise deployments.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		if url != "" {
			c.baseURL = url
		}
	}
}

// WithTimeout sets the HTTP request timeout.
// Panics if d is zero or negative.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		if d <= 0 {
			panic(fmt.Sprintf("ipclues: WithTimeout: duration must be positive, got %v", d))
		}
		c.httpClient.Timeout = d
	}
}

// WithUserAgent appends a custom string to the default User-Agent header.
// The final header takes the form: "ipclues-go/x.y.z <custom>".
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// WithHTTPClient replaces the default HTTP client.
// Use this to inject custom transports, proxies, or test mocks.
// The provided client must not be nil.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}
