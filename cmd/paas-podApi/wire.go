//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/zbnzl/paas-podApi/internal/biz"
	"github.com/zbnzl/paas-podApi/internal/clients"
	"github.com/zbnzl/paas-podApi/internal/conf"
	"github.com/zbnzl/paas-podApi/internal/data"
	"github.com/zbnzl/paas-podApi/internal/server"
	"github.com/zbnzl/paas-podApi/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(clients.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
