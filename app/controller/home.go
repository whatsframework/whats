package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whats/app/library/helper"
)

// HomeIndex
// @Summary 首页
// @Produce  json
// @Param authorization header string true "用户Token"
// @Router / [get]
func HomeIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.RespOne{
		Success: true,
		Message: "hello world",
	})
}
