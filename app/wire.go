//go:build wireinject
// +build wireinject

package app

import (
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/middleware"
	"go-zero-box/app/internal/svc"
	"go-zero-box/pkg"

	"github.com/google/wire"
)

type App struct {
	config     *config.Config
	middleware *middleware.Middleware
	svcCtx     *svc.ServiceContext
	pkg        *pkg.Pkg
}

func NewApp(config *config.Config, middleware *middleware.Middleware, svcCtx *svc.ServiceContext, pkg *pkg.Pkg) *App {
	return &App{config: config, middleware: middleware, svcCtx: svcCtx, pkg: pkg}
}

func initApp(c *config.Config) *App {
	wire.Build(
		config.Provider,
		middleware.Provider,
		svc.Provider,
		pkg.Provider,
		NewApp,
	)
	return nil
}
