package domain

type Event struct {
	repository EventRepository
	id         int64
	name       string
	method     string
	subscriber Subscriber
	withToken  bool
}

func NewEvent(name, method string, subscriber Subscriber) *Event {
	return &Event{
		name:       name,
		method:     method,
		subscriber: subscriber,
		withToken:  false,
	}
}

func (e *Event) WithId(id int64) *Event {
	e.id = id
	return e
}

func (e *Event) WithName(name string) *Event {
	e.name = name
	return e
}

func (e *Event) WithMethod(method string) *Event {
	e.method = method
	return e
}

func (e *Event) WithToken(withToken bool) *Event {
	e.withToken = withToken
	return e
}

func (e *Event) WithSubscriber(subscriber Subscriber) *Event {
	e.subscriber = subscriber
	return e
}

func (e *Event) WithRepository(repository EventRepository) *Event {
	e.repository = repository
	return e
}

func (e *Event) Id() int64 {
	return e.id
}

func (e *Event) Name() string {
	return e.name
}

func (e *Event) Method() string {
	return e.method
}

func (e *Event) HasToken() bool {
	return e.withToken
}

func (e *Event) NameSubscriber() string {
	return e.subscriber.Name()
}

func (e *Event) ServerSubscriber() string {
	return e.subscriber.Server()
}

func (e *Event) Create() error {
	return e.repository.Create(*e)
}

func (e *Event) Delete() error {
	return e.repository.Delete(e.id)
}

func (e *Event) Update() error {
	return e.repository.Update(*e)
}

func (e *Event) Exists() bool {
	return e.subscriber.name != ""
}

func ListEvents(repository EventRepository) []Event {
	return repository.AllEvents()
}

func FindEventById(id int64, repository EventRepository) Event {
	return repository.FindById(id)
}

func FindEventByName(name string, repository EventRepository) Event {
	return repository.FindByName(name)
}
