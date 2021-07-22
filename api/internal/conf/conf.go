package conf

import (
	"github.com/foursking/ztgo/core/log"
	"github.com/micro/go-micro/v2/config"
	"time"
)

func Init() {
	go watchLog()
}

func watchLog() {
	w, _ := config.Watch("log")
	for {
		v, _ := w.Next()
		var lc log.Options
		if err := v.Scan(&lc); err != nil {
			log.Errorf("watch log config error(%v)", err)
			time.Sleep(time.Second)
			continue
		}
		log.SetLevel(lc.Level)
	}
}
