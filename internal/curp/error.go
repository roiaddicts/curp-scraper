package curp

const (
	ErrCaptchaFailure    = "CAPTCHA_FAILURE"
	ErrEncodingError     = "ENCODING_ERROR"
	ErrRequestCreation   = "REQUEST_CREATION_ERROR"
	ErrRenapoHTTPError   = "RENAPO_HTTP_ERROR"
	ErrRenapoRateLimited = "RENAPO_RATE_LIMITED"
	ErrDecodeError       = "DECODE_ERROR"
	ErrCurpNotFound      = "CURP_NOT_FOUND"
	ErrInvalidCurp       = "INVALID_CURP"
)

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Code + ": " + e.Message
}
