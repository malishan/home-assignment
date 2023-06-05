package impl

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/flags"
	httpclient "github.com/malishan/home-assignment/utils/httpClient"
	"github.com/malishan/home-assignment/utils/logger"
	"golang.org/x/time/rate"
)

func TestHttpServiceImpl_BoredApiCall(t *testing.T) {

	configs.InitConfigClient(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": "../../../" + flags.BaseConfigPath() + "/" + flags.Env(),
			"configNames":      []string{constants.ApplicationConfigFile, constants.DatabaseConfigFile, constants.APIConfigFile, constants.LoggerConfigFile},
			"configType":       "yaml",
		},
	})
	logger.InitFileLogger(logger.FileConfig{})

	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/home/v1/activities", nil)

	type fields struct {
		HttpClient  httpclient.HTTPClient
		Ratelimiter *rate.Limiter
	}

	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		funcMock func()
		want     *model.BoredApiResponse
		want1    *errors.Error
	}{
		{
			name:   "config error",
			fields: fields{HttpClient: httpClientMock{}, Ratelimiter: rate.NewLimiter(rate.Every(1*time.Second), 10)},
			args:   args{ctx: ctx},
			funcMock: func() {
				getResponseMock = func(ctx context.Context, url string, headers, queryParam map[string]string, retrier *httpclient.Retrier) (*http.Response, error) {
					return nil, errors.New("conn error")
				}
			},
			want: nil,
			want1: &errors.Error{
				StatusCode: http.StatusInternalServerError,
				Code:       errors.InternalServerErrorCode,
				Message:    errors.InternalServerError,
				Details:    "boredApiCall: http req failed - {\"code\":\"\",\"message\":\"conn error\"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &HttpServiceImpl{
				HttpClient:  tt.fields.HttpClient,
				Ratelimiter: tt.fields.Ratelimiter,
			}
			tt.funcMock()
			got, got1 := impl.BoredApiCall(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpServiceImpl.BoredApiCall() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HttpServiceImpl.BoredApiCall() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
