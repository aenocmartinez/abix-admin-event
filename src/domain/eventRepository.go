package domain

type EventRepository interface {
	Create(event Event) error
	Delete(id int64) error
	Update(event Event) error
	AllEvents() []Event
	FindById(id int64) Event
	FindByName(name string) Event
}
