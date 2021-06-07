package router

import (
	_ "whats/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swaggerRouter swaggerRouter
func swaggerRouter(router *gin.Engine) {
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
