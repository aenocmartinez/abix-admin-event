package domain

import (
	"github.com/gin-gonic/gin"
)

type MethodPut struct {
}

func (g *MethodPut) Invoke(c *gin.Context, event Event) (json string) {
	return "En Invoke de PUT"
}
