package usecase

import (
	"encoding/json"
	"pulzo/src/dao/mysql"
	"pulzo/src/domain"
	"pulzo/src/infraestructure/jwt"
	"pulzo/src/view/dto"

	"github.com/gin-gonic/gin"
)

type AdminEventUseCase struct{}

func (useCase *AdminEventUseCase) Execute(c *gin.Context, event string) dto.EventAdminDto {
	var repository domain.EventRepository = mysql.NewEventDao()

	objEvent := domain.FindEventByName(event, c.Request.Method, repository)
	objEvent.WithRepository(repository)
	if !objEvent.Exists() {
		return dto.EventAdminDto{
			Status: "error",
			Error: dto.ErrorDto{
				Code:    202,
				Message: "el evento no existe",
			},
		}
	}

	if objEvent.HasToken() && !jwt.ValidateToken(c) {
		return dto.EventAdminDto{
			Status: "error",
			Error: dto.ErrorDto{
				Code:    401,
				Message: "Su sesión ha finalizado",
			},
		}
	}

	method := domain.MethodFactory(c.Request.Method)
	strJson := method.Invoke(c, objEvent)

	var jsonParsed interface{}
	json.Unmarshal([]byte(strJson), &jsonParsed)

	return dto.EventAdminDto{
		Status: "success",
		Response: dto.SuccessDto{
			Code: 200,
			Data: jsonParsed,
		},
	}
}
