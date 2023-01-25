package usecase

import (
	"errors"
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/view/dto"
)

type FindSubscriberUseCase struct{}

func (useCase *FindSubscriberUseCase) Execute(name string) (dto.SubscriberDto, error) {
	var dtoSubscriber dto.SubscriberDto
	var repository domain.SubscriberRepository = mysql.NewSubscriberDao()
	subscriber := domain.FindSubscriberByName(name, repository)
	if !subscriber.Exists() {
		return dtoSubscriber, errors.New("el suscriptor no existe")
	}

	dtoSubscriber.Name = subscriber.Name()
	dtoSubscriber.Server = subscriber.Server()

	return dtoSubscriber, nil
}
