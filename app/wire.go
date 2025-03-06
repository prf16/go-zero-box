//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/command"
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/queue"
	"go-zero-box/app/internal/svc"
)

type App struct {
	svcCtx  *svc.ServiceContext
	queue   *queue.Queue
	command *command.Command
}

func NewApp(svcCtx *svc.ServiceContext, queue *queue.Queue, command *command.Command) *App {
	return &App{svcCtx: svcCtx, queue: queue, command: command}
}

func initApp() *App {
	wire.Build(
		config.Provider,
		svc.Provider,
		queue.Provider,
		command.Provider,
		NewApp,
	)
	return &App{}
}
