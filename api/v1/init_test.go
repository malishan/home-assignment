package v1

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/business/home"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

var (
	getActivitiesMock func(ctx *gin.Context) ([]*model.BoredApiResponse, *errors.Error)
)

type providerMock struct {
}

func (impl providerMock) GetActivities(ctx *gin.Context) ([]*model.BoredApiResponse, *errors.Error) {
	return getActivitiesMock(ctx)
}

func TestInitHomeProvider(t *testing.T) {
	type args struct {
		provider home.HomeAPIProvider
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
			InitHomeProvider(tt.args.provider)
		})
	}
}
