package usecase

import (
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/view/dto"
)

type ListUsersUseCase struct{}

func (useCase *ListUsersUseCase) Execute() []dto.UserDto {
	var usersDto []dto.UserDto
	var repository domain.UserRepository = mysql.NewUserDao()

	users := domain.AllUsers(repository)

	for _, user := range users {
		usersDto = append(usersDto, dto.UserDto{
			Id:        user.Id(),
			Name:      user.Name(),
			Email:     user.Email(),
			State:     user.State(),
			CreatedAt: user.CreatedAt(),
		})
	}

	return usersDto
}
