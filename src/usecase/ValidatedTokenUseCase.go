package usecase

import (
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/infraestructure/jwt"

	"github.com/gin-gonic/gin"
)

type ValidatedTokenUseCase struct {
}

func (v *ValidatedTokenUseCase) Execute(c *gin.Context) bool {
	token := jwt.GetTokenRequest(c)

	if !jwt.VerifyToken(token) {
		return false
	}

	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserByToken(token, repository)

	return user.Exists()
}
