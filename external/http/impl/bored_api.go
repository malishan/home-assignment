package impl

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/configs"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
)

func (impl *HttpServiceImpl) BoredApiCall(ctx *gin.Context) (*model.BoredApiResponse, *errors.Error) {

	url, err := configs.Get().GetString(constants.APIConfigFile, constants.BoredApiUrlKey)
	if err != nil {
		return nil, &errors.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       errors.InternalServerErrorCode,
			Message:    errors.InternalServerError,
			Details:    "boredApiCall: fetch url config failed - " + err.Error(),
		}
	}

	limtErr := impl.Ratelimiter.Wait(ctx)
	if limtErr != nil {
		return nil, &errors.Error{
			StatusCode: http.StatusTooManyRequests,
			Code:       errors.InternalServerErrorCode,
			Message:    errors.InternalServerError,
			Details:    "boredApiCall: rate limit error - " + limtErr.Error(),
		}
	}

	res, err := impl.HttpClient.GetResponse(ctx.Request.Context(), url, nil, nil, nil)
	if err != nil {
		return nil, &errors.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       errors.InternalServerErrorCode,
			Message:    errors.InternalServerError,
			Details:    "boredApiCall: http req failed - " + err.Error(),
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &errors.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       errors.InternalServerErrorCode,
			Message:    errors.InternalServerError,
			Details:    "boredApiCall: read all failed - " + err.Error(),
		}
	}

	defer res.Body.Close()

	if string(body) == "" || res.StatusCode != http.StatusOK {
		return nil, &errors.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       errors.InternalServerErrorCode,
			Message:    errors.InternalServerError,
			Details:    "boredApiCall: incorrect response - " + string(body),
		}
	}

	var boredAPIResp model.BoredApiResponse

	err = json.Unmarshal(body, &boredAPIResp)
	if err != nil {
		return nil, &errors.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       errors.InternalServerErrorCode,
			Message:    errors.InternalServerError,
			Details:    "boredApiCall: resp unmarshal failed - " + err.Error(),
		}
	}

	return &boredAPIResp, nil
}
