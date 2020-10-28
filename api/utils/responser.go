package utils

// Responser struct
type Responser struct {
	StatusCode int64       `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
