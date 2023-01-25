package usecase

import (
	"errors"
	"log"
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/infraestructure/jwt"
)

type LoginUseCase struct{}

func (useCase *LoginUseCase) Execute(email, password string) (jwt.ResponseLogin, error) {
	var responseLogin jwt.ResponseLogin
	var repository domain.UserRepository = mysql.NewUserDao()

	var user domain.User = domain.FindUserByEmail(email, repository)
	if !user.Exists() {
		return responseLogin, errors.New("el usuario no existe")
	}

	if !user.IsActive() {
		return responseLogin, errors.New("el usuario está inactivo")
	}

	if !jwt.CheckPasswordHash(user.Password(), []byte(password)) {
		return responseLogin, errors.New("contraseña incorrecta")
	}

	tokendValid, err := jwt.GenerateJWT(email, "Admin")
	if err != nil {
		log.Println("LoginUseCase / GenerateJWT: ", err.Error())
	}

	user.WithToken(tokendValid).WithRepository(repository).UpdateToken()

	responseLogin.Email = email
	responseLogin.Token = tokendValid
	responseLogin.Id = user.Id()

	return responseLogin, nil
}
