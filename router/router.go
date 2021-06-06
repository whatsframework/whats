package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"whats/app/controller"

	_ "whats/docs"
)

// Routers Routers
func Routers(router *gin.Engine) *gin.Engine {
	NoMethodRoute(router)
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/", controller.HomeIndex)

	return router
}
