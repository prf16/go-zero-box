package svc

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/middleware"
	"go-zero-box/app/internal/model"
	"go-zero-box/app/internal/pkg/database"
	"go-zero-box/app/internal/pkg/redis"
	"go-zero-box/app/internal/services"
)

var Provider = wire.NewSet(
	NewServiceContext,
	model.Provider,
	services.Provider,
	middleware.Provider,
)

type ServiceContext struct {
	Redis    *redis.Redis
	Database *database.Database
	Config   *config.Config
	Model    *model.Model
	Service  *services.Services
	*middleware.Middleware
}

func NewServiceContext(redis *redis.Redis, database *database.Database, config *config.Config, model *model.Model, service *services.Services, middleware *middleware.Middleware) *ServiceContext {
	return &ServiceContext{Redis: redis, Database: database, Config: config, Model: model, Service: service, Middleware: middleware}
}
