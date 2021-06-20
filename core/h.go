package core

type H struct {
	Success      bool        `json:"success"`
	ErrorCode    int         `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Page         int         `json:"page,omitempty"`
	Limit        int         `json:"limit,omitempty"`
	Total        int64       `json:"total,omitempty"`
}

type E struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func NewE(errorCode int, errorMessage string) *E {
	return &E{Success: false, ErrorCode: errorCode, ErrorMessage: errorMessage}
}
