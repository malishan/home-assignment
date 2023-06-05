package psqlclient

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var client *sql.DB

func InitPsqlClient(config *PsqlClientConfig) error {

	err := getPsqlClientOptions(config)
	if err != nil {
		return err
	}

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.UserName, config.Password, config.DbName)

	// Create connection pool
	client, err = sql.Open(PsqlDriverName, psqlConn)
	if err != nil {
		return err
	}

	client.SetMaxIdleConns(config.MaxIdleConnections)
	client.SetMaxOpenConns(config.MaxOpenConnections)
	client.SetConnMaxLifetime(config.MaxLifetime)
	client.SetConnMaxIdleTime(config.MaxIdletime)

	if err = client.PingContext(context.Background()); err != nil {
		return err
	}

	return nil
}

func getPsqlClientOptions(config *PsqlClientConfig) error {

	if config == nil {
		return errors.New("psql client config not found")
	}

	if config.Host == "" {
		return errors.New("psql client host not found")
	}

	if config.DbName == "" {
		return errors.New("psql client name not found")
	}

	if config.UserName == "" {
		return errors.New("psql client username not found")
	}

	if config.Port == 0 {
		config.Port = DefaultPsqlPort
	}

	if config.MaxOpenConnections <= 0 {
		config.MaxOpenConnections = DefaultMaxOpenConnections
	}

	if config.MaxIdleConnections <= 0 {
		config.MaxIdleConnections = DefaultMaxIdleConnections
	}

	if config.MaxLifetime <= 0 {
		config.MaxLifetime = DefaultMaxLifetime
	}

	if config.MaxIdletime <= 0 {
		config.MaxIdletime = DefaultMaxIdletime
	}

	return nil
}

func GetClient() *sql.DB {
	return client
}
