package usecase

import (
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/view/dto"
)

type ListEventsUseCase struct{}

func (useCase *ListEventsUseCase) Execute() []dto.EventDto {
	var dtoEvents []dto.EventDto
	var repository domain.EventRepository = mysql.NewEventDao()

	events := domain.ListEvents(repository)
	for _, event := range events {
		dtoEvents = append(dtoEvents, dto.EventDto{
			Id:         event.Id(),
			Name:       event.Name(),
			Server:     event.ServerSubscriber(),
			Subscriber: event.NameSubscriber(),
			Method:     event.Method(),
			WithToken:  event.HasToken(),
		})
	}

	return dtoEvents
}
