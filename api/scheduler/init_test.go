package scheduler

import (
	"context"
	"testing"

	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/malishan/home-assignment/utils/logger"

	"github.com/malishan/home-assignment/business/scheduler"
)

var (
	pollActivityOperationMock func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error)
)

type providerMock struct {
}

func (impl providerMock) PollActivityOperation(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
	return pollActivityOperationMock(ctx)
}

func TestInitScheduler(t *testing.T) {
	type args struct {
		provider scheduler.ScheduleProvider
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "initialise provider",
			args: args{provider: providerMock{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitScheduler(tt.args.provider)
		})
	}
}

func TestStartCronJob(t *testing.T) {

	configs.InitConfigClient(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": "../../" + flags.BaseConfigPath() + "/" + flags.Env(),
			"configNames":      []string{constants.ApplicationConfigFile, constants.DatabaseConfigFile, constants.APIConfigFile, constants.LoggerConfigFile},
			"configType":       "yaml",
		},
	})

	logger.InitFileLogger(logger.FileConfig{EncodeLogsAsJson: true, FileLoggingEnabled: true, Directory: "../../Logs", Filename: "app.log", MaxSize: 100, Compress: true})

	schedulerHandler = providerMock{}

	tests := []struct {
		name     string
		mockFunc func()
		wantErr  bool
	}{
		//{
		//	name: "Provider Failed",
		//	mockFunc: func() {
		//		pollActivityOperationMock = func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
		//			return nil, &errors.Error{Details: "polling failed"}
		//		}
		//	},
		//	wantErr: true,
		//},
		{
			name: "Provider Successful",
			mockFunc: func() {
				pollActivityOperationMock = func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
					return []*model.SchedulerResponse{
						{Count: 1, Key: "7091374"},
					}, nil
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartCronJob(); (err != nil) != tt.wantErr {
				t.Errorf("StartCronJob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
