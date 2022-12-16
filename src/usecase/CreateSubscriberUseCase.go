package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type CreateSubscriberUseCase struct{}

func (useCase *CreateSubscriberUseCase) Execute(name, server string) (int, error) {
	var repository domain.SubscriberRepository = mysql.NewSubscriberDao()
	subscriber := domain.FindSubscriberByName(name, repository)
	if subscriber.Exists() {
		return 202, errors.New("el suscriptor ya existe")
	}

	subscriber.WithRepository(repository).WithName(name).WithServer(server)
	err := subscriber.Create()
	if err != nil {
		return 202, err
	}

	return 200, nil
}
