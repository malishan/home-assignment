package scheduler

import (
	"github.com/malishan/home-assignment/business/scheduler"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/robfig/cron/v3"
)

var schedulerHandler scheduler.ScheduleProvider

func InitScheduler(provider scheduler.ScheduleProvider) {
	schedulerHandler = provider
}

func StartCronJob() error {

	pollActivityTiming := configs.Get().GetStringD(constants.ApplicationConfigFile, constants.PollActivityCronTiming, "*/5 * * * *")

	cronJob := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))

	cronJob.AddFunc(pollActivityTiming, PollActivityOperation)

	cronJob.Start()

	return nil
}
