package controller

import (
	"abix360/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EventManager(c *gin.Context) {
	event := c.Query("event")
	if len(event) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el evento es obligatorio"})
		return
	}
	useCase := usecase.AdminEventUseCase{}
	eventAdminDto := useCase.Execute(c, event)
	if eventAdminDto.Status == "error" {
		c.JSON(eventAdminDto.Error.Code, gin.H{"error": eventAdminDto.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": eventAdminDto.Response})
}
