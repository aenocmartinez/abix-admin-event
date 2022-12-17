package formrequest

type UpdateEventFormRequest struct {
	Id         int64  `json:"id" binding:"required"`
	Name       string `json:"event" binding:"required"`
	Subscriber string `json:"subscriber" binding:"required"`
	Method     string `json:"method" binding:"required"`
}
