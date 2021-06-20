package controller

import (
	"github.com/gin-gonic/gin"
	"whats/core"
)

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (c *userController) login(ctx *gin.Context) {
	//service.NewUserService().Login()
	ctx.JSON(200, core.H{
		Success: true,
	})
}

func (c *userController) registry(ctx *gin.Context) {
	//service.NewUserService().Login()
	ctx.JSON(200, core.H{
		Success: true,
	})
}

func (c *userController) info(ctx *gin.Context) {
	//service.NewUserService().Info()
	ctx.JSON(200, core.H{
		Success: true,
	})
}

func (c *userController) Router(router *gin.RouterGroup) {
	v1 := router.Group("v1")
	{
		v1.POST("user/login", c.login)
		v1.POST("user/registry", c.registry)
		v1.GET("user/info", c.info)
	}
}
