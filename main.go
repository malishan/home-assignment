package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/malishan/home-assignment/api"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/flags"
	httpclient "github.com/malishan/home-assignment/utils/httpClient"
	"github.com/malishan/home-assignment/utils/logger"
	"github.com/malishan/home-assignment/utils/metrics"
	"github.com/malishan/home-assignment/utils/middlewares"
	psqlClient "github.com/malishan/home-assignment/utils/psqlClient"

	healthApi "github.com/malishan/home-assignment/api/health"
	scheduler "github.com/malishan/home-assignment/api/scheduler"
	schedulerJob "github.com/malishan/home-assignment/api/scheduler"
	homeApi "github.com/malishan/home-assignment/api/v1"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func init() {
	go handleSigTerm()
	ctx := context.Background()
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs + 1)

	initConfigs(ctx)
	initLogger(ctx)
	initDatabase(ctx)
	initHttpClient(ctx)
	initMetrics(ctx)
}

// @title 		Home Assignment
// @version 	1.0
// @description ## ## An assignment.
// @description
// @description X-User-Id 			e.g. => 'X-User-Id':'1234'
// @description X-Request-Id 		e.g. => 'X-Request-Id':'0bc86576911d7468f7bbd7d55fb2b72d'

// @termsOfService https://swagger.io/terms/

// @contact.name MD ALISHAN
// @contact.email ahmed.alishan3@gmail.com

// @BasePath /

func main() {

	ctx := context.Background()
	defer closeDatabase(ctx)

	components := LoadApplicationComponents()
	healthApi.InitHealthProvider(components.HealthAPIProvider)
	homeApi.InitHomeProvider(components.HomeAPIProvider)
	scheduler.InitScheduler(components.Scheduler)

	go startCron(ctx)

	startRouter()
}

func handleSigTerm() {
	done := make(chan os.Signal, 1)
	signal.Notify(done,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	sig := <-done

	switch sig {
	case syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT:
		closeDatabase(context.Background())
		os.Exit(0)
	}
}

func initConfigs(ctx context.Context) {

	err := configs.InitConfigClient(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": flags.BaseConfigPath() + "/" + flags.Env(),
			"configNames":      []string{constants.ApplicationConfigFile, constants.DatabaseConfigFile, constants.APIConfigFile, constants.LoggerConfigFile},
			"configType":       "yaml",
		},
	})

	if err != nil {
		logger.Fatal(ctx).Str("env", flags.Env()).Str("path", flags.BaseConfigPath()+"/"+flags.Env()).Stack().Err(err).Msg("unable to initialize config")
	}

	logger.Info(ctx).Str("env", flags.Env()).Str("path", flags.BaseConfigPath()+"/"+flags.Env()).Msg("initialized configs")
}

func initLogger(ctx context.Context) {

	err := logger.InitFileLogger(logger.FileConfig{
		EncodeLogsAsJson:      configs.Get().GetBoolD(constants.LoggerConfigFile, constants.EncodeLogAsJsonKey, true),
		ConsoleLoggingEnabled: configs.Get().GetBoolD(constants.LoggerConfigFile, constants.ConsoleLoggingEnabledKey, false),
		FileLoggingEnabled:    configs.Get().GetBoolD(constants.LoggerConfigFile, constants.FileLoggingEnabledKey, true),
		Directory:             configs.Get().GetStringD(constants.LoggerConfigFile, constants.DirectoryKey, "Log"),
		Filename:              configs.Get().GetStringD(constants.LoggerConfigFile, constants.FileNameKey, "app.log"),
		MaxSize:               int(configs.Get().GetIntD(constants.LoggerConfigFile, constants.MaxSizeInMBKey, 500)),
		MaxBackups:            int(configs.Get().GetIntD(constants.LoggerConfigFile, constants.MaxBackUpsKey, 1)),
		MaxAge:                int(configs.Get().GetIntD(constants.LoggerConfigFile, constants.MaxAgeInDaysKey, 5)),
		Compress:              configs.Get().GetBoolD(constants.LoggerConfigFile, constants.Compress, true),
	})

	if err != nil {
		logger.Fatal(ctx).Stack().Err(err).Msg("unable to initialize logger")
	}
}

