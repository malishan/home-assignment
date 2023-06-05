package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func createRequest(method, url string, headers map[string]string, query map[string]string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	q := request.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	request.URL.RawQuery = q.Encode()
	return request, err
}

func shouldRetry(attempt, retryCount int, response *http.Response, err error) bool {
	// check with attempt count
	if attempt >= retryCount {
		return false
	}
	// connect error
	if err != nil {
		return true
	}
	// server error
	if response.StatusCode == 0 || response.StatusCode >= 500 {
		return true
	}
	return false
}

func getWaitDuration(attempt int, retryWaitTime time.Duration, retryMaxWaitTime time.Duration) time.Duration {
	waitDuration := retryWaitTime * time.Duration(attempt+1)
	if waitDuration > retryMaxWaitTime {
		return retryMaxWaitTime
	}
	return waitDuration
}

func (c *Client) doWithTimeoutAndRetries(request *http.Request, timeout time.Duration, retryCount int, retryWaitTime time.Duration,
	retryMaxWaitTime time.Duration) (*http.Response, error) {

	// try to set the timeout if it is valid
	if timeout > 0 {
		ctx, _ := context.WithTimeout(request.Context(), timeout)
		// calling cancel is important otherwise it will lead to context leaks
		//defer cancel()
		request = request.WithContext(ctx)
	}

	ctx := request.Context()

	for attempt := 0; attempt <= retryCount; attempt++ {

		// start the timer
		start := time.Now()

		// do request
		response, err := c.httpClient.Do(request)

		if response == nil {
			c.metricLatencyAndStatusCode(request, fmt.Sprintf("failed count-%v", attempt), start, http.StatusInternalServerError)
		} else {
			c.metricLatencyAndStatusCode(request, fmt.Sprintf("count-%v", attempt), start, response.StatusCode)
		}

		if ctx.Err() != nil {
			return nil, err
		}

		// see if retry is required
		needsRetry := shouldRetry(attempt, retryCount, response, err)
		if !needsRetry {
			return response, err
		}

		// now that retry is required
		// calculate the wait duration
		waitDuration := getWaitDuration(attempt, retryWaitTime, retryMaxWaitTime)

		// now wait for this much duration or till context deadline
		select {
		case <-time.After(waitDuration):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return nil, errors.New("unable to process request")
}
