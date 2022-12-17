package dto

type EventAdminDto struct {
	Code  int         `json:"code"`
	Error error       `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
