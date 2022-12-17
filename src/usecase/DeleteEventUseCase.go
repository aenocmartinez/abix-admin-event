package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type DeleteEventUseCase struct{}

func (useCase *DeleteEventUseCase) Execute(id int64) (code int, err error) {
	var repository domain.EventRepository = mysql.NewEventDao()
	event := domain.FindEventById(id, repository)
	if !event.Exists() {
		return 202, errors.New("el evento no existe")
	}

	event.WithRepository(repository).WithId(id)
	err = event.Delete()
	if err != nil {
		return 202, err
	}

	return 200, nil
}
