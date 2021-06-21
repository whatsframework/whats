package middleware

import (
	"net/http"

	"whats/core"
	"whats/core/token"

	"github.com/gin-gonic/gin"
)

// Auth 通用权限判断
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("authorization")
		if tokenStr == "" {
			ctx.JSON(http.StatusOK, core.E{
				Success:      false,
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "token is empty",
			})
			ctx.Abort()
			return
		}
		uid, err := token.VerifyToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusOK, core.E{
				Success:      false,
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Set("uid", uid)
		ctx.Next()
	}
}
