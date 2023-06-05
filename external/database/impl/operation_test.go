package impl

import (
	"context"
	"database/sql"
	ers "errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
	psqlclient "github.com/malishan/home-assignment/utils/psqlClient"
)

func TestDbServiceImpl_PingStatus(t *testing.T) {

	//metrics.IniPrometheustMetrics(context.Background(),
	//	metrics.PsqlTimingEnable,
	//)

	type fields struct {
		DBClient psqlclient.PsqlClient
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockfunc func()
		want     *errors.Error
	}{
		{
			name:     "nil db client",
			fields:   fields{DBClient: nil},
			args:     args{ctx: context.Background()},
			mockfunc: func() {},
			want: &errors.Error{StatusCode: http.StatusInternalServerError,
				Code:    errors.InternalServerErrorCode,
				Message: errors.InternalServerError,
				Details: "pingStatus db - nil client",
			},
		},
		{
			name:   "ping failed",
			fields: fields{DBClient: dbClientMock{}},
			args:   args{ctx: context.Background()},
			mockfunc: func() {
				pingContextMock = func(ctx context.Context) error {
					return errors.New("conn error")
				}
			},
			want: &errors.Error{StatusCode: http.StatusInternalServerError,
				Code:    errors.InternalServerErrorCode,
				Message: errors.InternalServerError,
				Details: "pingStatus db - ping failed-{\"code\":\"\",\"message\":\"conn error\"}",
			},
		},
		{
			name:   "ping successful",
			fields: fields{DBClient: dbClientMock{}},
			args:   args{ctx: context.Background()},
			mockfunc: func() {
				pingContextMock = func(ctx context.Context) error {
					return nil
				}
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &DbServiceImpl{
				DBClient: tt.fields.DBClient,
			}
			tt.mockfunc()
			if got := impl.PingStatus(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbServiceImpl.PingStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbServiceImpl_InsertUserActivityRecord(t *testing.T) {

	//metrics.IniPrometheustMetrics(context.Background(),
	//	metrics.PsqlTimingEnable,
	//)

	type fields struct {
		DBClient psqlclient.PsqlClient
	}

	type args struct {
		ctx    context.Context
		id     string
		userId string
		record []*model.BoredApiResponse
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockfunc func()
		want     *errors.Error
	}{
		{
			name:     "nil db client",
			fields:   fields{DBClient: nil},
			args:     args{ctx: context.Background()},
			mockfunc: func() {},
			want: &errors.Error{StatusCode: http.StatusInternalServerError,
				Code:    errors.InternalServerErrorCode,
				Message: errors.InternalServerError,
				Details: "insertUserActivityRecord db - nil client",
			},
		},
		{
			name:     "invalid args",
			fields:   fields{DBClient: dbClientMock{}},
			args:     args{ctx: context.Background()},
			mockfunc: func() {},
			want: &errors.Error{StatusCode: http.StatusInternalServerError,
				Code:    errors.InternalServerErrorCode,
				Message: errors.InternalServerError,
				Details: "insertUserActivityRecord db - invalid func args",
			},
		},
		{
			name:   "begin transaction failed",
			fields: fields{DBClient: dbClientMock{}},
			args:   args{ctx: context.Background(), id: "ahqerhavqg", record: []*model.BoredApiResponse{{Key: "1234", Activity: "do something"}}},
			mockfunc: func() {
				beginTxMock = func(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
					return nil, errors.New("trx failed")
				}
			},
			want: &errors.Error{StatusCode: http.StatusInternalServerError,
				Code:    errors.InternalServerErrorCode,
				Message: errors.InternalServerError,
				Details: "insertUserActivityRecord db - begin trx failed-{\"code\":\"\",\"message\":\"trx failed\"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &DbServiceImpl{
				DBClient: tt.fields.DBClient,
			}
			tt.mockfunc()
			if got := impl.InsertUserActivityRecord(tt.args.ctx, tt.args.id, tt.args.userId, tt.args.record); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbServiceImpl.InsertUserActivityRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbServiceImpl_FetchUserActivityRecord(t *testing.T) {

	//metrics.IniPrometheustMetrics(context.Background(),
	//	metrics.PsqlTimingEnable,
	//)

	type fields struct {
		DBClient psqlclient.PsqlClient
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockfunc func()
		want     []*model.SchedulerResponse
		want1    *errors.Error
	}{
		{
			name:     "nil db client",
			fields:   fields{DBClient: nil},
			args:     args{ctx: context.Background()},
			mockfunc: func() {},
			want:     nil,
			want1: &errors.Error{StatusCode: http.StatusInternalServerError,
				Code:    errors.InternalServerErrorCode,
				Message: errors.InternalServerError,
				Details: "fetchUserActivityRecord db - nil client",
			},
		},
		{
			name:   "query failed",
			fields: fields{DBClient: dbClientMock{}},
			args:   args{ctx: context.Background()},
			mockfunc: func() {
				queryContextMock = func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
					return nil, ers.New("conn failed")
				}
			},
			want: nil,
			want1: &errors.Error{StatusCode: http.StatusInternalServerError,
				Code:    errors.InternalServerErrorCode,
				Message: errors.InternalServerError,
				Details: "fetchUserActivityRecord db - query failed-conn failed",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &DbServiceImpl{
				DBClient: tt.fields.DBClient,
			}
			tt.mockfunc()
			got, got1 := impl.FetchUserActivityRecord(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbServiceImpl.FetchUserActivityRecord() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DbServiceImpl.FetchUserActivityRecord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
