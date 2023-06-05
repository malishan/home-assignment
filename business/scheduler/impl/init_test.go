package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/malishan/home-assignment/external/database"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

var (
	pingStatusMock               func(ctx context.Context) *errors.Error
	insertUserActivityRecordMock func(ctx context.Context, id, userId string, record []*model.BoredApiResponse) *errors.Error
	fetchUserActivityRecordMock  func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error)
)

type dbMock struct {
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

func TestNewSchedulerService(t *testing.T) {
	type args struct {
		config *model.SchedulerConfig
		db     database.DbService
	}
	tests := []struct {
		name string
		args args
		want *SchedulerImpl
	}{
		{
			name: "initialise scheduler",
			args: args{config: &model.SchedulerConfig{}, db: dbMock{}},
			want: &SchedulerImpl{Config: &model.SchedulerConfig{}, DbService: dbMock{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSchedulerService(tt.args.config, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchedulerService() = %v, want %v", got, tt.want)
			}
		})
	}
}
