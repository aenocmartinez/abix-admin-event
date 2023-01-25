package main

import (
	"net/http"
	"pulzo/src/infraestructure/jwt"
	"pulzo/src/view/controller"

	"github.com/gin-gonic/gin"
)

func validateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "header no valid"})
		}
		c.Next()
	}
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.POST("/pulzo/v1/register", controller.Register)
	r.POST("/pulzo/v1/login", controller.Login)
	r.GET("/pulzo/v1/validate-token", controller.ValidatedToken)

	routes := r.Group("/pulzo/v1", validateHeader(), jwt.AuthorizeJWT())
	{
		routes.POST("/logout", controller.Logout)
		routes.PUT("/reset-password", controller.ResetPassword)
		routes.PUT("/update-info-personal", controller.UpdateInfoPersonal)
		routes.GET("/user/:id", controller.FindUser)
		routes.GET("/inactivate/:id", controller.InactivateUser)
		routes.GET("/users", controller.AllUsers)
		routes.GET("/activate-user/:id", controller.ActivateUser)

		// Subscriber
		routes.GET("/subscribers", controller.ListSubscribers)
		routes.POST("/subscriber", controller.CreateSubscriber)
		routes.DELETE("/subscriber", controller.DeleteSubscriber)
		routes.GET("/subscriber", controller.FindSubscriber)

		// Events
		routes.GET("/events", controller.ListEvents)
		routes.POST("/event", controller.CreateEvent)
		routes.PUT("/event", controller.UpdateEvent)
		routes.DELETE("/event/:id", controller.DeleteEvent)
		routes.GET("/event", controller.FindEvent)

		// EventManager
		routes.GET("/request", controller.EventManager)
		routes.POST("/request", controller.EventManager)
		routes.PUT("/request", controller.EventManager)
		routes.DELETE("/request/:event/:id", controller.EventManager)
	}

	r.Run(":8080")
}
