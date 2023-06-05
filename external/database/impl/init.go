package impl

import (
	"github.com/malishan/home-assignment/external/database"
	psqlclient "github.com/malishan/home-assignment/utils/psqlClient"
)

type DbServiceImpl struct {
	DBClient psqlclient.PsqlClient
}

func GetDBServiceInstance(dbClient psqlclient.PsqlClient) database.DbService {
	return &DbServiceImpl{
		DBClient: dbClient,
	}
}
