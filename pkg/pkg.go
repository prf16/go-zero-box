package pkg

import (
	"github.com/google/wire"
	"go-zero-box/pkg/asynqx"
	"go-zero-box/pkg/database"
	"go-zero-box/pkg/redis"
	"go-zero-box/pkg/rpc"
)

var Provider = wire.NewSet(
	NewPkg,
	database.Provider,
	redis.Provider,
	asynqx.Provider,
	rpc.Provider,
)

type Pkg struct {
	Database *database.Database
	Redis    *redis.Redis
	Asynqx   *asynqx.Asynq
	Rpc      *rpc.Rpc
}

func NewPkg(database *database.Database, redis *redis.Redis, asynqx *asynqx.Asynq, rpc *rpc.Rpc) *Pkg {
	return &Pkg{Database: database, Redis: redis, Asynqx: asynqx, Rpc: rpc}
}
