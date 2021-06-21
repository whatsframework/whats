package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whats/core"
)

// HomeIndex 首页
func HomeIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, core.H{
		Success: true,
	})
}
