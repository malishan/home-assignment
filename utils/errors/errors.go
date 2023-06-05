package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"

	"github.com/malishan/home-assignment/utils/constants"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// Error is the type which can be used as an implementation of error
type Error struct {
	StatusCode int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details,omitempty"`
	trace      string
}

// wrap is the method to be used to wrap the error
type wrap interface {
	Wrap(err error) error
}

// unwrap is the method to be used to unwrap the error
type unwrap interface {
	Unwrap() error
}

// is used to check if it matches the error provided
type is interface {
	Is(err error) bool
}

// as is used to cast error to interface
type as interface {
	As(interface{}) bool
}

// WithStatusCode is used to create a new error with the status code changed
func (e Error) WithStatusCode(statusCode int) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       e.Code,
		Message:    e.Message,
		Details:    e.Details,
	}
}

// WithCode is used to create a new error with the code changed
func (e Error) WithCode(code string) *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       code,
		Message:    e.Message,
		Details:    e.Details,
	}
}

// WithMessage is used to create a new error with the message changed
func (e Error) WithMessage(message string) *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    message,
		Details:    e.Details,
	}
}

// WithDetails is used to attach the details to the error response
func (e Error) WithDetails(details interface{}) *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    e.Message,
		Details:    details,
	}
}

// WithPrefixInDetails is used to attach certain string prefix to existing details to the error response
func (e Error) WithPrefixInDetails(prefix string) *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    e.Message,
		Details:    fmt.Sprintf(prefix+": %v", e.Details),
	}
}

// Value is used to get the reference to the value
func (e Error) Value() *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    e.Message,
		Details:    e.Details,
	}
}

// GetTrace is used to get the current trace
func (e Error) GetTrace() string {
	if e.trace == constants.Empty {
		return string(debug.Stack())
	}
	return e.trace
}

// Error is used to get the detail from the error
func (e *Error) Error() string {
	errData, _ := json.Marshal(e)
	return string(errData)
}

// Wrap is used to wrap the error
func (e *Error) Wrap(err error) error {
	if d, ok := err.(*Error); ok {
		e.Details = d
	} else {
		e.Details = New(err.Error())
	}
	return e
}

// Unwrap is used to unwrap the error to the details
func (e *Error) Unwrap() error {
	// first check if details exist
	if e.Details == nil {
		return nil
	}
	// next try to cast the details to error
	if u, ok := e.Details.(error); ok {
		return u
	}
	// otherwise nothing
	return nil
}

// Is used to check if it matches the error provided
func (e *Error) Is(err error) bool {
	return e.Error() == err.Error()
}

// As is used to case the error type to the target
func (e *Error) As(target interface{}) bool {
	if t, ok := target.(*Error); ok {
		t.StatusCode = e.StatusCode
		t.Code = e.Code
		t.Message = e.Message
		t.Details = e.Details
		return true
	}
	return false
}

// New is used to create a new error
func New(message string) error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

// Wrap is used to wrap the error into another
func Wrap(parent, err error) error {
	if e, ok := parent.(wrap); ok {
		return e.Wrap(err)
	}
	return parent
}

// Unwrap is used to unwrap the error to the details
func Unwrap(err error) error {
	if u, ok := err.(unwrap); ok {
		return u.Unwrap()
	}
	return nil
}

// As finds the first error in error's chain that matches target, and if so, sets
// target to that error value and returns true. Otherwise, it returns false.
func As(err error, target interface{}) bool {
	if target == nil {
		return false
	}
	val := reflect.ValueOf(target)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr || val.IsNil() {
		return false
	}
	if e := typ.Elem(); e.Kind() != reflect.Interface && e.Kind() != reflect.Struct && !e.Implements(errorType) {
		return false
	}
	targetType := typ.Elem()
	for err != nil {
		if reflect.TypeOf(err).AssignableTo(targetType) {
			val.Elem().Set(reflect.ValueOf(err))
			return true
		}
		if x, ok := err.(as); ok && x.As(target) {
			return true
		}
		err = Unwrap(err)
	}
	return false
}
