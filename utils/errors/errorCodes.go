package errors

const (
	SuccessInfoErrorCode = ""
	SuccessInfoMessage   = "SUCCESS"

	FailedInfoErrorCode = ""
	FailedInfoMessage   = "FAILED"
)

// common error
const (
	InternalServerErrorCode   = "CE0001"
	InvalidReqParameterCode   = "CE0002"
	MethodNotAllowedCode      = "CE0003"
	RequestDecodeErrorCode    = "CE0004"
	ContradictoryReqParamCode = "CE0005"

	InternalServerError   = "INTERNAL SERVER ERROR"
	InvalidReqParameters  = "INVALID REQUEST PARAMETERS"
	MethodNotAllowed      = "METHOD NOT ALLOWED"
	RequestDecodeError    = "REQUEST DECODING ERROR"
	ContradictoryReqParam = "CONTRADICTORY REQUEST PARAM"
)

// user error
const (
	UserNotFoundCode    = "UE0001"
	UserBlockedCode     = "UE0002"
	UserExistsErrorCode = "UE0003"
	DuplicateUserCode   = "UE0004"

	UserNotFound       = "USER NOT FOUND"
	UserBlockedError   = "USER IS BLOCKED"
	UserExistsError    = "USER ALREADY EXISTS"
	DuplicateUserError = "DUPLICATE USER FOUND"
)

// validation error
const (
	UserIDInvalidCode          = "VE0001"
	UserTypeInvalidCode        = "VE0002"
	UnauthorizedUserCode       = "VE0003"
	IDPassInvalidCode          = "VE0004"
	InvalidPassCode            = "VE0005"
	TokenExpiredCode           = "VE0006"
	RenewTokenExpiredCode      = "VE0007"
	InvalidTokenCode           = "VE0008"
	InvalidTokenDetailsCode    = "VE0009"
	SimilarLast3PassCode       = "VE0010"
	OTPUsedCode                = "VE0011"
	OTPInvalidCode             = "VE0012"
	OTPExpiredCode             = "VE0013"
	LoginLimitReachedCode      = "VE0014"
	RetryLoginLimitReachedCode = "VE0015"

	UserIDInvalid          = "USER ID IS INVALID"
	UserTypeInvalid        = "USER TYPE IS INVALID"
	UnauthorizedUser       = "UNAUTHORIZED USER"
	IDPassInvalid          = "ID-PASSWORD INVALID"
	InvalidPassword        = "INVALID PASSWORD"
	TokenExpired           = "TOKEN EXPIRED"
	RenewTokenExpired      = "RENEW TOKEN EXPIRED"
	InvalidToken           = "INVALID TOKEN"
	InvalidTokenDetails    = "TOKEN DETAILS ARE INVALID"
	SimilarLast3PassWord   = "NEW PASSWORD SHOULD BE DIFFERENT THAN THE LAST 3 PASSWORDS"
	OTPUsed                = "OTP IS ALREADY USED"
	OTPInvalid             = "OTP IS INVALID"
	OTPExpired             = "OTP IS EXPIRED"
	LoginLimitReached      = "MAX SUCCESS LOGIN LIMIT REACHED"
	RetryLoginLimitReached = "MAX INCORRECT LOGIN LIMIT REACHED"
)

// resource error
const (
	DataNotFoundCode   = "RE0001"
	InvalidDetailsCode = "RE0002"

	DataNotFound   = "DATA NOT FOUND"
	InvalidDetails = "DETAILS ARE INCORRECT"
)

// other error
const (
	OtherErrorCode = "OE0001"

	UserNameIncorrect            = "USER NAME SHOULD NOT BE BLANK OR SPECIAL CHARS AND LENGTH LESS THAN 15"
	FirstNameIncorrect           = "FIRST NAME SHOULD NOT BE BLANK OR ALPHANUMERIC AND LENGTH LESS THAN 12"
	LastNameIncorrect            = "LAST NAME SHOULD NOT BE BLANK OR ALPHANUMERIC AND LENGTH LESS THAN 12"
	FullNameIncorrect            = "FULL NAME COMBINED HAS INCORRECT LENGTH"
	EmailFormatIncorrect         = "EMAIL FORMAT IS INCORRECT OR LENGTH IS MORE THAN 40"
	EmailIncorrect               = "EMAIL IS INCORRECT"
	ContactNumberFormatIncorrect = "CONTACT NUMBER SHOULD NOT BE BLANK OR ALPHANUMERIC AND LENGTH LESS THAN 15"
	ContactNumberIncorrect       = "CONTACT NUMBER IS INCORRECT"
	PasswordFormatIncorrect      = "PASSWORD SHOULD BE ALPHA NUMERIC AND LENGTH BETWEEN 6 TO 12"
	PasswordIncorrect            = "PASSWORD IS INCORRECT"
)
