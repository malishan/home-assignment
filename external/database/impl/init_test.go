package impl

import (
	"context"
	"database/sql"
	"github.com/malishan/home-assignment/utils/metrics"
	"reflect"
	"testing"

	"github.com/malishan/home-assignment/external/database"
	psqlclient "github.com/malishan/home-assignment/utils/psqlClient"
)

var (
	closeMock           func() error
	statsMock           func() sql.DBStats
	pingContextMock     func(ctx context.Context) error
	prepareContextMock  func(ctx context.Context, query string) (*sql.Stmt, error)
	beginTxMock         func(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	queryContextMock    func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	queryRowContextMock func(ctx context.Context, query string, args ...interface{}) *sql.Row
	execContextMock     func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
)

type dbClientMock struct {
}

func (impl dbClientMock) Close() error {
	return closeMock()
}
func (impl dbClientMock) Stats() sql.DBStats {
	return statsMock()
}
func (impl dbClientMock) PingContext(ctx context.Context) error {
	return pingContextMock(ctx)
}
func (impl dbClientMock) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return prepareContextMock(ctx, query)
}
func (impl dbClientMock) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return beginTxMock(ctx, opts)
}
func (impl dbClientMock) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return queryContextMock(ctx, query, args...)
}
func (impl dbClientMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return queryRowContextMock(ctx, query, args...)
}
func (impl dbClientMock) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return execContextMock(ctx, query, args...)
}

func TestGetDBServiceInstance(t *testing.T) {

	metrics.IniPrometheustMetrics(context.Background(),
		metrics.PsqlTimingEnable,
	)

	type args struct {
		dbClient psqlclient.PsqlClient
	}

	tests := []struct {
		name string
		args args
		want database.DbService
	}{
		{
			name: "initialise db service",
			args: args{dbClient: dbClientMock{}},
			want: &DbServiceImpl{DBClient: dbClientMock{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDBServiceInstance(tt.args.dbClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBServiceInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
