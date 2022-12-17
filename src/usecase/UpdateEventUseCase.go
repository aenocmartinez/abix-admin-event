package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type UpdateEventUseCase struct{}

func (useCase *UpdateEventUseCase) Execute(updateEvent dto.EventDto) (code int, err error) {
	var repository domain.EventRepository = mysql.NewEventDao()
	event := domain.FindEventById(updateEvent.Id, repository)
	if !event.Exists() {
		return 202, errors.New("el evento no existe")
	}

	event.WithRepository(repository).WithId(updateEvent.Id)
	event.WithName(updateEvent.Name)
	event.WithMethod(updateEvent.Method)
	event.WithSubscriber(*domain.NewSubscriber(updateEvent.Subscriber))

	err = event.Update()
	if err != nil {
		return 500, err
	}
	return 200, nil
}
