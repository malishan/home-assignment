package impl

import (
	"time"

	"github.com/malishan/home-assignment/external/http"
	httpclient "github.com/malishan/home-assignment/utils/httpClient"
	"golang.org/x/time/rate"
)

type HttpServiceImpl struct {
	HttpClient  httpclient.HTTPClient
	Ratelimiter *rate.Limiter
}

// GetHttpServiceInstance : service handler for http operations
func GetHttpServiceInstance(rateLimitCount int, rateLimitDuration time.Duration, httpClient httpclient.HTTPClient) http.HttpService {
	return &HttpServiceImpl{
		HttpClient:  httpClient,
		Ratelimiter: rate.NewLimiter(rate.Every(rateLimitDuration), rateLimitCount),
	}
}
