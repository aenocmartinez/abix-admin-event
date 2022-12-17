package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
	"fmt"
)

type CreateEventUseCase struct{}

func (useCase *CreateEventUseCase) Execute(createEvent dto.EventDto) (code int, err error) {
	var repository domain.EventRepository = mysql.NewEventDao()
	event := domain.FindEventByName(createEvent.Name, repository)
	if event.Exists() {
		return 202, errors.New("el evento ya existe")
	}

	fmt.Println("createEvent.WithToken: ", createEvent.WithToken)

	event.WithRepository(repository).WithName(createEvent.Name).WithMethod(createEvent.Method).WithToken(createEvent.WithToken)
	event.WithSubscriber(*domain.NewSubscriber(createEvent.Subscriber))

	err = event.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
