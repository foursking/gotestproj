package main

import (
	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/conf"
	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/di"
	"github.com/foursking/ztgo/core"
	"github.com/foursking/ztgo/core/log"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	w := core.NewWebService()
	err := w.Init(web.BeforeStart(func() error {
		conf.Init()
		var lc log.Options
		if err := config.Get("log").Scan(&lc); err != nil {
			return err
		}
		log.Init(log.SetOptions(&lc))
		hs, err := di.NewServer()
		if err != nil {
			return err
		}
		return w.Init(web.Server(hs.Server), web.Handler(hs.Handler))
	}))
	if err != nil {
		log.Fatal(err)
	}
	if err = w.Run(); err != nil {
		log.Fatal(err)
	}
}
