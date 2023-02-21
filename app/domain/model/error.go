package model

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAppError(msg string, code int) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func HandleError(err error, message string, code int) error {
	appErr, ok := err.(*AppError)
	if !ok {
		appErr = NewAppError(message, code)
	}
	return appErr
}
