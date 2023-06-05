package health

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthRoutes(t *testing.T) {
	type args struct {
		group *gin.RouterGroup
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "initialize router",
			args: args{group: &gin.Default().RouterGroup},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HealthRoutes(tt.args.group)
		})
	}
}
