package domain

import (
	"github.com/gin-gonic/gin"
)

type MethodPost struct {
}

func (p *MethodPost) Invoke(c *gin.Context, event Event) (json string) {
	return "En Invoke de POST"
}
