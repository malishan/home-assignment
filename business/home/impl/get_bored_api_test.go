package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/external/database"
	"github.com/malishan/home-assignment/external/http"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/malishan/home-assignment/utils/logger"
)

func TestHomeAPIImpl_GetActivities(t *testing.T) {

	configs.InitConfigClient(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": "../../" + flags.BaseConfigPath() + "/" + flags.Env(),
			"configNames":      []string{constants.ApplicationConfigFile, constants.DatabaseConfigFile, constants.APIConfigFile, constants.LoggerConfigFile},
			"configType":       "yaml",
		},
	})

	logger.InitFileLogger(logger.FileConfig{})

	type fields struct {
		Config      *model.HomeConfig
		DbService   database.DbService
		HttpService http.HttpService
	}

	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockfunc func()
		want     []*model.BoredApiResponse
		want1    *errors.Error
	}{
		{
			name:   "http error",
			fields: fields{Config: &model.HomeConfig{ActivityCountInResponse: 1}, DbService: dbMock{}, HttpService: httpMock{}},
			args:   args{ctx: &gin.Context{}},
			mockfunc: func() {
				boredApiCallMock = func(ctx *gin.Context) (*model.BoredApiResponse, *errors.Error) {
					return nil, &errors.Error{Code: errors.InternalServerErrorCode}
				}
			},
			want:  nil,
			want1: &errors.Error{Code: errors.InternalServerErrorCode},
		},
		{
			name:   "duplicate keys",
			fields: fields{Config: &model.HomeConfig{ActivityCountInResponse: 2}, DbService: dbMock{}, HttpService: httpMock{}},
			args:   args{ctx: &gin.Context{}},
			mockfunc: func() {
				boredApiCallMock = func(ctx *gin.Context) (*model.BoredApiResponse, *errors.Error) {
					return &model.BoredApiResponse{Key: "1234", Activity: "do something"}, nil
				}
			},
			want:  nil,
			want1: &errors.Error{Code: errors.InvalidDetailsCode, Message: "DUPLICATE RECORD FOUND", Details: "duplicate key-1234"},
		},
		{
			name:   "db error",
			fields: fields{Config: &model.HomeConfig{ActivityCountInResponse: 1}, DbService: dbMock{}, HttpService: httpMock{}},
			args:   args{ctx: &gin.Context{}},
			mockfunc: func() {
				boredApiCallMock = func(ctx *gin.Context) (*model.BoredApiResponse, *errors.Error) {
					return &model.BoredApiResponse{Key: "1234", Activity: "do something"}, nil
				}
				insertUserActivityRecordMock = func(ctx context.Context, id, userId string, record []*model.BoredApiResponse) *errors.Error {
					return &errors.Error{Code: errors.InternalServerErrorCode, Message: errors.InternalServerError, Details: "insertion failed"}
				}
			},
			want:  []*model.BoredApiResponse{{Key: "1234", Activity: "do something"}},
			want1: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := HomeAPIImpl{
				Config:      tt.fields.Config,
				DbService:   tt.fields.DbService,
				HttpService: tt.fields.HttpService,
			}
			tt.mockfunc()
			got, got1 := impl.GetActivities(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HomeAPIImpl.GetActivities() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HomeAPIImpl.GetActivities() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
