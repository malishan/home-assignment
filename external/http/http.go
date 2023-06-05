package http

import (
	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

type HttpService interface {
	BoredApiCall(ctx *gin.Context) (*model.BoredApiResponse, *errors.Error)
}
