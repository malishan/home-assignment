package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/logger"
)

// GetActivities godoc
// @Summary 		Get Activities Api returns 3 distinct activity values.
// @Description 	It fetches the response from boredapi.com and returns three unique keys with their activity value
// @ID 				getActivities
// @Tags 			Home
// @Accept      	json
// @Produce 		json
// @Param X-User-Id 				header string false "X-User-Id Header"
// @Param X-Request-Id 				header string false "X-Request-Id Header"
// @Param X-Location 				header string false "X-Location Header"
// @Success 	200 	{object} 	model.APISuccessResponse{data=[]model.BoredApiResponse}
// @Failure 	400 	{object} 	model.APIFailureResponse
// @Failure 	500 	{object} 	model.APIFailureResponse
// @Router /home/v1/activities [get]
func getActivities(ctx *gin.Context) {

	logger.FileLogger.CtxInfo(ctx).Interface(constants.RequestHeaderLogParams, ctx.Request.Header).Msg("getBoredApi : req body")

	response, err := apiProvider.GetActivities(ctx)
	if err != nil {
		logger.FileLogger.CtxError(ctx).Stack().Err(errors.New(err.Details.(string))).Msg("getBoredApi : provider failed")
		ctx.JSON(err.StatusCode, model.APIFailureResponse{
			ErrorCode: err.Code,
			Message:   err.Message,
		})
		return
	}

	rawRsp, _ := json.Marshal(response)
	logger.FileLogger.CtxInfo(ctx).Str(constants.ResponseBodyLogParams, string(rawRsp)).Msg("getBoredApi : success")

	ctx.JSON(http.StatusOK, &model.APISuccessResponse{
		Status: true,
		Data:   response,
	})
}
