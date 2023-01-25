package usecase

import (
	"errors"
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/view/dto"
)

type FindUserUseCase struct{}

func (useCase *FindUserUseCase) Execute(id int64) (dto.InfoPersonalDTO, error) {
	var userDto dto.InfoPersonalDTO
	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserById(id, repository)
	if !user.Exists() {
		return userDto, errors.New("el usuario no existe")
	}

	userDto.Email = user.Email()
	userDto.Name = user.Name()
	userDto.Id = user.Id()
	userDto.State = user.State()
	userDto.CreatedAt = user.CreatedAt()

	return userDto, nil
}
