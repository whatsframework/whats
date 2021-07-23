# [whats framework](https://github.com/whatsframework/whats)

## Doc

[Whats](https://github.com/whatsframework/whats)
[Gin](https://github.com/gin-gonic/gin)
[Gorm](https://gorm.io/docs/)

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

## Config

### App Config
```dotenv
APP_NAME = whats
APP_DEBUG = true
HTTP_PORT = 2021
```

### DB Server Config

#### mysql Server

```dotenv
DB_CONNECTION=mysql
DB_DSN=root:123456@tcp(127.0.0.1:3306)/whats?charset=utf8mb4&parseTime=True&loc=Local
```

#### sqlite Server

```dotenv
DB_CONNECTION=sqlite
DB_DSN=whats.db
```

#### postgres Server

```dotenv
DB_CONNECTION=postgres
DB_DSN="host=localhost user=root password=123456 dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
```

#### sqlserver Server

```dotenv
DB_CONNECTION=sqlite
DB_DSN=sqlserver://root:123456@localhost:9930?database=whats
```

#### clickhouse Server

```dotenv
DB_CONNECTION=clickhouse
DB_DSN="tcp://localhost:9000?database=whats&username=root&password=123456&read_timeout=10&write_timeout=20
```

### Redis Server Config
```dotenv
REDIS_ADDR=127.0.0.1:6379
REDIS_PASSWORD=
```
### Jwt Config

```dotenv
JWT_SECRET_KEY = whats
JWT_SECRET_EXP = 7200
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
