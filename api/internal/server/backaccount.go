package server

import (
	"git.code.oa.com/qdgo/core/net/http"
	"github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
	http.JSON(ctx, "hello bossapi!", nil)
}