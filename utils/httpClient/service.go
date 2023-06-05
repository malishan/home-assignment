package httpclient

import (
	"context"
	"net/http"
	"time"
)

type Retrier struct {
	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
}

func (httpClientProvider *HTTPClientProvider) GetResponse(ctx context.Context, url string, headers map[string]string, queryParam map[string]string,
	retrier *Retrier) (*http.Response, error) {

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	deadline, ok := ctx.Deadline()
	if !ok {
		return client.GET(url, headers, queryParam)
	}

	if retrier == nil {
		return client.GETWithTimeout(url, headers, queryParam, deadline.Sub(time.Now()))
	}

	return client.GETWithTimeoutAndRetries(url, headers, queryParam, deadline.Sub(time.Now()), retrier.RetryCount, retrier.RetryWaitTime, retrier.RetryMaxWaitTime)
}
