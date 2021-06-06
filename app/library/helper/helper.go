package helper

// RespMany 输出json的实际结构
type RespMany struct {
	RespOne
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

// RespOne 输出json的实际结构
type RespOne struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    int         `json:"errorCode,omitempty"`
	Message      string      `json:"message,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
}

// RespErr 输出json的实际结构
type RespErr struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
