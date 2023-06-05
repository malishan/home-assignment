package impl

import (
	"context"
	netHttp "net/http"
	"reflect"
	"testing"
	"time"

	"github.com/malishan/home-assignment/external/http"
	httpclient "github.com/malishan/home-assignment/utils/httpClient"
	"github.com/malishan/home-assignment/utils/metrics"
	"golang.org/x/time/rate"
)

var (
	getResponseMock func(ctx context.Context, url string, headers map[string]string, queryParam map[string]string, retrier *httpclient.Retrier) (*netHttp.Response, error)
)

type httpClientMock struct {
}

func (impl httpClientMock) GetResponse(ctx context.Context, url string, headers map[string]string, queryParam map[string]string, retrier *httpclient.Retrier) (*netHttp.Response, error) {
	return getResponseMock(ctx, url, headers, queryParam, retrier)
}

func TestGetHttpServiceInstance(t *testing.T) {

	metrics.IniPrometheustMetrics(context.Background(),
		metrics.HTTPTotalRequestCounterEnable,
		metrics.HTTPResponseStatusCounterEnable,
		metrics.HTTPRequestCounterEnable,
		metrics.HTTPRequestTimingEnable,
	)

	type args struct {
		rateLimitCount    int
		rateLimitDuration time.Duration
		httpClient        httpclient.HTTPClient
	}

	tests := []struct {
		name string
		args args
		want http.HttpService
	}{
		{
			name: "initialise http service",
			args: args{rateLimitCount: 10, rateLimitDuration: 1 * time.Second, httpClient: httpClientMock{}},
			want: &HttpServiceImpl{HttpClient: httpClientMock{}, Ratelimiter: rate.NewLimiter(rate.Every(1*time.Second), 10)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHttpServiceInstance(tt.args.rateLimitCount, tt.args.rateLimitDuration, tt.args.httpClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHttpServiceInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
