package svc

import (
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/middleware"
	"go-zero-box/app/internal/svc/command"
	"go-zero-box/app/internal/svc/model"
	"go-zero-box/app/internal/svc/queue"
	"go-zero-box/app/internal/svc/services"
	"go-zero-box/pkg"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewServiceContext,
	command.Provider,
	model.Provider,
	queue.Provider,
	services.Provider,
)

type ServiceContext struct {
	*middleware.Middleware
	Config *config.Config
	Pkg    *pkg.Pkg

	Command  *command.Command
	Model    *model.Model
	Queue    *queue.Queue
	Services *services.Services
}

func NewServiceContext(middleware *middleware.Middleware, config *config.Config, pkg *pkg.Pkg, command *command.Command, model *model.Model, queue *queue.Queue, services *services.Services) *ServiceContext {
	return &ServiceContext{Middleware: middleware, Config: config, Pkg: pkg, Command: command, Model: model, Queue: queue, Services: services}
}
