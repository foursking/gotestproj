package server

import (
	"github.com/foursking/ztgo/core/net/http"
	"github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
	http.JSON(ctx, "hello bossapi!", nil)
}
