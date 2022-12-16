package main

import (
	"abix360/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/abix360/admin-event/v1/subscribers", controller.ListSubscribers)
	r.POST("/abix360/admin-event/v1/subscribers/create", controller.CreateSubscriber)
	r.DELETE("/abix360/admin-event/v1/subscribers/:name", controller.DeleteSubscriber)
	r.GET("/abix360/admin-event/v1/subscribers/:name", controller.FindSubscriber)

	// routes := r.Group("/abix360/v1", validateHeader(), abixjwt.AuthorizeJWT())
	// {
	// 	routes.POST("/logout", controller.Logout)
	// 	routes.PUT("/reset-password", controller.ResetPassword)
	// 	routes.PUT("/update-info-personal", controller.UpdateInfoPersonal)
	// 	routes.GET("/user/:id", controller.FindUser)
	// 	routes.GET("/unsuscribe/:id", controller.UnsuscribeUser)
	// }

	r.Run(":8081")
}
