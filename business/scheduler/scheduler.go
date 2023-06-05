package scheduler

import (
	"context"

	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

type ScheduleProvider interface {
	PollActivityOperation(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error)
}
