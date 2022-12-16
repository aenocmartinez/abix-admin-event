package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type DeleteSubscriberUseCase struct{}

func (useCase *DeleteSubscriberUseCase) Execute(name string) (int, error) {
	var repository domain.SubscriberRepository = mysql.NewSubscriberDao()
	subscriber := domain.FindSubscriberByName(name, repository)
	if !subscriber.Exists() {
		return 202, errors.New("el suscriptor no existe")
	}

	subscriber.WithRepository(repository).WithName(name)
	err := subscriber.Delete()
	if err != nil {
		return 202, err
	}

	return 200, nil
}
