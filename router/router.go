package router

import (
	"whats/app/controller"

	"github.com/gin-gonic/gin"
)

// Routers Gin Routers
func Routers(router *gin.Engine) {
	noMethodRoute(router)
	router.GET("/", controller.HomeIndex)

	api := router.Group("api")
	{
		controller.NewUserController().Router(api)
	}
}
