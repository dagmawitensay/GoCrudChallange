type AppError struct {
	ErrorMsg   string
	StatusCode int
}

func NewAppError(errorMsg string, statusCode int) *AppError {
	return &AppError{ErrorMsg: errorMsg, StatusCode: statusCode}
}