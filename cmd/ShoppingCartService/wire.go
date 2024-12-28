//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ShoppingCartService/internal/biz"
	"ShoppingCartService/internal/conf"
	"ShoppingCartService/internal/data"
	"ShoppingCartService/internal/registry"
	"ShoppingCartService/internal/server"
	"ShoppingCartService/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Mysql, *conf.Etcd, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, registry.ProviderSet, newApp))
}
