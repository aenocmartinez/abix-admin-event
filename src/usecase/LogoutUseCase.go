package usecase

import (
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/infraestructure/jwt"

	"errors"

	"github.com/gin-gonic/gin"
)

type LogoutUseCase struct{}

func (useCase *LogoutUseCase) Execute(c *gin.Context) (int, error) {
	var token string = jwt.GetTokenRequest(c)
	isValid := jwt.VerifyToken(token)
	if !isValid {
		return 202, errors.New("token no válido")
	}

	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserByToken(token, repository)
	if !user.Exists() {
		return 202, errors.New("su sesión ha caducado")
	}

	user.WithRepository(repository).WithToken("").UpdateToken()

	// if token != user.Token() {
	// 	return 202, errors.New("su sesión ha caducado 2")
	// }

	return 200, nil
}
