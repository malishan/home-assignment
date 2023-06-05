package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"runtime/debug"

	"github.com/google/uuid"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/logger"
)

func PollActivityOperation() {
	defer func() {
		if r := recover(); r != nil {
			logger.FileLogger.Error().Str("stack", string(debug.Stack())).Interface("recovery", r).Msg("pollActivityOperation : cron panic recovery")
		}
	}()

	ctx := context.WithValue(context.Background(), constants.IDLogParam, uuid.New().String())
	ctx = context.WithValue(ctx, constants.PathLogParam, "Cron-PollActivity")

	response, err := schedulerHandler.PollActivityOperation(ctx)
	if err != nil {
		logger.FileLogger.CtxError(ctx).Stack().Err(errors.New(err.Details.(string))).Msg("PollActivityOperation : provider failed")
		return
	}

	responseBytes, _ := json.Marshal(&response)

	logger.FileLogger.CtxInfo(ctx).Str(constants.ResponseBodyLogParams, string(responseBytes)).Msg("PollActivityOperation: success result")
}
