package core

import (
	"fmt"

	"whats/app/library/helper"

	"github.com/gin-gonic/gin"
)

// Recover 全局panic捕捉
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(200, helper.RespErr{
					Success:      false,
					ErrorCode:    500,
					ErrorMessage: fmt.Sprintf("System exception:%s", err),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
