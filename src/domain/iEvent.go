package domain

import "github.com/gin-gonic/gin"

type IEvent interface {
	Invoke(c *gin.Context, event Event) (json string)
}
