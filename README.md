# whatsframework 

## Directory Structure
```
├─app
│  ├─controller
│  ├─library
│  ├─middleware
│  ├─models
│  └─service
├─bootstrap
├─core
│  ├─cache
│  ├─config
│  ├─cron
│  ├─database
│  ├─token
│  └─utils
└─router
```

## Code Demo
### Controller Demo
```go
package controller

import (
	"net/http"
	
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
		v1.POST("user/login", c.login)
		v1.POST("user/registry", c.registry)
		v1.GET("user/info", c.info)
	}
}

```

### Service Demo
```go
package service

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) Login() {

}

func (u *userService) Registry() {

}

func (u *userService) Info() {

}
```

### Router Demo
```go
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
		controller.NewUserController().Router(api) // look here
	}
}

```