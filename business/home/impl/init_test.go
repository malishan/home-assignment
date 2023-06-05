package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/external/database"
	"github.com/malishan/home-assignment/external/http"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

var (
	pingStatusMock               func(ctx context.Context) *errors.Error
	insertUserActivityRecordMock func(ctx context.Context, id, userId string, record []*model.BoredApiResponse) *errors.Error
	fetchUserActivityRecordMock  func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error)
)

var (
	boredApiCallMock func(ctx *gin.Context) (*model.BoredApiResponse, *errors.Error)
)

type dbMock struct {
}

type httpMock struct {
}

func (impl dbMock) PingStatus(ctx context.Context) *errors.Error {
	return pingStatusMock(ctx)
}
func (impl dbMock) InsertUserActivityRecord(ctx context.Context, id, userId string, record []*model.BoredApiResponse) *errors.Error {
	return insertUserActivityRecordMock(ctx, id, userId, record)
}
func (impl dbMock) FetchUserActivityRecord(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
	return fetchUserActivityRecordMock(ctx)
}

func (impl httpMock) BoredApiCall(ctx *gin.Context) (*model.BoredApiResponse, *errors.Error) {
	return boredApiCallMock(ctx)
}

func TestNewHomeAPIService(t *testing.T) {
	type args struct {
		config      *model.HomeConfig
		db          database.DbService
		httpService http.HttpService
	}
	tests := []struct {
		name string
		args args
		want *HomeAPIImpl
	}{
		{
			name: "initialise home provider",
			args: args{config: &model.HomeConfig{}, db: dbMock{}, httpService: httpMock{}},
			want: &HomeAPIImpl{Config: &model.HomeConfig{}, DbService: dbMock{}, HttpService: httpMock{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHomeAPIService(tt.args.config, tt.args.db, tt.args.httpService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHomeAPIService() = %v, want %v", got, tt.want)
			}
		})
	}
}
