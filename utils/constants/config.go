package constants

// config file names
const (
	LoggerConfigFile      = "logger"
	ApplicationConfigFile = "application"
	DatabaseConfigFile    = "database"
	APIConfigFile         = "api"
)

// application configuration keywords
const (
	ApiTimeout               = "apiTimeout"
	ActivityApiResponseCount = "boredApiActivityCount"
	HttpRateLimitCount       = "httpApiRateLimitCount"
	HttpRateLimitDuration    = "httpApiRateLimitInSecond"

	HTTPConnectTimeoutInMillisKey        = "http.connectTimeoutInMillis"
	HTTPKeepAliveDurationInMillisKey     = "http.keepAliveDurationInMillis"
	HTTPMaxIdleConnectionsKey            = "http.maxIdleConnections"
	HTTPIdleConnectionTimeoutInMillisKey = "http.idleConnectionTimeoutInMillis"
	HTTPTlsHandshakeTimeoutInMillisKey   = "http.tlsHandshakeTimeoutInMillis"
	HTTPExpectContinueTimeoutInMillisKey = "http.expectContinueTimeoutInMillis"
	HTTPTimeoutInMillisKey               = "http.timeoutInMillis"

	PollActivityCronTiming = "pollActivityCronTiming"
)

// log configuration keywords
const (
	LogLevelKey              = "level"
	ConsoleLoggingEnabledKey = "consoleLoggingEnabled"
	FileLoggingEnabledKey    = "fileLoggingEnabled"
	DebugLoggingEnabledKey   = "debugLoggingEnabled"
	EncodeLogAsJsonKey       = "encodeLogsAsJson"
	DirectoryKey             = "directory"
	FileNameKey              = "fileName"
	MaxSizeInMBKey           = "maxSizeInMB"
	MaxBackUpsKey            = "maxBackups"
	MaxAgeInDaysKey          = "maxAgeInDays"
	Compress                 = "compress"
)

// db configuration keywords
const (
	PsqlDatabase        = "psqlDB"
	PsqlHostName        = PsqlDatabase + "." + "host"
	PsqlPort            = PsqlDatabase + "." + "port"
	PsqlDbName          = PsqlDatabase + "." + "dbName"
	PsqlUser            = PsqlDatabase + "." + "user"
	PsqlPassword        = PsqlDatabase + "." + "password"
	PsqlMaxOpenConn     = PsqlDatabase + "." + "maxOpenConnections"
	PsqlMaxIdleConn     = PsqlDatabase + "." + "maxIdleConnections"
	PsqlConnMaxLifetime = PsqlDatabase + "." + "connMaxLifetimeInHour"
	PsqlConnMaxIdleTime = PsqlDatabase + "." + "connMaxIdleTimeInMinute"
)

// api configuration keywords
const (
	BoredApi       = "boredApi"
	BoredApiUrlKey = BoredApi + "." + "url"
)
