package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/malishan/home-assignment/api/health"
	v1 "github.com/malishan/home-assignment/api/v1"
	"github.com/malishan/home-assignment/model"
	timeout "github.com/vearne/gin-timeout"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	customErr "github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/flags"
	"github.com/malishan/home-assignment/utils/logger"
	"github.com/malishan/home-assignment/utils/metrics"
	mw "github.com/malishan/home-assignment/utils/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/malishan/home-assignment/docs"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	router.Use(timeout.Timeout(
		timeout.WithTimeout(time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.ApiTimeout, 2))*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout),
		timeout.WithDefaultMsg(customErr.GetTimeoutError), // optional
		timeout.WithCallBack(func(r *http.Request) {
			logger.FileLogger.Error().Str(constants.RequestIDLogParam, r.Header.Get(constants.XRequestID)).Str(constants.UserIDLogParam, r.Header.Get(constants.XUserId)).
				Str(constants.PathLogParam, r.URL.Path).Stack().Err(errors.New("request timeout")).Msg("timeoutMiddleware : request timeout")
		}),
	))

	router.Use(mw.CORS(mw.CORSMiddlewareOptions{
		AllowedOrigins: "*",
		AllowedHeaders: "*",
		AllowedMethods: "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE",
	}))

	//api metrics
	router.Use(metrics.PrometheusMiddleware(model.ServiceName))

	//other metrics
	router.GET(constants.MetricRoute, gin.WrapH(promhttp.Handler()))

	//configure swagger
	router.GET(constants.SwaggerRoute, ginSwagger.WrapHandler(swaggerfiles.Handler))
	// if flags.Env() == constants.EnvDev || flags.Env() == constants.EnvTest || flags.Env() == constants.EnvIntegration || flags.Env() == constants.EnvPreProd {
	// 	router.GET(constant.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
	// }

	//version 1
	healthV1Router := router.Group(constants.HealthV1Route)
	health.HealthRoutes(healthV1Router)

	homeV1Router := router.Group(constants.HomeV1Route)
	v1.HomeRoutes(homeV1Router)

	homeSwagV1Router := router.Group(constants.HomeV1SwaggerDocRoute)
	v1.HomeSwaggerRoutes(homeSwagV1Router)

	logger.FileLogger.Info().Str("env", flags.Env()).Int("port", flags.Port()).Msg("Home Service - Initializing All Routes...")
	return router
}
