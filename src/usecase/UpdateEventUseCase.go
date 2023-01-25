package usecase

import (
	"errors"
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/view/dto"
)

type UpdateEventUseCase struct{}

func (useCase *UpdateEventUseCase) Execute(updateEvent dto.EventDto) (code int, err error) {
	var repository domain.EventRepository = mysql.NewEventDao()
	event := domain.FindEventById(updateEvent.Id, repository)
	event.WithRepository(repository)
	if !event.Exists() {
		return 202, errors.New("el evento no existe")
	}

	event.WithId(updateEvent.Id)
	event.WithName(updateEvent.Name)
	event.WithMethod(updateEvent.Method)
	event.WithSubscriber(*domain.NewSubscriber(updateEvent.Subscriber))
	event.WithToken(updateEvent.WithToken)

	err = event.Update()
	if err != nil {
		return 500, err
	}
	return 200, nil
}
