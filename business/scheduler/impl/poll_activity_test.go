package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/malishan/home-assignment/external/database"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/malishan/home-assignment/utils/logger"
)

func TestSchedulerImpl_PollActivityOperation(t *testing.T) {

	configs.InitConfigClient(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": "../../" + flags.BaseConfigPath() + "/" + flags.Env(),
			"configNames":      []string{constants.ApplicationConfigFile, constants.DatabaseConfigFile, constants.APIConfigFile, constants.LoggerConfigFile},
			"configType":       "yaml",
		},
	})

	logger.InitFileLogger(logger.FileConfig{})

	type fields struct {
		Config    *model.SchedulerConfig
		DbService database.DbService
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockfunc func()
		want     []*model.SchedulerResponse
		want1    *errors.Error
	}{
		{
			name:   "db error",
			fields: fields{Config: &model.SchedulerConfig{}, DbService: dbMock{}},
			args:   args{ctx: context.Background()},
			mockfunc: func() {
				fetchUserActivityRecordMock = func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
					return nil, &errors.Error{Code: errors.InternalServerErrorCode, Message: errors.InternalServerError, Details: "fetch failed"}
				}

			},
			want:  nil,
			want1: &errors.Error{Code: errors.InternalServerErrorCode, Message: errors.InternalServerError, Details: "fetch failed"},
		},
		{
			name:   "success",
			fields: fields{Config: &model.SchedulerConfig{}, DbService: dbMock{}},
			args:   args{ctx: context.Background()},
			mockfunc: func() {
				fetchUserActivityRecordMock = func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
					return []*model.SchedulerResponse{{Count: 1, Key: "1234"}}, nil
				}

			},
			want:  []*model.SchedulerResponse{{Count: 1, Key: "1234"}},
			want1: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := SchedulerImpl{
				Config:    tt.fields.Config,
				DbService: tt.fields.DbService,
			}
			tt.mockfunc()
			got, got1 := impl.PollActivityOperation(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchedulerImpl.PollActivityOperation() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SchedulerImpl.PollActivityOperation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
