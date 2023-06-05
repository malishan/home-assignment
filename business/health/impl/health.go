package impl

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
)

func (impl HealthAPIImpl) ResourceHealthStatus(ctx *gin.Context) []*model.HealthApiResponse {

	dbContext, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resp := make([]*model.HealthApiResponse, 0)

	err := impl.DbService.PingStatus(dbContext)
	if err == nil {
		resp = append(resp, &model.HealthApiResponse{Resource: "database", Status: "active", Message: ""})
	} else {
		resp = append(resp, &model.HealthApiResponse{Resource: "database", Status: "inactive", Message: "error: " + err.Details.(string)})
	}

	return resp
}
