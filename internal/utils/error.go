package utils

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewError(msg string, code int) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
	}
}
