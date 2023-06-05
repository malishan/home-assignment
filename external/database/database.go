package database

import (
	"context"

	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

type DbService interface {
	PingStatus(ctx context.Context) *errors.Error
	InsertUserActivityRecord(ctx context.Context, id, userId string, record []*model.BoredApiResponse) *errors.Error
	FetchUserActivityRecord(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error)
}
