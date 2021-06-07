package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// noMethodRoute 404 错误
func noMethodRoute(router *gin.Engine) {
	router.NoMethod(handleNotFound)
	router.NoRoute(handleNotFound)
}

// handleNotFound 结构化错误
func handleNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Success":      false,
		"ErrorMessage": fmt.Sprintf("Does Not Exist Router: %v %v", c.Request.Method, c.Request.URL.String()),
	})
	return
}
