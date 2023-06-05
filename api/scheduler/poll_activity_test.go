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
)

func TestPollActivityOperation(t *testing.T) {

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
	}{
		{
			name: "Panic Handler",
			mockFunc: func() {
				pollActivityOperationMock = func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
					return nil, &errors.Error{}
				}
			},
		},
		{
			name: "Provider Failed",
			mockFunc: func() {
				pollActivityOperationMock = func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
					return nil, &errors.Error{Details: "db operation failed"}
				}
			},
		},
		{
			name: "Provider Successful",
			mockFunc: func() {
				pollActivityOperationMock = func(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
					return []*model.SchedulerResponse{{Count: 1, Key: "6543"}}, nil
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			PollActivityOperation()
		})
	}
}
