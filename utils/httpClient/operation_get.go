package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

// GET is used to make a get request with the provided details
func (c *Client) GET(url string, headers map[string]string, query map[string]string) (*http.Response, error) {
	return c.GETWithTimeout(url, headers, query, 0)
}

// GETWithTimeout is used to make a get request with the provided details
// 0 timeout means default timeout will be used
func (c *Client) GETWithTimeout(url string, headers map[string]string, query map[string]string, timeout time.Duration) (*http.Response, error) {
	return c.GETWithTimeoutAndRetries(url, headers, query, timeout, 0, 0, 0)
}

// GETWithTimeoutAndRetries is used to make a get request with the provided details
// 0 timeout means default timeout will be used
func (c *Client) GETWithTimeoutAndRetries(url string, headers map[string]string, query map[string]string, timeout time.Duration,
	retryCount int, retryWaitTime time.Duration, retryMaxWaitTime time.Duration) (*http.Response, error) {
	// create a request
	request, err := createRequest(http.MethodGet, url, headers, query, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request %w", err)
	}

	// now time to execute with retry and backoff
	return c.doWithTimeoutAndRetries(request, timeout, retryCount, retryWaitTime, retryMaxWaitTime)
}
