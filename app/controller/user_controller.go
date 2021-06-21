package controller

import (
	"net/http"
	"whats/app/middleware"

	"whats/core"

	"github.com/gin-gonic/gin"
)

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (c *userController) login(ctx *gin.Context) {
	//service.NewUserService().Login()
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}

func (c *userController) registry(ctx *gin.Context) {
	//service.NewUserService().Login()
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}

func (c *userController) info(ctx *gin.Context) {
	//service.NewUserService().Info()
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}

func (c *userController) Router(router *gin.RouterGroup) {
	v1 := router.Group("v1")
	{
		v1.POST("user/registry", c.registry)
		v1.POST("user/login", c.login)
		// has auth
		auth := v1.Use(middleware.Auth())
		{
			auth.GET("user/info", c.info)
		}
	}
}
