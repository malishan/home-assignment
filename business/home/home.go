package home

import (
	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
)

type HomeAPIProvider interface {
	GetActivities(ctx *gin.Context) ([]*model.BoredApiResponse, *errors.Error)
}
