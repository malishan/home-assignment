package health

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/logger"
)

// HealthStatus godoc
// @Summary 		Provides resource health status.
// @Description 	Health status of resources like database, etc.
// @ID 				healthStatus
// @Tags 			Home
// @Accept      	json
// @Produce 		json
// @Param X-User-Id 				header string false "X-User-Id Header"
// @Param X-Request-Id 				header string false "X-Request-Id Header"
// @Param X-Location 				header string false "X-Location Header"
// @Success 	200 	{object} 	model.APISuccessResponse{data=[]model.HealthApiResponse}
// @Failure 	400 	{object} 	model.APIFailureResponse
// @Failure 	500 	{object} 	model.APIFailureResponse
// @Router /health/v1/status [get]
func healthStatus(ctx *gin.Context) {

	logger.FileLogger.CtxInfo(ctx).Interface(constants.RequestHeaderLogParams, ctx.Request.Header).Msg("healthStatus : req body")

	response := healthProvider.ResourceHealthStatus(ctx)

	rawRsp, _ := json.Marshal(response)
	logger.FileLogger.CtxInfo(ctx).Str(constants.ResponseBodyLogParams, string(rawRsp)).Msg("healthStatus : success")

	ctx.JSON(http.StatusOK, &model.APISuccessResponse{
		Status: true,
		Data:   response,
	})
}
