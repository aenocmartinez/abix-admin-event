package domain

import "strings"

func MethodFactory(methodType string) IEvent {
	var object IEvent
	switch strings.ToLower(methodType) {
	case "get":
		object = &MethodGet{}
	case "delete":
		object = &MethodDelete{}
	case "put":
		object = &MethodPut{}
	default:
		object = &MethodPost{}
	}
	return object
}
