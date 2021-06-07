package router

import (
	"github.com/gin-gonic/gin"
	"whats/app/controller"
	"whats/app/middleware"
)

// apiRouter apiRouter
func apiRouter(router *gin.Engine) {
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			// No Has Auth
			v1.GET("user/login", controller.HomeIndex)
			v1.GET("user/registered", controller.HomeIndex)

			// Has Auth
			v1Auth := v1.Use(middleware.Auth())
			{
				v1Auth.GET("user/token/renew", controller.HomeIndex)
				v1Auth.GET("user/info", controller.HomeIndex)
			}
		}
	}
}
