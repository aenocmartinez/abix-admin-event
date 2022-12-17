package dto

type EventDto struct {
	Id         int64  `json:"id"`
	Name       string `json:"event"`
	Subscriber string `json:"subscriber"`
	Method     string `json:"method"`
	Server     string `json:"server"`
}
