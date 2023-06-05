package errors

var ErrorCodeMap = map[string]string{
	InternalServerErrorCode:   InternalServerError,
	InvalidReqParameterCode:   InvalidReqParameters,
	RequestDecodeErrorCode:    RequestDecodeError,
	ContradictoryReqParamCode: ContradictoryReqParam,
}

var GetTimeoutError string = "{\"status\":false,\"errorcode\":\"OE0001\",\"message\":\"Activity API Not Available\"}"
