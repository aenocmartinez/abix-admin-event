package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type AdminEventUseCase struct{}

func (useCase *AdminEventUseCase) Execute(c *gin.Context, event string) dto.EventAdminDto {
	var repository domain.EventRepository = mysql.NewEventDao()

	objEvent := domain.NewEvent(event, c.Request.Method).WithRepository(repository)
	if !objEvent.Exists() {
		return dto.EventAdminDto{
			Status: "error",
			Error: dto.ErrorDto{
				Code:    202,
				Message: "el evento no existe",
			},
		}
	}

	if objEvent.HasToken() && !objEvent.HasValidToken(c) {
		return dto.EventAdminDto{
			Status: "error",
			Error: dto.ErrorDto{
				Code:    401,
				Message: "Token no válido",
			},
		}
	}

	method := domain.MethodFactory(c.Request.Method)
	strJson := method.Invoke(c, *objEvent)

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
