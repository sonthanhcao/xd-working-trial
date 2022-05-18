package errors

/*
Example error code is 5000102:
- 500 is HTTP status code (400, 401, 403, 500, ...)
- 01 is module represents for each handler
	+ 00 for common error for all handler
	+ 01 for health check handler
- 02 is actual error code, just auto increment and start at 1
*/

var (
	// Errors of module common
	// Format: ErrCommon<ERROR_NAME> = xxx00yy
	ErrCommonInternalServer    = ErrorCode("50000001")
	ErrCommonInvalidRequest    = ErrorCode("40000001")
	ErrCommonBindRequestError  = ErrorCode("40000002")
	ErrCommonUnauthorized      = ErrorCode("40100005")
	ErrShortURLNotFound        = ErrorCode("40400006")
	ErrAuthorizedNotPermission = ErrorCode("40100008")
	ErrCommonForbidden         = ErrorCode("40300000")
	ErrInternalGTNotFound      = ErrorCode("40400104")
	ErrInternalSaNotFound      = ErrorCode("40400206")
	ErrGTAssigned2AnotherRoute = ErrorCode("40000301")
)
