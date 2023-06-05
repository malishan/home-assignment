package v1

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/malishan/home-assignment/utils/constants"

	_ "github.com/malishan/home-assignment/docs"
)

// Home Routes
func HomeRoutes(group *gin.RouterGroup) {

	group.GET(constants.GetActivities, getActivities)

	group.GET(constants.SwaggerRoute, ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// Home Swagger Routes
func HomeSwaggerRoutes(group *gin.RouterGroup) {

	group.GET(constants.SwaggerRoute, ginSwagger.WrapHandler(swaggerfiles.Handler))
}
