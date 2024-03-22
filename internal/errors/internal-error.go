package errors

type InternalError struct {
	Message  string `json:"error_message"`
	HttpCode int    `json:"http_code"`
	Details []string `json:"details"`
}

func (e *InternalError) Error() string {
	return e.Message
}

func NewInternalError(message string, httpCode int, details []string) *InternalError {
	return &InternalError{message, httpCode, details}
}