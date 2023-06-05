package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/external/database"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

func TestHealthAPIImpl_ResourceHealthStatus(t *testing.T) {

	type fields struct {
		DbService database.DbService
	}

	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockfunc func()
		want     []*model.HealthApiResponse
	}{
		{
			name:   "Resource Failed",
			fields: fields{DbService: dbMock{}},
			args:   args{ctx: &gin.Context{}},
			mockfunc: func() {
				pingStatusMock = func(ctx context.Context) *errors.Error {
					return nil
				}
			},
			want: []*model.HealthApiResponse{{Resource: "database", Status: "active", Message: ""}},
		},
		{
			name:   "Resource Failed",
			fields: fields{DbService: dbMock{}},
			args:   args{ctx: &gin.Context{}},
			mockfunc: func() {
				pingStatusMock = func(ctx context.Context) *errors.Error {
					return &errors.Error{Details: "ping failed"}
				}
			},
			want: []*model.HealthApiResponse{{Resource: "database", Status: "inactive", Message: "error: ping failed"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()
			impl := HealthAPIImpl{
				DbService: tt.fields.DbService,
			}
			if got := impl.ResourceHealthStatus(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HealthAPIImpl.ResourceHealthStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
