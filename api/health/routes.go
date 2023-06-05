package health

import (
	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/utils/constants"
)

// Health Routes
func HealthRoutes(group *gin.RouterGroup) {

	group.GET(constants.HealthStatus, healthStatus)
}
