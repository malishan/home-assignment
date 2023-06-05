package impl

import (
	"context"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/logger"
)

func (impl HomeAPIImpl) GetActivities(ctx *gin.Context) ([]*model.BoredApiResponse, *errors.Error) {

	var (
		wg             sync.WaitGroup
		mutex          sync.RWMutex
		httpErr        *errors.Error
		distinctKeyMap = make(map[string]string)
	)

	//NOTE: ideally channel should be preferred to communicate via go-routines

	for i := 0; i < impl.Config.ActivityCountInResponse; i++ {
		i := i
		wg.Add(1)
		go func(w *sync.WaitGroup, m *sync.RWMutex) {
			defer w.Done()
			defer func() {
				if r := recover(); r != nil {
					logger.FileLogger.CtxError(ctx).Str("stack", string(debug.Stack())).Int("go-routine", i).Interface("recovery", r).Msg("getBoredApi : http call panic recovery")
				}
			}()
			resp, err := impl.HttpService.BoredApiCall(ctx)
			if err != nil {
				httpErr = err
			} else {
				m.RLock()
				_, ok := distinctKeyMap[resp.Key]
				if ok {
					httpErr = &errors.Error{Code: errors.InvalidDetailsCode, Message: "DUPLICATE RECORD FOUND", Details: "duplicate key-" + resp.Key}
				}
				m.RUnlock()

				m.Lock()
				distinctKeyMap[resp.Key] = resp.Activity
				m.Unlock()
			}
		}(&wg, &mutex)
	}

	wg.Wait()

	if httpErr != nil {
		return nil, httpErr
	}

	response := make([]*model.BoredApiResponse, 0)
	for key, val := range distinctKeyMap {
		response = append(response, &model.BoredApiResponse{Key: key, Activity: val})
	}

	go func(c *gin.Context, res []*model.BoredApiResponse) {
		defer func() {
			if r := recover(); r != nil {
				logger.FileLogger.CtxError(c).Str("stack", string(debug.Stack())).Interface("recovery", r).Msg("getBoredApi : db call panic recovery")
			}
		}()
		dbContext, cancel := context.WithTimeout(c, 1*time.Second)
		defer cancel()
		if err := impl.DbService.InsertUserActivityRecord(dbContext, ctx.GetString(constants.RequestIDLogParam), ctx.GetString(constants.UserIDLogParam), res); err != nil {
			logger.FileLogger.CtxError(c).Stack().Err(errors.New(err.Details.(string))).Msg("getBoredApi : db insertion failed")
		}
	}(ctx, response)

	return response, nil
}
