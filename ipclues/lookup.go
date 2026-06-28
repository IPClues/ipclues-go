package ipclues

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

// Lookup performs an IP intelligence lookup for the given IP address.
// Both IPv4 and IPv6 addresses are supported.
// The response fields populated depend on your plan tier.
//
// Returns ErrInvalidIP if ip is not a valid address string,
// ErrUnauthorized if the API key is invalid,
// ErrRateLimit if the request limit has been exceeded, and
// ErrNotFound if there is no record for the IP.
func (c *Client) Lookup(ctx context.Context, ip string) (*LookupResult, error) {
	if net.ParseIP(ip) == nil {
		return nil, ErrInvalidIP
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/lookup/ip/%s", c.baseURL, ip), nil)
	if err != nil {
		return nil, fmt.Errorf("ipclues: build request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgentHeader())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ipclues: http: %w", err)
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return nil, err
	}

	var result LookupResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("ipclues: decode response: %w", err)
	}

	return &result, nil
}

// apiErrorResponse is the shape of error bodies returned by the IPClues API.
type apiErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// checkResponse inspects an HTTP response and returns an appropriate error
// for any non-2xx status code.
func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))

	var apiErr apiErrorResponse
	if err := json.Unmarshal(body, &apiErr); err != nil || apiErr.Code == "" {
		// Non-JSON or empty body — synthesise from status code.
		apiErr.Code = strings.ToLower(http.StatusText(resp.StatusCode))
		apiErr.Message = string(body)
	}

	e := NewAPIError(resp.StatusCode, apiErr.Code, apiErr.Message)

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return fmt.Errorf("%w", ErrUnauthorized)
	case http.StatusTooManyRequests:
		return fmt.Errorf("%w", ErrRateLimit)
	case http.StatusNotFound:
		return fmt.Errorf("%w", ErrNotFound)
	default:
		return e
	}
}
