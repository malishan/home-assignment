package health

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"

	"github.com/malishan/home-assignment/business/health"
)

type Response struct {
	Status bool                      `json:"status" example:"false" enums:"true,false"`
	Data   []model.HealthApiResponse `json:"data"`
}

var (
	resourceHealthStatusMock func(ctx *gin.Context) []*model.HealthApiResponse
)

type providerMock struct {
}

func (impl providerMock) ResourceHealthStatus(ctx *gin.Context) []*model.HealthApiResponse {
	return resourceHealthStatusMock(ctx)
}

func TestInitHealthProvider(t *testing.T) {

	type args struct {
		provider health.HealthAPIProvider
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
			InitHealthProvider(tt.args.provider)
		})
	}
}
