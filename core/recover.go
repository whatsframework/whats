package core

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Recover 全局panic捕捉
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(recover())
				log.Println(err)
				//c.JSON(200, helper.Error{
				//	Success:      false,
				//	ErrorCode:    500,
				//	ErrorMessage: fmt.Sprintf("System exception:%s", err),
				//})
				c.Abort()
			}
		}()
		c.Next()
	}
}
