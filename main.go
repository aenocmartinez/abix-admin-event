package main

import (
	"abix360/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Subscriber
	r.GET("/abix360/admin-event/v1/subscribers", controller.ListSubscribers)
	r.POST("/abix360/admin-event/v1/subscriber", controller.CreateSubscriber)
	r.DELETE("/abix360/admin-event/v1/subscriber/:name", controller.DeleteSubscriber)
	r.GET("/abix360/admin-event/v1/subscriber/:name", controller.FindSubscriber)

	// Events
	r.GET("/abix360/admin-event/v1/events", controller.ListEvents)
	r.POST("/abix360/admin-event/v1/event", controller.CreateEvent)
	r.PUT("/abix360/admin-event/v1/event", controller.UpdateEvent)
	r.DELETE("/abix360/admin-event/v1/event/:id", controller.DeleteEvent)
	r.GET("/abix360/admin-event/v1/event/:id", controller.FindEvent)

	r.Run(":8081")
}
