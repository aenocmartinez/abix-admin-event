package domain

type Subscriber struct {
	repository SubscriberRepository
	name       string
	server     string
	createdAt  string
	updatedAt  string
}

func NewSubscriber(name, server string) *Subscriber {
	return &Subscriber{
		name:   name,
		server: server,
	}
}

func (s *Subscriber) WithCreatedAt(createdAt string) *Subscriber {
	s.createdAt = createdAt
	return s
}

func (s *Subscriber) WithUpdatedAt(updatedAt string) *Subscriber {
	s.updatedAt = updatedAt
	return s
}

func (s *Subscriber) Name() string {
	return s.name
}

func (s *Subscriber) Server() string {
	return s.server
}

func (s *Subscriber) WithName(name string) *Subscriber {
	s.name = name
	return s
}

func (s *Subscriber) WithServer(server string) *Subscriber {
	s.server = server
	return s
}

func (s *Subscriber) WithRepository(repository SubscriberRepository) *Subscriber {
	s.repository = repository
	return s
}

func (s *Subscriber) Exists() bool {
	return s.server != ""
}

func (s *Subscriber) Create() error {
	return s.repository.Create(*s)
}

func (s *Subscriber) Delete() error {
	return s.repository.Delete(s.name)
}

func ListSubscribers(repository SubscriberRepository) []Subscriber {
	return repository.AllSubscribers()
}

func FindSubscriberByName(name string, repository SubscriberRepository) Subscriber {
	return repository.FindByName(name)
}
