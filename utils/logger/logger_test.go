package logger

import (
	"context"
	"testing"

	"github.com/malishan/home-assignment/utils/constants"
)

func TestInitLogger(t *testing.T) {
	InitLogger(constants.TraceLevel)
	InitLogger(constants.DebugLevel)
	InitLogger(constants.InfoLevel)
	InitLogger(constants.WarnLevel)
	InitLogger(constants.ErrorLevel)
	InitLogger(constants.PanicLevel)
	InitLogger(constants.FatalLevel)
	Trace(context.Background())
	Debug(context.Background())
	Info(context.Background())
	Warn(context.Background())
	Error(context.Background())
	Panic(context.Background())
	Fatal(context.Background())
}