func initDatabase(ctx context.Context) {

	config := &psqlClient.PsqlClientConfig{
		Host:               configs.Get().GetStringD(constants.DatabaseConfigFile, constants.PsqlHostName, ""),
		Port:               int(configs.Get().GetIntD(constants.DatabaseConfigFile, constants.PsqlPort, 0)),
		DbName:             configs.Get().GetStringD(constants.DatabaseConfigFile, constants.PsqlDbName, ""),
		UserName:           configs.Get().GetStringD(constants.DatabaseConfigFile, constants.PsqlUser, ""),
		Password:           configs.Get().GetStringD(constants.DatabaseConfigFile, constants.PsqlPassword, ""),
		MaxOpenConnections: int(configs.Get().GetIntD(constants.DatabaseConfigFile, constants.PsqlMaxOpenConn, 0)),
		MaxIdleConnections: int(configs.Get().GetIntD(constants.DatabaseConfigFile, constants.PsqlMaxIdleConn, 0)),
		MaxLifetime:        time.Duration(configs.Get().GetIntD(constants.DatabaseConfigFile, constants.PsqlConnMaxLifetime, 0)) * time.Hour,
		MaxIdletime:        time.Duration(configs.Get().GetIntD(constants.DatabaseConfigFile, constants.PsqlConnMaxIdleTime, 0)) * time.Minute,
	}

	if err := psqlClient.InitPsqlClient(config); err != nil {
		logger.FileLogger.CtxFatal(ctx).Stack().Interface("config", config).Err(err).Msg("unable to initialize psql")
	}

	logger.FileLogger.CtxInfo(ctx).Interface("config", config).Msg("psql initialized")
}

func initHttpClient(ctx context.Context) {

	config := &httpclient.HttpConfig{
		ConnectTimeout:        time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HTTPConnectTimeoutInMillisKey, 3000)) * time.Millisecond,
		KeepAliveDuration:     time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HTTPKeepAliveDurationInMillisKey, 30000)) * time.Millisecond,
		MaxIdleConnections:    int(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HTTPMaxIdleConnectionsKey, 10)),
		IdleConnectionTimeout: time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HTTPIdleConnectionTimeoutInMillisKey, 90000)) * time.Millisecond,
		TLSHandshakeTimeout:   time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HTTPTlsHandshakeTimeoutInMillisKey, 10000)) * time.Millisecond,
		ExpectContinueTimeout: time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HTTPExpectContinueTimeoutInMillisKey, 1000)) * time.Millisecond,
		Timeout:               time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HTTPTimeoutInMillisKey, 1000)) * time.Millisecond,
	}

	if err := httpclient.InitHTTPClient(nil, *config); err != nil {
		logger.FileLogger.CtxFatal(ctx).Stack().Interface("config", config).Err(err).Msg("unable to initialize http client")
	}

	logger.FileLogger.CtxInfo(ctx).Interface("config", config).Msg("http client initialized")
}

func closeDatabase(ctx context.Context) {
	err := psqlClient.ClosePsqlClient()
	if err != nil {
		logger.FileLogger.CtxFatal(ctx).Str("env", flags.Env()).Stack().Err(err).Msg("error closing psql")
	}
	logger.FileLogger.CtxInfo(ctx).Msg("psql closed")
}

func initMetrics(ctx context.Context) {
	if err := metrics.IniPrometheustMetrics(ctx,
		metrics.HTTPTotalRequestCounterEnable,
		metrics.HTTPResponseStatusCounterEnable,
		metrics.HTTPRequestCounterEnable,

		metrics.HTTPRequestTimingEnable,
		metrics.PsqlTimingEnable,
	); err != nil {
		logger.FileLogger.CtxFatal(ctx).Stack().Err(err).Msg("unable to initialize prometheus metrics")
	}
	logger.FileLogger.CtxInfo(ctx).Msg("metrics initialized")
}

func startCron(ctx context.Context) {

	err := schedulerJob.StartCronJob()
	if err != nil {
		logger.FileLogger.CtxFatal(ctx).Stack().Err(err).Msg("unable to initialize cron jobs")
	}

	logger.FileLogger.CtxInfo(ctx).Msg("cron initialized")
}

func startRouter() {

	router := api.GetRouter(middlewares.Logger(middlewares.LoggerMiddlewareOptions{FileLogEnable: true}))

	err := router.Run(fmt.Sprintf(":%d", flags.Port()))
	if err != nil {
		logger.FileLogger.Fatal().Err(err).Msg("error starting router")
	}
}
