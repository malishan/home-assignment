package httpclient

import (
	"context"
	"net"
	"net/http"
	"net/http/cookiejar"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/publicsuffix"
)

// HttpConfig is the set of configurable parameters for the http client
type HttpConfig struct {
	// ConnectTimeout is the maximum amount of time a dial will wait for
	// connect to complete. If Deadline is also set, it may fail
	// earlier.
	ConnectTimeout time.Duration `json:"connectTimeout"`
	// KeepAliveDuration specifies the interval between keep-alive
	// probes for an active network connection.
	KeepAliveDuration time.Duration `json:"KeepAliveDuration"`
	// MaxIdleConnections controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConnections int `json:"maxIdleConnections"`
	// IdleConnectionTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	IdleConnectionTimeout time.Duration `json:"idleConnectionTimeout"`
	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake.
	TLSHandshakeTimeout time.Duration `json:"tlsHandshakeTimeout"`
	// ExpectContinueTimeout specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers.
	ExpectContinueTimeout time.Duration `json:"expectContinueTimeout"`
	// Timeout specifies a time limit for requests made by this
	// Client.
	Timeout time.Duration `json:"timeout"`
}

// Metric is the information used to tracking the performance
type Metric struct {
	Status          int    `json:"status"`
	Message         string `json:"message"`
	LatencyInMillis int64  `json:"latency"`
}

// Metrics provides the basic information for status and latency
type Metrics func(ctx context.Context, name string, m Metric)

// Client is the http client
type Client struct {
	httpClient *http.Client
	om         sync.Once
	m          Metrics
}

type HTTPClient interface {
	GetResponse(ctx context.Context, url string, headers map[string]string, queryParam map[string]string, retrier *Retrier) (*http.Response, error)
}

type HTTPClientProvider struct {
}

var enableMetrics bool
var client *Client

// InitHTTPClient is used to initialise the http client
func InitHTTPClient(m Metrics, config HttpConfig) error {

	var err error

	client, err = configureBasicHTTPClient(&config)
	if err != nil {
		return err
	}

	if enableMetrics {
		client = client.withMetrics(m)
	}

	return nil
}

func InitHTTPClientMetrics(enable bool) {
	enableMetrics = enable
}

func GetHttpClient() HTTPClient {
	return &HTTPClientProvider{}
}

func configureBasicHTTPClient(config *HttpConfig) (*Client, error) {

	cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	dialer := &net.Dialer{
		Timeout:   config.ConnectTimeout,
		KeepAlive: config.KeepAliveDuration,
	}

	client := Client{
		httpClient: &http.Client{
			Jar: cookieJar,
			Transport: &http.Transport{
				Proxy:                 http.ProxyFromEnvironment,
				DialContext:           dialer.DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          config.MaxIdleConnections,
				IdleConnTimeout:       config.IdleConnectionTimeout,
				TLSHandshakeTimeout:   config.TLSHandshakeTimeout,
				ExpectContinueTimeout: config.ExpectContinueTimeout,
				MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
			},
			// global timeout value for all requests
			Timeout: config.Timeout,
		},
	}

	return &client, nil
}

// WithMetrics is used to provide the metrics instance for the http client created
func (c *Client) withMetrics(m Metrics) *Client {
	if m != nil {
		c.om.Do(func() {
			c.m = m
		})
	}
	return c
}

func (c *Client) metricLatencyAndStatusCode(request *http.Request, message string, start time.Time, statusCode int) {
	if c.m != nil {
		c.m(request.Context(), request.URL.Path, Metric{Status: statusCode, Message: message, LatencyInMillis: time.Since(start).Milliseconds()})
	}
}
