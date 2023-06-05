package health

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
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/malishan/home-assignment/utils/logger"
)

func Test_healthStatus(t *testing.T) {

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
	ctx.Request = httptest.NewRequest(http.MethodGet, "/health/v1/status", nil)

	healthProvider = providerMock{}

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
			name: "Resouces Is Down",
			mockFunc: func() {
				resourceHealthStatusMock = func(ctx *gin.Context) []*model.HealthApiResponse {
					return []*model.HealthApiResponse{
						{Resource: "database", Status: "inactive", Message: "error: db ping failed"},
					}
				}
			},
			args:    args{ctx: ctx},
			wantErr: true,
		},
		{
			name: "Successful",
			mockFunc: func() {
				resourceHealthStatusMock = func(ctx *gin.Context) []*model.HealthApiResponse {
					return []*model.HealthApiResponse{
						{Resource: "database", Status: "active", Message: ""},
					}
				}
			},
			args:    args{ctx: ctx},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			healthStatus(tt.args.ctx)
			respBytes, _ := io.ReadAll(w.Body)
			var r Response
			json.Unmarshal(respBytes, &r)
			expected := r.Data[0].Status == "active"
			if expected == tt.wantErr {
				t.Errorf("Error expectation not met, want %v, got %v", tt.wantErr, expected)
			}
		})
	}
}
