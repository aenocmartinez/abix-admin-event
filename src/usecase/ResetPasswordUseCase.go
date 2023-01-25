package usecase

import (
	"errors"
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/infraestructure/jwt"
)

type ResetPasswordUseCase struct{}

func (useCase *ResetPasswordUseCase) Execute(id int64, password string) (int, error) {
	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserById(id, repository)
	if !user.Exists() {
		return 202, errors.New("el usuario no existe")
	}

	user.WithRepository(repository).WithPassword(jwt.HashAndSalt([]byte(password)))
	err := user.Update()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
