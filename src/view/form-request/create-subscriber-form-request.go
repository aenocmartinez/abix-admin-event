package formrequest

type CreateSubscriberFormRequest struct {
	Name   string `json:"name" binding:"required"`
	Server string `json:"server" binding:"required"`
}
