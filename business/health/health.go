package health

import (
	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
)

type HealthAPIProvider interface {
	ResourceHealthStatus(ctx *gin.Context) []*model.HealthApiResponse
}
