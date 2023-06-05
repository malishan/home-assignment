package metrics

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(HTTPRequestTiming.WithLabelValues(getHandlerName(c, serviceName)))
		HTTPTotalRequestCounter.WithLabelValues(getHandlerName(c, serviceName), serviceName).Inc()
		c.Next()
		HTTPResponseStatusCounter.WithLabelValues(strconv.Itoa(c.Writer.Status())).Inc()
		HTTPRequestCounter.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Host, getHandlerName(c, serviceName), c.Request.Method).Inc()
		timer.ObserveDuration()
	}
}

func getHandlerName(ctx *gin.Context, serviceName string) string {
	handlerName := ctx.HandlerName()
	if handlerName == constants.Empty {
		return constants.Empty
	}
	modifiedHandlerName := strings.ReplaceAll(handlerName, serviceName+"/api/", "")
	return modifiedHandlerName
}
