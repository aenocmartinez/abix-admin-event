package domain

type SubscriberRepository interface {
	Create(subscriber Subscriber) error
	Delete(name string) error
	FindByName(name string) Subscriber
	AllSubscribers() []Subscriber
}
