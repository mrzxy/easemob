package easemob

import "encoding/json"

type Error struct {
	Code int
	Data ErrorResponse
}

func (e *Error) Error() string {
	return e.Data.ErrorDescription
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetData() ErrorResponse {
	return e.Data
}

func NewEasemobError(code int, str []byte) error {
	err := ErrorResponse{}
	_ = json.Unmarshal(str, &err)
	return &Error{Code: code, Data: err}
}

type ErrorResponse struct {
	Errors           string `json:"error"`
	Timestamp        int64  `json:"timestamp"`
	Duration         int    `json:"duration"`
	Exception        string `json:"exception"`
	ErrorDescription string `json:"error_description"`
}

// Response 获取群组成员
type Response struct {
	Action          string                 `json:"action"`
	Application     string                 `json:"application"`
	Params          map[string]interface{} `json:"params"`
	Uri             string                 `json:"uri"`
	Entities        interface{}            `json:"entities"`
	Data            interface{}            `json:"data"`
	Timestamp       int64                  `json:"timestamp"`
	Duration        int                    `json:"duration"`
	Organization    string                 `json:"organization"`
	ApplicationName string                 `json:"applicationName"`
	Count           int                    `json:"count"`
}
