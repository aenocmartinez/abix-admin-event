package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type AdminEventUseCase struct{}

func (useCase *AdminEventUseCase) Execute(c *gin.Context, event string) dto.EventAdminDto {
	var repository domain.EventRepository = mysql.NewEventDao()
	objEvent := domain.FindEventByName(event, repository)
	if !objEvent.Exists() {
		return dto.EventAdminDto{
			Code:  202,
			Error: errors.New("el evento no existe"),
		}
	}
	method := domain.MethodFactory(c.Request.Method)
	strJson := method.Invoke(c, objEvent)

	var jsonParsed interface{}
	json.Unmarshal([]byte(strJson), &jsonParsed)

	return dto.EventAdminDto{
		Code:  200,
		Data:  jsonParsed,
		Error: nil,
	}
}
