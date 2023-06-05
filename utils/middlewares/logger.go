package middlewares

import (
	"time"

	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/logger"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

// LoggerMiddlewareOptions is the set of configurable allowed for log
type LoggerMiddlewareOptions struct {
	ConsoleLogEnable bool
	FileLogEnable    bool
}

// Logger is the middleware to be used for logging the request and response information
func Logger(options LoggerMiddlewareOptions) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		id := ctx.GetHeader(constants.XRequestID)

		if id == "" {
			// get a unique id
			uid, err := uuid.NewUUID()
			if err == nil {
				id = uid.String()
			}
			ctx.Request.Header.Add(constants.XRequestID, id)
		}

		// apply requestID in the context
		ctx.Set(constants.RequestIDLogParam, id)

		// apply correlationID in the context
		correlationID := ctx.GetHeader(constants.XCorrelationID)
		if correlationID != "" {
			ctx.Set(constants.CorrelationLogParam, correlationID)
		}

		// apply path in the context
		path := ctx.Request.URL.Path
		ctx.Set(constants.PathLogParam, path)

		// apply userID in the context
		userId := ctx.GetHeader(constants.XUserId)
		ctx.Set(constants.UserIDLogParam, userId)

		ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Start timer
		start := time.Now()

		// Process request
		ctx.Next()

		// stop timer
		end := time.Now()
		latency := end.Sub(start).Milliseconds()

		if options.ConsoleLogEnable {
			logger.Info(ctx).
				Time(constants.StartTimeLogParam, start).
				Str(constants.MethodLogParam, ctx.Request.Method).
				Interface(constants.RequestHeaderLogParams, ctx.Request.Header).
				Str(constants.ClientIPLogParam, ctx.ClientIP()).
				Int(constants.StatusCodeLogParam, ctx.Writer.Status()).
				Time(constants.EndTimeLogParam, end).
				Interface(constants.ResponseHeaderLogParams, ctx.Writer.Header()).
				Int64(constants.LatencyLogParam, latency).
				Str(constants.ErrorLogParam, ctx.Errors.ByType(gin.ErrorTypePrivate).String()).
				Send()
		}

		if options.FileLogEnable {
			logger.FileLogger.CtxInfo(ctx).
				Time(constants.StartTimeLogParam, start).
				Str(constants.MethodLogParam, ctx.Request.Method).
				Interface(constants.RequestHeaderLogParams, ctx.Request.Header).
				Str(constants.ClientIPLogParam, ctx.ClientIP()).
				Int(constants.StatusCodeLogParam, ctx.Writer.Status()).
				Time(constants.EndTimeLogParam, end).
				Interface(constants.ResponseHeaderLogParams, ctx.Writer.Header()).
				Int64(constants.LatencyLogParam, latency).
				Str(constants.ErrorLogParam, ctx.Errors.ByType(gin.ErrorTypePrivate).String()).
				Send()
		}
	}
}
