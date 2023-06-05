package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var HTTPRequestTiming, PsqlTiming *prometheus.HistogramVec

var HTTPTotalRequestCounter, HTTPResponseStatusCounter, HTTPRequestCounter *prometheus.CounterVec

func IniPrometheustMetrics(ctx context.Context, loader ...int) error {

	for _, load := range loader {

		switch load {
		case HTTPRequestTimingEnable:
			HTTPRequestTiming = promauto.NewHistogramVec(prometheus.HistogramOpts{
				Name:    "http_response_time_seconds",
				Help:    "Duration of HTTP requests.",
				Buckets: prometheus.LinearBuckets(0.1, 0.2, 15),
			}, []string{"path"})

		case PsqlTimingEnable:
			PsqlTiming = promauto.NewHistogramVec(prometheus.HistogramOpts{
				Name:    "psql_timing",
				Help:    "Psql timing",
				Buckets: prometheus.LinearBuckets(0.1, 0.2, 15),
			}, []string{"table", "method"})

		case HTTPTotalRequestCounterEnable:
			HTTPTotalRequestCounter = promauto.NewCounterVec(
				prometheus.CounterOpts{
					Name: "http_requests_total",
					Help: "Number of get requests.",
				},
				[]string{"path", "serviceName"},
			)

		case HTTPResponseStatusCounterEnable:
			HTTPResponseStatusCounter = promauto.NewCounterVec(
				prometheus.CounterOpts{
					Name: "response_status",
					Help: "Status of HTTP response",
				},
				[]string{"status"},
			)

		case HTTPRequestCounterEnable:
			HTTPRequestCounter = promauto.NewCounterVec(
				prometheus.CounterOpts{
					Name: "gin_requests_total",
					Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
				},
				[]string{"status", "host", "path", "method"},
			)
		}
	}

	return nil
}
