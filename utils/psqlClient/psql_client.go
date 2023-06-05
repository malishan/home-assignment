package psqlclient

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type PsqlClient interface {
	Close() error
	Stats() sql.DBStats
	PingContext(ctx context.Context) error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func GetPsqlClient() PsqlClient {
	return client
}

func ClosePsqlClient() error {
	if client == nil {
		return errors.New("client in closed")
	}

	return client.Close()
}
