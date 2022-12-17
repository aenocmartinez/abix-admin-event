package domain

import (
	"github.com/gin-gonic/gin"
)

type MethodDelete struct {
}

func (g *MethodDelete) Invoke(c *gin.Context, event Event) (json string) {
	return "En Invoke de DELETE"
}
