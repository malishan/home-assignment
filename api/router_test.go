package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/malishan/home-assignment/docs"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/malishan/home-assignment/utils/logger"
)

func TestGetRouter(t *testing.T) {

	configs.InitConfigClient(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": "../" + flags.BaseConfigPath() + "/" + flags.Env(),
			"configNames":      []string{constants.ApplicationConfigFile, constants.DatabaseConfigFile, constants.APIConfigFile, constants.LoggerConfigFile},
			"configType":       "yaml",
		},
	})

	logger.InitFileLogger(logger.FileConfig{EncodeLogsAsJson: true, FileLoggingEnabled: true, Directory: "../../Logs", Filename: "app.log", MaxSize: 100, Compress: true})

	type args struct {
		middlewares []gin.HandlerFunc
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Router Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRouter(tt.args.middlewares...)
		})
	}
}
