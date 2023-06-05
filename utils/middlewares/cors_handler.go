package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/utils/constants"
)

// CORSMiddlewareOptions contains a set of notion for setting cors middleware
type CORSMiddlewareOptions struct {
	AllowedOrigins   string
	allowCredentials string
	BlockCredentials bool
	AllowedHeaders   string
	AllowedMethods   string
}

// init is used to set defaults to the options
func (o *CORSMiddlewareOptions) init() {
	o.AllowedOrigins = strings.TrimSpace(o.AllowedOrigins)
	if o.AllowedOrigins == "" {
		o.AllowedOrigins = constants.AllHeaderValue
	}
	if o.BlockCredentials {
		o.allowCredentials = constants.FalseHeaderValue
	} else {
		o.allowCredentials = constants.TrueHeaderValue
	}
	o.AllowedHeaders = strings.TrimSpace(strings.ToUpper(o.AllowedHeaders))
	if o.AllowedHeaders == "" {
		o.AllowedHeaders = constants.AllHeaderValue
	}
	o.AllowedMethods = strings.TrimSpace(strings.ToUpper(o.AllowedMethods))
	if o.AllowedMethods == "" {
		o.AllowedMethods = constants.AllMethodsHeaderValue
	}
}

// CORS is used to allow CORS for the requests this is added to
func CORS(options CORSMiddlewareOptions) gin.HandlerFunc {
	options.init()
	return func(ctx *gin.Context) {
		ctx.Header(constants.AccessControlAllowOriginHeader, options.AllowedOrigins)
		ctx.Header(constants.AccessControlAllowCredentialsHeader, options.allowCredentials)
		ctx.Header(constants.AccessControlAllowHeadersHeader, options.AllowedHeaders)
		ctx.Header(constants.AccessControlAllowMethodsHeader, options.AllowedMethods)
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()

	}
}
