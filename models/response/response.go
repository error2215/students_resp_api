package response

import (
	"encoding/json"
	jsonErr "students_rest_api/models/error"
)

type response struct {
	Success bool               `json:"success"`
	Error   *jsonErr.JsonError `json:"error"`
	Data    interface{}        `json:"data"`
}

func New(errorCode int, errorMsg string, data interface{}) *response {
	if errorCode != 0 {
		return &response{
			Success: false,
			Error:   jsonErr.New(errorCode, errorMsg),
			Data:    data,
		}
	}
	return &response{
		Success: true,
		Error:   nil,
		Data:    data,
	}
}
func (r *response) ToString() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
