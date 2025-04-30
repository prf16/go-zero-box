package pkg

import (
	"github.com/google/wire"
	"go-zero-box/pkg/asynqx"
	"go-zero-box/pkg/database"
	"go-zero-box/pkg/redis"
)

var Provider = wire.NewSet(
	NewPkg,
	database.Provider,
	redis.Provider,
	asynqx.Provider,
)

type Pkg struct {
	Database *database.Database
	Redis    *redis.Redis
	Asynqx   *asynqx.Asynq
}

func NewPkg(database *database.Database, redis *redis.Redis, asynqx *asynqx.Asynq) *Pkg {
	return &Pkg{Database: database, Redis: redis, Asynqx: asynqx}
}
