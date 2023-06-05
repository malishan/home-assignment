package impl

import (
	"context"

	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

func (impl SchedulerImpl) PollActivityOperation(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {
	return impl.DbService.FetchUserActivityRecord(ctx)
}
