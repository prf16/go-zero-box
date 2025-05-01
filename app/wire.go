//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/command"
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/queue"
	"go-zero-box/app/internal/svc"
	"go-zero-box/pkg"
)

type App struct {
	config  *config.Config
	svcCtx  *svc.ServiceContext
	queue   *queue.Queue
	command *command.Command
	pkg     *pkg.Pkg
}

func NewApp(config *config.Config, svcCtx *svc.ServiceContext, queue *queue.Queue, command *command.Command, pkg *pkg.Pkg) *App {
	return &App{config: config, svcCtx: svcCtx, queue: queue, command: command, pkg: pkg}
}

func initApp(c *config.Config) *App {
	wire.Build(
		config.Provider,
		svc.Provider,
		queue.Provider,
		command.Provider,
		pkg.Provider,
		NewApp,
	)
	return &App{}
}
