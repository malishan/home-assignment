package impl

import (
	"github.com/malishan/home-assignment/external/database"
	"github.com/malishan/home-assignment/external/http"
	"github.com/malishan/home-assignment/model"
)

type HomeAPIImpl struct {
	Config      *model.HomeConfig
	DbService   database.DbService
	HttpService http.HttpService
}

func NewHomeAPIService(config *model.HomeConfig, db database.DbService, httpService http.HttpService) *HomeAPIImpl {

	return &HomeAPIImpl{
		Config:      config,
		DbService:   db,
		HttpService: httpService,
	}
}
