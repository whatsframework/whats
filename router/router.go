package router

import (
	"whats/app/controller"

	"github.com/gin-gonic/gin"
)

// Routers Gin Routers
func Routers(router *gin.Engine) *gin.Engine {
	router.GET("/", controller.HomeIndex)

	noMethodRoute(router)

	apiRouter(router)

	swaggerRouter(router)

	return router
}
