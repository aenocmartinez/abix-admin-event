package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type ListSubscribersUseCase struct{}

func (useCase *ListSubscribersUseCase) Execute() []dto.SubscriberDto {
	var dtoSubscribers []dto.SubscriberDto
	var repository domain.SubscriberRepository = mysql.NewSubscriberDao()

	subscribers := domain.ListSubscribers(repository)
	for _, subscriber := range subscribers {
		dtoSubscribers = append(dtoSubscribers, dto.SubscriberDto{
			Name:   subscriber.Name(),
			Server: subscriber.Server(),
		})
	}

	return dtoSubscribers
}
