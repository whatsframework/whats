package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NoMethodRoute 404 错误
func NoMethodRoute(router *gin.Engine) {
	router.NoMethod(HandleNotFound)
	router.NoRoute(HandleNotFound)
}

// HandleNotFound 结构化错误
func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Success":      false,
		"ErrorMessage": fmt.Sprintf("Does Not Exist Router: %v %v", c.Request.Method, c.Request.URL.String()),
	})
	return
}
