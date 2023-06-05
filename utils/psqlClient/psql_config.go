package psqlclient

import "time"

const PsqlDriverName = "postgres"
const DefaultPsqlPort = 5432
const DefaultMaxOpenConnections int = 0
const DefaultMaxIdleConnections int = 2
const DefaultMaxLifetime time.Duration = 1 * time.Hour
const DefaultMaxIdletime time.Duration = 0 * time.Second

type PsqlClientConfig struct {
	// host of the psql server
	Host string
	// Default: 5432
	Port int
	// name of the psql database to connect to
	DbName string
	// authorized user name to connect to the psql datase
	UserName string
	// authorized user password to connect to the psql datase
	Password string
	// MaxOpenConns sets the maximum number of open connections to the database.
	// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
	// MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit.
	// If n <= 0, then there is no limit on the number of open connections.
	// The default is 0 (unlimited).
	// MaxOpenConnections <= 0 means unlimited
	MaxOpenConnections int
	// MaxIdleConns sets the maximum number of connections in the idle connection pool.
	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
	// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
	// If n <= 0, no idle connections are retained.
	// The default max idle connections is currently 2. This may change in a future release.
	// MaxIdleConnections zero means defaultMaxIdleConns; negative means 0
	MaxIdleConnections int
	// ConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// Expired connections may be closed lazily before reuse.
	// If d <= 0, connections are not closed due to a connection's age.
	// MaxLifetime maximum amount of time a connection may be reused
	MaxLifetime time.Duration
	// ConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	// Expired connections may be closed lazily before reuse.
	// If d <= 0, connections are not closed due to a connection's idle time.
	// MaxIdletime maximum amount of time a connection may be idle before being closed
	MaxIdletime time.Duration
}
