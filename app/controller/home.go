package controller

import (
	"net/http"

	"whats/core"

	"github.com/gin-gonic/gin"
)

// HomeIndex 首页
func HomeIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}
