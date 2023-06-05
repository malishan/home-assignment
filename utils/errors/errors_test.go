package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	e := New("hello")
	assert.Error(t, e)
	assert.Equal(t, "{\"code\":\"\",\"message\":\"hello\"}", e.Error())

	er, ok := e.(*Error)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, er.StatusCode)
	assert.Equal(t, "hello", er.Message)
}

func TestError(t *testing.T) {
	e := (Error{Code: "naruto"}).Value().WithStatusCode(http.StatusOK).WithMessage("naruto").WithDetails("naruto")
	assert.Equal(t, "naruto", e.Code)
	assert.Equal(t, "naruto", e.Message)
	assert.Equal(t, "naruto", e.Details)
	assert.Equal(t, http.StatusOK, e.StatusCode)
	assert.Equal(t, "{\"code\":\"naruto\",\"message\":\"naruto\",\"details\":\"naruto\"}", e.Error())
}

func TestErrorWrapUnwrap(t *testing.T) {
	e := New("naruto")
	assert.Equal(t, "{\"code\":\"\",\"message\":\"naruto\"}", e.Error())
	e = Wrap(e, Wrap(New("hinata"), errors.New("boruto")))
	assert.Equal(t, "{\"code\":\"\",\"message\":\"naruto\",\"details\":{\"code\":\"\",\"message\":\"hinata\",\"details\":{\"code\":\"\",\"message\":\"boruto\"}}}", e.Error())
	e = Unwrap(e)
	assert.Error(t, e)
	assert.Equal(t, "{\"code\":\"\",\"message\":\"hinata\",\"details\":{\"code\":\"\",\"message\":\"boruto\"}}", e.Error())
	e = Unwrap(e)
	assert.Error(t, e)
	assert.Equal(t, "{\"code\":\"\",\"message\":\"boruto\"}", e.Error())
	e = errors.New("naruto")
	assert.Equal(t, e, Wrap(e, New("naruto")))
	assert.NoError(t, Unwrap(e))
	e = (Error{}).Value().WithDetails("naruto")
	assert.NoError(t, Unwrap(e))
	e = (Error{}).Value()
	assert.NoError(t, Unwrap(e))
}

func TestErrorAs(t *testing.T) {
	var e Error
	assert.False(t, As(New("naruto"), nil))
	assert.False(t, As(New("naruto"), e))
	assert.True(t, As(New("naruto"), &e))
	assert.Equal(t, "naruto", e.Message)
	assert.Equal(t, http.StatusInternalServerError, e.StatusCode)
	var x string
	assert.False(t, As(New("naruto"), &x))
	var y struct{}
	assert.False(t, As(New("naruto"), &y))
	assert.False(t, As(Wrap(New("naruto"), New("boruto")), &y))
}
