package controller

import (
	"net/http"
	"pulzo/src/usecase"
	"pulzo/src/view/dto"
	formrequest "pulzo/src/view/form-request"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var req formrequest.CreateEventFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.CreateEventUseCase{}
	code, err := useCase.Execute(dto.EventDto{
		Name:       req.Name,
		Subscriber: req.Subscriber,
		Method:     req.Method,
		WithToken:  req.WithToken,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "el evento se ha creado exitosamente"})
}

func DeleteEvent(c *gin.Context) {
	var strId string = c.Param("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.DeleteEventUseCase{}
	code, err := useCase.Execute(int64(id))
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "el evento se ha eliminado con éxito"})
}

func ListEvents(c *gin.Context) {
	useCase := usecase.ListEventsUseCase{}
	events := useCase.Execute()
	c.JSON(http.StatusOK, gin.H{"data": events})
}

func FindEvent(c *gin.Context) {
	var strId string = c.Query("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.FindEventUseCase{}
	event, err := useCase.Execute(int64(id))
	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}

func UpdateEvent(c *gin.Context) {
	var req formrequest.UpdateEventFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.UpdateEventUseCase{}
	code, err := useCase.Execute(dto.EventDto{
		Name:       req.Name,
		Subscriber: req.Subscriber,
		Method:     req.Method,
		Id:         req.Id,
		WithToken:  req.WithToken,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "el evento se ha actualizado exitosamente"})
}
