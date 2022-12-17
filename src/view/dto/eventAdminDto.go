package dto

type EventAdminDto struct {
	Status   string     `json:"status"`
	Response SuccessDto `json:"response,omitempty"`
	Error    ErrorDto   `json:"error,omitempty"`
}

type ErrorDto struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessDto struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"info,omitempty"`
}
