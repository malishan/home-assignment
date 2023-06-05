package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/malishan/home-assignment/utils/constants"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	r := gin.New()
	r.Use(CORS(CORSMiddlewareOptions{}))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, constants.AllHeaderValue, w.Header().Get(constants.AccessControlAllowOriginHeader))
	assert.Equal(t, constants.TrueHeaderValue, w.Header().Get(constants.AccessControlAllowCredentialsHeader))
	assert.Equal(t, constants.AllHeaderValue, w.Header().Get(constants.AccessControlAllowHeadersHeader))
	assert.Equal(t, constants.AllMethodsHeaderValue, w.Header().Get(constants.AccessControlAllowMethodsHeader))
}

func TestCORSOptions(t *testing.T) {
	r := gin.New()
	r.Use(CORS(CORSMiddlewareOptions{}))
	r.OPTIONS("/ping", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("OPTIONS", "/ping", nil)
	assert.NoError(t, err)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
	assert.Equal(t, constants.AllHeaderValue, w.Header().Get(constants.AccessControlAllowOriginHeader))
	assert.Equal(t, constants.TrueHeaderValue, w.Header().Get(constants.AccessControlAllowCredentialsHeader))
	assert.Equal(t, constants.AllHeaderValue, w.Header().Get(constants.AccessControlAllowHeadersHeader))
	assert.Equal(t, constants.AllMethodsHeaderValue, w.Header().Get(constants.AccessControlAllowMethodsHeader))
}

func TestCORSWithCustomParams(t *testing.T) {
	r := gin.New()
	r.Use(CORS(CORSMiddlewareOptions{
		AllowedOrigins:   "naruto",
		BlockCredentials: true,
		AllowedHeaders:   "naruto",
		AllowedMethods:   "naruto",
	}))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "naruto", w.Header().Get(constants.AccessControlAllowOriginHeader))
	assert.Equal(t, constants.FalseHeaderValue, w.Header().Get(constants.AccessControlAllowCredentialsHeader))
	assert.Equal(t, "NARUTO", w.Header().Get(constants.AccessControlAllowHeadersHeader))
	assert.Equal(t, "NARUTO", w.Header().Get(constants.AccessControlAllowMethodsHeader))
}
