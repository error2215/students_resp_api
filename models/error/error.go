package error

import "encoding/json"

type JsonError struct {
	Code    int
	Message string
}

func New(code int, message string) *JsonError {
	return &JsonError{
		Code:    code,
		Message: message,
	}
}

func NewInString(code int, message string) string {
	ret, _ := json.Marshal(JsonError{
		Code:    code,
		Message: message,
	})
	return string(ret)
}
