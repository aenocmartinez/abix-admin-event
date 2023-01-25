package usecase

import (
	"errors"
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/view/dto"
)

type FindEventUseCase struct{}

func (useCase *FindEventUseCase) Execute(id int64) (dto.EventDto, error) {
	var dtoEvent dto.EventDto
	var repository domain.EventRepository = mysql.NewEventDao()
	event := domain.FindEventById(id, repository)
	event.WithRepository(repository)
	if !event.Exists() {
		return dtoEvent, errors.New("el evento no existe")
	}

	dtoEvent.Name = event.Name()
	dtoEvent.Method = event.Method()
	dtoEvent.Subscriber = event.NameSubscriber()
	dtoEvent.Server = event.ServerSubscriber()
	dtoEvent.Id = event.Id()
	dtoEvent.WithToken = event.HasToken()

	return dtoEvent, nil
}
