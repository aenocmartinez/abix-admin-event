package formrequest

type CreateEventFormRequest struct {
	Name       string `json:"event" binding:"required"`
	Subscriber string `json:"subscriber" binding:"required"`
	Method     string `json:"method" binding:"required"`
	WithToken  bool   `json:"withToken"`
}
