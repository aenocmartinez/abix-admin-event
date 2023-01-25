package usecase

import (
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/infraestructure/jwt"

	"errors"
)

type RegisterUseCase struct{}

func (useCase *RegisterUseCase) Execute(name, email, password string) (int, error) {
	var repository domain.UserRepository = mysql.NewUserDao()

	user := domain.FindUserByEmail(email, repository)
	if user.Exists() {
		return 202, errors.New("el usuario ya se encuentra registrado")
	}

	user = *domain.NewUser(name, email).WithPassword(jwt.HashAndSalt([]byte(password))).WithRepository(repository)

	return 200, user.Create()
}
