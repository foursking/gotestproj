package server

import (
	"net/http"

	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/service"
	"git.code.oa.com/qdgo/core/log"
	xhttp "git.code.oa.com/qdgo/core/net/http"

	"github.com/gin-gonic/gin"
)

var (
	srv *service.Service
)

// New creates http server
func New(s *service.Service) *xhttp.Server {
	srv = s
	hs := xhttp.NewServer()
	route(hs)
	return hs
}

func route(hs *xhttp.Server) {
	hs.Ping(ping)
	hs.GET("/hello", hello)
	hs.GET("/testjson", parsejson)
	hs.GET("/testsubmit", submit)
	hs.POST("/genToken", genToken)
	hs.POST("/reflashToken", reflashToken)
}

func ping(ctx *gin.Context) {
	if err := srv.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
