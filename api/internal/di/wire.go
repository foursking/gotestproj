// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	xhttp "git.code.oa.com/qdgo/core/net/http"

	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/dao"
	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/server"
	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/service"

	"github.com/google/wire"
)

//go:generate wire

var daoProvider = wire.NewSet(dao.New)
var serviceProvider = wire.NewSet(service.New)
var serverProvider = wire.NewSet(server.New)

func NewServer() (*xhttp.Server, error) {
	panic(wire.Build(daoProvider, serviceProvider, serverProvider))
}
