package main

import (
	"time"

	healthBusiness "github.com/malishan/home-assignment/business/health"
	homeBusiness "github.com/malishan/home-assignment/business/home"
	schedulerBusiness "github.com/malishan/home-assignment/business/scheduler"

	health "github.com/malishan/home-assignment/business/health/impl"
	home "github.com/malishan/home-assignment/business/home/impl"
	scheduler "github.com/malishan/home-assignment/business/scheduler/impl"

	dbService "github.com/malishan/home-assignment/external/database/impl"
	httpService "github.com/malishan/home-assignment/external/http/impl"

	"github.com/malishan/home-assignment/model"

	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	httpclient "github.com/malishan/home-assignment/utils/httpClient"
	psqlclient "github.com/malishan/home-assignment/utils/psqlClient"
)

type Components struct {
	HealthAPIProvider healthBusiness.HealthAPIProvider
	HomeAPIProvider   homeBusiness.HomeAPIProvider
	Scheduler         schedulerBusiness.ScheduleProvider
}

func LoadApplicationComponents() Components {

	homeConfig := &model.HomeConfig{
		ActivityCountInResponse: int(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.ActivityApiResponseCount, 3)),
		RateLimitCount:          int(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HttpRateLimitCount, 60)),
		RateLimitDuration:       time.Duration(configs.Get().GetIntD(constants.ApplicationConfigFile, constants.HttpRateLimitDuration, 1)) * time.Second,
	}

	schedulerConfig := &model.SchedulerConfig{}

	dbClient := psqlclient.GetPsqlClient()
	httpClient := httpclient.GetHttpClient()

	dbService := dbService.GetDBServiceInstance(dbClient)
	httpService := httpService.GetHttpServiceInstance(homeConfig.RateLimitCount, homeConfig.RateLimitDuration, httpClient)

	return Components{
		HealthAPIProvider: health.NewHealthAPIService(dbService),
		HomeAPIProvider:   home.NewHomeAPIService(homeConfig, dbService, httpService),
		Scheduler:         scheduler.NewSchedulerService(schedulerConfig, dbService),
	}
}
