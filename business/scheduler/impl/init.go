package impl

import (
	"github.com/malishan/home-assignment/external/database"
	"github.com/malishan/home-assignment/model"
)

type SchedulerImpl struct {
	Config    *model.SchedulerConfig
	DbService database.DbService
}

func NewSchedulerService(config *model.SchedulerConfig, db database.DbService) *SchedulerImpl {

	return &SchedulerImpl{
		Config:    config,
		DbService: db,
	}
}
