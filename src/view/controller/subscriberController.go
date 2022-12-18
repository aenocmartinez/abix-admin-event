package controller

import (
	"abix360/src/usecase"
	formrequest "abix360/src/view/form-request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSubscriber(c *gin.Context) {
	var req formrequest.CreateSubscriberFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.CreateSubscriberUseCase{}
	code, err := useCase.Execute(req.Name, req.Server)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "el suscriptor se ha creado con éxito"})
}

func DeleteSubscriber(c *gin.Context) {
	var name string = c.Query("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.DeleteSubscriberUseCase{}
	code, err := useCase.Execute(name)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "el suscriptor se ha eliminado con éxito"})
}

func ListSubscribers(c *gin.Context) {
	useCase := usecase.ListSubscribersUseCase{}
	subscribers := useCase.Execute()
	c.JSON(http.StatusOK, gin.H{"data": subscribers})
}

func FindSubscriber(c *gin.Context) {
	var name string = c.Query("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.FindSubscriberUseCase{}
	subscriber, err := useCase.Execute(name)
	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscriber})
}
