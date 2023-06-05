package v1

import (
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/malishan/home-assignment/docs"
)

func TestHomeRoutes(t *testing.T) {
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
			HomeRoutes(tt.args.group)
		})
	}
}

func TestHomeSwaggerRoutes(t *testing.T) {
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
			HomeSwaggerRoutes(tt.args.group)
		})
	}
}
