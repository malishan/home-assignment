package logger

import (
	"context"
	"io"
	"runtime/debug"

	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
Log Levels:
	TRACE.
	DEBUG.
	INFO.
	WARN.
	ERROR.
	FATAL.
*/

type Level string

func (l Level) zeroLogLevel() zerolog.Level {
	switch l {
	case constants.TraceLevel:
		return zerolog.TraceLevel
	case constants.DebugLevel:
		return zerolog.DebugLevel
	case constants.InfoLevel:
		return zerolog.InfoLevel
	case constants.WarnLevel:
		return zerolog.WarnLevel
	case constants.ErrorLevel:
		return zerolog.ErrorLevel
	case constants.FatalLevel:
		return zerolog.FatalLevel
	case constants.PanicLevel:
		return zerolog.PanicLevel
	default:
		return zerolog.DebugLevel
	}
}

// InitLogger is used to initialize console logger
func InitLogger(level Level) {
	zerolog.ErrorStackMarshaler = getErrorStackMarshaller()
	zerolog.SetGlobalLevel(level.zeroLogLevel())
	log.Logger = log.With().Caller().Logger()
}

// InitLoggerWithWriter is used to initialize logger with a writer
func InitLoggerWithWriter(level Level, w io.Writer) {
	zerolog.ErrorStackMarshaler = getErrorStackMarshaller()
	zerolog.SetGlobalLevel(level.zeroLogLevel())
	log.Logger = zerolog.New(w).With().Caller().Timestamp().Logger()
}

// Trace is the console context logging for trace log
func Trace(ctx context.Context) *zerolog.Event {
	return withIDAndPath(ctx, log.Trace())
}

// Debug is the console context logging for debug log
func Debug(ctx context.Context) *zerolog.Event {
	return withIDAndPath(ctx, log.Debug())
}

// Info is the console context logging for info log
func Info(ctx context.Context) *zerolog.Event {
	return withIDAndPath(ctx, log.Info())
}

// Warn is the console context logging for warn log
func Warn(ctx context.Context) *zerolog.Event {
	return withIDAndPath(ctx, log.Warn())
}

// Error is the console context logging for error log
func Error(ctx context.Context) *zerolog.Event {
	return withIDAndPath(ctx, log.Error().Stack())
}

// Panic is the console context logging for panic log
func Panic(ctx context.Context) *zerolog.Event {
	return withIDAndPath(ctx, log.Panic().Stack())
}

// Fatal is the console context logging for fatal log
func Fatal(ctx context.Context) *zerolog.Event {
	return withIDAndPath(ctx, log.Fatal().Stack())
}

// ErrorWarn checks for the error object.
// In case it is corresponding to a 4XX status code, it logs it as warning.
// Otherwise, it logs it as an error.
func ErrorWarn(ctx context.Context, err error) *zerolog.Event {
	if e, ok := err.(*errors.Error); ok && e.StatusCode >= 400 && e.StatusCode < 500 {
		return Warn(ctx).Err(err)
	}
	return Error(ctx).Err(err)
}

func getErrorStackMarshaller() func(err error) interface{} {
	return func(err error) interface{} {
		return string(debug.Stack())
	}
}

func withIDAndPath(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		return event
	}
	id := ctx.Value(constants.RequestIDLogParam)
	if id != nil {
		event.Interface(constants.RequestIDLogParam, id)
	}
	path := ctx.Value(constants.PathLogParam)
	if path != nil {
		event.Interface(constants.PathLogParam, path)
	}
	correlationId := ctx.Value(constants.CorrelationLogParam)
	if correlationId != nil {
		event.Interface(constants.CorrelationLogParam, correlationId)
	}
	userId := ctx.Value(constants.UserIDLogParam)
	if userId != nil {
		event.Interface(constants.UserIDLogParam, userId)
	}
	return event
}
