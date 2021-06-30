package controller

import (
	"net/http"
	"whats/app/middleware"

	"whats/core"

	"github.com/gin-gonic/gin"
)

// UserController UserController
type UserController struct {
}

// NewUserController NewUserController
func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) login(ctx *gin.Context) {
	//service.NewUserService().Login()
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}

func (c *UserController) registry(ctx *gin.Context) {
	//service.NewUserService().Login()
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}

func (c *UserController) info(ctx *gin.Context) {
	//service.NewUserService().Info()
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}

// Router Router
func (c *UserController) Router(router *gin.RouterGroup) {
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
