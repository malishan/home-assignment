package impl

import "github.com/malishan/home-assignment/external/database"

type HealthAPIImpl struct {
	DbService database.DbService
}

func NewHealthAPIService(db database.DbService) *HealthAPIImpl {

	return &HealthAPIImpl{
		DbService: db,
	}
}
