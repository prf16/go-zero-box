package svc

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-box/app/internal/middleware"
	"go-zero-box/app/internal/model"
	"go-zero-box/app/internal/services"
)

var Provider = wire.NewSet(
	NewServiceContext,
	model.Provider,
	services.Provider,
	middleware.Provider,
)

type ServiceContext struct {
	Redis   *redis.Redis
	Model   *model.Model
	Service *services.Services
	*middleware.Middleware
}

func NewServiceContext(redis *redis.Redis, model *model.Model, service *services.Services, middleware *middleware.Middleware) *ServiceContext {
	return &ServiceContext{Redis: redis, Model: model, Service: service, Middleware: middleware}
}
