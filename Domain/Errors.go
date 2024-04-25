package Domain

type AppError struct {
	ErrorMsg   string
	StatusCode int
}

func(e *AppError) NewAppError(errorMsg string, statusCode int) *AppError {
	return &AppError{ErrorMsg: errorMsg, StatusCode: statusCode}
}