package svc

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/middleware"
	"go-zero-box/app/internal/model"
	"go-zero-box/app/internal/services"
	"go-zero-box/pkg"
)

var Provider = wire.NewSet(
	NewServiceContext,
	model.Provider,
	services.Provider,
	middleware.Provider,
)

type ServiceContext struct {
	Config  *config.Config
	Model   *model.Model
	Service *services.Services
	Pkg     *pkg.Pkg
	*middleware.Middleware
}

func NewServiceContext(config *config.Config, model *model.Model, service *services.Services, pkg *pkg.Pkg, middleware *middleware.Middleware) *ServiceContext {
	return &ServiceContext{Config: config, Model: model, Service: service, Pkg: pkg, Middleware: middleware}
}
