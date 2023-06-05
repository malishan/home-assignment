package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/malishan/home-assignment/utils/logger"
)

func Test_getActivities(t *testing.T) {

	configs.InitConfigClient(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": flags.BaseConfigPath() + "/" + flags.Env(),
			"configNames":      []string{constants.ApplicationConfigFile, constants.DatabaseConfigFile, constants.APIConfigFile, constants.LoggerConfigFile},
			"configType":       "yaml",
		},
	})

	logger.InitFileLogger(logger.FileConfig{})

	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/home/v1/activities", nil)

	apiProvider = providerMock{}

	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name     string
		mockFunc func()
		args     args
		wantErr  bool
	}{
		{
			name: "Provider Failed",
			mockFunc: func() {
				getActivitiesMock = func(ctx *gin.Context) ([]*model.BoredApiResponse, *errors.Error) {
					return nil, &errors.Error{Details: "provider failed error"}
				}
			},
			args:    args{ctx: ctx},
			wantErr: true,
		},
		{
			name: "Provider Successful",
			mockFunc: func() {
				getActivitiesMock = func(ctx *gin.Context) ([]*model.BoredApiResponse, *errors.Error) {
					return []*model.BoredApiResponse{{Key: "1234", Activity: "do something"}}, nil
				}
			},
			args:    args{ctx: ctx},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			getActivities(tt.args.ctx)
			respBytes, _ := io.ReadAll(w.Body)
			var r model.APIResponse
			json.Unmarshal(respBytes, &r)
			if r.Status == tt.wantErr {
				t.Errorf("Error expectation not met, want %v, got %v", tt.wantErr, r.Status)
			}
		})
	}
}
